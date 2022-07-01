//go:build !ignore_autogenerated
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

package v1beta2

import (
	unsafe "unsafe"

	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	conversion "k8s.io/apimachinery/pkg/conversion"
	runtime "k8s.io/apimachinery/pkg/runtime"
	v1 "k8s.io/kubernetes/cmd/kubeadm/app/apis/bootstraptoken/v1"
	kubeadm "k8s.io/kubernetes/cmd/kubeadm/app/apis/kubeadm"
)

func init() {
	localSchemeBuilder.Register(RegisterConversions)
}

// RegisterConversions adds conversion functions to the given scheme.
// Public to allow building arbitrary schemes.
func RegisterConversions(s *runtime.Scheme) error {
	if err := s.AddGeneratedConversionFunc((*APIEndpoint)(nil), (*kubeadm.APIEndpoint)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1beta2_APIEndpoint_To_kubeadm_APIEndpoint(a.(*APIEndpoint), b.(*kubeadm.APIEndpoint), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*kubeadm.APIEndpoint)(nil), (*APIEndpoint)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_kubeadm_APIEndpoint_To_v1beta2_APIEndpoint(a.(*kubeadm.APIEndpoint), b.(*APIEndpoint), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*APIServer)(nil), (*kubeadm.APIServer)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1beta2_APIServer_To_kubeadm_APIServer(a.(*APIServer), b.(*kubeadm.APIServer), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*kubeadm.APIServer)(nil), (*APIServer)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_kubeadm_APIServer_To_v1beta2_APIServer(a.(*kubeadm.APIServer), b.(*APIServer), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*BootstrapTokenDiscovery)(nil), (*kubeadm.BootstrapTokenDiscovery)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1beta2_BootstrapTokenDiscovery_To_kubeadm_BootstrapTokenDiscovery(a.(*BootstrapTokenDiscovery), b.(*kubeadm.BootstrapTokenDiscovery), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*kubeadm.BootstrapTokenDiscovery)(nil), (*BootstrapTokenDiscovery)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_kubeadm_BootstrapTokenDiscovery_To_v1beta2_BootstrapTokenDiscovery(a.(*kubeadm.BootstrapTokenDiscovery), b.(*BootstrapTokenDiscovery), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*kubeadm.ClusterConfiguration)(nil), (*ClusterConfiguration)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_kubeadm_ClusterConfiguration_To_v1beta2_ClusterConfiguration(a.(*kubeadm.ClusterConfiguration), b.(*ClusterConfiguration), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*ControlPlaneComponent)(nil), (*kubeadm.ControlPlaneComponent)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1beta2_ControlPlaneComponent_To_kubeadm_ControlPlaneComponent(a.(*ControlPlaneComponent), b.(*kubeadm.ControlPlaneComponent), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*kubeadm.ControlPlaneComponent)(nil), (*ControlPlaneComponent)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_kubeadm_ControlPlaneComponent_To_v1beta2_ControlPlaneComponent(a.(*kubeadm.ControlPlaneComponent), b.(*ControlPlaneComponent), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*DNS)(nil), (*kubeadm.DNS)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1beta2_DNS_To_kubeadm_DNS(a.(*DNS), b.(*kubeadm.DNS), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*kubeadm.DNS)(nil), (*DNS)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_kubeadm_DNS_To_v1beta2_DNS(a.(*kubeadm.DNS), b.(*DNS), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*Discovery)(nil), (*kubeadm.Discovery)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1beta2_Discovery_To_kubeadm_Discovery(a.(*Discovery), b.(*kubeadm.Discovery), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*kubeadm.Discovery)(nil), (*Discovery)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_kubeadm_Discovery_To_v1beta2_Discovery(a.(*kubeadm.Discovery), b.(*Discovery), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*Etcd)(nil), (*kubeadm.Etcd)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1beta2_Etcd_To_kubeadm_Etcd(a.(*Etcd), b.(*kubeadm.Etcd), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*kubeadm.Etcd)(nil), (*Etcd)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_kubeadm_Etcd_To_v1beta2_Etcd(a.(*kubeadm.Etcd), b.(*Etcd), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*ExternalEtcd)(nil), (*kubeadm.ExternalEtcd)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1beta2_ExternalEtcd_To_kubeadm_ExternalEtcd(a.(*ExternalEtcd), b.(*kubeadm.ExternalEtcd), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*kubeadm.ExternalEtcd)(nil), (*ExternalEtcd)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_kubeadm_ExternalEtcd_To_v1beta2_ExternalEtcd(a.(*kubeadm.ExternalEtcd), b.(*ExternalEtcd), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*FileDiscovery)(nil), (*kubeadm.FileDiscovery)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1beta2_FileDiscovery_To_kubeadm_FileDiscovery(a.(*FileDiscovery), b.(*kubeadm.FileDiscovery), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*kubeadm.FileDiscovery)(nil), (*FileDiscovery)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_kubeadm_FileDiscovery_To_v1beta2_FileDiscovery(a.(*kubeadm.FileDiscovery), b.(*FileDiscovery), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*HostPathMount)(nil), (*kubeadm.HostPathMount)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1beta2_HostPathMount_To_kubeadm_HostPathMount(a.(*HostPathMount), b.(*kubeadm.HostPathMount), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*kubeadm.HostPathMount)(nil), (*HostPathMount)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_kubeadm_HostPathMount_To_v1beta2_HostPathMount(a.(*kubeadm.HostPathMount), b.(*HostPathMount), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*ImageMeta)(nil), (*kubeadm.ImageMeta)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1beta2_ImageMeta_To_kubeadm_ImageMeta(a.(*ImageMeta), b.(*kubeadm.ImageMeta), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*kubeadm.ImageMeta)(nil), (*ImageMeta)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_kubeadm_ImageMeta_To_v1beta2_ImageMeta(a.(*kubeadm.ImageMeta), b.(*ImageMeta), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*JoinControlPlane)(nil), (*kubeadm.JoinControlPlane)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1beta2_JoinControlPlane_To_kubeadm_JoinControlPlane(a.(*JoinControlPlane), b.(*kubeadm.JoinControlPlane), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*kubeadm.JoinControlPlane)(nil), (*JoinControlPlane)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_kubeadm_JoinControlPlane_To_v1beta2_JoinControlPlane(a.(*kubeadm.JoinControlPlane), b.(*JoinControlPlane), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*LocalEtcd)(nil), (*kubeadm.LocalEtcd)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1beta2_LocalEtcd_To_kubeadm_LocalEtcd(a.(*LocalEtcd), b.(*kubeadm.LocalEtcd), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*kubeadm.LocalEtcd)(nil), (*LocalEtcd)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_kubeadm_LocalEtcd_To_v1beta2_LocalEtcd(a.(*kubeadm.LocalEtcd), b.(*LocalEtcd), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*Networking)(nil), (*kubeadm.Networking)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1beta2_Networking_To_kubeadm_Networking(a.(*Networking), b.(*kubeadm.Networking), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*kubeadm.Networking)(nil), (*Networking)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_kubeadm_Networking_To_v1beta2_Networking(a.(*kubeadm.Networking), b.(*Networking), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*NodeRegistrationOptions)(nil), (*kubeadm.NodeRegistrationOptions)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1beta2_NodeRegistrationOptions_To_kubeadm_NodeRegistrationOptions(a.(*NodeRegistrationOptions), b.(*kubeadm.NodeRegistrationOptions), scope)
	}); err != nil {
		return err
	}
	if err := s.AddConversionFunc((*kubeadm.InitConfiguration)(nil), (*InitConfiguration)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_kubeadm_InitConfiguration_To_v1beta2_InitConfiguration(a.(*kubeadm.InitConfiguration), b.(*InitConfiguration), scope)
	}); err != nil {
		return err
	}
	if err := s.AddConversionFunc((*kubeadm.JoinConfiguration)(nil), (*JoinConfiguration)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_kubeadm_JoinConfiguration_To_v1beta2_JoinConfiguration(a.(*kubeadm.JoinConfiguration), b.(*JoinConfiguration), scope)
	}); err != nil {
		return err
	}
	if err := s.AddConversionFunc((*kubeadm.NodeRegistrationOptions)(nil), (*NodeRegistrationOptions)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_kubeadm_NodeRegistrationOptions_To_v1beta2_NodeRegistrationOptions(a.(*kubeadm.NodeRegistrationOptions), b.(*NodeRegistrationOptions), scope)
	}); err != nil {
		return err
	}
	if err := s.AddConversionFunc((*v1.BootstrapToken)(nil), (*BootstrapToken)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1_BootstrapToken_To_v1beta2_BootstrapToken(a.(*v1.BootstrapToken), b.(*BootstrapToken), scope)
	}); err != nil {
		return err
	}
	if err := s.AddConversionFunc((*BootstrapToken)(nil), (*v1.BootstrapToken)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1beta2_BootstrapToken_To_v1_BootstrapToken(a.(*BootstrapToken), b.(*v1.BootstrapToken), scope)
	}); err != nil {
		return err
	}
	if err := s.AddConversionFunc((*ClusterConfiguration)(nil), (*kubeadm.ClusterConfiguration)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1beta2_ClusterConfiguration_To_kubeadm_ClusterConfiguration(a.(*ClusterConfiguration), b.(*kubeadm.ClusterConfiguration), scope)
	}); err != nil {
		return err
	}
	if err := s.AddConversionFunc((*InitConfiguration)(nil), (*kubeadm.InitConfiguration)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1beta2_InitConfiguration_To_kubeadm_InitConfiguration(a.(*InitConfiguration), b.(*kubeadm.InitConfiguration), scope)
	}); err != nil {
		return err
	}
	if err := s.AddConversionFunc((*JoinConfiguration)(nil), (*kubeadm.JoinConfiguration)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1beta2_JoinConfiguration_To_kubeadm_JoinConfiguration(a.(*JoinConfiguration), b.(*kubeadm.JoinConfiguration), scope)
	}); err != nil {
		return err
	}
	return nil
}

