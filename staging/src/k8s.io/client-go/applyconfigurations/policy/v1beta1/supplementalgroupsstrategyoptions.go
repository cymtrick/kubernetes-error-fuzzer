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

package v1beta1

import (
	v1beta1 "k8s.io/api/policy/v1beta1"
)

// SupplementalGroupsStrategyOptionsApplyConfiguration represents an declarative configuration of the SupplementalGroupsStrategyOptions type for use
// with apply.
type SupplementalGroupsStrategyOptionsApplyConfiguration struct {
	Rule   *v1beta1.SupplementalGroupsStrategyType `json:"rule,omitempty"`
	Ranges []IDRangeApplyConfiguration             `json:"ranges,omitempty"`
}

// SupplementalGroupsStrategyOptionsApplyConfiguration constructs an declarative configuration of the SupplementalGroupsStrategyOptions type for use with
// apply.
func SupplementalGroupsStrategyOptions() *SupplementalGroupsStrategyOptionsApplyConfiguration {
	return &SupplementalGroupsStrategyOptionsApplyConfiguration{}
}

// WithRule sets the Rule field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Rule field is set to the value of the last call.
func (b *SupplementalGroupsStrategyOptionsApplyConfiguration) WithRule(value v1beta1.SupplementalGroupsStrategyType) *SupplementalGroupsStrategyOptionsApplyConfiguration {
	b.Rule = &value
	return b
}

// WithRanges adds the given value to the Ranges field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the Ranges field.
func (b *SupplementalGroupsStrategyOptionsApplyConfiguration) WithRanges(values ...*IDRangeApplyConfiguration) *SupplementalGroupsStrategyOptionsApplyConfiguration {
	for i := range values {
		if values[i] == nil {
			panic("nil value passed to WithRanges")
		}
		b.Ranges = append(b.Ranges, *values[i])
	}
	return b
}
