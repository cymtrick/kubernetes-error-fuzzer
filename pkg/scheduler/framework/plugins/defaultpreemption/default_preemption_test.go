/*
Copyright 2020 The Kubernetes Authors.

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

package defaultpreemption

import (
	"context"
	"fmt"
	"reflect"
	"sort"
	"strings"
	"testing"
	"time"

	"github.com/google/go-cmp/cmp"
	v1 "k8s.io/api/core/v1"
	policy "k8s.io/api/policy/v1beta1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/util/sets"
	"k8s.io/client-go/informers"
	clientsetfake "k8s.io/client-go/kubernetes/fake"
	clienttesting "k8s.io/client-go/testing"
	"k8s.io/client-go/tools/events"
	extenderv1 "k8s.io/kube-scheduler/extender/v1"
	volumescheduling "k8s.io/kubernetes/pkg/controller/volume/scheduling"
	"k8s.io/kubernetes/pkg/scheduler/framework/plugins/defaultbinder"
	"k8s.io/kubernetes/pkg/scheduler/framework/plugins/interpodaffinity"
	"k8s.io/kubernetes/pkg/scheduler/framework/plugins/nodeaffinity"
	"k8s.io/kubernetes/pkg/scheduler/framework/plugins/nodelabel"
	"k8s.io/kubernetes/pkg/scheduler/framework/plugins/nodename"
	"k8s.io/kubernetes/pkg/scheduler/framework/plugins/noderesources"
	"k8s.io/kubernetes/pkg/scheduler/framework/plugins/nodeunschedulable"
	"k8s.io/kubernetes/pkg/scheduler/framework/plugins/podtopologyspread"
	"k8s.io/kubernetes/pkg/scheduler/framework/plugins/queuesort"
	"k8s.io/kubernetes/pkg/scheduler/framework/plugins/tainttoleration"
	"k8s.io/kubernetes/pkg/scheduler/framework/plugins/volumerestrictions"
	"k8s.io/kubernetes/pkg/scheduler/framework/plugins/volumezone"
	frameworkruntime "k8s.io/kubernetes/pkg/scheduler/framework/runtime"
	framework "k8s.io/kubernetes/pkg/scheduler/framework/v1alpha1"
	internalcache "k8s.io/kubernetes/pkg/scheduler/internal/cache"
	internalqueue "k8s.io/kubernetes/pkg/scheduler/internal/queue"
	st "k8s.io/kubernetes/pkg/scheduler/testing"
)

var (
	negPriority, lowPriority, midPriority, highPriority, veryHighPriority = int32(-100), int32(0), int32(100), int32(1000), int32(10000)

	smallRes = map[v1.ResourceName]string{
		v1.ResourceCPU:    "100m",
		v1.ResourceMemory: "100",
	}
	mediumRes = map[v1.ResourceName]string{
		v1.ResourceCPU:    "200m",
		v1.ResourceMemory: "200",
	}
	largeRes = map[v1.ResourceName]string{
		v1.ResourceCPU:    "300m",
		v1.ResourceMemory: "300",
	}
	veryLargeRes = map[v1.ResourceName]string{
		v1.ResourceCPU:    "500m",
		v1.ResourceMemory: "500",
	}

	epochTime  = metav1.NewTime(time.Unix(0, 0))
	epochTime1 = metav1.NewTime(time.Unix(0, 1))
	epochTime2 = metav1.NewTime(time.Unix(0, 2))
	epochTime3 = metav1.NewTime(time.Unix(0, 3))
	epochTime4 = metav1.NewTime(time.Unix(0, 4))
	epochTime5 = metav1.NewTime(time.Unix(0, 5))
	epochTime6 = metav1.NewTime(time.Unix(0, 6))
)

func mergeObjs(pod *v1.Pod, pods []*v1.Pod) []runtime.Object {
	var objs []runtime.Object
	if pod != nil {
		objs = append(objs, pod)
	}
	for i := range pods {
		objs = append(objs, pods[i])
	}
	return objs
}

func TestPostFilter(t *testing.T) {
	onePodRes := map[v1.ResourceName]string{v1.ResourcePods: "1"}
	tests := []struct {
		name                  string
		pod                   *v1.Pod
		pods                  []*v1.Pod
		nodes                 []*v1.Node
		filteredNodesStatuses framework.NodeToStatusMap
		extender              framework.Extender
		wantResult            *framework.PostFilterResult
		wantStatus            *framework.Status
	}{
		{
			name: "pod with higher priority can be made schedulable",
			pod:  st.MakePod().Name("p").UID("p").Priority(highPriority).Obj(),
			pods: []*v1.Pod{
				st.MakePod().Name("p1").UID("p1").Node("node1").Obj(),
			},
			nodes: []*v1.Node{
				st.MakeNode().Name("node1").Capacity(onePodRes).Obj(),
			},
			filteredNodesStatuses: framework.NodeToStatusMap{
				"node1": framework.NewStatus(framework.Unschedulable),
			},
			wantResult: &framework.PostFilterResult{NominatedNodeName: "node1"},
			wantStatus: framework.NewStatus(framework.Success),
		},
		{
			name: "pod with tied priority is still unschedulable",
			pod:  st.MakePod().Name("p").UID("p").Obj(),
			pods: []*v1.Pod{
				st.MakePod().Name("p1").UID("p1").Node("node1").Obj(),
			},
			nodes: []*v1.Node{
				st.MakeNode().Name("node1").Capacity(onePodRes).Obj(),
			},
			filteredNodesStatuses: framework.NodeToStatusMap{
				"node1": framework.NewStatus(framework.Unschedulable),
			},
			wantResult: nil,
			wantStatus: framework.NewStatus(framework.Unschedulable),
		},
		{
			name: "preemption should respect filteredNodesStatuses",
			pod:  st.MakePod().Name("p").UID("p").Priority(highPriority).Obj(),
			pods: []*v1.Pod{
				st.MakePod().Name("p1").UID("p1").Node("node1").Obj(),
			},
			nodes: []*v1.Node{
				st.MakeNode().Name("node1").Capacity(onePodRes).Obj(),
			},
			filteredNodesStatuses: framework.NodeToStatusMap{
				"node1": framework.NewStatus(framework.UnschedulableAndUnresolvable),
			},
			wantResult: nil,
			wantStatus: framework.NewStatus(framework.Unschedulable),
		},
		{
			name: "pod can be made schedulable on one node",
			pod:  st.MakePod().Name("p").UID("p").Priority(midPriority).Obj(),
			pods: []*v1.Pod{
				st.MakePod().Name("p1").UID("p1").Priority(highPriority).Node("node1").Obj(),
				st.MakePod().Name("p2").UID("p2").Priority(lowPriority).Node("node2").Obj(),
			},
			nodes: []*v1.Node{
				st.MakeNode().Name("node1").Capacity(onePodRes).Obj(),
				st.MakeNode().Name("node2").Capacity(onePodRes).Obj(),
			},
			filteredNodesStatuses: framework.NodeToStatusMap{
				"node1": framework.NewStatus(framework.Unschedulable),
				"node2": framework.NewStatus(framework.Unschedulable),
			},
			wantResult: &framework.PostFilterResult{NominatedNodeName: "node2"},
			wantStatus: framework.NewStatus(framework.Success),
		},
		{
			name: "preemption result filtered out by extenders",
			pod:  st.MakePod().Name("p").UID("p").Priority(highPriority).Obj(),
			pods: []*v1.Pod{
				st.MakePod().Name("p1").UID("p1").Node("machine1").Obj(),
				st.MakePod().Name("p2").UID("p2").Node("machine2").Obj(),
			},
			nodes: []*v1.Node{
				st.MakeNode().Name("machine1").Capacity(onePodRes).Obj(),
				st.MakeNode().Name("machine2").Capacity(onePodRes).Obj(),
			},
			filteredNodesStatuses: framework.NodeToStatusMap{
				"machine1": framework.NewStatus(framework.Unschedulable),
				"machine2": framework.NewStatus(framework.Unschedulable),
			},
			extender: &st.FakeExtender{Predicates: []st.FitPredicate{st.Machine1PredicateExtender}},
			wantResult: &framework.PostFilterResult{
				NominatedNodeName: "machine1",
			},
			wantStatus: framework.NewStatus(framework.Success),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			apiObjs := mergeObjs(tt.pod, tt.pods /*, tt.nodes */)
			cs := clientsetfake.NewSimpleClientset(apiObjs...)
			informerFactory := informers.NewSharedInformerFactory(cs, 0)
			// Register NodeResourceFit as the Filter & PreFilter plugin.
			registeredPlugins := []st.RegisterPluginFunc{
				st.RegisterQueueSortPlugin(queuesort.Name, queuesort.New),
				st.RegisterPluginAsExtensions(noderesources.FitName, noderesources.NewFit, "Filter", "PreFilter"),
				st.RegisterBindPlugin(defaultbinder.Name, defaultbinder.New),
			}
			var extenders []framework.Extender
			if tt.extender != nil {
				extenders = append(extenders, tt.extender)
			}
			f, err := st.NewFramework(registeredPlugins,
				frameworkruntime.WithClientSet(cs),
				frameworkruntime.WithEventRecorder(&events.FakeRecorder{}),
				frameworkruntime.WithInformerFactory(informerFactory),
				frameworkruntime.WithPodNominator(internalqueue.NewPodNominator()),
				frameworkruntime.WithExtenders(extenders),
				frameworkruntime.WithSnapshotSharedLister(internalcache.NewSnapshot(tt.pods, tt.nodes)),
			)
			if err != nil {
				t.Fatal(err)
			}

			state := framework.NewCycleState()
			// Ensure <state> is populated.
			if status := f.RunPreFilterPlugins(context.Background(), state, tt.pod); !status.IsSuccess() {
				t.Errorf("Unexpected PreFilter Status: %v", status)
			}

			p := DefaultPreemption{
				fh:        f,
				pdbLister: getPDBLister(informerFactory),
			}
			gotResult, gotStatus := p.PostFilter(context.TODO(), state, tt.pod, tt.filteredNodesStatuses)
			if !reflect.DeepEqual(gotStatus, tt.wantStatus) {
				t.Errorf("Status does not match: %v, want: %v", gotStatus, tt.wantStatus)
			}
			if diff := cmp.Diff(gotResult, tt.wantResult); diff != "" {
				t.Errorf("Unexpected postFilterResult (-want, +got): %s", diff)
			}
		})
	}
}