func autoConvert_v1beta2_APIEndpoint_To_kubeadm_APIEndpoint(in *APIEndpoint, out *kubeadm.APIEndpoint, s conversion.Scope) error {
	out.AdvertiseAddress = in.AdvertiseAddress
	out.BindPort = in.BindPort
	return nil
}

// Convert_v1beta2_APIEndpoint_To_kubeadm_APIEndpoint is an autogenerated conversion function.
func Convert_v1beta2_APIEndpoint_To_kubeadm_APIEndpoint(in *APIEndpoint, out *kubeadm.APIEndpoint, s conversion.Scope) error {
	return autoConvert_v1beta2_APIEndpoint_To_kubeadm_APIEndpoint(in, out, s)
}

func autoConvert_kubeadm_APIEndpoint_To_v1beta2_APIEndpoint(in *kubeadm.APIEndpoint, out *APIEndpoint, s conversion.Scope) error {
	out.AdvertiseAddress = in.AdvertiseAddress
	out.BindPort = in.BindPort
	return nil
}

// Convert_kubeadm_APIEndpoint_To_v1beta2_APIEndpoint is an autogenerated conversion function.
func Convert_kubeadm_APIEndpoint_To_v1beta2_APIEndpoint(in *kubeadm.APIEndpoint, out *APIEndpoint, s conversion.Scope) error {
	return autoConvert_kubeadm_APIEndpoint_To_v1beta2_APIEndpoint(in, out, s)
}

