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
	appsv1 "k8s.io/api/apps/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeDaemonSets implements DaemonSetInterface
type FakeDaemonSets struct {
	Fake *FakeAppsV1
	ns   string
}

var daemonsetsResource = schema.GroupVersionResource{Group: "apps", Version: "v1", Resource: "daemonsets"}

var daemonsetsKind = schema.GroupVersionKind{Group: "apps", Version: "v1", Kind: "DaemonSet"}

// Get takes name of the daemonSet, and returns the corresponding daemonSet object, and an error if there is any.
func (c *FakeDaemonSets) Get(name string, options v1.GetOptions) (result *appsv1.DaemonSet, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(daemonsetsResource, c.ns, name), &appsv1.DaemonSet{})

	if obj == nil {
		return nil, err
	}
	return obj.(*appsv1.DaemonSet), err
}

// List takes label and field selectors, and returns the list of DaemonSets that match those selectors.
func (c *FakeDaemonSets) List(opts v1.ListOptions) (result *appsv1.DaemonSetList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(daemonsetsResource, daemonsetsKind, c.ns, opts), &appsv1.DaemonSetList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &appsv1.DaemonSetList{ListMeta: obj.(*appsv1.DaemonSetList).ListMeta}
	for _, item := range obj.(*appsv1.DaemonSetList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested daemonSets.
func (c *FakeDaemonSets) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(daemonsetsResource, c.ns, opts))

}

// Create takes the representation of a daemonSet and creates it.  Returns the server's representation of the daemonSet, and an error, if there is any.
func (c *FakeDaemonSets) Create(daemonSet *appsv1.DaemonSet) (result *appsv1.DaemonSet, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(daemonsetsResource, c.ns, daemonSet), &appsv1.DaemonSet{})

	if obj == nil {
		return nil, err
	}
	return obj.(*appsv1.DaemonSet), err
}

// Update takes the representation of a daemonSet and updates it. Returns the server's representation of the daemonSet, and an error, if there is any.
func (c *FakeDaemonSets) Update(daemonSet *appsv1.DaemonSet) (result *appsv1.DaemonSet, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(daemonsetsResource, c.ns, daemonSet), &appsv1.DaemonSet{})

	if obj == nil {
		return nil, err
	}
	return obj.(*appsv1.DaemonSet), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeDaemonSets) UpdateStatus(daemonSet *appsv1.DaemonSet) (*appsv1.DaemonSet, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(daemonsetsResource, "status", c.ns, daemonSet), &appsv1.DaemonSet{})

	if obj == nil {
		return nil, err
	}
	return obj.(*appsv1.DaemonSet), err
}

// Delete takes name of the daemonSet and deletes it. Returns an error if one occurs.
func (c *FakeDaemonSets) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(daemonsetsResource, c.ns, name), &appsv1.DaemonSet{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeDaemonSets) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(daemonsetsResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &appsv1.DaemonSetList{})
	return err
}

// Patch applies the patch and returns the patched daemonSet.
func (c *FakeDaemonSets) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *appsv1.DaemonSet, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(daemonsetsResource, c.ns, name, pt, data, subresources...), &appsv1.DaemonSet{})

	if obj == nil {
		return nil, err
	}
	return obj.(*appsv1.DaemonSet), err
}
