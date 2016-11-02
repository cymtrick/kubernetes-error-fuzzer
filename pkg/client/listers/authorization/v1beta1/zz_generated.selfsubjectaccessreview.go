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

// This file was automatically generated by lister-gen with arguments: --input-dirs=[k8s.io/kubernetes/pkg/api,k8s.io/kubernetes/pkg/api/v1,k8s.io/kubernetes/pkg/apis/abac,k8s.io/kubernetes/pkg/apis/abac/v0,k8s.io/kubernetes/pkg/apis/abac/v1beta1,k8s.io/kubernetes/pkg/apis/apps,k8s.io/kubernetes/pkg/apis/apps/v1alpha1,k8s.io/kubernetes/pkg/apis/authentication,k8s.io/kubernetes/pkg/apis/authentication/v1beta1,k8s.io/kubernetes/pkg/apis/authorization,k8s.io/kubernetes/pkg/apis/authorization/v1beta1,k8s.io/kubernetes/pkg/apis/autoscaling,k8s.io/kubernetes/pkg/apis/autoscaling/v1,k8s.io/kubernetes/pkg/apis/batch,k8s.io/kubernetes/pkg/apis/batch/v1,k8s.io/kubernetes/pkg/apis/batch/v2alpha1,k8s.io/kubernetes/pkg/apis/certificates,k8s.io/kubernetes/pkg/apis/certificates/v1alpha1,k8s.io/kubernetes/pkg/apis/componentconfig,k8s.io/kubernetes/pkg/apis/componentconfig/v1alpha1,k8s.io/kubernetes/pkg/apis/extensions,k8s.io/kubernetes/pkg/apis/extensions/v1beta1,k8s.io/kubernetes/pkg/apis/imagepolicy,k8s.io/kubernetes/pkg/apis/imagepolicy/v1alpha1,k8s.io/kubernetes/pkg/apis/policy,k8s.io/kubernetes/pkg/apis/policy/v1alpha1,k8s.io/kubernetes/pkg/apis/rbac,k8s.io/kubernetes/pkg/apis/rbac/v1alpha1,k8s.io/kubernetes/pkg/apis/storage,k8s.io/kubernetes/pkg/apis/storage/v1beta1]

package v1beta1

import (
	"k8s.io/kubernetes/pkg/api/errors"
	v1 "k8s.io/kubernetes/pkg/api/v1"
	authorization "k8s.io/kubernetes/pkg/apis/authorization"
	v1beta1 "k8s.io/kubernetes/pkg/apis/authorization/v1beta1"
	"k8s.io/kubernetes/pkg/client/cache"
	"k8s.io/kubernetes/pkg/labels"
)

// SelfSubjectAccessReviewLister helps list SelfSubjectAccessReviews.
type SelfSubjectAccessReviewLister interface {
	// List lists all SelfSubjectAccessReviews in the indexer.
	List(selector labels.Selector) (ret []*v1beta1.SelfSubjectAccessReview, err error)
	// Get retrieves the SelfSubjectAccessReview from the index for a given name.
	Get(name string) (*v1beta1.SelfSubjectAccessReview, error)
}

// selfSubjectAccessReviewLister implements the SelfSubjectAccessReviewLister interface.
type selfSubjectAccessReviewLister struct {
	indexer cache.Indexer
}

// NewSelfSubjectAccessReviewLister returns a new SelfSubjectAccessReviewLister.
func NewSelfSubjectAccessReviewLister(indexer cache.Indexer) SelfSubjectAccessReviewLister {
	return &selfSubjectAccessReviewLister{indexer: indexer}
}

// List lists all SelfSubjectAccessReviews in the indexer.
func (s *selfSubjectAccessReviewLister) List(selector labels.Selector) (ret []*v1beta1.SelfSubjectAccessReview, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1beta1.SelfSubjectAccessReview))
	})
	return ret, err
}

// Get retrieves the SelfSubjectAccessReview from the index for a given name.
func (s *selfSubjectAccessReviewLister) Get(name string) (*v1beta1.SelfSubjectAccessReview, error) {
	key := &v1beta1.SelfSubjectAccessReview{ObjectMeta: v1.ObjectMeta{Name: name}}
	obj, exists, err := s.indexer.Get(key)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(authorization.Resource("selfsubjectaccessreview"), name)
	}
	return obj.(*v1beta1.SelfSubjectAccessReview), nil
}
