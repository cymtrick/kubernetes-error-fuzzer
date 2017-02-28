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

// This file was autogenerated by conversion-gen. Do not edit it manually!

package v1alpha1

import (
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	conversion "k8s.io/apimachinery/pkg/conversion"
	runtime "k8s.io/apimachinery/pkg/runtime"
	api "k8s.io/kubernetes/pkg/api"
	api_v1 "k8s.io/kubernetes/pkg/api/v1"
	componentconfig "k8s.io/kubernetes/pkg/apis/componentconfig"
	unsafe "unsafe"
)

func init() {
	SchemeBuilder.Register(RegisterConversions)
}

// RegisterConversions adds conversion functions to the given scheme.
// Public to allow building arbitrary schemes.
func RegisterConversions(scheme *runtime.Scheme) error {
	return scheme.AddGeneratedConversionFuncs(
		Convert_v1alpha1_KubeProxyConfiguration_To_componentconfig_KubeProxyConfiguration,
		Convert_componentconfig_KubeProxyConfiguration_To_v1alpha1_KubeProxyConfiguration,
		Convert_v1alpha1_KubeSchedulerConfiguration_To_componentconfig_KubeSchedulerConfiguration,
		Convert_componentconfig_KubeSchedulerConfiguration_To_v1alpha1_KubeSchedulerConfiguration,
		Convert_v1alpha1_KubeletAnonymousAuthentication_To_componentconfig_KubeletAnonymousAuthentication,
		Convert_componentconfig_KubeletAnonymousAuthentication_To_v1alpha1_KubeletAnonymousAuthentication,
		Convert_v1alpha1_KubeletAuthentication_To_componentconfig_KubeletAuthentication,
		Convert_componentconfig_KubeletAuthentication_To_v1alpha1_KubeletAuthentication,
		Convert_v1alpha1_KubeletAuthorization_To_componentconfig_KubeletAuthorization,
		Convert_componentconfig_KubeletAuthorization_To_v1alpha1_KubeletAuthorization,
		Convert_v1alpha1_KubeletConfiguration_To_componentconfig_KubeletConfiguration,
		Convert_componentconfig_KubeletConfiguration_To_v1alpha1_KubeletConfiguration,
		Convert_v1alpha1_KubeletWebhookAuthentication_To_componentconfig_KubeletWebhookAuthentication,
		Convert_componentconfig_KubeletWebhookAuthentication_To_v1alpha1_KubeletWebhookAuthentication,
		Convert_v1alpha1_KubeletWebhookAuthorization_To_componentconfig_KubeletWebhookAuthorization,
		Convert_componentconfig_KubeletWebhookAuthorization_To_v1alpha1_KubeletWebhookAuthorization,
		Convert_v1alpha1_KubeletX509Authentication_To_componentconfig_KubeletX509Authentication,
		Convert_componentconfig_KubeletX509Authentication_To_v1alpha1_KubeletX509Authentication,
		Convert_v1alpha1_LeaderElectionConfiguration_To_componentconfig_LeaderElectionConfiguration,
		Convert_componentconfig_LeaderElectionConfiguration_To_v1alpha1_LeaderElectionConfiguration,
	)
}

func autoConvert_v1alpha1_KubeProxyConfiguration_To_componentconfig_KubeProxyConfiguration(in *KubeProxyConfiguration, out *componentconfig.KubeProxyConfiguration, s conversion.Scope) error {
	out.BindAddress = in.BindAddress
	out.ClusterCIDR = in.ClusterCIDR
	out.HealthzBindAddress = in.HealthzBindAddress
	out.HealthzPort = in.HealthzPort
	out.HostnameOverride = in.HostnameOverride
	out.IPTablesMasqueradeBit = (*int32)(unsafe.Pointer(in.IPTablesMasqueradeBit))
	out.IPTablesSyncPeriod = in.IPTablesSyncPeriod
	out.IPTablesMinSyncPeriod = in.IPTablesMinSyncPeriod
	out.KubeconfigPath = in.KubeconfigPath
	out.MasqueradeAll = in.MasqueradeAll
	out.Master = in.Master
	out.OOMScoreAdj = (*int32)(unsafe.Pointer(in.OOMScoreAdj))
	out.Mode = componentconfig.ProxyMode(in.Mode)
	out.PortRange = in.PortRange
	out.ResourceContainer = in.ResourceContainer
	out.UDPIdleTimeout = in.UDPIdleTimeout
	out.ConntrackMax = in.ConntrackMax
	out.ConntrackMaxPerCore = in.ConntrackMaxPerCore
	out.ConntrackMin = in.ConntrackMin
	out.ConntrackTCPEstablishedTimeout = in.ConntrackTCPEstablishedTimeout
	out.ConntrackTCPCloseWaitTimeout = in.ConntrackTCPCloseWaitTimeout
	return nil
}

func Convert_v1alpha1_KubeProxyConfiguration_To_componentconfig_KubeProxyConfiguration(in *KubeProxyConfiguration, out *componentconfig.KubeProxyConfiguration, s conversion.Scope) error {
	return autoConvert_v1alpha1_KubeProxyConfiguration_To_componentconfig_KubeProxyConfiguration(in, out, s)
}