func autoConvert_v1beta2_APIServer_To_kubeadm_APIServer(in *APIServer, out *kubeadm.APIServer, s conversion.Scope) error {
	if err := Convert_v1beta2_ControlPlaneComponent_To_kubeadm_ControlPlaneComponent(&in.ControlPlaneComponent, &out.ControlPlaneComponent, s); err != nil {
		return err
	}
	out.CertSANs = *(*[]string)(unsafe.Pointer(&in.CertSANs))
	out.TimeoutForControlPlane = (*metav1.Duration)(unsafe.Pointer(in.TimeoutForControlPlane))
	return nil
}

// Convert_v1beta2_APIServer_To_kubeadm_APIServer is an autogenerated conversion function.
func Convert_v1beta2_APIServer_To_kubeadm_APIServer(in *APIServer, out *kubeadm.APIServer, s conversion.Scope) error {
	return autoConvert_v1beta2_APIServer_To_kubeadm_APIServer(in, out, s)
}

func autoConvert_kubeadm_APIServer_To_v1beta2_APIServer(in *kubeadm.APIServer, out *APIServer, s conversion.Scope) error {
	if err := Convert_kubeadm_ControlPlaneComponent_To_v1beta2_ControlPlaneComponent(&in.ControlPlaneComponent, &out.ControlPlaneComponent, s); err != nil {
		return err
	}
	out.CertSANs = *(*[]string)(unsafe.Pointer(&in.CertSANs))
	out.TimeoutForControlPlane = (*metav1.Duration)(unsafe.Pointer(in.TimeoutForControlPlane))
	return nil
}

// Convert_kubeadm_APIServer_To_v1beta2_APIServer is an autogenerated conversion function.
func Convert_kubeadm_APIServer_To_v1beta2_APIServer(in *kubeadm.APIServer, out *APIServer, s conversion.Scope) error {
	return autoConvert_kubeadm_APIServer_To_v1beta2_APIServer(in, out, s)
}

func autoConvert_v1beta2_BootstrapTokenDiscovery_To_kubeadm_BootstrapTokenDiscovery(in *BootstrapTokenDiscovery, out *kubeadm.BootstrapTokenDiscovery, s conversion.Scope) error {
	out.Token = in.Token
	out.APIServerEndpoint = in.APIServerEndpoint
	out.CACertHashes = *(*[]string)(unsafe.Pointer(&in.CACertHashes))
	out.UnsafeSkipCAVerification = in.UnsafeSkipCAVerification
	return nil
}

// Convert_v1beta2_BootstrapTokenDiscovery_To_kubeadm_BootstrapTokenDiscovery is an autogenerated conversion function.
func Convert_v1beta2_BootstrapTokenDiscovery_To_kubeadm_BootstrapTokenDiscovery(in *BootstrapTokenDiscovery, out *kubeadm.BootstrapTokenDiscovery, s conversion.Scope) error {
	return autoConvert_v1beta2_BootstrapTokenDiscovery_To_kubeadm_BootstrapTokenDiscovery(in, out, s)
}

func autoConvert_kubeadm_BootstrapTokenDiscovery_To_v1beta2_BootstrapTokenDiscovery(in *kubeadm.BootstrapTokenDiscovery, out *BootstrapTokenDiscovery, s conversion.Scope) error {
	out.Token = in.Token
	out.APIServerEndpoint = in.APIServerEndpoint
	out.CACertHashes = *(*[]string)(unsafe.Pointer(&in.CACertHashes))
	out.UnsafeSkipCAVerification = in.UnsafeSkipCAVerification
	return nil
}

// Convert_kubeadm_BootstrapTokenDiscovery_To_v1beta2_BootstrapTokenDiscovery is an autogenerated conversion function.
func Convert_kubeadm_BootstrapTokenDiscovery_To_v1beta2_BootstrapTokenDiscovery(in *kubeadm.BootstrapTokenDiscovery, out *BootstrapTokenDiscovery, s conversion.Scope) error {
	return autoConvert_kubeadm_BootstrapTokenDiscovery_To_v1beta2_BootstrapTokenDiscovery(in, out, s)
}

