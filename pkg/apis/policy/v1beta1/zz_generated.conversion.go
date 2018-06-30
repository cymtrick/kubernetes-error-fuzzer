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

// Code generated by conversion-gen. DO NOT EDIT.

package v1beta1

import (
	unsafe "unsafe"

	corev1 "k8s.io/api/core/v1"
	v1beta1 "k8s.io/api/policy/v1beta1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	conversion "k8s.io/apimachinery/pkg/conversion"
	runtime "k8s.io/apimachinery/pkg/runtime"
	intstr "k8s.io/apimachinery/pkg/util/intstr"
	core "k8s.io/kubernetes/pkg/apis/core"
	policy "k8s.io/kubernetes/pkg/apis/policy"
)

func init() {
	localSchemeBuilder.Register(RegisterConversions)
}

// RegisterConversions adds conversion functions to the given scheme.
// Public to allow building arbitrary schemes.
func RegisterConversions(scheme *runtime.Scheme) error {
	return scheme.AddGeneratedConversionFuncs(
		Convert_v1beta1_AllowedFlexVolume_To_policy_AllowedFlexVolume,
		Convert_policy_AllowedFlexVolume_To_v1beta1_AllowedFlexVolume,
		Convert_v1beta1_AllowedHostPath_To_policy_AllowedHostPath,
		Convert_policy_AllowedHostPath_To_v1beta1_AllowedHostPath,
		Convert_v1beta1_Eviction_To_policy_Eviction,
		Convert_policy_Eviction_To_v1beta1_Eviction,
		Convert_v1beta1_FSGroupStrategyOptions_To_policy_FSGroupStrategyOptions,
		Convert_policy_FSGroupStrategyOptions_To_v1beta1_FSGroupStrategyOptions,
		Convert_v1beta1_HostPortRange_To_policy_HostPortRange,
		Convert_policy_HostPortRange_To_v1beta1_HostPortRange,
		Convert_v1beta1_IDRange_To_policy_IDRange,
		Convert_policy_IDRange_To_v1beta1_IDRange,
		Convert_v1beta1_PodDisruptionBudget_To_policy_PodDisruptionBudget,
		Convert_policy_PodDisruptionBudget_To_v1beta1_PodDisruptionBudget,
		Convert_v1beta1_PodDisruptionBudgetList_To_policy_PodDisruptionBudgetList,
		Convert_policy_PodDisruptionBudgetList_To_v1beta1_PodDisruptionBudgetList,
		Convert_v1beta1_PodDisruptionBudgetSpec_To_policy_PodDisruptionBudgetSpec,
		Convert_policy_PodDisruptionBudgetSpec_To_v1beta1_PodDisruptionBudgetSpec,
		Convert_v1beta1_PodDisruptionBudgetStatus_To_policy_PodDisruptionBudgetStatus,
		Convert_policy_PodDisruptionBudgetStatus_To_v1beta1_PodDisruptionBudgetStatus,
		Convert_v1beta1_PodSecurityPolicy_To_policy_PodSecurityPolicy,
		Convert_policy_PodSecurityPolicy_To_v1beta1_PodSecurityPolicy,
		Convert_v1beta1_PodSecurityPolicyList_To_policy_PodSecurityPolicyList,
		Convert_policy_PodSecurityPolicyList_To_v1beta1_PodSecurityPolicyList,
		Convert_v1beta1_PodSecurityPolicySpec_To_policy_PodSecurityPolicySpec,
		Convert_policy_PodSecurityPolicySpec_To_v1beta1_PodSecurityPolicySpec,
		Convert_v1beta1_RunAsUserStrategyOptions_To_policy_RunAsUserStrategyOptions,
		Convert_policy_RunAsUserStrategyOptions_To_v1beta1_RunAsUserStrategyOptions,
		Convert_v1beta1_SELinuxStrategyOptions_To_policy_SELinuxStrategyOptions,
		Convert_policy_SELinuxStrategyOptions_To_v1beta1_SELinuxStrategyOptions,
		Convert_v1beta1_SupplementalGroupsStrategyOptions_To_policy_SupplementalGroupsStrategyOptions,
		Convert_policy_SupplementalGroupsStrategyOptions_To_v1beta1_SupplementalGroupsStrategyOptions,
	)
}

func autoConvert_v1beta1_AllowedFlexVolume_To_policy_AllowedFlexVolume(in *v1beta1.AllowedFlexVolume, out *policy.AllowedFlexVolume, s conversion.Scope) error {
	out.Driver = in.Driver
	return nil
}

// Convert_v1beta1_AllowedFlexVolume_To_policy_AllowedFlexVolume is an autogenerated conversion function.
func Convert_v1beta1_AllowedFlexVolume_To_policy_AllowedFlexVolume(in *v1beta1.AllowedFlexVolume, out *policy.AllowedFlexVolume, s conversion.Scope) error {
	return autoConvert_v1beta1_AllowedFlexVolume_To_policy_AllowedFlexVolume(in, out, s)
}

