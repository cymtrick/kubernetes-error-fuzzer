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

package authorization

import (
	reflect "reflect"

	conversion "k8s.io/apimachinery/pkg/conversion"
	runtime "k8s.io/apimachinery/pkg/runtime"
	api "k8s.io/kubernetes/pkg/api"
)

func init() {
	SchemeBuilder.Register(RegisterDeepCopies)
}

// RegisterDeepCopies adds deep-copy functions to the given scheme. Public
// to allow building arbitrary schemes.
func RegisterDeepCopies(scheme *runtime.Scheme) error {
	return scheme.AddGeneratedDeepCopyFuncs(
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_authorization_LocalSubjectAccessReview, InType: reflect.TypeOf(&LocalSubjectAccessReview{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_authorization_NonResourceAttributes, InType: reflect.TypeOf(&NonResourceAttributes{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_authorization_ResourceAttributes, InType: reflect.TypeOf(&ResourceAttributes{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_authorization_SelfSubjectAccessReview, InType: reflect.TypeOf(&SelfSubjectAccessReview{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_authorization_SelfSubjectAccessReviewSpec, InType: reflect.TypeOf(&SelfSubjectAccessReviewSpec{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_authorization_SubjectAccessReview, InType: reflect.TypeOf(&SubjectAccessReview{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_authorization_SubjectAccessReviewSpec, InType: reflect.TypeOf(&SubjectAccessReviewSpec{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_authorization_SubjectAccessReviewStatus, InType: reflect.TypeOf(&SubjectAccessReviewStatus{})},
	)
}

func DeepCopy_authorization_LocalSubjectAccessReview(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*LocalSubjectAccessReview)
		out := out.(*LocalSubjectAccessReview)
		*out = *in
		if err := api.DeepCopy_api_ObjectMeta(&in.ObjectMeta, &out.ObjectMeta, c); err != nil {
			return err
		}
		if err := DeepCopy_authorization_SubjectAccessReviewSpec(&in.Spec, &out.Spec, c); err != nil {
			return err
		}
		return nil
	}
}

func DeepCopy_authorization_NonResourceAttributes(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*NonResourceAttributes)
		out := out.(*NonResourceAttributes)
		*out = *in
		return nil
	}
}

func DeepCopy_authorization_ResourceAttributes(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*ResourceAttributes)
		out := out.(*ResourceAttributes)
		*out = *in
		return nil
	}
}

func DeepCopy_authorization_SelfSubjectAccessReview(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*SelfSubjectAccessReview)
		out := out.(*SelfSubjectAccessReview)
		*out = *in
		if err := api.DeepCopy_api_ObjectMeta(&in.ObjectMeta, &out.ObjectMeta, c); err != nil {
			return err
		}
		if err := DeepCopy_authorization_SelfSubjectAccessReviewSpec(&in.Spec, &out.Spec, c); err != nil {
			return err
		}
		return nil
	}
}

func DeepCopy_authorization_SelfSubjectAccessReviewSpec(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*SelfSubjectAccessReviewSpec)
		out := out.(*SelfSubjectAccessReviewSpec)
		*out = *in
		if in.ResourceAttributes != nil {
			in, out := &in.ResourceAttributes, &out.ResourceAttributes
			*out = new(ResourceAttributes)
			**out = **in
		}
		if in.NonResourceAttributes != nil {
			in, out := &in.NonResourceAttributes, &out.NonResourceAttributes
			*out = new(NonResourceAttributes)
			**out = **in
		}
		return nil
	}
}

func DeepCopy_authorization_SubjectAccessReview(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*SubjectAccessReview)
		out := out.(*SubjectAccessReview)
		*out = *in
		if err := api.DeepCopy_api_ObjectMeta(&in.ObjectMeta, &out.ObjectMeta, c); err != nil {
			return err
		}
		if err := DeepCopy_authorization_SubjectAccessReviewSpec(&in.Spec, &out.Spec, c); err != nil {
			return err
		}
		return nil
	}
}

func DeepCopy_authorization_SubjectAccessReviewSpec(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*SubjectAccessReviewSpec)
		out := out.(*SubjectAccessReviewSpec)
		*out = *in
		if in.ResourceAttributes != nil {
			in, out := &in.ResourceAttributes, &out.ResourceAttributes
			*out = new(ResourceAttributes)
			**out = **in
		}
		if in.NonResourceAttributes != nil {
			in, out := &in.NonResourceAttributes, &out.NonResourceAttributes
			*out = new(NonResourceAttributes)
			**out = **in
		}
		if in.Groups != nil {
			in, out := &in.Groups, &out.Groups
			*out = make([]string, len(*in))
			copy(*out, *in)
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
		}
		return nil
	}
}

func DeepCopy_authorization_SubjectAccessReviewStatus(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*SubjectAccessReviewStatus)
		out := out.(*SubjectAccessReviewStatus)
		*out = *in
		return nil
	}
}
