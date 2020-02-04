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

	authenticationv1 "k8s.io/api/authentication/v1"
	corev1 "k8s.io/api/core/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeServiceAccounts implements ServiceAccountInterface
type FakeServiceAccounts struct {
	Fake *FakeCoreV1
	ns   string
}

var serviceaccountsResource = schema.GroupVersionResource{Group: "", Version: "v1", Resource: "serviceaccounts"}

var serviceaccountsKind = schema.GroupVersionKind{Group: "", Version: "v1", Kind: "ServiceAccount"}

// Get takes name of the serviceAccount, and returns the corresponding serviceAccount object, and an error if there is any.
func (c *FakeServiceAccounts) Get(ctx context.Context, name string, options v1.GetOptions) (result *corev1.ServiceAccount, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(serviceaccountsResource, c.ns, name), &corev1.ServiceAccount{})

	if obj == nil {
		return nil, err
	}
	return obj.(*corev1.ServiceAccount), err
}

// List takes label and field selectors, and returns the list of ServiceAccounts that match those selectors.
func (c *FakeServiceAccounts) List(ctx context.Context, opts v1.ListOptions) (result *corev1.ServiceAccountList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(serviceaccountsResource, serviceaccountsKind, c.ns, opts), &corev1.ServiceAccountList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &corev1.ServiceAccountList{ListMeta: obj.(*corev1.ServiceAccountList).ListMeta}
	for _, item := range obj.(*corev1.ServiceAccountList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested serviceAccounts.
func (c *FakeServiceAccounts) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(serviceaccountsResource, c.ns, opts))

}

// Create takes the representation of a serviceAccount and creates it.  Returns the server's representation of the serviceAccount, and an error, if there is any.
func (c *FakeServiceAccounts) Create(ctx context.Context, serviceAccount *corev1.ServiceAccount) (result *corev1.ServiceAccount, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(serviceaccountsResource, c.ns, serviceAccount), &corev1.ServiceAccount{})

	if obj == nil {
		return nil, err
	}
	return obj.(*corev1.ServiceAccount), err
}

// Update takes the representation of a serviceAccount and updates it. Returns the server's representation of the serviceAccount, and an error, if there is any.
func (c *FakeServiceAccounts) Update(ctx context.Context, serviceAccount *corev1.ServiceAccount) (result *corev1.ServiceAccount, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(serviceaccountsResource, c.ns, serviceAccount), &corev1.ServiceAccount{})

	if obj == nil {
		return nil, err
	}
	return obj.(*corev1.ServiceAccount), err
}

// Delete takes name of the serviceAccount and deletes it. Returns an error if one occurs.
func (c *FakeServiceAccounts) Delete(ctx context.Context, name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(serviceaccountsResource, c.ns, name), &corev1.ServiceAccount{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeServiceAccounts) DeleteCollection(ctx context.Context, options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(serviceaccountsResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &corev1.ServiceAccountList{})
	return err
}

// Patch applies the patch and returns the patched serviceAccount.
func (c *FakeServiceAccounts) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, subresources ...string) (result *corev1.ServiceAccount, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(serviceaccountsResource, c.ns, name, pt, data, subresources...), &corev1.ServiceAccount{})

	if obj == nil {
		return nil, err
	}
	return obj.(*corev1.ServiceAccount), err
}

// CreateToken takes the representation of a tokenRequest and creates it.  Returns the server's representation of the tokenRequest, and an error, if there is any.
func (c *FakeServiceAccounts) CreateToken(ctx context.Context, serviceAccountName string, tokenRequest *authenticationv1.TokenRequest) (result *authenticationv1.TokenRequest, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateSubresourceAction(serviceaccountsResource, serviceAccountName, "token", c.ns, tokenRequest), &authenticationv1.TokenRequest{})

	if obj == nil {
		return nil, err
	}
	return obj.(*authenticationv1.TokenRequest), err
}
