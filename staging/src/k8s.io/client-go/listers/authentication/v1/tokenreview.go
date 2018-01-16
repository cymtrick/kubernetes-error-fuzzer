/*
Copyright 2018 The Kubernetes Authors.

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

// This file was automatically generated by lister-gen

package v1

import (
	v1 "k8s.io/api/authentication/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// TokenReviewLister helps list TokenReviews.
type TokenReviewLister interface {
	// List lists all TokenReviews in the indexer.
	List(selector labels.Selector) (ret []*v1.TokenReview, err error)
	// Get retrieves the TokenReview from the index for a given name.
	Get(name string) (*v1.TokenReview, error)
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
func (s *tokenReviewLister) List(selector labels.Selector) (ret []*v1.TokenReview, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1.TokenReview))
	})
	return ret, err
}

// Get retrieves the TokenReview from the index for a given name.
func (s *tokenReviewLister) Get(name string) (*v1.TokenReview, error) {
	obj, exists, err := s.indexer.GetByKey(name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1.Resource("tokenreview"), name)
	}
	return obj.(*v1.TokenReview), nil
}
