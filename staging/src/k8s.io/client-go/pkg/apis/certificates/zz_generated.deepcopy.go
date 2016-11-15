// +build !ignore_autogenerated

/*
Copyright 2016 The Kubernetes Authors.

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

// This file was autogenerated by deepcopy-gen. Do not edit it manually!

package certificates

import (
	api "k8s.io/client-go/pkg/api"
	conversion "k8s.io/client-go/pkg/conversion"
	runtime "k8s.io/client-go/pkg/runtime"
	reflect "reflect"
)

func init() {
	SchemeBuilder.Register(RegisterDeepCopies)
}

// RegisterDeepCopies adds deep-copy functions to the given scheme. Public
// to allow building arbitrary schemes.
func RegisterDeepCopies(scheme *runtime.Scheme) error {
	return scheme.AddGeneratedDeepCopyFuncs(
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_certificates_CertificateSigningRequest, InType: reflect.TypeOf(&CertificateSigningRequest{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_certificates_CertificateSigningRequestCondition, InType: reflect.TypeOf(&CertificateSigningRequestCondition{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_certificates_CertificateSigningRequestList, InType: reflect.TypeOf(&CertificateSigningRequestList{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_certificates_CertificateSigningRequestSpec, InType: reflect.TypeOf(&CertificateSigningRequestSpec{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_certificates_CertificateSigningRequestStatus, InType: reflect.TypeOf(&CertificateSigningRequestStatus{})},
	)
}

func DeepCopy_certificates_CertificateSigningRequest(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*CertificateSigningRequest)
		out := out.(*CertificateSigningRequest)
		out.TypeMeta = in.TypeMeta
		if err := api.DeepCopy_api_ObjectMeta(&in.ObjectMeta, &out.ObjectMeta, c); err != nil {
			return err
		}
		if err := DeepCopy_certificates_CertificateSigningRequestSpec(&in.Spec, &out.Spec, c); err != nil {
			return err
		}
		if err := DeepCopy_certificates_CertificateSigningRequestStatus(&in.Status, &out.Status, c); err != nil {
			return err
		}
		return nil
	}
}

func DeepCopy_certificates_CertificateSigningRequestCondition(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*CertificateSigningRequestCondition)
		out := out.(*CertificateSigningRequestCondition)
		out.Type = in.Type
		out.Reason = in.Reason
		out.Message = in.Message
		out.LastUpdateTime = in.LastUpdateTime.DeepCopy()
		return nil
	}
}

func DeepCopy_certificates_CertificateSigningRequestList(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*CertificateSigningRequestList)
		out := out.(*CertificateSigningRequestList)
		out.TypeMeta = in.TypeMeta
		out.ListMeta = in.ListMeta
		if in.Items != nil {
			in, out := &in.Items, &out.Items
			*out = make([]CertificateSigningRequest, len(*in))
			for i := range *in {
				if err := DeepCopy_certificates_CertificateSigningRequest(&(*in)[i], &(*out)[i], c); err != nil {
					return err
				}
			}
		} else {
			out.Items = nil
		}
		return nil
	}
}

func DeepCopy_certificates_CertificateSigningRequestSpec(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*CertificateSigningRequestSpec)
		out := out.(*CertificateSigningRequestSpec)
		if in.Request != nil {
			in, out := &in.Request, &out.Request
			*out = make([]byte, len(*in))
			copy(*out, *in)
		} else {
			out.Request = nil
		}
		out.Username = in.Username
		out.UID = in.UID
		if in.Groups != nil {
			in, out := &in.Groups, &out.Groups
			*out = make([]string, len(*in))
			copy(*out, *in)
		} else {
			out.Groups = nil
		}
		return nil
	}
}

func DeepCopy_certificates_CertificateSigningRequestStatus(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*CertificateSigningRequestStatus)
		out := out.(*CertificateSigningRequestStatus)
		if in.Conditions != nil {
			in, out := &in.Conditions, &out.Conditions
			*out = make([]CertificateSigningRequestCondition, len(*in))
			for i := range *in {
				if err := DeepCopy_certificates_CertificateSigningRequestCondition(&(*in)[i], &(*out)[i], c); err != nil {
					return err
				}
			}
		} else {
			out.Conditions = nil
		}
		if in.Certificate != nil {
			in, out := &in.Certificate, &out.Certificate
			*out = make([]byte, len(*in))
			copy(*out, *in)
		} else {
			out.Certificate = nil
		}
		return nil
	}
}
