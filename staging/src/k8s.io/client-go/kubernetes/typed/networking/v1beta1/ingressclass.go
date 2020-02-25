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

	v1beta1 "k8s.io/api/networking/v1beta1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	scheme "k8s.io/client-go/kubernetes/scheme"
	rest "k8s.io/client-go/rest"
)

// IngressClassesGetter has a method to return a IngressClassInterface.
// A group's client should implement this interface.
type IngressClassesGetter interface {
	IngressClasses() IngressClassInterface
}

// IngressClassInterface has methods to work with IngressClass resources.
type IngressClassInterface interface {
	Create(ctx context.Context, ingressClass *v1beta1.IngressClass, opts v1.CreateOptions) (*v1beta1.IngressClass, error)
	Update(ctx context.Context, ingressClass *v1beta1.IngressClass, opts v1.UpdateOptions) (*v1beta1.IngressClass, error)
	Delete(ctx context.Context, name string, opts *v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts *v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*v1beta1.IngressClass, error)
	List(ctx context.Context, opts v1.ListOptions) (*v1beta1.IngressClassList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1beta1.IngressClass, err error)
	IngressClassExpansion
}

// ingressClasses implements IngressClassInterface
type ingressClasses struct {
	client rest.Interface
}

// newIngressClasses returns a IngressClasses
func newIngressClasses(c *NetworkingV1beta1Client) *ingressClasses {
	return &ingressClasses{
		client: c.RESTClient(),
	}
}

// Get takes name of the ingressClass, and returns the corresponding ingressClass object, and an error if there is any.
func (c *ingressClasses) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1beta1.IngressClass, err error) {
	result = &v1beta1.IngressClass{}
	err = c.client.Get().
		Resource("ingressclasses").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of IngressClasses that match those selectors.
func (c *ingressClasses) List(ctx context.Context, opts v1.ListOptions) (result *v1beta1.IngressClassList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1beta1.IngressClassList{}
	err = c.client.Get().
		Resource("ingressclasses").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested ingressClasses.
func (c *ingressClasses) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Resource("ingressclasses").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a ingressClass and creates it.  Returns the server's representation of the ingressClass, and an error, if there is any.
func (c *ingressClasses) Create(ctx context.Context, ingressClass *v1beta1.IngressClass, opts v1.CreateOptions) (result *v1beta1.IngressClass, err error) {
	result = &v1beta1.IngressClass{}
	err = c.client.Post().
		Resource("ingressclasses").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(ingressClass).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a ingressClass and updates it. Returns the server's representation of the ingressClass, and an error, if there is any.
func (c *ingressClasses) Update(ctx context.Context, ingressClass *v1beta1.IngressClass, opts v1.UpdateOptions) (result *v1beta1.IngressClass, err error) {
	result = &v1beta1.IngressClass{}
	err = c.client.Put().
		Resource("ingressclasses").
		Name(ingressClass.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(ingressClass).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the ingressClass and deletes it. Returns an error if one occurs.
func (c *ingressClasses) Delete(ctx context.Context, name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Resource("ingressclasses").
		Name(name).
		Body(options).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *ingressClasses) DeleteCollection(ctx context.Context, options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Resource("ingressclasses").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched ingressClass.
func (c *ingressClasses) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1beta1.IngressClass, err error) {
	result = &v1beta1.IngressClass{}
	err = c.client.Patch(pt).
		Resource("ingressclasses").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}
