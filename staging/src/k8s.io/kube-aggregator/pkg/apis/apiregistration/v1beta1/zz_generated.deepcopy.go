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

package v1beta1

import (
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	conversion "k8s.io/apimachinery/pkg/conversion"
	runtime "k8s.io/apimachinery/pkg/runtime"
	reflect "reflect"
)

func init() {
	SchemeBuilder.Register(RegisterDeepCopies)
}

// RegisterDeepCopies adds deep-copy functions to the given scheme. Public
// to allow building arbitrary schemes.
func RegisterDeepCopies(scheme *runtime.Scheme) error {
	return scheme.AddGeneratedDeepCopyFuncs(
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_v1beta1_APIService, InType: reflect.TypeOf(&APIService{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_v1beta1_APIServiceCondition, InType: reflect.TypeOf(&APIServiceCondition{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_v1beta1_APIServiceList, InType: reflect.TypeOf(&APIServiceList{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_v1beta1_APIServiceSpec, InType: reflect.TypeOf(&APIServiceSpec{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_v1beta1_APIServiceStatus, InType: reflect.TypeOf(&APIServiceStatus{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_v1beta1_ServiceReference, InType: reflect.TypeOf(&ServiceReference{})},
	)
}

// DeepCopy_v1beta1_APIService is an autogenerated deepcopy function.
func DeepCopy_v1beta1_APIService(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*APIService)
		out := out.(*APIService)
		*out = *in
		if newVal, err := c.DeepCopy(&in.ObjectMeta); err != nil {
			return err
		} else {
			out.ObjectMeta = *newVal.(*v1.ObjectMeta)
		}
		if newVal, err := c.DeepCopy(&in.Spec); err != nil {
			return err
		} else {
			out.Spec = *newVal.(*APIServiceSpec)
		}
		if newVal, err := c.DeepCopy(&in.Status); err != nil {
			return err
		} else {
			out.Status = *newVal.(*APIServiceStatus)
		}
		return nil
	}
}

// DeepCopy_v1beta1_APIServiceCondition is an autogenerated deepcopy function.
func DeepCopy_v1beta1_APIServiceCondition(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*APIServiceCondition)
		out := out.(*APIServiceCondition)
		*out = *in
		out.LastTransitionTime = in.LastTransitionTime.DeepCopy()
		return nil
	}
}

// DeepCopy_v1beta1_APIServiceList is an autogenerated deepcopy function.
func DeepCopy_v1beta1_APIServiceList(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*APIServiceList)
		out := out.(*APIServiceList)
		*out = *in
		if in.Items != nil {
			in, out := &in.Items, &out.Items
			*out = make([]APIService, len(*in))
			for i := range *in {
				if newVal, err := c.DeepCopy(&(*in)[i]); err != nil {
					return err
				} else {
					(*out)[i] = *newVal.(*APIService)
				}
			}
		}
		return nil
	}
}

// DeepCopy_v1beta1_APIServiceSpec is an autogenerated deepcopy function.
func DeepCopy_v1beta1_APIServiceSpec(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*APIServiceSpec)
		out := out.(*APIServiceSpec)
		*out = *in
		if in.Service != nil {
			in, out := &in.Service, &out.Service
			*out = new(ServiceReference)
			**out = **in
		}
		if in.CABundle != nil {
			in, out := &in.CABundle, &out.CABundle
			*out = make([]byte, len(*in))
			copy(*out, *in)
		}
		return nil
	}
}

// DeepCopy_v1beta1_APIServiceStatus is an autogenerated deepcopy function.
func DeepCopy_v1beta1_APIServiceStatus(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*APIServiceStatus)
		out := out.(*APIServiceStatus)
		*out = *in
		if in.Conditions != nil {
			in, out := &in.Conditions, &out.Conditions
			*out = make([]APIServiceCondition, len(*in))
			for i := range *in {
				if newVal, err := c.DeepCopy(&(*in)[i]); err != nil {
					return err
				} else {
					(*out)[i] = *newVal.(*APIServiceCondition)
				}
			}
		}
		return nil
	}
}

// DeepCopy_v1beta1_ServiceReference is an autogenerated deepcopy function.
func DeepCopy_v1beta1_ServiceReference(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*ServiceReference)
		out := out.(*ServiceReference)
		*out = *in
		return nil
	}
}
