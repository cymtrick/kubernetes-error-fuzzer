/*
Copyright 2014 The Kubernetes Authors All rights reserved.

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

package resource

import (
	"fmt"
	"reflect"

	"k8s.io/kubernetes/pkg/api"
	"k8s.io/kubernetes/pkg/api/latest"
	"k8s.io/kubernetes/pkg/api/meta"
	"k8s.io/kubernetes/pkg/api/unversioned"
	"k8s.io/kubernetes/pkg/runtime"
	"k8s.io/kubernetes/pkg/util/errors"
	"k8s.io/kubernetes/pkg/util/sets"
	"k8s.io/kubernetes/pkg/watch"
)

// ErrMatchFunc can be used to filter errors that may not be true failures.
type ErrMatchFunc func(error) bool

// Result contains helper methods for dealing with the outcome of a Builder.
type Result struct {
	err     error
	visitor Visitor

	sources  []Visitor
	singular bool

	ignoreErrors []errors.Matcher

	// populated by a call to Infos
	info []*Info
}

// IgnoreErrors will filter errors that occur when by visiting the result
// (but not errors that occur by creating the result in the first place),
// eliminating any that match fns. This is best used in combination with
// Builder.ContinueOnError(), where the visitors accumulate errors and return
// them after visiting as a slice of errors. If no errors remain after
// filtering, the various visitor methods on Result will return nil for
// err.
func (r *Result) IgnoreErrors(fns ...ErrMatchFunc) *Result {
	for _, fn := range fns {
		r.ignoreErrors = append(r.ignoreErrors, errors.Matcher(fn))
	}
	return r
}

// Err returns one or more errors (via a util.ErrorList) that occurred prior
// to visiting the elements in the visitor. To see all errors including those
// that occur during visitation, invoke Infos().
func (r *Result) Err() error {
	return r.err
}

// Visit implements the Visitor interface on the items described in the Builder.
// Note that some visitor sources are not traversable more than once, or may
// return different results.  If you wish to operate on the same set of resources
// multiple times, use the Infos() method.
func (r *Result) Visit(fn VisitorFunc) error {
	if r.err != nil {
		return r.err
	}
	err := r.visitor.Visit(fn)
	return errors.FilterOut(err, r.ignoreErrors...)
}

// IntoSingular sets the provided boolean pointer to true if the Builder input
// reflected a single item, or multiple.
func (r *Result) IntoSingular(b *bool) *Result {
	*b = r.singular
	return r
}

// Infos returns an array of all of the resource infos retrieved via traversal.
// Will attempt to traverse the entire set of visitors only once, and will return
// a cached list on subsequent calls.
func (r *Result) Infos() ([]*Info, error) {
	if r.err != nil {
		return nil, r.err
	}
	if r.info != nil {
		return r.info, nil
	}

	infos := []*Info{}
	err := r.visitor.Visit(func(info *Info, err error) error {
		if err != nil {
			return err
		}
		infos = append(infos, info)
		return nil
	})
	err = errors.FilterOut(err, r.ignoreErrors...)

	r.info, r.err = infos, err
	return infos, err
}

// Object returns a single object representing the output of a single visit to all
// found resources.  If the Builder was a singular context (expected to return a
// single resource by user input) and only a single resource was found, the resource
// will be returned as is.  Otherwise, the returned resources will be part of an
// api.List. The ResourceVersion of the api.List will be set only if it is identical
// across all infos returned.
func (r *Result) Object() (runtime.Object, error) {
	infos, err := r.Infos()
	if err != nil {
		return nil, err
	}

	versions := sets.String{}
	objects := []runtime.Object{}
	for _, info := range infos {
		if info.Object != nil {
			objects = append(objects, info.Object)
			versions.Insert(info.ResourceVersion)
		}
	}

	if len(objects) == 1 {
		if r.singular {
			return objects[0], nil
		}
		// if the item is a list already, don't create another list
		if runtime.IsListType(objects[0]) {
			return objects[0], nil
		}
	}

	version := ""
	if len(versions) == 1 {
		version = versions.List()[0]
	}
	return &api.List{
		ListMeta: unversioned.ListMeta{
			ResourceVersion: version,
		},
		Items: objects,
	}, err
}

// ResourceMapping returns a single meta.RESTMapping representing the
// resources located by the builder, or an error if more than one
// mapping was found.
func (r *Result) ResourceMapping() (*meta.RESTMapping, error) {
	if r.err != nil {
		return nil, r.err
	}
	mappings := map[string]*meta.RESTMapping{}
	for i := range r.sources {
		m, ok := r.sources[i].(ResourceMapping)
		if !ok {
			return nil, fmt.Errorf("a resource mapping could not be loaded from %v", reflect.TypeOf(r.sources[i]))
		}
		mapping := m.ResourceMapping()
		mappings[mapping.Resource] = mapping
	}
	if len(mappings) != 1 {
		return nil, fmt.Errorf("expected only a single resource type")
	}
	for _, mapping := range mappings {
		return mapping, nil
	}
	return nil, nil
}

// Watch retrieves changes that occur on the server to the specified resource.
// It currently supports watching a single source - if the resource source
// (selectors or pure types) can be watched, they will be, otherwise the list
// will be visited (equivalent to the Infos() call) and if there is a single
// resource present, it will be watched, otherwise an error will be returned.
func (r *Result) Watch(resourceVersion string) (watch.Interface, error) {
	if r.err != nil {
		return nil, r.err
	}
	if len(r.sources) != 1 {
		return nil, fmt.Errorf("you may only watch a single resource or type of resource at a time")
	}
	w, ok := r.sources[0].(Watchable)
	if !ok {
		info, err := r.Infos()
		if err != nil {
			return nil, err
		}
		if len(info) != 1 {
			return nil, fmt.Errorf("watch is only supported on a single resource - %d resources were found", len(info))
		}
		return info[0].Watch(resourceVersion)
	}
	return w.Watch(resourceVersion)
}

// AsVersionedObject converts a list of infos into a single object - either a List containing
// the objects as children, or if only a single Object is present, as that object. The provided
// version will be preferred as the conversion target, but the Object's mapping version will be
// used if that version is not present.
func AsVersionedObject(infos []*Info, forceList bool, version string) (runtime.Object, error) {
	objects, err := AsVersionedObjects(infos, version)
	if err != nil {
		return nil, err
	}

	var object runtime.Object
	if len(objects) == 1 && !forceList {
		object = objects[0]
	} else {
		object = &api.List{Items: objects}
		converted, err := tryConvert(api.Scheme, object, version, latest.GroupOrDie("").Version)
		if err != nil {
			return nil, err
		}
		object = converted
	}
	return object, nil
}

// AsVersionedObjects converts a list of infos into versioned objects. The provided
// version will be preferred as the conversion target, but the Object's mapping version will be
// used if that version is not present.
func AsVersionedObjects(infos []*Info, version string) ([]runtime.Object, error) {
	objects := []runtime.Object{}
	for _, info := range infos {
		if info.Object == nil {
			continue
		}

		// objects that are not part of api.Scheme must be converted to JSON
		// TODO: convert to map[string]interface{}, attach to runtime.Unknown?
		if len(version) > 0 {
			if _, _, err := api.Scheme.ObjectVersionAndKind(info.Object); runtime.IsNotRegisteredError(err) {
				// TODO: ideally this would encode to version, but we don't expose multiple codecs here.
				data, err := info.Mapping.Codec.Encode(info.Object)
				if err != nil {
					return nil, err
				}
				objects = append(objects, &runtime.Unknown{RawJSON: data})
				continue
			}
		}

		converted, err := tryConvert(info.Mapping.ObjectConvertor, info.Object, version, info.Mapping.APIVersion)
		if err != nil {
			return nil, err
		}
		objects = append(objects, converted)
	}
	return objects, nil
}

// tryConvert attempts to convert the given object to the provided versions in order. This function assumes
// the object is in internal version.
func tryConvert(convertor runtime.ObjectConvertor, object runtime.Object, versions ...string) (runtime.Object, error) {
	var last error
	for _, version := range versions {
		if len(version) == 0 {
			return object, nil
		}
		obj, err := convertor.ConvertToVersion(object, version)
		if err != nil {
			last = err
			continue
		}
		return obj, nil
	}
	return nil, last
}
