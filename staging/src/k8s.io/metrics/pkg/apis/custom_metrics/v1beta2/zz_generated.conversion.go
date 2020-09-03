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

package v1beta2

import (
	unsafe "unsafe"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
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
	if err := s.AddGeneratedConversionFunc((*MetricIdentifier)(nil), (*custommetrics.MetricIdentifier)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1beta2_MetricIdentifier_To_custom_metrics_MetricIdentifier(a.(*MetricIdentifier), b.(*custommetrics.MetricIdentifier), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*custommetrics.MetricIdentifier)(nil), (*MetricIdentifier)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_custom_metrics_MetricIdentifier_To_v1beta2_MetricIdentifier(a.(*custommetrics.MetricIdentifier), b.(*MetricIdentifier), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*MetricListOptions)(nil), (*custommetrics.MetricListOptions)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1beta2_MetricListOptions_To_custom_metrics_MetricListOptions(a.(*MetricListOptions), b.(*custommetrics.MetricListOptions), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*custommetrics.MetricListOptions)(nil), (*MetricListOptions)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_custom_metrics_MetricListOptions_To_v1beta2_MetricListOptions(a.(*custommetrics.MetricListOptions), b.(*MetricListOptions), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*MetricValue)(nil), (*custommetrics.MetricValue)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1beta2_MetricValue_To_custom_metrics_MetricValue(a.(*MetricValue), b.(*custommetrics.MetricValue), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*custommetrics.MetricValue)(nil), (*MetricValue)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_custom_metrics_MetricValue_To_v1beta2_MetricValue(a.(*custommetrics.MetricValue), b.(*MetricValue), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*MetricValueList)(nil), (*custommetrics.MetricValueList)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1beta2_MetricValueList_To_custom_metrics_MetricValueList(a.(*MetricValueList), b.(*custommetrics.MetricValueList), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*custommetrics.MetricValueList)(nil), (*MetricValueList)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_custom_metrics_MetricValueList_To_v1beta2_MetricValueList(a.(*custommetrics.MetricValueList), b.(*MetricValueList), scope)
	}); err != nil {
		return err
	}
	return nil
}

func autoConvert_v1beta2_MetricIdentifier_To_custom_metrics_MetricIdentifier(in *MetricIdentifier, out *custommetrics.MetricIdentifier, s conversion.Scope) error {
	out.Name = in.Name
	out.Selector = (*v1.LabelSelector)(unsafe.Pointer(in.Selector))
	return nil
}

// Convert_v1beta2_MetricIdentifier_To_custom_metrics_MetricIdentifier is an autogenerated conversion function.
func Convert_v1beta2_MetricIdentifier_To_custom_metrics_MetricIdentifier(in *MetricIdentifier, out *custommetrics.MetricIdentifier, s conversion.Scope) error {
	return autoConvert_v1beta2_MetricIdentifier_To_custom_metrics_MetricIdentifier(in, out, s)
}

func autoConvert_custom_metrics_MetricIdentifier_To_v1beta2_MetricIdentifier(in *custommetrics.MetricIdentifier, out *MetricIdentifier, s conversion.Scope) error {
	out.Name = in.Name
	out.Selector = (*v1.LabelSelector)(unsafe.Pointer(in.Selector))
	return nil
}

// Convert_custom_metrics_MetricIdentifier_To_v1beta2_MetricIdentifier is an autogenerated conversion function.
func Convert_custom_metrics_MetricIdentifier_To_v1beta2_MetricIdentifier(in *custommetrics.MetricIdentifier, out *MetricIdentifier, s conversion.Scope) error {
	return autoConvert_custom_metrics_MetricIdentifier_To_v1beta2_MetricIdentifier(in, out, s)
}

func autoConvert_v1beta2_MetricListOptions_To_custom_metrics_MetricListOptions(in *MetricListOptions, out *custommetrics.MetricListOptions, s conversion.Scope) error {
	out.LabelSelector = in.LabelSelector
	out.MetricLabelSelector = in.MetricLabelSelector
	return nil
}

// Convert_v1beta2_MetricListOptions_To_custom_metrics_MetricListOptions is an autogenerated conversion function.
func Convert_v1beta2_MetricListOptions_To_custom_metrics_MetricListOptions(in *MetricListOptions, out *custommetrics.MetricListOptions, s conversion.Scope) error {
	return autoConvert_v1beta2_MetricListOptions_To_custom_metrics_MetricListOptions(in, out, s)
}

