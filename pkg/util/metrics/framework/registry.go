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

package framework

import (
	"github.com/blang/semver"
	"github.com/prometheus/client_golang/prometheus"
	dto "github.com/prometheus/client_model/go"
	"k8s.io/klog"
	"k8s.io/kubernetes/pkg/version"
)

// DefaultGlobalRegistry is a stub for the global registry which prometheus client
// currently uses.
var DefaultGlobalRegistry = NewKubeRegistry()

// KubeRegistry is a wrapper around a prometheus registry-type object. Upon initialization
// the kubernetes binary version information is loaded into the registry object, so that
// automatic behavior can be configured for metric versioning.
type KubeRegistry struct {
	PromRegistry
	version semver.Version
}

// Register registers a collectable metric, but it uses a global registry.
func Register(c KubeCollector) error {
	return DefaultGlobalRegistry.Register(c)
}

// MustRegister works like Register but registers any number of
// Collectors and panics upon the first registration that causes an
// error.
func MustRegister(cs ...KubeCollector) {
	DefaultGlobalRegistry.MustRegister(cs...)
}

// Register registers a new Collector to be included in metrics
// collection. It returns an error if the descriptors provided by the
// Collector are invalid or if they — in combination with descriptors of
// already registered Collectors — do not fulfill the consistency and
// uniqueness criteria described in the documentation of metric.Desc.
func (kr *KubeRegistry) Register(c KubeCollector) error {
	if c.Create(&kr.version) {
		return kr.PromRegistry.Register(c)
	}
	return nil
}

// MustRegister works like Register but registers any number of
// Collectors and panics upon the first registration that causes an
// error.
func (kr *KubeRegistry) MustRegister(cs ...KubeCollector) {
	metrics := make([]prometheus.Collector, 0, len(cs))
	for _, c := range cs {
		if c.Create(&kr.version) {
			metrics = append(metrics, c)
		}
	}
	kr.PromRegistry.MustRegister(metrics...)
}

// Unregister unregisters the Collector that equals the Collector passed
// in as an argument.  (Two Collectors are considered equal if their
// Describe method yields the same set of descriptors.) The function
// returns whether a Collector was unregistered. Note that an unchecked
// Collector cannot be unregistered (as its Describe method does not
// yield any descriptor).
func (kr *KubeRegistry) Unregister(collector KubeCollector) bool {
	return kr.PromRegistry.Unregister(collector)
}

// Gather calls the Collect method of the registered Collectors and then
// gathers the collected metrics into a lexicographically sorted slice
// of uniquely named MetricFamily protobufs. Gather ensures that the
// returned slice is valid and self-consistent so that it can be used
// for valid exposition. As an exception to the strict consistency
// requirements described for metric.Desc, Gather will tolerate
// different sets of label names for metrics of the same metric family.
func (kr *KubeRegistry) Gather() ([]*dto.MetricFamily, error) {
	return kr.PromRegistry.Gather()
}

// NewKubeRegistry creates a new kubernetes metric registry, loading in the kubernetes
// version information available to the binary.
func NewKubeRegistry() *KubeRegistry {
	v, err := parseVersion(version.Get())
	if err != nil {
		klog.Fatalf("Can't initialize a registry without a valid version %v", err)
	}
	if v == nil {
		klog.Fatalf("No valid version loaded for metrics registry")
	}
	return newKubeRegistry(semver.MustParse(*v))
}

// newKubeRegistry creates a new vanilla Registry without any Collectors
// pre-registered.
func newKubeRegistry(version semver.Version) *KubeRegistry {
	return &KubeRegistry{
		PromRegistry: prometheus.NewRegistry(),
		version:      version,
	}
}
