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

package cpumanager

import (
	"fmt"
	"reflect"
	"testing"

	"k8s.io/api/core/v1"
	"k8s.io/kubernetes/pkg/kubelet/cm/cpumanager/state"
	"k8s.io/kubernetes/pkg/kubelet/cm/cpumanager/topology"
	"k8s.io/kubernetes/pkg/kubelet/cm/cpuset"
	"k8s.io/kubernetes/pkg/kubelet/cm/topologymanager"
	"k8s.io/kubernetes/pkg/kubelet/cm/topologymanager/bitmask"
)

type staticPolicyTest struct {
	description     string
	topo            *topology.CPUTopology
	numReservedCPUs int
	containerID     string
	stAssignments   state.ContainerCPUAssignments
	stDefaultCPUSet cpuset.CPUSet
	pod             *v1.Pod
	expErr          error
	expCPUAlloc     bool
	expCSet         cpuset.CPUSet
	expPanic        bool
}

type staticPolicyMultiContainerTest struct {
	description      string
	topo             *topology.CPUTopology
	numReservedCPUs  int
	initContainerIDs []string
	containerIDs     []string
	stAssignments    state.ContainerCPUAssignments
	stDefaultCPUSet  cpuset.CPUSet
	pod              *v1.Pod
	expInitCSets     []cpuset.CPUSet
	expCSets         []cpuset.CPUSet
}

func TestStaticPolicyName(t *testing.T) {
	policy := NewStaticPolicy(topoSingleSocketHT, 1, topologymanager.NewFakeManager())

	policyName := policy.Name()
	if policyName != "static" {
		t.Errorf("StaticPolicy Name() error. expected: static, returned: %v",
			policyName)
	}
}

func TestStaticPolicyStart(t *testing.T) {
	testCases := []staticPolicyTest{
		{
			description: "non-corrupted state",
			topo:        topoDualSocketHT,
			stAssignments: state.ContainerCPUAssignments{
				"0": cpuset.NewCPUSet(0),
			},
			stDefaultCPUSet: cpuset.NewCPUSet(1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11),
			expCSet:         cpuset.NewCPUSet(1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11),
		},
		{
			description:     "empty cpuset",
			topo:            topoDualSocketHT,
			numReservedCPUs: 1,
			stAssignments:   state.ContainerCPUAssignments{},
			stDefaultCPUSet: cpuset.NewCPUSet(),
			expCSet:         cpuset.NewCPUSet(0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11),
		},
		{
			description:     "reserved cores 0 & 6 are not present in available cpuset",
			topo:            topoDualSocketHT,
			numReservedCPUs: 2,
			stAssignments:   state.ContainerCPUAssignments{},
			stDefaultCPUSet: cpuset.NewCPUSet(0, 1),
			expPanic:        true,
		},
		{
			description: "assigned core 2 is still present in available cpuset",
			topo:        topoDualSocketHT,
			stAssignments: state.ContainerCPUAssignments{
				"0": cpuset.NewCPUSet(0, 1, 2),
			},
			stDefaultCPUSet: cpuset.NewCPUSet(2, 3, 4, 5, 6, 7, 8, 9, 10, 11),
			expPanic:        true,
		},
		{
			description: "core 12 is not present in topology but is in state cpuset",
			topo:        topoDualSocketHT,
			stAssignments: state.ContainerCPUAssignments{
				"0": cpuset.NewCPUSet(0, 1, 2),
				"1": cpuset.NewCPUSet(3, 4),
			},
			stDefaultCPUSet: cpuset.NewCPUSet(5, 6, 7, 8, 9, 10, 11, 12),
			expPanic:        true,
		},
		{
			description: "core 11 is present in topology but is not in state cpuset",
			topo:        topoDualSocketHT,
			stAssignments: state.ContainerCPUAssignments{
				"0": cpuset.NewCPUSet(0, 1, 2),
				"1": cpuset.NewCPUSet(3, 4),
			},
			stDefaultCPUSet: cpuset.NewCPUSet(5, 6, 7, 8, 9, 10),
			expPanic:        true,
		},
	}
	for _, testCase := range testCases {
		t.Run(testCase.description, func(t *testing.T) {
			defer func() {
				if err := recover(); err != nil {
					if !testCase.expPanic {
						t.Errorf("unexpected panic occurred: %q", err)
					}
				} else if testCase.expPanic {
					t.Error("expected panic doesn't occurred")
				}
			}()
			policy := NewStaticPolicy(testCase.topo, testCase.numReservedCPUs, topologymanager.NewFakeManager()).(*staticPolicy)
			st := &mockState{
				assignments:   testCase.stAssignments,
				defaultCPUSet: testCase.stDefaultCPUSet,
			}
			policy.Start(st)

			if !testCase.stDefaultCPUSet.IsEmpty() {
				for cpuid := 1; cpuid < policy.topology.NumCPUs; cpuid++ {
					if !st.defaultCPUSet.Contains(cpuid) {
						t.Errorf("StaticPolicy Start() error. expected cpuid %d to be present in defaultCPUSet", cpuid)
					}
				}
			}
			if !st.GetDefaultCPUSet().Equals(testCase.expCSet) {
				t.Errorf("State CPUSet is different than expected. Have %q wants: %q", st.GetDefaultCPUSet(),
					testCase.expCSet)
			}

		})
	}
}

