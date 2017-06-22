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

// This file was automatically generated by lister-gen

package v1

import (
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	v1 "k8s.io/api/core/v1"
	"k8s.io/client-go/tools/cache"
)

// ServiceAccountLister helps list ServiceAccounts.
type ServiceAccountLister interface {
	// List lists all ServiceAccounts in the indexer.
	List(selector labels.Selector) (ret []*v1.ServiceAccount, err error)
	// ServiceAccounts returns an object that can list and get ServiceAccounts.
	ServiceAccounts(namespace string) ServiceAccountNamespaceLister
	ServiceAccountListerExpansion
}

// serviceAccountLister implements the ServiceAccountLister interface.
type serviceAccountLister struct {
	indexer cache.Indexer
}

// NewServiceAccountLister returns a new ServiceAccountLister.
func NewServiceAccountLister(indexer cache.Indexer) ServiceAccountLister {
	return &serviceAccountLister{indexer: indexer}
}

// List lists all ServiceAccounts in the indexer.
func (s *serviceAccountLister) List(selector labels.Selector) (ret []*v1.ServiceAccount, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1.ServiceAccount))
	})
	return ret, err
}

// ServiceAccounts returns an object that can list and get ServiceAccounts.
func (s *serviceAccountLister) ServiceAccounts(namespace string) ServiceAccountNamespaceLister {
	return serviceAccountNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// ServiceAccountNamespaceLister helps list and get ServiceAccounts.
type ServiceAccountNamespaceLister interface {
	// List lists all ServiceAccounts in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1.ServiceAccount, err error)
	// Get retrieves the ServiceAccount from the indexer for a given namespace and name.
	Get(name string) (*v1.ServiceAccount, error)
	ServiceAccountNamespaceListerExpansion
}

// serviceAccountNamespaceLister implements the ServiceAccountNamespaceLister
// interface.
type serviceAccountNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all ServiceAccounts in the indexer for a given namespace.
func (s serviceAccountNamespaceLister) List(selector labels.Selector) (ret []*v1.ServiceAccount, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1.ServiceAccount))
	})
	return ret, err
}

// Get retrieves the ServiceAccount from the indexer for a given namespace and name.
func (s serviceAccountNamespaceLister) Get(name string) (*v1.ServiceAccount, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1.Resource("serviceaccount"), name)
	}
	return obj.(*v1.ServiceAccount), nil
}
