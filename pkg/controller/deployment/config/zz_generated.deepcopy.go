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
func (in *DeploymentControllerConfiguration) DeepCopyInto(out *DeploymentControllerConfiguration) {
	*out = *in
	out.DeploymentControllerSyncPeriod = in.DeploymentControllerSyncPeriod
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DeploymentControllerConfiguration.
func (in *DeploymentControllerConfiguration) DeepCopy() *DeploymentControllerConfiguration {
	if in == nil {
		return nil
	}
	out := new(DeploymentControllerConfiguration)
	in.DeepCopyInto(out)
	return out
}
