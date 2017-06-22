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
	v1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// ServiceLister helps list Services.
type ServiceLister interface {
	// List lists all Services in the indexer.
	List(selector labels.Selector) (ret []*v1.Service, err error)
	// Services returns an object that can list and get Services.
	Services(namespace string) ServiceNamespaceLister
	ServiceListerExpansion
}

// serviceLister implements the ServiceLister interface.
type serviceLister struct {
	indexer cache.Indexer
}

// NewServiceLister returns a new ServiceLister.
func NewServiceLister(indexer cache.Indexer) ServiceLister {
	return &serviceLister{indexer: indexer}
}

// List lists all Services in the indexer.
func (s *serviceLister) List(selector labels.Selector) (ret []*v1.Service, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1.Service))
	})
	return ret, err
}

// Services returns an object that can list and get Services.
func (s *serviceLister) Services(namespace string) ServiceNamespaceLister {
	return serviceNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// ServiceNamespaceLister helps list and get Services.
type ServiceNamespaceLister interface {
	// List lists all Services in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1.Service, err error)
	// Get retrieves the Service from the indexer for a given namespace and name.
	Get(name string) (*v1.Service, error)
	ServiceNamespaceListerExpansion
}

// serviceNamespaceLister implements the ServiceNamespaceLister
// interface.
type serviceNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all Services in the indexer for a given namespace.
func (s serviceNamespaceLister) List(selector labels.Selector) (ret []*v1.Service, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1.Service))
	})
	return ret, err
}

// Get retrieves the Service from the indexer for a given namespace and name.
func (s serviceNamespaceLister) Get(name string) (*v1.Service, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1.Resource("service"), name)
	}
	return obj.(*v1.Service), nil
}
