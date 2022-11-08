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

package v1beta1

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Connection) DeepCopyInto(out *Connection) {
	*out = *in
	if in.Transport != nil {
		in, out := &in.Transport, &out.Transport
		*out = new(Transport)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Connection.
func (in *Connection) DeepCopy() *Connection {
	if in == nil {
		return nil
	}
	out := new(Connection)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EgressSelection) DeepCopyInto(out *EgressSelection) {
	*out = *in
	in.Connection.DeepCopyInto(&out.Connection)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EgressSelection.
func (in *EgressSelection) DeepCopy() *EgressSelection {
	if in == nil {
		return nil
	}
	out := new(EgressSelection)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EgressSelectorConfiguration) DeepCopyInto(out *EgressSelectorConfiguration) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	if in.EgressSelections != nil {
		in, out := &in.EgressSelections, &out.EgressSelections
		*out = make([]EgressSelection, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EgressSelectorConfiguration.
func (in *EgressSelectorConfiguration) DeepCopy() *EgressSelectorConfiguration {
	if in == nil {
		return nil
	}
	out := new(EgressSelectorConfiguration)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *EgressSelectorConfiguration) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TCPTransport) DeepCopyInto(out *TCPTransport) {
	*out = *in
	if in.TLSConfig != nil {
		in, out := &in.TLSConfig, &out.TLSConfig
		*out = new(TLSConfig)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TCPTransport.
func (in *TCPTransport) DeepCopy() *TCPTransport {
	if in == nil {
		return nil
	}
	out := new(TCPTransport)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TLSConfig) DeepCopyInto(out *TLSConfig) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TLSConfig.
func (in *TLSConfig) DeepCopy() *TLSConfig {
	if in == nil {
		return nil
	}
	out := new(TLSConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TracingConfiguration) DeepCopyInto(out *TracingConfiguration) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	if in.Endpoint != nil {
		in, out := &in.Endpoint, &out.Endpoint
		*out = new(string)
		**out = **in
	}
	if in.SamplingRatePerMillion != nil {
		in, out := &in.SamplingRatePerMillion, &out.SamplingRatePerMillion
		*out = new(int32)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TracingConfiguration.
func (in *TracingConfiguration) DeepCopy() *TracingConfiguration {
	if in == nil {
		return nil
	}
	out := new(TracingConfiguration)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *TracingConfiguration) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Transport) DeepCopyInto(out *Transport) {
	*out = *in
	if in.TCP != nil {
		in, out := &in.TCP, &out.TCP
		*out = new(TCPTransport)
		(*in).DeepCopyInto(*out)
	}
	if in.UDS != nil {
		in, out := &in.UDS, &out.UDS
		*out = new(UDSTransport)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Transport.
func (in *Transport) DeepCopy() *Transport {
	if in == nil {
		return nil
	}
	out := new(Transport)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *UDSTransport) DeepCopyInto(out *UDSTransport) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new UDSTransport.
func (in *UDSTransport) DeepCopy() *UDSTransport {
	if in == nil {
		return nil
	}
	out := new(UDSTransport)
	in.DeepCopyInto(out)
	return out
}
