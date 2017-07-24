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

// FakeSecrets implements SecretInterface
type FakeSecrets struct {
	Fake *FakeCoreV1
	ns   string
}

var secretsResource = schema.GroupVersionResource{Group: "", Version: "v1", Resource: "secrets"}

var secretsKind = schema.GroupVersionKind{Group: "", Version: "v1", Kind: "Secret"}

// Get takes name of the secret, and returns the corresponding secret object, and an error if there is any.
func (c *FakeSecrets) Get(name string, options v1.GetOptions) (result *core_v1.Secret, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(secretsResource, c.ns, name), &core_v1.Secret{})

	if obj == nil {
		return nil, err
	}
	return obj.(*core_v1.Secret), err
}

// List takes label and field selectors, and returns the list of Secrets that match those selectors.
func (c *FakeSecrets) List(opts v1.ListOptions) (result *core_v1.SecretList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(secretsResource, secretsKind, c.ns, opts), &core_v1.SecretList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &core_v1.SecretList{}
	for _, item := range obj.(*core_v1.SecretList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested secrets.
func (c *FakeSecrets) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(secretsResource, c.ns, opts))

}

// Create takes the representation of a secret and creates it.  Returns the server's representation of the secret, and an error, if there is any.
func (c *FakeSecrets) Create(secret *core_v1.Secret) (result *core_v1.Secret, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(secretsResource, c.ns, secret), &core_v1.Secret{})

	if obj == nil {
		return nil, err
	}
	return obj.(*core_v1.Secret), err
}

// Update takes the representation of a secret and updates it. Returns the server's representation of the secret, and an error, if there is any.
func (c *FakeSecrets) Update(secret *core_v1.Secret) (result *core_v1.Secret, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(secretsResource, c.ns, secret), &core_v1.Secret{})

	if obj == nil {
		return nil, err
	}
	return obj.(*core_v1.Secret), err
}

// Delete takes name of the secret and deletes it. Returns an error if one occurs.
func (c *FakeSecrets) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(secretsResource, c.ns, name), &core_v1.Secret{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeSecrets) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(secretsResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &core_v1.SecretList{})
	return err
}

// Patch applies the patch and returns the patched secret.
func (c *FakeSecrets) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *core_v1.Secret, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(secretsResource, c.ns, name, data, subresources...), &core_v1.Secret{})

	if obj == nil {
		return nil, err
	}
	return obj.(*core_v1.Secret), err
}
