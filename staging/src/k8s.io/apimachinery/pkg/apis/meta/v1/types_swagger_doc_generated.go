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
var map_APIGroup = map[string]string{
	"":                           "APIGroup contains the name, the supported versions, and the preferred version of a group.",
	"name":                       "name is the name of the group.",
	"versions":                   "versions are the versions supported in this group.",
	"preferredVersion":           "preferredVersion is the version preferred by the API server, which probably is the storage version.",
	"serverAddressByClientCIDRs": "a map of client CIDR to server address that is serving this group. This is to help clients reach servers in the most network-efficient way possible. Clients can use the appropriate server address as per the CIDR that they match. In case of multiple matches, clients should use the longest matching CIDR. The server returns only those CIDRs that it thinks that the client can match. For example: the master will return an internal IP CIDR only, if the client reaches the server using an internal IP. Server looks at X-Forwarded-For header or X-Real-Ip header or request.RemoteAddr (in that order) to get the client IP.",
}

func (APIGroup) SwaggerDoc() map[string]string {
	return map_APIGroup
}

var map_APIGroupList = map[string]string{
	"":       "APIGroupList is a list of APIGroup, to allow clients to discover the API at /apis.",
	"groups": "groups is a list of APIGroup.",
}

func (APIGroupList) SwaggerDoc() map[string]string {
	return map_APIGroupList
}

var map_APIResource = map[string]string{
	"":             "APIResource specifies the name of a resource and whether it is namespaced.",
	"name":         "name is the plural name of the resource.",
	"singularName": "singularName is the singular name of the resource.  This allows clients to handle plural and singular opaquely. The singularName is more correct for reporting status on a single item and both singular and plural are allowed from the kubectl CLI interface.",
	"namespaced":   "namespaced indicates if a resource is namespaced or not.",
	"group":        "group is the preferred group of the resource.  Empty implies the group of the containing resource list. For subresources, this may have a different value, for example: Scale\".",
	"version":      "version is the preferred version of the resource.  Empty implies the version of the containing resource list For subresources, this may have a different value, for example: v1 (while inside a v1beta1 version of the core resource's group)\".",
	"kind":         "kind is the kind for the resource (e.g. 'Foo' is the kind for a resource 'foo')",
	"verbs":        "verbs is a list of supported kube verbs (this includes get, list, watch, create, update, patch, delete, deletecollection, and proxy)",
	"shortNames":   "shortNames is a list of suggested short names of the resource.",
	"categories":   "categories is a list of the grouped resources this resource belongs to (e.g. 'all')",
}

func (APIResource) SwaggerDoc() map[string]string {
	return map_APIResource
}

var map_APIResourceList = map[string]string{
	"":             "APIResourceList is a list of APIResource, it is used to expose the name of the resources supported in a specific group and version, and if the resource is namespaced.",
	"groupVersion": "groupVersion is the group and version this APIResourceList is for.",
	"resources":    "resources contains the name of the resources and if they are namespaced.",
}

func (APIResourceList) SwaggerDoc() map[string]string {
	return map_APIResourceList
}

var map_APIVersions = map[string]string{
	"":                           "APIVersions lists the versions that are available, to allow clients to discover the API at /api, which is the root path of the legacy v1 API.",
	"versions":                   "versions are the api versions that are available.",
	"serverAddressByClientCIDRs": "a map of client CIDR to server address that is serving this group. This is to help clients reach servers in the most network-efficient way possible. Clients can use the appropriate server address as per the CIDR that they match. In case of multiple matches, clients should use the longest matching CIDR. The server returns only those CIDRs that it thinks that the client can match. For example: the master will return an internal IP CIDR only, if the client reaches the server using an internal IP. Server looks at X-Forwarded-For header or X-Real-Ip header or request.RemoteAddr (in that order) to get the client IP.",
}

func (APIVersions) SwaggerDoc() map[string]string {
	return map_APIVersions
}

var map_CreateOptions = map[string]string{
	"":       "CreateOptions may be provided when creating an API object.",
	"dryRun": "When present, indicates that modifications should not be persisted. An invalid or unrecognized dryRun directive will result in an error response and no further processing of the request. Valid values are: - All: all dry run stages will be processed",
}

