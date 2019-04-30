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

package stateful

import (
	"fmt"

	"k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/klog"
	framework "k8s.io/kubernetes/pkg/scheduler/framework/v1alpha1"
)

// MultipointExample is an example plugin that is executed at multiple extension points.
// This plugin is stateful. It receives arguments at initialization (NewMultipointPlugin)
// and changes its state when it is executed.
type MultipointExample struct {
	mpState map[int]string
	numRuns int
}

var _ = framework.ReservePlugin(&MultipointExample{})
var _ = framework.PrebindPlugin(&MultipointExample{})

// Name is the name of the plug used in Registry and configurations.
const Name = "multipoint-plugin-example"

// Name returns name of the plugin. It is used in logs, etc.
func (mp *MultipointExample) Name() string {
	return Name
}

// Reserve is the functions invoked by the framework at "reserve" extension point.
func (mp *MultipointExample) Reserve(pc *framework.PluginContext, pod *v1.Pod, nodeName string) *framework.Status {
	mp.numRuns++
	return nil
}

// Prebind is the functions invoked by the framework at "prebind" extension point.
func (mp *MultipointExample) Prebind(pc *framework.PluginContext, pod *v1.Pod, nodeName string) *framework.Status {
	mp.numRuns++
	if pod == nil {
		return framework.NewStatus(framework.Error, "pod must not be nil")
	}
	return nil
}

// New initializes a new plugin and returns it.
func New(config *runtime.Unknown, _ framework.FrameworkHandle) (framework.Plugin, error) {
	if config == nil {
		klog.Error("MultipointExample configuration cannot be empty")
		return nil, fmt.Errorf("MultipointExample configuration cannot be empty")
	}
	mp := MultipointExample{
		mpState: make(map[int]string),
	}
	return &mp, nil
}
