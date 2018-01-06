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

package fake

import (
	v1alpha1 "k8s.io/api/storage/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeVolumeAttachments implements VolumeAttachmentInterface
type FakeVolumeAttachments struct {
	Fake *FakeStorageV1alpha1
}

var volumeattachmentsResource = schema.GroupVersionResource{Group: "storage.k8s.io", Version: "v1alpha1", Resource: "volumeattachments"}

var volumeattachmentsKind = schema.GroupVersionKind{Group: "storage.k8s.io", Version: "v1alpha1", Kind: "VolumeAttachment"}

// Get takes name of the volumeAttachment, and returns the corresponding volumeAttachment object, and an error if there is any.
func (c *FakeVolumeAttachments) Get(name string, options v1.GetOptions) (result *v1alpha1.VolumeAttachment, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootGetAction(volumeattachmentsResource, name), &v1alpha1.VolumeAttachment{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.VolumeAttachment), err
}

// List takes label and field selectors, and returns the list of VolumeAttachments that match those selectors.
func (c *FakeVolumeAttachments) List(opts v1.ListOptions) (result *v1alpha1.VolumeAttachmentList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootListAction(volumeattachmentsResource, volumeattachmentsKind, opts), &v1alpha1.VolumeAttachmentList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.VolumeAttachmentList{}
	for _, item := range obj.(*v1alpha1.VolumeAttachmentList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested volumeAttachments.
func (c *FakeVolumeAttachments) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchAction(volumeattachmentsResource, opts))
}

// Create takes the representation of a volumeAttachment and creates it.  Returns the server's representation of the volumeAttachment, and an error, if there is any.
func (c *FakeVolumeAttachments) Create(volumeAttachment *v1alpha1.VolumeAttachment) (result *v1alpha1.VolumeAttachment, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateAction(volumeattachmentsResource, volumeAttachment), &v1alpha1.VolumeAttachment{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.VolumeAttachment), err
}

// Update takes the representation of a volumeAttachment and updates it. Returns the server's representation of the volumeAttachment, and an error, if there is any.
func (c *FakeVolumeAttachments) Update(volumeAttachment *v1alpha1.VolumeAttachment) (result *v1alpha1.VolumeAttachment, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateAction(volumeattachmentsResource, volumeAttachment), &v1alpha1.VolumeAttachment{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.VolumeAttachment), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeVolumeAttachments) UpdateStatus(volumeAttachment *v1alpha1.VolumeAttachment) (*v1alpha1.VolumeAttachment, error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateSubresourceAction(volumeattachmentsResource, "status", volumeAttachment), &v1alpha1.VolumeAttachment{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.VolumeAttachment), err
}

// Delete takes name of the volumeAttachment and deletes it. Returns an error if one occurs.
func (c *FakeVolumeAttachments) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteAction(volumeattachmentsResource, name), &v1alpha1.VolumeAttachment{})
	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeVolumeAttachments) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewRootDeleteCollectionAction(volumeattachmentsResource, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.VolumeAttachmentList{})
	return err
}

// Patch applies the patch and returns the patched volumeAttachment.
func (c *FakeVolumeAttachments) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.VolumeAttachment, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(volumeattachmentsResource, name, data, subresources...), &v1alpha1.VolumeAttachment{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.VolumeAttachment), err
}
