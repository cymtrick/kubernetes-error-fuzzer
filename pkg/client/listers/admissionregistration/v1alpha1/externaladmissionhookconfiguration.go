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

package v1alpha1

import (
	"k8s.io/apimachinery/pkg/api/errors"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
	v1alpha1 "k8s.io/api/admissionregistration/v1alpha1"
)

// ExternalAdmissionHookConfigurationLister helps list ExternalAdmissionHookConfigurations.
type ExternalAdmissionHookConfigurationLister interface {
	// List lists all ExternalAdmissionHookConfigurations in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.ExternalAdmissionHookConfiguration, err error)
	// Get retrieves the ExternalAdmissionHookConfiguration from the index for a given name.
	Get(name string) (*v1alpha1.ExternalAdmissionHookConfiguration, error)
	ExternalAdmissionHookConfigurationListerExpansion
}

// externalAdmissionHookConfigurationLister implements the ExternalAdmissionHookConfigurationLister interface.
type externalAdmissionHookConfigurationLister struct {
	indexer cache.Indexer
}

// NewExternalAdmissionHookConfigurationLister returns a new ExternalAdmissionHookConfigurationLister.
func NewExternalAdmissionHookConfigurationLister(indexer cache.Indexer) ExternalAdmissionHookConfigurationLister {
	return &externalAdmissionHookConfigurationLister{indexer: indexer}
}

// List lists all ExternalAdmissionHookConfigurations in the indexer.
func (s *externalAdmissionHookConfigurationLister) List(selector labels.Selector) (ret []*v1alpha1.ExternalAdmissionHookConfiguration, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.ExternalAdmissionHookConfiguration))
	})
	return ret, err
}

// Get retrieves the ExternalAdmissionHookConfiguration from the index for a given name.
func (s *externalAdmissionHookConfigurationLister) Get(name string) (*v1alpha1.ExternalAdmissionHookConfiguration, error) {
	key := &v1alpha1.ExternalAdmissionHookConfiguration{ObjectMeta: v1.ObjectMeta{Name: name}}
	obj, exists, err := s.indexer.Get(key)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("externaladmissionhookconfiguration"), name)
	}
	return obj.(*v1alpha1.ExternalAdmissionHookConfiguration), nil
}
