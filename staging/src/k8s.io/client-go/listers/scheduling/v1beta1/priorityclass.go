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

package v1beta1

import (
	v1beta1 "k8s.io/api/scheduling/v1beta1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// PriorityClassLister helps list PriorityClasses.
// All objects returned here must be treated as read-only.
type PriorityClassLister interface {
	// List lists all PriorityClasses in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1beta1.PriorityClass, err error)
	// Get retrieves the PriorityClass from the index for a given name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1beta1.PriorityClass, error)
	PriorityClassListerExpansion
}

// priorityClassLister implements the PriorityClassLister interface.
type priorityClassLister struct {
	indexer cache.Indexer
}

// NewPriorityClassLister returns a new PriorityClassLister.
func NewPriorityClassLister(indexer cache.Indexer) PriorityClassLister {
	return &priorityClassLister{indexer: indexer}
}

// List lists all PriorityClasses in the indexer.
func (s *priorityClassLister) List(selector labels.Selector) (ret []*v1beta1.PriorityClass, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1beta1.PriorityClass))
	})
	return ret, err
}

// Get retrieves the PriorityClass from the index for a given name.
func (s *priorityClassLister) Get(name string) (*v1beta1.PriorityClass, error) {
	obj, exists, err := s.indexer.GetByKey(name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1beta1.Resource("priorityclass"), name)
	}
	return obj.(*v1beta1.PriorityClass), nil
}
