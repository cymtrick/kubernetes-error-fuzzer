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

	v1alpha1 "k8s.io/api/auditregistration/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeAuditSinks implements AuditSinkInterface
type FakeAuditSinks struct {
	Fake *FakeAuditregistrationV1alpha1
}

var auditsinksResource = schema.GroupVersionResource{Group: "auditregistration.k8s.io", Version: "v1alpha1", Resource: "auditsinks"}

var auditsinksKind = schema.GroupVersionKind{Group: "auditregistration.k8s.io", Version: "v1alpha1", Kind: "AuditSink"}

// Get takes name of the auditSink, and returns the corresponding auditSink object, and an error if there is any.
func (c *FakeAuditSinks) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.AuditSink, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootGetAction(auditsinksResource, name), &v1alpha1.AuditSink{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AuditSink), err
}

// List takes label and field selectors, and returns the list of AuditSinks that match those selectors.
func (c *FakeAuditSinks) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.AuditSinkList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootListAction(auditsinksResource, auditsinksKind, opts), &v1alpha1.AuditSinkList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.AuditSinkList{ListMeta: obj.(*v1alpha1.AuditSinkList).ListMeta}
	for _, item := range obj.(*v1alpha1.AuditSinkList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested auditSinks.
func (c *FakeAuditSinks) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchAction(auditsinksResource, opts))
}

// Create takes the representation of a auditSink and creates it.  Returns the server's representation of the auditSink, and an error, if there is any.
func (c *FakeAuditSinks) Create(ctx context.Context, auditSink *v1alpha1.AuditSink, opts v1.CreateOptions) (result *v1alpha1.AuditSink, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateAction(auditsinksResource, auditSink), &v1alpha1.AuditSink{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AuditSink), err
}

// Update takes the representation of a auditSink and updates it. Returns the server's representation of the auditSink, and an error, if there is any.
func (c *FakeAuditSinks) Update(ctx context.Context, auditSink *v1alpha1.AuditSink, opts v1.UpdateOptions) (result *v1alpha1.AuditSink, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateAction(auditsinksResource, auditSink), &v1alpha1.AuditSink{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AuditSink), err
}

// Delete takes name of the auditSink and deletes it. Returns an error if one occurs.
func (c *FakeAuditSinks) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteAction(auditsinksResource, name), &v1alpha1.AuditSink{})
	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeAuditSinks) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewRootDeleteCollectionAction(auditsinksResource, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.AuditSinkList{})
	return err
}

// Patch applies the patch and returns the patched auditSink.
func (c *FakeAuditSinks) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.AuditSink, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(auditsinksResource, name, pt, data, subresources...), &v1alpha1.AuditSink{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AuditSink), err
}
