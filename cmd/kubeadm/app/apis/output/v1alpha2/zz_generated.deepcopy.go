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

// Code generated by deepcopy-gen. DO NOT EDIT.

package v1alpha2

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BootstrapToken) DeepCopyInto(out *BootstrapToken) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.BootstrapToken.DeepCopyInto(&out.BootstrapToken)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BootstrapToken.
func (in *BootstrapToken) DeepCopy() *BootstrapToken {
	if in == nil {
		return nil
	}
	out := new(BootstrapToken)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *BootstrapToken) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ComponentConfigVersionState) DeepCopyInto(out *ComponentConfigVersionState) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ComponentConfigVersionState.
func (in *ComponentConfigVersionState) DeepCopy() *ComponentConfigVersionState {
	if in == nil {
		return nil
	}
	out := new(ComponentConfigVersionState)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ComponentUpgradePlan) DeepCopyInto(out *ComponentUpgradePlan) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ComponentUpgradePlan.
func (in *ComponentUpgradePlan) DeepCopy() *ComponentUpgradePlan {
	if in == nil {
		return nil
	}
	out := new(ComponentUpgradePlan)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ComponentUpgradePlan) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Images) DeepCopyInto(out *Images) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	if in.Images != nil {
		in, out := &in.Images, &out.Images
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Images.
func (in *Images) DeepCopy() *Images {
	if in == nil {
		return nil
	}
	out := new(Images)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Images) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *UpgradePlan) DeepCopyInto(out *UpgradePlan) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	if in.Components != nil {
		in, out := &in.Components, &out.Components
		*out = make([]*ComponentUpgradePlan, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(ComponentUpgradePlan)
				**out = **in
			}
		}
	}
	if in.ConfigVersions != nil {
		in, out := &in.ConfigVersions, &out.ConfigVersions
		*out = make([]ComponentConfigVersionState, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new UpgradePlan.
func (in *UpgradePlan) DeepCopy() *UpgradePlan {
	if in == nil {
		return nil
	}
	out := new(UpgradePlan)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *UpgradePlan) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}
