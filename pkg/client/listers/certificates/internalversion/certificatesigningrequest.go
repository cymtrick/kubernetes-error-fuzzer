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

// This file was automatically generated by lister-gen with arguments: --input-dirs=[k8s.io/kubernetes/pkg/api,k8s.io/kubernetes/pkg/api/v1,k8s.io/kubernetes/pkg/apis/abac,k8s.io/kubernetes/pkg/apis/abac/v0,k8s.io/kubernetes/pkg/apis/abac/v1beta1,k8s.io/kubernetes/pkg/apis/apps,k8s.io/kubernetes/pkg/apis/apps/v1beta1,k8s.io/kubernetes/pkg/apis/authentication,k8s.io/kubernetes/pkg/apis/authentication/v1beta1,k8s.io/kubernetes/pkg/apis/authorization,k8s.io/kubernetes/pkg/apis/authorization/v1beta1,k8s.io/kubernetes/pkg/apis/autoscaling,k8s.io/kubernetes/pkg/apis/autoscaling/v1,k8s.io/kubernetes/pkg/apis/batch,k8s.io/kubernetes/pkg/apis/batch/v1,k8s.io/kubernetes/pkg/apis/batch/v2alpha1,k8s.io/kubernetes/pkg/apis/certificates,k8s.io/kubernetes/pkg/apis/certificates/v1alpha1,k8s.io/kubernetes/pkg/apis/componentconfig,k8s.io/kubernetes/pkg/apis/componentconfig/v1alpha1,k8s.io/kubernetes/pkg/apis/extensions,k8s.io/kubernetes/pkg/apis/extensions/v1beta1,k8s.io/kubernetes/pkg/apis/imagepolicy,k8s.io/kubernetes/pkg/apis/imagepolicy/v1alpha1,k8s.io/kubernetes/pkg/apis/meta/v1,k8s.io/kubernetes/pkg/apis/policy,k8s.io/kubernetes/pkg/apis/policy/v1alpha1,k8s.io/kubernetes/pkg/apis/policy/v1beta1,k8s.io/kubernetes/pkg/apis/rbac,k8s.io/kubernetes/pkg/apis/rbac/v1alpha1,k8s.io/kubernetes/pkg/apis/storage,k8s.io/kubernetes/pkg/apis/storage/v1beta1]

package internalversion

import (
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	api "k8s.io/kubernetes/pkg/api"
	certificates "k8s.io/kubernetes/pkg/apis/certificates"
	"k8s.io/kubernetes/pkg/client/cache"
)

// CertificateSigningRequestLister helps list CertificateSigningRequests.
type CertificateSigningRequestLister interface {
	// List lists all CertificateSigningRequests in the indexer.
	List(selector labels.Selector) (ret []*certificates.CertificateSigningRequest, err error)
	// Get retrieves the CertificateSigningRequest from the index for a given name.
	Get(name string) (*certificates.CertificateSigningRequest, error)
	CertificateSigningRequestListerExpansion
}

// certificateSigningRequestLister implements the CertificateSigningRequestLister interface.
type certificateSigningRequestLister struct {
	indexer cache.Indexer
}

// NewCertificateSigningRequestLister returns a new CertificateSigningRequestLister.
func NewCertificateSigningRequestLister(indexer cache.Indexer) CertificateSigningRequestLister {
	return &certificateSigningRequestLister{indexer: indexer}
}

// List lists all CertificateSigningRequests in the indexer.
func (s *certificateSigningRequestLister) List(selector labels.Selector) (ret []*certificates.CertificateSigningRequest, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*certificates.CertificateSigningRequest))
	})
	return ret, err
}

// Get retrieves the CertificateSigningRequest from the index for a given name.
func (s *certificateSigningRequestLister) Get(name string) (*certificates.CertificateSigningRequest, error) {
	key := &certificates.CertificateSigningRequest{ObjectMeta: api.ObjectMeta{Name: name}}
	obj, exists, err := s.indexer.Get(key)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(certificates.Resource("certificatesigningrequest"), name)
	}
	return obj.(*certificates.CertificateSigningRequest), nil
}
