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
// Those methods can be generated by using hack/update-generated-swagger-docs.sh

// AUTO-GENERATED FUNCTIONS START HERE. DO NOT EDIT.
var map_AuditSink = map[string]string{
	"":     "AuditSink represents a cluster level audit sink",
	"spec": "Spec defines the audit configuration spec",
}

func (AuditSink) SwaggerDoc() map[string]string {
	return map_AuditSink
}

var map_AuditSinkList = map[string]string{
	"":      "AuditSinkList is a list of AuditSink items.",
	"items": "List of audit configurations.",
}

func (AuditSinkList) SwaggerDoc() map[string]string {
	return map_AuditSinkList
}

var map_AuditSinkSpec = map[string]string{
	"":        "AuditSinkSpec holds the spec for the audit sink",
	"policy":  "Policy defines the policy for selecting which events should be sent to the webhook required",
	"webhook": "Webhook to send events required",
}

func (AuditSinkSpec) SwaggerDoc() map[string]string {
	return map_AuditSinkSpec
}

var map_Policy = map[string]string{
	"":       "Policy defines the configuration of how audit events are logged",
	"level":  "The Level that all requests are recorded at. available options: None, Metadata, Request, RequestResponse required",
	"stages": "Stages is a list of stages for which events are created.",
}

func (Policy) SwaggerDoc() map[string]string {
	return map_Policy
}

var map_ServiceReference = map[string]string{
	"":          "ServiceReference holds a reference to Service.legacy.k8s.io",
	"namespace": "`namespace` is the namespace of the service. Required",
	"name":      "`name` is the name of the service. Required",
	"path":      "`path` is an optional URL path which will be sent in any request to this service.",
	"port":      "If specified, the port on the service that hosting webhook. Default to 443 for backward compatibility. `port` should be a valid port number (1-65535, inclusive).",
}

func (ServiceReference) SwaggerDoc() map[string]string {
	return map_ServiceReference
}

var map_Webhook = map[string]string{
	"":             "Webhook holds the configuration of the webhook",
	"throttle":     "Throttle holds the options for throttling the webhook",
	"clientConfig": "ClientConfig holds the connection parameters for the webhook required",
}

func (Webhook) SwaggerDoc() map[string]string {
	return map_Webhook
}

var map_WebhookClientConfig = map[string]string{
	"":         "WebhookClientConfig contains the information to make a connection with the webhook",
	"url":      "`url` gives the location of the webhook, in standard URL form (`scheme://host:port/path`). Exactly one of `url` or `service` must be specified.\n\nThe `host` should not refer to a service running in the cluster; use the `service` field instead. The host might be resolved via external DNS in some apiservers (e.g., `kube-apiserver` cannot resolve in-cluster DNS as that would be a layering violation). `host` may also be an IP address.\n\nPlease note that using `localhost` or `127.0.0.1` as a `host` is risky unless you take great care to run this webhook on all hosts which run an apiserver which might need to make calls to this webhook. Such installs are likely to be non-portable, i.e., not easy to turn up in a new cluster.\n\nThe scheme must be \"https\"; the URL must begin with \"https://\".\n\nA path is optional, and if present may be any string permissible in a URL. You may use the path to pass an arbitrary string to the webhook, for example, a cluster identifier.\n\nAttempting to use a user or basic auth e.g. \"user:password@\" is not allowed. Fragments (\"#...\") and query parameters (\"?...\") are not allowed, either.",
	"service":  "`service` is a reference to the service for this webhook. Either `service` or `url` must be specified.\n\nIf the webhook is running within the cluster, then you should use `service`.",
	"caBundle": "`caBundle` is a PEM encoded CA bundle which will be used to validate the webhook's server certificate. If unspecified, system trust roots on the apiserver are used.",
}

func (WebhookClientConfig) SwaggerDoc() map[string]string {
	return map_WebhookClientConfig
}

var map_WebhookThrottleConfig = map[string]string{
	"":      "WebhookThrottleConfig holds the configuration for throttling events",
	"qps":   "ThrottleQPS maximum number of batches per second default 10 QPS",
	"burst": "ThrottleBurst is the maximum number of events sent at the same moment default 15 QPS",
}

func (WebhookThrottleConfig) SwaggerDoc() map[string]string {
	return map_WebhookThrottleConfig
}

// AUTO-GENERATED FUNCTIONS END HERE