func autoConvert_componentconfig_KubeProxyConfiguration_To_v1alpha1_KubeProxyConfiguration(in *componentconfig.KubeProxyConfiguration, out *KubeProxyConfiguration, s conversion.Scope) error {
	out.BindAddress = in.BindAddress
	out.ClusterCIDR = in.ClusterCIDR
	out.HealthzBindAddress = in.HealthzBindAddress
	out.HealthzPort = in.HealthzPort
	out.HostnameOverride = in.HostnameOverride
	out.IPTablesMasqueradeBit = (*int32)(unsafe.Pointer(in.IPTablesMasqueradeBit))
	out.IPTablesSyncPeriod = in.IPTablesSyncPeriod
	out.IPTablesMinSyncPeriod = in.IPTablesMinSyncPeriod
	out.KubeconfigPath = in.KubeconfigPath
	out.MasqueradeAll = in.MasqueradeAll
	out.Master = in.Master
	out.OOMScoreAdj = (*int32)(unsafe.Pointer(in.OOMScoreAdj))
	out.Mode = ProxyMode(in.Mode)
	out.PortRange = in.PortRange
	out.ResourceContainer = in.ResourceContainer
	out.UDPIdleTimeout = in.UDPIdleTimeout
	out.ConntrackMax = in.ConntrackMax
	out.ConntrackMaxPerCore = in.ConntrackMaxPerCore
	out.ConntrackMin = in.ConntrackMin
	out.ConntrackTCPEstablishedTimeout = in.ConntrackTCPEstablishedTimeout
	out.ConntrackTCPCloseWaitTimeout = in.ConntrackTCPCloseWaitTimeout
	return nil
}

func Convert_componentconfig_KubeProxyConfiguration_To_v1alpha1_KubeProxyConfiguration(in *componentconfig.KubeProxyConfiguration, out *KubeProxyConfiguration, s conversion.Scope) error {
	return autoConvert_componentconfig_KubeProxyConfiguration_To_v1alpha1_KubeProxyConfiguration(in, out, s)
}

func autoConvert_v1alpha1_KubeSchedulerConfiguration_To_componentconfig_KubeSchedulerConfiguration(in *KubeSchedulerConfiguration, out *componentconfig.KubeSchedulerConfiguration, s conversion.Scope) error {
	out.Port = int32(in.Port)
	out.Address = in.Address
	out.AlgorithmProvider = in.AlgorithmProvider
	out.PolicyConfigFile = in.PolicyConfigFile
	if err := v1.Convert_Pointer_bool_To_bool(&in.EnableProfiling, &out.EnableProfiling, s); err != nil {
		return err
	}
	out.EnableContentionProfiling = in.EnableContentionProfiling
	out.ContentType = in.ContentType
	out.KubeAPIQPS = in.KubeAPIQPS
	out.KubeAPIBurst = int32(in.KubeAPIBurst)
	out.SchedulerName = in.SchedulerName
	out.HardPodAffinitySymmetricWeight = in.HardPodAffinitySymmetricWeight
	out.FailureDomains = in.FailureDomains
	if err := Convert_v1alpha1_LeaderElectionConfiguration_To_componentconfig_LeaderElectionConfiguration(&in.LeaderElection, &out.LeaderElection, s); err != nil {
		return err
	}
	return nil
}

func Convert_v1alpha1_KubeSchedulerConfiguration_To_componentconfig_KubeSchedulerConfiguration(in *KubeSchedulerConfiguration, out *componentconfig.KubeSchedulerConfiguration, s conversion.Scope) error {
	return autoConvert_v1alpha1_KubeSchedulerConfiguration_To_componentconfig_KubeSchedulerConfiguration(in, out, s)
}

func autoConvert_componentconfig_KubeSchedulerConfiguration_To_v1alpha1_KubeSchedulerConfiguration(in *componentconfig.KubeSchedulerConfiguration, out *KubeSchedulerConfiguration, s conversion.Scope) error {
	out.Port = int(in.Port)
	out.Address = in.Address
	out.AlgorithmProvider = in.AlgorithmProvider
	out.PolicyConfigFile = in.PolicyConfigFile
	if err := v1.Convert_bool_To_Pointer_bool(&in.EnableProfiling, &out.EnableProfiling, s); err != nil {
		return err
	}
	out.EnableContentionProfiling = in.EnableContentionProfiling
	out.ContentType = in.ContentType
	out.KubeAPIQPS = in.KubeAPIQPS
	out.KubeAPIBurst = int(in.KubeAPIBurst)
	out.SchedulerName = in.SchedulerName
	out.HardPodAffinitySymmetricWeight = in.HardPodAffinitySymmetricWeight
	out.FailureDomains = in.FailureDomains
	if err := Convert_componentconfig_LeaderElectionConfiguration_To_v1alpha1_LeaderElectionConfiguration(&in.LeaderElection, &out.LeaderElection, s); err != nil {
		return err
	}
	return nil
}

func Convert_componentconfig_KubeSchedulerConfiguration_To_v1alpha1_KubeSchedulerConfiguration(in *componentconfig.KubeSchedulerConfiguration, out *KubeSchedulerConfiguration, s conversion.Scope) error {
	return autoConvert_componentconfig_KubeSchedulerConfiguration_To_v1alpha1_KubeSchedulerConfiguration(in, out, s)
}

func autoConvert_v1alpha1_KubeletAnonymousAuthentication_To_componentconfig_KubeletAnonymousAuthentication(in *KubeletAnonymousAuthentication, out *componentconfig.KubeletAnonymousAuthentication, s conversion.Scope) error {
	if err := v1.Convert_Pointer_bool_To_bool(&in.Enabled, &out.Enabled, s); err != nil {
		return err
	}
	return nil
}

func Convert_v1alpha1_KubeletAnonymousAuthentication_To_componentconfig_KubeletAnonymousAuthentication(in *KubeletAnonymousAuthentication, out *componentconfig.KubeletAnonymousAuthentication, s conversion.Scope) error {
	return autoConvert_v1alpha1_KubeletAnonymousAuthentication_To_componentconfig_KubeletAnonymousAuthentication(in, out, s)
}

