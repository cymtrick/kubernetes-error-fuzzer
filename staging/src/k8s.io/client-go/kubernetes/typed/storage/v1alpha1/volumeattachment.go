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
	v1alpha1 "k8s.io/api/storage/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	scheme "k8s.io/client-go/kubernetes/scheme"
	rest "k8s.io/client-go/rest"
)

// VolumeAttachmentsGetter has a method to return a VolumeAttachmentInterface.
// A group's client should implement this interface.
type VolumeAttachmentsGetter interface {
	VolumeAttachments() VolumeAttachmentInterface
}

// VolumeAttachmentInterface has methods to work with VolumeAttachment resources.
type VolumeAttachmentInterface interface {
	Create(*v1alpha1.VolumeAttachment) (*v1alpha1.VolumeAttachment, error)
	Update(*v1alpha1.VolumeAttachment) (*v1alpha1.VolumeAttachment, error)
	UpdateStatus(*v1alpha1.VolumeAttachment) (*v1alpha1.VolumeAttachment, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.VolumeAttachment, error)
	List(opts v1.ListOptions) (*v1alpha1.VolumeAttachmentList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.VolumeAttachment, err error)
	VolumeAttachmentExpansion
}

// volumeAttachments implements VolumeAttachmentInterface
type volumeAttachments struct {
	client rest.Interface
}

// newVolumeAttachments returns a VolumeAttachments
func newVolumeAttachments(c *StorageV1alpha1Client) *volumeAttachments {
	return &volumeAttachments{
		client: c.RESTClient(),
	}
}

// Get takes name of the volumeAttachment, and returns the corresponding volumeAttachment object, and an error if there is any.
func (c *volumeAttachments) Get(name string, options v1.GetOptions) (result *v1alpha1.VolumeAttachment, err error) {
	result = &v1alpha1.VolumeAttachment{}
	err = c.client.Get().
		Resource("volumeattachments").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of VolumeAttachments that match those selectors.
func (c *volumeAttachments) List(opts v1.ListOptions) (result *v1alpha1.VolumeAttachmentList, err error) {
	result = &v1alpha1.VolumeAttachmentList{}
	err = c.client.Get().
		Resource("volumeattachments").
		VersionedParams(&opts, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested volumeAttachments.
func (c *volumeAttachments) Watch(opts v1.ListOptions) (watch.Interface, error) {
	opts.Watch = true
	return c.client.Get().
		Resource("volumeattachments").
		VersionedParams(&opts, scheme.ParameterCodec).
		Watch()
}

// Create takes the representation of a volumeAttachment and creates it.  Returns the server's representation of the volumeAttachment, and an error, if there is any.
func (c *volumeAttachments) Create(volumeAttachment *v1alpha1.VolumeAttachment) (result *v1alpha1.VolumeAttachment, err error) {
	result = &v1alpha1.VolumeAttachment{}
	err = c.client.Post().
		Resource("volumeattachments").
		Body(volumeAttachment).
		Do().
		Into(result)
	return
}

// Update takes the representation of a volumeAttachment and updates it. Returns the server's representation of the volumeAttachment, and an error, if there is any.
func (c *volumeAttachments) Update(volumeAttachment *v1alpha1.VolumeAttachment) (result *v1alpha1.VolumeAttachment, err error) {
	result = &v1alpha1.VolumeAttachment{}
	err = c.client.Put().
		Resource("volumeattachments").
		Name(volumeAttachment.Name).
		Body(volumeAttachment).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *volumeAttachments) UpdateStatus(volumeAttachment *v1alpha1.VolumeAttachment) (result *v1alpha1.VolumeAttachment, err error) {
	result = &v1alpha1.VolumeAttachment{}
	err = c.client.Put().
		Resource("volumeattachments").
		Name(volumeAttachment.Name).
		SubResource("status").
		Body(volumeAttachment).
		Do().
		Into(result)
	return
}

// Delete takes name of the volumeAttachment and deletes it. Returns an error if one occurs.
func (c *volumeAttachments) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Resource("volumeattachments").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *volumeAttachments) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	return c.client.Delete().
		Resource("volumeattachments").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched volumeAttachment.
func (c *volumeAttachments) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.VolumeAttachment, err error) {
	result = &v1alpha1.VolumeAttachment{}
	err = c.client.Patch(pt).
		Resource("volumeattachments").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
