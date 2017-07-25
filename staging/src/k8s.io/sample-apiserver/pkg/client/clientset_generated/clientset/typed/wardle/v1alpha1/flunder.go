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

package v1alpha1

import (
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
	v1alpha1 "k8s.io/sample-apiserver/pkg/apis/wardle/v1alpha1"
	scheme "k8s.io/sample-apiserver/pkg/client/clientset_generated/clientset/scheme"
)

// FlundersGetter has a method to return a FlunderInterface.
// A group's client should implement this interface.
type FlundersGetter interface {
	Flunders(namespace string) FlunderInterface
}

// FlunderInterface has methods to work with Flunder resources.
type FlunderInterface interface {
	Create(*v1alpha1.Flunder) (*v1alpha1.Flunder, error)
	Update(*v1alpha1.Flunder) (*v1alpha1.Flunder, error)
	UpdateStatus(*v1alpha1.Flunder) (*v1alpha1.Flunder, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.Flunder, error)
	List(opts v1.ListOptions) (*v1alpha1.FlunderList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.Flunder, err error)
	FlunderExpansion
}

// flunders implements FlunderInterface
type flunders struct {
	client rest.Interface
	ns     string
}

// newFlunders returns a Flunders
func newFlunders(c *WardleV1alpha1Client, namespace string) *flunders {
	return &flunders{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the flunder, and returns the corresponding flunder object, and an error if there is any.
func (c *flunders) Get(name string, options v1.GetOptions) (result *v1alpha1.Flunder, err error) {
	result = &v1alpha1.Flunder{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("flunders").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of Flunders that match those selectors.
func (c *flunders) List(opts v1.ListOptions) (result *v1alpha1.FlunderList, err error) {
	result = &v1alpha1.FlunderList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("flunders").
		VersionedParams(&opts, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested flunders.
func (c *flunders) Watch(opts v1.ListOptions) (watch.Interface, error) {
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("flunders").
		VersionedParams(&opts, scheme.ParameterCodec).
		Watch()
}

// Create takes the representation of a flunder and creates it.  Returns the server's representation of the flunder, and an error, if there is any.
func (c *flunders) Create(flunder *v1alpha1.Flunder) (result *v1alpha1.Flunder, err error) {
	result = &v1alpha1.Flunder{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("flunders").
		Body(flunder).
		Do().
		Into(result)
	return
}

// Update takes the representation of a flunder and updates it. Returns the server's representation of the flunder, and an error, if there is any.
func (c *flunders) Update(flunder *v1alpha1.Flunder) (result *v1alpha1.Flunder, err error) {
	result = &v1alpha1.Flunder{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("flunders").
		Name(flunder.Name).
		Body(flunder).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *flunders) UpdateStatus(flunder *v1alpha1.Flunder) (result *v1alpha1.Flunder, err error) {
	result = &v1alpha1.Flunder{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("flunders").
		Name(flunder.Name).
		SubResource("status").
		Body(flunder).
		Do().
		Into(result)
	return
}

// Delete takes name of the flunder and deletes it. Returns an error if one occurs.
func (c *flunders) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("flunders").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *flunders) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("flunders").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched flunder.
func (c *flunders) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.Flunder, err error) {
	result = &v1alpha1.Flunder{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("flunders").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
