/*
Copyright 2015 The Kubernetes Authors.

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

package unstructured

import (
	"bytes"
	"fmt"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/types"
	utilruntime "k8s.io/apimachinery/pkg/util/runtime"
)

// Unstructured allows objects that do not have Golang structs registered to be manipulated
// generically. This can be used to deal with the API objects from a plug-in. Unstructured
// objects still have functioning TypeMeta features-- kind, version, etc.
//
// WARNING: This object has accessors for the v1 standard metadata. You *MUST NOT* use this
// type if you are dealing with objects that are not in the server meta v1 schema.
//
// TODO: make the serialization part of this type distinct from the field accessors.
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +k8s:deepcopy-gen=true
type Unstructured struct {
	// Object is a JSON compatible map with string, float, int, bool, []interface{}, or
	// map[string]interface{}
	// children.
	Object map[string]interface{}
}

var _ metav1.Object = &Unstructured{}
var _ runtime.Unstructured = &Unstructured{}

func (obj *Unstructured) GetObjectKind() schema.ObjectKind { return obj }

func (obj *Unstructured) IsList() bool {
	if obj.Object != nil {
		_, ok := obj.Object["items"]
		return ok
	}
	return false
}

func (obj *Unstructured) EachListItem(fn func(runtime.Object) error) error {
	if obj.Object == nil {
		return fmt.Errorf("content is not a list")
	}
	field, ok := obj.Object["items"]
	if !ok {
		return fmt.Errorf("content is not a list")
	}
	items, ok := field.([]interface{})
	if !ok {
		return nil
	}
	for _, item := range items {
		child, ok := item.(map[string]interface{})
		if !ok {
			return fmt.Errorf("items member is not an object")
		}
		if err := fn(&Unstructured{Object: child}); err != nil {
			return err
		}
	}
	return nil
}

func (obj *Unstructured) UnstructuredContent() map[string]interface{} {
	if obj.Object == nil {
		obj.Object = make(map[string]interface{})
	}
	return obj.Object
}

func (obj *Unstructured) SetUnstructuredContent(content map[string]interface{}) {
	obj.Object = content
}

// MarshalJSON ensures that the unstructured object produces proper
// JSON when passed to Go's standard JSON library.
func (u *Unstructured) MarshalJSON() ([]byte, error) {
	var buf bytes.Buffer
	err := UnstructuredJSONScheme.Encode(u, &buf)
	return buf.Bytes(), err
}

// UnmarshalJSON ensures that the unstructured object properly decodes
// JSON when passed to Go's standard JSON library.
func (u *Unstructured) UnmarshalJSON(b []byte) error {
	_, _, err := UnstructuredJSONScheme.Decode(b, nil, u)
	return err
}

func (in *Unstructured) DeepCopy() *Unstructured {
	if in == nil {
		return nil
	}
	out := new(Unstructured)
	*out = *in
	out.Object = runtime.DeepCopyJSON(in.Object)
	return out
}

func (u *Unstructured) setNestedField(value interface{}, fields ...string) {
	if u.Object == nil {
		u.Object = make(map[string]interface{})
	}
	SetNestedField(u.Object, value, fields...)
}

func (u *Unstructured) setNestedSlice(value []string, fields ...string) {
	if u.Object == nil {
		u.Object = make(map[string]interface{})
	}
	SetNestedStringSlice(u.Object, value, fields...)
}

func (u *Unstructured) setNestedMap(value map[string]string, fields ...string) {
	if u.Object == nil {
		u.Object = make(map[string]interface{})
	}
	SetNestedStringMap(u.Object, value, fields...)
}

func (u *Unstructured) GetOwnerReferences() []metav1.OwnerReference {
	field, ok := nestedFieldNoCopy(u.Object, "metadata", "ownerReferences")
	if !ok {
		return nil
	}
	original, ok := field.([]interface{})
	if !ok {
		return nil
	}
	ret := make([]metav1.OwnerReference, 0, len(original))
	for _, obj := range original {
		o, ok := obj.(map[string]interface{})
		if !ok {
			// expected map[string]interface{}, got something else
			return nil
		}
		ret = append(ret, extractOwnerReference(o))
	}
	return ret
}

func (u *Unstructured) SetOwnerReferences(references []metav1.OwnerReference) {
	newReferences := make([]interface{}, 0, len(references))
	for _, reference := range references {
		newReferences = append(newReferences, setOwnerReference(reference))
	}
	u.setNestedField(newReferences, "metadata", "ownerReferences")
}

func (u *Unstructured) GetAPIVersion() string {
	return getNestedString(u.Object, "apiVersion")
}

func (u *Unstructured) SetAPIVersion(version string) {
	u.setNestedField(version, "apiVersion")
}

func (u *Unstructured) GetKind() string {
	return getNestedString(u.Object, "kind")
}

func (u *Unstructured) SetKind(kind string) {
	u.setNestedField(kind, "kind")
}

func (u *Unstructured) GetNamespace() string {
	return getNestedString(u.Object, "metadata", "namespace")
}

func (u *Unstructured) SetNamespace(namespace string) {
	u.setNestedField(namespace, "metadata", "namespace")
}

func (u *Unstructured) GetName() string {
	return getNestedString(u.Object, "metadata", "name")
}

func (u *Unstructured) SetName(name string) {
	u.setNestedField(name, "metadata", "name")
}

func (u *Unstructured) GetGenerateName() string {
	return getNestedString(u.Object, "metadata", "generateName")
}

func (u *Unstructured) SetGenerateName(name string) {
	u.setNestedField(name, "metadata", "generateName")
}

func (u *Unstructured) GetUID() types.UID {
	return types.UID(getNestedString(u.Object, "metadata", "uid"))
}

func (u *Unstructured) SetUID(uid types.UID) {
	u.setNestedField(string(uid), "metadata", "uid")
}

func (u *Unstructured) GetResourceVersion() string {
	return getNestedString(u.Object, "metadata", "resourceVersion")
}

func (u *Unstructured) SetResourceVersion(version string) {
	u.setNestedField(version, "metadata", "resourceVersion")
}

func (u *Unstructured) GetGeneration() int64 {
	val, ok := NestedInt64(u.Object, "metadata", "generation")
	if !ok {
		return 0
	}
	return val
}

func (u *Unstructured) SetGeneration(generation int64) {
	u.setNestedField(generation, "metadata", "generation")
}

func (u *Unstructured) GetSelfLink() string {
	return getNestedString(u.Object, "metadata", "selfLink")
}

func (u *Unstructured) SetSelfLink(selfLink string) {
	u.setNestedField(selfLink, "metadata", "selfLink")
}

func (u *Unstructured) GetContinue() string {
	return getNestedString(u.Object, "metadata", "continue")
}

func (u *Unstructured) SetContinue(c string) {
	u.setNestedField(c, "metadata", "continue")
}

func (u *Unstructured) GetCreationTimestamp() metav1.Time {
	var timestamp metav1.Time
	timestamp.UnmarshalQueryParameter(getNestedString(u.Object, "metadata", "creationTimestamp"))
	return timestamp
}

func (u *Unstructured) SetCreationTimestamp(timestamp metav1.Time) {
	ts, _ := timestamp.MarshalQueryParameter()
	if len(ts) == 0 || timestamp.Time.IsZero() {
		RemoveNestedField(u.Object, "metadata", "creationTimestamp")
		return
	}
	u.setNestedField(ts, "metadata", "creationTimestamp")
}

func (u *Unstructured) GetDeletionTimestamp() *metav1.Time {
	var timestamp metav1.Time
	timestamp.UnmarshalQueryParameter(getNestedString(u.Object, "metadata", "deletionTimestamp"))
	if timestamp.IsZero() {
		return nil
	}
	return &timestamp
}

func (u *Unstructured) SetDeletionTimestamp(timestamp *metav1.Time) {
	if timestamp == nil {
		RemoveNestedField(u.Object, "metadata", "deletionTimestamp")
		return
	}
	ts, _ := timestamp.MarshalQueryParameter()
	u.setNestedField(ts, "metadata", "deletionTimestamp")
}

func (u *Unstructured) GetDeletionGracePeriodSeconds() *int64 {
	val, ok := NestedInt64(u.Object, "metadata", "deletionGracePeriodSeconds")
	if !ok {
		return nil
	}
	return &val
}

func (u *Unstructured) SetDeletionGracePeriodSeconds(deletionGracePeriodSeconds *int64) {
	if deletionGracePeriodSeconds == nil {
		RemoveNestedField(u.Object, "metadata", "deletionGracePeriodSeconds")
		return
	}
	u.setNestedField(*deletionGracePeriodSeconds, "metadata", "deletionGracePeriodSeconds")
}

func (u *Unstructured) GetLabels() map[string]string {
	m, ok := NestedStringMap(u.Object, "metadata", "labels")
	if !ok {
		return nil
	}
	return m
}

func (u *Unstructured) SetLabels(labels map[string]string) {
	u.setNestedMap(labels, "metadata", "labels")
}

func (u *Unstructured) GetAnnotations() map[string]string {
	m, ok := NestedStringMap(u.Object, "metadata", "annotations")
	if !ok {
		return nil
	}
	return m
}

func (u *Unstructured) SetAnnotations(annotations map[string]string) {
	u.setNestedMap(annotations, "metadata", "annotations")
}

func (u *Unstructured) SetGroupVersionKind(gvk schema.GroupVersionKind) {
	u.SetAPIVersion(gvk.GroupVersion().String())
	u.SetKind(gvk.Kind)
}

func (u *Unstructured) GroupVersionKind() schema.GroupVersionKind {
	gv, err := schema.ParseGroupVersion(u.GetAPIVersion())
	if err != nil {
		return schema.GroupVersionKind{}
	}
	gvk := gv.WithKind(u.GetKind())
	return gvk
}

func (u *Unstructured) GetInitializers() *metav1.Initializers {
	field, ok := nestedFieldNoCopy(u.Object, "metadata", "initializers")
	if !ok {
		return nil
	}
	obj, ok := field.(map[string]interface{})
	if !ok {
		// expected map[string]interface{}, got something else
		return nil
	}
	out := &metav1.Initializers{}
	if err := runtime.DefaultUnstructuredConverter.FromUnstructured(obj, out); err != nil {
		utilruntime.HandleError(fmt.Errorf("unable to retrieve initializers for object: %v", err))
	}
	return out
}

func (u *Unstructured) SetInitializers(initializers *metav1.Initializers) {
	if initializers == nil {
		RemoveNestedField(u.Object, "metadata", "initializers")
		return
	}
	out, err := runtime.DefaultUnstructuredConverter.ToUnstructured(initializers)
	if err != nil {
		utilruntime.HandleError(fmt.Errorf("unable to retrieve initializers for object: %v", err))
	}
	u.setNestedField(out, "metadata", "initializers")
}

func (u *Unstructured) GetFinalizers() []string {
	val, ok := NestedStringSlice(u.Object, "metadata", "finalizers")
	if !ok {
		return nil
	}
	return val
}

func (u *Unstructured) SetFinalizers(finalizers []string) {
	u.setNestedSlice(finalizers, "metadata", "finalizers")
}

func (u *Unstructured) GetClusterName() string {
	return getNestedString(u.Object, "metadata", "clusterName")
}

func (u *Unstructured) SetClusterName(clusterName string) {
	u.setNestedField(clusterName, "metadata", "clusterName")
}
