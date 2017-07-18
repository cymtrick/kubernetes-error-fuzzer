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

// This file was automatically generated by informer-gen

package v1alpha1

import (
	internalinterfaces "k8s.io/sample-apiserver/pkg/client/informers_generated/externalversions/internalinterfaces"
)

// Interface provides access to all the informers in this group version.
type Interface interface {
	// Fischers returns a FischerInformer.
	Fischers() FischerInformer
	// Flunders returns a FlunderInformer.
	Flunders() FlunderInformer
}

type version struct {
	internalinterfaces.SharedInformerFactory
}

// New returns a new Interface.
func New(f internalinterfaces.SharedInformerFactory) Interface {
	return &version{f}
}

// Fischers returns a FischerInformer.
func (v *version) Fischers() FischerInformer {
	return &fischerInformer{factory: v.SharedInformerFactory}
}

// Flunders returns a FlunderInformer.
func (v *version) Flunders() FlunderInformer {
	return &flunderInformer{factory: v.SharedInformerFactory}
}
