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

// Code generated by conversion-gen. DO NOT EDIT.

package v1alpha1

import (
	unsafe "unsafe"

	v1alpha1 "k8s.io/api/networking/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	conversion "k8s.io/apimachinery/pkg/conversion"
	runtime "k8s.io/apimachinery/pkg/runtime"
	networking "k8s.io/kubernetes/pkg/apis/networking"
)

func init() {
	localSchemeBuilder.Register(RegisterConversions)
}

// RegisterConversions adds conversion functions to the given scheme.
// Public to allow building arbitrary schemes.
func RegisterConversions(s *runtime.Scheme) error {
	if err := s.AddGeneratedConversionFunc((*v1alpha1.IPAddress)(nil), (*networking.IPAddress)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha1_IPAddress_To_networking_IPAddress(a.(*v1alpha1.IPAddress), b.(*networking.IPAddress), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*networking.IPAddress)(nil), (*v1alpha1.IPAddress)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_networking_IPAddress_To_v1alpha1_IPAddress(a.(*networking.IPAddress), b.(*v1alpha1.IPAddress), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*v1alpha1.IPAddressList)(nil), (*networking.IPAddressList)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha1_IPAddressList_To_networking_IPAddressList(a.(*v1alpha1.IPAddressList), b.(*networking.IPAddressList), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*networking.IPAddressList)(nil), (*v1alpha1.IPAddressList)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_networking_IPAddressList_To_v1alpha1_IPAddressList(a.(*networking.IPAddressList), b.(*v1alpha1.IPAddressList), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*v1alpha1.IPAddressSpec)(nil), (*networking.IPAddressSpec)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha1_IPAddressSpec_To_networking_IPAddressSpec(a.(*v1alpha1.IPAddressSpec), b.(*networking.IPAddressSpec), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*networking.IPAddressSpec)(nil), (*v1alpha1.IPAddressSpec)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_networking_IPAddressSpec_To_v1alpha1_IPAddressSpec(a.(*networking.IPAddressSpec), b.(*v1alpha1.IPAddressSpec), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*v1alpha1.ParentReference)(nil), (*networking.ParentReference)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha1_ParentReference_To_networking_ParentReference(a.(*v1alpha1.ParentReference), b.(*networking.ParentReference), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*networking.ParentReference)(nil), (*v1alpha1.ParentReference)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_networking_ParentReference_To_v1alpha1_ParentReference(a.(*networking.ParentReference), b.(*v1alpha1.ParentReference), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*v1alpha1.ServiceCIDR)(nil), (*networking.ServiceCIDR)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha1_ServiceCIDR_To_networking_ServiceCIDR(a.(*v1alpha1.ServiceCIDR), b.(*networking.ServiceCIDR), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*networking.ServiceCIDR)(nil), (*v1alpha1.ServiceCIDR)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_networking_ServiceCIDR_To_v1alpha1_ServiceCIDR(a.(*networking.ServiceCIDR), b.(*v1alpha1.ServiceCIDR), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*v1alpha1.ServiceCIDRList)(nil), (*networking.ServiceCIDRList)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha1_ServiceCIDRList_To_networking_ServiceCIDRList(a.(*v1alpha1.ServiceCIDRList), b.(*networking.ServiceCIDRList), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*networking.ServiceCIDRList)(nil), (*v1alpha1.ServiceCIDRList)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_networking_ServiceCIDRList_To_v1alpha1_ServiceCIDRList(a.(*networking.ServiceCIDRList), b.(*v1alpha1.ServiceCIDRList), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*v1alpha1.ServiceCIDRSpec)(nil), (*networking.ServiceCIDRSpec)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha1_ServiceCIDRSpec_To_networking_ServiceCIDRSpec(a.(*v1alpha1.ServiceCIDRSpec), b.(*networking.ServiceCIDRSpec), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*networking.ServiceCIDRSpec)(nil), (*v1alpha1.ServiceCIDRSpec)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_networking_ServiceCIDRSpec_To_v1alpha1_ServiceCIDRSpec(a.(*networking.ServiceCIDRSpec), b.(*v1alpha1.ServiceCIDRSpec), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*v1alpha1.ServiceCIDRStatus)(nil), (*networking.ServiceCIDRStatus)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha1_ServiceCIDRStatus_To_networking_ServiceCIDRStatus(a.(*v1alpha1.ServiceCIDRStatus), b.(*networking.ServiceCIDRStatus), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*networking.ServiceCIDRStatus)(nil), (*v1alpha1.ServiceCIDRStatus)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_networking_ServiceCIDRStatus_To_v1alpha1_ServiceCIDRStatus(a.(*networking.ServiceCIDRStatus), b.(*v1alpha1.ServiceCIDRStatus), scope)
	}); err != nil {
		return err
	}
	return nil
}