func TestStaticPolicyAdd(t *testing.T) {
	largeTopoBuilder := cpuset.NewBuilder()
	largeTopoSock0Builder := cpuset.NewBuilder()
	largeTopoSock1Builder := cpuset.NewBuilder()
	largeTopo := *topoQuadSocketFourWayHT
	for cpuid, val := range largeTopo.CPUDetails {
		largeTopoBuilder.Add(cpuid)
		if val.SocketID == 0 {
			largeTopoSock0Builder.Add(cpuid)
		} else if val.SocketID == 1 {
			largeTopoSock1Builder.Add(cpuid)
		}
	}
	largeTopoCPUSet := largeTopoBuilder.Result()
	largeTopoSock0CPUSet := largeTopoSock0Builder.Result()
	largeTopoSock1CPUSet := largeTopoSock1Builder.Result()

	testCases := []staticPolicyTest{
		{
			description:     "GuPodSingleCore, SingleSocketHT, ExpectError",
			topo:            topoSingleSocketHT,
			numReservedCPUs: 1,
			containerID:     "fakeID2",
			stAssignments:   state.ContainerCPUAssignments{},
			stDefaultCPUSet: cpuset.NewCPUSet(0, 1, 2, 3, 4, 5, 6, 7),
			pod:             makePod("8000m", "8000m"),
			expErr:          fmt.Errorf("not enough cpus available to satisfy request"),
			expCPUAlloc:     false,
			expCSet:         cpuset.NewCPUSet(),
		},
		{
			description:     "GuPodSingleCore, SingleSocketHT, ExpectAllocOneCPU",
			topo:            topoSingleSocketHT,
			numReservedCPUs: 1,
			containerID:     "fakeID2",
			stAssignments:   state.ContainerCPUAssignments{},
			stDefaultCPUSet: cpuset.NewCPUSet(0, 1, 2, 3, 4, 5, 6, 7),
			pod:             makePod("1000m", "1000m"),
			expErr:          nil,
			expCPUAlloc:     true,
			expCSet:         cpuset.NewCPUSet(4), // expect sibling of partial core
		},
		{
			description:     "GuPodMultipleCores, SingleSocketHT, ExpectAllocOneCore",
			topo:            topoSingleSocketHT,
			numReservedCPUs: 1,
			containerID:     "fakeID3",
			stAssignments: state.ContainerCPUAssignments{
				"fakeID100": cpuset.NewCPUSet(2, 3, 6, 7),
			},
			stDefaultCPUSet: cpuset.NewCPUSet(0, 1, 4, 5),
			pod:             makePod("2000m", "2000m"),
			expErr:          nil,
			expCPUAlloc:     true,
			expCSet:         cpuset.NewCPUSet(1, 5),
		},
		{
			description:     "GuPodMultipleCores, DualSocketHT, ExpectAllocOneSocket",
			topo:            topoDualSocketHT,
			numReservedCPUs: 1,
			containerID:     "fakeID3",
			stAssignments: state.ContainerCPUAssignments{
				"fakeID100": cpuset.NewCPUSet(2),
			},
			stDefaultCPUSet: cpuset.NewCPUSet(0, 1, 3, 4, 5, 6, 7, 8, 9, 10, 11),
			pod:             makePod("6000m", "6000m"),
			expErr:          nil,
			expCPUAlloc:     true,
			expCSet:         cpuset.NewCPUSet(1, 3, 5, 7, 9, 11),
		},
		{
			description:     "GuPodMultipleCores, DualSocketHT, ExpectAllocThreeCores",
			topo:            topoDualSocketHT,
			numReservedCPUs: 1,
			containerID:     "fakeID3",
			stAssignments: state.ContainerCPUAssignments{
				"fakeID100": cpuset.NewCPUSet(1, 5),
			},
			stDefaultCPUSet: cpuset.NewCPUSet(0, 2, 3, 4, 6, 7, 8, 9, 10, 11),
			pod:             makePod("6000m", "6000m"),
			expErr:          nil,
			expCPUAlloc:     true,
			expCSet:         cpuset.NewCPUSet(2, 3, 4, 8, 9, 10),
		},
		{
			description:     "GuPodMultipleCores, DualSocketNoHT, ExpectAllocOneSocket",
			topo:            topoDualSocketNoHT,
			numReservedCPUs: 1,
			containerID:     "fakeID1",
			stAssignments: state.ContainerCPUAssignments{
				"fakeID100": cpuset.NewCPUSet(),
			},
			stDefaultCPUSet: cpuset.NewCPUSet(0, 1, 3, 4, 5, 6, 7),
			pod:             makePod("4000m", "4000m"),
			expErr:          nil,
			expCPUAlloc:     true,
			expCSet:         cpuset.NewCPUSet(4, 5, 6, 7),
		},
		{
			description:     "GuPodMultipleCores, DualSocketNoHT, ExpectAllocFourCores",
			topo:            topoDualSocketNoHT,
			numReservedCPUs: 1,
			containerID:     "fakeID1",
			stAssignments: state.ContainerCPUAssignments{
				"fakeID100": cpuset.NewCPUSet(4, 5),
			},
			stDefaultCPUSet: cpuset.NewCPUSet(0, 1, 3, 6, 7),
			pod:             makePod("4000m", "4000m"),
			expErr:          nil,
			expCPUAlloc:     true,
			expCSet:         cpuset.NewCPUSet(1, 3, 6, 7),
		},
		{
			description:     "GuPodMultipleCores, DualSocketHT, ExpectAllocOneSocketOneCore",
			topo:            topoDualSocketHT,
			numReservedCPUs: 1,
			containerID:     "fakeID3",
			stAssignments: state.ContainerCPUAssignments{
				"fakeID100": cpuset.NewCPUSet(2),
			},
			stDefaultCPUSet: cpuset.NewCPUSet(0, 1, 3, 4, 5, 6, 7, 8, 9, 10, 11),
			pod:             makePod("8000m", "8000m"),
			expErr:          nil,
			expCPUAlloc:     true,
			expCSet:         cpuset.NewCPUSet(1, 3, 4, 5, 7, 9, 10, 11),
		},
		{
			description:     "NonGuPod, SingleSocketHT, NoAlloc",
			topo:            topoSingleSocketHT,
			numReservedCPUs: 1,
			containerID:     "fakeID1",
			stAssignments:   state.ContainerCPUAssignments{},
			stDefaultCPUSet: cpuset.NewCPUSet(0, 1, 2, 3, 4, 5, 6, 7),
			pod:             makePod("1000m", "2000m"),
			expErr:          nil,
			expCPUAlloc:     false,
			expCSet:         cpuset.NewCPUSet(),
		},
		{
			description:     "GuPodNonIntegerCore, SingleSocketHT, NoAlloc",
			topo:            topoSingleSocketHT,
			numReservedCPUs: 1,
			containerID:     "fakeID4",
			stAssignments:   state.ContainerCPUAssignments{},
			stDefaultCPUSet: cpuset.NewCPUSet(0, 1, 2, 3, 4, 5, 6, 7),
			pod:             makePod("977m", "977m"),
			expErr:          nil,
			expCPUAlloc:     false,
			expCSet:         cpuset.NewCPUSet(),
		},
		{
			description:     "GuPodMultipleCores, SingleSocketHT, NoAllocExpectError",
			topo:            topoSingleSocketHT,
			numReservedCPUs: 1,
			containerID:     "fakeID5",
			stAssignments: state.ContainerCPUAssignments{
				"fakeID100": cpuset.NewCPUSet(1, 2, 3, 4, 5, 6),
			},
			stDefaultCPUSet: cpuset.NewCPUSet(0, 7),
			pod:             makePod("2000m", "2000m"),
			expErr:          fmt.Errorf("not enough cpus available to satisfy request"),
			expCPUAlloc:     false,
			expCSet:         cpuset.NewCPUSet(),
		},
		{
			description:     "GuPodMultipleCores, DualSocketHT, NoAllocExpectError",
			topo:            topoDualSocketHT,
			numReservedCPUs: 1,
			containerID:     "fakeID5",
			stAssignments: state.ContainerCPUAssignments{
				"fakeID100": cpuset.NewCPUSet(1, 2, 3),
			},
			stDefaultCPUSet: cpuset.NewCPUSet(0, 4, 5, 6, 7, 8, 9, 10, 11),
			pod:             makePod("10000m", "10000m"),
			expErr:          fmt.Errorf("not enough cpus available to satisfy request"),
			expCPUAlloc:     false,
			expCSet:         cpuset.NewCPUSet(),
		},
		{
			// All the CPUs from Socket 0 are available. Some CPUs from each
			// Socket have been already assigned.
			// Expect all CPUs from Socket 0.
			description: "GuPodMultipleCores, topoQuadSocketFourWayHT, ExpectAllocSock0",
			topo:        topoQuadSocketFourWayHT,
			containerID: "fakeID5",
			stAssignments: state.ContainerCPUAssignments{
				"fakeID100": cpuset.NewCPUSet(3, 11, 4, 5, 6, 7),
			},
			stDefaultCPUSet: largeTopoCPUSet.Difference(cpuset.NewCPUSet(3, 11, 4, 5, 6, 7)),
			pod:             makePod("72000m", "72000m"),
			expErr:          nil,
			expCPUAlloc:     true,
			expCSet:         largeTopoSock0CPUSet,
		},
		{
			// Only 2 full cores from three Sockets and some partial cores are available.
			// Expect CPUs from the 2 full cores available from the three Sockets.
			description: "GuPodMultipleCores, topoQuadSocketFourWayHT, ExpectAllocAllFullCoresFromThreeSockets",
			topo:        topoQuadSocketFourWayHT,
			containerID: "fakeID5",
			stAssignments: state.ContainerCPUAssignments{
				"fakeID100": largeTopoCPUSet.Difference(cpuset.NewCPUSet(1, 25, 13, 38, 2, 9, 11, 35, 23, 48, 12, 51,
					53, 173, 113, 233, 54, 61)),
			},
			stDefaultCPUSet: cpuset.NewCPUSet(1, 25, 13, 38, 2, 9, 11, 35, 23, 48, 12, 51, 53, 173, 113, 233, 54, 61),
			pod:             makePod("12000m", "12000m"),
			expErr:          nil,
			expCPUAlloc:     true,
			expCSet:         cpuset.NewCPUSet(1, 25, 13, 38, 11, 35, 23, 48, 53, 173, 113, 233),
		},
		{
			// All CPUs from Socket 1, 1 full core and some partial cores are available.
			// Expect all CPUs from Socket 1 and the hyper-threads from the full core.
			description: "GuPodMultipleCores, topoQuadSocketFourWayHT, ExpectAllocAllSock1+FullCore",
			topo:        topoQuadSocketFourWayHT,
			containerID: "fakeID5",
			stAssignments: state.ContainerCPUAssignments{
				"fakeID100": largeTopoCPUSet.Difference(largeTopoSock1CPUSet.Union(cpuset.NewCPUSet(10, 34, 22, 47, 53,
					173, 61, 181, 108, 228, 115, 235))),
			},
			stDefaultCPUSet: largeTopoSock1CPUSet.Union(cpuset.NewCPUSet(10, 34, 22, 47, 53, 173, 61, 181, 108, 228,
				115, 235)),
			pod:         makePod("76000m", "76000m"),
			expErr:      nil,
			expCPUAlloc: true,
			expCSet:     largeTopoSock1CPUSet.Union(cpuset.NewCPUSet(10, 34, 22, 47)),
		},
		{
			// Only partial cores are available in the entire system.
			// Expect allocation of all the CPUs from the partial cores.
			description: "GuPodMultipleCores, topoQuadSocketFourWayHT, ExpectAllocCPUs",
			topo:        topoQuadSocketFourWayHT,
			containerID: "fakeID5",
			stAssignments: state.ContainerCPUAssignments{
				"fakeID100": largeTopoCPUSet.Difference(cpuset.NewCPUSet(10, 11, 53, 37, 55, 67, 52)),
			},
			stDefaultCPUSet: cpuset.NewCPUSet(10, 11, 53, 67, 52),
			pod:             makePod("5000m", "5000m"),
			expErr:          nil,
			expCPUAlloc:     true,
			expCSet:         cpuset.NewCPUSet(10, 11, 53, 67, 52),
		},
		{
			// Only 7 CPUs are available.
			// Pod requests 76 cores.
			// Error is expect since available CPUs are less than the request.
			description: "GuPodMultipleCores, topoQuadSocketFourWayHT, NoAlloc",
			topo:        topoQuadSocketFourWayHT,
			containerID: "fakeID5",
			stAssignments: state.ContainerCPUAssignments{
				"fakeID100": largeTopoCPUSet.Difference(cpuset.NewCPUSet(10, 11, 53, 37, 55, 67, 52)),
			},
			stDefaultCPUSet: cpuset.NewCPUSet(10, 11, 53, 37, 55, 67, 52),
			pod:             makePod("76000m", "76000m"),
			expErr:          fmt.Errorf("not enough cpus available to satisfy request"),
			expCPUAlloc:     false,
			expCSet:         cpuset.NewCPUSet(),
		},
	}

	for _, testCase := range testCases {
		policy := NewStaticPolicy(testCase.topo, testCase.numReservedCPUs, topologymanager.NewFakeManager())

		st := &mockState{
			assignments:   testCase.stAssignments,
			defaultCPUSet: testCase.stDefaultCPUSet,
		}

		container := &testCase.pod.Spec.Containers[0]
		err := policy.AddContainer(st, testCase.pod, container, testCase.containerID)
		if !reflect.DeepEqual(err, testCase.expErr) {
			t.Errorf("StaticPolicy AddContainer() error (%v). expected add error: %v but got: %v",
				testCase.description, testCase.expErr, err)
		}

		if testCase.expCPUAlloc {
			cset, found := st.assignments[testCase.containerID]
			if !found {
				t.Errorf("StaticPolicy AddContainer() error (%v). expected container id %v to be present in assignments %v",
					testCase.description, testCase.containerID, st.assignments)
			}

			if !reflect.DeepEqual(cset, testCase.expCSet) {
				t.Errorf("StaticPolicy AddContainer() error (%v). expected cpuset %v but got %v",
					testCase.description, testCase.expCSet, cset)
			}

			if !cset.Intersection(st.defaultCPUSet).IsEmpty() {
				t.Errorf("StaticPolicy AddContainer() error (%v). expected cpuset %v to be disoint from the shared cpuset %v",
					testCase.description, cset, st.defaultCPUSet)
			}
		}

		if !testCase.expCPUAlloc {
			_, found := st.assignments[testCase.containerID]
			if found {
				t.Errorf("StaticPolicy AddContainer() error (%v). Did not expect container id %v to be present in assignments %v",
					testCase.description, testCase.containerID, st.assignments)
			}
		}
	}
}

