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

package v1alpha1

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
var map_CertificateSigningRequest = map[string]string{
	"":       "Describes a certificate signing request",
	"spec":   "The certificate request itself and any additional information.",
	"status": "Derived information about the request.",
}

func (CertificateSigningRequest) SwaggerDoc() map[string]string {
	return map_CertificateSigningRequest
}

var map_CertificateSigningRequestCondition = map[string]string{
	"type":           "request approval state, currently Approved or Denied.",
	"reason":         "brief reason for the request state",
	"message":        "human readable message with details about the request state",
	"lastUpdateTime": "timestamp for the last update to this condition",
}

func (CertificateSigningRequestCondition) SwaggerDoc() map[string]string {
	return map_CertificateSigningRequestCondition
}

var map_CertificateSigningRequestSpec = map[string]string{
	"":         "This information is immutable after the request is created. Only the Request and ExtraInfo fields can be set on creation, other fields are derived by Kubernetes and cannot be modified by users.",
	"request":  "Base64-encoded PKCS#10 CSR data",
	"username": "Information about the requesting user (if relevant) See user.Info interface for details",
}

func (CertificateSigningRequestSpec) SwaggerDoc() map[string]string {
	return map_CertificateSigningRequestSpec
}

var map_CertificateSigningRequestStatus = map[string]string{
	"conditions":  "Conditions applied to the request, such as approval or denial.",
	"certificate": "If request was approved, the controller will place the issued certificate here.",
}

func (CertificateSigningRequestStatus) SwaggerDoc() map[string]string {
	return map_CertificateSigningRequestStatus
}

// AUTO-GENERATED FUNCTIONS END HERE
