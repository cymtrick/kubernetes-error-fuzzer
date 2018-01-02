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
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
	core "k8s.io/kubernetes/pkg/apis/core"
	internalclientset "k8s.io/kubernetes/pkg/client/clientset_generated/internalclientset"
	internalinterfaces "k8s.io/kubernetes/pkg/client/informers/informers_generated/internalversion/internalinterfaces"
	internalversion "k8s.io/kubernetes/pkg/client/listers/core/internalversion"
	time "time"
)

// PersistentVolumeClaimInformer provides access to a shared informer and lister for
// PersistentVolumeClaims.
type PersistentVolumeClaimInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() internalversion.PersistentVolumeClaimLister
}

type persistentVolumeClaimInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewPersistentVolumeClaimInformer constructs a new informer for PersistentVolumeClaim type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewPersistentVolumeClaimInformer(client internalclientset.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredPersistentVolumeClaimInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredPersistentVolumeClaimInformer constructs a new informer for PersistentVolumeClaim type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredPersistentVolumeClaimInformer(client internalclientset.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.Core().PersistentVolumeClaims(namespace).List(options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.Core().PersistentVolumeClaims(namespace).Watch(options)
			},
		},
		&core.PersistentVolumeClaim{},
		resyncPeriod,
		indexers,
	)
}

func (f *persistentVolumeClaimInformer) defaultInformer(client internalclientset.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredPersistentVolumeClaimInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *persistentVolumeClaimInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&core.PersistentVolumeClaim{}, f.defaultInformer)
}

func (f *persistentVolumeClaimInformer) Lister() internalversion.PersistentVolumeClaimLister {
	return internalversion.NewPersistentVolumeClaimLister(f.Informer().GetIndexer())
}
