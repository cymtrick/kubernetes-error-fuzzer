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
	conversion "k8s.io/apimachinery/pkg/conversion"
	runtime "k8s.io/apimachinery/pkg/runtime"
	config "k8s.io/cloud-provider/controllers/service/config"
)

func init() {
	localSchemeBuilder.Register(RegisterConversions)
}

// RegisterConversions adds conversion functions to the given scheme.
// Public to allow building arbitrary schemes.
func RegisterConversions(s *runtime.Scheme) error {
	if err := s.AddConversionFunc((*config.ServiceControllerConfiguration)(nil), (*ServiceControllerConfiguration)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_config_ServiceControllerConfiguration_To_v1alpha1_ServiceControllerConfiguration(a.(*config.ServiceControllerConfiguration), b.(*ServiceControllerConfiguration), scope)
	}); err != nil {
		return err
	}
	if err := s.AddConversionFunc((*ServiceControllerConfiguration)(nil), (*config.ServiceControllerConfiguration)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha1_ServiceControllerConfiguration_To_config_ServiceControllerConfiguration(a.(*ServiceControllerConfiguration), b.(*config.ServiceControllerConfiguration), scope)
	}); err != nil {
		return err
	}
	return nil
}

func autoConvert_v1alpha1_ServiceControllerConfiguration_To_config_ServiceControllerConfiguration(in *ServiceControllerConfiguration, out *config.ServiceControllerConfiguration, s conversion.Scope) error {
	out.ConcurrentServiceSyncs = in.ConcurrentServiceSyncs
	return nil
}

func autoConvert_config_ServiceControllerConfiguration_To_v1alpha1_ServiceControllerConfiguration(in *config.ServiceControllerConfiguration, out *ServiceControllerConfiguration, s conversion.Scope) error {
	out.ConcurrentServiceSyncs = in.ConcurrentServiceSyncs
	return nil
}