func TestStaticPolicyAddWithInitContainers(t *testing.T) {
	testCases := []staticPolicyMultiContainerTest{
		{
			description:      "No Guaranteed Init CPUs",
			topo:             topoSingleSocketHT,
			numReservedCPUs:  0,
			stAssignments:    state.ContainerCPUAssignments{},
			stDefaultCPUSet:  cpuset.NewCPUSet(0, 1, 2, 3, 4, 5, 6, 7),
			initContainerIDs: []string{"initFakeID"},
			containerIDs:     []string{"appFakeID"},
			pod: makeMultiContainerPod(
				[]struct{ request, limit string }{{"100m", "100m"}},
				[]struct{ request, limit string }{{"4000m", "4000m"}}),
			expInitCSets: []cpuset.CPUSet{
				cpuset.NewCPUSet()},
			expCSets: []cpuset.CPUSet{
				cpuset.NewCPUSet(0, 4, 1, 5)},
		},
		{
			description:      "Equal Number of Guaranteed CPUs",
			topo:             topoSingleSocketHT,
			numReservedCPUs:  0,
			stAssignments:    state.ContainerCPUAssignments{},
			stDefaultCPUSet:  cpuset.NewCPUSet(0, 1, 2, 3, 4, 5, 6, 7),
			initContainerIDs: []string{"initFakeID"},
			containerIDs:     []string{"appFakeID"},
			pod: makeMultiContainerPod(
				[]struct{ request, limit string }{{"4000m", "4000m"}},
				[]struct{ request, limit string }{{"4000m", "4000m"}}),
			expInitCSets: []cpuset.CPUSet{
				cpuset.NewCPUSet(0, 4, 1, 5)},
			expCSets: []cpuset.CPUSet{
				cpuset.NewCPUSet(0, 4, 1, 5)},
		},
		{
			description:      "More Init Container Guaranteed CPUs",
			topo:             topoSingleSocketHT,
			numReservedCPUs:  0,
			stAssignments:    state.ContainerCPUAssignments{},
			stDefaultCPUSet:  cpuset.NewCPUSet(0, 1, 2, 3, 4, 5, 6, 7),
			initContainerIDs: []string{"initFakeID"},
			containerIDs:     []string{"appFakeID"},
			pod: makeMultiContainerPod(
				[]struct{ request, limit string }{{"6000m", "6000m"}},
				[]struct{ request, limit string }{{"4000m", "4000m"}}),
			expInitCSets: []cpuset.CPUSet{
				cpuset.NewCPUSet(0, 4, 1, 5, 2, 6)},
			expCSets: []cpuset.CPUSet{
				cpuset.NewCPUSet(0, 4, 1, 5)},
		},
		{
			description:      "Less Init Container Guaranteed CPUs",
			topo:             topoSingleSocketHT,
			numReservedCPUs:  0,
			stAssignments:    state.ContainerCPUAssignments{},
			stDefaultCPUSet:  cpuset.NewCPUSet(0, 1, 2, 3, 4, 5, 6, 7),
			initContainerIDs: []string{"initFakeID"},
			containerIDs:     []string{"appFakeID"},
			pod: makeMultiContainerPod(
				[]struct{ request, limit string }{{"2000m", "2000m"}},
				[]struct{ request, limit string }{{"4000m", "4000m"}}),
			expInitCSets: []cpuset.CPUSet{
				cpuset.NewCPUSet(0, 4)},
			expCSets: []cpuset.CPUSet{
				cpuset.NewCPUSet(0, 4, 1, 5)},
		},
		{
			description:      "Multi Init Container Equal CPUs",
			topo:             topoSingleSocketHT,
			numReservedCPUs:  0,
			stAssignments:    state.ContainerCPUAssignments{},
			stDefaultCPUSet:  cpuset.NewCPUSet(0, 1, 2, 3, 4, 5, 6, 7),
			initContainerIDs: []string{"initFakeID-1", "initFakeID-2"},
			containerIDs:     []string{"appFakeID"},
			pod: makeMultiContainerPod(
				[]struct{ request, limit string }{
					{"2000m", "2000m"},
					{"2000m", "2000m"}},
				[]struct{ request, limit string }{
					{"2000m", "2000m"}}),
			expInitCSets: []cpuset.CPUSet{
				cpuset.NewCPUSet(0, 4),
				cpuset.NewCPUSet(0, 4)},
			expCSets: []cpuset.CPUSet{
				cpuset.NewCPUSet(0, 4)},
		},
		{
			description:      "Multi Init Container Less CPUs",
			topo:             topoSingleSocketHT,
			numReservedCPUs:  0,
			stAssignments:    state.ContainerCPUAssignments{},
			stDefaultCPUSet:  cpuset.NewCPUSet(0, 1, 2, 3, 4, 5, 6, 7),
			initContainerIDs: []string{"initFakeID-1", "initFakeID-2"},
			containerIDs:     []string{"appFakeID"},
			pod: makeMultiContainerPod(
				[]struct{ request, limit string }{
					{"4000m", "4000m"},
					{"4000m", "4000m"}},
				[]struct{ request, limit string }{
					{"2000m", "2000m"}}),
			expInitCSets: []cpuset.CPUSet{
				cpuset.NewCPUSet(0, 4, 1, 5),
				cpuset.NewCPUSet(0, 4, 1, 5)},
			expCSets: []cpuset.CPUSet{
				cpuset.NewCPUSet(0, 4)},
		},
		{
			description:      "Multi Init Container More CPUs",
			topo:             topoSingleSocketHT,
			numReservedCPUs:  0,
			stAssignments:    state.ContainerCPUAssignments{},
			stDefaultCPUSet:  cpuset.NewCPUSet(0, 1, 2, 3, 4, 5, 6, 7),
			initContainerIDs: []string{"initFakeID-1", "initFakeID-2"},
			containerIDs:     []string{"appFakeID"},
			pod: makeMultiContainerPod(
				[]struct{ request, limit string }{
					{"2000m", "2000m"},
					{"2000m", "2000m"}},
				[]struct{ request, limit string }{
					{"4000m", "4000m"}}),
			expInitCSets: []cpuset.CPUSet{
				cpuset.NewCPUSet(0, 4),
				cpuset.NewCPUSet(0, 4)},
			expCSets: []cpuset.CPUSet{
				cpuset.NewCPUSet(0, 4, 1, 5)},
		},
		{
			description:      "Multi Init Container Increasing CPUs",
			topo:             topoSingleSocketHT,
			numReservedCPUs:  0,
			stAssignments:    state.ContainerCPUAssignments{},
			stDefaultCPUSet:  cpuset.NewCPUSet(0, 1, 2, 3, 4, 5, 6, 7),
			initContainerIDs: []string{"initFakeID-1", "initFakeID-2"},
			containerIDs:     []string{"appFakeID"},
			pod: makeMultiContainerPod(
				[]struct{ request, limit string }{
					{"2000m", "2000m"},
					{"4000m", "4000m"}},
				[]struct{ request, limit string }{
					{"6000m", "6000m"}}),
			expInitCSets: []cpuset.CPUSet{
				cpuset.NewCPUSet(0, 4),
				cpuset.NewCPUSet(0, 4, 1, 5)},
			expCSets: []cpuset.CPUSet{
				cpuset.NewCPUSet(0, 4, 1, 5, 2, 6)},
		},
		{
			description:      "Multi Init, Multi App Container Split CPUs",
			topo:             topoSingleSocketHT,
			numReservedCPUs:  0,
			stAssignments:    state.ContainerCPUAssignments{},
			stDefaultCPUSet:  cpuset.NewCPUSet(0, 1, 2, 3, 4, 5, 6, 7),
			initContainerIDs: []string{"initFakeID-1", "initFakeID-2"},
			containerIDs:     []string{"appFakeID-1", "appFakeID-2"},
			pod: makeMultiContainerPod(
				[]struct{ request, limit string }{
					{"2000m", "2000m"},
					{"4000m", "4000m"}},
				[]struct{ request, limit string }{
					{"2000m", "2000m"},
					{"2000m", "2000m"}}),
			expInitCSets: []cpuset.CPUSet{
				cpuset.NewCPUSet(0, 4),
				cpuset.NewCPUSet(0, 4, 1, 5)},
			expCSets: []cpuset.CPUSet{
				cpuset.NewCPUSet(0, 4),
				cpuset.NewCPUSet(1, 5)},
		},
	}

	for _, testCase := range testCases {
		policy := NewStaticPolicy(testCase.topo, testCase.numReservedCPUs, topologymanager.NewFakeManager())

		st := &mockState{
			assignments:   testCase.stAssignments,
			defaultCPUSet: testCase.stDefaultCPUSet,
		}

		containers := append(
			testCase.pod.Spec.InitContainers,
			testCase.pod.Spec.Containers...)

		containerIDs := append(
			testCase.initContainerIDs,
			testCase.containerIDs...)

		expCSets := append(
			testCase.expInitCSets,
			testCase.expCSets...)

		for i := range containers {
			err := policy.AddContainer(st, testCase.pod, &containers[i], containerIDs[i])
			if err != nil {
				t.Errorf("StaticPolicy AddContainer() error (%v). unexpected error for container id: %v: %v",
					testCase.description, containerIDs[i], err)
			}

			cset, found := st.assignments[containerIDs[i]]
			if !expCSets[i].IsEmpty() && !found {
				t.Errorf("StaticPolicy AddContainer() error (%v). expected container id %v to be present in assignments %v",
					testCase.description, containerIDs[i], st.assignments)
			}

			if found && !cset.Equals(expCSets[i]) {
				t.Errorf("StaticPolicy AddContainer() error (%v). expected cpuset %v for container %v but got %v",
					testCase.description, expCSets[i], containerIDs[i], cset)
			}
		}
	}
}

