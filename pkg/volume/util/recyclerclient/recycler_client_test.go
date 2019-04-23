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

package recyclerclient

import (
	"fmt"
	"testing"

	v1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/watch"
	api "k8s.io/kubernetes/pkg/apis/core"
)

type testcase struct {
	// Input of the test
	name        string
	existingPod *v1.Pod
	createPod   *v1.Pod
	// eventSequence is list of events that are simulated during recycling. It
	// can be either event generated by a recycler pod or a state change of
	// the pod. (see newPodEvent and newEvent below).
	eventSequence []watch.Event

	// Expected output.
	// expectedEvents is list of events that were sent to the volume that was
	// recycled.
	expectedEvents []mockEvent
	expectedError  string
}

func newPodEvent(eventtype watch.EventType, name string, phase v1.PodPhase, message string) watch.Event {
	return watch.Event{
		Type:   eventtype,
		Object: newPod(name, phase, message),
	}
}

func newEvent(eventtype, message string) watch.Event {
	return watch.Event{
		Type: watch.Added,
		Object: &v1.Event{
			ObjectMeta: metav1.ObjectMeta{
				Namespace: metav1.NamespaceDefault,
			},
			Reason:  "MockEvent",
			Message: message,
			Type:    eventtype,
		},
	}
}

func newPod(name string, phase v1.PodPhase, message string) *v1.Pod {
	return &v1.Pod{
		ObjectMeta: metav1.ObjectMeta{
			Namespace: metav1.NamespaceDefault,
			Name:      name,
		},
		Status: v1.PodStatus{
			Phase:   phase,
			Message: message,
		},
	}
}

