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

package v1beta1

// This file contains a collection of methods that can be used from go-restful to
// generate Swagger API documentation for its models. Please read this PR for more
// information on the implementation: https://github.com/emicklei/go-restful/pull/215
//
// TODOs are ignored from the parser (e.g. TODO(andronat):... || TODO:...) if and only if
// they are on one line! For multiple line or blocks that you want to ignore use ---.
// Any context after a --- is ignored.
//
// Those methods can be generated by using hack/update-generated-swagger-docs.sh

// AUTO-GENERATED FUNCTIONS START HERE. DO NOT EDIT.
var map_AllowedCSIDriver = map[string]string{
	"":     "AllowedCSIDriver represents a single inline CSI Driver that is allowed to be used.",
	"name": "Name is the registered name of the CSI driver",
}

func (AllowedCSIDriver) SwaggerDoc() map[string]string {
	return map_AllowedCSIDriver
}

var map_AllowedFlexVolume = map[string]string{
	"":       "AllowedFlexVolume represents a single Flexvolume that is allowed to be used.",
	"driver": "driver is the name of the Flexvolume driver.",
}

func (AllowedFlexVolume) SwaggerDoc() map[string]string {
	return map_AllowedFlexVolume
}

var map_AllowedHostPath = map[string]string{
	"":           "AllowedHostPath defines the host volume conditions that will be enabled by a policy for pods to use. It requires the path prefix to be defined.",
	"pathPrefix": "pathPrefix is the path prefix that the host volume must match. It does not support `*`. Trailing slashes are trimmed when validating the path prefix with a host path.\n\nExamples: `/foo` would allow `/foo`, `/foo/` and `/foo/bar` `/foo` would not allow `/food` or `/etc/foo`",
	"readOnly":   "when set to true, will allow host volumes matching the pathPrefix only if all volume mounts are readOnly.",
}

func (AllowedHostPath) SwaggerDoc() map[string]string {
	return map_AllowedHostPath
}

var map_Eviction = map[string]string{
	"":              "Eviction evicts a pod from its node subject to certain policies and safety constraints. This is a subresource of Pod.  A request to cause such an eviction is created by POSTing to .../pods/<pod name>/evictions.",
	"metadata":      "ObjectMeta describes the pod that is being evicted.",
	"deleteOptions": "DeleteOptions may be provided",
}

func (Eviction) SwaggerDoc() map[string]string {
	return map_Eviction
}

var map_FSGroupStrategyOptions = map[string]string{
	"":       "FSGroupStrategyOptions defines the strategy type and options used to create the strategy.",
	"rule":   "rule is the strategy that will dictate what FSGroup is used in the SecurityContext.",
	"ranges": "ranges are the allowed ranges of fs groups.  If you would like to force a single fs group then supply a single range with the same start and end. Required for MustRunAs.",
}

func (FSGroupStrategyOptions) SwaggerDoc() map[string]string {
	return map_FSGroupStrategyOptions
}

var map_HostPortRange = map[string]string{
	"":    "HostPortRange defines a range of host ports that will be enabled by a policy for pods to use.  It requires both the start and end to be defined.",
	"min": "min is the start of the range, inclusive.",
	"max": "max is the end of the range, inclusive.",
}

func (HostPortRange) SwaggerDoc() map[string]string {
	return map_HostPortRange
}

var map_IDRange = map[string]string{
	"":    "IDRange provides a min/max of an allowed range of IDs.",
	"min": "min is the start of the range, inclusive.",
	"max": "max is the end of the range, inclusive.",
}

func (IDRange) SwaggerDoc() map[string]string {
	return map_IDRange
}

var map_PodDisruptionBudget = map[string]string{
	"":         "PodDisruptionBudget is an object to define the max disruption that can be caused to a collection of pods",
	"metadata": "Standard object's metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata",
	"spec":     "Specification of the desired behavior of the PodDisruptionBudget.",
	"status":   "Most recently observed status of the PodDisruptionBudget.",
}

func (PodDisruptionBudget) SwaggerDoc() map[string]string {
	return map_PodDisruptionBudget
}

