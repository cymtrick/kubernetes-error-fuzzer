/*
Copyright 2018 The Kubernetes Authors.

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
	core "k8s.io/kubernetes/pkg/apis/core"
)

// FakeEndpoints implements EndpointsInterface
type FakeEndpoints struct {
	Fake *FakeCore
	ns   string
}

var endpointsResource = schema.GroupVersionResource{Group: "", Version: "", Resource: "endpoints"}

var endpointsKind = schema.GroupVersionKind{Group: "", Version: "", Kind: "Endpoints"}

// Get takes name of the endpoints, and returns the corresponding endpoints object, and an error if there is any.
func (c *FakeEndpoints) Get(name string, options v1.GetOptions) (result *core.Endpoints, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(endpointsResource, c.ns, name), &core.Endpoints{})

	if obj == nil {
		return nil, err
	}
	return obj.(*core.Endpoints), err
}

// List takes label and field selectors, and returns the list of Endpoints that match those selectors.
func (c *FakeEndpoints) List(opts v1.ListOptions) (result *core.EndpointsList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(endpointsResource, endpointsKind, c.ns, opts), &core.EndpointsList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &core.EndpointsList{}
	for _, item := range obj.(*core.EndpointsList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested endpoints.
func (c *FakeEndpoints) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(endpointsResource, c.ns, opts))

}

// Create takes the representation of a endpoints and creates it.  Returns the server's representation of the endpoints, and an error, if there is any.
func (c *FakeEndpoints) Create(endpoints *core.Endpoints) (result *core.Endpoints, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(endpointsResource, c.ns, endpoints), &core.Endpoints{})

	if obj == nil {
		return nil, err
	}
	return obj.(*core.Endpoints), err
}

// Update takes the representation of a endpoints and updates it. Returns the server's representation of the endpoints, and an error, if there is any.
func (c *FakeEndpoints) Update(endpoints *core.Endpoints) (result *core.Endpoints, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(endpointsResource, c.ns, endpoints), &core.Endpoints{})

	if obj == nil {
		return nil, err
	}
	return obj.(*core.Endpoints), err
}

// Delete takes name of the endpoints and deletes it. Returns an error if one occurs.
func (c *FakeEndpoints) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(endpointsResource, c.ns, name), &core.Endpoints{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeEndpoints) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(endpointsResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &core.EndpointsList{})
	return err
}

// Patch applies the patch and returns the patched endpoints.
func (c *FakeEndpoints) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *core.Endpoints, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(endpointsResource, c.ns, name, data, subresources...), &core.Endpoints{})

	if obj == nil {
		return nil, err
	}
	return obj.(*core.Endpoints), err
}
