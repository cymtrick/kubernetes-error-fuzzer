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

	v1 "k8s.io/api/core/v1"
	v1alpha1 "k8s.io/api/discovery/v1alpha1"
	conversion "k8s.io/apimachinery/pkg/conversion"
	runtime "k8s.io/apimachinery/pkg/runtime"
	core "k8s.io/kubernetes/pkg/apis/core"
	discovery "k8s.io/kubernetes/pkg/apis/discovery"
)

func init() {
	localSchemeBuilder.Register(RegisterConversions)
}

// RegisterConversions adds conversion functions to the given scheme.
// Public to allow building arbitrary schemes.
func RegisterConversions(s *runtime.Scheme) error {
	if err := s.AddGeneratedConversionFunc((*v1alpha1.Endpoint)(nil), (*discovery.Endpoint)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha1_Endpoint_To_discovery_Endpoint(a.(*v1alpha1.Endpoint), b.(*discovery.Endpoint), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*discovery.Endpoint)(nil), (*v1alpha1.Endpoint)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_discovery_Endpoint_To_v1alpha1_Endpoint(a.(*discovery.Endpoint), b.(*v1alpha1.Endpoint), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*v1alpha1.EndpointConditions)(nil), (*discovery.EndpointConditions)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha1_EndpointConditions_To_discovery_EndpointConditions(a.(*v1alpha1.EndpointConditions), b.(*discovery.EndpointConditions), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*discovery.EndpointConditions)(nil), (*v1alpha1.EndpointConditions)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_discovery_EndpointConditions_To_v1alpha1_EndpointConditions(a.(*discovery.EndpointConditions), b.(*v1alpha1.EndpointConditions), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*v1alpha1.EndpointPort)(nil), (*discovery.EndpointPort)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha1_EndpointPort_To_discovery_EndpointPort(a.(*v1alpha1.EndpointPort), b.(*discovery.EndpointPort), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*discovery.EndpointPort)(nil), (*v1alpha1.EndpointPort)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_discovery_EndpointPort_To_v1alpha1_EndpointPort(a.(*discovery.EndpointPort), b.(*v1alpha1.EndpointPort), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*v1alpha1.EndpointSlice)(nil), (*discovery.EndpointSlice)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha1_EndpointSlice_To_discovery_EndpointSlice(a.(*v1alpha1.EndpointSlice), b.(*discovery.EndpointSlice), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*discovery.EndpointSlice)(nil), (*v1alpha1.EndpointSlice)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_discovery_EndpointSlice_To_v1alpha1_EndpointSlice(a.(*discovery.EndpointSlice), b.(*v1alpha1.EndpointSlice), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*v1alpha1.EndpointSliceList)(nil), (*discovery.EndpointSliceList)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha1_EndpointSliceList_To_discovery_EndpointSliceList(a.(*v1alpha1.EndpointSliceList), b.(*discovery.EndpointSliceList), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*discovery.EndpointSliceList)(nil), (*v1alpha1.EndpointSliceList)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_discovery_EndpointSliceList_To_v1alpha1_EndpointSliceList(a.(*discovery.EndpointSliceList), b.(*v1alpha1.EndpointSliceList), scope)
	}); err != nil {
		return err
	}
	return nil
}

func autoConvert_v1alpha1_Endpoint_To_discovery_Endpoint(in *v1alpha1.Endpoint, out *discovery.Endpoint, s conversion.Scope) error {
	out.Addresses = *(*[]string)(unsafe.Pointer(&in.Addresses))
	if err := Convert_v1alpha1_EndpointConditions_To_discovery_EndpointConditions(&in.Conditions, &out.Conditions, s); err != nil {
		return err
	}
	out.Hostname = (*string)(unsafe.Pointer(in.Hostname))
	out.TargetRef = (*core.ObjectReference)(unsafe.Pointer(in.TargetRef))
	out.Topology = *(*map[string]string)(unsafe.Pointer(&in.Topology))
	return nil
}

