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
	"testing"

	"k8s.io/kubernetes/pkg/api/v1"
	"k8s.io/kubernetes/pkg/client/clientset_generated/clientset/fake"
)

func checkSecret(t *testing.T, store *secretStore, ns, name string, shouldExist bool) {
	_, err := store.Get(ns, name)
	if shouldExist && err != nil {
		t.Errorf("unexpected actions: %#v", err)
	}
	if !shouldExist && (err == nil || !strings.Contains(err.Error(), "secret not registered")) {
		t.Errorf("unexpected actions: %#v", err)
	}
}

func TestSecretStore(t *testing.T) {
	fakeClient := &fake.Clientset{}
	store := newSecretStore(fakeClient)
	store.Add("ns1", "name1")
	store.Add("ns2", "name2")
	store.Add("ns1", "name1")
	store.Add("ns1", "name1")
	store.Delete("ns1", "name1")
	store.Delete("ns2", "name2")
	store.Add("ns3", "name3")

	// We expect one Get action per Add.
	actions := fakeClient.Actions()
	if len(actions) != 5 {
		t.Fatalf("unexpected actions: %#v", actions)
	}
	for _, a := range actions {
		if !a.Matches("get", "secrets") {
			t.Errorf("unexpected actions: %#v", a)
		}
	}

	checkSecret(t, store, "ns1", "name1", true)
	checkSecret(t, store, "ns2", "name2", false)
	checkSecret(t, store, "ns3", "name3", true)
	checkSecret(t, store, "ns4", "name4", false)
}

func TestSecretStoreRefresh(t *testing.T) {
	fakeClient := &fake.Clientset{}
	store := newSecretStore(fakeClient)

	for i := 0; i < 10; i++ {
		store.Add(fmt.Sprintf("ns-%d", i), fmt.Sprintf("name-%d", i))
	}
	fakeClient.ClearActions()

	store.Refresh()
	actions := fakeClient.Actions()
	if len(actions) != 10 {
		t.Fatalf("unexpected actions: %#v", actions)
	}
	for _, a := range actions {
		if !a.Matches("get", "secrets") {
			t.Errorf("unexpected actions: %#v", a)
		}
	}
}

type secretsToAttach struct {
	imagePullSecretNames    []string
	containerEnvSecretNames [][]string
}

func podWithSecrets(ns, name string, toAttach secretsToAttach) *v1.Pod {
	pod := &v1.Pod{
		ObjectMeta: v1.ObjectMeta{
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

func TestCachingSecretManager(t *testing.T) {
	fakeClient := &fake.Clientset{}
	secretStore := newSecretStore(fakeClient)
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
