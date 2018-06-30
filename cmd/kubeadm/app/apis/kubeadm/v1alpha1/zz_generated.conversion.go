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

package v1alpha1

import (
	unsafe "unsafe"

	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	conversion "k8s.io/apimachinery/pkg/conversion"
	runtime "k8s.io/apimachinery/pkg/runtime"
	kubeadm "k8s.io/kubernetes/cmd/kubeadm/app/apis/kubeadm"
	v1beta1 "k8s.io/kubernetes/pkg/kubelet/apis/kubeletconfig/v1beta1"
	kubeproxyconfigv1alpha1 "k8s.io/kubernetes/pkg/proxy/apis/kubeproxyconfig/v1alpha1"
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
	// WARNING: in.Endpoints requires manual conversion: does not exist in peer-type
	// WARNING: in.CAFile requires manual conversion: does not exist in peer-type
	// WARNING: in.CertFile requires manual conversion: does not exist in peer-type
	// WARNING: in.KeyFile requires manual conversion: does not exist in peer-type
	// WARNING: in.DataDir requires manual conversion: does not exist in peer-type
	// WARNING: in.ExtraArgs requires manual conversion: does not exist in peer-type
	// WARNING: in.Image requires manual conversion: does not exist in peer-type
	// WARNING: in.SelfHosted requires manual conversion: does not exist in peer-type
	// WARNING: in.ServerCertSANs requires manual conversion: does not exist in peer-type
	// WARNING: in.PeerCertSANs requires manual conversion: does not exist in peer-type
	return nil
}

func autoConvert_kubeadm_Etcd_To_v1alpha1_Etcd(in *kubeadm.Etcd, out *Etcd, s conversion.Scope) error {
	// WARNING: in.Local requires manual conversion: does not exist in peer-type
	// WARNING: in.External requires manual conversion: does not exist in peer-type
	return nil
}

func autoConvert_v1alpha1_HostPathMount_To_kubeadm_HostPathMount(in *HostPathMount, out *kubeadm.HostPathMount, s conversion.Scope) error {
	out.Name = in.Name
	out.HostPath = in.HostPath
	out.MountPath = in.MountPath
	out.Writable = in.Writable
	out.PathType = v1.HostPathType(in.PathType)
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
	out.Writable = in.Writable
	out.PathType = v1.HostPathType(in.PathType)
	return nil
}

// Convert_kubeadm_HostPathMount_To_v1alpha1_HostPathMount is an autogenerated conversion function.
func Convert_kubeadm_HostPathMount_To_v1alpha1_HostPathMount(in *kubeadm.HostPathMount, out *HostPathMount, s conversion.Scope) error {
	return autoConvert_kubeadm_HostPathMount_To_v1alpha1_HostPathMount(in, out, s)
}

func autoConvert_v1alpha1_KubeProxy_To_kubeadm_KubeProxy(in *KubeProxy, out *kubeadm.KubeProxy, s conversion.Scope) error {
	out.Config = (*kubeproxyconfigv1alpha1.KubeProxyConfiguration)(unsafe.Pointer(in.Config))
	return nil
}

// Convert_v1alpha1_KubeProxy_To_kubeadm_KubeProxy is an autogenerated conversion function.
func Convert_v1alpha1_KubeProxy_To_kubeadm_KubeProxy(in *KubeProxy, out *kubeadm.KubeProxy, s conversion.Scope) error {
	return autoConvert_v1alpha1_KubeProxy_To_kubeadm_KubeProxy(in, out, s)
}

