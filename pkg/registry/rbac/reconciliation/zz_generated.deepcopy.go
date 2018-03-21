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

package reconciliation

import (
	rbac "k8s.io/kubernetes/pkg/apis/rbac"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterRoleBindingAdapter) DeepCopyInto(out *ClusterRoleBindingAdapter) {
	*out = *in
	if in.ClusterRoleBinding != nil {
		in, out := &in.ClusterRoleBinding, &out.ClusterRoleBinding
		if *in == nil {
			*out = nil
		} else {
			*out = new(rbac.ClusterRoleBinding)
			(*in).DeepCopyInto(*out)
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterRoleBindingAdapter.
func (in *ClusterRoleBindingAdapter) DeepCopy() *ClusterRoleBindingAdapter {
	if in == nil {
		return nil
	}
	out := new(ClusterRoleBindingAdapter)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyRoleBinding is an autogenerated deepcopy function, copying the receiver, creating a new RoleBinding.
func (in ClusterRoleBindingAdapter) DeepCopyRoleBinding() RoleBinding {
	return *in.DeepCopy()
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterRoleRuleOwner) DeepCopyInto(out *ClusterRoleRuleOwner) {
	*out = *in
	if in.ClusterRole != nil {
		in, out := &in.ClusterRole, &out.ClusterRole
		if *in == nil {
			*out = nil
		} else {
			*out = new(rbac.ClusterRole)
			(*in).DeepCopyInto(*out)
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterRoleRuleOwner.
func (in *ClusterRoleRuleOwner) DeepCopy() *ClusterRoleRuleOwner {
	if in == nil {
		return nil
	}
	out := new(ClusterRoleRuleOwner)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyRuleOwner is an autogenerated deepcopy function, copying the receiver, creating a new RuleOwner.
func (in ClusterRoleRuleOwner) DeepCopyRuleOwner() RuleOwner {
	return *in.DeepCopy()
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RoleBindingAdapter) DeepCopyInto(out *RoleBindingAdapter) {
	*out = *in
	if in.RoleBinding != nil {
		in, out := &in.RoleBinding, &out.RoleBinding
		if *in == nil {
			*out = nil
		} else {
			*out = new(rbac.RoleBinding)
			(*in).DeepCopyInto(*out)
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RoleBindingAdapter.
func (in *RoleBindingAdapter) DeepCopy() *RoleBindingAdapter {
	if in == nil {
		return nil
	}
	out := new(RoleBindingAdapter)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyRoleBinding is an autogenerated deepcopy function, copying the receiver, creating a new RoleBinding.
func (in RoleBindingAdapter) DeepCopyRoleBinding() RoleBinding {
	return *in.DeepCopy()
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RoleRuleOwner) DeepCopyInto(out *RoleRuleOwner) {
	*out = *in
	if in.Role != nil {
		in, out := &in.Role, &out.Role
		if *in == nil {
			*out = nil
		} else {
			*out = new(rbac.Role)
			(*in).DeepCopyInto(*out)
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RoleRuleOwner.
func (in *RoleRuleOwner) DeepCopy() *RoleRuleOwner {
	if in == nil {
		return nil
	}
	out := new(RoleRuleOwner)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyRuleOwner is an autogenerated deepcopy function, copying the receiver, creating a new RuleOwner.
func (in RoleRuleOwner) DeepCopyRuleOwner() RuleOwner {
	return *in.DeepCopy()
}
