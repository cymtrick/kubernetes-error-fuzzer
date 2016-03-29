/*
Copyright 2015 The Kubernetes Authors All rights reserved.

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

// AUTO-GENERATED FUNCTIONS START HERE
var map_LocalSubjectAccessReview = map[string]string{
	"":       "LocalSubjectAccessReview checks whether or not a user or group can perform an action in a given namespace. Having a namespace scoped resource makes it much easier to grant namespace scoped policy that includes permissions checking.",
	"spec":   "Spec holds information about the request being evaluated.  spec.namespace must be equal to the namespace you made the request against.  If empty, it is defaulted.",
	"status": "Status is filled in by the server and indicates whether the request is allowed or not",
}

func (LocalSubjectAccessReview) SwaggerDoc() map[string]string {
	return map_LocalSubjectAccessReview
}

var map_NonResourceAttributes = map[string]string{
	"":     "NonResourceAttributes includes the authorization attributes available for non-resource requests to the Authorizer interface",
	"path": "Path is the URL path of the request",
	"verb": "Verb is the standard HTTP verb",
}

func (NonResourceAttributes) SwaggerDoc() map[string]string {
	return map_NonResourceAttributes
}

var map_ResourceAttributes = map[string]string{
	"":            "ResourceAttributes includes the authorization attributes available for resource requests to the Authorizer interface",
	"namespace":   "Namespace is the namespace of the action being requested.  Currently, there is no distinction between no namespace and all namespaces \"\" (empty) is defaulted for LocalSubjectAccessReviews \"\" (empty) is empty for cluster-scoped resources \"\" (empty) means \"all\" for namespace scoped resources from a SubjectAccessReview or SelfSubjectAccessReview",
	"verb":        "Verb is a kubernetes resource API verb, like: get, list, watch, create, update, delete, proxy.  \"*\" means all.",
	"group":       "Group is the API Group of the Resource.  \"*\" means all.",
	"version":     "Version is the API Version of the Resource.  \"*\" means all.",
	"resource":    "Resource is one of the existing resource types.  \"*\" means all.",
	"subresource": "Subresource is one of the existing resource types.  \"\" means none.",
	"name":        "Name is the name of the resource being requested for a \"get\" or deleted for a \"delete\". \"\" (empty) means all.",
}

func (ResourceAttributes) SwaggerDoc() map[string]string {
	return map_ResourceAttributes
}

var map_SelfSubjectAccessReview = map[string]string{
	"":       "SelfSubjectAccessReview checks whether or the current user can perform an action.  Not filling in a spec.namespace means \"in all namespaces\".  Self is a special case, because users should always be able to check whether they can perform an action",
	"spec":   "Spec holds information about the request being evaluated.  user and groups must be empty",
	"status": "Status is filled in by the server and indicates whether the request is allowed or not",
}

func (SelfSubjectAccessReview) SwaggerDoc() map[string]string {
	return map_SelfSubjectAccessReview
}

var map_SelfSubjectAccessReviewSpec = map[string]string{
	"":                      "SelfSubjectAccessReviewSpec is a description of the access request.  Exactly one of ResourceAuthorizationAttributes and NonResourceAuthorizationAttributes must be set",
	"resourceAttributes":    "ResourceAuthorizationAttributes describes information for a resource access request",
	"nonResourceAttributes": "NonResourceAttributes describes information for a non-resource access request",
}

func (SelfSubjectAccessReviewSpec) SwaggerDoc() map[string]string {
	return map_SelfSubjectAccessReviewSpec
}

var map_SubjectAccessReview = map[string]string{
	"":       "SubjectAccessReview checks whether or not a user or group can perform an action.",
	"spec":   "Spec holds information about the request being evaluated",
	"status": "Status is filled in by the server and indicates whether the request is allowed or not",
}

func (SubjectAccessReview) SwaggerDoc() map[string]string {
	return map_SubjectAccessReview
}

var map_SubjectAccessReviewSpec = map[string]string{
	"":                      "SubjectAccessReviewSpec is a description of the access request.  Exactly one of ResourceAuthorizationAttributes and NonResourceAuthorizationAttributes must be set",
	"resourceAttributes":    "ResourceAuthorizationAttributes describes information for a resource access request",
	"nonResourceAttributes": "NonResourceAttributes describes information for a non-resource access request",
	"user":                  "User is the user you're testing for. If you specify \"User\" but not \"Group\", then is it interpreted as \"What if User were not a member of any groups",
	"group":                 "Groups is the groups you're testing for.",
	"extra":                 "Extra corresponds to the user.Info.GetExtra() method from the authenticator.  Since that is input to the authorizer it needs a reflection here.",
}

func (SubjectAccessReviewSpec) SwaggerDoc() map[string]string {
	return map_SubjectAccessReviewSpec
}

var map_SubjectAccessReviewStatus = map[string]string{
	"":        "SubjectAccessReviewStatus",
	"allowed": "Allowed is required.  True if the action would be allowed, false otherwise.",
	"reason":  "Reason is optional.  It indicates why a request was allowed or denied.",
}

func (SubjectAccessReviewStatus) SwaggerDoc() map[string]string {
	return map_SubjectAccessReviewStatus
}

// AUTO-GENERATED FUNCTIONS END HERE
