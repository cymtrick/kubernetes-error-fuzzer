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

// CSINodeDriverApplyConfiguration represents an declarative configuration of the CSINodeDriver type for use
// with apply.
type CSINodeDriverApplyConfiguration struct {
	Name         *string                                `json:"name,omitempty"`
	NodeID       *string                                `json:"nodeID,omitempty"`
	TopologyKeys []string                               `json:"topologyKeys,omitempty"`
	Allocatable  *VolumeNodeResourcesApplyConfiguration `json:"allocatable,omitempty"`
}

// CSINodeDriverApplyConfiguration constructs an declarative configuration of the CSINodeDriver type for use with
// apply.
func CSINodeDriver() *CSINodeDriverApplyConfiguration {
	return &CSINodeDriverApplyConfiguration{}
}

// WithName sets the Name field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Name field is set to the value of the last call.
func (b *CSINodeDriverApplyConfiguration) WithName(value string) *CSINodeDriverApplyConfiguration {
	b.Name = &value
	return b
}

// WithNodeID sets the NodeID field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the NodeID field is set to the value of the last call.
func (b *CSINodeDriverApplyConfiguration) WithNodeID(value string) *CSINodeDriverApplyConfiguration {
	b.NodeID = &value
	return b
}

// WithTopologyKeys adds the given value to the TopologyKeys field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the TopologyKeys field.
func (b *CSINodeDriverApplyConfiguration) WithTopologyKeys(values ...string) *CSINodeDriverApplyConfiguration {
	for i := range values {
		b.TopologyKeys = append(b.TopologyKeys, values[i])
	}
	return b
}

// WithAllocatable sets the Allocatable field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Allocatable field is set to the value of the last call.
func (b *CSINodeDriverApplyConfiguration) WithAllocatable(value *VolumeNodeResourcesApplyConfiguration) *CSINodeDriverApplyConfiguration {
	b.Allocatable = value
	return b
}
