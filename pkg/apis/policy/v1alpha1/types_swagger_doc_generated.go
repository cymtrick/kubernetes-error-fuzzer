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
var map_Eviction = map[string]string{
	"":              "Eviction evicts a pod from its node subject to certain policies and safety constraints. This is a subresource of Pod.  A request to cause such an eviction is created by POSTing to .../pods/foo/evictions.",
	"metadata":      "ObjectMeta describes the pod that is being evicted.",
	"deleteOptions": "DeleteOptions may be provided",
}

func (Eviction) SwaggerDoc() map[string]string {
	return map_Eviction
}

var map_PodDisruptionBudget = map[string]string{
	"":       "PodDisruptionBudget is an object to define the max disruption that can be caused to a collection of pods",
	"spec":   "Specification of the desired behavior of the PodDisruptionBudget.",
	"status": "Most recently observed status of the PodDisruptionBudget.",
}

func (PodDisruptionBudget) SwaggerDoc() map[string]string {
	return map_PodDisruptionBudget
}

var map_PodDisruptionBudgetList = map[string]string{
	"": "PodDisruptionBudgetList is a collection of PodDisruptionBudgets.",
}

func (PodDisruptionBudgetList) SwaggerDoc() map[string]string {
	return map_PodDisruptionBudgetList
}

var map_PodDisruptionBudgetSpec = map[string]string{
	"":             "PodDisruptionBudgetSpec is a description of a PodDisruptionBudget.",
	"minAvailable": "The minimum number of pods that must be available simultaneously.  This can be either an integer or a string specifying a percentage, e.g. \"28%\".",
	"selector":     "Label query over pods whose evictions are managed by the disruption budget.",
}

func (PodDisruptionBudgetSpec) SwaggerDoc() map[string]string {
	return map_PodDisruptionBudgetSpec
}

var map_PodDisruptionBudgetStatus = map[string]string{
	"":                  "PodDisruptionBudgetStatus represents information about the status of a PodDisruptionBudget. Status may trail the actual state of a system.",
	"disruptionAllowed": "Whether or not a disruption is currently allowed.",
	"currentHealthy":    "current number of healthy pods",
	"desiredHealthy":    "minimum desired number of healthy pods",
	"expectedPods":      "total number of pods counted by this disruption budget",
}

func (PodDisruptionBudgetStatus) SwaggerDoc() map[string]string {
	return map_PodDisruptionBudgetStatus
}

// AUTO-GENERATED FUNCTIONS END HERE
