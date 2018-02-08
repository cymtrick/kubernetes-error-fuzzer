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

package v1alpha1

import (
	unsafe "unsafe"

	authentication_v1 "k8s.io/api/authentication/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	conversion "k8s.io/apimachinery/pkg/conversion"
	runtime "k8s.io/apimachinery/pkg/runtime"
	types "k8s.io/apimachinery/pkg/types"
	audit "k8s.io/apiserver/pkg/apis/audit"
)

func init() {
	localSchemeBuilder.Register(RegisterConversions)
}

// RegisterConversions adds conversion functions to the given scheme.
// Public to allow building arbitrary schemes.
func RegisterConversions(scheme *runtime.Scheme) error {
	return scheme.AddGeneratedConversionFuncs(
		Convert_v1alpha1_Event_To_audit_Event,
		Convert_audit_Event_To_v1alpha1_Event,
		Convert_v1alpha1_EventList_To_audit_EventList,
		Convert_audit_EventList_To_v1alpha1_EventList,
		Convert_v1alpha1_GroupResources_To_audit_GroupResources,
		Convert_audit_GroupResources_To_v1alpha1_GroupResources,
		Convert_v1alpha1_ObjectReference_To_audit_ObjectReference,
		Convert_audit_ObjectReference_To_v1alpha1_ObjectReference,
		Convert_v1alpha1_Policy_To_audit_Policy,
		Convert_audit_Policy_To_v1alpha1_Policy,
		Convert_v1alpha1_PolicyList_To_audit_PolicyList,
		Convert_audit_PolicyList_To_v1alpha1_PolicyList,
		Convert_v1alpha1_PolicyRule_To_audit_PolicyRule,
		Convert_audit_PolicyRule_To_v1alpha1_PolicyRule,
	)
}

func autoConvert_v1alpha1_Event_To_audit_Event(in *Event, out *audit.Event, s conversion.Scope) error {
	// WARNING: in.ObjectMeta requires manual conversion: does not exist in peer-type
	out.Level = audit.Level(in.Level)
	// WARNING: in.Timestamp requires manual conversion: does not exist in peer-type
	out.AuditID = types.UID(in.AuditID)
	out.Stage = audit.Stage(in.Stage)
	out.RequestURI = in.RequestURI
	out.Verb = in.Verb
	// TODO: Inefficient conversion - can we improve it?
	if err := s.Convert(&in.User, &out.User, 0); err != nil {
		return err
	}
	out.ImpersonatedUser = (*audit.UserInfo)(unsafe.Pointer(in.ImpersonatedUser))
	out.SourceIPs = *(*[]string)(unsafe.Pointer(&in.SourceIPs))
	if in.ObjectRef != nil {
		in, out := &in.ObjectRef, &out.ObjectRef
		*out = new(audit.ObjectReference)
		if err := Convert_v1alpha1_ObjectReference_To_audit_ObjectReference(*in, *out, s); err != nil {
			return err
		}
	} else {
		out.ObjectRef = nil
	}
	out.ResponseStatus = (*v1.Status)(unsafe.Pointer(in.ResponseStatus))
	out.RequestObject = (*runtime.Unknown)(unsafe.Pointer(in.RequestObject))
	out.ResponseObject = (*runtime.Unknown)(unsafe.Pointer(in.ResponseObject))
	out.RequestReceivedTimestamp = in.RequestReceivedTimestamp
	out.StageTimestamp = in.StageTimestamp
	out.Annotations = *(*map[string]string)(unsafe.Pointer(&in.Annotations))
	return nil
}

func autoConvert_audit_Event_To_v1alpha1_Event(in *audit.Event, out *Event, s conversion.Scope) error {
	out.Level = Level(in.Level)
	out.AuditID = types.UID(in.AuditID)
	out.Stage = Stage(in.Stage)
	out.RequestURI = in.RequestURI
	out.Verb = in.Verb
	// TODO: Inefficient conversion - can we improve it?
	if err := s.Convert(&in.User, &out.User, 0); err != nil {
		return err
	}
	out.ImpersonatedUser = (*authentication_v1.UserInfo)(unsafe.Pointer(in.ImpersonatedUser))
	out.SourceIPs = *(*[]string)(unsafe.Pointer(&in.SourceIPs))
	if in.ObjectRef != nil {
		in, out := &in.ObjectRef, &out.ObjectRef
		*out = new(ObjectReference)
		if err := Convert_audit_ObjectReference_To_v1alpha1_ObjectReference(*in, *out, s); err != nil {
			return err
		}
	} else {
		out.ObjectRef = nil
	}
	out.ResponseStatus = (*v1.Status)(unsafe.Pointer(in.ResponseStatus))
	out.RequestObject = (*runtime.Unknown)(unsafe.Pointer(in.RequestObject))
	out.ResponseObject = (*runtime.Unknown)(unsafe.Pointer(in.ResponseObject))
	out.RequestReceivedTimestamp = in.RequestReceivedTimestamp
	out.StageTimestamp = in.StageTimestamp
	out.Annotations = *(*map[string]string)(unsafe.Pointer(&in.Annotations))
	return nil
}

