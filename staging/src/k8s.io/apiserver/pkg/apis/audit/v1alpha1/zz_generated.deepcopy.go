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

// This file was autogenerated by deepcopy-gen. Do not edit it manually!

package v1alpha1

import (
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	conversion "k8s.io/apimachinery/pkg/conversion"
	runtime "k8s.io/apimachinery/pkg/runtime"
	authentication_v1 "k8s.io/api/authentication/v1"
	reflect "reflect"
)

func init() {
	SchemeBuilder.Register(RegisterDeepCopies)
}

// RegisterDeepCopies adds deep-copy functions to the given scheme. Public
// to allow building arbitrary schemes.
func RegisterDeepCopies(scheme *runtime.Scheme) error {
	return scheme.AddGeneratedDeepCopyFuncs(
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_v1alpha1_Event, InType: reflect.TypeOf(&Event{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_v1alpha1_EventList, InType: reflect.TypeOf(&EventList{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_v1alpha1_GroupResources, InType: reflect.TypeOf(&GroupResources{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_v1alpha1_ObjectReference, InType: reflect.TypeOf(&ObjectReference{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_v1alpha1_Policy, InType: reflect.TypeOf(&Policy{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_v1alpha1_PolicyList, InType: reflect.TypeOf(&PolicyList{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_v1alpha1_PolicyRule, InType: reflect.TypeOf(&PolicyRule{})},
	)
}

// DeepCopy_v1alpha1_Event is an autogenerated deepcopy function.
func DeepCopy_v1alpha1_Event(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*Event)
		out := out.(*Event)
		*out = *in
		if newVal, err := c.DeepCopy(&in.ObjectMeta); err != nil {
			return err
		} else {
			out.ObjectMeta = *newVal.(*v1.ObjectMeta)
		}
		out.Timestamp = in.Timestamp.DeepCopy()
		if newVal, err := c.DeepCopy(&in.User); err != nil {
			return err
		} else {
			out.User = *newVal.(*authentication_v1.UserInfo)
		}
		if in.ImpersonatedUser != nil {
			in, out := &in.ImpersonatedUser, &out.ImpersonatedUser
			if newVal, err := c.DeepCopy(*in); err != nil {
				return err
			} else {
				*out = newVal.(*authentication_v1.UserInfo)
			}
		}
		if in.SourceIPs != nil {
			in, out := &in.SourceIPs, &out.SourceIPs
			*out = make([]string, len(*in))
			copy(*out, *in)
		}
		if in.ObjectRef != nil {
			in, out := &in.ObjectRef, &out.ObjectRef
			*out = new(ObjectReference)
			**out = **in
		}
		if in.ResponseStatus != nil {
			in, out := &in.ResponseStatus, &out.ResponseStatus
			if newVal, err := c.DeepCopy(*in); err != nil {
				return err
			} else {
				*out = newVal.(*v1.Status)
			}
		}
		if in.RequestObject != nil {
			in, out := &in.RequestObject, &out.RequestObject
			if newVal, err := c.DeepCopy(*in); err != nil {
				return err
			} else {
				*out = newVal.(*runtime.Unknown)
			}
		}
		if in.ResponseObject != nil {
			in, out := &in.ResponseObject, &out.ResponseObject
			if newVal, err := c.DeepCopy(*in); err != nil {
				return err
			} else {
				*out = newVal.(*runtime.Unknown)
			}
		}
		return nil
	}
}

// DeepCopy_v1alpha1_EventList is an autogenerated deepcopy function.
func DeepCopy_v1alpha1_EventList(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*EventList)
		out := out.(*EventList)
		*out = *in
		if in.Items != nil {
			in, out := &in.Items, &out.Items
			*out = make([]Event, len(*in))
			for i := range *in {
				if newVal, err := c.DeepCopy(&(*in)[i]); err != nil {
					return err
				} else {
					(*out)[i] = *newVal.(*Event)
				}
			}
		}
		return nil
	}
}

// DeepCopy_v1alpha1_GroupResources is an autogenerated deepcopy function.
func DeepCopy_v1alpha1_GroupResources(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*GroupResources)
		out := out.(*GroupResources)
		*out = *in
		if in.Resources != nil {
			in, out := &in.Resources, &out.Resources
			*out = make([]string, len(*in))
			copy(*out, *in)
		}
		return nil
	}
}

// DeepCopy_v1alpha1_ObjectReference is an autogenerated deepcopy function.
func DeepCopy_v1alpha1_ObjectReference(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*ObjectReference)
		out := out.(*ObjectReference)
		*out = *in
		return nil
	}
}

// DeepCopy_v1alpha1_Policy is an autogenerated deepcopy function.
func DeepCopy_v1alpha1_Policy(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*Policy)
		out := out.(*Policy)
		*out = *in
		if newVal, err := c.DeepCopy(&in.ObjectMeta); err != nil {
			return err
		} else {
			out.ObjectMeta = *newVal.(*v1.ObjectMeta)
		}
		if in.Rules != nil {
			in, out := &in.Rules, &out.Rules
			*out = make([]PolicyRule, len(*in))
			for i := range *in {
				if newVal, err := c.DeepCopy(&(*in)[i]); err != nil {
					return err
				} else {
					(*out)[i] = *newVal.(*PolicyRule)
				}
			}
		}
		return nil
	}
}

// DeepCopy_v1alpha1_PolicyList is an autogenerated deepcopy function.
func DeepCopy_v1alpha1_PolicyList(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*PolicyList)
		out := out.(*PolicyList)
		*out = *in
		if in.Items != nil {
			in, out := &in.Items, &out.Items
			*out = make([]Policy, len(*in))
			for i := range *in {
				if newVal, err := c.DeepCopy(&(*in)[i]); err != nil {
					return err
				} else {
					(*out)[i] = *newVal.(*Policy)
				}
			}
		}
		return nil
	}
}

// DeepCopy_v1alpha1_PolicyRule is an autogenerated deepcopy function.
func DeepCopy_v1alpha1_PolicyRule(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*PolicyRule)
		out := out.(*PolicyRule)
		*out = *in
		if in.Users != nil {
			in, out := &in.Users, &out.Users
			*out = make([]string, len(*in))
			copy(*out, *in)
		}
		if in.UserGroups != nil {
			in, out := &in.UserGroups, &out.UserGroups
			*out = make([]string, len(*in))
			copy(*out, *in)
		}
		if in.Verbs != nil {
			in, out := &in.Verbs, &out.Verbs
			*out = make([]string, len(*in))
			copy(*out, *in)
		}
		if in.Resources != nil {
			in, out := &in.Resources, &out.Resources
			*out = make([]GroupResources, len(*in))
			for i := range *in {
				if newVal, err := c.DeepCopy(&(*in)[i]); err != nil {
					return err
				} else {
					(*out)[i] = *newVal.(*GroupResources)
				}
			}
		}
		if in.Namespaces != nil {
			in, out := &in.Namespaces, &out.Namespaces
			*out = make([]string, len(*in))
			copy(*out, *in)
		}
		if in.NonResourceURLs != nil {
			in, out := &in.NonResourceURLs, &out.NonResourceURLs
			*out = make([]string, len(*in))
			copy(*out, *in)
		}
		return nil
	}
}
