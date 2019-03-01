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

package v1

import (
	unsafe "unsafe"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	conversion "k8s.io/apimachinery/pkg/conversion"
	runtime "k8s.io/apimachinery/pkg/runtime"
	apiregistration "k8s.io/kube-aggregator/pkg/apis/apiregistration"
)

func init() {
	localSchemeBuilder.Register(RegisterConversions)
}

// RegisterConversions adds conversion functions to the given scheme.
// Public to allow building arbitrary schemes.
func RegisterConversions(s *runtime.Scheme) error {
	if err := s.AddGeneratedConversionFunc((*APIService)(nil), (*apiregistration.APIService)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1_APIService_To_apiregistration_APIService(a.(*APIService), b.(*apiregistration.APIService), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*apiregistration.APIService)(nil), (*APIService)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_apiregistration_APIService_To_v1_APIService(a.(*apiregistration.APIService), b.(*APIService), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*APIServiceCondition)(nil), (*apiregistration.APIServiceCondition)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1_APIServiceCondition_To_apiregistration_APIServiceCondition(a.(*APIServiceCondition), b.(*apiregistration.APIServiceCondition), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*apiregistration.APIServiceCondition)(nil), (*APIServiceCondition)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_apiregistration_APIServiceCondition_To_v1_APIServiceCondition(a.(*apiregistration.APIServiceCondition), b.(*APIServiceCondition), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*APIServiceList)(nil), (*apiregistration.APIServiceList)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1_APIServiceList_To_apiregistration_APIServiceList(a.(*APIServiceList), b.(*apiregistration.APIServiceList), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*apiregistration.APIServiceList)(nil), (*APIServiceList)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_apiregistration_APIServiceList_To_v1_APIServiceList(a.(*apiregistration.APIServiceList), b.(*APIServiceList), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*APIServiceSpec)(nil), (*apiregistration.APIServiceSpec)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1_APIServiceSpec_To_apiregistration_APIServiceSpec(a.(*APIServiceSpec), b.(*apiregistration.APIServiceSpec), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*apiregistration.APIServiceSpec)(nil), (*APIServiceSpec)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_apiregistration_APIServiceSpec_To_v1_APIServiceSpec(a.(*apiregistration.APIServiceSpec), b.(*APIServiceSpec), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*APIServiceStatus)(nil), (*apiregistration.APIServiceStatus)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1_APIServiceStatus_To_apiregistration_APIServiceStatus(a.(*APIServiceStatus), b.(*apiregistration.APIServiceStatus), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*apiregistration.APIServiceStatus)(nil), (*APIServiceStatus)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_apiregistration_APIServiceStatus_To_v1_APIServiceStatus(a.(*apiregistration.APIServiceStatus), b.(*APIServiceStatus), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*ServiceReference)(nil), (*apiregistration.ServiceReference)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1_ServiceReference_To_apiregistration_ServiceReference(a.(*ServiceReference), b.(*apiregistration.ServiceReference), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*apiregistration.ServiceReference)(nil), (*ServiceReference)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_apiregistration_ServiceReference_To_v1_ServiceReference(a.(*apiregistration.ServiceReference), b.(*ServiceReference), scope)
	}); err != nil {
		return err
	}
	return nil
}

func autoConvert_v1_APIService_To_apiregistration_APIService(in *APIService, out *apiregistration.APIService, s conversion.Scope) error {
	out.ObjectMeta = in.ObjectMeta
	if err := Convert_v1_APIServiceSpec_To_apiregistration_APIServiceSpec(&in.Spec, &out.Spec, s); err != nil {
		return err
	}
	if err := Convert_v1_APIServiceStatus_To_apiregistration_APIServiceStatus(&in.Status, &out.Status, s); err != nil {
		return err
	}
	return nil
}

// Convert_v1_APIService_To_apiregistration_APIService is an autogenerated conversion function.
func Convert_v1_APIService_To_apiregistration_APIService(in *APIService, out *apiregistration.APIService, s conversion.Scope) error {
	return autoConvert_v1_APIService_To_apiregistration_APIService(in, out, s)
}

func autoConvert_apiregistration_APIService_To_v1_APIService(in *apiregistration.APIService, out *APIService, s conversion.Scope) error {
	out.ObjectMeta = in.ObjectMeta
	if err := Convert_apiregistration_APIServiceSpec_To_v1_APIServiceSpec(&in.Spec, &out.Spec, s); err != nil {
		return err
	}
	if err := Convert_apiregistration_APIServiceStatus_To_v1_APIServiceStatus(&in.Status, &out.Status, s); err != nil {
		return err
	}
	return nil
}

