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
	unsafe "unsafe"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
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
	if err := s.AddGeneratedConversionFunc((*Cluster)(nil), (*clientauthentication.Cluster)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1_Cluster_To_clientauthentication_Cluster(a.(*Cluster), b.(*clientauthentication.Cluster), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*clientauthentication.Cluster)(nil), (*Cluster)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_clientauthentication_Cluster_To_v1_Cluster(a.(*clientauthentication.Cluster), b.(*Cluster), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*ExecCredential)(nil), (*clientauthentication.ExecCredential)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1_ExecCredential_To_clientauthentication_ExecCredential(a.(*ExecCredential), b.(*clientauthentication.ExecCredential), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*clientauthentication.ExecCredential)(nil), (*ExecCredential)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_clientauthentication_ExecCredential_To_v1_ExecCredential(a.(*clientauthentication.ExecCredential), b.(*ExecCredential), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*ExecCredentialSpec)(nil), (*clientauthentication.ExecCredentialSpec)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1_ExecCredentialSpec_To_clientauthentication_ExecCredentialSpec(a.(*ExecCredentialSpec), b.(*clientauthentication.ExecCredentialSpec), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*ExecCredentialStatus)(nil), (*clientauthentication.ExecCredentialStatus)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1_ExecCredentialStatus_To_clientauthentication_ExecCredentialStatus(a.(*ExecCredentialStatus), b.(*clientauthentication.ExecCredentialStatus), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*clientauthentication.ExecCredentialStatus)(nil), (*ExecCredentialStatus)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_clientauthentication_ExecCredentialStatus_To_v1_ExecCredentialStatus(a.(*clientauthentication.ExecCredentialStatus), b.(*ExecCredentialStatus), scope)
	}); err != nil {
		return err
	}
	if err := s.AddConversionFunc((*clientauthentication.ExecCredentialSpec)(nil), (*ExecCredentialSpec)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_clientauthentication_ExecCredentialSpec_To_v1_ExecCredentialSpec(a.(*clientauthentication.ExecCredentialSpec), b.(*ExecCredentialSpec), scope)
	}); err != nil {
		return err
	}
	return nil
}

func autoConvert_v1_Cluster_To_clientauthentication_Cluster(in *Cluster, out *clientauthentication.Cluster, s conversion.Scope) error {
	out.Server = in.Server
	out.TLSServerName = in.TLSServerName
	out.InsecureSkipTLSVerify = in.InsecureSkipTLSVerify
	out.CertificateAuthorityData = *(*[]byte)(unsafe.Pointer(&in.CertificateAuthorityData))
	out.ProxyURL = in.ProxyURL
	if err := runtime.Convert_runtime_RawExtension_To_runtime_Object(&in.Config, &out.Config, s); err != nil {
		return err
	}
	return nil
}

// Convert_v1_Cluster_To_clientauthentication_Cluster is an autogenerated conversion function.
func Convert_v1_Cluster_To_clientauthentication_Cluster(in *Cluster, out *clientauthentication.Cluster, s conversion.Scope) error {
	return autoConvert_v1_Cluster_To_clientauthentication_Cluster(in, out, s)
}

func autoConvert_clientauthentication_Cluster_To_v1_Cluster(in *clientauthentication.Cluster, out *Cluster, s conversion.Scope) error {
	out.Server = in.Server
	out.TLSServerName = in.TLSServerName
	out.InsecureSkipTLSVerify = in.InsecureSkipTLSVerify
	out.CertificateAuthorityData = *(*[]byte)(unsafe.Pointer(&in.CertificateAuthorityData))
	out.ProxyURL = in.ProxyURL
	if err := runtime.Convert_runtime_Object_To_runtime_RawExtension(&in.Config, &out.Config, s); err != nil {
		return err
	}
	return nil
}

// Convert_clientauthentication_Cluster_To_v1_Cluster is an autogenerated conversion function.
func Convert_clientauthentication_Cluster_To_v1_Cluster(in *clientauthentication.Cluster, out *Cluster, s conversion.Scope) error {
	return autoConvert_clientauthentication_Cluster_To_v1_Cluster(in, out, s)
}

func autoConvert_v1_ExecCredential_To_clientauthentication_ExecCredential(in *ExecCredential, out *clientauthentication.ExecCredential, s conversion.Scope) error {
	if err := Convert_v1_ExecCredentialSpec_To_clientauthentication_ExecCredentialSpec(&in.Spec, &out.Spec, s); err != nil {
		return err
	}
	out.Status = (*clientauthentication.ExecCredentialStatus)(unsafe.Pointer(in.Status))
	return nil
}

