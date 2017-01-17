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

// This file was autogenerated by conversion-gen. Do not edit it manually!

package v1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	conversion "k8s.io/apimachinery/pkg/conversion"
	runtime "k8s.io/apimachinery/pkg/runtime"
	autoscaling "k8s.io/client-go/pkg/apis/autoscaling"
	unsafe "unsafe"
)

func init() {
	SchemeBuilder.Register(RegisterConversions)
}

// RegisterConversions adds conversion functions to the given scheme.
// Public to allow building arbitrary schemes.
func RegisterConversions(scheme *runtime.Scheme) error {
	return scheme.AddGeneratedConversionFuncs(
		Convert_v1_CrossVersionObjectReference_To_autoscaling_CrossVersionObjectReference,
		Convert_autoscaling_CrossVersionObjectReference_To_v1_CrossVersionObjectReference,
		Convert_v1_HorizontalPodAutoscaler_To_autoscaling_HorizontalPodAutoscaler,
		Convert_autoscaling_HorizontalPodAutoscaler_To_v1_HorizontalPodAutoscaler,
		Convert_v1_HorizontalPodAutoscalerList_To_autoscaling_HorizontalPodAutoscalerList,
		Convert_autoscaling_HorizontalPodAutoscalerList_To_v1_HorizontalPodAutoscalerList,
		Convert_v1_HorizontalPodAutoscalerSpec_To_autoscaling_HorizontalPodAutoscalerSpec,
		Convert_autoscaling_HorizontalPodAutoscalerSpec_To_v1_HorizontalPodAutoscalerSpec,
		Convert_v1_HorizontalPodAutoscalerStatus_To_autoscaling_HorizontalPodAutoscalerStatus,
		Convert_autoscaling_HorizontalPodAutoscalerStatus_To_v1_HorizontalPodAutoscalerStatus,
		Convert_v1_Scale_To_autoscaling_Scale,
		Convert_autoscaling_Scale_To_v1_Scale,
		Convert_v1_ScaleSpec_To_autoscaling_ScaleSpec,
		Convert_autoscaling_ScaleSpec_To_v1_ScaleSpec,
		Convert_v1_ScaleStatus_To_autoscaling_ScaleStatus,
		Convert_autoscaling_ScaleStatus_To_v1_ScaleStatus,
	)
}

func autoConvert_v1_CrossVersionObjectReference_To_autoscaling_CrossVersionObjectReference(in *CrossVersionObjectReference, out *autoscaling.CrossVersionObjectReference, s conversion.Scope) error {
	out.Kind = in.Kind
	out.Name = in.Name
	out.APIVersion = in.APIVersion
	return nil
}

func Convert_v1_CrossVersionObjectReference_To_autoscaling_CrossVersionObjectReference(in *CrossVersionObjectReference, out *autoscaling.CrossVersionObjectReference, s conversion.Scope) error {
	return autoConvert_v1_CrossVersionObjectReference_To_autoscaling_CrossVersionObjectReference(in, out, s)
}

func autoConvert_autoscaling_CrossVersionObjectReference_To_v1_CrossVersionObjectReference(in *autoscaling.CrossVersionObjectReference, out *CrossVersionObjectReference, s conversion.Scope) error {
	out.Kind = in.Kind
	out.Name = in.Name
	out.APIVersion = in.APIVersion
	return nil
}

func Convert_autoscaling_CrossVersionObjectReference_To_v1_CrossVersionObjectReference(in *autoscaling.CrossVersionObjectReference, out *CrossVersionObjectReference, s conversion.Scope) error {
	return autoConvert_autoscaling_CrossVersionObjectReference_To_v1_CrossVersionObjectReference(in, out, s)
}

