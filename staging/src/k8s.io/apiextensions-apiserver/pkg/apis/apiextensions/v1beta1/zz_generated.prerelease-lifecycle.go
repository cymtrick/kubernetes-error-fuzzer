// +build !ignore_autogenerated

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

// Code generated by prerelease-lifecycle-gen. DO NOT EDIT.

package v1beta1

import (
	schema "k8s.io/apimachinery/pkg/runtime/schema"
)

// APILifecycleIntroduced is an autogenerated function, returning the release in which the API struct was introduced as int versions of major and minor for comparison.
// It is controlled by "k8s:prerelease-lifecycle-gen:introduced" tags in types.go.
func (in *ConversionReview) APILifecycleIntroduced() (major, minor int) {
	return 1, 13
}

// APILifecycleDeprecated is an autogenerated function, returning the release in which the API struct was or will be deprecated as int versions of major and minor for comparison.
// It is controlled by "k8s:prerelease-lifecycle-gen:deprecated" tags in types.go or  "k8s:prerelease-lifecycle-gen:introduced" plus three minor.
func (in *ConversionReview) APILifecycleDeprecated() (major, minor int) {
	return 1, 19
}

// APILifecycleReplacement is an autogenerated function, returning the group, version, and kind that should be used instead of this deprecated type.
// It is controlled by "k8s:prerelease-lifecycle-gen:replacement=<group>,<version>,<kind>" tags in types.go.
func (in *ConversionReview) APILifecycleReplacement() schema.GroupVersionKind {
	return schema.GroupVersionKind{Group: "apiextensions.k8s.io", Version: "v1", Kind: "ConversionReview"}
}

// APILifecycleRemoved is an autogenerated function, returning the release in which the API is no longer served as int versions of major and minor for comparison.
// It is controlled by "k8s:prerelease-lifecycle-gen:removed" tags in types.go or  "k8s:prerelease-lifecycle-gen:deprecated" plus three minor.
func (in *ConversionReview) APILifecycleRemoved() (major, minor int) {
	return 1, 22
}

// APILifecycleIntroduced is an autogenerated function, returning the release in which the API struct was introduced as int versions of major and minor for comparison.
// It is controlled by "k8s:prerelease-lifecycle-gen:introduced" tags in types.go.
func (in *CustomResourceDefinition) APILifecycleIntroduced() (major, minor int) {
	return 1, 7
}

// APILifecycleDeprecated is an autogenerated function, returning the release in which the API struct was or will be deprecated as int versions of major and minor for comparison.
// It is controlled by "k8s:prerelease-lifecycle-gen:deprecated" tags in types.go or  "k8s:prerelease-lifecycle-gen:introduced" plus three minor.
func (in *CustomResourceDefinition) APILifecycleDeprecated() (major, minor int) {
	return 1, 16
}

// APILifecycleReplacement is an autogenerated function, returning the group, version, and kind that should be used instead of this deprecated type.
// It is controlled by "k8s:prerelease-lifecycle-gen:replacement=<group>,<version>,<kind>" tags in types.go.
func (in *CustomResourceDefinition) APILifecycleReplacement() schema.GroupVersionKind {
	return schema.GroupVersionKind{Group: "apiextensions.k8s.io", Version: "v1", Kind: "CustomResourceDefinition"}
}

// APILifecycleRemoved is an autogenerated function, returning the release in which the API is no longer served as int versions of major and minor for comparison.
// It is controlled by "k8s:prerelease-lifecycle-gen:removed" tags in types.go or  "k8s:prerelease-lifecycle-gen:deprecated" plus three minor.
func (in *CustomResourceDefinition) APILifecycleRemoved() (major, minor int) {
	return 1, 22
}

// APILifecycleIntroduced is an autogenerated function, returning the release in which the API struct was introduced as int versions of major and minor for comparison.
// It is controlled by "k8s:prerelease-lifecycle-gen:introduced" tags in types.go.
func (in *CustomResourceDefinitionList) APILifecycleIntroduced() (major, minor int) {
	return 1, 7
}

// APILifecycleDeprecated is an autogenerated function, returning the release in which the API struct was or will be deprecated as int versions of major and minor for comparison.
// It is controlled by "k8s:prerelease-lifecycle-gen:deprecated" tags in types.go or  "k8s:prerelease-lifecycle-gen:introduced" plus three minor.
func (in *CustomResourceDefinitionList) APILifecycleDeprecated() (major, minor int) {
	return 1, 19
}

// APILifecycleReplacement is an autogenerated function, returning the group, version, and kind that should be used instead of this deprecated type.
// It is controlled by "k8s:prerelease-lifecycle-gen:replacement=<group>,<version>,<kind>" tags in types.go.
func (in *CustomResourceDefinitionList) APILifecycleReplacement() schema.GroupVersionKind {
	return schema.GroupVersionKind{Group: "apiextensions.k8s.io", Version: "v1", Kind: "CustomResourceDefinitionList"}
}

// APILifecycleRemoved is an autogenerated function, returning the release in which the API is no longer served as int versions of major and minor for comparison.
// It is controlled by "k8s:prerelease-lifecycle-gen:removed" tags in types.go or  "k8s:prerelease-lifecycle-gen:deprecated" plus three minor.
func (in *CustomResourceDefinitionList) APILifecycleRemoved() (major, minor int) {
	return 1, 22
}
