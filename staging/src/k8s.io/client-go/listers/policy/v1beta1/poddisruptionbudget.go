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
	v1beta1 "k8s.io/api/policy/v1beta1"
	"k8s.io/client-go/tools/cache"
)

// PodDisruptionBudgetLister helps list PodDisruptionBudgets.
type PodDisruptionBudgetLister interface {
	// List lists all PodDisruptionBudgets in the indexer.
	List(selector labels.Selector) (ret []*v1beta1.PodDisruptionBudget, err error)
	// PodDisruptionBudgets returns an object that can list and get PodDisruptionBudgets.
	PodDisruptionBudgets(namespace string) PodDisruptionBudgetNamespaceLister
	PodDisruptionBudgetListerExpansion
}

// podDisruptionBudgetLister implements the PodDisruptionBudgetLister interface.
type podDisruptionBudgetLister struct {
	indexer cache.Indexer
}

// NewPodDisruptionBudgetLister returns a new PodDisruptionBudgetLister.
func NewPodDisruptionBudgetLister(indexer cache.Indexer) PodDisruptionBudgetLister {
	return &podDisruptionBudgetLister{indexer: indexer}
}

// List lists all PodDisruptionBudgets in the indexer.
func (s *podDisruptionBudgetLister) List(selector labels.Selector) (ret []*v1beta1.PodDisruptionBudget, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1beta1.PodDisruptionBudget))
	})
	return ret, err
}

// PodDisruptionBudgets returns an object that can list and get PodDisruptionBudgets.
func (s *podDisruptionBudgetLister) PodDisruptionBudgets(namespace string) PodDisruptionBudgetNamespaceLister {
	return podDisruptionBudgetNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// PodDisruptionBudgetNamespaceLister helps list and get PodDisruptionBudgets.
type PodDisruptionBudgetNamespaceLister interface {
	// List lists all PodDisruptionBudgets in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1beta1.PodDisruptionBudget, err error)
	// Get retrieves the PodDisruptionBudget from the indexer for a given namespace and name.
	Get(name string) (*v1beta1.PodDisruptionBudget, error)
	PodDisruptionBudgetNamespaceListerExpansion
}

// podDisruptionBudgetNamespaceLister implements the PodDisruptionBudgetNamespaceLister
// interface.
type podDisruptionBudgetNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all PodDisruptionBudgets in the indexer for a given namespace.
func (s podDisruptionBudgetNamespaceLister) List(selector labels.Selector) (ret []*v1beta1.PodDisruptionBudget, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1beta1.PodDisruptionBudget))
	})
	return ret, err
}

// Get retrieves the PodDisruptionBudget from the indexer for a given namespace and name.
func (s podDisruptionBudgetNamespaceLister) Get(name string) (*v1beta1.PodDisruptionBudget, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1beta1.Resource("poddisruptionbudget"), name)
	}
	return obj.(*v1beta1.PodDisruptionBudget), nil
}
