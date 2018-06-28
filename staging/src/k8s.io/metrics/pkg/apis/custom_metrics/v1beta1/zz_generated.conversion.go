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

// Code generated by conversion-gen. DO NOT EDIT.

package v1beta1

import (
	unsafe "unsafe"

	conversion "k8s.io/apimachinery/pkg/conversion"
	runtime "k8s.io/apimachinery/pkg/runtime"
	custommetrics "k8s.io/metrics/pkg/apis/custom_metrics"
)

func init() {
	localSchemeBuilder.Register(RegisterConversions)
}

// RegisterConversions adds conversion functions to the given scheme.
// Public to allow building arbitrary schemes.
func RegisterConversions(s *runtime.Scheme) error {
	if err := s.AddGeneratedConversionFunc((*MetricValue)(nil), (*custommetrics.MetricValue)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1beta1_MetricValue_To_custom_metrics_MetricValue(a.(*MetricValue), b.(*custommetrics.MetricValue), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*custommetrics.MetricValue)(nil), (*MetricValue)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_custom_metrics_MetricValue_To_v1beta1_MetricValue(a.(*custommetrics.MetricValue), b.(*MetricValue), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*MetricValueList)(nil), (*custommetrics.MetricValueList)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1beta1_MetricValueList_To_custom_metrics_MetricValueList(a.(*MetricValueList), b.(*custommetrics.MetricValueList), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*custommetrics.MetricValueList)(nil), (*MetricValueList)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_custom_metrics_MetricValueList_To_v1beta1_MetricValueList(a.(*custommetrics.MetricValueList), b.(*MetricValueList), scope)
	}); err != nil {
		return err
	}
	return nil
}

func autoConvert_v1beta1_MetricValue_To_custom_metrics_MetricValue(in *MetricValue, out *custommetrics.MetricValue, s conversion.Scope) error {
	// TODO: Inefficient conversion - can we improve it?
	if err := s.Convert(&in.DescribedObject, &out.DescribedObject, 0); err != nil {
		return err
	}
	// WARNING: in.MetricName requires manual conversion: does not exist in peer-type
	out.Timestamp = in.Timestamp
	out.WindowSeconds = (*int64)(unsafe.Pointer(in.WindowSeconds))
	out.Value = in.Value
	return nil
}

<<<<<<< HEAD
// Convert_v1beta1_MetricValue_To_custom_metrics_MetricValue is an autogenerated conversion function.
func Convert_v1beta1_MetricValue_To_custom_metrics_MetricValue(in *MetricValue, out *custommetrics.MetricValue, s conversion.Scope) error {
	return autoConvert_v1beta1_MetricValue_To_custom_metrics_MetricValue(in, out, s)
}

func autoConvert_custom_metrics_MetricValue_To_v1beta1_MetricValue(in *custommetrics.MetricValue, out *MetricValue, s conversion.Scope) error {
=======
func autoConvert_custom_metrics_MetricValue_To_v1beta1_MetricValue(in *custom_metrics.MetricValue, out *MetricValue, s conversion.Scope) error {
>>>>>>> Update metrics API to include autoscaling/v2beta2 changes
	// TODO: Inefficient conversion - can we improve it?
	if err := s.Convert(&in.DescribedObject, &out.DescribedObject, 0); err != nil {
		return err
	}
	// WARNING: in.Metric requires manual conversion: does not exist in peer-type
	out.Timestamp = in.Timestamp
	out.WindowSeconds = (*int64)(unsafe.Pointer(in.WindowSeconds))
	out.Value = in.Value
	return nil
}

<<<<<<< HEAD
// Convert_custom_metrics_MetricValue_To_v1beta1_MetricValue is an autogenerated conversion function.
func Convert_custom_metrics_MetricValue_To_v1beta1_MetricValue(in *custommetrics.MetricValue, out *MetricValue, s conversion.Scope) error {
	return autoConvert_custom_metrics_MetricValue_To_v1beta1_MetricValue(in, out, s)
}

func autoConvert_v1beta1_MetricValueList_To_custom_metrics_MetricValueList(in *MetricValueList, out *custommetrics.MetricValueList, s conversion.Scope) error {
	out.ListMeta = in.ListMeta
	out.Items = *(*[]custommetrics.MetricValue)(unsafe.Pointer(&in.Items))
=======
func autoConvert_v1beta1_MetricValueList_To_custom_metrics_MetricValueList(in *MetricValueList, out *custom_metrics.MetricValueList, s conversion.Scope) error {
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]custom_metrics.MetricValue, len(*in))
		for i := range *in {
			if err := Convert_v1beta1_MetricValue_To_custom_metrics_MetricValue(&(*in)[i], &(*out)[i], s); err != nil {
				return err
			}
		}
	} else {
		out.Items = nil
	}
>>>>>>> Update metrics API to include autoscaling/v2beta2 changes
	return nil
}

// Convert_v1beta1_MetricValueList_To_custom_metrics_MetricValueList is an autogenerated conversion function.
func Convert_v1beta1_MetricValueList_To_custom_metrics_MetricValueList(in *MetricValueList, out *custommetrics.MetricValueList, s conversion.Scope) error {
	return autoConvert_v1beta1_MetricValueList_To_custom_metrics_MetricValueList(in, out, s)
}

func autoConvert_custom_metrics_MetricValueList_To_v1beta1_MetricValueList(in *custommetrics.MetricValueList, out *MetricValueList, s conversion.Scope) error {
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]MetricValue, len(*in))
		for i := range *in {
			if err := Convert_custom_metrics_MetricValue_To_v1beta1_MetricValue(&(*in)[i], &(*out)[i], s); err != nil {
				return err
			}
		}
	} else {
		out.Items = nil
	}
	return nil
}

// Convert_custom_metrics_MetricValueList_To_v1beta1_MetricValueList is an autogenerated conversion function.
func Convert_custom_metrics_MetricValueList_To_v1beta1_MetricValueList(in *custommetrics.MetricValueList, out *MetricValueList, s conversion.Scope) error {
	return autoConvert_custom_metrics_MetricValueList_To_v1beta1_MetricValueList(in, out, s)
}
