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
	settings "k8s.io/kubernetes/pkg/apis/settings"
)

// PodPresetLister helps list PodPresets.
type PodPresetLister interface {
	// List lists all PodPresets in the indexer.
	List(selector labels.Selector) (ret []*settings.PodPreset, err error)
	// PodPresets returns an object that can list and get PodPresets.
	PodPresets(namespace string) PodPresetNamespaceLister
	PodPresetListerExpansion
}

// podPresetLister implements the PodPresetLister interface.
type podPresetLister struct {
	indexer cache.Indexer
}

// NewPodPresetLister returns a new PodPresetLister.
func NewPodPresetLister(indexer cache.Indexer) PodPresetLister {
	return &podPresetLister{indexer: indexer}
}

// List lists all PodPresets in the indexer.
func (s *podPresetLister) List(selector labels.Selector) (ret []*settings.PodPreset, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*settings.PodPreset))
	})
	return ret, err
}

// PodPresets returns an object that can list and get PodPresets.
func (s *podPresetLister) PodPresets(namespace string) PodPresetNamespaceLister {
	return podPresetNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// PodPresetNamespaceLister helps list and get PodPresets.
type PodPresetNamespaceLister interface {
	// List lists all PodPresets in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*settings.PodPreset, err error)
	// Get retrieves the PodPreset from the indexer for a given namespace and name.
	Get(name string) (*settings.PodPreset, error)
	PodPresetNamespaceListerExpansion
}

// podPresetNamespaceLister implements the PodPresetNamespaceLister
// interface.
type podPresetNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all PodPresets in the indexer for a given namespace.
func (s podPresetNamespaceLister) List(selector labels.Selector) (ret []*settings.PodPreset, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*settings.PodPreset))
	})
	return ret, err
}

// Get retrieves the PodPreset from the indexer for a given namespace and name.
func (s podPresetNamespaceLister) Get(name string) (*settings.PodPreset, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(settings.Resource("podpreset"), name)
	}
	return obj.(*settings.PodPreset), nil
}
