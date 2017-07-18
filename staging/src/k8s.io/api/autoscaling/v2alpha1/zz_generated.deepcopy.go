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

package v2alpha1

import (
	resource "k8s.io/apimachinery/pkg/api/resource"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	conversion "k8s.io/apimachinery/pkg/conversion"
	runtime "k8s.io/apimachinery/pkg/runtime"
	reflect "reflect"
)

// Deprecated: register deep-copy functions.
func init() {
	SchemeBuilder.Register(RegisterDeepCopies)
}

// Deprecated: RegisterDeepCopies adds deep-copy functions to the given scheme. Public
// to allow building arbitrary schemes.
func RegisterDeepCopies(scheme *runtime.Scheme) error {
	return scheme.AddGeneratedDeepCopyFuncs(
		conversion.GeneratedDeepCopyFunc{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*CrossVersionObjectReference).DeepCopyInto(out.(*CrossVersionObjectReference))
			return nil
		}, InType: reflect.TypeOf(&CrossVersionObjectReference{})},
		conversion.GeneratedDeepCopyFunc{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*HorizontalPodAutoscaler).DeepCopyInto(out.(*HorizontalPodAutoscaler))
			return nil
		}, InType: reflect.TypeOf(&HorizontalPodAutoscaler{})},
		conversion.GeneratedDeepCopyFunc{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*HorizontalPodAutoscalerCondition).DeepCopyInto(out.(*HorizontalPodAutoscalerCondition))
			return nil
		}, InType: reflect.TypeOf(&HorizontalPodAutoscalerCondition{})},
		conversion.GeneratedDeepCopyFunc{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*HorizontalPodAutoscalerList).DeepCopyInto(out.(*HorizontalPodAutoscalerList))
			return nil
		}, InType: reflect.TypeOf(&HorizontalPodAutoscalerList{})},
		conversion.GeneratedDeepCopyFunc{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*HorizontalPodAutoscalerSpec).DeepCopyInto(out.(*HorizontalPodAutoscalerSpec))
			return nil
		}, InType: reflect.TypeOf(&HorizontalPodAutoscalerSpec{})},
		conversion.GeneratedDeepCopyFunc{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*HorizontalPodAutoscalerStatus).DeepCopyInto(out.(*HorizontalPodAutoscalerStatus))
			return nil
		}, InType: reflect.TypeOf(&HorizontalPodAutoscalerStatus{})},
		conversion.GeneratedDeepCopyFunc{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*MetricSpec).DeepCopyInto(out.(*MetricSpec))
			return nil
		}, InType: reflect.TypeOf(&MetricSpec{})},
		conversion.GeneratedDeepCopyFunc{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*MetricStatus).DeepCopyInto(out.(*MetricStatus))
			return nil
		}, InType: reflect.TypeOf(&MetricStatus{})},
		conversion.GeneratedDeepCopyFunc{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*ObjectMetricSource).DeepCopyInto(out.(*ObjectMetricSource))
			return nil
		}, InType: reflect.TypeOf(&ObjectMetricSource{})},
		conversion.GeneratedDeepCopyFunc{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*ObjectMetricStatus).DeepCopyInto(out.(*ObjectMetricStatus))
			return nil
		}, InType: reflect.TypeOf(&ObjectMetricStatus{})},
		conversion.GeneratedDeepCopyFunc{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*PodsMetricSource).DeepCopyInto(out.(*PodsMetricSource))
			return nil
		}, InType: reflect.TypeOf(&PodsMetricSource{})},
		conversion.GeneratedDeepCopyFunc{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*PodsMetricStatus).DeepCopyInto(out.(*PodsMetricStatus))
			return nil
		}, InType: reflect.TypeOf(&PodsMetricStatus{})},
		conversion.GeneratedDeepCopyFunc{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*ResourceMetricSource).DeepCopyInto(out.(*ResourceMetricSource))
			return nil
		}, InType: reflect.TypeOf(&ResourceMetricSource{})},
		conversion.GeneratedDeepCopyFunc{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*ResourceMetricStatus).DeepCopyInto(out.(*ResourceMetricStatus))
			return nil
		}, InType: reflect.TypeOf(&ResourceMetricStatus{})},
	)
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CrossVersionObjectReference) DeepCopyInto(out *CrossVersionObjectReference) {
	*out = *in
	return
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, creating a new CrossVersionObjectReference.
func (x *CrossVersionObjectReference) DeepCopy() *CrossVersionObjectReference {
	if x == nil {
		return nil
	}
	out := new(CrossVersionObjectReference)
	x.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HorizontalPodAutoscaler) DeepCopyInto(out *HorizontalPodAutoscaler) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, creating a new HorizontalPodAutoscaler.
func (x *HorizontalPodAutoscaler) DeepCopy() *HorizontalPodAutoscaler {
	if x == nil {
		return nil
	}
	out := new(HorizontalPodAutoscaler)
	x.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (x *HorizontalPodAutoscaler) DeepCopyObject() runtime.Object {
	if c := x.DeepCopy(); c != nil {
		return c
	} else {
		return nil
	}
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HorizontalPodAutoscalerCondition) DeepCopyInto(out *HorizontalPodAutoscalerCondition) {
	*out = *in
	in.LastTransitionTime.DeepCopyInto(&out.LastTransitionTime)
	return
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, creating a new HorizontalPodAutoscalerCondition.
func (x *HorizontalPodAutoscalerCondition) DeepCopy() *HorizontalPodAutoscalerCondition {
	if x == nil {
		return nil
	}
	out := new(HorizontalPodAutoscalerCondition)
	x.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HorizontalPodAutoscalerList) DeepCopyInto(out *HorizontalPodAutoscalerList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]HorizontalPodAutoscaler, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, creating a new HorizontalPodAutoscalerList.
func (x *HorizontalPodAutoscalerList) DeepCopy() *HorizontalPodAutoscalerList {
	if x == nil {
		return nil
	}
	out := new(HorizontalPodAutoscalerList)
	x.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (x *HorizontalPodAutoscalerList) DeepCopyObject() runtime.Object {
	if c := x.DeepCopy(); c != nil {
		return c
	} else {
		return nil
	}
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HorizontalPodAutoscalerSpec) DeepCopyInto(out *HorizontalPodAutoscalerSpec) {
	*out = *in
	out.ScaleTargetRef = in.ScaleTargetRef
	if in.MinReplicas != nil {
		in, out := &in.MinReplicas, &out.MinReplicas
		if *in == nil {
			*out = nil
		} else {
			*out = new(int32)
			**out = **in
		}
	}
	if in.Metrics != nil {
		in, out := &in.Metrics, &out.Metrics
		*out = make([]MetricSpec, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, creating a new HorizontalPodAutoscalerSpec.
func (x *HorizontalPodAutoscalerSpec) DeepCopy() *HorizontalPodAutoscalerSpec {
	if x == nil {
		return nil
	}
	out := new(HorizontalPodAutoscalerSpec)
	x.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HorizontalPodAutoscalerStatus) DeepCopyInto(out *HorizontalPodAutoscalerStatus) {
	*out = *in
	if in.ObservedGeneration != nil {
		in, out := &in.ObservedGeneration, &out.ObservedGeneration
		if *in == nil {
			*out = nil
		} else {
			*out = new(int64)
			**out = **in
		}
	}
	if in.LastScaleTime != nil {
		in, out := &in.LastScaleTime, &out.LastScaleTime
		if *in == nil {
			*out = nil
		} else {
			*out = new(v1.Time)
			(*in).DeepCopyInto(*out)
		}
	}
	if in.CurrentMetrics != nil {
		in, out := &in.CurrentMetrics, &out.CurrentMetrics
		*out = make([]MetricStatus, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]HorizontalPodAutoscalerCondition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, creating a new HorizontalPodAutoscalerStatus.
func (x *HorizontalPodAutoscalerStatus) DeepCopy() *HorizontalPodAutoscalerStatus {
	if x == nil {
		return nil
	}
	out := new(HorizontalPodAutoscalerStatus)
	x.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MetricSpec) DeepCopyInto(out *MetricSpec) {
	*out = *in
	if in.Object != nil {
		in, out := &in.Object, &out.Object
		if *in == nil {
			*out = nil
		} else {
			*out = new(ObjectMetricSource)
			(*in).DeepCopyInto(*out)
		}
	}
	if in.Pods != nil {
		in, out := &in.Pods, &out.Pods
		if *in == nil {
			*out = nil
		} else {
			*out = new(PodsMetricSource)
			(*in).DeepCopyInto(*out)
		}
	}
	if in.Resource != nil {
		in, out := &in.Resource, &out.Resource
		if *in == nil {
			*out = nil
		} else {
			*out = new(ResourceMetricSource)
			(*in).DeepCopyInto(*out)
		}
	}
	return
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, creating a new MetricSpec.
func (x *MetricSpec) DeepCopy() *MetricSpec {
	if x == nil {
		return nil
	}
	out := new(MetricSpec)
	x.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MetricStatus) DeepCopyInto(out *MetricStatus) {
	*out = *in
	if in.Object != nil {
		in, out := &in.Object, &out.Object
		if *in == nil {
			*out = nil
		} else {
			*out = new(ObjectMetricStatus)
			(*in).DeepCopyInto(*out)
		}
	}
	if in.Pods != nil {
		in, out := &in.Pods, &out.Pods
		if *in == nil {
			*out = nil
		} else {
			*out = new(PodsMetricStatus)
			(*in).DeepCopyInto(*out)
		}
	}
	if in.Resource != nil {
		in, out := &in.Resource, &out.Resource
		if *in == nil {
			*out = nil
		} else {
			*out = new(ResourceMetricStatus)
			(*in).DeepCopyInto(*out)
		}
	}
	return
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, creating a new MetricStatus.
func (x *MetricStatus) DeepCopy() *MetricStatus {
	if x == nil {
		return nil
	}
	out := new(MetricStatus)
	x.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ObjectMetricSource) DeepCopyInto(out *ObjectMetricSource) {
	*out = *in
	out.Target = in.Target
	out.TargetValue = in.TargetValue.DeepCopy()
	return
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, creating a new ObjectMetricSource.
func (x *ObjectMetricSource) DeepCopy() *ObjectMetricSource {
	if x == nil {
		return nil
	}
	out := new(ObjectMetricSource)
	x.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ObjectMetricStatus) DeepCopyInto(out *ObjectMetricStatus) {
	*out = *in
	out.Target = in.Target
	out.CurrentValue = in.CurrentValue.DeepCopy()
	return
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, creating a new ObjectMetricStatus.
func (x *ObjectMetricStatus) DeepCopy() *ObjectMetricStatus {
	if x == nil {
		return nil
	}
	out := new(ObjectMetricStatus)
	x.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PodsMetricSource) DeepCopyInto(out *PodsMetricSource) {
	*out = *in
	out.TargetAverageValue = in.TargetAverageValue.DeepCopy()
	return
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, creating a new PodsMetricSource.
func (x *PodsMetricSource) DeepCopy() *PodsMetricSource {
	if x == nil {
		return nil
	}
	out := new(PodsMetricSource)
	x.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PodsMetricStatus) DeepCopyInto(out *PodsMetricStatus) {
	*out = *in
	out.CurrentAverageValue = in.CurrentAverageValue.DeepCopy()
	return
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, creating a new PodsMetricStatus.
func (x *PodsMetricStatus) DeepCopy() *PodsMetricStatus {
	if x == nil {
		return nil
	}
	out := new(PodsMetricStatus)
	x.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ResourceMetricSource) DeepCopyInto(out *ResourceMetricSource) {
	*out = *in
	if in.TargetAverageUtilization != nil {
		in, out := &in.TargetAverageUtilization, &out.TargetAverageUtilization
		if *in == nil {
			*out = nil
		} else {
			*out = new(int32)
			**out = **in
		}
	}
	if in.TargetAverageValue != nil {
		in, out := &in.TargetAverageValue, &out.TargetAverageValue
		if *in == nil {
			*out = nil
		} else {
			*out = new(resource.Quantity)
			**out = (*in).DeepCopy()
		}
	}
	return
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, creating a new ResourceMetricSource.
func (x *ResourceMetricSource) DeepCopy() *ResourceMetricSource {
	if x == nil {
		return nil
	}
	out := new(ResourceMetricSource)
	x.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ResourceMetricStatus) DeepCopyInto(out *ResourceMetricStatus) {
	*out = *in
	if in.CurrentAverageUtilization != nil {
		in, out := &in.CurrentAverageUtilization, &out.CurrentAverageUtilization
		if *in == nil {
			*out = nil
		} else {
			*out = new(int32)
			**out = **in
		}
	}
	out.CurrentAverageValue = in.CurrentAverageValue.DeepCopy()
	return
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, creating a new ResourceMetricStatus.
func (x *ResourceMetricStatus) DeepCopy() *ResourceMetricStatus {
	if x == nil {
		return nil
	}
	out := new(ResourceMetricStatus)
	x.DeepCopyInto(out)
	return out
}
