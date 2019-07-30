/*
Copyright 2019 The Kubernetes Authors.

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

package v1alpha1

import (
	v1 "k8s.io/api/core/v1"
	discoveryv1alpha1 "k8s.io/api/discovery/v1alpha1"
	"k8s.io/apimachinery/pkg/runtime"
)

var (
	defaultAddressType = discoveryv1alpha1.AddressTypeIP
	defaultPortName    = ""
	defaultProtocol    = v1.ProtocolTCP
)

func addDefaultingFuncs(scheme *runtime.Scheme) error {
	return RegisterDefaults(scheme)
}

func SetDefaults_EndpointSlice(obj *discoveryv1alpha1.EndpointSlice) {
	if obj.AddressType == nil {
		obj.AddressType = &defaultAddressType
	}
}

func SetDefaults_EndpointPort(obj *discoveryv1alpha1.EndpointPort) {
	if obj.Name == nil {
		obj.Name = &defaultPortName
	}

	if obj.Protocol == nil {
		obj.Protocol = &defaultProtocol
	}
}
