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

package metrics

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type resourceUtilizationRatioTestCase struct {
	metrics           PodMetricsInfo
	requests          map[string]int64
	targetUtilization int32

	expectedUtilizationRatio   float64
	expectedCurrentUtilization int32
	expectedRawAverageValue    int64
	expectedErr                error
}

func (tc *resourceUtilizationRatioTestCase) runTest(t *testing.T) {
	actualUtilizationRatio, actualCurrentUtilization, actualRawAverageValue, actualErr := GetResourceUtilizationRatio(tc.metrics, tc.requests, tc.targetUtilization)

	if tc.expectedErr != nil {
		assert.Error(t, actualErr, "there should be an error getting the utilization ratio")
		assert.Contains(t, fmt.Sprintf("%v", actualErr), fmt.Sprintf("%v", tc.expectedErr), "the error message should be as expected")
		return
	}

	assert.NoError(t, actualErr, "there should be no error retrieving the utilization ratio")
	assert.Equal(t, tc.expectedUtilizationRatio, actualUtilizationRatio, "the utilization ratios should be as expected")
	assert.Equal(t, tc.expectedCurrentUtilization, actualCurrentUtilization, "the current utilization should be as expected")
	assert.Equal(t, tc.expectedRawAverageValue, actualRawAverageValue, "the raw average value should be as expected")
}

type metricUtilizationRatioTestCase struct {
	metrics           PodMetricsInfo
	targetUtilization int64

	expectedUtilizationRatio   float64
	expectedCurrentUtilization int64
}

func (tc *metricUtilizationRatioTestCase) runTest(t *testing.T) {
	actualUtilizationRatio, actualCurrentUtilization := GetMetricUtilizationRatio(tc.metrics, tc.targetUtilization)

	assert.Equal(t, tc.expectedUtilizationRatio, actualUtilizationRatio, "the utilization ratios should be as expected")
	assert.Equal(t, tc.expectedCurrentUtilization, actualCurrentUtilization, "the current utilization should be as expected")
}

func TestGetResourceUtilizationRatioBaseCase(t *testing.T) {
	tc := resourceUtilizationRatioTestCase{
		metrics: PodMetricsInfo{
			"test-pod-0": {Value: 50}, "test-pod-1": {Value: 76},
		},
		requests: map[string]int64{
			"test-pod-0": 100, "test-pod-1": 100,
		},
		targetUtilization:          50,
		expectedUtilizationRatio:   1.26,
		expectedCurrentUtilization: 63,
		expectedRawAverageValue:    63,
		expectedErr:                nil,
	}

	tc.runTest(t)
}

func TestGetResourceUtilizationRatioIgnorePodsWithNoRequest(t *testing.T) {
	tc := resourceUtilizationRatioTestCase{
		metrics: PodMetricsInfo{
			"test-pod-0": {Value: 50}, "test-pod-1": {Value: 76}, "test-pod-no-request": {Value: 100},
		},
		requests: map[string]int64{
			"test-pod-0": 100, "test-pod-1": 100,
		},
		targetUtilization:          50,
		expectedUtilizationRatio:   1.26,
		expectedCurrentUtilization: 63,
		expectedRawAverageValue:    63,
		expectedErr:                nil,
	}

	tc.runTest(t)
}

func TestGetResourceUtilizationRatioExtraRequest(t *testing.T) {
	tc := resourceUtilizationRatioTestCase{
		metrics: PodMetricsInfo{
			"test-pod-0": {Value: 50}, "test-pod-1": {Value: 76},
		},
		requests: map[string]int64{
			"test-pod-0": 100, "test-pod-1": 100, "test-pod-extra-request": 500,
		},
		targetUtilization:          50,
		expectedUtilizationRatio:   1.26,
		expectedCurrentUtilization: 63,
		expectedRawAverageValue:    63,
		expectedErr:                nil,
	}

	tc.runTest(t)
}

func TestGetResourceUtilizationRatioNoRequests(t *testing.T) {
	tc := resourceUtilizationRatioTestCase{
		metrics: PodMetricsInfo{
			"test-pod-0": {Value: 50}, "test-pod-1": {Value: 76},
		},
		requests:          map[string]int64{},
		targetUtilization: 50,

		expectedUtilizationRatio:   0,
		expectedCurrentUtilization: 0,
		expectedRawAverageValue:    0,
		expectedErr:                fmt.Errorf("no metrics returned matched known pods"),
	}

	tc.runTest(t)
}

func TestGetMetricUtilizationRatioBaseCase(t *testing.T) {
	tc := metricUtilizationRatioTestCase{
		metrics: PodMetricsInfo{
			"test-pod-0": {Value: 5000}, "test-pod-1": {Value: 10000},
		},
		targetUtilization:          10000,
		expectedUtilizationRatio:   .75,
		expectedCurrentUtilization: 7500,
	}

	tc.runTest(t)
}