func autoConvert_v1_HorizontalPodAutoscaler_To_autoscaling_HorizontalPodAutoscaler(in *HorizontalPodAutoscaler, out *autoscaling.HorizontalPodAutoscaler, s conversion.Scope) error {
	out.ObjectMeta = in.ObjectMeta
	if err := Convert_v1_HorizontalPodAutoscalerSpec_To_autoscaling_HorizontalPodAutoscalerSpec(&in.Spec, &out.Spec, s); err != nil {
		return err
	}
	if err := Convert_v1_HorizontalPodAutoscalerStatus_To_autoscaling_HorizontalPodAutoscalerStatus(&in.Status, &out.Status, s); err != nil {
		return err
	}
	return nil
}

func Convert_v1_HorizontalPodAutoscaler_To_autoscaling_HorizontalPodAutoscaler(in *HorizontalPodAutoscaler, out *autoscaling.HorizontalPodAutoscaler, s conversion.Scope) error {
	return autoConvert_v1_HorizontalPodAutoscaler_To_autoscaling_HorizontalPodAutoscaler(in, out, s)
}

func autoConvert_autoscaling_HorizontalPodAutoscaler_To_v1_HorizontalPodAutoscaler(in *autoscaling.HorizontalPodAutoscaler, out *HorizontalPodAutoscaler, s conversion.Scope) error {
	out.ObjectMeta = in.ObjectMeta
	if err := Convert_autoscaling_HorizontalPodAutoscalerSpec_To_v1_HorizontalPodAutoscalerSpec(&in.Spec, &out.Spec, s); err != nil {
		return err
	}
	if err := Convert_autoscaling_HorizontalPodAutoscalerStatus_To_v1_HorizontalPodAutoscalerStatus(&in.Status, &out.Status, s); err != nil {
		return err
	}
	return nil
}

func Convert_autoscaling_HorizontalPodAutoscaler_To_v1_HorizontalPodAutoscaler(in *autoscaling.HorizontalPodAutoscaler, out *HorizontalPodAutoscaler, s conversion.Scope) error {
	return autoConvert_autoscaling_HorizontalPodAutoscaler_To_v1_HorizontalPodAutoscaler(in, out, s)
}

func autoConvert_v1_HorizontalPodAutoscalerList_To_autoscaling_HorizontalPodAutoscalerList(in *HorizontalPodAutoscalerList, out *autoscaling.HorizontalPodAutoscalerList, s conversion.Scope) error {
	out.ListMeta = in.ListMeta
	out.Items = *(*[]autoscaling.HorizontalPodAutoscaler)(unsafe.Pointer(&in.Items))
	return nil
}

func Convert_v1_HorizontalPodAutoscalerList_To_autoscaling_HorizontalPodAutoscalerList(in *HorizontalPodAutoscalerList, out *autoscaling.HorizontalPodAutoscalerList, s conversion.Scope) error {
	return autoConvert_v1_HorizontalPodAutoscalerList_To_autoscaling_HorizontalPodAutoscalerList(in, out, s)
}

func autoConvert_autoscaling_HorizontalPodAutoscalerList_To_v1_HorizontalPodAutoscalerList(in *autoscaling.HorizontalPodAutoscalerList, out *HorizontalPodAutoscalerList, s conversion.Scope) error {
	out.ListMeta = in.ListMeta
	out.Items = *(*[]HorizontalPodAutoscaler)(unsafe.Pointer(&in.Items))
	return nil
}

func Convert_autoscaling_HorizontalPodAutoscalerList_To_v1_HorizontalPodAutoscalerList(in *autoscaling.HorizontalPodAutoscalerList, out *HorizontalPodAutoscalerList, s conversion.Scope) error {
	return autoConvert_autoscaling_HorizontalPodAutoscalerList_To_v1_HorizontalPodAutoscalerList(in, out, s)
}

func autoConvert_v1_HorizontalPodAutoscalerSpec_To_autoscaling_HorizontalPodAutoscalerSpec(in *HorizontalPodAutoscalerSpec, out *autoscaling.HorizontalPodAutoscalerSpec, s conversion.Scope) error {
	if err := Convert_v1_CrossVersionObjectReference_To_autoscaling_CrossVersionObjectReference(&in.ScaleTargetRef, &out.ScaleTargetRef, s); err != nil {
		return err
	}
	out.MinReplicas = (*int32)(unsafe.Pointer(in.MinReplicas))
	out.MaxReplicas = in.MaxReplicas
	out.TargetCPUUtilizationPercentage = (*int32)(unsafe.Pointer(in.TargetCPUUtilizationPercentage))
	return nil
}

