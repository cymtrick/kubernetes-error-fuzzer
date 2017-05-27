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

// AUTO-GENERATED FUNCTIONS START HERE
var map_NetworkPolicy = map[string]string{
	"":         "NetworkPolicy describes what network traffic is allowed for a set of Pods",
	"metadata": "Standard object's metadata. More info: http://releases.k8s.io/HEAD/docs/devel/api-conventions.md#metadata",
	"spec":     "Specification of the desired behavior for this NetworkPolicy.",
}

func (NetworkPolicy) SwaggerDoc() map[string]string {
	return map_NetworkPolicy
}

var map_NetworkPolicyIngressRule = map[string]string{
	"":      "NetworkPolicyIngressRule describes a particular set of traffic that is allowed to the pods matched by a NetworkPolicySpec's podSelector. The traffic must match both ports and from.",
	"ports": "List of ports which should be made accessible on the pods selected for this rule. Each item in this list is combined using a logical OR. If this field is empty or missing, this rule matches all ports (traffic not restricted by port). If this field is present and contains at least one item, then this rule allows traffic only if the traffic matches at least one port in the list.",
	"from":  "List of sources which should be able to access the pods selected for this rule. Items in this list are combined using a logical OR operation. If this field is empty or missing, this rule matches all sources (traffic not restricted by source). If this field is present and contains at least on item, this rule allows traffic only if the traffic matches at least one item in the from list.",
}

func (NetworkPolicyIngressRule) SwaggerDoc() map[string]string {
	return map_NetworkPolicyIngressRule
}

var map_NetworkPolicyList = map[string]string{
	"":         "NetworkPolicyList is a list of NetworkPolicy objects.",
	"metadata": "Standard list metadata. More info: http://releases.k8s.io/HEAD/docs/devel/api-conventions.md#metadata",
	"items":    "Items is a list of schema objects.",
}

func (NetworkPolicyList) SwaggerDoc() map[string]string {
	return map_NetworkPolicyList
}

var map_NetworkPolicyPeer = map[string]string{
	"":                  "NetworkPolicyPeer describes a peer to allow traffic from. Exactly one of its fields must be specified.",
	"podSelector":       "This is a label selector which selects Pods in this namespace. This field follows standard label selector semantics. If present but empty, this selector selects all pods in this namespace.",
	"namespaceSelector": "Selects Namespaces using cluster scoped-labels. This matches all pods in all namespaces selected by this label selector. This field follows standard label selector semantics. If present but empty, this selector selects all namespaces.",
}

func (NetworkPolicyPeer) SwaggerDoc() map[string]string {
	return map_NetworkPolicyPeer
}

var map_NetworkPolicyPort = map[string]string{
	"":         "NetworkPolicyPort describes a port to allow traffic on",
	"protocol": "The protocol (TCP or UDP) which traffic must match. If not specified, this field defaults to TCP.",
	"port":     "The port on the given protocol. This can either be a numerical or named port on a pod. If this field is not provided, this matches all port names and numbers.",
}

func (NetworkPolicyPort) SwaggerDoc() map[string]string {
	return map_NetworkPolicyPort
}

var map_NetworkPolicySpec = map[string]string{
	"":            "NetworkPolicySpec provides the specification of a NetworkPolicy",
	"podSelector": "Selects the pods to which this NetworkPolicy object applies. The array of ingress rules is applied to any pods selected by this field. Multiple network policies can select the same set of pods. In this case, the ingress rules for each are combined additively. This field is NOT optional and follows standard label selector semantics. An empty podSelector matches all pods in this namespace.",
	"ingress":     "List of ingress rules to be applied to the selected pods. Traffic is allowed to a pod if there are no NetworkPolicies selecting the pod (and cluster policy otherwise allows the traffic), OR if the traffic source is the pod's local node, OR if the traffic matches at least one ingress rule across all of the NetworkPolicy objects whose podSelector matches the pod. If this field is empty then this NetworkPolicy does not allow any traffic (and serves solely to ensure that the pods it selects are isolated by default)",
}

func (NetworkPolicySpec) SwaggerDoc() map[string]string {
	return map_NetworkPolicySpec
}

// AUTO-GENERATED FUNCTIONS END HERE
