/*
Copyright 2020 The Kubernetes Authors.

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

package endpointslicemirroring

import (
	"context"
	"fmt"
	"testing"
	"time"

	v1 "k8s.io/api/core/v1"
	discovery "k8s.io/api/discovery/v1beta1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/informers"
	"k8s.io/client-go/kubernetes/fake"
	"k8s.io/client-go/tools/cache"
	"k8s.io/client-go/tools/leaderelection/resourcelock"
	"k8s.io/kubernetes/pkg/controller"
)

// Most of the tests related to EndpointSlice allocation can be found in reconciler_test.go
// Tests here primarily focus on unique controller functionality before the reconciler begins

var alwaysReady = func() bool { return true }

type endpointSliceMirroringController struct {
	*Controller
	endpointsStore     cache.Store
	endpointSliceStore cache.Store
	serviceStore       cache.Store
}

func newController(batchPeriod time.Duration) (*fake.Clientset, *endpointSliceMirroringController) {
	client := newClientset()
	informerFactory := informers.NewSharedInformerFactory(client, controller.NoResyncPeriodFunc())

	esController := NewController(
		informerFactory.Core().V1().Endpoints(),
		informerFactory.Discovery().V1beta1().EndpointSlices(),
		informerFactory.Core().V1().Services(),
		int32(1000),
		client,
		batchPeriod)

	esController.endpointsSynced = alwaysReady
	esController.endpointSlicesSynced = alwaysReady
	esController.servicesSynced = alwaysReady

	return client, &endpointSliceMirroringController{
		esController,
		informerFactory.Core().V1().Endpoints().Informer().GetStore(),
		informerFactory.Discovery().V1beta1().EndpointSlices().Informer().GetStore(),
		informerFactory.Core().V1().Services().Informer().GetStore(),
	}
}

func TestSyncEndpoints(t *testing.T) {
	endpointsName := "testing-sync-endpoints"
	namespace := metav1.NamespaceDefault

	testCases := []struct {
		testName           string
		endpoints          *v1.Endpoints
		endpointSlices     []*discovery.EndpointSlice
		expectedNumActions int
		expectedNumSlices  int
	}{{
		testName: "Endpoints with no addresses",
		endpoints: &v1.Endpoints{
			Subsets: []v1.EndpointSubset{{
				Ports: []v1.EndpointPort{{Port: 80}},
			}},
		},
		endpointSlices:     []*discovery.EndpointSlice{},
		expectedNumActions: 0,
		expectedNumSlices:  0,
	}, {
		testName: "Endpoints with skip label true",
		endpoints: &v1.Endpoints{
			ObjectMeta: metav1.ObjectMeta{
				Labels: map[string]string{discovery.LabelSkipMirror: "true"},
			},
			Subsets: []v1.EndpointSubset{{
				Ports:     []v1.EndpointPort{{Port: 80}},
				Addresses: []v1.EndpointAddress{{IP: "10.0.0.1"}},
			}},
		},
		endpointSlices:     []*discovery.EndpointSlice{},
		expectedNumActions: 0,
		expectedNumSlices:  0,
	}, {
		testName: "Endpoints with skip label false",
		endpoints: &v1.Endpoints{
			ObjectMeta: metav1.ObjectMeta{
				Labels: map[string]string{discovery.LabelSkipMirror: "false"},
			},
			Subsets: []v1.EndpointSubset{{
				Ports:     []v1.EndpointPort{{Port: 80}},
				Addresses: []v1.EndpointAddress{{IP: "10.0.0.1"}},
			}},
		},
		endpointSlices:     []*discovery.EndpointSlice{},
		expectedNumActions: 1,
		expectedNumSlices:  1,
	}, {
		testName: "Existing EndpointSlices that need to be cleaned up",
		endpoints: &v1.Endpoints{
			Subsets: []v1.EndpointSubset{{
				Ports: []v1.EndpointPort{{Port: 80}},
			}},
		},
		endpointSlices: []*discovery.EndpointSlice{{
			ObjectMeta: metav1.ObjectMeta{
				Name: endpointsName + "-1",
				Labels: map[string]string{
					discovery.LabelServiceName: endpointsName,
					discovery.LabelManagedBy:   controllerName,
				},
			},
		}},
		expectedNumActions: 1,
		expectedNumSlices:  0,
	}, {
		testName: "Existing EndpointSlices managed by a different controller, no addresses to sync",
		endpoints: &v1.Endpoints{
			Subsets: []v1.EndpointSubset{{
				Ports: []v1.EndpointPort{{Port: 80}},
			}},
		},
		endpointSlices: []*discovery.EndpointSlice{{
			ObjectMeta: metav1.ObjectMeta{
				Name: endpointsName + "-1",
				Labels: map[string]string{
					discovery.LabelManagedBy: "something-else",
				},
			},
		}},
		expectedNumActions: 0,
		// This only queries for EndpointSlices managed by this controller.
		expectedNumSlices: 0,
	}, {
		testName: "Endpoints with 1000 addresses",
		endpoints: &v1.Endpoints{
			Subsets: []v1.EndpointSubset{{
				Ports:     []v1.EndpointPort{{Port: 80}},
				Addresses: generateAddresses(1000),
			}},
		},
		endpointSlices:     []*discovery.EndpointSlice{},
		expectedNumActions: 1,
		expectedNumSlices:  1,
	}, {
		testName: "Endpoints with 1001 addresses - 1 should not be mirrored",
		endpoints: &v1.Endpoints{
			Subsets: []v1.EndpointSubset{{
				Ports:     []v1.EndpointPort{{Port: 80}},
				Addresses: generateAddresses(1001),
			}},
		},
		endpointSlices:     []*discovery.EndpointSlice{},
		expectedNumActions: 1,
		expectedNumSlices:  1,
	}}

	for _, tc := range testCases {
		t.Run(tc.testName, func(t *testing.T) {
			client, esController := newController(time.Duration(0))
			tc.endpoints.Name = endpointsName
			tc.endpoints.Namespace = namespace
			esController.endpointsStore.Add(tc.endpoints)
			esController.serviceStore.Add(&v1.Service{ObjectMeta: metav1.ObjectMeta{
				Name:      endpointsName,
				Namespace: namespace,
			}})

			for _, epSlice := range tc.endpointSlices {
				epSlice.Namespace = namespace
				esController.endpointSliceStore.Add(epSlice)
				_, err := client.DiscoveryV1beta1().EndpointSlices(namespace).Create(context.TODO(), epSlice, metav1.CreateOptions{})
				if err != nil {
					t.Fatalf("Expected no error creating EndpointSlice, got %v", err)
				}
			}

			err := esController.syncEndpoints(fmt.Sprintf("%s/%s", namespace, endpointsName))
			if err != nil {
				t.Errorf("Unexpected error from syncEndpoints: %v", err)
			}

			numInitialActions := len(tc.endpointSlices)
			numExtraActions := len(client.Actions()) - numInitialActions
			if numExtraActions != tc.expectedNumActions {
				t.Fatalf("Expected %d additional client actions, got %d: %#v", tc.expectedNumActions, numExtraActions, client.Actions()[numInitialActions:])
			}

			endpointSlices := fetchEndpointSlices(t, client, namespace)
			expectEndpointSlices(t, tc.expectedNumSlices, int(defaultMaxEndpointsPerSubset), *tc.endpoints, endpointSlices)
		})
	}
}

func TestShouldMirror(t *testing.T) {
	svcWithSelector := &v1.Service{
		ObjectMeta: metav1.ObjectMeta{
			Name:      "with-selector",
			Namespace: "example1",
		},
		Spec: v1.ServiceSpec{
			Selector: map[string]string{"with": "selector"},
		},
	}
	svcWithoutSelector := &v1.Service{
		ObjectMeta: metav1.ObjectMeta{
			Name:      "without-selector",
			Namespace: "example1",
		},
		Spec: v1.ServiceSpec{},
	}

	testCases := []struct {
		testName     string
		endpoints    *v1.Endpoints
		service      *v1.Service
		shouldMirror bool
	}{{
		testName: "Service without selector with matching endpoints",
		service:  svcWithoutSelector,
		endpoints: &v1.Endpoints{
			ObjectMeta: metav1.ObjectMeta{
				Name:      svcWithoutSelector.Name,
				Namespace: svcWithoutSelector.Namespace,
			},
		},
		shouldMirror: true,
	}, {
		testName: "Service without selector, matching Endpoints with skip-mirror=true",
		service:  svcWithoutSelector,
		endpoints: &v1.Endpoints{
			ObjectMeta: metav1.ObjectMeta{
				Name:      svcWithSelector.Name,
				Namespace: svcWithSelector.Namespace,
				Labels: map[string]string{
					discovery.LabelSkipMirror: "true",
				},
			},
		},
		shouldMirror: false,
	}, {
		testName: "Service without selector, matching Endpoints with skip-mirror=invalid",
		service:  svcWithoutSelector,
		endpoints: &v1.Endpoints{
			ObjectMeta: metav1.ObjectMeta{
				Name:      svcWithoutSelector.Name,
				Namespace: svcWithoutSelector.Namespace,
				Labels: map[string]string{
					discovery.LabelSkipMirror: "invalid",
				},
			},
		},
		shouldMirror: true,
	}, {
		testName: "Service without selector, matching Endpoints with leader election annotation",
		service:  svcWithoutSelector,
		endpoints: &v1.Endpoints{
			ObjectMeta: metav1.ObjectMeta{
				Name:      svcWithSelector.Name,
				Namespace: svcWithSelector.Namespace,
				Annotations: map[string]string{
					resourcelock.LeaderElectionRecordAnnotationKey: "",
				},
			},
		},
		shouldMirror: false,
	}, {
		testName: "Service without selector, matching Endpoints without skip label in different namespace",
		service:  svcWithSelector,
		endpoints: &v1.Endpoints{
			ObjectMeta: metav1.ObjectMeta{
				Name:      svcWithSelector.Name,
				Namespace: svcWithSelector.Namespace + "different",
			},
		},
		shouldMirror: false,
	}, {
		testName:     "Service without selector or matching endpoints",
		service:      svcWithoutSelector,
		endpoints:    nil,
		shouldMirror: false,
	}, {
		testName: "Endpoints without matching Service",
		service:  nil,
		endpoints: &v1.Endpoints{
			ObjectMeta: metav1.ObjectMeta{
				Name:      svcWithoutSelector.Name,
				Namespace: svcWithoutSelector.Namespace,
			},
		},
		shouldMirror: false,
	}}

	for _, tc := range testCases {
		t.Run(tc.testName, func(t *testing.T) {
			_, c := newController(time.Duration(0))

			if tc.endpoints != nil {
				err := c.endpointsStore.Add(tc.endpoints)
				if err != nil {
					t.Fatalf("Error adding Endpoints to store: %v", err)
				}
			}

			if tc.service != nil {
				err := c.serviceStore.Add(tc.service)
				if err != nil {
					t.Fatalf("Error adding Service to store: %v", err)
				}
			}

			shouldMirror := c.shouldMirror(tc.endpoints)

			if shouldMirror != tc.shouldMirror {
				t.Errorf("Expected %t to be returned, got %t", tc.shouldMirror, shouldMirror)
			}
		})
	}
}

func TestEndpointSlicesMirroredForService(t *testing.T) {
	testCases := []struct {
		testName       string
		namespace      string
		name           string
		endpointSlice  *discovery.EndpointSlice
		expectedInList bool
	}{{
		testName:  "Service with matching EndpointSlice",
		namespace: "ns1",
		name:      "svc1",
		endpointSlice: &discovery.EndpointSlice{
			ObjectMeta: metav1.ObjectMeta{
				Name:      "example-1",
				Namespace: "ns1",
				Labels: map[string]string{
					discovery.LabelServiceName: "svc1",
					discovery.LabelManagedBy:   controllerName,
				},
			},
		},
		expectedInList: true,
	}, {
		testName:  "Service with EndpointSlice that has different namespace",
		namespace: "ns1",
		name:      "svc1",
		endpointSlice: &discovery.EndpointSlice{
			ObjectMeta: metav1.ObjectMeta{
				Name:      "example-1",
				Namespace: "ns2",
				Labels: map[string]string{
					discovery.LabelServiceName: "svc1",
					discovery.LabelManagedBy:   controllerName,
				},
			},
		},
		expectedInList: false,
	}, {
		testName:  "Service with EndpointSlice that has different service name",
		namespace: "ns1",
		name:      "svc1",
		endpointSlice: &discovery.EndpointSlice{
			ObjectMeta: metav1.ObjectMeta{
				Name:      "example-1",
				Namespace: "ns1",
				Labels: map[string]string{
					discovery.LabelServiceName: "svc2",
					discovery.LabelManagedBy:   controllerName,
				},
			},
		},
		expectedInList: false,
	}, {
		testName:  "Service with EndpointSlice that has different controller name",
		namespace: "ns1",
		name:      "svc1",
		endpointSlice: &discovery.EndpointSlice{
			ObjectMeta: metav1.ObjectMeta{
				Name:      "example-1",
				Namespace: "ns1",
				Labels: map[string]string{
					discovery.LabelServiceName: "svc1",
					discovery.LabelManagedBy:   controllerName + "foo",
				},
			},
		},
		expectedInList: false,
	}, {
		testName:  "Service with EndpointSlice that has missing controller name",
		namespace: "ns1",
		name:      "svc1",
		endpointSlice: &discovery.EndpointSlice{
			ObjectMeta: metav1.ObjectMeta{
				Name:      "example-1",
				Namespace: "ns1",
				Labels: map[string]string{
					discovery.LabelServiceName: "svc1",
				},
			},
		},
		expectedInList: false,
	}, {
		testName:  "Service with EndpointSlice that has missing service name",
		namespace: "ns1",
		name:      "svc1",
		endpointSlice: &discovery.EndpointSlice{
			ObjectMeta: metav1.ObjectMeta{
				Name:      "example-1",
				Namespace: "ns1",
				Labels: map[string]string{
					discovery.LabelManagedBy: controllerName,
				},
			},
		},
		expectedInList: false,
	}}

	for _, tc := range testCases {
		t.Run(tc.testName, func(t *testing.T) {
			_, c := newController(time.Duration(0))

			err := c.endpointSliceStore.Add(tc.endpointSlice)
			if err != nil {
				t.Fatalf("Error adding EndpointSlice to store: %v", err)
			}

			endpointSlices, err := endpointSlicesMirroredForService(c.endpointSliceLister, tc.namespace, tc.name)
			if err != nil {
				t.Fatalf("Expected no error, got %v", err)
			}

			if tc.expectedInList {
				if len(endpointSlices) != 1 {
					t.Fatalf("Expected 1 EndpointSlice to be in list, got %d", len(endpointSlices))
				}

				if endpointSlices[0].Name != tc.endpointSlice.Name {
					t.Fatalf("Expected %s EndpointSlice to be in list, got %s", tc.endpointSlice.Name, endpointSlices[0].Name)
				}
			} else {
				if len(endpointSlices) != 0 {
					t.Fatalf("Expected no EndpointSlices to be in list, got %d", len(endpointSlices))
				}
			}
		})
	}
}

func generateAddresses(num int) []v1.EndpointAddress {
	addresses := make([]v1.EndpointAddress, num)
	for i := 0; i < num; i++ {
		part1 := i / 255
		part2 := i % 255
		ip := fmt.Sprintf("10.0.%d.%d", part1, part2)
		addresses[i] = v1.EndpointAddress{IP: ip}
	}
	return addresses
}
