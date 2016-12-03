/*
Copyright 2016 The Kubernetes Authors.

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

// This file was automatically generated by lister-gen with arguments: --input-dirs=[k8s.io/kubernetes/pkg/api,k8s.io/kubernetes/pkg/api/v1,k8s.io/kubernetes/pkg/apis/abac,k8s.io/kubernetes/pkg/apis/abac/v0,k8s.io/kubernetes/pkg/apis/abac/v1beta1,k8s.io/kubernetes/pkg/apis/apps,k8s.io/kubernetes/pkg/apis/apps/v1beta1,k8s.io/kubernetes/pkg/apis/authentication,k8s.io/kubernetes/pkg/apis/authentication/v1beta1,k8s.io/kubernetes/pkg/apis/authorization,k8s.io/kubernetes/pkg/apis/authorization/v1beta1,k8s.io/kubernetes/pkg/apis/autoscaling,k8s.io/kubernetes/pkg/apis/autoscaling/v1,k8s.io/kubernetes/pkg/apis/batch,k8s.io/kubernetes/pkg/apis/batch/v1,k8s.io/kubernetes/pkg/apis/batch/v2alpha1,k8s.io/kubernetes/pkg/apis/certificates,k8s.io/kubernetes/pkg/apis/certificates/v1alpha1,k8s.io/kubernetes/pkg/apis/componentconfig,k8s.io/kubernetes/pkg/apis/componentconfig/v1alpha1,k8s.io/kubernetes/pkg/apis/extensions,k8s.io/kubernetes/pkg/apis/extensions/v1beta1,k8s.io/kubernetes/pkg/apis/imagepolicy,k8s.io/kubernetes/pkg/apis/imagepolicy/v1alpha1,k8s.io/kubernetes/pkg/apis/meta/v1,k8s.io/kubernetes/pkg/apis/policy,k8s.io/kubernetes/pkg/apis/policy/v1alpha1,k8s.io/kubernetes/pkg/apis/policy/v1beta1,k8s.io/kubernetes/pkg/apis/rbac,k8s.io/kubernetes/pkg/apis/rbac/v1alpha1,k8s.io/kubernetes/pkg/apis/storage,k8s.io/kubernetes/pkg/apis/storage/v1beta1]

package internalversion

import (
	"k8s.io/kubernetes/pkg/api/errors"
	batch "k8s.io/kubernetes/pkg/apis/batch"
	"k8s.io/kubernetes/pkg/client/cache"
	"k8s.io/kubernetes/pkg/labels"
)

// CronJobLister helps list CronJobs.
type CronJobLister interface {
	// List lists all CronJobs in the indexer.
	List(selector labels.Selector) (ret []*batch.CronJob, err error)
	// CronJobs returns an object that can list and get CronJobs.
	CronJobs(namespace string) CronJobNamespaceLister
	CronJobListerExpansion
}

// cronJobLister implements the CronJobLister interface.
type cronJobLister struct {
	indexer cache.Indexer
}

// NewCronJobLister returns a new CronJobLister.
func NewCronJobLister(indexer cache.Indexer) CronJobLister {
	return &cronJobLister{indexer: indexer}
}

// List lists all CronJobs in the indexer.
func (s *cronJobLister) List(selector labels.Selector) (ret []*batch.CronJob, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*batch.CronJob))
	})
	return ret, err
}

// CronJobs returns an object that can list and get CronJobs.
func (s *cronJobLister) CronJobs(namespace string) CronJobNamespaceLister {
	return cronJobNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// CronJobNamespaceLister helps list and get CronJobs.
type CronJobNamespaceLister interface {
	// List lists all CronJobs in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*batch.CronJob, err error)
	// Get retrieves the CronJob from the indexer for a given namespace and name.
	Get(name string) (*batch.CronJob, error)
	CronJobNamespaceListerExpansion
}

// cronJobNamespaceLister implements the CronJobNamespaceLister
// interface.
type cronJobNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all CronJobs in the indexer for a given namespace.
func (s cronJobNamespaceLister) List(selector labels.Selector) (ret []*batch.CronJob, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*batch.CronJob))
	})
	return ret, err
}

// Get retrieves the CronJob from the indexer for a given namespace and name.
func (s cronJobNamespaceLister) Get(name string) (*batch.CronJob, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(batch.Resource("cronjob"), name)
	}
	return obj.(*batch.CronJob), nil
}
