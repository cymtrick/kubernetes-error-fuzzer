/*
Copyright 2014 The Kubernetes Authors.

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

package endpoint

import (
	"context"

	metainternalversion "k8s.io/apimachinery/pkg/apis/meta/internalversion"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/watch"
	"k8s.io/apiserver/pkg/registry/rest"
	api "k8s.io/kubernetes/pkg/apis/core"
)

// Registry is an interface for things that know how to store endpoints.
type Registry interface {
	ListEndpoints(ctx context.Context, options *metainternalversion.ListOptions) (*api.EndpointsList, error)
	GetEndpoints(ctx context.Context, name string, options *metav1.GetOptions) (*api.Endpoints, error)
	WatchEndpoints(ctx context.Context, options *metainternalversion.ListOptions) (watch.Interface, error)
	UpdateEndpoints(ctx context.Context, e *api.Endpoints, createValidation rest.ValidateObjectFunc, updateValidation rest.ValidateObjectUpdateFunc, options *metav1.UpdateOptions) error
	DeleteEndpoints(ctx context.Context, name string) error
}

// storage puts strong typing around storage calls
type storage struct {
	rest.StandardStorage
}

// NewRegistry returns a new Registry interface for the given Storage. Any mismatched
// types will panic.
func NewRegistry(s rest.StandardStorage) Registry {
	return &storage{s}
}

func (s *storage) ListEndpoints(ctx context.Context, options *metainternalversion.ListOptions) (*api.EndpointsList, error) {
	obj, err := s.List(ctx, options)
	if err != nil {
		return nil, err
	}
	return obj.(*api.EndpointsList), nil
}

func (s *storage) WatchEndpoints(ctx context.Context, options *metainternalversion.ListOptions) (watch.Interface, error) {
	return s.Watch(ctx, options)
}

func (s *storage) GetEndpoints(ctx context.Context, name string, options *metav1.GetOptions) (*api.Endpoints, error) {
	obj, err := s.Get(ctx, name, options)
	if err != nil {
		return nil, err
	}
	return obj.(*api.Endpoints), nil
}

func (s *storage) UpdateEndpoints(ctx context.Context, endpoints *api.Endpoints, createValidation rest.ValidateObjectFunc, updateValidation rest.ValidateObjectUpdateFunc, options *metav1.UpdateOptions) error {
	_, _, err := s.Update(ctx, endpoints.Name, rest.DefaultUpdatedObjectInfo(endpoints), createValidation, updateValidation, false, options)
	return err
}

func (s *storage) DeleteEndpoints(ctx context.Context, name string) error {
	_, _, err := s.Delete(ctx, name, nil)
	return err
}
