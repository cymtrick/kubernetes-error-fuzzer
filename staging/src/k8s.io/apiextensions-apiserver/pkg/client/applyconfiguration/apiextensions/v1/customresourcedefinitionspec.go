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
	apiextensionsv1 "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1"
)

// CustomResourceDefinitionSpecApplyConfiguration represents an declarative configuration of the CustomResourceDefinitionSpec type for use
// with apply.
type CustomResourceDefinitionSpecApplyConfiguration struct {
	Group                 *string                                             `json:"group,omitempty"`
	Names                 *CustomResourceDefinitionNamesApplyConfiguration    `json:"names,omitempty"`
	Scope                 *apiextensionsv1.ResourceScope                      `json:"scope,omitempty"`
	Versions              []CustomResourceDefinitionVersionApplyConfiguration `json:"versions,omitempty"`
	Conversion            *CustomResourceConversionApplyConfiguration         `json:"conversion,omitempty"`
	PreserveUnknownFields *bool                                               `json:"preserveUnknownFields,omitempty"`
}

// CustomResourceDefinitionSpecApplyConfiguration constructs an declarative configuration of the CustomResourceDefinitionSpec type for use with
// apply.
func CustomResourceDefinitionSpec() *CustomResourceDefinitionSpecApplyConfiguration {
	return &CustomResourceDefinitionSpecApplyConfiguration{}
}

// WithGroup sets the Group field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Group field is set to the value of the last call.
func (b *CustomResourceDefinitionSpecApplyConfiguration) WithGroup(value string) *CustomResourceDefinitionSpecApplyConfiguration {
	b.Group = &value
	return b
}

// WithNames sets the Names field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Names field is set to the value of the last call.
func (b *CustomResourceDefinitionSpecApplyConfiguration) WithNames(value *CustomResourceDefinitionNamesApplyConfiguration) *CustomResourceDefinitionSpecApplyConfiguration {
	b.Names = value
	return b
}

// WithScope sets the Scope field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Scope field is set to the value of the last call.
func (b *CustomResourceDefinitionSpecApplyConfiguration) WithScope(value apiextensionsv1.ResourceScope) *CustomResourceDefinitionSpecApplyConfiguration {
	b.Scope = &value
	return b
}

// WithVersions adds the given value to the Versions field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the Versions field.
func (b *CustomResourceDefinitionSpecApplyConfiguration) WithVersions(values ...*CustomResourceDefinitionVersionApplyConfiguration) *CustomResourceDefinitionSpecApplyConfiguration {
	for i := range values {
		if values[i] == nil {
			panic("nil value passed to WithVersions")
		}
		b.Versions = append(b.Versions, *values[i])
	}
	return b
}

// WithConversion sets the Conversion field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Conversion field is set to the value of the last call.
func (b *CustomResourceDefinitionSpecApplyConfiguration) WithConversion(value *CustomResourceConversionApplyConfiguration) *CustomResourceDefinitionSpecApplyConfiguration {
	b.Conversion = value
	return b
}

// WithPreserveUnknownFields sets the PreserveUnknownFields field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the PreserveUnknownFields field is set to the value of the last call.
func (b *CustomResourceDefinitionSpecApplyConfiguration) WithPreserveUnknownFields(value bool) *CustomResourceDefinitionSpecApplyConfiguration {
	b.PreserveUnknownFields = &value
	return b
}
