/*
Copyright 2017 The Kubernetes Authors.

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

// This file was automatically generated by informer-gen

package internalversion

import (
	internalinterfaces "k8s.io/kubernetes/pkg/client/informers/informers_generated/internalversion/internalinterfaces"
)

// Interface provides access to all the informers in this group version.
type Interface interface {
	// ClusterRoles returns a ClusterRoleInformer.
	ClusterRoles() ClusterRoleInformer
	// ClusterRoleBindings returns a ClusterRoleBindingInformer.
	ClusterRoleBindings() ClusterRoleBindingInformer
	// Roles returns a RoleInformer.
	Roles() RoleInformer
	// RoleBindings returns a RoleBindingInformer.
	RoleBindings() RoleBindingInformer
}

type version struct {
	factory          internalinterfaces.SharedInformerFactory
	namespace        string
	tweakListOptions internalinterfaces.TweakListOptionsFunc
}

// New returns a new Interface.
func New(f internalinterfaces.SharedInformerFactory, namespace string, tweakListOptions internalinterfaces.TweakListOptionsFunc) Interface {
	return &version{factory: f, namespace: namespace, tweakListOptions: tweakListOptions}
}

// ClusterRoles returns a ClusterRoleInformer.
func (v *version) ClusterRoles() ClusterRoleInformer {
	return &clusterRoleInformer{factory: v.factory, tweakListOptions: v.tweakListOptions}
}

// ClusterRoleBindings returns a ClusterRoleBindingInformer.
func (v *version) ClusterRoleBindings() ClusterRoleBindingInformer {
	return &clusterRoleBindingInformer{factory: v.factory, tweakListOptions: v.tweakListOptions}
}

// Roles returns a RoleInformer.
func (v *version) Roles() RoleInformer {
	return &roleInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// RoleBindings returns a RoleBindingInformer.
func (v *version) RoleBindings() RoleBindingInformer {
	return &roleBindingInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}
