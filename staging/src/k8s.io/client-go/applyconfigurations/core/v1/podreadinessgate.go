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
	v1 "k8s.io/api/core/v1"
)

// PodReadinessGateApplyConfiguration represents an declarative configuration of the PodReadinessGate type for use
// with apply.
type PodReadinessGateApplyConfiguration struct {
	ConditionType *v1.PodConditionType `json:"conditionType,omitempty"`
}

// PodReadinessGateApplyConfiguration constructs an declarative configuration of the PodReadinessGate type for use with
// apply.
func PodReadinessGate() *PodReadinessGateApplyConfiguration {
	return &PodReadinessGateApplyConfiguration{}
}

// WithConditionType sets the ConditionType field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the ConditionType field is set to the value of the last call.
func (b *PodReadinessGateApplyConfiguration) WithConditionType(value v1.PodConditionType) *PodReadinessGateApplyConfiguration {
	b.ConditionType = &value
	return b
}
