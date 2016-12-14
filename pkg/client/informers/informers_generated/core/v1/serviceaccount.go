/*
Copyright 2016 The Kubernetes Authors.

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

// This file was automatically generated by informer-gen with arguments: --input-dirs=[k8s.io/kubernetes/pkg/api,k8s.io/kubernetes/pkg/api/v1,k8s.io/kubernetes/pkg/apis/abac,k8s.io/kubernetes/pkg/apis/abac/v0,k8s.io/kubernetes/pkg/apis/abac/v1beta1,k8s.io/kubernetes/pkg/apis/apps,k8s.io/kubernetes/pkg/apis/apps/v1beta1,k8s.io/kubernetes/pkg/apis/authentication,k8s.io/kubernetes/pkg/apis/authentication/v1beta1,k8s.io/kubernetes/pkg/apis/authorization,k8s.io/kubernetes/pkg/apis/authorization/v1beta1,k8s.io/kubernetes/pkg/apis/autoscaling,k8s.io/kubernetes/pkg/apis/autoscaling/v1,k8s.io/kubernetes/pkg/apis/batch,k8s.io/kubernetes/pkg/apis/batch/v1,k8s.io/kubernetes/pkg/apis/batch/v2alpha1,k8s.io/kubernetes/pkg/apis/certificates,k8s.io/kubernetes/pkg/apis/certificates/v1alpha1,k8s.io/kubernetes/pkg/apis/componentconfig,k8s.io/kubernetes/pkg/apis/componentconfig/v1alpha1,k8s.io/kubernetes/pkg/apis/extensions,k8s.io/kubernetes/pkg/apis/extensions/v1beta1,k8s.io/kubernetes/pkg/apis/imagepolicy,k8s.io/kubernetes/pkg/apis/imagepolicy/v1alpha1,k8s.io/kubernetes/pkg/apis/meta/v1,k8s.io/kubernetes/pkg/apis/policy,k8s.io/kubernetes/pkg/apis/policy/v1beta1,k8s.io/kubernetes/pkg/apis/rbac,k8s.io/kubernetes/pkg/apis/rbac/v1alpha1,k8s.io/kubernetes/pkg/apis/storage,k8s.io/kubernetes/pkg/apis/storage/v1beta1] --internal-clientset-package=k8s.io/kubernetes/pkg/client/clientset_generated/internalclientset --listers-package=k8s.io/kubernetes/pkg/client/listers --versioned-clientset-package=k8s.io/kubernetes/pkg/client/clientset_generated/clientset

package v1

import (
	api_v1 "k8s.io/kubernetes/pkg/api/v1"
	cache "k8s.io/kubernetes/pkg/client/cache"
	clientset "k8s.io/kubernetes/pkg/client/clientset_generated/clientset"
	internalinterfaces "k8s.io/kubernetes/pkg/client/informers/informers_generated/internalinterfaces"
	v1 "k8s.io/kubernetes/pkg/client/listers/core/v1"
	runtime "k8s.io/kubernetes/pkg/runtime"
	watch "k8s.io/kubernetes/pkg/watch"
	time "time"
)

// ServiceAccountInformer provides access to a shared informer and lister for
// ServiceAccounts.
type ServiceAccountInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1.ServiceAccountLister
}

type serviceAccountInformer struct {
	factory internalinterfaces.SharedInformerFactory
}

func newServiceAccountInformer(client clientset.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	sharedIndexInformer := cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options api_v1.ListOptions) (runtime.Object, error) {
				return client.CoreV1().ServiceAccounts(api_v1.NamespaceAll).List(options)
			},
			WatchFunc: func(options api_v1.ListOptions) (watch.Interface, error) {
				return client.CoreV1().ServiceAccounts(api_v1.NamespaceAll).Watch(options)
			},
		},
		&api_v1.ServiceAccount{},
		resyncPeriod,
		cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc},
	)

	return sharedIndexInformer
}

func (f *serviceAccountInformer) Informer() cache.SharedIndexInformer {
	return f.factory.VersionedInformerFor(&api_v1.ServiceAccount{}, newServiceAccountInformer)
}

func (f *serviceAccountInformer) Lister() v1.ServiceAccountLister {
	return v1.NewServiceAccountLister(f.Informer().GetIndexer())
}
