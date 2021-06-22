// +build !ignore_autogenerated

/*
Copyright The Kubernetes Authors.

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
	api "k8s.io/pod-security-admission/admission/api"
)

func init() {
	localSchemeBuilder.Register(RegisterConversions)
}

// RegisterConversions adds conversion functions to the given scheme.
// Public to allow building arbitrary schemes.
func RegisterConversions(s *runtime.Scheme) error {
	if err := s.AddGeneratedConversionFunc((*PodSecurityConfiguration)(nil), (*api.PodSecurityConfiguration)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha1_PodSecurityConfiguration_To_api_PodSecurityConfiguration(a.(*PodSecurityConfiguration), b.(*api.PodSecurityConfiguration), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*api.PodSecurityConfiguration)(nil), (*PodSecurityConfiguration)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_api_PodSecurityConfiguration_To_v1alpha1_PodSecurityConfiguration(a.(*api.PodSecurityConfiguration), b.(*PodSecurityConfiguration), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*PodSecurityDefaults)(nil), (*api.PodSecurityDefaults)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha1_PodSecurityDefaults_To_api_PodSecurityDefaults(a.(*PodSecurityDefaults), b.(*api.PodSecurityDefaults), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*api.PodSecurityDefaults)(nil), (*PodSecurityDefaults)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_api_PodSecurityDefaults_To_v1alpha1_PodSecurityDefaults(a.(*api.PodSecurityDefaults), b.(*PodSecurityDefaults), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*PodSecurityExemptions)(nil), (*api.PodSecurityExemptions)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha1_PodSecurityExemptions_To_api_PodSecurityExemptions(a.(*PodSecurityExemptions), b.(*api.PodSecurityExemptions), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*api.PodSecurityExemptions)(nil), (*PodSecurityExemptions)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_api_PodSecurityExemptions_To_v1alpha1_PodSecurityExemptions(a.(*api.PodSecurityExemptions), b.(*PodSecurityExemptions), scope)
	}); err != nil {
		return err
	}
	return nil
}

func autoConvert_v1alpha1_PodSecurityConfiguration_To_api_PodSecurityConfiguration(in *PodSecurityConfiguration, out *api.PodSecurityConfiguration, s conversion.Scope) error {
	if err := Convert_v1alpha1_PodSecurityDefaults_To_api_PodSecurityDefaults(&in.Defaults, &out.Defaults, s); err != nil {
		return err
	}
	if err := Convert_v1alpha1_PodSecurityExemptions_To_api_PodSecurityExemptions(&in.Exemptions, &out.Exemptions, s); err != nil {
		return err
	}
	return nil
}

// Convert_v1alpha1_PodSecurityConfiguration_To_api_PodSecurityConfiguration is an autogenerated conversion function.
func Convert_v1alpha1_PodSecurityConfiguration_To_api_PodSecurityConfiguration(in *PodSecurityConfiguration, out *api.PodSecurityConfiguration, s conversion.Scope) error {
	return autoConvert_v1alpha1_PodSecurityConfiguration_To_api_PodSecurityConfiguration(in, out, s)
}

func autoConvert_api_PodSecurityConfiguration_To_v1alpha1_PodSecurityConfiguration(in *api.PodSecurityConfiguration, out *PodSecurityConfiguration, s conversion.Scope) error {
	if err := Convert_api_PodSecurityDefaults_To_v1alpha1_PodSecurityDefaults(&in.Defaults, &out.Defaults, s); err != nil {
		return err
	}
	if err := Convert_api_PodSecurityExemptions_To_v1alpha1_PodSecurityExemptions(&in.Exemptions, &out.Exemptions, s); err != nil {
		return err
	}
	return nil
}

// Convert_api_PodSecurityConfiguration_To_v1alpha1_PodSecurityConfiguration is an autogenerated conversion function.
func Convert_api_PodSecurityConfiguration_To_v1alpha1_PodSecurityConfiguration(in *api.PodSecurityConfiguration, out *PodSecurityConfiguration, s conversion.Scope) error {
	return autoConvert_api_PodSecurityConfiguration_To_v1alpha1_PodSecurityConfiguration(in, out, s)
}

func autoConvert_v1alpha1_PodSecurityDefaults_To_api_PodSecurityDefaults(in *PodSecurityDefaults, out *api.PodSecurityDefaults, s conversion.Scope) error {
	out.Enforce = in.Enforce
	out.EnforceVersion = in.EnforceVersion
	out.Audit = in.Audit
	out.AuditVersion = in.AuditVersion
	out.Warn = in.Warn
	out.WarnVersion = in.WarnVersion
	return nil
}

// Convert_v1alpha1_PodSecurityDefaults_To_api_PodSecurityDefaults is an autogenerated conversion function.
func Convert_v1alpha1_PodSecurityDefaults_To_api_PodSecurityDefaults(in *PodSecurityDefaults, out *api.PodSecurityDefaults, s conversion.Scope) error {
	return autoConvert_v1alpha1_PodSecurityDefaults_To_api_PodSecurityDefaults(in, out, s)
}

func autoConvert_api_PodSecurityDefaults_To_v1alpha1_PodSecurityDefaults(in *api.PodSecurityDefaults, out *PodSecurityDefaults, s conversion.Scope) error {
	out.Enforce = in.Enforce
	out.EnforceVersion = in.EnforceVersion
	out.Audit = in.Audit
	out.AuditVersion = in.AuditVersion
	out.Warn = in.Warn
	out.WarnVersion = in.WarnVersion
	return nil
}

// Convert_api_PodSecurityDefaults_To_v1alpha1_PodSecurityDefaults is an autogenerated conversion function.
func Convert_api_PodSecurityDefaults_To_v1alpha1_PodSecurityDefaults(in *api.PodSecurityDefaults, out *PodSecurityDefaults, s conversion.Scope) error {
	return autoConvert_api_PodSecurityDefaults_To_v1alpha1_PodSecurityDefaults(in, out, s)
}

func autoConvert_v1alpha1_PodSecurityExemptions_To_api_PodSecurityExemptions(in *PodSecurityExemptions, out *api.PodSecurityExemptions, s conversion.Scope) error {
	out.Usernames = *(*[]string)(unsafe.Pointer(&in.Usernames))
	out.Namespaces = *(*[]string)(unsafe.Pointer(&in.Namespaces))
	out.RuntimeClasses = *(*[]string)(unsafe.Pointer(&in.RuntimeClasses))
	return nil
}

// Convert_v1alpha1_PodSecurityExemptions_To_api_PodSecurityExemptions is an autogenerated conversion function.
func Convert_v1alpha1_PodSecurityExemptions_To_api_PodSecurityExemptions(in *PodSecurityExemptions, out *api.PodSecurityExemptions, s conversion.Scope) error {
	return autoConvert_v1alpha1_PodSecurityExemptions_To_api_PodSecurityExemptions(in, out, s)
}

func autoConvert_api_PodSecurityExemptions_To_v1alpha1_PodSecurityExemptions(in *api.PodSecurityExemptions, out *PodSecurityExemptions, s conversion.Scope) error {
	out.Usernames = *(*[]string)(unsafe.Pointer(&in.Usernames))
	out.Namespaces = *(*[]string)(unsafe.Pointer(&in.Namespaces))
	out.RuntimeClasses = *(*[]string)(unsafe.Pointer(&in.RuntimeClasses))
	return nil
}

// Convert_api_PodSecurityExemptions_To_v1alpha1_PodSecurityExemptions is an autogenerated conversion function.
func Convert_api_PodSecurityExemptions_To_v1alpha1_PodSecurityExemptions(in *api.PodSecurityExemptions, out *PodSecurityExemptions, s conversion.Scope) error {
	return autoConvert_api_PodSecurityExemptions_To_v1alpha1_PodSecurityExemptions(in, out, s)
}