func autoConvert_v1beta2_ClusterConfiguration_To_kubeadm_ClusterConfiguration(in *ClusterConfiguration, out *kubeadm.ClusterConfiguration, s conversion.Scope) error {
	if err := Convert_v1beta2_Etcd_To_kubeadm_Etcd(&in.Etcd, &out.Etcd, s); err != nil {
		return err
	}
	if err := Convert_v1beta2_Networking_To_kubeadm_Networking(&in.Networking, &out.Networking, s); err != nil {
		return err
	}
	out.KubernetesVersion = in.KubernetesVersion
	out.ControlPlaneEndpoint = in.ControlPlaneEndpoint
	if err := Convert_v1beta2_APIServer_To_kubeadm_APIServer(&in.APIServer, &out.APIServer, s); err != nil {
		return err
	}
	if err := Convert_v1beta2_ControlPlaneComponent_To_kubeadm_ControlPlaneComponent(&in.ControllerManager, &out.ControllerManager, s); err != nil {
		return err
	}
	if err := Convert_v1beta2_ControlPlaneComponent_To_kubeadm_ControlPlaneComponent(&in.Scheduler, &out.Scheduler, s); err != nil {
		return err
	}
	if err := Convert_v1beta2_DNS_To_kubeadm_DNS(&in.DNS, &out.DNS, s); err != nil {
		return err
	}
	out.CertificatesDir = in.CertificatesDir
	out.ImageRepository = in.ImageRepository
	// WARNING: in.UseHyperKubeImage requires manual conversion: does not exist in peer-type
	out.FeatureGates = *(*map[string]bool)(unsafe.Pointer(&in.FeatureGates))
	out.ClusterName = in.ClusterName
	return nil
}

func autoConvert_kubeadm_ClusterConfiguration_To_v1beta2_ClusterConfiguration(in *kubeadm.ClusterConfiguration, out *ClusterConfiguration, s conversion.Scope) error {
	// INFO: in.ComponentConfigs opted out of conversion generation
	if err := Convert_kubeadm_Etcd_To_v1beta2_Etcd(&in.Etcd, &out.Etcd, s); err != nil {
		return err
	}
	if err := Convert_kubeadm_Networking_To_v1beta2_Networking(&in.Networking, &out.Networking, s); err != nil {
		return err
	}
	out.KubernetesVersion = in.KubernetesVersion
	// INFO: in.CIKubernetesVersion opted out of conversion generation
	out.ControlPlaneEndpoint = in.ControlPlaneEndpoint
	if err := Convert_kubeadm_APIServer_To_v1beta2_APIServer(&in.APIServer, &out.APIServer, s); err != nil {
		return err
	}
	if err := Convert_kubeadm_ControlPlaneComponent_To_v1beta2_ControlPlaneComponent(&in.ControllerManager, &out.ControllerManager, s); err != nil {
		return err
	}
	if err := Convert_kubeadm_ControlPlaneComponent_To_v1beta2_ControlPlaneComponent(&in.Scheduler, &out.Scheduler, s); err != nil {
		return err
	}
	if err := Convert_kubeadm_DNS_To_v1beta2_DNS(&in.DNS, &out.DNS, s); err != nil {
		return err
	}
	out.CertificatesDir = in.CertificatesDir
	out.ImageRepository = in.ImageRepository
	// INFO: in.CIImageRepository opted out of conversion generation
	out.FeatureGates = *(*map[string]bool)(unsafe.Pointer(&in.FeatureGates))
	out.ClusterName = in.ClusterName
	return nil
}

// Convert_kubeadm_ClusterConfiguration_To_v1beta2_ClusterConfiguration is an autogenerated conversion function.
func Convert_kubeadm_ClusterConfiguration_To_v1beta2_ClusterConfiguration(in *kubeadm.ClusterConfiguration, out *ClusterConfiguration, s conversion.Scope) error {
	return autoConvert_kubeadm_ClusterConfiguration_To_v1beta2_ClusterConfiguration(in, out, s)
}

func autoConvert_v1beta2_ControlPlaneComponent_To_kubeadm_ControlPlaneComponent(in *ControlPlaneComponent, out *kubeadm.ControlPlaneComponent, s conversion.Scope) error {
	out.ExtraArgs = *(*map[string]string)(unsafe.Pointer(&in.ExtraArgs))
	out.ExtraVolumes = *(*[]kubeadm.HostPathMount)(unsafe.Pointer(&in.ExtraVolumes))
	return nil
}

// Convert_v1beta2_ControlPlaneComponent_To_kubeadm_ControlPlaneComponent is an autogenerated conversion function.
func Convert_v1beta2_ControlPlaneComponent_To_kubeadm_ControlPlaneComponent(in *ControlPlaneComponent, out *kubeadm.ControlPlaneComponent, s conversion.Scope) error {
	return autoConvert_v1beta2_ControlPlaneComponent_To_kubeadm_ControlPlaneComponent(in, out, s)
}

func autoConvert_kubeadm_ControlPlaneComponent_To_v1beta2_ControlPlaneComponent(in *kubeadm.ControlPlaneComponent, out *ControlPlaneComponent, s conversion.Scope) error {
	out.ExtraArgs = *(*map[string]string)(unsafe.Pointer(&in.ExtraArgs))
	out.ExtraVolumes = *(*[]HostPathMount)(unsafe.Pointer(&in.ExtraVolumes))
	return nil
}

// Convert_kubeadm_ControlPlaneComponent_To_v1beta2_ControlPlaneComponent is an autogenerated conversion function.
func Convert_kubeadm_ControlPlaneComponent_To_v1beta2_ControlPlaneComponent(in *kubeadm.ControlPlaneComponent, out *ControlPlaneComponent, s conversion.Scope) error {
	return autoConvert_kubeadm_ControlPlaneComponent_To_v1beta2_ControlPlaneComponent(in, out, s)
}

func autoConvert_v1beta2_DNS_To_kubeadm_DNS(in *DNS, out *kubeadm.DNS, s conversion.Scope) error {
	out.Type = kubeadm.DNSAddOnType(in.Type)
	if err := Convert_v1beta2_ImageMeta_To_kubeadm_ImageMeta(&in.ImageMeta, &out.ImageMeta, s); err != nil {
		return err
	}
	return nil
}

