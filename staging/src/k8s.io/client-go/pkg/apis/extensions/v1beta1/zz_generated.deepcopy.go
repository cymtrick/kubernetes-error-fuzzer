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

package v1beta1

import (
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	conversion "k8s.io/apimachinery/pkg/conversion"
	runtime "k8s.io/apimachinery/pkg/runtime"
	intstr "k8s.io/apimachinery/pkg/util/intstr"
	api_v1 "k8s.io/client-go/pkg/api/v1"
	reflect "reflect"
)

func init() {
	SchemeBuilder.Register(RegisterDeepCopies)
}

// RegisterDeepCopies adds deep-copy functions to the given scheme. Public
// to allow building arbitrary schemes.
func RegisterDeepCopies(scheme *runtime.Scheme) error {
	return scheme.AddGeneratedDeepCopyFuncs(
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_v1beta1_APIVersion, InType: reflect.TypeOf(&APIVersion{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_v1beta1_CustomMetricCurrentStatus, InType: reflect.TypeOf(&CustomMetricCurrentStatus{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_v1beta1_CustomMetricCurrentStatusList, InType: reflect.TypeOf(&CustomMetricCurrentStatusList{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_v1beta1_CustomMetricTarget, InType: reflect.TypeOf(&CustomMetricTarget{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_v1beta1_CustomMetricTargetList, InType: reflect.TypeOf(&CustomMetricTargetList{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_v1beta1_DaemonSet, InType: reflect.TypeOf(&DaemonSet{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_v1beta1_DaemonSetList, InType: reflect.TypeOf(&DaemonSetList{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_v1beta1_DaemonSetSpec, InType: reflect.TypeOf(&DaemonSetSpec{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_v1beta1_DaemonSetStatus, InType: reflect.TypeOf(&DaemonSetStatus{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_v1beta1_DaemonSetUpdateStrategy, InType: reflect.TypeOf(&DaemonSetUpdateStrategy{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_v1beta1_Deployment, InType: reflect.TypeOf(&Deployment{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_v1beta1_DeploymentCondition, InType: reflect.TypeOf(&DeploymentCondition{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_v1beta1_DeploymentList, InType: reflect.TypeOf(&DeploymentList{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_v1beta1_DeploymentRollback, InType: reflect.TypeOf(&DeploymentRollback{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_v1beta1_DeploymentSpec, InType: reflect.TypeOf(&DeploymentSpec{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_v1beta1_DeploymentStatus, InType: reflect.TypeOf(&DeploymentStatus{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_v1beta1_DeploymentStrategy, InType: reflect.TypeOf(&DeploymentStrategy{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_v1beta1_FSGroupStrategyOptions, InType: reflect.TypeOf(&FSGroupStrategyOptions{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_v1beta1_HTTPIngressPath, InType: reflect.TypeOf(&HTTPIngressPath{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_v1beta1_HTTPIngressRuleValue, InType: reflect.TypeOf(&HTTPIngressRuleValue{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_v1beta1_HostPortRange, InType: reflect.TypeOf(&HostPortRange{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_v1beta1_IDRange, InType: reflect.TypeOf(&IDRange{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_v1beta1_Ingress, InType: reflect.TypeOf(&Ingress{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_v1beta1_IngressBackend, InType: reflect.TypeOf(&IngressBackend{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_v1beta1_IngressList, InType: reflect.TypeOf(&IngressList{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_v1beta1_IngressRule, InType: reflect.TypeOf(&IngressRule{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_v1beta1_IngressRuleValue, InType: reflect.TypeOf(&IngressRuleValue{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_v1beta1_IngressSpec, InType: reflect.TypeOf(&IngressSpec{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_v1beta1_IngressStatus, InType: reflect.TypeOf(&IngressStatus{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_v1beta1_IngressTLS, InType: reflect.TypeOf(&IngressTLS{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_v1beta1_NetworkPolicy, InType: reflect.TypeOf(&NetworkPolicy{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_v1beta1_NetworkPolicyIngressRule, InType: reflect.TypeOf(&NetworkPolicyIngressRule{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_v1beta1_NetworkPolicyList, InType: reflect.TypeOf(&NetworkPolicyList{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_v1beta1_NetworkPolicyPeer, InType: reflect.TypeOf(&NetworkPolicyPeer{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_v1beta1_NetworkPolicyPort, InType: reflect.TypeOf(&NetworkPolicyPort{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_v1beta1_NetworkPolicySpec, InType: reflect.TypeOf(&NetworkPolicySpec{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_v1beta1_PodSecurityPolicy, InType: reflect.TypeOf(&PodSecurityPolicy{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_v1beta1_PodSecurityPolicyList, InType: reflect.TypeOf(&PodSecurityPolicyList{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_v1beta1_PodSecurityPolicySpec, InType: reflect.TypeOf(&PodSecurityPolicySpec{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_v1beta1_ReplicaSet, InType: reflect.TypeOf(&ReplicaSet{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_v1beta1_ReplicaSetCondition, InType: reflect.TypeOf(&ReplicaSetCondition{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_v1beta1_ReplicaSetList, InType: reflect.TypeOf(&ReplicaSetList{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_v1beta1_ReplicaSetSpec, InType: reflect.TypeOf(&ReplicaSetSpec{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_v1beta1_ReplicaSetStatus, InType: reflect.TypeOf(&ReplicaSetStatus{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_v1beta1_ReplicationControllerDummy, InType: reflect.TypeOf(&ReplicationControllerDummy{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_v1beta1_RollbackConfig, InType: reflect.TypeOf(&RollbackConfig{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_v1beta1_RollingUpdateDaemonSet, InType: reflect.TypeOf(&RollingUpdateDaemonSet{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_v1beta1_RollingUpdateDeployment, InType: reflect.TypeOf(&RollingUpdateDeployment{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_v1beta1_RunAsUserStrategyOptions, InType: reflect.TypeOf(&RunAsUserStrategyOptions{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_v1beta1_SELinuxStrategyOptions, InType: reflect.TypeOf(&SELinuxStrategyOptions{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_v1beta1_Scale, InType: reflect.TypeOf(&Scale{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_v1beta1_ScaleSpec, InType: reflect.TypeOf(&ScaleSpec{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_v1beta1_ScaleStatus, InType: reflect.TypeOf(&ScaleStatus{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_v1beta1_SupplementalGroupsStrategyOptions, InType: reflect.TypeOf(&SupplementalGroupsStrategyOptions{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_v1beta1_ThirdPartyResource, InType: reflect.TypeOf(&ThirdPartyResource{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_v1beta1_ThirdPartyResourceData, InType: reflect.TypeOf(&ThirdPartyResourceData{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_v1beta1_ThirdPartyResourceDataList, InType: reflect.TypeOf(&ThirdPartyResourceDataList{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_v1beta1_ThirdPartyResourceList, InType: reflect.TypeOf(&ThirdPartyResourceList{})},
	)
}

// DeepCopy_v1beta1_APIVersion is an autogenerated deepcopy function.
func DeepCopy_v1beta1_APIVersion(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*APIVersion)
		out := out.(*APIVersion)
		*out = *in
		return nil
	}
}

// DeepCopy_v1beta1_CustomMetricCurrentStatus is an autogenerated deepcopy function.
func DeepCopy_v1beta1_CustomMetricCurrentStatus(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*CustomMetricCurrentStatus)
		out := out.(*CustomMetricCurrentStatus)
		*out = *in
		out.CurrentValue = in.CurrentValue.DeepCopy()
		return nil
	}
}

// DeepCopy_v1beta1_CustomMetricCurrentStatusList is an autogenerated deepcopy function.
func DeepCopy_v1beta1_CustomMetricCurrentStatusList(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*CustomMetricCurrentStatusList)
		out := out.(*CustomMetricCurrentStatusList)
		*out = *in
		if in.Items != nil {
			in, out := &in.Items, &out.Items
			*out = make([]CustomMetricCurrentStatus, len(*in))
			for i := range *in {
				if err := DeepCopy_v1beta1_CustomMetricCurrentStatus(&(*in)[i], &(*out)[i], c); err != nil {
					return err
				}
			}
		}
		return nil
	}
}

// DeepCopy_v1beta1_CustomMetricTarget is an autogenerated deepcopy function.
func DeepCopy_v1beta1_CustomMetricTarget(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*CustomMetricTarget)
		out := out.(*CustomMetricTarget)
		*out = *in
		out.TargetValue = in.TargetValue.DeepCopy()
		return nil
	}
}

// DeepCopy_v1beta1_CustomMetricTargetList is an autogenerated deepcopy function.
func DeepCopy_v1beta1_CustomMetricTargetList(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*CustomMetricTargetList)
		out := out.(*CustomMetricTargetList)
		*out = *in
		if in.Items != nil {
			in, out := &in.Items, &out.Items
			*out = make([]CustomMetricTarget, len(*in))
			for i := range *in {
				if err := DeepCopy_v1beta1_CustomMetricTarget(&(*in)[i], &(*out)[i], c); err != nil {
					return err
				}
			}
		}
		return nil
	}
}

// DeepCopy_v1beta1_DaemonSet is an autogenerated deepcopy function.
func DeepCopy_v1beta1_DaemonSet(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*DaemonSet)
		out := out.(*DaemonSet)
		*out = *in
		if newVal, err := c.DeepCopy(&in.ObjectMeta); err != nil {
			return err
		} else {
			out.ObjectMeta = *newVal.(*v1.ObjectMeta)
		}
		if err := DeepCopy_v1beta1_DaemonSetSpec(&in.Spec, &out.Spec, c); err != nil {
			return err
		}
		return nil
	}
}

// DeepCopy_v1beta1_DaemonSetList is an autogenerated deepcopy function.
func DeepCopy_v1beta1_DaemonSetList(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*DaemonSetList)
		out := out.(*DaemonSetList)
		*out = *in
		if in.Items != nil {
			in, out := &in.Items, &out.Items
			*out = make([]DaemonSet, len(*in))
			for i := range *in {
				if err := DeepCopy_v1beta1_DaemonSet(&(*in)[i], &(*out)[i], c); err != nil {
					return err
				}
			}
		}
		return nil
	}
}

// DeepCopy_v1beta1_DaemonSetSpec is an autogenerated deepcopy function.
func DeepCopy_v1beta1_DaemonSetSpec(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*DaemonSetSpec)
		out := out.(*DaemonSetSpec)
		*out = *in
		if in.Selector != nil {
			in, out := &in.Selector, &out.Selector
			if newVal, err := c.DeepCopy(*in); err != nil {
				return err
			} else {
				*out = newVal.(*v1.LabelSelector)
			}
		}
		if err := api_v1.DeepCopy_v1_PodTemplateSpec(&in.Template, &out.Template, c); err != nil {
			return err
		}
		if err := DeepCopy_v1beta1_DaemonSetUpdateStrategy(&in.UpdateStrategy, &out.UpdateStrategy, c); err != nil {
			return err
		}
		return nil
	}
}

// DeepCopy_v1beta1_DaemonSetStatus is an autogenerated deepcopy function.
func DeepCopy_v1beta1_DaemonSetStatus(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*DaemonSetStatus)
		out := out.(*DaemonSetStatus)
		*out = *in
		return nil
	}
}

// DeepCopy_v1beta1_DaemonSetUpdateStrategy is an autogenerated deepcopy function.
func DeepCopy_v1beta1_DaemonSetUpdateStrategy(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*DaemonSetUpdateStrategy)
		out := out.(*DaemonSetUpdateStrategy)
		*out = *in
		if in.RollingUpdate != nil {
			in, out := &in.RollingUpdate, &out.RollingUpdate
			*out = new(RollingUpdateDaemonSet)
			if err := DeepCopy_v1beta1_RollingUpdateDaemonSet(*in, *out, c); err != nil {
				return err
			}
		}
		return nil
	}
}

// DeepCopy_v1beta1_Deployment is an autogenerated deepcopy function.
func DeepCopy_v1beta1_Deployment(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*Deployment)
		out := out.(*Deployment)
		*out = *in
		if newVal, err := c.DeepCopy(&in.ObjectMeta); err != nil {
			return err
		} else {
			out.ObjectMeta = *newVal.(*v1.ObjectMeta)
		}
		if err := DeepCopy_v1beta1_DeploymentSpec(&in.Spec, &out.Spec, c); err != nil {
			return err
		}
		if err := DeepCopy_v1beta1_DeploymentStatus(&in.Status, &out.Status, c); err != nil {
			return err
		}
		return nil
	}
}

// DeepCopy_v1beta1_DeploymentCondition is an autogenerated deepcopy function.
func DeepCopy_v1beta1_DeploymentCondition(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*DeploymentCondition)
		out := out.(*DeploymentCondition)
		*out = *in
		out.LastUpdateTime = in.LastUpdateTime.DeepCopy()
		out.LastTransitionTime = in.LastTransitionTime.DeepCopy()
		return nil
	}
}

// DeepCopy_v1beta1_DeploymentList is an autogenerated deepcopy function.
func DeepCopy_v1beta1_DeploymentList(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*DeploymentList)
		out := out.(*DeploymentList)
		*out = *in
		if in.Items != nil {
			in, out := &in.Items, &out.Items
			*out = make([]Deployment, len(*in))
			for i := range *in {
				if err := DeepCopy_v1beta1_Deployment(&(*in)[i], &(*out)[i], c); err != nil {
					return err
				}
			}
		}
		return nil
	}
}

// DeepCopy_v1beta1_DeploymentRollback is an autogenerated deepcopy function.
func DeepCopy_v1beta1_DeploymentRollback(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*DeploymentRollback)
		out := out.(*DeploymentRollback)
		*out = *in
		if in.UpdatedAnnotations != nil {
			in, out := &in.UpdatedAnnotations, &out.UpdatedAnnotations
			*out = make(map[string]string)
			for key, val := range *in {
				(*out)[key] = val
			}
		}
		return nil
	}
}

// DeepCopy_v1beta1_DeploymentSpec is an autogenerated deepcopy function.
func DeepCopy_v1beta1_DeploymentSpec(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*DeploymentSpec)
		out := out.(*DeploymentSpec)
		*out = *in
		if in.Replicas != nil {
			in, out := &in.Replicas, &out.Replicas
			*out = new(int32)
			**out = **in
		}
		if in.Selector != nil {
			in, out := &in.Selector, &out.Selector
			if newVal, err := c.DeepCopy(*in); err != nil {
				return err
			} else {
				*out = newVal.(*v1.LabelSelector)
			}
		}
		if err := api_v1.DeepCopy_v1_PodTemplateSpec(&in.Template, &out.Template, c); err != nil {
			return err
		}
		if err := DeepCopy_v1beta1_DeploymentStrategy(&in.Strategy, &out.Strategy, c); err != nil {
			return err
		}
		if in.RevisionHistoryLimit != nil {
			in, out := &in.RevisionHistoryLimit, &out.RevisionHistoryLimit
			*out = new(int32)
			**out = **in
		}
		if in.RollbackTo != nil {
			in, out := &in.RollbackTo, &out.RollbackTo
			*out = new(RollbackConfig)
			**out = **in
		}
		if in.ProgressDeadlineSeconds != nil {
			in, out := &in.ProgressDeadlineSeconds, &out.ProgressDeadlineSeconds
			*out = new(int32)
			**out = **in
		}
		return nil
	}
}

// DeepCopy_v1beta1_DeploymentStatus is an autogenerated deepcopy function.
func DeepCopy_v1beta1_DeploymentStatus(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*DeploymentStatus)
		out := out.(*DeploymentStatus)
		*out = *in
		if in.Conditions != nil {
			in, out := &in.Conditions, &out.Conditions
			*out = make([]DeploymentCondition, len(*in))
			for i := range *in {
				if err := DeepCopy_v1beta1_DeploymentCondition(&(*in)[i], &(*out)[i], c); err != nil {
					return err
				}
			}
		}
		if in.CollisionCount != nil {
			in, out := &in.CollisionCount, &out.CollisionCount
			*out = new(int64)
			**out = **in
		}
		return nil
	}
}

// DeepCopy_v1beta1_DeploymentStrategy is an autogenerated deepcopy function.
func DeepCopy_v1beta1_DeploymentStrategy(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*DeploymentStrategy)
		out := out.(*DeploymentStrategy)
		*out = *in
		if in.RollingUpdate != nil {
			in, out := &in.RollingUpdate, &out.RollingUpdate
			*out = new(RollingUpdateDeployment)
			if err := DeepCopy_v1beta1_RollingUpdateDeployment(*in, *out, c); err != nil {
				return err
			}
		}
		return nil
	}
}

// DeepCopy_v1beta1_FSGroupStrategyOptions is an autogenerated deepcopy function.
func DeepCopy_v1beta1_FSGroupStrategyOptions(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*FSGroupStrategyOptions)
		out := out.(*FSGroupStrategyOptions)
		*out = *in
		if in.Ranges != nil {
			in, out := &in.Ranges, &out.Ranges
			*out = make([]IDRange, len(*in))
			copy(*out, *in)
		}
		return nil
	}
}

// DeepCopy_v1beta1_HTTPIngressPath is an autogenerated deepcopy function.
func DeepCopy_v1beta1_HTTPIngressPath(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*HTTPIngressPath)
		out := out.(*HTTPIngressPath)
		*out = *in
		return nil
	}
}

// DeepCopy_v1beta1_HTTPIngressRuleValue is an autogenerated deepcopy function.
func DeepCopy_v1beta1_HTTPIngressRuleValue(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*HTTPIngressRuleValue)
		out := out.(*HTTPIngressRuleValue)
		*out = *in
		if in.Paths != nil {
			in, out := &in.Paths, &out.Paths
			*out = make([]HTTPIngressPath, len(*in))
			copy(*out, *in)
		}
		return nil
	}
}

// DeepCopy_v1beta1_HostPortRange is an autogenerated deepcopy function.
func DeepCopy_v1beta1_HostPortRange(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*HostPortRange)
		out := out.(*HostPortRange)
		*out = *in
		return nil
	}
}

// DeepCopy_v1beta1_IDRange is an autogenerated deepcopy function.
func DeepCopy_v1beta1_IDRange(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*IDRange)
		out := out.(*IDRange)
		*out = *in
		return nil
	}
}

// DeepCopy_v1beta1_Ingress is an autogenerated deepcopy function.
func DeepCopy_v1beta1_Ingress(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*Ingress)
		out := out.(*Ingress)
		*out = *in
		if newVal, err := c.DeepCopy(&in.ObjectMeta); err != nil {
			return err
		} else {
			out.ObjectMeta = *newVal.(*v1.ObjectMeta)
		}
		if err := DeepCopy_v1beta1_IngressSpec(&in.Spec, &out.Spec, c); err != nil {
			return err
		}
		if err := DeepCopy_v1beta1_IngressStatus(&in.Status, &out.Status, c); err != nil {
			return err
		}
		return nil
	}
}

// DeepCopy_v1beta1_IngressBackend is an autogenerated deepcopy function.
func DeepCopy_v1beta1_IngressBackend(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*IngressBackend)
		out := out.(*IngressBackend)
		*out = *in
		return nil
	}
}

// DeepCopy_v1beta1_IngressList is an autogenerated deepcopy function.
func DeepCopy_v1beta1_IngressList(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*IngressList)
		out := out.(*IngressList)
		*out = *in
		if in.Items != nil {
			in, out := &in.Items, &out.Items
			*out = make([]Ingress, len(*in))
			for i := range *in {
				if err := DeepCopy_v1beta1_Ingress(&(*in)[i], &(*out)[i], c); err != nil {
					return err
				}
			}
		}
		return nil
	}
}

// DeepCopy_v1beta1_IngressRule is an autogenerated deepcopy function.
func DeepCopy_v1beta1_IngressRule(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*IngressRule)
		out := out.(*IngressRule)
		*out = *in
		if err := DeepCopy_v1beta1_IngressRuleValue(&in.IngressRuleValue, &out.IngressRuleValue, c); err != nil {
			return err
		}
		return nil
	}
}

// DeepCopy_v1beta1_IngressRuleValue is an autogenerated deepcopy function.
func DeepCopy_v1beta1_IngressRuleValue(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*IngressRuleValue)
		out := out.(*IngressRuleValue)
		*out = *in
		if in.HTTP != nil {
			in, out := &in.HTTP, &out.HTTP
			*out = new(HTTPIngressRuleValue)
			if err := DeepCopy_v1beta1_HTTPIngressRuleValue(*in, *out, c); err != nil {
				return err
			}
		}
		return nil
	}
}

// DeepCopy_v1beta1_IngressSpec is an autogenerated deepcopy function.
func DeepCopy_v1beta1_IngressSpec(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*IngressSpec)
		out := out.(*IngressSpec)
		*out = *in
		if in.Backend != nil {
			in, out := &in.Backend, &out.Backend
			*out = new(IngressBackend)
			**out = **in
		}
		if in.TLS != nil {
			in, out := &in.TLS, &out.TLS
			*out = make([]IngressTLS, len(*in))
			for i := range *in {
				if err := DeepCopy_v1beta1_IngressTLS(&(*in)[i], &(*out)[i], c); err != nil {
					return err
				}
			}
		}
		if in.Rules != nil {
			in, out := &in.Rules, &out.Rules
			*out = make([]IngressRule, len(*in))
			for i := range *in {
				if err := DeepCopy_v1beta1_IngressRule(&(*in)[i], &(*out)[i], c); err != nil {
					return err
				}
			}
		}
		return nil
	}
}

// DeepCopy_v1beta1_IngressStatus is an autogenerated deepcopy function.
func DeepCopy_v1beta1_IngressStatus(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*IngressStatus)
		out := out.(*IngressStatus)
		*out = *in
		if err := api_v1.DeepCopy_v1_LoadBalancerStatus(&in.LoadBalancer, &out.LoadBalancer, c); err != nil {
			return err
		}
		return nil
	}
}

// DeepCopy_v1beta1_IngressTLS is an autogenerated deepcopy function.
func DeepCopy_v1beta1_IngressTLS(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*IngressTLS)
		out := out.(*IngressTLS)
		*out = *in
		if in.Hosts != nil {
			in, out := &in.Hosts, &out.Hosts
			*out = make([]string, len(*in))
			copy(*out, *in)
		}
		return nil
	}
}

// DeepCopy_v1beta1_NetworkPolicy is an autogenerated deepcopy function.
func DeepCopy_v1beta1_NetworkPolicy(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*NetworkPolicy)
		out := out.(*NetworkPolicy)
		*out = *in
		if newVal, err := c.DeepCopy(&in.ObjectMeta); err != nil {
			return err
		} else {
			out.ObjectMeta = *newVal.(*v1.ObjectMeta)
		}
		if err := DeepCopy_v1beta1_NetworkPolicySpec(&in.Spec, &out.Spec, c); err != nil {
			return err
		}
		return nil
	}
}

// DeepCopy_v1beta1_NetworkPolicyIngressRule is an autogenerated deepcopy function.
func DeepCopy_v1beta1_NetworkPolicyIngressRule(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*NetworkPolicyIngressRule)
		out := out.(*NetworkPolicyIngressRule)
		*out = *in
		if in.Ports != nil {
			in, out := &in.Ports, &out.Ports
			*out = make([]NetworkPolicyPort, len(*in))
			for i := range *in {
				if err := DeepCopy_v1beta1_NetworkPolicyPort(&(*in)[i], &(*out)[i], c); err != nil {
					return err
				}
			}
		}
		if in.From != nil {
			in, out := &in.From, &out.From
			*out = make([]NetworkPolicyPeer, len(*in))
			for i := range *in {
				if err := DeepCopy_v1beta1_NetworkPolicyPeer(&(*in)[i], &(*out)[i], c); err != nil {
					return err
				}
			}
		}
		return nil
	}
}

// DeepCopy_v1beta1_NetworkPolicyList is an autogenerated deepcopy function.
func DeepCopy_v1beta1_NetworkPolicyList(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*NetworkPolicyList)
		out := out.(*NetworkPolicyList)
		*out = *in
		if in.Items != nil {
			in, out := &in.Items, &out.Items
			*out = make([]NetworkPolicy, len(*in))
			for i := range *in {
				if err := DeepCopy_v1beta1_NetworkPolicy(&(*in)[i], &(*out)[i], c); err != nil {
					return err
				}
			}
		}
		return nil
	}
}

// DeepCopy_v1beta1_NetworkPolicyPeer is an autogenerated deepcopy function.
func DeepCopy_v1beta1_NetworkPolicyPeer(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*NetworkPolicyPeer)
		out := out.(*NetworkPolicyPeer)
		*out = *in
		if in.PodSelector != nil {
			in, out := &in.PodSelector, &out.PodSelector
			if newVal, err := c.DeepCopy(*in); err != nil {
				return err
			} else {
				*out = newVal.(*v1.LabelSelector)
			}
		}
		if in.NamespaceSelector != nil {
			in, out := &in.NamespaceSelector, &out.NamespaceSelector
			if newVal, err := c.DeepCopy(*in); err != nil {
				return err
			} else {
				*out = newVal.(*v1.LabelSelector)
			}
		}
		return nil
	}
}

// DeepCopy_v1beta1_NetworkPolicyPort is an autogenerated deepcopy function.
func DeepCopy_v1beta1_NetworkPolicyPort(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*NetworkPolicyPort)
		out := out.(*NetworkPolicyPort)
		*out = *in
		if in.Protocol != nil {
			in, out := &in.Protocol, &out.Protocol
			*out = new(api_v1.Protocol)
			**out = **in
		}
		if in.Port != nil {
			in, out := &in.Port, &out.Port
			*out = new(intstr.IntOrString)
			**out = **in
		}
		return nil
	}
}

// DeepCopy_v1beta1_NetworkPolicySpec is an autogenerated deepcopy function.
func DeepCopy_v1beta1_NetworkPolicySpec(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*NetworkPolicySpec)
		out := out.(*NetworkPolicySpec)
		*out = *in
		if newVal, err := c.DeepCopy(&in.PodSelector); err != nil {
			return err
		} else {
			out.PodSelector = *newVal.(*v1.LabelSelector)
		}
		if in.Ingress != nil {
			in, out := &in.Ingress, &out.Ingress
			*out = make([]NetworkPolicyIngressRule, len(*in))
			for i := range *in {
				if err := DeepCopy_v1beta1_NetworkPolicyIngressRule(&(*in)[i], &(*out)[i], c); err != nil {
					return err
				}
			}
		}
		return nil
	}
}

// DeepCopy_v1beta1_PodSecurityPolicy is an autogenerated deepcopy function.
func DeepCopy_v1beta1_PodSecurityPolicy(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*PodSecurityPolicy)
		out := out.(*PodSecurityPolicy)
		*out = *in
		if newVal, err := c.DeepCopy(&in.ObjectMeta); err != nil {
			return err
		} else {
			out.ObjectMeta = *newVal.(*v1.ObjectMeta)
		}
		if err := DeepCopy_v1beta1_PodSecurityPolicySpec(&in.Spec, &out.Spec, c); err != nil {
			return err
		}
		return nil
	}
}

// DeepCopy_v1beta1_PodSecurityPolicyList is an autogenerated deepcopy function.
func DeepCopy_v1beta1_PodSecurityPolicyList(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*PodSecurityPolicyList)
		out := out.(*PodSecurityPolicyList)
		*out = *in
		if in.Items != nil {
			in, out := &in.Items, &out.Items
			*out = make([]PodSecurityPolicy, len(*in))
			for i := range *in {
				if err := DeepCopy_v1beta1_PodSecurityPolicy(&(*in)[i], &(*out)[i], c); err != nil {
					return err
				}
			}
		}
		return nil
	}
}

// DeepCopy_v1beta1_PodSecurityPolicySpec is an autogenerated deepcopy function.
func DeepCopy_v1beta1_PodSecurityPolicySpec(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*PodSecurityPolicySpec)
		out := out.(*PodSecurityPolicySpec)
		*out = *in
		if in.DefaultAddCapabilities != nil {
			in, out := &in.DefaultAddCapabilities, &out.DefaultAddCapabilities
			*out = make([]api_v1.Capability, len(*in))
			copy(*out, *in)
		}
		if in.RequiredDropCapabilities != nil {
			in, out := &in.RequiredDropCapabilities, &out.RequiredDropCapabilities
			*out = make([]api_v1.Capability, len(*in))
			copy(*out, *in)
		}
		if in.AllowedCapabilities != nil {
			in, out := &in.AllowedCapabilities, &out.AllowedCapabilities
			*out = make([]api_v1.Capability, len(*in))
			copy(*out, *in)
		}
		if in.Volumes != nil {
			in, out := &in.Volumes, &out.Volumes
			*out = make([]FSType, len(*in))
			copy(*out, *in)
		}
		if in.HostPorts != nil {
			in, out := &in.HostPorts, &out.HostPorts
			*out = make([]HostPortRange, len(*in))
			copy(*out, *in)
		}
		if err := DeepCopy_v1beta1_SELinuxStrategyOptions(&in.SELinux, &out.SELinux, c); err != nil {
			return err
		}
		if err := DeepCopy_v1beta1_RunAsUserStrategyOptions(&in.RunAsUser, &out.RunAsUser, c); err != nil {
			return err
		}
		if err := DeepCopy_v1beta1_SupplementalGroupsStrategyOptions(&in.SupplementalGroups, &out.SupplementalGroups, c); err != nil {
			return err
		}
		if err := DeepCopy_v1beta1_FSGroupStrategyOptions(&in.FSGroup, &out.FSGroup, c); err != nil {
			return err
		}
		return nil
	}
}

// DeepCopy_v1beta1_ReplicaSet is an autogenerated deepcopy function.
func DeepCopy_v1beta1_ReplicaSet(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*ReplicaSet)
		out := out.(*ReplicaSet)
		*out = *in
		if newVal, err := c.DeepCopy(&in.ObjectMeta); err != nil {
			return err
		} else {
			out.ObjectMeta = *newVal.(*v1.ObjectMeta)
		}
		if err := DeepCopy_v1beta1_ReplicaSetSpec(&in.Spec, &out.Spec, c); err != nil {
			return err
		}
		if err := DeepCopy_v1beta1_ReplicaSetStatus(&in.Status, &out.Status, c); err != nil {
			return err
		}
		return nil
	}
}

// DeepCopy_v1beta1_ReplicaSetCondition is an autogenerated deepcopy function.
func DeepCopy_v1beta1_ReplicaSetCondition(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*ReplicaSetCondition)
		out := out.(*ReplicaSetCondition)
		*out = *in
		out.LastTransitionTime = in.LastTransitionTime.DeepCopy()
		return nil
	}
}

// DeepCopy_v1beta1_ReplicaSetList is an autogenerated deepcopy function.
func DeepCopy_v1beta1_ReplicaSetList(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*ReplicaSetList)
		out := out.(*ReplicaSetList)
		*out = *in
		if in.Items != nil {
			in, out := &in.Items, &out.Items
			*out = make([]ReplicaSet, len(*in))
			for i := range *in {
				if err := DeepCopy_v1beta1_ReplicaSet(&(*in)[i], &(*out)[i], c); err != nil {
					return err
				}
			}
		}
		return nil
	}
}

// DeepCopy_v1beta1_ReplicaSetSpec is an autogenerated deepcopy function.
func DeepCopy_v1beta1_ReplicaSetSpec(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*ReplicaSetSpec)
		out := out.(*ReplicaSetSpec)
		*out = *in
		if in.Replicas != nil {
			in, out := &in.Replicas, &out.Replicas
			*out = new(int32)
			**out = **in
		}
		if in.Selector != nil {
			in, out := &in.Selector, &out.Selector
			if newVal, err := c.DeepCopy(*in); err != nil {
				return err
			} else {
				*out = newVal.(*v1.LabelSelector)
			}
		}
		if err := api_v1.DeepCopy_v1_PodTemplateSpec(&in.Template, &out.Template, c); err != nil {
			return err
		}
		return nil
	}
}

// DeepCopy_v1beta1_ReplicaSetStatus is an autogenerated deepcopy function.
func DeepCopy_v1beta1_ReplicaSetStatus(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*ReplicaSetStatus)
		out := out.(*ReplicaSetStatus)
		*out = *in
		if in.Conditions != nil {
			in, out := &in.Conditions, &out.Conditions
			*out = make([]ReplicaSetCondition, len(*in))
			for i := range *in {
				if err := DeepCopy_v1beta1_ReplicaSetCondition(&(*in)[i], &(*out)[i], c); err != nil {
					return err
				}
			}
		}
		return nil
	}
}

// DeepCopy_v1beta1_ReplicationControllerDummy is an autogenerated deepcopy function.
func DeepCopy_v1beta1_ReplicationControllerDummy(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*ReplicationControllerDummy)
		out := out.(*ReplicationControllerDummy)
		*out = *in
		return nil
	}
}

// DeepCopy_v1beta1_RollbackConfig is an autogenerated deepcopy function.
func DeepCopy_v1beta1_RollbackConfig(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*RollbackConfig)
		out := out.(*RollbackConfig)
		*out = *in
		return nil
	}
}

// DeepCopy_v1beta1_RollingUpdateDaemonSet is an autogenerated deepcopy function.
func DeepCopy_v1beta1_RollingUpdateDaemonSet(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*RollingUpdateDaemonSet)
		out := out.(*RollingUpdateDaemonSet)
		*out = *in
		if in.MaxUnavailable != nil {
			in, out := &in.MaxUnavailable, &out.MaxUnavailable
			*out = new(intstr.IntOrString)
			**out = **in
		}
		return nil
	}
}

// DeepCopy_v1beta1_RollingUpdateDeployment is an autogenerated deepcopy function.
func DeepCopy_v1beta1_RollingUpdateDeployment(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*RollingUpdateDeployment)
		out := out.(*RollingUpdateDeployment)
		*out = *in
		if in.MaxUnavailable != nil {
			in, out := &in.MaxUnavailable, &out.MaxUnavailable
			*out = new(intstr.IntOrString)
			**out = **in
		}
		if in.MaxSurge != nil {
			in, out := &in.MaxSurge, &out.MaxSurge
			*out = new(intstr.IntOrString)
			**out = **in
		}
		return nil
	}
}

// DeepCopy_v1beta1_RunAsUserStrategyOptions is an autogenerated deepcopy function.
func DeepCopy_v1beta1_RunAsUserStrategyOptions(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*RunAsUserStrategyOptions)
		out := out.(*RunAsUserStrategyOptions)
		*out = *in
		if in.Ranges != nil {
			in, out := &in.Ranges, &out.Ranges
			*out = make([]IDRange, len(*in))
			copy(*out, *in)
		}
		return nil
	}
}

// DeepCopy_v1beta1_SELinuxStrategyOptions is an autogenerated deepcopy function.
func DeepCopy_v1beta1_SELinuxStrategyOptions(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*SELinuxStrategyOptions)
		out := out.(*SELinuxStrategyOptions)
		*out = *in
		if in.SELinuxOptions != nil {
			in, out := &in.SELinuxOptions, &out.SELinuxOptions
			*out = new(api_v1.SELinuxOptions)
			**out = **in
		}
		return nil
	}
}

// DeepCopy_v1beta1_Scale is an autogenerated deepcopy function.
func DeepCopy_v1beta1_Scale(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*Scale)
		out := out.(*Scale)
		*out = *in
		if newVal, err := c.DeepCopy(&in.ObjectMeta); err != nil {
			return err
		} else {
			out.ObjectMeta = *newVal.(*v1.ObjectMeta)
		}
		if err := DeepCopy_v1beta1_ScaleStatus(&in.Status, &out.Status, c); err != nil {
			return err
		}
		return nil
	}
}

// DeepCopy_v1beta1_ScaleSpec is an autogenerated deepcopy function.
func DeepCopy_v1beta1_ScaleSpec(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*ScaleSpec)
		out := out.(*ScaleSpec)
		*out = *in
		return nil
	}
}

// DeepCopy_v1beta1_ScaleStatus is an autogenerated deepcopy function.
func DeepCopy_v1beta1_ScaleStatus(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*ScaleStatus)
		out := out.(*ScaleStatus)
		*out = *in
		if in.Selector != nil {
			in, out := &in.Selector, &out.Selector
			*out = make(map[string]string)
			for key, val := range *in {
				(*out)[key] = val
			}
		}
		return nil
	}
}

// DeepCopy_v1beta1_SupplementalGroupsStrategyOptions is an autogenerated deepcopy function.
func DeepCopy_v1beta1_SupplementalGroupsStrategyOptions(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*SupplementalGroupsStrategyOptions)
		out := out.(*SupplementalGroupsStrategyOptions)
		*out = *in
		if in.Ranges != nil {
			in, out := &in.Ranges, &out.Ranges
			*out = make([]IDRange, len(*in))
			copy(*out, *in)
		}
		return nil
	}
}

// DeepCopy_v1beta1_ThirdPartyResource is an autogenerated deepcopy function.
func DeepCopy_v1beta1_ThirdPartyResource(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*ThirdPartyResource)
		out := out.(*ThirdPartyResource)
		*out = *in
		if newVal, err := c.DeepCopy(&in.ObjectMeta); err != nil {
			return err
		} else {
			out.ObjectMeta = *newVal.(*v1.ObjectMeta)
		}
		if in.Versions != nil {
			in, out := &in.Versions, &out.Versions
			*out = make([]APIVersion, len(*in))
			copy(*out, *in)
		}
		return nil
	}
}

// DeepCopy_v1beta1_ThirdPartyResourceData is an autogenerated deepcopy function.
func DeepCopy_v1beta1_ThirdPartyResourceData(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*ThirdPartyResourceData)
		out := out.(*ThirdPartyResourceData)
		*out = *in
		if newVal, err := c.DeepCopy(&in.ObjectMeta); err != nil {
			return err
		} else {
			out.ObjectMeta = *newVal.(*v1.ObjectMeta)
		}
		if in.Data != nil {
			in, out := &in.Data, &out.Data
			*out = make([]byte, len(*in))
			copy(*out, *in)
		}
		return nil
	}
}

// DeepCopy_v1beta1_ThirdPartyResourceDataList is an autogenerated deepcopy function.
func DeepCopy_v1beta1_ThirdPartyResourceDataList(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*ThirdPartyResourceDataList)
		out := out.(*ThirdPartyResourceDataList)
		*out = *in
		if in.Items != nil {
			in, out := &in.Items, &out.Items
			*out = make([]ThirdPartyResourceData, len(*in))
			for i := range *in {
				if err := DeepCopy_v1beta1_ThirdPartyResourceData(&(*in)[i], &(*out)[i], c); err != nil {
					return err
				}
			}
		}
		return nil
	}
}

// DeepCopy_v1beta1_ThirdPartyResourceList is an autogenerated deepcopy function.
func DeepCopy_v1beta1_ThirdPartyResourceList(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*ThirdPartyResourceList)
		out := out.(*ThirdPartyResourceList)
		*out = *in
		if in.Items != nil {
			in, out := &in.Items, &out.Items
			*out = make([]ThirdPartyResource, len(*in))
			for i := range *in {
				if err := DeepCopy_v1beta1_ThirdPartyResource(&(*in)[i], &(*out)[i], c); err != nil {
					return err
				}
			}
		}
		return nil
	}
}
