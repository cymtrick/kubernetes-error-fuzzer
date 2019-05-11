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

package priorities

import (
	"reflect"
	"testing"

	"k8s.io/api/core/v1"
	schedulerapi "k8s.io/kubernetes/pkg/scheduler/api"
	schedulernodeinfo "k8s.io/kubernetes/pkg/scheduler/nodeinfo"
	st "k8s.io/kubernetes/pkg/scheduler/testing"
)

func Test_topologySpreadConstrantsMap_initialize(t *testing.T) {
	tests := []struct {
		name  string
		pod   *v1.Pod
		nodes []*v1.Node
		want  map[topologyPair][]string
	}{
		{
			name: "normal case",
			pod: st.MakePod().Name("p").Label("foo", "").
				SpreadConstraint(1, "zone", softSpread, st.MakeLabelSelector().Exists("foo").Obj()).
				SpreadConstraint(1, "node", softSpread, st.MakeLabelSelector().Exists("foo").Obj()).
				Obj(),
			nodes: []*v1.Node{
				st.MakeNode().Name("node-a").Label("zone", "zone1").Label("node", "node-a").Obj(),
				st.MakeNode().Name("node-b").Label("zone", "zone1").Label("node", "node-b").Obj(),
				st.MakeNode().Name("node-x").Label("zone", "zone2").Label("node", "node-x").Obj(),
			},
			want: map[topologyPair][]string{
				{key: "zone", value: "zone1"}:  {"node-a", "node-b"},
				{key: "zone", value: "zone2"}:  {"node-x"},
				{key: "node", value: "node-a"}: {"node-a"},
				{key: "node", value: "node-b"}: {"node-b"},
				{key: "node", value: "node-x"}: {"node-x"},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tMap := newTopologySpreadConstrantsMap(len(tt.nodes))
			tMap.initialize(tt.pod, tt.nodes)
			if !reflect.DeepEqual(tMap.topologyPairToNodeNames, tt.want) {
				t.Errorf("initilize().topologyPairToNodeNames = %#v, want %#v", tMap.topologyPairToNodeNames, tt.want)
			}
		})
	}
}

