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
	v2beta1 "k8s.io/api/autoscaling/v2beta1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeVerticalPodAutoscalers implements VerticalPodAutoscalerInterface
type FakeVerticalPodAutoscalers struct {
	Fake *FakeAutoscalingV2beta1
	ns   string
}

var verticalpodautoscalersResource = schema.GroupVersionResource{Group: "autoscaling", Version: "v2beta1", Resource: "verticalpodautoscalers"}

var verticalpodautoscalersKind = schema.GroupVersionKind{Group: "autoscaling", Version: "v2beta1", Kind: "VerticalPodAutoscaler"}

// Get takes name of the verticalPodAutoscaler, and returns the corresponding verticalPodAutoscaler object, and an error if there is any.
func (c *FakeVerticalPodAutoscalers) Get(name string, options v1.GetOptions) (result *v2beta1.VerticalPodAutoscaler, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(verticalpodautoscalersResource, c.ns, name), &v2beta1.VerticalPodAutoscaler{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v2beta1.VerticalPodAutoscaler), err
}

// List takes label and field selectors, and returns the list of VerticalPodAutoscalers that match those selectors.
func (c *FakeVerticalPodAutoscalers) List(opts v1.ListOptions) (result *v2beta1.VerticalPodAutoscalerList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(verticalpodautoscalersResource, verticalpodautoscalersKind, c.ns, opts), &v2beta1.VerticalPodAutoscalerList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v2beta1.VerticalPodAutoscalerList{ListMeta: obj.(*v2beta1.VerticalPodAutoscalerList).ListMeta}
	for _, item := range obj.(*v2beta1.VerticalPodAutoscalerList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested verticalPodAutoscalers.
func (c *FakeVerticalPodAutoscalers) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(verticalpodautoscalersResource, c.ns, opts))

}

// Create takes the representation of a verticalPodAutoscaler and creates it.  Returns the server's representation of the verticalPodAutoscaler, and an error, if there is any.
func (c *FakeVerticalPodAutoscalers) Create(verticalPodAutoscaler *v2beta1.VerticalPodAutoscaler) (result *v2beta1.VerticalPodAutoscaler, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(verticalpodautoscalersResource, c.ns, verticalPodAutoscaler), &v2beta1.VerticalPodAutoscaler{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v2beta1.VerticalPodAutoscaler), err
}

// Update takes the representation of a verticalPodAutoscaler and updates it. Returns the server's representation of the verticalPodAutoscaler, and an error, if there is any.
func (c *FakeVerticalPodAutoscalers) Update(verticalPodAutoscaler *v2beta1.VerticalPodAutoscaler) (result *v2beta1.VerticalPodAutoscaler, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(verticalpodautoscalersResource, c.ns, verticalPodAutoscaler), &v2beta1.VerticalPodAutoscaler{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v2beta1.VerticalPodAutoscaler), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeVerticalPodAutoscalers) UpdateStatus(verticalPodAutoscaler *v2beta1.VerticalPodAutoscaler) (*v2beta1.VerticalPodAutoscaler, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(verticalpodautoscalersResource, "status", c.ns, verticalPodAutoscaler), &v2beta1.VerticalPodAutoscaler{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v2beta1.VerticalPodAutoscaler), err
}

// Delete takes name of the verticalPodAutoscaler and deletes it. Returns an error if one occurs.
func (c *FakeVerticalPodAutoscalers) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(verticalpodautoscalersResource, c.ns, name), &v2beta1.VerticalPodAutoscaler{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeVerticalPodAutoscalers) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(verticalpodautoscalersResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &v2beta1.VerticalPodAutoscalerList{})
	return err
}

// Patch applies the patch and returns the patched verticalPodAutoscaler.
func (c *FakeVerticalPodAutoscalers) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v2beta1.VerticalPodAutoscaler, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(verticalpodautoscalersResource, c.ns, name, data, subresources...), &v2beta1.VerticalPodAutoscaler{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v2beta1.VerticalPodAutoscaler), err
}
