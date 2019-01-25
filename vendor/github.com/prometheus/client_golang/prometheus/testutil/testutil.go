// Copyright 2018 The Prometheus Authors
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Package testutil provides helpers to test code using the prometheus package
// of client_golang.
//
// While writing unit tests to verify correct instrumentation of your code, it's
// a common mistake to mostly test the instrumentation library instead of your
// own code. Rather than verifying that a prometheus.Counter's value has changed
// as expected or that it shows up in the exposition after registration, it is
// in general more robust and more faithful to the concept of unit tests to use
// mock implementations of the prometheus.Counter and prometheus.Registerer
// interfaces that simply assert that the Add or Register methods have been
// called with the expected arguments. However, this might be overkill in simple
// scenarios. The ToFloat64 function is provided for simple inspection of a
// single-value metric, but it has to be used with caution.
//
// End-to-end tests to verify all or larger parts of the metrics exposition can
// be implemented with the CollectAndCompare or GatherAndCompare functions. The
// most appropriate use is not so much testing instrumentation of your code, but
// testing custom prometheus.Collector implementations and in particular whole
// exporters, i.e. programs that retrieve telemetry data from a 3rd party source
// and convert it into Prometheus metrics.
package testutil

import (
	"bytes"
	"fmt"
	"io"
	"reflect"

	"github.com/prometheus/common/expfmt"

	dto "github.com/prometheus/client_model/go"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/internal"
)

// ToFloat64 collects all Metrics from the provided Collector. It expects that
// this results in exactly one Metric being collected, which must be a Gauge,
// Counter, or Untyped. In all other cases, ToFloat64 panics. ToFloat64 returns
// the value of the collected Metric.
//
// The Collector provided is typically a simple instance of Gauge or Counter, or
// – less commonly – a GaugeVec or CounterVec with exactly one element. But any
// Collector fulfilling the prerequisites described above will do.
//
// Use this function with caution. It is computationally very expensive and thus
// not suited at all to read values from Metrics in regular code. This is really
// only for testing purposes, and even for testing, other approaches are often
// more appropriate (see this package's documentation).
//
// A clear anti-pattern would be to use a metric type from the prometheus
// package to track values that are also needed for something else than the
// exposition of Prometheus metrics. For example, you would like to track the
// number of items in a queue because your code should reject queuing further
// items if a certain limit is reached. It is tempting to track the number of
// items in a prometheus.Gauge, as it is then easily available as a metric for
// exposition, too. However, then you would need to call ToFloat64 in your
// regular code, potentially quite often. The recommended way is to track the
// number of items conventionally (in the way you would have done it without
// considering Prometheus metrics) and then expose the number with a
// prometheus.GaugeFunc.
func ToFloat64(c prometheus.Collector) float64 {
	var (
		m      prometheus.Metric
		mCount int
		mChan  = make(chan prometheus.Metric)
		done   = make(chan struct{})
	)

	go func() {
		for m = range mChan {
			mCount++
		}
		close(done)
	}()

	c.Collect(mChan)
	close(mChan)
	<-done

	if mCount != 1 {
		panic(fmt.Errorf("collected %d metrics instead of exactly 1", mCount))
	}

	pb := &dto.Metric{}
	m.Write(pb)
	if pb.Gauge != nil {
		return pb.Gauge.GetValue()
	}
	if pb.Counter != nil {
		return pb.Counter.GetValue()
	}
	if pb.Untyped != nil {
		return pb.Untyped.GetValue()
	}
	panic(fmt.Errorf("collected a non-gauge/counter/untyped metric: %s", pb))
}

// CollectAndCompare registers the provided Collector with a newly created
// pedantic Registry. It then does the same as GatherAndCompare, gathering the
// metrics from the pedantic Registry.
func CollectAndCompare(c prometheus.Collector, expected io.Reader, metricNames ...string) error {
	reg := prometheus.NewPedanticRegistry()
	if err := reg.Register(c); err != nil {
		return fmt.Errorf("registering collector failed: %s", err)
	}
	return GatherAndCompare(reg, expected, metricNames...)
}

// GatherAndCompare gathers all metrics from the provided Gatherer and compares
// it to an expected output read from the provided Reader in the Prometheus text
// exposition format. If any metricNames are provided, only metrics with those
// names are compared.
func GatherAndCompare(g prometheus.Gatherer, expected io.Reader, metricNames ...string) error {
	metrics, err := g.Gather()
	if err != nil {
		return fmt.Errorf("gathering metrics failed: %s", err)
	}
	if metricNames != nil {
		metrics = filterMetrics(metrics, metricNames)
	}
	var tp expfmt.TextParser
	expectedMetrics, err := tp.TextToMetricFamilies(expected)
	if err != nil {
		return fmt.Errorf("parsing expected metrics failed: %s", err)
	}

	if !reflect.DeepEqual(metrics, internal.NormalizeMetricFamilies(expectedMetrics)) {
		// Encode the gathered output to the readable text format for comparison.
		var buf1 bytes.Buffer
		enc := expfmt.NewEncoder(&buf1, expfmt.FmtText)
		for _, mf := range metrics {
			if err := enc.Encode(mf); err != nil {
				return fmt.Errorf("encoding result failed: %s", err)
			}
		}
		// Encode normalized expected metrics again to generate them in the same ordering
		// the registry does to spot differences more easily.
		var buf2 bytes.Buffer
		enc = expfmt.NewEncoder(&buf2, expfmt.FmtText)
		for _, mf := range internal.NormalizeMetricFamilies(expectedMetrics) {
			if err := enc.Encode(mf); err != nil {
				return fmt.Errorf("encoding result failed: %s", err)
			}
		}

		return fmt.Errorf(`
metric output does not match expectation; want:

%s

got:

%s
`, buf2.String(), buf1.String())
	}
	return nil
}

func filterMetrics(metrics []*dto.MetricFamily, names []string) []*dto.MetricFamily {
	var filtered []*dto.MetricFamily
	for _, m := range metrics {
		for _, name := range names {
			if m.GetName() == name {
				filtered = append(filtered, m)
				break
			}
		}
	}
	return filtered
}
