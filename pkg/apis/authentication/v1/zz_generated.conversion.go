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

	v1 "k8s.io/api/authentication/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	conversion "k8s.io/apimachinery/pkg/conversion"
	runtime "k8s.io/apimachinery/pkg/runtime"
	types "k8s.io/apimachinery/pkg/types"
	authentication "k8s.io/kubernetes/pkg/apis/authentication"
)

func init() {
	localSchemeBuilder.Register(RegisterConversions)
}

// RegisterConversions adds conversion functions to the given scheme.
// Public to allow building arbitrary schemes.
func RegisterConversions(scheme *runtime.Scheme) error {
	return scheme.AddGeneratedConversionFuncs(
		Convert_v1_BoundObjectReference_To_authentication_BoundObjectReference,
		Convert_authentication_BoundObjectReference_To_v1_BoundObjectReference,
		Convert_v1_TokenRequest_To_authentication_TokenRequest,
		Convert_authentication_TokenRequest_To_v1_TokenRequest,
		Convert_v1_TokenRequestSpec_To_authentication_TokenRequestSpec,
		Convert_authentication_TokenRequestSpec_To_v1_TokenRequestSpec,
		Convert_v1_TokenRequestStatus_To_authentication_TokenRequestStatus,
		Convert_authentication_TokenRequestStatus_To_v1_TokenRequestStatus,
		Convert_v1_TokenReview_To_authentication_TokenReview,
		Convert_authentication_TokenReview_To_v1_TokenReview,
		Convert_v1_TokenReviewSpec_To_authentication_TokenReviewSpec,
		Convert_authentication_TokenReviewSpec_To_v1_TokenReviewSpec,
		Convert_v1_TokenReviewStatus_To_authentication_TokenReviewStatus,
		Convert_authentication_TokenReviewStatus_To_v1_TokenReviewStatus,
		Convert_v1_UserInfo_To_authentication_UserInfo,
		Convert_authentication_UserInfo_To_v1_UserInfo,
	)
}

func autoConvert_v1_BoundObjectReference_To_authentication_BoundObjectReference(in *v1.BoundObjectReference, out *authentication.BoundObjectReference, s conversion.Scope) error {
	out.Kind = in.Kind
	out.APIVersion = in.APIVersion
	out.Name = in.Name
	out.UID = types.UID(in.UID)
	return nil
}

// Convert_v1_BoundObjectReference_To_authentication_BoundObjectReference is an autogenerated conversion function.
func Convert_v1_BoundObjectReference_To_authentication_BoundObjectReference(in *v1.BoundObjectReference, out *authentication.BoundObjectReference, s conversion.Scope) error {
	return autoConvert_v1_BoundObjectReference_To_authentication_BoundObjectReference(in, out, s)
}

func autoConvert_authentication_BoundObjectReference_To_v1_BoundObjectReference(in *authentication.BoundObjectReference, out *v1.BoundObjectReference, s conversion.Scope) error {
	out.Kind = in.Kind
	out.APIVersion = in.APIVersion
	out.Name = in.Name
	out.UID = types.UID(in.UID)
	return nil
}

// Convert_authentication_BoundObjectReference_To_v1_BoundObjectReference is an autogenerated conversion function.
func Convert_authentication_BoundObjectReference_To_v1_BoundObjectReference(in *authentication.BoundObjectReference, out *v1.BoundObjectReference, s conversion.Scope) error {
	return autoConvert_authentication_BoundObjectReference_To_v1_BoundObjectReference(in, out, s)
}

func autoConvert_v1_TokenRequest_To_authentication_TokenRequest(in *v1.TokenRequest, out *authentication.TokenRequest, s conversion.Scope) error {
	out.ObjectMeta = in.ObjectMeta
	if err := Convert_v1_TokenRequestSpec_To_authentication_TokenRequestSpec(&in.Spec, &out.Spec, s); err != nil {
		return err
	}
	if err := Convert_v1_TokenRequestStatus_To_authentication_TokenRequestStatus(&in.Status, &out.Status, s); err != nil {
		return err
	}
	return nil
}

// Convert_v1_TokenRequest_To_authentication_TokenRequest is an autogenerated conversion function.
func Convert_v1_TokenRequest_To_authentication_TokenRequest(in *v1.TokenRequest, out *authentication.TokenRequest, s conversion.Scope) error {
	return autoConvert_v1_TokenRequest_To_authentication_TokenRequest(in, out, s)
}

func autoConvert_authentication_TokenRequest_To_v1_TokenRequest(in *authentication.TokenRequest, out *v1.TokenRequest, s conversion.Scope) error {
	out.ObjectMeta = in.ObjectMeta
	if err := Convert_authentication_TokenRequestSpec_To_v1_TokenRequestSpec(&in.Spec, &out.Spec, s); err != nil {
		return err
	}
	if err := Convert_authentication_TokenRequestStatus_To_v1_TokenRequestStatus(&in.Status, &out.Status, s); err != nil {
		return err
	}
	return nil
}

