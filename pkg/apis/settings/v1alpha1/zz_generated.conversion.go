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

package v1alpha1

import (
	unsafe "unsafe"

	settingsv1alpha1 "k8s.io/api/settings/v1alpha1"

	v1 "k8s.io/api/core/v1"
	conversion "k8s.io/apimachinery/pkg/conversion"
	runtime "k8s.io/apimachinery/pkg/runtime"
	api "k8s.io/kubernetes/pkg/api"
	settings "k8s.io/kubernetes/pkg/apis/settings"
)

func init() {
	localSchemeBuilder.Register(RegisterConversions)
}

// RegisterConversions adds conversion functions to the given scheme.
// Public to allow building arbitrary schemes.
func RegisterConversions(scheme *runtime.Scheme) error {
	return scheme.AddGeneratedConversionFuncs(
		Convert_v1alpha1_PodPreset_To_settings_PodPreset,
		Convert_settings_PodPreset_To_v1alpha1_PodPreset,
		Convert_v1alpha1_PodPresetList_To_settings_PodPresetList,
		Convert_settings_PodPresetList_To_v1alpha1_PodPresetList,
		Convert_v1alpha1_PodPresetSpec_To_settings_PodPresetSpec,
		Convert_settings_PodPresetSpec_To_v1alpha1_PodPresetSpec,
	)
}

func autoConvert_v1alpha1_PodPreset_To_settings_PodPreset(in *settingsv1alpha1.PodPreset, out *settings.PodPreset, s conversion.Scope) error {
	out.ObjectMeta = in.ObjectMeta
	if err := Convert_v1alpha1_PodPresetSpec_To_settings_PodPresetSpec(&in.Spec, &out.Spec, s); err != nil {
		return err
	}
	return nil
}

// Convert_v1alpha1_PodPreset_To_settings_PodPreset is an autogenerated conversion function.
func Convert_v1alpha1_PodPreset_To_settings_PodPreset(in *settingsv1alpha1.PodPreset, out *settings.PodPreset, s conversion.Scope) error {
	return autoConvert_v1alpha1_PodPreset_To_settings_PodPreset(in, out, s)
}

func autoConvert_settings_PodPreset_To_v1alpha1_PodPreset(in *settings.PodPreset, out *settingsv1alpha1.PodPreset, s conversion.Scope) error {
	out.ObjectMeta = in.ObjectMeta
	if err := Convert_settings_PodPresetSpec_To_v1alpha1_PodPresetSpec(&in.Spec, &out.Spec, s); err != nil {
		return err
	}
	return nil
}

// Convert_settings_PodPreset_To_v1alpha1_PodPreset is an autogenerated conversion function.
func Convert_settings_PodPreset_To_v1alpha1_PodPreset(in *settings.PodPreset, out *settingsv1alpha1.PodPreset, s conversion.Scope) error {
	return autoConvert_settings_PodPreset_To_v1alpha1_PodPreset(in, out, s)
}

func autoConvert_v1alpha1_PodPresetList_To_settings_PodPresetList(in *settingsv1alpha1.PodPresetList, out *settings.PodPresetList, s conversion.Scope) error {
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]settings.PodPreset, len(*in))
		for i := range *in {
			if err := Convert_v1alpha1_PodPreset_To_settings_PodPreset(&(*in)[i], &(*out)[i], s); err != nil {
				return err
			}
		}
	} else {
		out.Items = nil
	}
	return nil
}

// Convert_v1alpha1_PodPresetList_To_settings_PodPresetList is an autogenerated conversion function.
func Convert_v1alpha1_PodPresetList_To_settings_PodPresetList(in *settingsv1alpha1.PodPresetList, out *settings.PodPresetList, s conversion.Scope) error {
	return autoConvert_v1alpha1_PodPresetList_To_settings_PodPresetList(in, out, s)
}

func autoConvert_settings_PodPresetList_To_v1alpha1_PodPresetList(in *settings.PodPresetList, out *settingsv1alpha1.PodPresetList, s conversion.Scope) error {
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]settingsv1alpha1.PodPreset, len(*in))
		for i := range *in {
			if err := Convert_settings_PodPreset_To_v1alpha1_PodPreset(&(*in)[i], &(*out)[i], s); err != nil {
				return err
			}
		}
	} else {
		out.Items = make([]settingsv1alpha1.PodPreset, 0)
	}
	return nil
}

// Convert_settings_PodPresetList_To_v1alpha1_PodPresetList is an autogenerated conversion function.
func Convert_settings_PodPresetList_To_v1alpha1_PodPresetList(in *settings.PodPresetList, out *settingsv1alpha1.PodPresetList, s conversion.Scope) error {
	return autoConvert_settings_PodPresetList_To_v1alpha1_PodPresetList(in, out, s)
}

func autoConvert_v1alpha1_PodPresetSpec_To_settings_PodPresetSpec(in *settingsv1alpha1.PodPresetSpec, out *settings.PodPresetSpec, s conversion.Scope) error {
	out.Selector = in.Selector
	out.Env = *(*[]api.EnvVar)(unsafe.Pointer(&in.Env))
	out.EnvFrom = *(*[]api.EnvFromSource)(unsafe.Pointer(&in.EnvFrom))
	if in.Volumes != nil {
		in, out := &in.Volumes, &out.Volumes
		*out = make([]api.Volume, len(*in))
		for i := range *in {
			// TODO: Inefficient conversion - can we improve it?
			if err := s.Convert(&(*in)[i], &(*out)[i], 0); err != nil {
				return err
			}
		}
	} else {
		out.Volumes = nil
	}
	out.VolumeMounts = *(*[]api.VolumeMount)(unsafe.Pointer(&in.VolumeMounts))
	return nil
}

// Convert_v1alpha1_PodPresetSpec_To_settings_PodPresetSpec is an autogenerated conversion function.
func Convert_v1alpha1_PodPresetSpec_To_settings_PodPresetSpec(in *settingsv1alpha1.PodPresetSpec, out *settings.PodPresetSpec, s conversion.Scope) error {
	return autoConvert_v1alpha1_PodPresetSpec_To_settings_PodPresetSpec(in, out, s)
}

func autoConvert_settings_PodPresetSpec_To_v1alpha1_PodPresetSpec(in *settings.PodPresetSpec, out *settingsv1alpha1.PodPresetSpec, s conversion.Scope) error {
	out.Selector = in.Selector
	out.Env = *(*[]v1.EnvVar)(unsafe.Pointer(&in.Env))
	out.EnvFrom = *(*[]v1.EnvFromSource)(unsafe.Pointer(&in.EnvFrom))
	if in.Volumes != nil {
		in, out := &in.Volumes, &out.Volumes
		*out = make([]v1.Volume, len(*in))
		for i := range *in {
			// TODO: Inefficient conversion - can we improve it?
			if err := s.Convert(&(*in)[i], &(*out)[i], 0); err != nil {
				return err
			}
		}
	} else {
		out.Volumes = nil
	}
	out.VolumeMounts = *(*[]v1.VolumeMount)(unsafe.Pointer(&in.VolumeMounts))
	return nil
}

// Convert_settings_PodPresetSpec_To_v1alpha1_PodPresetSpec is an autogenerated conversion function.
func Convert_settings_PodPresetSpec_To_v1alpha1_PodPresetSpec(in *settings.PodPresetSpec, out *settingsv1alpha1.PodPresetSpec, s conversion.Scope) error {
	return autoConvert_settings_PodPresetSpec_To_v1alpha1_PodPresetSpec(in, out, s)
}
