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

// This file was automatically generated by informer-gen with arguments: --input-dirs=[k8s.io/kubernetes/cmd/kube-aggregator/pkg/apis/apiregistration,k8s.io/kubernetes/cmd/kube-aggregator/pkg/apis/apiregistration/v1alpha1] --internal-clientset-package=k8s.io/kubernetes/cmd/kube-aggregator/pkg/client/clientset_generated/internalclientset --listers-package=k8s.io/kubernetes/cmd/kube-aggregator/pkg/client/listers --output-package=k8s.io/kubernetes/cmd/kube-aggregator/pkg/client/informers --versioned-clientset-package=k8s.io/kubernetes/cmd/kube-aggregator/pkg/client/clientset_generated/clientset

package informers

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	clientset "k8s.io/kubernetes/cmd/kube-aggregator/pkg/client/clientset_generated/clientset"
	internalclientset "k8s.io/kubernetes/cmd/kube-aggregator/pkg/client/clientset_generated/internalclientset"
	apiregistration "k8s.io/kubernetes/cmd/kube-aggregator/pkg/client/informers/apiregistration"
	internalinterfaces "k8s.io/kubernetes/cmd/kube-aggregator/pkg/client/informers/internalinterfaces"
	cache "k8s.io/kubernetes/pkg/client/cache"
	reflect "reflect"
	sync "sync"
	time "time"
)

type sharedInformerFactory struct {
	internalClient  internalclientset.Interface
	versionedClient clientset.Interface
	lock            sync.Mutex
	defaultResync   time.Duration

	informers map[reflect.Type]cache.SharedIndexInformer
	// startedInformers is used for tracking which informers have been started.
	// This allows Start() to be called multiple times safely.
	startedInformers map[reflect.Type]bool
}

// NewSharedInformerFactory constructs a new instance of sharedInformerFactory
func NewSharedInformerFactory(internalClient internalclientset.Interface, versionedClient clientset.Interface, defaultResync time.Duration) SharedInformerFactory {
	return &sharedInformerFactory{
		internalClient:   internalClient,
		versionedClient:  versionedClient,
		defaultResync:    defaultResync,
		informers:        make(map[reflect.Type]cache.SharedIndexInformer),
		startedInformers: make(map[reflect.Type]bool),
	}
}

// Start initializes all requested informers.
func (f *sharedInformerFactory) Start(stopCh <-chan struct{}) {
	f.lock.Lock()
	defer f.lock.Unlock()

	for informerType, informer := range f.informers {
		if !f.startedInformers[informerType] {
			go informer.Run(stopCh)
			f.startedInformers[informerType] = true
		}
	}
}

// InternalInformerFor returns the SharedIndexInformer for obj using an internal
// client.
func (f *sharedInformerFactory) InternalInformerFor(obj runtime.Object, newFunc internalinterfaces.NewInternalInformerFunc) cache.SharedIndexInformer {
	f.lock.Lock()
	defer f.lock.Unlock()

	informerType := reflect.TypeOf(obj)
	informer, exists := f.informers[informerType]
	if exists {
		return informer
	}
	informer = newFunc(f.internalClient, f.defaultResync)
	f.informers[informerType] = informer

	return informer
}

// VersionedInformerFor returns the SharedIndexInformer for obj using a
// versioned client.
func (f *sharedInformerFactory) VersionedInformerFor(obj runtime.Object, newFunc internalinterfaces.NewVersionedInformerFunc) cache.SharedIndexInformer {
	f.lock.Lock()
	defer f.lock.Unlock()

	informerType := reflect.TypeOf(obj)
	informer, exists := f.informers[informerType]
	if exists {
		return informer
	}
	informer = newFunc(f.versionedClient, f.defaultResync)
	f.informers[informerType] = informer

	return informer
}

// SharedInformerFactory provides shared informers for resources in all known
// API group versions.
type SharedInformerFactory interface {
	internalinterfaces.SharedInformerFactory
	ForResource(resource schema.GroupVersionResource) (GenericInformer, error)

	Apiregistration() apiregistration.Interface
}

func (f *sharedInformerFactory) Apiregistration() apiregistration.Interface {
	return apiregistration.New(f)
}