// TestSelectNodesForPreemption tests dryRunPreemption. This test assumes
// that podsFitsOnNode works correctly and is tested separately.
func TestDryRunPreemption(t *testing.T) {
	tests := []struct {
		name                    string
		nodeNames               []string
		pod                     *v1.Pod
		pods                    []*v1.Pod
		registerPlugins         []st.RegisterPluginFunc
		pdbs                    []*policy.PodDisruptionBudget
		fakeFilterRC            framework.Code // return code for fake filter plugin
		expected                []Candidate
		expectedNumFilterCalled int32
	}{
		{
			name: "a pod that does not fit on any machine",
			registerPlugins: []st.RegisterPluginFunc{
				st.RegisterFilterPlugin("FalseFilter", st.NewFalseFilterPlugin),
			},
			nodeNames: []string{"node1", "node2"},
			pod:       st.MakePod().Name("p").UID("p").Priority(highPriority).Obj(),
			pods: []*v1.Pod{
				st.MakePod().Name("p1").UID("p1").Node("node1").Priority(midPriority).Obj(),
				st.MakePod().Name("p2").UID("p2").Node("node2").Priority(midPriority).Obj(),
			},
			expected:                nil,
			expectedNumFilterCalled: 2,
		},
		{
			name: "a pod that fits with no preemption",
			registerPlugins: []st.RegisterPluginFunc{
				st.RegisterFilterPlugin("TrueFilter", st.NewTrueFilterPlugin),
			},
			nodeNames: []string{"node1", "node2"},
			pod:       st.MakePod().Name("p").UID("p").Priority(highPriority).Obj(),
			pods: []*v1.Pod{
				st.MakePod().Name("p1").UID("p1").Node("node1").Priority(midPriority).Obj(),
				st.MakePod().Name("p2").UID("p2").Node("node2").Priority(midPriority).Obj(),
			},
			expected: []Candidate{
				&candidate{victims: &extenderv1.Victims{}, name: "node1"},
				&candidate{victims: &extenderv1.Victims{}, name: "node2"},
			},
			expectedNumFilterCalled: 4,
		},
		{
			name: "a pod that fits on one machine with no preemption",
			registerPlugins: []st.RegisterPluginFunc{
				st.RegisterFilterPlugin("MatchFilter", st.NewMatchFilterPlugin),
			},
			nodeNames: []string{"node1", "node2"},
			// Name the pod as "node1" to fit "MatchFilter" plugin.
			pod: st.MakePod().Name("node1").UID("node1").Priority(highPriority).Obj(),
			pods: []*v1.Pod{
				st.MakePod().Name("p1").UID("p1").Node("node1").Priority(midPriority).Obj(),
				st.MakePod().Name("p2").UID("p2").Node("node2").Priority(midPriority).Obj(),
			},
			expected: []Candidate{
				&candidate{victims: &extenderv1.Victims{}, name: "node1"},
			},
			expectedNumFilterCalled: 3,
		},
		{
			name: "a pod that fits on both machines when lower priority pods are preempted",
			registerPlugins: []st.RegisterPluginFunc{
				st.RegisterPluginAsExtensions(noderesources.FitName, noderesources.NewFit, "Filter", "PreFilter"),
			},
			nodeNames: []string{"node1", "node2"},
			pod:       st.MakePod().Name("p").UID("p").Priority(highPriority).Req(largeRes).Obj(),
			pods: []*v1.Pod{
				st.MakePod().Name("p1").UID("p1").Node("node1").Priority(midPriority).Req(largeRes).Obj(),
				st.MakePod().Name("p2").UID("p2").Node("node2").Priority(midPriority).Req(largeRes).Obj(),
			},
			expected: []Candidate{
				&candidate{
					victims: &extenderv1.Victims{
						Pods: []*v1.Pod{st.MakePod().Name("p1").UID("p1").Node("node1").Priority(midPriority).Req(largeRes).Obj()},
					},
					name: "node1",
				},
				&candidate{
					victims: &extenderv1.Victims{
						Pods: []*v1.Pod{st.MakePod().Name("p2").UID("p2").Node("node2").Priority(midPriority).Req(largeRes).Obj()},
					},
					name: "node2",
				},
			},
			expectedNumFilterCalled: 4,
		},
		{
			name: "a pod that would fit on the machines, but other pods running are higher priority, no preemption would happen",
			registerPlugins: []st.RegisterPluginFunc{
				st.RegisterPluginAsExtensions(noderesources.FitName, noderesources.NewFit, "Filter", "PreFilter"),
			},
			nodeNames: []string{"node1", "node2"},
			pod:       st.MakePod().Name("p").UID("p").Priority(lowPriority).Req(largeRes).Obj(),
			pods: []*v1.Pod{
				st.MakePod().Name("p1").UID("p1").Node("node1").Priority(midPriority).Req(largeRes).Obj(),
				st.MakePod().Name("p2").UID("p2").Node("node2").Priority(midPriority).Req(largeRes).Obj(),
			},
			expected:                nil,
			expectedNumFilterCalled: 0,
		},
		{
			name: "medium priority pod is preempted, but lower priority one stays as it is small",
			registerPlugins: []st.RegisterPluginFunc{
				st.RegisterPluginAsExtensions(noderesources.FitName, noderesources.NewFit, "Filter", "PreFilter"),
			},
			nodeNames: []string{"node1", "node2"},
			pod:       st.MakePod().Name("p").UID("p").Priority(highPriority).Req(largeRes).Obj(),
			pods: []*v1.Pod{
				st.MakePod().Name("p1.1").UID("p1.1").Node("node1").Priority(lowPriority).Req(smallRes).Obj(),
				st.MakePod().Name("p1.2").UID("p1.2").Node("node1").Priority(midPriority).Req(largeRes).Obj(),
				st.MakePod().Name("p2").UID("p2").Node("node2").Priority(midPriority).Req(largeRes).Obj(),
			},
			expected: []Candidate{
				&candidate{
					victims: &extenderv1.Victims{
						Pods: []*v1.Pod{st.MakePod().Name("p1.2").UID("p1.2").Node("node1").Priority(midPriority).Req(largeRes).Obj()},
					},
					name: "node1",
				},
				&candidate{
					victims: &extenderv1.Victims{
						Pods: []*v1.Pod{st.MakePod().Name("p2").UID("p2").Node("node2").Priority(midPriority).Req(largeRes).Obj()},
					},
					name: "node2",
				},
			},
			expectedNumFilterCalled: 5,
		},
		{
			name: "mixed priority pods are preempted",
			registerPlugins: []st.RegisterPluginFunc{
				st.RegisterPluginAsExtensions(noderesources.FitName, noderesources.NewFit, "Filter", "PreFilter"),
			},
			nodeNames: []string{"node1", "node2"},
			pod:       st.MakePod().Name("p").UID("p").Priority(highPriority).Req(largeRes).Obj(),
			pods: []*v1.Pod{
				st.MakePod().Name("p1.1").UID("p1.1").Node("node1").Priority(midPriority).Req(smallRes).Obj(),
				st.MakePod().Name("p1.2").UID("p1.2").Node("node1").Priority(lowPriority).Req(smallRes).Obj(),
				st.MakePod().Name("p1.3").UID("p1.3").Node("node1").Priority(midPriority).Req(mediumRes).Obj(),
				st.MakePod().Name("p1.4").UID("p1.4").Node("node1").Priority(highPriority).Req(smallRes).Obj(),
				st.MakePod().Name("p2").UID("p2").Node("node2").Priority(highPriority).Req(largeRes).Obj(),
			},
			expected: []Candidate{
				&candidate{
					victims: &extenderv1.Victims{
						Pods: []*v1.Pod{
							st.MakePod().Name("p1.2").UID("p1.2").Node("node1").Priority(lowPriority).Req(smallRes).Obj(),
							st.MakePod().Name("p1.3").UID("p1.3").Node("node1").Priority(midPriority).Req(mediumRes).Obj(),
						},
					},
					name: "node1",
				},
			},
			expectedNumFilterCalled: 4,
		},
		{
			name: "mixed priority pods are preempted, pick later StartTime one when priorities are equal",
			registerPlugins: []st.RegisterPluginFunc{
				st.RegisterPluginAsExtensions(noderesources.FitName, noderesources.NewFit, "Filter", "PreFilter"),
			},
			nodeNames: []string{"node1", "node2"},
			pod:       st.MakePod().Name("p").UID("p").Priority(highPriority).Req(largeRes).Obj(),
			pods: []*v1.Pod{
				st.MakePod().Name("p1.1").UID("p1.1").Node("node1").Priority(lowPriority).Req(smallRes).StartTime(epochTime5).Obj(),
				st.MakePod().Name("p1.2").UID("p1.2").Node("node1").Priority(lowPriority).Req(smallRes).StartTime(epochTime4).Obj(),
				st.MakePod().Name("p1.3").UID("p1.3").Node("node1").Priority(midPriority).Req(mediumRes).StartTime(epochTime3).Obj(),
				st.MakePod().Name("p1.4").UID("p1.4").Node("node1").Priority(highPriority).Req(smallRes).StartTime(epochTime2).Obj(),
				st.MakePod().Name("p2").UID("p2").Node("node2").Priority(highPriority).Req(largeRes).StartTime(epochTime1).Obj(),
			},
			expected: []Candidate{
				&candidate{
					victims: &extenderv1.Victims{
						Pods: []*v1.Pod{
							st.MakePod().Name("p1.1").UID("p1.1").Node("node1").Priority(lowPriority).Req(smallRes).StartTime(epochTime5).Obj(),
							st.MakePod().Name("p1.3").UID("p1.3").Node("node1").Priority(midPriority).Req(mediumRes).StartTime(epochTime3).Obj(),
						},
					},
					name: "node1",
				},
			},
			expectedNumFilterCalled: 4, // no preemption would happen on node2 and no filter call is counted.
		},
		{
			name: "pod with anti-affinity is preempted",
			registerPlugins: []st.RegisterPluginFunc{
				st.RegisterPluginAsExtensions(noderesources.FitName, noderesources.NewFit, "Filter", "PreFilter"),
				st.RegisterPluginAsExtensions(interpodaffinity.Name, interpodaffinity.New, "Filter", "PreFilter"),
			},
			nodeNames: []string{"node1", "node2"},
			pod:       st.MakePod().Name("p").UID("p").Label("foo", "").Priority(highPriority).Req(smallRes).Obj(),
			pods: []*v1.Pod{
				st.MakePod().Name("p1.1").UID("p1.1").Node("node1").Label("foo", "").Priority(lowPriority).Req(smallRes).
					PodAntiAffinityExists("foo", "hostname", st.PodAntiAffinityWithRequiredReq).Obj(),
				st.MakePod().Name("p1.2").UID("p1.2").Node("node1").Priority(midPriority).Req(smallRes).Obj(),
				st.MakePod().Name("p1.3").UID("p1.3").Node("node1").Priority(highPriority).Req(smallRes).Obj(),
				st.MakePod().Name("p2").UID("p2").Node("node2").Priority(highPriority).Req(smallRes).Obj(),
			},
			expected: []Candidate{
				&candidate{
					victims: &extenderv1.Victims{
						Pods: []*v1.Pod{
							st.MakePod().Name("p1.1").UID("p1.1").Node("node1").Label("foo", "").Priority(lowPriority).Req(smallRes).
								PodAntiAffinityExists("foo", "hostname", st.PodAntiAffinityWithRequiredReq).Obj(),
						},
					},
					name: "node1",
				},
			},
			expectedNumFilterCalled: 3, // no preemption would happen on node2 and no filter call is counted.
		},
		{
			name: "preemption to resolve pod topology spread filter failure",
			registerPlugins: []st.RegisterPluginFunc{
				st.RegisterPluginAsExtensions(podtopologyspread.Name, podtopologyspread.New, "PreFilter", "Filter"),
			},
			nodeNames: []string{"node-a/zone1", "node-b/zone1", "node-x/zone2"},
			pod: st.MakePod().Name("p").UID("p").Label("foo", "").Priority(highPriority).
				SpreadConstraint(1, "zone", v1.DoNotSchedule, st.MakeLabelSelector().Exists("foo").Obj()).
				SpreadConstraint(1, "hostname", v1.DoNotSchedule, st.MakeLabelSelector().Exists("foo").Obj()).
				Obj(),
			pods: []*v1.Pod{
				st.MakePod().Name("pod-a1").UID("pod-a1").Node("node-a").Label("foo", "").Priority(midPriority).Obj(),
				st.MakePod().Name("pod-a2").UID("pod-a2").Node("node-a").Label("foo", "").Priority(lowPriority).Obj(),
				st.MakePod().Name("pod-b1").UID("pod-b1").Node("node-b").Label("foo", "").Priority(lowPriority).Obj(),
				st.MakePod().Name("pod-x1").UID("pod-x1").Node("node-x").Label("foo", "").Priority(highPriority).Obj(),
				st.MakePod().Name("pod-x2").UID("pod-x2").Node("node-x").Label("foo", "").Priority(highPriority).Obj(),
			},
			expected: []Candidate{
				&candidate{
					victims: &extenderv1.Victims{
						Pods: []*v1.Pod{st.MakePod().Name("pod-a2").UID("pod-a2").Node("node-a").Label("foo", "").Priority(lowPriority).Obj()},
					},
					name: "node-a",
				},
				&candidate{
					victims: &extenderv1.Victims{
						Pods: []*v1.Pod{st.MakePod().Name("pod-b1").UID("pod-b1").Node("node-b").Label("foo", "").Priority(lowPriority).Obj()},
					},
					name: "node-b",
				},
			},
			expectedNumFilterCalled: 5, // node-a (3), node-b (2), node-x (0)
		},
		{
			name: "get Unschedulable in the preemption phase when the filter plugins filtering the nodes",
			registerPlugins: []st.RegisterPluginFunc{
				st.RegisterPluginAsExtensions(noderesources.FitName, noderesources.NewFit, "Filter", "PreFilter"),
			},
			nodeNames: []string{"node1", "node2"},
			pod:       st.MakePod().Name("p").UID("p").Priority(highPriority).Req(largeRes).Obj(),
			pods: []*v1.Pod{
				st.MakePod().Name("p1").UID("p1").Node("node1").Priority(midPriority).Req(largeRes).Obj(),
				st.MakePod().Name("p2").UID("p2").Node("node2").Priority(midPriority).Req(largeRes).Obj(),
			},
			fakeFilterRC:            framework.Unschedulable,
			expected:                nil,
			expectedNumFilterCalled: 2,
		},
		{
			name: "preemption with violation of same pdb",
			registerPlugins: []st.RegisterPluginFunc{
				st.RegisterPluginAsExtensions(noderesources.FitName, noderesources.NewFit, "Filter", "PreFilter"),
			},
			nodeNames: []string{"node1"},
			pod:       st.MakePod().Name("p").UID("p").Priority(highPriority).Req(veryLargeRes).Obj(),
			pods: []*v1.Pod{
				st.MakePod().Name("p1.1").UID("p1.1").Node("node1").Label("app", "foo").Priority(midPriority).Req(mediumRes).Obj(),
				st.MakePod().Name("p1.2").UID("p1.2").Node("node1").Label("app", "foo").Priority(midPriority).Req(mediumRes).Obj(),
			},
			pdbs: []*policy.PodDisruptionBudget{
				{
					Spec:   policy.PodDisruptionBudgetSpec{Selector: &metav1.LabelSelector{MatchLabels: map[string]string{"app": "foo"}}},
					Status: policy.PodDisruptionBudgetStatus{DisruptionsAllowed: 1},
				},
			},
			expected: []Candidate{
				&candidate{
					victims: &extenderv1.Victims{
						Pods: []*v1.Pod{
							st.MakePod().Name("p1.1").UID("p1.1").Node("node1").Label("app", "foo").Priority(midPriority).Req(mediumRes).Obj(),
							st.MakePod().Name("p1.2").UID("p1.2").Node("node1").Label("app", "foo").Priority(midPriority).Req(mediumRes).Obj(),
						},
						NumPDBViolations: 1,
					},
					name: "node1",
				},
			},
			expectedNumFilterCalled: 3,
		},
		{
			name: "preemption with violation of the pdb with pod whose eviction was processed, the victim doesn't belong to DisruptedPods",
			registerPlugins: []st.RegisterPluginFunc{
				st.RegisterPluginAsExtensions(noderesources.FitName, noderesources.NewFit, "Filter", "PreFilter"),
			},
			nodeNames: []string{"node1"},
			pod:       st.MakePod().Name("p").UID("p").Priority(highPriority).Req(veryLargeRes).Obj(),
			pods: []*v1.Pod{
				st.MakePod().Name("p1.1").UID("p1.1").Node("node1").Label("app", "foo").Priority(midPriority).Req(mediumRes).Obj(),
				st.MakePod().Name("p1.2").UID("p1.2").Node("node1").Label("app", "foo").Priority(midPriority).Req(mediumRes).Obj(),
			},
			pdbs: []*policy.PodDisruptionBudget{
				{
					Spec:   policy.PodDisruptionBudgetSpec{Selector: &metav1.LabelSelector{MatchLabels: map[string]string{"app": "foo"}}},
					Status: policy.PodDisruptionBudgetStatus{DisruptionsAllowed: 1, DisruptedPods: map[string]metav1.Time{"p2": {Time: time.Now()}}},
				},
			},
			expected: []Candidate{
				&candidate{
					victims: &extenderv1.Victims{
						Pods: []*v1.Pod{
							st.MakePod().Name("p1.1").UID("p1.1").Node("node1").Label("app", "foo").Priority(midPriority).Req(mediumRes).Obj(),
							st.MakePod().Name("p1.2").UID("p1.2").Node("node1").Label("app", "foo").Priority(midPriority).Req(mediumRes).Obj(),
						},
						NumPDBViolations: 1,
					},
					name: "node1",
				},
			},
			expectedNumFilterCalled: 3,
		},
		{
			name: "preemption with violation of the pdb with pod whose eviction was processed, the victim belongs to DisruptedPods",
			registerPlugins: []st.RegisterPluginFunc{
				st.RegisterPluginAsExtensions(noderesources.FitName, noderesources.NewFit, "Filter", "PreFilter"),
			},
			nodeNames: []string{"node1"},
			pod:       st.MakePod().Name("p").UID("p").Priority(highPriority).Req(veryLargeRes).Obj(),
			pods: []*v1.Pod{
				st.MakePod().Name("p1.1").UID("p1.1").Node("node1").Label("app", "foo").Priority(midPriority).Req(mediumRes).Obj(),
				st.MakePod().Name("p1.2").UID("p1.2").Node("node1").Label("app", "foo").Priority(midPriority).Req(mediumRes).Obj(),
			},
			pdbs: []*policy.PodDisruptionBudget{
				{
					Spec:   policy.PodDisruptionBudgetSpec{Selector: &metav1.LabelSelector{MatchLabels: map[string]string{"app": "foo"}}},
					Status: policy.PodDisruptionBudgetStatus{DisruptionsAllowed: 1, DisruptedPods: map[string]metav1.Time{"p1.2": {Time: time.Now()}}},
				},
			},
			expected: []Candidate{
				&candidate{
					victims: &extenderv1.Victims{
						Pods: []*v1.Pod{
							st.MakePod().Name("p1.1").UID("p1.1").Node("node1").Label("app", "foo").Priority(midPriority).Req(mediumRes).Obj(),
							st.MakePod().Name("p1.2").UID("p1.2").Node("node1").Label("app", "foo").Priority(midPriority).Req(mediumRes).Obj(),
						},
						NumPDBViolations: 0,
					},
					name: "node1",
				},
			},
			expectedNumFilterCalled: 3,
		},
		{
			name: "preemption with violation of the pdb with pod whose eviction was processed, the victim which belongs to DisruptedPods is treated as 'nonViolating'",
			registerPlugins: []st.RegisterPluginFunc{
				st.RegisterPluginAsExtensions(noderesources.FitName, noderesources.NewFit, "Filter", "PreFilter"),
			},
			nodeNames: []string{"node1"},
			pod:       st.MakePod().Name("p").UID("p").Priority(highPriority).Req(veryLargeRes).Obj(),
			pods: []*v1.Pod{
				st.MakePod().Name("p1.1").UID("p1.1").Node("node1").Label("app", "foo").Priority(midPriority).Req(mediumRes).Obj(),
				st.MakePod().Name("p1.2").UID("p1.2").Node("node1").Label("app", "foo").Priority(midPriority).Req(mediumRes).Obj(),
				st.MakePod().Name("p1.3").UID("p1.3").Node("node1").Label("app", "foo").Priority(midPriority).Req(mediumRes).Obj(),
			},
			pdbs: []*policy.PodDisruptionBudget{
				{
					Spec:   policy.PodDisruptionBudgetSpec{Selector: &metav1.LabelSelector{MatchLabels: map[string]string{"app": "foo"}}},
					Status: policy.PodDisruptionBudgetStatus{DisruptionsAllowed: 1, DisruptedPods: map[string]metav1.Time{"p1.3": {Time: time.Now()}}},
				},
			},
			expected: []Candidate{
				&candidate{
					victims: &extenderv1.Victims{
						Pods: []*v1.Pod{
							st.MakePod().Name("p1.1").UID("p1.1").Node("node1").Label("app", "foo").Priority(midPriority).Req(mediumRes).Obj(),
							st.MakePod().Name("p1.2").UID("p1.2").Node("node1").Label("app", "foo").Priority(midPriority).Req(mediumRes).Obj(),
							st.MakePod().Name("p1.3").UID("p1.3").Node("node1").Label("app", "foo").Priority(midPriority).Req(mediumRes).Obj(),
						},
						NumPDBViolations: 1,
					},
					name: "node1",
				},
			},
			expectedNumFilterCalled: 4,
		},
	}

	labelKeys := []string{"hostname", "zone", "region"}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			nodes := make([]*v1.Node, len(tt.nodeNames))
			fakeFilterRCMap := make(map[string]framework.Code, len(tt.nodeNames))
			for i, nodeName := range tt.nodeNames {
				nodeWrapper := st.MakeNode().Capacity(veryLargeRes)
				// Split node name by '/' to form labels in a format of
				// {"hostname": tpKeys[0], "zone": tpKeys[1], "region": tpKeys[2]}
				tpKeys := strings.Split(nodeName, "/")
				nodeWrapper.Name(tpKeys[0])
				for i, labelVal := range strings.Split(nodeName, "/") {
					nodeWrapper.Label(labelKeys[i], labelVal)
				}
				nodes[i] = nodeWrapper.Obj()
				fakeFilterRCMap[nodeName] = tt.fakeFilterRC
			}
			snapshot := internalcache.NewSnapshot(tt.pods, nodes)

			// For each test, register a FakeFilterPlugin along with essential plugins and tt.registerPlugins.
			fakePlugin := st.FakeFilterPlugin{
				FailedNodeReturnCodeMap: fakeFilterRCMap,
			}
			registeredPlugins := append([]st.RegisterPluginFunc{
				st.RegisterFilterPlugin(
					"FakeFilter",
					func(_ runtime.Object, fh framework.FrameworkHandle) (framework.Plugin, error) {
						return &fakePlugin, nil
					},
				)},
				st.RegisterQueueSortPlugin(queuesort.Name, queuesort.New),
				st.RegisterBindPlugin(defaultbinder.Name, defaultbinder.New),
			)
			registeredPlugins = append(registeredPlugins, tt.registerPlugins...)
			fwk, err := st.NewFramework(
				registeredPlugins,
				frameworkruntime.WithPodNominator(internalqueue.NewPodNominator()),
				frameworkruntime.WithSnapshotSharedLister(snapshot),
			)
			if err != nil {
				t.Fatal(err)
			}

			state := framework.NewCycleState()
			// Some tests rely on PreFilter plugin to compute its CycleState.
			if status := fwk.RunPreFilterPlugins(context.Background(), state, tt.pod); !status.IsSuccess() {
				t.Errorf("Unexpected PreFilter Status: %v", status)
			}

			nodeInfos, err := snapshot.NodeInfos().List()
			if err != nil {
				t.Fatal(err)
			}
			got := dryRunPreemption(context.Background(), fwk.PreemptHandle(), state, tt.pod, nodeInfos, tt.pdbs)
			if err != nil {
				t.Fatal(err)
			}
			// Sort the values (inner victims) and the candidate itself (by its NominatedNodeName).
			for i := range got {
				victims := got[i].Victims().Pods
				sort.Slice(victims, func(i, j int) bool {
					return victims[i].Name < victims[j].Name
				})
			}
			sort.Slice(got, func(i, j int) bool {
				return got[i].Name() < got[j].Name()
			})

			if tt.expectedNumFilterCalled != fakePlugin.NumFilterCalled {
				t.Errorf("expected fakePlugin.numFilterCalled is %d, but got %d", tt.expectedNumFilterCalled, fakePlugin.NumFilterCalled)
			}
			if diff := cmp.Diff(tt.expected, got, cmp.AllowUnexported(candidate{})); diff != "" {
				t.Errorf("Unexpected candidates (-want, +got): %s", diff)
			}
		})
	}
}