// Convert_v1beta2_DNS_To_kubeadm_DNS is an autogenerated conversion function.
func Convert_v1beta2_DNS_To_kubeadm_DNS(in *DNS, out *kubeadm.DNS, s conversion.Scope) error {
	return autoConvert_v1beta2_DNS_To_kubeadm_DNS(in, out, s)
}

func autoConvert_kubeadm_DNS_To_v1beta2_DNS(in *kubeadm.DNS, out *DNS, s conversion.Scope) error {
	out.Type = DNSAddOnType(in.Type)
	if err := Convert_kubeadm_ImageMeta_To_v1beta2_ImageMeta(&in.ImageMeta, &out.ImageMeta, s); err != nil {
		return err
	}
	return nil
}

// Convert_kubeadm_DNS_To_v1beta2_DNS is an autogenerated conversion function.
func Convert_kubeadm_DNS_To_v1beta2_DNS(in *kubeadm.DNS, out *DNS, s conversion.Scope) error {
	return autoConvert_kubeadm_DNS_To_v1beta2_DNS(in, out, s)
}

func autoConvert_v1beta2_Discovery_To_kubeadm_Discovery(in *Discovery, out *kubeadm.Discovery, s conversion.Scope) error {
	out.BootstrapToken = (*kubeadm.BootstrapTokenDiscovery)(unsafe.Pointer(in.BootstrapToken))
	out.File = (*kubeadm.FileDiscovery)(unsafe.Pointer(in.File))
	out.TLSBootstrapToken = in.TLSBootstrapToken
	out.Timeout = (*metav1.Duration)(unsafe.Pointer(in.Timeout))
	return nil
}

// Convert_v1beta2_Discovery_To_kubeadm_Discovery is an autogenerated conversion function.
func Convert_v1beta2_Discovery_To_kubeadm_Discovery(in *Discovery, out *kubeadm.Discovery, s conversion.Scope) error {
	return autoConvert_v1beta2_Discovery_To_kubeadm_Discovery(in, out, s)
}

func autoConvert_kubeadm_Discovery_To_v1beta2_Discovery(in *kubeadm.Discovery, out *Discovery, s conversion.Scope) error {
	out.BootstrapToken = (*BootstrapTokenDiscovery)(unsafe.Pointer(in.BootstrapToken))
	out.File = (*FileDiscovery)(unsafe.Pointer(in.File))
	out.TLSBootstrapToken = in.TLSBootstrapToken
	out.Timeout = (*metav1.Duration)(unsafe.Pointer(in.Timeout))
	return nil
}

// Convert_kubeadm_Discovery_To_v1beta2_Discovery is an autogenerated conversion function.
func Convert_kubeadm_Discovery_To_v1beta2_Discovery(in *kubeadm.Discovery, out *Discovery, s conversion.Scope) error {
	return autoConvert_kubeadm_Discovery_To_v1beta2_Discovery(in, out, s)
}

func autoConvert_v1beta2_Etcd_To_kubeadm_Etcd(in *Etcd, out *kubeadm.Etcd, s conversion.Scope) error {
	out.Local = (*kubeadm.LocalEtcd)(unsafe.Pointer(in.Local))
	out.External = (*kubeadm.ExternalEtcd)(unsafe.Pointer(in.External))
	return nil
}

// Convert_v1beta2_Etcd_To_kubeadm_Etcd is an autogenerated conversion function.
func Convert_v1beta2_Etcd_To_kubeadm_Etcd(in *Etcd, out *kubeadm.Etcd, s conversion.Scope) error {
	return autoConvert_v1beta2_Etcd_To_kubeadm_Etcd(in, out, s)
}

func autoConvert_kubeadm_Etcd_To_v1beta2_Etcd(in *kubeadm.Etcd, out *Etcd, s conversion.Scope) error {
	out.Local = (*LocalEtcd)(unsafe.Pointer(in.Local))
	out.External = (*ExternalEtcd)(unsafe.Pointer(in.External))
	return nil
}

// Convert_kubeadm_Etcd_To_v1beta2_Etcd is an autogenerated conversion function.
func Convert_kubeadm_Etcd_To_v1beta2_Etcd(in *kubeadm.Etcd, out *Etcd, s conversion.Scope) error {
	return autoConvert_kubeadm_Etcd_To_v1beta2_Etcd(in, out, s)
}

func autoConvert_v1beta2_ExternalEtcd_To_kubeadm_ExternalEtcd(in *ExternalEtcd, out *kubeadm.ExternalEtcd, s conversion.Scope) error {
	out.Endpoints = *(*[]string)(unsafe.Pointer(&in.Endpoints))
	out.CAFile = in.CAFile
	out.CertFile = in.CertFile
	out.KeyFile = in.KeyFile
	return nil
}

// Convert_v1beta2_ExternalEtcd_To_kubeadm_ExternalEtcd is an autogenerated conversion function.
func Convert_v1beta2_ExternalEtcd_To_kubeadm_ExternalEtcd(in *ExternalEtcd, out *kubeadm.ExternalEtcd, s conversion.Scope) error {
	return autoConvert_v1beta2_ExternalEtcd_To_kubeadm_ExternalEtcd(in, out, s)
}

func autoConvert_kubeadm_ExternalEtcd_To_v1beta2_ExternalEtcd(in *kubeadm.ExternalEtcd, out *ExternalEtcd, s conversion.Scope) error {
	out.Endpoints = *(*[]string)(unsafe.Pointer(&in.Endpoints))
	out.CAFile = in.CAFile
	out.CertFile = in.CertFile
	out.KeyFile = in.KeyFile
	return nil
}

