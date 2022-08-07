/*
Copyright The Kubernetes Authors.

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

// Code generated by lister-gen. DO NOT EDIT.

package v1alpha1

import (
	v1alpha1 "k8s.io/api/networking/v1alpha1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// ClusterCIDRLister helps list ClusterCIDRs.
// All objects returned here must be treated as read-only.
type ClusterCIDRLister interface {
	// List lists all ClusterCIDRs in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.ClusterCIDR, err error)
	// Get retrieves the ClusterCIDR from the index for a given name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.ClusterCIDR, error)
	ClusterCIDRListerExpansion
}

// clusterCIDRLister implements the ClusterCIDRLister interface.
type clusterCIDRLister struct {
	indexer cache.Indexer
}

// NewClusterCIDRLister returns a new ClusterCIDRLister.
func NewClusterCIDRLister(indexer cache.Indexer) ClusterCIDRLister {
	return &clusterCIDRLister{indexer: indexer}
}

// List lists all ClusterCIDRs in the indexer.
func (s *clusterCIDRLister) List(selector labels.Selector) (ret []*v1alpha1.ClusterCIDR, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.ClusterCIDR))
	})
	return ret, err
}

// Get retrieves the ClusterCIDR from the index for a given name.
func (s *clusterCIDRLister) Get(name string) (*v1alpha1.ClusterCIDR, error) {
	obj, exists, err := s.indexer.GetByKey(name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("clustercidr"), name)
	}
	return obj.(*v1alpha1.ClusterCIDR), nil
}
