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

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	conversion "k8s.io/apimachinery/pkg/conversion"
	runtime "k8s.io/apimachinery/pkg/runtime"
	v1alpha1 "k8s.io/kube-controller-manager/config/v1alpha1"
	config "k8s.io/kubernetes/pkg/controller/garbagecollector/config"
)

func init() {
	localSchemeBuilder.Register(RegisterConversions)
}

// RegisterConversions adds conversion functions to the given scheme.
// Public to allow building arbitrary schemes.
func RegisterConversions(s *runtime.Scheme) error {
	if err := s.AddGeneratedConversionFunc((*v1alpha1.GroupResource)(nil), (*config.GroupResource)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha1_GroupResource_To_config_GroupResource(a.(*v1alpha1.GroupResource), b.(*config.GroupResource), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*config.GroupResource)(nil), (*v1alpha1.GroupResource)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_config_GroupResource_To_v1alpha1_GroupResource(a.(*config.GroupResource), b.(*v1alpha1.GroupResource), scope)
	}); err != nil {
		return err
	}
	if err := s.AddConversionFunc((*config.GarbageCollectorControllerConfiguration)(nil), (*v1alpha1.GarbageCollectorControllerConfiguration)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_config_GarbageCollectorControllerConfiguration_To_v1alpha1_GarbageCollectorControllerConfiguration(a.(*config.GarbageCollectorControllerConfiguration), b.(*v1alpha1.GarbageCollectorControllerConfiguration), scope)
	}); err != nil {
		return err
	}
	if err := s.AddConversionFunc((*v1alpha1.GarbageCollectorControllerConfiguration)(nil), (*config.GarbageCollectorControllerConfiguration)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha1_GarbageCollectorControllerConfiguration_To_config_GarbageCollectorControllerConfiguration(a.(*v1alpha1.GarbageCollectorControllerConfiguration), b.(*config.GarbageCollectorControllerConfiguration), scope)
	}); err != nil {
		return err
	}
	return nil
}

func autoConvert_v1alpha1_GarbageCollectorControllerConfiguration_To_config_GarbageCollectorControllerConfiguration(in *v1alpha1.GarbageCollectorControllerConfiguration, out *config.GarbageCollectorControllerConfiguration, s conversion.Scope) error {
	if err := v1.Convert_Pointer_bool_To_bool(&in.EnableGarbageCollector, &out.EnableGarbageCollector, s); err != nil {
		return err
	}
	out.ConcurrentGCSyncs = in.ConcurrentGCSyncs
	out.GCIgnoredResources = *(*[]config.GroupResource)(unsafe.Pointer(&in.GCIgnoredResources))
	return nil
}

func autoConvert_config_GarbageCollectorControllerConfiguration_To_v1alpha1_GarbageCollectorControllerConfiguration(in *config.GarbageCollectorControllerConfiguration, out *v1alpha1.GarbageCollectorControllerConfiguration, s conversion.Scope) error {
	if err := v1.Convert_bool_To_Pointer_bool(&in.EnableGarbageCollector, &out.EnableGarbageCollector, s); err != nil {
		return err
	}
	out.ConcurrentGCSyncs = in.ConcurrentGCSyncs
	out.GCIgnoredResources = *(*[]v1alpha1.GroupResource)(unsafe.Pointer(&in.GCIgnoredResources))
	return nil
}

func autoConvert_v1alpha1_GroupResource_To_config_GroupResource(in *v1alpha1.GroupResource, out *config.GroupResource, s conversion.Scope) error {
	out.Group = in.Group
	out.Resource = in.Resource
	return nil
}

// Convert_v1alpha1_GroupResource_To_config_GroupResource is an autogenerated conversion function.
func Convert_v1alpha1_GroupResource_To_config_GroupResource(in *v1alpha1.GroupResource, out *config.GroupResource, s conversion.Scope) error {
	return autoConvert_v1alpha1_GroupResource_To_config_GroupResource(in, out, s)
}

func autoConvert_config_GroupResource_To_v1alpha1_GroupResource(in *config.GroupResource, out *v1alpha1.GroupResource, s conversion.Scope) error {
	out.Group = in.Group
	out.Resource = in.Resource
	return nil
}

// Convert_config_GroupResource_To_v1alpha1_GroupResource is an autogenerated conversion function.
func Convert_config_GroupResource_To_v1alpha1_GroupResource(in *config.GroupResource, out *v1alpha1.GroupResource, s conversion.Scope) error {
	return autoConvert_config_GroupResource_To_v1alpha1_GroupResource(in, out, s)
}