func autoConvert_v1alpha1_IPAddress_To_networking_IPAddress(in *v1alpha1.IPAddress, out *networking.IPAddress, s conversion.Scope) error {
	out.ObjectMeta = in.ObjectMeta
	if err := Convert_v1alpha1_IPAddressSpec_To_networking_IPAddressSpec(&in.Spec, &out.Spec, s); err != nil {
		return err
	}
	return nil
}

// Convert_v1alpha1_IPAddress_To_networking_IPAddress is an autogenerated conversion function.
func Convert_v1alpha1_IPAddress_To_networking_IPAddress(in *v1alpha1.IPAddress, out *networking.IPAddress, s conversion.Scope) error {
	return autoConvert_v1alpha1_IPAddress_To_networking_IPAddress(in, out, s)
}

func autoConvert_networking_IPAddress_To_v1alpha1_IPAddress(in *networking.IPAddress, out *v1alpha1.IPAddress, s conversion.Scope) error {
	out.ObjectMeta = in.ObjectMeta
	if err := Convert_networking_IPAddressSpec_To_v1alpha1_IPAddressSpec(&in.Spec, &out.Spec, s); err != nil {
		return err
	}
	return nil
}

// Convert_networking_IPAddress_To_v1alpha1_IPAddress is an autogenerated conversion function.
func Convert_networking_IPAddress_To_v1alpha1_IPAddress(in *networking.IPAddress, out *v1alpha1.IPAddress, s conversion.Scope) error {
	return autoConvert_networking_IPAddress_To_v1alpha1_IPAddress(in, out, s)
}

func autoConvert_v1alpha1_IPAddressList_To_networking_IPAddressList(in *v1alpha1.IPAddressList, out *networking.IPAddressList, s conversion.Scope) error {
	out.ListMeta = in.ListMeta
	out.Items = *(*[]networking.IPAddress)(unsafe.Pointer(&in.Items))
	return nil
}

// Convert_v1alpha1_IPAddressList_To_networking_IPAddressList is an autogenerated conversion function.
func Convert_v1alpha1_IPAddressList_To_networking_IPAddressList(in *v1alpha1.IPAddressList, out *networking.IPAddressList, s conversion.Scope) error {
	return autoConvert_v1alpha1_IPAddressList_To_networking_IPAddressList(in, out, s)
}

func autoConvert_networking_IPAddressList_To_v1alpha1_IPAddressList(in *networking.IPAddressList, out *v1alpha1.IPAddressList, s conversion.Scope) error {
	out.ListMeta = in.ListMeta
	out.Items = *(*[]v1alpha1.IPAddress)(unsafe.Pointer(&in.Items))
	return nil
}

// Convert_networking_IPAddressList_To_v1alpha1_IPAddressList is an autogenerated conversion function.
func Convert_networking_IPAddressList_To_v1alpha1_IPAddressList(in *networking.IPAddressList, out *v1alpha1.IPAddressList, s conversion.Scope) error {
	return autoConvert_networking_IPAddressList_To_v1alpha1_IPAddressList(in, out, s)
}

func autoConvert_v1alpha1_IPAddressSpec_To_networking_IPAddressSpec(in *v1alpha1.IPAddressSpec, out *networking.IPAddressSpec, s conversion.Scope) error {
	out.ParentRef = (*networking.ParentReference)(unsafe.Pointer(in.ParentRef))
	return nil
}

// Convert_v1alpha1_IPAddressSpec_To_networking_IPAddressSpec is an autogenerated conversion function.
func Convert_v1alpha1_IPAddressSpec_To_networking_IPAddressSpec(in *v1alpha1.IPAddressSpec, out *networking.IPAddressSpec, s conversion.Scope) error {
	return autoConvert_v1alpha1_IPAddressSpec_To_networking_IPAddressSpec(in, out, s)
}

func autoConvert_networking_IPAddressSpec_To_v1alpha1_IPAddressSpec(in *networking.IPAddressSpec, out *v1alpha1.IPAddressSpec, s conversion.Scope) error {
	out.ParentRef = (*v1alpha1.ParentReference)(unsafe.Pointer(in.ParentRef))
	return nil
}

