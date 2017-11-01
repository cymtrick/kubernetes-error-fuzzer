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

package v1beta1

import (
	internalinterfaces "k8s.io/client-go/informers/internalinterfaces"
)

// Interface provides access to all the informers in this group version.
type Interface interface {
	// DaemonSets returns a DaemonSetInformer.
	DaemonSets() DaemonSetInformer
	// Deployments returns a DeploymentInformer.
	Deployments() DeploymentInformer
	// Ingresses returns a IngressInformer.
	Ingresses() IngressInformer
	// PodSecurityPolicies returns a PodSecurityPolicyInformer.
	PodSecurityPolicies() PodSecurityPolicyInformer
	// ReplicaSets returns a ReplicaSetInformer.
	ReplicaSets() ReplicaSetInformer
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

// DaemonSets returns a DaemonSetInformer.
func (v *version) DaemonSets() DaemonSetInformer {
	return &daemonSetInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// Deployments returns a DeploymentInformer.
func (v *version) Deployments() DeploymentInformer {
	return &deploymentInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// Ingresses returns a IngressInformer.
func (v *version) Ingresses() IngressInformer {
	return &ingressInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// PodSecurityPolicies returns a PodSecurityPolicyInformer.
func (v *version) PodSecurityPolicies() PodSecurityPolicyInformer {
	return &podSecurityPolicyInformer{factory: v.factory, tweakListOptions: v.tweakListOptions}
}

// ReplicaSets returns a ReplicaSetInformer.
func (v *version) ReplicaSets() ReplicaSetInformer {
	return &replicaSetInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}
