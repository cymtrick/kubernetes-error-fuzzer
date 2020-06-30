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

// Code generated by deepcopy-gen. DO NOT EDIT.

package v1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AdmissionRequest) DeepCopyInto(out *AdmissionRequest) {
	*out = *in
	out.Kind = in.Kind
	out.Resource = in.Resource
	if in.RequestKind != nil {
		in, out := &in.RequestKind, &out.RequestKind
		*out = new(metav1.GroupVersionKind)
		**out = **in
	}
	if in.RequestResource != nil {
		in, out := &in.RequestResource, &out.RequestResource
		*out = new(metav1.GroupVersionResource)
		**out = **in
	}
	in.UserInfo.DeepCopyInto(&out.UserInfo)
	in.Object.DeepCopyInto(&out.Object)
	in.OldObject.DeepCopyInto(&out.OldObject)
	if in.DryRun != nil {
		in, out := &in.DryRun, &out.DryRun
		*out = new(bool)
		**out = **in
	}
	in.Options.DeepCopyInto(&out.Options)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AdmissionRequest.
func (in *AdmissionRequest) DeepCopy() *AdmissionRequest {
	if in == nil {
		return nil
	}
	out := new(AdmissionRequest)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AdmissionResponse) DeepCopyInto(out *AdmissionResponse) {
	*out = *in
	if in.Result != nil {
		in, out := &in.Result, &out.Result
		*out = new(metav1.Status)
		(*in).DeepCopyInto(*out)
	}
	if in.Patch != nil {
		in, out := &in.Patch, &out.Patch
		*out = make([]byte, len(*in))
		copy(*out, *in)
	}
	if in.PatchType != nil {
		in, out := &in.PatchType, &out.PatchType
		*out = new(PatchType)
		**out = **in
	}
	if in.AuditAnnotations != nil {
		in, out := &in.AuditAnnotations, &out.AuditAnnotations
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.Warnings != nil {
		in, out := &in.Warnings, &out.Warnings
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AdmissionResponse.
func (in *AdmissionResponse) DeepCopy() *AdmissionResponse {
	if in == nil {
		return nil
	}
	out := new(AdmissionResponse)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AdmissionReview) DeepCopyInto(out *AdmissionReview) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	if in.Request != nil {
		in, out := &in.Request, &out.Request
		*out = new(AdmissionRequest)
		(*in).DeepCopyInto(*out)
	}
	if in.Response != nil {
		in, out := &in.Response, &out.Response
		*out = new(AdmissionResponse)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AdmissionReview.
func (in *AdmissionReview) DeepCopy() *AdmissionReview {
	if in == nil {
		return nil
	}
	out := new(AdmissionReview)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *AdmissionReview) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}
