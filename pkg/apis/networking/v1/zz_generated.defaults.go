// +build !ignore_autogenerated

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

// This file was autogenerated by defaulter-gen. Do not edit it manually!

package v1

import (
	networkingv1 "k8s.io/api/networking/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// RegisterDefaults adds defaulters functions to the given scheme.
// Public to allow building arbitrary schemes.
// All generated defaulters are covering - they call all nested defaulters.
func RegisterDefaults(scheme *runtime.Scheme) error {
	scheme.AddTypeDefaultingFunc(&networkingv1.NetworkPolicy{}, func(obj interface{}) { SetObjectDefaults_NetworkPolicy(obj.(*networkingv1.NetworkPolicy)) })
	scheme.AddTypeDefaultingFunc(&networkingv1.NetworkPolicyList{}, func(obj interface{}) { SetObjectDefaults_NetworkPolicyList(obj.(*networkingv1.NetworkPolicyList)) })
	return nil
}

func SetObjectDefaults_NetworkPolicy(in *networkingv1.NetworkPolicy) {
	for i := range in.Spec.Ingress {
		a := &in.Spec.Ingress[i]
		for j := range a.Ports {
			b := &a.Ports[j]
			SetDefaults_NetworkPolicyPort(b)
		}
	}
}

func SetObjectDefaults_NetworkPolicyList(in *networkingv1.NetworkPolicyList) {
	for i := range in.Items {
		a := &in.Items[i]
		SetObjectDefaults_NetworkPolicy(a)
	}
}