func autoConvert_policy_AllowedFlexVolume_To_v1beta1_AllowedFlexVolume(in *policy.AllowedFlexVolume, out *v1beta1.AllowedFlexVolume, s conversion.Scope) error {
	out.Driver = in.Driver
	return nil
}

// Convert_policy_AllowedFlexVolume_To_v1beta1_AllowedFlexVolume is an autogenerated conversion function.
func Convert_policy_AllowedFlexVolume_To_v1beta1_AllowedFlexVolume(in *policy.AllowedFlexVolume, out *v1beta1.AllowedFlexVolume, s conversion.Scope) error {
	return autoConvert_policy_AllowedFlexVolume_To_v1beta1_AllowedFlexVolume(in, out, s)
}

func autoConvert_v1beta1_AllowedHostPath_To_policy_AllowedHostPath(in *v1beta1.AllowedHostPath, out *policy.AllowedHostPath, s conversion.Scope) error {
	out.PathPrefix = in.PathPrefix
	out.ReadOnly = in.ReadOnly
	return nil
}

// Convert_v1beta1_AllowedHostPath_To_policy_AllowedHostPath is an autogenerated conversion function.
func Convert_v1beta1_AllowedHostPath_To_policy_AllowedHostPath(in *v1beta1.AllowedHostPath, out *policy.AllowedHostPath, s conversion.Scope) error {
	return autoConvert_v1beta1_AllowedHostPath_To_policy_AllowedHostPath(in, out, s)
}

func autoConvert_policy_AllowedHostPath_To_v1beta1_AllowedHostPath(in *policy.AllowedHostPath, out *v1beta1.AllowedHostPath, s conversion.Scope) error {
	out.PathPrefix = in.PathPrefix
	out.ReadOnly = in.ReadOnly
	return nil
}

// Convert_policy_AllowedHostPath_To_v1beta1_AllowedHostPath is an autogenerated conversion function.
func Convert_policy_AllowedHostPath_To_v1beta1_AllowedHostPath(in *policy.AllowedHostPath, out *v1beta1.AllowedHostPath, s conversion.Scope) error {
	return autoConvert_policy_AllowedHostPath_To_v1beta1_AllowedHostPath(in, out, s)
}

func autoConvert_v1beta1_Eviction_To_policy_Eviction(in *v1beta1.Eviction, out *policy.Eviction, s conversion.Scope) error {
	out.ObjectMeta = in.ObjectMeta
	out.DeleteOptions = (*v1.DeleteOptions)(unsafe.Pointer(in.DeleteOptions))
	return nil
}

// Convert_v1beta1_Eviction_To_policy_Eviction is an autogenerated conversion function.
func Convert_v1beta1_Eviction_To_policy_Eviction(in *v1beta1.Eviction, out *policy.Eviction, s conversion.Scope) error {
	return autoConvert_v1beta1_Eviction_To_policy_Eviction(in, out, s)
}

func autoConvert_policy_Eviction_To_v1beta1_Eviction(in *policy.Eviction, out *v1beta1.Eviction, s conversion.Scope) error {
	out.ObjectMeta = in.ObjectMeta
	out.DeleteOptions = (*v1.DeleteOptions)(unsafe.Pointer(in.DeleteOptions))
	return nil
}

// Convert_policy_Eviction_To_v1beta1_Eviction is an autogenerated conversion function.
func Convert_policy_Eviction_To_v1beta1_Eviction(in *policy.Eviction, out *v1beta1.Eviction, s conversion.Scope) error {
	return autoConvert_policy_Eviction_To_v1beta1_Eviction(in, out, s)
}

func autoConvert_v1beta1_FSGroupStrategyOptions_To_policy_FSGroupStrategyOptions(in *v1beta1.FSGroupStrategyOptions, out *policy.FSGroupStrategyOptions, s conversion.Scope) error {
	out.Rule = policy.FSGroupStrategyType(in.Rule)
	out.Ranges = *(*[]policy.IDRange)(unsafe.Pointer(&in.Ranges))
	return nil
}

// Convert_v1beta1_FSGroupStrategyOptions_To_policy_FSGroupStrategyOptions is an autogenerated conversion function.
func Convert_v1beta1_FSGroupStrategyOptions_To_policy_FSGroupStrategyOptions(in *v1beta1.FSGroupStrategyOptions, out *policy.FSGroupStrategyOptions, s conversion.Scope) error {
	return autoConvert_v1beta1_FSGroupStrategyOptions_To_policy_FSGroupStrategyOptions(in, out, s)
}

