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

package v1alpha3

import (
	unsafe "unsafe"

	corev1 "k8s.io/api/core/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	conversion "k8s.io/apimachinery/pkg/conversion"
	runtime "k8s.io/apimachinery/pkg/runtime"
	kubeadm "k8s.io/kubernetes/cmd/kubeadm/app/apis/kubeadm"
)

func init() {
	localSchemeBuilder.Register(RegisterConversions)
}

// RegisterConversions adds conversion functions to the given scheme.
// Public to allow building arbitrary schemes.
func RegisterConversions(scheme *runtime.Scheme) error {
	return scheme.AddGeneratedConversionFuncs(
		Convert_v1alpha3_API_To_kubeadm_API,
		Convert_kubeadm_API_To_v1alpha3_API,
		Convert_v1alpha3_AuditPolicyConfiguration_To_kubeadm_AuditPolicyConfiguration,
		Convert_kubeadm_AuditPolicyConfiguration_To_v1alpha3_AuditPolicyConfiguration,
		Convert_v1alpha3_BootstrapToken_To_kubeadm_BootstrapToken,
		Convert_kubeadm_BootstrapToken_To_v1alpha3_BootstrapToken,
		Convert_v1alpha3_BootstrapTokenString_To_kubeadm_BootstrapTokenString,
		Convert_kubeadm_BootstrapTokenString_To_v1alpha3_BootstrapTokenString,
		Convert_v1alpha3_Etcd_To_kubeadm_Etcd,
		Convert_kubeadm_Etcd_To_v1alpha3_Etcd,
		Convert_v1alpha3_ExternalEtcd_To_kubeadm_ExternalEtcd,
		Convert_kubeadm_ExternalEtcd_To_v1alpha3_ExternalEtcd,
		Convert_v1alpha3_HostPathMount_To_kubeadm_HostPathMount,
		Convert_kubeadm_HostPathMount_To_v1alpha3_HostPathMount,
		Convert_v1alpha3_LocalEtcd_To_kubeadm_LocalEtcd,
		Convert_kubeadm_LocalEtcd_To_v1alpha3_LocalEtcd,
		Convert_v1alpha3_MasterConfiguration_To_kubeadm_MasterConfiguration,
		Convert_kubeadm_MasterConfiguration_To_v1alpha3_MasterConfiguration,
		Convert_v1alpha3_Networking_To_kubeadm_Networking,
		Convert_kubeadm_Networking_To_v1alpha3_Networking,
		Convert_v1alpha3_NodeConfiguration_To_kubeadm_NodeConfiguration,
		Convert_kubeadm_NodeConfiguration_To_v1alpha3_NodeConfiguration,
		Convert_v1alpha3_NodeRegistrationOptions_To_kubeadm_NodeRegistrationOptions,
		Convert_kubeadm_NodeRegistrationOptions_To_v1alpha3_NodeRegistrationOptions,
	)
}

func autoConvert_v1alpha3_API_To_kubeadm_API(in *API, out *kubeadm.API, s conversion.Scope) error {
	out.AdvertiseAddress = in.AdvertiseAddress
	out.ControlPlaneEndpoint = in.ControlPlaneEndpoint
	out.BindPort = in.BindPort
	return nil
}

// Convert_v1alpha3_API_To_kubeadm_API is an autogenerated conversion function.
func Convert_v1alpha3_API_To_kubeadm_API(in *API, out *kubeadm.API, s conversion.Scope) error {
	return autoConvert_v1alpha3_API_To_kubeadm_API(in, out, s)
}

func autoConvert_kubeadm_API_To_v1alpha3_API(in *kubeadm.API, out *API, s conversion.Scope) error {
	out.AdvertiseAddress = in.AdvertiseAddress
	out.ControlPlaneEndpoint = in.ControlPlaneEndpoint
	out.BindPort = in.BindPort
	return nil
}

// Convert_kubeadm_API_To_v1alpha3_API is an autogenerated conversion function.
func Convert_kubeadm_API_To_v1alpha3_API(in *kubeadm.API, out *API, s conversion.Scope) error {
	return autoConvert_kubeadm_API_To_v1alpha3_API(in, out, s)
}

