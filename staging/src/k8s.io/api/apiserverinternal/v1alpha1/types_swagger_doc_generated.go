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

package v1alpha1

// This file contains a collection of methods that can be used from go-restful to
// generate Swagger API documentation for its models. Please read this PR for more
// information on the implementation: https://github.com/emicklei/go-restful/pull/215
//
// TODOs are ignored from the parser (e.g. TODO(andronat):... || TODO:...) if and only if
// they are on one line! For multiple line or blocks that you want to ignore use ---.
// Any context after a --- is ignored.
//
// Those methods can be generated by using hack/update-codegen.sh

// AUTO-GENERATED FUNCTIONS START HERE. DO NOT EDIT.
var map_ServerStorageVersion = map[string]string{
	"":                  "An API server instance reports the version it can decode and the version it encodes objects to when persisting objects in the backend.",
	"apiServerID":       "The ID of the reporting API server.",
	"encodingVersion":   "The API server encodes the object to this version when persisting it in the backend (e.g., etcd).",
	"decodableVersions": "The API server can decode objects encoded in these versions. The encodingVersion must be included in the decodableVersions.",
	"servedVersions":    "The API server can serve these versions. DecodableVersions must include all ServedVersions.",
}

func (ServerStorageVersion) SwaggerDoc() map[string]string {
	return map_ServerStorageVersion
}

var map_StorageVersion = map[string]string{
	"":         "Storage version of a specific resource.",
	"metadata": "The name is <group>.<resource>.",
	"spec":     "Spec is an empty spec. It is here to comply with Kubernetes API style.",
	"status":   "API server instances report the version they can decode and the version they encode objects to when persisting objects in the backend.",
}

func (StorageVersion) SwaggerDoc() map[string]string {
	return map_StorageVersion
}

var map_StorageVersionCondition = map[string]string{
	"":                   "Describes the state of the storageVersion at a certain point.",
	"type":               "Type of the condition.",
	"status":             "Status of the condition, one of True, False, Unknown.",
	"observedGeneration": "If set, this represents the .metadata.generation that the condition was set based upon.",
	"lastTransitionTime": "Last time the condition transitioned from one status to another.",
	"reason":             "The reason for the condition's last transition.",
	"message":            "A human readable message indicating details about the transition.",
}

func (StorageVersionCondition) SwaggerDoc() map[string]string {
	return map_StorageVersionCondition
}

var map_StorageVersionList = map[string]string{
	"":         "A list of StorageVersions.",
	"metadata": "Standard list metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata",
	"items":    "Items holds a list of StorageVersion",
}

func (StorageVersionList) SwaggerDoc() map[string]string {
	return map_StorageVersionList
}

var map_StorageVersionSpec = map[string]string{
	"": "StorageVersionSpec is an empty spec.",
}

func (StorageVersionSpec) SwaggerDoc() map[string]string {
	return map_StorageVersionSpec
}

var map_StorageVersionStatus = map[string]string{
	"":                      "API server instances report the versions they can decode and the version they encode objects to when persisting objects in the backend.",
	"storageVersions":       "The reported versions per API server instance.",
	"commonEncodingVersion": "If all API server instances agree on the same encoding storage version, then this field is set to that version. Otherwise this field is left empty. API servers should finish updating its storageVersionStatus entry before serving write operations, so that this field will be in sync with the reality.",
	"conditions":            "The latest available observations of the storageVersion's state.",
}

func (StorageVersionStatus) SwaggerDoc() map[string]string {
	return map_StorageVersionStatus
}

// AUTO-GENERATED FUNCTIONS END HERE
