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

	v1beta2 "k8s.io/api/flowcontrol/v1beta2"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	flowcontrolv1beta2 "k8s.io/client-go/applyconfigurations/flowcontrol/v1beta2"
	testing "k8s.io/client-go/testing"
)

// FakePriorityLevelConfigurations implements PriorityLevelConfigurationInterface
type FakePriorityLevelConfigurations struct {
	Fake *FakeFlowcontrolV1beta2
}

var prioritylevelconfigurationsResource = schema.GroupVersionResource{Group: "flowcontrol.apiserver.k8s.io", Version: "v1beta2", Resource: "prioritylevelconfigurations"}

var prioritylevelconfigurationsKind = schema.GroupVersionKind{Group: "flowcontrol.apiserver.k8s.io", Version: "v1beta2", Kind: "PriorityLevelConfiguration"}

// Get takes name of the priorityLevelConfiguration, and returns the corresponding priorityLevelConfiguration object, and an error if there is any.
func (c *FakePriorityLevelConfigurations) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1beta2.PriorityLevelConfiguration, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootGetAction(prioritylevelconfigurationsResource, name), &v1beta2.PriorityLevelConfiguration{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta2.PriorityLevelConfiguration), err
}

// List takes label and field selectors, and returns the list of PriorityLevelConfigurations that match those selectors.
func (c *FakePriorityLevelConfigurations) List(ctx context.Context, opts v1.ListOptions) (result *v1beta2.PriorityLevelConfigurationList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootListAction(prioritylevelconfigurationsResource, prioritylevelconfigurationsKind, opts), &v1beta2.PriorityLevelConfigurationList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1beta2.PriorityLevelConfigurationList{ListMeta: obj.(*v1beta2.PriorityLevelConfigurationList).ListMeta}
	for _, item := range obj.(*v1beta2.PriorityLevelConfigurationList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested priorityLevelConfigurations.
func (c *FakePriorityLevelConfigurations) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchAction(prioritylevelconfigurationsResource, opts))
}

// Create takes the representation of a priorityLevelConfiguration and creates it.  Returns the server's representation of the priorityLevelConfiguration, and an error, if there is any.
func (c *FakePriorityLevelConfigurations) Create(ctx context.Context, priorityLevelConfiguration *v1beta2.PriorityLevelConfiguration, opts v1.CreateOptions) (result *v1beta2.PriorityLevelConfiguration, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateAction(prioritylevelconfigurationsResource, priorityLevelConfiguration), &v1beta2.PriorityLevelConfiguration{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta2.PriorityLevelConfiguration), err
}

// Update takes the representation of a priorityLevelConfiguration and updates it. Returns the server's representation of the priorityLevelConfiguration, and an error, if there is any.
func (c *FakePriorityLevelConfigurations) Update(ctx context.Context, priorityLevelConfiguration *v1beta2.PriorityLevelConfiguration, opts v1.UpdateOptions) (result *v1beta2.PriorityLevelConfiguration, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateAction(prioritylevelconfigurationsResource, priorityLevelConfiguration), &v1beta2.PriorityLevelConfiguration{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta2.PriorityLevelConfiguration), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakePriorityLevelConfigurations) UpdateStatus(ctx context.Context, priorityLevelConfiguration *v1beta2.PriorityLevelConfiguration, opts v1.UpdateOptions) (*v1beta2.PriorityLevelConfiguration, error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateSubresourceAction(prioritylevelconfigurationsResource, "status", priorityLevelConfiguration), &v1beta2.PriorityLevelConfiguration{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta2.PriorityLevelConfiguration), err
}

// Delete takes name of the priorityLevelConfiguration and deletes it. Returns an error if one occurs.
func (c *FakePriorityLevelConfigurations) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteAction(prioritylevelconfigurationsResource, name), &v1beta2.PriorityLevelConfiguration{})
	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakePriorityLevelConfigurations) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewRootDeleteCollectionAction(prioritylevelconfigurationsResource, listOpts)

	_, err := c.Fake.Invokes(action, &v1beta2.PriorityLevelConfigurationList{})
	return err
}

// Patch applies the patch and returns the patched priorityLevelConfiguration.
func (c *FakePriorityLevelConfigurations) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1beta2.PriorityLevelConfiguration, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(prioritylevelconfigurationsResource, name, pt, data, subresources...), &v1beta2.PriorityLevelConfiguration{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta2.PriorityLevelConfiguration), err
}

// Apply takes the given apply declarative configuration, applies it and returns the applied priorityLevelConfiguration.
func (c *FakePriorityLevelConfigurations) Apply(ctx context.Context, priorityLevelConfiguration *flowcontrolv1beta2.PriorityLevelConfigurationApplyConfiguration, opts v1.ApplyOptions) (result *v1beta2.PriorityLevelConfiguration, err error) {
	if priorityLevelConfiguration == nil {
		return nil, fmt.Errorf("priorityLevelConfiguration provided to Apply must not be nil")
	}
	data, err := json.Marshal(priorityLevelConfiguration)
	if err != nil {
		return nil, err
	}
	name := priorityLevelConfiguration.Name
	if name == nil {
		return nil, fmt.Errorf("priorityLevelConfiguration.Name must be provided to Apply")
	}
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(prioritylevelconfigurationsResource, *name, types.ApplyPatchType, data), &v1beta2.PriorityLevelConfiguration{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta2.PriorityLevelConfiguration), err
}

// ApplyStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating ApplyStatus().
func (c *FakePriorityLevelConfigurations) ApplyStatus(ctx context.Context, priorityLevelConfiguration *flowcontrolv1beta2.PriorityLevelConfigurationApplyConfiguration, opts v1.ApplyOptions) (result *v1beta2.PriorityLevelConfiguration, err error) {
	if priorityLevelConfiguration == nil {
		return nil, fmt.Errorf("priorityLevelConfiguration provided to Apply must not be nil")
	}
	data, err := json.Marshal(priorityLevelConfiguration)
	if err != nil {
		return nil, err
	}
	name := priorityLevelConfiguration.Name
	if name == nil {
		return nil, fmt.Errorf("priorityLevelConfiguration.Name must be provided to Apply")
	}
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(prioritylevelconfigurationsResource, *name, types.ApplyPatchType, data, "status"), &v1beta2.PriorityLevelConfiguration{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta2.PriorityLevelConfiguration), err
}
