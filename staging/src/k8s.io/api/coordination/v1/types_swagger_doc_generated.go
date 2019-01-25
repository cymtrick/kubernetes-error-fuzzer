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
var map_Lease = map[string]string{
	"":         "Lease defines a lease concept.",
	"metadata": "More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#metadata",
	"spec":     "Specification of the Lease. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#spec-and-status",
}

func (Lease) SwaggerDoc() map[string]string {
	return map_Lease
}

var map_LeaseList = map[string]string{
	"":         "LeaseList is a list of Lease objects.",
	"metadata": "Standard list metadata. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#metadata",
	"items":    "Items is a list of schema objects.",
}

func (LeaseList) SwaggerDoc() map[string]string {
	return map_LeaseList
}

var map_LeaseSpec = map[string]string{
	"":                     "LeaseSpec is a specification of a Lease.",
	"holderIdentity":       "holderIdentity contains the identity of the holder of a current lease.",
	"leaseDurationSeconds": "leaseDurationSeconds is a duration that candidates for a lease need to wait to force acquire it. This is measure against time of last observed RenewTime.",
	"acquireTime":          "acquireTime is a time when the current lease was acquired.",
	"renewTime":            "renewTime is a time when the current holder of a lease has last updated the lease.",
	"leaseTransitions":     "leaseTransitions is the number of transitions of a lease between holders.",
}

func (LeaseSpec) SwaggerDoc() map[string]string {
	return map_LeaseSpec
}

// AUTO-GENERATED FUNCTIONS END HERE
