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
	resourcequota "k8s.io/kubernetes/plugin/pkg/admission/resourcequota/apis/resourcequota"
)

func init() {
	localSchemeBuilder.Register(RegisterConversions)
}

// RegisterConversions adds conversion functions to the given scheme.
// Public to allow building arbitrary schemes.
func RegisterConversions(scheme *runtime.Scheme) error {
	return scheme.AddGeneratedConversionFuncs(
		Convert_v1alpha1_Configuration_To_resourcequota_Configuration,
		Convert_resourcequota_Configuration_To_v1alpha1_Configuration,
		Convert_v1alpha1_LimitedResource_To_resourcequota_LimitedResource,
		Convert_resourcequota_LimitedResource_To_v1alpha1_LimitedResource,
	)
}

func autoConvert_v1alpha1_Configuration_To_resourcequota_Configuration(in *Configuration, out *resourcequota.Configuration, s conversion.Scope) error {
	out.LimitedResources = *(*[]resourcequota.LimitedResource)(unsafe.Pointer(&in.LimitedResources))
	return nil
}

// Convert_v1alpha1_Configuration_To_resourcequota_Configuration is an autogenerated conversion function.
func Convert_v1alpha1_Configuration_To_resourcequota_Configuration(in *Configuration, out *resourcequota.Configuration, s conversion.Scope) error {
	return autoConvert_v1alpha1_Configuration_To_resourcequota_Configuration(in, out, s)
}

func autoConvert_resourcequota_Configuration_To_v1alpha1_Configuration(in *resourcequota.Configuration, out *Configuration, s conversion.Scope) error {
	out.LimitedResources = *(*[]LimitedResource)(unsafe.Pointer(&in.LimitedResources))
	return nil
}

// Convert_resourcequota_Configuration_To_v1alpha1_Configuration is an autogenerated conversion function.
func Convert_resourcequota_Configuration_To_v1alpha1_Configuration(in *resourcequota.Configuration, out *Configuration, s conversion.Scope) error {
	return autoConvert_resourcequota_Configuration_To_v1alpha1_Configuration(in, out, s)
}

func autoConvert_v1alpha1_LimitedResource_To_resourcequota_LimitedResource(in *LimitedResource, out *resourcequota.LimitedResource, s conversion.Scope) error {
	out.APIGroup = in.APIGroup
	out.Resource = in.Resource
	out.MatchContains = *(*[]string)(unsafe.Pointer(&in.MatchContains))
	return nil
}

// Convert_v1alpha1_LimitedResource_To_resourcequota_LimitedResource is an autogenerated conversion function.
func Convert_v1alpha1_LimitedResource_To_resourcequota_LimitedResource(in *LimitedResource, out *resourcequota.LimitedResource, s conversion.Scope) error {
	return autoConvert_v1alpha1_LimitedResource_To_resourcequota_LimitedResource(in, out, s)
}

func autoConvert_resourcequota_LimitedResource_To_v1alpha1_LimitedResource(in *resourcequota.LimitedResource, out *LimitedResource, s conversion.Scope) error {
	out.APIGroup = in.APIGroup
	out.Resource = in.Resource
	out.MatchContains = *(*[]string)(unsafe.Pointer(&in.MatchContains))
	return nil
}

// Convert_resourcequota_LimitedResource_To_v1alpha1_LimitedResource is an autogenerated conversion function.
func Convert_resourcequota_LimitedResource_To_v1alpha1_LimitedResource(in *resourcequota.LimitedResource, out *LimitedResource, s conversion.Scope) error {
	return autoConvert_resourcequota_LimitedResource_To_v1alpha1_LimitedResource(in, out, s)
}
