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

package v1beta1

import (
	"time"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
	v1beta1 "k8s.io/metrics/pkg/apis/metrics/v1beta1"
	scheme "k8s.io/metrics/pkg/client/clientset/versioned/scheme"
)

// NodeMetricsesGetter has a method to return a NodeMetricsInterface.
// A group's client should implement this interface.
type NodeMetricsesGetter interface {
	NodeMetricses() NodeMetricsInterface
}

// NodeMetricsInterface has methods to work with NodeMetrics resources.
type NodeMetricsInterface interface {
	Get(name string, options v1.GetOptions) (*v1beta1.NodeMetrics, error)
	List(opts v1.ListOptions) (*v1beta1.NodeMetricsList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	NodeMetricsExpansion
}

// nodeMetricses implements NodeMetricsInterface
type nodeMetricses struct {
	client rest.Interface
}

// newNodeMetricses returns a NodeMetricses
func newNodeMetricses(c *MetricsV1beta1Client) *nodeMetricses {
	return &nodeMetricses{
		client: c.RESTClient(),
	}
}

// Get takes name of the nodeMetrics, and returns the corresponding nodeMetrics object, and an error if there is any.
func (c *nodeMetricses) Get(name string, options v1.GetOptions) (result *v1beta1.NodeMetrics, err error) {
	result = &v1beta1.NodeMetrics{}
	err = c.client.Get().
		Resource("nodes").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of NodeMetricses that match those selectors.
func (c *nodeMetricses) List(opts v1.ListOptions) (result *v1beta1.NodeMetricsList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1beta1.NodeMetricsList{}
	err = c.client.Get().
		Resource("nodes").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested nodeMetricses.
func (c *nodeMetricses) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Resource("nodes").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}
