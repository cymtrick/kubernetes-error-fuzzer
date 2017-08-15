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

package parse

import (
	"fmt"
	"reflect"

	"k8s.io/kubernetes/pkg/kubectl/apply"
	"k8s.io/kubernetes/pkg/kubectl/cmd/util/openapi"
)

// Factory creates an Element by combining object values from recorded, local and remote sources with
// the metadata from an openapi schema.
type Factory struct {
	// Resources contains the openapi field metadata for the object models
	Resources openapi.Resources
}

// CreateElement returns an Element by collating the recorded, local and remote field values
func (b *Factory) CreateElement(recorded, local, remote map[string]interface{}) (apply.Element, error) {
	// Create an Item from the 3 values.  Use empty name for field
	visitor := &ElementBuildingVisitor{b.Resources}

	gvk, err := getCommonGroupVersionKind(recorded, local, remote)
	if err != nil {
		return nil, err
	}

	// Get the openapi object metadata
	s := visitor.resources.LookupResource(gvk)
	oapiKind, err := getKind(s)
	if err != nil {
		return nil, err
	}

	// Get the item for the type
	hasRecorded := recorded != nil
	hasLocal := local != nil
	hasRemote := remote != nil
	fieldName := ""
	item, err := visitor.getItem(oapiKind, fieldName,
		apply.RawElementData{recorded, local, remote},
		apply.HasElementData{hasRecorded, hasLocal, hasRemote})
	if err != nil {
		return nil, err
	}

	// Collate each field of the item into a combined Element
	return item.CreateElement(visitor)
}

// getItem returns the appropriate Item based on the underlying type of the arguments
func (v *ElementBuildingVisitor) getItem(s openapi.Schema, name string,
	data apply.RawElementData,
	isSet apply.HasElementData) (Item, error) {
	kind, err := getType(data.Recorded, data.Local, data.Remote)
	if err != nil {
		return nil, err
	}
	if kind == nil {
		// All of the items values are nil.
		return &emptyItem{Name: name}, nil
	}

	// Create an item matching the type
	switch kind.Kind() {
	case reflect.Bool, reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64, reflect.Uint,
		reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Float32, reflect.Float64,
		reflect.String:
		p, err := getPrimitive(s)
		if err != nil {
			return nil, fmt.Errorf("expected openapi Primitive, was %T for %v", s, kind)
		}
		return &primitiveItem{name, p, isSet, data}, nil
	case reflect.Array, reflect.Slice:
		a, err := getArray(s)
		if err != nil {
			return nil, fmt.Errorf("expected openapi Array, was %T for %v", s, kind)
		}
		return &listItem{name, a, isSet,
			apply.ListElementData{
				sliceCast(data.Recorded),
				sliceCast(data.Local),
				sliceCast(data.Remote),
			}}, nil
	case reflect.Map:
		if k, err := getKind(s); err == nil {
			return &typeItem{name, k, isSet,
				apply.MapElementData{
					mapCast(data.Recorded),
					mapCast(data.Local),
					mapCast(data.Remote),
				}}, nil
		}
		// If it looks like a map, and no openapi type is found, default to mapItem
		m, err := getMap(s)
		if err != nil {
			return nil, fmt.Errorf("expected openapi Kind or Map, was %T for %v", s, kind)
		}
		return &mapItem{name, m, isSet,
			apply.MapElementData{
				mapCast(data.Recorded),
				mapCast(data.Local),
				mapCast(data.Remote),
			}}, nil
	}
	return nil, fmt.Errorf("unsupported type type %v", kind)
}
