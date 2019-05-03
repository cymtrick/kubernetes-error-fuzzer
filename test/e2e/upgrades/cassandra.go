/*
Copyright 2017 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package upgrades

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"path/filepath"
	"sync"
	"time"

	"github.com/onsi/ginkgo"
	"github.com/onsi/gomega"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/version"
	"k8s.io/apimachinery/pkg/util/wait"
	"k8s.io/kubernetes/test/e2e/framework"
	e2elog "k8s.io/kubernetes/test/e2e/framework/log"
	"k8s.io/kubernetes/test/e2e/framework/testfiles"
)

const cassandraManifestPath = "test/e2e/testing-manifests/statefulset/cassandra"

// CassandraUpgradeTest ups and verifies that a Cassandra StatefulSet behaves
// well across upgrades.
type CassandraUpgradeTest struct {
	ip               string
	successfulWrites int
	ssTester         *framework.StatefulSetTester
}

// Name returns the tracking name of the test.
func (CassandraUpgradeTest) Name() string { return "cassandra-upgrade" }

// Skip returns true when this test can be skipped.
func (CassandraUpgradeTest) Skip(upgCtx UpgradeContext) bool {
	minVersion := version.MustParseSemantic("1.6.0")
	for _, vCtx := range upgCtx.Versions {
		if vCtx.Version.LessThan(minVersion) {
			return true
		}
	}
	return false
}

func cassandraKubectlCreate(ns, file string) {
	input := string(testfiles.ReadOrDie(filepath.Join(cassandraManifestPath, file), ginkgo.Fail))
	framework.RunKubectlOrDieInput(input, "create", "-f", "-", fmt.Sprintf("--namespace=%s", ns))
}

// Setup creates a Cassandra StatefulSet and a PDB. It also brings up a tester
// ReplicaSet and associated service and PDB to guarantee availability during
// the upgrade.
// It waits for the system to stabilize before adding two users to verify
// connectivity.
func (t *CassandraUpgradeTest) Setup(f *framework.Framework) {
	ns := f.Namespace.Name
	statefulsetPoll := 30 * time.Second
	statefulsetTimeout := 10 * time.Minute
	t.ssTester = framework.NewStatefulSetTester(f.ClientSet)

	ginkgo.By("Creating a PDB")
	cassandraKubectlCreate(ns, "pdb.yaml")

	ginkgo.By("Creating a Cassandra StatefulSet")
	t.ssTester.CreateStatefulSet(cassandraManifestPath, ns)

	ginkgo.By("Creating a cassandra-test-server deployment")
	cassandraKubectlCreate(ns, "tester.yaml")

	ginkgo.By("Getting the ingress IPs from the services")
	err := wait.PollImmediate(statefulsetPoll, statefulsetTimeout, func() (bool, error) {
		if t.ip = t.getServiceIP(f, ns, "test-server"); t.ip == "" {
			return false, nil
		}
		if _, err := t.listUsers(); err != nil {
			e2elog.Logf("Service endpoint is up but isn't responding")
			return false, nil
		}
		return true, nil
	})
	framework.ExpectNoError(err)
	e2elog.Logf("Service endpoint is up")

	ginkgo.By("Adding 2 dummy users")
	gomega.Expect(t.addUser("Alice")).NotTo(gomega.HaveOccurred())
	gomega.Expect(t.addUser("Bob")).NotTo(gomega.HaveOccurred())
	t.successfulWrites = 2

	ginkgo.By("Verifying that the users exist")
	users, err := t.listUsers()
	framework.ExpectNoError(err)
	gomega.Expect(len(users)).To(gomega.Equal(2))
}

// listUsers gets a list of users from the db via the tester service.
func (t *CassandraUpgradeTest) listUsers() ([]string, error) {
	r, err := http.Get(fmt.Sprintf("http://%s:8080/list", t.ip))
	if err != nil {
		return nil, err
	}
	defer r.Body.Close()
	if r.StatusCode != http.StatusOK {
		b, err := ioutil.ReadAll(r.Body)
		if err != nil {
			return nil, err
		}
		return nil, fmt.Errorf(string(b))
	}
	var names []string
	if err := json.NewDecoder(r.Body).Decode(&names); err != nil {
		return nil, err
	}
	return names, nil
}

// addUser adds a user to the db via the tester services.
func (t *CassandraUpgradeTest) addUser(name string) error {
	val := map[string][]string{"name": {name}}
	r, err := http.PostForm(fmt.Sprintf("http://%s:8080/add", t.ip), val)
	if err != nil {
		return err
	}
	defer r.Body.Close()
	if r.StatusCode != http.StatusOK {
		b, err := ioutil.ReadAll(r.Body)
		if err != nil {
			return err
		}
		return fmt.Errorf(string(b))
	}
	return nil
}

// getServiceIP is a helper method to extract the Ingress IP from the service.
func (t *CassandraUpgradeTest) getServiceIP(f *framework.Framework, ns, svcName string) string {
	svc, err := f.ClientSet.CoreV1().Services(ns).Get(svcName, metav1.GetOptions{})
	framework.ExpectNoError(err)
	ingress := svc.Status.LoadBalancer.Ingress
	if len(ingress) == 0 {
		return ""
	}
	return ingress[0].IP
}

// Test is called during the upgrade.
// It launches two goroutines, one continuously writes to the db and one reads
// from the db. Each attempt is tallied and at the end we verify if the success
// ratio is over a certain threshold (0.75). We also verify that we get
// at least the same number of rows back as we successfully wrote.
func (t *CassandraUpgradeTest) Test(f *framework.Framework, done <-chan struct{}, upgrade UpgradeType) {
	ginkgo.By("Continuously polling the database during upgrade.")
	var (
		success, failures, writeAttempts, lastUserCount int
		mu                                              sync.Mutex
		errors                                          = map[string]int{}
	)
	// Write loop.
	go wait.Until(func() {
		writeAttempts++
		if err := t.addUser(fmt.Sprintf("user-%d", writeAttempts)); err != nil {
			e2elog.Logf("Unable to add user: %v", err)
			mu.Lock()
			errors[err.Error()]++
			mu.Unlock()
			return
		}
		t.successfulWrites++
	}, 10*time.Millisecond, done)
	// Read loop.
	wait.Until(func() {
		users, err := t.listUsers()
		if err != nil {
			e2elog.Logf("Could not retrieve users: %v", err)
			failures++
			mu.Lock()
			errors[err.Error()]++
			mu.Unlock()
			return
		}
		success++
		lastUserCount = len(users)
	}, 10*time.Millisecond, done)
	e2elog.Logf("got %d users; want >=%d", lastUserCount, t.successfulWrites)

	gomega.Expect(lastUserCount >= t.successfulWrites).To(gomega.BeTrue())
	ratio := float64(success) / float64(success+failures)
	e2elog.Logf("Successful gets %d/%d=%v", success, success+failures, ratio)
	ratio = float64(t.successfulWrites) / float64(writeAttempts)
	e2elog.Logf("Successful writes %d/%d=%v", t.successfulWrites, writeAttempts, ratio)
	e2elog.Logf("Errors: %v", errors)
	// TODO(maisem): tweak this value once we have a few test runs.
	gomega.Expect(ratio > 0.75).To(gomega.BeTrue())
}

// Teardown does one final check of the data's availability.
func (t *CassandraUpgradeTest) Teardown(f *framework.Framework) {
	users, err := t.listUsers()
	framework.ExpectNoError(err)
	gomega.Expect(len(users) >= t.successfulWrites).To(gomega.BeTrue())
}