// Convert_authentication_TokenRequest_To_v1_TokenRequest is an autogenerated conversion function.
func Convert_authentication_TokenRequest_To_v1_TokenRequest(in *authentication.TokenRequest, out *v1.TokenRequest, s conversion.Scope) error {
	return autoConvert_authentication_TokenRequest_To_v1_TokenRequest(in, out, s)
}

func autoConvert_v1_TokenRequestSpec_To_authentication_TokenRequestSpec(in *v1.TokenRequestSpec, out *authentication.TokenRequestSpec, s conversion.Scope) error {
	out.Audiences = *(*[]string)(unsafe.Pointer(&in.Audiences))
	if err := metav1.Convert_Pointer_int64_To_int64(&in.ExpirationSeconds, &out.ExpirationSeconds, s); err != nil {
		return err
	}
	out.BoundObjectRef = (*authentication.BoundObjectReference)(unsafe.Pointer(in.BoundObjectRef))
	return nil
}

// Convert_v1_TokenRequestSpec_To_authentication_TokenRequestSpec is an autogenerated conversion function.
func Convert_v1_TokenRequestSpec_To_authentication_TokenRequestSpec(in *v1.TokenRequestSpec, out *authentication.TokenRequestSpec, s conversion.Scope) error {
	return autoConvert_v1_TokenRequestSpec_To_authentication_TokenRequestSpec(in, out, s)
}

func autoConvert_authentication_TokenRequestSpec_To_v1_TokenRequestSpec(in *authentication.TokenRequestSpec, out *v1.TokenRequestSpec, s conversion.Scope) error {
	out.Audiences = *(*[]string)(unsafe.Pointer(&in.Audiences))
	if err := metav1.Convert_int64_To_Pointer_int64(&in.ExpirationSeconds, &out.ExpirationSeconds, s); err != nil {
		return err
	}
	out.BoundObjectRef = (*v1.BoundObjectReference)(unsafe.Pointer(in.BoundObjectRef))
	return nil
}

// Convert_authentication_TokenRequestSpec_To_v1_TokenRequestSpec is an autogenerated conversion function.
func Convert_authentication_TokenRequestSpec_To_v1_TokenRequestSpec(in *authentication.TokenRequestSpec, out *v1.TokenRequestSpec, s conversion.Scope) error {
	return autoConvert_authentication_TokenRequestSpec_To_v1_TokenRequestSpec(in, out, s)
}

func autoConvert_v1_TokenRequestStatus_To_authentication_TokenRequestStatus(in *v1.TokenRequestStatus, out *authentication.TokenRequestStatus, s conversion.Scope) error {
	out.Token = in.Token
	out.ExpirationTimestamp = in.ExpirationTimestamp
	return nil
}

// Convert_v1_TokenRequestStatus_To_authentication_TokenRequestStatus is an autogenerated conversion function.
func Convert_v1_TokenRequestStatus_To_authentication_TokenRequestStatus(in *v1.TokenRequestStatus, out *authentication.TokenRequestStatus, s conversion.Scope) error {
	return autoConvert_v1_TokenRequestStatus_To_authentication_TokenRequestStatus(in, out, s)
}

func autoConvert_authentication_TokenRequestStatus_To_v1_TokenRequestStatus(in *authentication.TokenRequestStatus, out *v1.TokenRequestStatus, s conversion.Scope) error {
	out.Token = in.Token
	out.ExpirationTimestamp = in.ExpirationTimestamp
	return nil
}

// Convert_authentication_TokenRequestStatus_To_v1_TokenRequestStatus is an autogenerated conversion function.
func Convert_authentication_TokenRequestStatus_To_v1_TokenRequestStatus(in *authentication.TokenRequestStatus, out *v1.TokenRequestStatus, s conversion.Scope) error {
	return autoConvert_authentication_TokenRequestStatus_To_v1_TokenRequestStatus(in, out, s)
}

func autoConvert_v1_TokenReview_To_authentication_TokenReview(in *v1.TokenReview, out *authentication.TokenReview, s conversion.Scope) error {
	out.ObjectMeta = in.ObjectMeta
	if err := Convert_v1_TokenReviewSpec_To_authentication_TokenReviewSpec(&in.Spec, &out.Spec, s); err != nil {
		return err
	}
	if err := Convert_v1_TokenReviewStatus_To_authentication_TokenReviewStatus(&in.Status, &out.Status, s); err != nil {
		return err
	}
	return nil
}

// Convert_v1_TokenReview_To_authentication_TokenReview is an autogenerated conversion function.
func Convert_v1_TokenReview_To_authentication_TokenReview(in *v1.TokenReview, out *authentication.TokenReview, s conversion.Scope) error {
	return autoConvert_v1_TokenReview_To_authentication_TokenReview(in, out, s)
}

func autoConvert_authentication_TokenReview_To_v1_TokenReview(in *authentication.TokenReview, out *v1.TokenReview, s conversion.Scope) error {
	out.ObjectMeta = in.ObjectMeta
	if err := Convert_authentication_TokenReviewSpec_To_v1_TokenReviewSpec(&in.Spec, &out.Spec, s); err != nil {
		return err
	}
	if err := Convert_authentication_TokenReviewStatus_To_v1_TokenReviewStatus(&in.Status, &out.Status, s); err != nil {
		return err
	}
	return nil
}