// Convert_apiregistration_APIService_To_v1_APIService is an autogenerated conversion function.
func Convert_apiregistration_APIService_To_v1_APIService(in *apiregistration.APIService, out *APIService, s conversion.Scope) error {
	return autoConvert_apiregistration_APIService_To_v1_APIService(in, out, s)
}

func autoConvert_v1_APIServiceCondition_To_apiregistration_APIServiceCondition(in *APIServiceCondition, out *apiregistration.APIServiceCondition, s conversion.Scope) error {
	out.Type = apiregistration.APIServiceConditionType(in.Type)
	out.Status = apiregistration.ConditionStatus(in.Status)
	out.LastTransitionTime = in.LastTransitionTime
	out.Reason = in.Reason
	out.Message = in.Message
	return nil
}

// Convert_v1_APIServiceCondition_To_apiregistration_APIServiceCondition is an autogenerated conversion function.
func Convert_v1_APIServiceCondition_To_apiregistration_APIServiceCondition(in *APIServiceCondition, out *apiregistration.APIServiceCondition, s conversion.Scope) error {
	return autoConvert_v1_APIServiceCondition_To_apiregistration_APIServiceCondition(in, out, s)
}

func autoConvert_apiregistration_APIServiceCondition_To_v1_APIServiceCondition(in *apiregistration.APIServiceCondition, out *APIServiceCondition, s conversion.Scope) error {
	out.Type = APIServiceConditionType(in.Type)
	out.Status = ConditionStatus(in.Status)
	out.LastTransitionTime = in.LastTransitionTime
	out.Reason = in.Reason
	out.Message = in.Message
	return nil
}

// Convert_apiregistration_APIServiceCondition_To_v1_APIServiceCondition is an autogenerated conversion function.
func Convert_apiregistration_APIServiceCondition_To_v1_APIServiceCondition(in *apiregistration.APIServiceCondition, out *APIServiceCondition, s conversion.Scope) error {
	return autoConvert_apiregistration_APIServiceCondition_To_v1_APIServiceCondition(in, out, s)
}

func autoConvert_v1_APIServiceList_To_apiregistration_APIServiceList(in *APIServiceList, out *apiregistration.APIServiceList, s conversion.Scope) error {
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]apiregistration.APIService, len(*in))
		for i := range *in {
			if err := Convert_v1_APIService_To_apiregistration_APIService(&(*in)[i], &(*out)[i], s); err != nil {
				return err
			}
		}
	} else {
		out.Items = nil
	}
	return nil
}

// Convert_v1_APIServiceList_To_apiregistration_APIServiceList is an autogenerated conversion function.
func Convert_v1_APIServiceList_To_apiregistration_APIServiceList(in *APIServiceList, out *apiregistration.APIServiceList, s conversion.Scope) error {
	return autoConvert_v1_APIServiceList_To_apiregistration_APIServiceList(in, out, s)
}

func autoConvert_apiregistration_APIServiceList_To_v1_APIServiceList(in *apiregistration.APIServiceList, out *APIServiceList, s conversion.Scope) error {
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]APIService, len(*in))
		for i := range *in {
			if err := Convert_apiregistration_APIService_To_v1_APIService(&(*in)[i], &(*out)[i], s); err != nil {
				return err
			}
		}
	} else {
		out.Items = nil
	}
	return nil
}

// Convert_apiregistration_APIServiceList_To_v1_APIServiceList is an autogenerated conversion function.
func Convert_apiregistration_APIServiceList_To_v1_APIServiceList(in *apiregistration.APIServiceList, out *APIServiceList, s conversion.Scope) error {
	return autoConvert_apiregistration_APIServiceList_To_v1_APIServiceList(in, out, s)
}

func autoConvert_v1_APIServiceSpec_To_apiregistration_APIServiceSpec(in *APIServiceSpec, out *apiregistration.APIServiceSpec, s conversion.Scope) error {
	if in.Service != nil {
		in, out := &in.Service, &out.Service
		*out = new(apiregistration.ServiceReference)
		if err := Convert_v1_ServiceReference_To_apiregistration_ServiceReference(*in, *out, s); err != nil {
			return err
		}
	} else {
		out.Service = nil
	}
	out.Group = in.Group
	out.Version = in.Version
	out.InsecureSkipTLSVerify = in.InsecureSkipTLSVerify
	out.CABundle = *(*[]byte)(unsafe.Pointer(&in.CABundle))
	out.GroupPriorityMinimum = in.GroupPriorityMinimum
	out.VersionPriority = in.VersionPriority
	return nil
}

