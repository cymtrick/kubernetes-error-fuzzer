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

package externalversions

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	cache "k8s.io/client-go/tools/cache"
	clientset "k8s.io/kubernetes/pkg/client/clientset_generated/clientset"
	apps "k8s.io/kubernetes/pkg/client/informers/informers_generated/externalversions/apps"
	autoscaling "k8s.io/kubernetes/pkg/client/informers/informers_generated/externalversions/autoscaling"
	batch "k8s.io/kubernetes/pkg/client/informers/informers_generated/externalversions/batch"
	certificates "k8s.io/kubernetes/pkg/client/informers/informers_generated/externalversions/certificates"
	core "k8s.io/kubernetes/pkg/client/informers/informers_generated/externalversions/core"
	extensions "k8s.io/kubernetes/pkg/client/informers/informers_generated/externalversions/extensions"
	internalinterfaces "k8s.io/kubernetes/pkg/client/informers/informers_generated/externalversions/internalinterfaces"
	policy "k8s.io/kubernetes/pkg/client/informers/informers_generated/externalversions/policy"
	rbac "k8s.io/kubernetes/pkg/client/informers/informers_generated/externalversions/rbac"
	storage "k8s.io/kubernetes/pkg/client/informers/informers_generated/externalversions/storage"
	reflect "reflect"
	sync "sync"
	time "time"
)

type sharedInformerFactory struct {
	client        clientset.Interface
	lock          sync.Mutex
	defaultResync time.Duration

	informers map[reflect.Type]cache.SharedIndexInformer
	// startedInformers is used for tracking which informers have been started.
	// This allows Start() to be called multiple times safely.
	startedInformers map[reflect.Type]bool
}

// NewSharedInformerFactory constructs a new instance of sharedInformerFactory
func NewSharedInformerFactory(client clientset.Interface, defaultResync time.Duration) SharedInformerFactory {
	return &sharedInformerFactory{
		client:           client,
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
func (f *sharedInformerFactory) InformerFor(obj runtime.Object, newFunc internalinterfaces.NewInformerFunc) cache.SharedIndexInformer {
	f.lock.Lock()
	defer f.lock.Unlock()

	informerType := reflect.TypeOf(obj)
	informer, exists := f.informers[informerType]
	if exists {
		return informer
	}
	informer = newFunc(f.client, f.defaultResync)
	f.informers[informerType] = informer

	return informer
}

// SharedInformerFactory provides shared informers for resources in all known
// API group versions.
type SharedInformerFactory interface {
	internalinterfaces.SharedInformerFactory
	ForResource(resource schema.GroupVersionResource) (GenericInformer, error)

	Apps() apps.Interface
	Autoscaling() autoscaling.Interface
	Batch() batch.Interface
	Certificates() certificates.Interface
	Core() core.Interface
	Extensions() extensions.Interface
	Policy() policy.Interface
	Rbac() rbac.Interface
	Storage() storage.Interface
}

func (f *sharedInformerFactory) Apps() apps.Interface {
	return apps.New(f)
}

func (f *sharedInformerFactory) Autoscaling() autoscaling.Interface {
	return autoscaling.New(f)
}

func (f *sharedInformerFactory) Batch() batch.Interface {
	return batch.New(f)
}

func (f *sharedInformerFactory) Certificates() certificates.Interface {
	return certificates.New(f)
}

func (f *sharedInformerFactory) Core() core.Interface {
	return core.New(f)
}

func (f *sharedInformerFactory) Extensions() extensions.Interface {
	return extensions.New(f)
}

func (f *sharedInformerFactory) Policy() policy.Interface {
	return policy.New(f)
}

func (f *sharedInformerFactory) Rbac() rbac.Interface {
	return rbac.New(f)
}

func (f *sharedInformerFactory) Storage() storage.Interface {
	return storage.New(f)
}
