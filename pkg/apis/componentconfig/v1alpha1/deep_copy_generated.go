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

// This file was autogenerated by deepcopy-gen. Do not edit it manually!

package v1alpha1

import (
	api "k8s.io/kubernetes/pkg/api"
	unversioned "k8s.io/kubernetes/pkg/api/unversioned"
	conversion "k8s.io/kubernetes/pkg/conversion"
)

func init() {
	if err := api.Scheme.AddGeneratedDeepCopyFuncs(
		DeepCopy_v1alpha1_KubeProxyConfiguration,
		DeepCopy_v1alpha1_KubeSchedulerConfiguration,
		DeepCopy_v1alpha1_LeaderElectionConfiguration,
	); err != nil {
		// if one of the deep copy functions is malformed, detect it immediately.
		panic(err)
	}
}

func DeepCopy_v1alpha1_KubeProxyConfiguration(in KubeProxyConfiguration, out *KubeProxyConfiguration, c *conversion.Cloner) error {
	if err := unversioned.DeepCopy_unversioned_TypeMeta(in.TypeMeta, &out.TypeMeta, c); err != nil {
		return err
	}
	out.BindAddress = in.BindAddress
	out.ClusterCIDR = in.ClusterCIDR
	out.HealthzBindAddress = in.HealthzBindAddress
	out.HealthzPort = in.HealthzPort
	out.HostnameOverride = in.HostnameOverride
	if in.IPTablesMasqueradeBit != nil {
		in, out := in.IPTablesMasqueradeBit, &out.IPTablesMasqueradeBit
		*out = new(int32)
		**out = *in
	} else {
		out.IPTablesMasqueradeBit = nil
	}
	if err := unversioned.DeepCopy_unversioned_Duration(in.IPTablesSyncPeriod, &out.IPTablesSyncPeriod, c); err != nil {
		return err
	}
	out.KubeconfigPath = in.KubeconfigPath
	out.MasqueradeAll = in.MasqueradeAll
	out.Master = in.Master
	if in.OOMScoreAdj != nil {
		in, out := in.OOMScoreAdj, &out.OOMScoreAdj
		*out = new(int32)
		**out = *in
	} else {
		out.OOMScoreAdj = nil
	}
	out.Mode = in.Mode
	out.PortRange = in.PortRange
	out.ResourceContainer = in.ResourceContainer
	if err := unversioned.DeepCopy_unversioned_Duration(in.UDPIdleTimeout, &out.UDPIdleTimeout, c); err != nil {
		return err
	}
	out.ConntrackMax = in.ConntrackMax
	if err := unversioned.DeepCopy_unversioned_Duration(in.ConntrackTCPEstablishedTimeout, &out.ConntrackTCPEstablishedTimeout, c); err != nil {
		return err
	}
	return nil
}

func DeepCopy_v1alpha1_KubeSchedulerConfiguration(in KubeSchedulerConfiguration, out *KubeSchedulerConfiguration, c *conversion.Cloner) error {
	if err := unversioned.DeepCopy_unversioned_TypeMeta(in.TypeMeta, &out.TypeMeta, c); err != nil {
		return err
	}
	out.Port = in.Port
	out.Address = in.Address
	out.AlgorithmProvider = in.AlgorithmProvider
	out.PolicyConfigFile = in.PolicyConfigFile
	if in.EnableProfiling != nil {
		in, out := in.EnableProfiling, &out.EnableProfiling
		*out = new(bool)
		**out = *in
	} else {
		out.EnableProfiling = nil
	}
	out.ContentType = in.ContentType
	out.KubeAPIQPS = in.KubeAPIQPS
	out.KubeAPIBurst = in.KubeAPIBurst
	out.SchedulerName = in.SchedulerName
	out.HardPodAffinitySymmetricWeight = in.HardPodAffinitySymmetricWeight
	out.FailureDomains = in.FailureDomains
	if err := DeepCopy_v1alpha1_LeaderElectionConfiguration(in.LeaderElection, &out.LeaderElection, c); err != nil {
		return err
	}
	return nil
}

func DeepCopy_v1alpha1_LeaderElectionConfiguration(in LeaderElectionConfiguration, out *LeaderElectionConfiguration, c *conversion.Cloner) error {
	if in.LeaderElect != nil {
		in, out := in.LeaderElect, &out.LeaderElect
		*out = new(bool)
		**out = *in
	} else {
		out.LeaderElect = nil
	}
	if err := unversioned.DeepCopy_unversioned_Duration(in.LeaseDuration, &out.LeaseDuration, c); err != nil {
		return err
	}
	if err := unversioned.DeepCopy_unversioned_Duration(in.RenewDeadline, &out.RenewDeadline, c); err != nil {
		return err
	}
	if err := unversioned.DeepCopy_unversioned_Duration(in.RetryPeriod, &out.RetryPeriod, c); err != nil {
		return err
	}
	return nil
}