var map_PodDisruptionBudgetList = map[string]string{
	"":         "PodDisruptionBudgetList is a collection of PodDisruptionBudgets.",
	"metadata": "Standard object's metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata",
	"items":    "items list individual PodDisruptionBudget objects",
}

func (PodDisruptionBudgetList) SwaggerDoc() map[string]string {
	return map_PodDisruptionBudgetList
}

var map_PodDisruptionBudgetSpec = map[string]string{
	"":               "PodDisruptionBudgetSpec is a description of a PodDisruptionBudget.",
	"minAvailable":   "An eviction is allowed if at least \"minAvailable\" pods selected by \"selector\" will still be available after the eviction, i.e. even in the absence of the evicted pod.  So for example you can prevent all voluntary evictions by specifying \"100%\".",
	"selector":       "Label query over pods whose evictions are managed by the disruption budget. A null selector selects no pods. An empty selector ({}) also selects no pods, which differs from standard behavior of selecting all pods. In policy/v1, an empty selector will select all pods in the namespace.",
	"maxUnavailable": "An eviction is allowed if at most \"maxUnavailable\" pods selected by \"selector\" are unavailable after the eviction, i.e. even in absence of the evicted pod. For example, one can prevent all voluntary evictions by specifying 0. This is a mutually exclusive setting with \"minAvailable\".",
}

func (PodDisruptionBudgetSpec) SwaggerDoc() map[string]string {
	return map_PodDisruptionBudgetSpec
}

var map_PodDisruptionBudgetStatus = map[string]string{
	"":                   "PodDisruptionBudgetStatus represents information about the status of a PodDisruptionBudget. Status may trail the actual state of a system.",
	"observedGeneration": "Most recent generation observed when updating this PDB status. DisruptionsAllowed and other status information is valid only if observedGeneration equals to PDB's object generation.",
	"disruptedPods":      "DisruptedPods contains information about pods whose eviction was processed by the API server eviction subresource handler but has not yet been observed by the PodDisruptionBudget controller. A pod will be in this map from the time when the API server processed the eviction request to the time when the pod is seen by PDB controller as having been marked for deletion (or after a timeout). The key in the map is the name of the pod and the value is the time when the API server processed the eviction request. If the deletion didn't occur and a pod is still there it will be removed from the list automatically by PodDisruptionBudget controller after some time. If everything goes smooth this map should be empty for the most of the time. Large number of entries in the map may indicate problems with pod deletions.",
	"disruptionsAllowed": "Number of pod disruptions that are currently allowed.",
	"currentHealthy":     "current number of healthy pods",
	"desiredHealthy":     "minimum desired number of healthy pods",
	"expectedPods":       "total number of pods counted by this disruption budget",
	"conditions":         "Conditions contain conditions for PDB. The disruption controller sets the DisruptionAllowed condition. The following are known values for the reason field (additional reasons could be added in the future): - SyncFailed: The controller encountered an error and wasn't able to compute\n              the number of allowed disruptions. Therefore no disruptions are\n              allowed and the status of the condition will be False.\n- InsufficientPods: The number of pods are either at or below the number\n                    required by the PodDisruptionBudget. No disruptions are\n                    allowed and the status of the condition will be False.\n- SufficientPods: There are more pods than required by the PodDisruptionBudget.\n                  The condition will be True, and the number of allowed\n                  disruptions are provided by the disruptionsAllowed property.",
}

func (PodDisruptionBudgetStatus) SwaggerDoc() map[string]string {
	return map_PodDisruptionBudgetStatus
}

var map_PodSecurityPolicy = map[string]string{
	"":         "PodSecurityPolicy governs the ability to make requests that affect the Security Context that will be applied to a pod and container. Deprecated in 1.21.",
	"metadata": "Standard object's metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata",
	"spec":     "spec defines the policy enforced.",
}

func (PodSecurityPolicy) SwaggerDoc() map[string]string {
	return map_PodSecurityPolicy
}

var map_PodSecurityPolicyList = map[string]string{
	"":         "PodSecurityPolicyList is a list of PodSecurityPolicy objects.",
	"metadata": "Standard list metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata",
	"items":    "items is a list of schema objects.",
}