func autoConvert_componentconfig_KubeletAnonymousAuthentication_To_v1alpha1_KubeletAnonymousAuthentication(in *componentconfig.KubeletAnonymousAuthentication, out *KubeletAnonymousAuthentication, s conversion.Scope) error {
	if err := v1.Convert_bool_To_Pointer_bool(&in.Enabled, &out.Enabled, s); err != nil {
		return err
	}
	return nil
}

func Convert_componentconfig_KubeletAnonymousAuthentication_To_v1alpha1_KubeletAnonymousAuthentication(in *componentconfig.KubeletAnonymousAuthentication, out *KubeletAnonymousAuthentication, s conversion.Scope) error {
	return autoConvert_componentconfig_KubeletAnonymousAuthentication_To_v1alpha1_KubeletAnonymousAuthentication(in, out, s)
}

func autoConvert_v1alpha1_KubeletAuthentication_To_componentconfig_KubeletAuthentication(in *KubeletAuthentication, out *componentconfig.KubeletAuthentication, s conversion.Scope) error {
	if err := Convert_v1alpha1_KubeletX509Authentication_To_componentconfig_KubeletX509Authentication(&in.X509, &out.X509, s); err != nil {
		return err
	}
	if err := Convert_v1alpha1_KubeletWebhookAuthentication_To_componentconfig_KubeletWebhookAuthentication(&in.Webhook, &out.Webhook, s); err != nil {
		return err
	}
	if err := Convert_v1alpha1_KubeletAnonymousAuthentication_To_componentconfig_KubeletAnonymousAuthentication(&in.Anonymous, &out.Anonymous, s); err != nil {
		return err
	}
	return nil
}

func Convert_v1alpha1_KubeletAuthentication_To_componentconfig_KubeletAuthentication(in *KubeletAuthentication, out *componentconfig.KubeletAuthentication, s conversion.Scope) error {
	return autoConvert_v1alpha1_KubeletAuthentication_To_componentconfig_KubeletAuthentication(in, out, s)
}

func autoConvert_componentconfig_KubeletAuthentication_To_v1alpha1_KubeletAuthentication(in *componentconfig.KubeletAuthentication, out *KubeletAuthentication, s conversion.Scope) error {
	if err := Convert_componentconfig_KubeletX509Authentication_To_v1alpha1_KubeletX509Authentication(&in.X509, &out.X509, s); err != nil {
		return err
	}
	if err := Convert_componentconfig_KubeletWebhookAuthentication_To_v1alpha1_KubeletWebhookAuthentication(&in.Webhook, &out.Webhook, s); err != nil {
		return err
	}
	if err := Convert_componentconfig_KubeletAnonymousAuthentication_To_v1alpha1_KubeletAnonymousAuthentication(&in.Anonymous, &out.Anonymous, s); err != nil {
		return err
	}
	return nil
}

func Convert_componentconfig_KubeletAuthentication_To_v1alpha1_KubeletAuthentication(in *componentconfig.KubeletAuthentication, out *KubeletAuthentication, s conversion.Scope) error {
	return autoConvert_componentconfig_KubeletAuthentication_To_v1alpha1_KubeletAuthentication(in, out, s)
}

func autoConvert_v1alpha1_KubeletAuthorization_To_componentconfig_KubeletAuthorization(in *KubeletAuthorization, out *componentconfig.KubeletAuthorization, s conversion.Scope) error {
	out.Mode = componentconfig.KubeletAuthorizationMode(in.Mode)
	if err := Convert_v1alpha1_KubeletWebhookAuthorization_To_componentconfig_KubeletWebhookAuthorization(&in.Webhook, &out.Webhook, s); err != nil {
		return err
	}
	return nil
}

func Convert_v1alpha1_KubeletAuthorization_To_componentconfig_KubeletAuthorization(in *KubeletAuthorization, out *componentconfig.KubeletAuthorization, s conversion.Scope) error {
	return autoConvert_v1alpha1_KubeletAuthorization_To_componentconfig_KubeletAuthorization(in, out, s)
}

func autoConvert_componentconfig_KubeletAuthorization_To_v1alpha1_KubeletAuthorization(in *componentconfig.KubeletAuthorization, out *KubeletAuthorization, s conversion.Scope) error {
	out.Mode = KubeletAuthorizationMode(in.Mode)
	if err := Convert_componentconfig_KubeletWebhookAuthorization_To_v1alpha1_KubeletWebhookAuthorization(&in.Webhook, &out.Webhook, s); err != nil {
		return err
	}
	return nil
}

func Convert_componentconfig_KubeletAuthorization_To_v1alpha1_KubeletAuthorization(in *componentconfig.KubeletAuthorization, out *KubeletAuthorization, s conversion.Scope) error {
	return autoConvert_componentconfig_KubeletAuthorization_To_v1alpha1_KubeletAuthorization(in, out, s)
}