func autoConvert_policy_FSGroupStrategyOptions_To_v1beta1_FSGroupStrategyOptions(in *policy.FSGroupStrategyOptions, out *v1beta1.FSGroupStrategyOptions, s conversion.Scope) error {
	out.Rule = v1beta1.FSGroupStrategyType(in.Rule)
	out.Ranges = *(*[]v1beta1.IDRange)(unsafe.Pointer(&in.Ranges))
	return nil
}

// Convert_policy_FSGroupStrategyOptions_To_v1beta1_FSGroupStrategyOptions is an autogenerated conversion function.
func Convert_policy_FSGroupStrategyOptions_To_v1beta1_FSGroupStrategyOptions(in *policy.FSGroupStrategyOptions, out *v1beta1.FSGroupStrategyOptions, s conversion.Scope) error {
	return autoConvert_policy_FSGroupStrategyOptions_To_v1beta1_FSGroupStrategyOptions(in, out, s)
}

func autoConvert_v1beta1_HostPortRange_To_policy_HostPortRange(in *v1beta1.HostPortRange, out *policy.HostPortRange, s conversion.Scope) error {
	out.Min = in.Min
	out.Max = in.Max
	return nil
}

// Convert_v1beta1_HostPortRange_To_policy_HostPortRange is an autogenerated conversion function.
func Convert_v1beta1_HostPortRange_To_policy_HostPortRange(in *v1beta1.HostPortRange, out *policy.HostPortRange, s conversion.Scope) error {
	return autoConvert_v1beta1_HostPortRange_To_policy_HostPortRange(in, out, s)
}

func autoConvert_policy_HostPortRange_To_v1beta1_HostPortRange(in *policy.HostPortRange, out *v1beta1.HostPortRange, s conversion.Scope) error {
	out.Min = in.Min
	out.Max = in.Max
	return nil
}

// Convert_policy_HostPortRange_To_v1beta1_HostPortRange is an autogenerated conversion function.
func Convert_policy_HostPortRange_To_v1beta1_HostPortRange(in *policy.HostPortRange, out *v1beta1.HostPortRange, s conversion.Scope) error {
	return autoConvert_policy_HostPortRange_To_v1beta1_HostPortRange(in, out, s)
}

func autoConvert_v1beta1_IDRange_To_policy_IDRange(in *v1beta1.IDRange, out *policy.IDRange, s conversion.Scope) error {
	out.Min = in.Min
	out.Max = in.Max
	return nil
}

// Convert_v1beta1_IDRange_To_policy_IDRange is an autogenerated conversion function.
func Convert_v1beta1_IDRange_To_policy_IDRange(in *v1beta1.IDRange, out *policy.IDRange, s conversion.Scope) error {
	return autoConvert_v1beta1_IDRange_To_policy_IDRange(in, out, s)
}

func autoConvert_policy_IDRange_To_v1beta1_IDRange(in *policy.IDRange, out *v1beta1.IDRange, s conversion.Scope) error {
	out.Min = in.Min
	out.Max = in.Max
	return nil
}

// Convert_policy_IDRange_To_v1beta1_IDRange is an autogenerated conversion function.
func Convert_policy_IDRange_To_v1beta1_IDRange(in *policy.IDRange, out *v1beta1.IDRange, s conversion.Scope) error {
	return autoConvert_policy_IDRange_To_v1beta1_IDRange(in, out, s)
}

func autoConvert_v1beta1_PodDisruptionBudget_To_policy_PodDisruptionBudget(in *v1beta1.PodDisruptionBudget, out *policy.PodDisruptionBudget, s conversion.Scope) error {
	out.ObjectMeta = in.ObjectMeta
	if err := Convert_v1beta1_PodDisruptionBudgetSpec_To_policy_PodDisruptionBudgetSpec(&in.Spec, &out.Spec, s); err != nil {
		return err
	}
	if err := Convert_v1beta1_PodDisruptionBudgetStatus_To_policy_PodDisruptionBudgetStatus(&in.Status, &out.Status, s); err != nil {
		return err
	}
	return nil
}

// Convert_v1beta1_PodDisruptionBudget_To_policy_PodDisruptionBudget is an autogenerated conversion function.
func Convert_v1beta1_PodDisruptionBudget_To_policy_PodDisruptionBudget(in *v1beta1.PodDisruptionBudget, out *policy.PodDisruptionBudget, s conversion.Scope) error {
	return autoConvert_v1beta1_PodDisruptionBudget_To_policy_PodDisruptionBudget(in, out, s)
}

func autoConvert_policy_PodDisruptionBudget_To_v1beta1_PodDisruptionBudget(in *policy.PodDisruptionBudget, out *v1beta1.PodDisruptionBudget, s conversion.Scope) error {
	out.ObjectMeta = in.ObjectMeta
	if err := Convert_policy_PodDisruptionBudgetSpec_To_v1beta1_PodDisruptionBudgetSpec(&in.Spec, &out.Spec, s); err != nil {
		return err
	}
	if err := Convert_policy_PodDisruptionBudgetStatus_To_v1beta1_PodDisruptionBudgetStatus(&in.Status, &out.Status, s); err != nil {
		return err
	}
	return nil
}

