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

// VsphereVirtualDiskVolumeSourceApplyConfiguration represents an declarative configuration of the VsphereVirtualDiskVolumeSource type for use
// with apply.
type VsphereVirtualDiskVolumeSourceApplyConfiguration struct {
	VolumePath        *string `json:"volumePath,omitempty"`
	FSType            *string `json:"fsType,omitempty"`
	StoragePolicyName *string `json:"storagePolicyName,omitempty"`
	StoragePolicyID   *string `json:"storagePolicyID,omitempty"`
}

// VsphereVirtualDiskVolumeSourceApplyConfiguration constructs an declarative configuration of the VsphereVirtualDiskVolumeSource type for use with
// apply.
func VsphereVirtualDiskVolumeSource() *VsphereVirtualDiskVolumeSourceApplyConfiguration {
	return &VsphereVirtualDiskVolumeSourceApplyConfiguration{}
}

// WithVolumePath sets the VolumePath field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the VolumePath field is set to the value of the last call.
func (b *VsphereVirtualDiskVolumeSourceApplyConfiguration) WithVolumePath(value string) *VsphereVirtualDiskVolumeSourceApplyConfiguration {
	b.VolumePath = &value
	return b
}

// WithFSType sets the FSType field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the FSType field is set to the value of the last call.
func (b *VsphereVirtualDiskVolumeSourceApplyConfiguration) WithFSType(value string) *VsphereVirtualDiskVolumeSourceApplyConfiguration {
	b.FSType = &value
	return b
}

// WithStoragePolicyName sets the StoragePolicyName field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the StoragePolicyName field is set to the value of the last call.
func (b *VsphereVirtualDiskVolumeSourceApplyConfiguration) WithStoragePolicyName(value string) *VsphereVirtualDiskVolumeSourceApplyConfiguration {
	b.StoragePolicyName = &value
	return b
}

// WithStoragePolicyID sets the StoragePolicyID field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the StoragePolicyID field is set to the value of the last call.
func (b *VsphereVirtualDiskVolumeSourceApplyConfiguration) WithStoragePolicyID(value string) *VsphereVirtualDiskVolumeSourceApplyConfiguration {
	b.StoragePolicyID = &value
	return b
}
