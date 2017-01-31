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

package v1alpha1

import (
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
	rbac_v1alpha1 "k8s.io/kubernetes/pkg/apis/rbac/v1alpha1"
	clientset "k8s.io/kubernetes/pkg/client/clientset_generated/clientset"
	internalinterfaces "k8s.io/kubernetes/pkg/client/informers/informers_generated/internalinterfaces"
	v1alpha1 "k8s.io/kubernetes/pkg/client/listers/rbac/v1alpha1"
	time "time"
)

// RoleBindingInformer provides access to a shared informer and lister for
// RoleBindings.
type RoleBindingInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1alpha1.RoleBindingLister
}

type roleBindingInformer struct {
	factory internalinterfaces.SharedInformerFactory
}

func newRoleBindingInformer(client clientset.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	sharedIndexInformer := cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				return client.RbacV1alpha1().RoleBindings(v1.NamespaceAll).List(options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				return client.RbacV1alpha1().RoleBindings(v1.NamespaceAll).Watch(options)
			},
		},
		&rbac_v1alpha1.RoleBinding{},
		resyncPeriod,
		cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc},
	)

	return sharedIndexInformer
}

func (f *roleBindingInformer) Informer() cache.SharedIndexInformer {
	return f.factory.VersionedInformerFor(&rbac_v1alpha1.RoleBinding{}, newRoleBindingInformer)
}

func (f *roleBindingInformer) Lister() v1alpha1.RoleBindingLister {
	return v1alpha1.NewRoleBindingLister(f.Informer().GetIndexer())
}
