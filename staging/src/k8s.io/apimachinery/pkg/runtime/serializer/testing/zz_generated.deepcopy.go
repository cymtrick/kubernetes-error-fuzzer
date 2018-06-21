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

package testing

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ExternalInternalSame) DeepCopyInto(out *ExternalInternalSame) {
	*out = *in
	out.MyWeirdCustomEmbeddedVersionKindField = in.MyWeirdCustomEmbeddedVersionKindField
	out.A = in.A
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ExternalInternalSame.
func (in *ExternalInternalSame) DeepCopy() *ExternalInternalSame {
	if in == nil {
		return nil
	}
	out := new(ExternalInternalSame)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ExternalInternalSame) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ExternalTestType1) DeepCopyInto(out *ExternalTestType1) {
	*out = *in
	out.MyWeirdCustomEmbeddedVersionKindField = in.MyWeirdCustomEmbeddedVersionKindField
	if in.M != nil {
		in, out := &in.M, &out.M
		*out = make(map[string]int, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.N != nil {
		in, out := &in.N, &out.N
		*out = make(map[string]ExternalTestType2, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.O != nil {
		in, out := &in.O, &out.O
		*out = new(ExternalTestType2)
		**out = **in
	}
	if in.P != nil {
		in, out := &in.P, &out.P
		*out = make([]ExternalTestType2, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ExternalTestType1.
func (in *ExternalTestType1) DeepCopy() *ExternalTestType1 {
	if in == nil {
		return nil
	}
	out := new(ExternalTestType1)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ExternalTestType1) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ExternalTestType2) DeepCopyInto(out *ExternalTestType2) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ExternalTestType2.
func (in *ExternalTestType2) DeepCopy() *ExternalTestType2 {
	if in == nil {
		return nil
	}
	out := new(ExternalTestType2)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ExternalTestType2) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TestType1) DeepCopyInto(out *TestType1) {
	*out = *in
	out.MyWeirdCustomEmbeddedVersionKindField = in.MyWeirdCustomEmbeddedVersionKindField
	if in.M != nil {
		in, out := &in.M, &out.M
		*out = make(map[string]int, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.N != nil {
		in, out := &in.N, &out.N
		*out = make(map[string]TestType2, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.O != nil {
		in, out := &in.O, &out.O
		*out = new(TestType2)
		**out = **in
	}
	if in.P != nil {
		in, out := &in.P, &out.P
		*out = make([]TestType2, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TestType1.
func (in *TestType1) DeepCopy() *TestType1 {
	if in == nil {
		return nil
	}
	out := new(TestType1)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *TestType1) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TestType2) DeepCopyInto(out *TestType2) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TestType2.
func (in *TestType2) DeepCopy() *TestType2 {
	if in == nil {
		return nil
	}
	out := new(TestType2)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *TestType2) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}
