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

// This file was automatically generated by informer-gen with arguments: --input-dirs=[k8s.io/kubernetes/pkg/api,k8s.io/kubernetes/pkg/api/v1,k8s.io/kubernetes/pkg/apis/abac,k8s.io/kubernetes/pkg/apis/abac/v0,k8s.io/kubernetes/pkg/apis/abac/v1beta1,k8s.io/kubernetes/pkg/apis/apps,k8s.io/kubernetes/pkg/apis/apps/v1beta1,k8s.io/kubernetes/pkg/apis/authentication,k8s.io/kubernetes/pkg/apis/authentication/v1beta1,k8s.io/kubernetes/pkg/apis/authorization,k8s.io/kubernetes/pkg/apis/authorization/v1beta1,k8s.io/kubernetes/pkg/apis/autoscaling,k8s.io/kubernetes/pkg/apis/autoscaling/v1,k8s.io/kubernetes/pkg/apis/batch,k8s.io/kubernetes/pkg/apis/batch/v1,k8s.io/kubernetes/pkg/apis/batch/v2alpha1,k8s.io/kubernetes/pkg/apis/certificates,k8s.io/kubernetes/pkg/apis/certificates/v1alpha1,k8s.io/kubernetes/pkg/apis/componentconfig,k8s.io/kubernetes/pkg/apis/componentconfig/v1alpha1,k8s.io/kubernetes/pkg/apis/extensions,k8s.io/kubernetes/pkg/apis/extensions/v1beta1,k8s.io/kubernetes/pkg/apis/imagepolicy,k8s.io/kubernetes/pkg/apis/imagepolicy/v1alpha1,k8s.io/kubernetes/pkg/apis/policy,k8s.io/kubernetes/pkg/apis/policy/v1beta1,k8s.io/kubernetes/pkg/apis/rbac,k8s.io/kubernetes/pkg/apis/rbac/v1alpha1,k8s.io/kubernetes/pkg/apis/storage,k8s.io/kubernetes/pkg/apis/storage/v1beta1] --internal-clientset-package=k8s.io/kubernetes/pkg/client/clientset_generated/internalclientset --listers-package=k8s.io/kubernetes/pkg/client/listers --versioned-clientset-package=k8s.io/kubernetes/pkg/client/clientset_generated/clientset

package internalversion

import (
	internalinterfaces "k8s.io/kubernetes/pkg/client/informers/informers_generated/internalinterfaces"
)

// Interface provides access to all the informers in this group version.
type Interface interface {
	// DaemonSets returns a DaemonSetInformer.
	DaemonSets() DaemonSetInformer
	// Deployments returns a DeploymentInformer.
	Deployments() DeploymentInformer
	// Ingresses returns a IngressInformer.
	Ingresses() IngressInformer
	// NetworkPolicies returns a NetworkPolicyInformer.
	NetworkPolicies() NetworkPolicyInformer
	// PodSecurityPolicies returns a PodSecurityPolicyInformer.
	PodSecurityPolicies() PodSecurityPolicyInformer
	// ReplicaSets returns a ReplicaSetInformer.
	ReplicaSets() ReplicaSetInformer
	// ThirdPartyResources returns a ThirdPartyResourceInformer.
	ThirdPartyResources() ThirdPartyResourceInformer
}

type version struct {
	internalinterfaces.SharedInformerFactory
}

// New returns a new Interface.
func New(f internalinterfaces.SharedInformerFactory) Interface {
	return &version{f}
}

// DaemonSets returns a DaemonSetInformer.
func (v *version) DaemonSets() DaemonSetInformer {
	return &daemonSetInformer{factory: v.SharedInformerFactory}
}

// Deployments returns a DeploymentInformer.
func (v *version) Deployments() DeploymentInformer {
	return &deploymentInformer{factory: v.SharedInformerFactory}
}

// Ingresses returns a IngressInformer.
func (v *version) Ingresses() IngressInformer {
	return &ingressInformer{factory: v.SharedInformerFactory}
}

// NetworkPolicies returns a NetworkPolicyInformer.
func (v *version) NetworkPolicies() NetworkPolicyInformer {
	return &networkPolicyInformer{factory: v.SharedInformerFactory}
}

// PodSecurityPolicies returns a PodSecurityPolicyInformer.
func (v *version) PodSecurityPolicies() PodSecurityPolicyInformer {
	return &podSecurityPolicyInformer{factory: v.SharedInformerFactory}
}

// ReplicaSets returns a ReplicaSetInformer.
func (v *version) ReplicaSets() ReplicaSetInformer {
	return &replicaSetInformer{factory: v.SharedInformerFactory}
}

// ThirdPartyResources returns a ThirdPartyResourceInformer.
func (v *version) ThirdPartyResources() ThirdPartyResourceInformer {
	return &thirdPartyResourceInformer{factory: v.SharedInformerFactory}
}
