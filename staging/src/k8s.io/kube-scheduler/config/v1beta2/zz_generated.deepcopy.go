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

package v1beta2

import (
	corev1 "k8s.io/api/core/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	v1 "k8s.io/kube-scheduler/config/v1"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DefaultPreemptionArgs) DeepCopyInto(out *DefaultPreemptionArgs) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	if in.MinCandidateNodesPercentage != nil {
		in, out := &in.MinCandidateNodesPercentage, &out.MinCandidateNodesPercentage
		*out = new(int32)
		**out = **in
	}
	if in.MinCandidateNodesAbsolute != nil {
		in, out := &in.MinCandidateNodesAbsolute, &out.MinCandidateNodesAbsolute
		*out = new(int32)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DefaultPreemptionArgs.
func (in *DefaultPreemptionArgs) DeepCopy() *DefaultPreemptionArgs {
	if in == nil {
		return nil
	}
	out := new(DefaultPreemptionArgs)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *DefaultPreemptionArgs) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Extender) DeepCopyInto(out *Extender) {
	*out = *in
	if in.TLSConfig != nil {
		in, out := &in.TLSConfig, &out.TLSConfig
		*out = new(v1.ExtenderTLSConfig)
		(*in).DeepCopyInto(*out)
	}
	out.HTTPTimeout = in.HTTPTimeout
	if in.ManagedResources != nil {
		in, out := &in.ManagedResources, &out.ManagedResources
		*out = make([]v1.ExtenderManagedResource, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Extender.
func (in *Extender) DeepCopy() *Extender {
	if in == nil {
		return nil
	}
	out := new(Extender)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *InterPodAffinityArgs) DeepCopyInto(out *InterPodAffinityArgs) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	if in.HardPodAffinityWeight != nil {
		in, out := &in.HardPodAffinityWeight, &out.HardPodAffinityWeight
		*out = new(int32)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new InterPodAffinityArgs.
func (in *InterPodAffinityArgs) DeepCopy() *InterPodAffinityArgs {
	if in == nil {
		return nil
	}
	out := new(InterPodAffinityArgs)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *InterPodAffinityArgs) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KubeSchedulerConfiguration) DeepCopyInto(out *KubeSchedulerConfiguration) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	if in.Parallelism != nil {
		in, out := &in.Parallelism, &out.Parallelism
		*out = new(int32)
		**out = **in
	}
	in.LeaderElection.DeepCopyInto(&out.LeaderElection)
	out.ClientConnection = in.ClientConnection
	if in.HealthzBindAddress != nil {
		in, out := &in.HealthzBindAddress, &out.HealthzBindAddress
		*out = new(string)
		**out = **in
	}
	if in.MetricsBindAddress != nil {
		in, out := &in.MetricsBindAddress, &out.MetricsBindAddress
		*out = new(string)
		**out = **in
	}
	in.DebuggingConfiguration.DeepCopyInto(&out.DebuggingConfiguration)
	if in.PercentageOfNodesToScore != nil {
		in, out := &in.PercentageOfNodesToScore, &out.PercentageOfNodesToScore
		*out = new(int32)
		**out = **in
	}
	if in.PodInitialBackoffSeconds != nil {
		in, out := &in.PodInitialBackoffSeconds, &out.PodInitialBackoffSeconds
		*out = new(int64)
		**out = **in
	}
	if in.PodMaxBackoffSeconds != nil {
		in, out := &in.PodMaxBackoffSeconds, &out.PodMaxBackoffSeconds
		*out = new(int64)
		**out = **in
	}
	if in.Profiles != nil {
		in, out := &in.Profiles, &out.Profiles
		*out = make([]KubeSchedulerProfile, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Extenders != nil {
		in, out := &in.Extenders, &out.Extenders
		*out = make([]Extender, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KubeSchedulerConfiguration.
func (in *KubeSchedulerConfiguration) DeepCopy() *KubeSchedulerConfiguration {
	if in == nil {
		return nil
	}
	out := new(KubeSchedulerConfiguration)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *KubeSchedulerConfiguration) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KubeSchedulerProfile) DeepCopyInto(out *KubeSchedulerProfile) {
	*out = *in
	if in.SchedulerName != nil {
		in, out := &in.SchedulerName, &out.SchedulerName
		*out = new(string)
		**out = **in
	}
	if in.Plugins != nil {
		in, out := &in.Plugins, &out.Plugins
		*out = new(Plugins)
		(*in).DeepCopyInto(*out)
	}
	if in.PluginConfig != nil {
		in, out := &in.PluginConfig, &out.PluginConfig
		*out = make([]PluginConfig, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KubeSchedulerProfile.
func (in *KubeSchedulerProfile) DeepCopy() *KubeSchedulerProfile {
	if in == nil {
		return nil
	}
	out := new(KubeSchedulerProfile)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NodeAffinityArgs) DeepCopyInto(out *NodeAffinityArgs) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	if in.AddedAffinity != nil {
		in, out := &in.AddedAffinity, &out.AddedAffinity
		*out = new(corev1.NodeAffinity)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NodeAffinityArgs.
func (in *NodeAffinityArgs) DeepCopy() *NodeAffinityArgs {
	if in == nil {
		return nil
	}
	out := new(NodeAffinityArgs)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *NodeAffinityArgs) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NodeResourcesBalancedAllocationArgs) DeepCopyInto(out *NodeResourcesBalancedAllocationArgs) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	if in.Resources != nil {
		in, out := &in.Resources, &out.Resources
		*out = make([]ResourceSpec, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NodeResourcesBalancedAllocationArgs.
func (in *NodeResourcesBalancedAllocationArgs) DeepCopy() *NodeResourcesBalancedAllocationArgs {
	if in == nil {
		return nil
	}
	out := new(NodeResourcesBalancedAllocationArgs)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *NodeResourcesBalancedAllocationArgs) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NodeResourcesFitArgs) DeepCopyInto(out *NodeResourcesFitArgs) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	if in.IgnoredResources != nil {
		in, out := &in.IgnoredResources, &out.IgnoredResources
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.IgnoredResourceGroups != nil {
		in, out := &in.IgnoredResourceGroups, &out.IgnoredResourceGroups
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.ScoringStrategy != nil {
		in, out := &in.ScoringStrategy, &out.ScoringStrategy
		*out = new(ScoringStrategy)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NodeResourcesFitArgs.
func (in *NodeResourcesFitArgs) DeepCopy() *NodeResourcesFitArgs {
	if in == nil {
		return nil
	}
	out := new(NodeResourcesFitArgs)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *NodeResourcesFitArgs) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Plugin) DeepCopyInto(out *Plugin) {
	*out = *in
	if in.Weight != nil {
		in, out := &in.Weight, &out.Weight
		*out = new(int32)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Plugin.
func (in *Plugin) DeepCopy() *Plugin {
	if in == nil {
		return nil
	}
	out := new(Plugin)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PluginConfig) DeepCopyInto(out *PluginConfig) {
	*out = *in
	in.Args.DeepCopyInto(&out.Args)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PluginConfig.
func (in *PluginConfig) DeepCopy() *PluginConfig {
	if in == nil {
		return nil
	}
	out := new(PluginConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PluginSet) DeepCopyInto(out *PluginSet) {
	*out = *in
	if in.Enabled != nil {
		in, out := &in.Enabled, &out.Enabled
		*out = make([]Plugin, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Disabled != nil {
		in, out := &in.Disabled, &out.Disabled
		*out = make([]Plugin, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PluginSet.
func (in *PluginSet) DeepCopy() *PluginSet {
	if in == nil {
		return nil
	}
	out := new(PluginSet)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Plugins) DeepCopyInto(out *Plugins) {
	*out = *in
	in.QueueSort.DeepCopyInto(&out.QueueSort)
	in.PreFilter.DeepCopyInto(&out.PreFilter)
	in.Filter.DeepCopyInto(&out.Filter)
	in.PostFilter.DeepCopyInto(&out.PostFilter)
	in.PreScore.DeepCopyInto(&out.PreScore)
	in.Score.DeepCopyInto(&out.Score)
	in.Reserve.DeepCopyInto(&out.Reserve)
	in.Permit.DeepCopyInto(&out.Permit)
	in.PreBind.DeepCopyInto(&out.PreBind)
	in.Bind.DeepCopyInto(&out.Bind)
	in.PostBind.DeepCopyInto(&out.PostBind)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Plugins.
func (in *Plugins) DeepCopy() *Plugins {
	if in == nil {
		return nil
	}
	out := new(Plugins)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PodTopologySpreadArgs) DeepCopyInto(out *PodTopologySpreadArgs) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	if in.DefaultConstraints != nil {
		in, out := &in.DefaultConstraints, &out.DefaultConstraints
		*out = make([]corev1.TopologySpreadConstraint, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PodTopologySpreadArgs.
func (in *PodTopologySpreadArgs) DeepCopy() *PodTopologySpreadArgs {
	if in == nil {
		return nil
	}
	out := new(PodTopologySpreadArgs)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *PodTopologySpreadArgs) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RequestedToCapacityRatioParam) DeepCopyInto(out *RequestedToCapacityRatioParam) {
	*out = *in
	if in.Shape != nil {
		in, out := &in.Shape, &out.Shape
		*out = make([]UtilizationShapePoint, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RequestedToCapacityRatioParam.
func (in *RequestedToCapacityRatioParam) DeepCopy() *RequestedToCapacityRatioParam {
	if in == nil {
		return nil
	}
	out := new(RequestedToCapacityRatioParam)
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
func (in *ScoringStrategy) DeepCopyInto(out *ScoringStrategy) {
	*out = *in
	if in.Resources != nil {
		in, out := &in.Resources, &out.Resources
		*out = make([]ResourceSpec, len(*in))
		copy(*out, *in)
	}
	if in.RequestedToCapacityRatio != nil {
		in, out := &in.RequestedToCapacityRatio, &out.RequestedToCapacityRatio
		*out = new(RequestedToCapacityRatioParam)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ScoringStrategy.
func (in *ScoringStrategy) DeepCopy() *ScoringStrategy {
	if in == nil {
		return nil
	}
	out := new(ScoringStrategy)
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
func (in *VolumeBindingArgs) DeepCopyInto(out *VolumeBindingArgs) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	if in.BindTimeoutSeconds != nil {
		in, out := &in.BindTimeoutSeconds, &out.BindTimeoutSeconds
		*out = new(int64)
		**out = **in
	}
	if in.Shape != nil {
		in, out := &in.Shape, &out.Shape
		*out = make([]UtilizationShapePoint, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VolumeBindingArgs.
func (in *VolumeBindingArgs) DeepCopy() *VolumeBindingArgs {
	if in == nil {
		return nil
	}
	out := new(VolumeBindingArgs)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *VolumeBindingArgs) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}