func TestSelectBestCandidate(t *testing.T) {
	tests := []struct {
		name           string
		registerPlugin st.RegisterPluginFunc
		nodeNames      []string
		pod            *v1.Pod
		pods           []*v1.Pod
		expected       []string // any of the items is valid
	}{
		{
			name:           "No node needs preemption",
			registerPlugin: st.RegisterPluginAsExtensions(noderesources.FitName, noderesources.NewFit, "Filter", "PreFilter"),
			nodeNames:      []string{"node1"},
			pod:            st.MakePod().Name("p").UID("p").Priority(highPriority).Req(largeRes).Obj(),
			pods: []*v1.Pod{
				st.MakePod().Name("p1").UID("p1").Node("node1").Priority(midPriority).Req(smallRes).StartTime(epochTime).Obj(),
			},
			expected: []string{"node1"},
		},
		{
			name:           "a pod that fits on both nodes when lower priority pods are preempted",
			registerPlugin: st.RegisterPluginAsExtensions(noderesources.FitName, noderesources.NewFit, "Filter", "PreFilter"),
			nodeNames:      []string{"node1", "node2"},
			pod:            st.MakePod().Name("p").UID("p").Priority(highPriority).Req(largeRes).Obj(),
			pods: []*v1.Pod{
				st.MakePod().Name("p1").UID("p1").Node("node1").Priority(midPriority).Req(largeRes).StartTime(epochTime).Obj(),
				st.MakePod().Name("p2").UID("p2").Node("node2").Priority(midPriority).Req(largeRes).StartTime(epochTime).Obj(),
			},
			expected: []string{"node1", "node2"},
		},
		{
			name:           "machine with min highest priority pod is picked",
			registerPlugin: st.RegisterPluginAsExtensions(noderesources.FitName, noderesources.NewFit, "Filter", "PreFilter"),
			nodeNames:      []string{"node1", "node2", "node3"},
			pod:            st.MakePod().Name("p").UID("p").Priority(highPriority).Req(veryLargeRes).Obj(),
			pods: []*v1.Pod{
				st.MakePod().Name("p1.1").UID("p1.1").Node("node1").Priority(midPriority).Req(mediumRes).StartTime(epochTime).Obj(),
				st.MakePod().Name("p1.2").UID("p1.2").Node("node1").Priority(midPriority).Req(largeRes).StartTime(epochTime).Obj(),
				st.MakePod().Name("p2.1").UID("p2.1").Node("node2").Priority(midPriority).Req(mediumRes).StartTime(epochTime).Obj(),
				st.MakePod().Name("p2.2").UID("p2.2").Node("node2").Priority(lowPriority).Req(mediumRes).StartTime(epochTime).Obj(),
				st.MakePod().Name("p3.1").UID("p3.1").Node("node3").Priority(lowPriority).Req(mediumRes).StartTime(epochTime).Obj(),
				st.MakePod().Name("p3.2").UID("p3.2").Node("node3").Priority(lowPriority).Req(mediumRes).StartTime(epochTime).Obj(),
			},
			expected: []string{"node3"},
		},
		{
			name:           "when highest priorities are the same, minimum sum of priorities is picked",
			registerPlugin: st.RegisterPluginAsExtensions(noderesources.FitName, noderesources.NewFit, "Filter", "PreFilter"),
			nodeNames:      []string{"node1", "node2", "node3"},
			pod:            st.MakePod().Name("p").UID("p").Priority(highPriority).Req(veryLargeRes).Obj(),
			pods: []*v1.Pod{
				st.MakePod().Name("p1.1").UID("p1.1").Node("node1").Priority(midPriority).Req(mediumRes).StartTime(epochTime).Obj(),
				st.MakePod().Name("p1.2").UID("p1.2").Node("node1").Priority(midPriority).Req(largeRes).StartTime(epochTime).Obj(),
				st.MakePod().Name("p2.1").UID("p2.1").Node("node2").Priority(midPriority).Req(largeRes).StartTime(epochTime).Obj(),
				st.MakePod().Name("p2.2").UID("p2.2").Node("node2").Priority(lowPriority).Req(mediumRes).StartTime(epochTime).Obj(),
				st.MakePod().Name("p3.1").UID("p3.1").Node("node3").Priority(midPriority).Req(mediumRes).StartTime(epochTime).Obj(),
				st.MakePod().Name("p3.2").UID("p3.2").Node("node3").Priority(midPriority).Req(mediumRes).StartTime(epochTime).Obj(),
			},
			expected: []string{"node2"},
		},
		{
			name:           "when highest priority and sum are the same, minimum number of pods is picked",
			registerPlugin: st.RegisterPluginAsExtensions(noderesources.FitName, noderesources.NewFit, "Filter", "PreFilter"),
			nodeNames:      []string{"node1", "node2", "node3"},
			pod:            st.MakePod().Name("p").UID("p").Priority(highPriority).Req(veryLargeRes).Obj(),
			pods: []*v1.Pod{
				st.MakePod().Name("p1.1").UID("p1.1").Node("node1").Priority(midPriority).Req(smallRes).StartTime(epochTime).Obj(),
				st.MakePod().Name("p1.2").UID("p1.2").Node("node1").Priority(negPriority).Req(smallRes).StartTime(epochTime).Obj(),
				st.MakePod().Name("p1.3").UID("p1.3").Node("node1").Priority(midPriority).Req(smallRes).StartTime(epochTime).Obj(),
				st.MakePod().Name("p1.4").UID("p1.4").Node("node1").Priority(negPriority).Req(smallRes).StartTime(epochTime).Obj(),
				st.MakePod().Name("p2.1").UID("p2.1").Node("node2").Priority(midPriority).Req(largeRes).StartTime(epochTime).Obj(),
				st.MakePod().Name("p2.2").UID("p2.2").Node("node2").Priority(negPriority).Req(mediumRes).StartTime(epochTime).Obj(),
				st.MakePod().Name("p3.1").UID("p3.1").Node("node3").Priority(midPriority).Req(mediumRes).StartTime(epochTime).Obj(),
				st.MakePod().Name("p3.2").UID("p3.2").Node("node3").Priority(negPriority).Req(smallRes).StartTime(epochTime).Obj(),
				st.MakePod().Name("p3.3").UID("p3.3").Node("node3").Priority(lowPriority).Req(smallRes).StartTime(epochTime).Obj(),
			},
			expected: []string{"node2"},
		},
		{
			// pickOneNodeForPreemption adjusts pod priorities when finding the sum of the victims. This
			// test ensures that the logic works correctly.
			name:           "sum of adjusted priorities is considered",
			registerPlugin: st.RegisterPluginAsExtensions(noderesources.FitName, noderesources.NewFit, "Filter", "PreFilter"),
			nodeNames:      []string{"node1", "node2", "node3"},
			pod:            st.MakePod().Name("p").UID("p").Priority(highPriority).Req(veryLargeRes).Obj(),
			pods: []*v1.Pod{
				st.MakePod().Name("p1.1").UID("p1.1").Node("node1").Priority(midPriority).Req(smallRes).StartTime(epochTime).Obj(),
				st.MakePod().Name("p1.2").UID("p1.2").Node("node1").Priority(negPriority).Req(smallRes).StartTime(epochTime).Obj(),
				st.MakePod().Name("p1.3").UID("p1.3").Node("node1").Priority(negPriority).Req(smallRes).StartTime(epochTime).Obj(),
				st.MakePod().Name("p2.1").UID("p2.1").Node("node2").Priority(midPriority).Req(largeRes).StartTime(epochTime).Obj(),
				st.MakePod().Name("p2.2").UID("p2.2").Node("node2").Priority(negPriority).Req(mediumRes).StartTime(epochTime).Obj(),
				st.MakePod().Name("p3.1").UID("p3.1").Node("node3").Priority(midPriority).Req(mediumRes).StartTime(epochTime).Obj(),
				st.MakePod().Name("p3.2").UID("p3.2").Node("node3").Priority(negPriority).Req(smallRes).StartTime(epochTime).Obj(),
				st.MakePod().Name("p3.3").UID("p3.3").Node("node3").Priority(lowPriority).Req(smallRes).StartTime(epochTime).Obj(),
			},
			expected: []string{"node2"},
		},
		{
			name:           "non-overlapping lowest high priority, sum priorities, and number of pods",
			registerPlugin: st.RegisterPluginAsExtensions(noderesources.FitName, noderesources.NewFit, "Filter", "PreFilter"),
			nodeNames:      []string{"node1", "node2", "node3", "node4"},
			pod:            st.MakePod().Name("p").UID("p").Priority(veryHighPriority).Req(veryLargeRes).Obj(),
			pods: []*v1.Pod{
				st.MakePod().Name("p1.1").UID("p1.1").Node("node1").Priority(midPriority).Req(smallRes).StartTime(epochTime).Obj(),
				st.MakePod().Name("p1.2").UID("p1.2").Node("node1").Priority(lowPriority).Req(smallRes).StartTime(epochTime).Obj(),
				st.MakePod().Name("p1.3").UID("p1.3").Node("node1").Priority(lowPriority).Req(smallRes).StartTime(epochTime).Obj(),
				st.MakePod().Name("p2.1").UID("p2.1").Node("node2").Priority(highPriority).Req(largeRes).StartTime(epochTime).Obj(),
				st.MakePod().Name("p3.1").UID("p3.1").Node("node3").Priority(midPriority).Req(mediumRes).StartTime(epochTime).Obj(),
				st.MakePod().Name("p3.2").UID("p3.2").Node("node3").Priority(lowPriority).Req(smallRes).StartTime(epochTime).Obj(),
				st.MakePod().Name("p3.3").UID("p3.3").Node("node3").Priority(lowPriority).Req(smallRes).StartTime(epochTime).Obj(),
				st.MakePod().Name("p3.4").UID("p3.4").Node("node3").Priority(lowPriority).Req(mediumRes).StartTime(epochTime).Obj(),
				st.MakePod().Name("p4.1").UID("p4.1").Node("node4").Priority(midPriority).Req(mediumRes).StartTime(epochTime).Obj(),
				st.MakePod().Name("p4.2").UID("p4.2").Node("node4").Priority(midPriority).Req(smallRes).StartTime(epochTime).Obj(),
				st.MakePod().Name("p4.3").UID("p4.3").Node("node4").Priority(midPriority).Req(smallRes).StartTime(epochTime).Obj(),
				st.MakePod().Name("p4.4").UID("p4.4").Node("node4").Priority(negPriority).Req(smallRes).StartTime(epochTime).Obj(),
			},
			expected: []string{"node1"},
		},
		{
			name:           "same priority, same number of victims, different start time for each machine's pod",
			registerPlugin: st.RegisterPluginAsExtensions(noderesources.FitName, noderesources.NewFit, "Filter", "PreFilter"),
			nodeNames:      []string{"node1", "node2", "node3"},
			pod:            st.MakePod().Name("p").UID("p").Priority(highPriority).Req(veryLargeRes).Obj(),
			pods: []*v1.Pod{
				st.MakePod().Name("p1.1").UID("p1.1").Node("node1").Priority(midPriority).Req(mediumRes).StartTime(epochTime2).Obj(),
				st.MakePod().Name("p1.2").UID("p1.2").Node("node1").Priority(midPriority).Req(mediumRes).StartTime(epochTime2).Obj(),
				st.MakePod().Name("p2.1").UID("p2.1").Node("node2").Priority(midPriority).Req(mediumRes).StartTime(epochTime3).Obj(),
				st.MakePod().Name("p2.2").UID("p2.2").Node("node2").Priority(midPriority).Req(mediumRes).StartTime(epochTime3).Obj(),
				st.MakePod().Name("p3.1").UID("p3.1").Node("node3").Priority(midPriority).Req(mediumRes).StartTime(epochTime1).Obj(),
				st.MakePod().Name("p3.2").UID("p3.2").Node("node3").Priority(midPriority).Req(mediumRes).StartTime(epochTime1).Obj(),
			},
			expected: []string{"node2"},
		},
		{
			name:           "same priority, same number of victims, different start time for all pods",
			registerPlugin: st.RegisterPluginAsExtensions(noderesources.FitName, noderesources.NewFit, "Filter", "PreFilter"),
			nodeNames:      []string{"node1", "node2", "node3"},
			pod:            st.MakePod().Name("p").UID("p").Priority(highPriority).Req(veryLargeRes).Obj(),
			pods: []*v1.Pod{
				st.MakePod().Name("p1.1").UID("p1.1").Node("node1").Priority(midPriority).Req(mediumRes).StartTime(epochTime4).Obj(),
				st.MakePod().Name("p1.2").UID("p1.2").Node("node1").Priority(midPriority).Req(mediumRes).StartTime(epochTime2).Obj(),
				st.MakePod().Name("p2.1").UID("p2.1").Node("node2").Priority(midPriority).Req(mediumRes).StartTime(epochTime5).Obj(),
				st.MakePod().Name("p2.2").UID("p2.2").Node("node2").Priority(midPriority).Req(mediumRes).StartTime(epochTime1).Obj(),
				st.MakePod().Name("p3.1").UID("p3.1").Node("node3").Priority(midPriority).Req(mediumRes).StartTime(epochTime3).Obj(),
				st.MakePod().Name("p3.2").UID("p3.2").Node("node3").Priority(midPriority).Req(mediumRes).StartTime(epochTime6).Obj(),
			},
			expected: []string{"node3"},
		},
		{
			name:           "different priority, same number of victims, different start time for all pods",
			registerPlugin: st.RegisterPluginAsExtensions(noderesources.FitName, noderesources.NewFit, "Filter", "PreFilter"),
			nodeNames:      []string{"node1", "node2", "node3"},
			pod:            st.MakePod().Name("p").UID("p").Priority(highPriority).Req(veryLargeRes).Obj(),
			pods: []*v1.Pod{
				st.MakePod().Name("p1.1").UID("p1.1").Node("node1").Priority(lowPriority).Req(mediumRes).StartTime(epochTime4).Obj(),
				st.MakePod().Name("p1.2").UID("p1.2").Node("node1").Priority(midPriority).Req(mediumRes).StartTime(epochTime2).Obj(),
				st.MakePod().Name("p2.1").UID("p2.1").Node("node2").Priority(midPriority).Req(mediumRes).StartTime(epochTime6).Obj(),
				st.MakePod().Name("p2.2").UID("p2.2").Node("node2").Priority(lowPriority).Req(mediumRes).StartTime(epochTime1).Obj(),
				st.MakePod().Name("p3.1").UID("p3.1").Node("node3").Priority(lowPriority).Req(mediumRes).StartTime(epochTime3).Obj(),
				st.MakePod().Name("p3.2").UID("p3.2").Node("node3").Priority(midPriority).Req(mediumRes).StartTime(epochTime5).Obj(),
			},
			expected: []string{"node2"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			nodes := make([]*v1.Node, len(tt.nodeNames))
			for i, nodeName := range tt.nodeNames {
				nodes[i] = st.MakeNode().Name(nodeName).Capacity(veryLargeRes).Obj()
			}
			snapshot := internalcache.NewSnapshot(tt.pods, nodes)
			fwk, err := st.NewFramework(
				[]st.RegisterPluginFunc{
					tt.registerPlugin,
					st.RegisterQueueSortPlugin(queuesort.Name, queuesort.New),
					st.RegisterBindPlugin(defaultbinder.Name, defaultbinder.New),
				},
				frameworkruntime.WithPodNominator(internalqueue.NewPodNominator()),
				frameworkruntime.WithSnapshotSharedLister(snapshot),
			)
			if err != nil {
				t.Fatal(err)
			}

			state := framework.NewCycleState()
			// Some tests rely on PreFilter plugin to compute its CycleState.
			if status := fwk.RunPreFilterPlugins(context.Background(), state, tt.pod); !status.IsSuccess() {
				t.Errorf("Unexpected PreFilter Status: %v", status)
			}
			nodeInfos, err := snapshot.NodeInfos().List()
			if err != nil {
				t.Fatal(err)
			}
			candidates := dryRunPreemption(context.Background(), fwk.PreemptHandle(), state, tt.pod, nodeInfos, nil)
			s := SelectCandidate(candidates)
			found := false
			for _, nodeName := range tt.expected {
				if nodeName == s.Name() {
					found = true
					break
				}
			}
			if !found {
				t.Errorf("expect any node in %v, but got %v", tt.expected, s.Name())
			}
		})
	}
}

