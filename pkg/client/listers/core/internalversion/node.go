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

package internalversion

import (
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
	core "k8s.io/kubernetes/pkg/apis/core"
)

// NodeLister helps list Nodes.
type NodeLister interface {
	// List lists all Nodes in the indexer.
	List(selector labels.Selector) (ret []*core.Node, err error)
	// Get retrieves the Node from the index for a given name.
	Get(name string) (*core.Node, error)
	NodeListerExpansion
}

// nodeLister implements the NodeLister interface.
type nodeLister struct {
	indexer cache.Indexer
}

// NewNodeLister returns a new NodeLister.
func NewNodeLister(indexer cache.Indexer) NodeLister {
	return &nodeLister{indexer: indexer}
}

// List lists all Nodes in the indexer.
func (s *nodeLister) List(selector labels.Selector) (ret []*core.Node, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*core.Node))
	})
	return ret, err
}

// Get retrieves the Node from the index for a given name.
func (s *nodeLister) Get(name string) (*core.Node, error) {
	obj, exists, err := s.indexer.GetByKey(name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(core.Resource("node"), name)
	}
	return obj.(*core.Node), nil
}