func (CreateOptions) SwaggerDoc() map[string]string {
	return map_CreateOptions
}

var map_DeleteOptions = map[string]string{
	"":                   "DeleteOptions may be provided when deleting an API object.",
	"gracePeriodSeconds": "The duration in seconds before the object should be deleted. Value must be non-negative integer. The value zero indicates delete immediately. If this value is nil, the default grace period for the specified type will be used. Defaults to a per object value if not specified. zero means delete immediately.",
	"preconditions":      "Must be fulfilled before a deletion is carried out. If not possible, a 409 Conflict status will be returned.",
	"orphanDependents":   "Deprecated: please use the PropagationPolicy, this field will be deprecated in 1.7. Should the dependent objects be orphaned. If true/false, the \"orphan\" finalizer will be added to/removed from the object's finalizers list. Either this field or PropagationPolicy may be set, but not both.",
	"propagationPolicy":  "Whether and how garbage collection will be performed. Either this field or OrphanDependents may be set, but not both. The default policy is decided by the existing finalizer set in the metadata.finalizers and the resource-specific default policy. Acceptable values are: 'Orphan' - orphan the dependents; 'Background' - allow the garbage collector to delete the dependents in the background; 'Foreground' - a cascading policy that deletes all dependents in the foreground.",
	"dryRun":             "When present, indicates that modifications should not be persisted. An invalid or unrecognized dryRun directive will result in an error response and no further processing of the request. Valid values are: - All: all dry run stages will be processed",
}

func (DeleteOptions) SwaggerDoc() map[string]string {
	return map_DeleteOptions
}

var map_ExportOptions = map[string]string{
	"":       "ExportOptions is the query options to the standard REST get call.",
	"export": "Should this value be exported.  Export strips fields that a user can not specify.",
	"exact":  "Should the export be exact.  Exact export maintains cluster-specific fields like 'Namespace'.",
}

func (ExportOptions) SwaggerDoc() map[string]string {
	return map_ExportOptions
}

var map_GetOptions = map[string]string{
	"":                "GetOptions is the standard query options to the standard REST get call.",
	"resourceVersion": "When specified: - if unset, then the result is returned from remote storage based on quorum-read flag; - if it's 0, then we simply return what we currently have in cache, no guarantee; - if set to non zero, then the result is at least as fresh as given rv.",
}

func (GetOptions) SwaggerDoc() map[string]string {
	return map_GetOptions
}

var map_GroupVersionForDiscovery = map[string]string{
	"":             "GroupVersion contains the \"group/version\" and \"version\" string of a version. It is made a struct to keep extensibility.",
	"groupVersion": "groupVersion specifies the API group and version in the form \"group/version\"",
	"version":      "version specifies the version in the form of \"version\". This is to save the clients the trouble of splitting the GroupVersion.",
}

func (GroupVersionForDiscovery) SwaggerDoc() map[string]string {
	return map_GroupVersionForDiscovery
}

var map_Initializer = map[string]string{
	"":     "Initializer is information about an initializer that has not yet completed.",
	"name": "name of the process that is responsible for initializing this object.",
}

func (Initializer) SwaggerDoc() map[string]string {
	return map_Initializer
}

var map_Initializers = map[string]string{
	"":        "Initializers tracks the progress of initialization.",
	"pending": "Pending is a list of initializers that must execute in order before this object is visible. When the last pending initializer is removed, and no failing result is set, the initializers struct will be set to nil and the object is considered as initialized and visible to all clients.",
	"result":  "If result is set with the Failure field, the object will be persisted to storage and then deleted, ensuring that other clients can observe the deletion.",
}

func (Initializers) SwaggerDoc() map[string]string {
	return map_Initializers
}

