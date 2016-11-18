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

package eventsink

import (
	"testing"

	fake_fedclientset "k8s.io/kubernetes/federation/client/clientset_generated/federation_release_1_5/fake"
	. "k8s.io/kubernetes/federation/pkg/federation-controller/util/test"
	api_v1 "k8s.io/kubernetes/pkg/api/v1"
	"k8s.io/kubernetes/pkg/client/testing/core"
	"k8s.io/kubernetes/pkg/runtime"

	"github.com/stretchr/testify/assert"
)

func TestEventSink(t *testing.T) {
	fakeFederationClient := &fake_fedclientset.Clientset{}
	createdChan := make(chan runtime.Object, 100)
	fakeFederationClient.AddReactor("create", "events", func(action core.Action) (bool, runtime.Object, error) {
		createAction := action.(core.CreateAction)
		obj := createAction.GetObject()
		createdChan <- obj
		return true, obj, nil
	})
	updateChan := make(chan runtime.Object, 100)
	fakeFederationClient.AddReactor("update", "events", func(action core.Action) (bool, runtime.Object, error) {
		updateAction := action.(core.UpdateAction)
		obj := updateAction.GetObject()
		updateChan <- obj
		return true, obj, nil
	})

	event := api_v1.Event{
		ObjectMeta: api_v1.ObjectMeta{
			Name:      "bzium",
			Namespace: "ns",
		},
	}
	sink := NewFederatedEventSink(fakeFederationClient)
	eventUpdated, err := sink.Create(&event)
	assert.NoError(t, err)
	eventV1 := GetObjectFromChan(createdChan).(*api_v1.Event)
	assert.NotNil(t, eventV1)
	// Just some simple sanity checks.
	assert.Equal(t, event.Name, eventV1.Name)
	assert.Equal(t, event.Name, eventUpdated.Name)

	eventUpdated, err = sink.Update(&event)
	assert.NoError(t, err)
	eventV1 = GetObjectFromChan(updateChan).(*api_v1.Event)
	assert.NotNil(t, eventV1)
	// Just some simple sanity checks.
	assert.Equal(t, event.Name, eventV1.Name)
	assert.Equal(t, event.Name, eventUpdated.Name)
}
