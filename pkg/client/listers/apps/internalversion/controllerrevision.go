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

// Code generated by lister-gen. DO NOT EDIT.

package internalversion

import (
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
	apps "k8s.io/kubernetes/pkg/apis/apps"
)

// ControllerRevisionLister helps list ControllerRevisions.
type ControllerRevisionLister interface {
	// List lists all ControllerRevisions in the indexer.
	List(selector labels.Selector) (ret []*apps.ControllerRevision, err error)
	// ControllerRevisions returns an object that can list and get ControllerRevisions.
	ControllerRevisions(namespace string) ControllerRevisionNamespaceLister
	ControllerRevisionListerExpansion
}

// controllerRevisionLister implements the ControllerRevisionLister interface.
type controllerRevisionLister struct {
	indexer cache.Indexer
}

// NewControllerRevisionLister returns a new ControllerRevisionLister.
func NewControllerRevisionLister(indexer cache.Indexer) ControllerRevisionLister {
	return &controllerRevisionLister{indexer: indexer}
}

// List lists all ControllerRevisions in the indexer.
func (s *controllerRevisionLister) List(selector labels.Selector) (ret []*apps.ControllerRevision, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*apps.ControllerRevision))
	})
	return ret, err
}

// ControllerRevisions returns an object that can list and get ControllerRevisions.
func (s *controllerRevisionLister) ControllerRevisions(namespace string) ControllerRevisionNamespaceLister {
	return controllerRevisionNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// ControllerRevisionNamespaceLister helps list and get ControllerRevisions.
type ControllerRevisionNamespaceLister interface {
	// List lists all ControllerRevisions in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*apps.ControllerRevision, err error)
	// Get retrieves the ControllerRevision from the indexer for a given namespace and name.
	Get(name string) (*apps.ControllerRevision, error)
	ControllerRevisionNamespaceListerExpansion
}

// controllerRevisionNamespaceLister implements the ControllerRevisionNamespaceLister
// interface.
type controllerRevisionNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all ControllerRevisions in the indexer for a given namespace.
func (s controllerRevisionNamespaceLister) List(selector labels.Selector) (ret []*apps.ControllerRevision, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*apps.ControllerRevision))
	})
	return ret, err
}

// Get retrieves the ControllerRevision from the indexer for a given namespace and name.
func (s controllerRevisionNamespaceLister) Get(name string) (*apps.ControllerRevision, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(apps.Resource("controllerrevision"), name)
	}
	return obj.(*apps.ControllerRevision), nil
}
