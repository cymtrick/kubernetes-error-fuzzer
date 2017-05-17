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

package federation

import (
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	conversion "k8s.io/apimachinery/pkg/conversion"
	runtime "k8s.io/apimachinery/pkg/runtime"
	api "k8s.io/kubernetes/pkg/api"
	api_v1 "k8s.io/kubernetes/pkg/api/v1"
	reflect "reflect"
)

func init() {
	SchemeBuilder.Register(RegisterDeepCopies)
}

// RegisterDeepCopies adds deep-copy functions to the given scheme. Public
// to allow building arbitrary schemes.
func RegisterDeepCopies(scheme *runtime.Scheme) error {
	return scheme.AddGeneratedDeepCopyFuncs(
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_federation_Cluster, InType: reflect.TypeOf(&Cluster{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_federation_ClusterCondition, InType: reflect.TypeOf(&ClusterCondition{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_federation_ClusterList, InType: reflect.TypeOf(&ClusterList{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_federation_ClusterPreferences, InType: reflect.TypeOf(&ClusterPreferences{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_federation_ClusterServiceIngress, InType: reflect.TypeOf(&ClusterServiceIngress{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_federation_ClusterSpec, InType: reflect.TypeOf(&ClusterSpec{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_federation_ClusterStatus, InType: reflect.TypeOf(&ClusterStatus{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_federation_FederatedServiceIngress, InType: reflect.TypeOf(&FederatedServiceIngress{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_federation_ReplicaAllocationPreferences, InType: reflect.TypeOf(&ReplicaAllocationPreferences{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_federation_ServerAddressByClientCIDR, InType: reflect.TypeOf(&ServerAddressByClientCIDR{})},
	)
}

// DeepCopy_federation_Cluster is an autogenerated deepcopy function.
func DeepCopy_federation_Cluster(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*Cluster)
		out := out.(*Cluster)
		*out = *in
		if newVal, err := c.DeepCopy(&in.ObjectMeta); err != nil {
			return err
		} else {
			out.ObjectMeta = *newVal.(*v1.ObjectMeta)
		}
		if err := DeepCopy_federation_ClusterSpec(&in.Spec, &out.Spec, c); err != nil {
			return err
		}
		if err := DeepCopy_federation_ClusterStatus(&in.Status, &out.Status, c); err != nil {
			return err
		}
		return nil
	}
}

// DeepCopy_federation_ClusterCondition is an autogenerated deepcopy function.
func DeepCopy_federation_ClusterCondition(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*ClusterCondition)
		out := out.(*ClusterCondition)
		*out = *in
		out.LastProbeTime = in.LastProbeTime.DeepCopy()
		out.LastTransitionTime = in.LastTransitionTime.DeepCopy()
		return nil
	}
}

// DeepCopy_federation_ClusterList is an autogenerated deepcopy function.
func DeepCopy_federation_ClusterList(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*ClusterList)
		out := out.(*ClusterList)
		*out = *in
		if in.Items != nil {
			in, out := &in.Items, &out.Items
			*out = make([]Cluster, len(*in))
			for i := range *in {
				if err := DeepCopy_federation_Cluster(&(*in)[i], &(*out)[i], c); err != nil {
					return err
				}
			}
		}
		return nil
	}
}

// DeepCopy_federation_ClusterPreferences is an autogenerated deepcopy function.
func DeepCopy_federation_ClusterPreferences(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*ClusterPreferences)
		out := out.(*ClusterPreferences)
		*out = *in
		if in.MaxReplicas != nil {
			in, out := &in.MaxReplicas, &out.MaxReplicas
			*out = new(int64)
			**out = **in
		}
		return nil
	}
}

// DeepCopy_federation_ClusterServiceIngress is an autogenerated deepcopy function.
func DeepCopy_federation_ClusterServiceIngress(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*ClusterServiceIngress)
		out := out.(*ClusterServiceIngress)
		*out = *in
		if in.Items != nil {
			in, out := &in.Items, &out.Items
			*out = make([]api_v1.LoadBalancerIngress, len(*in))
			copy(*out, *in)
		}
		return nil
	}
}

// DeepCopy_federation_ClusterSpec is an autogenerated deepcopy function.
func DeepCopy_federation_ClusterSpec(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*ClusterSpec)
		out := out.(*ClusterSpec)
		*out = *in
		if in.ServerAddressByClientCIDRs != nil {
			in, out := &in.ServerAddressByClientCIDRs, &out.ServerAddressByClientCIDRs
			*out = make([]ServerAddressByClientCIDR, len(*in))
			copy(*out, *in)
		}
		if in.SecretRef != nil {
			in, out := &in.SecretRef, &out.SecretRef
			*out = new(api.LocalObjectReference)
			**out = **in
		}
		return nil
	}
}

// DeepCopy_federation_ClusterStatus is an autogenerated deepcopy function.
func DeepCopy_federation_ClusterStatus(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*ClusterStatus)
		out := out.(*ClusterStatus)
		*out = *in
		if in.Conditions != nil {
			in, out := &in.Conditions, &out.Conditions
			*out = make([]ClusterCondition, len(*in))
			for i := range *in {
				if err := DeepCopy_federation_ClusterCondition(&(*in)[i], &(*out)[i], c); err != nil {
					return err
				}
			}
		}
		if in.Zones != nil {
			in, out := &in.Zones, &out.Zones
			*out = make([]string, len(*in))
			copy(*out, *in)
		}
		return nil
	}
}

// DeepCopy_federation_FederatedServiceIngress is an autogenerated deepcopy function.
func DeepCopy_federation_FederatedServiceIngress(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*FederatedServiceIngress)
		out := out.(*FederatedServiceIngress)
		*out = *in
		if in.Items != nil {
			in, out := &in.Items, &out.Items
			*out = make([]ClusterServiceIngress, len(*in))
			for i := range *in {
				if err := DeepCopy_federation_ClusterServiceIngress(&(*in)[i], &(*out)[i], c); err != nil {
					return err
				}
			}
		}
		return nil
	}
}

// DeepCopy_federation_ReplicaAllocationPreferences is an autogenerated deepcopy function.
func DeepCopy_federation_ReplicaAllocationPreferences(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*ReplicaAllocationPreferences)
		out := out.(*ReplicaAllocationPreferences)
		*out = *in
		if in.Clusters != nil {
			in, out := &in.Clusters, &out.Clusters
			*out = make(map[string]ClusterPreferences)
			for key, val := range *in {
				newVal := new(ClusterPreferences)
				if err := DeepCopy_federation_ClusterPreferences(&val, newVal, c); err != nil {
					return err
				}
				(*out)[key] = *newVal
			}
		}
		return nil
	}
}

// DeepCopy_federation_ServerAddressByClientCIDR is an autogenerated deepcopy function.
func DeepCopy_federation_ServerAddressByClientCIDR(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*ServerAddressByClientCIDR)
		out := out.(*ServerAddressByClientCIDR)
		*out = *in
		return nil
	}
}
