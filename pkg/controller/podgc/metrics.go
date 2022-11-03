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

package podgc

import (
	"sync"

	"k8s.io/component-base/metrics"
	"k8s.io/component-base/metrics/legacyregistry"
)

const (
	podGCController = "pod_gc_collector"
)

var (
	deletingPodsTotal = metrics.NewCounterVec(
		&metrics.CounterOpts{
			Subsystem:      podGCController,
			Name:           "force_delete_pods_total",
			Help:           "Number of pods that are being forcefully deleted since the Pod GC Controller started.",
			StabilityLevel: metrics.ALPHA,
		},
		[]string{},
	)
	deletingPodsErrorTotal = metrics.NewCounterVec(
		&metrics.CounterOpts{
			Subsystem:      podGCController,
			Name:           "force_delete_pod_errors_total",
			Help:           "Number of errors encountered when forcefully deleting the pods since the Pod GC Controller started.",
			StabilityLevel: metrics.ALPHA,
		},
		[]string{},
	)
)

var registerMetrics sync.Once

// Register the metrics that are to be monitored.
func RegisterMetrics() {
	registerMetrics.Do(func() {
		legacyregistry.MustRegister(deletingPodsTotal)
		legacyregistry.MustRegister(deletingPodsErrorTotal)
	})
}