func autoConvert_v1alpha3_AuditPolicyConfiguration_To_kubeadm_AuditPolicyConfiguration(in *AuditPolicyConfiguration, out *kubeadm.AuditPolicyConfiguration, s conversion.Scope) error {
	out.Path = in.Path
	out.LogDir = in.LogDir
	out.LogMaxAge = (*int32)(unsafe.Pointer(in.LogMaxAge))
	return nil
}

// Convert_v1alpha3_AuditPolicyConfiguration_To_kubeadm_AuditPolicyConfiguration is an autogenerated conversion function.
func Convert_v1alpha3_AuditPolicyConfiguration_To_kubeadm_AuditPolicyConfiguration(in *AuditPolicyConfiguration, out *kubeadm.AuditPolicyConfiguration, s conversion.Scope) error {
	return autoConvert_v1alpha3_AuditPolicyConfiguration_To_kubeadm_AuditPolicyConfiguration(in, out, s)
}

func autoConvert_kubeadm_AuditPolicyConfiguration_To_v1alpha3_AuditPolicyConfiguration(in *kubeadm.AuditPolicyConfiguration, out *AuditPolicyConfiguration, s conversion.Scope) error {
	out.Path = in.Path
	out.LogDir = in.LogDir
	out.LogMaxAge = (*int32)(unsafe.Pointer(in.LogMaxAge))
	return nil
}

// Convert_kubeadm_AuditPolicyConfiguration_To_v1alpha3_AuditPolicyConfiguration is an autogenerated conversion function.
func Convert_kubeadm_AuditPolicyConfiguration_To_v1alpha3_AuditPolicyConfiguration(in *kubeadm.AuditPolicyConfiguration, out *AuditPolicyConfiguration, s conversion.Scope) error {
	return autoConvert_kubeadm_AuditPolicyConfiguration_To_v1alpha3_AuditPolicyConfiguration(in, out, s)
}

func autoConvert_v1alpha3_BootstrapToken_To_kubeadm_BootstrapToken(in *BootstrapToken, out *kubeadm.BootstrapToken, s conversion.Scope) error {
	out.Token = (*kubeadm.BootstrapTokenString)(unsafe.Pointer(in.Token))
	out.Description = in.Description
	out.TTL = (*v1.Duration)(unsafe.Pointer(in.TTL))
	out.Expires = (*v1.Time)(unsafe.Pointer(in.Expires))
	out.Usages = *(*[]string)(unsafe.Pointer(&in.Usages))
	out.Groups = *(*[]string)(unsafe.Pointer(&in.Groups))
	return nil
}

// Convert_v1alpha3_BootstrapToken_To_kubeadm_BootstrapToken is an autogenerated conversion function.
func Convert_v1alpha3_BootstrapToken_To_kubeadm_BootstrapToken(in *BootstrapToken, out *kubeadm.BootstrapToken, s conversion.Scope) error {
	return autoConvert_v1alpha3_BootstrapToken_To_kubeadm_BootstrapToken(in, out, s)
}

func autoConvert_kubeadm_BootstrapToken_To_v1alpha3_BootstrapToken(in *kubeadm.BootstrapToken, out *BootstrapToken, s conversion.Scope) error {
	out.Token = (*BootstrapTokenString)(unsafe.Pointer(in.Token))
	out.Description = in.Description
	out.TTL = (*v1.Duration)(unsafe.Pointer(in.TTL))
	out.Expires = (*v1.Time)(unsafe.Pointer(in.Expires))
	out.Usages = *(*[]string)(unsafe.Pointer(&in.Usages))
	out.Groups = *(*[]string)(unsafe.Pointer(&in.Groups))
	return nil
}

// Convert_kubeadm_BootstrapToken_To_v1alpha3_BootstrapToken is an autogenerated conversion function.
func Convert_kubeadm_BootstrapToken_To_v1alpha3_BootstrapToken(in *kubeadm.BootstrapToken, out *BootstrapToken, s conversion.Scope) error {
	return autoConvert_kubeadm_BootstrapToken_To_v1alpha3_BootstrapToken(in, out, s)
}