func TestPodEligibleToPreemptOthers(t *testing.T) {
	tests := []struct {
		name                string
		pod                 *v1.Pod
		pods                []*v1.Pod
		nodes               []string
		nominatedNodeStatus *framework.Status
		expected            bool
	}{
		{
			name:                "Pod with nominated node",
			pod:                 st.MakePod().Name("p_with_nominated_node").UID("p").Priority(highPriority).NominatedNodeName("node1").Obj(),
			pods:                []*v1.Pod{st.MakePod().Name("p1").UID("p1").Priority(lowPriority).Node("node1").Terminating().Obj()},
			nodes:               []string{"node1"},
			nominatedNodeStatus: framework.NewStatus(framework.UnschedulableAndUnresolvable, tainttoleration.ErrReasonNotMatch),
			expected:            true,
		},
		{
			name:                "Pod with nominated node, but without nominated node status",
			pod:                 st.MakePod().Name("p_without_status").UID("p").Priority(highPriority).NominatedNodeName("node1").Obj(),
			pods:                []*v1.Pod{st.MakePod().Name("p1").UID("p1").Priority(lowPriority).Node("node1").Terminating().Obj()},
			nodes:               []string{"node1"},
			nominatedNodeStatus: nil,
			expected:            false,
		},
		{
			name:                "Pod without nominated node",
			pod:                 st.MakePod().Name("p_without_nominated_node").UID("p").Priority(highPriority).Obj(),
			pods:                []*v1.Pod{},
			nodes:               []string{},
			nominatedNodeStatus: nil,
			expected:            true,
		},
		{
			name:                "Pod with 'PreemptNever' preemption policy",
			pod:                 st.MakePod().Name("p_with_preempt_never_policy").UID("p").Priority(highPriority).PreemptionPolicy(v1.PreemptNever).Obj(),
			pods:                []*v1.Pod{},
			nodes:               []string{},
			nominatedNodeStatus: nil,
			expected:            false,
		},
	}

	for _, test := range tests {
		var nodes []*v1.Node
		for _, n := range test.nodes {
			nodes = append(nodes, st.MakeNode().Name(n).Obj())
		}
		snapshot := internalcache.NewSnapshot(test.pods, nodes)
		if got := PodEligibleToPreemptOthers(test.pod, snapshot.NodeInfos(), test.nominatedNodeStatus); got != test.expected {
			t.Errorf("expected %t, got %t for pod: %s", test.expected, got, test.pod.Name)
		}
	}
}

