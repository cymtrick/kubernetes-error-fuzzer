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
	time "time"

	policyv1beta1 "k8s.io/api/policy/v1beta1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	internalinterfaces "k8s.io/client-go/informers/internalinterfaces"
	kubernetes "k8s.io/client-go/kubernetes"
	v1beta1 "k8s.io/client-go/listers/policy/v1beta1"
	cache "k8s.io/client-go/tools/cache"
)

// PodSecurityPolicyInformer provides access to a shared informer and lister for
// PodSecurityPolicies.
type PodSecurityPolicyInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1beta1.PodSecurityPolicyLister
}

type podSecurityPolicyInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
}

// NewPodSecurityPolicyInformer constructs a new informer for PodSecurityPolicy type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewPodSecurityPolicyInformer(client kubernetes.Interface, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredPodSecurityPolicyInformer(client, resyncPeriod, indexers, nil)
}

// NewFilteredPodSecurityPolicyInformer constructs a new informer for PodSecurityPolicy type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredPodSecurityPolicyInformer(client kubernetes.Interface, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.PolicyV1beta1().PodSecurityPolicies().List(options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.PolicyV1beta1().PodSecurityPolicies().Watch(options)
			},
		},
		&policyv1beta1.PodSecurityPolicy{},
		resyncPeriod,
		indexers,
	)
}

func (f *podSecurityPolicyInformer) defaultInformer(client kubernetes.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredPodSecurityPolicyInformer(client, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *podSecurityPolicyInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&policyv1beta1.PodSecurityPolicy{}, f.defaultInformer)
}

func (f *podSecurityPolicyInformer) Lister() v1beta1.PodSecurityPolicyLister {
	return v1beta1.NewPodSecurityPolicyLister(f.Informer().GetIndexer())
}
