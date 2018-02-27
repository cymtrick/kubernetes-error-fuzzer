// +build !ignore_autogenerated

/*
Copyright 2018 The Kubernetes Authors.

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

package v1alpha1

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
	v1beta1 "k8s.io/kubernetes/pkg/kubelet/apis/kubeletconfig/v1beta1"
	kubeproxyconfig_v1alpha1 "k8s.io/kubernetes/pkg/proxy/apis/kubeproxyconfig/v1alpha1"
)

// RegisterDefaults adds defaulters functions to the given scheme.
// Public to allow building arbitrary schemes.
// All generated defaulters are covering - they call all nested defaulters.
func RegisterDefaults(scheme *runtime.Scheme) error {
	scheme.AddTypeDefaultingFunc(&MasterConfiguration{}, func(obj interface{}) { SetObjectDefaults_MasterConfiguration(obj.(*MasterConfiguration)) })
	scheme.AddTypeDefaultingFunc(&NodeConfiguration{}, func(obj interface{}) { SetObjectDefaults_NodeConfiguration(obj.(*NodeConfiguration)) })
	return nil
}

func SetObjectDefaults_MasterConfiguration(in *MasterConfiguration) {
	SetDefaults_MasterConfiguration(in)
	if in.KubeProxy.Config != nil {
		kubeproxyconfig_v1alpha1.SetDefaults_KubeProxyConfiguration(in.KubeProxy.Config)
	}
	if in.KubeletConfiguration.BaseConfig != nil {
		v1beta1.SetDefaults_KubeletConfiguration(in.KubeletConfiguration.BaseConfig)
	}
}

func SetObjectDefaults_NodeConfiguration(in *NodeConfiguration) {
	SetDefaults_NodeConfiguration(in)
}
