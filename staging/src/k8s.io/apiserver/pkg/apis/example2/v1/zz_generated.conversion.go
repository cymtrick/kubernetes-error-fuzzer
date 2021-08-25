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

package v1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	conversion "k8s.io/apimachinery/pkg/conversion"
	runtime "k8s.io/apimachinery/pkg/runtime"
	example "k8s.io/apiserver/pkg/apis/example"
)

func init() {
	localSchemeBuilder.Register(RegisterConversions)
}

// RegisterConversions adds conversion functions to the given scheme.
// Public to allow building arbitrary schemes.
func RegisterConversions(s *runtime.Scheme) error {
	if err := s.AddGeneratedConversionFunc((*ReplicaSet)(nil), (*example.ReplicaSet)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1_ReplicaSet_To_example_ReplicaSet(a.(*ReplicaSet), b.(*example.ReplicaSet), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*example.ReplicaSet)(nil), (*ReplicaSet)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_example_ReplicaSet_To_v1_ReplicaSet(a.(*example.ReplicaSet), b.(*ReplicaSet), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*ReplicaSetStatus)(nil), (*example.ReplicaSetStatus)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1_ReplicaSetStatus_To_example_ReplicaSetStatus(a.(*ReplicaSetStatus), b.(*example.ReplicaSetStatus), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*example.ReplicaSetStatus)(nil), (*ReplicaSetStatus)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_example_ReplicaSetStatus_To_v1_ReplicaSetStatus(a.(*example.ReplicaSetStatus), b.(*ReplicaSetStatus), scope)
	}); err != nil {
		return err
	}
	if err := s.AddConversionFunc((*example.ReplicaSetSpec)(nil), (*ReplicaSetSpec)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_example_ReplicaSetSpec_To_v1_ReplicaSetSpec(a.(*example.ReplicaSetSpec), b.(*ReplicaSetSpec), scope)
	}); err != nil {
		return err
	}
	if err := s.AddConversionFunc((*ReplicaSetSpec)(nil), (*example.ReplicaSetSpec)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1_ReplicaSetSpec_To_example_ReplicaSetSpec(a.(*ReplicaSetSpec), b.(*example.ReplicaSetSpec), scope)
	}); err != nil {
		return err
	}
	return nil
}

func autoConvert_v1_ReplicaSet_To_example_ReplicaSet(in *ReplicaSet, out *example.ReplicaSet, s conversion.Scope) error {
	out.ObjectMeta = in.ObjectMeta
	if err := Convert_v1_ReplicaSetSpec_To_example_ReplicaSetSpec(&in.Spec, &out.Spec, s); err != nil {
		return err
	}
	if err := Convert_v1_ReplicaSetStatus_To_example_ReplicaSetStatus(&in.Status, &out.Status, s); err != nil {
		return err
	}
	return nil
}

// Convert_v1_ReplicaSet_To_example_ReplicaSet is an autogenerated conversion function.
func Convert_v1_ReplicaSet_To_example_ReplicaSet(in *ReplicaSet, out *example.ReplicaSet, s conversion.Scope) error {
	return autoConvert_v1_ReplicaSet_To_example_ReplicaSet(in, out, s)
}

func autoConvert_example_ReplicaSet_To_v1_ReplicaSet(in *example.ReplicaSet, out *ReplicaSet, s conversion.Scope) error {
	out.ObjectMeta = in.ObjectMeta
	if err := Convert_example_ReplicaSetSpec_To_v1_ReplicaSetSpec(&in.Spec, &out.Spec, s); err != nil {
		return err
	}
	if err := Convert_example_ReplicaSetStatus_To_v1_ReplicaSetStatus(&in.Status, &out.Status, s); err != nil {
		return err
	}
	return nil
}

// Convert_example_ReplicaSet_To_v1_ReplicaSet is an autogenerated conversion function.
func Convert_example_ReplicaSet_To_v1_ReplicaSet(in *example.ReplicaSet, out *ReplicaSet, s conversion.Scope) error {
	return autoConvert_example_ReplicaSet_To_v1_ReplicaSet(in, out, s)
}

func autoConvert_v1_ReplicaSetSpec_To_example_ReplicaSetSpec(in *ReplicaSetSpec, out *example.ReplicaSetSpec, s conversion.Scope) error {
	if err := metav1.Convert_Pointer_int32_To_int32(&in.Replicas, &out.Replicas, s); err != nil {
		return err
	}
	return nil
}

func autoConvert_example_ReplicaSetSpec_To_v1_ReplicaSetSpec(in *example.ReplicaSetSpec, out *ReplicaSetSpec, s conversion.Scope) error {
	if err := metav1.Convert_int32_To_Pointer_int32(&in.Replicas, &out.Replicas, s); err != nil {
		return err
	}
	return nil
}

func autoConvert_v1_ReplicaSetStatus_To_example_ReplicaSetStatus(in *ReplicaSetStatus, out *example.ReplicaSetStatus, s conversion.Scope) error {
	out.Replicas = in.Replicas
	return nil
}

// Convert_v1_ReplicaSetStatus_To_example_ReplicaSetStatus is an autogenerated conversion function.
func Convert_v1_ReplicaSetStatus_To_example_ReplicaSetStatus(in *ReplicaSetStatus, out *example.ReplicaSetStatus, s conversion.Scope) error {
	return autoConvert_v1_ReplicaSetStatus_To_example_ReplicaSetStatus(in, out, s)
}

func autoConvert_example_ReplicaSetStatus_To_v1_ReplicaSetStatus(in *example.ReplicaSetStatus, out *ReplicaSetStatus, s conversion.Scope) error {
	out.Replicas = in.Replicas
	return nil
}

// Convert_example_ReplicaSetStatus_To_v1_ReplicaSetStatus is an autogenerated conversion function.
func Convert_example_ReplicaSetStatus_To_v1_ReplicaSetStatus(in *example.ReplicaSetStatus, out *ReplicaSetStatus, s conversion.Scope) error {
	return autoConvert_example_ReplicaSetStatus_To_v1_ReplicaSetStatus(in, out, s)
}
