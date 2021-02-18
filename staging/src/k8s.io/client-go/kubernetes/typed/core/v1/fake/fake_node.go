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

	corev1 "k8s.io/api/core/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	applyconfigurationscorev1 "k8s.io/client-go/applyconfigurations/core/v1"
	testing "k8s.io/client-go/testing"
)

// FakeNodes implements NodeInterface
type FakeNodes struct {
	Fake *FakeCoreV1
}

var nodesResource = schema.GroupVersionResource{Group: "", Version: "v1", Resource: "nodes"}

var nodesKind = schema.GroupVersionKind{Group: "", Version: "v1", Kind: "Node"}

// Get takes name of the node, and returns the corresponding node object, and an error if there is any.
func (c *FakeNodes) Get(ctx context.Context, name string, options v1.GetOptions) (result *corev1.Node, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootGetAction(nodesResource, name), &corev1.Node{})
	if obj == nil {
		return nil, err
	}
	return obj.(*corev1.Node), err
}

// List takes label and field selectors, and returns the list of Nodes that match those selectors.
func (c *FakeNodes) List(ctx context.Context, opts v1.ListOptions) (result *corev1.NodeList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootListAction(nodesResource, nodesKind, opts), &corev1.NodeList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &corev1.NodeList{ListMeta: obj.(*corev1.NodeList).ListMeta}
	for _, item := range obj.(*corev1.NodeList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested nodes.
func (c *FakeNodes) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchAction(nodesResource, opts))
}

// Create takes the representation of a node and creates it.  Returns the server's representation of the node, and an error, if there is any.
func (c *FakeNodes) Create(ctx context.Context, node *corev1.Node, opts v1.CreateOptions) (result *corev1.Node, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateAction(nodesResource, node), &corev1.Node{})
	if obj == nil {
		return nil, err
	}
	return obj.(*corev1.Node), err
}

// Update takes the representation of a node and updates it. Returns the server's representation of the node, and an error, if there is any.
func (c *FakeNodes) Update(ctx context.Context, node *corev1.Node, opts v1.UpdateOptions) (result *corev1.Node, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateAction(nodesResource, node), &corev1.Node{})
	if obj == nil {
		return nil, err
	}
	return obj.(*corev1.Node), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeNodes) UpdateStatus(ctx context.Context, node *corev1.Node, opts v1.UpdateOptions) (*corev1.Node, error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateSubresourceAction(nodesResource, "status", node), &corev1.Node{})
	if obj == nil {
		return nil, err
	}
	return obj.(*corev1.Node), err
}

// Delete takes name of the node and deletes it. Returns an error if one occurs.
func (c *FakeNodes) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteAction(nodesResource, name), &corev1.Node{})
	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeNodes) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewRootDeleteCollectionAction(nodesResource, listOpts)

	_, err := c.Fake.Invokes(action, &corev1.NodeList{})
	return err
}

// Patch applies the patch and returns the patched node.
func (c *FakeNodes) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *corev1.Node, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(nodesResource, name, pt, data, subresources...), &corev1.Node{})
	if obj == nil {
		return nil, err
	}
	return obj.(*corev1.Node), err
}

// Apply takes the given apply declarative configuration, applies it and returns the applied node.
func (c *FakeNodes) Apply(ctx context.Context, node *applyconfigurationscorev1.NodeApplyConfiguration, opts v1.ApplyOptions) (result *corev1.Node, err error) {
	if node == nil {
		return nil, fmt.Errorf("node provided to Apply must not be nil")
	}
	data, err := json.Marshal(node)
	if err != nil {
		return nil, err
	}
	name := node.Name
	if name == nil {
		return nil, fmt.Errorf("node.Name must be provided to Apply")
	}
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(nodesResource, *name, types.ApplyPatchType, data), &corev1.Node{})
	if obj == nil {
		return nil, err
	}
	return obj.(*corev1.Node), err
}

// ApplyStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating ApplyStatus().
func (c *FakeNodes) ApplyStatus(ctx context.Context, node *applyconfigurationscorev1.NodeApplyConfiguration, opts v1.ApplyOptions) (result *corev1.Node, err error) {
	if node == nil {
		return nil, fmt.Errorf("node provided to Apply must not be nil")
	}
	data, err := json.Marshal(node)
	if err != nil {
		return nil, err
	}
	name := node.Name
	if name == nil {
		return nil, fmt.Errorf("node.Name must be provided to Apply")
	}
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(nodesResource, *name, types.ApplyPatchType, data, "status"), &corev1.Node{})
	if obj == nil {
		return nil, err
	}
	return obj.(*corev1.Node), err
}
