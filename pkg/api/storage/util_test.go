/*
Copyright 2022 The Kubernetes Authors.

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

package storage

import (
	"testing"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/sets"

	"k8s.io/kubernetes/pkg/apis/core"
	"k8s.io/kubernetes/pkg/apis/storage"
)

func TestStorageClassWarnings(t *testing.T) {
	testcases := []struct {
		name     string
		template *storage.StorageClass
		expected []string
	}{
		{
			name:     "null",
			template: nil,
			expected: nil,
		},
		{
			name: "no warning",
			template: &storage.StorageClass{
				ObjectMeta: metav1.ObjectMeta{
					Name: "foo",
				},
			},
			expected: nil,
		},
		{
			name: "warning",
			template: &storage.StorageClass{
				ObjectMeta: metav1.ObjectMeta{
					Name: "foo",
				},
				AllowedTopologies: []core.TopologySelectorTerm{
					{
						MatchLabelExpressions: []core.TopologySelectorLabelRequirement{
							{
								Key:    "beta.kubernetes.io/arch",
								Values: []string{"amd64"},
							},
							{
								Key:    "beta.kubernetes.io/os",
								Values: []string{"linux"},
							},
						},
					},
				},
			},
			expected: []string{
				`allowedTopologies[0].matchLabelExpressions[0].key: deprecated since v1.14; use "kubernetes.io/arch" instead`,
				`allowedTopologies[0].matchLabelExpressions[1].key: deprecated since v1.14; use "kubernetes.io/os" instead`,
			},
		},
	}

	for _, tc := range testcases {
		t.Run("podspec_"+tc.name, func(t *testing.T) {
			actual := sets.NewString(GetWarningsForStorageClass(tc.template)...)
			expected := sets.NewString(tc.expected...)
			for _, missing := range expected.Difference(actual).List() {
				t.Errorf("missing: %s", missing)
			}
			for _, extra := range actual.Difference(expected).List() {
				t.Errorf("extra: %s", extra)
			}
		})

	}
}

func TestCSIStorageCapacityWarnings(t *testing.T) {
	testcases := []struct {
		name     string
		template *storage.CSIStorageCapacity
		expected []string
	}{
		{
			name:     "null",
			template: nil,
			expected: nil,
		},
		{
			name: "no warning",
			template: &storage.CSIStorageCapacity{
				ObjectMeta: metav1.ObjectMeta{
					Name: "foo",
				},
			},
			expected: nil,
		},
		{
			name: "MatchLabels warning",
			template: &storage.CSIStorageCapacity{
				ObjectMeta: metav1.ObjectMeta{
					Name: "foo",
				},
				NodeTopology: &metav1.LabelSelector{
					MatchLabels: map[string]string{
						"beta.kubernetes.io/arch": "amd64",
						"beta.kubernetes.io/os":   "linux",
					},
				},
			},
			expected: []string{
				`nodeTopology.matchLabels.beta.kubernetes.io/arch: deprecated since v1.14; use "kubernetes.io/arch" instead`,
				`nodeTopology.matchLabels.beta.kubernetes.io/os: deprecated since v1.14; use "kubernetes.io/os" instead`,
			},
		},
		{
			name: "MatchExpressions warning",
			template: &storage.CSIStorageCapacity{
				ObjectMeta: metav1.ObjectMeta{
					Name: "foo",
				},
				NodeTopology: &metav1.LabelSelector{
					MatchExpressions: []metav1.LabelSelectorRequirement{
						{
							Key:      "beta.kubernetes.io/arch",
							Operator: metav1.LabelSelectorOpIn,
							Values:   []string{"amd64"},
						},
						{
							Key:      "beta.kubernetes.io/os",
							Operator: metav1.LabelSelectorOpIn,
							Values:   []string{"linux"},
						},
					},
				},
			},
			expected: []string{
				`nodeTopology.matchExpressions[0].key: beta.kubernetes.io/arch is deprecated since v1.14; use "kubernetes.io/arch" instead`,
				`nodeTopology.matchExpressions[1].key: beta.kubernetes.io/os is deprecated since v1.14; use "kubernetes.io/os" instead`,
			},
		},
	}

	for _, tc := range testcases {
		t.Run("podspec_"+tc.name, func(t *testing.T) {
			actual := sets.NewString(GetWarningsForCSIStorageCapacity(tc.template)...)
			expected := sets.NewString(tc.expected...)
			for _, missing := range expected.Difference(actual).List() {
				t.Errorf("missing: %s", missing)
			}
			for _, extra := range actual.Difference(expected).List() {
				t.Errorf("extra: %s", extra)
			}
		})

	}
}
