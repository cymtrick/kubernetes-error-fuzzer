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

package admissionregistration

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
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_admissionregistration_AdmissionHookClientConfig, InType: reflect.TypeOf(&AdmissionHookClientConfig{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_admissionregistration_ExternalAdmissionHook, InType: reflect.TypeOf(&ExternalAdmissionHook{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_admissionregistration_ExternalAdmissionHookConfiguration, InType: reflect.TypeOf(&ExternalAdmissionHookConfiguration{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_admissionregistration_ExternalAdmissionHookConfigurationList, InType: reflect.TypeOf(&ExternalAdmissionHookConfigurationList{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_admissionregistration_Initializer, InType: reflect.TypeOf(&Initializer{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_admissionregistration_InitializerConfiguration, InType: reflect.TypeOf(&InitializerConfiguration{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_admissionregistration_InitializerConfigurationList, InType: reflect.TypeOf(&InitializerConfigurationList{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_admissionregistration_Rule, InType: reflect.TypeOf(&Rule{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_admissionregistration_RuleWithOperations, InType: reflect.TypeOf(&RuleWithOperations{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_admissionregistration_ServiceReference, InType: reflect.TypeOf(&ServiceReference{})},
	)
}

// DeepCopy_admissionregistration_AdmissionHookClientConfig is an autogenerated deepcopy function.
func DeepCopy_admissionregistration_AdmissionHookClientConfig(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*AdmissionHookClientConfig)
		out := out.(*AdmissionHookClientConfig)
		*out = *in
		if in.CABundle != nil {
			in, out := &in.CABundle, &out.CABundle
			*out = make([]byte, len(*in))
			copy(*out, *in)
		}
		return nil
	}
}

// DeepCopy_admissionregistration_ExternalAdmissionHook is an autogenerated deepcopy function.
func DeepCopy_admissionregistration_ExternalAdmissionHook(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*ExternalAdmissionHook)
		out := out.(*ExternalAdmissionHook)
		*out = *in
		if err := DeepCopy_admissionregistration_AdmissionHookClientConfig(&in.ClientConfig, &out.ClientConfig, c); err != nil {
			return err
		}
		if in.Rules != nil {
			in, out := &in.Rules, &out.Rules
			*out = make([]RuleWithOperations, len(*in))
			for i := range *in {
				if err := DeepCopy_admissionregistration_RuleWithOperations(&(*in)[i], &(*out)[i], c); err != nil {
					return err
				}
			}
		}
		if in.FailurePolicy != nil {
			in, out := &in.FailurePolicy, &out.FailurePolicy
			*out = new(FailurePolicyType)
			**out = **in
		}
		return nil
	}
}

// DeepCopy_admissionregistration_ExternalAdmissionHookConfiguration is an autogenerated deepcopy function.
func DeepCopy_admissionregistration_ExternalAdmissionHookConfiguration(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*ExternalAdmissionHookConfiguration)
		out := out.(*ExternalAdmissionHookConfiguration)
		*out = *in
		if newVal, err := c.DeepCopy(&in.ObjectMeta); err != nil {
			return err
		} else {
			out.ObjectMeta = *newVal.(*v1.ObjectMeta)
		}
		if in.ExternalAdmissionHooks != nil {
			in, out := &in.ExternalAdmissionHooks, &out.ExternalAdmissionHooks
			*out = make([]ExternalAdmissionHook, len(*in))
			for i := range *in {
				if err := DeepCopy_admissionregistration_ExternalAdmissionHook(&(*in)[i], &(*out)[i], c); err != nil {
					return err
				}
			}
		}
		return nil
	}
}

// DeepCopy_admissionregistration_ExternalAdmissionHookConfigurationList is an autogenerated deepcopy function.
func DeepCopy_admissionregistration_ExternalAdmissionHookConfigurationList(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*ExternalAdmissionHookConfigurationList)
		out := out.(*ExternalAdmissionHookConfigurationList)
		*out = *in
		if in.Items != nil {
			in, out := &in.Items, &out.Items
			*out = make([]ExternalAdmissionHookConfiguration, len(*in))
			for i := range *in {
				if err := DeepCopy_admissionregistration_ExternalAdmissionHookConfiguration(&(*in)[i], &(*out)[i], c); err != nil {
					return err
				}
			}
		}
		return nil
	}
}

// DeepCopy_admissionregistration_Initializer is an autogenerated deepcopy function.
func DeepCopy_admissionregistration_Initializer(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*Initializer)
		out := out.(*Initializer)
		*out = *in
		if in.Rules != nil {
			in, out := &in.Rules, &out.Rules
			*out = make([]Rule, len(*in))
			for i := range *in {
				if err := DeepCopy_admissionregistration_Rule(&(*in)[i], &(*out)[i], c); err != nil {
					return err
				}
			}
		}
		if in.FailurePolicy != nil {
			in, out := &in.FailurePolicy, &out.FailurePolicy
			*out = new(FailurePolicyType)
			**out = **in
		}
		return nil
	}
}

// DeepCopy_admissionregistration_InitializerConfiguration is an autogenerated deepcopy function.
func DeepCopy_admissionregistration_InitializerConfiguration(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*InitializerConfiguration)
		out := out.(*InitializerConfiguration)
		*out = *in
		if newVal, err := c.DeepCopy(&in.ObjectMeta); err != nil {
			return err
		} else {
			out.ObjectMeta = *newVal.(*v1.ObjectMeta)
		}
		if in.Initializers != nil {
			in, out := &in.Initializers, &out.Initializers
			*out = make([]Initializer, len(*in))
			for i := range *in {
				if err := DeepCopy_admissionregistration_Initializer(&(*in)[i], &(*out)[i], c); err != nil {
					return err
				}
			}
		}
		return nil
	}
}

// DeepCopy_admissionregistration_InitializerConfigurationList is an autogenerated deepcopy function.
func DeepCopy_admissionregistration_InitializerConfigurationList(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*InitializerConfigurationList)
		out := out.(*InitializerConfigurationList)
		*out = *in
		if in.Items != nil {
			in, out := &in.Items, &out.Items
			*out = make([]InitializerConfiguration, len(*in))
			for i := range *in {
				if err := DeepCopy_admissionregistration_InitializerConfiguration(&(*in)[i], &(*out)[i], c); err != nil {
					return err
				}
			}
		}
		return nil
	}
}

// DeepCopy_admissionregistration_Rule is an autogenerated deepcopy function.
func DeepCopy_admissionregistration_Rule(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*Rule)
		out := out.(*Rule)
		*out = *in
		if in.APIGroups != nil {
			in, out := &in.APIGroups, &out.APIGroups
			*out = make([]string, len(*in))
			copy(*out, *in)
		}
		if in.APIVersions != nil {
			in, out := &in.APIVersions, &out.APIVersions
			*out = make([]string, len(*in))
			copy(*out, *in)
		}
		if in.Resources != nil {
			in, out := &in.Resources, &out.Resources
			*out = make([]string, len(*in))
			copy(*out, *in)
		}
		return nil
	}
}

// DeepCopy_admissionregistration_RuleWithOperations is an autogenerated deepcopy function.
func DeepCopy_admissionregistration_RuleWithOperations(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*RuleWithOperations)
		out := out.(*RuleWithOperations)
		*out = *in
		if in.Operations != nil {
			in, out := &in.Operations, &out.Operations
			*out = make([]OperationType, len(*in))
			copy(*out, *in)
		}
		if err := DeepCopy_admissionregistration_Rule(&in.Rule, &out.Rule, c); err != nil {
			return err
		}
		return nil
	}
}

// DeepCopy_admissionregistration_ServiceReference is an autogenerated deepcopy function.
func DeepCopy_admissionregistration_ServiceReference(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*ServiceReference)
		out := out.(*ServiceReference)
		*out = *in
		return nil
	}
}
