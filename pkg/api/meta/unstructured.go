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

package meta

import (
	"k8s.io/kubernetes/pkg/api/unversioned"
	"k8s.io/kubernetes/pkg/runtime"
)

// InterfacesForUnstructured returns VersionInterfaces suitable for
// dealing with runtime.Unstructured objects.
func InterfacesForUnstructured(unversioned.GroupVersion) (*VersionInterfaces, error) {
	return &VersionInterfaces{
		ObjectConvertor:  &runtime.UnstructuredObjectConverter{},
		MetadataAccessor: NewAccessor(),
	}, nil
}
