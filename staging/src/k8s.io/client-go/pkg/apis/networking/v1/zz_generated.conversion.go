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

// This file was autogenerated by conversion-gen. Do not edit it manually!

package v1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	conversion "k8s.io/apimachinery/pkg/conversion"
	runtime "k8s.io/apimachinery/pkg/runtime"
	intstr "k8s.io/apimachinery/pkg/util/intstr"
	api "k8s.io/client-go/pkg/api"
	api_v1 "k8s.io/client-go/pkg/api/v1"
	networking "k8s.io/client-go/pkg/apis/networking"
	unsafe "unsafe"
)

func init() {
	SchemeBuilder.Register(RegisterConversions)
}

// RegisterConversions adds conversion functions to the given scheme.
// Public to allow building arbitrary schemes.
func RegisterConversions(scheme *runtime.Scheme) error {
	return scheme.AddGeneratedConversionFuncs(
		Convert_v1_NetworkPolicy_To_networking_NetworkPolicy,
		Convert_networking_NetworkPolicy_To_v1_NetworkPolicy,
		Convert_v1_NetworkPolicyIngressRule_To_networking_NetworkPolicyIngressRule,
		Convert_networking_NetworkPolicyIngressRule_To_v1_NetworkPolicyIngressRule,
		Convert_v1_NetworkPolicyList_To_networking_NetworkPolicyList,
		Convert_networking_NetworkPolicyList_To_v1_NetworkPolicyList,
		Convert_v1_NetworkPolicyPeer_To_networking_NetworkPolicyPeer,
		Convert_networking_NetworkPolicyPeer_To_v1_NetworkPolicyPeer,
		Convert_v1_NetworkPolicyPort_To_networking_NetworkPolicyPort,
		Convert_networking_NetworkPolicyPort_To_v1_NetworkPolicyPort,
		Convert_v1_NetworkPolicySpec_To_networking_NetworkPolicySpec,
		Convert_networking_NetworkPolicySpec_To_v1_NetworkPolicySpec,
	)
}

func autoConvert_v1_NetworkPolicy_To_networking_NetworkPolicy(in *NetworkPolicy, out *networking.NetworkPolicy, s conversion.Scope) error {
	out.ObjectMeta = in.ObjectMeta
	if err := Convert_v1_NetworkPolicySpec_To_networking_NetworkPolicySpec(&in.Spec, &out.Spec, s); err != nil {
		return err
	}
	return nil
}

// Convert_v1_NetworkPolicy_To_networking_NetworkPolicy is an autogenerated conversion function.
func Convert_v1_NetworkPolicy_To_networking_NetworkPolicy(in *NetworkPolicy, out *networking.NetworkPolicy, s conversion.Scope) error {
	return autoConvert_v1_NetworkPolicy_To_networking_NetworkPolicy(in, out, s)
}

func autoConvert_networking_NetworkPolicy_To_v1_NetworkPolicy(in *networking.NetworkPolicy, out *NetworkPolicy, s conversion.Scope) error {
	out.ObjectMeta = in.ObjectMeta
	if err := Convert_networking_NetworkPolicySpec_To_v1_NetworkPolicySpec(&in.Spec, &out.Spec, s); err != nil {
		return err
	}
	return nil
}

// Convert_networking_NetworkPolicy_To_v1_NetworkPolicy is an autogenerated conversion function.
func Convert_networking_NetworkPolicy_To_v1_NetworkPolicy(in *networking.NetworkPolicy, out *NetworkPolicy, s conversion.Scope) error {
	return autoConvert_networking_NetworkPolicy_To_v1_NetworkPolicy(in, out, s)
}

func autoConvert_v1_NetworkPolicyIngressRule_To_networking_NetworkPolicyIngressRule(in *NetworkPolicyIngressRule, out *networking.NetworkPolicyIngressRule, s conversion.Scope) error {
	out.Ports = *(*[]networking.NetworkPolicyPort)(unsafe.Pointer(&in.Ports))
	out.From = *(*[]networking.NetworkPolicyPeer)(unsafe.Pointer(&in.From))
	return nil
}

