// +build !ignore_autogenerated

/*
Copyright 2017 The Kubernetes Authors.

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

package kubeletconfig

import (
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	conversion "k8s.io/apimachinery/pkg/conversion"
	runtime "k8s.io/apimachinery/pkg/runtime"
	api "k8s.io/kubernetes/pkg/api"
	reflect "reflect"
)

func init() {
	SchemeBuilder.Register(RegisterDeepCopies)
}

// RegisterDeepCopies adds deep-copy functions to the given scheme. Public
// to allow building arbitrary schemes.
//
// Deprecated: deepcopy registration will go away when static deepcopy is fully implemented.
func RegisterDeepCopies(scheme *runtime.Scheme) error {
	return scheme.AddGeneratedDeepCopyFuncs(
		conversion.GeneratedDeepCopyFunc{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*KubeletAnonymousAuthentication).DeepCopyInto(out.(*KubeletAnonymousAuthentication))
			return nil
		}, InType: reflect.TypeOf(&KubeletAnonymousAuthentication{})},
		conversion.GeneratedDeepCopyFunc{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*KubeletAuthentication).DeepCopyInto(out.(*KubeletAuthentication))
			return nil
		}, InType: reflect.TypeOf(&KubeletAuthentication{})},
		conversion.GeneratedDeepCopyFunc{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*KubeletAuthorization).DeepCopyInto(out.(*KubeletAuthorization))
			return nil
		}, InType: reflect.TypeOf(&KubeletAuthorization{})},
		conversion.GeneratedDeepCopyFunc{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*KubeletConfiguration).DeepCopyInto(out.(*KubeletConfiguration))
			return nil
		}, InType: reflect.TypeOf(&KubeletConfiguration{})},
		conversion.GeneratedDeepCopyFunc{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*KubeletWebhookAuthentication).DeepCopyInto(out.(*KubeletWebhookAuthentication))
			return nil
		}, InType: reflect.TypeOf(&KubeletWebhookAuthentication{})},
		conversion.GeneratedDeepCopyFunc{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*KubeletWebhookAuthorization).DeepCopyInto(out.(*KubeletWebhookAuthorization))
			return nil
		}, InType: reflect.TypeOf(&KubeletWebhookAuthorization{})},
		conversion.GeneratedDeepCopyFunc{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*KubeletX509Authentication).DeepCopyInto(out.(*KubeletX509Authentication))
			return nil
		}, InType: reflect.TypeOf(&KubeletX509Authentication{})},
	)
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KubeletAnonymousAuthentication) DeepCopyInto(out *KubeletAnonymousAuthentication) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KubeletAnonymousAuthentication.
func (in *KubeletAnonymousAuthentication) DeepCopy() *KubeletAnonymousAuthentication {
	if in == nil {
		return nil
	}
	out := new(KubeletAnonymousAuthentication)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KubeletAuthentication) DeepCopyInto(out *KubeletAuthentication) {
	*out = *in
	out.X509 = in.X509
	out.Webhook = in.Webhook
	out.Anonymous = in.Anonymous
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KubeletAuthentication.
func (in *KubeletAuthentication) DeepCopy() *KubeletAuthentication {
	if in == nil {
		return nil
	}
	out := new(KubeletAuthentication)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KubeletAuthorization) DeepCopyInto(out *KubeletAuthorization) {
	*out = *in
	out.Webhook = in.Webhook
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KubeletAuthorization.
func (in *KubeletAuthorization) DeepCopy() *KubeletAuthorization {
	if in == nil {
		return nil
	}
	out := new(KubeletAuthorization)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KubeletConfiguration) DeepCopyInto(out *KubeletConfiguration) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	if in.ConfigTrialDuration != nil {
		in, out := &in.ConfigTrialDuration, &out.ConfigTrialDuration
		if *in == nil {
			*out = nil
		} else {
			*out = new(v1.Duration)
			**out = **in
		}
	}
	out.SyncFrequency = in.SyncFrequency
	out.FileCheckFrequency = in.FileCheckFrequency
	out.HTTPCheckFrequency = in.HTTPCheckFrequency
	out.Authentication = in.Authentication
	out.Authorization = in.Authorization
	if in.HostNetworkSources != nil {
		in, out := &in.HostNetworkSources, &out.HostNetworkSources
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.HostPIDSources != nil {
		in, out := &in.HostPIDSources, &out.HostPIDSources
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.HostIPCSources != nil {
		in, out := &in.HostIPCSources, &out.HostIPCSources
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	out.MinimumGCAge = in.MinimumGCAge
	if in.ClusterDNS != nil {
		in, out := &in.ClusterDNS, &out.ClusterDNS
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	out.StreamingConnectionIdleTimeout = in.StreamingConnectionIdleTimeout
	out.NodeStatusUpdateFrequency = in.NodeStatusUpdateFrequency
	out.ImageMinimumGCAge = in.ImageMinimumGCAge
	out.VolumeStatsAggPeriod = in.VolumeStatsAggPeriod
	out.CPUManagerReconcilePeriod = in.CPUManagerReconcilePeriod
	out.RuntimeRequestTimeout = in.RuntimeRequestTimeout
	if in.RegisterWithTaints != nil {
		in, out := &in.RegisterWithTaints, &out.RegisterWithTaints
		*out = make([]api.Taint, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.NodeLabels != nil {
		in, out := &in.NodeLabels, &out.NodeLabels
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	out.EvictionPressureTransitionPeriod = in.EvictionPressureTransitionPeriod
	if in.ExperimentalQOSReserved != nil {
		in, out := &in.ExperimentalQOSReserved, &out.ExperimentalQOSReserved
		*out = make(ConfigurationMap, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.AllowedUnsafeSysctls != nil {
		in, out := &in.AllowedUnsafeSysctls, &out.AllowedUnsafeSysctls
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.FeatureGates != nil {
		in, out := &in.FeatureGates, &out.FeatureGates
		*out = make(map[string]bool, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.SystemReserved != nil {
		in, out := &in.SystemReserved, &out.SystemReserved
		*out = make(ConfigurationMap, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.KubeReserved != nil {
		in, out := &in.KubeReserved, &out.KubeReserved
		*out = make(ConfigurationMap, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.EnforceNodeAllocatable != nil {
		in, out := &in.EnforceNodeAllocatable, &out.EnforceNodeAllocatable
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KubeletConfiguration.
func (in *KubeletConfiguration) DeepCopy() *KubeletConfiguration {
	if in == nil {
		return nil
	}
	out := new(KubeletConfiguration)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *KubeletConfiguration) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	} else {
		return nil
	}
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KubeletWebhookAuthentication) DeepCopyInto(out *KubeletWebhookAuthentication) {
	*out = *in
	out.CacheTTL = in.CacheTTL
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KubeletWebhookAuthentication.
func (in *KubeletWebhookAuthentication) DeepCopy() *KubeletWebhookAuthentication {
	if in == nil {
		return nil
	}
	out := new(KubeletWebhookAuthentication)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KubeletWebhookAuthorization) DeepCopyInto(out *KubeletWebhookAuthorization) {
	*out = *in
	out.CacheAuthorizedTTL = in.CacheAuthorizedTTL
	out.CacheUnauthorizedTTL = in.CacheUnauthorizedTTL
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KubeletWebhookAuthorization.
func (in *KubeletWebhookAuthorization) DeepCopy() *KubeletWebhookAuthorization {
	if in == nil {
		return nil
	}
	out := new(KubeletWebhookAuthorization)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KubeletX509Authentication) DeepCopyInto(out *KubeletX509Authentication) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KubeletX509Authentication.
func (in *KubeletX509Authentication) DeepCopy() *KubeletX509Authentication {
	if in == nil {
		return nil
	}
	out := new(KubeletX509Authentication)
	in.DeepCopyInto(out)
	return out
}
