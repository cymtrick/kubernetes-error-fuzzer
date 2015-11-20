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

// A set of common functions needed by cmd/kubectl and pkg/kubectl packages.
package kubectl

import (
	"strings"

	"k8s.io/kubernetes/pkg/api"
	"k8s.io/kubernetes/pkg/api/meta"
)

const kubectlAnnotationPrefix = "kubectl.kubernetes.io/"

type NamespaceInfo struct {
	Namespace string
}

func listOfImages(spec *api.PodSpec) []string {
	var images []string
	for _, container := range spec.Containers {
		images = append(images, container.Image)
	}
	return images
}

func makeImageList(spec *api.PodSpec) string {
	return strings.Join(listOfImages(spec), ",")
}

// OutputVersionMapper is a RESTMapper that will prefer mappings that
// correspond to a preferred output version (if feasible)
type OutputVersionMapper struct {
	meta.RESTMapper
	OutputVersion string
}

// RESTMapping implements meta.RESTMapper by prepending the output version to the preferred version list.
func (m OutputVersionMapper) RESTMapping(kind string, versions ...string) (*meta.RESTMapping, error) {
	mapping, err := m.RESTMapper.RESTMapping(kind, m.OutputVersion)
	if err == nil {
		return mapping, nil
	}

	return m.RESTMapper.RESTMapping(kind, versions...)
}

// ShortcutExpander is a RESTMapper that can be used for Kubernetes
// resources.
type ShortcutExpander struct {
	meta.RESTMapper
}

// VersionAndKindForResource implements meta.RESTMapper. It expands the resource first, then invokes the wrapped
// mapper.
func (e ShortcutExpander) VersionAndKindForResource(resource string) (defaultVersion, kind string, err error) {
	resource = expandResourceShortcut(resource)
	defaultVersion, kind, err = e.RESTMapper.VersionAndKindForResource(resource)
	return defaultVersion, kind, err
}

// ResourceIsValid takes a string (kind) and checks if it's a valid resource.
// It expands the resource first, then invokes the wrapped mapper.
func (e ShortcutExpander) ResourceIsValid(resource string) bool {
	return e.RESTMapper.ResourceIsValid(expandResourceShortcut(resource))
}

// expandResourceShortcut will return the expanded version of resource
// (something that a pkg/api/meta.RESTMapper can understand), if it is
// indeed a shortcut. Otherwise, will return resource unmodified.
func expandResourceShortcut(resource string) string {
	shortForms := map[string]string{
		// Please keep this alphabetized
		"cs":     "componentstatuses",
		"ev":     "events",
		"ep":     "endpoints",
		"hpa":    "horizontalpodautoscalers",
		"limits": "limitranges",
		"no":     "nodes",
		"ns":     "namespaces",
		"po":     "pods",
		"pv":     "persistentvolumes",
		"pvc":    "persistentvolumeclaims",
		"quota":  "resourcequotas",
		"rc":     "replicationcontrollers",
		"ds":     "daemonsets",
		"svc":    "services",
		"ing":    "ingresses",
	}
	if expanded, ok := shortForms[resource]; ok {
		return expanded
	}
	return resource
}