// Convert_networking_IPAddressSpec_To_v1alpha1_IPAddressSpec is an autogenerated conversion function.
func Convert_networking_IPAddressSpec_To_v1alpha1_IPAddressSpec(in *networking.IPAddressSpec, out *v1alpha1.IPAddressSpec, s conversion.Scope) error {
	return autoConvert_networking_IPAddressSpec_To_v1alpha1_IPAddressSpec(in, out, s)
}

func autoConvert_v1alpha1_ParentReference_To_networking_ParentReference(in *v1alpha1.ParentReference, out *networking.ParentReference, s conversion.Scope) error {
	out.Group = in.Group
	out.Resource = in.Resource
	out.Namespace = in.Namespace
	out.Name = in.Name
	return nil
}

// Convert_v1alpha1_ParentReference_To_networking_ParentReference is an autogenerated conversion function.
func Convert_v1alpha1_ParentReference_To_networking_ParentReference(in *v1alpha1.ParentReference, out *networking.ParentReference, s conversion.Scope) error {
	return autoConvert_v1alpha1_ParentReference_To_networking_ParentReference(in, out, s)
}

func autoConvert_networking_ParentReference_To_v1alpha1_ParentReference(in *networking.ParentReference, out *v1alpha1.ParentReference, s conversion.Scope) error {
	out.Group = in.Group
	out.Resource = in.Resource
	out.Namespace = in.Namespace
	out.Name = in.Name
	return nil
}

// Convert_networking_ParentReference_To_v1alpha1_ParentReference is an autogenerated conversion function.
func Convert_networking_ParentReference_To_v1alpha1_ParentReference(in *networking.ParentReference, out *v1alpha1.ParentReference, s conversion.Scope) error {
	return autoConvert_networking_ParentReference_To_v1alpha1_ParentReference(in, out, s)
}

func autoConvert_v1alpha1_ServiceCIDR_To_networking_ServiceCIDR(in *v1alpha1.ServiceCIDR, out *networking.ServiceCIDR, s conversion.Scope) error {
	out.ObjectMeta = in.ObjectMeta
	if err := Convert_v1alpha1_ServiceCIDRSpec_To_networking_ServiceCIDRSpec(&in.Spec, &out.Spec, s); err != nil {
		return err
	}
	if err := Convert_v1alpha1_ServiceCIDRStatus_To_networking_ServiceCIDRStatus(&in.Status, &out.Status, s); err != nil {
		return err
	}
	return nil
}

// Convert_v1alpha1_ServiceCIDR_To_networking_ServiceCIDR is an autogenerated conversion function.
func Convert_v1alpha1_ServiceCIDR_To_networking_ServiceCIDR(in *v1alpha1.ServiceCIDR, out *networking.ServiceCIDR, s conversion.Scope) error {
	return autoConvert_v1alpha1_ServiceCIDR_To_networking_ServiceCIDR(in, out, s)
}

func autoConvert_networking_ServiceCIDR_To_v1alpha1_ServiceCIDR(in *networking.ServiceCIDR, out *v1alpha1.ServiceCIDR, s conversion.Scope) error {
	out.ObjectMeta = in.ObjectMeta
	if err := Convert_networking_ServiceCIDRSpec_To_v1alpha1_ServiceCIDRSpec(&in.Spec, &out.Spec, s); err != nil {
		return err
	}
	if err := Convert_networking_ServiceCIDRStatus_To_v1alpha1_ServiceCIDRStatus(&in.Status, &out.Status, s); err != nil {
		return err
	}
	return nil
}

// Convert_networking_ServiceCIDR_To_v1alpha1_ServiceCIDR is an autogenerated conversion function.
func Convert_networking_ServiceCIDR_To_v1alpha1_ServiceCIDR(in *networking.ServiceCIDR, out *v1alpha1.ServiceCIDR, s conversion.Scope) error {
	return autoConvert_networking_ServiceCIDR_To_v1alpha1_ServiceCIDR(in, out, s)
}

func autoConvert_v1alpha1_ServiceCIDRList_To_networking_ServiceCIDRList(in *v1alpha1.ServiceCIDRList, out *networking.ServiceCIDRList, s conversion.Scope) error {
	out.ListMeta = in.ListMeta
	out.Items = *(*[]networking.ServiceCIDR)(unsafe.Pointer(&in.Items))
	return nil
}