var map_LabelSelector = map[string]string{
	"":                 "A label selector is a label query over a set of resources. The result of matchLabels and matchExpressions are ANDed. An empty label selector matches all objects. A null label selector matches no objects.",
	"matchLabels":      "matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabels map is equivalent to an element of matchExpressions, whose key field is \"key\", the operator is \"In\", and the values array contains only \"value\". The requirements are ANDed.",
	"matchExpressions": "matchExpressions is a list of label selector requirements. The requirements are ANDed.",
}

func (LabelSelector) SwaggerDoc() map[string]string {
	return map_LabelSelector
}

var map_LabelSelectorRequirement = map[string]string{
	"":         "A label selector requirement is a selector that contains values, a key, and an operator that relates the key and values.",
	"key":      "key is the label key that the selector applies to.",
	"operator": "operator represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists and DoesNotExist.",
	"values":   "values is an array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. This array is replaced during a strategic merge patch.",
}

func (LabelSelectorRequirement) SwaggerDoc() map[string]string {
	return map_LabelSelectorRequirement
}

var map_List = map[string]string{
	"":         "List holds a list of objects, which may not be known by the server.",
	"metadata": "Standard list metadata. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds",
	"items":    "List of objects",
}

func (List) SwaggerDoc() map[string]string {
	return map_List
}

var map_ListMeta = map[string]string{
	"":                "ListMeta describes metadata that synthetic resources must have, including lists and various status objects. A resource may have only one of {ObjectMeta, ListMeta}.",
	"selfLink":        "selfLink is a URL representing this object. Populated by the system. Read-only.",
	"resourceVersion": "String that identifies the server's internal version of this object that can be used by clients to determine when objects have changed. Value must be treated as opaque by clients and passed unmodified back to the server. Populated by the system. Read-only. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#concurrency-control-and-consistency",
	"continue":        "continue may be set if the user set a limit on the number of items returned, and indicates that the server has more data available. The value is opaque and may be used to issue another request to the endpoint that served this list to retrieve the next set of available objects. Continuing a consistent list may not be possible if the server configuration has changed or more than a few minutes have passed. The resourceVersion field returned when using this continue value will be identical to the value in the first response, unless you have received this token from an error message.",
}

func (ListMeta) SwaggerDoc() map[string]string {
	return map_ListMeta
}

var map_ListOptions = map[string]string{
	"":                "ListOptions is the query options to a standard REST list call.",
	"labelSelector":   "A selector to restrict the list of returned objects by their labels. Defaults to everything.",
	"fieldSelector":   "A selector to restrict the list of returned objects by their fields. Defaults to everything.",
	"watch":           "Watch for changes to the described resources and return them as a stream of add, update, and remove notifications. Specify resourceVersion.",
	"resourceVersion": "When specified with a watch call, shows changes that occur after that particular version of a resource. Defaults to changes from the beginning of history. When specified for list: - if unset, then the result is returned from remote storage based on quorum-read flag; - if it's 0, then we simply return what we currently have in cache, no guarantee; - if set to non zero, then the result is at least as fresh as given rv.",
	"timeoutSeconds":  "Timeout for the list/watch call. This limits the duration of the call, regardless of any activity or inactivity.",
	"limit":           "limit is a maximum number of responses to return for a list call. If more items exist, the server will set the `continue` field on the list metadata to a value that can be used with the same initial query to retrieve the next set of results. Setting a limit may return fewer than the requested amount of items (up to zero items) in the event all requested objects are filtered out and clients should only use the presence of the continue field to determine whether more results are available. Servers may choose not to support the limit argument and will return all of the available results. If limit is specified and the continue field is empty, clients may assume that no more results are available. This field is not supported if watch is true.\n\nThe server guarantees that the objects returned when using continue will be identical to issuing a single list call without a limit - that is, no objects created, modified, or deleted after the first request is issued will be included in any subsequent continued requests. This is sometimes referred to as a consistent snapshot, and ensures that a client that is using limit to receive smaller chunks of a very large result can ensure they see all possible objects. If objects are updated during a chunked list the version of the object that was present at the time the first list result was calculated is returned.",
	"continue":        "The continue option should be set when retrieving more results from the server. Since this value is server defined, clients may only use the continue value from a previous query result with identical query parameters (except for the value of continue) and the server may reject a continue value it does not recognize. If the specified continue value is no longer valid whether due to expiration (generally five to fifteen minutes) or a configuration change on the server, the server will respond with a 410 ResourceExpired error together with a continue token. If the client needs a consistent list, it must restart their list without the continue field. Otherwise, the client may send another list request with the token received with the 410 error, the server will respond with a list starting from the next key, but from the latest snapshot, which is inconsistent from the previous list results - objects that are created, modified, or deleted after the first list request will be included in the response, as long as their keys are after the \"next key\".\n\nThis field is not supported when watch is true. Clients may start a watch from the last resourceVersion value returned by the server and not miss any modifications.",
}

