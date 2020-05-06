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
	clientauthentication "k8s.io/client-go/pkg/apis/clientauthentication"
)

func init() {
	localSchemeBuilder.Register(RegisterConversions)
}

// RegisterConversions adds conversion functions to the given scheme.
// Public to allow building arbitrary schemes.
func RegisterConversions(s *runtime.Scheme) error {
	if err := s.AddGeneratedConversionFunc((*ExecCredential)(nil), (*clientauthentication.ExecCredential)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha1_ExecCredential_To_clientauthentication_ExecCredential(a.(*ExecCredential), b.(*clientauthentication.ExecCredential), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*clientauthentication.ExecCredential)(nil), (*ExecCredential)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_clientauthentication_ExecCredential_To_v1alpha1_ExecCredential(a.(*clientauthentication.ExecCredential), b.(*ExecCredential), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*ExecCredentialSpec)(nil), (*clientauthentication.ExecCredentialSpec)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha1_ExecCredentialSpec_To_clientauthentication_ExecCredentialSpec(a.(*ExecCredentialSpec), b.(*clientauthentication.ExecCredentialSpec), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*ExecCredentialStatus)(nil), (*clientauthentication.ExecCredentialStatus)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha1_ExecCredentialStatus_To_clientauthentication_ExecCredentialStatus(a.(*ExecCredentialStatus), b.(*clientauthentication.ExecCredentialStatus), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*clientauthentication.ExecCredentialStatus)(nil), (*ExecCredentialStatus)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_clientauthentication_ExecCredentialStatus_To_v1alpha1_ExecCredentialStatus(a.(*clientauthentication.ExecCredentialStatus), b.(*ExecCredentialStatus), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*Response)(nil), (*clientauthentication.Response)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha1_Response_To_clientauthentication_Response(a.(*Response), b.(*clientauthentication.Response), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*clientauthentication.Response)(nil), (*Response)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_clientauthentication_Response_To_v1alpha1_Response(a.(*clientauthentication.Response), b.(*Response), scope)
	}); err != nil {
		return err
	}
	if err := s.AddConversionFunc((*clientauthentication.ExecCredentialSpec)(nil), (*ExecCredentialSpec)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_clientauthentication_ExecCredentialSpec_To_v1alpha1_ExecCredentialSpec(a.(*clientauthentication.ExecCredentialSpec), b.(*ExecCredentialSpec), scope)
	}); err != nil {
		return err
	}
	return nil
}

func autoConvert_v1alpha1_ExecCredential_To_clientauthentication_ExecCredential(in *ExecCredential, out *clientauthentication.ExecCredential, s conversion.Scope) error {
	if err := Convert_v1alpha1_ExecCredentialSpec_To_clientauthentication_ExecCredentialSpec(&in.Spec, &out.Spec, s); err != nil {
		return err
	}
	out.Status = (*clientauthentication.ExecCredentialStatus)(unsafe.Pointer(in.Status))
	return nil
}

// Convert_v1alpha1_ExecCredential_To_clientauthentication_ExecCredential is an autogenerated conversion function.
func Convert_v1alpha1_ExecCredential_To_clientauthentication_ExecCredential(in *ExecCredential, out *clientauthentication.ExecCredential, s conversion.Scope) error {
	return autoConvert_v1alpha1_ExecCredential_To_clientauthentication_ExecCredential(in, out, s)
}

func autoConvert_clientauthentication_ExecCredential_To_v1alpha1_ExecCredential(in *clientauthentication.ExecCredential, out *ExecCredential, s conversion.Scope) error {
	if err := Convert_clientauthentication_ExecCredentialSpec_To_v1alpha1_ExecCredentialSpec(&in.Spec, &out.Spec, s); err != nil {
		return err
	}
	out.Status = (*ExecCredentialStatus)(unsafe.Pointer(in.Status))
	return nil
}

// Convert_clientauthentication_ExecCredential_To_v1alpha1_ExecCredential is an autogenerated conversion function.
func Convert_clientauthentication_ExecCredential_To_v1alpha1_ExecCredential(in *clientauthentication.ExecCredential, out *ExecCredential, s conversion.Scope) error {
	return autoConvert_clientauthentication_ExecCredential_To_v1alpha1_ExecCredential(in, out, s)
}

