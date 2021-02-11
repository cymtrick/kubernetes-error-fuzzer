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

// GlusterfsPersistentVolumeSourceApplyConfiguration represents an declarative configuration of the GlusterfsPersistentVolumeSource type for use
// with apply.
type GlusterfsPersistentVolumeSourceApplyConfiguration struct {
	EndpointsName      *string `json:"endpoints,omitempty"`
	Path               *string `json:"path,omitempty"`
	ReadOnly           *bool   `json:"readOnly,omitempty"`
	EndpointsNamespace *string `json:"endpointsNamespace,omitempty"`
}

// GlusterfsPersistentVolumeSourceApplyConfiguration constructs an declarative configuration of the GlusterfsPersistentVolumeSource type for use with
// apply.
func GlusterfsPersistentVolumeSource() *GlusterfsPersistentVolumeSourceApplyConfiguration {
	return &GlusterfsPersistentVolumeSourceApplyConfiguration{}
}

// WithEndpointsName sets the EndpointsName field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the EndpointsName field is set to the value of the last call.
func (b *GlusterfsPersistentVolumeSourceApplyConfiguration) WithEndpointsName(value string) *GlusterfsPersistentVolumeSourceApplyConfiguration {
	b.EndpointsName = &value
	return b
}

// WithPath sets the Path field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Path field is set to the value of the last call.
func (b *GlusterfsPersistentVolumeSourceApplyConfiguration) WithPath(value string) *GlusterfsPersistentVolumeSourceApplyConfiguration {
	b.Path = &value
	return b
}

// WithReadOnly sets the ReadOnly field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the ReadOnly field is set to the value of the last call.
func (b *GlusterfsPersistentVolumeSourceApplyConfiguration) WithReadOnly(value bool) *GlusterfsPersistentVolumeSourceApplyConfiguration {
	b.ReadOnly = &value
	return b
}

// WithEndpointsNamespace sets the EndpointsNamespace field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the EndpointsNamespace field is set to the value of the last call.
func (b *GlusterfsPersistentVolumeSourceApplyConfiguration) WithEndpointsNamespace(value string) *GlusterfsPersistentVolumeSourceApplyConfiguration {
	b.EndpointsNamespace = &value
	return b
}