// Convert_policy_PodDisruptionBudget_To_v1beta1_PodDisruptionBudget is an autogenerated conversion function.
func Convert_policy_PodDisruptionBudget_To_v1beta1_PodDisruptionBudget(in *policy.PodDisruptionBudget, out *v1beta1.PodDisruptionBudget, s conversion.Scope) error {
	return autoConvert_policy_PodDisruptionBudget_To_v1beta1_PodDisruptionBudget(in, out, s)
}

func autoConvert_v1beta1_PodDisruptionBudgetList_To_policy_PodDisruptionBudgetList(in *v1beta1.PodDisruptionBudgetList, out *policy.PodDisruptionBudgetList, s conversion.Scope) error {
	out.ListMeta = in.ListMeta
	out.Items = *(*[]policy.PodDisruptionBudget)(unsafe.Pointer(&in.Items))
	return nil
}

// Convert_v1beta1_PodDisruptionBudgetList_To_policy_PodDisruptionBudgetList is an autogenerated conversion function.
func Convert_v1beta1_PodDisruptionBudgetList_To_policy_PodDisruptionBudgetList(in *v1beta1.PodDisruptionBudgetList, out *policy.PodDisruptionBudgetList, s conversion.Scope) error {
	return autoConvert_v1beta1_PodDisruptionBudgetList_To_policy_PodDisruptionBudgetList(in, out, s)
}

func autoConvert_policy_PodDisruptionBudgetList_To_v1beta1_PodDisruptionBudgetList(in *policy.PodDisruptionBudgetList, out *v1beta1.PodDisruptionBudgetList, s conversion.Scope) error {
	out.ListMeta = in.ListMeta
	out.Items = *(*[]v1beta1.PodDisruptionBudget)(unsafe.Pointer(&in.Items))
	return nil
}

// Convert_policy_PodDisruptionBudgetList_To_v1beta1_PodDisruptionBudgetList is an autogenerated conversion function.
func Convert_policy_PodDisruptionBudgetList_To_v1beta1_PodDisruptionBudgetList(in *policy.PodDisruptionBudgetList, out *v1beta1.PodDisruptionBudgetList, s conversion.Scope) error {
	return autoConvert_policy_PodDisruptionBudgetList_To_v1beta1_PodDisruptionBudgetList(in, out, s)
}

func autoConvert_v1beta1_PodDisruptionBudgetSpec_To_policy_PodDisruptionBudgetSpec(in *v1beta1.PodDisruptionBudgetSpec, out *policy.PodDisruptionBudgetSpec, s conversion.Scope) error {
	out.MinAvailable = (*intstr.IntOrString)(unsafe.Pointer(in.MinAvailable))
	out.Selector = (*v1.LabelSelector)(unsafe.Pointer(in.Selector))
	out.MaxUnavailable = (*intstr.IntOrString)(unsafe.Pointer(in.MaxUnavailable))
	return nil
}

// Convert_v1beta1_PodDisruptionBudgetSpec_To_policy_PodDisruptionBudgetSpec is an autogenerated conversion function.
func Convert_v1beta1_PodDisruptionBudgetSpec_To_policy_PodDisruptionBudgetSpec(in *v1beta1.PodDisruptionBudgetSpec, out *policy.PodDisruptionBudgetSpec, s conversion.Scope) error {
	return autoConvert_v1beta1_PodDisruptionBudgetSpec_To_policy_PodDisruptionBudgetSpec(in, out, s)
}

func autoConvert_policy_PodDisruptionBudgetSpec_To_v1beta1_PodDisruptionBudgetSpec(in *policy.PodDisruptionBudgetSpec, out *v1beta1.PodDisruptionBudgetSpec, s conversion.Scope) error {
	out.MinAvailable = (*intstr.IntOrString)(unsafe.Pointer(in.MinAvailable))
	out.Selector = (*v1.LabelSelector)(unsafe.Pointer(in.Selector))
	out.MaxUnavailable = (*intstr.IntOrString)(unsafe.Pointer(in.MaxUnavailable))
	return nil
}

// Convert_policy_PodDisruptionBudgetSpec_To_v1beta1_PodDisruptionBudgetSpec is an autogenerated conversion function.
func Convert_policy_PodDisruptionBudgetSpec_To_v1beta1_PodDisruptionBudgetSpec(in *policy.PodDisruptionBudgetSpec, out *v1beta1.PodDisruptionBudgetSpec, s conversion.Scope) error {
	return autoConvert_policy_PodDisruptionBudgetSpec_To_v1beta1_PodDisruptionBudgetSpec(in, out, s)
}

