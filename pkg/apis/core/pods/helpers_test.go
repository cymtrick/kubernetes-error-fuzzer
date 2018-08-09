/*
Copyright 2017 The Kubernetes Authors.

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

package pods

import (
	"reflect"
	"testing"

	"k8s.io/apimachinery/pkg/util/validation/field"
	utilfeature "k8s.io/apiserver/pkg/util/feature"
	featuregatetesting "k8s.io/component-base/featuregate/testing"
	api "k8s.io/kubernetes/pkg/apis/core"
	"k8s.io/kubernetes/pkg/features"
)

func TestVisitContainersWithPath(t *testing.T) {
	defer featuregatetesting.SetFeatureGateDuringTest(t, utilfeature.DefaultFeatureGate, features.EphemeralContainers, true)()

	testCases := []struct {
		description string
		haveSpec    *api.PodSpec
		wantNames   []string
	}{
		{
			"empty podspec",
			&api.PodSpec{},
			[]string{},
		},
		{
			"regular containers",
			&api.PodSpec{
				Containers: []api.Container{
					{Name: "c1"},
					{Name: "c2"},
				},
			},
			[]string{"spec.containers[0]", "spec.containers[1]"},
		},
		{
			"init containers",
			&api.PodSpec{
				InitContainers: []api.Container{
					{Name: "i1"},
					{Name: "i2"},
				},
			},
			[]string{"spec.initContainers[0]", "spec.initContainers[1]"},
		},
		{
			"regular and init containers",
			&api.PodSpec{
				Containers: []api.Container{
					{Name: "c1"},
					{Name: "c2"},
				},
				InitContainers: []api.Container{
					{Name: "i1"},
					{Name: "i2"},
				},
			},
			[]string{"spec.initContainers[0]", "spec.initContainers[1]", "spec.containers[0]", "spec.containers[1]"},
		},
		{
			"ephemeral containers",
			&api.PodSpec{
				Containers: []api.Container{
					{Name: "c1"},
					{Name: "c2"},
				},
				EphemeralContainers: []api.EphemeralContainer{
					{EphemeralContainerCommon: api.EphemeralContainerCommon{Name: "e1"}},
				},
			},
			[]string{"spec.containers[0]", "spec.containers[1]", "spec.ephemeralContainers[0]"},
		},
		{
			"all container types",
			&api.PodSpec{
				Containers: []api.Container{
					{Name: "c1"},
					{Name: "c2"},
				},
				InitContainers: []api.Container{
					{Name: "i1"},
					{Name: "i2"},
				},
				EphemeralContainers: []api.EphemeralContainer{
					{EphemeralContainerCommon: api.EphemeralContainerCommon{Name: "e1"}},
					{EphemeralContainerCommon: api.EphemeralContainerCommon{Name: "e2"}},
				},
			},
			[]string{"spec.initContainers[0]", "spec.initContainers[1]", "spec.containers[0]", "spec.containers[1]", "spec.ephemeralContainers[0]", "spec.ephemeralContainers[1]"},
		},
	}

	for _, tc := range testCases {
		gotNames := []string{}
		VisitContainersWithPath(tc.haveSpec, func(c *api.Container, p *field.Path) bool {
			gotNames = append(gotNames, p.String())
			return true
		})
		if !reflect.DeepEqual(gotNames, tc.wantNames) {
			t.Errorf("VisitContainersWithPath() for test case %q visited containers %q, wanted to visit %q", tc.description, gotNames, tc.wantNames)
		}
	}
}

func TestConvertDownwardAPIFieldLabel(t *testing.T) {
	testCases := []struct {
		version       string
		label         string
		value         string
		expectedErr   bool
		expectedLabel string
		expectedValue string
	}{
		{
			version:     "v2",
			label:       "metadata.name",
			value:       "test-pod",
			expectedErr: true,
		},
		{
			version:     "v1",
			label:       "invalid-label",
			value:       "value",
			expectedErr: true,
		},
		{
			version:       "v1",
			label:         "metadata.name",
			value:         "test-pod",
			expectedLabel: "metadata.name",
			expectedValue: "test-pod",
		},
		{
			version:       "v1",
			label:         "metadata.annotations",
			value:         "myValue",
			expectedLabel: "metadata.annotations",
			expectedValue: "myValue",
		},
		{
			version:       "v1",
			label:         "metadata.annotations['myKey']",
			value:         "myValue",
			expectedLabel: "metadata.annotations['myKey']",
			expectedValue: "myValue",
		},
		{
			version:       "v1",
			label:         "spec.host",
			value:         "127.0.0.1",
			expectedLabel: "spec.nodeName",
			expectedValue: "127.0.0.1",
		},
	}
	for _, tc := range testCases {
		label, value, err := ConvertDownwardAPIFieldLabel(tc.version, tc.label, tc.value)
		if err != nil {
			if tc.expectedErr {
				continue
			}
			t.Errorf("ConvertDownwardAPIFieldLabel(%s, %s, %s) failed: %s",
				tc.version, tc.label, tc.value, err)
		}
		if tc.expectedLabel != label || tc.expectedValue != value {
			t.Errorf("ConvertDownwardAPIFieldLabel(%s, %s, %s) = (%s, %s, nil), expected (%s, %s, nil)",
				tc.version, tc.label, tc.value, label, value, tc.expectedLabel, tc.expectedValue)
		}
	}
}
