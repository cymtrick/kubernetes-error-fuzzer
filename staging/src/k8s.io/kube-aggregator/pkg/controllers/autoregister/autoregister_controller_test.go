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

package autoregister

import (
	"fmt"
	"testing"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	clienttesting "k8s.io/client-go/testing"
	"k8s.io/client-go/tools/cache"
	"k8s.io/client-go/util/workqueue"
	"k8s.io/kube-aggregator/pkg/apis/apiregistration"
	"k8s.io/kube-aggregator/pkg/client/clientset_generated/internalclientset/fake"
	listers "k8s.io/kube-aggregator/pkg/client/listers/apiregistration/internalversion"
)

func newAutoRegisterManagedAPIService(name string) *apiregistration.APIService {
	return &apiregistration.APIService{
		ObjectMeta: metav1.ObjectMeta{Name: name, Labels: map[string]string{AutoRegisterManagedLabel: string("true")}},
	}
}

func newAutoRegisterManagedModifiedAPIService(name string) *apiregistration.APIService {
	return &apiregistration.APIService{
		ObjectMeta: metav1.ObjectMeta{Name: name, Labels: map[string]string{AutoRegisterManagedLabel: string("true")}},
		Spec: apiregistration.APIServiceSpec{
			Group: "something",
		},
	}
}

func newAPIService(name string) *apiregistration.APIService {
	return &apiregistration.APIService{
		ObjectMeta: metav1.ObjectMeta{Name: name},
	}
}

func checkForNothing(name string, client *fake.Clientset) error {
	if len(client.Actions()) > 0 {
		return fmt.Errorf("unexpected action: %v", client.Actions())
	}

	return nil
}

func checkForCreate(name string, client *fake.Clientset) error {
	if len(client.Actions()) == 0 {
		return nil
	}
	if len(client.Actions()) > 1 {
		return fmt.Errorf("unexpected action: %v", client.Actions())
	}

	action := client.Actions()[0]

	createAction, ok := action.(clienttesting.CreateAction)
	if !ok {
		return fmt.Errorf("unexpected action: %v", client.Actions())
	}
	apiService := createAction.GetObject().(*apiregistration.APIService)
	if apiService.Name != name || apiService.Labels[AutoRegisterManagedLabel] != "true" {
		return fmt.Errorf("bad name or label %v", createAction)
	}

	return nil
}

func checkForUpdate(name string, client *fake.Clientset) error {
	if len(client.Actions()) == 0 {
		return nil
	}
	if len(client.Actions()) > 1 {
		return fmt.Errorf("unexpected action: %v", client.Actions())
	}

	action := client.Actions()[0]
	updateAction, ok := action.(clienttesting.UpdateAction)
	if !ok {
		return fmt.Errorf("unexpected action: %v", client.Actions())
	}
	apiService := updateAction.GetObject().(*apiregistration.APIService)
	if apiService.Name != name || apiService.Labels[AutoRegisterManagedLabel] != "true" || apiService.Spec.Group != "" {
		return fmt.Errorf("bad name, label, or group %v", updateAction)
	}

	return nil
}

func checkForDelete(name string, client *fake.Clientset) error {
	if len(client.Actions()) == 0 {
		return nil
	}

	for _, action := range client.Actions() {
		deleteAction, ok := action.(clienttesting.DeleteAction)
		if !ok {
			return fmt.Errorf("unexpected action: %v", client.Actions())
		}
		if deleteAction.GetName() != name {
			return fmt.Errorf("bad name %v", deleteAction)
		}
	}

	return nil
}