// Convert_v1_APIServiceSpec_To_apiregistration_APIServiceSpec is an autogenerated conversion function.
func Convert_v1_APIServiceSpec_To_apiregistration_APIServiceSpec(in *APIServiceSpec, out *apiregistration.APIServiceSpec, s conversion.Scope) error {
	return autoConvert_v1_APIServiceSpec_To_apiregistration_APIServiceSpec(in, out, s)
}

func autoConvert_apiregistration_APIServiceSpec_To_v1_APIServiceSpec(in *apiregistration.APIServiceSpec, out *APIServiceSpec, s conversion.Scope) error {
	if in.Service != nil {
		in, out := &in.Service, &out.Service
		*out = new(ServiceReference)
		if err := Convert_apiregistration_ServiceReference_To_v1_ServiceReference(*in, *out, s); err != nil {
			return err
		}
	} else {
		out.Service = nil
	}
	out.Group = in.Group
	out.Version = in.Version
	out.InsecureSkipTLSVerify = in.InsecureSkipTLSVerify
	out.CABundle = *(*[]byte)(unsafe.Pointer(&in.CABundle))
	out.GroupPriorityMinimum = in.GroupPriorityMinimum
	out.VersionPriority = in.VersionPriority
	return nil
}

// Convert_apiregistration_APIServiceSpec_To_v1_APIServiceSpec is an autogenerated conversion function.
func Convert_apiregistration_APIServiceSpec_To_v1_APIServiceSpec(in *apiregistration.APIServiceSpec, out *APIServiceSpec, s conversion.Scope) error {
	return autoConvert_apiregistration_APIServiceSpec_To_v1_APIServiceSpec(in, out, s)
}

func autoConvert_v1_APIServiceStatus_To_apiregistration_APIServiceStatus(in *APIServiceStatus, out *apiregistration.APIServiceStatus, s conversion.Scope) error {
	out.Conditions = *(*[]apiregistration.APIServiceCondition)(unsafe.Pointer(&in.Conditions))
	return nil
}

// Convert_v1_APIServiceStatus_To_apiregistration_APIServiceStatus is an autogenerated conversion function.
func Convert_v1_APIServiceStatus_To_apiregistration_APIServiceStatus(in *APIServiceStatus, out *apiregistration.APIServiceStatus, s conversion.Scope) error {
	return autoConvert_v1_APIServiceStatus_To_apiregistration_APIServiceStatus(in, out, s)
}

func autoConvert_apiregistration_APIServiceStatus_To_v1_APIServiceStatus(in *apiregistration.APIServiceStatus, out *APIServiceStatus, s conversion.Scope) error {
	out.Conditions = *(*[]APIServiceCondition)(unsafe.Pointer(&in.Conditions))
	return nil
}

// Convert_apiregistration_APIServiceStatus_To_v1_APIServiceStatus is an autogenerated conversion function.
func Convert_apiregistration_APIServiceStatus_To_v1_APIServiceStatus(in *apiregistration.APIServiceStatus, out *APIServiceStatus, s conversion.Scope) error {
	return autoConvert_apiregistration_APIServiceStatus_To_v1_APIServiceStatus(in, out, s)
}

func autoConvert_v1_ServiceReference_To_apiregistration_ServiceReference(in *ServiceReference, out *apiregistration.ServiceReference, s conversion.Scope) error {
	out.Namespace = in.Namespace
	out.Name = in.Name
	if err := metav1.Convert_Pointer_int32_To_int32(&in.Port, &out.Port, s); err != nil {
		return err
	}
	return nil
}

// Convert_v1_ServiceReference_To_apiregistration_ServiceReference is an autogenerated conversion function.
func Convert_v1_ServiceReference_To_apiregistration_ServiceReference(in *ServiceReference, out *apiregistration.ServiceReference, s conversion.Scope) error {
	return autoConvert_v1_ServiceReference_To_apiregistration_ServiceReference(in, out, s)
}

func autoConvert_apiregistration_ServiceReference_To_v1_ServiceReference(in *apiregistration.ServiceReference, out *ServiceReference, s conversion.Scope) error {
	out.Namespace = in.Namespace
	out.Name = in.Name
	if err := metav1.Convert_int32_To_Pointer_int32(&in.Port, &out.Port, s); err != nil {
		return err
	}
	return nil
}

// Convert_apiregistration_ServiceReference_To_v1_ServiceReference is an autogenerated conversion function.
func Convert_apiregistration_ServiceReference_To_v1_ServiceReference(in *apiregistration.ServiceReference, out *ServiceReference, s conversion.Scope) error {
	return autoConvert_apiregistration_ServiceReference_To_v1_ServiceReference(in, out, s)
}
