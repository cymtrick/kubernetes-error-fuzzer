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

package v1alpha1

import (
	api "k8s.io/kubernetes/pkg/api"
	unversioned "k8s.io/kubernetes/pkg/api/unversioned"
	v1 "k8s.io/kubernetes/pkg/api/v1"
	conversion "k8s.io/kubernetes/pkg/conversion"
)

func init() {
	if err := api.Scheme.AddGeneratedDeepCopyFuncs(
		DeepCopy_v1alpha1_CertificateSigningRequest,
		DeepCopy_v1alpha1_CertificateSigningRequestCondition,
		DeepCopy_v1alpha1_CertificateSigningRequestList,
		DeepCopy_v1alpha1_CertificateSigningRequestSpec,
		DeepCopy_v1alpha1_CertificateSigningRequestStatus,
	); err != nil {
		// if one of the deep copy functions is malformed, detect it immediately.
		panic(err)
	}
}

func DeepCopy_v1alpha1_CertificateSigningRequest(in CertificateSigningRequest, out *CertificateSigningRequest, c *conversion.Cloner) error {
	if err := unversioned.DeepCopy_unversioned_TypeMeta(in.TypeMeta, &out.TypeMeta, c); err != nil {
		return err
	}
	if err := v1.DeepCopy_v1_ObjectMeta(in.ObjectMeta, &out.ObjectMeta, c); err != nil {
		return err
	}
	if err := DeepCopy_v1alpha1_CertificateSigningRequestSpec(in.Spec, &out.Spec, c); err != nil {
		return err
	}
	if err := DeepCopy_v1alpha1_CertificateSigningRequestStatus(in.Status, &out.Status, c); err != nil {
		return err
	}
	return nil
}

func DeepCopy_v1alpha1_CertificateSigningRequestCondition(in CertificateSigningRequestCondition, out *CertificateSigningRequestCondition, c *conversion.Cloner) error {
	out.Type = in.Type
	out.Reason = in.Reason
	out.Message = in.Message
	if err := unversioned.DeepCopy_unversioned_Time(in.LastUpdateTime, &out.LastUpdateTime, c); err != nil {
		return err
	}
	return nil
}

func DeepCopy_v1alpha1_CertificateSigningRequestList(in CertificateSigningRequestList, out *CertificateSigningRequestList, c *conversion.Cloner) error {
	if err := unversioned.DeepCopy_unversioned_TypeMeta(in.TypeMeta, &out.TypeMeta, c); err != nil {
		return err
	}
	if err := unversioned.DeepCopy_unversioned_ListMeta(in.ListMeta, &out.ListMeta, c); err != nil {
		return err
	}
	if in.Items != nil {
		in, out := in.Items, &out.Items
		*out = make([]CertificateSigningRequest, len(in))
		for i := range in {
			if err := DeepCopy_v1alpha1_CertificateSigningRequest(in[i], &(*out)[i], c); err != nil {
				return err
			}
		}
	} else {
		out.Items = nil
	}
	return nil
}

func DeepCopy_v1alpha1_CertificateSigningRequestSpec(in CertificateSigningRequestSpec, out *CertificateSigningRequestSpec, c *conversion.Cloner) error {
	if in.Request != nil {
		in, out := in.Request, &out.Request
		*out = make([]byte, len(in))
		copy(*out, in)
	} else {
		out.Request = nil
	}
	out.Username = in.Username
	out.UID = in.UID
	if in.Groups != nil {
		in, out := in.Groups, &out.Groups
		*out = make([]string, len(in))
		copy(*out, in)
	} else {
		out.Groups = nil
	}
	return nil
}

func DeepCopy_v1alpha1_CertificateSigningRequestStatus(in CertificateSigningRequestStatus, out *CertificateSigningRequestStatus, c *conversion.Cloner) error {
	if in.Conditions != nil {
		in, out := in.Conditions, &out.Conditions
		*out = make([]CertificateSigningRequestCondition, len(in))
		for i := range in {
			if err := DeepCopy_v1alpha1_CertificateSigningRequestCondition(in[i], &(*out)[i], c); err != nil {
				return err
			}
		}
	} else {
		out.Conditions = nil
	}
	if in.Certificate != nil {
		in, out := in.Certificate, &out.Certificate
		*out = make([]byte, len(in))
		copy(*out, in)
	} else {
		out.Certificate = nil
	}
	return nil
}
