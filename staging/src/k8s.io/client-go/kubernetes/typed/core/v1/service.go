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
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	scheme "k8s.io/client-go/kubernetes/scheme"
	rest "k8s.io/client-go/rest"
)

// ServicesGetter has a method to return a ServiceInterface.
// A group's client should implement this interface.
type ServicesGetter interface {
	Services(namespace string) ServiceInterface
}

// ServiceInterface has methods to work with Service resources.
type ServiceInterface interface {
	Create(*v1.Service) (*v1.Service, error)
	Update(*v1.Service) (*v1.Service, error)
	UpdateStatus(*v1.Service) (*v1.Service, error)
	Delete(name string, options *metav1.DeleteOptions) error
	Get(name string, options metav1.GetOptions) (*v1.Service, error)
	List(opts metav1.ListOptions) (*v1.ServiceList, error)
	Watch(opts metav1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1.Service, err error)
	ServiceExpansion
}

// services implements ServiceInterface
type services struct {
	client rest.Interface
	ns     string
}

// newServices returns a Services
func newServices(c *CoreV1Client, namespace string) *services {
	return &services{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the service, and returns the corresponding service object, and an error if there is any.
func (c *services) Get(name string, options metav1.GetOptions) (result *v1.Service, err error) {
	result = &v1.Service{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("services").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of Services that match those selectors.
func (c *services) List(opts metav1.ListOptions) (result *v1.ServiceList, err error) {
	result = &v1.ServiceList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("services").
		VersionedParams(&opts, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested services.
func (c *services) Watch(opts metav1.ListOptions) (watch.Interface, error) {
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("services").
		VersionedParams(&opts, scheme.ParameterCodec).
		Watch()
}

// Create takes the representation of a service and creates it.  Returns the server's representation of the service, and an error, if there is any.
func (c *services) Create(service *v1.Service) (result *v1.Service, err error) {
	result = &v1.Service{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("services").
		Body(service).
		Do().
		Into(result)
	return
}

// Update takes the representation of a service and updates it. Returns the server's representation of the service, and an error, if there is any.
func (c *services) Update(service *v1.Service) (result *v1.Service, err error) {
	result = &v1.Service{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("services").
		Name(service.Name).
		Body(service).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *services) UpdateStatus(service *v1.Service) (result *v1.Service, err error) {
	result = &v1.Service{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("services").
		Name(service.Name).
		SubResource("status").
		Body(service).
		Do().
		Into(result)
	return
}

// Delete takes name of the service and deletes it. Returns an error if one occurs.
func (c *services) Delete(name string, options *metav1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("services").
		Name(name).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched service.
func (c *services) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1.Service, err error) {
	result = &v1.Service{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("services").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
