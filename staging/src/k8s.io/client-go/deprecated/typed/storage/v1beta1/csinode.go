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
	"context"
	"time"

	v1beta1 "k8s.io/api/storage/v1beta1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	scheme "k8s.io/client-go/deprecated/scheme"
	rest "k8s.io/client-go/rest"
)

// CSINodesGetter has a method to return a CSINodeInterface.
// A group's client should implement this interface.
type CSINodesGetter interface {
	CSINodes() CSINodeInterface
}

// CSINodeInterface has methods to work with CSINode resources.
type CSINodeInterface interface {
	Create(*v1beta1.CSINode) (*v1beta1.CSINode, error)
	Update(*v1beta1.CSINode) (*v1beta1.CSINode, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1beta1.CSINode, error)
	List(opts v1.ListOptions) (*v1beta1.CSINodeList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1beta1.CSINode, err error)
	CSINodeExpansion
}

// cSINodes implements CSINodeInterface
type cSINodes struct {
	client rest.Interface
}

// newCSINodes returns a CSINodes
func newCSINodes(c *StorageV1beta1Client) *cSINodes {
	return &cSINodes{
		client: c.RESTClient(),
	}
}

// Get takes name of the cSINode, and returns the corresponding cSINode object, and an error if there is any.
func (c *cSINodes) Get(name string, options v1.GetOptions) (result *v1beta1.CSINode, err error) {
	result = &v1beta1.CSINode{}
	err = c.client.Get().
		Resource("csinodes").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(context.TODO()).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of CSINodes that match those selectors.
func (c *cSINodes) List(opts v1.ListOptions) (result *v1beta1.CSINodeList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1beta1.CSINodeList{}
	err = c.client.Get().
		Resource("csinodes").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(context.TODO()).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested cSINodes.
func (c *cSINodes) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Resource("csinodes").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(context.TODO())
}

// Create takes the representation of a cSINode and creates it.  Returns the server's representation of the cSINode, and an error, if there is any.
func (c *cSINodes) Create(cSINode *v1beta1.CSINode) (result *v1beta1.CSINode, err error) {
	result = &v1beta1.CSINode{}
	err = c.client.Post().
		Resource("csinodes").
		Body(cSINode).
		Do(context.TODO()).
		Into(result)
	return
}

// Update takes the representation of a cSINode and updates it. Returns the server's representation of the cSINode, and an error, if there is any.
func (c *cSINodes) Update(cSINode *v1beta1.CSINode) (result *v1beta1.CSINode, err error) {
	result = &v1beta1.CSINode{}
	err = c.client.Put().
		Resource("csinodes").
		Name(cSINode.Name).
		Body(cSINode).
		Do(context.TODO()).
		Into(result)
	return
}

// Delete takes name of the cSINode and deletes it. Returns an error if one occurs.
func (c *cSINodes) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Resource("csinodes").
		Name(name).
		Body(options).
		Do(context.TODO()).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *cSINodes) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Resource("csinodes").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do(context.TODO()).
		Error()
}

// Patch applies the patch and returns the patched cSINode.
func (c *cSINodes) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1beta1.CSINode, err error) {
	result = &v1beta1.CSINode{}
	err = c.client.Patch(pt).
		Resource("csinodes").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do(context.TODO()).
		Into(result)
	return
}