func TestStaticPolicyRemove(t *testing.T) {
	testCases := []staticPolicyTest{
		{
			description: "SingleSocketHT, DeAllocOneContainer",
			topo:        topoSingleSocketHT,
			containerID: "fakeID1",
			stAssignments: state.ContainerCPUAssignments{
				"fakeID1": cpuset.NewCPUSet(1, 2, 3),
			},
			stDefaultCPUSet: cpuset.NewCPUSet(4, 5, 6, 7),
			expCSet:         cpuset.NewCPUSet(1, 2, 3, 4, 5, 6, 7),
		},
		{
			description: "SingleSocketHT, DeAllocOneContainer, BeginEmpty",
			topo:        topoSingleSocketHT,
			containerID: "fakeID1",
			stAssignments: state.ContainerCPUAssignments{
				"fakeID1": cpuset.NewCPUSet(1, 2, 3),
				"fakeID2": cpuset.NewCPUSet(4, 5, 6, 7),
			},
			stDefaultCPUSet: cpuset.NewCPUSet(),
			expCSet:         cpuset.NewCPUSet(1, 2, 3),
		},
		{
			description: "SingleSocketHT, DeAllocTwoContainer",
			topo:        topoSingleSocketHT,
			containerID: "fakeID1",
			stAssignments: state.ContainerCPUAssignments{
				"fakeID1": cpuset.NewCPUSet(1, 3, 5),
				"fakeID2": cpuset.NewCPUSet(2, 4),
			},
			stDefaultCPUSet: cpuset.NewCPUSet(6, 7),
			expCSet:         cpuset.NewCPUSet(1, 3, 5, 6, 7),
		},
		{
			description: "SingleSocketHT, NoDeAlloc",
			topo:        topoSingleSocketHT,
			containerID: "fakeID2",
			stAssignments: state.ContainerCPUAssignments{
				"fakeID1": cpuset.NewCPUSet(1, 3, 5),
			},
			stDefaultCPUSet: cpuset.NewCPUSet(2, 4, 6, 7),
			expCSet:         cpuset.NewCPUSet(2, 4, 6, 7),
		},
	}

	for _, testCase := range testCases {
		policy := NewStaticPolicy(testCase.topo, testCase.numReservedCPUs, topologymanager.NewFakeManager())

		st := &mockState{
			assignments:   testCase.stAssignments,
			defaultCPUSet: testCase.stDefaultCPUSet,
		}

		policy.RemoveContainer(st, testCase.containerID)

		if !reflect.DeepEqual(st.defaultCPUSet, testCase.expCSet) {
			t.Errorf("StaticPolicy RemoveContainer() error (%v). expected default cpuset %v but got %v",
				testCase.description, testCase.expCSet, st.defaultCPUSet)
		}

		if _, found := st.assignments[testCase.containerID]; found {
			t.Errorf("StaticPolicy RemoveContainer() error (%v). expected containerID %v not be in assignments %v",
				testCase.description, testCase.containerID, st.assignments)
		}
	}
}

