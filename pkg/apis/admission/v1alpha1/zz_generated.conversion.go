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

package v1alpha1

import (
	v1alpha1 "k8s.io/api/admission/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	conversion "k8s.io/apimachinery/pkg/conversion"
	runtime "k8s.io/apimachinery/pkg/runtime"
	types "k8s.io/apimachinery/pkg/types"
	admission "k8s.io/kubernetes/pkg/apis/admission"
	unsafe "unsafe"
)

func init() {
	localSchemeBuilder.Register(RegisterConversions)
}

// RegisterConversions adds conversion functions to the given scheme.
// Public to allow building arbitrary schemes.
func RegisterConversions(scheme *runtime.Scheme) error {
	return scheme.AddGeneratedConversionFuncs(
		Convert_v1alpha1_AdmissionRequest_To_admission_AdmissionRequest,
		Convert_admission_AdmissionRequest_To_v1alpha1_AdmissionRequest,
		Convert_v1alpha1_AdmissionResponse_To_admission_AdmissionResponse,
		Convert_admission_AdmissionResponse_To_v1alpha1_AdmissionResponse,
		Convert_v1alpha1_AdmissionReview_To_admission_AdmissionReview,
		Convert_admission_AdmissionReview_To_v1alpha1_AdmissionReview,
	)
}

func autoConvert_v1alpha1_AdmissionRequest_To_admission_AdmissionRequest(in *v1alpha1.AdmissionRequest, out *admission.AdmissionRequest, s conversion.Scope) error {
	out.UID = types.UID(in.UID)
	out.Kind = in.Kind
	out.Resource = in.Resource
	out.SubResource = in.SubResource
	out.Name = in.Name
	out.Namespace = in.Namespace
	out.Operation = admission.Operation(in.Operation)
	// TODO: Inefficient conversion - can we improve it?
	if err := s.Convert(&in.UserInfo, &out.UserInfo, 0); err != nil {
		return err
	}
	if err := runtime.Convert_runtime_RawExtension_To_runtime_Object(&in.Object, &out.Object, s); err != nil {
		return err
	}
	if err := runtime.Convert_runtime_RawExtension_To_runtime_Object(&in.OldObject, &out.OldObject, s); err != nil {
		return err
	}
	return nil
}

// Convert_v1alpha1_AdmissionRequest_To_admission_AdmissionRequest is an autogenerated conversion function.
func Convert_v1alpha1_AdmissionRequest_To_admission_AdmissionRequest(in *v1alpha1.AdmissionRequest, out *admission.AdmissionRequest, s conversion.Scope) error {
	return autoConvert_v1alpha1_AdmissionRequest_To_admission_AdmissionRequest(in, out, s)
}

func autoConvert_admission_AdmissionRequest_To_v1alpha1_AdmissionRequest(in *admission.AdmissionRequest, out *v1alpha1.AdmissionRequest, s conversion.Scope) error {
	out.UID = types.UID(in.UID)
	out.Kind = in.Kind
	out.Resource = in.Resource
	out.SubResource = in.SubResource
	out.Name = in.Name
	out.Namespace = in.Namespace
	out.Operation = v1alpha1.Operation(in.Operation)
	// TODO: Inefficient conversion - can we improve it?
	if err := s.Convert(&in.UserInfo, &out.UserInfo, 0); err != nil {
		return err
	}
	if err := runtime.Convert_runtime_Object_To_runtime_RawExtension(&in.Object, &out.Object, s); err != nil {
		return err
	}
	if err := runtime.Convert_runtime_Object_To_runtime_RawExtension(&in.OldObject, &out.OldObject, s); err != nil {
		return err
	}
	return nil
}

// Convert_admission_AdmissionRequest_To_v1alpha1_AdmissionRequest is an autogenerated conversion function.
func Convert_admission_AdmissionRequest_To_v1alpha1_AdmissionRequest(in *admission.AdmissionRequest, out *v1alpha1.AdmissionRequest, s conversion.Scope) error {
	return autoConvert_admission_AdmissionRequest_To_v1alpha1_AdmissionRequest(in, out, s)
}

