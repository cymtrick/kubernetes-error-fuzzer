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
	"k8s.io/kubernetes/pkg/kubectl/apply"
	"k8s.io/kubernetes/pkg/kubectl/cmd/util/openapi"
)

// typeElement builds a new mapElement from a typeItem
func (v ElementBuildingVisitor) typeElement(meta apply.FieldMetaImpl, item *typeItem) (*apply.TypeElement, error) {
	// Function to get the schema of a field from its key
	var fn schemaFn = func(key string) openapi.Schema {
		if item.Type != nil && item.Type.Fields != nil {
			return item.Type.Fields[key]
		}
		return nil
	}

	// Collect same fields from multiple maps into a map of elements
	values, err := v.createMapValues(fn, meta, item.HasElementData, item.MapElementData)
	if err != nil {
		return nil, err
	}

	// Return the result
	return &apply.TypeElement{
		FieldMetaImpl:  meta,
		HasElementData: item.HasElementData,
		MapElementData: item.MapElementData,
		Values:         values,
	}, nil
}