func autoConvert_v1alpha1_EventList_To_audit_EventList(in *EventList, out *audit.EventList, s conversion.Scope) error {
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]audit.Event, len(*in))
		for i := range *in {
			if err := Convert_v1alpha1_Event_To_audit_Event(&(*in)[i], &(*out)[i], s); err != nil {
				return err
			}
		}
	} else {
		out.Items = nil
	}
	return nil
}

// Convert_v1alpha1_EventList_To_audit_EventList is an autogenerated conversion function.
func Convert_v1alpha1_EventList_To_audit_EventList(in *EventList, out *audit.EventList, s conversion.Scope) error {
	return autoConvert_v1alpha1_EventList_To_audit_EventList(in, out, s)
}

func autoConvert_audit_EventList_To_v1alpha1_EventList(in *audit.EventList, out *EventList, s conversion.Scope) error {
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Event, len(*in))
		for i := range *in {
			if err := Convert_audit_Event_To_v1alpha1_Event(&(*in)[i], &(*out)[i], s); err != nil {
				return err
			}
		}
	} else {
		out.Items = nil
	}
	return nil
}

// Convert_audit_EventList_To_v1alpha1_EventList is an autogenerated conversion function.
func Convert_audit_EventList_To_v1alpha1_EventList(in *audit.EventList, out *EventList, s conversion.Scope) error {
	return autoConvert_audit_EventList_To_v1alpha1_EventList(in, out, s)
}

func autoConvert_v1alpha1_GroupResources_To_audit_GroupResources(in *GroupResources, out *audit.GroupResources, s conversion.Scope) error {
	out.Group = in.Group
	out.Resources = *(*[]string)(unsafe.Pointer(&in.Resources))
	out.ResourceNames = *(*[]string)(unsafe.Pointer(&in.ResourceNames))
	return nil
}

// Convert_v1alpha1_GroupResources_To_audit_GroupResources is an autogenerated conversion function.
func Convert_v1alpha1_GroupResources_To_audit_GroupResources(in *GroupResources, out *audit.GroupResources, s conversion.Scope) error {
	return autoConvert_v1alpha1_GroupResources_To_audit_GroupResources(in, out, s)
}

func autoConvert_audit_GroupResources_To_v1alpha1_GroupResources(in *audit.GroupResources, out *GroupResources, s conversion.Scope) error {
	out.Group = in.Group
	out.Resources = *(*[]string)(unsafe.Pointer(&in.Resources))
	out.ResourceNames = *(*[]string)(unsafe.Pointer(&in.ResourceNames))
	return nil
}

// Convert_audit_GroupResources_To_v1alpha1_GroupResources is an autogenerated conversion function.
func Convert_audit_GroupResources_To_v1alpha1_GroupResources(in *audit.GroupResources, out *GroupResources, s conversion.Scope) error {
	return autoConvert_audit_GroupResources_To_v1alpha1_GroupResources(in, out, s)
}

func autoConvert_v1alpha1_ObjectReference_To_audit_ObjectReference(in *ObjectReference, out *audit.ObjectReference, s conversion.Scope) error {
	out.Resource = in.Resource
	out.Namespace = in.Namespace
	out.Name = in.Name
	out.UID = types.UID(in.UID)
	out.APIVersion = in.APIVersion
	out.ResourceVersion = in.ResourceVersion
	out.Subresource = in.Subresource
	return nil
}

func autoConvert_audit_ObjectReference_To_v1alpha1_ObjectReference(in *audit.ObjectReference, out *ObjectReference, s conversion.Scope) error {
	out.Resource = in.Resource
	out.Namespace = in.Namespace
	out.Name = in.Name
	out.UID = types.UID(in.UID)
	// WARNING: in.APIGroup requires manual conversion: does not exist in peer-type
	out.APIVersion = in.APIVersion
	out.ResourceVersion = in.ResourceVersion
	out.Subresource = in.Subresource
	return nil
}

func autoConvert_v1alpha1_Policy_To_audit_Policy(in *Policy, out *audit.Policy, s conversion.Scope) error {
	out.ObjectMeta = in.ObjectMeta
	out.Rules = *(*[]audit.PolicyRule)(unsafe.Pointer(&in.Rules))
	out.OmitStages = *(*[]audit.Stage)(unsafe.Pointer(&in.OmitStages))
	return nil
}

