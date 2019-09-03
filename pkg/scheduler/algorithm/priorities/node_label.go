/*
Copyright 2016 The Kubernetes Authors.

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

package priorities

import (
	"fmt"

	"k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/labels"
	schedulerapi "k8s.io/kubernetes/pkg/scheduler/api"
	schedulernodeinfo "k8s.io/kubernetes/pkg/scheduler/nodeinfo"
)

// NodeLabelPrioritizer contains information to calculate node label priority.
type NodeLabelPrioritizer struct {
	label    string
	presence bool
}

// NewNodeLabelPriority creates a NodeLabelPrioritizer.
func NewNodeLabelPriority(label string, presence bool) (PriorityMapFunction, PriorityReduceFunction) {
	labelPrioritizer := &NodeLabelPrioritizer{
		label:    label,
		presence: presence,
	}
	return labelPrioritizer.CalculateNodeLabelPriorityMap, nil
}

// CalculateNodeLabelPriorityMap checks whether a particular label exists on a node or not, regardless of its value.
// If presence is true, prioritizes nodes that have the specified label, regardless of value.
// If presence is false, prioritizes nodes that do not have the specified label.
func (n *NodeLabelPrioritizer) CalculateNodeLabelPriorityMap(pod *v1.Pod, meta interface{}, nodeInfo *schedulernodeinfo.NodeInfo) (schedulerapi.HostPriority, error) {
	node := nodeInfo.Node()
	if node == nil {
		return schedulerapi.HostPriority{}, fmt.Errorf("node not found")
	}

	exists := labels.Set(node.Labels).Has(n.label)
	score := 0
	if (exists && n.presence) || (!exists && !n.presence) {
		score = schedulerapi.MaxPriority
	}
	return schedulerapi.HostPriority{
		Host:  node.Name,
		Score: int64(score),
	}, nil
}
