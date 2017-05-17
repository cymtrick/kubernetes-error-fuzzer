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

package v1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
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
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_v1_Pod, InType: reflect.TypeOf(&Pod{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_v1_PodCondition, InType: reflect.TypeOf(&PodCondition{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_v1_PodList, InType: reflect.TypeOf(&PodList{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_v1_PodSpec, InType: reflect.TypeOf(&PodSpec{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_v1_PodStatus, InType: reflect.TypeOf(&PodStatus{})},
	)
}

// DeepCopy_v1_Pod is an autogenerated deepcopy function.
func DeepCopy_v1_Pod(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*Pod)
		out := out.(*Pod)
		*out = *in
		if newVal, err := c.DeepCopy(&in.ObjectMeta); err != nil {
			return err
		} else {
			out.ObjectMeta = *newVal.(*meta_v1.ObjectMeta)
		}
		if newVal, err := c.DeepCopy(&in.Spec); err != nil {
			return err
		} else {
			out.Spec = *newVal.(*PodSpec)
		}
		if newVal, err := c.DeepCopy(&in.Status); err != nil {
			return err
		} else {
			out.Status = *newVal.(*PodStatus)
		}
		return nil
	}
}

// DeepCopy_v1_PodCondition is an autogenerated deepcopy function.
func DeepCopy_v1_PodCondition(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*PodCondition)
		out := out.(*PodCondition)
		*out = *in
		out.LastProbeTime = in.LastProbeTime.DeepCopy()
		out.LastTransitionTime = in.LastTransitionTime.DeepCopy()
		return nil
	}
}

// DeepCopy_v1_PodList is an autogenerated deepcopy function.
func DeepCopy_v1_PodList(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*PodList)
		out := out.(*PodList)
		*out = *in
		if in.Items != nil {
			in, out := &in.Items, &out.Items
			*out = make([]Pod, len(*in))
			for i := range *in {
				if newVal, err := c.DeepCopy(&(*in)[i]); err != nil {
					return err
				} else {
					(*out)[i] = *newVal.(*Pod)
				}
			}
		}
		return nil
	}
}

// DeepCopy_v1_PodSpec is an autogenerated deepcopy function.
func DeepCopy_v1_PodSpec(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*PodSpec)
		out := out.(*PodSpec)
		*out = *in
		if in.TerminationGracePeriodSeconds != nil {
			in, out := &in.TerminationGracePeriodSeconds, &out.TerminationGracePeriodSeconds
			*out = new(int64)
			**out = **in
		}
		if in.ActiveDeadlineSeconds != nil {
			in, out := &in.ActiveDeadlineSeconds, &out.ActiveDeadlineSeconds
			*out = new(int64)
			**out = **in
		}
		if in.NodeSelector != nil {
			in, out := &in.NodeSelector, &out.NodeSelector
			*out = make(map[string]string)
			for key, val := range *in {
				(*out)[key] = val
			}
		}
		return nil
	}
}

// DeepCopy_v1_PodStatus is an autogenerated deepcopy function.
func DeepCopy_v1_PodStatus(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*PodStatus)
		out := out.(*PodStatus)
		*out = *in
		if in.Conditions != nil {
			in, out := &in.Conditions, &out.Conditions
			*out = make([]PodCondition, len(*in))
			for i := range *in {
				if newVal, err := c.DeepCopy(&(*in)[i]); err != nil {
					return err
				} else {
					(*out)[i] = *newVal.(*PodCondition)
				}
			}
		}
		if in.StartTime != nil {
			in, out := &in.StartTime, &out.StartTime
			*out = new(meta_v1.Time)
			**out = (*in).DeepCopy()
		}
		return nil
	}
}