// Convert_kubeadm_ExternalEtcd_To_v1beta2_ExternalEtcd is an autogenerated conversion function.
func Convert_kubeadm_ExternalEtcd_To_v1beta2_ExternalEtcd(in *kubeadm.ExternalEtcd, out *ExternalEtcd, s conversion.Scope) error {
	return autoConvert_kubeadm_ExternalEtcd_To_v1beta2_ExternalEtcd(in, out, s)
}

func autoConvert_v1beta2_FileDiscovery_To_kubeadm_FileDiscovery(in *FileDiscovery, out *kubeadm.FileDiscovery, s conversion.Scope) error {
	out.KubeConfigPath = in.KubeConfigPath
	return nil
}

// Convert_v1beta2_FileDiscovery_To_kubeadm_FileDiscovery is an autogenerated conversion function.
func Convert_v1beta2_FileDiscovery_To_kubeadm_FileDiscovery(in *FileDiscovery, out *kubeadm.FileDiscovery, s conversion.Scope) error {
	return autoConvert_v1beta2_FileDiscovery_To_kubeadm_FileDiscovery(in, out, s)
}

func autoConvert_kubeadm_FileDiscovery_To_v1beta2_FileDiscovery(in *kubeadm.FileDiscovery, out *FileDiscovery, s conversion.Scope) error {
	out.KubeConfigPath = in.KubeConfigPath
	return nil
}

// Convert_kubeadm_FileDiscovery_To_v1beta2_FileDiscovery is an autogenerated conversion function.
func Convert_kubeadm_FileDiscovery_To_v1beta2_FileDiscovery(in *kubeadm.FileDiscovery, out *FileDiscovery, s conversion.Scope) error {
	return autoConvert_kubeadm_FileDiscovery_To_v1beta2_FileDiscovery(in, out, s)
}

func autoConvert_v1beta2_HostPathMount_To_kubeadm_HostPathMount(in *HostPathMount, out *kubeadm.HostPathMount, s conversion.Scope) error {
	out.Name = in.Name
	out.HostPath = in.HostPath
	out.MountPath = in.MountPath
	out.ReadOnly = in.ReadOnly
	out.PathType = corev1.HostPathType(in.PathType)
	return nil
}

// Convert_v1beta2_HostPathMount_To_kubeadm_HostPathMount is an autogenerated conversion function.
func Convert_v1beta2_HostPathMount_To_kubeadm_HostPathMount(in *HostPathMount, out *kubeadm.HostPathMount, s conversion.Scope) error {
	return autoConvert_v1beta2_HostPathMount_To_kubeadm_HostPathMount(in, out, s)
}

func autoConvert_kubeadm_HostPathMount_To_v1beta2_HostPathMount(in *kubeadm.HostPathMount, out *HostPathMount, s conversion.Scope) error {
	out.Name = in.Name
	out.HostPath = in.HostPath
	out.MountPath = in.MountPath
	out.ReadOnly = in.ReadOnly
	out.PathType = corev1.HostPathType(in.PathType)
	return nil
}

// Convert_kubeadm_HostPathMount_To_v1beta2_HostPathMount is an autogenerated conversion function.
func Convert_kubeadm_HostPathMount_To_v1beta2_HostPathMount(in *kubeadm.HostPathMount, out *HostPathMount, s conversion.Scope) error {
	return autoConvert_kubeadm_HostPathMount_To_v1beta2_HostPathMount(in, out, s)
}

func autoConvert_v1beta2_ImageMeta_To_kubeadm_ImageMeta(in *ImageMeta, out *kubeadm.ImageMeta, s conversion.Scope) error {
	out.ImageRepository = in.ImageRepository
	out.ImageTag = in.ImageTag
	return nil
}

// Convert_v1beta2_ImageMeta_To_kubeadm_ImageMeta is an autogenerated conversion function.
func Convert_v1beta2_ImageMeta_To_kubeadm_ImageMeta(in *ImageMeta, out *kubeadm.ImageMeta, s conversion.Scope) error {
	return autoConvert_v1beta2_ImageMeta_To_kubeadm_ImageMeta(in, out, s)
}

func autoConvert_kubeadm_ImageMeta_To_v1beta2_ImageMeta(in *kubeadm.ImageMeta, out *ImageMeta, s conversion.Scope) error {
	out.ImageRepository = in.ImageRepository
	out.ImageTag = in.ImageTag
	return nil
}

// Convert_kubeadm_ImageMeta_To_v1beta2_ImageMeta is an autogenerated conversion function.
func Convert_kubeadm_ImageMeta_To_v1beta2_ImageMeta(in *kubeadm.ImageMeta, out *ImageMeta, s conversion.Scope) error {
	return autoConvert_kubeadm_ImageMeta_To_v1beta2_ImageMeta(in, out, s)
}

