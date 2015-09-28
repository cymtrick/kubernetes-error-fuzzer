// +build cgo,linux

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

package oom

import (
	"testing"
)

// Converts a sequence of PID lists into a PID lister.
// The PID lister returns pidListSequence[i] on the ith call. If i >= length of pidListSequence
// then return the last element of pidListSequence (the sequence is considered to have) stabilized.
func sequenceToPidLister(pidListSequence [][]int) func(string) ([]int, error) {
	var numCalls int
	return func(cgroupName string) ([]int, error) {
		numCalls++
		if len(pidListSequence) == 0 {
			return []int{}, nil
		} else if numCalls > len(pidListSequence) {
			return pidListSequence[len(pidListSequence)-1], nil
		}
		return pidListSequence[numCalls-1], nil
	}
}

// Tests that applyOOMScoreAdjContainer correctly applies OOM scores to relevant processes, or
// returns the right error.
func applyOOMScoreAdjContainerTester(pidListSequence [][]int, maxTries int, appliedPids []int, expectedError bool, t *testing.T) {
	pidOOMs := make(map[int]bool)

	// Mock ApplyOOMScoreAdj and pidLister.
	oomAdjuster := NewOOMAdjuster()
	oomAdjuster.ApplyOOMScoreAdj = func(pid int, oomScoreAdj int) error {
		pidOOMs[pid] = true
		return nil
	}
	oomAdjuster.pidLister = sequenceToPidLister(pidListSequence)
	err := oomAdjuster.ApplyOOMScoreAdjContainer("", 100, maxTries)

	// Check error value.
	if expectedError && err == nil {
		t.Errorf("Expected error %+v when running ApplyOOMScoreAdjContainer but got no error", expectedError)
		return
	} else if !expectedError && err != nil {
		t.Errorf("Expected no error but got error %+v when running ApplyOOMScoreAdjContainer", err)
		return
	} else if err != nil {
		return
	}

	// Check that OOM scores were applied to the right processes.
	if len(appliedPids) != len(pidOOMs) {
		t.Errorf("Applied OOM scores to incorrect number of processes")
		return
	}
	for _, pid := range appliedPids {
		if !pidOOMs[pid] {
			t.Errorf("Failed to apply OOM scores to process %d", pid)
		}
	}
}

func TestOOMScoreAdjContainer(t *testing.T) {
	pidListSequenceEmpty := [][]int{}
	applyOOMScoreAdjContainerTester(pidListSequenceEmpty, 3, nil, true, t)

	pidListSequence1 := [][]int{
		{1, 2},
	}
	applyOOMScoreAdjContainerTester(pidListSequence1, 1, nil, true, t)
	applyOOMScoreAdjContainerTester(pidListSequence1, 2, []int{1, 2}, false, t)
	applyOOMScoreAdjContainerTester(pidListSequence1, 3, []int{1, 2}, false, t)

	pidListSequence3 := [][]int{
		{1, 2},
		{1, 2, 4, 5},
		{2, 1, 4, 5, 3},
	}
	applyOOMScoreAdjContainerTester(pidListSequence3, 1, nil, true, t)
	applyOOMScoreAdjContainerTester(pidListSequence3, 2, nil, true, t)
	applyOOMScoreAdjContainerTester(pidListSequence3, 3, nil, true, t)
	applyOOMScoreAdjContainerTester(pidListSequence3, 4, []int{1, 2, 3, 4, 5}, false, t)

	pidListSequenceLag := [][]int{
		{},
		{},
		{},
		{1, 2, 4},
		{1, 2, 4, 5},
	}
	for i := 1; i < 5; i++ {
		applyOOMScoreAdjContainerTester(pidListSequenceLag, i, nil, true, t)
	}
	applyOOMScoreAdjContainerTester(pidListSequenceLag, 6, []int{1, 2, 4, 5}, false, t)
}