func (PodSecurityPolicyList) SwaggerDoc() map[string]string {
	return map_PodSecurityPolicyList
}

var map_PodSecurityPolicySpec = map[string]string{
	"":                                "PodSecurityPolicySpec defines the policy enforced.",
	"privileged":                      "privileged determines if a pod can request to be run as privileged.",
	"defaultAddCapabilities":          "defaultAddCapabilities is the default set of capabilities that will be added to the container unless the pod spec specifically drops the capability.  You may not list a capability in both defaultAddCapabilities and requiredDropCapabilities. Capabilities added here are implicitly allowed, and need not be included in the allowedCapabilities list.",
	"requiredDropCapabilities":        "requiredDropCapabilities are the capabilities that will be dropped from the container.  These are required to be dropped and cannot be added.",
	"allowedCapabilities":             "allowedCapabilities is a list of capabilities that can be requested to add to the container. Capabilities in this field may be added at the pod author's discretion. You must not list a capability in both allowedCapabilities and requiredDropCapabilities.",
	"volumes":                         "volumes is an allowlist of volume plugins. Empty indicates that no volumes may be used. To allow all volumes you may use '*'.",
	"hostNetwork":                     "hostNetwork determines if the policy allows the use of HostNetwork in the pod spec.",
	"hostPorts":                       "hostPorts determines which host port ranges are allowed to be exposed.",
	"hostPID":                         "hostPID determines if the policy allows the use of HostPID in the pod spec.",
	"hostIPC":                         "hostIPC determines if the policy allows the use of HostIPC in the pod spec.",
	"seLinux":                         "seLinux is the strategy that will dictate the allowable labels that may be set.",
	"runAsUser":                       "runAsUser is the strategy that will dictate the allowable RunAsUser values that may be set.",
	"runAsGroup":                      "RunAsGroup is the strategy that will dictate the allowable RunAsGroup values that may be set. If this field is omitted, the pod's RunAsGroup can take any value. This field requires the RunAsGroup feature gate to be enabled.",
	"supplementalGroups":              "supplementalGroups is the strategy that will dictate what supplemental groups are used by the SecurityContext.",
	"fsGroup":                         "fsGroup is the strategy that will dictate what fs group is used by the SecurityContext.",
	"readOnlyRootFilesystem":          "readOnlyRootFilesystem when set to true will force containers to run with a read only root file system.  If the container specifically requests to run with a non-read only root file system the PSP should deny the pod. If set to false the container may run with a read only root file system if it wishes but it will not be forced to.",
	"defaultAllowPrivilegeEscalation": "defaultAllowPrivilegeEscalation controls the default setting for whether a process can gain more privileges than its parent process.",
	"allowPrivilegeEscalation":        "allowPrivilegeEscalation determines if a pod can request to allow privilege escalation. If unspecified, defaults to true.",
	"allowedHostPaths":                "allowedHostPaths is an allowlist of host paths. Empty indicates that all host paths may be used.",
	"allowedFlexVolumes":              "allowedFlexVolumes is an allowlist of Flexvolumes.  Empty or nil indicates that all Flexvolumes may be used.  This parameter is effective only when the usage of the Flexvolumes is allowed in the \"volumes\" field.",
	"allowedCSIDrivers":               "AllowedCSIDrivers is an allowlist of inline CSI drivers that must be explicitly set to be embedded within a pod spec. An empty value indicates that any CSI driver can be used for inline ephemeral volumes.",
	"allowedUnsafeSysctls":            "allowedUnsafeSysctls is a list of explicitly allowed unsafe sysctls, defaults to none. Each entry is either a plain sysctl name or ends in \"*\" in which case it is considered as a prefix of allowed sysctls. Single * means all unsafe sysctls are allowed. Kubelet has to allowlist all allowed unsafe sysctls explicitly to avoid rejection.\n\nExamples: e.g. \"foo/*\" allows \"foo/bar\", \"foo/baz\", etc. e.g. \"foo.*\" allows \"foo.bar\", \"foo.baz\", etc.",
	"forbiddenSysctls":                "forbiddenSysctls is a list of explicitly forbidden sysctls, defaults to none. Each entry is either a plain sysctl name or ends in \"*\" in which case it is considered as a prefix of forbidden sysctls. Single * means all sysctls are forbidden.\n\nExamples: e.g. \"foo/*\" forbids \"foo/bar\", \"foo/baz\", etc. e.g. \"foo.*\" forbids \"foo.bar\", \"foo.baz\", etc.",
	"allowedProcMountTypes":           "AllowedProcMountTypes is an allowlist of allowed ProcMountTypes. Empty or nil indicates that only the DefaultProcMountType may be used. This requires the ProcMountType feature flag to be enabled.",
	"runtimeClass":                    "runtimeClass is the strategy that will dictate the allowable RuntimeClasses for a pod. If this field is omitted, the pod's runtimeClassName field is unrestricted. Enforcement of this field depends on the RuntimeClass feature gate being enabled.",
}

