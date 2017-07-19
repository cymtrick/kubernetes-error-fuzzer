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
	core_v1 "k8s.io/api/core/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeNamespaces implements NamespaceInterface
type FakeNamespaces struct {
	Fake *FakeCoreV1
}

var namespacesResource = schema.GroupVersionResource{Group: "", Version: "v1", Resource: "namespaces"}

var namespacesKind = schema.GroupVersionKind{Group: "", Version: "v1", Kind: "Namespace"}

func (c *FakeNamespaces) Get(name string, options v1.GetOptions) (result *core_v1.Namespace, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootGetAction(namespacesResource, name), &core_v1.Namespace{})
	if obj == nil {
		return nil, err
	}
	return obj.(*core_v1.Namespace), err
}

func (c *FakeNamespaces) List(opts v1.ListOptions) (result *core_v1.NamespaceList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootListAction(namespacesResource, namespacesKind, opts), &core_v1.NamespaceList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &core_v1.NamespaceList{}
	for _, item := range obj.(*core_v1.NamespaceList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested namespaces.
func (c *FakeNamespaces) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchAction(namespacesResource, opts))
}

func (c *FakeNamespaces) Create(namespace *core_v1.Namespace) (result *core_v1.Namespace, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateAction(namespacesResource, namespace), &core_v1.Namespace{})
	if obj == nil {
		return nil, err
	}
	return obj.(*core_v1.Namespace), err
}

func (c *FakeNamespaces) Update(namespace *core_v1.Namespace) (result *core_v1.Namespace, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateAction(namespacesResource, namespace), &core_v1.Namespace{})
	if obj == nil {
		return nil, err
	}
	return obj.(*core_v1.Namespace), err
}

func (c *FakeNamespaces) UpdateStatus(namespace *core_v1.Namespace) (*core_v1.Namespace, error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateSubresourceAction(namespacesResource, "status", namespace), &core_v1.Namespace{})
	if obj == nil {
		return nil, err
	}
	return obj.(*core_v1.Namespace), err
}

func (c *FakeNamespaces) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteAction(namespacesResource, name), &core_v1.Namespace{})
	return err
}

func (c *FakeNamespaces) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewRootDeleteCollectionAction(namespacesResource, listOptions)

	_, err := c.Fake.Invokes(action, &core_v1.NamespaceList{})
	return err
}

// Patch applies the patch and returns the patched namespace.
func (c *FakeNamespaces) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *core_v1.Namespace, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(namespacesResource, name, data, subresources...), &core_v1.Namespace{})
	if obj == nil {
		return nil, err
	}
	return obj.(*core_v1.Namespace), err
}
