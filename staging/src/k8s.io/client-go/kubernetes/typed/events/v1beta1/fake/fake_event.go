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

	v1beta1 "k8s.io/api/events/v1beta1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	eventsv1beta1 "k8s.io/client-go/applyconfigurations/events/v1beta1"
	testing "k8s.io/client-go/testing"
)

// FakeEvents implements EventInterface
type FakeEvents struct {
	Fake *FakeEventsV1beta1
	ns   string
}

var eventsResource = v1beta1.SchemeGroupVersion.WithResource("events")

var eventsKind = v1beta1.SchemeGroupVersion.WithKind("Event")

// Get takes name of the event, and returns the corresponding event object, and an error if there is any.
func (c *FakeEvents) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1beta1.Event, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(eventsResource, c.ns, name), &v1beta1.Event{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.Event), err
}

// List takes label and field selectors, and returns the list of Events that match those selectors.
func (c *FakeEvents) List(ctx context.Context, opts v1.ListOptions) (result *v1beta1.EventList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(eventsResource, eventsKind, c.ns, opts), &v1beta1.EventList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1beta1.EventList{ListMeta: obj.(*v1beta1.EventList).ListMeta}
	for _, item := range obj.(*v1beta1.EventList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested events.
func (c *FakeEvents) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(eventsResource, c.ns, opts))

}

// Create takes the representation of a event and creates it.  Returns the server's representation of the event, and an error, if there is any.
func (c *FakeEvents) Create(ctx context.Context, event *v1beta1.Event, opts v1.CreateOptions) (result *v1beta1.Event, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(eventsResource, c.ns, event), &v1beta1.Event{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.Event), err
}

// Update takes the representation of a event and updates it. Returns the server's representation of the event, and an error, if there is any.
func (c *FakeEvents) Update(ctx context.Context, event *v1beta1.Event, opts v1.UpdateOptions) (result *v1beta1.Event, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(eventsResource, c.ns, event), &v1beta1.Event{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.Event), err
}

// Delete takes name of the event and deletes it. Returns an error if one occurs.
func (c *FakeEvents) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteActionWithOptions(eventsResource, c.ns, name, opts), &v1beta1.Event{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeEvents) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(eventsResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1beta1.EventList{})
	return err
}

// Patch applies the patch and returns the patched event.
func (c *FakeEvents) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1beta1.Event, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(eventsResource, c.ns, name, pt, data, subresources...), &v1beta1.Event{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.Event), err
}

// Apply takes the given apply declarative configuration, applies it and returns the applied event.
func (c *FakeEvents) Apply(ctx context.Context, event *eventsv1beta1.EventApplyConfiguration, opts v1.ApplyOptions) (result *v1beta1.Event, err error) {
	if event == nil {
		return nil, fmt.Errorf("event provided to Apply must not be nil")
	}
	data, err := json.Marshal(event)
	if err != nil {
		return nil, err
	}
	name := event.Name
	if name == nil {
		return nil, fmt.Errorf("event.Name must be provided to Apply")
	}
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(eventsResource, c.ns, *name, types.ApplyPatchType, data), &v1beta1.Event{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.Event), err
}