func autoConvert_v1alpha1_KubeletConfiguration_To_componentconfig_KubeletConfiguration(in *KubeletConfiguration, out *componentconfig.KubeletConfiguration, s conversion.Scope) error {
	out.PodManifestPath = in.PodManifestPath
	out.SyncFrequency = in.SyncFrequency
	out.FileCheckFrequency = in.FileCheckFrequency
	out.HTTPCheckFrequency = in.HTTPCheckFrequency
	out.ManifestURL = in.ManifestURL
	out.ManifestURLHeader = in.ManifestURLHeader
	if err := v1.Convert_Pointer_bool_To_bool(&in.EnableServer, &out.EnableServer, s); err != nil {
		return err
	}
	out.Address = in.Address
	out.Port = in.Port
	out.ReadOnlyPort = in.ReadOnlyPort
	out.TLSCertFile = in.TLSCertFile
	out.TLSPrivateKeyFile = in.TLSPrivateKeyFile
	out.CertDirectory = in.CertDirectory
	if err := Convert_v1alpha1_KubeletAuthentication_To_componentconfig_KubeletAuthentication(&in.Authentication, &out.Authentication, s); err != nil {
		return err
	}
	if err := Convert_v1alpha1_KubeletAuthorization_To_componentconfig_KubeletAuthorization(&in.Authorization, &out.Authorization, s); err != nil {
		return err
	}
	out.HostnameOverride = in.HostnameOverride
	out.PodInfraContainerImage = in.PodInfraContainerImage
	out.DockerEndpoint = in.DockerEndpoint
	out.RootDirectory = in.RootDirectory
	out.SeccompProfileRoot = in.SeccompProfileRoot
	if err := v1.Convert_Pointer_bool_To_bool(&in.AllowPrivileged, &out.AllowPrivileged, s); err != nil {
		return err
	}
	out.HostNetworkSources = *(*[]string)(unsafe.Pointer(&in.HostNetworkSources))
	out.HostPIDSources = *(*[]string)(unsafe.Pointer(&in.HostPIDSources))
	out.HostIPCSources = *(*[]string)(unsafe.Pointer(&in.HostIPCSources))
	if err := v1.Convert_Pointer_int32_To_int32(&in.RegistryPullQPS, &out.RegistryPullQPS, s); err != nil {
		return err
	}
	out.RegistryBurst = in.RegistryBurst
	if err := v1.Convert_Pointer_int32_To_int32(&in.EventRecordQPS, &out.EventRecordQPS, s); err != nil {
		return err
	}
	out.EventBurst = in.EventBurst
	if err := v1.Convert_Pointer_bool_To_bool(&in.EnableDebuggingHandlers, &out.EnableDebuggingHandlers, s); err != nil {
		return err
	}
	out.EnableContentionProfiling = in.EnableContentionProfiling
	out.MinimumGCAge = in.MinimumGCAge
	out.MaxPerPodContainerCount = in.MaxPerPodContainerCount
	if err := v1.Convert_Pointer_int32_To_int32(&in.MaxContainerCount, &out.MaxContainerCount, s); err != nil {
		return err
	}
	out.CAdvisorPort = in.CAdvisorPort
	out.HealthzPort = in.HealthzPort
	out.HealthzBindAddress = in.HealthzBindAddress
	if err := v1.Convert_Pointer_int32_To_int32(&in.OOMScoreAdj, &out.OOMScoreAdj, s); err != nil {
		return err
	}
	if err := v1.Convert_Pointer_bool_To_bool(&in.RegisterNode, &out.RegisterNode, s); err != nil {
		return err
	}
	out.ClusterDomain = in.ClusterDomain
	out.MasterServiceNamespace = in.MasterServiceNamespace
	out.ClusterDNS = *(*[]string)(unsafe.Pointer(&in.ClusterDNS))
	out.StreamingConnectionIdleTimeout = in.StreamingConnectionIdleTimeout
	out.NodeStatusUpdateFrequency = in.NodeStatusUpdateFrequency
	out.ImageMinimumGCAge = in.ImageMinimumGCAge
	if err := v1.Convert_Pointer_int32_To_int32(&in.ImageGCHighThresholdPercent, &out.ImageGCHighThresholdPercent, s); err != nil {
		return err
	}
	if err := v1.Convert_Pointer_int32_To_int32(&in.ImageGCLowThresholdPercent, &out.ImageGCLowThresholdPercent, s); err != nil {
		return err
	}
	out.LowDiskSpaceThresholdMB = in.LowDiskSpaceThresholdMB
	out.VolumeStatsAggPeriod = in.VolumeStatsAggPeriod
	out.NetworkPluginName = in.NetworkPluginName
	out.NetworkPluginDir = in.NetworkPluginDir
	out.CNIConfDir = in.CNIConfDir
	out.CNIBinDir = in.CNIBinDir
	out.NetworkPluginMTU = in.NetworkPluginMTU
	out.VolumePluginDir = in.VolumePluginDir
	out.CloudProvider = in.CloudProvider
	out.CloudConfigFile = in.CloudConfigFile
	out.KubeletCgroups = in.KubeletCgroups
	out.RuntimeCgroups = in.RuntimeCgroups
	out.SystemCgroups = in.SystemCgroups
	out.CgroupRoot = in.CgroupRoot
	if err := v1.Convert_Pointer_bool_To_bool(&in.CgroupsPerQOS, &out.CgroupsPerQOS, s); err != nil {
		return err
	}
	out.CgroupDriver = in.CgroupDriver
	out.ContainerRuntime = in.ContainerRuntime
	out.RemoteRuntimeEndpoint = in.RemoteRuntimeEndpoint
	out.RemoteImageEndpoint = in.RemoteImageEndpoint
	out.RuntimeRequestTimeout = in.RuntimeRequestTimeout
	out.ImagePullProgressDeadline = in.ImagePullProgressDeadline
	out.RktPath = in.RktPath
	out.ExperimentalMounterPath = in.ExperimentalMounterPath
	out.RktAPIEndpoint = in.RktAPIEndpoint
	out.RktStage1Image = in.RktStage1Image
	if err := v1.Convert_Pointer_string_To_string(&in.LockFilePath, &out.LockFilePath, s); err != nil {
		return err
	}
	out.ExitOnLockContention = in.ExitOnLockContention
	out.HairpinMode = in.HairpinMode
	out.BabysitDaemons = in.BabysitDaemons
	out.MaxPods = in.MaxPods
	out.DockerExecHandlerName = in.DockerExecHandlerName
	out.PodCIDR = in.PodCIDR
	out.ResolverConfig = in.ResolverConfig
	if err := v1.Convert_Pointer_bool_To_bool(&in.CPUCFSQuota, &out.CPUCFSQuota, s); err != nil {
		return err
	}
	if err := v1.Convert_Pointer_bool_To_bool(&in.Containerized, &out.Containerized, s); err != nil {
		return err
	}
	out.MaxOpenFiles = in.MaxOpenFiles
	if err := v1.Convert_Pointer_bool_To_bool(&in.RegisterSchedulable, &out.RegisterSchedulable, s); err != nil {
		return err
	}
	out.RegisterWithTaints = *(*[]api.Taint)(unsafe.Pointer(&in.RegisterWithTaints))
	out.ContentType = in.ContentType
	if err := v1.Convert_Pointer_int32_To_int32(&in.KubeAPIQPS, &out.KubeAPIQPS, s); err != nil {
		return err
	}
	out.KubeAPIBurst = in.KubeAPIBurst
	if err := v1.Convert_Pointer_bool_To_bool(&in.SerializeImagePulls, &out.SerializeImagePulls, s); err != nil {
		return err
	}
	out.OutOfDiskTransitionFrequency = in.OutOfDiskTransitionFrequency
	out.NodeIP = in.NodeIP
	out.NodeLabels = *(*map[string]string)(unsafe.Pointer(&in.NodeLabels))
	out.NonMasqueradeCIDR = in.NonMasqueradeCIDR
	out.EnableCustomMetrics = in.EnableCustomMetrics
	if err := v1.Convert_Pointer_string_To_string(&in.EvictionHard, &out.EvictionHard, s); err != nil {
		return err
	}
	out.EvictionSoft = in.EvictionSoft
	out.EvictionSoftGracePeriod = in.EvictionSoftGracePeriod
	out.EvictionPressureTransitionPeriod = in.EvictionPressureTransitionPeriod
	out.EvictionMaxPodGracePeriod = in.EvictionMaxPodGracePeriod
	out.EvictionMinimumReclaim = in.EvictionMinimumReclaim
	if err := v1.Convert_Pointer_bool_To_bool(&in.ExperimentalKernelMemcgNotification, &out.ExperimentalKernelMemcgNotification, s); err != nil {
		return err
	}
	out.PodsPerCore = in.PodsPerCore
	if err := v1.Convert_Pointer_bool_To_bool(&in.EnableControllerAttachDetach, &out.EnableControllerAttachDetach, s); err != nil {
		return err
	}
	out.ExperimentalQOSReserved = *(*componentconfig.ConfigurationMap)(unsafe.Pointer(&in.ExperimentalQOSReserved))
	out.ProtectKernelDefaults = in.ProtectKernelDefaults
	if err := v1.Convert_Pointer_bool_To_bool(&in.MakeIPTablesUtilChains, &out.MakeIPTablesUtilChains, s); err != nil {
		return err
	}
	if err := v1.Convert_Pointer_int32_To_int32(&in.IPTablesMasqueradeBit, &out.IPTablesMasqueradeBit, s); err != nil {
		return err
	}
	if err := v1.Convert_Pointer_int32_To_int32(&in.IPTablesDropBit, &out.IPTablesDropBit, s); err != nil {
		return err
	}
	out.AllowedUnsafeSysctls = *(*[]string)(unsafe.Pointer(&in.AllowedUnsafeSysctls))
	out.FeatureGates = in.FeatureGates
	if err := v1.Convert_Pointer_bool_To_bool(&in.EnableCRI, &out.EnableCRI, s); err != nil {
		return err
	}
	out.ExperimentalFailSwapOn = in.ExperimentalFailSwapOn
	out.ExperimentalCheckNodeCapabilitiesBeforeMount = in.ExperimentalCheckNodeCapabilitiesBeforeMount
	out.KeepTerminatedPodVolumes = in.KeepTerminatedPodVolumes
	out.SystemReserved = *(*componentconfig.ConfigurationMap)(unsafe.Pointer(&in.SystemReserved))
	out.KubeReserved = *(*componentconfig.ConfigurationMap)(unsafe.Pointer(&in.KubeReserved))
	out.SystemReservedCgroup = in.SystemReservedCgroup
	out.KubeReservedCgroup = in.KubeReservedCgroup
	out.EnforceNodeAllocatable = *(*[]string)(unsafe.Pointer(&in.EnforceNodeAllocatable))
	out.ExperimentalNodeAllocatableIgnoreEvictionThreshold = in.ExperimentalNodeAllocatableIgnoreEvictionThreshold
	return nil
}

