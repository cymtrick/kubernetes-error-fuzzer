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

package v1beta1

import (
	extensions_v1beta1 "k8s.io/api/extensions/v1beta1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	internalinterfaces "k8s.io/client-go/informers/internalinterfaces"
	kubernetes "k8s.io/client-go/kubernetes"
	v1beta1 "k8s.io/client-go/listers/extensions/v1beta1"
	cache "k8s.io/client-go/tools/cache"
	time "time"
)

// ThirdPartyResourceInformer provides access to a shared informer and lister for
// ThirdPartyResources.
type ThirdPartyResourceInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1beta1.ThirdPartyResourceLister
}

type thirdPartyResourceInformer struct {
	factory internalinterfaces.SharedInformerFactory
}

func newThirdPartyResourceInformer(client kubernetes.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	sharedIndexInformer := cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				return client.ExtensionsV1beta1().ThirdPartyResources().List(options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				return client.ExtensionsV1beta1().ThirdPartyResources().Watch(options)
			},
		},
		&extensions_v1beta1.ThirdPartyResource{},
		resyncPeriod,
		cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc},
	)

	return sharedIndexInformer
}

func (f *thirdPartyResourceInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&extensions_v1beta1.ThirdPartyResource{}, newThirdPartyResourceInformer)
}

func (f *thirdPartyResourceInformer) Lister() v1beta1.ThirdPartyResourceLister {
	return v1beta1.NewThirdPartyResourceLister(f.Informer().GetIndexer())
}