func TestSync(t *testing.T) {
	tests := []struct {
		name               string
		apiServiceName     string
		addAPIServices     []*apiregistration.APIService
		updateAPIServices  []*apiregistration.APIService
		addSyncAPIServices []*apiregistration.APIService
		delSyncAPIServices []string
		expectedResults    func(name string, client *fake.Clientset) error
	}{
		{
			name:               "adding an API service which isn't auto-managed does nothing",
			apiServiceName:     "foo",
			addAPIServices:     []*apiregistration.APIService{newAPIService("foo")},
			updateAPIServices:  []*apiregistration.APIService{},
			addSyncAPIServices: []*apiregistration.APIService{},
			delSyncAPIServices: []string{},
			expectedResults:    checkForNothing,
		},
		{
			name:               "adding one to auto-register should create",
			apiServiceName:     "foo",
			addAPIServices:     []*apiregistration.APIService{},
			updateAPIServices:  []*apiregistration.APIService{},
			addSyncAPIServices: []*apiregistration.APIService{newAPIService("foo")},
			delSyncAPIServices: []string{},
			expectedResults:    checkForCreate,
		},
		{
			name:               "duplicate AddAPIServiceToSync don't panic",
			apiServiceName:     "foo",
			addAPIServices:     []*apiregistration.APIService{newAutoRegisterManagedAPIService("foo")},
			updateAPIServices:  []*apiregistration.APIService{},
			addSyncAPIServices: []*apiregistration.APIService{newAutoRegisterManagedAPIService("foo"), newAutoRegisterManagedAPIService("foo")},
			delSyncAPIServices: []string{},
			expectedResults:    checkForNothing,
		},
		{
			name:               "duplicate RemoveAPIServiceToSync don't panic",
			apiServiceName:     "foo",
			addAPIServices:     []*apiregistration.APIService{newAutoRegisterManagedAPIService("foo")},
			updateAPIServices:  []*apiregistration.APIService{},
			addSyncAPIServices: []*apiregistration.APIService{},
			delSyncAPIServices: []string{"foo", "foo"},
			expectedResults:    checkForDelete,
		},
		{
			name:               "removing auto-manged then RemoveAPIService should not touch APIService",
			apiServiceName:     "foo",
			addAPIServices:     []*apiregistration.APIService{},
			updateAPIServices:  []*apiregistration.APIService{newAPIService("foo")},
			addSyncAPIServices: []*apiregistration.APIService{},
			delSyncAPIServices: []string{"foo"},
			expectedResults:    checkForNothing,
		},
		{
			name:               "create managed apiservice without a matching request",
			apiServiceName:     "foo",
			addAPIServices:     []*apiregistration.APIService{newAPIService("foo")},
			updateAPIServices:  []*apiregistration.APIService{newAutoRegisterManagedAPIService("foo")},
			addSyncAPIServices: []*apiregistration.APIService{},
			delSyncAPIServices: []string{},
			expectedResults:    checkForDelete,
		},
		{
			name:               "modifying it should result in stomping",
			apiServiceName:     "foo",
			addAPIServices:     []*apiregistration.APIService{},
			updateAPIServices:  []*apiregistration.APIService{newAutoRegisterManagedModifiedAPIService("foo")},
			addSyncAPIServices: []*apiregistration.APIService{newAutoRegisterManagedAPIService("foo")},
			delSyncAPIServices: []string{},
			expectedResults:    checkForUpdate,
		},
	}

	for _, test := range tests {
		fakeClient := fake.NewSimpleClientset()
		apiServiceIndexer := cache.NewIndexer(cache.MetaNamespaceKeyFunc, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc})

		c := autoRegisterController{
			apiServiceClient:  fakeClient.Apiregistration(),
			apiServiceLister:  listers.NewAPIServiceLister(apiServiceIndexer),
			apiServicesToSync: map[string]*apiregistration.APIService{},
			queue:             workqueue.NewNamedRateLimitingQueue(workqueue.DefaultControllerRateLimiter(), "autoregister"),
		}

		for _, obj := range test.addAPIServices {
			apiServiceIndexer.Add(obj)
		}

		for _, obj := range test.updateAPIServices {
			apiServiceIndexer.Update(obj)
		}

		for _, obj := range test.addSyncAPIServices {
			c.AddAPIServiceToSync(obj)
		}

		for _, objName := range test.delSyncAPIServices {
			c.RemoveAPIServiceToSync(objName)
		}

		c.checkAPIService(test.apiServiceName)

		//compare the expected results
		err := test.expectedResults(test.apiServiceName, fakeClient)
		if err != nil {
			t.Errorf("%s %v", test.name, err)
		}
	}
}