func autoConvert_v1alpha1_AdmissionResponse_To_admission_AdmissionResponse(in *v1alpha1.AdmissionResponse, out *admission.AdmissionResponse, s conversion.Scope) error {
	out.UID = types.UID(in.UID)
	out.Allowed = in.Allowed
	out.Result = (*v1.Status)(unsafe.Pointer(in.Result))
	out.Patch = *(*[]byte)(unsafe.Pointer(&in.Patch))
	out.PatchType = (*admission.PatchType)(unsafe.Pointer(in.PatchType))
	return nil
}

// Convert_v1alpha1_AdmissionResponse_To_admission_AdmissionResponse is an autogenerated conversion function.
func Convert_v1alpha1_AdmissionResponse_To_admission_AdmissionResponse(in *v1alpha1.AdmissionResponse, out *admission.AdmissionResponse, s conversion.Scope) error {
	return autoConvert_v1alpha1_AdmissionResponse_To_admission_AdmissionResponse(in, out, s)
}

func autoConvert_admission_AdmissionResponse_To_v1alpha1_AdmissionResponse(in *admission.AdmissionResponse, out *v1alpha1.AdmissionResponse, s conversion.Scope) error {
	out.UID = types.UID(in.UID)
	out.Allowed = in.Allowed
	out.Result = (*v1.Status)(unsafe.Pointer(in.Result))
	out.Patch = *(*[]byte)(unsafe.Pointer(&in.Patch))
	out.PatchType = (*v1alpha1.PatchType)(unsafe.Pointer(in.PatchType))
	return nil
}

// Convert_admission_AdmissionResponse_To_v1alpha1_AdmissionResponse is an autogenerated conversion function.
func Convert_admission_AdmissionResponse_To_v1alpha1_AdmissionResponse(in *admission.AdmissionResponse, out *v1alpha1.AdmissionResponse, s conversion.Scope) error {
	return autoConvert_admission_AdmissionResponse_To_v1alpha1_AdmissionResponse(in, out, s)
}

func autoConvert_v1alpha1_AdmissionReview_To_admission_AdmissionReview(in *v1alpha1.AdmissionReview, out *admission.AdmissionReview, s conversion.Scope) error {
	if in.Request != nil {
		in, out := &in.Request, &out.Request
		*out = new(admission.AdmissionRequest)
		if err := Convert_v1alpha1_AdmissionRequest_To_admission_AdmissionRequest(*in, *out, s); err != nil {
			return err
		}
	} else {
		out.Request = nil
	}
	out.Response = (*admission.AdmissionResponse)(unsafe.Pointer(in.Response))
	return nil
}

// Convert_v1alpha1_AdmissionReview_To_admission_AdmissionReview is an autogenerated conversion function.
func Convert_v1alpha1_AdmissionReview_To_admission_AdmissionReview(in *v1alpha1.AdmissionReview, out *admission.AdmissionReview, s conversion.Scope) error {
	return autoConvert_v1alpha1_AdmissionReview_To_admission_AdmissionReview(in, out, s)
}

func autoConvert_admission_AdmissionReview_To_v1alpha1_AdmissionReview(in *admission.AdmissionReview, out *v1alpha1.AdmissionReview, s conversion.Scope) error {
	if in.Request != nil {
		in, out := &in.Request, &out.Request
		*out = new(v1alpha1.AdmissionRequest)
		if err := Convert_admission_AdmissionRequest_To_v1alpha1_AdmissionRequest(*in, *out, s); err != nil {
			return err
		}
	} else {
		out.Request = nil
	}
	out.Response = (*v1alpha1.AdmissionResponse)(unsafe.Pointer(in.Response))
	return nil
}

// Convert_admission_AdmissionReview_To_v1alpha1_AdmissionReview is an autogenerated conversion function.
func Convert_admission_AdmissionReview_To_v1alpha1_AdmissionReview(in *admission.AdmissionReview, out *v1alpha1.AdmissionReview, s conversion.Scope) error {
	return autoConvert_admission_AdmissionReview_To_v1alpha1_AdmissionReview(in, out, s)
}
