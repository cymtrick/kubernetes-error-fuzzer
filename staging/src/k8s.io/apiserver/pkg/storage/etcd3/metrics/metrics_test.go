/*
Copyright 2023 The Kubernetes Authors.

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
	"errors"
	"strings"
	"testing"
	"time"

	"k8s.io/component-base/metrics"
	"k8s.io/component-base/metrics/testutil"
)

func TestRecordDecodeError(t *testing.T) {
	registry := metrics.NewKubeRegistry()
	defer registry.Reset()
	registry.Register(decodeErrorCounts)
	resourceName := "pods"
	testedMetrics := "apiserver_storage_decode_errors_total"
	testCases := []struct {
		desc     string
		resource string
		want     string
	}{
		{
			desc:     "test success",
			resource: resourceName,
			want: `
		# HELP apiserver_storage_decode_errors_total [ALPHA] Number of stored object decode errors split by object type
		# TYPE apiserver_storage_decode_errors_total counter
		apiserver_storage_decode_errors_total{resource="pods"} 1
`,
		},
	}

	for _, test := range testCases {
		t.Run(test.desc, func(t *testing.T) {
			RecordDecodeError(test.resource)
			if err := testutil.GatherAndCompare(registry, strings.NewReader(test.want), testedMetrics); err != nil {
				t.Fatal(err)
			}
		})
	}
}

func TestRecordEtcdRequest(t *testing.T) {
	registry := metrics.NewKubeRegistry()

	// modify default sinceInSeconds to constant NOW
	sinceInSeconds = func(t time.Time) float64 {
		return time.Unix(0, 300*int64(time.Millisecond)).Sub(t).Seconds()
	}

	testedMetrics := []metrics.Registerable{
		etcdRequestCounts,
		etcdRequestErrorCounts,
		etcdRequestLatency,
	}

	testedMetricsName := make([]string, 0, len(testedMetrics))
	for _, m := range testedMetrics {
		registry.MustRegister(m)
		testedMetricsName = append(testedMetricsName, m.FQName())
	}

	testCases := []struct {
		desc      string
		operation string
		typ       string
		err       error
		startTime time.Time
		want      string
	}{
		{
			desc:      "success_request",
			operation: "foo",
			typ:       "bar",
			err:       nil,
			startTime: time.Unix(0, 0), // 0.3s
			want: `# HELP etcd_request_duration_seconds [ALPHA] Etcd request latency in seconds for each operation and object type.
# TYPE etcd_request_duration_seconds histogram
etcd_request_duration_seconds_bucket{operation="foo",type="bar",le="0.005"} 0
etcd_request_duration_seconds_bucket{operation="foo",type="bar",le="0.025"} 0
etcd_request_duration_seconds_bucket{operation="foo",type="bar",le="0.05"} 0
etcd_request_duration_seconds_bucket{operation="foo",type="bar",le="0.1"} 0
etcd_request_duration_seconds_bucket{operation="foo",type="bar",le="0.2"} 0
etcd_request_duration_seconds_bucket{operation="foo",type="bar",le="0.4"} 1
etcd_request_duration_seconds_bucket{operation="foo",type="bar",le="0.6"} 1
etcd_request_duration_seconds_bucket{operation="foo",type="bar",le="0.8"} 1
etcd_request_duration_seconds_bucket{operation="foo",type="bar",le="1"} 1
etcd_request_duration_seconds_bucket{operation="foo",type="bar",le="1.25"} 1
etcd_request_duration_seconds_bucket{operation="foo",type="bar",le="1.5"} 1
etcd_request_duration_seconds_bucket{operation="foo",type="bar",le="2"} 1
etcd_request_duration_seconds_bucket{operation="foo",type="bar",le="3"} 1
etcd_request_duration_seconds_bucket{operation="foo",type="bar",le="4"} 1
etcd_request_duration_seconds_bucket{operation="foo",type="bar",le="5"} 1
etcd_request_duration_seconds_bucket{operation="foo",type="bar",le="6"} 1
etcd_request_duration_seconds_bucket{operation="foo",type="bar",le="8"} 1
etcd_request_duration_seconds_bucket{operation="foo",type="bar",le="10"} 1
etcd_request_duration_seconds_bucket{operation="foo",type="bar",le="15"} 1
etcd_request_duration_seconds_bucket{operation="foo",type="bar",le="20"} 1
etcd_request_duration_seconds_bucket{operation="foo",type="bar",le="30"} 1
etcd_request_duration_seconds_bucket{operation="foo",type="bar",le="45"} 1
etcd_request_duration_seconds_bucket{operation="foo",type="bar",le="60"} 1
etcd_request_duration_seconds_bucket{operation="foo",type="bar",le="+Inf"} 1
etcd_request_duration_seconds_sum{operation="foo",type="bar"} 0.3
etcd_request_duration_seconds_count{operation="foo",type="bar"} 1
# HELP etcd_requests_total [ALPHA] Etcd request counts for each operation and object type.
# TYPE etcd_requests_total counter
etcd_requests_total{operation="foo",type="bar"} 1
`,
		},
		{
			desc:      "failed_request",
			operation: "foo",
			typ:       "bar",
			err:       errors.New("some error"),
			startTime: time.Unix(0, 0), // 0.3s
			want: `# HELP etcd_request_duration_seconds [ALPHA] Etcd request latency in seconds for each operation and object type.
# TYPE etcd_request_duration_seconds histogram
etcd_request_duration_seconds_bucket{operation="foo",type="bar",le="0.005"} 0
etcd_request_duration_seconds_bucket{operation="foo",type="bar",le="0.025"} 0
etcd_request_duration_seconds_bucket{operation="foo",type="bar",le="0.05"} 0
etcd_request_duration_seconds_bucket{operation="foo",type="bar",le="0.1"} 0
etcd_request_duration_seconds_bucket{operation="foo",type="bar",le="0.2"} 0
etcd_request_duration_seconds_bucket{operation="foo",type="bar",le="0.4"} 1
etcd_request_duration_seconds_bucket{operation="foo",type="bar",le="0.6"} 1
etcd_request_duration_seconds_bucket{operation="foo",type="bar",le="0.8"} 1
etcd_request_duration_seconds_bucket{operation="foo",type="bar",le="1"} 1
etcd_request_duration_seconds_bucket{operation="foo",type="bar",le="1.25"} 1
etcd_request_duration_seconds_bucket{operation="foo",type="bar",le="1.5"} 1
etcd_request_duration_seconds_bucket{operation="foo",type="bar",le="2"} 1
etcd_request_duration_seconds_bucket{operation="foo",type="bar",le="3"} 1
etcd_request_duration_seconds_bucket{operation="foo",type="bar",le="4"} 1
etcd_request_duration_seconds_bucket{operation="foo",type="bar",le="5"} 1
etcd_request_duration_seconds_bucket{operation="foo",type="bar",le="6"} 1
etcd_request_duration_seconds_bucket{operation="foo",type="bar",le="8"} 1
etcd_request_duration_seconds_bucket{operation="foo",type="bar",le="10"} 1
etcd_request_duration_seconds_bucket{operation="foo",type="bar",le="15"} 1
etcd_request_duration_seconds_bucket{operation="foo",type="bar",le="20"} 1
etcd_request_duration_seconds_bucket{operation="foo",type="bar",le="30"} 1
etcd_request_duration_seconds_bucket{operation="foo",type="bar",le="45"} 1
etcd_request_duration_seconds_bucket{operation="foo",type="bar",le="60"} 1
etcd_request_duration_seconds_bucket{operation="foo",type="bar",le="+Inf"} 1
etcd_request_duration_seconds_sum{operation="foo",type="bar"} 0.3
etcd_request_duration_seconds_count{operation="foo",type="bar"} 1
# HELP etcd_requests_total [ALPHA] Etcd request counts for each operation and object type.
# TYPE etcd_requests_total counter
etcd_requests_total{operation="foo",type="bar"} 1
# HELP etcd_request_errors_total [ALPHA] Etcd failed request counts for each operation and object type.
# TYPE etcd_request_errors_total counter
etcd_request_errors_total{operation="foo",type="bar"} 1
`,
		},
	}

	for _, test := range testCases {
		t.Run(test.desc, func(t *testing.T) {
			defer registry.Reset()
			RecordEtcdRequest(test.operation, test.typ, test.err, test.startTime)
			if err := testutil.GatherAndCompare(registry, strings.NewReader(test.want), testedMetricsName...); err != nil {
				t.Fatal(err)
			}
		})
	}
}
