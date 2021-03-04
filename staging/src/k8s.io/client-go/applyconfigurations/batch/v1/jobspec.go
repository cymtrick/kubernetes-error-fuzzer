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
	batchv1 "k8s.io/api/batch/v1"
	corev1 "k8s.io/client-go/applyconfigurations/core/v1"
	v1 "k8s.io/client-go/applyconfigurations/meta/v1"
)

// JobSpecApplyConfiguration represents an declarative configuration of the JobSpec type for use
// with apply.
type JobSpecApplyConfiguration struct {
	Parallelism             *int32                                    `json:"parallelism,omitempty"`
	Completions             *int32                                    `json:"completions,omitempty"`
	ActiveDeadlineSeconds   *int64                                    `json:"activeDeadlineSeconds,omitempty"`
	BackoffLimit            *int32                                    `json:"backoffLimit,omitempty"`
	Selector                *v1.LabelSelectorApplyConfiguration       `json:"selector,omitempty"`
	ManualSelector          *bool                                     `json:"manualSelector,omitempty"`
	Template                *corev1.PodTemplateSpecApplyConfiguration `json:"template,omitempty"`
	TTLSecondsAfterFinished *int32                                    `json:"ttlSecondsAfterFinished,omitempty"`
	CompletionMode          *batchv1.CompletionMode                   `json:"completionMode,omitempty"`
}

// JobSpecApplyConfiguration constructs an declarative configuration of the JobSpec type for use with
// apply.
func JobSpec() *JobSpecApplyConfiguration {
	return &JobSpecApplyConfiguration{}
}

// WithParallelism sets the Parallelism field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Parallelism field is set to the value of the last call.
func (b *JobSpecApplyConfiguration) WithParallelism(value int32) *JobSpecApplyConfiguration {
	b.Parallelism = &value
	return b
}

// WithCompletions sets the Completions field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Completions field is set to the value of the last call.
func (b *JobSpecApplyConfiguration) WithCompletions(value int32) *JobSpecApplyConfiguration {
	b.Completions = &value
	return b
}

// WithActiveDeadlineSeconds sets the ActiveDeadlineSeconds field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the ActiveDeadlineSeconds field is set to the value of the last call.
func (b *JobSpecApplyConfiguration) WithActiveDeadlineSeconds(value int64) *JobSpecApplyConfiguration {
	b.ActiveDeadlineSeconds = &value
	return b
}

// WithBackoffLimit sets the BackoffLimit field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the BackoffLimit field is set to the value of the last call.
func (b *JobSpecApplyConfiguration) WithBackoffLimit(value int32) *JobSpecApplyConfiguration {
	b.BackoffLimit = &value
	return b
}

// WithSelector sets the Selector field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Selector field is set to the value of the last call.
func (b *JobSpecApplyConfiguration) WithSelector(value *v1.LabelSelectorApplyConfiguration) *JobSpecApplyConfiguration {
	b.Selector = value
	return b
}

// WithManualSelector sets the ManualSelector field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the ManualSelector field is set to the value of the last call.
func (b *JobSpecApplyConfiguration) WithManualSelector(value bool) *JobSpecApplyConfiguration {
	b.ManualSelector = &value
	return b
}

// WithTemplate sets the Template field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Template field is set to the value of the last call.
func (b *JobSpecApplyConfiguration) WithTemplate(value *corev1.PodTemplateSpecApplyConfiguration) *JobSpecApplyConfiguration {
	b.Template = value
	return b
}

// WithTTLSecondsAfterFinished sets the TTLSecondsAfterFinished field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the TTLSecondsAfterFinished field is set to the value of the last call.
func (b *JobSpecApplyConfiguration) WithTTLSecondsAfterFinished(value int32) *JobSpecApplyConfiguration {
	b.TTLSecondsAfterFinished = &value
	return b
}

// WithCompletionMode sets the CompletionMode field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the CompletionMode field is set to the value of the last call.
func (b *JobSpecApplyConfiguration) WithCompletionMode(value batchv1.CompletionMode) *JobSpecApplyConfiguration {
	b.CompletionMode = &value
	return b
}