func (PodSecurityPolicySpec) SwaggerDoc() map[string]string {
	return map_PodSecurityPolicySpec
}

var map_RunAsGroupStrategyOptions = map[string]string{
	"":       "RunAsGroupStrategyOptions defines the strategy type and any options used to create the strategy.",
	"rule":   "rule is the strategy that will dictate the allowable RunAsGroup values that may be set.",
	"ranges": "ranges are the allowed ranges of gids that may be used. If you would like to force a single gid then supply a single range with the same start and end. Required for MustRunAs.",
}

func (RunAsGroupStrategyOptions) SwaggerDoc() map[string]string {
	return map_RunAsGroupStrategyOptions
}

var map_RunAsUserStrategyOptions = map[string]string{
	"":       "RunAsUserStrategyOptions defines the strategy type and any options used to create the strategy.",
	"rule":   "rule is the strategy that will dictate the allowable RunAsUser values that may be set.",
	"ranges": "ranges are the allowed ranges of uids that may be used. If you would like to force a single uid then supply a single range with the same start and end. Required for MustRunAs.",
}

func (RunAsUserStrategyOptions) SwaggerDoc() map[string]string {
	return map_RunAsUserStrategyOptions
}

var map_RuntimeClassStrategyOptions = map[string]string{
	"":                         "RuntimeClassStrategyOptions define the strategy that will dictate the allowable RuntimeClasses for a pod.",
	"allowedRuntimeClassNames": "allowedRuntimeClassNames is an allowlist of RuntimeClass names that may be specified on a pod. A value of \"*\" means that any RuntimeClass name is allowed, and must be the only item in the list. An empty list requires the RuntimeClassName field to be unset.",
	"defaultRuntimeClassName":  "defaultRuntimeClassName is the default RuntimeClassName to set on the pod. The default MUST be allowed by the allowedRuntimeClassNames list. A value of nil does not mutate the Pod.",
}

func (RuntimeClassStrategyOptions) SwaggerDoc() map[string]string {
	return map_RuntimeClassStrategyOptions
}

var map_SELinuxStrategyOptions = map[string]string{
	"":               "SELinuxStrategyOptions defines the strategy type and any options used to create the strategy.",
	"rule":           "rule is the strategy that will dictate the allowable labels that may be set.",
	"seLinuxOptions": "seLinuxOptions required to run as; required for MustRunAs More info: https://kubernetes.io/docs/tasks/configure-pod-container/security-context/",
}

func (SELinuxStrategyOptions) SwaggerDoc() map[string]string {
	return map_SELinuxStrategyOptions
}

var map_SupplementalGroupsStrategyOptions = map[string]string{
	"":       "SupplementalGroupsStrategyOptions defines the strategy type and options used to create the strategy.",
	"rule":   "rule is the strategy that will dictate what supplemental groups is used in the SecurityContext.",
	"ranges": "ranges are the allowed ranges of supplemental groups.  If you would like to force a single supplemental group then supply a single range with the same start and end. Required for MustRunAs.",
}

func (SupplementalGroupsStrategyOptions) SwaggerDoc() map[string]string {
	return map_SupplementalGroupsStrategyOptions
}

// AUTO-GENERATED FUNCTIONS END HERE
