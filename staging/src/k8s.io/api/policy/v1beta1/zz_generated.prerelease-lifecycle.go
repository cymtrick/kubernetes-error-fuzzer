//go:build !ignore_autogenerated
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
func (in *Eviction) APILifecycleIntroduced() (major, minor int) {
	return 1, 5
}

// APILifecycleDeprecated is an autogenerated function, returning the release in which the API struct was or will be deprecated as int versions of major and minor for comparison.
// It is controlled by "k8s:prerelease-lifecycle-gen:deprecated" tags in types.go or  "k8s:prerelease-lifecycle-gen:introduced" plus three minor.
func (in *Eviction) APILifecycleDeprecated() (major, minor int) {
	return 1, 22
}

// APILifecycleRemoved is an autogenerated function, returning the release in which the API is no longer served as int versions of major and minor for comparison.
// It is controlled by "k8s:prerelease-lifecycle-gen:removed" tags in types.go or  "k8s:prerelease-lifecycle-gen:deprecated" plus three minor.
func (in *Eviction) APILifecycleRemoved() (major, minor int) {
	return 1, 25
}

// APILifecycleIntroduced is an autogenerated function, returning the release in which the API struct was introduced as int versions of major and minor for comparison.
// It is controlled by "k8s:prerelease-lifecycle-gen:introduced" tags in types.go.
func (in *PodDisruptionBudget) APILifecycleIntroduced() (major, minor int) {
	return 1, 5
}

// APILifecycleDeprecated is an autogenerated function, returning the release in which the API struct was or will be deprecated as int versions of major and minor for comparison.
// It is controlled by "k8s:prerelease-lifecycle-gen:deprecated" tags in types.go or  "k8s:prerelease-lifecycle-gen:introduced" plus three minor.
func (in *PodDisruptionBudget) APILifecycleDeprecated() (major, minor int) {
	return 1, 21
}

// APILifecycleReplacement is an autogenerated function, returning the group, version, and kind that should be used instead of this deprecated type.
// It is controlled by "k8s:prerelease-lifecycle-gen:replacement=<group>,<version>,<kind>" tags in types.go.
func (in *PodDisruptionBudget) APILifecycleReplacement() schema.GroupVersionKind {
	return schema.GroupVersionKind{Group: "policy", Version: "v1", Kind: "PodDisruptionBudget"}
}

// APILifecycleRemoved is an autogenerated function, returning the release in which the API is no longer served as int versions of major and minor for comparison.
// It is controlled by "k8s:prerelease-lifecycle-gen:removed" tags in types.go or  "k8s:prerelease-lifecycle-gen:deprecated" plus three minor.
func (in *PodDisruptionBudget) APILifecycleRemoved() (major, minor int) {
	return 1, 25
}

// APILifecycleIntroduced is an autogenerated function, returning the release in which the API struct was introduced as int versions of major and minor for comparison.
// It is controlled by "k8s:prerelease-lifecycle-gen:introduced" tags in types.go.
func (in *PodDisruptionBudgetList) APILifecycleIntroduced() (major, minor int) {
	return 1, 5
}

// APILifecycleDeprecated is an autogenerated function, returning the release in which the API struct was or will be deprecated as int versions of major and minor for comparison.
// It is controlled by "k8s:prerelease-lifecycle-gen:deprecated" tags in types.go or  "k8s:prerelease-lifecycle-gen:introduced" plus three minor.
func (in *PodDisruptionBudgetList) APILifecycleDeprecated() (major, minor int) {
	return 1, 21
}

// APILifecycleReplacement is an autogenerated function, returning the group, version, and kind that should be used instead of this deprecated type.
// It is controlled by "k8s:prerelease-lifecycle-gen:replacement=<group>,<version>,<kind>" tags in types.go.
func (in *PodDisruptionBudgetList) APILifecycleReplacement() schema.GroupVersionKind {
	return schema.GroupVersionKind{Group: "policy", Version: "v1", Kind: "PodDisruptionBudgetList"}
}

// APILifecycleRemoved is an autogenerated function, returning the release in which the API is no longer served as int versions of major and minor for comparison.
// It is controlled by "k8s:prerelease-lifecycle-gen:removed" tags in types.go or  "k8s:prerelease-lifecycle-gen:deprecated" plus three minor.
func (in *PodDisruptionBudgetList) APILifecycleRemoved() (major, minor int) {
	return 1, 25
}
