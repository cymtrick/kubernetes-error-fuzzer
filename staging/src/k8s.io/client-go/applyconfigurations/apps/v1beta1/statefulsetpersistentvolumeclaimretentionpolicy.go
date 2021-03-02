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
	v1beta1 "k8s.io/api/apps/v1beta1"
)

// StatefulSetPersistentVolumeClaimRetentionPolicyApplyConfiguration represents an declarative configuration of the StatefulSetPersistentVolumeClaimRetentionPolicy type for use
// with apply.
type StatefulSetPersistentVolumeClaimRetentionPolicyApplyConfiguration struct {
	WhenDeleted *v1beta1.PersistentVolumeClaimRetentionPolicyType `json:"whenDeleted,omitempty"`
	WhenScaled  *v1beta1.PersistentVolumeClaimRetentionPolicyType `json:"whenScaled,omitempty"`
}

// StatefulSetPersistentVolumeClaimRetentionPolicyApplyConfiguration constructs an declarative configuration of the StatefulSetPersistentVolumeClaimRetentionPolicy type for use with
// apply.
func StatefulSetPersistentVolumeClaimRetentionPolicy() *StatefulSetPersistentVolumeClaimRetentionPolicyApplyConfiguration {
	return &StatefulSetPersistentVolumeClaimRetentionPolicyApplyConfiguration{}
}

// WithWhenDeleted sets the WhenDeleted field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the WhenDeleted field is set to the value of the last call.
func (b *StatefulSetPersistentVolumeClaimRetentionPolicyApplyConfiguration) WithWhenDeleted(value v1beta1.PersistentVolumeClaimRetentionPolicyType) *StatefulSetPersistentVolumeClaimRetentionPolicyApplyConfiguration {
	b.WhenDeleted = &value
	return b
}

// WithWhenScaled sets the WhenScaled field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the WhenScaled field is set to the value of the last call.
func (b *StatefulSetPersistentVolumeClaimRetentionPolicyApplyConfiguration) WithWhenScaled(value v1beta1.PersistentVolumeClaimRetentionPolicyType) *StatefulSetPersistentVolumeClaimRetentionPolicyApplyConfiguration {
	b.WhenScaled = &value
	return b
}
