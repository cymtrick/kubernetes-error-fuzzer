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
	unsafe "unsafe"

	authorizationv1 "k8s.io/api/authorization/v1"

	conversion "k8s.io/apimachinery/pkg/conversion"
	runtime "k8s.io/apimachinery/pkg/runtime"
	authorization "k8s.io/kubernetes/pkg/apis/authorization"
)

func init() {
	localSchemeBuilder.Register(RegisterConversions)
}

// RegisterConversions adds conversion functions to the given scheme.
// Public to allow building arbitrary schemes.
func RegisterConversions(scheme *runtime.Scheme) error {
	return scheme.AddGeneratedConversionFuncs(
		Convert_v1_LocalSubjectAccessReview_To_authorization_LocalSubjectAccessReview,
		Convert_authorization_LocalSubjectAccessReview_To_v1_LocalSubjectAccessReview,
		Convert_v1_NonResourceAttributes_To_authorization_NonResourceAttributes,
		Convert_authorization_NonResourceAttributes_To_v1_NonResourceAttributes,
		Convert_v1_ResourceAttributes_To_authorization_ResourceAttributes,
		Convert_authorization_ResourceAttributes_To_v1_ResourceAttributes,
		Convert_v1_SelfSubjectAccessReview_To_authorization_SelfSubjectAccessReview,
		Convert_authorization_SelfSubjectAccessReview_To_v1_SelfSubjectAccessReview,
		Convert_v1_SelfSubjectAccessReviewSpec_To_authorization_SelfSubjectAccessReviewSpec,
		Convert_authorization_SelfSubjectAccessReviewSpec_To_v1_SelfSubjectAccessReviewSpec,
		Convert_v1_SubjectAccessReview_To_authorization_SubjectAccessReview,
		Convert_authorization_SubjectAccessReview_To_v1_SubjectAccessReview,
		Convert_v1_SubjectAccessReviewSpec_To_authorization_SubjectAccessReviewSpec,
		Convert_authorization_SubjectAccessReviewSpec_To_v1_SubjectAccessReviewSpec,
		Convert_v1_SubjectAccessReviewStatus_To_authorization_SubjectAccessReviewStatus,
		Convert_authorization_SubjectAccessReviewStatus_To_v1_SubjectAccessReviewStatus,
	)
}

func autoConvert_v1_LocalSubjectAccessReview_To_authorization_LocalSubjectAccessReview(in *authorizationv1.LocalSubjectAccessReview, out *authorization.LocalSubjectAccessReview, s conversion.Scope) error {
	out.ObjectMeta = in.ObjectMeta
	if err := Convert_v1_SubjectAccessReviewSpec_To_authorization_SubjectAccessReviewSpec(&in.Spec, &out.Spec, s); err != nil {
		return err
	}
	if err := Convert_v1_SubjectAccessReviewStatus_To_authorization_SubjectAccessReviewStatus(&in.Status, &out.Status, s); err != nil {
		return err
	}
	return nil
}

// Convert_v1_LocalSubjectAccessReview_To_authorization_LocalSubjectAccessReview is an autogenerated conversion function.
func Convert_v1_LocalSubjectAccessReview_To_authorization_LocalSubjectAccessReview(in *authorizationv1.LocalSubjectAccessReview, out *authorization.LocalSubjectAccessReview, s conversion.Scope) error {
	return autoConvert_v1_LocalSubjectAccessReview_To_authorization_LocalSubjectAccessReview(in, out, s)
}

func autoConvert_authorization_LocalSubjectAccessReview_To_v1_LocalSubjectAccessReview(in *authorization.LocalSubjectAccessReview, out *authorizationv1.LocalSubjectAccessReview, s conversion.Scope) error {
	out.ObjectMeta = in.ObjectMeta
	if err := Convert_authorization_SubjectAccessReviewSpec_To_v1_SubjectAccessReviewSpec(&in.Spec, &out.Spec, s); err != nil {
		return err
	}
	if err := Convert_authorization_SubjectAccessReviewStatus_To_v1_SubjectAccessReviewStatus(&in.Status, &out.Status, s); err != nil {
		return err
	}
	return nil
}

// Convert_authorization_LocalSubjectAccessReview_To_v1_LocalSubjectAccessReview is an autogenerated conversion function.
func Convert_authorization_LocalSubjectAccessReview_To_v1_LocalSubjectAccessReview(in *authorization.LocalSubjectAccessReview, out *authorizationv1.LocalSubjectAccessReview, s conversion.Scope) error {
	return autoConvert_authorization_LocalSubjectAccessReview_To_v1_LocalSubjectAccessReview(in, out, s)
}

