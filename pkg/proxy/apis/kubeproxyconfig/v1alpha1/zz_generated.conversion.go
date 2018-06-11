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

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	conversion "k8s.io/apimachinery/pkg/conversion"
	runtime "k8s.io/apimachinery/pkg/runtime"
	kubeproxyconfig "k8s.io/kubernetes/pkg/proxy/apis/kubeproxyconfig"
)

func init() {
	localSchemeBuilder.Register(RegisterConversions)
}

// RegisterConversions adds conversion functions to the given scheme.
// Public to allow building arbitrary schemes.
func RegisterConversions(s *runtime.Scheme) error {
	if err := s.AddGeneratedConversionFunc((*ClientConnectionConfiguration)(nil), (*kubeproxyconfig.ClientConnectionConfiguration)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha1_ClientConnectionConfiguration_To_kubeproxyconfig_ClientConnectionConfiguration(a.(*ClientConnectionConfiguration), b.(*kubeproxyconfig.ClientConnectionConfiguration), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*kubeproxyconfig.ClientConnectionConfiguration)(nil), (*ClientConnectionConfiguration)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_kubeproxyconfig_ClientConnectionConfiguration_To_v1alpha1_ClientConnectionConfiguration(a.(*kubeproxyconfig.ClientConnectionConfiguration), b.(*ClientConnectionConfiguration), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*KubeProxyConfiguration)(nil), (*kubeproxyconfig.KubeProxyConfiguration)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha1_KubeProxyConfiguration_To_kubeproxyconfig_KubeProxyConfiguration(a.(*KubeProxyConfiguration), b.(*kubeproxyconfig.KubeProxyConfiguration), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*kubeproxyconfig.KubeProxyConfiguration)(nil), (*KubeProxyConfiguration)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_kubeproxyconfig_KubeProxyConfiguration_To_v1alpha1_KubeProxyConfiguration(a.(*kubeproxyconfig.KubeProxyConfiguration), b.(*KubeProxyConfiguration), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*KubeProxyConntrackConfiguration)(nil), (*kubeproxyconfig.KubeProxyConntrackConfiguration)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha1_KubeProxyConntrackConfiguration_To_kubeproxyconfig_KubeProxyConntrackConfiguration(a.(*KubeProxyConntrackConfiguration), b.(*kubeproxyconfig.KubeProxyConntrackConfiguration), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*kubeproxyconfig.KubeProxyConntrackConfiguration)(nil), (*KubeProxyConntrackConfiguration)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_kubeproxyconfig_KubeProxyConntrackConfiguration_To_v1alpha1_KubeProxyConntrackConfiguration(a.(*kubeproxyconfig.KubeProxyConntrackConfiguration), b.(*KubeProxyConntrackConfiguration), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*KubeProxyIPTablesConfiguration)(nil), (*kubeproxyconfig.KubeProxyIPTablesConfiguration)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha1_KubeProxyIPTablesConfiguration_To_kubeproxyconfig_KubeProxyIPTablesConfiguration(a.(*KubeProxyIPTablesConfiguration), b.(*kubeproxyconfig.KubeProxyIPTablesConfiguration), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*kubeproxyconfig.KubeProxyIPTablesConfiguration)(nil), (*KubeProxyIPTablesConfiguration)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_kubeproxyconfig_KubeProxyIPTablesConfiguration_To_v1alpha1_KubeProxyIPTablesConfiguration(a.(*kubeproxyconfig.KubeProxyIPTablesConfiguration), b.(*KubeProxyIPTablesConfiguration), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*KubeProxyIPVSConfiguration)(nil), (*kubeproxyconfig.KubeProxyIPVSConfiguration)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha1_KubeProxyIPVSConfiguration_To_kubeproxyconfig_KubeProxyIPVSConfiguration(a.(*KubeProxyIPVSConfiguration), b.(*kubeproxyconfig.KubeProxyIPVSConfiguration), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*kubeproxyconfig.KubeProxyIPVSConfiguration)(nil), (*KubeProxyIPVSConfiguration)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_kubeproxyconfig_KubeProxyIPVSConfiguration_To_v1alpha1_KubeProxyIPVSConfiguration(a.(*kubeproxyconfig.KubeProxyIPVSConfiguration), b.(*KubeProxyIPVSConfiguration), scope)
	}); err != nil {
		return err
	}
	return nil
}

