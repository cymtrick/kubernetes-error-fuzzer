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

package v1

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
var map_CSIDriver = map[string]string{
	"":         "CSIDriver captures information about a Container Storage Interface (CSI) volume driver deployed on the cluster. Kubernetes attach detach controller uses this object to determine whether attach is required. Kubelet uses this object to determine whether pod information needs to be passed on mount. CSIDriver objects are non-namespaced.",
	"metadata": "Standard object metadata. metadata.Name indicates the name of the CSI driver that this object refers to; it MUST be the same name returned by the CSI GetPluginName() call for that driver. The driver name must be 63 characters or less, beginning and ending with an alphanumeric character ([a-z0-9A-Z]) with dashes (-), dots (.), and alphanumerics between. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata",
	"spec":     "Specification of the CSI Driver.",
}

func (CSIDriver) SwaggerDoc() map[string]string {
	return map_CSIDriver
}

var map_CSIDriverList = map[string]string{
	"":         "CSIDriverList is a collection of CSIDriver objects.",
	"metadata": "Standard list metadata More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata",
	"items":    "items is the list of CSIDriver",
}

func (CSIDriverList) SwaggerDoc() map[string]string {
	return map_CSIDriverList
}

var map_CSIDriverSpec = map[string]string{
	"":                     "CSIDriverSpec is the specification of a CSIDriver.",
	"attachRequired":       "attachRequired indicates this CSI volume driver requires an attach operation (because it implements the CSI ControllerPublishVolume() method), and that the Kubernetes attach detach controller should call the attach volume interface which checks the volumeattachment status and waits until the volume is attached before proceeding to mounting. The CSI external-attacher coordinates with CSI volume driver and updates the volumeattachment status when the attach operation is complete. If the CSIDriverRegistry feature gate is enabled and the value is specified to false, the attach operation will be skipped. Otherwise the attach operation will be called.\n\nThis field is immutable.",
	"podInfoOnMount":       "If set to true, podInfoOnMount indicates this CSI volume driver requires additional pod information (like podName, podUID, etc.) during mount operations. If set to false, pod information will not be passed on mount. Default is false. The CSI driver specifies podInfoOnMount as part of driver deployment. If true, Kubelet will pass pod information as VolumeContext in the CSI NodePublishVolume() calls. The CSI driver is responsible for parsing and validating the information passed in as VolumeContext. The following VolumeConext will be passed if podInfoOnMount is set to true. This list might grow, but the prefix will be used. \"csi.storage.k8s.io/pod.name\": pod.Name \"csi.storage.k8s.io/pod.namespace\": pod.Namespace \"csi.storage.k8s.io/pod.uid\": string(pod.UID) \"csi.storage.k8s.io/ephemeral\": \"true\" if the volume is an ephemeral inline volume\n                                defined by a CSIVolumeSource, otherwise \"false\"\n\n\"csi.storage.k8s.io/ephemeral\" is a new feature in Kubernetes 1.16. It is only required for drivers which support both the \"Persistent\" and \"Ephemeral\" VolumeLifecycleMode. Other drivers can leave pod info disabled and/or ignore this field. As Kubernetes 1.15 doesn't support this field, drivers can only support one mode when deployed on such a cluster and the deployment determines which mode that is, for example via a command line parameter of the driver.\n\nThis field is immutable.",
	"volumeLifecycleModes": "volumeLifecycleModes defines what kind of volumes this CSI volume driver supports. The default if the list is empty is \"Persistent\", which is the usage defined by the CSI specification and implemented in Kubernetes via the usual PV/PVC mechanism. The other mode is \"Ephemeral\". In this mode, volumes are defined inline inside the pod spec with CSIVolumeSource and their lifecycle is tied to the lifecycle of that pod. A driver has to be aware of this because it is only going to get a NodePublishVolume call for such a volume. For more information about implementing this mode, see https://kubernetes-csi.github.io/docs/ephemeral-local-volumes.html A driver can support one or more of these modes and more modes may be added in the future. This field is beta.\n\nThis field is immutable.",
	"storageCapacity":      "If set to true, storageCapacity indicates that the CSI volume driver wants pod scheduling to consider the storage capacity that the driver deployment will report by creating CSIStorageCapacity objects with capacity information.\n\nThe check can be enabled immediately when deploying a driver. In that case, provisioning new volumes with late binding will pause until the driver deployment has published some suitable CSIStorageCapacity object.\n\nAlternatively, the driver can be deployed with the field unset or false and it can be flipped later when storage capacity information has been published.\n\nThis field is immutable.\n\nThis is a beta field and only available when the CSIStorageCapacity feature is enabled. The default is false.",
	"fsGroupPolicy":        "Defines if the underlying volume supports changing ownership and permission of the volume before being mounted. Refer to the specific FSGroupPolicy values for additional details. This field is alpha-level, and is only honored by servers that enable the CSIVolumeFSGroupPolicy feature gate.\n\nThis field is immutable.",
	"tokenRequests":        "TokenRequests indicates the CSI driver needs pods' service account tokens it is mounting volume for to do necessary authentication. Kubelet will pass the tokens in VolumeContext in the CSI NodePublishVolume calls. The CSI driver should parse and validate the following VolumeContext: \"csi.storage.k8s.io/serviceAccount.tokens\": {\n  \"<audience>\": {\n    \"token\": <token>,\n    \"expirationTimestamp\": <expiration timestamp in RFC3339>,\n  },\n  ...\n}\n\nNote: Audience in each TokenRequest should be different and at most one token is empty string. To receive a new token after expiry, RequiresRepublish can be used to trigger NodePublishVolume periodically.\n\nThis is a beta feature and only available when the CSIServiceAccountToken feature is enabled.",
	"requiresRepublish":    "RequiresRepublish indicates the CSI driver wants `NodePublishVolume` being periodically called to reflect any possible change in the mounted volume. This field defaults to false.\n\nNote: After a successful initial NodePublishVolume call, subsequent calls to NodePublishVolume should only update the contents of the volume. New mount points will not be seen by a running container.\n\nThis is a beta feature and only available when the CSIServiceAccountToken feature is enabled.",
}

