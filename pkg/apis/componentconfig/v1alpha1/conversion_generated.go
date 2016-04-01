// +build !ignore_autogenerated

/*
Copyright 2015 The Kubernetes Authors All rights reserved.

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

// DO NOT EDIT. THIS FILE IS AUTO-GENERATED BY $KUBEROOT/hack/update-generated-conversions.sh

package v1alpha1

import (
	reflect "reflect"

	api "k8s.io/kubernetes/pkg/api"
	componentconfig "k8s.io/kubernetes/pkg/apis/componentconfig"
	conversion "k8s.io/kubernetes/pkg/conversion"
)

func autoConvert_componentconfig_KubeProxyConfiguration_To_v1alpha1_KubeProxyConfiguration(in *componentconfig.KubeProxyConfiguration, out *KubeProxyConfiguration, s conversion.Scope) error {
	if defaulting, found := s.DefaultingInterface(reflect.TypeOf(*in)); found {
		defaulting.(func(*componentconfig.KubeProxyConfiguration))(in)
	}
	if err := api.Convert_unversioned_TypeMeta_To_unversioned_TypeMeta(&in.TypeMeta, &out.TypeMeta, s); err != nil {
		return err
	}
	out.BindAddress = in.BindAddress
	out.HealthzBindAddress = in.HealthzBindAddress
	out.HealthzPort = int32(in.HealthzPort)
	out.HostnameOverride = in.HostnameOverride
	if in.IPTablesMasqueradeBit != nil {
		out.IPTablesMasqueradeBit = new(int32)
		*out.IPTablesMasqueradeBit = int32(*in.IPTablesMasqueradeBit)
	} else {
		out.IPTablesMasqueradeBit = nil
	}
	if err := s.Convert(&in.IPTablesSyncPeriod, &out.IPTablesSyncPeriod, 0); err != nil {
		return err
	}
	out.KubeconfigPath = in.KubeconfigPath
	out.MasqueradeAll = in.MasqueradeAll
	out.Master = in.Master
	if in.OOMScoreAdj != nil {
		out.OOMScoreAdj = new(int32)
		*out.OOMScoreAdj = int32(*in.OOMScoreAdj)
	} else {
		out.OOMScoreAdj = nil
	}
	out.Mode = ProxyMode(in.Mode)
	out.PortRange = in.PortRange
	out.ResourceContainer = in.ResourceContainer
	if err := s.Convert(&in.UDPIdleTimeout, &out.UDPIdleTimeout, 0); err != nil {
		return err
	}
	out.ConntrackMax = int32(in.ConntrackMax)
	if err := s.Convert(&in.ConntrackTCPEstablishedTimeout, &out.ConntrackTCPEstablishedTimeout, 0); err != nil {
		return err
	}
	return nil
}

func Convert_componentconfig_KubeProxyConfiguration_To_v1alpha1_KubeProxyConfiguration(in *componentconfig.KubeProxyConfiguration, out *KubeProxyConfiguration, s conversion.Scope) error {
	return autoConvert_componentconfig_KubeProxyConfiguration_To_v1alpha1_KubeProxyConfiguration(in, out, s)
}

func autoConvert_componentconfig_KubeSchedulerConfiguration_To_v1alpha1_KubeSchedulerConfiguration(in *componentconfig.KubeSchedulerConfiguration, out *KubeSchedulerConfiguration, s conversion.Scope) error {
	if defaulting, found := s.DefaultingInterface(reflect.TypeOf(*in)); found {
		defaulting.(func(*componentconfig.KubeSchedulerConfiguration))(in)
	}
	if err := api.Convert_unversioned_TypeMeta_To_unversioned_TypeMeta(&in.TypeMeta, &out.TypeMeta, s); err != nil {
		return err
	}
	out.Port = in.Port
	out.Address = in.Address
	out.AlgorithmProvider = in.AlgorithmProvider
	out.PolicyConfigFile = in.PolicyConfigFile
	if err := api.Convert_bool_To_Pointer_bool(&in.EnableProfiling, &out.EnableProfiling, s); err != nil {
		return err
	}
	out.KubeAPIQPS = in.KubeAPIQPS
	out.KubeAPIBurst = in.KubeAPIBurst
	out.SchedulerName = in.SchedulerName
	if err := Convert_componentconfig_LeaderElectionConfiguration_To_v1alpha1_LeaderElectionConfiguration(&in.LeaderElection, &out.LeaderElection, s); err != nil {
		return err
	}
	return nil
}

func Convert_componentconfig_KubeSchedulerConfiguration_To_v1alpha1_KubeSchedulerConfiguration(in *componentconfig.KubeSchedulerConfiguration, out *KubeSchedulerConfiguration, s conversion.Scope) error {
	return autoConvert_componentconfig_KubeSchedulerConfiguration_To_v1alpha1_KubeSchedulerConfiguration(in, out, s)
}

func autoConvert_componentconfig_LeaderElectionConfiguration_To_v1alpha1_LeaderElectionConfiguration(in *componentconfig.LeaderElectionConfiguration, out *LeaderElectionConfiguration, s conversion.Scope) error {
	if defaulting, found := s.DefaultingInterface(reflect.TypeOf(*in)); found {
		defaulting.(func(*componentconfig.LeaderElectionConfiguration))(in)
	}
	if err := api.Convert_bool_To_Pointer_bool(&in.LeaderElect, &out.LeaderElect, s); err != nil {
		return err
	}
	if err := s.Convert(&in.LeaseDuration, &out.LeaseDuration, 0); err != nil {
		return err
	}
	if err := s.Convert(&in.RenewDeadline, &out.RenewDeadline, 0); err != nil {
		return err
	}
	if err := s.Convert(&in.RetryPeriod, &out.RetryPeriod, 0); err != nil {
		return err
	}
	return nil
}

func Convert_componentconfig_LeaderElectionConfiguration_To_v1alpha1_LeaderElectionConfiguration(in *componentconfig.LeaderElectionConfiguration, out *LeaderElectionConfiguration, s conversion.Scope) error {
	return autoConvert_componentconfig_LeaderElectionConfiguration_To_v1alpha1_LeaderElectionConfiguration(in, out, s)
}

func autoConvert_v1alpha1_KubeProxyConfiguration_To_componentconfig_KubeProxyConfiguration(in *KubeProxyConfiguration, out *componentconfig.KubeProxyConfiguration, s conversion.Scope) error {
	if defaulting, found := s.DefaultingInterface(reflect.TypeOf(*in)); found {
		defaulting.(func(*KubeProxyConfiguration))(in)
	}
	if err := api.Convert_unversioned_TypeMeta_To_unversioned_TypeMeta(&in.TypeMeta, &out.TypeMeta, s); err != nil {
		return err
	}
	out.BindAddress = in.BindAddress
	out.HealthzBindAddress = in.HealthzBindAddress
	out.HealthzPort = int(in.HealthzPort)
	out.HostnameOverride = in.HostnameOverride
	if in.IPTablesMasqueradeBit != nil {
		out.IPTablesMasqueradeBit = new(int)
		*out.IPTablesMasqueradeBit = int(*in.IPTablesMasqueradeBit)
	} else {
		out.IPTablesMasqueradeBit = nil
	}
	if err := s.Convert(&in.IPTablesSyncPeriod, &out.IPTablesSyncPeriod, 0); err != nil {
		return err
	}
	out.KubeconfigPath = in.KubeconfigPath
	out.MasqueradeAll = in.MasqueradeAll
	out.Master = in.Master
	if in.OOMScoreAdj != nil {
		out.OOMScoreAdj = new(int)
		*out.OOMScoreAdj = int(*in.OOMScoreAdj)
	} else {
		out.OOMScoreAdj = nil
	}
	out.Mode = componentconfig.ProxyMode(in.Mode)
	out.PortRange = in.PortRange
	out.ResourceContainer = in.ResourceContainer
	if err := s.Convert(&in.UDPIdleTimeout, &out.UDPIdleTimeout, 0); err != nil {
		return err
	}
	out.ConntrackMax = int(in.ConntrackMax)
	if err := s.Convert(&in.ConntrackTCPEstablishedTimeout, &out.ConntrackTCPEstablishedTimeout, 0); err != nil {
		return err
	}
	return nil
}

func Convert_v1alpha1_KubeProxyConfiguration_To_componentconfig_KubeProxyConfiguration(in *KubeProxyConfiguration, out *componentconfig.KubeProxyConfiguration, s conversion.Scope) error {
	return autoConvert_v1alpha1_KubeProxyConfiguration_To_componentconfig_KubeProxyConfiguration(in, out, s)
}

func autoConvert_v1alpha1_KubeSchedulerConfiguration_To_componentconfig_KubeSchedulerConfiguration(in *KubeSchedulerConfiguration, out *componentconfig.KubeSchedulerConfiguration, s conversion.Scope) error {
	if defaulting, found := s.DefaultingInterface(reflect.TypeOf(*in)); found {
		defaulting.(func(*KubeSchedulerConfiguration))(in)
	}
	if err := api.Convert_unversioned_TypeMeta_To_unversioned_TypeMeta(&in.TypeMeta, &out.TypeMeta, s); err != nil {
		return err
	}
	out.Port = in.Port
	out.Address = in.Address
	out.AlgorithmProvider = in.AlgorithmProvider
	out.PolicyConfigFile = in.PolicyConfigFile
	if err := api.Convert_Pointer_bool_To_bool(&in.EnableProfiling, &out.EnableProfiling, s); err != nil {
		return err
	}
	out.KubeAPIQPS = in.KubeAPIQPS
	out.KubeAPIBurst = in.KubeAPIBurst
	out.SchedulerName = in.SchedulerName
	if err := Convert_v1alpha1_LeaderElectionConfiguration_To_componentconfig_LeaderElectionConfiguration(&in.LeaderElection, &out.LeaderElection, s); err != nil {
		return err
	}
	return nil
}

func Convert_v1alpha1_KubeSchedulerConfiguration_To_componentconfig_KubeSchedulerConfiguration(in *KubeSchedulerConfiguration, out *componentconfig.KubeSchedulerConfiguration, s conversion.Scope) error {
	return autoConvert_v1alpha1_KubeSchedulerConfiguration_To_componentconfig_KubeSchedulerConfiguration(in, out, s)
}

func autoConvert_v1alpha1_LeaderElectionConfiguration_To_componentconfig_LeaderElectionConfiguration(in *LeaderElectionConfiguration, out *componentconfig.LeaderElectionConfiguration, s conversion.Scope) error {
	if defaulting, found := s.DefaultingInterface(reflect.TypeOf(*in)); found {
		defaulting.(func(*LeaderElectionConfiguration))(in)
	}
	if err := api.Convert_Pointer_bool_To_bool(&in.LeaderElect, &out.LeaderElect, s); err != nil {
		return err
	}
	if err := s.Convert(&in.LeaseDuration, &out.LeaseDuration, 0); err != nil {
		return err
	}
	if err := s.Convert(&in.RenewDeadline, &out.RenewDeadline, 0); err != nil {
		return err
	}
	if err := s.Convert(&in.RetryPeriod, &out.RetryPeriod, 0); err != nil {
		return err
	}
	return nil
}

func Convert_v1alpha1_LeaderElectionConfiguration_To_componentconfig_LeaderElectionConfiguration(in *LeaderElectionConfiguration, out *componentconfig.LeaderElectionConfiguration, s conversion.Scope) error {
	return autoConvert_v1alpha1_LeaderElectionConfiguration_To_componentconfig_LeaderElectionConfiguration(in, out, s)
}

func init() {
	err := api.Scheme.AddGeneratedConversionFuncs(
		autoConvert_componentconfig_KubeProxyConfiguration_To_v1alpha1_KubeProxyConfiguration,
		autoConvert_componentconfig_KubeSchedulerConfiguration_To_v1alpha1_KubeSchedulerConfiguration,
		autoConvert_componentconfig_LeaderElectionConfiguration_To_v1alpha1_LeaderElectionConfiguration,
		autoConvert_v1alpha1_KubeProxyConfiguration_To_componentconfig_KubeProxyConfiguration,
		autoConvert_v1alpha1_KubeSchedulerConfiguration_To_componentconfig_KubeSchedulerConfiguration,
		autoConvert_v1alpha1_LeaderElectionConfiguration_To_componentconfig_LeaderElectionConfiguration,
	)
	if err != nil {
		// If one of the conversion functions is malformed, detect it immediately.
		panic(err)
	}
}
