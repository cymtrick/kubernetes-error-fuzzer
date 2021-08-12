//go:build !ignore_autogenerated
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

package v1alpha1

import (
	unsafe "unsafe"

	v1 "k8s.io/api/core/v1"
	conversion "k8s.io/apimachinery/pkg/conversion"
	runtime "k8s.io/apimachinery/pkg/runtime"
	metrics "k8s.io/metrics/pkg/apis/metrics"
)

func init() {
	localSchemeBuilder.Register(RegisterConversions)
}

// RegisterConversions adds conversion functions to the given scheme.
// Public to allow building arbitrary schemes.
func RegisterConversions(s *runtime.Scheme) error {
	if err := s.AddGeneratedConversionFunc((*ContainerMetrics)(nil), (*metrics.ContainerMetrics)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha1_ContainerMetrics_To_metrics_ContainerMetrics(a.(*ContainerMetrics), b.(*metrics.ContainerMetrics), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*metrics.ContainerMetrics)(nil), (*ContainerMetrics)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_metrics_ContainerMetrics_To_v1alpha1_ContainerMetrics(a.(*metrics.ContainerMetrics), b.(*ContainerMetrics), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*NodeMetrics)(nil), (*metrics.NodeMetrics)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha1_NodeMetrics_To_metrics_NodeMetrics(a.(*NodeMetrics), b.(*metrics.NodeMetrics), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*metrics.NodeMetrics)(nil), (*NodeMetrics)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_metrics_NodeMetrics_To_v1alpha1_NodeMetrics(a.(*metrics.NodeMetrics), b.(*NodeMetrics), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*NodeMetricsList)(nil), (*metrics.NodeMetricsList)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha1_NodeMetricsList_To_metrics_NodeMetricsList(a.(*NodeMetricsList), b.(*metrics.NodeMetricsList), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*metrics.NodeMetricsList)(nil), (*NodeMetricsList)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_metrics_NodeMetricsList_To_v1alpha1_NodeMetricsList(a.(*metrics.NodeMetricsList), b.(*NodeMetricsList), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*PodMetrics)(nil), (*metrics.PodMetrics)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha1_PodMetrics_To_metrics_PodMetrics(a.(*PodMetrics), b.(*metrics.PodMetrics), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*metrics.PodMetrics)(nil), (*PodMetrics)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_metrics_PodMetrics_To_v1alpha1_PodMetrics(a.(*metrics.PodMetrics), b.(*PodMetrics), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*PodMetricsList)(nil), (*metrics.PodMetricsList)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha1_PodMetricsList_To_metrics_PodMetricsList(a.(*PodMetricsList), b.(*metrics.PodMetricsList), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*metrics.PodMetricsList)(nil), (*PodMetricsList)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_metrics_PodMetricsList_To_v1alpha1_PodMetricsList(a.(*metrics.PodMetricsList), b.(*PodMetricsList), scope)
	}); err != nil {
		return err
	}
	return nil
}

func autoConvert_v1alpha1_ContainerMetrics_To_metrics_ContainerMetrics(in *ContainerMetrics, out *metrics.ContainerMetrics, s conversion.Scope) error {
	out.Name = in.Name
	out.Usage = *(*v1.ResourceList)(unsafe.Pointer(&in.Usage))
	return nil
}

// Convert_v1alpha1_ContainerMetrics_To_metrics_ContainerMetrics is an autogenerated conversion function.
func Convert_v1alpha1_ContainerMetrics_To_metrics_ContainerMetrics(in *ContainerMetrics, out *metrics.ContainerMetrics, s conversion.Scope) error {
	return autoConvert_v1alpha1_ContainerMetrics_To_metrics_ContainerMetrics(in, out, s)
}

func autoConvert_metrics_ContainerMetrics_To_v1alpha1_ContainerMetrics(in *metrics.ContainerMetrics, out *ContainerMetrics, s conversion.Scope) error {
	out.Name = in.Name
	out.Usage = *(*v1.ResourceList)(unsafe.Pointer(&in.Usage))
	return nil
}

// Convert_metrics_ContainerMetrics_To_v1alpha1_ContainerMetrics is an autogenerated conversion function.
func Convert_metrics_ContainerMetrics_To_v1alpha1_ContainerMetrics(in *metrics.ContainerMetrics, out *ContainerMetrics, s conversion.Scope) error {
	return autoConvert_metrics_ContainerMetrics_To_v1alpha1_ContainerMetrics(in, out, s)
}

func autoConvert_v1alpha1_NodeMetrics_To_metrics_NodeMetrics(in *NodeMetrics, out *metrics.NodeMetrics, s conversion.Scope) error {
	out.ObjectMeta = in.ObjectMeta
	out.Timestamp = in.Timestamp
	out.Window = in.Window
	out.Usage = *(*v1.ResourceList)(unsafe.Pointer(&in.Usage))
	return nil
}

// Convert_v1alpha1_NodeMetrics_To_metrics_NodeMetrics is an autogenerated conversion function.
func Convert_v1alpha1_NodeMetrics_To_metrics_NodeMetrics(in *NodeMetrics, out *metrics.NodeMetrics, s conversion.Scope) error {
	return autoConvert_v1alpha1_NodeMetrics_To_metrics_NodeMetrics(in, out, s)
}

