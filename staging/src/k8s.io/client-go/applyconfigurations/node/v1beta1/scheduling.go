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
	v1 "k8s.io/client-go/applyconfigurations/core/v1"
)

// SchedulingApplyConfiguration represents an declarative configuration of the Scheduling type for use
// with apply.
type SchedulingApplyConfiguration struct {
	NodeSelector map[string]string                 `json:"nodeSelector,omitempty"`
	Tolerations  []v1.TolerationApplyConfiguration `json:"tolerations,omitempty"`
}

// SchedulingApplyConfiguration constructs an declarative configuration of the Scheduling type for use with
// apply.
func Scheduling() *SchedulingApplyConfiguration {
	return &SchedulingApplyConfiguration{}
}

// WithNodeSelector puts the entries into the NodeSelector field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, the entries provided by each call will be put on the NodeSelector field,
// overwriting an existing map entries in NodeSelector field with the same key.
func (b *SchedulingApplyConfiguration) WithNodeSelector(entries map[string]string) *SchedulingApplyConfiguration {
	if b.NodeSelector == nil && len(entries) > 0 {
		b.NodeSelector = make(map[string]string, len(entries))
	}
	for k, v := range entries {
		b.NodeSelector[k] = v
	}
	return b
}

// WithTolerations adds the given value to the Tolerations field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the Tolerations field.
func (b *SchedulingApplyConfiguration) WithTolerations(values ...*v1.TolerationApplyConfiguration) *SchedulingApplyConfiguration {
	for i := range values {
		if values[i] == nil {
			panic("nil value passed to WithTolerations")
		}
		b.Tolerations = append(b.Tolerations, *values[i])
	}
	return b
}
