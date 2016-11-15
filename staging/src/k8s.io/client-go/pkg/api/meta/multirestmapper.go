/*
Copyright 2014 The Kubernetes Authors.

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
	"fmt"
	"strings"

	"k8s.io/client-go/pkg/api/unversioned"
	utilerrors "k8s.io/client-go/pkg/util/errors"
	"k8s.io/client-go/pkg/util/sets"
)

// MultiRESTMapper is a wrapper for multiple RESTMappers.
type MultiRESTMapper []RESTMapper

func (m MultiRESTMapper) String() string {
	nested := []string{}
	for _, t := range m {
		currString := fmt.Sprintf("%v", t)
		splitStrings := strings.Split(currString, "\n")
		nested = append(nested, strings.Join(splitStrings, "\n\t"))
	}

	return fmt.Sprintf("MultiRESTMapper{\n\t%s\n}", strings.Join(nested, "\n\t"))
}

// ResourceSingularizer converts a REST resource name from plural to singular (e.g., from pods to pod)
// This implementation supports multiple REST schemas and return the first match.
func (m MultiRESTMapper) ResourceSingularizer(resource string) (singular string, err error) {
	for _, t := range m {
		singular, err = t.ResourceSingularizer(resource)
		if err == nil {
			return
		}
	}
	return
}

func (m MultiRESTMapper) ResourcesFor(resource unversioned.GroupVersionResource) ([]unversioned.GroupVersionResource, error) {
	allGVRs := []unversioned.GroupVersionResource{}
	for _, t := range m {
		gvrs, err := t.ResourcesFor(resource)
		// ignore "no match" errors, but any other error percolates back up
		if IsNoMatchError(err) {
			continue
		}
		if err != nil {
			return nil, err
		}

		// walk the existing values to de-dup
		for _, curr := range gvrs {
			found := false
			for _, existing := range allGVRs {
				if curr == existing {
					found = true
					break
				}
			}

			if !found {
				allGVRs = append(allGVRs, curr)
			}
		}
	}

	if len(allGVRs) == 0 {
		return nil, &NoResourceMatchError{PartialResource: resource}
	}

	return allGVRs, nil
}

func (m MultiRESTMapper) KindsFor(resource unversioned.GroupVersionResource) (gvk []unversioned.GroupVersionKind, err error) {
	allGVKs := []unversioned.GroupVersionKind{}
	for _, t := range m {
		gvks, err := t.KindsFor(resource)
		// ignore "no match" errors, but any other error percolates back up
		if IsNoMatchError(err) {
			continue
		}
		if err != nil {
			return nil, err
		}

		// walk the existing values to de-dup
		for _, curr := range gvks {
			found := false
			for _, existing := range allGVKs {
				if curr == existing {
					found = true
					break
				}
			}

			if !found {
				allGVKs = append(allGVKs, curr)
			}
		}
	}

	if len(allGVKs) == 0 {
		return nil, &NoResourceMatchError{PartialResource: resource}
	}

	return allGVKs, nil
}

func (m MultiRESTMapper) ResourceFor(resource unversioned.GroupVersionResource) (unversioned.GroupVersionResource, error) {
	resources, err := m.ResourcesFor(resource)
	if err != nil {
		return unversioned.GroupVersionResource{}, err
	}
	if len(resources) == 1 {
		return resources[0], nil
	}

	return unversioned.GroupVersionResource{}, &AmbiguousResourceError{PartialResource: resource, MatchingResources: resources}
}

func (m MultiRESTMapper) KindFor(resource unversioned.GroupVersionResource) (unversioned.GroupVersionKind, error) {
	kinds, err := m.KindsFor(resource)
	if err != nil {
		return unversioned.GroupVersionKind{}, err
	}
	if len(kinds) == 1 {
		return kinds[0], nil
	}

	return unversioned.GroupVersionKind{}, &AmbiguousResourceError{PartialResource: resource, MatchingKinds: kinds}
}

// RESTMapping provides the REST mapping for the resource based on the
// kind and version. This implementation supports multiple REST schemas and
// return the first match.
func (m MultiRESTMapper) RESTMapping(gk unversioned.GroupKind, versions ...string) (*RESTMapping, error) {
	allMappings := []*RESTMapping{}
	errors := []error{}

	for _, t := range m {
		currMapping, err := t.RESTMapping(gk, versions...)
		// ignore "no match" errors, but any other error percolates back up
		if IsNoMatchError(err) {
			continue
		}
		if err != nil {
			errors = append(errors, err)
			continue
		}

		allMappings = append(allMappings, currMapping)
	}

	// if we got exactly one mapping, then use it even if other requested failed
	if len(allMappings) == 1 {
		return allMappings[0], nil
	}
	if len(allMappings) > 1 {
		var kinds []unversioned.GroupVersionKind
		for _, m := range allMappings {
			kinds = append(kinds, m.GroupVersionKind)
		}
		return nil, &AmbiguousKindError{PartialKind: gk.WithVersion(""), MatchingKinds: kinds}
	}
	if len(errors) > 0 {
		return nil, utilerrors.NewAggregate(errors)
	}
	return nil, &NoKindMatchError{PartialKind: gk.WithVersion("")}
}

// RESTMappings returns all possible RESTMappings for the provided group kind, or an error
// if the type is not recognized.
func (m MultiRESTMapper) RESTMappings(gk unversioned.GroupKind) ([]*RESTMapping, error) {
	var allMappings []*RESTMapping
	var errors []error

	for _, t := range m {
		currMappings, err := t.RESTMappings(gk)
		// ignore "no match" errors, but any other error percolates back up
		if IsNoMatchError(err) {
			continue
		}
		if err != nil {
			errors = append(errors, err)
			continue
		}
		allMappings = append(allMappings, currMappings...)
	}
	if len(errors) > 0 {
		return nil, utilerrors.NewAggregate(errors)
	}
	if len(allMappings) == 0 {
		return nil, &NoKindMatchError{PartialKind: gk.WithVersion("")}
	}
	return allMappings, nil
}

// AliasesForResource finds the first alias response for the provided mappers.
func (m MultiRESTMapper) AliasesForResource(alias string) ([]string, bool) {
	seenAliases := sets.NewString()
	allAliases := []string{}
	handled := false

	for _, t := range m {
		if currAliases, currOk := t.AliasesForResource(alias); currOk {
			for _, currAlias := range currAliases {
				if !seenAliases.Has(currAlias) {
					allAliases = append(allAliases, currAlias)
					seenAliases.Insert(currAlias)
				}
			}
			handled = true
		}
	}
	return allAliases, handled
}
