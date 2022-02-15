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

package topologymanager

import (
	"k8s.io/klog/v2"
	"k8s.io/kubernetes/pkg/kubelet/cm/topologymanager/bitmask"
)

// Policy interface for Topology Manager Pod Admit Result
type Policy interface {
	// Returns Policy Name
	Name() string
	// Returns a merged TopologyHint based on input from hint providers
	// and a Pod Admit Handler Response based on hints and policy type
	Merge(providersHints []map[string][]TopologyHint) (TopologyHint, bool)
}

// Merge a TopologyHints permutation to a single hint by performing a bitwise-AND
// of their affinity masks. The hint shall be preferred if all hits in the permutation
// are preferred.
func mergePermutation(numaNodes []int, permutation []TopologyHint) TopologyHint {
	// Get the NUMANodeAffinity from each hint in the permutation and see if any
	// of them encode unpreferred allocations.
	preferred := true
	defaultAffinity, _ := bitmask.NewBitMask(numaNodes...)
	var numaAffinities []bitmask.BitMask
	for _, hint := range permutation {
		// Only consider hints that have an actual NUMANodeAffinity set.
		if hint.NUMANodeAffinity != nil {
			numaAffinities = append(numaAffinities, hint.NUMANodeAffinity)
			// Only mark preferred if all affinities are equal.
			if !hint.NUMANodeAffinity.IsEqual(numaAffinities[0]) {
				preferred = false
			}
		}
		// Only mark preferred if all affinities are preferred.
		if !hint.Preferred {
			preferred = false
		}
	}

	// Merge the affinities using a bitwise-and operation.
	mergedAffinity := bitmask.And(defaultAffinity, numaAffinities...)
	// Build a mergedHint from the merged affinity mask, setting preferred as
	// appropriate based on the logic above.
	return TopologyHint{mergedAffinity, preferred}
}

func filterProvidersHints(providersHints []map[string][]TopologyHint) [][]TopologyHint {
	// Loop through all hint providers and save an accumulated list of the
	// hints returned by each hint provider. If no hints are provided, assume
	// that provider has no preference for topology-aware allocation.
	var allProviderHints [][]TopologyHint
	for _, hints := range providersHints {
		// If hints is nil, insert a single, preferred any-numa hint into allProviderHints.
		if len(hints) == 0 {
			klog.InfoS("Hint Provider has no preference for NUMA affinity with any resource")
			allProviderHints = append(allProviderHints, []TopologyHint{{nil, true}})
			continue
		}

		// Otherwise, accumulate the hints for each resource type into allProviderHints.
		for resource := range hints {
			if hints[resource] == nil {
				klog.InfoS("Hint Provider has no preference for NUMA affinity with resource", "resource", resource)
				allProviderHints = append(allProviderHints, []TopologyHint{{nil, true}})
				continue
			}

			if len(hints[resource]) == 0 {
				klog.InfoS("Hint Provider has no possible NUMA affinities for resource", "resource", resource)
				allProviderHints = append(allProviderHints, []TopologyHint{{nil, false}})
				continue
			}

			allProviderHints = append(allProviderHints, hints[resource])
		}
	}
	return allProviderHints
}

func compareHints(current *TopologyHint, candidate *TopologyHint) *TopologyHint {
	// Only consider candidates that result in a NUMANodeAffinity > 0 to
	// replace the current bestHint.
	if candidate.NUMANodeAffinity.Count() == 0 {
		return current
	}

	// If no current bestHint is set, return the candidate as the bestHint.
	if current == nil {
		return candidate
	}

	// If the current bestHint is non-preferred and the candidate hint is
	// preferred, always choose the preferred hint over the non-preferred one.
	if !current.Preferred && candidate.Preferred {
		return candidate
	}

	// If the current bestHint is preferred and the candidate hint is
	// non-preferred, never update the bestHint, regardless of the
	// candidate hint's narowness.
	if current.Preferred && !candidate.Preferred {
		return current
	}

	// If the current bestHint and the candidate hint have the same
	// preference, only consider candidate hints that have a narrower
	// NUMANodeAffinity than the NUMANodeAffinity in the current bestHint.
	if !candidate.NUMANodeAffinity.IsNarrowerThan(current.NUMANodeAffinity) {
		return current
	}

	// In all other cases, update the bestHint to the candidate hint.
	return candidate
}

func mergeFilteredHints(numaNodes []int, filteredHints [][]TopologyHint) TopologyHint {
	var bestHint *TopologyHint
	iterateAllProviderTopologyHints(filteredHints, func(permutation []TopologyHint) {
		// Get the NUMANodeAffinity from each hint in the permutation and see if any
		// of them encode unpreferred allocations.
		mergedHint := mergePermutation(numaNodes, permutation)

		// Compare the current bestHint with the candidate mergedHint and
		// update bestHint if appropriate.
		bestHint = compareHints(bestHint, &mergedHint)
	})

	if bestHint == nil {
		defaultAffinity, _ := bitmask.NewBitMask(numaNodes...)
		bestHint = &TopologyHint{defaultAffinity, false}
	}

	return *bestHint
}

// Iterate over all permutations of hints in 'allProviderHints [][]TopologyHint'.
//
// This procedure is implemented as a recursive function over the set of hints
// in 'allproviderHints[i]'. It applies the function 'callback' to each
// permutation as it is found. It is the equivalent of:
//
// for i := 0; i < len(providerHints[0]); i++
//     for j := 0; j < len(providerHints[1]); j++
//         for k := 0; k < len(providerHints[2]); k++
//             ...
//             for z := 0; z < len(providerHints[-1]); z++
//                 permutation := []TopologyHint{
//                     providerHints[0][i],
//                     providerHints[1][j],
//                     providerHints[2][k],
//                     ...
//                     providerHints[-1][z]
//                 }
//                 callback(permutation)
func iterateAllProviderTopologyHints(allProviderHints [][]TopologyHint, callback func([]TopologyHint)) {
	// Internal helper function to accumulate the permutation before calling the callback.
	var iterate func(i int, accum []TopologyHint)
	iterate = func(i int, accum []TopologyHint) {
		// Base case: we have looped through all providers and have a full permutation.
		if i == len(allProviderHints) {
			callback(accum)
			return
		}

		// Loop through all hints for provider 'i', and recurse to build the
		// the permutation of this hint with all hints from providers 'i++'.
		for j := range allProviderHints[i] {
			iterate(i+1, append(accum, allProviderHints[i][j]))
		}
	}
	iterate(0, []TopologyHint{})
}
