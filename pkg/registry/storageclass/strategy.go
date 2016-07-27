/*
Copyright 2016 The Kubernetes Authors.

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

package storageclass

import (
	"fmt"

	"k8s.io/kubernetes/pkg/api"
	"k8s.io/kubernetes/pkg/apis/extensions"
	"k8s.io/kubernetes/pkg/apis/extensions/validation"
	"k8s.io/kubernetes/pkg/fields"
	"k8s.io/kubernetes/pkg/labels"
	"k8s.io/kubernetes/pkg/registry/generic"
	"k8s.io/kubernetes/pkg/runtime"
	"k8s.io/kubernetes/pkg/util/validation/field"
)

// storageClassStrategy implements behavior for StorageClass objects
type storageClassStrategy struct {
	runtime.ObjectTyper
	api.NameGenerator
}

// Strategy is the default logic that applies when creating and updating
// StorageClass objects via the REST API.
var Strategy = storageClassStrategy{api.Scheme, api.SimpleNameGenerator}

func (storageClassStrategy) NamespaceScoped() bool {
	return false
}

// ResetBeforeCreate clears the Status field which is not allowed to be set by end users on creation.
func (storageClassStrategy) PrepareForCreate(obj runtime.Object) {
	_ = obj.(*extensions.StorageClass)
}

func (storageClassStrategy) Validate(ctx api.Context, obj runtime.Object) field.ErrorList {
	storageClass := obj.(*extensions.StorageClass)
	return validation.ValidateStorageClass(storageClass)
}

// Canonicalize normalizes the object after validation.
func (storageClassStrategy) Canonicalize(obj runtime.Object) {
}

func (storageClassStrategy) AllowCreateOnUpdate() bool {
	return false
}

// PrepareForUpdate sets the Status fields which is not allowed to be set by an end user updating a PV
func (storageClassStrategy) PrepareForUpdate(obj, old runtime.Object) {
	_ = obj.(*extensions.StorageClass)
	_ = old.(*extensions.StorageClass)
}

func (storageClassStrategy) ValidateUpdate(ctx api.Context, obj, old runtime.Object) field.ErrorList {
	errorList := validation.ValidateStorageClass(obj.(*extensions.StorageClass))
	return append(errorList, validation.ValidateStorageClassUpdate(obj.(*extensions.StorageClass), old.(*extensions.StorageClass))...)
}

func (storageClassStrategy) AllowUnconditionalUpdate() bool {
	return true
}

// MatchStorageClass returns a generic matcher for a given label and field selector.
func MatchStorageClasses(label labels.Selector, field fields.Selector) generic.Matcher {
	return &generic.SelectionPredicate{
		Label: label,
		Field: field,
		GetAttrs: func(obj runtime.Object) (labels.Set, fields.Set, error) {
			cls, ok := obj.(*extensions.StorageClass)
			if !ok {
				return nil, nil, fmt.Errorf("given object is not of type StorageClass")
			}

			return labels.Set(cls.ObjectMeta.Labels), StorageClassToSelectableFields(cls), nil
		},
	}
}

// StorageClassToSelectableFields returns a label set that represents the object
func StorageClassToSelectableFields(storageClass *extensions.StorageClass) fields.Set {
	return generic.ObjectMetaFieldsSet(storageClass.ObjectMeta, false)
}
