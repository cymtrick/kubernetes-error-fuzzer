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

// This file was autogenerated by conversion-gen. Do not edit it manually!

package v1beta1

import (
	unsafe "unsafe"

	conversion "k8s.io/apimachinery/pkg/conversion"
	runtime "k8s.io/apimachinery/pkg/runtime"
	apiregistration "k8s.io/kube-aggregator/pkg/apis/apiregistration"
)

func init() {
	localSchemeBuilder.Register(RegisterConversions)
}

// RegisterConversions adds conversion functions to the given scheme.
// Public to allow building arbitrary schemes.
func RegisterConversions(scheme *runtime.Scheme) error {
	return scheme.AddGeneratedConversionFuncs(
		Convert_v1beta1_APIService_To_apiregistration_APIService,
		Convert_apiregistration_APIService_To_v1beta1_APIService,
		Convert_v1beta1_APIServiceCondition_To_apiregistration_APIServiceCondition,
		Convert_apiregistration_APIServiceCondition_To_v1beta1_APIServiceCondition,
		Convert_v1beta1_APIServiceList_To_apiregistration_APIServiceList,
		Convert_apiregistration_APIServiceList_To_v1beta1_APIServiceList,
		Convert_v1beta1_APIServiceSpec_To_apiregistration_APIServiceSpec,
		Convert_apiregistration_APIServiceSpec_To_v1beta1_APIServiceSpec,
		Convert_v1beta1_APIServiceStatus_To_apiregistration_APIServiceStatus,
		Convert_apiregistration_APIServiceStatus_To_v1beta1_APIServiceStatus,
		Convert_v1beta1_ServiceReference_To_apiregistration_ServiceReference,
		Convert_apiregistration_ServiceReference_To_v1beta1_ServiceReference,
	)
}

func autoConvert_v1beta1_APIService_To_apiregistration_APIService(in *APIService, out *apiregistration.APIService, s conversion.Scope) error {
	out.ObjectMeta = in.ObjectMeta
	if err := Convert_v1beta1_APIServiceSpec_To_apiregistration_APIServiceSpec(&in.Spec, &out.Spec, s); err != nil {
		return err
	}
	if err := Convert_v1beta1_APIServiceStatus_To_apiregistration_APIServiceStatus(&in.Status, &out.Status, s); err != nil {
		return err
	}
	return nil
}

// Convert_v1beta1_APIService_To_apiregistration_APIService is an autogenerated conversion function.
func Convert_v1beta1_APIService_To_apiregistration_APIService(in *APIService, out *apiregistration.APIService, s conversion.Scope) error {
	return autoConvert_v1beta1_APIService_To_apiregistration_APIService(in, out, s)
}

func autoConvert_apiregistration_APIService_To_v1beta1_APIService(in *apiregistration.APIService, out *APIService, s conversion.Scope) error {
	out.ObjectMeta = in.ObjectMeta
	if err := Convert_apiregistration_APIServiceSpec_To_v1beta1_APIServiceSpec(&in.Spec, &out.Spec, s); err != nil {
		return err
	}
	if err := Convert_apiregistration_APIServiceStatus_To_v1beta1_APIServiceStatus(&in.Status, &out.Status, s); err != nil {
		return err
	}
	return nil
}

// Convert_apiregistration_APIService_To_v1beta1_APIService is an autogenerated conversion function.
func Convert_apiregistration_APIService_To_v1beta1_APIService(in *apiregistration.APIService, out *APIService, s conversion.Scope) error {
	return autoConvert_apiregistration_APIService_To_v1beta1_APIService(in, out, s)
}

func autoConvert_v1beta1_APIServiceCondition_To_apiregistration_APIServiceCondition(in *APIServiceCondition, out *apiregistration.APIServiceCondition, s conversion.Scope) error {
	out.Type = apiregistration.APIServiceConditionType(in.Type)
	out.Status = apiregistration.ConditionStatus(in.Status)
	out.LastTransitionTime = in.LastTransitionTime
	out.Reason = in.Reason
	out.Message = in.Message
	return nil
}

