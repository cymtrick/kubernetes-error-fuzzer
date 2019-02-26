/*
Copyright 2019 The Kubernetes Authors.

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

package fieldmanager_test

import (
	"errors"
	"testing"

	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/meta"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apiserver/pkg/endpoints/handlers/fieldmanager"
)

type fakeObjectConvertor struct{}

func (c *fakeObjectConvertor) Convert(in, out, context interface{}) error {
	out = in
	return nil
}

func (c *fakeObjectConvertor) ConvertToVersion(in runtime.Object, _ runtime.GroupVersioner) (runtime.Object, error) {
	return in, nil
}

func (c *fakeObjectConvertor) ConvertFieldLabel(_ schema.GroupVersionKind, _, _ string) (string, string, error) {
	return "", "", errors.New("not implemented")
}

type fakeObjectDefaulter struct{}

func (d *fakeObjectDefaulter) Default(in runtime.Object) {}

func NewTestFieldManager(t *testing.T) *fieldmanager.FieldManager {
	gv := schema.GroupVersion{
		Group:   "apps",
		Version: "v1",
	}

	return fieldmanager.NewCRDFieldManager(
		&fakeObjectConvertor{},
		&fakeObjectDefaulter{},
		gv,
		gv,
	)
}

func TestFieldManagerCreation(t *testing.T) {
	if NewTestFieldManager(t) == nil {
		t.Fatal("failed to create FieldManager")
	}
}

func TestApplyStripsFields(t *testing.T) {
	f := NewTestFieldManager(t)

	obj := &corev1.Pod{}

	newObj, err := f.Apply(obj, []byte(`{
		"apiVersion": "v1",
		"kind": "Deployment",
		"metadata": {
			"name": "b",
			"namespace": "b",
			"creationTimestamp": "2016-05-19T09:59:00Z",
			"selfLink": "b",
			"uid": "b",
			"clusterName": "b",
			"generation": 0,
			"managedFields": [],
			"resourceVersion": "b"
		}
	}`), false)
	if err != nil {
		t.Fatalf("failed to apply object: %v", err)
	}

	accessor, err := meta.Accessor(newObj)
	if err != nil {
		t.Fatalf("couldn't get accessor: %v", err)
	}

	if m := accessor.GetManagedFields(); len(m) != 0 {
		t.Fatalf("fields did not get stripped on apply: %v", m)
	}
}
func TestApplyDoesNotStripLabels(t *testing.T) {
	f := NewTestFieldManager(t)

	obj := &corev1.Pod{}

	newObj, err := f.Apply(obj, []byte(`{
		"apiVersion": "v1",
		"kind": "Pod",
		"metadata": {
			"labels": {
				"a": "b"
			},
		}
	}`), false)
	if err != nil {
		t.Fatalf("failed to apply object: %v", err)
	}

	accessor, err := meta.Accessor(newObj)
	if err != nil {
		t.Fatalf("couldn't get accessor: %v", err)
	}

	if m := accessor.GetManagedFields(); len(m) != 1 {
		t.Fatalf("labels shouldn't get stripped on apply: %v", m)
	}
}
