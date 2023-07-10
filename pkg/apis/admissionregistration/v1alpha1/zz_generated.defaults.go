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

// Code generated by defaulter-gen. DO NOT EDIT.

package v1alpha1

import (
	v1alpha1 "k8s.io/api/admissionregistration/v1alpha1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	v1 "k8s.io/kubernetes/pkg/apis/admissionregistration/v1"
)

// RegisterDefaults adds defaulters functions to the given scheme.
// Public to allow building arbitrary schemes.
// All generated defaulters are covering - they call all nested defaulters.
func RegisterDefaults(scheme *runtime.Scheme) error {
	scheme.AddTypeDefaultingFunc(&v1alpha1.ValidatingAdmissionPolicy{}, func(obj interface{}) {
		SetObjectDefaults_ValidatingAdmissionPolicy(obj.(*v1alpha1.ValidatingAdmissionPolicy))
	})
	scheme.AddTypeDefaultingFunc(&v1alpha1.ValidatingAdmissionPolicyBinding{}, func(obj interface{}) {
		SetObjectDefaults_ValidatingAdmissionPolicyBinding(obj.(*v1alpha1.ValidatingAdmissionPolicyBinding))
	})
	scheme.AddTypeDefaultingFunc(&v1alpha1.ValidatingAdmissionPolicyBindingList{}, func(obj interface{}) {
		SetObjectDefaults_ValidatingAdmissionPolicyBindingList(obj.(*v1alpha1.ValidatingAdmissionPolicyBindingList))
	})
	scheme.AddTypeDefaultingFunc(&v1alpha1.ValidatingAdmissionPolicyList{}, func(obj interface{}) {
		SetObjectDefaults_ValidatingAdmissionPolicyList(obj.(*v1alpha1.ValidatingAdmissionPolicyList))
	})
	return nil
}

func SetObjectDefaults_ValidatingAdmissionPolicy(in *v1alpha1.ValidatingAdmissionPolicy) {
	SetDefaults_ValidatingAdmissionPolicySpec(&in.Spec)
	if in.Spec.MatchConstraints != nil {
		SetDefaults_MatchResources(in.Spec.MatchConstraints)
		for i := range in.Spec.MatchConstraints.ResourceRules {
			a := &in.Spec.MatchConstraints.ResourceRules[i]
			v1.SetDefaults_Rule(&a.RuleWithOperations.Rule)
		}
		for i := range in.Spec.MatchConstraints.ExcludeResourceRules {
			a := &in.Spec.MatchConstraints.ExcludeResourceRules[i]
			v1.SetDefaults_Rule(&a.RuleWithOperations.Rule)
		}
	}
}

func SetObjectDefaults_ValidatingAdmissionPolicyBinding(in *v1alpha1.ValidatingAdmissionPolicyBinding) {
	if in.Spec.ParamRef != nil {
		SetDefaults_ParamRef(in.Spec.ParamRef)
	}
	if in.Spec.MatchResources != nil {
		SetDefaults_MatchResources(in.Spec.MatchResources)
		for i := range in.Spec.MatchResources.ResourceRules {
			a := &in.Spec.MatchResources.ResourceRules[i]
			v1.SetDefaults_Rule(&a.RuleWithOperations.Rule)
		}
		for i := range in.Spec.MatchResources.ExcludeResourceRules {
			a := &in.Spec.MatchResources.ExcludeResourceRules[i]
			v1.SetDefaults_Rule(&a.RuleWithOperations.Rule)
		}
	}
}

func SetObjectDefaults_ValidatingAdmissionPolicyBindingList(in *v1alpha1.ValidatingAdmissionPolicyBindingList) {
	for i := range in.Items {
		a := &in.Items[i]
		SetObjectDefaults_ValidatingAdmissionPolicyBinding(a)
	}
}

func SetObjectDefaults_ValidatingAdmissionPolicyList(in *v1alpha1.ValidatingAdmissionPolicyList) {
	for i := range in.Items {
		a := &in.Items[i]
		SetObjectDefaults_ValidatingAdmissionPolicy(a)
	}
}