func autoConvert_v1alpha1_ClientConnectionConfiguration_To_kubeproxyconfig_ClientConnectionConfiguration(in *ClientConnectionConfiguration, out *kubeproxyconfig.ClientConnectionConfiguration, s conversion.Scope) error {
	out.KubeConfigFile = in.KubeConfigFile
	out.AcceptContentTypes = in.AcceptContentTypes
	out.ContentType = in.ContentType
	out.QPS = in.QPS
	out.Burst = int32(in.Burst)
	return nil
}

// Convert_v1alpha1_ClientConnectionConfiguration_To_kubeproxyconfig_ClientConnectionConfiguration is an autogenerated conversion function.
func Convert_v1alpha1_ClientConnectionConfiguration_To_kubeproxyconfig_ClientConnectionConfiguration(in *ClientConnectionConfiguration, out *kubeproxyconfig.ClientConnectionConfiguration, s conversion.Scope) error {
	return autoConvert_v1alpha1_ClientConnectionConfiguration_To_kubeproxyconfig_ClientConnectionConfiguration(in, out, s)
}

func autoConvert_kubeproxyconfig_ClientConnectionConfiguration_To_v1alpha1_ClientConnectionConfiguration(in *kubeproxyconfig.ClientConnectionConfiguration, out *ClientConnectionConfiguration, s conversion.Scope) error {
	out.KubeConfigFile = in.KubeConfigFile
	out.AcceptContentTypes = in.AcceptContentTypes
	out.ContentType = in.ContentType
	out.QPS = in.QPS
	out.Burst = int(in.Burst)
	return nil
}

// Convert_kubeproxyconfig_ClientConnectionConfiguration_To_v1alpha1_ClientConnectionConfiguration is an autogenerated conversion function.
func Convert_kubeproxyconfig_ClientConnectionConfiguration_To_v1alpha1_ClientConnectionConfiguration(in *kubeproxyconfig.ClientConnectionConfiguration, out *ClientConnectionConfiguration, s conversion.Scope) error {
	return autoConvert_kubeproxyconfig_ClientConnectionConfiguration_To_v1alpha1_ClientConnectionConfiguration(in, out, s)
}

func autoConvert_v1alpha1_KubeProxyConfiguration_To_kubeproxyconfig_KubeProxyConfiguration(in *KubeProxyConfiguration, out *kubeproxyconfig.KubeProxyConfiguration, s conversion.Scope) error {
	out.FeatureGates = *(*map[string]bool)(unsafe.Pointer(&in.FeatureGates))
	out.BindAddress = in.BindAddress
	out.HealthzBindAddress = in.HealthzBindAddress
	out.MetricsBindAddress = in.MetricsBindAddress
	out.EnableProfiling = in.EnableProfiling
	out.ClusterCIDR = in.ClusterCIDR
	out.HostnameOverride = in.HostnameOverride
	if err := Convert_v1alpha1_ClientConnectionConfiguration_To_kubeproxyconfig_ClientConnectionConfiguration(&in.ClientConnection, &out.ClientConnection, s); err != nil {
		return err
	}
	if err := Convert_v1alpha1_KubeProxyIPTablesConfiguration_To_kubeproxyconfig_KubeProxyIPTablesConfiguration(&in.IPTables, &out.IPTables, s); err != nil {
		return err
	}
	if err := Convert_v1alpha1_KubeProxyIPVSConfiguration_To_kubeproxyconfig_KubeProxyIPVSConfiguration(&in.IPVS, &out.IPVS, s); err != nil {
		return err
	}
	out.OOMScoreAdj = (*int32)(unsafe.Pointer(in.OOMScoreAdj))
	out.Mode = kubeproxyconfig.ProxyMode(in.Mode)
	out.PortRange = in.PortRange
	out.ResourceContainer = in.ResourceContainer
	out.UDPIdleTimeout = in.UDPIdleTimeout
	if err := Convert_v1alpha1_KubeProxyConntrackConfiguration_To_kubeproxyconfig_KubeProxyConntrackConfiguration(&in.Conntrack, &out.Conntrack, s); err != nil {
		return err
	}
	out.ConfigSyncPeriod = in.ConfigSyncPeriod
	out.NodePortAddresses = *(*[]string)(unsafe.Pointer(&in.NodePortAddresses))
	out.SCTPUserSpaceNode = in.SCTPUserSpaceNode
	return nil
}

