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

// EnvVarApplyConfiguration represents an declarative configuration of the EnvVar type for use
// with apply.
type EnvVarApplyConfiguration struct {
	Name      *string                         `json:"name,omitempty"`
	Value     *string                         `json:"value,omitempty"`
	ValueFrom *EnvVarSourceApplyConfiguration `json:"valueFrom,omitempty"`
}

// EnvVarApplyConfiguration constructs an declarative configuration of the EnvVar type for use with
// apply.
func EnvVar() *EnvVarApplyConfiguration {
	return &EnvVarApplyConfiguration{}
}

// WithName sets the Name field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Name field is set to the value of the last call.
func (b *EnvVarApplyConfiguration) WithName(value string) *EnvVarApplyConfiguration {
	b.Name = &value
	return b
}

// WithValue sets the Value field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Value field is set to the value of the last call.
func (b *EnvVarApplyConfiguration) WithValue(value string) *EnvVarApplyConfiguration {
	b.Value = &value
	return b
}

// WithValueFrom sets the ValueFrom field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the ValueFrom field is set to the value of the last call.
func (b *EnvVarApplyConfiguration) WithValueFrom(value *EnvVarSourceApplyConfiguration) *EnvVarApplyConfiguration {
	b.ValueFrom = value
	return b
}