func TestCalculateEvenPodsSpreadPriority(t *testing.T) {
	tests := []struct {
		name         string
		pod          *v1.Pod
		existingPods []*v1.Pod
		nodes        []*v1.Node
		failedNodes  []*v1.Node // nodes + failedNodes = all nodes
		want         schedulerapi.HostPriorityList
	}{
		// Explanation on the Legend:
		// a) X/Y means there are X matching pods on node1 and Y on node2, both nodes are candidates
		//   (i.e. they have passed all predicates)
		// b) X/~Y~ means there are X matching pods on node1 and Y on node2, but node Y is NOT a candidate
		// c) X/?Y? means there are X matching pods on node1 and Y on node2, both nodes are candidates
		//    but node2 either i) doesn't have all required topologyKeys present, or ii) doesn't match
		//    incoming pod's nodeSelector/nodeAffinity
		{
			// if there is only one candidate node, it should be scored to 10
			name: "one constraint on node, no existing pods",
			pod: st.MakePod().Name("p").Label("foo", "").
				SpreadConstraint(1, "node", softSpread, st.MakeLabelSelector().Exists("foo").Obj()).
				Obj(),
			nodes: []*v1.Node{
				st.MakeNode().Name("node-a").Label("node", "node-a").Obj(),
				st.MakeNode().Name("node-b").Label("node", "node-b").Obj(),
			},
			want: []schedulerapi.HostPriority{
				{Host: "node-a", Score: 10},
				{Host: "node-b", Score: 10},
			},
		},
		{
			// if there is only one candidate node, it should be scored to 10
			name: "one constraint on node, only one node is candidate",
			pod: st.MakePod().Name("p").Label("foo", "").
				SpreadConstraint(1, "node", softSpread, st.MakeLabelSelector().Exists("foo").Obj()).
				Obj(),
			existingPods: []*v1.Pod{
				st.MakePod().Name("p-a1").Node("node-a").Label("foo", "").Obj(),
				st.MakePod().Name("p-a2").Node("node-a").Label("foo", "").Obj(),
				st.MakePod().Name("p-b1").Node("node-b").Label("foo", "").Obj(),
			},
			nodes: []*v1.Node{
				st.MakeNode().Name("node-a").Label("node", "node-a").Obj(),
			},
			failedNodes: []*v1.Node{
				st.MakeNode().Name("node-b").Label("node", "node-b").Obj(),
			},
			want: []schedulerapi.HostPriority{
				{Host: "node-a", Score: 10},
			},
		},
		{
			// matching pods spread as 2/1/0/3, total = 6
			// after reversing, it's 4/5/6/3
			// so scores = 40/6, 50/6, 60/6, 30/6
			name: "one constraint on node, all 4 nodes are candidates",
			pod: st.MakePod().Name("p").Label("foo", "").
				SpreadConstraint(1, "node", softSpread, st.MakeLabelSelector().Exists("foo").Obj()).
				Obj(),
			existingPods: []*v1.Pod{
				st.MakePod().Name("p-a1").Node("node-a").Label("foo", "").Obj(),
				st.MakePod().Name("p-a2").Node("node-a").Label("foo", "").Obj(),
				st.MakePod().Name("p-b1").Node("node-b").Label("foo", "").Obj(),
				st.MakePod().Name("p-d1").Node("node-d").Label("foo", "").Obj(),
				st.MakePod().Name("p-d2").Node("node-d").Label("foo", "").Obj(),
				st.MakePod().Name("p-d3").Node("node-d").Label("foo", "").Obj(),
			},
			nodes: []*v1.Node{
				st.MakeNode().Name("node-a").Label("node", "node-a").Obj(),
				st.MakeNode().Name("node-b").Label("node", "node-b").Obj(),
				st.MakeNode().Name("node-c").Label("node", "node-c").Obj(),
				st.MakeNode().Name("node-d").Label("node", "node-d").Obj(),
			},
			failedNodes: []*v1.Node{},
			want: []schedulerapi.HostPriority{
				{Host: "node-a", Score: 6},
				{Host: "node-b", Score: 8},
				{Host: "node-c", Score: 10},
				{Host: "node-d", Score: 5},
			},
		},
		{
			// matching pods spread as 4/2/1/~3~, total = 4+2+1 = 7 (as node4 is not a candidate)
			// after reversing, it's 3/5/6
			// so scores = 30/6, 50/6, 60/6
			name: "one constraint on node, 3 out of 4 nodes are candidates",
			pod: st.MakePod().Name("p").Label("foo", "").
				SpreadConstraint(1, "node", softSpread, st.MakeLabelSelector().Exists("foo").Obj()).
				Obj(),
			existingPods: []*v1.Pod{
				st.MakePod().Name("p-a1").Node("node-a").Label("foo", "").Obj(),
				st.MakePod().Name("p-a2").Node("node-a").Label("foo", "").Obj(),
				st.MakePod().Name("p-a3").Node("node-a").Label("foo", "").Obj(),
				st.MakePod().Name("p-a4").Node("node-a").Label("foo", "").Obj(),
				st.MakePod().Name("p-b1").Node("node-b").Label("foo", "").Obj(),
				st.MakePod().Name("p-b2").Node("node-b").Label("foo", "").Obj(),
				st.MakePod().Name("p-x1").Node("node-x").Label("foo", "").Obj(),
				st.MakePod().Name("p-y1").Node("node-y").Label("foo", "").Obj(),
				st.MakePod().Name("p-y2").Node("node-y").Label("foo", "").Obj(),
				st.MakePod().Name("p-y3").Node("node-y").Label("foo", "").Obj(),
			},
			nodes: []*v1.Node{
				st.MakeNode().Name("node-a").Label("node", "node-a").Obj(),
				st.MakeNode().Name("node-b").Label("node", "node-b").Obj(),
				st.MakeNode().Name("node-x").Label("node", "node-x").Obj(),
			},
			failedNodes: []*v1.Node{
				st.MakeNode().Name("node-y").Label("node", "node-y").Obj(),
			},
			want: []schedulerapi.HostPriority{
				{Host: "node-a", Score: 5},
				{Host: "node-b", Score: 8},
				{Host: "node-x", Score: 10},
			},
		},
		{
			// matching pods spread as 4/?2?/1/~3~, total = 4+?+1 = 5 (as node2 is problematic)
			// after reversing, it's 1/?/4
			// so scores = 10/4, 0, 40/4
			name: "one constraint on node, 3 out of 4 nodes are candidates",
			pod: st.MakePod().Name("p").Label("foo", "").
				SpreadConstraint(1, "node", softSpread, st.MakeLabelSelector().Exists("foo").Obj()).
				Obj(),
			existingPods: []*v1.Pod{
				st.MakePod().Name("p-a1").Node("node-a").Label("foo", "").Obj(),
				st.MakePod().Name("p-a2").Node("node-a").Label("foo", "").Obj(),
				st.MakePod().Name("p-a3").Node("node-a").Label("foo", "").Obj(),
				st.MakePod().Name("p-a4").Node("node-a").Label("foo", "").Obj(),
				st.MakePod().Name("p-b1").Node("node-b").Label("foo", "").Obj(),
				st.MakePod().Name("p-b2").Node("node-b").Label("foo", "").Obj(),
				st.MakePod().Name("p-x1").Node("node-x").Label("foo", "").Obj(),
				st.MakePod().Name("p-y1").Node("node-y").Label("foo", "").Obj(),
				st.MakePod().Name("p-y2").Node("node-y").Label("foo", "").Obj(),
				st.MakePod().Name("p-y3").Node("node-y").Label("foo", "").Obj(),
			},
			nodes: []*v1.Node{
				st.MakeNode().Name("node-a").Label("node", "node-a").Obj(),
				st.MakeNode().Name("node-b").Label("n", "node-b").Obj(), // label `n` doesn't match topologyKey
				st.MakeNode().Name("node-x").Label("node", "node-x").Obj(),
			},
			failedNodes: []*v1.Node{
				st.MakeNode().Name("node-y").Label("node", "node-y").Obj(),
			},
			want: []schedulerapi.HostPriority{
				{Host: "node-a", Score: 2},
				{Host: "node-b", Score: 0},
				{Host: "node-x", Score: 10},
			},
		},
		{
			// matching pods spread as 4/2/1/~3~, total = 6+6+4 = 16 (as topologyKey is zone instead of node)
			// after reversing, it's 10/10/12
			// so scores = 100/12, 100/12, 120/12
			name: "one constraint on zone, 3 out of 4 nodes are candidates",
			pod: st.MakePod().Name("p").Label("foo", "").
				SpreadConstraint(1, "zone", softSpread, st.MakeLabelSelector().Exists("foo").Obj()).
				Obj(),
			existingPods: []*v1.Pod{
				st.MakePod().Name("p-a1").Node("node-a").Label("foo", "").Obj(),
				st.MakePod().Name("p-a2").Node("node-a").Label("foo", "").Obj(),
				st.MakePod().Name("p-a3").Node("node-a").Label("foo", "").Obj(),
				st.MakePod().Name("p-a4").Node("node-a").Label("foo", "").Obj(),
				st.MakePod().Name("p-b1").Node("node-b").Label("foo", "").Obj(),
				st.MakePod().Name("p-b2").Node("node-b").Label("foo", "").Obj(),
				st.MakePod().Name("p-x1").Node("node-x").Label("foo", "").Obj(),
				st.MakePod().Name("p-y1").Node("node-y").Label("foo", "").Obj(),
				st.MakePod().Name("p-y2").Node("node-y").Label("foo", "").Obj(),
				st.MakePod().Name("p-y3").Node("node-y").Label("foo", "").Obj(),
			},
			nodes: []*v1.Node{
				st.MakeNode().Name("node-a").Label("zone", "zone1").Label("node", "node-a").Obj(),
				st.MakeNode().Name("node-b").Label("zone", "zone1").Label("node", "node-b").Obj(),
				st.MakeNode().Name("node-x").Label("zone", "zone2").Label("node", "node-x").Obj(),
			},
			failedNodes: []*v1.Node{
				st.MakeNode().Name("node-y").Label("zone", "zone2").Label("node", "node-y").Obj(),
			},
			want: []schedulerapi.HostPriority{
				{Host: "node-a", Score: 8},
				{Host: "node-b", Score: 8},
				{Host: "node-x", Score: 10},
			},
		},
		{
			// matching pods spread as 2/~1~/2/~4~, total = 2+3 + 2+6 = 13 (zone and node should be both sumed up)
			// after reversing, it's 8/5
			// so scores = 80/8, 50/8
			name: "two constraint on zone and node, 2 out of 4 nodes are candidates",
			pod: st.MakePod().Name("p").Label("foo", "").
				SpreadConstraint(1, "zone", softSpread, st.MakeLabelSelector().Exists("foo").Obj()).
				SpreadConstraint(1, "node", softSpread, st.MakeLabelSelector().Exists("foo").Obj()).
				Obj(),
			existingPods: []*v1.Pod{
				st.MakePod().Name("p-a1").Node("node-a").Label("foo", "").Obj(),
				st.MakePod().Name("p-a2").Node("node-a").Label("foo", "").Obj(),
				st.MakePod().Name("p-b1").Node("node-b").Label("foo", "").Obj(),
				st.MakePod().Name("p-x1").Node("node-x").Label("foo", "").Obj(),
				st.MakePod().Name("p-x2").Node("node-x").Label("foo", "").Obj(),
				st.MakePod().Name("p-y1").Node("node-y").Label("foo", "").Obj(),
				st.MakePod().Name("p-y2").Node("node-y").Label("foo", "").Obj(),
				st.MakePod().Name("p-y3").Node("node-y").Label("foo", "").Obj(),
				st.MakePod().Name("p-y4").Node("node-y").Label("foo", "").Obj(),
			},
			nodes: []*v1.Node{
				st.MakeNode().Name("node-a").Label("zone", "zone1").Label("node", "node-a").Obj(),
				st.MakeNode().Name("node-x").Label("zone", "zone2").Label("node", "node-x").Obj(),
			},
			failedNodes: []*v1.Node{
				st.MakeNode().Name("node-b").Label("zone", "zone1").Label("node", "node-b").Obj(),
				st.MakeNode().Name("node-y").Label("zone", "zone2").Label("node", "node-y").Obj(),
			},
			want: []schedulerapi.HostPriority{
				{Host: "node-a", Score: 10},
				{Host: "node-x", Score: 6},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			allNodes := append([]*v1.Node{}, tt.nodes...)
			allNodes = append(allNodes, tt.failedNodes...)
			nodeNameToInfo := schedulernodeinfo.CreateNodeNameToInfoMap(tt.existingPods, allNodes)

			got, _ := CalculateEvenPodsSpreadPriority(tt.pod, nodeNameToInfo, tt.nodes)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CalculateEvenPodsSpreadPriority() = %#v, want %#v", got, tt.want)
			}
		})
	}
}

var (
	hardSpread = v1.DoNotSchedule
	softSpread = v1.ScheduleAnyway
)