// Convert_v1alpha1_KubeProxyConfiguration_To_kubeproxyconfig_KubeProxyConfiguration is an autogenerated conversion function.
func Convert_v1alpha1_KubeProxyConfiguration_To_kubeproxyconfig_KubeProxyConfiguration(in *KubeProxyConfiguration, out *kubeproxyconfig.KubeProxyConfiguration, s conversion.Scope) error {
	return autoConvert_v1alpha1_KubeProxyConfiguration_To_kubeproxyconfig_KubeProxyConfiguration(in, out, s)
}

func autoConvert_kubeproxyconfig_KubeProxyConfiguration_To_v1alpha1_KubeProxyConfiguration(in *kubeproxyconfig.KubeProxyConfiguration, out *KubeProxyConfiguration, s conversion.Scope) error {
	out.FeatureGates = *(*map[string]bool)(unsafe.Pointer(&in.FeatureGates))
	out.BindAddress = in.BindAddress
	out.HealthzBindAddress = in.HealthzBindAddress
	out.MetricsBindAddress = in.MetricsBindAddress
	out.EnableProfiling = in.EnableProfiling
	out.ClusterCIDR = in.ClusterCIDR
	out.HostnameOverride = in.HostnameOverride
	if err := Convert_kubeproxyconfig_ClientConnectionConfiguration_To_v1alpha1_ClientConnectionConfiguration(&in.ClientConnection, &out.ClientConnection, s); err != nil {
		return err
	}
	if err := Convert_kubeproxyconfig_KubeProxyIPTablesConfiguration_To_v1alpha1_KubeProxyIPTablesConfiguration(&in.IPTables, &out.IPTables, s); err != nil {
		return err
	}
	if err := Convert_kubeproxyconfig_KubeProxyIPVSConfiguration_To_v1alpha1_KubeProxyIPVSConfiguration(&in.IPVS, &out.IPVS, s); err != nil {
		return err
	}
	out.OOMScoreAdj = (*int32)(unsafe.Pointer(in.OOMScoreAdj))
	out.Mode = ProxyMode(in.Mode)
	out.PortRange = in.PortRange
	out.ResourceContainer = in.ResourceContainer
	out.UDPIdleTimeout = in.UDPIdleTimeout
	if err := Convert_kubeproxyconfig_KubeProxyConntrackConfiguration_To_v1alpha1_KubeProxyConntrackConfiguration(&in.Conntrack, &out.Conntrack, s); err != nil {
		return err
	}
	out.ConfigSyncPeriod = in.ConfigSyncPeriod
	out.NodePortAddresses = *(*[]string)(unsafe.Pointer(&in.NodePortAddresses))
	out.SCTPUserSpaceNode = in.SCTPUserSpaceNode
	return nil
}

// Convert_kubeproxyconfig_KubeProxyConfiguration_To_v1alpha1_KubeProxyConfiguration is an autogenerated conversion function.
func Convert_kubeproxyconfig_KubeProxyConfiguration_To_v1alpha1_KubeProxyConfiguration(in *kubeproxyconfig.KubeProxyConfiguration, out *KubeProxyConfiguration, s conversion.Scope) error {
	return autoConvert_kubeproxyconfig_KubeProxyConfiguration_To_v1alpha1_KubeProxyConfiguration(in, out, s)
}