func Convert_v1alpha1_KubeletConfiguration_To_componentconfig_KubeletConfiguration(in *KubeletConfiguration, out *componentconfig.KubeletConfiguration, s conversion.Scope) error {
	return autoConvert_v1alpha1_KubeletConfiguration_To_componentconfig_KubeletConfiguration(in, out, s)
}

func autoConvert_componentconfig_KubeletConfiguration_To_v1alpha1_KubeletConfiguration(in *componentconfig.KubeletConfiguration, out *KubeletConfiguration, s conversion.Scope) error {
	out.PodManifestPath = in.PodManifestPath
	out.SyncFrequency = in.SyncFrequency
	out.FileCheckFrequency = in.FileCheckFrequency
	out.HTTPCheckFrequency = in.HTTPCheckFrequency
	out.ManifestURL = in.ManifestURL
	out.ManifestURLHeader = in.ManifestURLHeader
	if err := v1.Convert_bool_To_Pointer_bool(&in.EnableServer, &out.EnableServer, s); err != nil {
		return err
	}
	out.Address = in.Address
	out.Port = in.Port
	out.ReadOnlyPort = in.ReadOnlyPort
	out.TLSCertFile = in.TLSCertFile
	out.TLSPrivateKeyFile = in.TLSPrivateKeyFile
	out.CertDirectory = in.CertDirectory
	if err := Convert_componentconfig_KubeletAuthentication_To_v1alpha1_KubeletAuthentication(&in.Authentication, &out.Authentication, s); err != nil {
		return err
	}
	if err := Convert_componentconfig_KubeletAuthorization_To_v1alpha1_KubeletAuthorization(&in.Authorization, &out.Authorization, s); err != nil {
		return err
	}
	out.HostnameOverride = in.HostnameOverride
	out.PodInfraContainerImage = in.PodInfraContainerImage
	out.DockerEndpoint = in.DockerEndpoint
	out.RootDirectory = in.RootDirectory
	out.SeccompProfileRoot = in.SeccompProfileRoot
	if err := v1.Convert_bool_To_Pointer_bool(&in.AllowPrivileged, &out.AllowPrivileged, s); err != nil {
		return err
	}
	out.HostNetworkSources = *(*[]string)(unsafe.Pointer(&in.HostNetworkSources))
	out.HostPIDSources = *(*[]string)(unsafe.Pointer(&in.HostPIDSources))
	out.HostIPCSources = *(*[]string)(unsafe.Pointer(&in.HostIPCSources))
	if err := v1.Convert_int32_To_Pointer_int32(&in.RegistryPullQPS, &out.RegistryPullQPS, s); err != nil {
		return err
	}
	out.RegistryBurst = in.RegistryBurst
	if err := v1.Convert_int32_To_Pointer_int32(&in.EventRecordQPS, &out.EventRecordQPS, s); err != nil {
		return err
	}
	out.EventBurst = in.EventBurst
	if err := v1.Convert_bool_To_Pointer_bool(&in.EnableDebuggingHandlers, &out.EnableDebuggingHandlers, s); err != nil {
		return err
	}
	out.EnableContentionProfiling = in.EnableContentionProfiling
	out.MinimumGCAge = in.MinimumGCAge
	out.MaxPerPodContainerCount = in.MaxPerPodContainerCount
	if err := v1.Convert_int32_To_Pointer_int32(&in.MaxContainerCount, &out.MaxContainerCount, s); err != nil {
		return err
	}
	out.CAdvisorPort = in.CAdvisorPort
	out.HealthzPort = in.HealthzPort
	out.HealthzBindAddress = in.HealthzBindAddress
	if err := v1.Convert_int32_To_Pointer_int32(&in.OOMScoreAdj, &out.OOMScoreAdj, s); err != nil {
		return err
	}
	if err := v1.Convert_bool_To_Pointer_bool(&in.RegisterNode, &out.RegisterNode, s); err != nil {
		return err
	}
	out.ClusterDomain = in.ClusterDomain
	out.MasterServiceNamespace = in.MasterServiceNamespace
	out.ClusterDNS = *(*[]string)(unsafe.Pointer(&in.ClusterDNS))
	out.StreamingConnectionIdleTimeout = in.StreamingConnectionIdleTimeout
	out.NodeStatusUpdateFrequency = in.NodeStatusUpdateFrequency
	out.ImageMinimumGCAge = in.ImageMinimumGCAge
	if err := v1.Convert_int32_To_Pointer_int32(&in.ImageGCHighThresholdPercent, &out.ImageGCHighThresholdPercent, s); err != nil {
		return err
	}
	if err := v1.Convert_int32_To_Pointer_int32(&in.ImageGCLowThresholdPercent, &out.ImageGCLowThresholdPercent, s); err != nil {
		return err
	}
	out.LowDiskSpaceThresholdMB = in.LowDiskSpaceThresholdMB
	out.VolumeStatsAggPeriod = in.VolumeStatsAggPeriod
	out.NetworkPluginName = in.NetworkPluginName
	out.NetworkPluginMTU = in.NetworkPluginMTU
	out.NetworkPluginDir = in.NetworkPluginDir
	out.CNIConfDir = in.CNIConfDir
	out.CNIBinDir = in.CNIBinDir
	out.VolumePluginDir = in.VolumePluginDir
	out.CloudProvider = in.CloudProvider
	out.CloudConfigFile = in.CloudConfigFile
	out.KubeletCgroups = in.KubeletCgroups
	if err := v1.Convert_bool_To_Pointer_bool(&in.CgroupsPerQOS, &out.CgroupsPerQOS, s); err != nil {
		return err
	}
	out.CgroupDriver = in.CgroupDriver
	out.RuntimeCgroups = in.RuntimeCgroups
	out.SystemCgroups = in.SystemCgroups
	out.CgroupRoot = in.CgroupRoot
	out.ContainerRuntime = in.ContainerRuntime
	out.RemoteRuntimeEndpoint = in.RemoteRuntimeEndpoint
	out.RemoteImageEndpoint = in.RemoteImageEndpoint
	out.RuntimeRequestTimeout = in.RuntimeRequestTimeout
	out.ImagePullProgressDeadline = in.ImagePullProgressDeadline
	out.RktPath = in.RktPath
	out.ExperimentalMounterPath = in.ExperimentalMounterPath
	out.RktAPIEndpoint = in.RktAPIEndpoint
	out.RktStage1Image = in.RktStage1Image
	if err := v1.Convert_string_To_Pointer_string(&in.LockFilePath, &out.LockFilePath, s); err != nil {
		return err
	}
	out.ExitOnLockContention = in.ExitOnLockContention
	out.HairpinMode = in.HairpinMode
	out.BabysitDaemons = in.BabysitDaemons
	out.MaxPods = in.MaxPods
	out.DockerExecHandlerName = in.DockerExecHandlerName
	out.PodCIDR = in.PodCIDR
	out.ResolverConfig = in.ResolverConfig
	if err := v1.Convert_bool_To_Pointer_bool(&in.CPUCFSQuota, &out.CPUCFSQuota, s); err != nil {
		return err
	}
	if err := v1.Convert_bool_To_Pointer_bool(&in.Containerized, &out.Containerized, s); err != nil {
		return err
	}
	out.MaxOpenFiles = in.MaxOpenFiles
	if err := v1.Convert_bool_To_Pointer_bool(&in.RegisterSchedulable, &out.RegisterSchedulable, s); err != nil {
		return err
	}
	out.RegisterWithTaints = *(*[]api_v1.Taint)(unsafe.Pointer(&in.RegisterWithTaints))
	out.ContentType = in.ContentType
	if err := v1.Convert_int32_To_Pointer_int32(&in.KubeAPIQPS, &out.KubeAPIQPS, s); err != nil {
		return err
	}
	out.KubeAPIBurst = in.KubeAPIBurst
	if err := v1.Convert_bool_To_Pointer_bool(&in.SerializeImagePulls, &out.SerializeImagePulls, s); err != nil {
		return err
	}
	out.OutOfDiskTransitionFrequency = in.OutOfDiskTransitionFrequency
	out.NodeIP = in.NodeIP
	out.NodeLabels = *(*map[string]string)(unsafe.Pointer(&in.NodeLabels))
	out.NonMasqueradeCIDR = in.NonMasqueradeCIDR
	out.EnableCustomMetrics = in.EnableCustomMetrics
	if err := v1.Convert_string_To_Pointer_string(&in.EvictionHard, &out.EvictionHard, s); err != nil {
		return err
	}
	out.EvictionSoft = in.EvictionSoft
	out.EvictionSoftGracePeriod = in.EvictionSoftGracePeriod
	out.EvictionPressureTransitionPeriod = in.EvictionPressureTransitionPeriod
	out.EvictionMaxPodGracePeriod = in.EvictionMaxPodGracePeriod
	out.EvictionMinimumReclaim = in.EvictionMinimumReclaim
	if err := v1.Convert_bool_To_Pointer_bool(&in.ExperimentalKernelMemcgNotification, &out.ExperimentalKernelMemcgNotification, s); err != nil {
		return err
	}
	out.PodsPerCore = in.PodsPerCore
	if err := v1.Convert_bool_To_Pointer_bool(&in.EnableControllerAttachDetach, &out.EnableControllerAttachDetach, s); err != nil {
		return err
	}
	out.ExperimentalQOSReserved = *(*map[string]string)(unsafe.Pointer(&in.ExperimentalQOSReserved))
	out.ProtectKernelDefaults = in.ProtectKernelDefaults
	if err := v1.Convert_bool_To_Pointer_bool(&in.MakeIPTablesUtilChains, &out.MakeIPTablesUtilChains, s); err != nil {
		return err
	}
	if err := v1.Convert_int32_To_Pointer_int32(&in.IPTablesMasqueradeBit, &out.IPTablesMasqueradeBit, s); err != nil {
		return err
	}
	if err := v1.Convert_int32_To_Pointer_int32(&in.IPTablesDropBit, &out.IPTablesDropBit, s); err != nil {
		return err
	}
	out.AllowedUnsafeSysctls = *(*[]string)(unsafe.Pointer(&in.AllowedUnsafeSysctls))
	out.FeatureGates = in.FeatureGates
	if err := v1.Convert_bool_To_Pointer_bool(&in.EnableCRI, &out.EnableCRI, s); err != nil {
		return err
	}
	out.ExperimentalFailSwapOn = in.ExperimentalFailSwapOn
	out.ExperimentalCheckNodeCapabilitiesBeforeMount = in.ExperimentalCheckNodeCapabilitiesBeforeMount
	out.KeepTerminatedPodVolumes = in.KeepTerminatedPodVolumes
	out.SystemReserved = *(*map[string]string)(unsafe.Pointer(&in.SystemReserved))
	out.KubeReserved = *(*map[string]string)(unsafe.Pointer(&in.KubeReserved))
	out.SystemReservedCgroup = in.SystemReservedCgroup
	out.KubeReservedCgroup = in.KubeReservedCgroup
	out.EnforceNodeAllocatable = *(*[]string)(unsafe.Pointer(&in.EnforceNodeAllocatable))
	out.ExperimentalNodeAllocatableIgnoreEvictionThreshold = in.ExperimentalNodeAllocatableIgnoreEvictionThreshold
	return nil
}

