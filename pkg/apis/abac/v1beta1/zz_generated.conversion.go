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

package v1beta1

import (
	conversion "k8s.io/apimachinery/pkg/conversion"
	runtime "k8s.io/apimachinery/pkg/runtime"
	abac "k8s.io/kubernetes/pkg/apis/abac"
)

func init() {
	localSchemeBuilder.Register(RegisterConversions)
}

// RegisterConversions adds conversion functions to the given scheme.
// Public to allow building arbitrary schemes.
func RegisterConversions(s *runtime.Scheme) error {
	if err := s.AddGeneratedConversionFunc((*Policy)(nil), (*abac.Policy)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1beta1_Policy_To_abac_Policy(a.(*Policy), b.(*abac.Policy), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*abac.Policy)(nil), (*Policy)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_abac_Policy_To_v1beta1_Policy(a.(*abac.Policy), b.(*Policy), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*PolicySpec)(nil), (*abac.PolicySpec)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1beta1_PolicySpec_To_abac_PolicySpec(a.(*PolicySpec), b.(*abac.PolicySpec), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*abac.PolicySpec)(nil), (*PolicySpec)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_abac_PolicySpec_To_v1beta1_PolicySpec(a.(*abac.PolicySpec), b.(*PolicySpec), scope)
	}); err != nil {
		return err
	}
	return nil
}

func autoConvert_v1beta1_Policy_To_abac_Policy(in *Policy, out *abac.Policy, s conversion.Scope) error {
	if err := Convert_v1beta1_PolicySpec_To_abac_PolicySpec(&in.Spec, &out.Spec, s); err != nil {
		return err
	}
	return nil
}

// Convert_v1beta1_Policy_To_abac_Policy is an autogenerated conversion function.
func Convert_v1beta1_Policy_To_abac_Policy(in *Policy, out *abac.Policy, s conversion.Scope) error {
	return autoConvert_v1beta1_Policy_To_abac_Policy(in, out, s)
}

func autoConvert_abac_Policy_To_v1beta1_Policy(in *abac.Policy, out *Policy, s conversion.Scope) error {
	if err := Convert_abac_PolicySpec_To_v1beta1_PolicySpec(&in.Spec, &out.Spec, s); err != nil {
		return err
	}
	return nil
}

// Convert_abac_Policy_To_v1beta1_Policy is an autogenerated conversion function.
func Convert_abac_Policy_To_v1beta1_Policy(in *abac.Policy, out *Policy, s conversion.Scope) error {
	return autoConvert_abac_Policy_To_v1beta1_Policy(in, out, s)
}

func autoConvert_v1beta1_PolicySpec_To_abac_PolicySpec(in *PolicySpec, out *abac.PolicySpec, s conversion.Scope) error {
	out.User = in.User
	out.Group = in.Group
	out.Readonly = in.Readonly
	out.APIGroup = in.APIGroup
	out.Resource = in.Resource
	out.Namespace = in.Namespace
	out.NonResourcePath = in.NonResourcePath
	return nil
}

// Convert_v1beta1_PolicySpec_To_abac_PolicySpec is an autogenerated conversion function.
func Convert_v1beta1_PolicySpec_To_abac_PolicySpec(in *PolicySpec, out *abac.PolicySpec, s conversion.Scope) error {
	return autoConvert_v1beta1_PolicySpec_To_abac_PolicySpec(in, out, s)
}

func autoConvert_abac_PolicySpec_To_v1beta1_PolicySpec(in *abac.PolicySpec, out *PolicySpec, s conversion.Scope) error {
	out.User = in.User
	out.Group = in.Group
	out.Readonly = in.Readonly
	out.APIGroup = in.APIGroup
	out.Resource = in.Resource
	out.Namespace = in.Namespace
	out.NonResourcePath = in.NonResourcePath
	return nil
}

// Convert_abac_PolicySpec_To_v1beta1_PolicySpec is an autogenerated conversion function.
func Convert_abac_PolicySpec_To_v1beta1_PolicySpec(in *abac.PolicySpec, out *PolicySpec, s conversion.Scope) error {
	return autoConvert_abac_PolicySpec_To_v1beta1_PolicySpec(in, out, s)
}
