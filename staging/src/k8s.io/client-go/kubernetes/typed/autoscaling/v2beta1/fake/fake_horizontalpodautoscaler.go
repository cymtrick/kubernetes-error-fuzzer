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

	v2beta1 "k8s.io/api/autoscaling/v2beta1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeHorizontalPodAutoscalers implements HorizontalPodAutoscalerInterface
type FakeHorizontalPodAutoscalers struct {
	Fake *FakeAutoscalingV2beta1
	ns   string
}

var horizontalpodautoscalersResource = schema.GroupVersionResource{Group: "autoscaling", Version: "v2beta1", Resource: "horizontalpodautoscalers"}

var horizontalpodautoscalersKind = schema.GroupVersionKind{Group: "autoscaling", Version: "v2beta1", Kind: "HorizontalPodAutoscaler"}

// Get takes name of the horizontalPodAutoscaler, and returns the corresponding horizontalPodAutoscaler object, and an error if there is any.
func (c *FakeHorizontalPodAutoscalers) Get(ctx context.Context, name string, options v1.GetOptions) (result *v2beta1.HorizontalPodAutoscaler, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(horizontalpodautoscalersResource, c.ns, name), &v2beta1.HorizontalPodAutoscaler{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v2beta1.HorizontalPodAutoscaler), err
}

// List takes label and field selectors, and returns the list of HorizontalPodAutoscalers that match those selectors.
func (c *FakeHorizontalPodAutoscalers) List(ctx context.Context, opts v1.ListOptions) (result *v2beta1.HorizontalPodAutoscalerList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(horizontalpodautoscalersResource, horizontalpodautoscalersKind, c.ns, opts), &v2beta1.HorizontalPodAutoscalerList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v2beta1.HorizontalPodAutoscalerList{ListMeta: obj.(*v2beta1.HorizontalPodAutoscalerList).ListMeta}
	for _, item := range obj.(*v2beta1.HorizontalPodAutoscalerList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested horizontalPodAutoscalers.
func (c *FakeHorizontalPodAutoscalers) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(horizontalpodautoscalersResource, c.ns, opts))

}

// Create takes the representation of a horizontalPodAutoscaler and creates it.  Returns the server's representation of the horizontalPodAutoscaler, and an error, if there is any.
func (c *FakeHorizontalPodAutoscalers) Create(ctx context.Context, horizontalPodAutoscaler *v2beta1.HorizontalPodAutoscaler) (result *v2beta1.HorizontalPodAutoscaler, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(horizontalpodautoscalersResource, c.ns, horizontalPodAutoscaler), &v2beta1.HorizontalPodAutoscaler{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v2beta1.HorizontalPodAutoscaler), err
}

// Update takes the representation of a horizontalPodAutoscaler and updates it. Returns the server's representation of the horizontalPodAutoscaler, and an error, if there is any.
func (c *FakeHorizontalPodAutoscalers) Update(ctx context.Context, horizontalPodAutoscaler *v2beta1.HorizontalPodAutoscaler) (result *v2beta1.HorizontalPodAutoscaler, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(horizontalpodautoscalersResource, c.ns, horizontalPodAutoscaler), &v2beta1.HorizontalPodAutoscaler{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v2beta1.HorizontalPodAutoscaler), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeHorizontalPodAutoscalers) UpdateStatus(ctx context.Context, horizontalPodAutoscaler *v2beta1.HorizontalPodAutoscaler) (*v2beta1.HorizontalPodAutoscaler, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(horizontalpodautoscalersResource, "status", c.ns, horizontalPodAutoscaler), &v2beta1.HorizontalPodAutoscaler{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v2beta1.HorizontalPodAutoscaler), err
}

// Delete takes name of the horizontalPodAutoscaler and deletes it. Returns an error if one occurs.
func (c *FakeHorizontalPodAutoscalers) Delete(ctx context.Context, name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(horizontalpodautoscalersResource, c.ns, name), &v2beta1.HorizontalPodAutoscaler{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeHorizontalPodAutoscalers) DeleteCollection(ctx context.Context, options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(horizontalpodautoscalersResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &v2beta1.HorizontalPodAutoscalerList{})
	return err
}

// Patch applies the patch and returns the patched horizontalPodAutoscaler.
func (c *FakeHorizontalPodAutoscalers) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, subresources ...string) (result *v2beta1.HorizontalPodAutoscaler, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(horizontalpodautoscalersResource, c.ns, name, pt, data, subresources...), &v2beta1.HorizontalPodAutoscaler{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v2beta1.HorizontalPodAutoscaler), err
}
