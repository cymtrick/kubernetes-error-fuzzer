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

// Code generated by client-gen. DO NOT EDIT.

package v1alpha2

import (
	"net/http"

	v1alpha2 "k8s.io/api/resource/v1alpha2"
	"k8s.io/client-go/kubernetes/scheme"
	rest "k8s.io/client-go/rest"
)

type ResourceV1alpha2Interface interface {
	RESTClient() rest.Interface
	PodSchedulingContextsGetter
	ResourceClaimsGetter
	ResourceClaimTemplatesGetter
	ResourceClassesGetter
}

// ResourceV1alpha2Client is used to interact with features provided by the resource.k8s.io group.
type ResourceV1alpha2Client struct {
	restClient rest.Interface
}

func (c *ResourceV1alpha2Client) PodSchedulingContexts(namespace string) PodSchedulingContextInterface {
	return newPodSchedulingContexts(c, namespace)
}

func (c *ResourceV1alpha2Client) ResourceClaims(namespace string) ResourceClaimInterface {
	return newResourceClaims(c, namespace)
}

func (c *ResourceV1alpha2Client) ResourceClaimTemplates(namespace string) ResourceClaimTemplateInterface {
	return newResourceClaimTemplates(c, namespace)
}

func (c *ResourceV1alpha2Client) ResourceClasses() ResourceClassInterface {
	return newResourceClasses(c)
}

// NewForConfig creates a new ResourceV1alpha2Client for the given config.
// NewForConfig is equivalent to NewForConfigAndClient(c, httpClient),
// where httpClient was generated with rest.HTTPClientFor(c).
func NewForConfig(c *rest.Config) (*ResourceV1alpha2Client, error) {
	config := *c
	if err := setConfigDefaults(&config); err != nil {
		return nil, err
	}
	httpClient, err := rest.HTTPClientFor(&config)
	if err != nil {
		return nil, err
	}
	return NewForConfigAndClient(&config, httpClient)
}

// NewForConfigAndClient creates a new ResourceV1alpha2Client for the given config and http client.
// Note the http client provided takes precedence over the configured transport values.
func NewForConfigAndClient(c *rest.Config, h *http.Client) (*ResourceV1alpha2Client, error) {
	config := *c
	if err := setConfigDefaults(&config); err != nil {
		return nil, err
	}
	client, err := rest.RESTClientForConfigAndClient(&config, h)
	if err != nil {
		return nil, err
	}
	return &ResourceV1alpha2Client{client}, nil
}

// NewForConfigOrDie creates a new ResourceV1alpha2Client for the given config and
// panics if there is an error in the config.
func NewForConfigOrDie(c *rest.Config) *ResourceV1alpha2Client {
	client, err := NewForConfig(c)
	if err != nil {
		panic(err)
	}
	return client
}

// New creates a new ResourceV1alpha2Client for the given RESTClient.
func New(c rest.Interface) *ResourceV1alpha2Client {
	return &ResourceV1alpha2Client{c}
}

func setConfigDefaults(config *rest.Config) error {
	gv := v1alpha2.SchemeGroupVersion
	config.GroupVersion = &gv
	config.APIPath = "/apis"
	config.NegotiatedSerializer = scheme.Codecs.WithoutConversion()

	if config.UserAgent == "" {
		config.UserAgent = rest.DefaultKubernetesUserAgent()
	}

	return nil
}

// RESTClient returns a RESTClient that is used to communicate
// with API server by this client implementation.
func (c *ResourceV1alpha2Client) RESTClient() rest.Interface {
	if c == nil {
		return nil
	}
	return c.restClient
}