// Convert_v1beta1_APIServiceCondition_To_apiregistration_APIServiceCondition is an autogenerated conversion function.
func Convert_v1beta1_APIServiceCondition_To_apiregistration_APIServiceCondition(in *APIServiceCondition, out *apiregistration.APIServiceCondition, s conversion.Scope) error {
	return autoConvert_v1beta1_APIServiceCondition_To_apiregistration_APIServiceCondition(in, out, s)
}

func autoConvert_apiregistration_APIServiceCondition_To_v1beta1_APIServiceCondition(in *apiregistration.APIServiceCondition, out *APIServiceCondition, s conversion.Scope) error {
	out.Type = APIServiceConditionType(in.Type)
	out.Status = ConditionStatus(in.Status)
	out.LastTransitionTime = in.LastTransitionTime
	out.Reason = in.Reason
	out.Message = in.Message
	return nil
}

// Convert_apiregistration_APIServiceCondition_To_v1beta1_APIServiceCondition is an autogenerated conversion function.
func Convert_apiregistration_APIServiceCondition_To_v1beta1_APIServiceCondition(in *apiregistration.APIServiceCondition, out *APIServiceCondition, s conversion.Scope) error {
	return autoConvert_apiregistration_APIServiceCondition_To_v1beta1_APIServiceCondition(in, out, s)
}

func autoConvert_v1beta1_APIServiceList_To_apiregistration_APIServiceList(in *APIServiceList, out *apiregistration.APIServiceList, s conversion.Scope) error {
	out.ListMeta = in.ListMeta
	out.Items = *(*[]apiregistration.APIService)(unsafe.Pointer(&in.Items))
	return nil
}

// Convert_v1beta1_APIServiceList_To_apiregistration_APIServiceList is an autogenerated conversion function.
func Convert_v1beta1_APIServiceList_To_apiregistration_APIServiceList(in *APIServiceList, out *apiregistration.APIServiceList, s conversion.Scope) error {
	return autoConvert_v1beta1_APIServiceList_To_apiregistration_APIServiceList(in, out, s)
}

func autoConvert_apiregistration_APIServiceList_To_v1beta1_APIServiceList(in *apiregistration.APIServiceList, out *APIServiceList, s conversion.Scope) error {
	out.ListMeta = in.ListMeta
	out.Items = *(*[]APIService)(unsafe.Pointer(&in.Items))
	return nil
}

// Convert_apiregistration_APIServiceList_To_v1beta1_APIServiceList is an autogenerated conversion function.
func Convert_apiregistration_APIServiceList_To_v1beta1_APIServiceList(in *apiregistration.APIServiceList, out *APIServiceList, s conversion.Scope) error {
	return autoConvert_apiregistration_APIServiceList_To_v1beta1_APIServiceList(in, out, s)
}

func autoConvert_v1beta1_APIServiceSpec_To_apiregistration_APIServiceSpec(in *APIServiceSpec, out *apiregistration.APIServiceSpec, s conversion.Scope) error {
	out.Service = (*apiregistration.ServiceReference)(unsafe.Pointer(in.Service))
	out.Group = in.Group
	out.Version = in.Version
	out.InsecureSkipTLSVerify = in.InsecureSkipTLSVerify
	out.CABundle = *(*[]byte)(unsafe.Pointer(&in.CABundle))
	out.GroupPriorityMinimum = in.GroupPriorityMinimum
	out.VersionPriority = in.VersionPriority
	return nil
}

// Convert_v1beta1_APIServiceSpec_To_apiregistration_APIServiceSpec is an autogenerated conversion function.
func Convert_v1beta1_APIServiceSpec_To_apiregistration_APIServiceSpec(in *APIServiceSpec, out *apiregistration.APIServiceSpec, s conversion.Scope) error {
	return autoConvert_v1beta1_APIServiceSpec_To_apiregistration_APIServiceSpec(in, out, s)
}

