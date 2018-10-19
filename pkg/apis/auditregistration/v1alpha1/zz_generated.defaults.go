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

package v1alpha1

import (
	v1alpha1 "k8s.io/api/auditregistration/v1alpha1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// RegisterDefaults adds defaulters functions to the given scheme.
// Public to allow building arbitrary schemes.
// All generated defaulters are covering - they call all nested defaulters.
func RegisterDefaults(scheme *runtime.Scheme) error {
	scheme.AddTypeDefaultingFunc(&v1alpha1.AuditSink{}, func(obj interface{}) { SetObjectDefaults_AuditSink(obj.(*v1alpha1.AuditSink)) })
	scheme.AddTypeDefaultingFunc(&v1alpha1.AuditSinkList{}, func(obj interface{}) { SetObjectDefaults_AuditSinkList(obj.(*v1alpha1.AuditSinkList)) })
	return nil
}

func SetObjectDefaults_AuditSink(in *v1alpha1.AuditSink) {
	SetDefaults_AuditSink(in)
}

func SetObjectDefaults_AuditSinkList(in *v1alpha1.AuditSinkList) {
	for i := range in.Items {
		a := &in.Items[i]
		SetObjectDefaults_AuditSink(a)
	}
}