// Convert_v1_NetworkPolicyIngressRule_To_networking_NetworkPolicyIngressRule is an autogenerated conversion function.
func Convert_v1_NetworkPolicyIngressRule_To_networking_NetworkPolicyIngressRule(in *NetworkPolicyIngressRule, out *networking.NetworkPolicyIngressRule, s conversion.Scope) error {
	return autoConvert_v1_NetworkPolicyIngressRule_To_networking_NetworkPolicyIngressRule(in, out, s)
}

func autoConvert_networking_NetworkPolicyIngressRule_To_v1_NetworkPolicyIngressRule(in *networking.NetworkPolicyIngressRule, out *NetworkPolicyIngressRule, s conversion.Scope) error {
	out.Ports = *(*[]NetworkPolicyPort)(unsafe.Pointer(&in.Ports))
	out.From = *(*[]NetworkPolicyPeer)(unsafe.Pointer(&in.From))
	return nil
}

// Convert_networking_NetworkPolicyIngressRule_To_v1_NetworkPolicyIngressRule is an autogenerated conversion function.
func Convert_networking_NetworkPolicyIngressRule_To_v1_NetworkPolicyIngressRule(in *networking.NetworkPolicyIngressRule, out *NetworkPolicyIngressRule, s conversion.Scope) error {
	return autoConvert_networking_NetworkPolicyIngressRule_To_v1_NetworkPolicyIngressRule(in, out, s)
}

func autoConvert_v1_NetworkPolicyList_To_networking_NetworkPolicyList(in *NetworkPolicyList, out *networking.NetworkPolicyList, s conversion.Scope) error {
	out.ListMeta = in.ListMeta
	out.Items = *(*[]networking.NetworkPolicy)(unsafe.Pointer(&in.Items))
	return nil
}

// Convert_v1_NetworkPolicyList_To_networking_NetworkPolicyList is an autogenerated conversion function.
func Convert_v1_NetworkPolicyList_To_networking_NetworkPolicyList(in *NetworkPolicyList, out *networking.NetworkPolicyList, s conversion.Scope) error {
	return autoConvert_v1_NetworkPolicyList_To_networking_NetworkPolicyList(in, out, s)
}

func autoConvert_networking_NetworkPolicyList_To_v1_NetworkPolicyList(in *networking.NetworkPolicyList, out *NetworkPolicyList, s conversion.Scope) error {
	out.ListMeta = in.ListMeta
	if in.Items == nil {
		out.Items = make([]NetworkPolicy, 0)
	} else {
		out.Items = *(*[]NetworkPolicy)(unsafe.Pointer(&in.Items))
	}
	return nil
}

// Convert_networking_NetworkPolicyList_To_v1_NetworkPolicyList is an autogenerated conversion function.
func Convert_networking_NetworkPolicyList_To_v1_NetworkPolicyList(in *networking.NetworkPolicyList, out *NetworkPolicyList, s conversion.Scope) error {
	return autoConvert_networking_NetworkPolicyList_To_v1_NetworkPolicyList(in, out, s)
}

func autoConvert_v1_NetworkPolicyPeer_To_networking_NetworkPolicyPeer(in *NetworkPolicyPeer, out *networking.NetworkPolicyPeer, s conversion.Scope) error {
	out.PodSelector = (*meta_v1.LabelSelector)(unsafe.Pointer(in.PodSelector))
	out.NamespaceSelector = (*meta_v1.LabelSelector)(unsafe.Pointer(in.NamespaceSelector))
	return nil
}

// Convert_v1_NetworkPolicyPeer_To_networking_NetworkPolicyPeer is an autogenerated conversion function.
func Convert_v1_NetworkPolicyPeer_To_networking_NetworkPolicyPeer(in *NetworkPolicyPeer, out *networking.NetworkPolicyPeer, s conversion.Scope) error {
	return autoConvert_v1_NetworkPolicyPeer_To_networking_NetworkPolicyPeer(in, out, s)
}

