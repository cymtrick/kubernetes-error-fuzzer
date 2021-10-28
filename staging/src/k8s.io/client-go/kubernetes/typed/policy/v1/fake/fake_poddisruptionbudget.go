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

	policyv1 "k8s.io/api/policy/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	applyconfigurationspolicyv1 "k8s.io/client-go/applyconfigurations/policy/v1"
	testing "k8s.io/client-go/testing"
)

// FakePodDisruptionBudgets implements PodDisruptionBudgetInterface
type FakePodDisruptionBudgets struct {
	Fake *FakePolicyV1
	ns   string
}

var poddisruptionbudgetsResource = schema.GroupVersionResource{Group: "policy", Version: "v1", Resource: "poddisruptionbudgets"}

var poddisruptionbudgetsKind = schema.GroupVersionKind{Group: "policy", Version: "v1", Kind: "PodDisruptionBudget"}

// Get takes name of the podDisruptionBudget, and returns the corresponding podDisruptionBudget object, and an error if there is any.
func (c *FakePodDisruptionBudgets) Get(ctx context.Context, name string, options v1.GetOptions) (result *policyv1.PodDisruptionBudget, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(poddisruptionbudgetsResource, c.ns, name), &policyv1.PodDisruptionBudget{})

	if obj == nil {
		return nil, err
	}
	return obj.(*policyv1.PodDisruptionBudget), err
}

// List takes label and field selectors, and returns the list of PodDisruptionBudgets that match those selectors.
func (c *FakePodDisruptionBudgets) List(ctx context.Context, opts v1.ListOptions) (result *policyv1.PodDisruptionBudgetList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(poddisruptionbudgetsResource, poddisruptionbudgetsKind, c.ns, opts), &policyv1.PodDisruptionBudgetList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &policyv1.PodDisruptionBudgetList{ListMeta: obj.(*policyv1.PodDisruptionBudgetList).ListMeta}
	for _, item := range obj.(*policyv1.PodDisruptionBudgetList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested podDisruptionBudgets.
func (c *FakePodDisruptionBudgets) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(poddisruptionbudgetsResource, c.ns, opts))

}

// Create takes the representation of a podDisruptionBudget and creates it.  Returns the server's representation of the podDisruptionBudget, and an error, if there is any.
func (c *FakePodDisruptionBudgets) Create(ctx context.Context, podDisruptionBudget *policyv1.PodDisruptionBudget, opts v1.CreateOptions) (result *policyv1.PodDisruptionBudget, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(poddisruptionbudgetsResource, c.ns, podDisruptionBudget), &policyv1.PodDisruptionBudget{})

	if obj == nil {
		return nil, err
	}
	return obj.(*policyv1.PodDisruptionBudget), err
}

// Update takes the representation of a podDisruptionBudget and updates it. Returns the server's representation of the podDisruptionBudget, and an error, if there is any.
func (c *FakePodDisruptionBudgets) Update(ctx context.Context, podDisruptionBudget *policyv1.PodDisruptionBudget, opts v1.UpdateOptions) (result *policyv1.PodDisruptionBudget, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(poddisruptionbudgetsResource, c.ns, podDisruptionBudget), &policyv1.PodDisruptionBudget{})

	if obj == nil {
		return nil, err
	}
	return obj.(*policyv1.PodDisruptionBudget), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakePodDisruptionBudgets) UpdateStatus(ctx context.Context, podDisruptionBudget *policyv1.PodDisruptionBudget, opts v1.UpdateOptions) (*policyv1.PodDisruptionBudget, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(poddisruptionbudgetsResource, "status", c.ns, podDisruptionBudget), &policyv1.PodDisruptionBudget{})

	if obj == nil {
		return nil, err
	}
	return obj.(*policyv1.PodDisruptionBudget), err
}

// Delete takes name of the podDisruptionBudget and deletes it. Returns an error if one occurs.
func (c *FakePodDisruptionBudgets) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteActionWithOptions(poddisruptionbudgetsResource, c.ns, name, opts), &policyv1.PodDisruptionBudget{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakePodDisruptionBudgets) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(poddisruptionbudgetsResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &policyv1.PodDisruptionBudgetList{})
	return err
}

// Patch applies the patch and returns the patched podDisruptionBudget.
func (c *FakePodDisruptionBudgets) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *policyv1.PodDisruptionBudget, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(poddisruptionbudgetsResource, c.ns, name, pt, data, subresources...), &policyv1.PodDisruptionBudget{})

	if obj == nil {
		return nil, err
	}
	return obj.(*policyv1.PodDisruptionBudget), err
}

// Apply takes the given apply declarative configuration, applies it and returns the applied podDisruptionBudget.
func (c *FakePodDisruptionBudgets) Apply(ctx context.Context, podDisruptionBudget *applyconfigurationspolicyv1.PodDisruptionBudgetApplyConfiguration, opts v1.ApplyOptions) (result *policyv1.PodDisruptionBudget, err error) {
	if podDisruptionBudget == nil {
		return nil, fmt.Errorf("podDisruptionBudget provided to Apply must not be nil")
	}
	data, err := json.Marshal(podDisruptionBudget)
	if err != nil {
		return nil, err
	}
	name := podDisruptionBudget.Name
	if name == nil {
		return nil, fmt.Errorf("podDisruptionBudget.Name must be provided to Apply")
	}
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(poddisruptionbudgetsResource, c.ns, *name, types.ApplyPatchType, data), &policyv1.PodDisruptionBudget{})

	if obj == nil {
		return nil, err
	}
	return obj.(*policyv1.PodDisruptionBudget), err
}

// ApplyStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating ApplyStatus().
func (c *FakePodDisruptionBudgets) ApplyStatus(ctx context.Context, podDisruptionBudget *applyconfigurationspolicyv1.PodDisruptionBudgetApplyConfiguration, opts v1.ApplyOptions) (result *policyv1.PodDisruptionBudget, err error) {
	if podDisruptionBudget == nil {
		return nil, fmt.Errorf("podDisruptionBudget provided to Apply must not be nil")
	}
	data, err := json.Marshal(podDisruptionBudget)
	if err != nil {
		return nil, err
	}
	name := podDisruptionBudget.Name
	if name == nil {
		return nil, fmt.Errorf("podDisruptionBudget.Name must be provided to Apply")
	}
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(poddisruptionbudgetsResource, c.ns, *name, types.ApplyPatchType, data, "status"), &policyv1.PodDisruptionBudget{})

	if obj == nil {
		return nil, err
	}
	return obj.(*policyv1.PodDisruptionBudget), err
}
