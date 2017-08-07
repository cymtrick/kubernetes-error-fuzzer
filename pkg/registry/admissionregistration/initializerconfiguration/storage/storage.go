/*
Copyright 2015 The Kubernetes Authors.

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

package storage

import (
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apiserver/pkg/registry/generic"
	genericregistry "k8s.io/apiserver/pkg/registry/generic/registry"
	"k8s.io/kubernetes/pkg/api"
	"k8s.io/kubernetes/pkg/apis/admissionregistration"
	"k8s.io/kubernetes/pkg/registry/admissionregistration/initializerconfiguration"
	"k8s.io/kubernetes/pkg/registry/cachesize"
)

// rest implements a RESTStorage for pod disruption budgets against etcd
type REST struct {
	*genericregistry.Store
}

// NewREST returns a RESTStorage object that will work against pod disruption budgets.
func NewREST(optsGetter generic.RESTOptionsGetter) *REST {
	store := &genericregistry.Store{
		Copier:      api.Scheme,
		NewFunc:     func() runtime.Object { return &admissionregistration.InitializerConfiguration{} },
		NewListFunc: func() runtime.Object { return &admissionregistration.InitializerConfigurationList{} },
		ObjectNameFunc: func(obj runtime.Object) (string, error) {
			return obj.(*admissionregistration.InitializerConfiguration).Name, nil
		},
		PredicateFunc:            initializerconfiguration.MatchInitializerConfiguration,
		DefaultQualifiedResource: admissionregistration.Resource("initializerconfigurations"),
		WatchCacheSize:           cachesize.GetWatchCacheSizeByResource("initializerconfigurations"),

		CreateStrategy: initializerconfiguration.Strategy,
		UpdateStrategy: initializerconfiguration.Strategy,
		DeleteStrategy: initializerconfiguration.Strategy,
	}
	options := &generic.StoreOptions{RESTOptions: optsGetter, AttrFunc: initializerconfiguration.GetAttrs}
	if err := store.CompleteWithOptions(options); err != nil {
		panic(err) // TODO: Propagate error up
	}
	return &REST{store}
}