func (CSIDriverSpec) SwaggerDoc() map[string]string {
	return map_CSIDriverSpec
}

var map_CSINode = map[string]string{
	"":         "CSINode holds information about all CSI drivers installed on a node. CSI drivers do not need to create the CSINode object directly. As long as they use the node-driver-registrar sidecar container, the kubelet will automatically populate the CSINode object for the CSI driver as part of kubelet plugin registration. CSINode has the same name as a node. If the object is missing, it means either there are no CSI Drivers available on the node, or the Kubelet version is low enough that it doesn't create this object. CSINode has an OwnerReference that points to the corresponding node object.",
	"metadata": "metadata.name must be the Kubernetes node name.",
	"spec":     "spec is the specification of CSINode",
}

func (CSINode) SwaggerDoc() map[string]string {
	return map_CSINode
}

var map_CSINodeDriver = map[string]string{
	"":             "CSINodeDriver holds information about the specification of one CSI driver installed on a node",
	"name":         "This is the name of the CSI driver that this object refers to. This MUST be the same name returned by the CSI GetPluginName() call for that driver.",
	"nodeID":       "nodeID of the node from the driver point of view. This field enables Kubernetes to communicate with storage systems that do not share the same nomenclature for nodes. For example, Kubernetes may refer to a given node as \"node1\", but the storage system may refer to the same node as \"nodeA\". When Kubernetes issues a command to the storage system to attach a volume to a specific node, it can use this field to refer to the node name using the ID that the storage system will understand, e.g. \"nodeA\" instead of \"node1\". This field is required.",
	"topologyKeys": "topologyKeys is the list of keys supported by the driver. When a driver is initialized on a cluster, it provides a set of topology keys that it understands (e.g. \"company.com/zone\", \"company.com/region\"). When a driver is initialized on a node, it provides the same topology keys along with values. Kubelet will expose these topology keys as labels on its own node object. When Kubernetes does topology aware provisioning, it can use this list to determine which labels it should retrieve from the node object and pass back to the driver. It is possible for different nodes to use different topology keys. This can be empty if driver does not support topology.",
	"allocatable":  "allocatable represents the volume resources of a node that are available for scheduling. This field is beta.",
}

func (CSINodeDriver) SwaggerDoc() map[string]string {
	return map_CSINodeDriver
}

var map_CSINodeList = map[string]string{
	"":         "CSINodeList is a collection of CSINode objects.",
	"metadata": "Standard list metadata More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata",
	"items":    "items is the list of CSINode",
}

func (CSINodeList) SwaggerDoc() map[string]string {
	return map_CSINodeList
}

var map_CSINodeSpec = map[string]string{
	"":        "CSINodeSpec holds information about the specification of all CSI drivers installed on a node",
	"drivers": "drivers is a list of information of all CSI Drivers existing on a node. If all drivers in the list are uninstalled, this can become empty.",
}

func (CSINodeSpec) SwaggerDoc() map[string]string {
	return map_CSINodeSpec
}

var map_StorageClass = map[string]string{
	"":                     "StorageClass describes the parameters for a class of storage for which PersistentVolumes can be dynamically provisioned.\n\nStorageClasses are non-namespaced; the name of the storage class according to etcd is in ObjectMeta.Name.",
	"metadata":             "Standard object's metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata",
	"provisioner":          "Provisioner indicates the type of the provisioner.",
	"parameters":           "Parameters holds the parameters for the provisioner that should create volumes of this storage class.",
	"reclaimPolicy":        "Dynamically provisioned PersistentVolumes of this storage class are created with this reclaimPolicy. Defaults to Delete.",
	"mountOptions":         "Dynamically provisioned PersistentVolumes of this storage class are created with these mountOptions, e.g. [\"ro\", \"soft\"]. Not validated - mount of the PVs will simply fail if one is invalid.",
	"allowVolumeExpansion": "AllowVolumeExpansion shows whether the storage class allow volume expand",
	"volumeBindingMode":    "VolumeBindingMode indicates how PersistentVolumeClaims should be provisioned and bound.  When unset, VolumeBindingImmediate is used. This field is only honored by servers that enable the VolumeScheduling feature.",
	"allowedTopologies":    "Restrict the node topologies where volumes can be dynamically provisioned. Each volume plugin defines its own supported topology specifications. An empty TopologySelectorTerm list means there is no topology restriction. This field is only honored by servers that enable the VolumeScheduling feature.",
}

