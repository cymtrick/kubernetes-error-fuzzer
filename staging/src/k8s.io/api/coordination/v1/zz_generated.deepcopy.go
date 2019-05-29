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
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Lease) DeepCopyInto(out *Lease) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Lease.
func (in *Lease) DeepCopy() *Lease {
	if in == nil {
		return nil
	}
	out := new(Lease)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Lease) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LeaseList) DeepCopyInto(out *LeaseList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Lease, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LeaseList.
func (in *LeaseList) DeepCopy() *LeaseList {
	if in == nil {
		return nil
	}
	out := new(LeaseList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *LeaseList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LeaseSpec) DeepCopyInto(out *LeaseSpec) {
	*out = *in
	if in.HolderIdentity != nil {
		in, out := &in.HolderIdentity, &out.HolderIdentity
		*out = new(string)
		**out = **in
	}
	if in.LeaseDurationSeconds != nil {
		in, out := &in.LeaseDurationSeconds, &out.LeaseDurationSeconds
		*out = new(int32)
		**out = **in
	}
	if in.AcquireTime != nil {
		in, out := &in.AcquireTime, &out.AcquireTime
		*out = (*in).DeepCopy()
	}
	if in.RenewTime != nil {
		in, out := &in.RenewTime, &out.RenewTime
		*out = (*in).DeepCopy()
	}
	if in.LeaseTransitions != nil {
		in, out := &in.LeaseTransitions, &out.LeaseTransitions
		*out = new(int32)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LeaseSpec.
func (in *LeaseSpec) DeepCopy() *LeaseSpec {
	if in == nil {
		return nil
	}
	out := new(LeaseSpec)
	in.DeepCopyInto(out)
	return out
}
