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

package app

import (
	"net/http"
	"testing"
	"time"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/discovery"
	fakediscovery "k8s.io/client-go/discovery/fake"
	"k8s.io/client-go/informers"
	clientset "k8s.io/client-go/kubernetes"
	fakeclientset "k8s.io/client-go/kubernetes/fake"
	restclient "k8s.io/client-go/rest"
)

// TestClientBuilder inherits ClientBuilder and can accept a given fake clientset.
type TestClientBuilder struct {
	clientset clientset.Interface
}

func (TestClientBuilder) Config(name string) (*restclient.Config, error) { return nil, nil }
func (TestClientBuilder) ConfigOrDie(name string) *restclient.Config {
	return &restclient.Config{}
}

func (TestClientBuilder) Client(name string) (clientset.Interface, error) { return nil, nil }
func (m TestClientBuilder) ClientOrDie(name string) clientset.Interface {
	return m.clientset
}

// FakeDiscoveryWithError inherits DiscoveryInterface(via FakeDiscovery) with some methods accepting testing data.
type FakeDiscoveryWithError struct {
	fakediscovery.FakeDiscovery
	PossibleResources []*metav1.APIResourceList
	Err               error
}

func (d FakeDiscoveryWithError) ServerPreferredNamespacedResources() ([]*metav1.APIResourceList, error) {
	return d.PossibleResources, d.Err
}

// FakeDiscoveryWithError inherits Clientset(via FakeClientset) with overridden Discovery method.
type FakeClientSet struct {
	fakeclientset.Clientset
	DiscoveryObj *FakeDiscoveryWithError
}

func (c *FakeClientSet) Discovery() discovery.DiscoveryInterface {
	return c.DiscoveryObj
}

func (c *FakeClientSet) GetPossibleResources() []*metav1.APIResourceList {
	return c.DiscoveryObj.PossibleResources
}

// Create a fake Clientset with its Discovery method overridden.
func NewFakeClientset(fakeDiscovery FakeDiscoveryWithError) *FakeClientSet {
	cs := &FakeClientSet{}
	cs.DiscoveryObj = &fakeDiscovery
	return cs
}

func possibleDiscoveryResource() []*metav1.APIResourceList {
	return []*metav1.APIResourceList{
		{
			GroupVersion: "create/v1",
			APIResources: []metav1.APIResource{
				{
					Name:       "jobs",
					Verbs:      []string{"create", "list", "watch", "delete"},
					ShortNames: []string{"jz"},
					Categories: []string{"all"},
				},
			},
		},
	}
}

type controllerInitFunc func(ControllerContext) (http.Handler, bool, error)

func TestController_DiscoveryError(t *testing.T) {
	controllerInitFuncMap := map[string]controllerInitFunc{
		"ResourceQuotaController":    startResourceQuotaController,
		"GarbageCollectorController": startGarbageCollectorController,
	}

	tcs := map[string]struct {
		discoveryError    error
		expectedErr       bool
		possibleResources []*metav1.APIResourceList
	}{
		"No Discovery Error": {
			discoveryError:    nil,
			possibleResources: possibleDiscoveryResource(),
			expectedErr:       false,
		},
		"Discovery Calls Partially Failed": {
			discoveryError:    new(discovery.ErrGroupDiscoveryFailed),
			possibleResources: possibleDiscoveryResource(),
			expectedErr:       false,
		},
	}
	for name, test := range tcs {
		testDiscovery := FakeDiscoveryWithError{Err: test.discoveryError, PossibleResources: test.possibleResources}
		testClientset := NewFakeClientset(testDiscovery)
		testClientBuilder := TestClientBuilder{clientset: testClientset}
		ctx := ControllerContext{
			ClientBuilder:    testClientBuilder,
			InformerFactory:  informers.NewSharedInformerFactoryWithOptions(testClientset, time.Duration(1)),
			InformersStarted: make(chan struct{}),
		}
		for funcName, controllerInit := range controllerInitFuncMap {
			_, _, err := controllerInit(ctx)
			if test.expectedErr != (err != nil) {
				t.Errorf("%v test failed for use case: %v", funcName, name)
			}
		}
		_, _, err := startModifiedNamespaceController(
			ctx, testClientset, testClientBuilder.ConfigOrDie("namespace-controller"))
		if test.expectedErr != (err != nil) {
			t.Errorf("Namespace Controller test failed for use case: %v", name)
		}
	}
}