// Convert_v1alpha1_ServiceCIDRList_To_networking_ServiceCIDRList is an autogenerated conversion function.
func Convert_v1alpha1_ServiceCIDRList_To_networking_ServiceCIDRList(in *v1alpha1.ServiceCIDRList, out *networking.ServiceCIDRList, s conversion.Scope) error {
	return autoConvert_v1alpha1_ServiceCIDRList_To_networking_ServiceCIDRList(in, out, s)
}

func autoConvert_networking_ServiceCIDRList_To_v1alpha1_ServiceCIDRList(in *networking.ServiceCIDRList, out *v1alpha1.ServiceCIDRList, s conversion.Scope) error {
	out.ListMeta = in.ListMeta
	out.Items = *(*[]v1alpha1.ServiceCIDR)(unsafe.Pointer(&in.Items))
	return nil
}

// Convert_networking_ServiceCIDRList_To_v1alpha1_ServiceCIDRList is an autogenerated conversion function.
func Convert_networking_ServiceCIDRList_To_v1alpha1_ServiceCIDRList(in *networking.ServiceCIDRList, out *v1alpha1.ServiceCIDRList, s conversion.Scope) error {
	return autoConvert_networking_ServiceCIDRList_To_v1alpha1_ServiceCIDRList(in, out, s)
}

func autoConvert_v1alpha1_ServiceCIDRSpec_To_networking_ServiceCIDRSpec(in *v1alpha1.ServiceCIDRSpec, out *networking.ServiceCIDRSpec, s conversion.Scope) error {
	out.IPv4 = in.IPv4
	out.IPv6 = in.IPv6
	return nil
}

// Convert_v1alpha1_ServiceCIDRSpec_To_networking_ServiceCIDRSpec is an autogenerated conversion function.
func Convert_v1alpha1_ServiceCIDRSpec_To_networking_ServiceCIDRSpec(in *v1alpha1.ServiceCIDRSpec, out *networking.ServiceCIDRSpec, s conversion.Scope) error {
	return autoConvert_v1alpha1_ServiceCIDRSpec_To_networking_ServiceCIDRSpec(in, out, s)
}

func autoConvert_networking_ServiceCIDRSpec_To_v1alpha1_ServiceCIDRSpec(in *networking.ServiceCIDRSpec, out *v1alpha1.ServiceCIDRSpec, s conversion.Scope) error {
	out.IPv4 = in.IPv4
	out.IPv6 = in.IPv6
	return nil
}

// Convert_networking_ServiceCIDRSpec_To_v1alpha1_ServiceCIDRSpec is an autogenerated conversion function.
func Convert_networking_ServiceCIDRSpec_To_v1alpha1_ServiceCIDRSpec(in *networking.ServiceCIDRSpec, out *v1alpha1.ServiceCIDRSpec, s conversion.Scope) error {
	return autoConvert_networking_ServiceCIDRSpec_To_v1alpha1_ServiceCIDRSpec(in, out, s)
}

func autoConvert_v1alpha1_ServiceCIDRStatus_To_networking_ServiceCIDRStatus(in *v1alpha1.ServiceCIDRStatus, out *networking.ServiceCIDRStatus, s conversion.Scope) error {
	out.Conditions = *(*[]v1.Condition)(unsafe.Pointer(&in.Conditions))
	return nil
}

// Convert_v1alpha1_ServiceCIDRStatus_To_networking_ServiceCIDRStatus is an autogenerated conversion function.
func Convert_v1alpha1_ServiceCIDRStatus_To_networking_ServiceCIDRStatus(in *v1alpha1.ServiceCIDRStatus, out *networking.ServiceCIDRStatus, s conversion.Scope) error {
	return autoConvert_v1alpha1_ServiceCIDRStatus_To_networking_ServiceCIDRStatus(in, out, s)
}

func autoConvert_networking_ServiceCIDRStatus_To_v1alpha1_ServiceCIDRStatus(in *networking.ServiceCIDRStatus, out *v1alpha1.ServiceCIDRStatus, s conversion.Scope) error {
	out.Conditions = *(*[]v1.Condition)(unsafe.Pointer(&in.Conditions))
	return nil
}

// Convert_networking_ServiceCIDRStatus_To_v1alpha1_ServiceCIDRStatus is an autogenerated conversion function.
func Convert_networking_ServiceCIDRStatus_To_v1alpha1_ServiceCIDRStatus(in *networking.ServiceCIDRStatus, out *v1alpha1.ServiceCIDRStatus, s conversion.Scope) error {
	return autoConvert_networking_ServiceCIDRStatus_To_v1alpha1_ServiceCIDRStatus(in, out, s)
}
