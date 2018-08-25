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

// Code generated by deepcopy-gen. DO NOT EDIT.

package v1alpha1

import (
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KubeProxyConfiguration) DeepCopyInto(out *KubeProxyConfiguration) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	if in.FeatureGates != nil {
		in, out := &in.FeatureGates, &out.FeatureGates
		*out = make(map[string]bool, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	out.ClientConnection = in.ClientConnection
	in.IPTables.DeepCopyInto(&out.IPTables)
	in.IPVS.DeepCopyInto(&out.IPVS)
	if in.OOMScoreAdj != nil {
		in, out := &in.OOMScoreAdj, &out.OOMScoreAdj
		*out = new(int32)
		**out = **in
	}
	out.UDPIdleTimeout = in.UDPIdleTimeout
	in.Conntrack.DeepCopyInto(&out.Conntrack)
	out.ConfigSyncPeriod = in.ConfigSyncPeriod
	if in.NodePortAddresses != nil {
		in, out := &in.NodePortAddresses, &out.NodePortAddresses
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KubeProxyConfiguration.
func (in *KubeProxyConfiguration) DeepCopy() *KubeProxyConfiguration {
	if in == nil {
		return nil
	}
	out := new(KubeProxyConfiguration)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *KubeProxyConfiguration) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KubeProxyConntrackConfiguration) DeepCopyInto(out *KubeProxyConntrackConfiguration) {
	*out = *in
	if in.Max != nil {
		in, out := &in.Max, &out.Max
		*out = new(int32)
		**out = **in
	}
	if in.MaxPerCore != nil {
		in, out := &in.MaxPerCore, &out.MaxPerCore
		*out = new(int32)
		**out = **in
	}
	if in.Min != nil {
		in, out := &in.Min, &out.Min
		*out = new(int32)
		**out = **in
	}
	if in.TCPEstablishedTimeout != nil {
		in, out := &in.TCPEstablishedTimeout, &out.TCPEstablishedTimeout
		*out = new(v1.Duration)
		**out = **in
	}
	if in.TCPCloseWaitTimeout != nil {
		in, out := &in.TCPCloseWaitTimeout, &out.TCPCloseWaitTimeout
		*out = new(v1.Duration)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KubeProxyConntrackConfiguration.
func (in *KubeProxyConntrackConfiguration) DeepCopy() *KubeProxyConntrackConfiguration {
	if in == nil {
		return nil
	}
	out := new(KubeProxyConntrackConfiguration)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KubeProxyIPTablesConfiguration) DeepCopyInto(out *KubeProxyIPTablesConfiguration) {
	*out = *in
	if in.MasqueradeBit != nil {
		in, out := &in.MasqueradeBit, &out.MasqueradeBit
		*out = new(int32)
		**out = **in
	}
	out.SyncPeriod = in.SyncPeriod
	out.MinSyncPeriod = in.MinSyncPeriod
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KubeProxyIPTablesConfiguration.
func (in *KubeProxyIPTablesConfiguration) DeepCopy() *KubeProxyIPTablesConfiguration {
	if in == nil {
		return nil
	}
	out := new(KubeProxyIPTablesConfiguration)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KubeProxyIPVSConfiguration) DeepCopyInto(out *KubeProxyIPVSConfiguration) {
	*out = *in
	out.SyncPeriod = in.SyncPeriod
	out.MinSyncPeriod = in.MinSyncPeriod
	if in.ExcludeCIDRs != nil {
		in, out := &in.ExcludeCIDRs, &out.ExcludeCIDRs
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KubeProxyIPVSConfiguration.
func (in *KubeProxyIPVSConfiguration) DeepCopy() *KubeProxyIPVSConfiguration {
	if in == nil {
		return nil
	}
	out := new(KubeProxyIPVSConfiguration)
	in.DeepCopyInto(out)
	return out
}
