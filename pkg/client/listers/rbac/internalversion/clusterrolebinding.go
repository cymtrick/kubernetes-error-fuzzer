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

// This file was automatically generated by lister-gen with arguments: --input-dirs=[k8s.io/kubernetes/pkg/api,k8s.io/kubernetes/pkg/api/v1,k8s.io/kubernetes/pkg/apis/abac,k8s.io/kubernetes/pkg/apis/abac/v0,k8s.io/kubernetes/pkg/apis/abac/v1beta1,k8s.io/kubernetes/pkg/apis/apps,k8s.io/kubernetes/pkg/apis/apps/v1beta1,k8s.io/kubernetes/pkg/apis/authentication,k8s.io/kubernetes/pkg/apis/authentication/v1beta1,k8s.io/kubernetes/pkg/apis/authorization,k8s.io/kubernetes/pkg/apis/authorization/v1beta1,k8s.io/kubernetes/pkg/apis/autoscaling,k8s.io/kubernetes/pkg/apis/autoscaling/v1,k8s.io/kubernetes/pkg/apis/batch,k8s.io/kubernetes/pkg/apis/batch/v1,k8s.io/kubernetes/pkg/apis/batch/v2alpha1,k8s.io/kubernetes/pkg/apis/certificates,k8s.io/kubernetes/pkg/apis/certificates/v1alpha1,k8s.io/kubernetes/pkg/apis/componentconfig,k8s.io/kubernetes/pkg/apis/componentconfig/v1alpha1,k8s.io/kubernetes/pkg/apis/extensions,k8s.io/kubernetes/pkg/apis/extensions/v1beta1,k8s.io/kubernetes/pkg/apis/imagepolicy,k8s.io/kubernetes/pkg/apis/imagepolicy/v1alpha1,k8s.io/kubernetes/pkg/apis/policy,k8s.io/kubernetes/pkg/apis/policy/v1alpha1,k8s.io/kubernetes/pkg/apis/policy/v1beta1,k8s.io/kubernetes/pkg/apis/rbac,k8s.io/kubernetes/pkg/apis/rbac/v1alpha1,k8s.io/kubernetes/pkg/apis/storage,k8s.io/kubernetes/pkg/apis/storage/v1beta1]

package internalversion

import (
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	api "k8s.io/kubernetes/pkg/api"
	rbac "k8s.io/kubernetes/pkg/apis/rbac"
	"k8s.io/kubernetes/pkg/client/cache"
)

// ClusterRoleBindingLister helps list ClusterRoleBindings.
type ClusterRoleBindingLister interface {
	// List lists all ClusterRoleBindings in the indexer.
	List(selector labels.Selector) (ret []*rbac.ClusterRoleBinding, err error)
	// Get retrieves the ClusterRoleBinding from the index for a given name.
	Get(name string) (*rbac.ClusterRoleBinding, error)
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
func (s *clusterRoleBindingLister) List(selector labels.Selector) (ret []*rbac.ClusterRoleBinding, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*rbac.ClusterRoleBinding))
	})
	return ret, err
}

// Get retrieves the ClusterRoleBinding from the index for a given name.
func (s *clusterRoleBindingLister) Get(name string) (*rbac.ClusterRoleBinding, error) {
	key := &rbac.ClusterRoleBinding{ObjectMeta: api.ObjectMeta{Name: name}}
	obj, exists, err := s.indexer.Get(key)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(rbac.Resource("clusterrolebinding"), name)
	}
	return obj.(*rbac.ClusterRoleBinding), nil
}
