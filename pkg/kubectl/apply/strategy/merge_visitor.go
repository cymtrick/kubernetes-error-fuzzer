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

package strategy

import (
	"fmt"

	"k8s.io/kubernetes/pkg/kubectl/apply"
)

func createMergeStrategy(options Options, strategic *delegatingStrategy) mergeStrategy {
	return mergeStrategy{
		strategic,
		options,
	}
}

// mergeStrategy creates a patch to merge a local file value into a remote field value
type mergeStrategy struct {
	strategic *delegatingStrategy
	options   Options
}

// MergeList creates a patch to merge a local list field value into a remote list field value
func (v mergeStrategy) MergeList(e apply.ListElement) (apply.Result, error) {
	// No merge logic if adding or deleting a field
	if result, done := v.doAddOrDelete(e); done {
		return result, nil
	}

	// Merge each item in the list and append it to the list
	merged := []interface{}{}
	for _, value := range e.Values {
		// Recursively merge the list element before adding the value to the list
		m, err := value.Merge(v.strategic)
		if err != nil {
			return apply.Result{}, err
		}

		fmt.Printf("\nResult %+v\n%+v\n", m.MergedResult, value)
		switch m.Operation {
		case apply.SET:
			// Keep the list item value
			merged = append(merged, m.MergedResult)
		case apply.DROP:
			// Drop the list item value
		default:
			panic(fmt.Errorf("Unexpected result operation type %+v", m))
		}
	}

	if len(merged) == 0 {
		// If the list is empty, return a nil entry
		return apply.Result{Operation: apply.SET, MergedResult: nil}, nil
	}
	// Return the merged list, and tell the caller to keep it
	return apply.Result{Operation: apply.SET, MergedResult: merged}, nil
}

// MergeMap creates a patch to merge a local map field into a remote map field
func (v mergeStrategy) MergeMap(e apply.MapElement) (apply.Result, error) {
	return v.doMergeMap(e)
}

// MergeType creates a patch to merge a local map field into a remote map field
func (v mergeStrategy) MergeType(e apply.TypeElement) (apply.Result, error) {
	return v.doMergeMap(e)
}

// do merges a recorded, local and remote map into a new object
func (v mergeStrategy) doMergeMap(e apply.MapValuesElement) (apply.Result, error) {
	// No merge logic if adding or deleting a field
	if result, done := v.doAddOrDelete(e); done {
		return result, nil
	}

	// Merge each item in the list
	merged := map[string]interface{}{}
	for key, value := range e.GetValues() {
		// Recursively merge the map element before adding the value to the map
		result, err := value.Merge(v.strategic)
		if err != nil {
			return apply.Result{}, err
		}

		switch result.Operation {
		case apply.SET:
			// Keep the map item value
			merged[key] = result.MergedResult
		case apply.DROP:
			// Drop the map item value
		default:
			panic(fmt.Errorf("Unexpected result operation type %+v", result))
		}
	}

	// Return the merged map, and tell the caller to keep it
	if len(merged) == 0 {
		// Special case the empty map to set the field value to nil, but keep the field key
		// This is how the tests expect the structures to look when parsed from yaml
		return apply.Result{Operation: apply.SET, MergedResult: nil}, nil
	}
	return apply.Result{Operation: apply.SET, MergedResult: merged}, nil
}

func (v mergeStrategy) doAddOrDelete(e apply.Element) (apply.Result, bool) {
	if apply.IsAdd(e) {
		return apply.Result{Operation: apply.SET, MergedResult: e.GetLocal()}, true
	}

	// Delete the List
	if apply.IsDrop(e) {
		return apply.Result{Operation: apply.DROP}, true
	}

	return apply.Result{}, false
}

// MergePrimitive returns and error.  Primitive elements can't be merged, only replaced.
func (v mergeStrategy) MergePrimitive(diff apply.PrimitiveElement) (apply.Result, error) {
	return apply.Result{}, fmt.Errorf("Cannot merge primitive element %v", diff.Name)
}

// MergeEmpty
func (v mergeStrategy) MergeEmpty(diff apply.EmptyElement) (apply.Result, error) {
	return apply.Result{Operation: apply.SET}, nil
}

var _ apply.Strategy = &mergeStrategy{}
