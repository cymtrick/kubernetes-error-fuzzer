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
	v1alpha1 "k8s.io/api/admissionregistration/v1alpha1"
	conversion "k8s.io/apimachinery/pkg/conversion"
	runtime "k8s.io/apimachinery/pkg/runtime"
	admissionregistration "k8s.io/kubernetes/pkg/apis/admissionregistration"
	unsafe "unsafe"
)

func init() {
	localSchemeBuilder.Register(RegisterConversions)
}

// RegisterConversions adds conversion functions to the given scheme.
// Public to allow building arbitrary schemes.
func RegisterConversions(scheme *runtime.Scheme) error {
	return scheme.AddGeneratedConversionFuncs(
		Convert_v1alpha1_Initializer_To_admissionregistration_Initializer,
		Convert_admissionregistration_Initializer_To_v1alpha1_Initializer,
		Convert_v1alpha1_InitializerConfiguration_To_admissionregistration_InitializerConfiguration,
		Convert_admissionregistration_InitializerConfiguration_To_v1alpha1_InitializerConfiguration,
		Convert_v1alpha1_InitializerConfigurationList_To_admissionregistration_InitializerConfigurationList,
		Convert_admissionregistration_InitializerConfigurationList_To_v1alpha1_InitializerConfigurationList,
		Convert_v1alpha1_Rule_To_admissionregistration_Rule,
		Convert_admissionregistration_Rule_To_v1alpha1_Rule,
	)
}

func autoConvert_v1alpha1_Initializer_To_admissionregistration_Initializer(in *v1alpha1.Initializer, out *admissionregistration.Initializer, s conversion.Scope) error {
	out.Name = in.Name
	out.Rules = *(*[]admissionregistration.Rule)(unsafe.Pointer(&in.Rules))
	return nil
}

// Convert_v1alpha1_Initializer_To_admissionregistration_Initializer is an autogenerated conversion function.
func Convert_v1alpha1_Initializer_To_admissionregistration_Initializer(in *v1alpha1.Initializer, out *admissionregistration.Initializer, s conversion.Scope) error {
	return autoConvert_v1alpha1_Initializer_To_admissionregistration_Initializer(in, out, s)
}

func autoConvert_admissionregistration_Initializer_To_v1alpha1_Initializer(in *admissionregistration.Initializer, out *v1alpha1.Initializer, s conversion.Scope) error {
	out.Name = in.Name
	out.Rules = *(*[]v1alpha1.Rule)(unsafe.Pointer(&in.Rules))
	return nil
}

// Convert_admissionregistration_Initializer_To_v1alpha1_Initializer is an autogenerated conversion function.
func Convert_admissionregistration_Initializer_To_v1alpha1_Initializer(in *admissionregistration.Initializer, out *v1alpha1.Initializer, s conversion.Scope) error {
	return autoConvert_admissionregistration_Initializer_To_v1alpha1_Initializer(in, out, s)
}

func autoConvert_v1alpha1_InitializerConfiguration_To_admissionregistration_InitializerConfiguration(in *v1alpha1.InitializerConfiguration, out *admissionregistration.InitializerConfiguration, s conversion.Scope) error {
	out.ObjectMeta = in.ObjectMeta
	out.Initializers = *(*[]admissionregistration.Initializer)(unsafe.Pointer(&in.Initializers))
	return nil
}

// Convert_v1alpha1_InitializerConfiguration_To_admissionregistration_InitializerConfiguration is an autogenerated conversion function.
func Convert_v1alpha1_InitializerConfiguration_To_admissionregistration_InitializerConfiguration(in *v1alpha1.InitializerConfiguration, out *admissionregistration.InitializerConfiguration, s conversion.Scope) error {
	return autoConvert_v1alpha1_InitializerConfiguration_To_admissionregistration_InitializerConfiguration(in, out, s)
}

func autoConvert_admissionregistration_InitializerConfiguration_To_v1alpha1_InitializerConfiguration(in *admissionregistration.InitializerConfiguration, out *v1alpha1.InitializerConfiguration, s conversion.Scope) error {
	out.ObjectMeta = in.ObjectMeta
	out.Initializers = *(*[]v1alpha1.Initializer)(unsafe.Pointer(&in.Initializers))
	return nil
}

