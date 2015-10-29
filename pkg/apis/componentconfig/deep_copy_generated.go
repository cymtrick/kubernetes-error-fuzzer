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

// DO NOT EDIT. THIS FILE IS AUTO-GENERATED BY $KUBEROOT/hack/update-generated-deep-copies.sh.

package componentconfig

import (
	api "k8s.io/kubernetes/pkg/api"
	unversioned "k8s.io/kubernetes/pkg/api/unversioned"
	conversion "k8s.io/kubernetes/pkg/conversion"
)

func deepCopy_unversioned_TypeMeta(in unversioned.TypeMeta, out *unversioned.TypeMeta, c *conversion.Cloner) error {
	out.Kind = in.Kind
	out.APIVersion = in.APIVersion
	return nil
}

func deepCopy_componentconfig_KubeProxyConfiguration(in KubeProxyConfiguration, out *KubeProxyConfiguration, c *conversion.Cloner) error {
	if err := deepCopy_unversioned_TypeMeta(in.TypeMeta, &out.TypeMeta, c); err != nil {
		return err
	}
	out.BindAddress = in.BindAddress
	out.CleanupIPTables = in.CleanupIPTables
	out.HealthzBindAddress = in.HealthzBindAddress
	out.HealthzPort = in.HealthzPort
	out.HostnameOverride = in.HostnameOverride
	out.IPTablesSyncePeriodSeconds = in.IPTablesSyncePeriodSeconds
	out.KubeAPIBurst = in.KubeAPIBurst
	out.KubeAPIQPS = in.KubeAPIQPS
	out.KubeconfigPath = in.KubeconfigPath
	out.MasqueradeAll = in.MasqueradeAll
	out.Master = in.Master
	if in.OOMScoreAdj != nil {
		out.OOMScoreAdj = new(int)
		*out.OOMScoreAdj = *in.OOMScoreAdj
	} else {
		out.OOMScoreAdj = nil
	}
	out.Mode = in.Mode
	out.PortRange = in.PortRange
	out.ResourceContainer = in.ResourceContainer
	out.UDPTimeoutMilliseconds = in.UDPTimeoutMilliseconds
	return nil
}

func init() {
	err := api.Scheme.AddGeneratedDeepCopyFuncs(
		deepCopy_unversioned_TypeMeta,
		deepCopy_componentconfig_KubeProxyConfiguration,
	)
	if err != nil {
		// if one of the deep copy functions is malformed, detect it immediately.
		panic(err)
	}
}
