/*
Copyright 2015 The Kubernetes Authors.

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

package meta

import (
	"errors"
	"fmt"
	"reflect"
	"sync"

	"k8s.io/apimachinery/pkg/conversion"
	"k8s.io/apimachinery/pkg/runtime"
)

var (
	// isListCache maintains a cache of types that are checked for lists
	// which is used by IsListType.
	// TODO: remove and replace with an interface check
	isListCache = struct {
		lock   sync.RWMutex
		byType map[reflect.Type]bool
	}{
		byType: make(map[reflect.Type]bool, 1024),
	}
)

// IsListType returns true if the provided Object has a slice called Items.
// TODO: Replace the code in this check with an interface comparison by
//   creating and enforcing that lists implement a list accessor.
func IsListType(obj runtime.Object) bool {
	switch t := obj.(type) {
	case runtime.Unstructured:
		return t.IsList()
	}
	t := reflect.TypeOf(obj)

	isListCache.lock.RLock()
	ok, exists := isListCache.byType[t]
	isListCache.lock.RUnlock()

	if !exists {
		_, err := getItemsPtr(obj)
		ok = err == nil

		// cache only the first 1024 types
		isListCache.lock.Lock()
		if len(isListCache.byType) < 1024 {
			isListCache.byType[t] = ok
		}
		isListCache.lock.Unlock()
	}

	return ok
}

var (
	errExpectFieldItems = errors.New("no Items field in this object")
	errExpectSliceItems = errors.New("Items field must be a slice of objects")
)

// GetItemsPtr returns a pointer to the list object's Items member.
// If 'list' doesn't have an Items member, it's not really a list type
// and an error will be returned.
// This function will either return a pointer to a slice, or an error, but not both.
// TODO: this will be replaced with an interface in the future
func GetItemsPtr(list runtime.Object) (interface{}, error) {
	obj, err := getItemsPtr(list)
	if err != nil {
		return nil, fmt.Errorf("%T is not a list: %v", err)
	}
	return obj, nil
}

// getItemsPtr returns a pointer to the list object's Items member or an error.
func getItemsPtr(list runtime.Object) (interface{}, error) {
	v, err := conversion.EnforcePtr(list)
	if err != nil {
		return nil, err
	}

	items := v.FieldByName("Items")
	if !items.IsValid() {
		return nil, errExpectFieldItems
	}
	switch items.Kind() {
	case reflect.Interface, reflect.Ptr:
		target := reflect.TypeOf(items.Interface()).Elem()
		if target.Kind() != reflect.Slice {
			return nil, errExpectSliceItems
		}
		return items.Interface(), nil
	case reflect.Slice:
		return items.Addr().Interface(), nil
	default:
		return nil, errExpectSliceItems
	}
}

// EachListItem invokes fn on each runtime.Object in the list. Any error immediately terminates
// the loop.
func EachListItem(obj runtime.Object, fn func(runtime.Object) error) error {
	if unstructured, ok := obj.(runtime.Unstructured); ok {
		return unstructured.EachListItem(fn)
	}
	// TODO: Change to an interface call?
	itemsPtr, err := GetItemsPtr(obj)
	if err != nil {
		return err
	}
	items, err := conversion.EnforcePtr(itemsPtr)
	if err != nil {
		return err
	}
	len := items.Len()
	if len == 0 {
		return nil
	}
	takeAddr := false
	if elemType := items.Type().Elem(); elemType.Kind() != reflect.Ptr && elemType.Kind() != reflect.Interface {
		if !items.Index(0).CanAddr() {
			return fmt.Errorf("unable to take address of items in %T for EachListItem", obj)
		}
		takeAddr = true
	}

	for i := 0; i < len; i++ {
		raw := items.Index(i)
		if takeAddr {
			raw = raw.Addr()
		}
		switch item := raw.Interface().(type) {
		case *runtime.RawExtension:
			if err := fn(item.Object); err != nil {
				return err
			}
		case runtime.Object:
			if err := fn(item); err != nil {
				return err
			}
		default:
			obj, ok := item.(runtime.Object)
			if !ok {
				return fmt.Errorf("%v: item[%v]: Expected object, got %#v(%s)", obj, i, raw.Interface(), raw.Kind())
			}
			if err := fn(obj); err != nil {
				return err
			}
		}
	}
	return nil
}

// ExtractList returns obj's Items element as an array of runtime.Objects.
// Returns an error if obj is not a List type (does not have an Items member).
func ExtractList(obj runtime.Object) ([]runtime.Object, error) {
	itemsPtr, err := GetItemsPtr(obj)
	if err != nil {
		return nil, err
	}
	items, err := conversion.EnforcePtr(itemsPtr)
	if err != nil {
		return nil, err
	}
	list := make([]runtime.Object, items.Len())
	for i := range list {
		raw := items.Index(i)
		switch item := raw.Interface().(type) {
		case runtime.RawExtension:
			switch {
			case item.Object != nil:
				list[i] = item.Object
			case item.Raw != nil:
				// TODO: Set ContentEncoding and ContentType correctly.
				list[i] = &runtime.Unknown{Raw: item.Raw}
			default:
				list[i] = nil
			}
		case runtime.Object:
			list[i] = item
		default:
			var found bool
			if list[i], found = raw.Addr().Interface().(runtime.Object); !found {
				return nil, fmt.Errorf("%v: item[%v]: Expected object, got %#v(%s)", obj, i, raw.Interface(), raw.Kind())
			}
		}
	}
	return list, nil
}

// objectSliceType is the type of a slice of Objects
var objectSliceType = reflect.TypeOf([]runtime.Object{})

// LenList returns the length of this list or 0 if it is not a list.
func LenList(list runtime.Object) int {
	itemsPtr, err := GetItemsPtr(list)
	if err != nil {
		return 0
	}
	items, err := conversion.EnforcePtr(itemsPtr)
	if err != nil {
		return 0
	}
	return items.Len()
}

// SetList sets the given list object's Items member have the elements given in
// objects.
// Returns an error if list is not a List type (does not have an Items member),
// or if any of the objects are not of the right type.
func SetList(list runtime.Object, objects []runtime.Object) error {
	itemsPtr, err := GetItemsPtr(list)
	if err != nil {
		return err
	}
	items, err := conversion.EnforcePtr(itemsPtr)
	if err != nil {
		return err
	}
	if items.Type() == objectSliceType {
		items.Set(reflect.ValueOf(objects))
		return nil
	}
	slice := reflect.MakeSlice(items.Type(), len(objects), len(objects))
	for i := range objects {
		dest := slice.Index(i)
		if dest.Type() == reflect.TypeOf(runtime.RawExtension{}) {
			dest = dest.FieldByName("Object")
		}

		// check to see if you're directly assignable
		if reflect.TypeOf(objects[i]).AssignableTo(dest.Type()) {
			dest.Set(reflect.ValueOf(objects[i]))
			continue
		}

		src, err := conversion.EnforcePtr(objects[i])
		if err != nil {
			return err
		}
		if src.Type().AssignableTo(dest.Type()) {
			dest.Set(src)
		} else if src.Type().ConvertibleTo(dest.Type()) {
			dest.Set(src.Convert(dest.Type()))
		} else {
			return fmt.Errorf("item[%d]: can't assign or convert %v into %v", i, src.Type(), dest.Type())
		}
	}
	items.Set(slice)
	return nil
}
