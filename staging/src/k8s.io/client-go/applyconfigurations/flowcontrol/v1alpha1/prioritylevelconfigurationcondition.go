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

package v1alpha1

import (
	v1alpha1 "k8s.io/api/flowcontrol/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// PriorityLevelConfigurationConditionApplyConfiguration represents an declarative configuration of the PriorityLevelConfigurationCondition type for use
// with apply.
type PriorityLevelConfigurationConditionApplyConfiguration struct {
	Type               *v1alpha1.PriorityLevelConfigurationConditionType `json:"type,omitempty"`
	Status             *v1alpha1.ConditionStatus                         `json:"status,omitempty"`
	LastTransitionTime *v1.Time                                          `json:"lastTransitionTime,omitempty"`
	Reason             *string                                           `json:"reason,omitempty"`
	Message            *string                                           `json:"message,omitempty"`
}

// PriorityLevelConfigurationConditionApplyConfiguration constructs an declarative configuration of the PriorityLevelConfigurationCondition type for use with
// apply.
func PriorityLevelConfigurationCondition() *PriorityLevelConfigurationConditionApplyConfiguration {
	return &PriorityLevelConfigurationConditionApplyConfiguration{}
}

// WithType sets the Type field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Type field is set to the value of the last call.
func (b *PriorityLevelConfigurationConditionApplyConfiguration) WithType(value v1alpha1.PriorityLevelConfigurationConditionType) *PriorityLevelConfigurationConditionApplyConfiguration {
	b.Type = &value
	return b
}

// WithStatus sets the Status field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Status field is set to the value of the last call.
func (b *PriorityLevelConfigurationConditionApplyConfiguration) WithStatus(value v1alpha1.ConditionStatus) *PriorityLevelConfigurationConditionApplyConfiguration {
	b.Status = &value
	return b
}

// WithLastTransitionTime sets the LastTransitionTime field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the LastTransitionTime field is set to the value of the last call.
func (b *PriorityLevelConfigurationConditionApplyConfiguration) WithLastTransitionTime(value v1.Time) *PriorityLevelConfigurationConditionApplyConfiguration {
	b.LastTransitionTime = &value
	return b
}

// WithReason sets the Reason field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Reason field is set to the value of the last call.
func (b *PriorityLevelConfigurationConditionApplyConfiguration) WithReason(value string) *PriorityLevelConfigurationConditionApplyConfiguration {
	b.Reason = &value
	return b
}

// WithMessage sets the Message field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Message field is set to the value of the last call.
func (b *PriorityLevelConfigurationConditionApplyConfiguration) WithMessage(value string) *PriorityLevelConfigurationConditionApplyConfiguration {
	b.Message = &value
	return b
}
