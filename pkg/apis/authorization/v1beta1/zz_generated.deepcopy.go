// +build !ignore_autogenerated

/*
Copyright 2016 The Kubernetes Authors.

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

package v1beta1

import (
	api "k8s.io/kubernetes/pkg/api"
	v1 "k8s.io/kubernetes/pkg/api/v1"
	conversion "k8s.io/kubernetes/pkg/conversion"
	reflect "reflect"
)

func init() {
	if err := api.Scheme.AddGeneratedDeepCopyFuncs(
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_v1beta1_LocalSubjectAccessReview, InType: reflect.TypeOf(&LocalSubjectAccessReview{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_v1beta1_NonResourceAttributes, InType: reflect.TypeOf(&NonResourceAttributes{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_v1beta1_ResourceAttributes, InType: reflect.TypeOf(&ResourceAttributes{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_v1beta1_SelfSubjectAccessReview, InType: reflect.TypeOf(&SelfSubjectAccessReview{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_v1beta1_SelfSubjectAccessReviewSpec, InType: reflect.TypeOf(&SelfSubjectAccessReviewSpec{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_v1beta1_SubjectAccessReview, InType: reflect.TypeOf(&SubjectAccessReview{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_v1beta1_SubjectAccessReviewSpec, InType: reflect.TypeOf(&SubjectAccessReviewSpec{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_v1beta1_SubjectAccessReviewStatus, InType: reflect.TypeOf(&SubjectAccessReviewStatus{})},
	); err != nil {
		// if one of the deep copy functions is malformed, detect it immediately.
		panic(err)
	}
}

func DeepCopy_v1beta1_LocalSubjectAccessReview(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*LocalSubjectAccessReview)
		out := out.(*LocalSubjectAccessReview)
		out.TypeMeta = in.TypeMeta
		if err := v1.DeepCopy_v1_ObjectMeta(&in.ObjectMeta, &out.ObjectMeta, c); err != nil {
			return err
		}
		if err := DeepCopy_v1beta1_SubjectAccessReviewSpec(&in.Spec, &out.Spec, c); err != nil {
			return err
		}
		out.Status = in.Status
		return nil
	}
}

func DeepCopy_v1beta1_NonResourceAttributes(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*NonResourceAttributes)
		out := out.(*NonResourceAttributes)
		out.Path = in.Path
		out.Verb = in.Verb
		return nil
	}
}

func DeepCopy_v1beta1_ResourceAttributes(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*ResourceAttributes)
		out := out.(*ResourceAttributes)
		out.Namespace = in.Namespace
		out.Verb = in.Verb
		out.Group = in.Group
		out.Version = in.Version
		out.Resource = in.Resource
		out.Subresource = in.Subresource
		out.Name = in.Name
		return nil
	}
}

func DeepCopy_v1beta1_SelfSubjectAccessReview(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*SelfSubjectAccessReview)
		out := out.(*SelfSubjectAccessReview)
		out.TypeMeta = in.TypeMeta
		if err := v1.DeepCopy_v1_ObjectMeta(&in.ObjectMeta, &out.ObjectMeta, c); err != nil {
			return err
		}
		if err := DeepCopy_v1beta1_SelfSubjectAccessReviewSpec(&in.Spec, &out.Spec, c); err != nil {
			return err
		}
		out.Status = in.Status
		return nil
	}
}

func DeepCopy_v1beta1_SelfSubjectAccessReviewSpec(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*SelfSubjectAccessReviewSpec)
		out := out.(*SelfSubjectAccessReviewSpec)
		if in.ResourceAttributes != nil {
			in, out := &in.ResourceAttributes, &out.ResourceAttributes
			*out = new(ResourceAttributes)
			**out = **in
		} else {
			out.ResourceAttributes = nil
		}
		if in.NonResourceAttributes != nil {
			in, out := &in.NonResourceAttributes, &out.NonResourceAttributes
			*out = new(NonResourceAttributes)
			**out = **in
		} else {
			out.NonResourceAttributes = nil
		}
		return nil
	}
}

func DeepCopy_v1beta1_SubjectAccessReview(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*SubjectAccessReview)
		out := out.(*SubjectAccessReview)
		out.TypeMeta = in.TypeMeta
		if err := v1.DeepCopy_v1_ObjectMeta(&in.ObjectMeta, &out.ObjectMeta, c); err != nil {
			return err
		}
		if err := DeepCopy_v1beta1_SubjectAccessReviewSpec(&in.Spec, &out.Spec, c); err != nil {
			return err
		}
		out.Status = in.Status
		return nil
	}
}

func DeepCopy_v1beta1_SubjectAccessReviewSpec(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*SubjectAccessReviewSpec)
		out := out.(*SubjectAccessReviewSpec)
		if in.ResourceAttributes != nil {
			in, out := &in.ResourceAttributes, &out.ResourceAttributes
			*out = new(ResourceAttributes)
			**out = **in
		} else {
			out.ResourceAttributes = nil
		}
		if in.NonResourceAttributes != nil {
			in, out := &in.NonResourceAttributes, &out.NonResourceAttributes
			*out = new(NonResourceAttributes)
			**out = **in
		} else {
			out.NonResourceAttributes = nil
		}
		out.User = in.User
		if in.Groups != nil {
			in, out := &in.Groups, &out.Groups
			*out = make([]string, len(*in))
			copy(*out, *in)
		} else {
			out.Groups = nil
		}
		if in.Extra != nil {
			in, out := &in.Extra, &out.Extra
			*out = make(map[string]ExtraValue)
			for key, val := range *in {
				if newVal, err := c.DeepCopy(&val); err != nil {
					return err
				} else {
					(*out)[key] = *newVal.(*ExtraValue)
				}
			}
		} else {
			out.Extra = nil
		}
		return nil
	}
}

func DeepCopy_v1beta1_SubjectAccessReviewStatus(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*SubjectAccessReviewStatus)
		out := out.(*SubjectAccessReviewStatus)
		out.Allowed = in.Allowed
		out.Reason = in.Reason
		out.EvaluationError = in.EvaluationError
		return nil
	}
}
