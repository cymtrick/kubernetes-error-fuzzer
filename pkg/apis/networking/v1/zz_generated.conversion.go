// +build !ignore_autogenerated

/*
Copyright 2018 The Kubernetes Authors.

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

// Code generated by conversion-gen. DO NOT EDIT.

package v1

import (
	unsafe "unsafe"

	core_v1 "k8s.io/api/core/v1"
	v1 "k8s.io/api/networking/v1"
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	conversion "k8s.io/apimachinery/pkg/conversion"
	runtime "k8s.io/apimachinery/pkg/runtime"
	intstr "k8s.io/apimachinery/pkg/util/intstr"
	core "k8s.io/kubernetes/pkg/apis/core"
	networking "k8s.io/kubernetes/pkg/apis/networking"
)

func init() {
	localSchemeBuilder.Register(RegisterConversions)
}

// RegisterConversions adds conversion functions to the given scheme.
// Public to allow building arbitrary schemes.
func RegisterConversions(scheme *runtime.Scheme) error {
	return scheme.AddGeneratedConversionFuncs(
		Convert_v1_IPBlock_To_networking_IPBlock,
		Convert_networking_IPBlock_To_v1_IPBlock,
		Convert_v1_NetworkPolicy_To_networking_NetworkPolicy,
		Convert_networking_NetworkPolicy_To_v1_NetworkPolicy,
		Convert_v1_NetworkPolicyEgressRule_To_networking_NetworkPolicyEgressRule,
		Convert_networking_NetworkPolicyEgressRule_To_v1_NetworkPolicyEgressRule,
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

func autoConvert_v1_IPBlock_To_networking_IPBlock(in *v1.IPBlock, out *networking.IPBlock, s conversion.Scope) error {
	out.CIDR = in.CIDR
	out.Except = *(*[]string)(unsafe.Pointer(&in.Except))
	return nil
}

// Convert_v1_IPBlock_To_networking_IPBlock is an autogenerated conversion function.
func Convert_v1_IPBlock_To_networking_IPBlock(in *v1.IPBlock, out *networking.IPBlock, s conversion.Scope) error {
	return autoConvert_v1_IPBlock_To_networking_IPBlock(in, out, s)
}

func autoConvert_networking_IPBlock_To_v1_IPBlock(in *networking.IPBlock, out *v1.IPBlock, s conversion.Scope) error {
	out.CIDR = in.CIDR
	out.Except = *(*[]string)(unsafe.Pointer(&in.Except))
	return nil
}

// Convert_networking_IPBlock_To_v1_IPBlock is an autogenerated conversion function.
func Convert_networking_IPBlock_To_v1_IPBlock(in *networking.IPBlock, out *v1.IPBlock, s conversion.Scope) error {
	return autoConvert_networking_IPBlock_To_v1_IPBlock(in, out, s)
}

func autoConvert_v1_NetworkPolicy_To_networking_NetworkPolicy(in *v1.NetworkPolicy, out *networking.NetworkPolicy, s conversion.Scope) error {
	out.ObjectMeta = in.ObjectMeta
	if err := Convert_v1_NetworkPolicySpec_To_networking_NetworkPolicySpec(&in.Spec, &out.Spec, s); err != nil {
		return err
	}
	return nil
}

// Convert_v1_NetworkPolicy_To_networking_NetworkPolicy is an autogenerated conversion function.
func Convert_v1_NetworkPolicy_To_networking_NetworkPolicy(in *v1.NetworkPolicy, out *networking.NetworkPolicy, s conversion.Scope) error {
	return autoConvert_v1_NetworkPolicy_To_networking_NetworkPolicy(in, out, s)
}

func autoConvert_networking_NetworkPolicy_To_v1_NetworkPolicy(in *networking.NetworkPolicy, out *v1.NetworkPolicy, s conversion.Scope) error {
	out.ObjectMeta = in.ObjectMeta
	if err := Convert_networking_NetworkPolicySpec_To_v1_NetworkPolicySpec(&in.Spec, &out.Spec, s); err != nil {
		return err
	}
	return nil
}

// Convert_networking_NetworkPolicy_To_v1_NetworkPolicy is an autogenerated conversion function.
func Convert_networking_NetworkPolicy_To_v1_NetworkPolicy(in *networking.NetworkPolicy, out *v1.NetworkPolicy, s conversion.Scope) error {
	return autoConvert_networking_NetworkPolicy_To_v1_NetworkPolicy(in, out, s)
}

func autoConvert_v1_NetworkPolicyEgressRule_To_networking_NetworkPolicyEgressRule(in *v1.NetworkPolicyEgressRule, out *networking.NetworkPolicyEgressRule, s conversion.Scope) error {
	out.Ports = *(*[]networking.NetworkPolicyPort)(unsafe.Pointer(&in.Ports))
	out.To = *(*[]networking.NetworkPolicyPeer)(unsafe.Pointer(&in.To))
	return nil
}

// Convert_v1_NetworkPolicyEgressRule_To_networking_NetworkPolicyEgressRule is an autogenerated conversion function.
func Convert_v1_NetworkPolicyEgressRule_To_networking_NetworkPolicyEgressRule(in *v1.NetworkPolicyEgressRule, out *networking.NetworkPolicyEgressRule, s conversion.Scope) error {
	return autoConvert_v1_NetworkPolicyEgressRule_To_networking_NetworkPolicyEgressRule(in, out, s)
}

func autoConvert_networking_NetworkPolicyEgressRule_To_v1_NetworkPolicyEgressRule(in *networking.NetworkPolicyEgressRule, out *v1.NetworkPolicyEgressRule, s conversion.Scope) error {
	out.Ports = *(*[]v1.NetworkPolicyPort)(unsafe.Pointer(&in.Ports))
	out.To = *(*[]v1.NetworkPolicyPeer)(unsafe.Pointer(&in.To))
	return nil
}

// Convert_networking_NetworkPolicyEgressRule_To_v1_NetworkPolicyEgressRule is an autogenerated conversion function.
func Convert_networking_NetworkPolicyEgressRule_To_v1_NetworkPolicyEgressRule(in *networking.NetworkPolicyEgressRule, out *v1.NetworkPolicyEgressRule, s conversion.Scope) error {
	return autoConvert_networking_NetworkPolicyEgressRule_To_v1_NetworkPolicyEgressRule(in, out, s)
}

func autoConvert_v1_NetworkPolicyIngressRule_To_networking_NetworkPolicyIngressRule(in *v1.NetworkPolicyIngressRule, out *networking.NetworkPolicyIngressRule, s conversion.Scope) error {
	out.Ports = *(*[]networking.NetworkPolicyPort)(unsafe.Pointer(&in.Ports))
	out.From = *(*[]networking.NetworkPolicyPeer)(unsafe.Pointer(&in.From))
	return nil
}

// Convert_v1_NetworkPolicyIngressRule_To_networking_NetworkPolicyIngressRule is an autogenerated conversion function.
func Convert_v1_NetworkPolicyIngressRule_To_networking_NetworkPolicyIngressRule(in *v1.NetworkPolicyIngressRule, out *networking.NetworkPolicyIngressRule, s conversion.Scope) error {
	return autoConvert_v1_NetworkPolicyIngressRule_To_networking_NetworkPolicyIngressRule(in, out, s)
}

func autoConvert_networking_NetworkPolicyIngressRule_To_v1_NetworkPolicyIngressRule(in *networking.NetworkPolicyIngressRule, out *v1.NetworkPolicyIngressRule, s conversion.Scope) error {
	out.Ports = *(*[]v1.NetworkPolicyPort)(unsafe.Pointer(&in.Ports))
	out.From = *(*[]v1.NetworkPolicyPeer)(unsafe.Pointer(&in.From))
	return nil
}

// Convert_networking_NetworkPolicyIngressRule_To_v1_NetworkPolicyIngressRule is an autogenerated conversion function.
func Convert_networking_NetworkPolicyIngressRule_To_v1_NetworkPolicyIngressRule(in *networking.NetworkPolicyIngressRule, out *v1.NetworkPolicyIngressRule, s conversion.Scope) error {
	return autoConvert_networking_NetworkPolicyIngressRule_To_v1_NetworkPolicyIngressRule(in, out, s)
}

func autoConvert_v1_NetworkPolicyList_To_networking_NetworkPolicyList(in *v1.NetworkPolicyList, out *networking.NetworkPolicyList, s conversion.Scope) error {
	out.ListMeta = in.ListMeta
	out.Items = *(*[]networking.NetworkPolicy)(unsafe.Pointer(&in.Items))
	return nil
}

// Convert_v1_NetworkPolicyList_To_networking_NetworkPolicyList is an autogenerated conversion function.
func Convert_v1_NetworkPolicyList_To_networking_NetworkPolicyList(in *v1.NetworkPolicyList, out *networking.NetworkPolicyList, s conversion.Scope) error {
	return autoConvert_v1_NetworkPolicyList_To_networking_NetworkPolicyList(in, out, s)
}

func autoConvert_networking_NetworkPolicyList_To_v1_NetworkPolicyList(in *networking.NetworkPolicyList, out *v1.NetworkPolicyList, s conversion.Scope) error {
	out.ListMeta = in.ListMeta
	out.Items = *(*[]v1.NetworkPolicy)(unsafe.Pointer(&in.Items))
	return nil
}

// Convert_networking_NetworkPolicyList_To_v1_NetworkPolicyList is an autogenerated conversion function.
func Convert_networking_NetworkPolicyList_To_v1_NetworkPolicyList(in *networking.NetworkPolicyList, out *v1.NetworkPolicyList, s conversion.Scope) error {
	return autoConvert_networking_NetworkPolicyList_To_v1_NetworkPolicyList(in, out, s)
}

func autoConvert_v1_NetworkPolicyPeer_To_networking_NetworkPolicyPeer(in *v1.NetworkPolicyPeer, out *networking.NetworkPolicyPeer, s conversion.Scope) error {
	out.PodSelector = (*meta_v1.LabelSelector)(unsafe.Pointer(in.PodSelector))
	out.NamespaceSelector = (*meta_v1.LabelSelector)(unsafe.Pointer(in.NamespaceSelector))
	out.IPBlock = (*networking.IPBlock)(unsafe.Pointer(in.IPBlock))
	return nil
}

// Convert_v1_NetworkPolicyPeer_To_networking_NetworkPolicyPeer is an autogenerated conversion function.
func Convert_v1_NetworkPolicyPeer_To_networking_NetworkPolicyPeer(in *v1.NetworkPolicyPeer, out *networking.NetworkPolicyPeer, s conversion.Scope) error {
	return autoConvert_v1_NetworkPolicyPeer_To_networking_NetworkPolicyPeer(in, out, s)
}

func autoConvert_networking_NetworkPolicyPeer_To_v1_NetworkPolicyPeer(in *networking.NetworkPolicyPeer, out *v1.NetworkPolicyPeer, s conversion.Scope) error {
	out.PodSelector = (*meta_v1.LabelSelector)(unsafe.Pointer(in.PodSelector))
	out.NamespaceSelector = (*meta_v1.LabelSelector)(unsafe.Pointer(in.NamespaceSelector))
	out.IPBlock = (*v1.IPBlock)(unsafe.Pointer(in.IPBlock))
	return nil
}

// Convert_networking_NetworkPolicyPeer_To_v1_NetworkPolicyPeer is an autogenerated conversion function.
func Convert_networking_NetworkPolicyPeer_To_v1_NetworkPolicyPeer(in *networking.NetworkPolicyPeer, out *v1.NetworkPolicyPeer, s conversion.Scope) error {
	return autoConvert_networking_NetworkPolicyPeer_To_v1_NetworkPolicyPeer(in, out, s)
}

func autoConvert_v1_NetworkPolicyPort_To_networking_NetworkPolicyPort(in *v1.NetworkPolicyPort, out *networking.NetworkPolicyPort, s conversion.Scope) error {
	out.Protocol = (*core.Protocol)(unsafe.Pointer(in.Protocol))
	out.Port = (*intstr.IntOrString)(unsafe.Pointer(in.Port))
	return nil
}

// Convert_v1_NetworkPolicyPort_To_networking_NetworkPolicyPort is an autogenerated conversion function.
func Convert_v1_NetworkPolicyPort_To_networking_NetworkPolicyPort(in *v1.NetworkPolicyPort, out *networking.NetworkPolicyPort, s conversion.Scope) error {
	return autoConvert_v1_NetworkPolicyPort_To_networking_NetworkPolicyPort(in, out, s)
}

func autoConvert_networking_NetworkPolicyPort_To_v1_NetworkPolicyPort(in *networking.NetworkPolicyPort, out *v1.NetworkPolicyPort, s conversion.Scope) error {
	out.Protocol = (*core_v1.Protocol)(unsafe.Pointer(in.Protocol))
	out.Port = (*intstr.IntOrString)(unsafe.Pointer(in.Port))
	return nil
}

// Convert_networking_NetworkPolicyPort_To_v1_NetworkPolicyPort is an autogenerated conversion function.
func Convert_networking_NetworkPolicyPort_To_v1_NetworkPolicyPort(in *networking.NetworkPolicyPort, out *v1.NetworkPolicyPort, s conversion.Scope) error {
	return autoConvert_networking_NetworkPolicyPort_To_v1_NetworkPolicyPort(in, out, s)
}

func autoConvert_v1_NetworkPolicySpec_To_networking_NetworkPolicySpec(in *v1.NetworkPolicySpec, out *networking.NetworkPolicySpec, s conversion.Scope) error {
	out.PodSelector = in.PodSelector
	out.Ingress = *(*[]networking.NetworkPolicyIngressRule)(unsafe.Pointer(&in.Ingress))
	out.Egress = *(*[]networking.NetworkPolicyEgressRule)(unsafe.Pointer(&in.Egress))
	out.PolicyTypes = *(*[]networking.PolicyType)(unsafe.Pointer(&in.PolicyTypes))
	return nil
}

// Convert_v1_NetworkPolicySpec_To_networking_NetworkPolicySpec is an autogenerated conversion function.
func Convert_v1_NetworkPolicySpec_To_networking_NetworkPolicySpec(in *v1.NetworkPolicySpec, out *networking.NetworkPolicySpec, s conversion.Scope) error {
	return autoConvert_v1_NetworkPolicySpec_To_networking_NetworkPolicySpec(in, out, s)
}

func autoConvert_networking_NetworkPolicySpec_To_v1_NetworkPolicySpec(in *networking.NetworkPolicySpec, out *v1.NetworkPolicySpec, s conversion.Scope) error {
	out.PodSelector = in.PodSelector
	out.Ingress = *(*[]v1.NetworkPolicyIngressRule)(unsafe.Pointer(&in.Ingress))
	out.Egress = *(*[]v1.NetworkPolicyEgressRule)(unsafe.Pointer(&in.Egress))
	out.PolicyTypes = *(*[]v1.PolicyType)(unsafe.Pointer(&in.PolicyTypes))
	return nil
}

// Convert_networking_NetworkPolicySpec_To_v1_NetworkPolicySpec is an autogenerated conversion function.
func Convert_networking_NetworkPolicySpec_To_v1_NetworkPolicySpec(in *networking.NetworkPolicySpec, out *v1.NetworkPolicySpec, s conversion.Scope) error {
	return autoConvert_networking_NetworkPolicySpec_To_v1_NetworkPolicySpec(in, out, s)
}