// Convert_v1alpha1_Endpoint_To_discovery_Endpoint is an autogenerated conversion function.
func Convert_v1alpha1_Endpoint_To_discovery_Endpoint(in *v1alpha1.Endpoint, out *discovery.Endpoint, s conversion.Scope) error {
	return autoConvert_v1alpha1_Endpoint_To_discovery_Endpoint(in, out, s)
}

func autoConvert_discovery_Endpoint_To_v1alpha1_Endpoint(in *discovery.Endpoint, out *v1alpha1.Endpoint, s conversion.Scope) error {
	out.Addresses = *(*[]string)(unsafe.Pointer(&in.Addresses))
	if err := Convert_discovery_EndpointConditions_To_v1alpha1_EndpointConditions(&in.Conditions, &out.Conditions, s); err != nil {
		return err
	}
	out.Hostname = (*string)(unsafe.Pointer(in.Hostname))
	out.TargetRef = (*v1.ObjectReference)(unsafe.Pointer(in.TargetRef))
	out.Topology = *(*map[string]string)(unsafe.Pointer(&in.Topology))
	return nil
}

// Convert_discovery_Endpoint_To_v1alpha1_Endpoint is an autogenerated conversion function.
func Convert_discovery_Endpoint_To_v1alpha1_Endpoint(in *discovery.Endpoint, out *v1alpha1.Endpoint, s conversion.Scope) error {
	return autoConvert_discovery_Endpoint_To_v1alpha1_Endpoint(in, out, s)
}

func autoConvert_v1alpha1_EndpointConditions_To_discovery_EndpointConditions(in *v1alpha1.EndpointConditions, out *discovery.EndpointConditions, s conversion.Scope) error {
	out.Ready = (*bool)(unsafe.Pointer(in.Ready))
	return nil
}

// Convert_v1alpha1_EndpointConditions_To_discovery_EndpointConditions is an autogenerated conversion function.
func Convert_v1alpha1_EndpointConditions_To_discovery_EndpointConditions(in *v1alpha1.EndpointConditions, out *discovery.EndpointConditions, s conversion.Scope) error {
	return autoConvert_v1alpha1_EndpointConditions_To_discovery_EndpointConditions(in, out, s)
}

func autoConvert_discovery_EndpointConditions_To_v1alpha1_EndpointConditions(in *discovery.EndpointConditions, out *v1alpha1.EndpointConditions, s conversion.Scope) error {
	out.Ready = (*bool)(unsafe.Pointer(in.Ready))
	return nil
}

// Convert_discovery_EndpointConditions_To_v1alpha1_EndpointConditions is an autogenerated conversion function.
func Convert_discovery_EndpointConditions_To_v1alpha1_EndpointConditions(in *discovery.EndpointConditions, out *v1alpha1.EndpointConditions, s conversion.Scope) error {
	return autoConvert_discovery_EndpointConditions_To_v1alpha1_EndpointConditions(in, out, s)
}

func autoConvert_v1alpha1_EndpointPort_To_discovery_EndpointPort(in *v1alpha1.EndpointPort, out *discovery.EndpointPort, s conversion.Scope) error {
	out.Name = (*string)(unsafe.Pointer(in.Name))
	out.Protocol = (*core.Protocol)(unsafe.Pointer(in.Protocol))
	out.Port = (*int32)(unsafe.Pointer(in.Port))
	out.AppProtocol = (*string)(unsafe.Pointer(in.AppProtocol))
	return nil
}

// Convert_v1alpha1_EndpointPort_To_discovery_EndpointPort is an autogenerated conversion function.
func Convert_v1alpha1_EndpointPort_To_discovery_EndpointPort(in *v1alpha1.EndpointPort, out *discovery.EndpointPort, s conversion.Scope) error {
	return autoConvert_v1alpha1_EndpointPort_To_discovery_EndpointPort(in, out, s)
}

func autoConvert_discovery_EndpointPort_To_v1alpha1_EndpointPort(in *discovery.EndpointPort, out *v1alpha1.EndpointPort, s conversion.Scope) error {
	out.Name = (*string)(unsafe.Pointer(in.Name))
	out.Protocol = (*v1.Protocol)(unsafe.Pointer(in.Protocol))
	out.Port = (*int32)(unsafe.Pointer(in.Port))
	out.AppProtocol = (*string)(unsafe.Pointer(in.AppProtocol))
	return nil
}