func autoConvert_v1beta1_PodDisruptionBudgetStatus_To_policy_PodDisruptionBudgetStatus(in *v1beta1.PodDisruptionBudgetStatus, out *policy.PodDisruptionBudgetStatus, s conversion.Scope) error {
	out.ObservedGeneration = in.ObservedGeneration
	out.DisruptedPods = *(*map[string]v1.Time)(unsafe.Pointer(&in.DisruptedPods))
	out.PodDisruptionsAllowed = in.PodDisruptionsAllowed
	out.CurrentHealthy = in.CurrentHealthy
	out.DesiredHealthy = in.DesiredHealthy
	out.ExpectedPods = in.ExpectedPods
	return nil
}

// Convert_v1beta1_PodDisruptionBudgetStatus_To_policy_PodDisruptionBudgetStatus is an autogenerated conversion function.
func Convert_v1beta1_PodDisruptionBudgetStatus_To_policy_PodDisruptionBudgetStatus(in *v1beta1.PodDisruptionBudgetStatus, out *policy.PodDisruptionBudgetStatus, s conversion.Scope) error {
	return autoConvert_v1beta1_PodDisruptionBudgetStatus_To_policy_PodDisruptionBudgetStatus(in, out, s)
}

func autoConvert_policy_PodDisruptionBudgetStatus_To_v1beta1_PodDisruptionBudgetStatus(in *policy.PodDisruptionBudgetStatus, out *v1beta1.PodDisruptionBudgetStatus, s conversion.Scope) error {
	out.ObservedGeneration = in.ObservedGeneration
	out.DisruptedPods = *(*map[string]v1.Time)(unsafe.Pointer(&in.DisruptedPods))
	out.PodDisruptionsAllowed = in.PodDisruptionsAllowed
	out.CurrentHealthy = in.CurrentHealthy
	out.DesiredHealthy = in.DesiredHealthy
	out.ExpectedPods = in.ExpectedPods
	return nil
}

// Convert_policy_PodDisruptionBudgetStatus_To_v1beta1_PodDisruptionBudgetStatus is an autogenerated conversion function.
func Convert_policy_PodDisruptionBudgetStatus_To_v1beta1_PodDisruptionBudgetStatus(in *policy.PodDisruptionBudgetStatus, out *v1beta1.PodDisruptionBudgetStatus, s conversion.Scope) error {
	return autoConvert_policy_PodDisruptionBudgetStatus_To_v1beta1_PodDisruptionBudgetStatus(in, out, s)
}

func autoConvert_v1beta1_PodSecurityPolicy_To_policy_PodSecurityPolicy(in *v1beta1.PodSecurityPolicy, out *policy.PodSecurityPolicy, s conversion.Scope) error {
	out.ObjectMeta = in.ObjectMeta
	if err := Convert_v1beta1_PodSecurityPolicySpec_To_policy_PodSecurityPolicySpec(&in.Spec, &out.Spec, s); err != nil {
		return err
	}
	return nil
}

// Convert_v1beta1_PodSecurityPolicy_To_policy_PodSecurityPolicy is an autogenerated conversion function.
func Convert_v1beta1_PodSecurityPolicy_To_policy_PodSecurityPolicy(in *v1beta1.PodSecurityPolicy, out *policy.PodSecurityPolicy, s conversion.Scope) error {
	return autoConvert_v1beta1_PodSecurityPolicy_To_policy_PodSecurityPolicy(in, out, s)
}

func autoConvert_policy_PodSecurityPolicy_To_v1beta1_PodSecurityPolicy(in *policy.PodSecurityPolicy, out *v1beta1.PodSecurityPolicy, s conversion.Scope) error {
	out.ObjectMeta = in.ObjectMeta
	if err := Convert_policy_PodSecurityPolicySpec_To_v1beta1_PodSecurityPolicySpec(&in.Spec, &out.Spec, s); err != nil {
		return err
	}
	return nil
}

// Convert_policy_PodSecurityPolicy_To_v1beta1_PodSecurityPolicy is an autogenerated conversion function.
func Convert_policy_PodSecurityPolicy_To_v1beta1_PodSecurityPolicy(in *policy.PodSecurityPolicy, out *v1beta1.PodSecurityPolicy, s conversion.Scope) error {
	return autoConvert_policy_PodSecurityPolicy_To_v1beta1_PodSecurityPolicy(in, out, s)
}

func autoConvert_v1beta1_PodSecurityPolicyList_To_policy_PodSecurityPolicyList(in *v1beta1.PodSecurityPolicyList, out *policy.PodSecurityPolicyList, s conversion.Scope) error {
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]policy.PodSecurityPolicy, len(*in))
		for i := range *in {
			if err := Convert_v1beta1_PodSecurityPolicy_To_policy_PodSecurityPolicy(&(*in)[i], &(*out)[i], s); err != nil {
				return err
			}
		}
	} else {
		out.Items = nil
	}
	return nil
}

