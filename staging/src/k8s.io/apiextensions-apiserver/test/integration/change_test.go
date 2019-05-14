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

package integration

import (
	"fmt"
	"sync"
	"testing"
	"time"

	apiextensionsv1beta1 "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1beta1"
	"k8s.io/apiextensions-apiserver/pkg/client/clientset/clientset"
	"k8s.io/apiextensions-apiserver/test/integration/fixtures"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/dynamic"
)

func TestChangeCRD(t *testing.T) {
	tearDown, config, _, err := fixtures.StartDefaultServer(t)
	if err != nil {
		t.Fatal(err)
	}
	defer tearDown()
	config.QPS = 1000
	config.Burst = 1000
	apiExtensionsClient, err := clientset.NewForConfig(config)
	if err != nil {
		t.Fatal(err)
	}
	dynamicClient, err := dynamic.NewForConfig(config)
	if err != nil {
		t.Fatal(err)
	}

	noxuDefinition := fixtures.NewNoxuCustomResourceDefinition(apiextensionsv1beta1.NamespaceScoped)
	noxuDefinition, err = fixtures.CreateNewCustomResourceDefinition(noxuDefinition, apiExtensionsClient, dynamicClient)
	if err != nil {
		t.Fatal(err)
	}

	ns := "default"
	noxuNamespacedResourceClient := newNamespacedCustomResourceVersionedClient(ns, dynamicClient, noxuDefinition, "v1beta1")

	stopChan := make(chan struct{})

	wg := &sync.WaitGroup{}

	// Set up loop to modify CRD in the background
	wg.Add(1)
	go func() {
		defer wg.Done()
		for {
			select {
			case <-stopChan:
				return
			default:
			}

			noxuDefinitionToUpdate, err := apiExtensionsClient.ApiextensionsV1beta1().CustomResourceDefinitions().Get(noxuDefinition.Name, metav1.GetOptions{})
			if err != nil {
				t.Fatal(err)
			}
			if len(noxuDefinitionToUpdate.Spec.Versions) == 1 {
				v2 := noxuDefinitionToUpdate.Spec.Versions[0]
				v2.Name = "v2"
				v2.Served = true
				v2.Storage = false
				noxuDefinitionToUpdate.Spec.Versions = append(noxuDefinitionToUpdate.Spec.Versions, v2)
			} else {
				noxuDefinitionToUpdate.Spec.Versions = noxuDefinitionToUpdate.Spec.Versions[0:1]
			}
			if _, err := apiExtensionsClient.ApiextensionsV1beta1().CustomResourceDefinitions().Update(noxuDefinitionToUpdate); err != nil && !apierrors.IsConflict(err) {
				t.Fatal(err)
			}
			time.Sleep(10 * time.Millisecond)
		}
	}()

	// Set up 100 loops creating and reading custom resources
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			noxuInstanceToCreate := fixtures.NewNoxuInstance(ns, fmt.Sprintf("foo-%d", i))
			if _, err := noxuNamespacedResourceClient.Create(noxuInstanceToCreate, metav1.CreateOptions{}); err != nil {
				t.Fatal(err)
			}
			for {
				select {
				case <-stopChan:
					return
				default:
					if _, err := noxuNamespacedResourceClient.Get(noxuInstanceToCreate.GetName(), metav1.GetOptions{}); err != nil {
						t.Fatal(err)
					}
				}
				time.Sleep(10 * time.Millisecond)
			}
		}(i)
	}

	// Let all the established get request loops soak
	time.Sleep(5 * time.Second)

	// Tear down
	close(stopChan)

	// Let loops drain
	wg.Wait()
}