func Convert_v1_HorizontalPodAutoscalerSpec_To_autoscaling_HorizontalPodAutoscalerSpec(in *HorizontalPodAutoscalerSpec, out *autoscaling.HorizontalPodAutoscalerSpec, s conversion.Scope) error {
	return autoConvert_v1_HorizontalPodAutoscalerSpec_To_autoscaling_HorizontalPodAutoscalerSpec(in, out, s)
}

func autoConvert_autoscaling_HorizontalPodAutoscalerSpec_To_v1_HorizontalPodAutoscalerSpec(in *autoscaling.HorizontalPodAutoscalerSpec, out *HorizontalPodAutoscalerSpec, s conversion.Scope) error {
	if err := Convert_autoscaling_CrossVersionObjectReference_To_v1_CrossVersionObjectReference(&in.ScaleTargetRef, &out.ScaleTargetRef, s); err != nil {
		return err
	}
	out.MinReplicas = (*int32)(unsafe.Pointer(in.MinReplicas))
	out.MaxReplicas = in.MaxReplicas
	out.TargetCPUUtilizationPercentage = (*int32)(unsafe.Pointer(in.TargetCPUUtilizationPercentage))
	return nil
}

func Convert_autoscaling_HorizontalPodAutoscalerSpec_To_v1_HorizontalPodAutoscalerSpec(in *autoscaling.HorizontalPodAutoscalerSpec, out *HorizontalPodAutoscalerSpec, s conversion.Scope) error {
	return autoConvert_autoscaling_HorizontalPodAutoscalerSpec_To_v1_HorizontalPodAutoscalerSpec(in, out, s)
}

func autoConvert_v1_HorizontalPodAutoscalerStatus_To_autoscaling_HorizontalPodAutoscalerStatus(in *HorizontalPodAutoscalerStatus, out *autoscaling.HorizontalPodAutoscalerStatus, s conversion.Scope) error {
	out.ObservedGeneration = (*int64)(unsafe.Pointer(in.ObservedGeneration))
	out.LastScaleTime = (*meta_v1.Time)(unsafe.Pointer(in.LastScaleTime))
	out.CurrentReplicas = in.CurrentReplicas
	out.DesiredReplicas = in.DesiredReplicas
	out.CurrentCPUUtilizationPercentage = (*int32)(unsafe.Pointer(in.CurrentCPUUtilizationPercentage))
	return nil
}

func Convert_v1_HorizontalPodAutoscalerStatus_To_autoscaling_HorizontalPodAutoscalerStatus(in *HorizontalPodAutoscalerStatus, out *autoscaling.HorizontalPodAutoscalerStatus, s conversion.Scope) error {
	return autoConvert_v1_HorizontalPodAutoscalerStatus_To_autoscaling_HorizontalPodAutoscalerStatus(in, out, s)
}

func autoConvert_autoscaling_HorizontalPodAutoscalerStatus_To_v1_HorizontalPodAutoscalerStatus(in *autoscaling.HorizontalPodAutoscalerStatus, out *HorizontalPodAutoscalerStatus, s conversion.Scope) error {
	out.ObservedGeneration = (*int64)(unsafe.Pointer(in.ObservedGeneration))
	out.LastScaleTime = (*meta_v1.Time)(unsafe.Pointer(in.LastScaleTime))
	out.CurrentReplicas = in.CurrentReplicas
	out.DesiredReplicas = in.DesiredReplicas
	out.CurrentCPUUtilizationPercentage = (*int32)(unsafe.Pointer(in.CurrentCPUUtilizationPercentage))
	return nil
}