func autoConvert_custom_metrics_MetricListOptions_To_v1beta2_MetricListOptions(in *custommetrics.MetricListOptions, out *MetricListOptions, s conversion.Scope) error {
	out.LabelSelector = in.LabelSelector
	out.MetricLabelSelector = in.MetricLabelSelector
	return nil
}

// Convert_custom_metrics_MetricListOptions_To_v1beta2_MetricListOptions is an autogenerated conversion function.
func Convert_custom_metrics_MetricListOptions_To_v1beta2_MetricListOptions(in *custommetrics.MetricListOptions, out *MetricListOptions, s conversion.Scope) error {
	return autoConvert_custom_metrics_MetricListOptions_To_v1beta2_MetricListOptions(in, out, s)
}

func autoConvert_v1beta2_MetricValue_To_custom_metrics_MetricValue(in *MetricValue, out *custommetrics.MetricValue, s conversion.Scope) error {
	if err := custommetrics.Convert_v1_ObjectReference_To_custom_metrics_ObjectReference(&in.DescribedObject, &out.DescribedObject, s); err != nil {
		return err
	}
	if err := Convert_v1beta2_MetricIdentifier_To_custom_metrics_MetricIdentifier(&in.Metric, &out.Metric, s); err != nil {
		return err
	}
	out.Timestamp = in.Timestamp
	out.WindowSeconds = (*int64)(unsafe.Pointer(in.WindowSeconds))
	out.Value = in.Value
	return nil
}

// Convert_v1beta2_MetricValue_To_custom_metrics_MetricValue is an autogenerated conversion function.
func Convert_v1beta2_MetricValue_To_custom_metrics_MetricValue(in *MetricValue, out *custommetrics.MetricValue, s conversion.Scope) error {
	return autoConvert_v1beta2_MetricValue_To_custom_metrics_MetricValue(in, out, s)
}

func autoConvert_custom_metrics_MetricValue_To_v1beta2_MetricValue(in *custommetrics.MetricValue, out *MetricValue, s conversion.Scope) error {
	if err := custommetrics.Convert_custom_metrics_ObjectReference_To_v1_ObjectReference(&in.DescribedObject, &out.DescribedObject, s); err != nil {
		return err
	}
	if err := Convert_custom_metrics_MetricIdentifier_To_v1beta2_MetricIdentifier(&in.Metric, &out.Metric, s); err != nil {
		return err
	}
	out.Timestamp = in.Timestamp
	out.WindowSeconds = (*int64)(unsafe.Pointer(in.WindowSeconds))
	out.Value = in.Value
	return nil
}

// Convert_custom_metrics_MetricValue_To_v1beta2_MetricValue is an autogenerated conversion function.
func Convert_custom_metrics_MetricValue_To_v1beta2_MetricValue(in *custommetrics.MetricValue, out *MetricValue, s conversion.Scope) error {
	return autoConvert_custom_metrics_MetricValue_To_v1beta2_MetricValue(in, out, s)
}

func autoConvert_v1beta2_MetricValueList_To_custom_metrics_MetricValueList(in *MetricValueList, out *custommetrics.MetricValueList, s conversion.Scope) error {
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]custommetrics.MetricValue, len(*in))
		for i := range *in {
			if err := Convert_v1beta2_MetricValue_To_custom_metrics_MetricValue(&(*in)[i], &(*out)[i], s); err != nil {
				return err
			}
		}
	} else {
		out.Items = nil
	}
	return nil
}

// Convert_v1beta2_MetricValueList_To_custom_metrics_MetricValueList is an autogenerated conversion function.
func Convert_v1beta2_MetricValueList_To_custom_metrics_MetricValueList(in *MetricValueList, out *custommetrics.MetricValueList, s conversion.Scope) error {
	return autoConvert_v1beta2_MetricValueList_To_custom_metrics_MetricValueList(in, out, s)
}

func autoConvert_custom_metrics_MetricValueList_To_v1beta2_MetricValueList(in *custommetrics.MetricValueList, out *MetricValueList, s conversion.Scope) error {
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]MetricValue, len(*in))
		for i := range *in {
			if err := Convert_custom_metrics_MetricValue_To_v1beta2_MetricValue(&(*in)[i], &(*out)[i], s); err != nil {
				return err
			}
		}
	} else {
		out.Items = nil
	}
	return nil
}

// Convert_custom_metrics_MetricValueList_To_v1beta2_MetricValueList is an autogenerated conversion function.
func Convert_custom_metrics_MetricValueList_To_v1beta2_MetricValueList(in *custommetrics.MetricValueList, out *MetricValueList, s conversion.Scope) error {
	return autoConvert_custom_metrics_MetricValueList_To_v1beta2_MetricValueList(in, out, s)
}