func (ListOptions) SwaggerDoc() map[string]string {
	return map_ListOptions
}

var map_ObjectMeta = map[string]string{
	"":                           "ObjectMeta is metadata that all persisted resources must have, which includes all objects users must create.",
	"name":                       "Name must be unique within a namespace. Is required when creating resources, although some resources may allow a client to request the generation of an appropriate name automatically. Name is primarily intended for creation idempotence and configuration definition. Cannot be updated. More info: http://kubernetes.io/docs/user-guide/identifiers#names",
	"generateName":               "GenerateName is an optional prefix, used by the server, to generate a unique name ONLY IF the Name field has not been provided. If this field is used, the name returned to the client will be different than the name passed. This value will also be combined with a unique suffix. The provided value has the same validation rules as the Name field, and may be truncated by the length of the suffix required to make the value unique on the server.\n\nIf this field is specified and the generated name exists, the server will NOT return a 409 - instead, it will either return 201 Created or 500 with Reason ServerTimeout indicating a unique name could not be found in the time allotted, and the client should retry (optionally after the time indicated in the Retry-After header).\n\nApplied only if Name is not specified. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#idempotency",
	"namespace":                  "Namespace defines the space within each name must be unique. An empty namespace is equivalent to the \"default\" namespace, but \"default\" is the canonical representation. Not all objects are required to be scoped to a namespace - the value of this field for those objects will be empty.\n\nMust be a DNS_LABEL. Cannot be updated. More info: http://kubernetes.io/docs/user-guide/namespaces",
	"selfLink":                   "SelfLink is a URL representing this object. Populated by the system. Read-only.",
	"uid":                        "UID is the unique in time and space value for this object. It is typically generated by the server on successful creation of a resource and is not allowed to change on PUT operations.\n\nPopulated by the system. Read-only. More info: http://kubernetes.io/docs/user-guide/identifiers#uids",
	"resourceVersion":            "An opaque value that represents the internal version of this object that can be used by clients to determine when objects have changed. May be used for optimistic concurrency, change detection, and the watch operation on a resource or set of resources. Clients must treat these values as opaque and passed unmodified back to the server. They may only be valid for a particular resource or set of resources.\n\nPopulated by the system. Read-only. Value must be treated as opaque by clients and . More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#concurrency-control-and-consistency",
	"generation":                 "A sequence number representing a specific generation of the desired state. Populated by the system. Read-only.",
	"creationTimestamp":          "CreationTimestamp is a timestamp representing the server time when this object was created. It is not guaranteed to be set in happens-before order across separate operations. Clients may not set this value. It is represented in RFC3339 form and is in UTC.\n\nPopulated by the system. Read-only. Null for lists. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#metadata",
	"deletionTimestamp":          "DeletionTimestamp is RFC 3339 date and time at which this resource will be deleted. This field is set by the server when a graceful deletion is requested by the user, and is not directly settable by a client. The resource is expected to be deleted (no longer visible from resource lists, and not reachable by name) after the time in this field, once the finalizers list is empty. As long as the finalizers list contains items, deletion is blocked. Once the deletionTimestamp is set, this value may not be unset or be set further into the future, although it may be shortened or the resource may be deleted prior to this time. For example, a user may request that a pod is deleted in 30 seconds. The Kubelet will react by sending a graceful termination signal to the containers in the pod. After that 30 seconds, the Kubelet will send a hard termination signal (SIGKILL) to the container and after cleanup, remove the pod from the API. In the presence of network partitions, this object may still exist after this timestamp, until an administrator or automated process can determine the resource is fully terminated. If not set, graceful deletion of the object has not been requested.\n\nPopulated by the system when a graceful deletion is requested. Read-only. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#metadata",
	"deletionGracePeriodSeconds": "Number of seconds allowed for this object to gracefully terminate before it will be removed from the system. Only set when deletionTimestamp is also set. May only be shortened. Read-only.",
	"labels":                     "Map of string keys and values that can be used to organize and categorize (scope and select) objects. May match selectors of replication controllers and services. More info: http://kubernetes.io/docs/user-guide/labels",
	"annotations":                "Annotations is an unstructured key value map stored with a resource that may be set by external tools to store and retrieve arbitrary metadata. They are not queryable and should be preserved when modifying objects. More info: http://kubernetes.io/docs/user-guide/annotations",
	"ownerReferences":            "List of objects depended by this object. If ALL objects in the list have been deleted, this object will be garbage collected. If this object is managed by a controller, then an entry in this list will point to this controller, with the controller field set to true. There cannot be more than one managing controller.",
	"initializers":               "An initializer is a controller which enforces some system invariant at object creation time. This field is a list of initializers that have not yet acted on this object. If nil or empty, this object has been completely initialized. Otherwise, the object is considered uninitialized and is hidden (in list/watch and get calls) from clients that haven't explicitly asked to observe uninitialized objects.\n\nWhen an object is created, the system will populate this list with the current set of initializers. Only privileged users may set or modify this list. Once it is empty, it may not be modified further by any user.",
	"finalizers":                 "Must be empty before the object is deleted from the registry. Each entry is an identifier for the responsible component that will remove the entry from the list. If the deletionTimestamp of the object is non-nil, entries in this list can only be removed.",
	"clusterName":                "The name of the cluster which the object belongs to. This is used to distinguish resources with same name and namespace in different clusters. This field is not set anywhere right now and apiserver is going to ignore it if set in create or update request.",
}

