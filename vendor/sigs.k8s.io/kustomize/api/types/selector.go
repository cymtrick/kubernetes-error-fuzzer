// Copyright 2019 The Kubernetes Authors.
// SPDX-License-Identifier: Apache-2.0

package types

import (
	"regexp"

	"sigs.k8s.io/kustomize/api/resid"
)

// Selector specifies a set of resources.
// Any resource that matches intersection of all conditions
// is included in this set.
type Selector struct {
	// KrmId refers to a GVKN/Ns of a resource.
	KrmId `json:",inline,omitempty" yaml:",inline,omitempty"`

	// AnnotationSelector is a string that follows the label selection expression
	// https://kubernetes.io/docs/concepts/overview/working-with-objects/labels/#api
	// It matches with the resource annotations.
	AnnotationSelector string `json:"annotationSelector,omitempty" yaml:"annotationSelector,omitempty"`

	// LabelSelector is a string that follows the label selection expression
	// https://kubernetes.io/docs/concepts/overview/working-with-objects/labels/#api
	// It matches with the resource labels.
	LabelSelector string `json:"labelSelector,omitempty" yaml:"labelSelector,omitempty"`
}

// KrmId refers to a GVKN/Ns of a resource.
type KrmId struct {
	resid.Gvk `json:",inline,omitempty" yaml:",inline,omitempty"`
	Name      string `json:"name,omitempty" yaml:"name,omitempty"`
	Namespace string `json:"namespace,omitempty" yaml:"namespace,omitempty"`
}

// Match returns true if id selects other, i.e. id's fields
// either match other's or are empty
func (id *KrmId) Match(other *KrmId) bool {
	return (id.Group == "" || id.Group == other.Group) &&
		(id.Version == "" || id.Version == other.Version) &&
		(id.Kind == "" || id.Kind == other.Kind) &&
		(id.Name == "" || id.Name == other.Name) &&
		(id.Namespace == "" || id.Namespace == other.Namespace)
}

// SelectorRegex is a Selector with regex in GVK
// Any resource that matches intersection of all conditions
// is included in this set.
type SelectorRegex struct {
	selector       *Selector
	groupRegex     *regexp.Regexp
	versionRegex   *regexp.Regexp
	kindRegex      *regexp.Regexp
	nameRegex      *regexp.Regexp
	namespaceRegex *regexp.Regexp
}

// NewSelectorRegex returns a pointer to a new SelectorRegex
// which uses the same condition as s.
func NewSelectorRegex(s *Selector) (*SelectorRegex, error) {
	sr := new(SelectorRegex)
	var err error
	sr.selector = s
	sr.groupRegex, err = regexp.Compile(anchorRegex(s.Gvk.Group))
	if err != nil {
		return nil, err
	}
	sr.versionRegex, err = regexp.Compile(anchorRegex(s.Gvk.Version))
	if err != nil {
		return nil, err
	}
	sr.kindRegex, err = regexp.Compile(anchorRegex(s.Gvk.Kind))
	if err != nil {
		return nil, err
	}
	sr.nameRegex, err = regexp.Compile(anchorRegex(s.Name))
	if err != nil {
		return nil, err
	}
	sr.namespaceRegex, err = regexp.Compile(anchorRegex(s.Namespace))
	if err != nil {
		return nil, err
	}
	return sr, nil
}

func anchorRegex(pattern string) string {
	if pattern == "" {
		return pattern
	}
	return "^(?:" + pattern + ")$"
}

// MatchGvk return true if gvk can be matched by s.
func (s *SelectorRegex) MatchGvk(gvk resid.Gvk) bool {
	if len(s.selector.Gvk.Group) > 0 {
		if !s.groupRegex.MatchString(gvk.Group) {
			return false
		}
	}
	if len(s.selector.Gvk.Version) > 0 {
		if !s.versionRegex.MatchString(gvk.Version) {
			return false
		}
	}
	if len(s.selector.Gvk.Kind) > 0 {
		if !s.kindRegex.MatchString(gvk.Kind) {
			return false
		}
	}
	return true
}

// MatchName returns true if the name in selector is
// empty or the n can be matches by the name in selector
func (s *SelectorRegex) MatchName(n string) bool {
	if s.selector.Name == "" {
		return true
	}
	return s.nameRegex.MatchString(n)
}

// MatchNamespace returns true if the namespace in selector is
// empty or the ns can be matches by the namespace in selector
func (s *SelectorRegex) MatchNamespace(ns string) bool {
	if s.selector.Namespace == "" {
		return true
	}
	return s.namespaceRegex.MatchString(ns)
}
