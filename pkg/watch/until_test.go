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

package watch

import (
	"errors"
	"strings"
	"testing"
	"time"

	"k8s.io/kubernetes/pkg/api"
	"k8s.io/kubernetes/pkg/runtime"
	"k8s.io/kubernetes/pkg/util/wait"
)

func TestUntil(t *testing.T) {
	fw := NewFake()
	go func() {
		var obj *api.Pod
		fw.Add(obj)
		fw.Modify(obj)
	}()
	conditions := []ConditionFunc{
		func(event Event) (bool, error) { return event.Type == Added, nil },
		func(event Event) (bool, error) { return event.Type == Modified, nil },
	}

	timeout := time.Minute
	lastEvent, err := Until(timeout, fw, conditions...)
	if err != nil {
		t.Fatalf("expected nil error, got %#v", err)
	}
	if lastEvent == nil {
		t.Fatal("expected an event")
	}
	if lastEvent.Type != Modified {
		t.Fatalf("expected MODIFIED event type, got %v", lastEvent.Type)
	}
	if got, isPod := lastEvent.Object.(*api.Pod); !isPod {
		t.Fatalf("expected a pod event, got %#v", got)
	}
}

func TestUntilMultipleConditions(t *testing.T) {
	fw := NewFake()
	go func() {
		var obj *api.Pod
		fw.Add(obj)
	}()
	conditions := []ConditionFunc{
		func(event Event) (bool, error) { return event.Type == Added, nil },
		func(event Event) (bool, error) { return event.Type == Added, nil },
	}

	timeout := time.Minute
	lastEvent, err := Until(timeout, fw, conditions...)
	if err != nil {
		t.Fatalf("expected nil error, got %#v", err)
	}
	if lastEvent == nil {
		t.Fatal("expected an event")
	}
	if lastEvent.Type != Added {
		t.Fatalf("expected MODIFIED event type, got %v", lastEvent.Type)
	}
	if got, isPod := lastEvent.Object.(*api.Pod); !isPod {
		t.Fatalf("expected a pod event, got %#v", got)
	}
}

func TestUntilMultipleConditionsFail(t *testing.T) {
	fw := NewFake()
	go func() {
		var obj *api.Pod
		fw.Add(obj)
	}()
	conditions := []ConditionFunc{
		func(event Event) (bool, error) { return event.Type == Added, nil },
		func(event Event) (bool, error) { return event.Type == Added, nil },
		func(event Event) (bool, error) { return event.Type == Deleted, nil },
	}

	timeout := 10 * time.Second
	lastEvent, err := Until(timeout, fw, conditions...)
	if err != wait.ErrWaitTimeout {
		t.Fatalf("expected ErrWaitTimeout error, got %#v", err)
	}
	if lastEvent == nil {
		t.Fatal("expected an event")
	}
	if lastEvent.Type != Added {
		t.Fatalf("expected ADDED event type, got %v", lastEvent.Type)
	}
	if got, isPod := lastEvent.Object.(*api.Pod); !isPod {
		t.Fatalf("expected a pod event, got %#v", got)
	}
}

func TestUntilTimeout(t *testing.T) {
	fw := NewFake()
	go func() {
		var obj *api.Pod
		fw.Add(obj)
		fw.Modify(obj)
	}()
	conditions := []ConditionFunc{
		func(event Event) (bool, error) {
			return event.Type == Added, nil
		},
		func(event Event) (bool, error) {
			return event.Type == Modified, nil
		},
	}

	timeout := time.Duration(0)
	lastEvent, err := Until(timeout, fw, conditions...)
	if err != nil {
		t.Fatalf("expected nil error, got %#v", err)
	}
	if lastEvent == nil {
		t.Fatal("expected an event")
	}
	if lastEvent.Type != Modified {
		t.Fatalf("expected MODIFIED event type, got %v", lastEvent.Type)
	}
	if got, isPod := lastEvent.Object.(*api.Pod); !isPod {
		t.Fatalf("expected a pod event, got %#v", got)
	}
}

func TestUntilErrorCondition(t *testing.T) {
	fw := NewFake()
	go func() {
		var obj *api.Pod
		fw.Add(obj)
	}()
	expected := "something bad"
	conditions := []ConditionFunc{
		func(event Event) (bool, error) { return event.Type == Added, nil },
		func(event Event) (bool, error) { return false, errors.New(expected) },
	}

	timeout := time.Minute
	_, err := Until(timeout, fw, conditions...)
	if err == nil {
		t.Fatal("expected an error")
	}
	if !strings.Contains(err.Error(), expected) {
		t.Fatalf("expected %q in error string, got %q", expected, err.Error())
	}
}

type lw struct {
	list  runtime.Object
	watch Interface
}

func (w lw) List(options api.ListOptions) (runtime.Object, error) {
	return w.list, nil
}

func (w lw) Watch(options api.ListOptions) (Interface, error) {
	return w.watch, nil
}

func TestListWatchUntil(t *testing.T) {
	fw := NewFake()
	go func() {
		var obj *api.Pod
		fw.Modify(obj)
	}()
	listwatch := lw{
		list:  &api.PodList{Items: []api.Pod{{}}},
		watch: fw,
	}

	conditions := []ConditionFunc{
		func(event Event) (bool, error) {
			t.Logf("got %#v", event)
			return event.Type == Added, nil
		},
		func(event Event) (bool, error) {
			t.Logf("got %#v", event)
			return event.Type == Modified, nil
		},
	}

	timeout := 10 * time.Second
	lastEvent, err := ListWatchUntil(timeout, listwatch, conditions...)
	if err != nil {
		t.Fatalf("expected nil error, got %#v", err)
	}
	if lastEvent == nil {
		t.Fatal("expected an event")
	}
	if lastEvent.Type != Modified {
		t.Fatalf("expected MODIFIED event type, got %v", lastEvent.Type)
	}
	if got, isPod := lastEvent.Object.(*api.Pod); !isPod {
		t.Fatalf("expected a pod event, got %#v", got)
	}
}