// Convert_v1beta1_PodSecurityPolicyList_To_policy_PodSecurityPolicyList is an autogenerated conversion function.
func Convert_v1beta1_PodSecurityPolicyList_To_policy_PodSecurityPolicyList(in *v1beta1.PodSecurityPolicyList, out *policy.PodSecurityPolicyList, s conversion.Scope) error {
	return autoConvert_v1beta1_PodSecurityPolicyList_To_policy_PodSecurityPolicyList(in, out, s)
}

func autoConvert_policy_PodSecurityPolicyList_To_v1beta1_PodSecurityPolicyList(in *policy.PodSecurityPolicyList, out *v1beta1.PodSecurityPolicyList, s conversion.Scope) error {
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]v1beta1.PodSecurityPolicy, len(*in))
		for i := range *in {
			if err := Convert_policy_PodSecurityPolicy_To_v1beta1_PodSecurityPolicy(&(*in)[i], &(*out)[i], s); err != nil {
				return err
			}
		}
	} else {
		out.Items = nil
	}
	return nil
}

// Convert_policy_PodSecurityPolicyList_To_v1beta1_PodSecurityPolicyList is an autogenerated conversion function.
func Convert_policy_PodSecurityPolicyList_To_v1beta1_PodSecurityPolicyList(in *policy.PodSecurityPolicyList, out *v1beta1.PodSecurityPolicyList, s conversion.Scope) error {
	return autoConvert_policy_PodSecurityPolicyList_To_v1beta1_PodSecurityPolicyList(in, out, s)
}

func autoConvert_v1beta1_PodSecurityPolicySpec_To_policy_PodSecurityPolicySpec(in *v1beta1.PodSecurityPolicySpec, out *policy.PodSecurityPolicySpec, s conversion.Scope) error {
	out.Privileged = in.Privileged
	out.DefaultAddCapabilities = *(*[]core.Capability)(unsafe.Pointer(&in.DefaultAddCapabilities))
	out.RequiredDropCapabilities = *(*[]core.Capability)(unsafe.Pointer(&in.RequiredDropCapabilities))
	out.AllowedCapabilities = *(*[]core.Capability)(unsafe.Pointer(&in.AllowedCapabilities))
	out.Volumes = *(*[]policy.FSType)(unsafe.Pointer(&in.Volumes))
	out.HostNetwork = in.HostNetwork
	out.HostPorts = *(*[]policy.HostPortRange)(unsafe.Pointer(&in.HostPorts))
	out.HostPID = in.HostPID
	out.HostIPC = in.HostIPC
	if err := Convert_v1beta1_SELinuxStrategyOptions_To_policy_SELinuxStrategyOptions(&in.SELinux, &out.SELinux, s); err != nil {
		return err
	}
	if err := Convert_v1beta1_RunAsUserStrategyOptions_To_policy_RunAsUserStrategyOptions(&in.RunAsUser, &out.RunAsUser, s); err != nil {
		return err
	}
	if err := Convert_v1beta1_SupplementalGroupsStrategyOptions_To_policy_SupplementalGroupsStrategyOptions(&in.SupplementalGroups, &out.SupplementalGroups, s); err != nil {
		return err
	}
	if err := Convert_v1beta1_FSGroupStrategyOptions_To_policy_FSGroupStrategyOptions(&in.FSGroup, &out.FSGroup, s); err != nil {
		return err
	}
	out.ReadOnlyRootFilesystem = in.ReadOnlyRootFilesystem
	out.DefaultAllowPrivilegeEscalation = (*bool)(unsafe.Pointer(in.DefaultAllowPrivilegeEscalation))
	if err := v1.Convert_Pointer_bool_To_bool(&in.AllowPrivilegeEscalation, &out.AllowPrivilegeEscalation, s); err != nil {
		return err
	}
	out.AllowedHostPaths = *(*[]policy.AllowedHostPath)(unsafe.Pointer(&in.AllowedHostPaths))
	out.AllowedFlexVolumes = *(*[]policy.AllowedFlexVolume)(unsafe.Pointer(&in.AllowedFlexVolumes))
	out.AllowedUnsafeSysctls = *(*[]string)(unsafe.Pointer(&in.AllowedUnsafeSysctls))
	out.ForbiddenSysctls = *(*[]string)(unsafe.Pointer(&in.ForbiddenSysctls))
	return nil
}

// Convert_v1beta1_PodSecurityPolicySpec_To_policy_PodSecurityPolicySpec is an autogenerated conversion function.
func Convert_v1beta1_PodSecurityPolicySpec_To_policy_PodSecurityPolicySpec(in *v1beta1.PodSecurityPolicySpec, out *policy.PodSecurityPolicySpec, s conversion.Scope) error {
	return autoConvert_v1beta1_PodSecurityPolicySpec_To_policy_PodSecurityPolicySpec(in, out, s)
}

