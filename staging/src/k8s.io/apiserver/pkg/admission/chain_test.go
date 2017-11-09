/*
Copyright 2015 The Kubernetes Authors.

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

package admission

import (
	"fmt"
	"testing"

	"k8s.io/apimachinery/pkg/runtime/schema"
)

type FakeHandler struct {
	*Handler
	name                     string
	admit, admitCalled       bool
	validate, validateCalled bool
}

func (h *FakeHandler) Admit(a Attributes) (err error) {
	h.admitCalled = true
	if h.admit {
		return nil
	}
	return fmt.Errorf("Don't admit")
}

func (h *FakeHandler) Validate(a Attributes) (err error) {
	h.validateCalled = true
	if h.validate {
		return nil
	}
	return fmt.Errorf("Don't validate")
}

func makeHandler(name string, accept bool, ops ...Operation) Interface {
	return &FakeHandler{
		name:     name,
		admit:    accept,
		validate: accept,
		Handler:  NewHandler(ops...),
	}
}

func TestAdmitAndValidate(t *testing.T) {
	tests := []struct {
		name      string
		operation Operation
		chain     chainAdmissionHandler
		accept    bool
		calls     map[string]bool
	}{
		{
			name:      "all accept",
			operation: Create,
			chain: []Interface{
				makeHandler("a", true, Update, Delete, Create),
				makeHandler("b", true, Delete, Create),
				makeHandler("c", true, Create),
			},
			calls:  map[string]bool{"a": true, "b": true, "c": true},
			accept: true,
		},
		{
			name:      "ignore handler",
			operation: Create,
			chain: []Interface{
				makeHandler("a", true, Update, Delete, Create),
				makeHandler("b", false, Delete),
				makeHandler("c", true, Create),
			},
			calls:  map[string]bool{"a": true, "c": true},
			accept: true,
		},
		{
			name:      "ignore all",
			operation: Connect,
			chain: []Interface{
				makeHandler("a", true, Update, Delete, Create),
				makeHandler("b", false, Delete),
				makeHandler("c", true, Create),
			},
			calls:  map[string]bool{},
			accept: true,
		},
		{
			name:      "reject one",
			operation: Delete,
			chain: []Interface{
				makeHandler("a", true, Update, Delete, Create),
				makeHandler("b", false, Delete),
				makeHandler("c", true, Create),
			},
			calls:  map[string]bool{"a": true, "b": true},
			accept: false,
		},
	}
	for _, test := range tests {
		// call admit and check that validate was not called at all
		err := test.chain.Admit(NewAttributesRecord(nil, nil, schema.GroupVersionKind{}, "", "", schema.GroupVersionResource{}, "", test.operation, nil))
		accepted := (err == nil)
		if accepted != test.accept {
			t.Errorf("%s: unexpected result of admit call: %v\n", test.name, accepted)
		}
		for _, h := range test.chain {
			fake := h.(*FakeHandler)
			_, shouldBeCalled := test.calls[fake.name]
			if shouldBeCalled != fake.admitCalled {
				t.Errorf("%s: admit handler %s not called as expected: %v", test.name, fake.name, fake.admitCalled)
				continue
			}
			if fake.validateCalled {
				t.Errorf("%s: validate handler %s called during admit", test.name, fake.name)
			}

			// reset value for validation test
			fake.admitCalled = false
		}

		// call validate and check that admit was not called at all
		err = test.chain.Validate(NewAttributesRecord(nil, nil, schema.GroupVersionKind{}, "", "", schema.GroupVersionResource{}, "", test.operation, nil))
		accepted = (err == nil)
		if accepted != test.accept {
			t.Errorf("%s: unexpected result of validate call: %v\n", test.name, accepted)
		}
		for _, h := range test.chain {
			fake := h.(*FakeHandler)
			_, shouldBeCalled := test.calls[fake.name]
			if shouldBeCalled != fake.validateCalled {
				t.Errorf("%s: validate handler %s not called as expected: %v", test.name, fake.name, fake.validateCalled)
				continue
			}

			if fake.admitCalled {
				t.Errorf("%s: admit handler %s called during admit", test.name, fake.name)
			}
		}
	}
}

func TestHandles(t *testing.T) {
	chain := chainAdmissionHandler{
		makeHandler("a", true, Update, Delete, Create),
		makeHandler("b", true, Delete, Create),
		makeHandler("c", true, Create),
	}

	tests := []struct {
		name      string
		operation Operation
		chain     chainAdmissionHandler
		expected  bool
	}{
		{
			name:      "all handle",
			operation: Create,
			expected:  true,
		},
		{
			name:      "none handle",
			operation: Connect,
			expected:  false,
		},
		{
			name:      "some handle",
			operation: Delete,
			expected:  true,
		},
	}
	for _, test := range tests {
		handles := chain.Handles(test.operation)
		if handles != test.expected {
			t.Errorf("Unexpected handles result. Expected: %v. Actual: %v", test.expected, handles)
		}
	}
}
