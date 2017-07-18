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

package v1alpha1

import (
	conversion "k8s.io/apimachinery/pkg/conversion"
	runtime "k8s.io/apimachinery/pkg/runtime"
	reflect "reflect"
)

// Deprecated: register deep-copy functions.
func init() {
	SchemeBuilder.Register(RegisterDeepCopies)
}

// Deprecated: RegisterDeepCopies adds deep-copy functions to the given scheme. Public
// to allow building arbitrary schemes.
func RegisterDeepCopies(scheme *runtime.Scheme) error {
	return scheme.AddGeneratedDeepCopyFuncs(
		conversion.GeneratedDeepCopyFunc{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*AdmissionHookClientConfig).DeepCopyInto(out.(*AdmissionHookClientConfig))
			return nil
		}, InType: reflect.TypeOf(&AdmissionHookClientConfig{})},
		conversion.GeneratedDeepCopyFunc{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*ExternalAdmissionHook).DeepCopyInto(out.(*ExternalAdmissionHook))
			return nil
		}, InType: reflect.TypeOf(&ExternalAdmissionHook{})},
		conversion.GeneratedDeepCopyFunc{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*ExternalAdmissionHookConfiguration).DeepCopyInto(out.(*ExternalAdmissionHookConfiguration))
			return nil
		}, InType: reflect.TypeOf(&ExternalAdmissionHookConfiguration{})},
		conversion.GeneratedDeepCopyFunc{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*ExternalAdmissionHookConfigurationList).DeepCopyInto(out.(*ExternalAdmissionHookConfigurationList))
			return nil
		}, InType: reflect.TypeOf(&ExternalAdmissionHookConfigurationList{})},
		conversion.GeneratedDeepCopyFunc{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*Initializer).DeepCopyInto(out.(*Initializer))
			return nil
		}, InType: reflect.TypeOf(&Initializer{})},
		conversion.GeneratedDeepCopyFunc{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*InitializerConfiguration).DeepCopyInto(out.(*InitializerConfiguration))
			return nil
		}, InType: reflect.TypeOf(&InitializerConfiguration{})},
		conversion.GeneratedDeepCopyFunc{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*InitializerConfigurationList).DeepCopyInto(out.(*InitializerConfigurationList))
			return nil
		}, InType: reflect.TypeOf(&InitializerConfigurationList{})},
		conversion.GeneratedDeepCopyFunc{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*Rule).DeepCopyInto(out.(*Rule))
			return nil
		}, InType: reflect.TypeOf(&Rule{})},
		conversion.GeneratedDeepCopyFunc{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*RuleWithOperations).DeepCopyInto(out.(*RuleWithOperations))
			return nil
		}, InType: reflect.TypeOf(&RuleWithOperations{})},
		conversion.GeneratedDeepCopyFunc{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*ServiceReference).DeepCopyInto(out.(*ServiceReference))
			return nil
		}, InType: reflect.TypeOf(&ServiceReference{})},
	)
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AdmissionHookClientConfig) DeepCopyInto(out *AdmissionHookClientConfig) {
	*out = *in
	out.Service = in.Service
	if in.CABundle != nil {
		in, out := &in.CABundle, &out.CABundle
		*out = make([]byte, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, creating a new AdmissionHookClientConfig.
func (x *AdmissionHookClientConfig) DeepCopy() *AdmissionHookClientConfig {
	if x == nil {
		return nil
	}
	out := new(AdmissionHookClientConfig)
	x.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ExternalAdmissionHook) DeepCopyInto(out *ExternalAdmissionHook) {
	*out = *in
	in.ClientConfig.DeepCopyInto(&out.ClientConfig)
	if in.Rules != nil {
		in, out := &in.Rules, &out.Rules
		*out = make([]RuleWithOperations, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.FailurePolicy != nil {
		in, out := &in.FailurePolicy, &out.FailurePolicy
		if *in == nil {
			*out = nil
		} else {
			*out = new(FailurePolicyType)
			**out = **in
		}
	}
	return
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, creating a new ExternalAdmissionHook.
func (x *ExternalAdmissionHook) DeepCopy() *ExternalAdmissionHook {
	if x == nil {
		return nil
	}
	out := new(ExternalAdmissionHook)
	x.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ExternalAdmissionHookConfiguration) DeepCopyInto(out *ExternalAdmissionHookConfiguration) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	if in.ExternalAdmissionHooks != nil {
		in, out := &in.ExternalAdmissionHooks, &out.ExternalAdmissionHooks
		*out = make([]ExternalAdmissionHook, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, creating a new ExternalAdmissionHookConfiguration.
func (x *ExternalAdmissionHookConfiguration) DeepCopy() *ExternalAdmissionHookConfiguration {
	if x == nil {
		return nil
	}
	out := new(ExternalAdmissionHookConfiguration)
	x.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (x *ExternalAdmissionHookConfiguration) DeepCopyObject() runtime.Object {
	if c := x.DeepCopy(); c != nil {
		return c
	} else {
		return nil
	}
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ExternalAdmissionHookConfigurationList) DeepCopyInto(out *ExternalAdmissionHookConfigurationList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ExternalAdmissionHookConfiguration, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, creating a new ExternalAdmissionHookConfigurationList.
func (x *ExternalAdmissionHookConfigurationList) DeepCopy() *ExternalAdmissionHookConfigurationList {
	if x == nil {
		return nil
	}
	out := new(ExternalAdmissionHookConfigurationList)
	x.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (x *ExternalAdmissionHookConfigurationList) DeepCopyObject() runtime.Object {
	if c := x.DeepCopy(); c != nil {
		return c
	} else {
		return nil
	}
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Initializer) DeepCopyInto(out *Initializer) {
	*out = *in
	if in.Rules != nil {
		in, out := &in.Rules, &out.Rules
		*out = make([]Rule, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.FailurePolicy != nil {
		in, out := &in.FailurePolicy, &out.FailurePolicy
		if *in == nil {
			*out = nil
		} else {
			*out = new(FailurePolicyType)
			**out = **in
		}
	}
	return
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, creating a new Initializer.
func (x *Initializer) DeepCopy() *Initializer {
	if x == nil {
		return nil
	}
	out := new(Initializer)
	x.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *InitializerConfiguration) DeepCopyInto(out *InitializerConfiguration) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	if in.Initializers != nil {
		in, out := &in.Initializers, &out.Initializers
		*out = make([]Initializer, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, creating a new InitializerConfiguration.
func (x *InitializerConfiguration) DeepCopy() *InitializerConfiguration {
	if x == nil {
		return nil
	}
	out := new(InitializerConfiguration)
	x.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (x *InitializerConfiguration) DeepCopyObject() runtime.Object {
	if c := x.DeepCopy(); c != nil {
		return c
	} else {
		return nil
	}
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *InitializerConfigurationList) DeepCopyInto(out *InitializerConfigurationList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]InitializerConfiguration, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, creating a new InitializerConfigurationList.
func (x *InitializerConfigurationList) DeepCopy() *InitializerConfigurationList {
	if x == nil {
		return nil
	}
	out := new(InitializerConfigurationList)
	x.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (x *InitializerConfigurationList) DeepCopyObject() runtime.Object {
	if c := x.DeepCopy(); c != nil {
		return c
	} else {
		return nil
	}
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Rule) DeepCopyInto(out *Rule) {
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
	return
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, creating a new Rule.
func (x *Rule) DeepCopy() *Rule {
	if x == nil {
		return nil
	}
	out := new(Rule)
	x.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RuleWithOperations) DeepCopyInto(out *RuleWithOperations) {
	*out = *in
	if in.Operations != nil {
		in, out := &in.Operations, &out.Operations
		*out = make([]OperationType, len(*in))
		copy(*out, *in)
	}
	in.Rule.DeepCopyInto(&out.Rule)
	return
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, creating a new RuleWithOperations.
func (x *RuleWithOperations) DeepCopy() *RuleWithOperations {
	if x == nil {
		return nil
	}
	out := new(RuleWithOperations)
	x.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ServiceReference) DeepCopyInto(out *ServiceReference) {
	*out = *in
	return
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, creating a new ServiceReference.
func (x *ServiceReference) DeepCopy() *ServiceReference {
	if x == nil {
		return nil
	}
	out := new(ServiceReference)
	x.DeepCopyInto(out)
	return out
}