// Convert_discovery_EndpointPort_To_v1alpha1_EndpointPort is an autogenerated conversion function.
func Convert_discovery_EndpointPort_To_v1alpha1_EndpointPort(in *discovery.EndpointPort, out *v1alpha1.EndpointPort, s conversion.Scope) error {
	return autoConvert_discovery_EndpointPort_To_v1alpha1_EndpointPort(in, out, s)
}

func autoConvert_v1alpha1_EndpointSlice_To_discovery_EndpointSlice(in *v1alpha1.EndpointSlice, out *discovery.EndpointSlice, s conversion.Scope) error {
	out.ObjectMeta = in.ObjectMeta
	out.AddressType = (*discovery.AddressType)(unsafe.Pointer(in.AddressType))
	out.Endpoints = *(*[]discovery.Endpoint)(unsafe.Pointer(&in.Endpoints))
	out.Ports = *(*[]discovery.EndpointPort)(unsafe.Pointer(&in.Ports))
	return nil
}

// Convert_v1alpha1_EndpointSlice_To_discovery_EndpointSlice is an autogenerated conversion function.
func Convert_v1alpha1_EndpointSlice_To_discovery_EndpointSlice(in *v1alpha1.EndpointSlice, out *discovery.EndpointSlice, s conversion.Scope) error {
	return autoConvert_v1alpha1_EndpointSlice_To_discovery_EndpointSlice(in, out, s)
}

func autoConvert_discovery_EndpointSlice_To_v1alpha1_EndpointSlice(in *discovery.EndpointSlice, out *v1alpha1.EndpointSlice, s conversion.Scope) error {
	out.ObjectMeta = in.ObjectMeta
	out.AddressType = (*v1alpha1.AddressType)(unsafe.Pointer(in.AddressType))
	out.Endpoints = *(*[]v1alpha1.Endpoint)(unsafe.Pointer(&in.Endpoints))
	out.Ports = *(*[]v1alpha1.EndpointPort)(unsafe.Pointer(&in.Ports))
	return nil
}

// Convert_discovery_EndpointSlice_To_v1alpha1_EndpointSlice is an autogenerated conversion function.
func Convert_discovery_EndpointSlice_To_v1alpha1_EndpointSlice(in *discovery.EndpointSlice, out *v1alpha1.EndpointSlice, s conversion.Scope) error {
	return autoConvert_discovery_EndpointSlice_To_v1alpha1_EndpointSlice(in, out, s)
}

func autoConvert_v1alpha1_EndpointSliceList_To_discovery_EndpointSliceList(in *v1alpha1.EndpointSliceList, out *discovery.EndpointSliceList, s conversion.Scope) error {
	out.ListMeta = in.ListMeta
	out.Items = *(*[]discovery.EndpointSlice)(unsafe.Pointer(&in.Items))
	return nil
}

// Convert_v1alpha1_EndpointSliceList_To_discovery_EndpointSliceList is an autogenerated conversion function.
func Convert_v1alpha1_EndpointSliceList_To_discovery_EndpointSliceList(in *v1alpha1.EndpointSliceList, out *discovery.EndpointSliceList, s conversion.Scope) error {
	return autoConvert_v1alpha1_EndpointSliceList_To_discovery_EndpointSliceList(in, out, s)
}

func autoConvert_discovery_EndpointSliceList_To_v1alpha1_EndpointSliceList(in *discovery.EndpointSliceList, out *v1alpha1.EndpointSliceList, s conversion.Scope) error {
	out.ListMeta = in.ListMeta
	out.Items = *(*[]v1alpha1.EndpointSlice)(unsafe.Pointer(&in.Items))
	return nil
}

// Convert_discovery_EndpointSliceList_To_v1alpha1_EndpointSliceList is an autogenerated conversion function.
func Convert_discovery_EndpointSliceList_To_v1alpha1_EndpointSliceList(in *discovery.EndpointSliceList, out *v1alpha1.EndpointSliceList, s conversion.Scope) error {
	return autoConvert_discovery_EndpointSliceList_To_v1alpha1_EndpointSliceList(in, out, s)
}
