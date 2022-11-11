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
	policyv1 "k8s.io/api/policy/v1"
	intstr "k8s.io/apimachinery/pkg/util/intstr"
	v1 "k8s.io/client-go/applyconfigurations/meta/v1"
)

// PodDisruptionBudgetSpecApplyConfiguration represents an declarative configuration of the PodDisruptionBudgetSpec type for use
// with apply.
type PodDisruptionBudgetSpecApplyConfiguration struct {
	MinAvailable               *intstr.IntOrString                      `json:"minAvailable,omitempty"`
	Selector                   *v1.LabelSelectorApplyConfiguration      `json:"selector,omitempty"`
	MaxUnavailable             *intstr.IntOrString                      `json:"maxUnavailable,omitempty"`
	UnhealthyPodEvictionPolicy *policyv1.UnhealthyPodEvictionPolicyType `json:"unhealthyPodEvictionPolicy,omitempty"`
}

// PodDisruptionBudgetSpecApplyConfiguration constructs an declarative configuration of the PodDisruptionBudgetSpec type for use with
// apply.
func PodDisruptionBudgetSpec() *PodDisruptionBudgetSpecApplyConfiguration {
	return &PodDisruptionBudgetSpecApplyConfiguration{}
}

// WithMinAvailable sets the MinAvailable field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the MinAvailable field is set to the value of the last call.
func (b *PodDisruptionBudgetSpecApplyConfiguration) WithMinAvailable(value intstr.IntOrString) *PodDisruptionBudgetSpecApplyConfiguration {
	b.MinAvailable = &value
	return b
}

// WithSelector sets the Selector field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Selector field is set to the value of the last call.
func (b *PodDisruptionBudgetSpecApplyConfiguration) WithSelector(value *v1.LabelSelectorApplyConfiguration) *PodDisruptionBudgetSpecApplyConfiguration {
	b.Selector = value
	return b
}

// WithMaxUnavailable sets the MaxUnavailable field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the MaxUnavailable field is set to the value of the last call.
func (b *PodDisruptionBudgetSpecApplyConfiguration) WithMaxUnavailable(value intstr.IntOrString) *PodDisruptionBudgetSpecApplyConfiguration {
	b.MaxUnavailable = &value
	return b
}

// WithUnhealthyPodEvictionPolicy sets the UnhealthyPodEvictionPolicy field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the UnhealthyPodEvictionPolicy field is set to the value of the last call.
func (b *PodDisruptionBudgetSpecApplyConfiguration) WithUnhealthyPodEvictionPolicy(value policyv1.UnhealthyPodEvictionPolicyType) *PodDisruptionBudgetSpecApplyConfiguration {
	b.UnhealthyPodEvictionPolicy = &value
	return b
}
