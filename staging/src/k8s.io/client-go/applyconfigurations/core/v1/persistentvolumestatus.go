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
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// PersistentVolumeStatusApplyConfiguration represents an declarative configuration of the PersistentVolumeStatus type for use
// with apply.
type PersistentVolumeStatusApplyConfiguration struct {
	Phase                   *v1.PersistentVolumePhase `json:"phase,omitempty"`
	Message                 *string                   `json:"message,omitempty"`
	Reason                  *string                   `json:"reason,omitempty"`
	LastPhaseTransitionTime *metav1.Time              `json:"lastPhaseTransitionTime,omitempty"`
}

// PersistentVolumeStatusApplyConfiguration constructs an declarative configuration of the PersistentVolumeStatus type for use with
// apply.
func PersistentVolumeStatus() *PersistentVolumeStatusApplyConfiguration {
	return &PersistentVolumeStatusApplyConfiguration{}
}

// WithPhase sets the Phase field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Phase field is set to the value of the last call.
func (b *PersistentVolumeStatusApplyConfiguration) WithPhase(value v1.PersistentVolumePhase) *PersistentVolumeStatusApplyConfiguration {
	b.Phase = &value
	return b
}

// WithMessage sets the Message field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Message field is set to the value of the last call.
func (b *PersistentVolumeStatusApplyConfiguration) WithMessage(value string) *PersistentVolumeStatusApplyConfiguration {
	b.Message = &value
	return b
}

// WithReason sets the Reason field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Reason field is set to the value of the last call.
func (b *PersistentVolumeStatusApplyConfiguration) WithReason(value string) *PersistentVolumeStatusApplyConfiguration {
	b.Reason = &value
	return b
}

// WithLastPhaseTransitionTime sets the LastPhaseTransitionTime field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the LastPhaseTransitionTime field is set to the value of the last call.
func (b *PersistentVolumeStatusApplyConfiguration) WithLastPhaseTransitionTime(value metav1.Time) *PersistentVolumeStatusApplyConfiguration {
	b.LastPhaseTransitionTime = &value
	return b
}