func Convert_componentconfig_KubeletConfiguration_To_v1alpha1_KubeletConfiguration(in *componentconfig.KubeletConfiguration, out *KubeletConfiguration, s conversion.Scope) error {
	return autoConvert_componentconfig_KubeletConfiguration_To_v1alpha1_KubeletConfiguration(in, out, s)
}

func autoConvert_v1alpha1_KubeletWebhookAuthentication_To_componentconfig_KubeletWebhookAuthentication(in *KubeletWebhookAuthentication, out *componentconfig.KubeletWebhookAuthentication, s conversion.Scope) error {
	if err := v1.Convert_Pointer_bool_To_bool(&in.Enabled, &out.Enabled, s); err != nil {
		return err
	}
	out.CacheTTL = in.CacheTTL
	return nil
}

func Convert_v1alpha1_KubeletWebhookAuthentication_To_componentconfig_KubeletWebhookAuthentication(in *KubeletWebhookAuthentication, out *componentconfig.KubeletWebhookAuthentication, s conversion.Scope) error {
	return autoConvert_v1alpha1_KubeletWebhookAuthentication_To_componentconfig_KubeletWebhookAuthentication(in, out, s)
}

func autoConvert_componentconfig_KubeletWebhookAuthentication_To_v1alpha1_KubeletWebhookAuthentication(in *componentconfig.KubeletWebhookAuthentication, out *KubeletWebhookAuthentication, s conversion.Scope) error {
	if err := v1.Convert_bool_To_Pointer_bool(&in.Enabled, &out.Enabled, s); err != nil {
		return err
	}
	out.CacheTTL = in.CacheTTL
	return nil
}

