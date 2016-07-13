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

package v1alpha1

import (
	api "k8s.io/kubernetes/pkg/api"
	unversioned "k8s.io/kubernetes/pkg/api/unversioned"
	v1 "k8s.io/kubernetes/pkg/api/v1"
	conversion "k8s.io/kubernetes/pkg/conversion"
	reflect "reflect"
)

func init() {
	if err := api.Scheme.AddGeneratedDeepCopyFuncs(
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_v1alpha1_PetSet, InType: reflect.TypeOf(func() *PetSet { var x *PetSet; return x }())},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_v1alpha1_PetSetList, InType: reflect.TypeOf(func() *PetSetList { var x *PetSetList; return x }())},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_v1alpha1_PetSetSpec, InType: reflect.TypeOf(func() *PetSetSpec { var x *PetSetSpec; return x }())},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_v1alpha1_PetSetStatus, InType: reflect.TypeOf(func() *PetSetStatus { var x *PetSetStatus; return x }())},
	); err != nil {
		// if one of the deep copy functions is malformed, detect it immediately.
		panic(err)
	}
}

func DeepCopy_v1alpha1_PetSet(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*PetSet)
		out := out.(*PetSet)
		out.TypeMeta = in.TypeMeta
		if err := v1.DeepCopy_v1_ObjectMeta(&in.ObjectMeta, &out.ObjectMeta, c); err != nil {
			return err
		}
		if err := DeepCopy_v1alpha1_PetSetSpec(&in.Spec, &out.Spec, c); err != nil {
			return err
		}
		if err := DeepCopy_v1alpha1_PetSetStatus(&in.Status, &out.Status, c); err != nil {
			return err
		}
		return nil
	}
}

func DeepCopy_v1alpha1_PetSetList(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*PetSetList)
		out := out.(*PetSetList)
		out.TypeMeta = in.TypeMeta
		out.ListMeta = in.ListMeta
		if in.Items != nil {
			in, out := &in.Items, &out.Items
			*out = make([]PetSet, len(*in))
			for i := range *in {
				if err := DeepCopy_v1alpha1_PetSet(&(*in)[i], &(*out)[i], c); err != nil {
					return err
				}
			}
		} else {
			out.Items = nil
		}
		return nil
	}
}

func DeepCopy_v1alpha1_PetSetSpec(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*PetSetSpec)
		out := out.(*PetSetSpec)
		if in.Replicas != nil {
			in, out := &in.Replicas, &out.Replicas
			*out = new(int32)
			**out = **in
		} else {
			out.Replicas = nil
		}
		if in.Selector != nil {
			in, out := &in.Selector, &out.Selector
			*out = new(unversioned.LabelSelector)
			if err := unversioned.DeepCopy_unversioned_LabelSelector(*in, *out, c); err != nil {
				return err
			}
		} else {
			out.Selector = nil
		}
		if err := v1.DeepCopy_v1_PodTemplateSpec(&in.Template, &out.Template, c); err != nil {
			return err
		}
		if in.VolumeClaimTemplates != nil {
			in, out := &in.VolumeClaimTemplates, &out.VolumeClaimTemplates
			*out = make([]v1.PersistentVolumeClaim, len(*in))
			for i := range *in {
				if err := v1.DeepCopy_v1_PersistentVolumeClaim(&(*in)[i], &(*out)[i], c); err != nil {
					return err
				}
			}
		} else {
			out.VolumeClaimTemplates = nil
		}
		out.ServiceName = in.ServiceName
		return nil
	}
}

func DeepCopy_v1alpha1_PetSetStatus(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*PetSetStatus)
		out := out.(*PetSetStatus)
		if in.ObservedGeneration != nil {
			in, out := &in.ObservedGeneration, &out.ObservedGeneration
			*out = new(int64)
			**out = **in
		} else {
			out.ObservedGeneration = nil
		}
		out.Replicas = in.Replicas
		return nil
	}
}