func autoConvert_v1_NonResourceAttributes_To_authorization_NonResourceAttributes(in *authorizationv1.NonResourceAttributes, out *authorization.NonResourceAttributes, s conversion.Scope) error {
	out.Path = in.Path
	out.Verb = in.Verb
	return nil
}

// Convert_v1_NonResourceAttributes_To_authorization_NonResourceAttributes is an autogenerated conversion function.
func Convert_v1_NonResourceAttributes_To_authorization_NonResourceAttributes(in *authorizationv1.NonResourceAttributes, out *authorization.NonResourceAttributes, s conversion.Scope) error {
	return autoConvert_v1_NonResourceAttributes_To_authorization_NonResourceAttributes(in, out, s)
}

func autoConvert_authorization_NonResourceAttributes_To_v1_NonResourceAttributes(in *authorization.NonResourceAttributes, out *authorizationv1.NonResourceAttributes, s conversion.Scope) error {
	out.Path = in.Path
	out.Verb = in.Verb
	return nil
}

// Convert_authorization_NonResourceAttributes_To_v1_NonResourceAttributes is an autogenerated conversion function.
func Convert_authorization_NonResourceAttributes_To_v1_NonResourceAttributes(in *authorization.NonResourceAttributes, out *authorizationv1.NonResourceAttributes, s conversion.Scope) error {
	return autoConvert_authorization_NonResourceAttributes_To_v1_NonResourceAttributes(in, out, s)
}

func autoConvert_v1_ResourceAttributes_To_authorization_ResourceAttributes(in *authorizationv1.ResourceAttributes, out *authorization.ResourceAttributes, s conversion.Scope) error {
	out.Namespace = in.Namespace
	out.Verb = in.Verb
	out.Group = in.Group
	out.Version = in.Version
	out.Resource = in.Resource
	out.Subresource = in.Subresource
	out.Name = in.Name
	return nil
}

// Convert_v1_ResourceAttributes_To_authorization_ResourceAttributes is an autogenerated conversion function.
func Convert_v1_ResourceAttributes_To_authorization_ResourceAttributes(in *authorizationv1.ResourceAttributes, out *authorization.ResourceAttributes, s conversion.Scope) error {
	return autoConvert_v1_ResourceAttributes_To_authorization_ResourceAttributes(in, out, s)
}

func autoConvert_authorization_ResourceAttributes_To_v1_ResourceAttributes(in *authorization.ResourceAttributes, out *authorizationv1.ResourceAttributes, s conversion.Scope) error {
	out.Namespace = in.Namespace
	out.Verb = in.Verb
	out.Group = in.Group
	out.Version = in.Version
	out.Resource = in.Resource
	out.Subresource = in.Subresource
	out.Name = in.Name
	return nil
}

// Convert_authorization_ResourceAttributes_To_v1_ResourceAttributes is an autogenerated conversion function.
func Convert_authorization_ResourceAttributes_To_v1_ResourceAttributes(in *authorization.ResourceAttributes, out *authorizationv1.ResourceAttributes, s conversion.Scope) error {
	return autoConvert_authorization_ResourceAttributes_To_v1_ResourceAttributes(in, out, s)
}

func autoConvert_v1_SelfSubjectAccessReview_To_authorization_SelfSubjectAccessReview(in *authorizationv1.SelfSubjectAccessReview, out *authorization.SelfSubjectAccessReview, s conversion.Scope) error {
	out.ObjectMeta = in.ObjectMeta
	if err := Convert_v1_SelfSubjectAccessReviewSpec_To_authorization_SelfSubjectAccessReviewSpec(&in.Spec, &out.Spec, s); err != nil {
		return err
	}
	if err := Convert_v1_SubjectAccessReviewStatus_To_authorization_SubjectAccessReviewStatus(&in.Status, &out.Status, s); err != nil {
		return err
	}
	return nil
}

// Convert_v1_SelfSubjectAccessReview_To_authorization_SelfSubjectAccessReview is an autogenerated conversion function.
func Convert_v1_SelfSubjectAccessReview_To_authorization_SelfSubjectAccessReview(in *authorizationv1.SelfSubjectAccessReview, out *authorization.SelfSubjectAccessReview, s conversion.Scope) error {
	return autoConvert_v1_SelfSubjectAccessReview_To_authorization_SelfSubjectAccessReview(in, out, s)
}

func autoConvert_authorization_SelfSubjectAccessReview_To_v1_SelfSubjectAccessReview(in *authorization.SelfSubjectAccessReview, out *authorizationv1.SelfSubjectAccessReview, s conversion.Scope) error {
	out.ObjectMeta = in.ObjectMeta
	if err := Convert_authorization_SelfSubjectAccessReviewSpec_To_v1_SelfSubjectAccessReviewSpec(&in.Spec, &out.Spec, s); err != nil {
		return err
	}
	if err := Convert_authorization_SubjectAccessReviewStatus_To_v1_SubjectAccessReviewStatus(&in.Status, &out.Status, s); err != nil {
		return err
	}
	return nil
}