func (StorageClass) SwaggerDoc() map[string]string {
	return map_StorageClass
}

var map_StorageClassList = map[string]string{
	"":         "StorageClassList is a collection of storage classes.",
	"metadata": "Standard list metadata More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata",
	"items":    "Items is the list of StorageClasses",
}

func (StorageClassList) SwaggerDoc() map[string]string {
	return map_StorageClassList
}

var map_TokenRequest = map[string]string{
	"":                  "TokenRequest contains parameters of a service account token.",
	"audience":          "Audience is the intended audience of the token in \"TokenRequestSpec\". It will default to the audiences of kube apiserver.",
	"expirationSeconds": "ExpirationSeconds is the duration of validity of the token in \"TokenRequestSpec\". It has the same default value of \"ExpirationSeconds\" in \"TokenRequestSpec\".",
}

func (TokenRequest) SwaggerDoc() map[string]string {
	return map_TokenRequest
}

var map_VolumeAttachment = map[string]string{
	"":         "VolumeAttachment captures the intent to attach or detach the specified volume to/from the specified node.\n\nVolumeAttachment objects are non-namespaced.",
	"metadata": "Standard object metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata",
	"spec":     "Specification of the desired attach/detach volume behavior. Populated by the Kubernetes system.",
	"status":   "Status of the VolumeAttachment request. Populated by the entity completing the attach or detach operation, i.e. the external-attacher.",
}

func (VolumeAttachment) SwaggerDoc() map[string]string {
	return map_VolumeAttachment
}

var map_VolumeAttachmentList = map[string]string{
	"":         "VolumeAttachmentList is a collection of VolumeAttachment objects.",
	"metadata": "Standard list metadata More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata",
	"items":    "Items is the list of VolumeAttachments",
}

func (VolumeAttachmentList) SwaggerDoc() map[string]string {
	return map_VolumeAttachmentList
}

var map_VolumeAttachmentSource = map[string]string{
	"":                     "VolumeAttachmentSource represents a volume that should be attached. Right now only PersistenVolumes can be attached via external attacher, in future we may allow also inline volumes in pods. Exactly one member can be set.",
	"persistentVolumeName": "Name of the persistent volume to attach.",
}

func (VolumeAttachmentSource) SwaggerDoc() map[string]string {
	return map_VolumeAttachmentSource
}

var map_VolumeAttachmentSpec = map[string]string{
	"":         "VolumeAttachmentSpec is the specification of a VolumeAttachment request.",
	"attacher": "Attacher indicates the name of the volume driver that MUST handle this request. This is the name returned by GetPluginName().",
	"source":   "Source represents the volume that should be attached.",
	"nodeName": "The node that the volume should be attached to.",
}

func (VolumeAttachmentSpec) SwaggerDoc() map[string]string {
	return map_VolumeAttachmentSpec
}

var map_VolumeAttachmentStatus = map[string]string{
	"":                   "VolumeAttachmentStatus is the status of a VolumeAttachment request.",
	"attached":           "Indicates the volume is successfully attached. This field must only be set by the entity completing the attach operation, i.e. the external-attacher.",
	"attachmentMetadata": "Upon successful attach, this field is populated with any information returned by the attach operation that must be passed into subsequent WaitForAttach or Mount calls. This field must only be set by the entity completing the attach operation, i.e. the external-attacher.",
	"attachError":        "The last error encountered during attach operation, if any. This field must only be set by the entity completing the attach operation, i.e. the external-attacher.",
	"detachError":        "The last error encountered during detach operation, if any. This field must only be set by the entity completing the detach operation, i.e. the external-attacher.",
}

func (VolumeAttachmentStatus) SwaggerDoc() map[string]string {
	return map_VolumeAttachmentStatus
}

var map_VolumeError = map[string]string{
	"":        "VolumeError captures an error encountered during a volume operation.",
	"time":    "Time the error was encountered.",
	"message": "String detailing the error encountered during Attach or Detach operation. This string may be logged, so it should not contain sensitive information.",
}

func (VolumeError) SwaggerDoc() map[string]string {
	return map_VolumeError
}

var map_VolumeNodeResources = map[string]string{
	"":      "VolumeNodeResources is a set of resource limits for scheduling of volumes.",
	"count": "Maximum number of unique volumes managed by the CSI driver that can be used on a node. A volume that is both attached and mounted on a node is considered to be used once, not twice. The same rule applies for a unique volume that is shared among multiple pods on the same node. If this field is not specified, then the supported number of volumes on this node is unbounded.",
}

func (VolumeNodeResources) SwaggerDoc() map[string]string {
	return map_VolumeNodeResources
}

// AUTO-GENERATED FUNCTIONS END HERE
