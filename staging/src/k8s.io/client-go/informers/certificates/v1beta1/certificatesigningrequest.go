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

package v1beta1

import (
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	internalinterfaces "k8s.io/client-go/informers/internalinterfaces"
	kubernetes "k8s.io/client-go/kubernetes"
	v1beta1 "k8s.io/client-go/listers/certificates/v1beta1"
	certificates_v1beta1 "k8s.io/api/certificates/v1beta1"
	cache "k8s.io/client-go/tools/cache"
	time "time"
)

// CertificateSigningRequestInformer provides access to a shared informer and lister for
// CertificateSigningRequests.
type CertificateSigningRequestInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1beta1.CertificateSigningRequestLister
}

type certificateSigningRequestInformer struct {
	factory internalinterfaces.SharedInformerFactory
}

func newCertificateSigningRequestInformer(client kubernetes.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	sharedIndexInformer := cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				return client.CertificatesV1beta1().CertificateSigningRequests().List(options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				return client.CertificatesV1beta1().CertificateSigningRequests().Watch(options)
			},
		},
		&certificates_v1beta1.CertificateSigningRequest{},
		resyncPeriod,
		cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc},
	)

	return sharedIndexInformer
}

func (f *certificateSigningRequestInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&certificates_v1beta1.CertificateSigningRequest{}, newCertificateSigningRequestInformer)
}

func (f *certificateSigningRequestInformer) Lister() v1beta1.CertificateSigningRequestLister {
	return v1beta1.NewCertificateSigningRequestLister(f.Informer().GetIndexer())
}
