/*
Copyright 2024 The Kubernetes Authors.

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

package generic

import (
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/types"
)

type PolicyAccessor interface {
	GetName() string
	GetNamespace() string
	GetParamKind() *schema.GroupVersionKind
}

type BindingAccessor interface {
	GetName() string
	GetNamespace() string

	// GetPolicyName returns the name of the (Validating/Mutating)AdmissionPolicy,
	// which is cluster-scoped, so namespace is usually left blank.
	// But we leave the door open to add a namespaced vesion in the future
	GetPolicyName() types.NamespacedName
}
