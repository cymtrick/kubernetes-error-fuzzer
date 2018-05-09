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

package restmapper

import (
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/client-go/discovery"
)

// CategoryExpander maps category strings to GroupResouces.
// Categories are classification or 'tag' of a group of resources.
type CategoryExpander interface {
	Expand(category string) ([]schema.GroupResource, bool)
}

// SimpleCategoryExpander implements CategoryExpander interface
// using a static mapping of categories to GroupResource mapping.
type SimpleCategoryExpander struct {
	Expansions map[string][]schema.GroupResource
}

func (e SimpleCategoryExpander) Expand(category string) ([]schema.GroupResource, bool) {
	ret, ok := e.Expansions[category]
	return ret, ok
}

// discoveryCategoryExpander struct lets a REST Client wrapper (discoveryClient) to retrieve list of APIResourceList,
// and then convert to fallbackExpander
type discoveryCategoryExpander struct {
	fallbackExpander CategoryExpander
	discoveryClient  discovery.DiscoveryInterface
}

// NewDiscoveryCategoryExpander returns a category expander that makes use of the "categories" fields from
// the API, found through the discovery client. In case of any error or no category found (which likely
// means we're at a cluster prior to categories support, fallback to the expander provided.
func NewDiscoveryCategoryExpander(fallbackExpander CategoryExpander, client discovery.DiscoveryInterface) (discoveryCategoryExpander, error) {
	if client == nil {
		panic("Please provide discovery client to shortcut expander")
	}
	return discoveryCategoryExpander{fallbackExpander: fallbackExpander, discoveryClient: client}, nil
}

func (e discoveryCategoryExpander) Expand(category string) ([]schema.GroupResource, bool) {
	// Get all supported resources for groups and versions from server, if no resource found, fallback anyway.
	apiResourceLists, _ := e.discoveryClient.ServerResources()
	if len(apiResourceLists) == 0 {
		return e.fallbackExpander.Expand(category)
	}

	discoveredExpansions := map[string][]schema.GroupResource{}

	for _, apiResourceList := range apiResourceLists {
		gv, err := schema.ParseGroupVersion(apiResourceList.GroupVersion)
		if err != nil {
			return e.fallbackExpander.Expand(category)
		}
		// Collect GroupVersions by categories
		for _, apiResource := range apiResourceList.APIResources {
			if categories := apiResource.Categories; len(categories) > 0 {
				for _, category := range categories {
					groupResource := schema.GroupResource{
						Group:    gv.Group,
						Resource: apiResource.Name,
					}
					discoveredExpansions[category] = append(discoveredExpansions[category], groupResource)
				}
			}
		}
	}

	if len(discoveredExpansions) == 0 {
		// We don't know if the server really don't have any resource with categories,
		// or we're on a cluster version prior to categories support. Anyways, fallback.
		return e.fallbackExpander.Expand(category)
	}

	ret, ok := discoveredExpansions[category]
	return ret, ok
}

// discoveryFilteredExpander expands the given CategoryExpander (delegate) to filter group and resource returned from server
type discoveryFilteredExpander struct {
	delegate CategoryExpander

	discoveryClient discovery.DiscoveryInterface
}

// NewDiscoveryFilteredExpander returns a category expander that filters the returned groupresources
// by what the server has available
func NewDiscoveryFilteredExpander(delegate CategoryExpander, client discovery.DiscoveryInterface) (discoveryFilteredExpander, error) {
	if client == nil {
		panic("Please provide discovery client to shortcut expander")
	}
	return discoveryFilteredExpander{delegate: delegate, discoveryClient: client}, nil
}

func (e discoveryFilteredExpander) Expand(category string) ([]schema.GroupResource, bool) {
	delegateExpansion, ok := e.delegate.Expand(category)

	// Check if we have access to server resources
	apiResources, err := e.discoveryClient.ServerResources()
	if err != nil {
		return delegateExpansion, ok
	}

	availableResources, err := discovery.GroupVersionResources(apiResources)
	if err != nil {
		return delegateExpansion, ok
	}

	available := []schema.GroupResource{}
	for _, requestedResource := range delegateExpansion {
		for availableResource := range availableResources {
			if requestedResource.Group == availableResource.Group &&
				requestedResource.Resource == availableResource.Resource {
				available = append(available, requestedResource)
				break
			}
		}
	}

	return available, ok
}

// UnionCategoryExpander implements CategoryExpander interface.
// It maps given category string to union of expansions returned by all the CategoryExpanders in the list.
type UnionCategoryExpander []CategoryExpander

func (u UnionCategoryExpander) Expand(category string) ([]schema.GroupResource, bool) {
	ret := []schema.GroupResource{}
	ok := false

	// Expand the category for each CategoryExpander in the list and merge/combine the results.
	for _, expansion := range u {
		curr, currOk := expansion.Expand(category)

		for _, currGR := range curr {
			found := false
			for _, existing := range ret {
				if existing == currGR {
					found = true
					break
				}
			}
			if !found {
				ret = append(ret, currGR)
			}
		}
		ok = ok || currOk
	}

	return ret, ok
}
