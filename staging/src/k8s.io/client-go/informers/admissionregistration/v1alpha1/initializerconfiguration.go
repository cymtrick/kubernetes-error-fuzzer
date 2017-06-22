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

// This file was automatically generated by informer-gen

package v1alpha1

import (
	admissionregistration_v1alpha1 "k8s.io/api/admissionregistration/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	internalinterfaces "k8s.io/client-go/informers/internalinterfaces"
	kubernetes "k8s.io/client-go/kubernetes"
	v1alpha1 "k8s.io/client-go/listers/admissionregistration/v1alpha1"
	cache "k8s.io/client-go/tools/cache"
	time "time"
)

// InitializerConfigurationInformer provides access to a shared informer and lister for
// InitializerConfigurations.
type InitializerConfigurationInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1alpha1.InitializerConfigurationLister
}

type initializerConfigurationInformer struct {
	factory internalinterfaces.SharedInformerFactory
}

func newInitializerConfigurationInformer(client kubernetes.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	sharedIndexInformer := cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				return client.AdmissionregistrationV1alpha1().InitializerConfigurations().List(options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				return client.AdmissionregistrationV1alpha1().InitializerConfigurations().Watch(options)
			},
		},
		&admissionregistration_v1alpha1.InitializerConfiguration{},
		resyncPeriod,
		cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc},
	)

	return sharedIndexInformer
}

func (f *initializerConfigurationInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&admissionregistration_v1alpha1.InitializerConfiguration{}, newInitializerConfigurationInformer)
}

func (f *initializerConfigurationInformer) Lister() v1alpha1.InitializerConfigurationLister {
	return v1alpha1.NewInitializerConfigurationLister(f.Informer().GetIndexer())
}
