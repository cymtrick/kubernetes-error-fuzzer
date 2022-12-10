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
var map_PriorityClass = map[string]string{
	"":                 "PriorityClass defines mapping from a priority class name to the priority integer value. The value can be any valid integer.",
	"metadata":         "Standard object's metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata",
	"value":            "value represents the integer value of this priority class. This is the actual priority that pods receive when they have the name of this class in their pod spec.",
	"globalDefault":    "globalDefault specifies whether this PriorityClass should be considered as the default priority for pods that do not have any priority class. Only one PriorityClass can be marked as `globalDefault`. However, if more than one PriorityClasses exists with their `globalDefault` field set to true, the smallest value of such global default PriorityClasses will be used as the default priority.",
	"description":      "description is an arbitrary string that usually provides guidelines on when this priority class should be used.",
	"preemptionPolicy": "preemptionPolicy is the Policy for preempting pods with lower priority. One of Never, PreemptLowerPriority. Defaults to PreemptLowerPriority if unset.",
}

func (PriorityClass) SwaggerDoc() map[string]string {
	return map_PriorityClass
}

var map_PriorityClassList = map[string]string{
	"":         "PriorityClassList is a collection of priority classes.",
	"metadata": "Standard list metadata More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata",
	"items":    "items is the list of PriorityClasses",
}

func (PriorityClassList) SwaggerDoc() map[string]string {
	return map_PriorityClassList
}

// AUTO-GENERATED FUNCTIONS END HERE
