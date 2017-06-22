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

package v1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	internalinterfaces "k8s.io/client-go/informers/internalinterfaces"
	kubernetes "k8s.io/client-go/kubernetes"
	v1 "k8s.io/client-go/listers/core/v1"
	api_v1 "k8s.io/api/core/v1"
	cache "k8s.io/client-go/tools/cache"
	time "time"
)

// PersistentVolumeInformer provides access to a shared informer and lister for
// PersistentVolumes.
type PersistentVolumeInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1.PersistentVolumeLister
}

type persistentVolumeInformer struct {
	factory internalinterfaces.SharedInformerFactory
}

func newPersistentVolumeInformer(client kubernetes.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	sharedIndexInformer := cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options meta_v1.ListOptions) (runtime.Object, error) {
				return client.CoreV1().PersistentVolumes().List(options)
			},
			WatchFunc: func(options meta_v1.ListOptions) (watch.Interface, error) {
				return client.CoreV1().PersistentVolumes().Watch(options)
			},
		},
		&api_v1.PersistentVolume{},
		resyncPeriod,
		cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc},
	)

	return sharedIndexInformer
}

func (f *persistentVolumeInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&api_v1.PersistentVolume{}, newPersistentVolumeInformer)
}

func (f *persistentVolumeInformer) Lister() v1.PersistentVolumeLister {
	return v1.NewPersistentVolumeLister(f.Informer().GetIndexer())
}
