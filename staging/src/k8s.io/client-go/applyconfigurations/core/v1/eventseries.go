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
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EventSeriesApplyConfiguration represents an declarative configuration of the EventSeries type for use
// with apply.
type EventSeriesApplyConfiguration struct {
	Count            *int32        `json:"count,omitempty"`
	LastObservedTime *v1.MicroTime `json:"lastObservedTime,omitempty"`
}

// EventSeriesApplyConfiguration constructs an declarative configuration of the EventSeries type for use with
// apply.
func EventSeries() *EventSeriesApplyConfiguration {
	return &EventSeriesApplyConfiguration{}
}

// WithCount sets the Count field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Count field is set to the value of the last call.
func (b *EventSeriesApplyConfiguration) WithCount(value int32) *EventSeriesApplyConfiguration {
	b.Count = &value
	return b
}

// WithLastObservedTime sets the LastObservedTime field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the LastObservedTime field is set to the value of the last call.
func (b *EventSeriesApplyConfiguration) WithLastObservedTime(value v1.MicroTime) *EventSeriesApplyConfiguration {
	b.LastObservedTime = &value
	return b
}
