/*
Copyright 2018 The Kubernetes Authors.

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

package testing

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/kubernetes/pkg/kubelet/runtimeclass"
	nodev1alpha1 "k8s.io/node-api/pkg/apis/node/v1alpha1"
	nodeapiclient "k8s.io/node-api/pkg/client/clientset/versioned"
	nodeapifake "k8s.io/node-api/pkg/client/clientset/versioned/fake"
)

const (
	// SandboxRuntimeClass is a valid RuntimeClass pre-populated in the populated dynamic client.
	SandboxRuntimeClass = "sandbox"
	// SandboxRuntimeHandler is the handler associated with the SandboxRuntimeClass.
	SandboxRuntimeHandler = "kata-containers"

	// EmptyRuntimeClass is a valid RuntimeClass without a handler pre-populated in the populated dynamic client.
	EmptyRuntimeClass = "native"
)

// NewPopulatedClient creates a node-api client for use with the runtimeclass.Manager,
// and populates it with a few test RuntimeClass objects.
func NewPopulatedClient() nodeapiclient.Interface {
	return nodeapifake.NewSimpleClientset(
		NewRuntimeClass(EmptyRuntimeClass, ""),
		NewRuntimeClass(SandboxRuntimeClass, SandboxRuntimeHandler),
	)
}

// StartManagerSync starts the manager, and waits for the informer cache to sync.
// Returns a function to stop the manager, which should be called with a defer:
//     defer StartManagerSync(t, m)()
func StartManagerSync(m *runtimeclass.Manager) func() {
	stopCh := make(chan struct{})
	m.Start(stopCh)
	m.WaitForCacheSync(stopCh)
	return func() {
		close(stopCh)
	}
}

// NewRuntimeClass is a helper to generate a RuntimeClass resource with
// the given name & handler.
func NewRuntimeClass(name, handler string) *nodev1alpha1.RuntimeClass {
	return &nodev1alpha1.RuntimeClass{
		ObjectMeta: metav1.ObjectMeta{
			Name: name,
		},
		Spec: nodev1alpha1.RuntimeClassSpec{
			RuntimeHandler: &handler,
		},
	}
}

// NewUnstructuredRuntimeClass is a helper to generate an unstructured RuntimeClass resource with
// the given name & handler.
func NewUnstructuredRuntimeClass(name, handler string) *unstructured.Unstructured {
	return &unstructured.Unstructured{
		Object: map[string]interface{}{
			"apiVersion": "node.k8s.io/v1alpha1",
			"kind":       "RuntimeClass",
			"metadata": map[string]interface{}{
				"name": name,
			},
			"spec": map[string]interface{}{
				"runtimeHandler": handler,
			},
		},
	}
}