func TestRecyclerPod(t *testing.T) {
	tests := []testcase{
		{
			// Test recycler success with some events
			name:      "RecyclerSuccess",
			createPod: newPod("podRecyclerSuccess", v1.PodPending, ""),
			eventSequence: []watch.Event{
				// Pod gets Running and Succeeded
				newPodEvent(watch.Added, "podRecyclerSuccess", v1.PodPending, ""),
				newEvent(v1.EventTypeNormal, "Successfully assigned recycler-for-podRecyclerSuccess to 127.0.0.1"),
				newEvent(v1.EventTypeNormal, "Pulling image \"k8s.gcr.io/busybox\""),
				newEvent(v1.EventTypeNormal, "Successfully pulled image \"k8s.gcr.io/busybox\""),
				newEvent(v1.EventTypeNormal, "Created container with docker id 83d929aeac82"),
				newEvent(v1.EventTypeNormal, "Started container with docker id 83d929aeac82"),
				newPodEvent(watch.Modified, "podRecyclerSuccess", v1.PodRunning, ""),
				newPodEvent(watch.Modified, "podRecyclerSuccess", v1.PodSucceeded, ""),
			},
			expectedEvents: []mockEvent{
				{v1.EventTypeNormal, "Successfully assigned recycler-for-podRecyclerSuccess to 127.0.0.1"},
				{v1.EventTypeNormal, "Pulling image \"k8s.gcr.io/busybox\""},
				{v1.EventTypeNormal, "Successfully pulled image \"k8s.gcr.io/busybox\""},
				{v1.EventTypeNormal, "Created container with docker id 83d929aeac82"},
				{v1.EventTypeNormal, "Started container with docker id 83d929aeac82"},
			},
			expectedError: "",
		},
		{
			// Test recycler failure with some events
			name:      "RecyclerFailure",
			createPod: newPod("podRecyclerFailure", v1.PodPending, ""),
			eventSequence: []watch.Event{
				// Pod gets Running and Succeeded
				newPodEvent(watch.Added, "podRecyclerFailure", v1.PodPending, ""),
				newEvent(v1.EventTypeNormal, "Successfully assigned recycler-for-podRecyclerFailure to 127.0.0.1"),
				newEvent(v1.EventTypeWarning, "Unable to mount volumes for pod \"recycler-for-podRecyclerFailure_default(3c9809e5-347c-11e6-a79b-3c970e965218)\": timeout expired waiting for volumes to attach/mount"),
				newEvent(v1.EventTypeWarning, "Error syncing pod, skipping: timeout expired waiting for volumes to attach/mount for pod \"default\"/\"recycler-for-podRecyclerFailure\". list of unattached/unmounted"),
				newPodEvent(watch.Modified, "podRecyclerFailure", v1.PodRunning, ""),
				newPodEvent(watch.Modified, "podRecyclerFailure", v1.PodFailed, "Pod was active on the node longer than specified deadline"),
			},
			expectedEvents: []mockEvent{
				{v1.EventTypeNormal, "Successfully assigned recycler-for-podRecyclerFailure to 127.0.0.1"},
				{v1.EventTypeWarning, "Unable to mount volumes for pod \"recycler-for-podRecyclerFailure_default(3c9809e5-347c-11e6-a79b-3c970e965218)\": timeout expired waiting for volumes to attach/mount"},
				{v1.EventTypeWarning, "Error syncing pod, skipping: timeout expired waiting for volumes to attach/mount for pod \"default\"/\"recycler-for-podRecyclerFailure\". list of unattached/unmounted"},
			},
			expectedError: "failed to recycle volume: Pod was active on the node longer than specified deadline",
		},
		{
			// Recycler pod gets deleted
			name:      "RecyclerDeleted",
			createPod: newPod("podRecyclerDeleted", v1.PodPending, ""),
			eventSequence: []watch.Event{
				// Pod gets Running and Succeeded
				newPodEvent(watch.Added, "podRecyclerDeleted", v1.PodPending, ""),
				newEvent(v1.EventTypeNormal, "Successfully assigned recycler-for-podRecyclerDeleted to 127.0.0.1"),
				newPodEvent(watch.Deleted, "podRecyclerDeleted", v1.PodPending, ""),
			},
			expectedEvents: []mockEvent{
				{v1.EventTypeNormal, "Successfully assigned recycler-for-podRecyclerDeleted to 127.0.0.1"},
			},
			expectedError: "failed to recycle volume: recycler pod was deleted",
		},
		{
			// Another recycler pod is already running
			name:          "RecyclerRunning",
			existingPod:   newPod("podOldRecycler", v1.PodRunning, ""),
			createPod:     newPod("podNewRecycler", v1.PodFailed, "mock message"),
			eventSequence: []watch.Event{},
			expectedError: "old recycler pod found, will retry later",
		},
	}

	for _, test := range tests {
		t.Logf("Test %q", test.name)
		client := &mockRecyclerClient{
			events: test.eventSequence,
			pod:    test.existingPod,
		}
		err := internalRecycleVolumeByWatchingPodUntilCompletion(test.createPod.Name, test.createPod, client)
		receivedError := ""
		if err != nil {
			receivedError = err.Error()
		}
		if receivedError != test.expectedError {
			t.Errorf("Test %q failed, expected error %q, got %q", test.name, test.expectedError, receivedError)
			continue
		}
		if !client.deletedCalled {
			t.Errorf("Test %q failed, expected deferred client.Delete to be called on recycler pod", test.name)
			continue
		}
		for i, expectedEvent := range test.expectedEvents {
			if len(client.receivedEvents) <= i {
				t.Errorf("Test %q failed, expected event %d: %q not received", test.name, i, expectedEvent.message)
				continue
			}
			receivedEvent := client.receivedEvents[i]
			if expectedEvent.eventtype != receivedEvent.eventtype {
				t.Errorf("Test %q failed, event %d does not match: expected eventtype %q, got %q", test.name, i, expectedEvent.eventtype, receivedEvent.eventtype)
			}
			if expectedEvent.message != receivedEvent.message {
				t.Errorf("Test %q failed, event %d does not match: expected message %q, got %q", test.name, i, expectedEvent.message, receivedEvent.message)
			}
		}
		for i := len(test.expectedEvents); i < len(client.receivedEvents); i++ {
			t.Errorf("Test %q failed, unexpected event received: %s, %q", test.name, client.receivedEvents[i].eventtype, client.receivedEvents[i].message)
		}
	}
}

type mockRecyclerClient struct {
	pod            *v1.Pod
	deletedCalled  bool
	receivedEvents []mockEvent
	events         []watch.Event
}

type mockEvent struct {
	eventtype, message string
}

func (c *mockRecyclerClient) CreatePod(pod *v1.Pod) (*v1.Pod, error) {
	if c.pod == nil {
		c.pod = pod
		return c.pod, nil
	}
	// Simulate "already exists" error
	return nil, errors.NewAlreadyExists(api.Resource("pods"), pod.Name)
}

func (c *mockRecyclerClient) GetPod(name, namespace string) (*v1.Pod, error) {
	if c.pod != nil {
		return c.pod, nil
	}
	return nil, fmt.Errorf("pod does not exist")
}

func (c *mockRecyclerClient) DeletePod(name, namespace string) error {
	c.deletedCalled = true
	return nil
}

func (c *mockRecyclerClient) WatchPod(name, namespace string, stopChannel chan struct{}) (<-chan watch.Event, error) {
	eventCh := make(chan watch.Event, 0)
	go func() {
		for _, e := range c.events {
			eventCh <- e
		}
	}()
	return eventCh, nil
}

func (c *mockRecyclerClient) Event(eventtype, message string) {
	c.receivedEvents = append(c.receivedEvents, mockEvent{eventtype, message})
}
