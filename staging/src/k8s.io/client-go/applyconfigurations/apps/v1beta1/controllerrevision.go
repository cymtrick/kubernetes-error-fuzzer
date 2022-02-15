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

import (
	v1beta1 "k8s.io/api/apps/v1beta1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	types "k8s.io/apimachinery/pkg/types"
	managedfields "k8s.io/apimachinery/pkg/util/managedfields"
	internal "k8s.io/client-go/applyconfigurations/internal"
	v1 "k8s.io/client-go/applyconfigurations/meta/v1"
)

// ControllerRevisionApplyConfiguration represents an declarative configuration of the ControllerRevision type for use
// with apply.
type ControllerRevisionApplyConfiguration struct {
	v1.TypeMetaApplyConfiguration    `json:",inline"`
	*v1.ObjectMetaApplyConfiguration `json:"metadata,omitempty"`
	Data                             *runtime.RawExtension `json:"data,omitempty"`
	Revision                         *int64                `json:"revision,omitempty"`
}

// ControllerRevision constructs an declarative configuration of the ControllerRevision type for use with
// apply.
func ControllerRevision(name, namespace string) *ControllerRevisionApplyConfiguration {
	b := &ControllerRevisionApplyConfiguration{}
	b.WithName(name)
	b.WithNamespace(namespace)
	b.WithKind("ControllerRevision")
	b.WithAPIVersion("apps/v1beta1")
	return b
}

// ExtractControllerRevision extracts the applied configuration owned by fieldManager from
// controllerRevision. If no managedFields are found in controllerRevision for fieldManager, a
// ControllerRevisionApplyConfiguration is returned with only the Name, Namespace (if applicable),
// APIVersion and Kind populated. It is possible that no managed fields were found for because other
// field managers have taken ownership of all the fields previously owned by fieldManager, or because
// the fieldManager never owned fields any fields.
// controllerRevision must be a unmodified ControllerRevision API object that was retrieved from the Kubernetes API.
// ExtractControllerRevision provides a way to perform a extract/modify-in-place/apply workflow.
// Note that an extracted apply configuration will contain fewer fields than what the fieldManager previously
// applied if another fieldManager has updated or force applied any of the previously applied fields.
// Experimental!
func ExtractControllerRevision(controllerRevision *v1beta1.ControllerRevision, fieldManager string) (*ControllerRevisionApplyConfiguration, error) {
	return extractControllerRevision(controllerRevision, fieldManager, "")
}

// ExtractControllerRevisionStatus is the same as ExtractControllerRevision except
// that it extracts the status subresource applied configuration.
// Experimental!
func ExtractControllerRevisionStatus(controllerRevision *v1beta1.ControllerRevision, fieldManager string) (*ControllerRevisionApplyConfiguration, error) {
	return extractControllerRevision(controllerRevision, fieldManager, "status")
}

func extractControllerRevision(controllerRevision *v1beta1.ControllerRevision, fieldManager string, subresource string) (*ControllerRevisionApplyConfiguration, error) {
	b := &ControllerRevisionApplyConfiguration{}
	err := managedfields.ExtractInto(controllerRevision, internal.Parser().Type("io.k8s.api.apps.v1beta1.ControllerRevision"), fieldManager, b, subresource)
	if err != nil {
		return nil, err
	}
	b.WithName(controllerRevision.Name)
	b.WithNamespace(controllerRevision.Namespace)

	b.WithKind("ControllerRevision")
	b.WithAPIVersion("apps/v1beta1")
	return b, nil
}

// WithKind sets the Kind field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Kind field is set to the value of the last call.
func (b *ControllerRevisionApplyConfiguration) WithKind(value string) *ControllerRevisionApplyConfiguration {
	b.Kind = &value
	return b
}

// WithAPIVersion sets the APIVersion field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the APIVersion field is set to the value of the last call.
func (b *ControllerRevisionApplyConfiguration) WithAPIVersion(value string) *ControllerRevisionApplyConfiguration {
	b.APIVersion = &value
	return b
}

// WithName sets the Name field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Name field is set to the value of the last call.
func (b *ControllerRevisionApplyConfiguration) WithName(value string) *ControllerRevisionApplyConfiguration {
	b.ensureObjectMetaApplyConfigurationExists()
	b.Name = &value
	return b
}

// WithGenerateName sets the GenerateName field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the GenerateName field is set to the value of the last call.
func (b *ControllerRevisionApplyConfiguration) WithGenerateName(value string) *ControllerRevisionApplyConfiguration {
	b.ensureObjectMetaApplyConfigurationExists()
	b.GenerateName = &value
	return b
}

// WithNamespace sets the Namespace field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Namespace field is set to the value of the last call.
func (b *ControllerRevisionApplyConfiguration) WithNamespace(value string) *ControllerRevisionApplyConfiguration {
	b.ensureObjectMetaApplyConfigurationExists()
	b.Namespace = &value
	return b
}

// WithUID sets the UID field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the UID field is set to the value of the last call.
func (b *ControllerRevisionApplyConfiguration) WithUID(value types.UID) *ControllerRevisionApplyConfiguration {
	b.ensureObjectMetaApplyConfigurationExists()
	b.UID = &value
	return b
}

