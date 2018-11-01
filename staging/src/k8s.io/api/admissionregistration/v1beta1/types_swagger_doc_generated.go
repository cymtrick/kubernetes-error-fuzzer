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

// AUTO-GENERATED FUNCTIONS START HERE. DO NOT EDIT.
var map_MutatingWebhookConfiguration = map[string]string{
	"":         "MutatingWebhookConfiguration describes the configuration of and admission webhook that accept or reject and may change the object.",
	"metadata": "Standard object metadata; More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#metadata.",
	"webhooks": "Webhooks is a list of webhooks and the affected resources and operations.",
}

func (MutatingWebhookConfiguration) SwaggerDoc() map[string]string {
	return map_MutatingWebhookConfiguration
}

var map_MutatingWebhookConfigurationList = map[string]string{
	"":         "MutatingWebhookConfigurationList is a list of MutatingWebhookConfiguration.",
	"metadata": "Standard list metadata. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds",
	"items":    "List of MutatingWebhookConfiguration.",
}

func (MutatingWebhookConfigurationList) SwaggerDoc() map[string]string {
	return map_MutatingWebhookConfigurationList
}

var map_Rule = map[string]string{
	"":            "Rule is a tuple of APIGroups, APIVersion, and Resources.It is recommended to make sure that all the tuple expansions are valid.",
	"apiGroups":   "APIGroups is the API groups the resources belong to. '*' is all groups. If '*' is present, the length of the slice must be one. Required.",
	"apiVersions": "APIVersions is the API versions the resources belong to. '*' is all versions. If '*' is present, the length of the slice must be one. Required.",
	"resources":   "Resources is a list of resources this rule applies to.\n\nFor example: 'pods' means pods. 'pods/log' means the log subresource of pods. '*' means all resources, but not subresources. 'pods/*' means all subresources of pods. '*/scale' means all scale subresources. '*/*' means all resources and their subresources.\n\nIf wildcard is present, the validation rule will ensure resources do not overlap with each other.\n\nDepending on the enclosing object, subresources might not be allowed. Required.",
}

func (Rule) SwaggerDoc() map[string]string {
	return map_Rule
}

var map_RuleWithOperations = map[string]string{
	"":           "RuleWithOperations is a tuple of Operations and Resources. It is recommended to make sure that all the tuple expansions are valid.",
	"operations": "Operations is the operations the admission hook cares about - CREATE, UPDATE, or * for all operations. If '*' is present, the length of the slice must be one. Required.",
}

func (RuleWithOperations) SwaggerDoc() map[string]string {
	return map_RuleWithOperations
}

var map_ServiceReference = map[string]string{
	"":          "ServiceReference holds a reference to Service.legacy.k8s.io",
	"namespace": "`namespace` is the namespace of the service. Required",
	"name":      "`name` is the name of the service. Required",
	"path":      "`path` is an optional URL path which will be sent in any request to this service.",
}

func (ServiceReference) SwaggerDoc() map[string]string {
	return map_ServiceReference
}

var map_ValidatingWebhookConfiguration = map[string]string{
	"":         "ValidatingWebhookConfiguration describes the configuration of and admission webhook that accept or reject and object without changing it.",
	"metadata": "Standard object metadata; More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#metadata.",
	"webhooks": "Webhooks is a list of webhooks and the affected resources and operations.",
}

func (ValidatingWebhookConfiguration) SwaggerDoc() map[string]string {
	return map_ValidatingWebhookConfiguration
}

var map_ValidatingWebhookConfigurationList = map[string]string{
	"":         "ValidatingWebhookConfigurationList is a list of ValidatingWebhookConfiguration.",
	"metadata": "Standard list metadata. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds",
	"items":    "List of ValidatingWebhookConfiguration.",
}

func (ValidatingWebhookConfigurationList) SwaggerDoc() map[string]string {
	return map_ValidatingWebhookConfigurationList
}

