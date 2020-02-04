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

package fake

import (
	"context"

	v1beta1 "k8s.io/api/storage/v1beta1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeCSINodes implements CSINodeInterface
type FakeCSINodes struct {
	Fake *FakeStorageV1beta1
}

var csinodesResource = schema.GroupVersionResource{Group: "storage.k8s.io", Version: "v1beta1", Resource: "csinodes"}

var csinodesKind = schema.GroupVersionKind{Group: "storage.k8s.io", Version: "v1beta1", Kind: "CSINode"}

// Get takes name of the cSINode, and returns the corresponding cSINode object, and an error if there is any.
func (c *FakeCSINodes) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1beta1.CSINode, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootGetAction(csinodesResource, name), &v1beta1.CSINode{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.CSINode), err
}

// List takes label and field selectors, and returns the list of CSINodes that match those selectors.
func (c *FakeCSINodes) List(ctx context.Context, opts v1.ListOptions) (result *v1beta1.CSINodeList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootListAction(csinodesResource, csinodesKind, opts), &v1beta1.CSINodeList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1beta1.CSINodeList{ListMeta: obj.(*v1beta1.CSINodeList).ListMeta}
	for _, item := range obj.(*v1beta1.CSINodeList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested cSINodes.
func (c *FakeCSINodes) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchAction(csinodesResource, opts))
}

// Create takes the representation of a cSINode and creates it.  Returns the server's representation of the cSINode, and an error, if there is any.
func (c *FakeCSINodes) Create(ctx context.Context, cSINode *v1beta1.CSINode) (result *v1beta1.CSINode, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateAction(csinodesResource, cSINode), &v1beta1.CSINode{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.CSINode), err
}

// Update takes the representation of a cSINode and updates it. Returns the server's representation of the cSINode, and an error, if there is any.
func (c *FakeCSINodes) Update(ctx context.Context, cSINode *v1beta1.CSINode) (result *v1beta1.CSINode, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateAction(csinodesResource, cSINode), &v1beta1.CSINode{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.CSINode), err
}

// Delete takes name of the cSINode and deletes it. Returns an error if one occurs.
func (c *FakeCSINodes) Delete(ctx context.Context, name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteAction(csinodesResource, name), &v1beta1.CSINode{})
	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeCSINodes) DeleteCollection(ctx context.Context, options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewRootDeleteCollectionAction(csinodesResource, listOptions)

	_, err := c.Fake.Invokes(action, &v1beta1.CSINodeList{})
	return err
}

// Patch applies the patch and returns the patched cSINode.
func (c *FakeCSINodes) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, subresources ...string) (result *v1beta1.CSINode, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(csinodesResource, name, pt, data, subresources...), &v1beta1.CSINode{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.CSINode), err
}
