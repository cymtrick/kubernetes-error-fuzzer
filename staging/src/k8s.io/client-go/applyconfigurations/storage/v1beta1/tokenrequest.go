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

// TokenRequestApplyConfiguration represents an declarative configuration of the TokenRequest type for use
// with apply.
type TokenRequestApplyConfiguration struct {
	Audience          *string `json:"audience,omitempty"`
	ExpirationSeconds *int64  `json:"expirationSeconds,omitempty"`
}

// TokenRequestApplyConfiguration constructs an declarative configuration of the TokenRequest type for use with
// apply.
func TokenRequest() *TokenRequestApplyConfiguration {
	return &TokenRequestApplyConfiguration{}
}

// WithAudience sets the Audience field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Audience field is set to the value of the last call.
func (b *TokenRequestApplyConfiguration) WithAudience(value string) *TokenRequestApplyConfiguration {
	b.Audience = &value
	return b
}

// WithExpirationSeconds sets the ExpirationSeconds field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the ExpirationSeconds field is set to the value of the last call.
func (b *TokenRequestApplyConfiguration) WithExpirationSeconds(value int64) *TokenRequestApplyConfiguration {
	b.ExpirationSeconds = &value
	return b
}
