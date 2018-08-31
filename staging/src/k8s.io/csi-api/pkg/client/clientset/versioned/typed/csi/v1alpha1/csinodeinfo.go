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

package v1alpha1

import (
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
	v1alpha1 "k8s.io/csi-api/pkg/apis/csi/v1alpha1"
	scheme "k8s.io/csi-api/pkg/client/clientset/versioned/scheme"
)

// CSINodeInfosGetter has a method to return a CSINodeInfoInterface.
// A group's client should implement this interface.
type CSINodeInfosGetter interface {
	CSINodeInfos() CSINodeInfoInterface
}

// CSINodeInfoInterface has methods to work with CSINodeInfo resources.
type CSINodeInfoInterface interface {
	Create(*v1alpha1.CSINodeInfo) (*v1alpha1.CSINodeInfo, error)
	Update(*v1alpha1.CSINodeInfo) (*v1alpha1.CSINodeInfo, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.CSINodeInfo, error)
	List(opts v1.ListOptions) (*v1alpha1.CSINodeInfoList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.CSINodeInfo, err error)
	CSINodeInfoExpansion
}

// cSINodeInfos implements CSINodeInfoInterface
type cSINodeInfos struct {
	client rest.Interface
}

// newCSINodeInfos returns a CSINodeInfos
func newCSINodeInfos(c *CsiV1alpha1Client) *cSINodeInfos {
	return &cSINodeInfos{
		client: c.RESTClient(),
	}
}

// Get takes name of the cSINodeInfo, and returns the corresponding cSINodeInfo object, and an error if there is any.
func (c *cSINodeInfos) Get(name string, options v1.GetOptions) (result *v1alpha1.CSINodeInfo, err error) {
	result = &v1alpha1.CSINodeInfo{}
	err = c.client.Get().
		Resource("csinodeinfos").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of CSINodeInfos that match those selectors.
func (c *cSINodeInfos) List(opts v1.ListOptions) (result *v1alpha1.CSINodeInfoList, err error) {
	result = &v1alpha1.CSINodeInfoList{}
	err = c.client.Get().
		Resource("csinodeinfos").
		VersionedParams(&opts, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested cSINodeInfos.
func (c *cSINodeInfos) Watch(opts v1.ListOptions) (watch.Interface, error) {
	opts.Watch = true
	return c.client.Get().
		Resource("csinodeinfos").
		VersionedParams(&opts, scheme.ParameterCodec).
		Watch()
}

// Create takes the representation of a cSINodeInfo and creates it.  Returns the server's representation of the cSINodeInfo, and an error, if there is any.
func (c *cSINodeInfos) Create(cSINodeInfo *v1alpha1.CSINodeInfo) (result *v1alpha1.CSINodeInfo, err error) {
	result = &v1alpha1.CSINodeInfo{}
	err = c.client.Post().
		Resource("csinodeinfos").
		Body(cSINodeInfo).
		Do().
		Into(result)
	return
}

// Update takes the representation of a cSINodeInfo and updates it. Returns the server's representation of the cSINodeInfo, and an error, if there is any.
func (c *cSINodeInfos) Update(cSINodeInfo *v1alpha1.CSINodeInfo) (result *v1alpha1.CSINodeInfo, err error) {
	result = &v1alpha1.CSINodeInfo{}
	err = c.client.Put().
		Resource("csinodeinfos").
		Name(cSINodeInfo.Name).
		Body(cSINodeInfo).
		Do().
		Into(result)
	return
}

// Delete takes name of the cSINodeInfo and deletes it. Returns an error if one occurs.
func (c *cSINodeInfos) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Resource("csinodeinfos").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *cSINodeInfos) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	return c.client.Delete().
		Resource("csinodeinfos").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched cSINodeInfo.
func (c *cSINodeInfos) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.CSINodeInfo, err error) {
	result = &v1alpha1.CSINodeInfo{}
	err = c.client.Patch(pt).
		Resource("csinodeinfos").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
