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

package internalversion

import (
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
	api "k8s.io/kubernetes/pkg/api"
	internalclientset "k8s.io/kubernetes/pkg/client/clientset_generated/internalclientset"
	internalinterfaces "k8s.io/kubernetes/pkg/client/informers/informers_generated/internalversion/internalinterfaces"
	internalversion "k8s.io/kubernetes/pkg/client/listers/core/internalversion"
	time "time"
)

// PersistentVolumeInformer provides access to a shared informer and lister for
// PersistentVolumes.
type PersistentVolumeInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() internalversion.PersistentVolumeLister
}

type persistentVolumeInformer struct {
	factory internalinterfaces.SharedInformerFactory
}

// NewPersistentVolumeInformer constructs a new informer for PersistentVolume type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewPersistentVolumeInformer(client internalclientset.Interface, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				return client.Core().PersistentVolumes().List(options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				return client.Core().PersistentVolumes().Watch(options)
			},
		},
		&api.PersistentVolume{},
		resyncPeriod,
		indexers,
	)
}

func defaultPersistentVolumeInformer(client internalclientset.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewPersistentVolumeInformer(client, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc})
}

func (f *persistentVolumeInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&api.PersistentVolume{}, defaultPersistentVolumeInformer)
}

func (f *persistentVolumeInformer) Lister() internalversion.PersistentVolumeLister {
	return internalversion.NewPersistentVolumeLister(f.Informer().GetIndexer())
}