func autoConvert_v1alpha3_BootstrapTokenString_To_kubeadm_BootstrapTokenString(in *BootstrapTokenString, out *kubeadm.BootstrapTokenString, s conversion.Scope) error {
	out.ID = in.ID
	out.Secret = in.Secret
	return nil
}

// Convert_v1alpha3_BootstrapTokenString_To_kubeadm_BootstrapTokenString is an autogenerated conversion function.
func Convert_v1alpha3_BootstrapTokenString_To_kubeadm_BootstrapTokenString(in *BootstrapTokenString, out *kubeadm.BootstrapTokenString, s conversion.Scope) error {
	return autoConvert_v1alpha3_BootstrapTokenString_To_kubeadm_BootstrapTokenString(in, out, s)
}

func autoConvert_kubeadm_BootstrapTokenString_To_v1alpha3_BootstrapTokenString(in *kubeadm.BootstrapTokenString, out *BootstrapTokenString, s conversion.Scope) error {
	out.ID = in.ID
	out.Secret = in.Secret
	return nil
}

// Convert_kubeadm_BootstrapTokenString_To_v1alpha3_BootstrapTokenString is an autogenerated conversion function.
func Convert_kubeadm_BootstrapTokenString_To_v1alpha3_BootstrapTokenString(in *kubeadm.BootstrapTokenString, out *BootstrapTokenString, s conversion.Scope) error {
	return autoConvert_kubeadm_BootstrapTokenString_To_v1alpha3_BootstrapTokenString(in, out, s)
}

func autoConvert_v1alpha3_Etcd_To_kubeadm_Etcd(in *Etcd, out *kubeadm.Etcd, s conversion.Scope) error {
	out.Local = (*kubeadm.LocalEtcd)(unsafe.Pointer(in.Local))
	out.External = (*kubeadm.ExternalEtcd)(unsafe.Pointer(in.External))
	return nil
}

// Convert_v1alpha3_Etcd_To_kubeadm_Etcd is an autogenerated conversion function.
func Convert_v1alpha3_Etcd_To_kubeadm_Etcd(in *Etcd, out *kubeadm.Etcd, s conversion.Scope) error {
	return autoConvert_v1alpha3_Etcd_To_kubeadm_Etcd(in, out, s)
}

func autoConvert_kubeadm_Etcd_To_v1alpha3_Etcd(in *kubeadm.Etcd, out *Etcd, s conversion.Scope) error {
	out.Local = (*LocalEtcd)(unsafe.Pointer(in.Local))
	out.External = (*ExternalEtcd)(unsafe.Pointer(in.External))
	return nil
}

// Convert_kubeadm_Etcd_To_v1alpha3_Etcd is an autogenerated conversion function.
func Convert_kubeadm_Etcd_To_v1alpha3_Etcd(in *kubeadm.Etcd, out *Etcd, s conversion.Scope) error {
	return autoConvert_kubeadm_Etcd_To_v1alpha3_Etcd(in, out, s)
}

func autoConvert_v1alpha3_ExternalEtcd_To_kubeadm_ExternalEtcd(in *ExternalEtcd, out *kubeadm.ExternalEtcd, s conversion.Scope) error {
	out.Endpoints = *(*[]string)(unsafe.Pointer(&in.Endpoints))
	out.CAFile = in.CAFile
	out.CertFile = in.CertFile
	out.KeyFile = in.KeyFile
	return nil
}

// Convert_v1alpha3_ExternalEtcd_To_kubeadm_ExternalEtcd is an autogenerated conversion function.
func Convert_v1alpha3_ExternalEtcd_To_kubeadm_ExternalEtcd(in *ExternalEtcd, out *kubeadm.ExternalEtcd, s conversion.Scope) error {
	return autoConvert_v1alpha3_ExternalEtcd_To_kubeadm_ExternalEtcd(in, out, s)
}

func autoConvert_kubeadm_ExternalEtcd_To_v1alpha3_ExternalEtcd(in *kubeadm.ExternalEtcd, out *ExternalEtcd, s conversion.Scope) error {
	out.Endpoints = *(*[]string)(unsafe.Pointer(&in.Endpoints))
	out.CAFile = in.CAFile
	out.CertFile = in.CertFile
	out.KeyFile = in.KeyFile
	return nil
}