// WithResourceVersion sets the ResourceVersion field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the ResourceVersion field is set to the value of the last call.
func (b *ControllerRevisionApplyConfiguration) WithResourceVersion(value string) *ControllerRevisionApplyConfiguration {
	b.ensureObjectMetaApplyConfigurationExists()
	b.ResourceVersion = &value
	return b
}

// WithGeneration sets the Generation field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Generation field is set to the value of the last call.
func (b *ControllerRevisionApplyConfiguration) WithGeneration(value int64) *ControllerRevisionApplyConfiguration {
	b.ensureObjectMetaApplyConfigurationExists()
	b.Generation = &value
	return b
}

// WithCreationTimestamp sets the CreationTimestamp field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the CreationTimestamp field is set to the value of the last call.
func (b *ControllerRevisionApplyConfiguration) WithCreationTimestamp(value metav1.Time) *ControllerRevisionApplyConfiguration {
	b.ensureObjectMetaApplyConfigurationExists()
	b.CreationTimestamp = &value
	return b
}

// WithDeletionTimestamp sets the DeletionTimestamp field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the DeletionTimestamp field is set to the value of the last call.
func (b *ControllerRevisionApplyConfiguration) WithDeletionTimestamp(value metav1.Time) *ControllerRevisionApplyConfiguration {
	b.ensureObjectMetaApplyConfigurationExists()
	b.DeletionTimestamp = &value
	return b
}

// WithDeletionGracePeriodSeconds sets the DeletionGracePeriodSeconds field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the DeletionGracePeriodSeconds field is set to the value of the last call.
func (b *ControllerRevisionApplyConfiguration) WithDeletionGracePeriodSeconds(value int64) *ControllerRevisionApplyConfiguration {
	b.ensureObjectMetaApplyConfigurationExists()
	b.DeletionGracePeriodSeconds = &value
	return b
}

// WithLabels puts the entries into the Labels field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, the entries provided by each call will be put on the Labels field,
// overwriting an existing map entries in Labels field with the same key.
func (b *ControllerRevisionApplyConfiguration) WithLabels(entries map[string]string) *ControllerRevisionApplyConfiguration {
	b.ensureObjectMetaApplyConfigurationExists()
	if b.Labels == nil && len(entries) > 0 {
		b.Labels = make(map[string]string, len(entries))
	}
	for k, v := range entries {
		b.Labels[k] = v
	}
	return b
}

// WithAnnotations puts the entries into the Annotations field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, the entries provided by each call will be put on the Annotations field,
// overwriting an existing map entries in Annotations field with the same key.
func (b *ControllerRevisionApplyConfiguration) WithAnnotations(entries map[string]string) *ControllerRevisionApplyConfiguration {
	b.ensureObjectMetaApplyConfigurationExists()
	if b.Annotations == nil && len(entries) > 0 {
		b.Annotations = make(map[string]string, len(entries))
	}
	for k, v := range entries {
		b.Annotations[k] = v
	}
	return b
}

// WithOwnerReferences adds the given value to the OwnerReferences field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the OwnerReferences field.
func (b *ControllerRevisionApplyConfiguration) WithOwnerReferences(values ...*v1.OwnerReferenceApplyConfiguration) *ControllerRevisionApplyConfiguration {
	b.ensureObjectMetaApplyConfigurationExists()
	for i := range values {
		if values[i] == nil {
			panic("nil value passed to WithOwnerReferences")
		}
		b.OwnerReferences = append(b.OwnerReferences, *values[i])
	}
	return b
}

// WithFinalizers adds the given value to the Finalizers field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the Finalizers field.
func (b *ControllerRevisionApplyConfiguration) WithFinalizers(values ...string) *ControllerRevisionApplyConfiguration {
	b.ensureObjectMetaApplyConfigurationExists()
	for i := range values {
		b.Finalizers = append(b.Finalizers, values[i])
	}
	return b
}

// WithClusterName sets the ClusterName field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the ClusterName field is set to the value of the last call.
func (b *ControllerRevisionApplyConfiguration) WithClusterName(value string) *ControllerRevisionApplyConfiguration {
	b.ensureObjectMetaApplyConfigurationExists()
	b.ClusterName = &value
	return b
}

func (b *ControllerRevisionApplyConfiguration) ensureObjectMetaApplyConfigurationExists() {
	if b.ObjectMetaApplyConfiguration == nil {
		b.ObjectMetaApplyConfiguration = &v1.ObjectMetaApplyConfiguration{}
	}
}

// WithData sets the Data field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Data field is set to the value of the last call.
func (b *ControllerRevisionApplyConfiguration) WithData(value runtime.RawExtension) *ControllerRevisionApplyConfiguration {
	b.Data = &value
	return b
}

// WithRevision sets the Revision field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Revision field is set to the value of the last call.
func (b *ControllerRevisionApplyConfiguration) WithRevision(value int64) *ControllerRevisionApplyConfiguration {
	b.Revision = &value
	return b
}
