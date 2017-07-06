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

package api

import (
	v1 "k8s.io/api/core/v1"
	conversion "k8s.io/apimachinery/pkg/conversion"
	runtime "k8s.io/apimachinery/pkg/runtime"
	rest "k8s.io/client-go/rest"
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
			in.(*ExtenderArgs).DeepCopyInto(out.(*ExtenderArgs))
			return nil
		}, InType: reflect.TypeOf(&ExtenderArgs{})},
		conversion.GeneratedDeepCopyFunc{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*ExtenderBindingArgs).DeepCopyInto(out.(*ExtenderBindingArgs))
			return nil
		}, InType: reflect.TypeOf(&ExtenderBindingArgs{})},
		conversion.GeneratedDeepCopyFunc{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*ExtenderBindingResult).DeepCopyInto(out.(*ExtenderBindingResult))
			return nil
		}, InType: reflect.TypeOf(&ExtenderBindingResult{})},
		conversion.GeneratedDeepCopyFunc{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*ExtenderConfig).DeepCopyInto(out.(*ExtenderConfig))
			return nil
		}, InType: reflect.TypeOf(&ExtenderConfig{})},
		conversion.GeneratedDeepCopyFunc{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*ExtenderFilterResult).DeepCopyInto(out.(*ExtenderFilterResult))
			return nil
		}, InType: reflect.TypeOf(&ExtenderFilterResult{})},
		conversion.GeneratedDeepCopyFunc{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*HostPriority).DeepCopyInto(out.(*HostPriority))
			return nil
		}, InType: reflect.TypeOf(&HostPriority{})},
		conversion.GeneratedDeepCopyFunc{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*LabelPreference).DeepCopyInto(out.(*LabelPreference))
			return nil
		}, InType: reflect.TypeOf(&LabelPreference{})},
		conversion.GeneratedDeepCopyFunc{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*LabelsPresence).DeepCopyInto(out.(*LabelsPresence))
			return nil
		}, InType: reflect.TypeOf(&LabelsPresence{})},
		conversion.GeneratedDeepCopyFunc{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*Policy).DeepCopyInto(out.(*Policy))
			return nil
		}, InType: reflect.TypeOf(&Policy{})},
		conversion.GeneratedDeepCopyFunc{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*PredicateArgument).DeepCopyInto(out.(*PredicateArgument))
			return nil
		}, InType: reflect.TypeOf(&PredicateArgument{})},
		conversion.GeneratedDeepCopyFunc{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*PredicatePolicy).DeepCopyInto(out.(*PredicatePolicy))
			return nil
		}, InType: reflect.TypeOf(&PredicatePolicy{})},
		conversion.GeneratedDeepCopyFunc{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*PriorityArgument).DeepCopyInto(out.(*PriorityArgument))
			return nil
		}, InType: reflect.TypeOf(&PriorityArgument{})},
		conversion.GeneratedDeepCopyFunc{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*PriorityPolicy).DeepCopyInto(out.(*PriorityPolicy))
			return nil
		}, InType: reflect.TypeOf(&PriorityPolicy{})},
		conversion.GeneratedDeepCopyFunc{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*ServiceAffinity).DeepCopyInto(out.(*ServiceAffinity))
			return nil
		}, InType: reflect.TypeOf(&ServiceAffinity{})},
		conversion.GeneratedDeepCopyFunc{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*ServiceAntiAffinity).DeepCopyInto(out.(*ServiceAntiAffinity))
			return nil
		}, InType: reflect.TypeOf(&ServiceAntiAffinity{})},
	)
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ExtenderArgs) DeepCopyInto(out *ExtenderArgs) {
	*out = *in
	in.Pod.DeepCopyInto(&out.Pod)
	if in.Nodes != nil {
		in, out := &in.Nodes, &out.Nodes
		if *in == nil {
			*out = nil
		} else {
			*out = new(v1.NodeList)
			(*in).DeepCopyInto(*out)
		}
	}
	if in.NodeNames != nil {
		in, out := &in.NodeNames, &out.NodeNames
		if *in == nil {
			*out = nil
		} else {
			*out = new([]string)
			if **in != nil {
				in, out := *in, *out
				*out = make([]string, len(*in))
				copy(*out, *in)
			}
		}
	}
	return
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, creating a new ExtenderArgs.
func (x *ExtenderArgs) DeepCopy() *ExtenderArgs {
	if x == nil {
		return nil
	}
	out := new(ExtenderArgs)
	x.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ExtenderBindingArgs) DeepCopyInto(out *ExtenderBindingArgs) {
	*out = *in
	return
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, creating a new ExtenderBindingArgs.
func (x *ExtenderBindingArgs) DeepCopy() *ExtenderBindingArgs {
	if x == nil {
		return nil
	}
	out := new(ExtenderBindingArgs)
	x.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ExtenderBindingResult) DeepCopyInto(out *ExtenderBindingResult) {
	*out = *in
	return
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, creating a new ExtenderBindingResult.
func (x *ExtenderBindingResult) DeepCopy() *ExtenderBindingResult {
	if x == nil {
		return nil
	}
	out := new(ExtenderBindingResult)
	x.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ExtenderConfig) DeepCopyInto(out *ExtenderConfig) {
	*out = *in
	if in.TLSConfig != nil {
		in, out := &in.TLSConfig, &out.TLSConfig
		if *in == nil {
			*out = nil
		} else {
			*out = new(rest.TLSClientConfig)
			(*in).DeepCopyInto(*out)
		}
	}
	return
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, creating a new ExtenderConfig.
func (x *ExtenderConfig) DeepCopy() *ExtenderConfig {
	if x == nil {
		return nil
	}
	out := new(ExtenderConfig)
	x.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ExtenderFilterResult) DeepCopyInto(out *ExtenderFilterResult) {
	*out = *in
	if in.Nodes != nil {
		in, out := &in.Nodes, &out.Nodes
		if *in == nil {
			*out = nil
		} else {
			*out = new(v1.NodeList)
			(*in).DeepCopyInto(*out)
		}
	}
	if in.NodeNames != nil {
		in, out := &in.NodeNames, &out.NodeNames
		if *in == nil {
			*out = nil
		} else {
			*out = new([]string)
			if **in != nil {
				in, out := *in, *out
				*out = make([]string, len(*in))
				copy(*out, *in)
			}
		}
	}
	if in.FailedNodes != nil {
		in, out := &in.FailedNodes, &out.FailedNodes
		*out = make(FailedNodesMap, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	return
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, creating a new ExtenderFilterResult.
func (x *ExtenderFilterResult) DeepCopy() *ExtenderFilterResult {
	if x == nil {
		return nil
	}
	out := new(ExtenderFilterResult)
	x.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HostPriority) DeepCopyInto(out *HostPriority) {
	*out = *in
	return
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, creating a new HostPriority.
func (x *HostPriority) DeepCopy() *HostPriority {
	if x == nil {
		return nil
	}
	out := new(HostPriority)
	x.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LabelPreference) DeepCopyInto(out *LabelPreference) {
	*out = *in
	return
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, creating a new LabelPreference.
func (x *LabelPreference) DeepCopy() *LabelPreference {
	if x == nil {
		return nil
	}
	out := new(LabelPreference)
	x.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LabelsPresence) DeepCopyInto(out *LabelsPresence) {
	*out = *in
	if in.Labels != nil {
		in, out := &in.Labels, &out.Labels
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, creating a new LabelsPresence.
func (x *LabelsPresence) DeepCopy() *LabelsPresence {
	if x == nil {
		return nil
	}
	out := new(LabelsPresence)
	x.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Policy) DeepCopyInto(out *Policy) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	if in.Predicates != nil {
		in, out := &in.Predicates, &out.Predicates
		*out = make([]PredicatePolicy, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Priorities != nil {
		in, out := &in.Priorities, &out.Priorities
		*out = make([]PriorityPolicy, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.ExtenderConfigs != nil {
		in, out := &in.ExtenderConfigs, &out.ExtenderConfigs
		*out = make([]ExtenderConfig, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, creating a new Policy.
func (x *Policy) DeepCopy() *Policy {
	if x == nil {
		return nil
	}
	out := new(Policy)
	x.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (x *Policy) DeepCopyObject() runtime.Object {
	if c := x.DeepCopy(); c != nil {
		return c
	} else {
		return nil
	}
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PredicateArgument) DeepCopyInto(out *PredicateArgument) {
	*out = *in
	if in.ServiceAffinity != nil {
		in, out := &in.ServiceAffinity, &out.ServiceAffinity
		if *in == nil {
			*out = nil
		} else {
			*out = new(ServiceAffinity)
			(*in).DeepCopyInto(*out)
		}
	}
	if in.LabelsPresence != nil {
		in, out := &in.LabelsPresence, &out.LabelsPresence
		if *in == nil {
			*out = nil
		} else {
			*out = new(LabelsPresence)
			(*in).DeepCopyInto(*out)
		}
	}
	return
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, creating a new PredicateArgument.
func (x *PredicateArgument) DeepCopy() *PredicateArgument {
	if x == nil {
		return nil
	}
	out := new(PredicateArgument)
	x.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PredicatePolicy) DeepCopyInto(out *PredicatePolicy) {
	*out = *in
	if in.Argument != nil {
		in, out := &in.Argument, &out.Argument
		if *in == nil {
			*out = nil
		} else {
			*out = new(PredicateArgument)
			(*in).DeepCopyInto(*out)
		}
	}
	return
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, creating a new PredicatePolicy.
func (x *PredicatePolicy) DeepCopy() *PredicatePolicy {
	if x == nil {
		return nil
	}
	out := new(PredicatePolicy)
	x.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PriorityArgument) DeepCopyInto(out *PriorityArgument) {
	*out = *in
	if in.ServiceAntiAffinity != nil {
		in, out := &in.ServiceAntiAffinity, &out.ServiceAntiAffinity
		if *in == nil {
			*out = nil
		} else {
			*out = new(ServiceAntiAffinity)
			**out = **in
		}
	}
	if in.LabelPreference != nil {
		in, out := &in.LabelPreference, &out.LabelPreference
		if *in == nil {
			*out = nil
		} else {
			*out = new(LabelPreference)
			**out = **in
		}
	}
	return
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, creating a new PriorityArgument.
func (x *PriorityArgument) DeepCopy() *PriorityArgument {
	if x == nil {
		return nil
	}
	out := new(PriorityArgument)
	x.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PriorityPolicy) DeepCopyInto(out *PriorityPolicy) {
	*out = *in
	if in.Argument != nil {
		in, out := &in.Argument, &out.Argument
		if *in == nil {
			*out = nil
		} else {
			*out = new(PriorityArgument)
			(*in).DeepCopyInto(*out)
		}
	}
	return
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, creating a new PriorityPolicy.
func (x *PriorityPolicy) DeepCopy() *PriorityPolicy {
	if x == nil {
		return nil
	}
	out := new(PriorityPolicy)
	x.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ServiceAffinity) DeepCopyInto(out *ServiceAffinity) {
	*out = *in
	if in.Labels != nil {
		in, out := &in.Labels, &out.Labels
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, creating a new ServiceAffinity.
func (x *ServiceAffinity) DeepCopy() *ServiceAffinity {
	if x == nil {
		return nil
	}
	out := new(ServiceAffinity)
	x.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ServiceAntiAffinity) DeepCopyInto(out *ServiceAntiAffinity) {
	*out = *in
	return
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, creating a new ServiceAntiAffinity.
func (x *ServiceAntiAffinity) DeepCopy() *ServiceAntiAffinity {
	if x == nil {
		return nil
	}
	out := new(ServiceAntiAffinity)
	x.DeepCopyInto(out)
	return out
}
