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

package v1beta1

import (
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
	v1beta1 "k8s.io/kubernetes/pkg/apis/extensions/v1beta1"
)

// IngressLister helps list Ingresses.
type IngressLister interface {
	// List lists all Ingresses in the indexer.
	List(selector labels.Selector) (ret []*v1beta1.Ingress, err error)
	// Ingresses returns an object that can list and get Ingresses.
	Ingresses(namespace string) IngressNamespaceLister
	IngressListerExpansion
}

// ingressLister implements the IngressLister interface.
type ingressLister struct {
	indexer cache.Indexer
}

// NewIngressLister returns a new IngressLister.
func NewIngressLister(indexer cache.Indexer) IngressLister {
	return &ingressLister{indexer: indexer}
}

// List lists all Ingresses in the indexer.
func (s *ingressLister) List(selector labels.Selector) (ret []*v1beta1.Ingress, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1beta1.Ingress))
	})
	return ret, err
}

// Ingresses returns an object that can list and get Ingresses.
func (s *ingressLister) Ingresses(namespace string) IngressNamespaceLister {
	return ingressNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// IngressNamespaceLister helps list and get Ingresses.
type IngressNamespaceLister interface {
	// List lists all Ingresses in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1beta1.Ingress, err error)
	// Get retrieves the Ingress from the indexer for a given namespace and name.
	Get(name string) (*v1beta1.Ingress, error)
	IngressNamespaceListerExpansion
}

// ingressNamespaceLister implements the IngressNamespaceLister
// interface.
type ingressNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all Ingresses in the indexer for a given namespace.
func (s ingressNamespaceLister) List(selector labels.Selector) (ret []*v1beta1.Ingress, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1beta1.Ingress))
	})
	return ret, err
}

// Get retrieves the Ingress from the indexer for a given namespace and name.
func (s ingressNamespaceLister) Get(name string) (*v1beta1.Ingress, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1beta1.Resource("ingress"), name)
	}
	return obj.(*v1beta1.Ingress), nil
}
