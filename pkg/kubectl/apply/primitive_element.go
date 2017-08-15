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

package apply

import (
	"fmt"
)

// PrimitiveElement contains the recorded, local and remote values for a field
// of type primitive
type PrimitiveElement struct {
	// FieldMetaImpl contains metadata about the field from openapi
	FieldMetaImpl

	// HasElementData contains whether the field was set
	HasElementData

	// RawElementData contains the values the field was set to
	RawElementData
}

// Merge implements Element.Merge
func (e PrimitiveElement) Merge(v Strategy) (Result, error) {
	return v.MergePrimitive(e)
}

// String returns a string representation of the PrimitiveElement
func (e PrimitiveElement) String() string {
	return fmt.Sprintf("name: %s recorded: %v local: %v remote: %v", e.Name, e.Recorded, e.Local, e.Remote)
}

var _ Element = &PrimitiveElement{}