func TestTopologyAwareAllocateCPUs(t *testing.T) {
	testCases := []struct {
		description     string
		topo            *topology.CPUTopology
		stAssignments   state.ContainerCPUAssignments
		stDefaultCPUSet cpuset.CPUSet
		numRequested    int
		socketMask      bitmask.BitMask
		expCSet         cpuset.CPUSet
	}{
		{
			description:     "Request 2 CPUs, No BitMask",
			topo:            topoDualSocketHT,
			stAssignments:   state.ContainerCPUAssignments{},
			stDefaultCPUSet: cpuset.NewCPUSet(0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11),
			numRequested:    2,
			socketMask:      nil,
			expCSet:         cpuset.NewCPUSet(0, 6),
		},
		{
			description:     "Request 2 CPUs, BitMask on Socket 0",
			topo:            topoDualSocketHT,
			stAssignments:   state.ContainerCPUAssignments{},
			stDefaultCPUSet: cpuset.NewCPUSet(0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11),
			numRequested:    2,
			socketMask: func() bitmask.BitMask {
				mask, _ := bitmask.NewBitMask(0)
				return mask
			}(),
			expCSet: cpuset.NewCPUSet(0, 6),
		},
		{
			description:     "Request 2 CPUs, BitMask on Socket 1",
			topo:            topoDualSocketHT,
			stAssignments:   state.ContainerCPUAssignments{},
			stDefaultCPUSet: cpuset.NewCPUSet(0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11),
			numRequested:    2,
			socketMask: func() bitmask.BitMask {
				mask, _ := bitmask.NewBitMask(1)
				return mask
			}(),
			expCSet: cpuset.NewCPUSet(1, 7),
		},
		{
			description:     "Request 8 CPUs, BitMask on Socket 0",
			topo:            topoDualSocketHT,
			stAssignments:   state.ContainerCPUAssignments{},
			stDefaultCPUSet: cpuset.NewCPUSet(0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11),
			numRequested:    8,
			socketMask: func() bitmask.BitMask {
				mask, _ := bitmask.NewBitMask(0)
				return mask
			}(),
			expCSet: cpuset.NewCPUSet(0, 6, 2, 8, 4, 10, 1, 7),
		},
		{
			description:     "Request 8 CPUs, BitMask on Socket 1",
			topo:            topoDualSocketHT,
			stAssignments:   state.ContainerCPUAssignments{},
			stDefaultCPUSet: cpuset.NewCPUSet(0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11),
			numRequested:    8,
			socketMask: func() bitmask.BitMask {
				mask, _ := bitmask.NewBitMask(1)
				return mask
			}(),
			expCSet: cpuset.NewCPUSet(1, 7, 3, 9, 5, 11, 0, 6),
		},
	}
	for _, tc := range testCases {
		policy := NewStaticPolicy(tc.topo, 0, topologymanager.NewFakeManager()).(*staticPolicy)
		st := &mockState{
			assignments:   tc.stAssignments,
			defaultCPUSet: tc.stDefaultCPUSet,
		}
		policy.Start(st)

		cset, err := policy.allocateCPUs(st, tc.numRequested, tc.socketMask)
		if err != nil {
			t.Errorf("StaticPolicy allocateCPUs() error (%v). expected CPUSet %v not error %v",
				tc.description, tc.expCSet, err)
			continue
		}

		if !reflect.DeepEqual(tc.expCSet, cset) {
			t.Errorf("StaticPolicy allocateCPUs() error (%v). expected CPUSet %v but got %v",
				tc.description, tc.expCSet, cset)
		}
	}
}
