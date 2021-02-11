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

// DeploymentStatusApplyConfiguration represents an declarative configuration of the DeploymentStatus type for use
// with apply.
type DeploymentStatusApplyConfiguration struct {
	ObservedGeneration  *int64                                  `json:"observedGeneration,omitempty"`
	Replicas            *int32                                  `json:"replicas,omitempty"`
	UpdatedReplicas     *int32                                  `json:"updatedReplicas,omitempty"`
	ReadyReplicas       *int32                                  `json:"readyReplicas,omitempty"`
	AvailableReplicas   *int32                                  `json:"availableReplicas,omitempty"`
	UnavailableReplicas *int32                                  `json:"unavailableReplicas,omitempty"`
	Conditions          []DeploymentConditionApplyConfiguration `json:"conditions,omitempty"`
	CollisionCount      *int32                                  `json:"collisionCount,omitempty"`
}

// DeploymentStatusApplyConfiguration constructs an declarative configuration of the DeploymentStatus type for use with
// apply.
func DeploymentStatus() *DeploymentStatusApplyConfiguration {
	return &DeploymentStatusApplyConfiguration{}
}

// WithObservedGeneration sets the ObservedGeneration field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the ObservedGeneration field is set to the value of the last call.
func (b *DeploymentStatusApplyConfiguration) WithObservedGeneration(value int64) *DeploymentStatusApplyConfiguration {
	b.ObservedGeneration = &value
	return b
}

// WithReplicas sets the Replicas field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Replicas field is set to the value of the last call.
func (b *DeploymentStatusApplyConfiguration) WithReplicas(value int32) *DeploymentStatusApplyConfiguration {
	b.Replicas = &value
	return b
}

// WithUpdatedReplicas sets the UpdatedReplicas field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the UpdatedReplicas field is set to the value of the last call.
func (b *DeploymentStatusApplyConfiguration) WithUpdatedReplicas(value int32) *DeploymentStatusApplyConfiguration {
	b.UpdatedReplicas = &value
	return b
}

// WithReadyReplicas sets the ReadyReplicas field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the ReadyReplicas field is set to the value of the last call.
func (b *DeploymentStatusApplyConfiguration) WithReadyReplicas(value int32) *DeploymentStatusApplyConfiguration {
	b.ReadyReplicas = &value
	return b
}

// WithAvailableReplicas sets the AvailableReplicas field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the AvailableReplicas field is set to the value of the last call.
func (b *DeploymentStatusApplyConfiguration) WithAvailableReplicas(value int32) *DeploymentStatusApplyConfiguration {
	b.AvailableReplicas = &value
	return b
}

// WithUnavailableReplicas sets the UnavailableReplicas field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the UnavailableReplicas field is set to the value of the last call.
func (b *DeploymentStatusApplyConfiguration) WithUnavailableReplicas(value int32) *DeploymentStatusApplyConfiguration {
	b.UnavailableReplicas = &value
	return b
}

// WithConditions adds the given value to the Conditions field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the Conditions field.
func (b *DeploymentStatusApplyConfiguration) WithConditions(values ...*DeploymentConditionApplyConfiguration) *DeploymentStatusApplyConfiguration {
	for i := range values {
		if values[i] == nil {
			panic("nil value passed to WithConditions")
		}
		b.Conditions = append(b.Conditions, *values[i])
	}
	return b
}

// WithCollisionCount sets the CollisionCount field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the CollisionCount field is set to the value of the last call.
func (b *DeploymentStatusApplyConfiguration) WithCollisionCount(value int32) *DeploymentStatusApplyConfiguration {
	b.CollisionCount = &value
	return b
}
