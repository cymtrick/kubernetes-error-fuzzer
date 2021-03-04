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

// Code generated by applyconfiguration-gen. DO NOT EDIT.

package v1

import (
	resource "k8s.io/apimachinery/pkg/api/resource"
)

// ResourceFieldSelectorApplyConfiguration represents an declarative configuration of the ResourceFieldSelector type for use
// with apply.
type ResourceFieldSelectorApplyConfiguration struct {
	ContainerName *string            `json:"containerName,omitempty"`
	Resource      *string            `json:"resource,omitempty"`
	Divisor       *resource.Quantity `json:"divisor,omitempty"`
}

// ResourceFieldSelectorApplyConfiguration constructs an declarative configuration of the ResourceFieldSelector type for use with
// apply.
func ResourceFieldSelector() *ResourceFieldSelectorApplyConfiguration {
	return &ResourceFieldSelectorApplyConfiguration{}
}

// WithContainerName sets the ContainerName field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the ContainerName field is set to the value of the last call.
func (b *ResourceFieldSelectorApplyConfiguration) WithContainerName(value string) *ResourceFieldSelectorApplyConfiguration {
	b.ContainerName = &value
	return b
}

// WithResource sets the Resource field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Resource field is set to the value of the last call.
func (b *ResourceFieldSelectorApplyConfiguration) WithResource(value string) *ResourceFieldSelectorApplyConfiguration {
	b.Resource = &value
	return b
}

// WithDivisor sets the Divisor field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Divisor field is set to the value of the last call.
func (b *ResourceFieldSelectorApplyConfiguration) WithDivisor(value resource.Quantity) *ResourceFieldSelectorApplyConfiguration {
	b.Divisor = &value
	return b
}
