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

package app

import (
	"testing"
)

func TestValueOfAllocatableResources(t *testing.T) {
	testCases := []struct {
		kubeReserved   map[string]string
		systemReserved map[string]string
		errorExpected  bool
		name           string
	}{
		{
			kubeReserved:   map[string]string{"cpu": "200m", "memory": "-150G", "ephemeral-storage": "10Gi"},
			systemReserved: map[string]string{"cpu": "200m", "memory": "15Ki"},
			errorExpected:  true,
			name:           "negative quantity value",
		},
		{
			kubeReserved:   map[string]string{"cpu": "200m", "memory": "150Gi", "ephemeral-storage": "10Gi"},
			systemReserved: map[string]string{"cpu": "200m", "memory": "15Ky"},
			errorExpected:  true,
			name:           "invalid quantity unit",
		},
		{
			kubeReserved:   map[string]string{"cpu": "200m", "memory": "15G", "ephemeral-storage": "10Gi"},
			systemReserved: map[string]string{"cpu": "200m", "memory": "15Ki"},
			errorExpected:  false,
			name:           "Valid resource quantity",
		},
	}

	for _, test := range testCases {
		_, err1 := parseResourceList(test.kubeReserved)
		_, err2 := parseResourceList(test.systemReserved)
		if test.errorExpected {
			if err1 == nil && err2 == nil {
				t.Errorf("%s: error expected", test.name)
			}
		} else {
			if err1 != nil || err2 != nil {
				t.Errorf("%s: unexpected error: %v, %v", test.name, err1, err2)
			}
		}
	}
}

func TestValueOfReservedMemoryConfig(t *testing.T) {
	testCases := []struct {
		config        []map[string]string
		errorExpected bool
		name          string
	}{
		{
			config:        []map[string]string{{"numa-node": "0", "type": "memory", "limit": "2Gi"}},
			errorExpected: false,
			name:          "Valid resource quantity",
		},
		{
			config:        []map[string]string{{"numa-node": "0", "type": "memory", "limit": "2000m"}, {"numa-node": "1", "type": "memory", "limit": "1Gi"}},
			errorExpected: false,
			name:          "Valid resource quantity",
		},
		{
			config:        []map[string]string{{"type": "memory", "limit": "2Gi"}},
			errorExpected: true,
			name:          "Missing key",
		},
		{
			config:        []map[string]string{{"numa-node": "one", "type": "memory", "limit": "2Gi"}},
			errorExpected: true,
			name:          "Wrong 'numa-node' value",
		},
		{
			config:        []map[string]string{{"numa-node": "0", "type": "not-memory", "limit": "2Gi"}},
			errorExpected: true,
			name:          "Wrong 'memory' value",
		},
		{
			config:        []map[string]string{{"numa-node": "0", "type": "memory", "limit": "2Gigs"}},
			errorExpected: true,
			name:          "Wrong 'limit' value",
		},
		{
			config:        []map[string]string{{"numa-node": "-1", "type": "memory", "limit": "2Gigs"}},
			errorExpected: true,
			name:          "Invalid 'numa-node' number",
		},
	}

	for _, test := range testCases {
		_, err := parseReservedMemoryConfig(test.config)
		if test.errorExpected {
			if err == nil {
				t.Errorf("%s: error expected", test.name)
			}
		} else {
			if err != nil {
				t.Errorf("%s: unexpected error: %v", test.name, err)
			}
		}
	}
}