func TestNodesWherePreemptionMightHelp(t *testing.T) {
	// Prepare 4 nodes names.
	nodeNames := []string{"node1", "node2", "node3", "node4"}

	tests := []struct {
		name          string
		nodesStatuses framework.NodeToStatusMap
		expected      map[string]bool // set of expected node names. Value is ignored.
	}{
		{
			name: "No node should be attempted",
			nodesStatuses: framework.NodeToStatusMap{
				"node1": framework.NewStatus(framework.UnschedulableAndUnresolvable, nodeaffinity.ErrReason),
				"node2": framework.NewStatus(framework.UnschedulableAndUnresolvable, nodename.ErrReason),
				"node3": framework.NewStatus(framework.UnschedulableAndUnresolvable, tainttoleration.ErrReasonNotMatch),
				"node4": framework.NewStatus(framework.UnschedulableAndUnresolvable, nodelabel.ErrReasonPresenceViolated),
			},
			expected: map[string]bool{},
		},
		{
			name: "ErrReasonAffinityNotMatch should be tried as it indicates that the pod is unschedulable due to inter-pod affinity or anti-affinity",
			nodesStatuses: framework.NodeToStatusMap{
				"node1": framework.NewStatus(framework.Unschedulable, interpodaffinity.ErrReasonAffinityNotMatch),
				"node2": framework.NewStatus(framework.UnschedulableAndUnresolvable, nodename.ErrReason),
				"node3": framework.NewStatus(framework.UnschedulableAndUnresolvable, nodeunschedulable.ErrReasonUnschedulable),
			},
			expected: map[string]bool{"node1": true, "node4": true},
		},
		{
			name: "pod with both pod affinity and anti-affinity should be tried",
			nodesStatuses: framework.NodeToStatusMap{
				"node1": framework.NewStatus(framework.Unschedulable, interpodaffinity.ErrReasonAffinityNotMatch),
				"node2": framework.NewStatus(framework.UnschedulableAndUnresolvable, nodename.ErrReason),
			},
			expected: map[string]bool{"node1": true, "node3": true, "node4": true},
		},
		{
			name: "ErrReasonAffinityRulesNotMatch should not be tried as it indicates that the pod is unschedulable due to inter-pod affinity, but ErrReasonAffinityNotMatch should be tried as it indicates that the pod is unschedulable due to inter-pod affinity or anti-affinity",
			nodesStatuses: framework.NodeToStatusMap{
				"node1": framework.NewStatus(framework.UnschedulableAndUnresolvable, interpodaffinity.ErrReasonAffinityRulesNotMatch),
				"node2": framework.NewStatus(framework.Unschedulable, interpodaffinity.ErrReasonAffinityNotMatch),
			},
			expected: map[string]bool{"node2": true, "node3": true, "node4": true},
		},
		{
			name: "Mix of failed predicates works fine",
			nodesStatuses: framework.NodeToStatusMap{
				"node1": framework.NewStatus(framework.UnschedulableAndUnresolvable, volumerestrictions.ErrReasonDiskConflict),
				"node2": framework.NewStatus(framework.Unschedulable, fmt.Sprintf("Insufficient %v", v1.ResourceMemory)),
			},
			expected: map[string]bool{"node2": true, "node3": true, "node4": true},
		},
		{
			name: "Node condition errors should be considered unresolvable",
			nodesStatuses: framework.NodeToStatusMap{
				"node1": framework.NewStatus(framework.UnschedulableAndUnresolvable, nodeunschedulable.ErrReasonUnknownCondition),
			},
			expected: map[string]bool{"node2": true, "node3": true, "node4": true},
		},
		{
			name: "ErrVolume... errors should not be tried as it indicates that the pod is unschedulable due to no matching volumes for pod on node",
			nodesStatuses: framework.NodeToStatusMap{
				"node1": framework.NewStatus(framework.UnschedulableAndUnresolvable, volumezone.ErrReasonConflict),
				"node2": framework.NewStatus(framework.UnschedulableAndUnresolvable, string(volumescheduling.ErrReasonNodeConflict)),
				"node3": framework.NewStatus(framework.UnschedulableAndUnresolvable, string(volumescheduling.ErrReasonBindConflict)),
			},
			expected: map[string]bool{"node4": true},
		},
		{
			name: "ErrReasonConstraintsNotMatch should be tried as it indicates that the pod is unschedulable due to topology spread constraints",
			nodesStatuses: framework.NodeToStatusMap{
				"node1": framework.NewStatus(framework.Unschedulable, podtopologyspread.ErrReasonConstraintsNotMatch),
				"node2": framework.NewStatus(framework.UnschedulableAndUnresolvable, nodename.ErrReason),
				"node3": framework.NewStatus(framework.Unschedulable, podtopologyspread.ErrReasonConstraintsNotMatch),
			},
			expected: map[string]bool{"node1": true, "node3": true, "node4": true},
		},
		{
			name: "UnschedulableAndUnresolvable status should be skipped but Unschedulable should be tried",
			nodesStatuses: framework.NodeToStatusMap{
				"node2": framework.NewStatus(framework.UnschedulableAndUnresolvable, ""),
				"node3": framework.NewStatus(framework.Unschedulable, ""),
				"node4": framework.NewStatus(framework.UnschedulableAndUnresolvable, ""),
			},
			expected: map[string]bool{"node1": true, "node3": true},
		},
		{
			name: "ErrReasonNodeLabelNotMatch should not be tried as it indicates that the pod is unschedulable due to node doesn't have the required label",
			nodesStatuses: framework.NodeToStatusMap{
				"node2": framework.NewStatus(framework.UnschedulableAndUnresolvable, podtopologyspread.ErrReasonNodeLabelNotMatch),
				"node3": framework.NewStatus(framework.Unschedulable, ""),
				"node4": framework.NewStatus(framework.UnschedulableAndUnresolvable, ""),
			},
			expected: map[string]bool{"node1": true, "node3": true},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var nodeInfos []*framework.NodeInfo
			for _, name := range nodeNames {
				ni := framework.NewNodeInfo()
				ni.SetNode(st.MakeNode().Name(name).Obj())
				nodeInfos = append(nodeInfos, ni)
			}
			nodes := nodesWherePreemptionMightHelp(nodeInfos, tt.nodesStatuses)
			if len(tt.expected) != len(nodes) {
				t.Errorf("number of nodes is not the same as expected. exptectd: %d, got: %d. Nodes: %v", len(tt.expected), len(nodes), nodes)
			}
			for _, node := range nodes {
				name := node.Node().Name
				if _, found := tt.expected[name]; !found {
					t.Errorf("node %v is not expected.", name)
				}
			}
		})
	}
}

