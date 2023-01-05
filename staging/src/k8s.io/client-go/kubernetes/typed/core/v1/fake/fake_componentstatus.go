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

	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	corev1 "k8s.io/client-go/applyconfigurations/core/v1"
	testing "k8s.io/client-go/testing"
)

// FakeComponentStatuses implements ComponentStatusInterface
type FakeComponentStatuses struct {
	Fake *FakeCoreV1
}

var componentstatusesResource = v1.SchemeGroupVersion.WithResource("componentstatuses")

var componentstatusesKind = v1.SchemeGroupVersion.WithKind("ComponentStatus")

// Get takes name of the componentStatus, and returns the corresponding componentStatus object, and an error if there is any.
func (c *FakeComponentStatuses) Get(ctx context.Context, name string, options metav1.GetOptions) (result *v1.ComponentStatus, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootGetAction(componentstatusesResource, name), &v1.ComponentStatus{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1.ComponentStatus), err
}

// List takes label and field selectors, and returns the list of ComponentStatuses that match those selectors.
func (c *FakeComponentStatuses) List(ctx context.Context, opts metav1.ListOptions) (result *v1.ComponentStatusList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootListAction(componentstatusesResource, componentstatusesKind, opts), &v1.ComponentStatusList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1.ComponentStatusList{ListMeta: obj.(*v1.ComponentStatusList).ListMeta}
	for _, item := range obj.(*v1.ComponentStatusList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested componentStatuses.
func (c *FakeComponentStatuses) Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchAction(componentstatusesResource, opts))
}

// Create takes the representation of a componentStatus and creates it.  Returns the server's representation of the componentStatus, and an error, if there is any.
func (c *FakeComponentStatuses) Create(ctx context.Context, componentStatus *v1.ComponentStatus, opts metav1.CreateOptions) (result *v1.ComponentStatus, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateAction(componentstatusesResource, componentStatus), &v1.ComponentStatus{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1.ComponentStatus), err
}

// Update takes the representation of a componentStatus and updates it. Returns the server's representation of the componentStatus, and an error, if there is any.
func (c *FakeComponentStatuses) Update(ctx context.Context, componentStatus *v1.ComponentStatus, opts metav1.UpdateOptions) (result *v1.ComponentStatus, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateAction(componentstatusesResource, componentStatus), &v1.ComponentStatus{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1.ComponentStatus), err
}

// Delete takes name of the componentStatus and deletes it. Returns an error if one occurs.
func (c *FakeComponentStatuses) Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteActionWithOptions(componentstatusesResource, name, opts), &v1.ComponentStatus{})
	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeComponentStatuses) DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error {
	action := testing.NewRootDeleteCollectionAction(componentstatusesResource, listOpts)

	_, err := c.Fake.Invokes(action, &v1.ComponentStatusList{})
	return err
}

// Patch applies the patch and returns the patched componentStatus.
func (c *FakeComponentStatuses) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *v1.ComponentStatus, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(componentstatusesResource, name, pt, data, subresources...), &v1.ComponentStatus{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1.ComponentStatus), err
}

// Apply takes the given apply declarative configuration, applies it and returns the applied componentStatus.
func (c *FakeComponentStatuses) Apply(ctx context.Context, componentStatus *corev1.ComponentStatusApplyConfiguration, opts metav1.ApplyOptions) (result *v1.ComponentStatus, err error) {
	if componentStatus == nil {
		return nil, fmt.Errorf("componentStatus provided to Apply must not be nil")
	}
	data, err := json.Marshal(componentStatus)
	if err != nil {
		return nil, err
	}
	name := componentStatus.Name
	if name == nil {
		return nil, fmt.Errorf("componentStatus.Name must be provided to Apply")
	}
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(componentstatusesResource, *name, types.ApplyPatchType, data), &v1.ComponentStatus{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1.ComponentStatus), err
}