func autoConvert_v1alpha1_KubeProxyConntrackConfiguration_To_kubeproxyconfig_KubeProxyConntrackConfiguration(in *KubeProxyConntrackConfiguration, out *kubeproxyconfig.KubeProxyConntrackConfiguration, s conversion.Scope) error {
	out.Max = (*int32)(unsafe.Pointer(in.Max))
	out.MaxPerCore = (*int32)(unsafe.Pointer(in.MaxPerCore))
	out.Min = (*int32)(unsafe.Pointer(in.Min))
	out.TCPEstablishedTimeout = (*v1.Duration)(unsafe.Pointer(in.TCPEstablishedTimeout))
	out.TCPCloseWaitTimeout = (*v1.Duration)(unsafe.Pointer(in.TCPCloseWaitTimeout))
	return nil
}

// Convert_v1alpha1_KubeProxyConntrackConfiguration_To_kubeproxyconfig_KubeProxyConntrackConfiguration is an autogenerated conversion function.
func Convert_v1alpha1_KubeProxyConntrackConfiguration_To_kubeproxyconfig_KubeProxyConntrackConfiguration(in *KubeProxyConntrackConfiguration, out *kubeproxyconfig.KubeProxyConntrackConfiguration, s conversion.Scope) error {
	return autoConvert_v1alpha1_KubeProxyConntrackConfiguration_To_kubeproxyconfig_KubeProxyConntrackConfiguration(in, out, s)
}

func autoConvert_kubeproxyconfig_KubeProxyConntrackConfiguration_To_v1alpha1_KubeProxyConntrackConfiguration(in *kubeproxyconfig.KubeProxyConntrackConfiguration, out *KubeProxyConntrackConfiguration, s conversion.Scope) error {
	out.Max = (*int32)(unsafe.Pointer(in.Max))
	out.MaxPerCore = (*int32)(unsafe.Pointer(in.MaxPerCore))
	out.Min = (*int32)(unsafe.Pointer(in.Min))
	out.TCPEstablishedTimeout = (*v1.Duration)(unsafe.Pointer(in.TCPEstablishedTimeout))
	out.TCPCloseWaitTimeout = (*v1.Duration)(unsafe.Pointer(in.TCPCloseWaitTimeout))
	return nil
}

// Convert_kubeproxyconfig_KubeProxyConntrackConfiguration_To_v1alpha1_KubeProxyConntrackConfiguration is an autogenerated conversion function.
func Convert_kubeproxyconfig_KubeProxyConntrackConfiguration_To_v1alpha1_KubeProxyConntrackConfiguration(in *kubeproxyconfig.KubeProxyConntrackConfiguration, out *KubeProxyConntrackConfiguration, s conversion.Scope) error {
	return autoConvert_kubeproxyconfig_KubeProxyConntrackConfiguration_To_v1alpha1_KubeProxyConntrackConfiguration(in, out, s)
}

func autoConvert_v1alpha1_KubeProxyIPTablesConfiguration_To_kubeproxyconfig_KubeProxyIPTablesConfiguration(in *KubeProxyIPTablesConfiguration, out *kubeproxyconfig.KubeProxyIPTablesConfiguration, s conversion.Scope) error {
	out.MasqueradeBit = (*int32)(unsafe.Pointer(in.MasqueradeBit))
	out.MasqueradeAll = in.MasqueradeAll
	out.SyncPeriod = in.SyncPeriod
	out.MinSyncPeriod = in.MinSyncPeriod
	return nil
}

// Convert_v1alpha1_KubeProxyIPTablesConfiguration_To_kubeproxyconfig_KubeProxyIPTablesConfiguration is an autogenerated conversion function.
func Convert_v1alpha1_KubeProxyIPTablesConfiguration_To_kubeproxyconfig_KubeProxyIPTablesConfiguration(in *KubeProxyIPTablesConfiguration, out *kubeproxyconfig.KubeProxyIPTablesConfiguration, s conversion.Scope) error {
	return autoConvert_v1alpha1_KubeProxyIPTablesConfiguration_To_kubeproxyconfig_KubeProxyIPTablesConfiguration(in, out, s)
}