// Convert_authorization_SelfSubjectAccessReview_To_v1_SelfSubjectAccessReview is an autogenerated conversion function.
func Convert_authorization_SelfSubjectAccessReview_To_v1_SelfSubjectAccessReview(in *authorization.SelfSubjectAccessReview, out *authorizationv1.SelfSubjectAccessReview, s conversion.Scope) error {
	return autoConvert_authorization_SelfSubjectAccessReview_To_v1_SelfSubjectAccessReview(in, out, s)
}

func autoConvert_v1_SelfSubjectAccessReviewSpec_To_authorization_SelfSubjectAccessReviewSpec(in *authorizationv1.SelfSubjectAccessReviewSpec, out *authorization.SelfSubjectAccessReviewSpec, s conversion.Scope) error {
	out.ResourceAttributes = (*authorization.ResourceAttributes)(unsafe.Pointer(in.ResourceAttributes))
	out.NonResourceAttributes = (*authorization.NonResourceAttributes)(unsafe.Pointer(in.NonResourceAttributes))
	return nil
}

// Convert_v1_SelfSubjectAccessReviewSpec_To_authorization_SelfSubjectAccessReviewSpec is an autogenerated conversion function.
func Convert_v1_SelfSubjectAccessReviewSpec_To_authorization_SelfSubjectAccessReviewSpec(in *authorizationv1.SelfSubjectAccessReviewSpec, out *authorization.SelfSubjectAccessReviewSpec, s conversion.Scope) error {
	return autoConvert_v1_SelfSubjectAccessReviewSpec_To_authorization_SelfSubjectAccessReviewSpec(in, out, s)
}

func autoConvert_authorization_SelfSubjectAccessReviewSpec_To_v1_SelfSubjectAccessReviewSpec(in *authorization.SelfSubjectAccessReviewSpec, out *authorizationv1.SelfSubjectAccessReviewSpec, s conversion.Scope) error {
	out.ResourceAttributes = (*authorizationv1.ResourceAttributes)(unsafe.Pointer(in.ResourceAttributes))
	out.NonResourceAttributes = (*authorizationv1.NonResourceAttributes)(unsafe.Pointer(in.NonResourceAttributes))
	return nil
}

// Convert_authorization_SelfSubjectAccessReviewSpec_To_v1_SelfSubjectAccessReviewSpec is an autogenerated conversion function.
func Convert_authorization_SelfSubjectAccessReviewSpec_To_v1_SelfSubjectAccessReviewSpec(in *authorization.SelfSubjectAccessReviewSpec, out *authorizationv1.SelfSubjectAccessReviewSpec, s conversion.Scope) error {
	return autoConvert_authorization_SelfSubjectAccessReviewSpec_To_v1_SelfSubjectAccessReviewSpec(in, out, s)
}

func autoConvert_v1_SubjectAccessReview_To_authorization_SubjectAccessReview(in *authorizationv1.SubjectAccessReview, out *authorization.SubjectAccessReview, s conversion.Scope) error {
	out.ObjectMeta = in.ObjectMeta
	if err := Convert_v1_SubjectAccessReviewSpec_To_authorization_SubjectAccessReviewSpec(&in.Spec, &out.Spec, s); err != nil {
		return err
	}
	if err := Convert_v1_SubjectAccessReviewStatus_To_authorization_SubjectAccessReviewStatus(&in.Status, &out.Status, s); err != nil {
		return err
	}
	return nil
}

// Convert_v1_SubjectAccessReview_To_authorization_SubjectAccessReview is an autogenerated conversion function.
func Convert_v1_SubjectAccessReview_To_authorization_SubjectAccessReview(in *authorizationv1.SubjectAccessReview, out *authorization.SubjectAccessReview, s conversion.Scope) error {
	return autoConvert_v1_SubjectAccessReview_To_authorization_SubjectAccessReview(in, out, s)
}

func autoConvert_authorization_SubjectAccessReview_To_v1_SubjectAccessReview(in *authorization.SubjectAccessReview, out *authorizationv1.SubjectAccessReview, s conversion.Scope) error {
	out.ObjectMeta = in.ObjectMeta
	if err := Convert_authorization_SubjectAccessReviewSpec_To_v1_SubjectAccessReviewSpec(&in.Spec, &out.Spec, s); err != nil {
		return err
	}
	if err := Convert_authorization_SubjectAccessReviewStatus_To_v1_SubjectAccessReviewStatus(&in.Status, &out.Status, s); err != nil {
		return err
	}
	return nil
}

