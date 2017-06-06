/*
Copyright 2017 The Kubernetes Authors.

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

package internalversion

import (
	rest "k8s.io/client-go/rest"
	"k8s.io/kubernetes/pkg/client/clientset_generated/internalclientset/scheme"
)

type AppsInterface interface {
	RESTClient() rest.Interface
	ControllerRevisionsGetter
	StatefulSetsGetter
}

// AppsClient is used to interact with features provided by the apps group.
type AppsClient struct {
	restClient rest.Interface
}

func (c *AppsClient) ControllerRevisions(namespace string) ControllerRevisionInterface {
	return newControllerRevisions(c, namespace)
}

func (c *AppsClient) StatefulSets(namespace string) StatefulSetInterface {
	return newStatefulSets(c, namespace)
}

// NewForConfig creates a new AppsClient for the given config.
func NewForConfig(c *rest.Config) (*AppsClient, error) {
	config := *c
	if err := setConfigDefaults(&config); err != nil {
		return nil, err
	}
	client, err := rest.RESTClientFor(&config)
	if err != nil {
		return nil, err
	}
	return &AppsClient{client}, nil
}

// NewForConfigOrDie creates a new AppsClient for the given config and
// panics if there is an error in the config.
func NewForConfigOrDie(c *rest.Config) *AppsClient {
	client, err := NewForConfig(c)
	if err != nil {
		panic(err)
	}
	return client
}

// New creates a new AppsClient for the given RESTClient.
func New(c rest.Interface) *AppsClient {
	return &AppsClient{c}
}

func setConfigDefaults(config *rest.Config) error {
	g, err := scheme.Registry.Group("apps")
	if err != nil {
		return err
	}

	config.APIPath = "/apis"
	if config.UserAgent == "" {
		config.UserAgent = rest.DefaultKubernetesUserAgent()
	}
	if config.GroupVersion == nil || config.GroupVersion.Group != g.GroupVersion.Group {
		gv := g.GroupVersion
		config.GroupVersion = &gv
	}
	config.NegotiatedSerializer = scheme.Codecs

	if config.QPS == 0 {
		config.QPS = 5
	}
	if config.Burst == 0 {
		config.Burst = 10
	}

	return nil
}

// RESTClient returns a RESTClient that is used to communicate
// with API server by this client implementation.
func (c *AppsClient) RESTClient() rest.Interface {
	if c == nil {
		return nil
	}
	return c.restClient
}
