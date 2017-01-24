/*
Copyright 2016 The Kubernetes Authors.

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

package secret

import (
	"fmt"
	"strings"
	"sync"
	"testing"
	"time"

	"k8s.io/kubernetes/pkg/api/v1"
	"k8s.io/kubernetes/pkg/client/clientset_generated/clientset/fake"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/util/clock"

	"github.com/stretchr/testify/assert"
)

func checkSecret(t *testing.T, store *secretStore, ns, name string, shouldExist bool) {
	_, err := store.Get(ns, name)
	if shouldExist && err != nil {
		t.Errorf("unexpected actions: %#v", err)
	}
	if !shouldExist && (err == nil || !strings.Contains(err.Error(), fmt.Sprintf("secret %q/%q not registered", ns, name))) {
		t.Errorf("unexpected actions: %#v", err)
	}
}

func TestSecretStore(t *testing.T) {
	fakeClient := &fake.Clientset{}
	store := newSecretStore(fakeClient, clock.RealClock{}, 0)
	store.Add("ns1", "name1")
	store.Add("ns2", "name2")
	store.Add("ns1", "name1")
	store.Add("ns1", "name1")
	store.Delete("ns1", "name1")
	store.Delete("ns2", "name2")
	store.Add("ns3", "name3")

	// Adds don't issue Get requests.
	actions := fakeClient.Actions()
	assert.Equal(t, 0, len(actions), "unexpected actions: %#v", actions)
	// Should issue Get request
	store.Get("ns1", "name1")
	// Shouldn't issue Get request, as secret is not registered
	store.Get("ns2", "name2")
	// Should issue Get request
	store.Get("ns3", "name3")

	actions = fakeClient.Actions()
	assert.Equal(t, 2, len(actions), "unexpected actions: %#v", actions)

	for _, a := range actions {
		assert.True(t, a.Matches("get", "secrets"), "unexpected actions: %#v", a)
	}

	checkSecret(t, store, "ns1", "name1", true)
	checkSecret(t, store, "ns2", "name2", false)
	checkSecret(t, store, "ns3", "name3", true)
	checkSecret(t, store, "ns4", "name4", false)
}

func TestSecretStoreGetAlwaysRefresh(t *testing.T) {
	fakeClient := &fake.Clientset{}
	fakeClock := clock.NewFakeClock(time.Now())
	store := newSecretStore(fakeClient, fakeClock, 0)

	for i := 0; i < 10; i++ {
		store.Add(fmt.Sprintf("ns-%d", i), fmt.Sprintf("name-%d", i))
	}
	fakeClient.ClearActions()

	wg := sync.WaitGroup{}
	wg.Add(100)
	for i := 0; i < 100; i++ {
		go func(i int) {
			store.Get(fmt.Sprintf("ns-%d", i%10), fmt.Sprintf("name-%d", i%10))
			wg.Done()
		}(i)
	}
	wg.Wait()
	actions := fakeClient.Actions()
	assert.Equal(t, 100, len(actions), "unexpected actions: %#v", actions)

	for _, a := range actions {
		assert.True(t, a.Matches("get", "secrets"), "unexpected actions: %#v", a)
	}
}

func TestSecretStoreGetNeverRefresh(t *testing.T) {
	fakeClient := &fake.Clientset{}
	fakeClock := clock.NewFakeClock(time.Now())
	store := newSecretStore(fakeClient, fakeClock, time.Minute)

	for i := 0; i < 10; i++ {
		store.Add(fmt.Sprintf("ns-%d", i), fmt.Sprintf("name-%d", i))
	}
	fakeClient.ClearActions()

	wg := sync.WaitGroup{}
	wg.Add(100)
	for i := 0; i < 100; i++ {
		go func(i int) {
			store.Get(fmt.Sprintf("ns-%d", i%10), fmt.Sprintf("name-%d", i%10))
			wg.Done()
		}(i)
	}
	wg.Wait()
	actions := fakeClient.Actions()
	// Only first Get, should forward the Get request.
	assert.Equal(t, 10, len(actions), "unexpected actions: %#v", actions)
}

type secretsToAttach struct {
	imagePullSecretNames    []string
	containerEnvSecretNames [][]string
}

func podWithSecrets(ns, name string, toAttach secretsToAttach) *v1.Pod {
	pod := &v1.Pod{
		ObjectMeta: metav1.ObjectMeta{
			Namespace: ns,
			Name:      name,
		},
		Spec: v1.PodSpec{},
	}
	for _, name := range toAttach.imagePullSecretNames {
		pod.Spec.ImagePullSecrets = append(
			pod.Spec.ImagePullSecrets, v1.LocalObjectReference{Name: name})
	}
	for i, names := range toAttach.containerEnvSecretNames {
		container := v1.Container{
			Name: fmt.Sprintf("container-%d", i),
		}
		for _, name := range names {
			envSource := &v1.EnvVarSource{
				SecretKeyRef: &v1.SecretKeySelector{
					LocalObjectReference: v1.LocalObjectReference{
						Name: name,
					},
				},
			}
			container.Env = append(container.Env, v1.EnvVar{ValueFrom: envSource})
		}
		pod.Spec.Containers = append(pod.Spec.Containers, container)
	}
	return pod
}

func TestCacheInvalidation(t *testing.T) {
	fakeClient := &fake.Clientset{}
	fakeClock := clock.NewFakeClock(time.Now())
	store := newSecretStore(fakeClient, fakeClock, time.Minute)
	manager := &cachingSecretManager{
		secretStore:    store,
		registeredPods: make(map[objectKey]*v1.Pod),
	}

	// Create a pod with some secrets.
	s1 := secretsToAttach{
		imagePullSecretNames:    []string{"s1"},
		containerEnvSecretNames: [][]string{{"s1"}, {"s2"}},
	}
	manager.RegisterPod(podWithSecrets("ns1", "name1", s1))
	// Fetch both secrets - this should triggger get operations.
	store.Get("ns1", "s1")
	store.Get("ns1", "s2")
	actions := fakeClient.Actions()
	assert.Equal(t, 2, len(actions), "unexpected actions: %#v", actions)
	fakeClient.ClearActions()

	// Update a pod with a new secret.
	s2 := secretsToAttach{
		imagePullSecretNames:    []string{"s1"},
		containerEnvSecretNames: [][]string{{"s1"}, {"s2"}, {"s3"}},
	}
	manager.RegisterPod(podWithSecrets("ns1", "name1", s2))
	// All secrets should be invalidated - this should trigger get operations.
	store.Get("ns1", "s1")
	store.Get("ns1", "s2")
	store.Get("ns1", "s3")
	actions = fakeClient.Actions()
	assert.Equal(t, 3, len(actions), "unexpected actions: %#v", actions)
	fakeClient.ClearActions()

	// Create a new pod that is refencing the first two secrets - those should
	// be invalidated.
	manager.RegisterPod(podWithSecrets("ns1", "name2", s1))
	store.Get("ns1", "s1")
	store.Get("ns1", "s2")
	store.Get("ns1", "s3")
	actions = fakeClient.Actions()
	assert.Equal(t, 2, len(actions), "unexpected actions: %#v", actions)
	fakeClient.ClearActions()
}

func TestCacheRefcounts(t *testing.T) {
	fakeClient := &fake.Clientset{}
	fakeClock := clock.NewFakeClock(time.Now())
	store := newSecretStore(fakeClient, fakeClock, time.Minute)
	manager := &cachingSecretManager{
		secretStore:    store,
		registeredPods: make(map[objectKey]*v1.Pod),
	}

	s1 := secretsToAttach{
		imagePullSecretNames:    []string{"s1"},
		containerEnvSecretNames: [][]string{{"s1"}, {"s2"}, {"s3"}},
	}
	manager.RegisterPod(podWithSecrets("ns1", "name1", s1))
	manager.RegisterPod(podWithSecrets("ns1", "name2", s1))
	s2 := secretsToAttach{
		imagePullSecretNames:    []string{"s2"},
		containerEnvSecretNames: [][]string{{"s4"}, {"s5"}},
	}
	manager.RegisterPod(podWithSecrets("ns1", "name2", s2))
	manager.RegisterPod(podWithSecrets("ns1", "name3", s2))
	manager.RegisterPod(podWithSecrets("ns1", "name4", s2))
	manager.UnregisterPod(podWithSecrets("ns1", "name3", s2))
	s3 := secretsToAttach{
		imagePullSecretNames:    []string{"s1"},
		containerEnvSecretNames: [][]string{{"s3"}, {"s5"}},
	}
	manager.RegisterPod(podWithSecrets("ns1", "name5", s3))
	manager.RegisterPod(podWithSecrets("ns1", "name6", s3))
	s4 := secretsToAttach{
		imagePullSecretNames:    []string{"s3"},
		containerEnvSecretNames: [][]string{{"s6"}},
	}
	manager.RegisterPod(podWithSecrets("ns1", "name7", s4))
	manager.UnregisterPod(podWithSecrets("ns1", "name7", s4))

	// Also check the Add + Update + Remove scenario.
	manager.RegisterPod(podWithSecrets("ns1", "other-name", s1))
	manager.RegisterPod(podWithSecrets("ns1", "other-name", s2))
	manager.UnregisterPod(podWithSecrets("ns1", "other-name", s2))

	// Now we have: 1 pod with s1, 2 pods with s2 and 2 pods with s3, 0 pods with s4.
	verify := func(ns, name string, count int) bool {
		store.lock.Lock()
		defer store.lock.Unlock()
		item, ok := store.items[objectKey{ns, name}]
		if !ok {
			return count == 0
		}
		return item.refCount == count
	}
	assert.True(t, verify("ns1", "s1", 3))
	assert.True(t, verify("ns1", "s2", 3))
	assert.True(t, verify("ns1", "s3", 3))
	assert.True(t, verify("ns1", "s4", 2))
	assert.True(t, verify("ns1", "s5", 4))
	assert.True(t, verify("ns1", "s6", 0))
	assert.True(t, verify("ns1", "s7", 0))
}

func TestCachingSecretManager(t *testing.T) {
	fakeClient := &fake.Clientset{}
	secretStore := newSecretStore(fakeClient, clock.RealClock{}, 0)
	manager := &cachingSecretManager{
		secretStore:    secretStore,
		registeredPods: make(map[objectKey]*v1.Pod),
	}

	// Create a pod with some secrets.
	s1 := secretsToAttach{
		imagePullSecretNames:    []string{"s1"},
		containerEnvSecretNames: [][]string{{"s1"}, {"s2"}},
	}
	manager.RegisterPod(podWithSecrets("ns1", "name1", s1))
	// Update the pod with a different secrets.
	s2 := secretsToAttach{
		imagePullSecretNames:    []string{"s1"},
		containerEnvSecretNames: [][]string{{"s3"}, {"s4"}},
	}
	manager.RegisterPod(podWithSecrets("ns1", "name1", s2))
	// Create another pod, but with same secrets in different namespace.
	manager.RegisterPod(podWithSecrets("ns2", "name2", s2))
	// Create and delete a pod with some other secrets.
	s3 := secretsToAttach{
		imagePullSecretNames:    []string{"s5"},
		containerEnvSecretNames: [][]string{{"s6"}},
	}
	manager.RegisterPod(podWithSecrets("ns3", "name", s3))
	manager.UnregisterPod(podWithSecrets("ns3", "name", s3))

	// We should have only: s1, s3 and s4 secrets in namespaces: ns1 and ns2.
	for _, ns := range []string{"ns1", "ns2", "ns3"} {
		for _, secret := range []string{"s1", "s2", "s3", "s4", "s5", "s6"} {
			shouldExist :=
				(secret == "s1" || secret == "s3" || secret == "s4") && (ns == "ns1" || ns == "ns2")
			checkSecret(t, secretStore, ns, secret, shouldExist)
		}
	}
}