func autoConvert_policy_PodSecurityPolicySpec_To_v1beta1_PodSecurityPolicySpec(in *policy.PodSecurityPolicySpec, out *v1beta1.PodSecurityPolicySpec, s conversion.Scope) error {
	out.Privileged = in.Privileged
	out.DefaultAddCapabilities = *(*[]corev1.Capability)(unsafe.Pointer(&in.DefaultAddCapabilities))
	out.RequiredDropCapabilities = *(*[]corev1.Capability)(unsafe.Pointer(&in.RequiredDropCapabilities))
	out.AllowedCapabilities = *(*[]corev1.Capability)(unsafe.Pointer(&in.AllowedCapabilities))
	out.Volumes = *(*[]v1beta1.FSType)(unsafe.Pointer(&in.Volumes))
	out.HostNetwork = in.HostNetwork
	out.HostPorts = *(*[]v1beta1.HostPortRange)(unsafe.Pointer(&in.HostPorts))
	out.HostPID = in.HostPID
	out.HostIPC = in.HostIPC
	if err := Convert_policy_SELinuxStrategyOptions_To_v1beta1_SELinuxStrategyOptions(&in.SELinux, &out.SELinux, s); err != nil {
		return err
	}
	if err := Convert_policy_RunAsUserStrategyOptions_To_v1beta1_RunAsUserStrategyOptions(&in.RunAsUser, &out.RunAsUser, s); err != nil {
		return err
	}
	if err := Convert_policy_SupplementalGroupsStrategyOptions_To_v1beta1_SupplementalGroupsStrategyOptions(&in.SupplementalGroups, &out.SupplementalGroups, s); err != nil {
		return err
	}
	if err := Convert_policy_FSGroupStrategyOptions_To_v1beta1_FSGroupStrategyOptions(&in.FSGroup, &out.FSGroup, s); err != nil {
		return err
	}
	out.ReadOnlyRootFilesystem = in.ReadOnlyRootFilesystem
	out.DefaultAllowPrivilegeEscalation = (*bool)(unsafe.Pointer(in.DefaultAllowPrivilegeEscalation))
	if err := v1.Convert_bool_To_Pointer_bool(&in.AllowPrivilegeEscalation, &out.AllowPrivilegeEscalation, s); err != nil {
		return err
	}
	out.AllowedHostPaths = *(*[]v1beta1.AllowedHostPath)(unsafe.Pointer(&in.AllowedHostPaths))
	out.AllowedFlexVolumes = *(*[]v1beta1.AllowedFlexVolume)(unsafe.Pointer(&in.AllowedFlexVolumes))
	out.AllowedUnsafeSysctls = *(*[]string)(unsafe.Pointer(&in.AllowedUnsafeSysctls))
	out.ForbiddenSysctls = *(*[]string)(unsafe.Pointer(&in.ForbiddenSysctls))
	return nil
}

// Convert_policy_PodSecurityPolicySpec_To_v1beta1_PodSecurityPolicySpec is an autogenerated conversion function.
func Convert_policy_PodSecurityPolicySpec_To_v1beta1_PodSecurityPolicySpec(in *policy.PodSecurityPolicySpec, out *v1beta1.PodSecurityPolicySpec, s conversion.Scope) error {
	return autoConvert_policy_PodSecurityPolicySpec_To_v1beta1_PodSecurityPolicySpec(in, out, s)
}

func autoConvert_v1beta1_RunAsUserStrategyOptions_To_policy_RunAsUserStrategyOptions(in *v1beta1.RunAsUserStrategyOptions, out *policy.RunAsUserStrategyOptions, s conversion.Scope) error {
	out.Rule = policy.RunAsUserStrategy(in.Rule)
	out.Ranges = *(*[]policy.IDRange)(unsafe.Pointer(&in.Ranges))
	return nil
}

// Convert_v1beta1_RunAsUserStrategyOptions_To_policy_RunAsUserStrategyOptions is an autogenerated conversion function.
func Convert_v1beta1_RunAsUserStrategyOptions_To_policy_RunAsUserStrategyOptions(in *v1beta1.RunAsUserStrategyOptions, out *policy.RunAsUserStrategyOptions, s conversion.Scope) error {
	return autoConvert_v1beta1_RunAsUserStrategyOptions_To_policy_RunAsUserStrategyOptions(in, out, s)
}

func autoConvert_policy_RunAsUserStrategyOptions_To_v1beta1_RunAsUserStrategyOptions(in *policy.RunAsUserStrategyOptions, out *v1beta1.RunAsUserStrategyOptions, s conversion.Scope) error {
	out.Rule = v1beta1.RunAsUserStrategy(in.Rule)
	out.Ranges = *(*[]v1beta1.IDRange)(unsafe.Pointer(&in.Ranges))
	return nil
}

