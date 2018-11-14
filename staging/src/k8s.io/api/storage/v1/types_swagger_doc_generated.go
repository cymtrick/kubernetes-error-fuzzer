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
var map_StorageClass = map[string]string{
	"":                     "StorageClass describes the parameters for a class of storage for which PersistentVolumes can be dynamically provisioned.\n\nStorageClasses are non-namespaced; the name of the storage class according to etcd is in ObjectMeta.Name.",
	"metadata":             "Standard object's metadata. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#metadata",
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
	"metadata": "Standard list metadata More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#metadata",
	"items":    "Items is the list of StorageClasses",
}

func (StorageClassList) SwaggerDoc() map[string]string {
	return map_StorageClassList
}

var map_VolumeAttachment = map[string]string{
	"":         "VolumeAttachment captures the intent to attach or detach the specified volume to/from the specified node.\n\nVolumeAttachment objects are non-namespaced.",
	"metadata": "Standard object metadata. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#metadata",
	"spec":     "Specification of the desired attach/detach volume behavior. Populated by the Kubernetes system.",
	"status":   "Status of the VolumeAttachment request. Populated by the entity completing the attach or detach operation, i.e. the external-attacher.",
}

func (VolumeAttachment) SwaggerDoc() map[string]string {
	return map_VolumeAttachment
}

var map_VolumeAttachmentList = map[string]string{
	"":         "VolumeAttachmentList is a collection of VolumeAttachment objects.",
	"metadata": "Standard list metadata More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#metadata",
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
	"message": "String detailing the error encountered during Attach or Detach operation. This string maybe logged, so it should not contain sensitive information.",
}

func (VolumeError) SwaggerDoc() map[string]string {
	return map_VolumeError
}

// AUTO-GENERATED FUNCTIONS END HERE