func autoConvert_kubeproxyconfig_KubeProxyIPTablesConfiguration_To_v1alpha1_KubeProxyIPTablesConfiguration(in *kubeproxyconfig.KubeProxyIPTablesConfiguration, out *KubeProxyIPTablesConfiguration, s conversion.Scope) error {
	out.MasqueradeBit = (*int32)(unsafe.Pointer(in.MasqueradeBit))
	out.MasqueradeAll = in.MasqueradeAll
	out.SyncPeriod = in.SyncPeriod
	out.MinSyncPeriod = in.MinSyncPeriod
	return nil
}

// Convert_kubeproxyconfig_KubeProxyIPTablesConfiguration_To_v1alpha1_KubeProxyIPTablesConfiguration is an autogenerated conversion function.
func Convert_kubeproxyconfig_KubeProxyIPTablesConfiguration_To_v1alpha1_KubeProxyIPTablesConfiguration(in *kubeproxyconfig.KubeProxyIPTablesConfiguration, out *KubeProxyIPTablesConfiguration, s conversion.Scope) error {
	return autoConvert_kubeproxyconfig_KubeProxyIPTablesConfiguration_To_v1alpha1_KubeProxyIPTablesConfiguration(in, out, s)
}

func autoConvert_v1alpha1_KubeProxyIPVSConfiguration_To_kubeproxyconfig_KubeProxyIPVSConfiguration(in *KubeProxyIPVSConfiguration, out *kubeproxyconfig.KubeProxyIPVSConfiguration, s conversion.Scope) error {
	out.SyncPeriod = in.SyncPeriod
	out.MinSyncPeriod = in.MinSyncPeriod
	out.Scheduler = in.Scheduler
	out.ExcludeCIDRs = *(*[]string)(unsafe.Pointer(&in.ExcludeCIDRs))
	return nil
}

// Convert_v1alpha1_KubeProxyIPVSConfiguration_To_kubeproxyconfig_KubeProxyIPVSConfiguration is an autogenerated conversion function.
func Convert_v1alpha1_KubeProxyIPVSConfiguration_To_kubeproxyconfig_KubeProxyIPVSConfiguration(in *KubeProxyIPVSConfiguration, out *kubeproxyconfig.KubeProxyIPVSConfiguration, s conversion.Scope) error {
	return autoConvert_v1alpha1_KubeProxyIPVSConfiguration_To_kubeproxyconfig_KubeProxyIPVSConfiguration(in, out, s)
}

func autoConvert_kubeproxyconfig_KubeProxyIPVSConfiguration_To_v1alpha1_KubeProxyIPVSConfiguration(in *kubeproxyconfig.KubeProxyIPVSConfiguration, out *KubeProxyIPVSConfiguration, s conversion.Scope) error {
	out.SyncPeriod = in.SyncPeriod
	out.MinSyncPeriod = in.MinSyncPeriod
	out.Scheduler = in.Scheduler
	out.ExcludeCIDRs = *(*[]string)(unsafe.Pointer(&in.ExcludeCIDRs))
	return nil
}

// Convert_kubeproxyconfig_KubeProxyIPVSConfiguration_To_v1alpha1_KubeProxyIPVSConfiguration is an autogenerated conversion function.
func Convert_kubeproxyconfig_KubeProxyIPVSConfiguration_To_v1alpha1_KubeProxyIPVSConfiguration(in *kubeproxyconfig.KubeProxyIPVSConfiguration, out *KubeProxyIPVSConfiguration, s conversion.Scope) error {
	return autoConvert_kubeproxyconfig_KubeProxyIPVSConfiguration_To_v1alpha1_KubeProxyIPVSConfiguration(in, out, s)
}
