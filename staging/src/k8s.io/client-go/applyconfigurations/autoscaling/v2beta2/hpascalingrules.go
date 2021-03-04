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

package v2beta2

import (
	v2beta2 "k8s.io/api/autoscaling/v2beta2"
)

// HPAScalingRulesApplyConfiguration represents an declarative configuration of the HPAScalingRules type for use
// with apply.
type HPAScalingRulesApplyConfiguration struct {
	StabilizationWindowSeconds *int32                               `json:"stabilizationWindowSeconds,omitempty"`
	SelectPolicy               *v2beta2.ScalingPolicySelect         `json:"selectPolicy,omitempty"`
	Policies                   []HPAScalingPolicyApplyConfiguration `json:"policies,omitempty"`
}

// HPAScalingRulesApplyConfiguration constructs an declarative configuration of the HPAScalingRules type for use with
// apply.
func HPAScalingRules() *HPAScalingRulesApplyConfiguration {
	return &HPAScalingRulesApplyConfiguration{}
}

// WithStabilizationWindowSeconds sets the StabilizationWindowSeconds field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the StabilizationWindowSeconds field is set to the value of the last call.
func (b *HPAScalingRulesApplyConfiguration) WithStabilizationWindowSeconds(value int32) *HPAScalingRulesApplyConfiguration {
	b.StabilizationWindowSeconds = &value
	return b
}

// WithSelectPolicy sets the SelectPolicy field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the SelectPolicy field is set to the value of the last call.
func (b *HPAScalingRulesApplyConfiguration) WithSelectPolicy(value v2beta2.ScalingPolicySelect) *HPAScalingRulesApplyConfiguration {
	b.SelectPolicy = &value
	return b
}

// WithPolicies adds the given value to the Policies field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the Policies field.
func (b *HPAScalingRulesApplyConfiguration) WithPolicies(values ...*HPAScalingPolicyApplyConfiguration) *HPAScalingRulesApplyConfiguration {
	for i := range values {
		if values[i] == nil {
			panic("nil value passed to WithPolicies")
		}
		b.Policies = append(b.Policies, *values[i])
	}
	return b
}
