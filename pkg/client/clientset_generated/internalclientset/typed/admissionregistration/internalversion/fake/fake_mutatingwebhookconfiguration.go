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
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
	admissionregistration "k8s.io/kubernetes/pkg/apis/admissionregistration"
)

// FakeMutatingWebhookConfigurations implements MutatingWebhookConfigurationInterface
type FakeMutatingWebhookConfigurations struct {
	Fake *FakeAdmissionregistration
}

var mutatingwebhookconfigurationsResource = schema.GroupVersionResource{Group: "admissionregistration.k8s.io", Version: "", Resource: "mutatingwebhookconfigurations"}

var mutatingwebhookconfigurationsKind = schema.GroupVersionKind{Group: "admissionregistration.k8s.io", Version: "", Kind: "MutatingWebhookConfiguration"}

// Get takes name of the mutatingWebhookConfiguration, and returns the corresponding mutatingWebhookConfiguration object, and an error if there is any.
func (c *FakeMutatingWebhookConfigurations) Get(name string, options v1.GetOptions) (result *admissionregistration.MutatingWebhookConfiguration, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootGetAction(mutatingwebhookconfigurationsResource, name), &admissionregistration.MutatingWebhookConfiguration{})
	if obj == nil {
		return nil, err
	}
	return obj.(*admissionregistration.MutatingWebhookConfiguration), err
}

// List takes label and field selectors, and returns the list of MutatingWebhookConfigurations that match those selectors.
func (c *FakeMutatingWebhookConfigurations) List(opts v1.ListOptions) (result *admissionregistration.MutatingWebhookConfigurationList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootListAction(mutatingwebhookconfigurationsResource, mutatingwebhookconfigurationsKind, opts), &admissionregistration.MutatingWebhookConfigurationList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &admissionregistration.MutatingWebhookConfigurationList{}
	for _, item := range obj.(*admissionregistration.MutatingWebhookConfigurationList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested mutatingWebhookConfigurations.
func (c *FakeMutatingWebhookConfigurations) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchAction(mutatingwebhookconfigurationsResource, opts))
}

// Create takes the representation of a mutatingWebhookConfiguration and creates it.  Returns the server's representation of the mutatingWebhookConfiguration, and an error, if there is any.
func (c *FakeMutatingWebhookConfigurations) Create(mutatingWebhookConfiguration *admissionregistration.MutatingWebhookConfiguration) (result *admissionregistration.MutatingWebhookConfiguration, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateAction(mutatingwebhookconfigurationsResource, mutatingWebhookConfiguration), &admissionregistration.MutatingWebhookConfiguration{})
	if obj == nil {
		return nil, err
	}
	return obj.(*admissionregistration.MutatingWebhookConfiguration), err
}

// Update takes the representation of a mutatingWebhookConfiguration and updates it. Returns the server's representation of the mutatingWebhookConfiguration, and an error, if there is any.
func (c *FakeMutatingWebhookConfigurations) Update(mutatingWebhookConfiguration *admissionregistration.MutatingWebhookConfiguration) (result *admissionregistration.MutatingWebhookConfiguration, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateAction(mutatingwebhookconfigurationsResource, mutatingWebhookConfiguration), &admissionregistration.MutatingWebhookConfiguration{})
	if obj == nil {
		return nil, err
	}
	return obj.(*admissionregistration.MutatingWebhookConfiguration), err
}

// Delete takes name of the mutatingWebhookConfiguration and deletes it. Returns an error if one occurs.
func (c *FakeMutatingWebhookConfigurations) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteAction(mutatingwebhookconfigurationsResource, name), &admissionregistration.MutatingWebhookConfiguration{})
	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeMutatingWebhookConfigurations) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewRootDeleteCollectionAction(mutatingwebhookconfigurationsResource, listOptions)

	_, err := c.Fake.Invokes(action, &admissionregistration.MutatingWebhookConfigurationList{})
	return err
}

// Patch applies the patch and returns the patched mutatingWebhookConfiguration.
func (c *FakeMutatingWebhookConfigurations) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *admissionregistration.MutatingWebhookConfiguration, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(mutatingwebhookconfigurationsResource, name, data, subresources...), &admissionregistration.MutatingWebhookConfiguration{})
	if obj == nil {
		return nil, err
	}
	return obj.(*admissionregistration.MutatingWebhookConfiguration), err
}