// Convert_v1alpha1_Policy_To_audit_Policy is an autogenerated conversion function.
func Convert_v1alpha1_Policy_To_audit_Policy(in *Policy, out *audit.Policy, s conversion.Scope) error {
	return autoConvert_v1alpha1_Policy_To_audit_Policy(in, out, s)
}

func autoConvert_audit_Policy_To_v1alpha1_Policy(in *audit.Policy, out *Policy, s conversion.Scope) error {
	out.ObjectMeta = in.ObjectMeta
	out.Rules = *(*[]PolicyRule)(unsafe.Pointer(&in.Rules))
	out.OmitStages = *(*[]Stage)(unsafe.Pointer(&in.OmitStages))
	return nil
}

// Convert_audit_Policy_To_v1alpha1_Policy is an autogenerated conversion function.
func Convert_audit_Policy_To_v1alpha1_Policy(in *audit.Policy, out *Policy, s conversion.Scope) error {
	return autoConvert_audit_Policy_To_v1alpha1_Policy(in, out, s)
}

func autoConvert_v1alpha1_PolicyList_To_audit_PolicyList(in *PolicyList, out *audit.PolicyList, s conversion.Scope) error {
	out.ListMeta = in.ListMeta
	out.Items = *(*[]audit.Policy)(unsafe.Pointer(&in.Items))
	return nil
}

// Convert_v1alpha1_PolicyList_To_audit_PolicyList is an autogenerated conversion function.
func Convert_v1alpha1_PolicyList_To_audit_PolicyList(in *PolicyList, out *audit.PolicyList, s conversion.Scope) error {
	return autoConvert_v1alpha1_PolicyList_To_audit_PolicyList(in, out, s)
}

func autoConvert_audit_PolicyList_To_v1alpha1_PolicyList(in *audit.PolicyList, out *PolicyList, s conversion.Scope) error {
	out.ListMeta = in.ListMeta
	out.Items = *(*[]Policy)(unsafe.Pointer(&in.Items))
	return nil
}

// Convert_audit_PolicyList_To_v1alpha1_PolicyList is an autogenerated conversion function.
func Convert_audit_PolicyList_To_v1alpha1_PolicyList(in *audit.PolicyList, out *PolicyList, s conversion.Scope) error {
	return autoConvert_audit_PolicyList_To_v1alpha1_PolicyList(in, out, s)
}

func autoConvert_v1alpha1_PolicyRule_To_audit_PolicyRule(in *PolicyRule, out *audit.PolicyRule, s conversion.Scope) error {
	out.Level = audit.Level(in.Level)
	out.Users = *(*[]string)(unsafe.Pointer(&in.Users))
	out.UserGroups = *(*[]string)(unsafe.Pointer(&in.UserGroups))
	out.Verbs = *(*[]string)(unsafe.Pointer(&in.Verbs))
	out.Resources = *(*[]audit.GroupResources)(unsafe.Pointer(&in.Resources))
	out.Namespaces = *(*[]string)(unsafe.Pointer(&in.Namespaces))
	out.NonResourceURLs = *(*[]string)(unsafe.Pointer(&in.NonResourceURLs))
	out.OmitStages = *(*[]audit.Stage)(unsafe.Pointer(&in.OmitStages))
	return nil
}

// Convert_v1alpha1_PolicyRule_To_audit_PolicyRule is an autogenerated conversion function.
func Convert_v1alpha1_PolicyRule_To_audit_PolicyRule(in *PolicyRule, out *audit.PolicyRule, s conversion.Scope) error {
	return autoConvert_v1alpha1_PolicyRule_To_audit_PolicyRule(in, out, s)
}

func autoConvert_audit_PolicyRule_To_v1alpha1_PolicyRule(in *audit.PolicyRule, out *PolicyRule, s conversion.Scope) error {
	out.Level = Level(in.Level)
	out.Users = *(*[]string)(unsafe.Pointer(&in.Users))
	out.UserGroups = *(*[]string)(unsafe.Pointer(&in.UserGroups))
	out.Verbs = *(*[]string)(unsafe.Pointer(&in.Verbs))
	out.Resources = *(*[]GroupResources)(unsafe.Pointer(&in.Resources))
	out.Namespaces = *(*[]string)(unsafe.Pointer(&in.Namespaces))
	out.NonResourceURLs = *(*[]string)(unsafe.Pointer(&in.NonResourceURLs))
	out.OmitStages = *(*[]Stage)(unsafe.Pointer(&in.OmitStages))
	return nil
}

// Convert_audit_PolicyRule_To_v1alpha1_PolicyRule is an autogenerated conversion function.
func Convert_audit_PolicyRule_To_v1alpha1_PolicyRule(in *audit.PolicyRule, out *PolicyRule, s conversion.Scope) error {
	return autoConvert_audit_PolicyRule_To_v1alpha1_PolicyRule(in, out, s)
}
