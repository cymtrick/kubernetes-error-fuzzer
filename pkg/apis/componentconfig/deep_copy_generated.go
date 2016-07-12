// +build !ignore_autogenerated

/*
Copyright 2016 The Kubernetes Authors.

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

package componentconfig

import (
	api "k8s.io/kubernetes/pkg/api"
	conversion "k8s.io/kubernetes/pkg/conversion"
	reflect "reflect"
)

func init() {
	if err := api.Scheme.AddGeneratedDeepCopyFuncs(
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_componentconfig_IPVar, InType: reflect.TypeOf(func() *IPVar { var x *IPVar; return x }())},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_componentconfig_KubeControllerManagerConfiguration, InType: reflect.TypeOf(func() *KubeControllerManagerConfiguration { var x *KubeControllerManagerConfiguration; return x }())},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_componentconfig_KubeProxyConfiguration, InType: reflect.TypeOf(func() *KubeProxyConfiguration { var x *KubeProxyConfiguration; return x }())},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_componentconfig_KubeSchedulerConfiguration, InType: reflect.TypeOf(func() *KubeSchedulerConfiguration { var x *KubeSchedulerConfiguration; return x }())},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_componentconfig_KubeletConfiguration, InType: reflect.TypeOf(func() *KubeletConfiguration { var x *KubeletConfiguration; return x }())},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_componentconfig_LeaderElectionConfiguration, InType: reflect.TypeOf(func() *LeaderElectionConfiguration { var x *LeaderElectionConfiguration; return x }())},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_componentconfig_PersistentVolumeRecyclerConfiguration, InType: reflect.TypeOf(func() *PersistentVolumeRecyclerConfiguration { var x *PersistentVolumeRecyclerConfiguration; return x }())},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_componentconfig_PortRangeVar, InType: reflect.TypeOf(func() *PortRangeVar { var x *PortRangeVar; return x }())},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_componentconfig_VolumeConfiguration, InType: reflect.TypeOf(func() *VolumeConfiguration { var x *VolumeConfiguration; return x }())},
	); err != nil {
		// if one of the deep copy functions is malformed, detect it immediately.
		panic(err)
	}
}

func DeepCopy_componentconfig_IPVar(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*IPVar)
		out := out.(*IPVar)
		if in.Val != nil {
			in, out := &in.Val, &out.Val
			*out = new(string)
			**out = **in
		} else {
			out.Val = nil
		}
		return nil
	}
}

func DeepCopy_componentconfig_KubeControllerManagerConfiguration(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*KubeControllerManagerConfiguration)
		out := out.(*KubeControllerManagerConfiguration)
		out.TypeMeta = in.TypeMeta
		out.Port = in.Port
		out.Address = in.Address
		out.CloudProvider = in.CloudProvider
		out.CloudConfigFile = in.CloudConfigFile
		out.ConcurrentEndpointSyncs = in.ConcurrentEndpointSyncs
		out.ConcurrentRSSyncs = in.ConcurrentRSSyncs
		out.ConcurrentRCSyncs = in.ConcurrentRCSyncs
		out.ConcurrentResourceQuotaSyncs = in.ConcurrentResourceQuotaSyncs
		out.ConcurrentDeploymentSyncs = in.ConcurrentDeploymentSyncs
		out.ConcurrentDaemonSetSyncs = in.ConcurrentDaemonSetSyncs
		out.ConcurrentJobSyncs = in.ConcurrentJobSyncs
		out.ConcurrentNamespaceSyncs = in.ConcurrentNamespaceSyncs
		out.ConcurrentSATokenSyncs = in.ConcurrentSATokenSyncs
		out.LookupCacheSizeForRC = in.LookupCacheSizeForRC
		out.LookupCacheSizeForRS = in.LookupCacheSizeForRS
		out.LookupCacheSizeForDaemonSet = in.LookupCacheSizeForDaemonSet
		out.ServiceSyncPeriod = in.ServiceSyncPeriod
		out.NodeSyncPeriod = in.NodeSyncPeriod
		out.ResourceQuotaSyncPeriod = in.ResourceQuotaSyncPeriod
		out.NamespaceSyncPeriod = in.NamespaceSyncPeriod
		out.PVClaimBinderSyncPeriod = in.PVClaimBinderSyncPeriod
		out.MinResyncPeriod = in.MinResyncPeriod
		out.TerminatedPodGCThreshold = in.TerminatedPodGCThreshold
		out.HorizontalPodAutoscalerSyncPeriod = in.HorizontalPodAutoscalerSyncPeriod
		out.DeploymentControllerSyncPeriod = in.DeploymentControllerSyncPeriod
		out.PodEvictionTimeout = in.PodEvictionTimeout
		out.DeletingPodsQps = in.DeletingPodsQps
		out.DeletingPodsBurst = in.DeletingPodsBurst
		out.NodeMonitorGracePeriod = in.NodeMonitorGracePeriod
		out.RegisterRetryCount = in.RegisterRetryCount
		out.NodeStartupGracePeriod = in.NodeStartupGracePeriod
		out.NodeMonitorPeriod = in.NodeMonitorPeriod
		out.ServiceAccountKeyFile = in.ServiceAccountKeyFile
		out.EnableProfiling = in.EnableProfiling
		out.ClusterName = in.ClusterName
		out.ClusterCIDR = in.ClusterCIDR
		out.ServiceCIDR = in.ServiceCIDR
		out.NodeCIDRMaskSize = in.NodeCIDRMaskSize
		out.AllocateNodeCIDRs = in.AllocateNodeCIDRs
		out.ConfigureCloudRoutes = in.ConfigureCloudRoutes
		out.RootCAFile = in.RootCAFile
		out.ContentType = in.ContentType
		out.KubeAPIQPS = in.KubeAPIQPS
		out.KubeAPIBurst = in.KubeAPIBurst
		out.LeaderElection = in.LeaderElection
		out.VolumeConfiguration = in.VolumeConfiguration
		out.ControllerStartInterval = in.ControllerStartInterval
		out.EnableGarbageCollector = in.EnableGarbageCollector
		return nil
	}
}

func DeepCopy_componentconfig_KubeProxyConfiguration(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*KubeProxyConfiguration)
		out := out.(*KubeProxyConfiguration)
		out.TypeMeta = in.TypeMeta
		out.BindAddress = in.BindAddress
		out.ClusterCIDR = in.ClusterCIDR
		out.HealthzBindAddress = in.HealthzBindAddress
		out.HealthzPort = in.HealthzPort
		out.HostnameOverride = in.HostnameOverride
		if in.IPTablesMasqueradeBit != nil {
			in, out := &in.IPTablesMasqueradeBit, &out.IPTablesMasqueradeBit
			*out = new(int32)
			**out = **in
		} else {
			out.IPTablesMasqueradeBit = nil
		}
		out.IPTablesSyncPeriod = in.IPTablesSyncPeriod
		out.KubeconfigPath = in.KubeconfigPath
		out.MasqueradeAll = in.MasqueradeAll
		out.Master = in.Master
		if in.OOMScoreAdj != nil {
			in, out := &in.OOMScoreAdj, &out.OOMScoreAdj
			*out = new(int32)
			**out = **in
		} else {
			out.OOMScoreAdj = nil
		}
		out.Mode = in.Mode
		out.PortRange = in.PortRange
		out.ResourceContainer = in.ResourceContainer
		out.UDPIdleTimeout = in.UDPIdleTimeout
		out.ConntrackMax = in.ConntrackMax
		out.ConntrackTCPEstablishedTimeout = in.ConntrackTCPEstablishedTimeout
		return nil
	}
}

func DeepCopy_componentconfig_KubeSchedulerConfiguration(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*KubeSchedulerConfiguration)
		out := out.(*KubeSchedulerConfiguration)
		out.TypeMeta = in.TypeMeta
		out.Port = in.Port
		out.Address = in.Address
		out.AlgorithmProvider = in.AlgorithmProvider
		out.PolicyConfigFile = in.PolicyConfigFile
		out.EnableProfiling = in.EnableProfiling
		out.ContentType = in.ContentType
		out.KubeAPIQPS = in.KubeAPIQPS
		out.KubeAPIBurst = in.KubeAPIBurst
		out.SchedulerName = in.SchedulerName
		out.HardPodAffinitySymmetricWeight = in.HardPodAffinitySymmetricWeight
		out.FailureDomains = in.FailureDomains
		out.LeaderElection = in.LeaderElection
		return nil
	}
}

func DeepCopy_componentconfig_KubeletConfiguration(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*KubeletConfiguration)
		out := out.(*KubeletConfiguration)
		out.Config = in.Config
		out.SyncFrequency = in.SyncFrequency
		out.FileCheckFrequency = in.FileCheckFrequency
		out.HTTPCheckFrequency = in.HTTPCheckFrequency
		out.ManifestURL = in.ManifestURL
		out.ManifestURLHeader = in.ManifestURLHeader
		out.EnableServer = in.EnableServer
		out.Address = in.Address
		out.Port = in.Port
		out.ReadOnlyPort = in.ReadOnlyPort
		out.TLSCertFile = in.TLSCertFile
		out.TLSPrivateKeyFile = in.TLSPrivateKeyFile
		out.CertDirectory = in.CertDirectory
		out.HostnameOverride = in.HostnameOverride
		out.PodInfraContainerImage = in.PodInfraContainerImage
		out.DockerEndpoint = in.DockerEndpoint
		out.RootDirectory = in.RootDirectory
		out.SeccompProfileRoot = in.SeccompProfileRoot
		out.AllowPrivileged = in.AllowPrivileged
		out.HostNetworkSources = in.HostNetworkSources
		out.HostPIDSources = in.HostPIDSources
		out.HostIPCSources = in.HostIPCSources
		out.RegistryPullQPS = in.RegistryPullQPS
		out.RegistryBurst = in.RegistryBurst
		out.EventRecordQPS = in.EventRecordQPS
		out.EventBurst = in.EventBurst
		out.EnableDebuggingHandlers = in.EnableDebuggingHandlers
		out.MinimumGCAge = in.MinimumGCAge
		out.MaxPerPodContainerCount = in.MaxPerPodContainerCount
		out.MaxContainerCount = in.MaxContainerCount
		out.CAdvisorPort = in.CAdvisorPort
		out.HealthzPort = in.HealthzPort
		out.HealthzBindAddress = in.HealthzBindAddress
		out.OOMScoreAdj = in.OOMScoreAdj
		out.RegisterNode = in.RegisterNode
		out.ClusterDomain = in.ClusterDomain
		out.MasterServiceNamespace = in.MasterServiceNamespace
		out.ClusterDNS = in.ClusterDNS
		out.StreamingConnectionIdleTimeout = in.StreamingConnectionIdleTimeout
		out.NodeStatusUpdateFrequency = in.NodeStatusUpdateFrequency
		out.ImageMinimumGCAge = in.ImageMinimumGCAge
		out.ImageGCHighThresholdPercent = in.ImageGCHighThresholdPercent
		out.ImageGCLowThresholdPercent = in.ImageGCLowThresholdPercent
		out.LowDiskSpaceThresholdMB = in.LowDiskSpaceThresholdMB
		out.VolumeStatsAggPeriod = in.VolumeStatsAggPeriod
		out.NetworkPluginName = in.NetworkPluginName
		out.NetworkPluginDir = in.NetworkPluginDir
		out.VolumePluginDir = in.VolumePluginDir
		out.CloudProvider = in.CloudProvider
		out.CloudConfigFile = in.CloudConfigFile
		out.KubeletCgroups = in.KubeletCgroups
		out.RuntimeCgroups = in.RuntimeCgroups
		out.SystemCgroups = in.SystemCgroups
		out.CgroupRoot = in.CgroupRoot
		out.ContainerRuntime = in.ContainerRuntime
		out.RuntimeRequestTimeout = in.RuntimeRequestTimeout
		out.RktPath = in.RktPath
		out.RktAPIEndpoint = in.RktAPIEndpoint
		out.RktStage1Image = in.RktStage1Image
		out.LockFilePath = in.LockFilePath
		out.ExitOnLockContention = in.ExitOnLockContention
		out.ConfigureCBR0 = in.ConfigureCBR0
		out.HairpinMode = in.HairpinMode
		out.BabysitDaemons = in.BabysitDaemons
		out.MaxPods = in.MaxPods
		out.NvidiaGPUs = in.NvidiaGPUs
		out.DockerExecHandlerName = in.DockerExecHandlerName
		out.PodCIDR = in.PodCIDR
		out.ResolverConfig = in.ResolverConfig
		out.CPUCFSQuota = in.CPUCFSQuota
		out.Containerized = in.Containerized
		out.MaxOpenFiles = in.MaxOpenFiles
		out.ReconcileCIDR = in.ReconcileCIDR
		out.RegisterSchedulable = in.RegisterSchedulable
		out.ContentType = in.ContentType
		out.KubeAPIQPS = in.KubeAPIQPS
		out.KubeAPIBurst = in.KubeAPIBurst
		out.SerializeImagePulls = in.SerializeImagePulls
		out.ExperimentalFlannelOverlay = in.ExperimentalFlannelOverlay
		out.OutOfDiskTransitionFrequency = in.OutOfDiskTransitionFrequency
		out.NodeIP = in.NodeIP
		if in.NodeLabels != nil {
			in, out := &in.NodeLabels, &out.NodeLabels
			*out = make(map[string]string)
			for key, val := range *in {
				(*out)[key] = val
			}
		} else {
			out.NodeLabels = nil
		}
		out.NonMasqueradeCIDR = in.NonMasqueradeCIDR
		out.EnableCustomMetrics = in.EnableCustomMetrics
		out.EvictionHard = in.EvictionHard
		out.EvictionSoft = in.EvictionSoft
		out.EvictionSoftGracePeriod = in.EvictionSoftGracePeriod
		out.EvictionPressureTransitionPeriod = in.EvictionPressureTransitionPeriod
		out.EvictionMaxPodGracePeriod = in.EvictionMaxPodGracePeriod
		out.PodsPerCore = in.PodsPerCore
		out.EnableControllerAttachDetach = in.EnableControllerAttachDetach
		return nil
	}
}

func DeepCopy_componentconfig_LeaderElectionConfiguration(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*LeaderElectionConfiguration)
		out := out.(*LeaderElectionConfiguration)
		out.LeaderElect = in.LeaderElect
		out.LeaseDuration = in.LeaseDuration
		out.RenewDeadline = in.RenewDeadline
		out.RetryPeriod = in.RetryPeriod
		return nil
	}
}

func DeepCopy_componentconfig_PersistentVolumeRecyclerConfiguration(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*PersistentVolumeRecyclerConfiguration)
		out := out.(*PersistentVolumeRecyclerConfiguration)
		out.MaximumRetry = in.MaximumRetry
		out.MinimumTimeoutNFS = in.MinimumTimeoutNFS
		out.PodTemplateFilePathNFS = in.PodTemplateFilePathNFS
		out.IncrementTimeoutNFS = in.IncrementTimeoutNFS
		out.PodTemplateFilePathHostPath = in.PodTemplateFilePathHostPath
		out.MinimumTimeoutHostPath = in.MinimumTimeoutHostPath
		out.IncrementTimeoutHostPath = in.IncrementTimeoutHostPath
		return nil
	}
}

func DeepCopy_componentconfig_PortRangeVar(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*PortRangeVar)
		out := out.(*PortRangeVar)
		if in.Val != nil {
			in, out := &in.Val, &out.Val
			*out = new(string)
			**out = **in
		} else {
			out.Val = nil
		}
		return nil
	}
}

func DeepCopy_componentconfig_VolumeConfiguration(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*VolumeConfiguration)
		out := out.(*VolumeConfiguration)
		out.EnableHostPathProvisioning = in.EnableHostPathProvisioning
		out.EnableDynamicProvisioning = in.EnableDynamicProvisioning
		out.PersistentVolumeRecyclerConfiguration = in.PersistentVolumeRecyclerConfiguration
		out.FlexVolumePluginDir = in.FlexVolumePluginDir
		return nil
	}
}
