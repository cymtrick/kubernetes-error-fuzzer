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

package apimachinery

import (
	"k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1beta1"
	"k8s.io/apiextensions-apiserver/pkg/client/clientset/clientset"
	"k8s.io/apiextensions-apiserver/test/integration/fixtures"
	utilversion "k8s.io/apimachinery/pkg/util/version"
	"k8s.io/kubernetes/test/e2e/framework"
	e2elog "k8s.io/kubernetes/test/e2e/framework/log"

	"github.com/onsi/ginkgo"
)

var crdVersion = utilversion.MustParseSemantic("v1.7.0")

var _ = SIGDescribe("CustomResourceDefinition resources", func() {

	f := framework.NewDefaultFramework("custom-resource-definition")

	ginkgo.Context("Simple CustomResourceDefinition", func() {
		/*
			Release : v1.9
			Testname: Custom Resource Definition, create
			Description: Create a API extension client, define a random custom resource definition, create the custom resource. API server MUST be able to create the custom resource.
		*/
		framework.ConformanceIt("creating/deleting custom resource definition objects works ", func() {

			config, err := framework.LoadConfig()
			if err != nil {
				e2elog.Failf("failed to load config: %v", err)
			}

			apiExtensionClient, err := clientset.NewForConfig(config)
			if err != nil {
				e2elog.Failf("failed to initialize apiExtensionClient: %v", err)
			}

			randomDefinition := fixtures.NewRandomNameCustomResourceDefinition(v1beta1.ClusterScoped)

			//create CRD and waits for the resource to be recognized and available.
			randomDefinition, err = fixtures.CreateNewCustomResourceDefinition(randomDefinition, apiExtensionClient, f.DynamicClient)
			if err != nil {
				e2elog.Failf("failed to create CustomResourceDefinition: %v", err)
			}

			defer func() {
				err = fixtures.DeleteCustomResourceDefinition(randomDefinition, apiExtensionClient)
				if err != nil {
					e2elog.Failf("failed to delete CustomResourceDefinition: %v", err)
				}
			}()
		})
	})
})
