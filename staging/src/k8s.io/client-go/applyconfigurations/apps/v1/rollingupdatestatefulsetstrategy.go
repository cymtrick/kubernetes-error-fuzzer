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
	intstr "k8s.io/apimachinery/pkg/util/intstr"
)

// RollingUpdateStatefulSetStrategyApplyConfiguration represents an declarative configuration of the RollingUpdateStatefulSetStrategy type for use
// with apply.
type RollingUpdateStatefulSetStrategyApplyConfiguration struct {
	Partition      *int32              `json:"partition,omitempty"`
	MaxUnavailable *intstr.IntOrString `json:"maxUnavailable,omitempty"`
}

// RollingUpdateStatefulSetStrategyApplyConfiguration constructs an declarative configuration of the RollingUpdateStatefulSetStrategy type for use with
// apply.
func RollingUpdateStatefulSetStrategy() *RollingUpdateStatefulSetStrategyApplyConfiguration {
	return &RollingUpdateStatefulSetStrategyApplyConfiguration{}
}

// WithPartition sets the Partition field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Partition field is set to the value of the last call.
func (b *RollingUpdateStatefulSetStrategyApplyConfiguration) WithPartition(value int32) *RollingUpdateStatefulSetStrategyApplyConfiguration {
	b.Partition = &value
	return b
}

// WithMaxUnavailable sets the MaxUnavailable field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the MaxUnavailable field is set to the value of the last call.
func (b *RollingUpdateStatefulSetStrategyApplyConfiguration) WithMaxUnavailable(value intstr.IntOrString) *RollingUpdateStatefulSetStrategyApplyConfiguration {
	b.MaxUnavailable = &value
	return b
}