func autoConvert_kubeadm_KubeProxy_To_v1alpha1_KubeProxy(in *kubeadm.KubeProxy, out *KubeProxy, s conversion.Scope) error {
	out.Config = (*kubeproxyconfigv1alpha1.KubeProxyConfiguration)(unsafe.Pointer(in.Config))
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
	// WARNING: in.CloudProvider requires manual conversion: does not exist in peer-type
	// WARNING: in.NodeName requires manual conversion: does not exist in peer-type
	// WARNING: in.AuthorizationModes requires manual conversion: does not exist in peer-type
	// WARNING: in.NoTaintMaster requires manual conversion: does not exist in peer-type
	// WARNING: in.PrivilegedPods requires manual conversion: does not exist in peer-type
	// WARNING: in.Token requires manual conversion: does not exist in peer-type
	// WARNING: in.TokenTTL requires manual conversion: does not exist in peer-type
	// WARNING: in.TokenUsages requires manual conversion: does not exist in peer-type
	// WARNING: in.TokenGroups requires manual conversion: does not exist in peer-type
	// WARNING: in.CRISocket requires manual conversion: does not exist in peer-type
	out.APIServerExtraArgs = *(*map[string]string)(unsafe.Pointer(&in.APIServerExtraArgs))
	out.ControllerManagerExtraArgs = *(*map[string]string)(unsafe.Pointer(&in.ControllerManagerExtraArgs))
	out.SchedulerExtraArgs = *(*map[string]string)(unsafe.Pointer(&in.SchedulerExtraArgs))
	out.APIServerExtraVolumes = *(*[]kubeadm.HostPathMount)(unsafe.Pointer(&in.APIServerExtraVolumes))
	out.ControllerManagerExtraVolumes = *(*[]kubeadm.HostPathMount)(unsafe.Pointer(&in.ControllerManagerExtraVolumes))
	out.SchedulerExtraVolumes = *(*[]kubeadm.HostPathMount)(unsafe.Pointer(&in.SchedulerExtraVolumes))
	out.APIServerCertSANs = *(*[]string)(unsafe.Pointer(&in.APIServerCertSANs))
	out.CertificatesDir = in.CertificatesDir
	out.ImageRepository = in.ImageRepository
	// WARNING: in.ImagePullPolicy requires manual conversion: does not exist in peer-type
	out.UnifiedControlPlaneImage = in.UnifiedControlPlaneImage
	if err := Convert_v1alpha1_AuditPolicyConfiguration_To_kubeadm_AuditPolicyConfiguration(&in.AuditPolicyConfiguration, &out.AuditPolicyConfiguration, s); err != nil {
		return err
	}
	out.FeatureGates = *(*map[string]bool)(unsafe.Pointer(&in.FeatureGates))
	out.ClusterName = in.ClusterName
	return nil
}

func autoConvert_kubeadm_MasterConfiguration_To_v1alpha1_MasterConfiguration(in *kubeadm.MasterConfiguration, out *MasterConfiguration, s conversion.Scope) error {
	// WARNING: in.BootstrapTokens requires manual conversion: does not exist in peer-type
	// WARNING: in.NodeRegistration requires manual conversion: does not exist in peer-type
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
	if err := Convert_kubeadm_AuditPolicyConfiguration_To_v1alpha1_AuditPolicyConfiguration(&in.AuditPolicyConfiguration, &out.AuditPolicyConfiguration, s); err != nil {
		return err
	}
	out.FeatureGates = *(*map[string]bool)(unsafe.Pointer(&in.FeatureGates))
	out.ClusterName = in.ClusterName
	return nil
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
	out.DiscoveryTimeout = (*metav1.Duration)(unsafe.Pointer(in.DiscoveryTimeout))
	// WARNING: in.NodeName requires manual conversion: does not exist in peer-type
	out.TLSBootstrapToken = in.TLSBootstrapToken
	out.Token = in.Token
	// WARNING: in.CRISocket requires manual conversion: does not exist in peer-type
	out.ClusterName = in.ClusterName
	out.DiscoveryTokenCACertHashes = *(*[]string)(unsafe.Pointer(&in.DiscoveryTokenCACertHashes))
	out.DiscoveryTokenUnsafeSkipCAVerification = in.DiscoveryTokenUnsafeSkipCAVerification
	out.FeatureGates = *(*map[string]bool)(unsafe.Pointer(&in.FeatureGates))
	return nil
}

func autoConvert_kubeadm_NodeConfiguration_To_v1alpha1_NodeConfiguration(in *kubeadm.NodeConfiguration, out *NodeConfiguration, s conversion.Scope) error {
	// WARNING: in.NodeRegistration requires manual conversion: does not exist in peer-type
	out.CACertPath = in.CACertPath
	out.DiscoveryFile = in.DiscoveryFile
	out.DiscoveryToken = in.DiscoveryToken
	out.DiscoveryTokenAPIServers = *(*[]string)(unsafe.Pointer(&in.DiscoveryTokenAPIServers))
	out.DiscoveryTimeout = (*metav1.Duration)(unsafe.Pointer(in.DiscoveryTimeout))
	out.TLSBootstrapToken = in.TLSBootstrapToken
	out.Token = in.Token
	out.ClusterName = in.ClusterName
	out.DiscoveryTokenCACertHashes = *(*[]string)(unsafe.Pointer(&in.DiscoveryTokenCACertHashes))
	out.DiscoveryTokenUnsafeSkipCAVerification = in.DiscoveryTokenUnsafeSkipCAVerification
	out.FeatureGates = *(*map[string]bool)(unsafe.Pointer(&in.FeatureGates))
	return nil
}