// Convert_kubeadm_ExternalEtcd_To_v1alpha3_ExternalEtcd is an autogenerated conversion function.
func Convert_kubeadm_ExternalEtcd_To_v1alpha3_ExternalEtcd(in *kubeadm.ExternalEtcd, out *ExternalEtcd, s conversion.Scope) error {
	return autoConvert_kubeadm_ExternalEtcd_To_v1alpha3_ExternalEtcd(in, out, s)
}

func autoConvert_v1alpha3_HostPathMount_To_kubeadm_HostPathMount(in *HostPathMount, out *kubeadm.HostPathMount, s conversion.Scope) error {
	out.Name = in.Name
	out.HostPath = in.HostPath
	out.MountPath = in.MountPath
	out.Writable = in.Writable
	out.PathType = corev1.HostPathType(in.PathType)
	return nil
}

// Convert_v1alpha3_HostPathMount_To_kubeadm_HostPathMount is an autogenerated conversion function.
func Convert_v1alpha3_HostPathMount_To_kubeadm_HostPathMount(in *HostPathMount, out *kubeadm.HostPathMount, s conversion.Scope) error {
	return autoConvert_v1alpha3_HostPathMount_To_kubeadm_HostPathMount(in, out, s)
}

func autoConvert_kubeadm_HostPathMount_To_v1alpha3_HostPathMount(in *kubeadm.HostPathMount, out *HostPathMount, s conversion.Scope) error {
	out.Name = in.Name
	out.HostPath = in.HostPath
	out.MountPath = in.MountPath
	out.Writable = in.Writable
	out.PathType = corev1.HostPathType(in.PathType)
	return nil
}

// Convert_kubeadm_HostPathMount_To_v1alpha3_HostPathMount is an autogenerated conversion function.
func Convert_kubeadm_HostPathMount_To_v1alpha3_HostPathMount(in *kubeadm.HostPathMount, out *HostPathMount, s conversion.Scope) error {
	return autoConvert_kubeadm_HostPathMount_To_v1alpha3_HostPathMount(in, out, s)
}

func autoConvert_v1alpha3_LocalEtcd_To_kubeadm_LocalEtcd(in *LocalEtcd, out *kubeadm.LocalEtcd, s conversion.Scope) error {
	out.Image = in.Image
	out.DataDir = in.DataDir
	out.ExtraArgs = *(*map[string]string)(unsafe.Pointer(&in.ExtraArgs))
	out.ServerCertSANs = *(*[]string)(unsafe.Pointer(&in.ServerCertSANs))
	out.PeerCertSANs = *(*[]string)(unsafe.Pointer(&in.PeerCertSANs))
	return nil
}

// Convert_v1alpha3_LocalEtcd_To_kubeadm_LocalEtcd is an autogenerated conversion function.
func Convert_v1alpha3_LocalEtcd_To_kubeadm_LocalEtcd(in *LocalEtcd, out *kubeadm.LocalEtcd, s conversion.Scope) error {
	return autoConvert_v1alpha3_LocalEtcd_To_kubeadm_LocalEtcd(in, out, s)
}

func autoConvert_kubeadm_LocalEtcd_To_v1alpha3_LocalEtcd(in *kubeadm.LocalEtcd, out *LocalEtcd, s conversion.Scope) error {
	out.Image = in.Image
	out.DataDir = in.DataDir
	out.ExtraArgs = *(*map[string]string)(unsafe.Pointer(&in.ExtraArgs))
	out.ServerCertSANs = *(*[]string)(unsafe.Pointer(&in.ServerCertSANs))
	out.PeerCertSANs = *(*[]string)(unsafe.Pointer(&in.PeerCertSANs))
	return nil
}

// Convert_kubeadm_LocalEtcd_To_v1alpha3_LocalEtcd is an autogenerated conversion function.
func Convert_kubeadm_LocalEtcd_To_v1alpha3_LocalEtcd(in *kubeadm.LocalEtcd, out *LocalEtcd, s conversion.Scope) error {
	return autoConvert_kubeadm_LocalEtcd_To_v1alpha3_LocalEtcd(in, out, s)
}

