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

package fake

import (
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
	core "k8s.io/kubernetes/pkg/apis/core"
)

// FakeLimitRanges implements LimitRangeInterface
type FakeLimitRanges struct {
	Fake *FakeCore
	ns   string
}

var limitrangesResource = schema.GroupVersionResource{Group: "", Version: "", Resource: "limitranges"}

var limitrangesKind = schema.GroupVersionKind{Group: "", Version: "", Kind: "LimitRange"}

// Get takes name of the limitRange, and returns the corresponding limitRange object, and an error if there is any.
func (c *FakeLimitRanges) Get(name string, options v1.GetOptions) (result *core.LimitRange, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(limitrangesResource, c.ns, name), &core.LimitRange{})

	if obj == nil {
		return nil, err
	}
	return obj.(*core.LimitRange), err
}

// List takes label and field selectors, and returns the list of LimitRanges that match those selectors.
func (c *FakeLimitRanges) List(opts v1.ListOptions) (result *core.LimitRangeList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(limitrangesResource, limitrangesKind, c.ns, opts), &core.LimitRangeList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &core.LimitRangeList{}
	for _, item := range obj.(*core.LimitRangeList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested limitRanges.
func (c *FakeLimitRanges) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(limitrangesResource, c.ns, opts))

}

// Create takes the representation of a limitRange and creates it.  Returns the server's representation of the limitRange, and an error, if there is any.
func (c *FakeLimitRanges) Create(limitRange *core.LimitRange) (result *core.LimitRange, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(limitrangesResource, c.ns, limitRange), &core.LimitRange{})

	if obj == nil {
		return nil, err
	}
	return obj.(*core.LimitRange), err
}

// Update takes the representation of a limitRange and updates it. Returns the server's representation of the limitRange, and an error, if there is any.
func (c *FakeLimitRanges) Update(limitRange *core.LimitRange) (result *core.LimitRange, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(limitrangesResource, c.ns, limitRange), &core.LimitRange{})

	if obj == nil {
		return nil, err
	}
	return obj.(*core.LimitRange), err
}

// Delete takes name of the limitRange and deletes it. Returns an error if one occurs.
func (c *FakeLimitRanges) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(limitrangesResource, c.ns, name), &core.LimitRange{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeLimitRanges) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(limitrangesResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &core.LimitRangeList{})
	return err
}

// Patch applies the patch and returns the patched limitRange.
func (c *FakeLimitRanges) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *core.LimitRange, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(limitrangesResource, c.ns, name, data, subresources...), &core.LimitRange{})

	if obj == nil {
		return nil, err
	}
	return obj.(*core.LimitRange), err
}