func Convert_autoscaling_HorizontalPodAutoscalerStatus_To_v1_HorizontalPodAutoscalerStatus(in *autoscaling.HorizontalPodAutoscalerStatus, out *HorizontalPodAutoscalerStatus, s conversion.Scope) error {
	return autoConvert_autoscaling_HorizontalPodAutoscalerStatus_To_v1_HorizontalPodAutoscalerStatus(in, out, s)
}

func autoConvert_v1_Scale_To_autoscaling_Scale(in *Scale, out *autoscaling.Scale, s conversion.Scope) error {
	out.ObjectMeta = in.ObjectMeta
	if err := Convert_v1_ScaleSpec_To_autoscaling_ScaleSpec(&in.Spec, &out.Spec, s); err != nil {
		return err
	}
	if err := Convert_v1_ScaleStatus_To_autoscaling_ScaleStatus(&in.Status, &out.Status, s); err != nil {
		return err
	}
	return nil
}

func Convert_v1_Scale_To_autoscaling_Scale(in *Scale, out *autoscaling.Scale, s conversion.Scope) error {
	return autoConvert_v1_Scale_To_autoscaling_Scale(in, out, s)
}

func autoConvert_autoscaling_Scale_To_v1_Scale(in *autoscaling.Scale, out *Scale, s conversion.Scope) error {
	out.ObjectMeta = in.ObjectMeta
	if err := Convert_autoscaling_ScaleSpec_To_v1_ScaleSpec(&in.Spec, &out.Spec, s); err != nil {
		return err
	}
	if err := Convert_autoscaling_ScaleStatus_To_v1_ScaleStatus(&in.Status, &out.Status, s); err != nil {
		return err
	}
	return nil
}

func Convert_autoscaling_Scale_To_v1_Scale(in *autoscaling.Scale, out *Scale, s conversion.Scope) error {
	return autoConvert_autoscaling_Scale_To_v1_Scale(in, out, s)
}

func autoConvert_v1_ScaleSpec_To_autoscaling_ScaleSpec(in *ScaleSpec, out *autoscaling.ScaleSpec, s conversion.Scope) error {
	out.Replicas = in.Replicas
	return nil
}

func Convert_v1_ScaleSpec_To_autoscaling_ScaleSpec(in *ScaleSpec, out *autoscaling.ScaleSpec, s conversion.Scope) error {
	return autoConvert_v1_ScaleSpec_To_autoscaling_ScaleSpec(in, out, s)
}

func autoConvert_autoscaling_ScaleSpec_To_v1_ScaleSpec(in *autoscaling.ScaleSpec, out *ScaleSpec, s conversion.Scope) error {
	out.Replicas = in.Replicas
	return nil
}

func Convert_autoscaling_ScaleSpec_To_v1_ScaleSpec(in *autoscaling.ScaleSpec, out *ScaleSpec, s conversion.Scope) error {
	return autoConvert_autoscaling_ScaleSpec_To_v1_ScaleSpec(in, out, s)
}

func autoConvert_v1_ScaleStatus_To_autoscaling_ScaleStatus(in *ScaleStatus, out *autoscaling.ScaleStatus, s conversion.Scope) error {
	out.Replicas = in.Replicas
	out.Selector = in.Selector
	return nil
}

func Convert_v1_ScaleStatus_To_autoscaling_ScaleStatus(in *ScaleStatus, out *autoscaling.ScaleStatus, s conversion.Scope) error {
	return autoConvert_v1_ScaleStatus_To_autoscaling_ScaleStatus(in, out, s)
}

func autoConvert_autoscaling_ScaleStatus_To_v1_ScaleStatus(in *autoscaling.ScaleStatus, out *ScaleStatus, s conversion.Scope) error {
	out.Replicas = in.Replicas
	out.Selector = in.Selector
	return nil
}

func Convert_autoscaling_ScaleStatus_To_v1_ScaleStatus(in *autoscaling.ScaleStatus, out *ScaleStatus, s conversion.Scope) error {
	return autoConvert_autoscaling_ScaleStatus_To_v1_ScaleStatus(in, out, s)
}
