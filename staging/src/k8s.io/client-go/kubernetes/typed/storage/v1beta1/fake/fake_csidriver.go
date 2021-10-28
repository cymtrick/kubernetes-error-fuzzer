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
	json "encoding/json"
	"fmt"

	v1beta1 "k8s.io/api/storage/v1beta1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	storagev1beta1 "k8s.io/client-go/applyconfigurations/storage/v1beta1"
	testing "k8s.io/client-go/testing"
)

// FakeCSIDrivers implements CSIDriverInterface
type FakeCSIDrivers struct {
	Fake *FakeStorageV1beta1
}

var csidriversResource = schema.GroupVersionResource{Group: "storage.k8s.io", Version: "v1beta1", Resource: "csidrivers"}

var csidriversKind = schema.GroupVersionKind{Group: "storage.k8s.io", Version: "v1beta1", Kind: "CSIDriver"}

// Get takes name of the cSIDriver, and returns the corresponding cSIDriver object, and an error if there is any.
func (c *FakeCSIDrivers) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1beta1.CSIDriver, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootGetAction(csidriversResource, name), &v1beta1.CSIDriver{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.CSIDriver), err
}

// List takes label and field selectors, and returns the list of CSIDrivers that match those selectors.
func (c *FakeCSIDrivers) List(ctx context.Context, opts v1.ListOptions) (result *v1beta1.CSIDriverList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootListAction(csidriversResource, csidriversKind, opts), &v1beta1.CSIDriverList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1beta1.CSIDriverList{ListMeta: obj.(*v1beta1.CSIDriverList).ListMeta}
	for _, item := range obj.(*v1beta1.CSIDriverList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested cSIDrivers.
func (c *FakeCSIDrivers) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchAction(csidriversResource, opts))
}

// Create takes the representation of a cSIDriver and creates it.  Returns the server's representation of the cSIDriver, and an error, if there is any.
func (c *FakeCSIDrivers) Create(ctx context.Context, cSIDriver *v1beta1.CSIDriver, opts v1.CreateOptions) (result *v1beta1.CSIDriver, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateAction(csidriversResource, cSIDriver), &v1beta1.CSIDriver{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.CSIDriver), err
}

// Update takes the representation of a cSIDriver and updates it. Returns the server's representation of the cSIDriver, and an error, if there is any.
func (c *FakeCSIDrivers) Update(ctx context.Context, cSIDriver *v1beta1.CSIDriver, opts v1.UpdateOptions) (result *v1beta1.CSIDriver, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateAction(csidriversResource, cSIDriver), &v1beta1.CSIDriver{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.CSIDriver), err
}

// Delete takes name of the cSIDriver and deletes it. Returns an error if one occurs.
func (c *FakeCSIDrivers) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteActionWithOptions(csidriversResource, name, opts), &v1beta1.CSIDriver{})
	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeCSIDrivers) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewRootDeleteCollectionAction(csidriversResource, listOpts)

	_, err := c.Fake.Invokes(action, &v1beta1.CSIDriverList{})
	return err
}

// Patch applies the patch and returns the patched cSIDriver.
func (c *FakeCSIDrivers) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1beta1.CSIDriver, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(csidriversResource, name, pt, data, subresources...), &v1beta1.CSIDriver{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.CSIDriver), err
}

// Apply takes the given apply declarative configuration, applies it and returns the applied cSIDriver.
func (c *FakeCSIDrivers) Apply(ctx context.Context, cSIDriver *storagev1beta1.CSIDriverApplyConfiguration, opts v1.ApplyOptions) (result *v1beta1.CSIDriver, err error) {
	if cSIDriver == nil {
		return nil, fmt.Errorf("cSIDriver provided to Apply must not be nil")
	}
	data, err := json.Marshal(cSIDriver)
	if err != nil {
		return nil, err
	}
	name := cSIDriver.Name
	if name == nil {
		return nil, fmt.Errorf("cSIDriver.Name must be provided to Apply")
	}
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(csidriversResource, *name, types.ApplyPatchType, data), &v1beta1.CSIDriver{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.CSIDriver), err
}
