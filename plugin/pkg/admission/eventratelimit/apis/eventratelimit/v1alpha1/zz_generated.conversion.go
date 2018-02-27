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

// Code generated by conversion-gen. DO NOT EDIT.

package v1alpha1

import (
	unsafe "unsafe"

	conversion "k8s.io/apimachinery/pkg/conversion"
	runtime "k8s.io/apimachinery/pkg/runtime"
	eventratelimit "k8s.io/kubernetes/plugin/pkg/admission/eventratelimit/apis/eventratelimit"
)

func init() {
	localSchemeBuilder.Register(RegisterConversions)
}

// RegisterConversions adds conversion functions to the given scheme.
// Public to allow building arbitrary schemes.
func RegisterConversions(scheme *runtime.Scheme) error {
	return scheme.AddGeneratedConversionFuncs(
		Convert_v1alpha1_Configuration_To_eventratelimit_Configuration,
		Convert_eventratelimit_Configuration_To_v1alpha1_Configuration,
		Convert_v1alpha1_Limit_To_eventratelimit_Limit,
		Convert_eventratelimit_Limit_To_v1alpha1_Limit,
	)
}

func autoConvert_v1alpha1_Configuration_To_eventratelimit_Configuration(in *Configuration, out *eventratelimit.Configuration, s conversion.Scope) error {
	out.Limits = *(*[]eventratelimit.Limit)(unsafe.Pointer(&in.Limits))
	return nil
}

// Convert_v1alpha1_Configuration_To_eventratelimit_Configuration is an autogenerated conversion function.
func Convert_v1alpha1_Configuration_To_eventratelimit_Configuration(in *Configuration, out *eventratelimit.Configuration, s conversion.Scope) error {
	return autoConvert_v1alpha1_Configuration_To_eventratelimit_Configuration(in, out, s)
}

func autoConvert_eventratelimit_Configuration_To_v1alpha1_Configuration(in *eventratelimit.Configuration, out *Configuration, s conversion.Scope) error {
	out.Limits = *(*[]Limit)(unsafe.Pointer(&in.Limits))
	return nil
}

// Convert_eventratelimit_Configuration_To_v1alpha1_Configuration is an autogenerated conversion function.
func Convert_eventratelimit_Configuration_To_v1alpha1_Configuration(in *eventratelimit.Configuration, out *Configuration, s conversion.Scope) error {
	return autoConvert_eventratelimit_Configuration_To_v1alpha1_Configuration(in, out, s)
}

func autoConvert_v1alpha1_Limit_To_eventratelimit_Limit(in *Limit, out *eventratelimit.Limit, s conversion.Scope) error {
	out.Type = eventratelimit.LimitType(in.Type)
	out.QPS = in.QPS
	out.Burst = in.Burst
	out.CacheSize = in.CacheSize
	return nil
}

// Convert_v1alpha1_Limit_To_eventratelimit_Limit is an autogenerated conversion function.
func Convert_v1alpha1_Limit_To_eventratelimit_Limit(in *Limit, out *eventratelimit.Limit, s conversion.Scope) error {
	return autoConvert_v1alpha1_Limit_To_eventratelimit_Limit(in, out, s)
}

func autoConvert_eventratelimit_Limit_To_v1alpha1_Limit(in *eventratelimit.Limit, out *Limit, s conversion.Scope) error {
	out.Type = LimitType(in.Type)
	out.QPS = in.QPS
	out.Burst = in.Burst
	out.CacheSize = in.CacheSize
	return nil
}

// Convert_eventratelimit_Limit_To_v1alpha1_Limit is an autogenerated conversion function.
func Convert_eventratelimit_Limit_To_v1alpha1_Limit(in *eventratelimit.Limit, out *Limit, s conversion.Scope) error {
	return autoConvert_eventratelimit_Limit_To_v1alpha1_Limit(in, out, s)
}