// Convert_admissionregistration_InitializerConfiguration_To_v1alpha1_InitializerConfiguration is an autogenerated conversion function.
func Convert_admissionregistration_InitializerConfiguration_To_v1alpha1_InitializerConfiguration(in *admissionregistration.InitializerConfiguration, out *v1alpha1.InitializerConfiguration, s conversion.Scope) error {
	return autoConvert_admissionregistration_InitializerConfiguration_To_v1alpha1_InitializerConfiguration(in, out, s)
}

func autoConvert_v1alpha1_InitializerConfigurationList_To_admissionregistration_InitializerConfigurationList(in *v1alpha1.InitializerConfigurationList, out *admissionregistration.InitializerConfigurationList, s conversion.Scope) error {
	out.ListMeta = in.ListMeta
	out.Items = *(*[]admissionregistration.InitializerConfiguration)(unsafe.Pointer(&in.Items))
	return nil
}

// Convert_v1alpha1_InitializerConfigurationList_To_admissionregistration_InitializerConfigurationList is an autogenerated conversion function.
func Convert_v1alpha1_InitializerConfigurationList_To_admissionregistration_InitializerConfigurationList(in *v1alpha1.InitializerConfigurationList, out *admissionregistration.InitializerConfigurationList, s conversion.Scope) error {
	return autoConvert_v1alpha1_InitializerConfigurationList_To_admissionregistration_InitializerConfigurationList(in, out, s)
}

func autoConvert_admissionregistration_InitializerConfigurationList_To_v1alpha1_InitializerConfigurationList(in *admissionregistration.InitializerConfigurationList, out *v1alpha1.InitializerConfigurationList, s conversion.Scope) error {
	out.ListMeta = in.ListMeta
	out.Items = *(*[]v1alpha1.InitializerConfiguration)(unsafe.Pointer(&in.Items))
	return nil
}

// Convert_admissionregistration_InitializerConfigurationList_To_v1alpha1_InitializerConfigurationList is an autogenerated conversion function.
func Convert_admissionregistration_InitializerConfigurationList_To_v1alpha1_InitializerConfigurationList(in *admissionregistration.InitializerConfigurationList, out *v1alpha1.InitializerConfigurationList, s conversion.Scope) error {
	return autoConvert_admissionregistration_InitializerConfigurationList_To_v1alpha1_InitializerConfigurationList(in, out, s)
}

func autoConvert_v1alpha1_Rule_To_admissionregistration_Rule(in *v1alpha1.Rule, out *admissionregistration.Rule, s conversion.Scope) error {
	out.APIGroups = *(*[]string)(unsafe.Pointer(&in.APIGroups))
	out.APIVersions = *(*[]string)(unsafe.Pointer(&in.APIVersions))
	out.Resources = *(*[]string)(unsafe.Pointer(&in.Resources))
	return nil
}

// Convert_v1alpha1_Rule_To_admissionregistration_Rule is an autogenerated conversion function.
func Convert_v1alpha1_Rule_To_admissionregistration_Rule(in *v1alpha1.Rule, out *admissionregistration.Rule, s conversion.Scope) error {
	return autoConvert_v1alpha1_Rule_To_admissionregistration_Rule(in, out, s)
}

func autoConvert_admissionregistration_Rule_To_v1alpha1_Rule(in *admissionregistration.Rule, out *v1alpha1.Rule, s conversion.Scope) error {
	out.APIGroups = *(*[]string)(unsafe.Pointer(&in.APIGroups))
	out.APIVersions = *(*[]string)(unsafe.Pointer(&in.APIVersions))
	out.Resources = *(*[]string)(unsafe.Pointer(&in.Resources))
	return nil
}

// Convert_admissionregistration_Rule_To_v1alpha1_Rule is an autogenerated conversion function.
func Convert_admissionregistration_Rule_To_v1alpha1_Rule(in *admissionregistration.Rule, out *v1alpha1.Rule, s conversion.Scope) error {
	return autoConvert_admissionregistration_Rule_To_v1alpha1_Rule(in, out, s)
}