func autoConvert_networking_NetworkPolicyPeer_To_v1_NetworkPolicyPeer(in *networking.NetworkPolicyPeer, out *NetworkPolicyPeer, s conversion.Scope) error {
	out.PodSelector = (*meta_v1.LabelSelector)(unsafe.Pointer(in.PodSelector))
	out.NamespaceSelector = (*meta_v1.LabelSelector)(unsafe.Pointer(in.NamespaceSelector))
	return nil
}

// Convert_networking_NetworkPolicyPeer_To_v1_NetworkPolicyPeer is an autogenerated conversion function.
func Convert_networking_NetworkPolicyPeer_To_v1_NetworkPolicyPeer(in *networking.NetworkPolicyPeer, out *NetworkPolicyPeer, s conversion.Scope) error {
	return autoConvert_networking_NetworkPolicyPeer_To_v1_NetworkPolicyPeer(in, out, s)
}

func autoConvert_v1_NetworkPolicyPort_To_networking_NetworkPolicyPort(in *NetworkPolicyPort, out *networking.NetworkPolicyPort, s conversion.Scope) error {
	out.Protocol = (*api.Protocol)(unsafe.Pointer(in.Protocol))
	out.Port = (*intstr.IntOrString)(unsafe.Pointer(in.Port))
	return nil
}

// Convert_v1_NetworkPolicyPort_To_networking_NetworkPolicyPort is an autogenerated conversion function.
func Convert_v1_NetworkPolicyPort_To_networking_NetworkPolicyPort(in *NetworkPolicyPort, out *networking.NetworkPolicyPort, s conversion.Scope) error {
	return autoConvert_v1_NetworkPolicyPort_To_networking_NetworkPolicyPort(in, out, s)
}

func autoConvert_networking_NetworkPolicyPort_To_v1_NetworkPolicyPort(in *networking.NetworkPolicyPort, out *NetworkPolicyPort, s conversion.Scope) error {
	out.Protocol = (*api_v1.Protocol)(unsafe.Pointer(in.Protocol))
	out.Port = (*intstr.IntOrString)(unsafe.Pointer(in.Port))
	return nil
}

// Convert_networking_NetworkPolicyPort_To_v1_NetworkPolicyPort is an autogenerated conversion function.
func Convert_networking_NetworkPolicyPort_To_v1_NetworkPolicyPort(in *networking.NetworkPolicyPort, out *NetworkPolicyPort, s conversion.Scope) error {
	return autoConvert_networking_NetworkPolicyPort_To_v1_NetworkPolicyPort(in, out, s)
}

func autoConvert_v1_NetworkPolicySpec_To_networking_NetworkPolicySpec(in *NetworkPolicySpec, out *networking.NetworkPolicySpec, s conversion.Scope) error {
	out.PodSelector = in.PodSelector
	out.Ingress = *(*[]networking.NetworkPolicyIngressRule)(unsafe.Pointer(&in.Ingress))
	return nil
}

// Convert_v1_NetworkPolicySpec_To_networking_NetworkPolicySpec is an autogenerated conversion function.
func Convert_v1_NetworkPolicySpec_To_networking_NetworkPolicySpec(in *NetworkPolicySpec, out *networking.NetworkPolicySpec, s conversion.Scope) error {
	return autoConvert_v1_NetworkPolicySpec_To_networking_NetworkPolicySpec(in, out, s)
}

func autoConvert_networking_NetworkPolicySpec_To_v1_NetworkPolicySpec(in *networking.NetworkPolicySpec, out *NetworkPolicySpec, s conversion.Scope) error {
	out.PodSelector = in.PodSelector
	out.Ingress = *(*[]NetworkPolicyIngressRule)(unsafe.Pointer(&in.Ingress))
	return nil
}

// Convert_networking_NetworkPolicySpec_To_v1_NetworkPolicySpec is an autogenerated conversion function.
func Convert_networking_NetworkPolicySpec_To_v1_NetworkPolicySpec(in *networking.NetworkPolicySpec, out *NetworkPolicySpec, s conversion.Scope) error {
	return autoConvert_networking_NetworkPolicySpec_To_v1_NetworkPolicySpec(in, out, s)
}