// Convert_authorization_SubjectAccessReview_To_v1_SubjectAccessReview is an autogenerated conversion function.
func Convert_authorization_SubjectAccessReview_To_v1_SubjectAccessReview(in *authorization.SubjectAccessReview, out *authorizationv1.SubjectAccessReview, s conversion.Scope) error {
	return autoConvert_authorization_SubjectAccessReview_To_v1_SubjectAccessReview(in, out, s)
}

func autoConvert_v1_SubjectAccessReviewSpec_To_authorization_SubjectAccessReviewSpec(in *authorizationv1.SubjectAccessReviewSpec, out *authorization.SubjectAccessReviewSpec, s conversion.Scope) error {
	out.ResourceAttributes = (*authorization.ResourceAttributes)(unsafe.Pointer(in.ResourceAttributes))
	out.NonResourceAttributes = (*authorization.NonResourceAttributes)(unsafe.Pointer(in.NonResourceAttributes))
	out.User = in.User
	out.Groups = *(*[]string)(unsafe.Pointer(&in.Groups))
	out.Extra = *(*map[string]authorization.ExtraValue)(unsafe.Pointer(&in.Extra))
	return nil
}

// Convert_v1_SubjectAccessReviewSpec_To_authorization_SubjectAccessReviewSpec is an autogenerated conversion function.
func Convert_v1_SubjectAccessReviewSpec_To_authorization_SubjectAccessReviewSpec(in *authorizationv1.SubjectAccessReviewSpec, out *authorization.SubjectAccessReviewSpec, s conversion.Scope) error {
	return autoConvert_v1_SubjectAccessReviewSpec_To_authorization_SubjectAccessReviewSpec(in, out, s)
}

func autoConvert_authorization_SubjectAccessReviewSpec_To_v1_SubjectAccessReviewSpec(in *authorization.SubjectAccessReviewSpec, out *authorizationv1.SubjectAccessReviewSpec, s conversion.Scope) error {
	out.ResourceAttributes = (*authorizationv1.ResourceAttributes)(unsafe.Pointer(in.ResourceAttributes))
	out.NonResourceAttributes = (*authorizationv1.NonResourceAttributes)(unsafe.Pointer(in.NonResourceAttributes))
	out.User = in.User
	out.Groups = *(*[]string)(unsafe.Pointer(&in.Groups))
	out.Extra = *(*map[string]authorizationv1.ExtraValue)(unsafe.Pointer(&in.Extra))
	return nil
}

// Convert_authorization_SubjectAccessReviewSpec_To_v1_SubjectAccessReviewSpec is an autogenerated conversion function.
func Convert_authorization_SubjectAccessReviewSpec_To_v1_SubjectAccessReviewSpec(in *authorization.SubjectAccessReviewSpec, out *authorizationv1.SubjectAccessReviewSpec, s conversion.Scope) error {
	return autoConvert_authorization_SubjectAccessReviewSpec_To_v1_SubjectAccessReviewSpec(in, out, s)
}

func autoConvert_v1_SubjectAccessReviewStatus_To_authorization_SubjectAccessReviewStatus(in *authorizationv1.SubjectAccessReviewStatus, out *authorization.SubjectAccessReviewStatus, s conversion.Scope) error {
	out.Allowed = in.Allowed
	out.Reason = in.Reason
	out.EvaluationError = in.EvaluationError
	return nil
}

// Convert_v1_SubjectAccessReviewStatus_To_authorization_SubjectAccessReviewStatus is an autogenerated conversion function.
func Convert_v1_SubjectAccessReviewStatus_To_authorization_SubjectAccessReviewStatus(in *authorizationv1.SubjectAccessReviewStatus, out *authorization.SubjectAccessReviewStatus, s conversion.Scope) error {
	return autoConvert_v1_SubjectAccessReviewStatus_To_authorization_SubjectAccessReviewStatus(in, out, s)
}

func autoConvert_authorization_SubjectAccessReviewStatus_To_v1_SubjectAccessReviewStatus(in *authorization.SubjectAccessReviewStatus, out *authorizationv1.SubjectAccessReviewStatus, s conversion.Scope) error {
	out.Allowed = in.Allowed
	out.Reason = in.Reason
	out.EvaluationError = in.EvaluationError
	return nil
}

// Convert_authorization_SubjectAccessReviewStatus_To_v1_SubjectAccessReviewStatus is an autogenerated conversion function.
func Convert_authorization_SubjectAccessReviewStatus_To_v1_SubjectAccessReviewStatus(in *authorization.SubjectAccessReviewStatus, out *authorizationv1.SubjectAccessReviewStatus, s conversion.Scope) error {
	return autoConvert_authorization_SubjectAccessReviewStatus_To_v1_SubjectAccessReviewStatus(in, out, s)
}
