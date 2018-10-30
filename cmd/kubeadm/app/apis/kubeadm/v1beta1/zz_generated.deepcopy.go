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
	corev1 "k8s.io/api/core/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *APIEndpoint) DeepCopyInto(out *APIEndpoint) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new APIEndpoint.
func (in *APIEndpoint) DeepCopy() *APIEndpoint {
	if in == nil {
		return nil
	}
	out := new(APIEndpoint)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AuditPolicyConfiguration) DeepCopyInto(out *AuditPolicyConfiguration) {
	*out = *in
	if in.LogMaxAge != nil {
		in, out := &in.LogMaxAge, &out.LogMaxAge
		*out = new(int32)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AuditPolicyConfiguration.
func (in *AuditPolicyConfiguration) DeepCopy() *AuditPolicyConfiguration {
	if in == nil {
		return nil
	}
	out := new(AuditPolicyConfiguration)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BootstrapToken) DeepCopyInto(out *BootstrapToken) {
	*out = *in
	if in.Token != nil {
		in, out := &in.Token, &out.Token
		*out = new(BootstrapTokenString)
		**out = **in
	}
	if in.TTL != nil {
		in, out := &in.TTL, &out.TTL
		*out = new(v1.Duration)
		**out = **in
	}
	if in.Expires != nil {
		in, out := &in.Expires, &out.Expires
		*out = (*in).DeepCopy()
	}
	if in.Usages != nil {
		in, out := &in.Usages, &out.Usages
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Groups != nil {
		in, out := &in.Groups, &out.Groups
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
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

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BootstrapTokenDiscovery) DeepCopyInto(out *BootstrapTokenDiscovery) {
	*out = *in
	if in.CACertHashes != nil {
		in, out := &in.CACertHashes, &out.CACertHashes
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BootstrapTokenDiscovery.
func (in *BootstrapTokenDiscovery) DeepCopy() *BootstrapTokenDiscovery {
	if in == nil {
		return nil
	}
	out := new(BootstrapTokenDiscovery)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BootstrapTokenString) DeepCopyInto(out *BootstrapTokenString) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BootstrapTokenString.
func (in *BootstrapTokenString) DeepCopy() *BootstrapTokenString {
	if in == nil {
		return nil
	}
	out := new(BootstrapTokenString)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterConfiguration) DeepCopyInto(out *ClusterConfiguration) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.Etcd.DeepCopyInto(&out.Etcd)
	out.Networking = in.Networking
	if in.APIServerExtraArgs != nil {
		in, out := &in.APIServerExtraArgs, &out.APIServerExtraArgs
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.ControllerManagerExtraArgs != nil {
		in, out := &in.ControllerManagerExtraArgs, &out.ControllerManagerExtraArgs
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.SchedulerExtraArgs != nil {
		in, out := &in.SchedulerExtraArgs, &out.SchedulerExtraArgs
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.APIServerExtraVolumes != nil {
		in, out := &in.APIServerExtraVolumes, &out.APIServerExtraVolumes
		*out = make([]HostPathMount, len(*in))
		copy(*out, *in)
	}
	if in.ControllerManagerExtraVolumes != nil {
		in, out := &in.ControllerManagerExtraVolumes, &out.ControllerManagerExtraVolumes
		*out = make([]HostPathMount, len(*in))
		copy(*out, *in)
	}
	if in.SchedulerExtraVolumes != nil {
		in, out := &in.SchedulerExtraVolumes, &out.SchedulerExtraVolumes
		*out = make([]HostPathMount, len(*in))
		copy(*out, *in)
	}
	if in.APIServerCertSANs != nil {
		in, out := &in.APIServerCertSANs, &out.APIServerCertSANs
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	in.AuditPolicyConfiguration.DeepCopyInto(&out.AuditPolicyConfiguration)
	if in.FeatureGates != nil {
		in, out := &in.FeatureGates, &out.FeatureGates
		*out = make(map[string]bool, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterConfiguration.
func (in *ClusterConfiguration) DeepCopy() *ClusterConfiguration {
	if in == nil {
		return nil
	}
	out := new(ClusterConfiguration)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ClusterConfiguration) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterStatus) DeepCopyInto(out *ClusterStatus) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	if in.APIEndpoints != nil {
		in, out := &in.APIEndpoints, &out.APIEndpoints
		*out = make(map[string]APIEndpoint, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterStatus.
func (in *ClusterStatus) DeepCopy() *ClusterStatus {
	if in == nil {
		return nil
	}
	out := new(ClusterStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ClusterStatus) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Discovery) DeepCopyInto(out *Discovery) {
	*out = *in
	if in.BootstrapToken != nil {
		in, out := &in.BootstrapToken, &out.BootstrapToken
		*out = new(BootstrapTokenDiscovery)
		(*in).DeepCopyInto(*out)
	}
	if in.File != nil {
		in, out := &in.File, &out.File
		*out = new(FileDiscovery)
		**out = **in
	}
	if in.Timeout != nil {
		in, out := &in.Timeout, &out.Timeout
		*out = new(v1.Duration)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Discovery.
func (in *Discovery) DeepCopy() *Discovery {
	if in == nil {
		return nil
	}
	out := new(Discovery)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Etcd) DeepCopyInto(out *Etcd) {
	*out = *in
	if in.Local != nil {
		in, out := &in.Local, &out.Local
		*out = new(LocalEtcd)
		(*in).DeepCopyInto(*out)
	}
	if in.External != nil {
		in, out := &in.External, &out.External
		*out = new(ExternalEtcd)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Etcd.
func (in *Etcd) DeepCopy() *Etcd {
	if in == nil {
		return nil
	}
	out := new(Etcd)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ExternalEtcd) DeepCopyInto(out *ExternalEtcd) {
	*out = *in
	if in.Endpoints != nil {
		in, out := &in.Endpoints, &out.Endpoints
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ExternalEtcd.
func (in *ExternalEtcd) DeepCopy() *ExternalEtcd {
	if in == nil {
		return nil
	}
	out := new(ExternalEtcd)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FileDiscovery) DeepCopyInto(out *FileDiscovery) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FileDiscovery.
func (in *FileDiscovery) DeepCopy() *FileDiscovery {
	if in == nil {
		return nil
	}
	out := new(FileDiscovery)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HostPathMount) DeepCopyInto(out *HostPathMount) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HostPathMount.
func (in *HostPathMount) DeepCopy() *HostPathMount {
	if in == nil {
		return nil
	}
	out := new(HostPathMount)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *InitConfiguration) DeepCopyInto(out *InitConfiguration) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ClusterConfiguration.DeepCopyInto(&out.ClusterConfiguration)
	if in.BootstrapTokens != nil {
		in, out := &in.BootstrapTokens, &out.BootstrapTokens
		*out = make([]BootstrapToken, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	in.NodeRegistration.DeepCopyInto(&out.NodeRegistration)
	out.APIEndpoint = in.APIEndpoint
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new InitConfiguration.
func (in *InitConfiguration) DeepCopy() *InitConfiguration {
	if in == nil {
		return nil
	}
	out := new(InitConfiguration)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *InitConfiguration) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *JoinConfiguration) DeepCopyInto(out *JoinConfiguration) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.NodeRegistration.DeepCopyInto(&out.NodeRegistration)
	in.Discovery.DeepCopyInto(&out.Discovery)
	out.APIEndpoint = in.APIEndpoint
	if in.FeatureGates != nil {
		in, out := &in.FeatureGates, &out.FeatureGates
		*out = make(map[string]bool, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new JoinConfiguration.
func (in *JoinConfiguration) DeepCopy() *JoinConfiguration {
	if in == nil {
		return nil
	}
	out := new(JoinConfiguration)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *JoinConfiguration) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LocalEtcd) DeepCopyInto(out *LocalEtcd) {
	*out = *in
	if in.ExtraArgs != nil {
		in, out := &in.ExtraArgs, &out.ExtraArgs
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.ServerCertSANs != nil {
		in, out := &in.ServerCertSANs, &out.ServerCertSANs
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.PeerCertSANs != nil {
		in, out := &in.PeerCertSANs, &out.PeerCertSANs
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LocalEtcd.
func (in *LocalEtcd) DeepCopy() *LocalEtcd {
	if in == nil {
		return nil
	}
	out := new(LocalEtcd)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Networking) DeepCopyInto(out *Networking) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Networking.
func (in *Networking) DeepCopy() *Networking {
	if in == nil {
		return nil
	}
	out := new(Networking)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NodeRegistrationOptions) DeepCopyInto(out *NodeRegistrationOptions) {
	*out = *in
	if in.Taints != nil {
		in, out := &in.Taints, &out.Taints
		*out = make([]corev1.Taint, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.KubeletExtraArgs != nil {
		in, out := &in.KubeletExtraArgs, &out.KubeletExtraArgs
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NodeRegistrationOptions.
func (in *NodeRegistrationOptions) DeepCopy() *NodeRegistrationOptions {
	if in == nil {
		return nil
	}
	out := new(NodeRegistrationOptions)
	in.DeepCopyInto(out)
	return out
}