// Convert_v1_ExecCredential_To_clientauthentication_ExecCredential is an autogenerated conversion function.
func Convert_v1_ExecCredential_To_clientauthentication_ExecCredential(in *ExecCredential, out *clientauthentication.ExecCredential, s conversion.Scope) error {
	return autoConvert_v1_ExecCredential_To_clientauthentication_ExecCredential(in, out, s)
}

func autoConvert_clientauthentication_ExecCredential_To_v1_ExecCredential(in *clientauthentication.ExecCredential, out *ExecCredential, s conversion.Scope) error {
	if err := Convert_clientauthentication_ExecCredentialSpec_To_v1_ExecCredentialSpec(&in.Spec, &out.Spec, s); err != nil {
		return err
	}
	out.Status = (*ExecCredentialStatus)(unsafe.Pointer(in.Status))
	return nil
}

// Convert_clientauthentication_ExecCredential_To_v1_ExecCredential is an autogenerated conversion function.
func Convert_clientauthentication_ExecCredential_To_v1_ExecCredential(in *clientauthentication.ExecCredential, out *ExecCredential, s conversion.Scope) error {
	return autoConvert_clientauthentication_ExecCredential_To_v1_ExecCredential(in, out, s)
}

func autoConvert_v1_ExecCredentialSpec_To_clientauthentication_ExecCredentialSpec(in *ExecCredentialSpec, out *clientauthentication.ExecCredentialSpec, s conversion.Scope) error {
	if in.Cluster != nil {
		in, out := &in.Cluster, &out.Cluster
		*out = new(clientauthentication.Cluster)
		if err := Convert_v1_Cluster_To_clientauthentication_Cluster(*in, *out, s); err != nil {
			return err
		}
	} else {
		out.Cluster = nil
	}
	out.Interactive = in.Interactive
	return nil
}

// Convert_v1_ExecCredentialSpec_To_clientauthentication_ExecCredentialSpec is an autogenerated conversion function.
func Convert_v1_ExecCredentialSpec_To_clientauthentication_ExecCredentialSpec(in *ExecCredentialSpec, out *clientauthentication.ExecCredentialSpec, s conversion.Scope) error {
	return autoConvert_v1_ExecCredentialSpec_To_clientauthentication_ExecCredentialSpec(in, out, s)
}

func autoConvert_clientauthentication_ExecCredentialSpec_To_v1_ExecCredentialSpec(in *clientauthentication.ExecCredentialSpec, out *ExecCredentialSpec, s conversion.Scope) error {
	// WARNING: in.Response requires manual conversion: does not exist in peer-type
	out.Interactive = in.Interactive
	if in.Cluster != nil {
		in, out := &in.Cluster, &out.Cluster
		*out = new(Cluster)
		if err := Convert_clientauthentication_Cluster_To_v1_Cluster(*in, *out, s); err != nil {
			return err
		}
	} else {
		out.Cluster = nil
	}
	return nil
}

func autoConvert_v1_ExecCredentialStatus_To_clientauthentication_ExecCredentialStatus(in *ExecCredentialStatus, out *clientauthentication.ExecCredentialStatus, s conversion.Scope) error {
	out.ExpirationTimestamp = (*metav1.Time)(unsafe.Pointer(in.ExpirationTimestamp))
	out.Token = in.Token
	out.ClientCertificateData = in.ClientCertificateData
	out.ClientKeyData = in.ClientKeyData
	return nil
}

// Convert_v1_ExecCredentialStatus_To_clientauthentication_ExecCredentialStatus is an autogenerated conversion function.
func Convert_v1_ExecCredentialStatus_To_clientauthentication_ExecCredentialStatus(in *ExecCredentialStatus, out *clientauthentication.ExecCredentialStatus, s conversion.Scope) error {
	return autoConvert_v1_ExecCredentialStatus_To_clientauthentication_ExecCredentialStatus(in, out, s)
}

func autoConvert_clientauthentication_ExecCredentialStatus_To_v1_ExecCredentialStatus(in *clientauthentication.ExecCredentialStatus, out *ExecCredentialStatus, s conversion.Scope) error {
	out.ExpirationTimestamp = (*metav1.Time)(unsafe.Pointer(in.ExpirationTimestamp))
	out.Token = in.Token
	out.ClientCertificateData = in.ClientCertificateData
	out.ClientKeyData = in.ClientKeyData
	return nil
}

// Convert_clientauthentication_ExecCredentialStatus_To_v1_ExecCredentialStatus is an autogenerated conversion function.
func Convert_clientauthentication_ExecCredentialStatus_To_v1_ExecCredentialStatus(in *clientauthentication.ExecCredentialStatus, out *ExecCredentialStatus, s conversion.Scope) error {
	return autoConvert_clientauthentication_ExecCredentialStatus_To_v1_ExecCredentialStatus(in, out, s)
}
