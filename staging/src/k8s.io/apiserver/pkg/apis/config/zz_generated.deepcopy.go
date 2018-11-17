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

package config

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AESConfiguration) DeepCopyInto(out *AESConfiguration) {
	*out = *in
	if in.Keys != nil {
		in, out := &in.Keys, &out.Keys
		*out = make([]Key, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AESConfiguration.
func (in *AESConfiguration) DeepCopy() *AESConfiguration {
	if in == nil {
		return nil
	}
	out := new(AESConfiguration)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DebuggingConfiguration) DeepCopyInto(out *DebuggingConfiguration) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DebuggingConfiguration.
func (in *DebuggingConfiguration) DeepCopy() *DebuggingConfiguration {
	if in == nil {
		return nil
	}
	out := new(DebuggingConfiguration)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EncryptionConfiguration) DeepCopyInto(out *EncryptionConfiguration) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	if in.Resources != nil {
		in, out := &in.Resources, &out.Resources
		*out = make([]ResourceConfiguration, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EncryptionConfiguration.
func (in *EncryptionConfiguration) DeepCopy() *EncryptionConfiguration {
	if in == nil {
		return nil
	}
	out := new(EncryptionConfiguration)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *EncryptionConfiguration) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IdentityConfiguration) DeepCopyInto(out *IdentityConfiguration) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IdentityConfiguration.
func (in *IdentityConfiguration) DeepCopy() *IdentityConfiguration {
	if in == nil {
		return nil
	}
	out := new(IdentityConfiguration)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KMSConfiguration) DeepCopyInto(out *KMSConfiguration) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KMSConfiguration.
func (in *KMSConfiguration) DeepCopy() *KMSConfiguration {
	if in == nil {
		return nil
	}
	out := new(KMSConfiguration)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Key) DeepCopyInto(out *Key) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Key.
func (in *Key) DeepCopy() *Key {
	if in == nil {
		return nil
	}
	out := new(Key)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LeaderElectionConfiguration) DeepCopyInto(out *LeaderElectionConfiguration) {
	*out = *in
	out.LeaseDuration = in.LeaseDuration
	out.RenewDeadline = in.RenewDeadline
	out.RetryPeriod = in.RetryPeriod
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LeaderElectionConfiguration.
func (in *LeaderElectionConfiguration) DeepCopy() *LeaderElectionConfiguration {
	if in == nil {
		return nil
	}
	out := new(LeaderElectionConfiguration)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ProviderConfiguration) DeepCopyInto(out *ProviderConfiguration) {
	*out = *in
	if in.AESGCM != nil {
		in, out := &in.AESGCM, &out.AESGCM
		*out = new(AESConfiguration)
		(*in).DeepCopyInto(*out)
	}
	if in.AESCBC != nil {
		in, out := &in.AESCBC, &out.AESCBC
		*out = new(AESConfiguration)
		(*in).DeepCopyInto(*out)
	}
	if in.Secretbox != nil {
		in, out := &in.Secretbox, &out.Secretbox
		*out = new(SecretboxConfiguration)
		(*in).DeepCopyInto(*out)
	}
	if in.Identity != nil {
		in, out := &in.Identity, &out.Identity
		*out = new(IdentityConfiguration)
		**out = **in
	}
	if in.KMS != nil {
		in, out := &in.KMS, &out.KMS
		*out = new(KMSConfiguration)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ProviderConfiguration.
func (in *ProviderConfiguration) DeepCopy() *ProviderConfiguration {
	if in == nil {
		return nil
	}
	out := new(ProviderConfiguration)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ResourceConfiguration) DeepCopyInto(out *ResourceConfiguration) {
	*out = *in
	if in.Resources != nil {
		in, out := &in.Resources, &out.Resources
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Providers != nil {
		in, out := &in.Providers, &out.Providers
		*out = make([]ProviderConfiguration, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ResourceConfiguration.
func (in *ResourceConfiguration) DeepCopy() *ResourceConfiguration {
	if in == nil {
		return nil
	}
	out := new(ResourceConfiguration)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SecretboxConfiguration) DeepCopyInto(out *SecretboxConfiguration) {
	*out = *in
	if in.Keys != nil {
		in, out := &in.Keys, &out.Keys
		*out = make([]Key, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SecretboxConfiguration.
func (in *SecretboxConfiguration) DeepCopy() *SecretboxConfiguration {
	if in == nil {
		return nil
	}
	out := new(SecretboxConfiguration)
	in.DeepCopyInto(out)
	return out
}
