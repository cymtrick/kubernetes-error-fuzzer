// +build linux

/*
Copyright 2015 The Kubernetes Authors All rights reserved.

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

package volume

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"
)

const expectedBlockSize = 4096

// TestMetricsDuGetCapacity tests that MetricsDu can read disk usage
// for path
func TestMetricsDuGetCapacity(t *testing.T) {
	tmpDir, err := ioutil.TempDir(os.TempDir(), "metrics_du_test")
	if err != nil {
		t.Fatalf("Can't make a tmp dir: %v", err)
	}
	defer os.RemoveAll(tmpDir)
	metrics := NewMetricsDu(tmpDir)

	expectedEmptyDirUsage, err := FindEmptyDirectoryUsageOnTmpfs()
	if err != nil {
		t.Errorf("Unexpected error finding expected empty directory usage on tmpfs: %v", err)
	}

	actual, err := metrics.GetMetrics()
	if err != nil {
		t.Errorf("Unexpected error when calling GetMetrics %v", err)
	}
	if e, a := expectedEmptyDirUsage.Value(), actual.Used.Value(); e != a {
		t.Errorf("Unexpected value for empty directory; expected %v, got %v", e, a)
	}

	// TODO(pwittroc): Figure out a way to test these values for correctness, maybe by formatting and mounting a file
	// as a filesystem
	if a := actual.Capacity.Value(); a <= 0 {
		t.Errorf("Expected Capacity %d to be greater than 0.", a)
	}
	if a := actual.Available.Value(); a <= 0 {
		t.Errorf("Expected Available %d to be greater than 0.", a)
	}

	// Write a file and expect Used to increase
	ioutil.WriteFile(filepath.Join(tmpDir, "f1"), []byte("Hello World"), os.ModeTemporary)
	actual, err = metrics.GetMetrics()
	if err != nil {
		t.Errorf("Unexpected error when calling GetMetrics %v", err)
	}
	if e, a := (expectedEmptyDirUsage.Value() + expectedBlockSize), actual.Used.Value(); e != a {
		t.Errorf("Unexpected Used for directory with file.  Expected %v, got %d.", e, a)
	}
}

// TestMetricsDuRequireInit tests that if MetricsDu is not initialized with a path, GetMetrics
// returns an error
func TestMetricsDuRequirePath(t *testing.T) {
	metrics := &metricsDu{}
	actual, err := metrics.GetMetrics()
	expected := &Metrics{}
	if *actual != *expected {
		t.Errorf("Expected empty Metrics from uninitialized MetricsDu, actual %v", *actual)
	}
	if err == nil {
		t.Errorf("Expected error when calling GetMetrics on uninitialized MetricsDu, actual nil")
	}
}

// TestMetricsDuRealDirectory tests that if MetricsDu is initialized to a non-existent path, GetMetrics
// returns an error
func TestMetricsDuRequireRealDirectory(t *testing.T) {
	metrics := NewMetricsDu("/not/a/real/directory")
	actual, err := metrics.GetMetrics()
	expected := &Metrics{}
	if *actual != *expected {
		t.Errorf("Expected empty Metrics from incorrectly initialized MetricsDu, actual %v", *actual)
	}
	if err == nil {
		t.Errorf("Expected error when calling GetMetrics on incorrectly initialized MetricsDu, actual nil")
	}
}