func autoConvert_apiregistration_APIServiceSpec_To_v1beta1_APIServiceSpec(in *apiregistration.APIServiceSpec, out *APIServiceSpec, s conversion.Scope) error {
	out.Service = (*ServiceReference)(unsafe.Pointer(in.Service))
	out.Group = in.Group
	out.Version = in.Version
	out.InsecureSkipTLSVerify = in.InsecureSkipTLSVerify
	out.CABundle = *(*[]byte)(unsafe.Pointer(&in.CABundle))
	out.GroupPriorityMinimum = in.GroupPriorityMinimum
	out.VersionPriority = in.VersionPriority
	return nil
}

// Convert_apiregistration_APIServiceSpec_To_v1beta1_APIServiceSpec is an autogenerated conversion function.
func Convert_apiregistration_APIServiceSpec_To_v1beta1_APIServiceSpec(in *apiregistration.APIServiceSpec, out *APIServiceSpec, s conversion.Scope) error {
	return autoConvert_apiregistration_APIServiceSpec_To_v1beta1_APIServiceSpec(in, out, s)
}

func autoConvert_v1beta1_APIServiceStatus_To_apiregistration_APIServiceStatus(in *APIServiceStatus, out *apiregistration.APIServiceStatus, s conversion.Scope) error {
	out.Conditions = *(*[]apiregistration.APIServiceCondition)(unsafe.Pointer(&in.Conditions))
	return nil
}

// Convert_v1beta1_APIServiceStatus_To_apiregistration_APIServiceStatus is an autogenerated conversion function.
func Convert_v1beta1_APIServiceStatus_To_apiregistration_APIServiceStatus(in *APIServiceStatus, out *apiregistration.APIServiceStatus, s conversion.Scope) error {
	return autoConvert_v1beta1_APIServiceStatus_To_apiregistration_APIServiceStatus(in, out, s)
}

func autoConvert_apiregistration_APIServiceStatus_To_v1beta1_APIServiceStatus(in *apiregistration.APIServiceStatus, out *APIServiceStatus, s conversion.Scope) error {
	out.Conditions = *(*[]APIServiceCondition)(unsafe.Pointer(&in.Conditions))
	return nil
}

// Convert_apiregistration_APIServiceStatus_To_v1beta1_APIServiceStatus is an autogenerated conversion function.
func Convert_apiregistration_APIServiceStatus_To_v1beta1_APIServiceStatus(in *apiregistration.APIServiceStatus, out *APIServiceStatus, s conversion.Scope) error {
	return autoConvert_apiregistration_APIServiceStatus_To_v1beta1_APIServiceStatus(in, out, s)
}

func autoConvert_v1beta1_ServiceReference_To_apiregistration_ServiceReference(in *ServiceReference, out *apiregistration.ServiceReference, s conversion.Scope) error {
	out.Namespace = in.Namespace
	out.Name = in.Name
	return nil
}

// Convert_v1beta1_ServiceReference_To_apiregistration_ServiceReference is an autogenerated conversion function.
func Convert_v1beta1_ServiceReference_To_apiregistration_ServiceReference(in *ServiceReference, out *apiregistration.ServiceReference, s conversion.Scope) error {
	return autoConvert_v1beta1_ServiceReference_To_apiregistration_ServiceReference(in, out, s)
}

func autoConvert_apiregistration_ServiceReference_To_v1beta1_ServiceReference(in *apiregistration.ServiceReference, out *ServiceReference, s conversion.Scope) error {
	out.Namespace = in.Namespace
	out.Name = in.Name
	return nil
}

// Convert_apiregistration_ServiceReference_To_v1beta1_ServiceReference is an autogenerated conversion function.
func Convert_apiregistration_ServiceReference_To_v1beta1_ServiceReference(in *apiregistration.ServiceReference, out *ServiceReference, s conversion.Scope) error {
	return autoConvert_apiregistration_ServiceReference_To_v1beta1_ServiceReference(in, out, s)
}
