/*
Copyright 2016 The Kubernetes Authors All rights reserved.

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

package capabilities

import (
	"k8s.io/kubernetes/pkg/api"
	"k8s.io/kubernetes/pkg/util/validation/field"
)

// CapabilitiesStrategy defines the interface for all cap constraint strategies.
type CapabilitiesStrategy interface {
	// Generate creates the capabilities based on policy rules.
	Generate(pod *api.Pod, container *api.Container) (*api.Capabilities, error)
	// Validate ensures that the specified values fall within the range of the strategy.
	Validate(pod *api.Pod, container *api.Container) field.ErrorList
}
