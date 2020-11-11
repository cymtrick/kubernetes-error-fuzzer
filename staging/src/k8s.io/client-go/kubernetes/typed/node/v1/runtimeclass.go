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

	v1 "k8s.io/api/node/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	scheme "k8s.io/client-go/kubernetes/scheme"
	rest "k8s.io/client-go/rest"
)

// RuntimeClassesGetter has a method to return a RuntimeClassInterface.
// A group's client should implement this interface.
type RuntimeClassesGetter interface {
	RuntimeClasses() RuntimeClassInterface
}

// RuntimeClassInterface has methods to work with RuntimeClass resources.
type RuntimeClassInterface interface {
	Create(ctx context.Context, runtimeClass *v1.RuntimeClass, opts metav1.CreateOptions) (*v1.RuntimeClass, error)
	Update(ctx context.Context, runtimeClass *v1.RuntimeClass, opts metav1.UpdateOptions) (*v1.RuntimeClass, error)
	Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error
	Get(ctx context.Context, name string, opts metav1.GetOptions) (*v1.RuntimeClass, error)
	List(ctx context.Context, opts metav1.ListOptions) (*v1.RuntimeClassList, error)
	Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *v1.RuntimeClass, err error)
	RuntimeClassExpansion
}

// runtimeClasses implements RuntimeClassInterface
type runtimeClasses struct {
	client rest.Interface
}

// newRuntimeClasses returns a RuntimeClasses
func newRuntimeClasses(c *NodeV1Client) *runtimeClasses {
	return &runtimeClasses{
		client: c.RESTClient(),
	}
}

// Get takes name of the runtimeClass, and returns the corresponding runtimeClass object, and an error if there is any.
func (c *runtimeClasses) Get(ctx context.Context, name string, options metav1.GetOptions) (result *v1.RuntimeClass, err error) {
	result = &v1.RuntimeClass{}
	err = c.client.Get().
		Resource("runtimeclasses").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of RuntimeClasses that match those selectors.
func (c *runtimeClasses) List(ctx context.Context, opts metav1.ListOptions) (result *v1.RuntimeClassList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1.RuntimeClassList{}
	err = c.client.Get().
		Resource("runtimeclasses").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested runtimeClasses.
func (c *runtimeClasses) Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Resource("runtimeclasses").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a runtimeClass and creates it.  Returns the server's representation of the runtimeClass, and an error, if there is any.
func (c *runtimeClasses) Create(ctx context.Context, runtimeClass *v1.RuntimeClass, opts metav1.CreateOptions) (result *v1.RuntimeClass, err error) {
	result = &v1.RuntimeClass{}
	err = c.client.Post().
		Resource("runtimeclasses").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(runtimeClass).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a runtimeClass and updates it. Returns the server's representation of the runtimeClass, and an error, if there is any.
func (c *runtimeClasses) Update(ctx context.Context, runtimeClass *v1.RuntimeClass, opts metav1.UpdateOptions) (result *v1.RuntimeClass, err error) {
	result = &v1.RuntimeClass{}
	err = c.client.Put().
		Resource("runtimeclasses").
		Name(runtimeClass.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(runtimeClass).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the runtimeClass and deletes it. Returns an error if one occurs.
func (c *runtimeClasses) Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error {
	return c.client.Delete().
		Resource("runtimeclasses").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *runtimeClasses) DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Resource("runtimeclasses").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched runtimeClass.
func (c *runtimeClasses) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *v1.RuntimeClass, err error) {
	result = &v1.RuntimeClass{}
	err = c.client.Patch(pt).
		Resource("runtimeclasses").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}
