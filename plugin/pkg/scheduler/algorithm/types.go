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

package algorithm

import (
	"k8s.io/kubernetes/pkg/api"
	schedulerapi "k8s.io/kubernetes/plugin/pkg/scheduler/api"
	"k8s.io/kubernetes/plugin/pkg/scheduler/schedulercache"
)

// FitPredicate is a function that indicates if a pod fits into an existing node.
// The failure information is given by the error.
type FitPredicate func(pod *api.Pod, meta interface{}, nodeInfo *schedulercache.NodeInfo) (bool, []PredicateFailureReason, error)

type PriorityFunction func(pod *api.Pod, nodeNameToInfo map[string]*schedulercache.NodeInfo, nodes []*api.Node) (schedulerapi.HostPriorityList, error)

type PriorityConfig struct {
	Function PriorityFunction
	Weight   int
}

type PredicateFailureReason interface {
	GetReason() string
}
