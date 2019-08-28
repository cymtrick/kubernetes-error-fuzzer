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

// Code generated by deepcopy-gen. DO NOT EDIT.

package v1

import (
	corev1 "k8s.io/api/core/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ExtenderArgs) DeepCopyInto(out *ExtenderArgs) {
	*out = *in
	if in.Pod != nil {
		in, out := &in.Pod, &out.Pod
		*out = new(corev1.Pod)
		(*in).DeepCopyInto(*out)
	}
	if in.Nodes != nil {
		in, out := &in.Nodes, &out.Nodes
		*out = new(corev1.NodeList)
		(*in).DeepCopyInto(*out)
	}
	if in.NodeNames != nil {
		in, out := &in.NodeNames, &out.NodeNames
		*out = new([]string)
		if **in != nil {
			in, out := *in, *out
			*out = make([]string, len(*in))
			copy(*out, *in)
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ExtenderArgs.
func (in *ExtenderArgs) DeepCopy() *ExtenderArgs {
	if in == nil {
		return nil
	}
	out := new(ExtenderArgs)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ExtenderBindingArgs) DeepCopyInto(out *ExtenderBindingArgs) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ExtenderBindingArgs.
func (in *ExtenderBindingArgs) DeepCopy() *ExtenderBindingArgs {
	if in == nil {
		return nil
	}
	out := new(ExtenderBindingArgs)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ExtenderBindingResult) DeepCopyInto(out *ExtenderBindingResult) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ExtenderBindingResult.
func (in *ExtenderBindingResult) DeepCopy() *ExtenderBindingResult {
	if in == nil {
		return nil
	}
	out := new(ExtenderBindingResult)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ExtenderConfig) DeepCopyInto(out *ExtenderConfig) {
	*out = *in
	if in.TLSConfig != nil {
		in, out := &in.TLSConfig, &out.TLSConfig
		*out = new(ExtenderTLSConfig)
		(*in).DeepCopyInto(*out)
	}
	if in.ManagedResources != nil {
		in, out := &in.ManagedResources, &out.ManagedResources
		*out = make([]ExtenderManagedResource, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ExtenderConfig.
func (in *ExtenderConfig) DeepCopy() *ExtenderConfig {
	if in == nil {
		return nil
	}
	out := new(ExtenderConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ExtenderFilterResult) DeepCopyInto(out *ExtenderFilterResult) {
	*out = *in
	if in.Nodes != nil {
		in, out := &in.Nodes, &out.Nodes
		*out = new(corev1.NodeList)
		(*in).DeepCopyInto(*out)
	}
	if in.NodeNames != nil {
		in, out := &in.NodeNames, &out.NodeNames
		*out = new([]string)
		if **in != nil {
			in, out := *in, *out
			*out = make([]string, len(*in))
			copy(*out, *in)
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

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ExtenderFilterResult.
func (in *ExtenderFilterResult) DeepCopy() *ExtenderFilterResult {
	if in == nil {
		return nil
	}
	out := new(ExtenderFilterResult)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ExtenderManagedResource) DeepCopyInto(out *ExtenderManagedResource) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ExtenderManagedResource.
func (in *ExtenderManagedResource) DeepCopy() *ExtenderManagedResource {
	if in == nil {
		return nil
	}
	out := new(ExtenderManagedResource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ExtenderPreemptionArgs) DeepCopyInto(out *ExtenderPreemptionArgs) {
	*out = *in
	if in.Pod != nil {
		in, out := &in.Pod, &out.Pod
		*out = new(corev1.Pod)
		(*in).DeepCopyInto(*out)
	}
	if in.NodeNameToVictims != nil {
		in, out := &in.NodeNameToVictims, &out.NodeNameToVictims
		*out = make(map[string]*Victims, len(*in))
		for key, val := range *in {
			var outVal *Victims
			if val == nil {
				(*out)[key] = nil
			} else {
				in, out := &val, &outVal
				*out = new(Victims)
				(*in).DeepCopyInto(*out)
			}
			(*out)[key] = outVal
		}
	}
	if in.NodeNameToMetaVictims != nil {
		in, out := &in.NodeNameToMetaVictims, &out.NodeNameToMetaVictims
		*out = make(map[string]*MetaVictims, len(*in))
		for key, val := range *in {
			var outVal *MetaVictims
			if val == nil {
				(*out)[key] = nil
			} else {
				in, out := &val, &outVal
				*out = new(MetaVictims)
				(*in).DeepCopyInto(*out)
			}
			(*out)[key] = outVal
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ExtenderPreemptionArgs.
func (in *ExtenderPreemptionArgs) DeepCopy() *ExtenderPreemptionArgs {
	if in == nil {
		return nil
	}
	out := new(ExtenderPreemptionArgs)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ExtenderPreemptionResult) DeepCopyInto(out *ExtenderPreemptionResult) {
	*out = *in
	if in.NodeNameToMetaVictims != nil {
		in, out := &in.NodeNameToMetaVictims, &out.NodeNameToMetaVictims
		*out = make(map[string]*MetaVictims, len(*in))
		for key, val := range *in {
			var outVal *MetaVictims
			if val == nil {
				(*out)[key] = nil
			} else {
				in, out := &val, &outVal
				*out = new(MetaVictims)
				(*in).DeepCopyInto(*out)
			}
			(*out)[key] = outVal
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ExtenderPreemptionResult.
func (in *ExtenderPreemptionResult) DeepCopy() *ExtenderPreemptionResult {
	if in == nil {
		return nil
	}
	out := new(ExtenderPreemptionResult)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ExtenderTLSConfig) DeepCopyInto(out *ExtenderTLSConfig) {
	*out = *in
	if in.CertData != nil {
		in, out := &in.CertData, &out.CertData
		*out = make([]byte, len(*in))
		copy(*out, *in)
	}
	if in.KeyData != nil {
		in, out := &in.KeyData, &out.KeyData
		*out = make([]byte, len(*in))
		copy(*out, *in)
	}
	if in.CAData != nil {
		in, out := &in.CAData, &out.CAData
		*out = make([]byte, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ExtenderTLSConfig.
func (in *ExtenderTLSConfig) DeepCopy() *ExtenderTLSConfig {
	if in == nil {
		return nil
	}
	out := new(ExtenderTLSConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in FailedNodesMap) DeepCopyInto(out *FailedNodesMap) {
	{
		in := &in
		*out = make(FailedNodesMap, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
		return
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FailedNodesMap.
func (in FailedNodesMap) DeepCopy() FailedNodesMap {
	if in == nil {
		return nil
	}
	out := new(FailedNodesMap)
	in.DeepCopyInto(out)
	return *out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HostPriority) DeepCopyInto(out *HostPriority) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HostPriority.
func (in *HostPriority) DeepCopy() *HostPriority {
	if in == nil {
		return nil
	}
	out := new(HostPriority)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in HostPriorityList) DeepCopyInto(out *HostPriorityList) {
	{
		in := &in
		*out = make(HostPriorityList, len(*in))
		copy(*out, *in)
		return
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HostPriorityList.
func (in HostPriorityList) DeepCopy() HostPriorityList {
	if in == nil {
		return nil
	}
	out := new(HostPriorityList)
	in.DeepCopyInto(out)
	return *out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LabelPreference) DeepCopyInto(out *LabelPreference) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LabelPreference.
func (in *LabelPreference) DeepCopy() *LabelPreference {
	if in == nil {
		return nil
	}
	out := new(LabelPreference)
	in.DeepCopyInto(out)
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

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LabelsPresence.
func (in *LabelsPresence) DeepCopy() *LabelsPresence {
	if in == nil {
		return nil
	}
	out := new(LabelsPresence)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MetaPod) DeepCopyInto(out *MetaPod) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MetaPod.
func (in *MetaPod) DeepCopy() *MetaPod {
	if in == nil {
		return nil
	}
	out := new(MetaPod)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MetaVictims) DeepCopyInto(out *MetaVictims) {
	*out = *in
	if in.Pods != nil {
		in, out := &in.Pods, &out.Pods
		*out = make([]*MetaPod, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(MetaPod)
				**out = **in
			}
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MetaVictims.
func (in *MetaVictims) DeepCopy() *MetaVictims {
	if in == nil {
		return nil
	}
	out := new(MetaVictims)
	in.DeepCopyInto(out)
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

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Policy.
func (in *Policy) DeepCopy() *Policy {
	if in == nil {
		return nil
	}
	out := new(Policy)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Policy) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PredicateArgument) DeepCopyInto(out *PredicateArgument) {
	*out = *in
	if in.ServiceAffinity != nil {
		in, out := &in.ServiceAffinity, &out.ServiceAffinity
		*out = new(ServiceAffinity)
		(*in).DeepCopyInto(*out)
	}
	if in.LabelsPresence != nil {
		in, out := &in.LabelsPresence, &out.LabelsPresence
		*out = new(LabelsPresence)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PredicateArgument.
func (in *PredicateArgument) DeepCopy() *PredicateArgument {
	if in == nil {
		return nil
	}
	out := new(PredicateArgument)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PredicatePolicy) DeepCopyInto(out *PredicatePolicy) {
	*out = *in
	if in.Argument != nil {
		in, out := &in.Argument, &out.Argument
		*out = new(PredicateArgument)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PredicatePolicy.
func (in *PredicatePolicy) DeepCopy() *PredicatePolicy {
	if in == nil {
		return nil
	}
	out := new(PredicatePolicy)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PriorityArgument) DeepCopyInto(out *PriorityArgument) {
	*out = *in
	if in.ServiceAntiAffinity != nil {
		in, out := &in.ServiceAntiAffinity, &out.ServiceAntiAffinity
		*out = new(ServiceAntiAffinity)
		**out = **in
	}
	if in.LabelPreference != nil {
		in, out := &in.LabelPreference, &out.LabelPreference
		*out = new(LabelPreference)
		**out = **in
	}
	if in.RequestedToCapacityRatioArguments != nil {
		in, out := &in.RequestedToCapacityRatioArguments, &out.RequestedToCapacityRatioArguments
		*out = new(RequestedToCapacityRatioArguments)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PriorityArgument.
func (in *PriorityArgument) DeepCopy() *PriorityArgument {
	if in == nil {
		return nil
	}
	out := new(PriorityArgument)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PriorityPolicy) DeepCopyInto(out *PriorityPolicy) {
	*out = *in
	if in.Argument != nil {
		in, out := &in.Argument, &out.Argument
		*out = new(PriorityArgument)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PriorityPolicy.
func (in *PriorityPolicy) DeepCopy() *PriorityPolicy {
	if in == nil {
		return nil
	}
	out := new(PriorityPolicy)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RequestedToCapacityRatioArguments) DeepCopyInto(out *RequestedToCapacityRatioArguments) {
	*out = *in
	if in.UtilizationShape != nil {
		in, out := &in.UtilizationShape, &out.UtilizationShape
		*out = make([]UtilizationShapePoint, len(*in))
		copy(*out, *in)
	}
	if in.Resources != nil {
		in, out := &in.Resources, &out.Resources
		*out = make([]ResourceSpec, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RequestedToCapacityRatioArguments.
func (in *RequestedToCapacityRatioArguments) DeepCopy() *RequestedToCapacityRatioArguments {
	if in == nil {
		return nil
	}
	out := new(RequestedToCapacityRatioArguments)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ResourceSpec) DeepCopyInto(out *ResourceSpec) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ResourceSpec.
func (in *ResourceSpec) DeepCopy() *ResourceSpec {
	if in == nil {
		return nil
	}
	out := new(ResourceSpec)
	in.DeepCopyInto(out)
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

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ServiceAffinity.
func (in *ServiceAffinity) DeepCopy() *ServiceAffinity {
	if in == nil {
		return nil
	}
	out := new(ServiceAffinity)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ServiceAntiAffinity) DeepCopyInto(out *ServiceAntiAffinity) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ServiceAntiAffinity.
func (in *ServiceAntiAffinity) DeepCopy() *ServiceAntiAffinity {
	if in == nil {
		return nil
	}
	out := new(ServiceAntiAffinity)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *UtilizationShapePoint) DeepCopyInto(out *UtilizationShapePoint) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new UtilizationShapePoint.
func (in *UtilizationShapePoint) DeepCopy() *UtilizationShapePoint {
	if in == nil {
		return nil
	}
	out := new(UtilizationShapePoint)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Victims) DeepCopyInto(out *Victims) {
	*out = *in
	if in.Pods != nil {
		in, out := &in.Pods, &out.Pods
		*out = make([]*corev1.Pod, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(corev1.Pod)
				(*in).DeepCopyInto(*out)
			}
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Victims.
func (in *Victims) DeepCopy() *Victims {
	if in == nil {
		return nil
	}
	out := new(Victims)
	in.DeepCopyInto(out)
	return out
}
