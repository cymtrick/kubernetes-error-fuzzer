// Copyright 2018 Google LLC
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

package traits

import "github.com/google/cel-go/common/types/ref"

// Lister interface which aggregates the traits of a list.
type Lister interface {
	ref.Val
	Adder
	Container
	Indexer
	Iterable
	Sizer
}

// MutableLister interface which emits an immutable result after an intermediate computation.
type MutableLister interface {
	Lister
	ToImmutableList() Lister
}
