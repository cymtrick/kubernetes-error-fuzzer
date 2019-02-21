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

package metrics

import (
	"sync"
	"time"

	"github.com/prometheus/client_golang/prometheus"
)

const (
	// NetworkPluginOperationsKey is the key for operation count metrics.
	NetworkPluginOperationsKey = "network_plugin_operations"
	// NetworkPluginOperationsLatencyKey is the key for the operation latency metrics.
	NetworkPluginOperationsLatencyKey = "network_plugin_operations_latency_seconds"
	// DeprecatedNetworkPluginOperationsLatencyKey is the deprecated key for the operation latency metrics.
	DeprecatedNetworkPluginOperationsLatencyKey = "network_plugin_operations_latency_microseconds"

	// Keep the "kubelet" subsystem for backward compatibility.
	kubeletSubsystem = "kubelet"
)

var (
	// NetworkPluginOperationsLatency collects operation latency numbers by operation
	// type.
	NetworkPluginOperationsLatency = prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Subsystem: kubeletSubsystem,
			Name:      NetworkPluginOperationsLatencyKey,
			Help:      "Latency in seconds of network plugin operations. Broken down by operation type.",
			Buckets:   prometheus.DefBuckets,
		},
		[]string{"operation_type"},
	)

	// DeprecatedNetworkPluginOperationsLatency collects operation latency numbers by operation
	// type.
	DeprecatedNetworkPluginOperationsLatency = prometheus.NewSummaryVec(
		prometheus.SummaryOpts{
			Subsystem: kubeletSubsystem,
			Name:      DeprecatedNetworkPluginOperationsLatencyKey,
			Help:      "(Deprecated) Latency in microseconds of network plugin operations. Broken down by operation type.",
		},
		[]string{"operation_type"},
	)
)

var registerMetrics sync.Once

// Register all metrics.
func Register() {
	registerMetrics.Do(func() {
		prometheus.MustRegister(NetworkPluginOperationsLatency)
		prometheus.MustRegister(DeprecatedNetworkPluginOperationsLatency)
	})
}

// SinceInMicroseconds gets the time since the specified start in microseconds.
func SinceInMicroseconds(start time.Time) float64 {
	return float64(time.Since(start).Nanoseconds() / time.Microsecond.Nanoseconds())
}

// SinceInSeconds gets the time since the specified start in seconds.
func SinceInSeconds(start time.Time) float64 {
	return time.Since(start).Seconds()
}
