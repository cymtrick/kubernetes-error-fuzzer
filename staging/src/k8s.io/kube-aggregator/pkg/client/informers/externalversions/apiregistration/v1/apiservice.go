/*
Copyright The Kubernetes Authors.

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

// Code generated by informer-gen. DO NOT EDIT.

package v1

import (
	"context"
	time "time"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
	apiregistrationv1 "k8s.io/kube-aggregator/pkg/apis/apiregistration/v1"
	clientset "k8s.io/kube-aggregator/pkg/client/clientset_generated/clientset"
	internalinterfaces "k8s.io/kube-aggregator/pkg/client/informers/externalversions/internalinterfaces"
	v1 "k8s.io/kube-aggregator/pkg/client/listers/apiregistration/v1"
)

// APIServiceInformer provides access to a shared informer and lister for
// APIServices.
type APIServiceInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1.APIServiceLister
}

type aPIServiceInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
}

// NewAPIServiceInformer constructs a new informer for APIService type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewAPIServiceInformer(client clientset.Interface, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredAPIServiceInformer(client, resyncPeriod, indexers, nil)
}

// NewFilteredAPIServiceInformer constructs a new informer for APIService type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredAPIServiceInformer(client clientset.Interface, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options metav1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.ApiregistrationV1().APIServices().List(context.TODO(), options)
			},
			WatchFunc: func(options metav1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.ApiregistrationV1().APIServices().Watch(context.TODO(), options)
			},
		},
		&apiregistrationv1.APIService{},
		resyncPeriod,
		indexers,
	)
}

func (f *aPIServiceInformer) defaultInformer(client clientset.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredAPIServiceInformer(client, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *aPIServiceInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&apiregistrationv1.APIService{}, f.defaultInformer)
}

func (f *aPIServiceInformer) Lister() v1.APIServiceLister {
	return v1.NewAPIServiceLister(f.Informer().GetIndexer())
}
