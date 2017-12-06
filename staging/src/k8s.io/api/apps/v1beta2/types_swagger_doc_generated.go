/*
Copyright 2016 The Kubernetes Authors.

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

package v1beta2

// This file contains a collection of methods that can be used from go-restful to
// generate Swagger API documentation for its models. Please read this PR for more
// information on the implementation: https://github.com/emicklei/go-restful/pull/215
//
// TODOs are ignored from the parser (e.g. TODO(andronat):... || TODO:...) if and only if
// they are on one line! For multiple line or blocks that you want to ignore use ---.
// Any context after a --- is ignored.
//
// Those methods can be generated by using hack/update-generated-swagger-docs.sh

// AUTO-GENERATED FUNCTIONS START HERE
var map_ControllerRevision = map[string]string{
	"":         "DEPRECATED - This group version of ControllerRevision is deprecated by apps/v1/ControllerRevision. See the release notes for more information. ControllerRevision implements an immutable snapshot of state data. Clients are responsible for serializing and deserializing the objects that contain their internal state. Once a ControllerRevision has been successfully created, it can not be updated. The API Server will fail validation of all requests that attempt to mutate the Data field. ControllerRevisions may, however, be deleted. Note that, due to its use by both the DaemonSet and StatefulSet controllers for update and rollback, this object is beta. However, it may be subject to name and representation changes in future releases, and clients should not depend on its stability. It is primarily for internal use by controllers.",
	"metadata": "Standard object's metadata. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#metadata",
	"data":     "Data is the serialized representation of the state.",
	"revision": "Revision indicates the revision of the state represented by Data.",
}

func (ControllerRevision) SwaggerDoc() map[string]string {
	return map_ControllerRevision
}

var map_ControllerRevisionList = map[string]string{
	"":         "ControllerRevisionList is a resource containing a list of ControllerRevision objects.",
	"metadata": "More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#metadata",
	"items":    "Items is the list of ControllerRevisions",
}

func (ControllerRevisionList) SwaggerDoc() map[string]string {
	return map_ControllerRevisionList
}

var map_DaemonSet = map[string]string{
	"":         "DEPRECATED - This group version of DaemonSet is deprecated by apps/v1/DaemonSet. See the release notes for more information. DaemonSet represents the configuration of a daemon set.",
	"metadata": "Standard object's metadata. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#metadata",
	"spec":     "The desired behavior of this daemon set. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#spec-and-status",
	"status":   "The current status of this daemon set. This data may be out of date by some window of time. Populated by the system. Read-only. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#spec-and-status",
}

func (DaemonSet) SwaggerDoc() map[string]string {
	return map_DaemonSet
}

var map_DaemonSetCondition = map[string]string{
	"":                   "DaemonSetCondition describes the state of a DaemonSet at a certain point.",
	"type":               "Type of DaemonSet condition.",
	"status":             "Status of the condition, one of True, False, Unknown.",
	"lastTransitionTime": "Last time the condition transitioned from one status to another.",
	"reason":             "The reason for the condition's last transition.",
	"message":            "A human readable message indicating details about the transition.",
}

func (DaemonSetCondition) SwaggerDoc() map[string]string {
	return map_DaemonSetCondition
}

var map_DaemonSetList = map[string]string{
	"":         "DaemonSetList is a collection of daemon sets.",
	"metadata": "Standard list metadata. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#metadata",
	"items":    "A list of daemon sets.",
}

func (DaemonSetList) SwaggerDoc() map[string]string {
	return map_DaemonSetList
}

var map_DaemonSetSpec = map[string]string{
	"":                     "DaemonSetSpec is the specification of a daemon set.",
	"selector":             "A label query over pods that are managed by the daemon set. Must match in order to be controlled. It must match the pod template's labels. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/labels/#label-selectors",
	"template":             "An object that describes the pod that will be created. The DaemonSet will create exactly one copy of this pod on every node that matches the template's node selector (or on every node if no node selector is specified). More info: https://kubernetes.io/docs/concepts/workloads/controllers/replicationcontroller#pod-template",
	"updateStrategy":       "An update strategy to replace existing DaemonSet pods with new pods.",
	"minReadySeconds":      "The minimum number of seconds for which a newly created DaemonSet pod should be ready without any of its container crashing, for it to be considered available. Defaults to 0 (pod will be considered available as soon as it is ready).",
	"revisionHistoryLimit": "The number of old history to retain to allow rollback. This is a pointer to distinguish between explicit zero and not specified. Defaults to 10.",
}

func (DaemonSetSpec) SwaggerDoc() map[string]string {
	return map_DaemonSetSpec
}

var map_DaemonSetStatus = map[string]string{
	"": "DaemonSetStatus represents the current status of a daemon set.",
	"currentNumberScheduled": "The number of nodes that are running at least 1 daemon pod and are supposed to run the daemon pod. More info: https://kubernetes.io/docs/concepts/workloads/controllers/daemonset/",
	"numberMisscheduled":     "The number of nodes that are running the daemon pod, but are not supposed to run the daemon pod. More info: https://kubernetes.io/docs/concepts/workloads/controllers/daemonset/",
	"desiredNumberScheduled": "The total number of nodes that should be running the daemon pod (including nodes correctly running the daemon pod). More info: https://kubernetes.io/docs/concepts/workloads/controllers/daemonset/",
	"numberReady":            "The number of nodes that should be running the daemon pod and have one or more of the daemon pod running and ready.",
	"observedGeneration":     "The most recent generation observed by the daemon set controller.",
	"updatedNumberScheduled": "The total number of nodes that are running updated daemon pod",
	"numberAvailable":        "The number of nodes that should be running the daemon pod and have one or more of the daemon pod running and available (ready for at least spec.minReadySeconds)",
	"numberUnavailable":      "The number of nodes that should be running the daemon pod and have none of the daemon pod running and available (ready for at least spec.minReadySeconds)",
	"collisionCount":         "Count of hash collisions for the DaemonSet. The DaemonSet controller uses this field as a collision avoidance mechanism when it needs to create the name for the newest ControllerRevision.",
	"conditions":             "Represents the latest available observations of a DaemonSet's current state.",
}

func (DaemonSetStatus) SwaggerDoc() map[string]string {
	return map_DaemonSetStatus
}

var map_DaemonSetUpdateStrategy = map[string]string{
	"":              "DaemonSetUpdateStrategy is a struct used to control the update strategy for a DaemonSet.",
	"type":          "Type of daemon set update. Can be \"RollingUpdate\" or \"OnDelete\". Default is RollingUpdate.",
	"rollingUpdate": "Rolling update config params. Present only if type = \"RollingUpdate\".",
}

func (DaemonSetUpdateStrategy) SwaggerDoc() map[string]string {
	return map_DaemonSetUpdateStrategy
}

var map_Deployment = map[string]string{
	"":         "DEPRECATED - This group version of Deployment is deprecated by apps/v1/Deployment. See the release notes for more information. Deployment enables declarative updates for Pods and ReplicaSets.",
	"metadata": "Standard object metadata.",
	"spec":     "Specification of the desired behavior of the Deployment.",
	"status":   "Most recently observed status of the Deployment.",
}

func (Deployment) SwaggerDoc() map[string]string {
	return map_Deployment
}

var map_DeploymentCondition = map[string]string{
	"":                   "DeploymentCondition describes the state of a deployment at a certain point.",
	"type":               "Type of deployment condition.",
	"status":             "Status of the condition, one of True, False, Unknown.",
	"lastUpdateTime":     "The last time this condition was updated.",
	"lastTransitionTime": "Last time the condition transitioned from one status to another.",
	"reason":             "The reason for the condition's last transition.",
	"message":            "A human readable message indicating details about the transition.",
}

func (DeploymentCondition) SwaggerDoc() map[string]string {
	return map_DeploymentCondition
}

var map_DeploymentList = map[string]string{
	"":         "DeploymentList is a list of Deployments.",
	"metadata": "Standard list metadata.",
	"items":    "Items is the list of Deployments.",
}

func (DeploymentList) SwaggerDoc() map[string]string {
	return map_DeploymentList
}

var map_DeploymentSpec = map[string]string{
	"":                        "DeploymentSpec is the specification of the desired behavior of the Deployment.",
	"replicas":                "Number of desired pods. This is a pointer to distinguish between explicit zero and not specified. Defaults to 1.",
	"selector":                "Label selector for pods. Existing ReplicaSets whose pods are selected by this will be the ones affected by this deployment. It must match the pod template's labels.",
	"template":                "Template describes the pods that will be created.",
	"strategy":                "The deployment strategy to use to replace existing pods with new ones.",
	"minReadySeconds":         "Minimum number of seconds for which a newly created pod should be ready without any of its container crashing, for it to be considered available. Defaults to 0 (pod will be considered available as soon as it is ready)",
	"revisionHistoryLimit":    "The number of old ReplicaSets to retain to allow rollback. This is a pointer to distinguish between explicit zero and not specified. Defaults to 10.",
	"paused":                  "Indicates that the deployment is paused.",
	"progressDeadlineSeconds": "The maximum time in seconds for a deployment to make progress before it is considered to be failed. The deployment controller will continue to process failed deployments and a condition with a ProgressDeadlineExceeded reason will be surfaced in the deployment status. Note that progress will not be estimated during the time a deployment is paused. Defaults to 600s.",
}

func (DeploymentSpec) SwaggerDoc() map[string]string {
	return map_DeploymentSpec
}

var map_DeploymentStatus = map[string]string{
	"":                    "DeploymentStatus is the most recently observed status of the Deployment.",
	"observedGeneration":  "The generation observed by the deployment controller.",
	"replicas":            "Total number of non-terminated pods targeted by this deployment (their labels match the selector).",
	"updatedReplicas":     "Total number of non-terminated pods targeted by this deployment that have the desired template spec.",
	"readyReplicas":       "Total number of ready pods targeted by this deployment.",
	"availableReplicas":   "Total number of available pods (ready for at least minReadySeconds) targeted by this deployment.",
	"unavailableReplicas": "Total number of unavailable pods targeted by this deployment. This is the total number of pods that are still required for the deployment to have 100% available capacity. They may either be pods that are running but not yet available or pods that still have not been created.",
	"conditions":          "Represents the latest available observations of a deployment's current state.",
	"collisionCount":      "Count of hash collisions for the Deployment. The Deployment controller uses this field as a collision avoidance mechanism when it needs to create the name for the newest ReplicaSet.",
}

func (DeploymentStatus) SwaggerDoc() map[string]string {
	return map_DeploymentStatus
}

var map_DeploymentStrategy = map[string]string{
	"":              "DeploymentStrategy describes how to replace existing pods with new ones.",
	"type":          "Type of deployment. Can be \"Recreate\" or \"RollingUpdate\". Default is RollingUpdate.",
	"rollingUpdate": "Rolling update config params. Present only if DeploymentStrategyType = RollingUpdate.",
}

func (DeploymentStrategy) SwaggerDoc() map[string]string {
	return map_DeploymentStrategy
}

var map_ReplicaSet = map[string]string{
	"":         "DEPRECATED - This group version of ReplicaSet is deprecated by apps/v1/ReplicaSet. See the release notes for more information. ReplicaSet ensures that a specified number of pod replicas are running at any given time.",
	"metadata": "If the Labels of a ReplicaSet are empty, they are defaulted to be the same as the Pod(s) that the ReplicaSet manages. Standard object's metadata. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#metadata",
	"spec":     "Spec defines the specification of the desired behavior of the ReplicaSet. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#spec-and-status",
	"status":   "Status is the most recently observed status of the ReplicaSet. This data may be out of date by some window of time. Populated by the system. Read-only. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#spec-and-status",
}

func (ReplicaSet) SwaggerDoc() map[string]string {
	return map_ReplicaSet
}

var map_ReplicaSetCondition = map[string]string{
	"":                   "ReplicaSetCondition describes the state of a replica set at a certain point.",
	"type":               "Type of replica set condition.",
	"status":             "Status of the condition, one of True, False, Unknown.",
	"lastTransitionTime": "The last time the condition transitioned from one status to another.",
	"reason":             "The reason for the condition's last transition.",
	"message":            "A human readable message indicating details about the transition.",
}

func (ReplicaSetCondition) SwaggerDoc() map[string]string {
	return map_ReplicaSetCondition
}

var map_ReplicaSetList = map[string]string{
	"":         "ReplicaSetList is a collection of ReplicaSets.",
	"metadata": "Standard list metadata. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds",
	"items":    "List of ReplicaSets. More info: https://kubernetes.io/docs/concepts/workloads/controllers/replicationcontroller",
}

func (ReplicaSetList) SwaggerDoc() map[string]string {
	return map_ReplicaSetList
}

var map_ReplicaSetSpec = map[string]string{
	"":                "ReplicaSetSpec is the specification of a ReplicaSet.",
	"replicas":        "Replicas is the number of desired replicas. This is a pointer to distinguish between explicit zero and unspecified. Defaults to 1. More info: https://kubernetes.io/docs/concepts/workloads/controllers/replicationcontroller/#what-is-a-replicationcontroller",
	"minReadySeconds": "Minimum number of seconds for which a newly created pod should be ready without any of its container crashing, for it to be considered available. Defaults to 0 (pod will be considered available as soon as it is ready)",
	"selector":        "Selector is a label query over pods that should match the replica count. Label keys and values that must match in order to be controlled by this replica set. It must match the pod template's labels. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/labels/#label-selectors",
	"template":        "Template is the object that describes the pod that will be created if insufficient replicas are detected. More info: https://kubernetes.io/docs/concepts/workloads/controllers/replicationcontroller#pod-template",
}

func (ReplicaSetSpec) SwaggerDoc() map[string]string {
	return map_ReplicaSetSpec
}

var map_ReplicaSetStatus = map[string]string{
	"":                     "ReplicaSetStatus represents the current status of a ReplicaSet.",
	"replicas":             "Replicas is the most recently oberved number of replicas. More info: https://kubernetes.io/docs/concepts/workloads/controllers/replicationcontroller/#what-is-a-replicationcontroller",
	"fullyLabeledReplicas": "The number of pods that have labels matching the labels of the pod template of the replicaset.",
	"readyReplicas":        "The number of ready replicas for this replica set.",
	"availableReplicas":    "The number of available replicas (ready for at least minReadySeconds) for this replica set.",
	"observedGeneration":   "ObservedGeneration reflects the generation of the most recently observed ReplicaSet.",
	"conditions":           "Represents the latest available observations of a replica set's current state.",
}

func (ReplicaSetStatus) SwaggerDoc() map[string]string {
	return map_ReplicaSetStatus
}

var map_RollingUpdateDaemonSet = map[string]string{
	"":               "Spec to control the desired behavior of daemon set rolling update.",
	"maxUnavailable": "The maximum number of DaemonSet pods that can be unavailable during the update. Value can be an absolute number (ex: 5) or a percentage of total number of DaemonSet pods at the start of the update (ex: 10%). Absolute number is calculated from percentage by rounding up. This cannot be 0. Default value is 1. Example: when this is set to 30%, at most 30% of the total number of nodes that should be running the daemon pod (i.e. status.desiredNumberScheduled) can have their pods stopped for an update at any given time. The update starts by stopping at most 30% of those DaemonSet pods and then brings up new DaemonSet pods in their place. Once the new pods are available, it then proceeds onto other DaemonSet pods, thus ensuring that at least 70% of original number of DaemonSet pods are available at all times during the update.",
}

func (RollingUpdateDaemonSet) SwaggerDoc() map[string]string {
	return map_RollingUpdateDaemonSet
}

var map_RollingUpdateDeployment = map[string]string{
	"":               "Spec to control the desired behavior of rolling update.",
	"maxUnavailable": "The maximum number of pods that can be unavailable during the update. Value can be an absolute number (ex: 5) or a percentage of desired pods (ex: 10%). Absolute number is calculated from percentage by rounding down. This can not be 0 if MaxSurge is 0. Defaults to 25%. Example: when this is set to 30%, the old RC can be scaled down to 70% of desired pods immediately when the rolling update starts. Once new pods are ready, old RC can be scaled down further, followed by scaling up the new RC, ensuring that the total number of pods available at all times during the update is at least 70% of desired pods.",
	"maxSurge":       "The maximum number of pods that can be scheduled above the desired number of pods. Value can be an absolute number (ex: 5) or a percentage of desired pods (ex: 10%). This can not be 0 if MaxUnavailable is 0. Absolute number is calculated from percentage by rounding up. Defaults to 25%. Example: when this is set to 30%, the new RC can be scaled up immediately when the rolling update starts, such that the total number of old and new pods do not exceed 130% of desired pods. Once old pods have been killed, new RC can be scaled up further, ensuring that total number of pods running at any time during the update is atmost 130% of desired pods.",
}

func (RollingUpdateDeployment) SwaggerDoc() map[string]string {
	return map_RollingUpdateDeployment
}

var map_RollingUpdateStatefulSetStrategy = map[string]string{
	"":          "RollingUpdateStatefulSetStrategy is used to communicate parameter for RollingUpdateStatefulSetStrategyType.",
	"partition": "Partition indicates the ordinal at which the StatefulSet should be partitioned. Default value is 0.",
}

func (RollingUpdateStatefulSetStrategy) SwaggerDoc() map[string]string {
	return map_RollingUpdateStatefulSetStrategy
}

var map_Scale = map[string]string{
	"":         "Scale represents a scaling request for a resource.",
	"metadata": "Standard object metadata; More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#metadata.",
	"spec":     "defines the behavior of the scale. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#spec-and-status.",
	"status":   "current status of the scale. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#spec-and-status. Read-only.",
}

func (Scale) SwaggerDoc() map[string]string {
	return map_Scale
}

var map_ScaleSpec = map[string]string{
	"":         "ScaleSpec describes the attributes of a scale subresource",
	"replicas": "desired number of instances for the scaled object.",
}

func (ScaleSpec) SwaggerDoc() map[string]string {
	return map_ScaleSpec
}

var map_ScaleStatus = map[string]string{
	"":               "ScaleStatus represents the current status of a scale subresource.",
	"replicas":       "actual number of observed instances of the scaled object.",
	"selector":       "label query over pods that should match the replicas count. More info: http://kubernetes.io/docs/user-guide/labels#label-selectors",
	"targetSelector": "label selector for pods that should match the replicas count. This is a serializated version of both map-based and more expressive set-based selectors. This is done to avoid introspection in the clients. The string will be in the same format as the query-param syntax. If the target type only supports map-based selectors, both this field and map-based selector field are populated. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/labels/#label-selectors",
}

func (ScaleStatus) SwaggerDoc() map[string]string {
	return map_ScaleStatus
}

var map_StatefulSet = map[string]string{
	"":       "DEPRECATED - This group version of StatefulSet is deprecated by apps/v1/StatefulSet. See the release notes for more information. StatefulSet represents a set of pods with consistent identities. Identities are defined as:\n - Network: A single stable DNS and hostname.\n - Storage: As many VolumeClaims as requested.\nThe StatefulSet guarantees that a given network identity will always map to the same storage identity.",
	"spec":   "Spec defines the desired identities of pods in this set.",
	"status": "Status is the current status of Pods in this StatefulSet. This data may be out of date by some window of time.",
}

func (StatefulSet) SwaggerDoc() map[string]string {
	return map_StatefulSet
}

var map_StatefulSetCondition = map[string]string{
	"":                   "StatefulSetCondition describes the state of a statefulset at a certain point.",
	"type":               "Type of statefulset condition.",
	"status":             "Status of the condition, one of True, False, Unknown.",
	"lastTransitionTime": "Last time the condition transitioned from one status to another.",
	"reason":             "The reason for the condition's last transition.",
	"message":            "A human readable message indicating details about the transition.",
}

func (StatefulSetCondition) SwaggerDoc() map[string]string {
	return map_StatefulSetCondition
}

var map_StatefulSetList = map[string]string{
	"": "StatefulSetList is a collection of StatefulSets.",
}

func (StatefulSetList) SwaggerDoc() map[string]string {
	return map_StatefulSetList
}

var map_StatefulSetSpec = map[string]string{
	"":                     "A StatefulSetSpec is the specification of a StatefulSet.",
	"replicas":             "replicas is the desired number of replicas of the given Template. These are replicas in the sense that they are instantiations of the same Template, but individual replicas also have a consistent identity. If unspecified, defaults to 1.",
	"selector":             "selector is a label query over pods that should match the replica count. It must match the pod template's labels. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/labels/#label-selectors",
	"template":             "template is the object that describes the pod that will be created if insufficient replicas are detected. Each pod stamped out by the StatefulSet will fulfill this Template, but have a unique identity from the rest of the StatefulSet.",
	"volumeClaimTemplates": "volumeClaimTemplates is a list of claims that pods are allowed to reference. The StatefulSet controller is responsible for mapping network identities to claims in a way that maintains the identity of a pod. Every claim in this list must have at least one matching (by name) volumeMount in one container in the template. A claim in this list takes precedence over any volumes in the template, with the same name.",
	"serviceName":          "serviceName is the name of the service that governs this StatefulSet. This service must exist before the StatefulSet, and is responsible for the network identity of the set. Pods get DNS/hostnames that follow the pattern: pod-specific-string.serviceName.default.svc.cluster.local where \"pod-specific-string\" is managed by the StatefulSet controller.",
	"podManagementPolicy":  "podManagementPolicy controls how pods are created during initial scale up, when replacing pods on nodes, or when scaling down. The default policy is `OrderedReady`, where pods are created in increasing order (pod-0, then pod-1, etc) and the controller will wait until each pod is ready before continuing. When scaling down, the pods are removed in the opposite order. The alternative policy is `Parallel` which will create pods in parallel to match the desired scale without waiting, and on scale down will delete all pods at once.",
	"updateStrategy":       "updateStrategy indicates the StatefulSetUpdateStrategy that will be employed to update Pods in the StatefulSet when a revision is made to Template.",
	"revisionHistoryLimit": "revisionHistoryLimit is the maximum number of revisions that will be maintained in the StatefulSet's revision history. The revision history consists of all revisions not represented by a currently applied StatefulSetSpec version. The default value is 10.",
}

func (StatefulSetSpec) SwaggerDoc() map[string]string {
	return map_StatefulSetSpec
}

var map_StatefulSetStatus = map[string]string{
	"":                   "StatefulSetStatus represents the current state of a StatefulSet.",
	"observedGeneration": "observedGeneration is the most recent generation observed for this StatefulSet. It corresponds to the StatefulSet's generation, which is updated on mutation by the API Server.",
	"replicas":           "replicas is the number of Pods created by the StatefulSet controller.",
	"readyReplicas":      "readyReplicas is the number of Pods created by the StatefulSet controller that have a Ready Condition.",
	"currentReplicas":    "currentReplicas is the number of Pods created by the StatefulSet controller from the StatefulSet version indicated by currentRevision.",
	"updatedReplicas":    "updatedReplicas is the number of Pods created by the StatefulSet controller from the StatefulSet version indicated by updateRevision.",
	"currentRevision":    "currentRevision, if not empty, indicates the version of the StatefulSet used to generate Pods in the sequence [0,currentReplicas).",
	"updateRevision":     "updateRevision, if not empty, indicates the version of the StatefulSet used to generate Pods in the sequence [replicas-updatedReplicas,replicas)",
	"collisionCount":     "collisionCount is the count of hash collisions for the StatefulSet. The StatefulSet controller uses this field as a collision avoidance mechanism when it needs to create the name for the newest ControllerRevision.",
	"conditions":         "Represents the latest available observations of a statefulset's current state.",
}

func (StatefulSetStatus) SwaggerDoc() map[string]string {
	return map_StatefulSetStatus
}

var map_StatefulSetUpdateStrategy = map[string]string{
	"":              "StatefulSetUpdateStrategy indicates the strategy that the StatefulSet controller will use to perform updates. It includes any additional parameters necessary to perform the update for the indicated strategy.",
	"type":          "Type indicates the type of the StatefulSetUpdateStrategy. Default is RollingUpdate.",
	"rollingUpdate": "RollingUpdate is used to communicate parameters when Type is RollingUpdateStatefulSetStrategyType.",
}

func (StatefulSetUpdateStrategy) SwaggerDoc() map[string]string {
	return map_StatefulSetUpdateStrategy
}

// AUTO-GENERATED FUNCTIONS END HERE
