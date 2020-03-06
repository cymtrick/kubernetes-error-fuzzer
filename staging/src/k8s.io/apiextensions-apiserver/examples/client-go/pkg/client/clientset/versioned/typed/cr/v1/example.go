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

	v1 "k8s.io/apiextensions-apiserver/examples/client-go/pkg/apis/cr/v1"
	scheme "k8s.io/apiextensions-apiserver/examples/client-go/pkg/client/clientset/versioned/scheme"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// ExamplesGetter has a method to return a ExampleInterface.
// A group's client should implement this interface.
type ExamplesGetter interface {
	Examples(namespace string) ExampleInterface
}

// ExampleInterface has methods to work with Example resources.
type ExampleInterface interface {
	Create(ctx context.Context, example *v1.Example, opts metav1.CreateOptions) (*v1.Example, error)
	Update(ctx context.Context, example *v1.Example, opts metav1.UpdateOptions) (*v1.Example, error)
	Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error
	Get(ctx context.Context, name string, opts metav1.GetOptions) (*v1.Example, error)
	List(ctx context.Context, opts metav1.ListOptions) (*v1.ExampleList, error)
	Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *v1.Example, err error)
	ExampleExpansion
}

// examples implements ExampleInterface
type examples struct {
	client rest.Interface
	ns     string
}

// newExamples returns a Examples
func newExamples(c *CrV1Client, namespace string) *examples {
	return &examples{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the example, and returns the corresponding example object, and an error if there is any.
func (c *examples) Get(ctx context.Context, name string, options metav1.GetOptions) (result *v1.Example, err error) {
	result = &v1.Example{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("examples").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of Examples that match those selectors.
func (c *examples) List(ctx context.Context, opts metav1.ListOptions) (result *v1.ExampleList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1.ExampleList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("examples").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested examples.
func (c *examples) Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("examples").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a example and creates it.  Returns the server's representation of the example, and an error, if there is any.
func (c *examples) Create(ctx context.Context, example *v1.Example, opts metav1.CreateOptions) (result *v1.Example, err error) {
	result = &v1.Example{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("examples").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(example).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a example and updates it. Returns the server's representation of the example, and an error, if there is any.
func (c *examples) Update(ctx context.Context, example *v1.Example, opts metav1.UpdateOptions) (result *v1.Example, err error) {
	result = &v1.Example{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("examples").
		Name(example.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(example).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the example and deletes it. Returns an error if one occurs.
func (c *examples) Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("examples").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *examples) DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("examples").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched example.
func (c *examples) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *v1.Example, err error) {
	result = &v1.Example{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("examples").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}