// Convert_policy_RunAsUserStrategyOptions_To_v1beta1_RunAsUserStrategyOptions is an autogenerated conversion function.
func Convert_policy_RunAsUserStrategyOptions_To_v1beta1_RunAsUserStrategyOptions(in *policy.RunAsUserStrategyOptions, out *v1beta1.RunAsUserStrategyOptions, s conversion.Scope) error {
	return autoConvert_policy_RunAsUserStrategyOptions_To_v1beta1_RunAsUserStrategyOptions(in, out, s)
}

func autoConvert_v1beta1_SELinuxStrategyOptions_To_policy_SELinuxStrategyOptions(in *v1beta1.SELinuxStrategyOptions, out *policy.SELinuxStrategyOptions, s conversion.Scope) error {
	out.Rule = policy.SELinuxStrategy(in.Rule)
	out.SELinuxOptions = (*core.SELinuxOptions)(unsafe.Pointer(in.SELinuxOptions))
	return nil
}

// Convert_v1beta1_SELinuxStrategyOptions_To_policy_SELinuxStrategyOptions is an autogenerated conversion function.
func Convert_v1beta1_SELinuxStrategyOptions_To_policy_SELinuxStrategyOptions(in *v1beta1.SELinuxStrategyOptions, out *policy.SELinuxStrategyOptions, s conversion.Scope) error {
	return autoConvert_v1beta1_SELinuxStrategyOptions_To_policy_SELinuxStrategyOptions(in, out, s)
}

func autoConvert_policy_SELinuxStrategyOptions_To_v1beta1_SELinuxStrategyOptions(in *policy.SELinuxStrategyOptions, out *v1beta1.SELinuxStrategyOptions, s conversion.Scope) error {
	out.Rule = v1beta1.SELinuxStrategy(in.Rule)
	out.SELinuxOptions = (*corev1.SELinuxOptions)(unsafe.Pointer(in.SELinuxOptions))
	return nil
}

// Convert_policy_SELinuxStrategyOptions_To_v1beta1_SELinuxStrategyOptions is an autogenerated conversion function.
func Convert_policy_SELinuxStrategyOptions_To_v1beta1_SELinuxStrategyOptions(in *policy.SELinuxStrategyOptions, out *v1beta1.SELinuxStrategyOptions, s conversion.Scope) error {
	return autoConvert_policy_SELinuxStrategyOptions_To_v1beta1_SELinuxStrategyOptions(in, out, s)
}

func autoConvert_v1beta1_SupplementalGroupsStrategyOptions_To_policy_SupplementalGroupsStrategyOptions(in *v1beta1.SupplementalGroupsStrategyOptions, out *policy.SupplementalGroupsStrategyOptions, s conversion.Scope) error {
	out.Rule = policy.SupplementalGroupsStrategyType(in.Rule)
	out.Ranges = *(*[]policy.IDRange)(unsafe.Pointer(&in.Ranges))
	return nil
}

// Convert_v1beta1_SupplementalGroupsStrategyOptions_To_policy_SupplementalGroupsStrategyOptions is an autogenerated conversion function.
func Convert_v1beta1_SupplementalGroupsStrategyOptions_To_policy_SupplementalGroupsStrategyOptions(in *v1beta1.SupplementalGroupsStrategyOptions, out *policy.SupplementalGroupsStrategyOptions, s conversion.Scope) error {
	return autoConvert_v1beta1_SupplementalGroupsStrategyOptions_To_policy_SupplementalGroupsStrategyOptions(in, out, s)
}

func autoConvert_policy_SupplementalGroupsStrategyOptions_To_v1beta1_SupplementalGroupsStrategyOptions(in *policy.SupplementalGroupsStrategyOptions, out *v1beta1.SupplementalGroupsStrategyOptions, s conversion.Scope) error {
	out.Rule = v1beta1.SupplementalGroupsStrategyType(in.Rule)
	out.Ranges = *(*[]v1beta1.IDRange)(unsafe.Pointer(&in.Ranges))
	return nil
}

// Convert_policy_SupplementalGroupsStrategyOptions_To_v1beta1_SupplementalGroupsStrategyOptions is an autogenerated conversion function.
func Convert_policy_SupplementalGroupsStrategyOptions_To_v1beta1_SupplementalGroupsStrategyOptions(in *policy.SupplementalGroupsStrategyOptions, out *v1beta1.SupplementalGroupsStrategyOptions, s conversion.Scope) error {
	return autoConvert_policy_SupplementalGroupsStrategyOptions_To_v1beta1_SupplementalGroupsStrategyOptions(in, out, s)
}