func autoConvert_metrics_NodeMetrics_To_v1alpha1_NodeMetrics(in *metrics.NodeMetrics, out *NodeMetrics, s conversion.Scope) error {
	out.ObjectMeta = in.ObjectMeta
	out.Timestamp = in.Timestamp
	out.Window = in.Window
	out.Usage = *(*v1.ResourceList)(unsafe.Pointer(&in.Usage))
	return nil
}

// Convert_metrics_NodeMetrics_To_v1alpha1_NodeMetrics is an autogenerated conversion function.
func Convert_metrics_NodeMetrics_To_v1alpha1_NodeMetrics(in *metrics.NodeMetrics, out *NodeMetrics, s conversion.Scope) error {
	return autoConvert_metrics_NodeMetrics_To_v1alpha1_NodeMetrics(in, out, s)
}

func autoConvert_v1alpha1_NodeMetricsList_To_metrics_NodeMetricsList(in *NodeMetricsList, out *metrics.NodeMetricsList, s conversion.Scope) error {
	out.ListMeta = in.ListMeta
	out.Items = *(*[]metrics.NodeMetrics)(unsafe.Pointer(&in.Items))
	return nil
}

// Convert_v1alpha1_NodeMetricsList_To_metrics_NodeMetricsList is an autogenerated conversion function.
func Convert_v1alpha1_NodeMetricsList_To_metrics_NodeMetricsList(in *NodeMetricsList, out *metrics.NodeMetricsList, s conversion.Scope) error {
	return autoConvert_v1alpha1_NodeMetricsList_To_metrics_NodeMetricsList(in, out, s)
}

func autoConvert_metrics_NodeMetricsList_To_v1alpha1_NodeMetricsList(in *metrics.NodeMetricsList, out *NodeMetricsList, s conversion.Scope) error {
	out.ListMeta = in.ListMeta
	out.Items = *(*[]NodeMetrics)(unsafe.Pointer(&in.Items))
	return nil
}

// Convert_metrics_NodeMetricsList_To_v1alpha1_NodeMetricsList is an autogenerated conversion function.
func Convert_metrics_NodeMetricsList_To_v1alpha1_NodeMetricsList(in *metrics.NodeMetricsList, out *NodeMetricsList, s conversion.Scope) error {
	return autoConvert_metrics_NodeMetricsList_To_v1alpha1_NodeMetricsList(in, out, s)
}

func autoConvert_v1alpha1_PodMetrics_To_metrics_PodMetrics(in *PodMetrics, out *metrics.PodMetrics, s conversion.Scope) error {
	out.ObjectMeta = in.ObjectMeta
	out.Timestamp = in.Timestamp
	out.Window = in.Window
	out.Containers = *(*[]metrics.ContainerMetrics)(unsafe.Pointer(&in.Containers))
	return nil
}

// Convert_v1alpha1_PodMetrics_To_metrics_PodMetrics is an autogenerated conversion function.
func Convert_v1alpha1_PodMetrics_To_metrics_PodMetrics(in *PodMetrics, out *metrics.PodMetrics, s conversion.Scope) error {
	return autoConvert_v1alpha1_PodMetrics_To_metrics_PodMetrics(in, out, s)
}

func autoConvert_metrics_PodMetrics_To_v1alpha1_PodMetrics(in *metrics.PodMetrics, out *PodMetrics, s conversion.Scope) error {
	out.ObjectMeta = in.ObjectMeta
	out.Timestamp = in.Timestamp
	out.Window = in.Window
	out.Containers = *(*[]ContainerMetrics)(unsafe.Pointer(&in.Containers))
	return nil
}

// Convert_metrics_PodMetrics_To_v1alpha1_PodMetrics is an autogenerated conversion function.
func Convert_metrics_PodMetrics_To_v1alpha1_PodMetrics(in *metrics.PodMetrics, out *PodMetrics, s conversion.Scope) error {
	return autoConvert_metrics_PodMetrics_To_v1alpha1_PodMetrics(in, out, s)
}

func autoConvert_v1alpha1_PodMetricsList_To_metrics_PodMetricsList(in *PodMetricsList, out *metrics.PodMetricsList, s conversion.Scope) error {
	out.ListMeta = in.ListMeta
	out.Items = *(*[]metrics.PodMetrics)(unsafe.Pointer(&in.Items))
	return nil
}

// Convert_v1alpha1_PodMetricsList_To_metrics_PodMetricsList is an autogenerated conversion function.
func Convert_v1alpha1_PodMetricsList_To_metrics_PodMetricsList(in *PodMetricsList, out *metrics.PodMetricsList, s conversion.Scope) error {
	return autoConvert_v1alpha1_PodMetricsList_To_metrics_PodMetricsList(in, out, s)
}

func autoConvert_metrics_PodMetricsList_To_v1alpha1_PodMetricsList(in *metrics.PodMetricsList, out *PodMetricsList, s conversion.Scope) error {
	out.ListMeta = in.ListMeta
	out.Items = *(*[]PodMetrics)(unsafe.Pointer(&in.Items))
	return nil
}

// Convert_metrics_PodMetricsList_To_v1alpha1_PodMetricsList is an autogenerated conversion function.
func Convert_metrics_PodMetricsList_To_v1alpha1_PodMetricsList(in *metrics.PodMetricsList, out *PodMetricsList, s conversion.Scope) error {
	return autoConvert_metrics_PodMetricsList_To_v1alpha1_PodMetricsList(in, out, s)
}
