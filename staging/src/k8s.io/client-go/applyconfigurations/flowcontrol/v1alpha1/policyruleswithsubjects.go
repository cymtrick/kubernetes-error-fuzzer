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

// PolicyRulesWithSubjectsApplyConfiguration represents an declarative configuration of the PolicyRulesWithSubjects type for use
// with apply.
type PolicyRulesWithSubjectsApplyConfiguration struct {
	Subjects         []SubjectApplyConfiguration               `json:"subjects,omitempty"`
	ResourceRules    []ResourcePolicyRuleApplyConfiguration    `json:"resourceRules,omitempty"`
	NonResourceRules []NonResourcePolicyRuleApplyConfiguration `json:"nonResourceRules,omitempty"`
}

// PolicyRulesWithSubjectsApplyConfiguration constructs an declarative configuration of the PolicyRulesWithSubjects type for use with
// apply.
func PolicyRulesWithSubjects() *PolicyRulesWithSubjectsApplyConfiguration {
	return &PolicyRulesWithSubjectsApplyConfiguration{}
}

// WithSubjects adds the given value to the Subjects field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the Subjects field.
func (b *PolicyRulesWithSubjectsApplyConfiguration) WithSubjects(values ...*SubjectApplyConfiguration) *PolicyRulesWithSubjectsApplyConfiguration {
	for i := range values {
		if values[i] == nil {
			panic("nil value passed to WithSubjects")
		}
		b.Subjects = append(b.Subjects, *values[i])
	}
	return b
}

// WithResourceRules adds the given value to the ResourceRules field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the ResourceRules field.
func (b *PolicyRulesWithSubjectsApplyConfiguration) WithResourceRules(values ...*ResourcePolicyRuleApplyConfiguration) *PolicyRulesWithSubjectsApplyConfiguration {
	for i := range values {
		if values[i] == nil {
			panic("nil value passed to WithResourceRules")
		}
		b.ResourceRules = append(b.ResourceRules, *values[i])
	}
	return b
}

// WithNonResourceRules adds the given value to the NonResourceRules field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the NonResourceRules field.
func (b *PolicyRulesWithSubjectsApplyConfiguration) WithNonResourceRules(values ...*NonResourcePolicyRuleApplyConfiguration) *PolicyRulesWithSubjectsApplyConfiguration {
	for i := range values {
		if values[i] == nil {
			panic("nil value passed to WithNonResourceRules")
		}
		b.NonResourceRules = append(b.NonResourceRules, *values[i])
	}
	return b
}