func autoConvert_v1alpha3_MasterConfiguration_To_kubeadm_MasterConfiguration(in *MasterConfiguration, out *kubeadm.MasterConfiguration, s conversion.Scope) error {
	out.BootstrapTokens = *(*[]kubeadm.BootstrapToken)(unsafe.Pointer(&in.BootstrapTokens))
	if err := Convert_v1alpha3_NodeRegistrationOptions_To_kubeadm_NodeRegistrationOptions(&in.NodeRegistration, &out.NodeRegistration, s); err != nil {
		return err
	}
	if err := Convert_v1alpha3_API_To_kubeadm_API(&in.API, &out.API, s); err != nil {
		return err
	}
	// WARNING: in.KubeProxy requires manual conversion: does not exist in peer-type
	if err := Convert_v1alpha3_Etcd_To_kubeadm_Etcd(&in.Etcd, &out.Etcd, s); err != nil {
		return err
	}
	// WARNING: in.KubeletConfiguration requires manual conversion: does not exist in peer-type
	if err := Convert_v1alpha3_Networking_To_kubeadm_Networking(&in.Networking, &out.Networking, s); err != nil {
		return err
	}
	out.KubernetesVersion = in.KubernetesVersion
	out.APIServerExtraArgs = *(*map[string]string)(unsafe.Pointer(&in.APIServerExtraArgs))
	out.ControllerManagerExtraArgs = *(*map[string]string)(unsafe.Pointer(&in.ControllerManagerExtraArgs))
	out.SchedulerExtraArgs = *(*map[string]string)(unsafe.Pointer(&in.SchedulerExtraArgs))
	out.APIServerExtraVolumes = *(*[]kubeadm.HostPathMount)(unsafe.Pointer(&in.APIServerExtraVolumes))
	out.ControllerManagerExtraVolumes = *(*[]kubeadm.HostPathMount)(unsafe.Pointer(&in.ControllerManagerExtraVolumes))
	out.SchedulerExtraVolumes = *(*[]kubeadm.HostPathMount)(unsafe.Pointer(&in.SchedulerExtraVolumes))
	out.APIServerCertSANs = *(*[]string)(unsafe.Pointer(&in.APIServerCertSANs))
	out.CertificatesDir = in.CertificatesDir
	out.ImageRepository = in.ImageRepository
	out.UnifiedControlPlaneImage = in.UnifiedControlPlaneImage
	if err := Convert_v1alpha3_AuditPolicyConfiguration_To_kubeadm_AuditPolicyConfiguration(&in.AuditPolicyConfiguration, &out.AuditPolicyConfiguration, s); err != nil {
		return err
	}
	out.FeatureGates = *(*map[string]bool)(unsafe.Pointer(&in.FeatureGates))
	out.ClusterName = in.ClusterName
	return nil
}

func autoConvert_kubeadm_MasterConfiguration_To_v1alpha3_MasterConfiguration(in *kubeadm.MasterConfiguration, out *MasterConfiguration, s conversion.Scope) error {
	out.BootstrapTokens = *(*[]BootstrapToken)(unsafe.Pointer(&in.BootstrapTokens))
	if err := Convert_kubeadm_NodeRegistrationOptions_To_v1alpha3_NodeRegistrationOptions(&in.NodeRegistration, &out.NodeRegistration, s); err != nil {
		return err
	}
	// WARNING: in.ComponentConfigs requires manual conversion: does not exist in peer-type
	if err := Convert_kubeadm_API_To_v1alpha3_API(&in.API, &out.API, s); err != nil {
		return err
	}
	if err := Convert_kubeadm_Etcd_To_v1alpha3_Etcd(&in.Etcd, &out.Etcd, s); err != nil {
		return err
	}
	if err := Convert_kubeadm_Networking_To_v1alpha3_Networking(&in.Networking, &out.Networking, s); err != nil {
		return err
	}
	out.KubernetesVersion = in.KubernetesVersion
	out.APIServerExtraArgs = *(*map[string]string)(unsafe.Pointer(&in.APIServerExtraArgs))
	out.ControllerManagerExtraArgs = *(*map[string]string)(unsafe.Pointer(&in.ControllerManagerExtraArgs))
	out.SchedulerExtraArgs = *(*map[string]string)(unsafe.Pointer(&in.SchedulerExtraArgs))
	out.APIServerExtraVolumes = *(*[]HostPathMount)(unsafe.Pointer(&in.APIServerExtraVolumes))
	out.ControllerManagerExtraVolumes = *(*[]HostPathMount)(unsafe.Pointer(&in.ControllerManagerExtraVolumes))
	out.SchedulerExtraVolumes = *(*[]HostPathMount)(unsafe.Pointer(&in.SchedulerExtraVolumes))
	out.APIServerCertSANs = *(*[]string)(unsafe.Pointer(&in.APIServerCertSANs))
	out.CertificatesDir = in.CertificatesDir
	out.ImageRepository = in.ImageRepository
	// INFO: in.CIImageRepository opted out of conversion generation
	out.UnifiedControlPlaneImage = in.UnifiedControlPlaneImage
	if err := Convert_kubeadm_AuditPolicyConfiguration_To_v1alpha3_AuditPolicyConfiguration(&in.AuditPolicyConfiguration, &out.AuditPolicyConfiguration, s); err != nil {
		return err
	}
	out.FeatureGates = *(*map[string]bool)(unsafe.Pointer(&in.FeatureGates))
	out.ClusterName = in.ClusterName
	return nil
}

