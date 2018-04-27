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

	v1 "k8s.io/api/authorization/v1"
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
		Convert_v1_NonResourceRule_To_authorization_NonResourceRule,
		Convert_authorization_NonResourceRule_To_v1_NonResourceRule,
		Convert_v1_ResourceAttributes_To_authorization_ResourceAttributes,
		Convert_authorization_ResourceAttributes_To_v1_ResourceAttributes,
		Convert_v1_ResourceRule_To_authorization_ResourceRule,
		Convert_authorization_ResourceRule_To_v1_ResourceRule,
		Convert_v1_SelfSubjectAccessReview_To_authorization_SelfSubjectAccessReview,
		Convert_authorization_SelfSubjectAccessReview_To_v1_SelfSubjectAccessReview,
		Convert_v1_SelfSubjectAccessReviewSpec_To_authorization_SelfSubjectAccessReviewSpec,
		Convert_authorization_SelfSubjectAccessReviewSpec_To_v1_SelfSubjectAccessReviewSpec,
		Convert_v1_SelfSubjectRulesReview_To_authorization_SelfSubjectRulesReview,
		Convert_authorization_SelfSubjectRulesReview_To_v1_SelfSubjectRulesReview,
		Convert_v1_SelfSubjectRulesReviewSpec_To_authorization_SelfSubjectRulesReviewSpec,
		Convert_authorization_SelfSubjectRulesReviewSpec_To_v1_SelfSubjectRulesReviewSpec,
		Convert_v1_SubjectAccessReview_To_authorization_SubjectAccessReview,
		Convert_authorization_SubjectAccessReview_To_v1_SubjectAccessReview,
		Convert_v1_SubjectAccessReviewSpec_To_authorization_SubjectAccessReviewSpec,
		Convert_authorization_SubjectAccessReviewSpec_To_v1_SubjectAccessReviewSpec,
		Convert_v1_SubjectAccessReviewStatus_To_authorization_SubjectAccessReviewStatus,
		Convert_authorization_SubjectAccessReviewStatus_To_v1_SubjectAccessReviewStatus,
		Convert_v1_SubjectRulesReviewStatus_To_authorization_SubjectRulesReviewStatus,
		Convert_authorization_SubjectRulesReviewStatus_To_v1_SubjectRulesReviewStatus,
	)
}

