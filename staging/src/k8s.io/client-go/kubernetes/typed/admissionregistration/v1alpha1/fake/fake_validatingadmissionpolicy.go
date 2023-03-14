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

	v1alpha1 "k8s.io/api/admissionregistration/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	admissionregistrationv1alpha1 "k8s.io/client-go/applyconfigurations/admissionregistration/v1alpha1"
	testing "k8s.io/client-go/testing"
)

// FakeValidatingAdmissionPolicies implements ValidatingAdmissionPolicyInterface
type FakeValidatingAdmissionPolicies struct {
	Fake *FakeAdmissionregistrationV1alpha1
}

var validatingadmissionpoliciesResource = v1alpha1.SchemeGroupVersion.WithResource("validatingadmissionpolicies")

var validatingadmissionpoliciesKind = v1alpha1.SchemeGroupVersion.WithKind("ValidatingAdmissionPolicy")

// Get takes name of the validatingAdmissionPolicy, and returns the corresponding validatingAdmissionPolicy object, and an error if there is any.
func (c *FakeValidatingAdmissionPolicies) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.ValidatingAdmissionPolicy, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootGetAction(validatingadmissionpoliciesResource, name), &v1alpha1.ValidatingAdmissionPolicy{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ValidatingAdmissionPolicy), err
}

// List takes label and field selectors, and returns the list of ValidatingAdmissionPolicies that match those selectors.
func (c *FakeValidatingAdmissionPolicies) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.ValidatingAdmissionPolicyList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootListAction(validatingadmissionpoliciesResource, validatingadmissionpoliciesKind, opts), &v1alpha1.ValidatingAdmissionPolicyList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.ValidatingAdmissionPolicyList{ListMeta: obj.(*v1alpha1.ValidatingAdmissionPolicyList).ListMeta}
	for _, item := range obj.(*v1alpha1.ValidatingAdmissionPolicyList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested validatingAdmissionPolicies.
func (c *FakeValidatingAdmissionPolicies) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchAction(validatingadmissionpoliciesResource, opts))
}

// Create takes the representation of a validatingAdmissionPolicy and creates it.  Returns the server's representation of the validatingAdmissionPolicy, and an error, if there is any.
func (c *FakeValidatingAdmissionPolicies) Create(ctx context.Context, validatingAdmissionPolicy *v1alpha1.ValidatingAdmissionPolicy, opts v1.CreateOptions) (result *v1alpha1.ValidatingAdmissionPolicy, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateAction(validatingadmissionpoliciesResource, validatingAdmissionPolicy), &v1alpha1.ValidatingAdmissionPolicy{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ValidatingAdmissionPolicy), err
}

// Update takes the representation of a validatingAdmissionPolicy and updates it. Returns the server's representation of the validatingAdmissionPolicy, and an error, if there is any.
func (c *FakeValidatingAdmissionPolicies) Update(ctx context.Context, validatingAdmissionPolicy *v1alpha1.ValidatingAdmissionPolicy, opts v1.UpdateOptions) (result *v1alpha1.ValidatingAdmissionPolicy, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateAction(validatingadmissionpoliciesResource, validatingAdmissionPolicy), &v1alpha1.ValidatingAdmissionPolicy{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ValidatingAdmissionPolicy), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeValidatingAdmissionPolicies) UpdateStatus(ctx context.Context, validatingAdmissionPolicy *v1alpha1.ValidatingAdmissionPolicy, opts v1.UpdateOptions) (*v1alpha1.ValidatingAdmissionPolicy, error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateSubresourceAction(validatingadmissionpoliciesResource, "status", validatingAdmissionPolicy), &v1alpha1.ValidatingAdmissionPolicy{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ValidatingAdmissionPolicy), err
}

// Delete takes name of the validatingAdmissionPolicy and deletes it. Returns an error if one occurs.
func (c *FakeValidatingAdmissionPolicies) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteActionWithOptions(validatingadmissionpoliciesResource, name, opts), &v1alpha1.ValidatingAdmissionPolicy{})
	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeValidatingAdmissionPolicies) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewRootDeleteCollectionAction(validatingadmissionpoliciesResource, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.ValidatingAdmissionPolicyList{})
	return err
}

// Patch applies the patch and returns the patched validatingAdmissionPolicy.
func (c *FakeValidatingAdmissionPolicies) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.ValidatingAdmissionPolicy, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(validatingadmissionpoliciesResource, name, pt, data, subresources...), &v1alpha1.ValidatingAdmissionPolicy{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ValidatingAdmissionPolicy), err
}

// Apply takes the given apply declarative configuration, applies it and returns the applied validatingAdmissionPolicy.
func (c *FakeValidatingAdmissionPolicies) Apply(ctx context.Context, validatingAdmissionPolicy *admissionregistrationv1alpha1.ValidatingAdmissionPolicyApplyConfiguration, opts v1.ApplyOptions) (result *v1alpha1.ValidatingAdmissionPolicy, err error) {
	if validatingAdmissionPolicy == nil {
		return nil, fmt.Errorf("validatingAdmissionPolicy provided to Apply must not be nil")
	}
	data, err := json.Marshal(validatingAdmissionPolicy)
	if err != nil {
		return nil, err
	}
	name := validatingAdmissionPolicy.Name
	if name == nil {
		return nil, fmt.Errorf("validatingAdmissionPolicy.Name must be provided to Apply")
	}
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(validatingadmissionpoliciesResource, *name, types.ApplyPatchType, data), &v1alpha1.ValidatingAdmissionPolicy{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ValidatingAdmissionPolicy), err
}

// ApplyStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating ApplyStatus().
func (c *FakeValidatingAdmissionPolicies) ApplyStatus(ctx context.Context, validatingAdmissionPolicy *admissionregistrationv1alpha1.ValidatingAdmissionPolicyApplyConfiguration, opts v1.ApplyOptions) (result *v1alpha1.ValidatingAdmissionPolicy, err error) {
	if validatingAdmissionPolicy == nil {
		return nil, fmt.Errorf("validatingAdmissionPolicy provided to Apply must not be nil")
	}
	data, err := json.Marshal(validatingAdmissionPolicy)
	if err != nil {
		return nil, err
	}
	name := validatingAdmissionPolicy.Name
	if name == nil {
		return nil, fmt.Errorf("validatingAdmissionPolicy.Name must be provided to Apply")
	}
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(validatingadmissionpoliciesResource, *name, types.ApplyPatchType, data, "status"), &v1alpha1.ValidatingAdmissionPolicy{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ValidatingAdmissionPolicy), err
}