func (ObjectMeta) SwaggerDoc() map[string]string {
	return map_ObjectMeta
}

var map_OwnerReference = map[string]string{
	"":                   "OwnerReference contains enough information to let you identify an owning object. An owning object must be in the same namespace as the dependent, or be cluster-scoped, so there is no namespace field.",
	"apiVersion":         "API version of the referent.",
	"kind":               "Kind of the referent. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds",
	"name":               "Name of the referent. More info: http://kubernetes.io/docs/user-guide/identifiers#names",
	"uid":                "UID of the referent. More info: http://kubernetes.io/docs/user-guide/identifiers#uids",
	"controller":         "If true, this reference points to the managing controller.",
	"blockOwnerDeletion": "If true, AND if the owner has the \"foregroundDeletion\" finalizer, then the owner cannot be deleted from the key-value store until this reference is removed. Defaults to false. To set this field, a user needs \"delete\" permission of the owner, otherwise 422 (Unprocessable Entity) will be returned.",
}

func (OwnerReference) SwaggerDoc() map[string]string {
	return map_OwnerReference
}

var map_Patch = map[string]string{
	"": "Patch is provided to give a concrete name and type to the Kubernetes PATCH request body.",
}

func (Patch) SwaggerDoc() map[string]string {
	return map_Patch
}

var map_Preconditions = map[string]string{
	"":    "Preconditions must be fulfilled before an operation (update, delete, etc.) is carried out.",
	"uid": "Specifies the target UID.",
}

func (Preconditions) SwaggerDoc() map[string]string {
	return map_Preconditions
}

var map_RootPaths = map[string]string{
	"":      "RootPaths lists the paths available at root. For example: \"/healthz\", \"/apis\".",
	"paths": "paths are the paths available at root.",
}

func (RootPaths) SwaggerDoc() map[string]string {
	return map_RootPaths
}