// Convert_authentication_TokenReview_To_v1_TokenReview is an autogenerated conversion function.
func Convert_authentication_TokenReview_To_v1_TokenReview(in *authentication.TokenReview, out *v1.TokenReview, s conversion.Scope) error {
	return autoConvert_authentication_TokenReview_To_v1_TokenReview(in, out, s)
}

func autoConvert_v1_TokenReviewSpec_To_authentication_TokenReviewSpec(in *v1.TokenReviewSpec, out *authentication.TokenReviewSpec, s conversion.Scope) error {
	out.Token = in.Token
	return nil
}

// Convert_v1_TokenReviewSpec_To_authentication_TokenReviewSpec is an autogenerated conversion function.
func Convert_v1_TokenReviewSpec_To_authentication_TokenReviewSpec(in *v1.TokenReviewSpec, out *authentication.TokenReviewSpec, s conversion.Scope) error {
	return autoConvert_v1_TokenReviewSpec_To_authentication_TokenReviewSpec(in, out, s)
}

func autoConvert_authentication_TokenReviewSpec_To_v1_TokenReviewSpec(in *authentication.TokenReviewSpec, out *v1.TokenReviewSpec, s conversion.Scope) error {
	out.Token = in.Token
	return nil
}

// Convert_authentication_TokenReviewSpec_To_v1_TokenReviewSpec is an autogenerated conversion function.
func Convert_authentication_TokenReviewSpec_To_v1_TokenReviewSpec(in *authentication.TokenReviewSpec, out *v1.TokenReviewSpec, s conversion.Scope) error {
	return autoConvert_authentication_TokenReviewSpec_To_v1_TokenReviewSpec(in, out, s)
}

func autoConvert_v1_TokenReviewStatus_To_authentication_TokenReviewStatus(in *v1.TokenReviewStatus, out *authentication.TokenReviewStatus, s conversion.Scope) error {
	out.Authenticated = in.Authenticated
	if err := Convert_v1_UserInfo_To_authentication_UserInfo(&in.User, &out.User, s); err != nil {
		return err
	}
	out.Error = in.Error
	return nil
}

// Convert_v1_TokenReviewStatus_To_authentication_TokenReviewStatus is an autogenerated conversion function.
func Convert_v1_TokenReviewStatus_To_authentication_TokenReviewStatus(in *v1.TokenReviewStatus, out *authentication.TokenReviewStatus, s conversion.Scope) error {
	return autoConvert_v1_TokenReviewStatus_To_authentication_TokenReviewStatus(in, out, s)
}

func autoConvert_authentication_TokenReviewStatus_To_v1_TokenReviewStatus(in *authentication.TokenReviewStatus, out *v1.TokenReviewStatus, s conversion.Scope) error {
	out.Authenticated = in.Authenticated
	if err := Convert_authentication_UserInfo_To_v1_UserInfo(&in.User, &out.User, s); err != nil {
		return err
	}
	out.Error = in.Error
	return nil
}

// Convert_authentication_TokenReviewStatus_To_v1_TokenReviewStatus is an autogenerated conversion function.
func Convert_authentication_TokenReviewStatus_To_v1_TokenReviewStatus(in *authentication.TokenReviewStatus, out *v1.TokenReviewStatus, s conversion.Scope) error {
	return autoConvert_authentication_TokenReviewStatus_To_v1_TokenReviewStatus(in, out, s)
}

func autoConvert_v1_UserInfo_To_authentication_UserInfo(in *v1.UserInfo, out *authentication.UserInfo, s conversion.Scope) error {
	out.Username = in.Username
	out.UID = in.UID
	out.Groups = *(*[]string)(unsafe.Pointer(&in.Groups))
	out.Extra = *(*map[string]authentication.ExtraValue)(unsafe.Pointer(&in.Extra))
	return nil
}

// Convert_v1_UserInfo_To_authentication_UserInfo is an autogenerated conversion function.
func Convert_v1_UserInfo_To_authentication_UserInfo(in *v1.UserInfo, out *authentication.UserInfo, s conversion.Scope) error {
	return autoConvert_v1_UserInfo_To_authentication_UserInfo(in, out, s)
}

func autoConvert_authentication_UserInfo_To_v1_UserInfo(in *authentication.UserInfo, out *v1.UserInfo, s conversion.Scope) error {
	out.Username = in.Username
	out.UID = in.UID
	out.Groups = *(*[]string)(unsafe.Pointer(&in.Groups))
	out.Extra = *(*map[string]v1.ExtraValue)(unsafe.Pointer(&in.Extra))
	return nil
}

// Convert_authentication_UserInfo_To_v1_UserInfo is an autogenerated conversion function.
func Convert_authentication_UserInfo_To_v1_UserInfo(in *authentication.UserInfo, out *v1.UserInfo, s conversion.Scope) error {
	return autoConvert_authentication_UserInfo_To_v1_UserInfo(in, out, s)
}