func TestPreempt(t *testing.T) {
	tests := []struct {
		name           string
		pod            *v1.Pod
		pods           []*v1.Pod
		extenders      []*st.FakeExtender
		nodeNames      []string
		registerPlugin st.RegisterPluginFunc
		expectedNode   string
		expectedPods   []string // list of preempted pods
	}{
		{
			name: "basic preemption logic",
			pod:  st.MakePod().Name("p").UID("p").Priority(highPriority).Req(veryLargeRes).PreemptionPolicy(v1.PreemptLowerPriority).Obj(),
			pods: []*v1.Pod{
				st.MakePod().Name("p1.1").UID("p1.1").Node("node1").Priority(lowPriority).Req(smallRes).Obj(),
				st.MakePod().Name("p1.2").UID("p1.2").Node("node1").Priority(lowPriority).Req(smallRes).Obj(),
				st.MakePod().Name("p2.1").UID("p2.1").Node("node2").Priority(highPriority).Req(largeRes).Obj(),
				st.MakePod().Name("p3.1").UID("p3.1").Node("node3").Priority(midPriority).Req(mediumRes).Obj(),
			},
			nodeNames:      []string{"node1", "node2", "node3"},
			registerPlugin: st.RegisterPluginAsExtensions(noderesources.FitName, noderesources.NewFit, "Filter", "PreFilter"),
			expectedNode:   "node1",
			expectedPods:   []string{"p1.1", "p1.2"},
		},
		{
			name: "preemption for topology spread constraints",
			pod: st.MakePod().Name("p").UID("p").Label("foo", "").Priority(highPriority).
				SpreadConstraint(1, "zone", v1.DoNotSchedule, st.MakeLabelSelector().Exists("foo").Obj()).
				SpreadConstraint(1, "hostname", v1.DoNotSchedule, st.MakeLabelSelector().Exists("foo").Obj()).
				Obj(),
			pods: []*v1.Pod{
				st.MakePod().Name("p-a1").UID("p-a1").Node("node-a").Label("foo", "").Priority(highPriority).Obj(),
				st.MakePod().Name("p-a2").UID("p-a2").Node("node-a").Label("foo", "").Priority(highPriority).Obj(),
				st.MakePod().Name("p-b1").UID("p-b1").Node("node-b").Label("foo", "").Priority(lowPriority).Obj(),
				st.MakePod().Name("p-x1").UID("p-x1").Node("node-x").Label("foo", "").Priority(highPriority).Obj(),
				st.MakePod().Name("p-x2").UID("p-x2").Node("node-x").Label("foo", "").Priority(highPriority).Obj(),
			},
			nodeNames:      []string{"node-a/zone1", "node-b/zone1", "node-x/zone2"},
			registerPlugin: st.RegisterPluginAsExtensions(podtopologyspread.Name, podtopologyspread.New, "PreFilter", "Filter"),
			expectedNode:   "node-b",
			expectedPods:   []string{"p-b1"},
		},
		{
			name: "Scheduler extenders allow only machine1, otherwise machine3 would have been chosen",
			pod:  st.MakePod().Name("p").UID("p").Priority(highPriority).Req(veryLargeRes).PreemptionPolicy(v1.PreemptLowerPriority).Obj(),
			pods: []*v1.Pod{
				st.MakePod().Name("p1.1").UID("p1.1").Node("machine1").Priority(midPriority).Req(smallRes).Obj(),
				st.MakePod().Name("p1.2").UID("p1.2").Node("machine1").Priority(lowPriority).Req(smallRes).Obj(),
				st.MakePod().Name("p2.1").UID("p2.1").Node("machine3").Priority(midPriority).Req(largeRes).Obj(),
			},
			nodeNames: []string{"machine1", "machine2", "machine3"},
			extenders: []*st.FakeExtender{
				{Predicates: []st.FitPredicate{st.TruePredicateExtender}},
				{Predicates: []st.FitPredicate{st.Machine1PredicateExtender}},
			},
			registerPlugin: st.RegisterPluginAsExtensions(noderesources.FitName, noderesources.NewFit, "Filter", "PreFilter"),
			expectedNode:   "machine1",
			expectedPods:   []string{"p1.1", "p1.2"},
		},
		{
			name: "Scheduler extenders do not allow any preemption",
			pod:  st.MakePod().Name("p").UID("p").Priority(highPriority).Req(veryLargeRes).PreemptionPolicy(v1.PreemptLowerPriority).Obj(),
			pods: []*v1.Pod{
				st.MakePod().Name("p1.1").UID("p1.1").Node("node1").Priority(midPriority).Req(smallRes).Obj(),
				st.MakePod().Name("p1.2").UID("p1.2").Node("node1").Priority(lowPriority).Req(smallRes).Obj(),
				st.MakePod().Name("p2.1").UID("p2.1").Node("node2").Priority(midPriority).Req(largeRes).Obj(),
			},
			nodeNames: []string{"node1", "node2", "node3"},
			extenders: []*st.FakeExtender{
				{Predicates: []st.FitPredicate{st.FalsePredicateExtender}},
			},
			registerPlugin: st.RegisterPluginAsExtensions(noderesources.FitName, noderesources.NewFit, "Filter", "PreFilter"),
			expectedNode:   "",
			expectedPods:   []string{},
		},
		{
			name: "One scheduler extender allows only machine1, the other returns error but ignorable. Only machine1 would be chosen",
			pod:  st.MakePod().Name("p").UID("p").Priority(highPriority).Req(veryLargeRes).PreemptionPolicy(v1.PreemptLowerPriority).Obj(),
			pods: []*v1.Pod{
				st.MakePod().Name("p1.1").UID("p1.1").Node("machine1").Priority(midPriority).Req(smallRes).Obj(),
				st.MakePod().Name("p1.2").UID("p1.2").Node("machine1").Priority(lowPriority).Req(smallRes).Obj(),
				st.MakePod().Name("p2.1").UID("p2.1").Node("machine2").Priority(midPriority).Req(largeRes).Obj(),
			},
			nodeNames: []string{"machine1", "machine2", "machine3"},
			extenders: []*st.FakeExtender{
				{Predicates: []st.FitPredicate{st.ErrorPredicateExtender}, Ignorable: true},
				{Predicates: []st.FitPredicate{st.Machine1PredicateExtender}},
			},
			registerPlugin: st.RegisterPluginAsExtensions(noderesources.FitName, noderesources.NewFit, "Filter", "PreFilter"),
			expectedNode:   "machine1",
			expectedPods:   []string{"p1.1", "p1.2"},
		},
		{
			name: "One scheduler extender allows only machine1, but it is not interested in given pod, otherwise machine1 would have been chosen",
			pod:  st.MakePod().Name("p").UID("p").Priority(highPriority).Req(veryLargeRes).PreemptionPolicy(v1.PreemptLowerPriority).Obj(),
			pods: []*v1.Pod{
				st.MakePod().Name("p1.1").UID("p1.1").Node("machine1").Priority(midPriority).Req(smallRes).Obj(),
				st.MakePod().Name("p1.2").UID("p1.2").Node("machine1").Priority(lowPriority).Req(smallRes).Obj(),
				st.MakePod().Name("p2.1").UID("p2.1").Node("machine2").Priority(midPriority).Req(largeRes).Obj(),
			},
			nodeNames: []string{"machine1", "machine2"},
			extenders: []*st.FakeExtender{
				{Predicates: []st.FitPredicate{st.Machine1PredicateExtender}, UnInterested: true},
				{Predicates: []st.FitPredicate{st.TruePredicateExtender}},
			},
			registerPlugin: st.RegisterPluginAsExtensions(noderesources.FitName, noderesources.NewFit, "Filter", "PreFilter"),
			// sum of priorities of all victims on machine1 is larger than machine2, machine2 is chosen.
			expectedNode: "machine2",
			expectedPods: []string{"p2.1"},
		},
		{
			name: "no preempting in pod",
			pod:  st.MakePod().Name("p").UID("p").Priority(highPriority).Req(veryLargeRes).PreemptionPolicy(v1.PreemptNever).Obj(),
			pods: []*v1.Pod{
				st.MakePod().Name("p1.1").UID("p1.1").Node("node1").Priority(lowPriority).Req(smallRes).Obj(),
				st.MakePod().Name("p1.2").UID("p1.2").Node("node1").Priority(lowPriority).Req(smallRes).Obj(),
				st.MakePod().Name("p2.1").UID("p2.1").Node("node2").Priority(highPriority).Req(largeRes).Obj(),
				st.MakePod().Name("p3.1").UID("p3.1").Node("node3").Priority(midPriority).Req(mediumRes).Obj(),
			},
			nodeNames:      []string{"node1", "node2", "node3"},
			registerPlugin: st.RegisterPluginAsExtensions(noderesources.FitName, noderesources.NewFit, "Filter", "PreFilter"),
			expectedNode:   "",
			expectedPods:   nil,
		},
		{
			name: "PreemptionPolicy is nil",
			pod:  st.MakePod().Name("p").UID("p").Priority(highPriority).Req(veryLargeRes).Obj(),
			pods: []*v1.Pod{
				st.MakePod().Name("p1.1").UID("p1.1").Node("node1").Priority(lowPriority).Req(smallRes).Obj(),
				st.MakePod().Name("p1.2").UID("p1.2").Node("node1").Priority(lowPriority).Req(smallRes).Obj(),
				st.MakePod().Name("p2.1").UID("p2.1").Node("node2").Priority(highPriority).Req(largeRes).Obj(),
				st.MakePod().Name("p3.1").UID("p3.1").Node("node3").Priority(midPriority).Req(mediumRes).Obj(),
			},
			nodeNames:      []string{"node1", "node2", "node3"},
			registerPlugin: st.RegisterPluginAsExtensions(noderesources.FitName, noderesources.NewFit, "Filter", "PreFilter"),
			expectedNode:   "node1",
			expectedPods:   []string{"p1.1", "p1.2"},
		},
	}

	labelKeys := []string{"hostname", "zone", "region"}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			apiObjs := mergeObjs(test.pod, test.pods)
			client := clientsetfake.NewSimpleClientset(apiObjs...)
			deletedPodNames := make(sets.String)
			client.PrependReactor("delete", "pods", func(action clienttesting.Action) (bool, runtime.Object, error) {
				deletedPodNames.Insert(action.(clienttesting.DeleteAction).GetName())
				return true, nil, nil
			})

			stop := make(chan struct{})
			defer close(stop)

			cache := internalcache.New(time.Duration(0), stop)
			for _, pod := range test.pods {
				cache.AddPod(pod)
			}
			cachedNodeInfoMap := map[string]*framework.NodeInfo{}
			nodes := make([]*v1.Node, len(test.nodeNames))
			for i, name := range test.nodeNames {
				node := st.MakeNode().Name(name).Capacity(veryLargeRes).Obj()
				// Split node name by '/' to form labels in a format of
				// {"hostname": node.Name[0], "zone": node.Name[1], "region": node.Name[2]}
				node.ObjectMeta.Labels = make(map[string]string)
				for i, label := range strings.Split(node.Name, "/") {
					node.ObjectMeta.Labels[labelKeys[i]] = label
				}
				node.Name = node.ObjectMeta.Labels["hostname"]
				cache.AddNode(node)
				nodes[i] = node

				// Set nodeInfo to extenders to mock extenders' cache for preemption.
				cachedNodeInfo := framework.NewNodeInfo()
				cachedNodeInfo.SetNode(node)
				cachedNodeInfoMap[node.Name] = cachedNodeInfo
			}
			var extenders []framework.Extender
			for _, extender := range test.extenders {
				// Set nodeInfoMap as extenders cached node information.
				extender.CachedNodeNameToInfo = cachedNodeInfoMap
				extenders = append(extenders, extender)
			}

			informerFactory := informers.NewSharedInformerFactory(client, 0)
			fwk, err := st.NewFramework(
				[]st.RegisterPluginFunc{
					test.registerPlugin,
					st.RegisterQueueSortPlugin(queuesort.Name, queuesort.New),
					st.RegisterBindPlugin(defaultbinder.Name, defaultbinder.New),
				},
				frameworkruntime.WithClientSet(client),
				frameworkruntime.WithEventRecorder(&events.FakeRecorder{}),
				frameworkruntime.WithExtenders(extenders),
				frameworkruntime.WithPodNominator(internalqueue.NewPodNominator()),
				frameworkruntime.WithSnapshotSharedLister(internalcache.NewSnapshot(test.pods, nodes)),
				frameworkruntime.WithInformerFactory(informerFactory),
			)
			if err != nil {
				t.Fatal(err)
			}

			state := framework.NewCycleState()
			// Some tests rely on PreFilter plugin to compute its CycleState.
			preFilterStatus := fwk.RunPreFilterPlugins(context.Background(), state, test.pod)
			if !preFilterStatus.IsSuccess() {
				t.Errorf("Unexpected preFilterStatus: %v", preFilterStatus)
			}
			// Call preempt and check the expected results.
			pl := DefaultPreemption{
				fh:        fwk,
				pdbLister: getPDBLister(informerFactory),
			}
			node, err := pl.preempt(context.Background(), state, test.pod, make(framework.NodeToStatusMap))
			if err != nil {
				t.Errorf("unexpected error in preemption: %v", err)
			}
			if len(node) != 0 && node != test.expectedNode {
				t.Errorf("expected node: %v, got: %v", test.expectedNode, node)
			}
			if len(node) == 0 && len(test.expectedNode) != 0 {
				t.Errorf("expected node: %v, got: nothing", test.expectedNode)
			}
			if len(deletedPodNames) != len(test.expectedPods) {
				t.Errorf("expected %v pods, got %v.", len(test.expectedPods), len(deletedPodNames))
			}
			for victimName := range deletedPodNames {
				found := false
				for _, expPod := range test.expectedPods {
					if expPod == victimName {
						found = true
						break
					}
				}
				if !found {
					t.Errorf("pod %v is not expected to be a victim.", victimName)
				}
			}
			test.pod.Status.NominatedNodeName = node
			client.CoreV1().Pods(test.pod.Namespace).Update(context.TODO(), test.pod, metav1.UpdateOptions{})

			// Manually set the deleted Pods' deletionTimestamp to non-nil.
			for _, pod := range test.pods {
				if deletedPodNames.Has(pod.Name) {
					now := metav1.Now()
					pod.DeletionTimestamp = &now
					deletedPodNames.Delete(pod.Name)
				}
			}

			// Call preempt again and make sure it doesn't preempt any more pods.
			node, err = pl.preempt(context.Background(), state, test.pod, make(framework.NodeToStatusMap))
			if err != nil {
				t.Errorf("unexpected error in preemption: %v", err)
			}
			if len(node) != 0 && len(deletedPodNames) > 0 {
				t.Errorf("didn't expect any more preemption. Node %v is selected for preemption.", node)
			}
		})
	}
}
