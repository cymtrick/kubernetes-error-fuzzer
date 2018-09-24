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

package v1

import (
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
	v1 "k8s.io/code-generator/_examples/MixedCase/apis/example/v1"
)

// ClusterTestTypeLister helps list ClusterTestTypes.
type ClusterTestTypeLister interface {
	// List lists all ClusterTestTypes in the indexer.
	List(selector labels.Selector) (ret []*v1.ClusterTestType, err error)
	// Get retrieves the ClusterTestType from the index for a given name.
	Get(name string) (*v1.ClusterTestType, error)
	ClusterTestTypeListerExpansion
}

// clusterTestTypeLister implements the ClusterTestTypeLister interface.
type clusterTestTypeLister struct {
	indexer cache.Indexer
}

// NewClusterTestTypeLister returns a new ClusterTestTypeLister.
func NewClusterTestTypeLister(indexer cache.Indexer) ClusterTestTypeLister {
	return &clusterTestTypeLister{indexer: indexer}
}

// List lists all ClusterTestTypes in the indexer.
func (s *clusterTestTypeLister) List(selector labels.Selector) (ret []*v1.ClusterTestType, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1.ClusterTestType))
	})
	return ret, err
}

// Get retrieves the ClusterTestType from the index for a given name.
func (s *clusterTestTypeLister) Get(name string) (*v1.ClusterTestType, error) {
	obj, exists, err := s.indexer.GetByKey(name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1.Resource("clustertesttype"), name)
	}
	return obj.(*v1.ClusterTestType), nil
}
