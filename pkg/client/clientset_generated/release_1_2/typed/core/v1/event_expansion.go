/*
Copyright 2016 The Kubernetes Authors.

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

package v1

import (
	"fmt"

	"k8s.io/kubernetes/pkg/api"
	"k8s.io/kubernetes/pkg/api/v1"
	"k8s.io/kubernetes/pkg/fields"
	"k8s.io/kubernetes/pkg/runtime"
)

// The EventExpansion interface allows manually adding extra methods to the EventInterface.
type EventExpansion interface {
	// CreateWithEventNamespace is the same as a Create, except that it sends the request to the event.Namespace.
	CreateWithEventNamespace(event *v1.Event) (*v1.Event, error)
	// UpdateWithEventNamespace is the same as a Update, except that it sends the request to the event.Namespace.
	UpdateWithEventNamespace(event *v1.Event) (*v1.Event, error)
	Patch(event *v1.Event, data []byte) (*v1.Event, error)
	// Search finds events about the specified object
	Search(objOrRef runtime.Object) (*v1.EventList, error)
	// Returns the appropriate field selector based on the API version being used to communicate with the server.
	// The returned field selector can be used with List and Watch to filter desired events.
	GetFieldSelector(involvedObjectName, involvedObjectNamespace, involvedObjectKind, involvedObjectUID *string) fields.Selector
}

// CreateWithEventNamespace makes a new event. Returns the copy of the event the server returns,
// or an error. The namespace to create the event within is deduced from the
// event; it must either match this event client's namespace, or this event
// client must have been created with the "" namespace.
func (e *events) CreateWithEventNamespace(event *v1.Event) (*v1.Event, error) {
	if e.ns != "" && event.Namespace != e.ns {
		return nil, fmt.Errorf("can't create an event with namespace '%v' in namespace '%v'", event.Namespace, e.ns)
	}
	result := &v1.Event{}
	err := e.client.Post().
		NamespaceIfScoped(event.Namespace, len(event.Namespace) > 0).
		Resource("events").
		Body(event).
		Do().
		Into(result)
	return result, err
}

// UpdateWithEventNamespace modifies an existing event. It returns the copy of the event that the server returns,
// or an error. The namespace and key to update the event within is deduced from the event. The
// namespace must either match this event client's namespace, or this event client must have been
// created with the "" namespace. Update also requires the ResourceVersion to be set in the event
// object.
func (e *events) UpdateWithEventNamespace(event *v1.Event) (*v1.Event, error) {
	result := &v1.Event{}
	err := e.client.Put().
		NamespaceIfScoped(event.Namespace, len(event.Namespace) > 0).
		Resource("events").
		Name(event.Name).
		Body(event).
		Do().
		Into(result)
	return result, err
}

// Patch modifies an existing event. It returns the copy of the event that the server returns, or an
// error. The namespace and name of the target event is deduced from the incompleteEvent. The
// namespace must either match this event client's namespace, or this event client must have been
// created with the "" namespace.
func (e *events) Patch(incompleteEvent *v1.Event, data []byte) (*v1.Event, error) {
	result := &v1.Event{}
	err := e.client.Patch(api.StrategicMergePatchType).
		NamespaceIfScoped(incompleteEvent.Namespace, len(incompleteEvent.Namespace) > 0).
		Resource("events").
		Name(incompleteEvent.Name).
		Body(data).
		Do().
		Into(result)
	return result, err
}

// Search finds events about the specified object. The namespace of the
// object must match this event's client namespace unless the event client
// was made with the "" namespace.
func (e *events) Search(objOrRef runtime.Object) (*v1.EventList, error) {
	ref, err := api.GetReference(objOrRef)
	if err != nil {
		return nil, err
	}
	if e.ns != "" && ref.Namespace != e.ns {
		return nil, fmt.Errorf("won't be able to find any events of namespace '%v' in namespace '%v'", ref.Namespace, e.ns)
	}
	stringRefKind := string(ref.Kind)
	var refKind *string
	if stringRefKind != "" {
		refKind = &stringRefKind
	}
	stringRefUID := string(ref.UID)
	var refUID *string
	if stringRefUID != "" {
		refUID = &stringRefUID
	}
	fieldSelector := e.GetFieldSelector(&ref.Name, &ref.Namespace, refKind, refUID)
	return e.List(api.ListOptions{FieldSelector: fieldSelector})
}

// Returns the appropriate field selector based on the API version being used to communicate with the server.
// The returned field selector can be used with List and Watch to filter desired events.
func (e *events) GetFieldSelector(involvedObjectName, involvedObjectNamespace, involvedObjectKind, involvedObjectUID *string) fields.Selector {
	apiVersion := e.client.APIVersion().String()
	field := fields.Set{}
	if involvedObjectName != nil {
		field[GetInvolvedObjectNameFieldLabel(apiVersion)] = *involvedObjectName
	}
	if involvedObjectNamespace != nil {
		field["involvedObject.namespace"] = *involvedObjectNamespace
	}
	if involvedObjectKind != nil {
		field["involvedObject.kind"] = *involvedObjectKind
	}
	if involvedObjectUID != nil {
		field["involvedObject.uid"] = *involvedObjectUID
	}
	return field.AsSelector()
}

// Returns the appropriate field label to use for name of the involved object as per the given API version.
func GetInvolvedObjectNameFieldLabel(version string) string {
	return "involvedObject.name"
}

// TODO: This is a temporary arrangement and will be removed once all clients are moved to use the clientset.
type EventSinkImpl struct {
	Interface EventInterface
}

func (e *EventSinkImpl) Create(event *v1.Event) (*v1.Event, error) {
	return e.Interface.CreateWithEventNamespace(event)
}

func (e *EventSinkImpl) Update(event *v1.Event) (*v1.Event, error) {
	return e.Interface.UpdateWithEventNamespace(event)
}

func (e *EventSinkImpl) Patch(event *v1.Event, data []byte) (*v1.Event, error) {
	return e.Interface.Patch(event, data)
}
