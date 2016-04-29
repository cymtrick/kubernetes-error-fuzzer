// +build !ignore_autogenerated

/*
Copyright 2016 The Kubernetes Authors All rights reserved.

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
	api "k8s.io/kubernetes/pkg/api"
	componentconfig "k8s.io/kubernetes/pkg/apis/componentconfig"
	conversion "k8s.io/kubernetes/pkg/conversion"
)

func init() {
	if err := api.Scheme.AddGeneratedConversionFuncs(
		Convert_v1alpha1_KubeProxyConfiguration_To_componentconfig_KubeProxyConfiguration,
		Convert_componentconfig_KubeProxyConfiguration_To_v1alpha1_KubeProxyConfiguration,
		Convert_v1alpha1_KubeSchedulerConfiguration_To_componentconfig_KubeSchedulerConfiguration,
		Convert_componentconfig_KubeSchedulerConfiguration_To_v1alpha1_KubeSchedulerConfiguration,
		Convert_v1alpha1_LeaderElectionConfiguration_To_componentconfig_LeaderElectionConfiguration,
		Convert_componentconfig_LeaderElectionConfiguration_To_v1alpha1_LeaderElectionConfiguration,
	); err != nil {
		// if one of the conversion functions is malformed, detect it immediately.
		panic(err)
	}
}

func autoConvert_v1alpha1_KubeProxyConfiguration_To_componentconfig_KubeProxyConfiguration(in *KubeProxyConfiguration, out *componentconfig.KubeProxyConfiguration, s conversion.Scope) error {
	SetDefaults_KubeProxyConfiguration(in)
	if err := api.Convert_unversioned_TypeMeta_To_unversioned_TypeMeta(&in.TypeMeta, &out.TypeMeta, s); err != nil {
		return err
	}
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
	// TODO: Inefficient conversion - can we improve it?
	if err := s.Convert(&in.IPTablesSyncPeriod, &out.IPTablesSyncPeriod, 0); err != nil {
		return err
	}
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
	out.Mode = componentconfig.ProxyMode(in.Mode)
	out.PortRange = in.PortRange
	out.ResourceContainer = in.ResourceContainer
	// TODO: Inefficient conversion - can we improve it?
	if err := s.Convert(&in.UDPIdleTimeout, &out.UDPIdleTimeout, 0); err != nil {
		return err
	}
	out.ConntrackMax = in.ConntrackMax
	// TODO: Inefficient conversion - can we improve it?
	if err := s.Convert(&in.ConntrackTCPEstablishedTimeout, &out.ConntrackTCPEstablishedTimeout, 0); err != nil {
		return err
	}
	return nil
}

func Convert_v1alpha1_KubeProxyConfiguration_To_componentconfig_KubeProxyConfiguration(in *KubeProxyConfiguration, out *componentconfig.KubeProxyConfiguration, s conversion.Scope) error {
	return autoConvert_v1alpha1_KubeProxyConfiguration_To_componentconfig_KubeProxyConfiguration(in, out, s)
}

func autoConvert_componentconfig_KubeProxyConfiguration_To_v1alpha1_KubeProxyConfiguration(in *componentconfig.KubeProxyConfiguration, out *KubeProxyConfiguration, s conversion.Scope) error {
	if err := api.Convert_unversioned_TypeMeta_To_unversioned_TypeMeta(&in.TypeMeta, &out.TypeMeta, s); err != nil {
		return err
	}
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
	// TODO: Inefficient conversion - can we improve it?
	if err := s.Convert(&in.IPTablesSyncPeriod, &out.IPTablesSyncPeriod, 0); err != nil {
		return err
	}
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
	out.Mode = ProxyMode(in.Mode)
	out.PortRange = in.PortRange
	out.ResourceContainer = in.ResourceContainer
	// TODO: Inefficient conversion - can we improve it?
	if err := s.Convert(&in.UDPIdleTimeout, &out.UDPIdleTimeout, 0); err != nil {
		return err
	}
	out.ConntrackMax = in.ConntrackMax
	// TODO: Inefficient conversion - can we improve it?
	if err := s.Convert(&in.ConntrackTCPEstablishedTimeout, &out.ConntrackTCPEstablishedTimeout, 0); err != nil {
		return err
	}
	return nil
}

func Convert_componentconfig_KubeProxyConfiguration_To_v1alpha1_KubeProxyConfiguration(in *componentconfig.KubeProxyConfiguration, out *KubeProxyConfiguration, s conversion.Scope) error {
	return autoConvert_componentconfig_KubeProxyConfiguration_To_v1alpha1_KubeProxyConfiguration(in, out, s)
}

func autoConvert_v1alpha1_KubeSchedulerConfiguration_To_componentconfig_KubeSchedulerConfiguration(in *KubeSchedulerConfiguration, out *componentconfig.KubeSchedulerConfiguration, s conversion.Scope) error {
	SetDefaults_KubeSchedulerConfiguration(in)
	if err := api.Convert_unversioned_TypeMeta_To_unversioned_TypeMeta(&in.TypeMeta, &out.TypeMeta, s); err != nil {
		return err
	}
	out.Port = int32(in.Port)
	out.Address = in.Address
	out.AlgorithmProvider = in.AlgorithmProvider
	out.PolicyConfigFile = in.PolicyConfigFile
	if err := api.Convert_Pointer_bool_To_bool(&in.EnableProfiling, &out.EnableProfiling, s); err != nil {
		return err
	}
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
	if err := api.Convert_unversioned_TypeMeta_To_unversioned_TypeMeta(&in.TypeMeta, &out.TypeMeta, s); err != nil {
		return err
	}
	out.Port = int(in.Port)
	out.Address = in.Address
	out.AlgorithmProvider = in.AlgorithmProvider
	out.PolicyConfigFile = in.PolicyConfigFile
	if err := api.Convert_bool_To_Pointer_bool(&in.EnableProfiling, &out.EnableProfiling, s); err != nil {
		return err
	}
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

func autoConvert_v1alpha1_LeaderElectionConfiguration_To_componentconfig_LeaderElectionConfiguration(in *LeaderElectionConfiguration, out *componentconfig.LeaderElectionConfiguration, s conversion.Scope) error {
	SetDefaults_LeaderElectionConfiguration(in)
	if err := api.Convert_Pointer_bool_To_bool(&in.LeaderElect, &out.LeaderElect, s); err != nil {
		return err
	}
	// TODO: Inefficient conversion - can we improve it?
	if err := s.Convert(&in.LeaseDuration, &out.LeaseDuration, 0); err != nil {
		return err
	}
	// TODO: Inefficient conversion - can we improve it?
	if err := s.Convert(&in.RenewDeadline, &out.RenewDeadline, 0); err != nil {
		return err
	}
	// TODO: Inefficient conversion - can we improve it?
	if err := s.Convert(&in.RetryPeriod, &out.RetryPeriod, 0); err != nil {
		return err
	}
	return nil
}

func Convert_v1alpha1_LeaderElectionConfiguration_To_componentconfig_LeaderElectionConfiguration(in *LeaderElectionConfiguration, out *componentconfig.LeaderElectionConfiguration, s conversion.Scope) error {
	return autoConvert_v1alpha1_LeaderElectionConfiguration_To_componentconfig_LeaderElectionConfiguration(in, out, s)
}

func autoConvert_componentconfig_LeaderElectionConfiguration_To_v1alpha1_LeaderElectionConfiguration(in *componentconfig.LeaderElectionConfiguration, out *LeaderElectionConfiguration, s conversion.Scope) error {
	if err := api.Convert_bool_To_Pointer_bool(&in.LeaderElect, &out.LeaderElect, s); err != nil {
		return err
	}
	// TODO: Inefficient conversion - can we improve it?
	if err := s.Convert(&in.LeaseDuration, &out.LeaseDuration, 0); err != nil {
		return err
	}
	// TODO: Inefficient conversion - can we improve it?
	if err := s.Convert(&in.RenewDeadline, &out.RenewDeadline, 0); err != nil {
		return err
	}
	// TODO: Inefficient conversion - can we improve it?
	if err := s.Convert(&in.RetryPeriod, &out.RetryPeriod, 0); err != nil {
		return err
	}
	return nil
}

func Convert_componentconfig_LeaderElectionConfiguration_To_v1alpha1_LeaderElectionConfiguration(in *componentconfig.LeaderElectionConfiguration, out *LeaderElectionConfiguration, s conversion.Scope) error {
	return autoConvert_componentconfig_LeaderElectionConfiguration_To_v1alpha1_LeaderElectionConfiguration(in, out, s)
}
