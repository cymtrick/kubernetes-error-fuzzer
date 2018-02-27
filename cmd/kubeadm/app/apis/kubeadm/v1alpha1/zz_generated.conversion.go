// +build !ignore_autogenerated

/*
Copyright 2018 The Kubernetes Authors.

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

	core_v1 "k8s.io/api/core/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	conversion "k8s.io/apimachinery/pkg/conversion"
	runtime "k8s.io/apimachinery/pkg/runtime"
	kubeadm "k8s.io/kubernetes/cmd/kubeadm/app/apis/kubeadm"
	v1beta1 "k8s.io/kubernetes/pkg/kubelet/apis/kubeletconfig/v1beta1"
	kubeproxyconfig_v1alpha1 "k8s.io/kubernetes/pkg/proxy/apis/kubeproxyconfig/v1alpha1"
)

func init() {
	localSchemeBuilder.Register(RegisterConversions)
}

// RegisterConversions adds conversion functions to the given scheme.
// Public to allow building arbitrary schemes.
func RegisterConversions(scheme *runtime.Scheme) error {
	return scheme.AddGeneratedConversionFuncs(
		Convert_v1alpha1_API_To_kubeadm_API,
		Convert_kubeadm_API_To_v1alpha1_API,
		Convert_v1alpha1_AuditPolicyConfiguration_To_kubeadm_AuditPolicyConfiguration,
		Convert_kubeadm_AuditPolicyConfiguration_To_v1alpha1_AuditPolicyConfiguration,
		Convert_v1alpha1_Etcd_To_kubeadm_Etcd,
		Convert_kubeadm_Etcd_To_v1alpha1_Etcd,
		Convert_v1alpha1_HostPathMount_To_kubeadm_HostPathMount,
		Convert_kubeadm_HostPathMount_To_v1alpha1_HostPathMount,
		Convert_v1alpha1_KubeProxy_To_kubeadm_KubeProxy,
		Convert_kubeadm_KubeProxy_To_v1alpha1_KubeProxy,
		Convert_v1alpha1_KubeletConfiguration_To_kubeadm_KubeletConfiguration,
		Convert_kubeadm_KubeletConfiguration_To_v1alpha1_KubeletConfiguration,
		Convert_v1alpha1_MasterConfiguration_To_kubeadm_MasterConfiguration,
		Convert_kubeadm_MasterConfiguration_To_v1alpha1_MasterConfiguration,
		Convert_v1alpha1_Networking_To_kubeadm_Networking,
		Convert_kubeadm_Networking_To_v1alpha1_Networking,
		Convert_v1alpha1_NodeConfiguration_To_kubeadm_NodeConfiguration,
		Convert_kubeadm_NodeConfiguration_To_v1alpha1_NodeConfiguration,
		Convert_v1alpha1_SelfHostedEtcd_To_kubeadm_SelfHostedEtcd,
		Convert_kubeadm_SelfHostedEtcd_To_v1alpha1_SelfHostedEtcd,
		Convert_v1alpha1_TokenDiscovery_To_kubeadm_TokenDiscovery,
		Convert_kubeadm_TokenDiscovery_To_v1alpha1_TokenDiscovery,
	)
}

func autoConvert_v1alpha1_API_To_kubeadm_API(in *API, out *kubeadm.API, s conversion.Scope) error {
	out.AdvertiseAddress = in.AdvertiseAddress
	out.ControlPlaneEndpoint = in.ControlPlaneEndpoint
	out.BindPort = in.BindPort
	return nil
}

// Convert_v1alpha1_API_To_kubeadm_API is an autogenerated conversion function.
func Convert_v1alpha1_API_To_kubeadm_API(in *API, out *kubeadm.API, s conversion.Scope) error {
	return autoConvert_v1alpha1_API_To_kubeadm_API(in, out, s)
}

func autoConvert_kubeadm_API_To_v1alpha1_API(in *kubeadm.API, out *API, s conversion.Scope) error {
	out.AdvertiseAddress = in.AdvertiseAddress
	out.ControlPlaneEndpoint = in.ControlPlaneEndpoint
	out.BindPort = in.BindPort
	return nil
}

// Convert_kubeadm_API_To_v1alpha1_API is an autogenerated conversion function.
func Convert_kubeadm_API_To_v1alpha1_API(in *kubeadm.API, out *API, s conversion.Scope) error {
	return autoConvert_kubeadm_API_To_v1alpha1_API(in, out, s)
}

func autoConvert_v1alpha1_AuditPolicyConfiguration_To_kubeadm_AuditPolicyConfiguration(in *AuditPolicyConfiguration, out *kubeadm.AuditPolicyConfiguration, s conversion.Scope) error {
	out.Path = in.Path
	out.LogDir = in.LogDir
	out.LogMaxAge = (*int32)(unsafe.Pointer(in.LogMaxAge))
	return nil
}

// Convert_v1alpha1_AuditPolicyConfiguration_To_kubeadm_AuditPolicyConfiguration is an autogenerated conversion function.
func Convert_v1alpha1_AuditPolicyConfiguration_To_kubeadm_AuditPolicyConfiguration(in *AuditPolicyConfiguration, out *kubeadm.AuditPolicyConfiguration, s conversion.Scope) error {
	return autoConvert_v1alpha1_AuditPolicyConfiguration_To_kubeadm_AuditPolicyConfiguration(in, out, s)
}

func autoConvert_kubeadm_AuditPolicyConfiguration_To_v1alpha1_AuditPolicyConfiguration(in *kubeadm.AuditPolicyConfiguration, out *AuditPolicyConfiguration, s conversion.Scope) error {
	out.Path = in.Path
	out.LogDir = in.LogDir
	out.LogMaxAge = (*int32)(unsafe.Pointer(in.LogMaxAge))
	return nil
}

// Convert_kubeadm_AuditPolicyConfiguration_To_v1alpha1_AuditPolicyConfiguration is an autogenerated conversion function.
func Convert_kubeadm_AuditPolicyConfiguration_To_v1alpha1_AuditPolicyConfiguration(in *kubeadm.AuditPolicyConfiguration, out *AuditPolicyConfiguration, s conversion.Scope) error {
	return autoConvert_kubeadm_AuditPolicyConfiguration_To_v1alpha1_AuditPolicyConfiguration(in, out, s)
}

func autoConvert_v1alpha1_Etcd_To_kubeadm_Etcd(in *Etcd, out *kubeadm.Etcd, s conversion.Scope) error {
	out.Endpoints = *(*[]string)(unsafe.Pointer(&in.Endpoints))
	out.CAFile = in.CAFile
	out.CertFile = in.CertFile
	out.KeyFile = in.KeyFile
	out.DataDir = in.DataDir
	out.ExtraArgs = *(*map[string]string)(unsafe.Pointer(&in.ExtraArgs))
	out.Image = in.Image
	out.SelfHosted = (*kubeadm.SelfHostedEtcd)(unsafe.Pointer(in.SelfHosted))
	out.ServerCertSANs = *(*[]string)(unsafe.Pointer(&in.ServerCertSANs))
	out.PeerCertSANs = *(*[]string)(unsafe.Pointer(&in.PeerCertSANs))
	return nil
}

// Convert_v1alpha1_Etcd_To_kubeadm_Etcd is an autogenerated conversion function.
func Convert_v1alpha1_Etcd_To_kubeadm_Etcd(in *Etcd, out *kubeadm.Etcd, s conversion.Scope) error {
	return autoConvert_v1alpha1_Etcd_To_kubeadm_Etcd(in, out, s)
}

func autoConvert_kubeadm_Etcd_To_v1alpha1_Etcd(in *kubeadm.Etcd, out *Etcd, s conversion.Scope) error {
	out.Endpoints = *(*[]string)(unsafe.Pointer(&in.Endpoints))
	out.CAFile = in.CAFile
	out.CertFile = in.CertFile
	out.KeyFile = in.KeyFile
	out.DataDir = in.DataDir
	out.ExtraArgs = *(*map[string]string)(unsafe.Pointer(&in.ExtraArgs))
	out.Image = in.Image
	out.SelfHosted = (*SelfHostedEtcd)(unsafe.Pointer(in.SelfHosted))
	out.ServerCertSANs = *(*[]string)(unsafe.Pointer(&in.ServerCertSANs))
	out.PeerCertSANs = *(*[]string)(unsafe.Pointer(&in.PeerCertSANs))
	return nil
}

// Convert_kubeadm_Etcd_To_v1alpha1_Etcd is an autogenerated conversion function.
func Convert_kubeadm_Etcd_To_v1alpha1_Etcd(in *kubeadm.Etcd, out *Etcd, s conversion.Scope) error {
	return autoConvert_kubeadm_Etcd_To_v1alpha1_Etcd(in, out, s)
}

func autoConvert_v1alpha1_HostPathMount_To_kubeadm_HostPathMount(in *HostPathMount, out *kubeadm.HostPathMount, s conversion.Scope) error {
	out.Name = in.Name
	out.HostPath = in.HostPath
	out.MountPath = in.MountPath
	return nil
}

// Convert_v1alpha1_HostPathMount_To_kubeadm_HostPathMount is an autogenerated conversion function.
func Convert_v1alpha1_HostPathMount_To_kubeadm_HostPathMount(in *HostPathMount, out *kubeadm.HostPathMount, s conversion.Scope) error {
	return autoConvert_v1alpha1_HostPathMount_To_kubeadm_HostPathMount(in, out, s)
}

func autoConvert_kubeadm_HostPathMount_To_v1alpha1_HostPathMount(in *kubeadm.HostPathMount, out *HostPathMount, s conversion.Scope) error {
	out.Name = in.Name
	out.HostPath = in.HostPath
	out.MountPath = in.MountPath
	return nil
}

// Convert_kubeadm_HostPathMount_To_v1alpha1_HostPathMount is an autogenerated conversion function.
func Convert_kubeadm_HostPathMount_To_v1alpha1_HostPathMount(in *kubeadm.HostPathMount, out *HostPathMount, s conversion.Scope) error {
	return autoConvert_kubeadm_HostPathMount_To_v1alpha1_HostPathMount(in, out, s)
}

func autoConvert_v1alpha1_KubeProxy_To_kubeadm_KubeProxy(in *KubeProxy, out *kubeadm.KubeProxy, s conversion.Scope) error {
	out.Config = (*kubeproxyconfig_v1alpha1.KubeProxyConfiguration)(unsafe.Pointer(in.Config))
	return nil
}

// Convert_v1alpha1_KubeProxy_To_kubeadm_KubeProxy is an autogenerated conversion function.
func Convert_v1alpha1_KubeProxy_To_kubeadm_KubeProxy(in *KubeProxy, out *kubeadm.KubeProxy, s conversion.Scope) error {
	return autoConvert_v1alpha1_KubeProxy_To_kubeadm_KubeProxy(in, out, s)
}

func autoConvert_kubeadm_KubeProxy_To_v1alpha1_KubeProxy(in *kubeadm.KubeProxy, out *KubeProxy, s conversion.Scope) error {
	out.Config = (*kubeproxyconfig_v1alpha1.KubeProxyConfiguration)(unsafe.Pointer(in.Config))
	return nil
}

// Convert_kubeadm_KubeProxy_To_v1alpha1_KubeProxy is an autogenerated conversion function.
func Convert_kubeadm_KubeProxy_To_v1alpha1_KubeProxy(in *kubeadm.KubeProxy, out *KubeProxy, s conversion.Scope) error {
	return autoConvert_kubeadm_KubeProxy_To_v1alpha1_KubeProxy(in, out, s)
}

func autoConvert_v1alpha1_KubeletConfiguration_To_kubeadm_KubeletConfiguration(in *KubeletConfiguration, out *kubeadm.KubeletConfiguration, s conversion.Scope) error {
	out.BaseConfig = (*v1beta1.KubeletConfiguration)(unsafe.Pointer(in.BaseConfig))
	return nil
}

// Convert_v1alpha1_KubeletConfiguration_To_kubeadm_KubeletConfiguration is an autogenerated conversion function.
func Convert_v1alpha1_KubeletConfiguration_To_kubeadm_KubeletConfiguration(in *KubeletConfiguration, out *kubeadm.KubeletConfiguration, s conversion.Scope) error {
	return autoConvert_v1alpha1_KubeletConfiguration_To_kubeadm_KubeletConfiguration(in, out, s)
}

func autoConvert_kubeadm_KubeletConfiguration_To_v1alpha1_KubeletConfiguration(in *kubeadm.KubeletConfiguration, out *KubeletConfiguration, s conversion.Scope) error {
	out.BaseConfig = (*v1beta1.KubeletConfiguration)(unsafe.Pointer(in.BaseConfig))
	return nil
}

// Convert_kubeadm_KubeletConfiguration_To_v1alpha1_KubeletConfiguration is an autogenerated conversion function.
func Convert_kubeadm_KubeletConfiguration_To_v1alpha1_KubeletConfiguration(in *kubeadm.KubeletConfiguration, out *KubeletConfiguration, s conversion.Scope) error {
	return autoConvert_kubeadm_KubeletConfiguration_To_v1alpha1_KubeletConfiguration(in, out, s)
}

func autoConvert_v1alpha1_MasterConfiguration_To_kubeadm_MasterConfiguration(in *MasterConfiguration, out *kubeadm.MasterConfiguration, s conversion.Scope) error {
	if err := Convert_v1alpha1_API_To_kubeadm_API(&in.API, &out.API, s); err != nil {
		return err
	}
	if err := Convert_v1alpha1_KubeProxy_To_kubeadm_KubeProxy(&in.KubeProxy, &out.KubeProxy, s); err != nil {
		return err
	}
	if err := Convert_v1alpha1_Etcd_To_kubeadm_Etcd(&in.Etcd, &out.Etcd, s); err != nil {
		return err
	}
	if err := Convert_v1alpha1_KubeletConfiguration_To_kubeadm_KubeletConfiguration(&in.KubeletConfiguration, &out.KubeletConfiguration, s); err != nil {
		return err
	}
	if err := Convert_v1alpha1_Networking_To_kubeadm_Networking(&in.Networking, &out.Networking, s); err != nil {
		return err
	}
	out.KubernetesVersion = in.KubernetesVersion
	out.CloudProvider = in.CloudProvider
	out.NodeName = in.NodeName
	out.AuthorizationModes = *(*[]string)(unsafe.Pointer(&in.AuthorizationModes))
	out.NoTaintMaster = in.NoTaintMaster
	out.PrivilegedPods = in.PrivilegedPods
	out.Token = in.Token
	out.TokenTTL = (*v1.Duration)(unsafe.Pointer(in.TokenTTL))
	out.CRISocket = in.CRISocket
	out.APIServerExtraArgs = *(*map[string]string)(unsafe.Pointer(&in.APIServerExtraArgs))
	out.ControllerManagerExtraArgs = *(*map[string]string)(unsafe.Pointer(&in.ControllerManagerExtraArgs))
	out.SchedulerExtraArgs = *(*map[string]string)(unsafe.Pointer(&in.SchedulerExtraArgs))
	out.APIServerExtraVolumes = *(*[]kubeadm.HostPathMount)(unsafe.Pointer(&in.APIServerExtraVolumes))
	out.ControllerManagerExtraVolumes = *(*[]kubeadm.HostPathMount)(unsafe.Pointer(&in.ControllerManagerExtraVolumes))
	out.SchedulerExtraVolumes = *(*[]kubeadm.HostPathMount)(unsafe.Pointer(&in.SchedulerExtraVolumes))
	out.APIServerCertSANs = *(*[]string)(unsafe.Pointer(&in.APIServerCertSANs))
	out.CertificatesDir = in.CertificatesDir
	out.ImageRepository = in.ImageRepository
	out.ImagePullPolicy = core_v1.PullPolicy(in.ImagePullPolicy)
	out.UnifiedControlPlaneImage = in.UnifiedControlPlaneImage
	if err := Convert_v1alpha1_AuditPolicyConfiguration_To_kubeadm_AuditPolicyConfiguration(&in.AuditPolicyConfiguration, &out.AuditPolicyConfiguration, s); err != nil {
		return err
	}
	out.FeatureGates = *(*map[string]bool)(unsafe.Pointer(&in.FeatureGates))
	return nil
}

// Convert_v1alpha1_MasterConfiguration_To_kubeadm_MasterConfiguration is an autogenerated conversion function.
func Convert_v1alpha1_MasterConfiguration_To_kubeadm_MasterConfiguration(in *MasterConfiguration, out *kubeadm.MasterConfiguration, s conversion.Scope) error {
	return autoConvert_v1alpha1_MasterConfiguration_To_kubeadm_MasterConfiguration(in, out, s)
}

func autoConvert_kubeadm_MasterConfiguration_To_v1alpha1_MasterConfiguration(in *kubeadm.MasterConfiguration, out *MasterConfiguration, s conversion.Scope) error {
	if err := Convert_kubeadm_API_To_v1alpha1_API(&in.API, &out.API, s); err != nil {
		return err
	}
	if err := Convert_kubeadm_KubeProxy_To_v1alpha1_KubeProxy(&in.KubeProxy, &out.KubeProxy, s); err != nil {
		return err
	}
	if err := Convert_kubeadm_Etcd_To_v1alpha1_Etcd(&in.Etcd, &out.Etcd, s); err != nil {
		return err
	}
	if err := Convert_kubeadm_KubeletConfiguration_To_v1alpha1_KubeletConfiguration(&in.KubeletConfiguration, &out.KubeletConfiguration, s); err != nil {
		return err
	}
	if err := Convert_kubeadm_Networking_To_v1alpha1_Networking(&in.Networking, &out.Networking, s); err != nil {
		return err
	}
	out.KubernetesVersion = in.KubernetesVersion
	out.CloudProvider = in.CloudProvider
	out.NodeName = in.NodeName
	out.AuthorizationModes = *(*[]string)(unsafe.Pointer(&in.AuthorizationModes))
	out.NoTaintMaster = in.NoTaintMaster
	out.PrivilegedPods = in.PrivilegedPods
	out.Token = in.Token
	out.TokenTTL = (*v1.Duration)(unsafe.Pointer(in.TokenTTL))
	out.CRISocket = in.CRISocket
	out.APIServerExtraArgs = *(*map[string]string)(unsafe.Pointer(&in.APIServerExtraArgs))
	out.ControllerManagerExtraArgs = *(*map[string]string)(unsafe.Pointer(&in.ControllerManagerExtraArgs))
	out.SchedulerExtraArgs = *(*map[string]string)(unsafe.Pointer(&in.SchedulerExtraArgs))
	out.APIServerExtraVolumes = *(*[]HostPathMount)(unsafe.Pointer(&in.APIServerExtraVolumes))
	out.ControllerManagerExtraVolumes = *(*[]HostPathMount)(unsafe.Pointer(&in.ControllerManagerExtraVolumes))
	out.SchedulerExtraVolumes = *(*[]HostPathMount)(unsafe.Pointer(&in.SchedulerExtraVolumes))
	out.APIServerCertSANs = *(*[]string)(unsafe.Pointer(&in.APIServerCertSANs))
	out.CertificatesDir = in.CertificatesDir
	out.ImagePullPolicy = core_v1.PullPolicy(in.ImagePullPolicy)
	out.ImageRepository = in.ImageRepository
	// INFO: in.CIImageRepository opted out of conversion generation
	out.UnifiedControlPlaneImage = in.UnifiedControlPlaneImage
	if err := Convert_kubeadm_AuditPolicyConfiguration_To_v1alpha1_AuditPolicyConfiguration(&in.AuditPolicyConfiguration, &out.AuditPolicyConfiguration, s); err != nil {
		return err
	}
	out.FeatureGates = *(*map[string]bool)(unsafe.Pointer(&in.FeatureGates))
	return nil
}

// Convert_kubeadm_MasterConfiguration_To_v1alpha1_MasterConfiguration is an autogenerated conversion function.
func Convert_kubeadm_MasterConfiguration_To_v1alpha1_MasterConfiguration(in *kubeadm.MasterConfiguration, out *MasterConfiguration, s conversion.Scope) error {
	return autoConvert_kubeadm_MasterConfiguration_To_v1alpha1_MasterConfiguration(in, out, s)
}

func autoConvert_v1alpha1_Networking_To_kubeadm_Networking(in *Networking, out *kubeadm.Networking, s conversion.Scope) error {
	out.ServiceSubnet = in.ServiceSubnet
	out.PodSubnet = in.PodSubnet
	out.DNSDomain = in.DNSDomain
	return nil
}

// Convert_v1alpha1_Networking_To_kubeadm_Networking is an autogenerated conversion function.
func Convert_v1alpha1_Networking_To_kubeadm_Networking(in *Networking, out *kubeadm.Networking, s conversion.Scope) error {
	return autoConvert_v1alpha1_Networking_To_kubeadm_Networking(in, out, s)
}

func autoConvert_kubeadm_Networking_To_v1alpha1_Networking(in *kubeadm.Networking, out *Networking, s conversion.Scope) error {
	out.ServiceSubnet = in.ServiceSubnet
	out.PodSubnet = in.PodSubnet
	out.DNSDomain = in.DNSDomain
	return nil
}

// Convert_kubeadm_Networking_To_v1alpha1_Networking is an autogenerated conversion function.
func Convert_kubeadm_Networking_To_v1alpha1_Networking(in *kubeadm.Networking, out *Networking, s conversion.Scope) error {
	return autoConvert_kubeadm_Networking_To_v1alpha1_Networking(in, out, s)
}

func autoConvert_v1alpha1_NodeConfiguration_To_kubeadm_NodeConfiguration(in *NodeConfiguration, out *kubeadm.NodeConfiguration, s conversion.Scope) error {
	out.CACertPath = in.CACertPath
	out.DiscoveryFile = in.DiscoveryFile
	out.DiscoveryToken = in.DiscoveryToken
	out.DiscoveryTokenAPIServers = *(*[]string)(unsafe.Pointer(&in.DiscoveryTokenAPIServers))
	out.NodeName = in.NodeName
	out.TLSBootstrapToken = in.TLSBootstrapToken
	out.Token = in.Token
	out.CRISocket = in.CRISocket
	out.DiscoveryTokenCACertHashes = *(*[]string)(unsafe.Pointer(&in.DiscoveryTokenCACertHashes))
	out.DiscoveryTokenUnsafeSkipCAVerification = in.DiscoveryTokenUnsafeSkipCAVerification
	out.FeatureGates = *(*map[string]bool)(unsafe.Pointer(&in.FeatureGates))
	return nil
}

// Convert_v1alpha1_NodeConfiguration_To_kubeadm_NodeConfiguration is an autogenerated conversion function.
func Convert_v1alpha1_NodeConfiguration_To_kubeadm_NodeConfiguration(in *NodeConfiguration, out *kubeadm.NodeConfiguration, s conversion.Scope) error {
	return autoConvert_v1alpha1_NodeConfiguration_To_kubeadm_NodeConfiguration(in, out, s)
}

func autoConvert_kubeadm_NodeConfiguration_To_v1alpha1_NodeConfiguration(in *kubeadm.NodeConfiguration, out *NodeConfiguration, s conversion.Scope) error {
	out.CACertPath = in.CACertPath
	out.DiscoveryFile = in.DiscoveryFile
	out.DiscoveryToken = in.DiscoveryToken
	out.DiscoveryTokenAPIServers = *(*[]string)(unsafe.Pointer(&in.DiscoveryTokenAPIServers))
	out.NodeName = in.NodeName
	out.TLSBootstrapToken = in.TLSBootstrapToken
	out.Token = in.Token
	out.CRISocket = in.CRISocket
	out.DiscoveryTokenCACertHashes = *(*[]string)(unsafe.Pointer(&in.DiscoveryTokenCACertHashes))
	out.DiscoveryTokenUnsafeSkipCAVerification = in.DiscoveryTokenUnsafeSkipCAVerification
	out.FeatureGates = *(*map[string]bool)(unsafe.Pointer(&in.FeatureGates))
	return nil
}

// Convert_kubeadm_NodeConfiguration_To_v1alpha1_NodeConfiguration is an autogenerated conversion function.
func Convert_kubeadm_NodeConfiguration_To_v1alpha1_NodeConfiguration(in *kubeadm.NodeConfiguration, out *NodeConfiguration, s conversion.Scope) error {
	return autoConvert_kubeadm_NodeConfiguration_To_v1alpha1_NodeConfiguration(in, out, s)
}

func autoConvert_v1alpha1_SelfHostedEtcd_To_kubeadm_SelfHostedEtcd(in *SelfHostedEtcd, out *kubeadm.SelfHostedEtcd, s conversion.Scope) error {
	out.CertificatesDir = in.CertificatesDir
	out.ClusterServiceName = in.ClusterServiceName
	out.EtcdVersion = in.EtcdVersion
	out.OperatorVersion = in.OperatorVersion
	return nil
}

// Convert_v1alpha1_SelfHostedEtcd_To_kubeadm_SelfHostedEtcd is an autogenerated conversion function.
func Convert_v1alpha1_SelfHostedEtcd_To_kubeadm_SelfHostedEtcd(in *SelfHostedEtcd, out *kubeadm.SelfHostedEtcd, s conversion.Scope) error {
	return autoConvert_v1alpha1_SelfHostedEtcd_To_kubeadm_SelfHostedEtcd(in, out, s)
}

func autoConvert_kubeadm_SelfHostedEtcd_To_v1alpha1_SelfHostedEtcd(in *kubeadm.SelfHostedEtcd, out *SelfHostedEtcd, s conversion.Scope) error {
	out.CertificatesDir = in.CertificatesDir
	out.ClusterServiceName = in.ClusterServiceName
	out.EtcdVersion = in.EtcdVersion
	out.OperatorVersion = in.OperatorVersion
	return nil
}

// Convert_kubeadm_SelfHostedEtcd_To_v1alpha1_SelfHostedEtcd is an autogenerated conversion function.
func Convert_kubeadm_SelfHostedEtcd_To_v1alpha1_SelfHostedEtcd(in *kubeadm.SelfHostedEtcd, out *SelfHostedEtcd, s conversion.Scope) error {
	return autoConvert_kubeadm_SelfHostedEtcd_To_v1alpha1_SelfHostedEtcd(in, out, s)
}

func autoConvert_v1alpha1_TokenDiscovery_To_kubeadm_TokenDiscovery(in *TokenDiscovery, out *kubeadm.TokenDiscovery, s conversion.Scope) error {
	out.ID = in.ID
	out.Secret = in.Secret
	return nil
}

// Convert_v1alpha1_TokenDiscovery_To_kubeadm_TokenDiscovery is an autogenerated conversion function.
func Convert_v1alpha1_TokenDiscovery_To_kubeadm_TokenDiscovery(in *TokenDiscovery, out *kubeadm.TokenDiscovery, s conversion.Scope) error {
	return autoConvert_v1alpha1_TokenDiscovery_To_kubeadm_TokenDiscovery(in, out, s)
}

func autoConvert_kubeadm_TokenDiscovery_To_v1alpha1_TokenDiscovery(in *kubeadm.TokenDiscovery, out *TokenDiscovery, s conversion.Scope) error {
	out.ID = in.ID
	out.Secret = in.Secret
	return nil
}

// Convert_kubeadm_TokenDiscovery_To_v1alpha1_TokenDiscovery is an autogenerated conversion function.
func Convert_kubeadm_TokenDiscovery_To_v1alpha1_TokenDiscovery(in *kubeadm.TokenDiscovery, out *TokenDiscovery, s conversion.Scope) error {
	return autoConvert_kubeadm_TokenDiscovery_To_v1alpha1_TokenDiscovery(in, out, s)
}
