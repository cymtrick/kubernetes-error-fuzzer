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

// Code generated by defaulter-gen. DO NOT EDIT.

package v1beta1

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// RegisterDefaults adds defaulters functions to the given scheme.
// Public to allow building arbitrary schemes.
// All generated defaulters are covering - they call all nested defaulters.
func RegisterDefaults(scheme *runtime.Scheme) error {
	scheme.AddTypeDefaultingFunc(&ClusterConfiguration{}, func(obj interface{}) { SetObjectDefaults_ClusterConfiguration(obj.(*ClusterConfiguration)) })
	scheme.AddTypeDefaultingFunc(&ClusterStatus{}, func(obj interface{}) { SetObjectDefaults_ClusterStatus(obj.(*ClusterStatus)) })
	scheme.AddTypeDefaultingFunc(&InitConfiguration{}, func(obj interface{}) { SetObjectDefaults_InitConfiguration(obj.(*InitConfiguration)) })
	scheme.AddTypeDefaultingFunc(&JoinConfiguration{}, func(obj interface{}) { SetObjectDefaults_JoinConfiguration(obj.(*JoinConfiguration)) })
	return nil
}

func SetObjectDefaults_ClusterConfiguration(in *ClusterConfiguration) {
	SetDefaults_ClusterConfiguration(in)
}

func SetObjectDefaults_ClusterStatus(in *ClusterStatus) {
}

func SetObjectDefaults_InitConfiguration(in *InitConfiguration) {
	SetDefaults_InitConfiguration(in)
	SetObjectDefaults_ClusterConfiguration(&in.ClusterConfiguration)
	for i := range in.BootstrapTokens {
		a := &in.BootstrapTokens[i]
		SetDefaults_BootstrapToken(a)
	}
	SetDefaults_NodeRegistrationOptions(&in.NodeRegistration)
	SetDefaults_APIEndpoint(&in.APIEndpoint)
}

func SetObjectDefaults_JoinConfiguration(in *JoinConfiguration) {
	SetDefaults_JoinConfiguration(in)
	SetDefaults_NodeRegistrationOptions(&in.NodeRegistration)
	SetDefaults_Discovery(&in.Discovery)
	if in.Discovery.File != nil {
		SetDefaults_FileDiscovery(in.Discovery.File)
	}
	SetDefaults_APIEndpoint(&in.APIEndpoint)
}