var map_Webhook = map[string]string{
	"":                  "Webhook describes an admission webhook and the resources and operations it applies to.",
	"name":              "The name of the admission webhook. Name should be fully qualified, e.g., imagepolicy.kubernetes.io, where \"imagepolicy\" is the name of the webhook, and kubernetes.io is the name of the organization. Required.",
	"clientConfig":      "ClientConfig defines how to communicate with the hook. Required",
	"rules":             "Rules describes what operations on what resources/subresources the webhook cares about. The webhook cares about an operation if it matches _any_ Rule. However, in order to prevent ValidatingAdmissionWebhooks and MutatingAdmissionWebhooks from putting the cluster in a state which cannot be recovered from without completely disabling the plugin, ValidatingAdmissionWebhooks and MutatingAdmissionWebhooks are never called on admission requests for ValidatingWebhookConfiguration and MutatingWebhookConfiguration objects.",
	"failurePolicy":     "FailurePolicy defines how unrecognized errors from the admission endpoint are handled - allowed values are Ignore or Fail. Defaults to Ignore.",
	"namespaceSelector": "NamespaceSelector decides whether to run the webhook on an object based on whether the namespace for that object matches the selector. If the object itself is a namespace, the matching is performed on object.metadata.labels. If the object is another cluster scoped resource, it never skips the webhook.\n\nFor example, to run the webhook on any objects whose namespace is not associated with \"runlevel\" of \"0\" or \"1\";  you will set the selector as follows: \"namespaceSelector\": {\n  \"matchExpressions\": [\n    {\n      \"key\": \"runlevel\",\n      \"operator\": \"NotIn\",\n      \"values\": [\n        \"0\",\n        \"1\"\n      ]\n    }\n  ]\n}\n\nIf instead you want to only run the webhook on any objects whose namespace is associated with the \"environment\" of \"prod\" or \"staging\"; you will set the selector as follows: \"namespaceSelector\": {\n  \"matchExpressions\": [\n    {\n      \"key\": \"environment\",\n      \"operator\": \"In\",\n      \"values\": [\n        \"prod\",\n        \"staging\"\n      ]\n    }\n  ]\n}\n\nSee https://kubernetes.io/docs/concepts/overview/working-with-objects/labels/ for more examples of label selectors.\n\nDefault to the empty LabelSelector, which matches everything.",
	"sideEffects":       "SideEffects states whether this webhookk has side effects. Acceptable values are: Unknown, None, Some, NoneOnDryRun Webhooks with side effects MUST implement a reconciliation system, since a request may be rejected by a future step in the admission change and the side effects therefore need to be undone. Requests with the dryRun attribute will be auto-rejected if they match a webhook with sideEffects == Unknown or Some. Defaults to Unknown.",
}

func (Webhook) SwaggerDoc() map[string]string {
	return map_Webhook
}

var map_WebhookClientConfig = map[string]string{
	"":         "WebhookClientConfig contains the information to make a TLS connection with the webhook",
	"url":      "`url` gives the location of the webhook, in standard URL form (`scheme://host:port/path`). Exactly one of `url` or `service` must be specified.\n\nThe `host` should not refer to a service running in the cluster; use the `service` field instead. The host might be resolved via external DNS in some apiservers (e.g., `kube-apiserver` cannot resolve in-cluster DNS as that would be a layering violation). `host` may also be an IP address.\n\nPlease note that using `localhost` or `127.0.0.1` as a `host` is risky unless you take great care to run this webhook on all hosts which run an apiserver which might need to make calls to this webhook. Such installs are likely to be non-portable, i.e., not easy to turn up in a new cluster.\n\nThe scheme must be \"https\"; the URL must begin with \"https://\".\n\nA path is optional, and if present may be any string permissible in a URL. You may use the path to pass an arbitrary string to the webhook, for example, a cluster identifier.\n\nAttempting to use a user or basic auth e.g. \"user:password@\" is not allowed. Fragments (\"#...\") and query parameters (\"?...\") are not allowed, either.",
	"service":  "`service` is a reference to the service for this webhook. Either `service` or `url` must be specified.\n\nIf the webhook is running within the cluster, then you should use `service`.\n\nPort 443 will be used if it is open, otherwise it is an error.",
	"caBundle": "`caBundle` is a PEM encoded CA bundle which will be used to validate the webhook's server certificate. If unspecified, system trust roots on the apiserver are used.",
}

func (WebhookClientConfig) SwaggerDoc() map[string]string {
	return map_WebhookClientConfig
}

// AUTO-GENERATED FUNCTIONS END HERE
