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

// PodDNSConfigApplyConfiguration represents an declarative configuration of the PodDNSConfig type for use
// with apply.
type PodDNSConfigApplyConfiguration struct {
	Nameservers []string                               `json:"nameservers,omitempty"`
	Searches    []string                               `json:"searches,omitempty"`
	Options     []PodDNSConfigOptionApplyConfiguration `json:"options,omitempty"`
}

// PodDNSConfigApplyConfiguration constructs an declarative configuration of the PodDNSConfig type for use with
// apply.
func PodDNSConfig() *PodDNSConfigApplyConfiguration {
	return &PodDNSConfigApplyConfiguration{}
}

// WithNameservers adds the given value to the Nameservers field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the Nameservers field.
func (b *PodDNSConfigApplyConfiguration) WithNameservers(values ...string) *PodDNSConfigApplyConfiguration {
	for i := range values {
		b.Nameservers = append(b.Nameservers, values[i])
	}
	return b
}

// WithSearches adds the given value to the Searches field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the Searches field.
func (b *PodDNSConfigApplyConfiguration) WithSearches(values ...string) *PodDNSConfigApplyConfiguration {
	for i := range values {
		b.Searches = append(b.Searches, values[i])
	}
	return b
}

// WithOptions adds the given value to the Options field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the Options field.
func (b *PodDNSConfigApplyConfiguration) WithOptions(values ...*PodDNSConfigOptionApplyConfiguration) *PodDNSConfigApplyConfiguration {
	for i := range values {
		if values[i] == nil {
			panic("nil value passed to WithOptions")
		}
		b.Options = append(b.Options, *values[i])
	}
	return b
}