func Convert_componentconfig_KubeletWebhookAuthentication_To_v1alpha1_KubeletWebhookAuthentication(in *componentconfig.KubeletWebhookAuthentication, out *KubeletWebhookAuthentication, s conversion.Scope) error {
	return autoConvert_componentconfig_KubeletWebhookAuthentication_To_v1alpha1_KubeletWebhookAuthentication(in, out, s)
}

func autoConvert_v1alpha1_KubeletWebhookAuthorization_To_componentconfig_KubeletWebhookAuthorization(in *KubeletWebhookAuthorization, out *componentconfig.KubeletWebhookAuthorization, s conversion.Scope) error {
	out.CacheAuthorizedTTL = in.CacheAuthorizedTTL
	out.CacheUnauthorizedTTL = in.CacheUnauthorizedTTL
	return nil
}

func Convert_v1alpha1_KubeletWebhookAuthorization_To_componentconfig_KubeletWebhookAuthorization(in *KubeletWebhookAuthorization, out *componentconfig.KubeletWebhookAuthorization, s conversion.Scope) error {
	return autoConvert_v1alpha1_KubeletWebhookAuthorization_To_componentconfig_KubeletWebhookAuthorization(in, out, s)
}

func autoConvert_componentconfig_KubeletWebhookAuthorization_To_v1alpha1_KubeletWebhookAuthorization(in *componentconfig.KubeletWebhookAuthorization, out *KubeletWebhookAuthorization, s conversion.Scope) error {
	out.CacheAuthorizedTTL = in.CacheAuthorizedTTL
	out.CacheUnauthorizedTTL = in.CacheUnauthorizedTTL
	return nil
}

