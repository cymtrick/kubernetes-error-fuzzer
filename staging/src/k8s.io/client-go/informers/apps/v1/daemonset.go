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
	apps_v1 "k8s.io/api/apps/v1"
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	internalinterfaces "k8s.io/client-go/informers/internalinterfaces"
	kubernetes "k8s.io/client-go/kubernetes"
	v1 "k8s.io/client-go/listers/apps/v1"
	cache "k8s.io/client-go/tools/cache"
	time "time"
)

// DaemonSetInformer provides access to a shared informer and lister for
// DaemonSets.
type DaemonSetInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1.DaemonSetLister
}

type daemonSetInformer struct {
	factory internalinterfaces.SharedInformerFactory
}

// NewDaemonSetInformer constructs a new informer for DaemonSet type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewDaemonSetInformer(client kubernetes.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options meta_v1.ListOptions) (runtime.Object, error) {
				return client.AppsV1().DaemonSets(namespace).List(options)
			},
			WatchFunc: func(options meta_v1.ListOptions) (watch.Interface, error) {
				return client.AppsV1().DaemonSets(namespace).Watch(options)
			},
		},
		&apps_v1.DaemonSet{},
		resyncPeriod,
		indexers,
	)
}

func defaultDaemonSetInformer(client kubernetes.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewDaemonSetInformer(client, meta_v1.NamespaceAll, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc})
}

func (f *daemonSetInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&apps_v1.DaemonSet{}, defaultDaemonSetInformer)
}

func (f *daemonSetInformer) Lister() v1.DaemonSetLister {
	return v1.NewDaemonSetLister(f.Informer().GetIndexer())
}
