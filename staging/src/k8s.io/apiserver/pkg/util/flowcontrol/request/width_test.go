/*
Copyright 2021 The Kubernetes Authors.

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

package request

import (
	"errors"
	"net/http"
	"testing"
	"time"

	apirequest "k8s.io/apiserver/pkg/endpoints/request"
)

func TestWorkEstimator(t *testing.T) {
	tests := []struct {
		name                      string
		requestURI                string
		requestInfo               *apirequest.RequestInfo
		counts                    map[string]int64
		countErr                  error
		watchCount                int
		initialSeatsExpected      uint
		finalSeatsExpected        uint
		additionalLatencyExpected time.Duration
	}{
		{
			name:                 "request has no RequestInfo",
			requestURI:           "http://server/apis/",
			requestInfo:          nil,
			initialSeatsExpected: maximumSeats,
		},
		{
			name:       "request verb is not list",
			requestURI: "http://server/apis/",
			requestInfo: &apirequest.RequestInfo{
				Verb: "get",
			},
			initialSeatsExpected: minimumSeats,
		},
		{
			name:       "request verb is list, conversion to ListOptions returns error",
			requestURI: "http://server/apis/foo.bar/v1/events?limit=invalid",
			requestInfo: &apirequest.RequestInfo{
				Verb:     "list",
				APIGroup: "foo.bar",
				Resource: "events",
			},
			counts: map[string]int64{
				"events.foo.bar": 799,
			},
			initialSeatsExpected: maximumSeats,
		},
		{
			name:       "request verb is list, has limit and resource version is 1",
			requestURI: "http://server/apis/foo.bar/v1/events?limit=399&resourceVersion=1",
			requestInfo: &apirequest.RequestInfo{
				Verb:     "list",
				APIGroup: "foo.bar",
				Resource: "events",
			},
			counts: map[string]int64{
				"events.foo.bar": 699,
			},
			initialSeatsExpected: 8,
		},
		{
			name:       "request verb is list, limit not set",
			requestURI: "http://server/apis/foo.bar/v1/events?resourceVersion=1",
			requestInfo: &apirequest.RequestInfo{
				Verb:     "list",
				APIGroup: "foo.bar",
				Resource: "events",
			},
			counts: map[string]int64{
				"events.foo.bar": 699,
			},
			initialSeatsExpected: 7,
		},
		{
			name:       "request verb is list, resource version not set",
			requestURI: "http://server/apis/foo.bar/v1/events?limit=399",
			requestInfo: &apirequest.RequestInfo{
				Verb:     "list",
				APIGroup: "foo.bar",
				Resource: "events",
			},
			counts: map[string]int64{
				"events.foo.bar": 699,
			},
			initialSeatsExpected: 8,
		},
		{
			name:       "request verb is list, no query parameters, count known",
			requestURI: "http://server/apis/foo.bar/v1/events",
			requestInfo: &apirequest.RequestInfo{
				Verb:     "list",
				APIGroup: "foo.bar",
				Resource: "events",
			},
			counts: map[string]int64{
				"events.foo.bar": 399,
			},
			initialSeatsExpected: 8,
		},
		{
			name:       "request verb is list, no query parameters, count not known",
			requestURI: "http://server/apis/foo.bar/v1/events",
			requestInfo: &apirequest.RequestInfo{
				Verb:     "list",
				APIGroup: "foo.bar",
				Resource: "events",
			},
			countErr:             ObjectCountNotFoundErr,
			initialSeatsExpected: maximumSeats,
		},
		{
			name:       "request verb is list, continuation is set",
			requestURI: "http://server/apis/foo.bar/v1/events?continue=token&limit=399",
			requestInfo: &apirequest.RequestInfo{
				Verb:     "list",
				APIGroup: "foo.bar",
				Resource: "events",
			},
			counts: map[string]int64{
				"events.foo.bar": 699,
			},
			initialSeatsExpected: 8,
		},
		{
			name:       "request verb is list, resource version is zero",
			requestURI: "http://server/apis/foo.bar/v1/events?limit=299&resourceVersion=0",
			requestInfo: &apirequest.RequestInfo{
				Verb:     "list",
				APIGroup: "foo.bar",
				Resource: "events",
			},
			counts: map[string]int64{
				"events.foo.bar": 399,
			},
			initialSeatsExpected: 4,
		},
		{
			name:       "request verb is list, resource version is zero, no limit",
			requestURI: "http://server/apis/foo.bar/v1/events?resourceVersion=0",
			requestInfo: &apirequest.RequestInfo{
				Verb:     "list",
				APIGroup: "foo.bar",
				Resource: "events",
			},
			counts: map[string]int64{
				"events.foo.bar": 799,
			},
			initialSeatsExpected: 8,
		},
		{
			name:       "request verb is list, resource version match is Exact",
			requestURI: "http://server/apis/foo.bar/v1/events?resourceVersion=foo&resourceVersionMatch=Exact&limit=399",
			requestInfo: &apirequest.RequestInfo{
				Verb:     "list",
				APIGroup: "foo.bar",
				Resource: "events",
			},
			counts: map[string]int64{
				"events.foo.bar": 699,
			},
			initialSeatsExpected: 8,
		},
		{
			name:       "request verb is list, resource version match is NotOlderThan, limit not specified",
			requestURI: "http://server/apis/foo.bar/v1/events?resourceVersion=foo&resourceVersionMatch=NotOlderThan",
			requestInfo: &apirequest.RequestInfo{
				Verb:     "list",
				APIGroup: "foo.bar",
				Resource: "events",
			},
			counts: map[string]int64{
				"events.foo.bar": 799,
			},
			initialSeatsExpected: 8,
		},
		{
			name:       "request verb is list, maximum is capped",
			requestURI: "http://server/apis/foo.bar/v1/events?resourceVersion=foo",
			requestInfo: &apirequest.RequestInfo{
				Verb:     "list",
				APIGroup: "foo.bar",
				Resource: "events",
			},
			counts: map[string]int64{
				"events.foo.bar": 1999,
			},
			initialSeatsExpected: maximumSeats,
		},
		{
			name:       "request verb is list, list from cache, count not known",
			requestURI: "http://server/apis/foo.bar/v1/events?resourceVersion=0&limit=799",
			requestInfo: &apirequest.RequestInfo{
				Verb:     "list",
				APIGroup: "foo.bar",
				Resource: "events",
			},
			countErr:             ObjectCountNotFoundErr,
			initialSeatsExpected: maximumSeats,
		},
		{
			name:       "request verb is list, object count is stale",
			requestURI: "http://server/apis/foo.bar/v1/events?limit=499",
			requestInfo: &apirequest.RequestInfo{
				Verb:     "list",
				APIGroup: "foo.bar",
				Resource: "events",
			},
			counts: map[string]int64{
				"events.foo.bar": 799,
			},
			countErr:             ObjectCountStaleErr,
			initialSeatsExpected: maximumSeats,
		},
		{
			name:       "request verb is list, object count is not found",
			requestURI: "http://server/apis/foo.bar/v1/events?limit=499",
			requestInfo: &apirequest.RequestInfo{
				Verb:     "list",
				APIGroup: "foo.bar",
				Resource: "events",
			},
			countErr:             ObjectCountNotFoundErr,
			initialSeatsExpected: maximumSeats,
		},
		{
			name:       "request verb is list, count getter throws unknown error",
			requestURI: "http://server/apis/foo.bar/v1/events?limit=499",
			requestInfo: &apirequest.RequestInfo{
				Verb:     "list",
				APIGroup: "foo.bar",
				Resource: "events",
			},
			countErr:             errors.New("unknown error"),
			initialSeatsExpected: maximumSeats,
		},
		{
			name:       "request verb is create, no watches",
			requestURI: "http://server/apis/foo.bar/v1/foos",
			requestInfo: &apirequest.RequestInfo{
				Verb:     "create",
				APIGroup: "foo.bar",
				Resource: "foos",
			},
			initialSeatsExpected:      1,
			finalSeatsExpected:        0,
			additionalLatencyExpected: 0,
		},
		{
			name:       "request verb is create, watches registered",
			requestURI: "http://server/apis/foo.bar/v1/foos",
			requestInfo: &apirequest.RequestInfo{
				Verb:     "create",
				APIGroup: "foo.bar",
				Resource: "foos",
			},
			watchCount:                29,
			initialSeatsExpected:      1,
			finalSeatsExpected:        3,
			additionalLatencyExpected: 5 * time.Millisecond,
		},
		{
			name:       "request verb is create, watches registered, no additional latency",
			requestURI: "http://server/apis/foo.bar/v1/foos",
			requestInfo: &apirequest.RequestInfo{
				Verb:     "create",
				APIGroup: "foo.bar",
				Resource: "foos",
			},
			watchCount:                5,
			initialSeatsExpected:      1,
			finalSeatsExpected:        0,
			additionalLatencyExpected: 0,
		},
		{
			name:       "request verb is create, watches registered, maximum is capped",
			requestURI: "http://server/apis/foo.bar/v1/foos",
			requestInfo: &apirequest.RequestInfo{
				Verb:     "create",
				APIGroup: "foo.bar",
				Resource: "foos",
			},
			watchCount:                199,
			initialSeatsExpected:      1,
			finalSeatsExpected:        10,
			additionalLatencyExpected: 10 * time.Millisecond,
		},
		{
			name:       "request verb is update, no watches",
			requestURI: "http://server/apis/foo.bar/v1/foos/myfoo",
			requestInfo: &apirequest.RequestInfo{
				Verb:     "update",
				APIGroup: "foo.bar",
				Resource: "foos",
			},
			initialSeatsExpected:      1,
			finalSeatsExpected:        0,
			additionalLatencyExpected: 0,
		},
		{
			name:       "request verb is update, watches registered",
			requestURI: "http://server/apis/foor.bar/v1/foos/myfoo",
			requestInfo: &apirequest.RequestInfo{
				Verb:     "update",
				APIGroup: "foo.bar",
				Resource: "foos",
			},
			watchCount:                29,
			initialSeatsExpected:      1,
			finalSeatsExpected:        3,
			additionalLatencyExpected: 5 * time.Millisecond,
		},
		{
			name:       "request verb is patch, no watches",
			requestURI: "http://server/apis/foo.bar/v1/foos/myfoo",
			requestInfo: &apirequest.RequestInfo{
				Verb:     "patch",
				APIGroup: "foo.bar",
				Resource: "foos",
			},
			initialSeatsExpected:      1,
			finalSeatsExpected:        0,
			additionalLatencyExpected: 0,
		},
		{
			name:       "request verb is patch, watches registered",
			requestURI: "http://server/apis/foo.bar/v1/foos/myfoo",
			requestInfo: &apirequest.RequestInfo{
				Verb:     "patch",
				APIGroup: "foo.bar",
				Resource: "foos",
			},
			watchCount:                29,
			initialSeatsExpected:      1,
			finalSeatsExpected:        3,
			additionalLatencyExpected: 5 * time.Millisecond,
		},
		{
			name:       "request verb is delete, no watches",
			requestURI: "http://server/apis/foo.bar/v1/foos/myfoo",
			requestInfo: &apirequest.RequestInfo{
				Verb:     "delete",
				APIGroup: "foo.bar",
				Resource: "foos",
			},
			initialSeatsExpected:      1,
			finalSeatsExpected:        0,
			additionalLatencyExpected: 0,
		},
		{
			name:       "request verb is delete, watches registered",
			requestURI: "http://server/apis/foo.bar/v1/foos/myfoo",
			requestInfo: &apirequest.RequestInfo{
				Verb:     "delete",
				APIGroup: "foo.bar",
				Resource: "foos",
			},
			watchCount:                29,
			initialSeatsExpected:      1,
			finalSeatsExpected:        3,
			additionalLatencyExpected: 5 * time.Millisecond,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			counts := test.counts
			if len(counts) == 0 {
				counts = map[string]int64{}
			}
			countsFn := func(key string) (int64, error) {
				return counts[key], test.countErr
			}
			watchCountsFn := func(_ *apirequest.RequestInfo) int {
				return test.watchCount
			}

			// TODO(wojtek-t): Simplify it once we enable mutating work estimator
			// by default.
			testEstimator := &workEstimator{
				listWorkEstimator:     newListWorkEstimator(countsFn),
				mutatingWorkEstimator: newTestMutatingWorkEstimator(watchCountsFn, true),
			}
			estimator := WorkEstimatorFunc(testEstimator.estimate)

			req, err := http.NewRequest("GET", test.requestURI, nil)
			if err != nil {
				t.Fatalf("Failed to create new HTTP request - %v", err)
			}

			if test.requestInfo != nil {
				req = req.WithContext(apirequest.WithRequestInfo(req.Context(), test.requestInfo))
			}

			workestimateGot := estimator.EstimateWork(req)
			if test.initialSeatsExpected != workestimateGot.InitialSeats {
				t.Errorf("Expected work estimate to match: %d initial seats, but got: %d", test.initialSeatsExpected, workestimateGot.InitialSeats)
			}
			if test.finalSeatsExpected != workestimateGot.FinalSeats {
				t.Errorf("Expected work estimate to match: %d final seats, but got: %d", test.finalSeatsExpected, workestimateGot.FinalSeats)
			}
			if test.additionalLatencyExpected != workestimateGot.AdditionalLatency {
				t.Errorf("Expected work estimate to match additional latency: %v, but got: %v", test.additionalLatencyExpected, workestimateGot.AdditionalLatency)
			}
		})
	}
}