func Convert_componentconfig_KubeletWebhookAuthorization_To_v1alpha1_KubeletWebhookAuthorization(in *componentconfig.KubeletWebhookAuthorization, out *KubeletWebhookAuthorization, s conversion.Scope) error {
	return autoConvert_componentconfig_KubeletWebhookAuthorization_To_v1alpha1_KubeletWebhookAuthorization(in, out, s)
}

func autoConvert_v1alpha1_KubeletX509Authentication_To_componentconfig_KubeletX509Authentication(in *KubeletX509Authentication, out *componentconfig.KubeletX509Authentication, s conversion.Scope) error {
	out.ClientCAFile = in.ClientCAFile
	return nil
}

func Convert_v1alpha1_KubeletX509Authentication_To_componentconfig_KubeletX509Authentication(in *KubeletX509Authentication, out *componentconfig.KubeletX509Authentication, s conversion.Scope) error {
	return autoConvert_v1alpha1_KubeletX509Authentication_To_componentconfig_KubeletX509Authentication(in, out, s)
}

func autoConvert_componentconfig_KubeletX509Authentication_To_v1alpha1_KubeletX509Authentication(in *componentconfig.KubeletX509Authentication, out *KubeletX509Authentication, s conversion.Scope) error {
	out.ClientCAFile = in.ClientCAFile
	return nil
}

func Convert_componentconfig_KubeletX509Authentication_To_v1alpha1_KubeletX509Authentication(in *componentconfig.KubeletX509Authentication, out *KubeletX509Authentication, s conversion.Scope) error {
	return autoConvert_componentconfig_KubeletX509Authentication_To_v1alpha1_KubeletX509Authentication(in, out, s)
}

func autoConvert_v1alpha1_LeaderElectionConfiguration_To_componentconfig_LeaderElectionConfiguration(in *LeaderElectionConfiguration, out *componentconfig.LeaderElectionConfiguration, s conversion.Scope) error {
	if err := v1.Convert_Pointer_bool_To_bool(&in.LeaderElect, &out.LeaderElect, s); err != nil {
		return err
	}
	out.LeaseDuration = in.LeaseDuration
	out.RenewDeadline = in.RenewDeadline
	out.RetryPeriod = in.RetryPeriod
	return nil
}

func Convert_v1alpha1_LeaderElectionConfiguration_To_componentconfig_LeaderElectionConfiguration(in *LeaderElectionConfiguration, out *componentconfig.LeaderElectionConfiguration, s conversion.Scope) error {
	return autoConvert_v1alpha1_LeaderElectionConfiguration_To_componentconfig_LeaderElectionConfiguration(in, out, s)
}

func autoConvert_componentconfig_LeaderElectionConfiguration_To_v1alpha1_LeaderElectionConfiguration(in *componentconfig.LeaderElectionConfiguration, out *LeaderElectionConfiguration, s conversion.Scope) error {
	if err := v1.Convert_bool_To_Pointer_bool(&in.LeaderElect, &out.LeaderElect, s); err != nil {
		return err
	}
	out.LeaseDuration = in.LeaseDuration
	out.RenewDeadline = in.RenewDeadline
	out.RetryPeriod = in.RetryPeriod
	return nil
}

func Convert_componentconfig_LeaderElectionConfiguration_To_v1alpha1_LeaderElectionConfiguration(in *componentconfig.LeaderElectionConfiguration, out *LeaderElectionConfiguration, s conversion.Scope) error {
	return autoConvert_componentconfig_LeaderElectionConfiguration_To_v1alpha1_LeaderElectionConfiguration(in, out, s)
}
