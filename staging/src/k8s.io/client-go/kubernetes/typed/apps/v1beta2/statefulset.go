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

package v1beta2

import (
	"context"
	"time"

	v1beta2 "k8s.io/api/apps/v1beta2"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	scheme "k8s.io/client-go/kubernetes/scheme"
	rest "k8s.io/client-go/rest"
)

// StatefulSetsGetter has a method to return a StatefulSetInterface.
// A group's client should implement this interface.
type StatefulSetsGetter interface {
	StatefulSets(namespace string) StatefulSetInterface
}

// StatefulSetInterface has methods to work with StatefulSet resources.
type StatefulSetInterface interface {
	Create(ctx context.Context, statefulSet *v1beta2.StatefulSet, opts v1.CreateOptions) (*v1beta2.StatefulSet, error)
	Update(ctx context.Context, statefulSet *v1beta2.StatefulSet, opts v1.UpdateOptions) (*v1beta2.StatefulSet, error)
	UpdateStatus(ctx context.Context, statefulSet *v1beta2.StatefulSet, opts v1.UpdateOptions) (*v1beta2.StatefulSet, error)
	Delete(ctx context.Context, name string, opts *v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts *v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*v1beta2.StatefulSet, error)
	List(ctx context.Context, opts v1.ListOptions) (*v1beta2.StatefulSetList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1beta2.StatefulSet, err error)
	GetScale(ctx context.Context, statefulSetName string, options v1.GetOptions) (*v1beta2.Scale, error)
	UpdateScale(ctx context.Context, statefulSetName string, scale *v1beta2.Scale, opts v1.UpdateOptions) (*v1beta2.Scale, error)

	StatefulSetExpansion
}

// statefulSets implements StatefulSetInterface
type statefulSets struct {
	client rest.Interface
	ns     string
}

// newStatefulSets returns a StatefulSets
func newStatefulSets(c *AppsV1beta2Client, namespace string) *statefulSets {
	return &statefulSets{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the statefulSet, and returns the corresponding statefulSet object, and an error if there is any.
func (c *statefulSets) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1beta2.StatefulSet, err error) {
	result = &v1beta2.StatefulSet{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("statefulsets").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of StatefulSets that match those selectors.
func (c *statefulSets) List(ctx context.Context, opts v1.ListOptions) (result *v1beta2.StatefulSetList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1beta2.StatefulSetList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("statefulsets").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested statefulSets.
func (c *statefulSets) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("statefulsets").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a statefulSet and creates it.  Returns the server's representation of the statefulSet, and an error, if there is any.
func (c *statefulSets) Create(ctx context.Context, statefulSet *v1beta2.StatefulSet, opts v1.CreateOptions) (result *v1beta2.StatefulSet, err error) {
	result = &v1beta2.StatefulSet{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("statefulsets").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(statefulSet).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a statefulSet and updates it. Returns the server's representation of the statefulSet, and an error, if there is any.
func (c *statefulSets) Update(ctx context.Context, statefulSet *v1beta2.StatefulSet, opts v1.UpdateOptions) (result *v1beta2.StatefulSet, err error) {
	result = &v1beta2.StatefulSet{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("statefulsets").
		Name(statefulSet.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(statefulSet).
		Do(ctx).
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *statefulSets) UpdateStatus(ctx context.Context, statefulSet *v1beta2.StatefulSet, opts v1.UpdateOptions) (result *v1beta2.StatefulSet, err error) {
	result = &v1beta2.StatefulSet{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("statefulsets").
		Name(statefulSet.Name).
		SubResource("status").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(statefulSet).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the statefulSet and deletes it. Returns an error if one occurs.
func (c *statefulSets) Delete(ctx context.Context, name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("statefulsets").
		Name(name).
		Body(options).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *statefulSets) DeleteCollection(ctx context.Context, options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("statefulsets").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched statefulSet.
func (c *statefulSets) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1beta2.StatefulSet, err error) {
	result = &v1beta2.StatefulSet{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("statefulsets").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}

// GetScale takes name of the statefulSet, and returns the corresponding v1beta2.Scale object, and an error if there is any.
func (c *statefulSets) GetScale(ctx context.Context, statefulSetName string, options v1.GetOptions) (result *v1beta2.Scale, err error) {
	result = &v1beta2.Scale{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("statefulsets").
		Name(statefulSetName).
		SubResource("scale").
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// UpdateScale takes the top resource name and the representation of a scale and updates it. Returns the server's representation of the scale, and an error, if there is any.
func (c *statefulSets) UpdateScale(ctx context.Context, statefulSetName string, scale *v1beta2.Scale, opts v1.UpdateOptions) (result *v1beta2.Scale, err error) {
	result = &v1beta2.Scale{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("statefulsets").
		Name(statefulSetName).
		SubResource("scale").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(scale).
		Do(ctx).
		Into(result)
	return
}
