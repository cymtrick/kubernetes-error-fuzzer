/*
Copyright 2018 The Kubernetes Authors.

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
	v1alpha1 "k8s.io/sample-controller/pkg/apis/samplecontroller/v1alpha1"
	scheme "k8s.io/sample-controller/pkg/client/clientset/versioned/scheme"
)

// FoosGetter has a method to return a FooInterface.
// A group's client should implement this interface.
type FoosGetter interface {
	Foos(namespace string) FooInterface
}

// FooInterface has methods to work with Foo resources.
type FooInterface interface {
	Create(*v1alpha1.Foo) (*v1alpha1.Foo, error)
	Update(*v1alpha1.Foo) (*v1alpha1.Foo, error)
	UpdateStatus(*v1alpha1.Foo) (*v1alpha1.Foo, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.Foo, error)
	List(opts v1.ListOptions) (*v1alpha1.FooList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.Foo, err error)
	FooExpansion
}

// foos implements FooInterface
type foos struct {
	client rest.Interface
	ns     string
}

// newFoos returns a Foos
func newFoos(c *SamplecontrollerV1alpha1Client, namespace string) *foos {
	return &foos{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the foo, and returns the corresponding foo object, and an error if there is any.
func (c *foos) Get(name string, options v1.GetOptions) (result *v1alpha1.Foo, err error) {
	result = &v1alpha1.Foo{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("foos").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of Foos that match those selectors.
func (c *foos) List(opts v1.ListOptions) (result *v1alpha1.FooList, err error) {
	result = &v1alpha1.FooList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("foos").
		VersionedParams(&opts, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested foos.
func (c *foos) Watch(opts v1.ListOptions) (watch.Interface, error) {
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("foos").
		VersionedParams(&opts, scheme.ParameterCodec).
		Watch()
}

// Create takes the representation of a foo and creates it.  Returns the server's representation of the foo, and an error, if there is any.
func (c *foos) Create(foo *v1alpha1.Foo) (result *v1alpha1.Foo, err error) {
	result = &v1alpha1.Foo{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("foos").
		Body(foo).
		Do().
		Into(result)
	return
}

// Update takes the representation of a foo and updates it. Returns the server's representation of the foo, and an error, if there is any.
func (c *foos) Update(foo *v1alpha1.Foo) (result *v1alpha1.Foo, err error) {
	result = &v1alpha1.Foo{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("foos").
		Name(foo.Name).
		Body(foo).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *foos) UpdateStatus(foo *v1alpha1.Foo) (result *v1alpha1.Foo, err error) {
	result = &v1alpha1.Foo{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("foos").
		Name(foo.Name).
		SubResource("status").
		Body(foo).
		Do().
		Into(result)
	return
}

// Delete takes name of the foo and deletes it. Returns an error if one occurs.
func (c *foos) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("foos").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *foos) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("foos").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched foo.
func (c *foos) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.Foo, err error) {
	result = &v1alpha1.Foo{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("foos").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
