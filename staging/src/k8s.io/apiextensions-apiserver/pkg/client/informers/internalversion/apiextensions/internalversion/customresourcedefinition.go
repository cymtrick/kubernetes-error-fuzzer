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

// This file was automatically generated by informer-gen

package internalversion

import (
	time "time"

	apiextensions "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions"
	internalclientset "k8s.io/apiextensions-apiserver/pkg/client/clientset/internalclientset"
	internalinterfaces "k8s.io/apiextensions-apiserver/pkg/client/informers/internalversion/internalinterfaces"
	internalversion "k8s.io/apiextensions-apiserver/pkg/client/listers/apiextensions/internalversion"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// CustomResourceDefinitionInformer provides access to a shared informer and lister for
// CustomResourceDefinitions.
type CustomResourceDefinitionInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() internalversion.CustomResourceDefinitionLister
}

type customResourceDefinitionInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
}

// NewCustomResourceDefinitionInformer constructs a new informer for CustomResourceDefinition type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewCustomResourceDefinitionInformer(client internalclientset.Interface, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredCustomResourceDefinitionInformer(client, resyncPeriod, indexers, nil)
}

// NewFilteredCustomResourceDefinitionInformer constructs a new informer for CustomResourceDefinition type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredCustomResourceDefinitionInformer(client internalclientset.Interface, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.Apiextensions().CustomResourceDefinitions().List(options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.Apiextensions().CustomResourceDefinitions().Watch(options)
			},
		},
		&apiextensions.CustomResourceDefinition{},
		resyncPeriod,
		indexers,
	)
}

func (f *customResourceDefinitionInformer) defaultInformer(client internalclientset.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredCustomResourceDefinitionInformer(client, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *customResourceDefinitionInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&apiextensions.CustomResourceDefinition{}, f.defaultInformer)
}

func (f *customResourceDefinitionInformer) Lister() internalversion.CustomResourceDefinitionLister {
	return internalversion.NewCustomResourceDefinitionLister(f.Informer().GetIndexer())
}
