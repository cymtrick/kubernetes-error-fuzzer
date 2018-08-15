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

package componentconfig

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AttachDetachControllerConfiguration) DeepCopyInto(out *AttachDetachControllerConfiguration) {
	*out = *in
	out.ReconcilerSyncLoopPeriod = in.ReconcilerSyncLoopPeriod
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AttachDetachControllerConfiguration.
func (in *AttachDetachControllerConfiguration) DeepCopy() *AttachDetachControllerConfiguration {
	if in == nil {
		return nil
	}
	out := new(AttachDetachControllerConfiguration)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CSRSigningControllerConfiguration) DeepCopyInto(out *CSRSigningControllerConfiguration) {
	*out = *in
	out.ClusterSigningDuration = in.ClusterSigningDuration
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CSRSigningControllerConfiguration.
func (in *CSRSigningControllerConfiguration) DeepCopy() *CSRSigningControllerConfiguration {
	if in == nil {
		return nil
	}
	out := new(CSRSigningControllerConfiguration)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CloudControllerManagerConfiguration) DeepCopyInto(out *CloudControllerManagerConfiguration) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.CloudProvider = in.CloudProvider
	out.Debugging = in.Debugging
	out.GenericComponent = in.GenericComponent
	out.KubeCloudShared = in.KubeCloudShared
	out.ServiceController = in.ServiceController
	out.NodeStatusUpdateFrequency = in.NodeStatusUpdateFrequency
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CloudControllerManagerConfiguration.
func (in *CloudControllerManagerConfiguration) DeepCopy() *CloudControllerManagerConfiguration {
	if in == nil {
		return nil
	}
	out := new(CloudControllerManagerConfiguration)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *CloudControllerManagerConfiguration) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CloudProviderConfiguration) DeepCopyInto(out *CloudProviderConfiguration) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CloudProviderConfiguration.
func (in *CloudProviderConfiguration) DeepCopy() *CloudProviderConfiguration {
	if in == nil {
		return nil
	}
	out := new(CloudProviderConfiguration)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DaemonSetControllerConfiguration) DeepCopyInto(out *DaemonSetControllerConfiguration) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DaemonSetControllerConfiguration.
func (in *DaemonSetControllerConfiguration) DeepCopy() *DaemonSetControllerConfiguration {
	if in == nil {
		return nil
	}
	out := new(DaemonSetControllerConfiguration)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DeploymentControllerConfiguration) DeepCopyInto(out *DeploymentControllerConfiguration) {
	*out = *in
	out.DeploymentControllerSyncPeriod = in.DeploymentControllerSyncPeriod
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DeploymentControllerConfiguration.
func (in *DeploymentControllerConfiguration) DeepCopy() *DeploymentControllerConfiguration {
	if in == nil {
		return nil
	}
	out := new(DeploymentControllerConfiguration)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DeprecatedControllerConfiguration) DeepCopyInto(out *DeprecatedControllerConfiguration) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DeprecatedControllerConfiguration.
func (in *DeprecatedControllerConfiguration) DeepCopy() *DeprecatedControllerConfiguration {
	if in == nil {
		return nil
	}
	out := new(DeprecatedControllerConfiguration)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EndPointControllerConfiguration) DeepCopyInto(out *EndPointControllerConfiguration) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EndPointControllerConfiguration.
func (in *EndPointControllerConfiguration) DeepCopy() *EndPointControllerConfiguration {
	if in == nil {
		return nil
	}
	out := new(EndPointControllerConfiguration)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GarbageCollectorControllerConfiguration) DeepCopyInto(out *GarbageCollectorControllerConfiguration) {
	*out = *in
	if in.GCIgnoredResources != nil {
		in, out := &in.GCIgnoredResources, &out.GCIgnoredResources
		*out = make([]GroupResource, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GarbageCollectorControllerConfiguration.
func (in *GarbageCollectorControllerConfiguration) DeepCopy() *GarbageCollectorControllerConfiguration {
	if in == nil {
		return nil
	}
	out := new(GarbageCollectorControllerConfiguration)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GenericComponentConfiguration) DeepCopyInto(out *GenericComponentConfiguration) {
	*out = *in
	out.MinResyncPeriod = in.MinResyncPeriod
	out.ControllerStartInterval = in.ControllerStartInterval
	out.LeaderElection = in.LeaderElection
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GenericComponentConfiguration.
func (in *GenericComponentConfiguration) DeepCopy() *GenericComponentConfiguration {
	if in == nil {
		return nil
	}
	out := new(GenericComponentConfiguration)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GroupResource) DeepCopyInto(out *GroupResource) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GroupResource.
func (in *GroupResource) DeepCopy() *GroupResource {
	if in == nil {
		return nil
	}
	out := new(GroupResource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HPAControllerConfiguration) DeepCopyInto(out *HPAControllerConfiguration) {
	*out = *in
	out.HorizontalPodAutoscalerSyncPeriod = in.HorizontalPodAutoscalerSyncPeriod
	out.HorizontalPodAutoscalerUpscaleForbiddenWindow = in.HorizontalPodAutoscalerUpscaleForbiddenWindow
	out.HorizontalPodAutoscalerDownscaleForbiddenWindow = in.HorizontalPodAutoscalerDownscaleForbiddenWindow
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HPAControllerConfiguration.
func (in *HPAControllerConfiguration) DeepCopy() *HPAControllerConfiguration {
	if in == nil {
		return nil
	}
	out := new(HPAControllerConfiguration)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IPPortVar) DeepCopyInto(out *IPPortVar) {
	*out = *in
	if in.Val != nil {
		in, out := &in.Val, &out.Val
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IPPortVar.
func (in *IPPortVar) DeepCopy() *IPPortVar {
	if in == nil {
		return nil
	}
	out := new(IPPortVar)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IPVar) DeepCopyInto(out *IPVar) {
	*out = *in
	if in.Val != nil {
		in, out := &in.Val, &out.Val
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IPVar.
func (in *IPVar) DeepCopy() *IPVar {
	if in == nil {
		return nil
	}
	out := new(IPVar)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *JobControllerConfiguration) DeepCopyInto(out *JobControllerConfiguration) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new JobControllerConfiguration.
func (in *JobControllerConfiguration) DeepCopy() *JobControllerConfiguration {
	if in == nil {
		return nil
	}
	out := new(JobControllerConfiguration)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KubeCloudSharedConfiguration) DeepCopyInto(out *KubeCloudSharedConfiguration) {
	*out = *in
	out.RouteReconciliationPeriod = in.RouteReconciliationPeriod
	out.NodeMonitorPeriod = in.NodeMonitorPeriod
	out.NodeSyncPeriod = in.NodeSyncPeriod
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KubeCloudSharedConfiguration.
func (in *KubeCloudSharedConfiguration) DeepCopy() *KubeCloudSharedConfiguration {
	if in == nil {
		return nil
	}
	out := new(KubeCloudSharedConfiguration)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KubeControllerManagerConfiguration) DeepCopyInto(out *KubeControllerManagerConfiguration) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.CloudProvider = in.CloudProvider
	out.Debugging = in.Debugging
	out.GenericComponent = in.GenericComponent
	out.KubeCloudShared = in.KubeCloudShared
	out.AttachDetachController = in.AttachDetachController
	out.CSRSigningController = in.CSRSigningController
	out.DaemonSetController = in.DaemonSetController
	out.DeploymentController = in.DeploymentController
	out.DeprecatedController = in.DeprecatedController
	out.EndPointController = in.EndPointController
	in.GarbageCollectorController.DeepCopyInto(&out.GarbageCollectorController)
	out.HPAController = in.HPAController
	out.JobController = in.JobController
	out.NamespaceController = in.NamespaceController
	out.NodeIpamController = in.NodeIpamController
	out.NodeLifecycleController = in.NodeLifecycleController
	out.PersistentVolumeBinderController = in.PersistentVolumeBinderController
	out.PodGCController = in.PodGCController
	out.ReplicaSetController = in.ReplicaSetController
	out.ReplicationController = in.ReplicationController
	out.ResourceQuotaController = in.ResourceQuotaController
	out.SAController = in.SAController
	out.ServiceController = in.ServiceController
	if in.Controllers != nil {
		in, out := &in.Controllers, &out.Controllers
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KubeControllerManagerConfiguration.
func (in *KubeControllerManagerConfiguration) DeepCopy() *KubeControllerManagerConfiguration {
	if in == nil {
		return nil
	}
	out := new(KubeControllerManagerConfiguration)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *KubeControllerManagerConfiguration) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KubeSchedulerConfiguration) DeepCopyInto(out *KubeSchedulerConfiguration) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.DebuggingConfiguration = in.DebuggingConfiguration
	in.AlgorithmSource.DeepCopyInto(&out.AlgorithmSource)
	out.LeaderElection = in.LeaderElection
	out.ClientConnection = in.ClientConnection
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
func (in *KubeSchedulerLeaderElectionConfiguration) DeepCopyInto(out *KubeSchedulerLeaderElectionConfiguration) {
	*out = *in
	out.LeaderElectionConfiguration = in.LeaderElectionConfiguration
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KubeSchedulerLeaderElectionConfiguration.
func (in *KubeSchedulerLeaderElectionConfiguration) DeepCopy() *KubeSchedulerLeaderElectionConfiguration {
	if in == nil {
		return nil
	}
	out := new(KubeSchedulerLeaderElectionConfiguration)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NamespaceControllerConfiguration) DeepCopyInto(out *NamespaceControllerConfiguration) {
	*out = *in
	out.NamespaceSyncPeriod = in.NamespaceSyncPeriod
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NamespaceControllerConfiguration.
func (in *NamespaceControllerConfiguration) DeepCopy() *NamespaceControllerConfiguration {
	if in == nil {
		return nil
	}
	out := new(NamespaceControllerConfiguration)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NodeIpamControllerConfiguration) DeepCopyInto(out *NodeIpamControllerConfiguration) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NodeIpamControllerConfiguration.
func (in *NodeIpamControllerConfiguration) DeepCopy() *NodeIpamControllerConfiguration {
	if in == nil {
		return nil
	}
	out := new(NodeIpamControllerConfiguration)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NodeLifecycleControllerConfiguration) DeepCopyInto(out *NodeLifecycleControllerConfiguration) {
	*out = *in
	out.NodeStartupGracePeriod = in.NodeStartupGracePeriod
	out.NodeMonitorGracePeriod = in.NodeMonitorGracePeriod
	out.PodEvictionTimeout = in.PodEvictionTimeout
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NodeLifecycleControllerConfiguration.
func (in *NodeLifecycleControllerConfiguration) DeepCopy() *NodeLifecycleControllerConfiguration {
	if in == nil {
		return nil
	}
	out := new(NodeLifecycleControllerConfiguration)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PersistentVolumeBinderControllerConfiguration) DeepCopyInto(out *PersistentVolumeBinderControllerConfiguration) {
	*out = *in
	out.PVClaimBinderSyncPeriod = in.PVClaimBinderSyncPeriod
	out.VolumeConfiguration = in.VolumeConfiguration
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PersistentVolumeBinderControllerConfiguration.
func (in *PersistentVolumeBinderControllerConfiguration) DeepCopy() *PersistentVolumeBinderControllerConfiguration {
	if in == nil {
		return nil
	}
	out := new(PersistentVolumeBinderControllerConfiguration)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PersistentVolumeRecyclerConfiguration) DeepCopyInto(out *PersistentVolumeRecyclerConfiguration) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PersistentVolumeRecyclerConfiguration.
func (in *PersistentVolumeRecyclerConfiguration) DeepCopy() *PersistentVolumeRecyclerConfiguration {
	if in == nil {
		return nil
	}
	out := new(PersistentVolumeRecyclerConfiguration)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PodGCControllerConfiguration) DeepCopyInto(out *PodGCControllerConfiguration) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PodGCControllerConfiguration.
func (in *PodGCControllerConfiguration) DeepCopy() *PodGCControllerConfiguration {
	if in == nil {
		return nil
	}
	out := new(PodGCControllerConfiguration)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PortRangeVar) DeepCopyInto(out *PortRangeVar) {
	*out = *in
	if in.Val != nil {
		in, out := &in.Val, &out.Val
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PortRangeVar.
func (in *PortRangeVar) DeepCopy() *PortRangeVar {
	if in == nil {
		return nil
	}
	out := new(PortRangeVar)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ReplicaSetControllerConfiguration) DeepCopyInto(out *ReplicaSetControllerConfiguration) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ReplicaSetControllerConfiguration.
func (in *ReplicaSetControllerConfiguration) DeepCopy() *ReplicaSetControllerConfiguration {
	if in == nil {
		return nil
	}
	out := new(ReplicaSetControllerConfiguration)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ReplicationControllerConfiguration) DeepCopyInto(out *ReplicationControllerConfiguration) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ReplicationControllerConfiguration.
func (in *ReplicationControllerConfiguration) DeepCopy() *ReplicationControllerConfiguration {
	if in == nil {
		return nil
	}
	out := new(ReplicationControllerConfiguration)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ResourceQuotaControllerConfiguration) DeepCopyInto(out *ResourceQuotaControllerConfiguration) {
	*out = *in
	out.ResourceQuotaSyncPeriod = in.ResourceQuotaSyncPeriod
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ResourceQuotaControllerConfiguration.
func (in *ResourceQuotaControllerConfiguration) DeepCopy() *ResourceQuotaControllerConfiguration {
	if in == nil {
		return nil
	}
	out := new(ResourceQuotaControllerConfiguration)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SAControllerConfiguration) DeepCopyInto(out *SAControllerConfiguration) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SAControllerConfiguration.
func (in *SAControllerConfiguration) DeepCopy() *SAControllerConfiguration {
	if in == nil {
		return nil
	}
	out := new(SAControllerConfiguration)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SchedulerAlgorithmSource) DeepCopyInto(out *SchedulerAlgorithmSource) {
	*out = *in
	if in.Policy != nil {
		in, out := &in.Policy, &out.Policy
		*out = new(SchedulerPolicySource)
		(*in).DeepCopyInto(*out)
	}
	if in.Provider != nil {
		in, out := &in.Provider, &out.Provider
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SchedulerAlgorithmSource.
func (in *SchedulerAlgorithmSource) DeepCopy() *SchedulerAlgorithmSource {
	if in == nil {
		return nil
	}
	out := new(SchedulerAlgorithmSource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SchedulerPolicyConfigMapSource) DeepCopyInto(out *SchedulerPolicyConfigMapSource) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SchedulerPolicyConfigMapSource.
func (in *SchedulerPolicyConfigMapSource) DeepCopy() *SchedulerPolicyConfigMapSource {
	if in == nil {
		return nil
	}
	out := new(SchedulerPolicyConfigMapSource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SchedulerPolicyFileSource) DeepCopyInto(out *SchedulerPolicyFileSource) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SchedulerPolicyFileSource.
func (in *SchedulerPolicyFileSource) DeepCopy() *SchedulerPolicyFileSource {
	if in == nil {
		return nil
	}
	out := new(SchedulerPolicyFileSource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SchedulerPolicySource) DeepCopyInto(out *SchedulerPolicySource) {
	*out = *in
	if in.File != nil {
		in, out := &in.File, &out.File
		*out = new(SchedulerPolicyFileSource)
		**out = **in
	}
	if in.ConfigMap != nil {
		in, out := &in.ConfigMap, &out.ConfigMap
		*out = new(SchedulerPolicyConfigMapSource)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SchedulerPolicySource.
func (in *SchedulerPolicySource) DeepCopy() *SchedulerPolicySource {
	if in == nil {
		return nil
	}
	out := new(SchedulerPolicySource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ServiceControllerConfiguration) DeepCopyInto(out *ServiceControllerConfiguration) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ServiceControllerConfiguration.
func (in *ServiceControllerConfiguration) DeepCopy() *ServiceControllerConfiguration {
	if in == nil {
		return nil
	}
	out := new(ServiceControllerConfiguration)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VolumeConfiguration) DeepCopyInto(out *VolumeConfiguration) {
	*out = *in
	out.PersistentVolumeRecyclerConfiguration = in.PersistentVolumeRecyclerConfiguration
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VolumeConfiguration.
func (in *VolumeConfiguration) DeepCopy() *VolumeConfiguration {
	if in == nil {
		return nil
	}
	out := new(VolumeConfiguration)
	in.DeepCopyInto(out)
	return out
}
