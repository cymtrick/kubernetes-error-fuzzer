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

// StorageVersionStatusApplyConfiguration represents an declarative configuration of the StorageVersionStatus type for use
// with apply.
type StorageVersionStatusApplyConfiguration struct {
	StorageVersions       []ServerStorageVersionApplyConfiguration    `json:"storageVersions,omitempty"`
	CommonEncodingVersion *string                                     `json:"commonEncodingVersion,omitempty"`
	Conditions            []StorageVersionConditionApplyConfiguration `json:"conditions,omitempty"`
}

// StorageVersionStatusApplyConfiguration constructs an declarative configuration of the StorageVersionStatus type for use with
// apply.
func StorageVersionStatus() *StorageVersionStatusApplyConfiguration {
	return &StorageVersionStatusApplyConfiguration{}
}

// WithStorageVersions adds the given value to the StorageVersions field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the StorageVersions field.
func (b *StorageVersionStatusApplyConfiguration) WithStorageVersions(values ...*ServerStorageVersionApplyConfiguration) *StorageVersionStatusApplyConfiguration {
	for i := range values {
		if values[i] == nil {
			panic("nil value passed to WithStorageVersions")
		}
		b.StorageVersions = append(b.StorageVersions, *values[i])
	}
	return b
}

// WithCommonEncodingVersion sets the CommonEncodingVersion field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the CommonEncodingVersion field is set to the value of the last call.
func (b *StorageVersionStatusApplyConfiguration) WithCommonEncodingVersion(value string) *StorageVersionStatusApplyConfiguration {
	b.CommonEncodingVersion = &value
	return b
}

// WithConditions adds the given value to the Conditions field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the Conditions field.
func (b *StorageVersionStatusApplyConfiguration) WithConditions(values ...*StorageVersionConditionApplyConfiguration) *StorageVersionStatusApplyConfiguration {
	for i := range values {
		if values[i] == nil {
			panic("nil value passed to WithConditions")
		}
		b.Conditions = append(b.Conditions, *values[i])
	}
	return b
}
