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

package v1

import (
	"context"
	"time"

	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	scheme "k8s.io/client-go/deprecated/scheme"
	rest "k8s.io/client-go/rest"
)

// EndpointsGetter has a method to return a EndpointsInterface.
// A group's client should implement this interface.
type EndpointsGetter interface {
	Endpoints(namespace string) EndpointsInterface
}

// EndpointsInterface has methods to work with Endpoints resources.
type EndpointsInterface interface {
	Create(*v1.Endpoints) (*v1.Endpoints, error)
	Update(*v1.Endpoints) (*v1.Endpoints, error)
	Delete(name string, options *metav1.DeleteOptions) error
	DeleteCollection(options *metav1.DeleteOptions, listOptions metav1.ListOptions) error
	Get(name string, options metav1.GetOptions) (*v1.Endpoints, error)
	List(opts metav1.ListOptions) (*v1.EndpointsList, error)
	Watch(opts metav1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1.Endpoints, err error)
	EndpointsExpansion
}

// endpoints implements EndpointsInterface
type endpoints struct {
	client rest.Interface
	ns     string
}

// newEndpoints returns a Endpoints
func newEndpoints(c *CoreV1Client, namespace string) *endpoints {
	return &endpoints{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the endpoints, and returns the corresponding endpoints object, and an error if there is any.
func (c *endpoints) Get(name string, options metav1.GetOptions) (result *v1.Endpoints, err error) {
	result = &v1.Endpoints{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("endpoints").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(context.TODO()).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of Endpoints that match those selectors.
func (c *endpoints) List(opts metav1.ListOptions) (result *v1.EndpointsList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1.EndpointsList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("endpoints").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(context.TODO()).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested endpoints.
func (c *endpoints) Watch(opts metav1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("endpoints").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(context.TODO())
}

// Create takes the representation of a endpoints and creates it.  Returns the server's representation of the endpoints, and an error, if there is any.
func (c *endpoints) Create(endpoints *v1.Endpoints) (result *v1.Endpoints, err error) {
	result = &v1.Endpoints{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("endpoints").
		Body(endpoints).
		Do(context.TODO()).
		Into(result)
	return
}

// Update takes the representation of a endpoints and updates it. Returns the server's representation of the endpoints, and an error, if there is any.
func (c *endpoints) Update(endpoints *v1.Endpoints) (result *v1.Endpoints, err error) {
	result = &v1.Endpoints{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("endpoints").
		Name(endpoints.Name).
		Body(endpoints).
		Do(context.TODO()).
		Into(result)
	return
}

// Delete takes name of the endpoints and deletes it. Returns an error if one occurs.
func (c *endpoints) Delete(name string, options *metav1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("endpoints").
		Name(name).
		Body(options).
		Do(context.TODO()).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *endpoints) DeleteCollection(options *metav1.DeleteOptions, listOptions metav1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("endpoints").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do(context.TODO()).
		Error()
}

// Patch applies the patch and returns the patched endpoints.
func (c *endpoints) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1.Endpoints, err error) {
	result = &v1.Endpoints{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("endpoints").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do(context.TODO()).
		Into(result)
	return
}
