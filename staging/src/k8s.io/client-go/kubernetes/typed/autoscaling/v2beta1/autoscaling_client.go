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

package v2beta1

import (
	v2beta1 "k8s.io/api/autoscaling/v2beta1"
	serializer "k8s.io/apimachinery/pkg/runtime/serializer"
	"k8s.io/client-go/kubernetes/scheme"
	rest "k8s.io/client-go/rest"
)

type AutoscalingV2beta1Interface interface {
	RESTClient() rest.Interface
	HorizontalPodAutoscalersGetter
	VerticalPodAutoscalersGetter
}

// AutoscalingV2beta1Client is used to interact with features provided by the autoscaling group.
type AutoscalingV2beta1Client struct {
	restClient rest.Interface
}

func (c *AutoscalingV2beta1Client) HorizontalPodAutoscalers(namespace string) HorizontalPodAutoscalerInterface {
	return newHorizontalPodAutoscalers(c, namespace)
}

func (c *AutoscalingV2beta1Client) VerticalPodAutoscalers(namespace string) VerticalPodAutoscalerInterface {
	return newVerticalPodAutoscalers(c, namespace)
}

// NewForConfig creates a new AutoscalingV2beta1Client for the given config.
func NewForConfig(c *rest.Config) (*AutoscalingV2beta1Client, error) {
	config := *c
	if err := setConfigDefaults(&config); err != nil {
		return nil, err
	}
	client, err := rest.RESTClientFor(&config)
	if err != nil {
		return nil, err
	}
	return &AutoscalingV2beta1Client{client}, nil
}

// NewForConfigOrDie creates a new AutoscalingV2beta1Client for the given config and
// panics if there is an error in the config.
func NewForConfigOrDie(c *rest.Config) *AutoscalingV2beta1Client {
	client, err := NewForConfig(c)
	if err != nil {
		panic(err)
	}
	return client
}

// New creates a new AutoscalingV2beta1Client for the given RESTClient.
func New(c rest.Interface) *AutoscalingV2beta1Client {
	return &AutoscalingV2beta1Client{c}
}

func setConfigDefaults(config *rest.Config) error {
	gv := v2beta1.SchemeGroupVersion
	config.GroupVersion = &gv
	config.APIPath = "/apis"
	config.NegotiatedSerializer = serializer.DirectCodecFactory{CodecFactory: scheme.Codecs}

	if config.UserAgent == "" {
		config.UserAgent = rest.DefaultKubernetesUserAgent()
	}

	return nil
}

// RESTClient returns a RESTClient that is used to communicate
// with API server by this client implementation.
func (c *AutoscalingV2beta1Client) RESTClient() rest.Interface {
	if c == nil {
		return nil
	}
	return c.restClient
}
