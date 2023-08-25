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

// CustomResourceDefinitionNamesApplyConfiguration represents an declarative configuration of the CustomResourceDefinitionNames type for use
// with apply.
type CustomResourceDefinitionNamesApplyConfiguration struct {
	Plural     *string  `json:"plural,omitempty"`
	Singular   *string  `json:"singular,omitempty"`
	ShortNames []string `json:"shortNames,omitempty"`
	Kind       *string  `json:"kind,omitempty"`
	ListKind   *string  `json:"listKind,omitempty"`
	Categories []string `json:"categories,omitempty"`
}

// CustomResourceDefinitionNamesApplyConfiguration constructs an declarative configuration of the CustomResourceDefinitionNames type for use with
// apply.
func CustomResourceDefinitionNames() *CustomResourceDefinitionNamesApplyConfiguration {
	return &CustomResourceDefinitionNamesApplyConfiguration{}
}

// WithPlural sets the Plural field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Plural field is set to the value of the last call.
func (b *CustomResourceDefinitionNamesApplyConfiguration) WithPlural(value string) *CustomResourceDefinitionNamesApplyConfiguration {
	b.Plural = &value
	return b
}

// WithSingular sets the Singular field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Singular field is set to the value of the last call.
func (b *CustomResourceDefinitionNamesApplyConfiguration) WithSingular(value string) *CustomResourceDefinitionNamesApplyConfiguration {
	b.Singular = &value
	return b
}

// WithShortNames adds the given value to the ShortNames field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the ShortNames field.
func (b *CustomResourceDefinitionNamesApplyConfiguration) WithShortNames(values ...string) *CustomResourceDefinitionNamesApplyConfiguration {
	for i := range values {
		b.ShortNames = append(b.ShortNames, values[i])
	}
	return b
}

// WithKind sets the Kind field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Kind field is set to the value of the last call.
func (b *CustomResourceDefinitionNamesApplyConfiguration) WithKind(value string) *CustomResourceDefinitionNamesApplyConfiguration {
	b.Kind = &value
	return b
}

// WithListKind sets the ListKind field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the ListKind field is set to the value of the last call.
func (b *CustomResourceDefinitionNamesApplyConfiguration) WithListKind(value string) *CustomResourceDefinitionNamesApplyConfiguration {
	b.ListKind = &value
	return b
}

// WithCategories adds the given value to the Categories field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the Categories field.
func (b *CustomResourceDefinitionNamesApplyConfiguration) WithCategories(values ...string) *CustomResourceDefinitionNamesApplyConfiguration {
	for i := range values {
		b.Categories = append(b.Categories, values[i])
	}
	return b
}
