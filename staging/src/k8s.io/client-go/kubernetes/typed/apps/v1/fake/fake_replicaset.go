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
	apps_v1 "k8s.io/api/apps/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeReplicaSets implements ReplicaSetInterface
type FakeReplicaSets struct {
	Fake *FakeAppsV1
	ns   string
}

var replicasetsResource = schema.GroupVersionResource{Group: "apps", Version: "v1", Resource: "replicasets"}

var replicasetsKind = schema.GroupVersionKind{Group: "apps", Version: "v1", Kind: "ReplicaSet"}

// Get takes name of the replicaSet, and returns the corresponding replicaSet object, and an error if there is any.
func (c *FakeReplicaSets) Get(name string, options v1.GetOptions) (result *apps_v1.ReplicaSet, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(replicasetsResource, c.ns, name), &apps_v1.ReplicaSet{})

	if obj == nil {
		return nil, err
	}
	return obj.(*apps_v1.ReplicaSet), err
}

// List takes label and field selectors, and returns the list of ReplicaSets that match those selectors.
func (c *FakeReplicaSets) List(opts v1.ListOptions) (result *apps_v1.ReplicaSetList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(replicasetsResource, replicasetsKind, c.ns, opts), &apps_v1.ReplicaSetList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &apps_v1.ReplicaSetList{}
	for _, item := range obj.(*apps_v1.ReplicaSetList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested replicaSets.
func (c *FakeReplicaSets) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(replicasetsResource, c.ns, opts))

}

// Create takes the representation of a replicaSet and creates it.  Returns the server's representation of the replicaSet, and an error, if there is any.
func (c *FakeReplicaSets) Create(replicaSet *apps_v1.ReplicaSet) (result *apps_v1.ReplicaSet, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(replicasetsResource, c.ns, replicaSet), &apps_v1.ReplicaSet{})

	if obj == nil {
		return nil, err
	}
	return obj.(*apps_v1.ReplicaSet), err
}

// Update takes the representation of a replicaSet and updates it. Returns the server's representation of the replicaSet, and an error, if there is any.
func (c *FakeReplicaSets) Update(replicaSet *apps_v1.ReplicaSet) (result *apps_v1.ReplicaSet, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(replicasetsResource, c.ns, replicaSet), &apps_v1.ReplicaSet{})

	if obj == nil {
		return nil, err
	}
	return obj.(*apps_v1.ReplicaSet), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeReplicaSets) UpdateStatus(replicaSet *apps_v1.ReplicaSet) (*apps_v1.ReplicaSet, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(replicasetsResource, "status", c.ns, replicaSet), &apps_v1.ReplicaSet{})

	if obj == nil {
		return nil, err
	}
	return obj.(*apps_v1.ReplicaSet), err
}

// Delete takes name of the replicaSet and deletes it. Returns an error if one occurs.
func (c *FakeReplicaSets) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(replicasetsResource, c.ns, name), &apps_v1.ReplicaSet{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeReplicaSets) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(replicasetsResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &apps_v1.ReplicaSetList{})
	return err
}

// Patch applies the patch and returns the patched replicaSet.
func (c *FakeReplicaSets) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *apps_v1.ReplicaSet, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(replicasetsResource, c.ns, name, data, subresources...), &apps_v1.ReplicaSet{})

	if obj == nil {
		return nil, err
	}
	return obj.(*apps_v1.ReplicaSet), err
}
