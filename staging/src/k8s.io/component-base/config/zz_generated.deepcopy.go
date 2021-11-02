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

package config

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClientConnectionConfiguration) DeepCopyInto(out *ClientConnectionConfiguration) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClientConnectionConfiguration.
func (in *ClientConnectionConfiguration) DeepCopy() *ClientConnectionConfiguration {
	if in == nil {
		return nil
	}
	out := new(ClientConnectionConfiguration)
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
func (in *FormatOptions) DeepCopyInto(out *FormatOptions) {
	*out = *in
	in.JSON.DeepCopyInto(&out.JSON)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FormatOptions.
func (in *FormatOptions) DeepCopy() *FormatOptions {
	if in == nil {
		return nil
	}
	out := new(FormatOptions)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *JSONOptions) DeepCopyInto(out *JSONOptions) {
	*out = *in
	in.InfoBufferSize.DeepCopyInto(&out.InfoBufferSize)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new JSONOptions.
func (in *JSONOptions) DeepCopy() *JSONOptions {
	if in == nil {
		return nil
	}
	out := new(JSONOptions)
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
func (in *LoggingConfiguration) DeepCopyInto(out *LoggingConfiguration) {
	*out = *in
	if in.VModule != nil {
		in, out := &in.VModule, &out.VModule
		*out = make(VModuleConfiguration, len(*in))
		copy(*out, *in)
	}
	in.Options.DeepCopyInto(&out.Options)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LoggingConfiguration.
func (in *LoggingConfiguration) DeepCopy() *LoggingConfiguration {
	if in == nil {
		return nil
	}
	out := new(LoggingConfiguration)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in VModuleConfiguration) DeepCopyInto(out *VModuleConfiguration) {
	{
		in := &in
		*out = make(VModuleConfiguration, len(*in))
		copy(*out, *in)
		return
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VModuleConfiguration.
func (in VModuleConfiguration) DeepCopy() VModuleConfiguration {
	if in == nil {
		return nil
	}
	out := new(VModuleConfiguration)
	in.DeepCopyInto(out)
	return *out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VModuleItem) DeepCopyInto(out *VModuleItem) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VModuleItem.
func (in *VModuleItem) DeepCopy() *VModuleItem {
	if in == nil {
		return nil
	}
	out := new(VModuleItem)
	in.DeepCopyInto(out)
	return out
}
