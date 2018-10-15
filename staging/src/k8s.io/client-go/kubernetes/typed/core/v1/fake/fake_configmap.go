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
	corev1 "k8s.io/api/core/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeConfigMaps implements ConfigMapInterface
type FakeConfigMaps struct {
	Fake *FakeCoreV1
	ns   string
}

var configmapsResource = schema.GroupVersionResource{Group: "", Version: "v1", Resource: "configmaps"}

var configmapsKind = schema.GroupVersionKind{Group: "", Version: "v1", Kind: "ConfigMap"}

// Get takes name of the configMap, and returns the corresponding configMap object, and an error if there is any.
func (c *FakeConfigMaps) Get(name string, options v1.GetOptions) (result *corev1.ConfigMap, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(configmapsResource, c.ns, name), &corev1.ConfigMap{})

	if obj == nil {
		return nil, err
	}
	return obj.(*corev1.ConfigMap), err
}

// List takes label and field selectors, and returns the list of ConfigMaps that match those selectors.
func (c *FakeConfigMaps) List(opts v1.ListOptions) (result *corev1.ConfigMapList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(configmapsResource, configmapsKind, c.ns, opts), &corev1.ConfigMapList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &corev1.ConfigMapList{ListMeta: obj.(*corev1.ConfigMapList).ListMeta}
	for _, item := range obj.(*corev1.ConfigMapList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested configMaps.
func (c *FakeConfigMaps) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(configmapsResource, c.ns, opts))

}

// Create takes the representation of a configMap and creates it.  Returns the server's representation of the configMap, and an error, if there is any.
func (c *FakeConfigMaps) Create(configMap *corev1.ConfigMap) (result *corev1.ConfigMap, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(configmapsResource, c.ns, configMap), &corev1.ConfigMap{})

	if obj == nil {
		return nil, err
	}
	return obj.(*corev1.ConfigMap), err
}

// Update takes the representation of a configMap and updates it. Returns the server's representation of the configMap, and an error, if there is any.
func (c *FakeConfigMaps) Update(configMap *corev1.ConfigMap) (result *corev1.ConfigMap, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(configmapsResource, c.ns, configMap), &corev1.ConfigMap{})

	if obj == nil {
		return nil, err
	}
	return obj.(*corev1.ConfigMap), err
}

// Delete takes name of the configMap and deletes it. Returns an error if one occurs.
func (c *FakeConfigMaps) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(configmapsResource, c.ns, name), &corev1.ConfigMap{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeConfigMaps) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(configmapsResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &corev1.ConfigMapList{})
	return err
}

// Patch applies the patch and returns the patched configMap.
func (c *FakeConfigMaps) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *corev1.ConfigMap, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(configmapsResource, c.ns, name, pt, data, subresources...), &corev1.ConfigMap{})

	if obj == nil {
		return nil, err
	}
	return obj.(*corev1.ConfigMap), err
}