func autoConvert_v1alpha3_Networking_To_kubeadm_Networking(in *Networking, out *kubeadm.Networking, s conversion.Scope) error {
	out.ServiceSubnet = in.ServiceSubnet
	out.PodSubnet = in.PodSubnet
	out.DNSDomain = in.DNSDomain
	return nil
}

// Convert_v1alpha3_Networking_To_kubeadm_Networking is an autogenerated conversion function.
func Convert_v1alpha3_Networking_To_kubeadm_Networking(in *Networking, out *kubeadm.Networking, s conversion.Scope) error {
	return autoConvert_v1alpha3_Networking_To_kubeadm_Networking(in, out, s)
}

func autoConvert_kubeadm_Networking_To_v1alpha3_Networking(in *kubeadm.Networking, out *Networking, s conversion.Scope) error {
	out.ServiceSubnet = in.ServiceSubnet
	out.PodSubnet = in.PodSubnet
	out.DNSDomain = in.DNSDomain
	return nil
}

// Convert_kubeadm_Networking_To_v1alpha3_Networking is an autogenerated conversion function.
func Convert_kubeadm_Networking_To_v1alpha3_Networking(in *kubeadm.Networking, out *Networking, s conversion.Scope) error {
	return autoConvert_kubeadm_Networking_To_v1alpha3_Networking(in, out, s)
}

func autoConvert_v1alpha3_NodeConfiguration_To_kubeadm_NodeConfiguration(in *NodeConfiguration, out *kubeadm.NodeConfiguration, s conversion.Scope) error {
	if err := Convert_v1alpha3_NodeRegistrationOptions_To_kubeadm_NodeRegistrationOptions(&in.NodeRegistration, &out.NodeRegistration, s); err != nil {
		return err
	}
	out.CACertPath = in.CACertPath
	out.DiscoveryFile = in.DiscoveryFile
	out.DiscoveryToken = in.DiscoveryToken
	out.DiscoveryTokenAPIServers = *(*[]string)(unsafe.Pointer(&in.DiscoveryTokenAPIServers))
	out.DiscoveryTimeout = (*v1.Duration)(unsafe.Pointer(in.DiscoveryTimeout))
	out.TLSBootstrapToken = in.TLSBootstrapToken
	out.Token = in.Token
	out.ClusterName = in.ClusterName
	out.DiscoveryTokenCACertHashes = *(*[]string)(unsafe.Pointer(&in.DiscoveryTokenCACertHashes))
	out.DiscoveryTokenUnsafeSkipCAVerification = in.DiscoveryTokenUnsafeSkipCAVerification
	out.FeatureGates = *(*map[string]bool)(unsafe.Pointer(&in.FeatureGates))
	return nil
}

