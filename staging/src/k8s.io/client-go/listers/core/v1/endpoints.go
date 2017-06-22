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

// EndpointsLister helps list Endpoints.
type EndpointsLister interface {
	// List lists all Endpoints in the indexer.
	List(selector labels.Selector) (ret []*v1.Endpoints, err error)
	// Endpoints returns an object that can list and get Endpoints.
	Endpoints(namespace string) EndpointsNamespaceLister
	EndpointsListerExpansion
}

// endpointsLister implements the EndpointsLister interface.
type endpointsLister struct {
	indexer cache.Indexer
}

// NewEndpointsLister returns a new EndpointsLister.
func NewEndpointsLister(indexer cache.Indexer) EndpointsLister {
	return &endpointsLister{indexer: indexer}
}

// List lists all Endpoints in the indexer.
func (s *endpointsLister) List(selector labels.Selector) (ret []*v1.Endpoints, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1.Endpoints))
	})
	return ret, err
}

// Endpoints returns an object that can list and get Endpoints.
func (s *endpointsLister) Endpoints(namespace string) EndpointsNamespaceLister {
	return endpointsNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// EndpointsNamespaceLister helps list and get Endpoints.
type EndpointsNamespaceLister interface {
	// List lists all Endpoints in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1.Endpoints, err error)
	// Get retrieves the Endpoints from the indexer for a given namespace and name.
	Get(name string) (*v1.Endpoints, error)
	EndpointsNamespaceListerExpansion
}

// endpointsNamespaceLister implements the EndpointsNamespaceLister
// interface.
type endpointsNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all Endpoints in the indexer for a given namespace.
func (s endpointsNamespaceLister) List(selector labels.Selector) (ret []*v1.Endpoints, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1.Endpoints))
	})
	return ret, err
}

// Get retrieves the Endpoints from the indexer for a given namespace and name.
func (s endpointsNamespaceLister) Get(name string) (*v1.Endpoints, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1.Resource("endpoints"), name)
	}
	return obj.(*v1.Endpoints), nil
}
