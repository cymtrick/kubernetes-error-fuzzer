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

// VolumeAttachmentStatusApplyConfiguration represents an declarative configuration of the VolumeAttachmentStatus type for use
// with apply.
type VolumeAttachmentStatusApplyConfiguration struct {
	Attached           *bool                          `json:"attached,omitempty"`
	AttachmentMetadata map[string]string              `json:"attachmentMetadata,omitempty"`
	AttachError        *VolumeErrorApplyConfiguration `json:"attachError,omitempty"`
	DetachError        *VolumeErrorApplyConfiguration `json:"detachError,omitempty"`
}

// VolumeAttachmentStatusApplyConfiguration constructs an declarative configuration of the VolumeAttachmentStatus type for use with
// apply.
func VolumeAttachmentStatus() *VolumeAttachmentStatusApplyConfiguration {
	return &VolumeAttachmentStatusApplyConfiguration{}
}

// WithAttached sets the Attached field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Attached field is set to the value of the last call.
func (b *VolumeAttachmentStatusApplyConfiguration) WithAttached(value bool) *VolumeAttachmentStatusApplyConfiguration {
	b.Attached = &value
	return b
}

// WithAttachmentMetadata puts the entries into the AttachmentMetadata field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, the entries provided by each call will be put on the AttachmentMetadata field,
// overwriting an existing map entries in AttachmentMetadata field with the same key.
func (b *VolumeAttachmentStatusApplyConfiguration) WithAttachmentMetadata(entries map[string]string) *VolumeAttachmentStatusApplyConfiguration {
	if b.AttachmentMetadata == nil && len(entries) > 0 {
		b.AttachmentMetadata = make(map[string]string, len(entries))
	}
	for k, v := range entries {
		b.AttachmentMetadata[k] = v
	}
	return b
}

// WithAttachError sets the AttachError field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the AttachError field is set to the value of the last call.
func (b *VolumeAttachmentStatusApplyConfiguration) WithAttachError(value *VolumeErrorApplyConfiguration) *VolumeAttachmentStatusApplyConfiguration {
	b.AttachError = value
	return b
}

// WithDetachError sets the DetachError field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the DetachError field is set to the value of the last call.
func (b *VolumeAttachmentStatusApplyConfiguration) WithDetachError(value *VolumeErrorApplyConfiguration) *VolumeAttachmentStatusApplyConfiguration {
	b.DetachError = value
	return b
}