// Convert_v1alpha3_NodeConfiguration_To_kubeadm_NodeConfiguration is an autogenerated conversion function.
func Convert_v1alpha3_NodeConfiguration_To_kubeadm_NodeConfiguration(in *NodeConfiguration, out *kubeadm.NodeConfiguration, s conversion.Scope) error {
	return autoConvert_v1alpha3_NodeConfiguration_To_kubeadm_NodeConfiguration(in, out, s)
}

func autoConvert_kubeadm_NodeConfiguration_To_v1alpha3_NodeConfiguration(in *kubeadm.NodeConfiguration, out *NodeConfiguration, s conversion.Scope) error {
	if err := Convert_kubeadm_NodeRegistrationOptions_To_v1alpha3_NodeRegistrationOptions(&in.NodeRegistration, &out.NodeRegistration, s); err != nil {
		return err
	}
	out.CACertPath = in.CACertPath
	out.DiscoveryFile = in.DiscoveryFile
	out.DiscoveryToken = in.DiscoveryToken
	out.DiscoveryTokenAPIServers = *(*[]string)(unsafe.Pointer(&in.DiscoveryTokenAPIServers))
	out.DiscoveryTimeout = (*v1.Duration)(unsafe.Pointer(in.DiscoveryTimeout))
	out.TLSBootstrapToken = in.TLSBootstrapToken
	out.Token = in.Token
	out.ClusterName = in.ClusterName
	out.DiscoveryTokenCACertHashes = *(*[]string)(unsafe.Pointer(&in.DiscoveryTokenCACertHashes))
	out.DiscoveryTokenUnsafeSkipCAVerification = in.DiscoveryTokenUnsafeSkipCAVerification
	out.FeatureGates = *(*map[string]bool)(unsafe.Pointer(&in.FeatureGates))
	return nil
}

// Convert_kubeadm_NodeConfiguration_To_v1alpha3_NodeConfiguration is an autogenerated conversion function.
func Convert_kubeadm_NodeConfiguration_To_v1alpha3_NodeConfiguration(in *kubeadm.NodeConfiguration, out *NodeConfiguration, s conversion.Scope) error {
	return autoConvert_kubeadm_NodeConfiguration_To_v1alpha3_NodeConfiguration(in, out, s)
}

func autoConvert_v1alpha3_NodeRegistrationOptions_To_kubeadm_NodeRegistrationOptions(in *NodeRegistrationOptions, out *kubeadm.NodeRegistrationOptions, s conversion.Scope) error {
	out.Name = in.Name
	out.CRISocket = in.CRISocket
	out.Taints = *(*[]corev1.Taint)(unsafe.Pointer(&in.Taints))
	out.KubeletExtraArgs = *(*map[string]string)(unsafe.Pointer(&in.KubeletExtraArgs))
	return nil
}

// Convert_v1alpha3_NodeRegistrationOptions_To_kubeadm_NodeRegistrationOptions is an autogenerated conversion function.
func Convert_v1alpha3_NodeRegistrationOptions_To_kubeadm_NodeRegistrationOptions(in *NodeRegistrationOptions, out *kubeadm.NodeRegistrationOptions, s conversion.Scope) error {
	return autoConvert_v1alpha3_NodeRegistrationOptions_To_kubeadm_NodeRegistrationOptions(in, out, s)
}

func autoConvert_kubeadm_NodeRegistrationOptions_To_v1alpha3_NodeRegistrationOptions(in *kubeadm.NodeRegistrationOptions, out *NodeRegistrationOptions, s conversion.Scope) error {
	out.Name = in.Name
	out.CRISocket = in.CRISocket
	out.Taints = *(*[]corev1.Taint)(unsafe.Pointer(&in.Taints))
	out.KubeletExtraArgs = *(*map[string]string)(unsafe.Pointer(&in.KubeletExtraArgs))
	return nil
}

// Convert_kubeadm_NodeRegistrationOptions_To_v1alpha3_NodeRegistrationOptions is an autogenerated conversion function.
func Convert_kubeadm_NodeRegistrationOptions_To_v1alpha3_NodeRegistrationOptions(in *kubeadm.NodeRegistrationOptions, out *NodeRegistrationOptions, s conversion.Scope) error {
	return autoConvert_kubeadm_NodeRegistrationOptions_To_v1alpha3_NodeRegistrationOptions(in, out, s)
}
