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
// Those methods can be generated by using hack/update-codegen.sh

// AUTO-GENERATED FUNCTIONS START HERE. DO NOT EDIT.
var map_BoundObjectReference = map[string]string{
	"":           "BoundObjectReference is a reference to an object that a token is bound to.",
	"kind":       "Kind of the referent. Valid kinds are 'Pod' and 'Secret'.",
	"apiVersion": "API version of the referent.",
	"name":       "Name of the referent.",
	"uid":        "UID of the referent.",
}

func (BoundObjectReference) SwaggerDoc() map[string]string {
	return map_BoundObjectReference
}

var map_TokenRequest = map[string]string{
	"":         "TokenRequest requests a token for a given service account.",
	"metadata": "Standard object's metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata",
	"spec":     "Spec holds information about the request being evaluated",
	"status":   "Status is filled in by the server and indicates whether the token can be authenticated.",
}

func (TokenRequest) SwaggerDoc() map[string]string {
	return map_TokenRequest
}

var map_TokenRequestSpec = map[string]string{
	"":                  "TokenRequestSpec contains client provided parameters of a token request.",
	"audiences":         "Audiences are the intendend audiences of the token. A recipient of a token must identify themself with an identifier in the list of audiences of the token, and otherwise should reject the token. A token issued for multiple audiences may be used to authenticate against any of the audiences listed but implies a high degree of trust between the target audiences.",
	"expirationSeconds": "ExpirationSeconds is the requested duration of validity of the request. The token issuer may return a token with a different validity duration so a client needs to check the 'expiration' field in a response.",
	"boundObjectRef":    "BoundObjectRef is a reference to an object that the token will be bound to. The token will only be valid for as long as the bound object exists. NOTE: The API server's TokenReview endpoint will validate the BoundObjectRef, but other audiences may not. Keep ExpirationSeconds small if you want prompt revocation.",
}

func (TokenRequestSpec) SwaggerDoc() map[string]string {
	return map_TokenRequestSpec
}

var map_TokenRequestStatus = map[string]string{
	"":                    "TokenRequestStatus is the result of a token request.",
	"token":               "Token is the opaque bearer token.",
	"expirationTimestamp": "ExpirationTimestamp is the time of expiration of the returned token.",
}

func (TokenRequestStatus) SwaggerDoc() map[string]string {
	return map_TokenRequestStatus
}

var map_TokenReview = map[string]string{
	"":         "TokenReview attempts to authenticate a token to a known user. Note: TokenReview requests may be cached by the webhook token authenticator plugin in the kube-apiserver.",
	"metadata": "Standard object's metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata",
	"spec":     "Spec holds information about the request being evaluated",
	"status":   "Status is filled in by the server and indicates whether the request can be authenticated.",
}

func (TokenReview) SwaggerDoc() map[string]string {
	return map_TokenReview
}

var map_TokenReviewSpec = map[string]string{
	"":          "TokenReviewSpec is a description of the token authentication request.",
	"token":     "Token is the opaque bearer token.",
	"audiences": "Audiences is a list of the identifiers that the resource server presented with the token identifies as. Audience-aware token authenticators will verify that the token was intended for at least one of the audiences in this list. If no audiences are provided, the audience will default to the audience of the Kubernetes apiserver.",
}

func (TokenReviewSpec) SwaggerDoc() map[string]string {
	return map_TokenReviewSpec
}

var map_TokenReviewStatus = map[string]string{
	"":              "TokenReviewStatus is the result of the token authentication request.",
	"authenticated": "Authenticated indicates that the token was associated with a known user.",
	"user":          "User is the UserInfo associated with the provided token.",
	"audiences":     "Audiences are audience identifiers chosen by the authenticator that are compatible with both the TokenReview and token. An identifier is any identifier in the intersection of the TokenReviewSpec audiences and the token's audiences. A client of the TokenReview API that sets the spec.audiences field should validate that a compatible audience identifier is returned in the status.audiences field to ensure that the TokenReview server is audience aware. If a TokenReview returns an empty status.audience field where status.authenticated is \"true\", the token is valid against the audience of the Kubernetes API server.",
	"error":         "Error indicates that the token couldn't be checked",
}

func (TokenReviewStatus) SwaggerDoc() map[string]string {
	return map_TokenReviewStatus
}

var map_UserInfo = map[string]string{
	"":         "UserInfo holds the information about the user needed to implement the user.Info interface.",
	"username": "The name that uniquely identifies this user among all active users.",
	"uid":      "A unique value that identifies this user across time. If this user is deleted and another user by the same name is added, they will have different UIDs.",
	"groups":   "The names of groups this user is a part of.",
	"extra":    "Any additional information provided by the authenticator.",
}

func (UserInfo) SwaggerDoc() map[string]string {
	return map_UserInfo
}

// AUTO-GENERATED FUNCTIONS END HERE
