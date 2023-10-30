// Copyright 2023 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package ext

import (
	"math"

	"github.com/google/cel-go/cel"
	"github.com/google/cel-go/checker"
	"github.com/google/cel-go/common/types"
	"github.com/google/cel-go/common/types/ref"
	"github.com/google/cel-go/common/types/traits"
	"github.com/google/cel-go/interpreter"
)

// Sets returns a cel.EnvOption to configure namespaced set relationship
// functions.
//
// There is no set type within CEL, and while one may be introduced in the
// future, there are cases where a `list` type is known to behave like a set.
// For such cases, this library provides some basic functionality for
// determining set containment, equivalence, and intersection.
//
// # Sets.Contains
//
// Returns whether the first list argument contains all elements in the second
// list argument. The list may contain elements of any type and standard CEL
// equality is used to determine whether a value exists in both lists. If the
// second list is empty, the result will always return true.
//
//	sets.contains(list(T), list(T)) -> bool
//
// Examples:
//
//	sets.contains([], []) // true
//	sets.contains([], [1]) // false
//	sets.contains([1, 2, 3, 4], [2, 3]) // true
//	sets.contains([1, 2.0, 3u], [1.0, 2u, 3]) // true
//
// # Sets.Equivalent
//
// Returns whether the first and second list are set equivalent. Lists are set
// equivalent if for every item in the first list, there is an element in the
// second which is equal. The lists may not be of the same size as they do not
// guarantee the elements within them are unique, so size does not factor into
// the computation.
//
// Examples:
//
//	sets.equivalent([], []) // true
//	sets.equivalent([1], [1, 1]) // true
//	sets.equivalent([1], [1u, 1.0]) // true
//	sets.equivalent([1, 2, 3], [3u, 2.0, 1]) // true
//
// # Sets.Intersects
//
// Returns whether the first list has at least one element whose value is equal
// to an element in the second list. If either list is empty, the result will
// be false.
//
// Examples:
//
//	sets.intersects([1], []) // false
//	sets.intersects([1], [1, 2]) // true
//	sets.intersects([[1], [2, 3]], [[1, 2], [2, 3.0]]) // true
func Sets() cel.EnvOption {
	return cel.Lib(setsLib{})
}

type setsLib struct{}

// LibraryName implements the SingletonLibrary interface method.
func (setsLib) LibraryName() string {
	return "cel.lib.ext.sets"
}

// CompileOptions implements the Library interface method.
func (setsLib) CompileOptions() []cel.EnvOption {
	listType := cel.ListType(cel.TypeParamType("T"))
	return []cel.EnvOption{
		cel.Function("sets.contains",
			cel.Overload("list_sets_contains_list", []*cel.Type{listType, listType}, cel.BoolType,
				cel.BinaryBinding(setsContains))),
		cel.Function("sets.equivalent",
			cel.Overload("list_sets_equivalent_list", []*cel.Type{listType, listType}, cel.BoolType,
				cel.BinaryBinding(setsEquivalent))),
		cel.Function("sets.intersects",
			cel.Overload("list_sets_intersects_list", []*cel.Type{listType, listType}, cel.BoolType,
				cel.BinaryBinding(setsIntersects))),
		cel.CostEstimatorOptions(
			checker.OverloadCostEstimate("list_sets_contains_list", estimateSetsCost(1)),
			checker.OverloadCostEstimate("list_sets_intersects_list", estimateSetsCost(1)),
			// equivalence requires potentially two m*n comparisons to ensure each list is contained by the other
			checker.OverloadCostEstimate("list_sets_equivalent_list", estimateSetsCost(2)),
		),
	}
}

// ProgramOptions implements the Library interface method.
func (setsLib) ProgramOptions() []cel.ProgramOption {
	return []cel.ProgramOption{
		cel.CostTrackerOptions(
			interpreter.OverloadCostTracker("list_sets_contains_list", trackSetsCost(1)),
			interpreter.OverloadCostTracker("list_sets_intersects_list", trackSetsCost(1)),
			interpreter.OverloadCostTracker("list_sets_equivalent_list", trackSetsCost(2)),
		),
	}
}

func setsIntersects(listA, listB ref.Val) ref.Val {
	lA := listA.(traits.Lister)
	lB := listB.(traits.Lister)
	it := lA.Iterator()
	for it.HasNext() == types.True {
		exists := lB.Contains(it.Next())
		if exists == types.True {
			return types.True
		}
	}
	return types.False
}

func setsContains(list, sublist ref.Val) ref.Val {
	l := list.(traits.Lister)
	sub := sublist.(traits.Lister)
	it := sub.Iterator()
	for it.HasNext() == types.True {
		exists := l.Contains(it.Next())
		if exists != types.True {
			return exists
		}
	}
	return types.True
}

func setsEquivalent(listA, listB ref.Val) ref.Val {
	aContainsB := setsContains(listA, listB)
	if aContainsB != types.True {
		return aContainsB
	}
	return setsContains(listB, listA)
}

func estimateSetsCost(costFactor float64) checker.FunctionEstimator {
	return func(estimator checker.CostEstimator, target *checker.AstNode, args []checker.AstNode) *checker.CallEstimate {
		if len(args) == 2 {
			arg0Size := estimateSize(estimator, args[0])
			arg1Size := estimateSize(estimator, args[1])
			costEstimate := arg0Size.Multiply(arg1Size).MultiplyByCostFactor(costFactor).Add(callCostEstimate)
			return &checker.CallEstimate{CostEstimate: costEstimate}
		}
		return nil
	}
}

func estimateSize(estimator checker.CostEstimator, node checker.AstNode) checker.SizeEstimate {
	if l := node.ComputedSize(); l != nil {
		return *l
	}
	if l := estimator.EstimateSize(node); l != nil {
		return *l
	}
	return checker.SizeEstimate{Min: 0, Max: math.MaxUint64}
}

func trackSetsCost(costFactor float64) interpreter.FunctionTracker {
	return func(args []ref.Val, _ ref.Val) *uint64 {
		lhsSize := actualSize(args[0])
		rhsSize := actualSize(args[1])
		cost := callCost + uint64(float64(lhsSize*rhsSize)*costFactor)
		return &cost
	}
}

func actualSize(value ref.Val) uint64 {
	if sz, ok := value.(traits.Sizer); ok {
		return uint64(sz.Size().(types.Int))
	}
	return 1
}

var (
	callCostEstimate = checker.CostEstimate{Min: 1, Max: 1}
	callCost         = uint64(1)
)