func autoConvert_v1_LocalSubjectAccessReview_To_authorization_LocalSubjectAccessReview(in *v1.LocalSubjectAccessReview, out *authorization.LocalSubjectAccessReview, s conversion.Scope) error {
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
func Convert_v1_LocalSubjectAccessReview_To_authorization_LocalSubjectAccessReview(in *v1.LocalSubjectAccessReview, out *authorization.LocalSubjectAccessReview, s conversion.Scope) error {
	return autoConvert_v1_LocalSubjectAccessReview_To_authorization_LocalSubjectAccessReview(in, out, s)
}

func autoConvert_authorization_LocalSubjectAccessReview_To_v1_LocalSubjectAccessReview(in *authorization.LocalSubjectAccessReview, out *v1.LocalSubjectAccessReview, s conversion.Scope) error {
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
func Convert_authorization_LocalSubjectAccessReview_To_v1_LocalSubjectAccessReview(in *authorization.LocalSubjectAccessReview, out *v1.LocalSubjectAccessReview, s conversion.Scope) error {
	return autoConvert_authorization_LocalSubjectAccessReview_To_v1_LocalSubjectAccessReview(in, out, s)
}

func autoConvert_v1_NonResourceAttributes_To_authorization_NonResourceAttributes(in *v1.NonResourceAttributes, out *authorization.NonResourceAttributes, s conversion.Scope) error {
	out.Path = in.Path
	out.Verb = in.Verb
	return nil
}

// Convert_v1_NonResourceAttributes_To_authorization_NonResourceAttributes is an autogenerated conversion function.
func Convert_v1_NonResourceAttributes_To_authorization_NonResourceAttributes(in *v1.NonResourceAttributes, out *authorization.NonResourceAttributes, s conversion.Scope) error {
	return autoConvert_v1_NonResourceAttributes_To_authorization_NonResourceAttributes(in, out, s)
}

func autoConvert_authorization_NonResourceAttributes_To_v1_NonResourceAttributes(in *authorization.NonResourceAttributes, out *v1.NonResourceAttributes, s conversion.Scope) error {
	out.Path = in.Path
	out.Verb = in.Verb
	return nil
}

// Convert_authorization_NonResourceAttributes_To_v1_NonResourceAttributes is an autogenerated conversion function.
func Convert_authorization_NonResourceAttributes_To_v1_NonResourceAttributes(in *authorization.NonResourceAttributes, out *v1.NonResourceAttributes, s conversion.Scope) error {
	return autoConvert_authorization_NonResourceAttributes_To_v1_NonResourceAttributes(in, out, s)
}

func autoConvert_v1_NonResourceRule_To_authorization_NonResourceRule(in *v1.NonResourceRule, out *authorization.NonResourceRule, s conversion.Scope) error {
	out.Verbs = *(*[]string)(unsafe.Pointer(&in.Verbs))
	out.NonResourceURLs = *(*[]string)(unsafe.Pointer(&in.NonResourceURLs))
	return nil
}

// Convert_v1_NonResourceRule_To_authorization_NonResourceRule is an autogenerated conversion function.
func Convert_v1_NonResourceRule_To_authorization_NonResourceRule(in *v1.NonResourceRule, out *authorization.NonResourceRule, s conversion.Scope) error {
	return autoConvert_v1_NonResourceRule_To_authorization_NonResourceRule(in, out, s)
}

func autoConvert_authorization_NonResourceRule_To_v1_NonResourceRule(in *authorization.NonResourceRule, out *v1.NonResourceRule, s conversion.Scope) error {
	out.Verbs = *(*[]string)(unsafe.Pointer(&in.Verbs))
	out.NonResourceURLs = *(*[]string)(unsafe.Pointer(&in.NonResourceURLs))
	return nil
}

// Convert_authorization_NonResourceRule_To_v1_NonResourceRule is an autogenerated conversion function.
func Convert_authorization_NonResourceRule_To_v1_NonResourceRule(in *authorization.NonResourceRule, out *v1.NonResourceRule, s conversion.Scope) error {
	return autoConvert_authorization_NonResourceRule_To_v1_NonResourceRule(in, out, s)
}

func autoConvert_v1_ResourceAttributes_To_authorization_ResourceAttributes(in *v1.ResourceAttributes, out *authorization.ResourceAttributes, s conversion.Scope) error {
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
func Convert_v1_ResourceAttributes_To_authorization_ResourceAttributes(in *v1.ResourceAttributes, out *authorization.ResourceAttributes, s conversion.Scope) error {
	return autoConvert_v1_ResourceAttributes_To_authorization_ResourceAttributes(in, out, s)
}

func autoConvert_authorization_ResourceAttributes_To_v1_ResourceAttributes(in *authorization.ResourceAttributes, out *v1.ResourceAttributes, s conversion.Scope) error {
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
func Convert_authorization_ResourceAttributes_To_v1_ResourceAttributes(in *authorization.ResourceAttributes, out *v1.ResourceAttributes, s conversion.Scope) error {
	return autoConvert_authorization_ResourceAttributes_To_v1_ResourceAttributes(in, out, s)
}

func autoConvert_v1_ResourceRule_To_authorization_ResourceRule(in *v1.ResourceRule, out *authorization.ResourceRule, s conversion.Scope) error {
	out.Verbs = *(*[]string)(unsafe.Pointer(&in.Verbs))
	out.APIGroups = *(*[]string)(unsafe.Pointer(&in.APIGroups))
	out.Resources = *(*[]string)(unsafe.Pointer(&in.Resources))
	out.ResourceNames = *(*[]string)(unsafe.Pointer(&in.ResourceNames))
	return nil
}

// Convert_v1_ResourceRule_To_authorization_ResourceRule is an autogenerated conversion function.
func Convert_v1_ResourceRule_To_authorization_ResourceRule(in *v1.ResourceRule, out *authorization.ResourceRule, s conversion.Scope) error {
	return autoConvert_v1_ResourceRule_To_authorization_ResourceRule(in, out, s)
}

func autoConvert_authorization_ResourceRule_To_v1_ResourceRule(in *authorization.ResourceRule, out *v1.ResourceRule, s conversion.Scope) error {
	out.Verbs = *(*[]string)(unsafe.Pointer(&in.Verbs))
	out.APIGroups = *(*[]string)(unsafe.Pointer(&in.APIGroups))
	out.Resources = *(*[]string)(unsafe.Pointer(&in.Resources))
	out.ResourceNames = *(*[]string)(unsafe.Pointer(&in.ResourceNames))
	return nil
}

// Convert_authorization_ResourceRule_To_v1_ResourceRule is an autogenerated conversion function.
func Convert_authorization_ResourceRule_To_v1_ResourceRule(in *authorization.ResourceRule, out *v1.ResourceRule, s conversion.Scope) error {
	return autoConvert_authorization_ResourceRule_To_v1_ResourceRule(in, out, s)
}

func autoConvert_v1_SelfSubjectAccessReview_To_authorization_SelfSubjectAccessReview(in *v1.SelfSubjectAccessReview, out *authorization.SelfSubjectAccessReview, s conversion.Scope) error {
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
func Convert_v1_SelfSubjectAccessReview_To_authorization_SelfSubjectAccessReview(in *v1.SelfSubjectAccessReview, out *authorization.SelfSubjectAccessReview, s conversion.Scope) error {
	return autoConvert_v1_SelfSubjectAccessReview_To_authorization_SelfSubjectAccessReview(in, out, s)
}

func autoConvert_authorization_SelfSubjectAccessReview_To_v1_SelfSubjectAccessReview(in *authorization.SelfSubjectAccessReview, out *v1.SelfSubjectAccessReview, s conversion.Scope) error {
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
func Convert_authorization_SelfSubjectAccessReview_To_v1_SelfSubjectAccessReview(in *authorization.SelfSubjectAccessReview, out *v1.SelfSubjectAccessReview, s conversion.Scope) error {
	return autoConvert_authorization_SelfSubjectAccessReview_To_v1_SelfSubjectAccessReview(in, out, s)
}

func autoConvert_v1_SelfSubjectAccessReviewSpec_To_authorization_SelfSubjectAccessReviewSpec(in *v1.SelfSubjectAccessReviewSpec, out *authorization.SelfSubjectAccessReviewSpec, s conversion.Scope) error {
	out.ResourceAttributes = (*authorization.ResourceAttributes)(unsafe.Pointer(in.ResourceAttributes))
	out.NonResourceAttributes = (*authorization.NonResourceAttributes)(unsafe.Pointer(in.NonResourceAttributes))
	return nil
}

// Convert_v1_SelfSubjectAccessReviewSpec_To_authorization_SelfSubjectAccessReviewSpec is an autogenerated conversion function.
func Convert_v1_SelfSubjectAccessReviewSpec_To_authorization_SelfSubjectAccessReviewSpec(in *v1.SelfSubjectAccessReviewSpec, out *authorization.SelfSubjectAccessReviewSpec, s conversion.Scope) error {
	return autoConvert_v1_SelfSubjectAccessReviewSpec_To_authorization_SelfSubjectAccessReviewSpec(in, out, s)
}

func autoConvert_authorization_SelfSubjectAccessReviewSpec_To_v1_SelfSubjectAccessReviewSpec(in *authorization.SelfSubjectAccessReviewSpec, out *v1.SelfSubjectAccessReviewSpec, s conversion.Scope) error {
	out.ResourceAttributes = (*v1.ResourceAttributes)(unsafe.Pointer(in.ResourceAttributes))
	out.NonResourceAttributes = (*v1.NonResourceAttributes)(unsafe.Pointer(in.NonResourceAttributes))
	return nil
}

// Convert_authorization_SelfSubjectAccessReviewSpec_To_v1_SelfSubjectAccessReviewSpec is an autogenerated conversion function.
func Convert_authorization_SelfSubjectAccessReviewSpec_To_v1_SelfSubjectAccessReviewSpec(in *authorization.SelfSubjectAccessReviewSpec, out *v1.SelfSubjectAccessReviewSpec, s conversion.Scope) error {
	return autoConvert_authorization_SelfSubjectAccessReviewSpec_To_v1_SelfSubjectAccessReviewSpec(in, out, s)
}

func autoConvert_v1_SelfSubjectRulesReview_To_authorization_SelfSubjectRulesReview(in *v1.SelfSubjectRulesReview, out *authorization.SelfSubjectRulesReview, s conversion.Scope) error {
	out.ObjectMeta = in.ObjectMeta
	if err := Convert_v1_SelfSubjectRulesReviewSpec_To_authorization_SelfSubjectRulesReviewSpec(&in.Spec, &out.Spec, s); err != nil {
		return err
	}
	if err := Convert_v1_SubjectRulesReviewStatus_To_authorization_SubjectRulesReviewStatus(&in.Status, &out.Status, s); err != nil {
		return err
	}
	return nil
}

// Convert_v1_SelfSubjectRulesReview_To_authorization_SelfSubjectRulesReview is an autogenerated conversion function.
func Convert_v1_SelfSubjectRulesReview_To_authorization_SelfSubjectRulesReview(in *v1.SelfSubjectRulesReview, out *authorization.SelfSubjectRulesReview, s conversion.Scope) error {
	return autoConvert_v1_SelfSubjectRulesReview_To_authorization_SelfSubjectRulesReview(in, out, s)
}

func autoConvert_authorization_SelfSubjectRulesReview_To_v1_SelfSubjectRulesReview(in *authorization.SelfSubjectRulesReview, out *v1.SelfSubjectRulesReview, s conversion.Scope) error {
	out.ObjectMeta = in.ObjectMeta
	if err := Convert_authorization_SelfSubjectRulesReviewSpec_To_v1_SelfSubjectRulesReviewSpec(&in.Spec, &out.Spec, s); err != nil {
		return err
	}
	if err := Convert_authorization_SubjectRulesReviewStatus_To_v1_SubjectRulesReviewStatus(&in.Status, &out.Status, s); err != nil {
		return err
	}
	return nil
}

// Convert_authorization_SelfSubjectRulesReview_To_v1_SelfSubjectRulesReview is an autogenerated conversion function.
func Convert_authorization_SelfSubjectRulesReview_To_v1_SelfSubjectRulesReview(in *authorization.SelfSubjectRulesReview, out *v1.SelfSubjectRulesReview, s conversion.Scope) error {
	return autoConvert_authorization_SelfSubjectRulesReview_To_v1_SelfSubjectRulesReview(in, out, s)
}

func autoConvert_v1_SelfSubjectRulesReviewSpec_To_authorization_SelfSubjectRulesReviewSpec(in *v1.SelfSubjectRulesReviewSpec, out *authorization.SelfSubjectRulesReviewSpec, s conversion.Scope) error {
	out.Namespace = in.Namespace
	return nil
}

// Convert_v1_SelfSubjectRulesReviewSpec_To_authorization_SelfSubjectRulesReviewSpec is an autogenerated conversion function.
func Convert_v1_SelfSubjectRulesReviewSpec_To_authorization_SelfSubjectRulesReviewSpec(in *v1.SelfSubjectRulesReviewSpec, out *authorization.SelfSubjectRulesReviewSpec, s conversion.Scope) error {
	return autoConvert_v1_SelfSubjectRulesReviewSpec_To_authorization_SelfSubjectRulesReviewSpec(in, out, s)
}

func autoConvert_authorization_SelfSubjectRulesReviewSpec_To_v1_SelfSubjectRulesReviewSpec(in *authorization.SelfSubjectRulesReviewSpec, out *v1.SelfSubjectRulesReviewSpec, s conversion.Scope) error {
	out.Namespace = in.Namespace
	return nil
}

// Convert_authorization_SelfSubjectRulesReviewSpec_To_v1_SelfSubjectRulesReviewSpec is an autogenerated conversion function.
func Convert_authorization_SelfSubjectRulesReviewSpec_To_v1_SelfSubjectRulesReviewSpec(in *authorization.SelfSubjectRulesReviewSpec, out *v1.SelfSubjectRulesReviewSpec, s conversion.Scope) error {
	return autoConvert_authorization_SelfSubjectRulesReviewSpec_To_v1_SelfSubjectRulesReviewSpec(in, out, s)
}

func autoConvert_v1_SubjectAccessReview_To_authorization_SubjectAccessReview(in *v1.SubjectAccessReview, out *authorization.SubjectAccessReview, s conversion.Scope) error {
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
func Convert_v1_SubjectAccessReview_To_authorization_SubjectAccessReview(in *v1.SubjectAccessReview, out *authorization.SubjectAccessReview, s conversion.Scope) error {
	return autoConvert_v1_SubjectAccessReview_To_authorization_SubjectAccessReview(in, out, s)
}

func autoConvert_authorization_SubjectAccessReview_To_v1_SubjectAccessReview(in *authorization.SubjectAccessReview, out *v1.SubjectAccessReview, s conversion.Scope) error {
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
func Convert_authorization_SubjectAccessReview_To_v1_SubjectAccessReview(in *authorization.SubjectAccessReview, out *v1.SubjectAccessReview, s conversion.Scope) error {
	return autoConvert_authorization_SubjectAccessReview_To_v1_SubjectAccessReview(in, out, s)
}

func autoConvert_v1_SubjectAccessReviewSpec_To_authorization_SubjectAccessReviewSpec(in *v1.SubjectAccessReviewSpec, out *authorization.SubjectAccessReviewSpec, s conversion.Scope) error {
	out.ResourceAttributes = (*authorization.ResourceAttributes)(unsafe.Pointer(in.ResourceAttributes))
	out.NonResourceAttributes = (*authorization.NonResourceAttributes)(unsafe.Pointer(in.NonResourceAttributes))
	out.User = in.User
	out.Groups = *(*[]string)(unsafe.Pointer(&in.Groups))
	out.Extra = *(*map[string]authorization.ExtraValue)(unsafe.Pointer(&in.Extra))
	out.UID = in.UID
	return nil
}

// Convert_v1_SubjectAccessReviewSpec_To_authorization_SubjectAccessReviewSpec is an autogenerated conversion function.
func Convert_v1_SubjectAccessReviewSpec_To_authorization_SubjectAccessReviewSpec(in *v1.SubjectAccessReviewSpec, out *authorization.SubjectAccessReviewSpec, s conversion.Scope) error {
	return autoConvert_v1_SubjectAccessReviewSpec_To_authorization_SubjectAccessReviewSpec(in, out, s)
}

func autoConvert_authorization_SubjectAccessReviewSpec_To_v1_SubjectAccessReviewSpec(in *authorization.SubjectAccessReviewSpec, out *v1.SubjectAccessReviewSpec, s conversion.Scope) error {
	out.ResourceAttributes = (*v1.ResourceAttributes)(unsafe.Pointer(in.ResourceAttributes))
	out.NonResourceAttributes = (*v1.NonResourceAttributes)(unsafe.Pointer(in.NonResourceAttributes))
	out.User = in.User
	out.Groups = *(*[]string)(unsafe.Pointer(&in.Groups))
	out.Extra = *(*map[string]v1.ExtraValue)(unsafe.Pointer(&in.Extra))
	out.UID = in.UID
	return nil
}

// Convert_authorization_SubjectAccessReviewSpec_To_v1_SubjectAccessReviewSpec is an autogenerated conversion function.
func Convert_authorization_SubjectAccessReviewSpec_To_v1_SubjectAccessReviewSpec(in *authorization.SubjectAccessReviewSpec, out *v1.SubjectAccessReviewSpec, s conversion.Scope) error {
	return autoConvert_authorization_SubjectAccessReviewSpec_To_v1_SubjectAccessReviewSpec(in, out, s)
}

func autoConvert_v1_SubjectAccessReviewStatus_To_authorization_SubjectAccessReviewStatus(in *v1.SubjectAccessReviewStatus, out *authorization.SubjectAccessReviewStatus, s conversion.Scope) error {
	out.Allowed = in.Allowed
	out.Denied = in.Denied
	out.Reason = in.Reason
	out.EvaluationError = in.EvaluationError
	return nil
}

// Convert_v1_SubjectAccessReviewStatus_To_authorization_SubjectAccessReviewStatus is an autogenerated conversion function.
func Convert_v1_SubjectAccessReviewStatus_To_authorization_SubjectAccessReviewStatus(in *v1.SubjectAccessReviewStatus, out *authorization.SubjectAccessReviewStatus, s conversion.Scope) error {
	return autoConvert_v1_SubjectAccessReviewStatus_To_authorization_SubjectAccessReviewStatus(in, out, s)
}

func autoConvert_authorization_SubjectAccessReviewStatus_To_v1_SubjectAccessReviewStatus(in *authorization.SubjectAccessReviewStatus, out *v1.SubjectAccessReviewStatus, s conversion.Scope) error {
	out.Allowed = in.Allowed
	out.Denied = in.Denied
	out.Reason = in.Reason
	out.EvaluationError = in.EvaluationError
	return nil
}

// Convert_authorization_SubjectAccessReviewStatus_To_v1_SubjectAccessReviewStatus is an autogenerated conversion function.
func Convert_authorization_SubjectAccessReviewStatus_To_v1_SubjectAccessReviewStatus(in *authorization.SubjectAccessReviewStatus, out *v1.SubjectAccessReviewStatus, s conversion.Scope) error {
	return autoConvert_authorization_SubjectAccessReviewStatus_To_v1_SubjectAccessReviewStatus(in, out, s)
}

func autoConvert_v1_SubjectRulesReviewStatus_To_authorization_SubjectRulesReviewStatus(in *v1.SubjectRulesReviewStatus, out *authorization.SubjectRulesReviewStatus, s conversion.Scope) error {
	out.ResourceRules = *(*[]authorization.ResourceRule)(unsafe.Pointer(&in.ResourceRules))
	out.NonResourceRules = *(*[]authorization.NonResourceRule)(unsafe.Pointer(&in.NonResourceRules))
	out.Incomplete = in.Incomplete
	out.EvaluationError = in.EvaluationError
	return nil
}

// Convert_v1_SubjectRulesReviewStatus_To_authorization_SubjectRulesReviewStatus is an autogenerated conversion function.
func Convert_v1_SubjectRulesReviewStatus_To_authorization_SubjectRulesReviewStatus(in *v1.SubjectRulesReviewStatus, out *authorization.SubjectRulesReviewStatus, s conversion.Scope) error {
	return autoConvert_v1_SubjectRulesReviewStatus_To_authorization_SubjectRulesReviewStatus(in, out, s)
}

func autoConvert_authorization_SubjectRulesReviewStatus_To_v1_SubjectRulesReviewStatus(in *authorization.SubjectRulesReviewStatus, out *v1.SubjectRulesReviewStatus, s conversion.Scope) error {
	out.ResourceRules = *(*[]v1.ResourceRule)(unsafe.Pointer(&in.ResourceRules))
	out.NonResourceRules = *(*[]v1.NonResourceRule)(unsafe.Pointer(&in.NonResourceRules))
	out.Incomplete = in.Incomplete
	out.EvaluationError = in.EvaluationError
	return nil
}

// Convert_authorization_SubjectRulesReviewStatus_To_v1_SubjectRulesReviewStatus is an autogenerated conversion function.
func Convert_authorization_SubjectRulesReviewStatus_To_v1_SubjectRulesReviewStatus(in *authorization.SubjectRulesReviewStatus, out *v1.SubjectRulesReviewStatus, s conversion.Scope) error {
	return autoConvert_authorization_SubjectRulesReviewStatus_To_v1_SubjectRulesReviewStatus(in, out, s)
}
