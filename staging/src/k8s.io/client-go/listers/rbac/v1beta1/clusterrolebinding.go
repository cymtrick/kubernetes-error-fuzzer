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

package v1beta1

import (
	v1beta1 "k8s.io/api/rbac/v1beta1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// ClusterRoleBindingLister helps list ClusterRoleBindings.
type ClusterRoleBindingLister interface {
	// List lists all ClusterRoleBindings in the indexer.
	List(selector labels.Selector) (ret []*v1beta1.ClusterRoleBinding, err error)
	// Get retrieves the ClusterRoleBinding from the index for a given name.
	Get(name string) (*v1beta1.ClusterRoleBinding, error)
	ClusterRoleBindingListerExpansion
}

// clusterRoleBindingLister implements the ClusterRoleBindingLister interface.
type clusterRoleBindingLister struct {
	indexer cache.Indexer
}

// NewClusterRoleBindingLister returns a new ClusterRoleBindingLister.
func NewClusterRoleBindingLister(indexer cache.Indexer) ClusterRoleBindingLister {
	return &clusterRoleBindingLister{indexer: indexer}
}

// List lists all ClusterRoleBindings in the indexer.
func (s *clusterRoleBindingLister) List(selector labels.Selector) (ret []*v1beta1.ClusterRoleBinding, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1beta1.ClusterRoleBinding))
	})
	return ret, err
}

// Get retrieves the ClusterRoleBinding from the index for a given name.
func (s *clusterRoleBindingLister) Get(name string) (*v1beta1.ClusterRoleBinding, error) {
	obj, exists, err := s.indexer.GetByKey(name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1beta1.Resource("clusterrolebinding"), name)
	}
	return obj.(*v1beta1.ClusterRoleBinding), nil
}
