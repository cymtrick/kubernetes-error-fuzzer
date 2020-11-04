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

package v1beta1

import (
	"context"
	time "time"

	flowcontrolv1beta1 "k8s.io/api/flowcontrol/v1beta1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	internalinterfaces "k8s.io/client-go/informers/internalinterfaces"
	kubernetes "k8s.io/client-go/kubernetes"
	v1beta1 "k8s.io/client-go/listers/flowcontrol/v1beta1"
	cache "k8s.io/client-go/tools/cache"
)

// PriorityLevelConfigurationInformer provides access to a shared informer and lister for
// PriorityLevelConfigurations.
type PriorityLevelConfigurationInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1beta1.PriorityLevelConfigurationLister
}

type priorityLevelConfigurationInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
}

// NewPriorityLevelConfigurationInformer constructs a new informer for PriorityLevelConfiguration type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewPriorityLevelConfigurationInformer(client kubernetes.Interface, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredPriorityLevelConfigurationInformer(client, resyncPeriod, indexers, nil)
}

// NewFilteredPriorityLevelConfigurationInformer constructs a new informer for PriorityLevelConfiguration type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredPriorityLevelConfigurationInformer(client kubernetes.Interface, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.FlowcontrolV1beta1().PriorityLevelConfigurations().List(context.TODO(), options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.FlowcontrolV1beta1().PriorityLevelConfigurations().Watch(context.TODO(), options)
			},
		},
		&flowcontrolv1beta1.PriorityLevelConfiguration{},
		resyncPeriod,
		indexers,
	)
}

func (f *priorityLevelConfigurationInformer) defaultInformer(client kubernetes.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredPriorityLevelConfigurationInformer(client, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *priorityLevelConfigurationInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&flowcontrolv1beta1.PriorityLevelConfiguration{}, f.defaultInformer)
}

func (f *priorityLevelConfigurationInformer) Lister() v1beta1.PriorityLevelConfigurationLister {
	return v1beta1.NewPriorityLevelConfigurationLister(f.Informer().GetIndexer())
}
