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

package v1beta2

// StatefulSetStatusApplyConfiguration represents an declarative configuration of the StatefulSetStatus type for use
// with apply.
type StatefulSetStatusApplyConfiguration struct {
	ObservedGeneration *int64                                   `json:"observedGeneration,omitempty"`
	Replicas           *int32                                   `json:"replicas,omitempty"`
	ReadyReplicas      *int32                                   `json:"readyReplicas,omitempty"`
	CurrentReplicas    *int32                                   `json:"currentReplicas,omitempty"`
	UpdatedReplicas    *int32                                   `json:"updatedReplicas,omitempty"`
	CurrentRevision    *string                                  `json:"currentRevision,omitempty"`
	UpdateRevision     *string                                  `json:"updateRevision,omitempty"`
	CollisionCount     *int32                                   `json:"collisionCount,omitempty"`
	Conditions         []StatefulSetConditionApplyConfiguration `json:"conditions,omitempty"`
}

// StatefulSetStatusApplyConfiguration constructs an declarative configuration of the StatefulSetStatus type for use with
// apply.
func StatefulSetStatus() *StatefulSetStatusApplyConfiguration {
	return &StatefulSetStatusApplyConfiguration{}
}

// WithObservedGeneration sets the ObservedGeneration field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the ObservedGeneration field is set to the value of the last call.
func (b *StatefulSetStatusApplyConfiguration) WithObservedGeneration(value int64) *StatefulSetStatusApplyConfiguration {
	b.ObservedGeneration = &value
	return b
}

// WithReplicas sets the Replicas field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Replicas field is set to the value of the last call.
func (b *StatefulSetStatusApplyConfiguration) WithReplicas(value int32) *StatefulSetStatusApplyConfiguration {
	b.Replicas = &value
	return b
}

// WithReadyReplicas sets the ReadyReplicas field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the ReadyReplicas field is set to the value of the last call.
func (b *StatefulSetStatusApplyConfiguration) WithReadyReplicas(value int32) *StatefulSetStatusApplyConfiguration {
	b.ReadyReplicas = &value
	return b
}

// WithCurrentReplicas sets the CurrentReplicas field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the CurrentReplicas field is set to the value of the last call.
func (b *StatefulSetStatusApplyConfiguration) WithCurrentReplicas(value int32) *StatefulSetStatusApplyConfiguration {
	b.CurrentReplicas = &value
	return b
}

// WithUpdatedReplicas sets the UpdatedReplicas field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the UpdatedReplicas field is set to the value of the last call.
func (b *StatefulSetStatusApplyConfiguration) WithUpdatedReplicas(value int32) *StatefulSetStatusApplyConfiguration {
	b.UpdatedReplicas = &value
	return b
}

// WithCurrentRevision sets the CurrentRevision field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the CurrentRevision field is set to the value of the last call.
func (b *StatefulSetStatusApplyConfiguration) WithCurrentRevision(value string) *StatefulSetStatusApplyConfiguration {
	b.CurrentRevision = &value
	return b
}

// WithUpdateRevision sets the UpdateRevision field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the UpdateRevision field is set to the value of the last call.
func (b *StatefulSetStatusApplyConfiguration) WithUpdateRevision(value string) *StatefulSetStatusApplyConfiguration {
	b.UpdateRevision = &value
	return b
}

// WithCollisionCount sets the CollisionCount field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the CollisionCount field is set to the value of the last call.
func (b *StatefulSetStatusApplyConfiguration) WithCollisionCount(value int32) *StatefulSetStatusApplyConfiguration {
	b.CollisionCount = &value
	return b
}

// WithConditions adds the given value to the Conditions field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the Conditions field.
func (b *StatefulSetStatusApplyConfiguration) WithConditions(values ...*StatefulSetConditionApplyConfiguration) *StatefulSetStatusApplyConfiguration {
	for i := range values {
		if values[i] == nil {
			panic("nil value passed to WithConditions")
		}
		b.Conditions = append(b.Conditions, *values[i])
	}
	return b
}