var map_ServerAddressByClientCIDR = map[string]string{
	"":              "ServerAddressByClientCIDR helps the client to determine the server address that they should use, depending on the clientCIDR that they match.",
	"clientCIDR":    "The CIDR with which clients can match their IP to figure out the server address that they should use.",
	"serverAddress": "Address of this server, suitable for a client that matches the above CIDR. This can be a hostname, hostname:port, IP or IP:port.",
}

func (ServerAddressByClientCIDR) SwaggerDoc() map[string]string {
	return map_ServerAddressByClientCIDR
}

var map_Status = map[string]string{
	"":         "Status is a return value for calls that don't return other objects.",
	"metadata": "Standard list metadata. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds",
	"status":   "Status of the operation. One of: \"Success\" or \"Failure\". More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#spec-and-status",
	"message":  "A human-readable description of the status of this operation.",
	"reason":   "A machine-readable description of why this operation is in the \"Failure\" status. If this value is empty there is no information available. A Reason clarifies an HTTP status code but does not override it.",
	"details":  "Extended data associated with the reason.  Each reason may define its own extended details. This field is optional and the data returned is not guaranteed to conform to any schema except that defined by the reason type.",
	"code":     "Suggested HTTP return code for this status, 0 if not set.",
}

func (Status) SwaggerDoc() map[string]string {
	return map_Status
}

var map_StatusCause = map[string]string{
	"":        "StatusCause provides more information about an api.Status failure, including cases when multiple errors are encountered.",
	"reason":  "A machine-readable description of the cause of the error. If this value is empty there is no information available.",
	"message": "A human-readable description of the cause of the error.  This field may be presented as-is to a reader.",
	"field":   "The field of the resource that has caused this error, as named by its JSON serialization. May include dot and postfix notation for nested attributes. Arrays are zero-indexed.  Fields may appear more than once in an array of causes due to fields having multiple errors. Optional.\n\nExamples:\n  \"name\" - the field \"name\" on the current resource\n  \"items[0].name\" - the field \"name\" on the first array entry in \"items\"",
}

func (StatusCause) SwaggerDoc() map[string]string {
	return map_StatusCause
}

var map_StatusDetails = map[string]string{
	"":                  "StatusDetails is a set of additional properties that MAY be set by the server to provide additional information about a response. The Reason field of a Status object defines what attributes will be set. Clients must ignore fields that do not match the defined type of each attribute, and should assume that any attribute may be empty, invalid, or under defined.",
	"name":              "The name attribute of the resource associated with the status StatusReason (when there is a single name which can be described).",
	"group":             "The group attribute of the resource associated with the status StatusReason.",
	"kind":              "The kind attribute of the resource associated with the status StatusReason. On some operations may differ from the requested resource Kind. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds",
	"uid":               "UID of the resource. (when there is a single resource which can be described). More info: http://kubernetes.io/docs/user-guide/identifiers#uids",
	"causes":            "The Causes array includes more details associated with the StatusReason failure. Not all StatusReasons may provide detailed causes.",
	"retryAfterSeconds": "If specified, the time in seconds before the operation should be retried. Some errors may indicate the client must take an alternate action - for those errors this field may indicate how long to wait before taking the alternate action.",
}

func (StatusDetails) SwaggerDoc() map[string]string {
	return map_StatusDetails
}

var map_TypeMeta = map[string]string{
	"":           "TypeMeta describes an individual object in an API response or request with strings representing the type of the object and its API schema version. Structures that are versioned or persisted should inline TypeMeta.",
	"kind":       "Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds",
	"apiVersion": "APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#resources",
}

func (TypeMeta) SwaggerDoc() map[string]string {
	return map_TypeMeta
}

var map_UpdateOptions = map[string]string{
	"":       "UpdateOptions may be provided when updating an API object.",
	"dryRun": "When present, indicates that modifications should not be persisted. An invalid or unrecognized dryRun directive will result in an error response and no further processing of the request. Valid values are: - All: all dry run stages will be processed",
}

func (UpdateOptions) SwaggerDoc() map[string]string {
	return map_UpdateOptions
}

// AUTO-GENERATED FUNCTIONS END HERE
