/*
Copyright 2019 The Kubernetes Authors.

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

package prebind

import (
	"k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/runtime"

	framework "k8s.io/kubernetes/pkg/scheduler/framework/v1alpha1"
)

// NoOpFilter is a plugin that implements the filter plugin and always returns Success.
// This is just to make sure that all code dependencies are properly addressed while
// working on legitimate plugins.
type NoOpFilter struct{}

var _ = framework.FilterPlugin(NoOpFilter{})

// Name is the name of the plugin used in Registry and configurations.
const Name = "noop-filter"

// Name returns name of the plugin. It is used in logs, etc.
func (n NoOpFilter) Name() string {
	return Name
}

// Filter invoked at the filter extension point.
func (n NoOpFilter) Filter(pc *framework.PluginContext, pod *v1.Pod, nodeName string) *framework.Status {
	return nil
}

// New initializes a new plugin and returns it.
func New(_ *runtime.Unknown, _ framework.FrameworkHandle) (framework.Plugin, error) {
	return &NoOpFilter{}, nil
}
