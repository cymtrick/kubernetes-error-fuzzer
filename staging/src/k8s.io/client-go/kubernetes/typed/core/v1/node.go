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
	scheme "k8s.io/client-go/kubernetes/scheme"
	rest "k8s.io/client-go/rest"
)

// NodesGetter has a method to return a NodeInterface.
// A group's client should implement this interface.
type NodesGetter interface {
	Nodes() NodeInterface
}

// NodeInterface has methods to work with Node resources.
type NodeInterface interface {
	Create(*v1.Node) (*v1.Node, error)
	Update(*v1.Node) (*v1.Node, error)
	UpdateStatus(*v1.Node) (*v1.Node, error)
	Delete(name string, options *metav1.DeleteOptions) error
	DeleteCollection(options *metav1.DeleteOptions, listOptions metav1.ListOptions) error
	Get(name string, options metav1.GetOptions) (*v1.Node, error)
	List(opts metav1.ListOptions) (*v1.NodeList, error)
	Watch(opts metav1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1.Node, err error)
	NodeExpansion
}

// nodes implements NodeInterface
type nodes struct {
	client rest.Interface
}

// newNodes returns a Nodes
func newNodes(c *CoreV1Client) *nodes {
	return &nodes{
		client: c.RESTClient(),
	}
}

// Get takes name of the node, and returns the corresponding node object, and an error if there is any.
func (c *nodes) Get(name string, options metav1.GetOptions) (result *v1.Node, err error) {
	result = &v1.Node{}
	err = c.client.Get().
		Resource("nodes").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(context.TODO()).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of Nodes that match those selectors.
func (c *nodes) List(opts metav1.ListOptions) (result *v1.NodeList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1.NodeList{}
	err = c.client.Get().
		Resource("nodes").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(context.TODO()).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested nodes.
func (c *nodes) Watch(opts metav1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Resource("nodes").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(context.TODO())
}

// Create takes the representation of a node and creates it.  Returns the server's representation of the node, and an error, if there is any.
func (c *nodes) Create(node *v1.Node) (result *v1.Node, err error) {
	result = &v1.Node{}
	err = c.client.Post().
		Resource("nodes").
		Body(node).
		Do(context.TODO()).
		Into(result)
	return
}

// Update takes the representation of a node and updates it. Returns the server's representation of the node, and an error, if there is any.
func (c *nodes) Update(node *v1.Node) (result *v1.Node, err error) {
	result = &v1.Node{}
	err = c.client.Put().
		Resource("nodes").
		Name(node.Name).
		Body(node).
		Do(context.TODO()).
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *nodes) UpdateStatus(node *v1.Node) (result *v1.Node, err error) {
	result = &v1.Node{}
	err = c.client.Put().
		Resource("nodes").
		Name(node.Name).
		SubResource("status").
		Body(node).
		Do(context.TODO()).
		Into(result)
	return
}

// Delete takes name of the node and deletes it. Returns an error if one occurs.
func (c *nodes) Delete(name string, options *metav1.DeleteOptions) error {
	return c.client.Delete().
		Resource("nodes").
		Name(name).
		Body(options).
		Do(context.TODO()).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *nodes) DeleteCollection(options *metav1.DeleteOptions, listOptions metav1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Resource("nodes").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do(context.TODO()).
		Error()
}

// Patch applies the patch and returns the patched node.
func (c *nodes) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1.Node, err error) {
	result = &v1.Node{}
	err = c.client.Patch(pt).
		Resource("nodes").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do(context.TODO()).
		Into(result)
	return
}
