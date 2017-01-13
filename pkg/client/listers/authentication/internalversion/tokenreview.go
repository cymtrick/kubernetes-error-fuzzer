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
	authentication "k8s.io/kubernetes/pkg/apis/authentication"
	"k8s.io/kubernetes/pkg/client/cache"
)

// TokenReviewLister helps list TokenReviews.
type TokenReviewLister interface {
	// List lists all TokenReviews in the indexer.
	List(selector labels.Selector) (ret []*authentication.TokenReview, err error)
	// Get retrieves the TokenReview from the index for a given name.
	Get(name string) (*authentication.TokenReview, error)
	TokenReviewListerExpansion
}

// tokenReviewLister implements the TokenReviewLister interface.
type tokenReviewLister struct {
	indexer cache.Indexer
}

// NewTokenReviewLister returns a new TokenReviewLister.
func NewTokenReviewLister(indexer cache.Indexer) TokenReviewLister {
	return &tokenReviewLister{indexer: indexer}
}

// List lists all TokenReviews in the indexer.
func (s *tokenReviewLister) List(selector labels.Selector) (ret []*authentication.TokenReview, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*authentication.TokenReview))
	})
	return ret, err
}

// Get retrieves the TokenReview from the index for a given name.
func (s *tokenReviewLister) Get(name string) (*authentication.TokenReview, error) {
	key := &authentication.TokenReview{ObjectMeta: api.ObjectMeta{Name: name}}
	obj, exists, err := s.indexer.Get(key)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(authentication.Resource("tokenreview"), name)
	}
	return obj.(*authentication.TokenReview), nil
}
