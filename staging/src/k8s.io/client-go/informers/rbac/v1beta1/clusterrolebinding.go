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

	rbacv1beta1 "k8s.io/api/rbac/v1beta1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	internalinterfaces "k8s.io/client-go/informers/internalinterfaces"
	kubernetes "k8s.io/client-go/kubernetes"
	v1beta1 "k8s.io/client-go/listers/rbac/v1beta1"
	cache "k8s.io/client-go/tools/cache"
)

// ClusterRoleBindingInformer provides access to a shared informer and lister for
// ClusterRoleBindings.
type ClusterRoleBindingInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1beta1.ClusterRoleBindingLister
}

type clusterRoleBindingInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
}

// NewClusterRoleBindingInformer constructs a new informer for ClusterRoleBinding type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewClusterRoleBindingInformer(client kubernetes.Interface, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredClusterRoleBindingInformer(client, resyncPeriod, indexers, nil)
}

// NewFilteredClusterRoleBindingInformer constructs a new informer for ClusterRoleBinding type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredClusterRoleBindingInformer(client kubernetes.Interface, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.RbacV1beta1().ClusterRoleBindings().List(context.TODO(), options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.RbacV1beta1().ClusterRoleBindings().Watch(context.TODO(), options)
			},
		},
		&rbacv1beta1.ClusterRoleBinding{},
		resyncPeriod,
		indexers,
	)
}

func (f *clusterRoleBindingInformer) defaultInformer(client kubernetes.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredClusterRoleBindingInformer(client, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *clusterRoleBindingInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&rbacv1beta1.ClusterRoleBinding{}, f.defaultInformer)
}

func (f *clusterRoleBindingInformer) Lister() v1beta1.ClusterRoleBindingLister {
	return v1beta1.NewClusterRoleBindingLister(f.Informer().GetIndexer())
}