func autoConvert_v1beta2_InitConfiguration_To_kubeadm_InitConfiguration(in *InitConfiguration, out *kubeadm.InitConfiguration, s conversion.Scope) error {
	if in.BootstrapTokens != nil {
		in, out := &in.BootstrapTokens, &out.BootstrapTokens
		*out = make([]v1.BootstrapToken, len(*in))
		for i := range *in {
			if err := Convert_v1beta2_BootstrapToken_To_v1_BootstrapToken(&(*in)[i], &(*out)[i], s); err != nil {
				return err
			}
		}
	} else {
		out.BootstrapTokens = nil
	}
	if err := Convert_v1beta2_NodeRegistrationOptions_To_kubeadm_NodeRegistrationOptions(&in.NodeRegistration, &out.NodeRegistration, s); err != nil {
		return err
	}
	if err := Convert_v1beta2_APIEndpoint_To_kubeadm_APIEndpoint(&in.LocalAPIEndpoint, &out.LocalAPIEndpoint, s); err != nil {
		return err
	}
	out.CertificateKey = in.CertificateKey
	return nil
}

func autoConvert_kubeadm_InitConfiguration_To_v1beta2_InitConfiguration(in *kubeadm.InitConfiguration, out *InitConfiguration, s conversion.Scope) error {
	// WARNING: in.ClusterConfiguration requires manual conversion: does not exist in peer-type
	if in.BootstrapTokens != nil {
		in, out := &in.BootstrapTokens, &out.BootstrapTokens
		*out = make([]BootstrapToken, len(*in))
		for i := range *in {
			if err := Convert_v1_BootstrapToken_To_v1beta2_BootstrapToken(&(*in)[i], &(*out)[i], s); err != nil {
				return err
			}
		}
	} else {
		out.BootstrapTokens = nil
	}
	if err := Convert_kubeadm_NodeRegistrationOptions_To_v1beta2_NodeRegistrationOptions(&in.NodeRegistration, &out.NodeRegistration, s); err != nil {
		return err
	}
	if err := Convert_kubeadm_APIEndpoint_To_v1beta2_APIEndpoint(&in.LocalAPIEndpoint, &out.LocalAPIEndpoint, s); err != nil {
		return err
	}
	out.CertificateKey = in.CertificateKey
	// WARNING: in.SkipPhases requires manual conversion: does not exist in peer-type
	// WARNING: in.Patches requires manual conversion: does not exist in peer-type
	return nil
}

func autoConvert_v1beta2_JoinConfiguration_To_kubeadm_JoinConfiguration(in *JoinConfiguration, out *kubeadm.JoinConfiguration, s conversion.Scope) error {
	if err := Convert_v1beta2_NodeRegistrationOptions_To_kubeadm_NodeRegistrationOptions(&in.NodeRegistration, &out.NodeRegistration, s); err != nil {
		return err
	}
	out.CACertPath = in.CACertPath
	if err := Convert_v1beta2_Discovery_To_kubeadm_Discovery(&in.Discovery, &out.Discovery, s); err != nil {
		return err
	}
	out.ControlPlane = (*kubeadm.JoinControlPlane)(unsafe.Pointer(in.ControlPlane))
	return nil
}

func autoConvert_kubeadm_JoinConfiguration_To_v1beta2_JoinConfiguration(in *kubeadm.JoinConfiguration, out *JoinConfiguration, s conversion.Scope) error {
	if err := Convert_kubeadm_NodeRegistrationOptions_To_v1beta2_NodeRegistrationOptions(&in.NodeRegistration, &out.NodeRegistration, s); err != nil {
		return err
	}
	out.CACertPath = in.CACertPath
	if err := Convert_kubeadm_Discovery_To_v1beta2_Discovery(&in.Discovery, &out.Discovery, s); err != nil {
		return err
	}
	out.ControlPlane = (*JoinControlPlane)(unsafe.Pointer(in.ControlPlane))
	// WARNING: in.SkipPhases requires manual conversion: does not exist in peer-type
	// WARNING: in.Patches requires manual conversion: does not exist in peer-type
	return nil
}

func autoConvert_v1beta2_JoinControlPlane_To_kubeadm_JoinControlPlane(in *JoinControlPlane, out *kubeadm.JoinControlPlane, s conversion.Scope) error {
	if err := Convert_v1beta2_APIEndpoint_To_kubeadm_APIEndpoint(&in.LocalAPIEndpoint, &out.LocalAPIEndpoint, s); err != nil {
		return err
	}
	out.CertificateKey = in.CertificateKey
	return nil
}

// Convert_v1beta2_JoinControlPlane_To_kubeadm_JoinControlPlane is an autogenerated conversion function.
func Convert_v1beta2_JoinControlPlane_To_kubeadm_JoinControlPlane(in *JoinControlPlane, out *kubeadm.JoinControlPlane, s conversion.Scope) error {
	return autoConvert_v1beta2_JoinControlPlane_To_kubeadm_JoinControlPlane(in, out, s)
}

func autoConvert_kubeadm_JoinControlPlane_To_v1beta2_JoinControlPlane(in *kubeadm.JoinControlPlane, out *JoinControlPlane, s conversion.Scope) error {
	if err := Convert_kubeadm_APIEndpoint_To_v1beta2_APIEndpoint(&in.LocalAPIEndpoint, &out.LocalAPIEndpoint, s); err != nil {
		return err
	}
	out.CertificateKey = in.CertificateKey
	return nil
}

// Convert_kubeadm_JoinControlPlane_To_v1beta2_JoinControlPlane is an autogenerated conversion function.
func Convert_kubeadm_JoinControlPlane_To_v1beta2_JoinControlPlane(in *kubeadm.JoinControlPlane, out *JoinControlPlane, s conversion.Scope) error {
	return autoConvert_kubeadm_JoinControlPlane_To_v1beta2_JoinControlPlane(in, out, s)
}

