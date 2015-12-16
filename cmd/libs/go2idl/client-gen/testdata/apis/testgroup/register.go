/*
Copyright 2015 The Kubernetes Authors All rights reserved.

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

package testgroup

import (
	"k8s.io/kubernetes/pkg/api"
	"k8s.io/kubernetes/pkg/api/unversioned"
)

var SchemeGroupVersion = unversioned.GroupVersion{Group: "testgroup", Version: ""}

func init() {
	// Register the API.
	addKnownTypes()
}

// Adds the list of known types to api.Scheme.
func addKnownTypes() {
	api.Scheme.AddKnownTypes(SchemeGroupVersion,
		&TestType{},
		&TestTypeList{},
	)

	api.Scheme.AddKnownTypes(SchemeGroupVersion,
		&unversioned.ListOptions{})
}

func (obj *TestType) GetObjectKind() unversioned.ObjectKind     { return &obj.TypeMeta }
func (obj *TestTypeList) GetObjectKind() unversioned.ObjectKind { return &obj.TypeMeta }