func autoConvert_v1alpha1_ExecCredentialSpec_To_clientauthentication_ExecCredentialSpec(in *ExecCredentialSpec, out *clientauthentication.ExecCredentialSpec, s conversion.Scope) error {
	out.Response = (*clientauthentication.Response)(unsafe.Pointer(in.Response))
	out.Interactive = in.Interactive
	return nil
}

// Convert_v1alpha1_ExecCredentialSpec_To_clientauthentication_ExecCredentialSpec is an autogenerated conversion function.
func Convert_v1alpha1_ExecCredentialSpec_To_clientauthentication_ExecCredentialSpec(in *ExecCredentialSpec, out *clientauthentication.ExecCredentialSpec, s conversion.Scope) error {
	return autoConvert_v1alpha1_ExecCredentialSpec_To_clientauthentication_ExecCredentialSpec(in, out, s)
}

func autoConvert_clientauthentication_ExecCredentialSpec_To_v1alpha1_ExecCredentialSpec(in *clientauthentication.ExecCredentialSpec, out *ExecCredentialSpec, s conversion.Scope) error {
	out.Response = (*Response)(unsafe.Pointer(in.Response))
	out.Interactive = in.Interactive
	// WARNING: in.Cluster requires manual conversion: does not exist in peer-type
	return nil
}

func autoConvert_v1alpha1_ExecCredentialStatus_To_clientauthentication_ExecCredentialStatus(in *ExecCredentialStatus, out *clientauthentication.ExecCredentialStatus, s conversion.Scope) error {
	out.ExpirationTimestamp = (*v1.Time)(unsafe.Pointer(in.ExpirationTimestamp))
	out.Token = in.Token
	out.ClientCertificateData = in.ClientCertificateData
	out.ClientKeyData = in.ClientKeyData
	return nil
}

// Convert_v1alpha1_ExecCredentialStatus_To_clientauthentication_ExecCredentialStatus is an autogenerated conversion function.
func Convert_v1alpha1_ExecCredentialStatus_To_clientauthentication_ExecCredentialStatus(in *ExecCredentialStatus, out *clientauthentication.ExecCredentialStatus, s conversion.Scope) error {
	return autoConvert_v1alpha1_ExecCredentialStatus_To_clientauthentication_ExecCredentialStatus(in, out, s)
}

func autoConvert_clientauthentication_ExecCredentialStatus_To_v1alpha1_ExecCredentialStatus(in *clientauthentication.ExecCredentialStatus, out *ExecCredentialStatus, s conversion.Scope) error {
	out.ExpirationTimestamp = (*v1.Time)(unsafe.Pointer(in.ExpirationTimestamp))
	out.Token = in.Token
	out.ClientCertificateData = in.ClientCertificateData
	out.ClientKeyData = in.ClientKeyData
	return nil
}

// Convert_clientauthentication_ExecCredentialStatus_To_v1alpha1_ExecCredentialStatus is an autogenerated conversion function.
func Convert_clientauthentication_ExecCredentialStatus_To_v1alpha1_ExecCredentialStatus(in *clientauthentication.ExecCredentialStatus, out *ExecCredentialStatus, s conversion.Scope) error {
	return autoConvert_clientauthentication_ExecCredentialStatus_To_v1alpha1_ExecCredentialStatus(in, out, s)
}

func autoConvert_v1alpha1_Response_To_clientauthentication_Response(in *Response, out *clientauthentication.Response, s conversion.Scope) error {
	out.Header = *(*map[string][]string)(unsafe.Pointer(&in.Header))
	out.Code = in.Code
	return nil
}

// Convert_v1alpha1_Response_To_clientauthentication_Response is an autogenerated conversion function.
func Convert_v1alpha1_Response_To_clientauthentication_Response(in *Response, out *clientauthentication.Response, s conversion.Scope) error {
	return autoConvert_v1alpha1_Response_To_clientauthentication_Response(in, out, s)
}

func autoConvert_clientauthentication_Response_To_v1alpha1_Response(in *clientauthentication.Response, out *Response, s conversion.Scope) error {
	out.Header = *(*map[string][]string)(unsafe.Pointer(&in.Header))
	out.Code = in.Code
	return nil
}

// Convert_clientauthentication_Response_To_v1alpha1_Response is an autogenerated conversion function.
func Convert_clientauthentication_Response_To_v1alpha1_Response(in *clientauthentication.Response, out *Response, s conversion.Scope) error {
	return autoConvert_clientauthentication_Response_To_v1alpha1_Response(in, out, s)
}
