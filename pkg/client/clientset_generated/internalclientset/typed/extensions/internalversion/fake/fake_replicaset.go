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
	autoscaling "k8s.io/kubernetes/pkg/apis/autoscaling"
	extensions "k8s.io/kubernetes/pkg/apis/extensions"
)

// FakeReplicaSets implements ReplicaSetInterface
type FakeReplicaSets struct {
	Fake *FakeExtensions
	ns   string
}

var replicasetsResource = schema.GroupVersionResource{Group: "extensions", Version: "", Resource: "replicasets"}

var replicasetsKind = schema.GroupVersionKind{Group: "extensions", Version: "", Kind: "ReplicaSet"}

// Get takes name of the replicaSet, and returns the corresponding replicaSet object, and an error if there is any.
func (c *FakeReplicaSets) Get(name string, options v1.GetOptions) (result *extensions.ReplicaSet, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(replicasetsResource, c.ns, name), &extensions.ReplicaSet{})

	if obj == nil {
		return nil, err
	}
	return obj.(*extensions.ReplicaSet), err
}

// List takes label and field selectors, and returns the list of ReplicaSets that match those selectors.
func (c *FakeReplicaSets) List(opts v1.ListOptions) (result *extensions.ReplicaSetList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(replicasetsResource, replicasetsKind, c.ns, opts), &extensions.ReplicaSetList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &extensions.ReplicaSetList{}
	for _, item := range obj.(*extensions.ReplicaSetList).Items {
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
func (c *FakeReplicaSets) Create(replicaSet *extensions.ReplicaSet) (result *extensions.ReplicaSet, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(replicasetsResource, c.ns, replicaSet), &extensions.ReplicaSet{})

	if obj == nil {
		return nil, err
	}
	return obj.(*extensions.ReplicaSet), err
}

// Update takes the representation of a replicaSet and updates it. Returns the server's representation of the replicaSet, and an error, if there is any.
func (c *FakeReplicaSets) Update(replicaSet *extensions.ReplicaSet) (result *extensions.ReplicaSet, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(replicasetsResource, c.ns, replicaSet), &extensions.ReplicaSet{})

	if obj == nil {
		return nil, err
	}
	return obj.(*extensions.ReplicaSet), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeReplicaSets) UpdateStatus(replicaSet *extensions.ReplicaSet) (*extensions.ReplicaSet, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(replicasetsResource, "status", c.ns, replicaSet), &extensions.ReplicaSet{})

	if obj == nil {
		return nil, err
	}
	return obj.(*extensions.ReplicaSet), err
}

// Delete takes name of the replicaSet and deletes it. Returns an error if one occurs.
func (c *FakeReplicaSets) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(replicasetsResource, c.ns, name), &extensions.ReplicaSet{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeReplicaSets) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(replicasetsResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &extensions.ReplicaSetList{})
	return err
}

// Patch applies the patch and returns the patched replicaSet.
func (c *FakeReplicaSets) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *extensions.ReplicaSet, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(replicasetsResource, c.ns, name, data, subresources...), &extensions.ReplicaSet{})

	if obj == nil {
		return nil, err
	}
	return obj.(*extensions.ReplicaSet), err
}

// GetScale takes name of the replicaSet, and returns the corresponding scale object, and an error if there is any.
func (c *FakeReplicaSets) GetScale(replicaSetName string, options v1.GetOptions) (result *autoscaling.Scale, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetSubresourceAction(replicasetsResource, c.ns, "scale", replicaSetName), &autoscaling.Scale{})

	if obj == nil {
		return nil, err
	}
	return obj.(*autoscaling.Scale), err
}

// UpdateScale takes the representation of a scale and updates it. Returns the server's representation of the scale, and an error, if there is any.
func (c *FakeReplicaSets) UpdateScale(replicaSetName string, scale *autoscaling.Scale) (result *autoscaling.Scale, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(replicasetsResource, "scale", c.ns, scale), &autoscaling.Scale{})

	if obj == nil {
		return nil, err
	}
	return obj.(*autoscaling.Scale), err
}