func autoConvert_v1beta2_LocalEtcd_To_kubeadm_LocalEtcd(in *LocalEtcd, out *kubeadm.LocalEtcd, s conversion.Scope) error {
	if err := Convert_v1beta2_ImageMeta_To_kubeadm_ImageMeta(&in.ImageMeta, &out.ImageMeta, s); err != nil {
		return err
	}
	out.DataDir = in.DataDir
	out.ExtraArgs = *(*map[string]string)(unsafe.Pointer(&in.ExtraArgs))
	out.ServerCertSANs = *(*[]string)(unsafe.Pointer(&in.ServerCertSANs))
	out.PeerCertSANs = *(*[]string)(unsafe.Pointer(&in.PeerCertSANs))
	return nil
}

// Convert_v1beta2_LocalEtcd_To_kubeadm_LocalEtcd is an autogenerated conversion function.
func Convert_v1beta2_LocalEtcd_To_kubeadm_LocalEtcd(in *LocalEtcd, out *kubeadm.LocalEtcd, s conversion.Scope) error {
	return autoConvert_v1beta2_LocalEtcd_To_kubeadm_LocalEtcd(in, out, s)
}

func autoConvert_kubeadm_LocalEtcd_To_v1beta2_LocalEtcd(in *kubeadm.LocalEtcd, out *LocalEtcd, s conversion.Scope) error {
	if err := Convert_kubeadm_ImageMeta_To_v1beta2_ImageMeta(&in.ImageMeta, &out.ImageMeta, s); err != nil {
		return err
	}
	out.DataDir = in.DataDir
	out.ExtraArgs = *(*map[string]string)(unsafe.Pointer(&in.ExtraArgs))
	out.ServerCertSANs = *(*[]string)(unsafe.Pointer(&in.ServerCertSANs))
	out.PeerCertSANs = *(*[]string)(unsafe.Pointer(&in.PeerCertSANs))
	return nil
}

// Convert_kubeadm_LocalEtcd_To_v1beta2_LocalEtcd is an autogenerated conversion function.
func Convert_kubeadm_LocalEtcd_To_v1beta2_LocalEtcd(in *kubeadm.LocalEtcd, out *LocalEtcd, s conversion.Scope) error {
	return autoConvert_kubeadm_LocalEtcd_To_v1beta2_LocalEtcd(in, out, s)
}

func autoConvert_v1beta2_Networking_To_kubeadm_Networking(in *Networking, out *kubeadm.Networking, s conversion.Scope) error {
	out.ServiceSubnet = in.ServiceSubnet
	out.PodSubnet = in.PodSubnet
	out.DNSDomain = in.DNSDomain
	return nil
}

// Convert_v1beta2_Networking_To_kubeadm_Networking is an autogenerated conversion function.
func Convert_v1beta2_Networking_To_kubeadm_Networking(in *Networking, out *kubeadm.Networking, s conversion.Scope) error {
	return autoConvert_v1beta2_Networking_To_kubeadm_Networking(in, out, s)
}

func autoConvert_kubeadm_Networking_To_v1beta2_Networking(in *kubeadm.Networking, out *Networking, s conversion.Scope) error {
	out.ServiceSubnet = in.ServiceSubnet
	out.PodSubnet = in.PodSubnet
	out.DNSDomain = in.DNSDomain
	return nil
}

// Convert_kubeadm_Networking_To_v1beta2_Networking is an autogenerated conversion function.
func Convert_kubeadm_Networking_To_v1beta2_Networking(in *kubeadm.Networking, out *Networking, s conversion.Scope) error {
	return autoConvert_kubeadm_Networking_To_v1beta2_Networking(in, out, s)
}

func autoConvert_v1beta2_NodeRegistrationOptions_To_kubeadm_NodeRegistrationOptions(in *NodeRegistrationOptions, out *kubeadm.NodeRegistrationOptions, s conversion.Scope) error {
	out.Name = in.Name
	out.CRISocket = in.CRISocket
	out.Taints = *(*[]corev1.Taint)(unsafe.Pointer(&in.Taints))
	out.KubeletExtraArgs = *(*map[string]string)(unsafe.Pointer(&in.KubeletExtraArgs))
	out.IgnorePreflightErrors = *(*[]string)(unsafe.Pointer(&in.IgnorePreflightErrors))
	return nil
}

// Convert_v1beta2_NodeRegistrationOptions_To_kubeadm_NodeRegistrationOptions is an autogenerated conversion function.
func Convert_v1beta2_NodeRegistrationOptions_To_kubeadm_NodeRegistrationOptions(in *NodeRegistrationOptions, out *kubeadm.NodeRegistrationOptions, s conversion.Scope) error {
	return autoConvert_v1beta2_NodeRegistrationOptions_To_kubeadm_NodeRegistrationOptions(in, out, s)
}

func autoConvert_kubeadm_NodeRegistrationOptions_To_v1beta2_NodeRegistrationOptions(in *kubeadm.NodeRegistrationOptions, out *NodeRegistrationOptions, s conversion.Scope) error {
	out.Name = in.Name
	out.CRISocket = in.CRISocket
	out.Taints = *(*[]corev1.Taint)(unsafe.Pointer(&in.Taints))
	out.KubeletExtraArgs = *(*map[string]string)(unsafe.Pointer(&in.KubeletExtraArgs))
	out.IgnorePreflightErrors = *(*[]string)(unsafe.Pointer(&in.IgnorePreflightErrors))
	// WARNING: in.ImagePullPolicy requires manual conversion: does not exist in peer-type
	return nil
}
