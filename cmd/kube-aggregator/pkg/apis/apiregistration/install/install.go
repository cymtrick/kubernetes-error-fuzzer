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

package install

import (
	"k8s.io/apimachinery/pkg/apimachinery/registered"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/util/sets"
	"k8s.io/kubernetes/cmd/kube-aggregator/pkg/apis/apiregistration"
	"k8s.io/kubernetes/cmd/kube-aggregator/pkg/apis/apiregistration/v1alpha1"
	"k8s.io/kubernetes/pkg/api"
	"k8s.io/kubernetes/pkg/apimachinery/announced"
)

func init() {
	Install(api.Registry, api.Scheme)
}

// Install registers the API group and adds types to a scheme
func Install(registry *registered.APIRegistrationManager, scheme *runtime.Scheme) {
	if err := announced.NewGroupMetaFactory(
		&announced.GroupMetaFactoryArgs{
			GroupName:                  apiregistration.GroupName,
			RootScopedKinds:            sets.NewString("APIService"),
			VersionPreferenceOrder:     []string{v1alpha1.SchemeGroupVersion.Version},
			ImportPrefix:               "k8s.io/kubernetes/cmd/kube-aggregator/pkg/apis/apiregistration",
			AddInternalObjectsToScheme: apiregistration.AddToScheme,
		},
		announced.VersionToSchemeFunc{
			v1alpha1.SchemeGroupVersion.Version: v1alpha1.AddToScheme,
		},
	).Announce().RegisterAndEnable(registry, scheme); err != nil {
		panic(err)
	}
}
