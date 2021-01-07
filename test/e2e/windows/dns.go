/*
Copyright 2019 The Kubernetes Authors.

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

package windows

import (
	"context"
	"regexp"
	"strings"

	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/kubernetes/test/e2e/framework"
	e2epod "k8s.io/kubernetes/test/e2e/framework/pod"
	e2eskipper "k8s.io/kubernetes/test/e2e/framework/skipper"

	"github.com/onsi/ginkgo"
)

var _ = SIGDescribe("[Feature:Windows] DNS", func() {

	ginkgo.BeforeEach(func() {
		e2eskipper.SkipUnlessNodeOSDistroIs("windows")
	})

	f := framework.NewDefaultFramework("dns")

	ginkgo.It("should support configurable pod DNS servers", func() {
		ginkgo.By("Preparing a test DNS service with injected DNS names...")
		testInjectedIP := "1.1.1.1"
		testSearchPath := "resolv.conf.local"

		ginkgo.By("Creating a pod with dnsPolicy=None and customized dnsConfig...")
		testUtilsPod := e2epod.NewAgnhostPod(f.Namespace.Name, "e2e-dns-utils", nil, nil, nil)
		testUtilsPod.Spec.DNSPolicy = v1.DNSNone
		testUtilsPod.Spec.DNSConfig = &v1.PodDNSConfig{
			Nameservers: []string{testInjectedIP},
			Searches:    []string{testSearchPath},
		}
		testUtilsPod.Spec.NodeSelector = map[string]string{
			"kubernetes.io/os": "windows",
		}
		testUtilsPod, err := f.ClientSet.CoreV1().Pods(f.Namespace.Name).Create(context.TODO(), testUtilsPod, metav1.CreateOptions{})
		framework.ExpectNoError(err)
		framework.Logf("Created pod %v", testUtilsPod)
		defer func() {
			framework.Logf("Deleting pod %s...", testUtilsPod.Name)
			if err := f.ClientSet.CoreV1().Pods(f.Namespace.Name).Delete(context.TODO(), testUtilsPod.Name, *metav1.NewDeleteOptions(0)); err != nil {
				framework.Failf("Failed to delete pod %s: %v", testUtilsPod.Name, err)
			}
		}()
		framework.ExpectNoError(e2epod.WaitForPodNameRunningInNamespace(f.ClientSet, testUtilsPod.Name, f.Namespace.Name), "failed to wait for pod %s to be running", testUtilsPod.Name)

		ginkgo.By("Verifying customized DNS option is configured on pod...")
		cmd := []string{"ipconfig", "/all"}
		stdout, _, err := f.ExecWithOptions(framework.ExecOptions{
			Command:       cmd,
			Namespace:     f.Namespace.Name,
			PodName:       testUtilsPod.Name,
			ContainerName: "agnhost-container",
			CaptureStdout: true,
			CaptureStderr: true,
		})
		framework.ExpectNoError(err)

		framework.Logf("ipconfig /all:\n%s", stdout)
		dnsRegex, err := regexp.Compile(`DNS Servers[\s*.]*:(\s*[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3})+`)
		framework.ExpectNoError(err)

		if dnsRegex.MatchString(stdout) {
			match := dnsRegex.FindString(stdout)

			if !strings.Contains(match, testInjectedIP) {
				framework.Failf("customized DNS options not found in ipconfig /all, got: %s", match)
			}
		} else {
			framework.Failf("cannot find DNS server info in ipconfig /all output: \n%s", stdout)
		}
		// TODO: Add more test cases for other DNSPolicies.
	})
})
