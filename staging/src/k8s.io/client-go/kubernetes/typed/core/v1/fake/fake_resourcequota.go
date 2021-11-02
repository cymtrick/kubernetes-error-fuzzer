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

// FakeResourceQuotas implements ResourceQuotaInterface
type FakeResourceQuotas struct {
	Fake *FakeCoreV1
	ns   string
}

var resourcequotasResource = schema.GroupVersionResource{Group: "", Version: "v1", Resource: "resourcequotas"}

var resourcequotasKind = schema.GroupVersionKind{Group: "", Version: "v1", Kind: "ResourceQuota"}

// Get takes name of the resourceQuota, and returns the corresponding resourceQuota object, and an error if there is any.
func (c *FakeResourceQuotas) Get(ctx context.Context, name string, options v1.GetOptions) (result *corev1.ResourceQuota, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(resourcequotasResource, c.ns, name), &corev1.ResourceQuota{})

	if obj == nil {
		return nil, err
	}
	return obj.(*corev1.ResourceQuota), err
}

// List takes label and field selectors, and returns the list of ResourceQuotas that match those selectors.
func (c *FakeResourceQuotas) List(ctx context.Context, opts v1.ListOptions) (result *corev1.ResourceQuotaList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(resourcequotasResource, resourcequotasKind, c.ns, opts), &corev1.ResourceQuotaList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &corev1.ResourceQuotaList{ListMeta: obj.(*corev1.ResourceQuotaList).ListMeta}
	for _, item := range obj.(*corev1.ResourceQuotaList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested resourceQuotas.
func (c *FakeResourceQuotas) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(resourcequotasResource, c.ns, opts))

}

// Create takes the representation of a resourceQuota and creates it.  Returns the server's representation of the resourceQuota, and an error, if there is any.
func (c *FakeResourceQuotas) Create(ctx context.Context, resourceQuota *corev1.ResourceQuota, opts v1.CreateOptions) (result *corev1.ResourceQuota, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(resourcequotasResource, c.ns, resourceQuota), &corev1.ResourceQuota{})

	if obj == nil {
		return nil, err
	}
	return obj.(*corev1.ResourceQuota), err
}

// Update takes the representation of a resourceQuota and updates it. Returns the server's representation of the resourceQuota, and an error, if there is any.
func (c *FakeResourceQuotas) Update(ctx context.Context, resourceQuota *corev1.ResourceQuota, opts v1.UpdateOptions) (result *corev1.ResourceQuota, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(resourcequotasResource, c.ns, resourceQuota), &corev1.ResourceQuota{})

	if obj == nil {
		return nil, err
	}
	return obj.(*corev1.ResourceQuota), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeResourceQuotas) UpdateStatus(ctx context.Context, resourceQuota *corev1.ResourceQuota, opts v1.UpdateOptions) (*corev1.ResourceQuota, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(resourcequotasResource, "status", c.ns, resourceQuota), &corev1.ResourceQuota{})

	if obj == nil {
		return nil, err
	}
	return obj.(*corev1.ResourceQuota), err
}

// Delete takes name of the resourceQuota and deletes it. Returns an error if one occurs.
func (c *FakeResourceQuotas) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteActionWithOptions(resourcequotasResource, c.ns, name, opts), &corev1.ResourceQuota{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeResourceQuotas) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(resourcequotasResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &corev1.ResourceQuotaList{})
	return err
}

// Patch applies the patch and returns the patched resourceQuota.
func (c *FakeResourceQuotas) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *corev1.ResourceQuota, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(resourcequotasResource, c.ns, name, pt, data, subresources...), &corev1.ResourceQuota{})

	if obj == nil {
		return nil, err
	}
	return obj.(*corev1.ResourceQuota), err
}

// Apply takes the given apply declarative configuration, applies it and returns the applied resourceQuota.
func (c *FakeResourceQuotas) Apply(ctx context.Context, resourceQuota *applyconfigurationscorev1.ResourceQuotaApplyConfiguration, opts v1.ApplyOptions) (result *corev1.ResourceQuota, err error) {
	if resourceQuota == nil {
		return nil, fmt.Errorf("resourceQuota provided to Apply must not be nil")
	}
	data, err := json.Marshal(resourceQuota)
	if err != nil {
		return nil, err
	}
	name := resourceQuota.Name
	if name == nil {
		return nil, fmt.Errorf("resourceQuota.Name must be provided to Apply")
	}
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(resourcequotasResource, c.ns, *name, types.ApplyPatchType, data), &corev1.ResourceQuota{})

	if obj == nil {
		return nil, err
	}
	return obj.(*corev1.ResourceQuota), err
}

// ApplyStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating ApplyStatus().
func (c *FakeResourceQuotas) ApplyStatus(ctx context.Context, resourceQuota *applyconfigurationscorev1.ResourceQuotaApplyConfiguration, opts v1.ApplyOptions) (result *corev1.ResourceQuota, err error) {
	if resourceQuota == nil {
		return nil, fmt.Errorf("resourceQuota provided to Apply must not be nil")
	}
	data, err := json.Marshal(resourceQuota)
	if err != nil {
		return nil, err
	}
	name := resourceQuota.Name
	if name == nil {
		return nil, fmt.Errorf("resourceQuota.Name must be provided to Apply")
	}
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(resourcequotasResource, c.ns, *name, types.ApplyPatchType, data, "status"), &corev1.ResourceQuota{})

	if obj == nil {
		return nil, err
	}
	return obj.(*corev1.ResourceQuota), err
}
