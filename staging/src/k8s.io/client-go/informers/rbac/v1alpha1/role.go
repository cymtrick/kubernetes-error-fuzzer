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
	rbac_v1alpha1 "k8s.io/api/rbac/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	internalinterfaces "k8s.io/client-go/informers/internalinterfaces"
	kubernetes "k8s.io/client-go/kubernetes"
	v1alpha1 "k8s.io/client-go/listers/rbac/v1alpha1"
	cache "k8s.io/client-go/tools/cache"
	time "time"
)

// RoleInformer provides access to a shared informer and lister for
// Roles.
type RoleInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1alpha1.RoleLister
}

type roleInformer struct {
	factory internalinterfaces.SharedInformerFactory
}

func newRoleInformer(client kubernetes.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	sharedIndexInformer := cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				return client.RbacV1alpha1().Roles(v1.NamespaceAll).List(options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				return client.RbacV1alpha1().Roles(v1.NamespaceAll).Watch(options)
			},
		},
		&rbac_v1alpha1.Role{},
		resyncPeriod,
		cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc},
	)

	return sharedIndexInformer
}

func (f *roleInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&rbac_v1alpha1.Role{}, newRoleInformer)
}

func (f *roleInformer) Lister() v1alpha1.RoleLister {
	return v1alpha1.NewRoleLister(f.Informer().GetIndexer())
}
