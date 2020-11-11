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

package devicemanager

import (
	"fmt"
	"reflect"
	"sort"
	"testing"

	v1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/resource"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/apimachinery/pkg/util/sets"
	pluginapi "k8s.io/kubelet/pkg/apis/deviceplugin/v1beta1"
	"k8s.io/kubernetes/pkg/kubelet/cm/topologymanager"
	"k8s.io/kubernetes/pkg/kubelet/cm/topologymanager/bitmask"
)

type mockAffinityStore struct {
	hint topologymanager.TopologyHint
}

func (m *mockAffinityStore) GetAffinity(podUID string, containerName string) topologymanager.TopologyHint {
	return m.hint
}

func makeNUMADevice(id string, numa int) pluginapi.Device {
	return pluginapi.Device{
		ID:       id,
		Topology: &pluginapi.TopologyInfo{Nodes: []*pluginapi.NUMANode{{ID: int64(numa)}}},
	}
}

func makeSocketMask(sockets ...int) bitmask.BitMask {
	mask, _ := bitmask.NewBitMask(sockets...)
	return mask
}

func TestGetTopologyHints(t *testing.T) {
	tcases := []struct {
		description      string
		podUID           string
		containerName    string
		request          map[string]string
		devices          map[string][]pluginapi.Device
		allocatedDevices map[string]map[string]map[string][]string
		expectedHints    map[string][]topologymanager.TopologyHint
	}{
		{
			description:   "Single Request, no alignment",
			podUID:        "fakePod",
			containerName: "fakeContainer",
			request: map[string]string{
				"testdevice": "1",
			},
			devices: map[string][]pluginapi.Device{
				"testdevice": {
					{ID: "Dev1"},
					{ID: "Dev2"},
				},
			},
			expectedHints: map[string][]topologymanager.TopologyHint{
				"testdevice": nil,
			},
		},
		{
			description:   "Single Request, only one with alignment",
			podUID:        "fakePod",
			containerName: "fakeContainer",
			request: map[string]string{
				"testdevice": "1",
			},
			devices: map[string][]pluginapi.Device{
				"testdevice": {
					{ID: "Dev1"},
					makeNUMADevice("Dev2", 1),
				},
			},
			expectedHints: map[string][]topologymanager.TopologyHint{
				"testdevice": {
					{
						NUMANodeAffinity: makeSocketMask(1),
						Preferred:        true,
					},
					{
						NUMANodeAffinity: makeSocketMask(0, 1),
						Preferred:        false,
					},
				},
			},
		},
		{
			description:   "Single Request, one device per socket",
			podUID:        "fakePod",
			containerName: "fakeContainer",
			request: map[string]string{
				"testdevice": "1",
			},
			devices: map[string][]pluginapi.Device{
				"testdevice": {
					makeNUMADevice("Dev1", 0),
					makeNUMADevice("Dev2", 1),
				},
			},
			expectedHints: map[string][]topologymanager.TopologyHint{
				"testdevice": {
					{
						NUMANodeAffinity: makeSocketMask(0),
						Preferred:        true,
					},
					{
						NUMANodeAffinity: makeSocketMask(1),
						Preferred:        true,
					},
					{
						NUMANodeAffinity: makeSocketMask(0, 1),
						Preferred:        false,
					},
				},
			},
		},
		{
			description:   "Request for 2, one device per socket",
			podUID:        "fakePod",
			containerName: "fakeContainer",
			request: map[string]string{
				"testdevice": "2",
			},
			devices: map[string][]pluginapi.Device{
				"testdevice": {
					makeNUMADevice("Dev1", 0),
					makeNUMADevice("Dev2", 1),
				},
			},
			expectedHints: map[string][]topologymanager.TopologyHint{
				"testdevice": {
					{
						NUMANodeAffinity: makeSocketMask(0, 1),
						Preferred:        true,
					},
				},
			},
		},
		{
			description:   "Request for 2, 2 devices per socket",
			podUID:        "fakePod",
			containerName: "fakeContainer",
			request: map[string]string{
				"testdevice": "2",
			},
			devices: map[string][]pluginapi.Device{
				"testdevice": {
					makeNUMADevice("Dev1", 0),
					makeNUMADevice("Dev2", 1),
					makeNUMADevice("Dev3", 0),
					makeNUMADevice("Dev4", 1),
				},
			},
			expectedHints: map[string][]topologymanager.TopologyHint{
				"testdevice": {
					{
						NUMANodeAffinity: makeSocketMask(0),
						Preferred:        true,
					},
					{
						NUMANodeAffinity: makeSocketMask(1),
						Preferred:        true,
					},
					{
						NUMANodeAffinity: makeSocketMask(0, 1),
						Preferred:        false,
					},
				},
			},
		},
		{
			description:   "Request for 2, optimal on 1 NUMA node, forced cross-NUMA",
			podUID:        "fakePod",
			containerName: "fakeContainer",
			request: map[string]string{
				"testdevice": "2",
			},
			devices: map[string][]pluginapi.Device{
				"testdevice": {
					makeNUMADevice("Dev1", 0),
					makeNUMADevice("Dev2", 1),
					makeNUMADevice("Dev3", 0),
					makeNUMADevice("Dev4", 1),
				},
			},
			allocatedDevices: map[string]map[string]map[string][]string{
				"fakePod": {
					"fakeOtherContainer": {
						"testdevice": {"Dev1", "Dev2"},
					},
				},
			},
			expectedHints: map[string][]topologymanager.TopologyHint{
				"testdevice": {
					{
						NUMANodeAffinity: makeSocketMask(0, 1),
						Preferred:        false,
					},
				},
			},
		},
		{
			description:   "2 device types, mixed configuration",
			podUID:        "fakePod",
			containerName: "fakeContainer",
			request: map[string]string{
				"testdevice1": "2",
				"testdevice2": "1",
			},
			devices: map[string][]pluginapi.Device{
				"testdevice1": {
					makeNUMADevice("Dev1", 0),
					makeNUMADevice("Dev2", 1),
					makeNUMADevice("Dev3", 0),
					makeNUMADevice("Dev4", 1),
				},
				"testdevice2": {
					makeNUMADevice("Dev1", 0),
				},
			},
			expectedHints: map[string][]topologymanager.TopologyHint{
				"testdevice1": {
					{
						NUMANodeAffinity: makeSocketMask(0),
						Preferred:        true,
					},
					{
						NUMANodeAffinity: makeSocketMask(1),
						Preferred:        true,
					},
					{
						NUMANodeAffinity: makeSocketMask(0, 1),
						Preferred:        false,
					},
				},
				"testdevice2": {
					{
						NUMANodeAffinity: makeSocketMask(0),
						Preferred:        true,
					},
					{
						NUMANodeAffinity: makeSocketMask(0, 1),
						Preferred:        false,
					},
				},
			},
		},
		{
			description:   "Single device type, more requested than available",
			podUID:        "fakePod",
			containerName: "fakeContainer",
			request: map[string]string{
				"testdevice": "6",
			},
			devices: map[string][]pluginapi.Device{
				"testdevice": {
					makeNUMADevice("Dev1", 0),
					makeNUMADevice("Dev2", 0),
					makeNUMADevice("Dev3", 1),
					makeNUMADevice("Dev4", 1),
				},
			},
			expectedHints: map[string][]topologymanager.TopologyHint{
				"testdevice": {},
			},
		},
		{
			description:   "Single device type, all already allocated to container",
			podUID:        "fakePod",
			containerName: "fakeContainer",
			request: map[string]string{
				"testdevice": "2",
			},
			devices: map[string][]pluginapi.Device{
				"testdevice": {
					makeNUMADevice("Dev1", 0),
					makeNUMADevice("Dev2", 0),
				},
			},
			allocatedDevices: map[string]map[string]map[string][]string{
				"fakePod": {
					"fakeContainer": {
						"testdevice": {"Dev1", "Dev2"},
					},
				},
			},
			expectedHints: map[string][]topologymanager.TopologyHint{
				"testdevice": {
					{
						NUMANodeAffinity: makeSocketMask(0),
						Preferred:        true,
					},
					{
						NUMANodeAffinity: makeSocketMask(0, 1),
						Preferred:        false,
					},
				},
			},
		},
		{
			description:   "Single device type, less already allocated to container than requested",
			podUID:        "fakePod",
			containerName: "fakeContainer",
			request: map[string]string{
				"testdevice": "4",
			},
			devices: map[string][]pluginapi.Device{
				"testdevice": {
					makeNUMADevice("Dev1", 0),
					makeNUMADevice("Dev2", 0),
					makeNUMADevice("Dev3", 1),
					makeNUMADevice("Dev4", 1),
				},
			},
			allocatedDevices: map[string]map[string]map[string][]string{
				"fakePod": {
					"fakeContainer": {
						"testdevice": {"Dev1", "Dev2"},
					},
				},
			},
			expectedHints: map[string][]topologymanager.TopologyHint{
				"testdevice": {},
			},
		},
		{
			description:   "Single device type, more already allocated to container than requested",
			podUID:        "fakePod",
			containerName: "fakeContainer",
			request: map[string]string{
				"testdevice": "2",
			},
			devices: map[string][]pluginapi.Device{
				"testdevice": {
					makeNUMADevice("Dev1", 0),
					makeNUMADevice("Dev2", 0),
					makeNUMADevice("Dev3", 1),
					makeNUMADevice("Dev4", 1),
				},
			},
			allocatedDevices: map[string]map[string]map[string][]string{
				"fakePod": {
					"fakeContainer": {
						"testdevice": {"Dev1", "Dev2", "Dev3", "Dev4"},
					},
				},
			},
			expectedHints: map[string][]topologymanager.TopologyHint{
				"testdevice": {},
			},
		},
	}

	for _, tc := range tcases {
		resourceList := v1.ResourceList{}
		for r := range tc.request {
			resourceList[v1.ResourceName(r)] = resource.MustParse(tc.request[r])
		}

		pod := makePod(resourceList)
		pod.UID = types.UID(tc.podUID)
		pod.Spec.Containers[0].Name = tc.containerName

		m := ManagerImpl{
			allDevices:       make(map[string]map[string]pluginapi.Device),
			healthyDevices:   make(map[string]sets.String),
			allocatedDevices: make(map[string]sets.String),
			podDevices:       newPodDevices(),
			sourcesReady:     &sourcesReadyStub{},
			activePods:       func() []*v1.Pod { return []*v1.Pod{pod} },
			numaNodes:        []int{0, 1},
		}

		for r := range tc.devices {
			m.allDevices[r] = make(map[string]pluginapi.Device)
			m.healthyDevices[r] = sets.NewString()

			for _, d := range tc.devices[r] {
				m.allDevices[r][d.ID] = d
				m.healthyDevices[r].Insert(d.ID)
			}
		}

		for p := range tc.allocatedDevices {
			for c := range tc.allocatedDevices[p] {
				for r, devices := range tc.allocatedDevices[p][c] {
					m.podDevices.insert(p, c, r, constructDevices(devices), nil)

					m.allocatedDevices[r] = sets.NewString()
					for _, d := range devices {
						m.allocatedDevices[r].Insert(d)
					}
				}
			}
		}

		hints := m.GetTopologyHints(pod, &pod.Spec.Containers[0])

		for r := range tc.expectedHints {
			sort.SliceStable(hints[r], func(i, j int) bool {
				return hints[r][i].LessThan(hints[r][j])
			})
			sort.SliceStable(tc.expectedHints[r], func(i, j int) bool {
				return tc.expectedHints[r][i].LessThan(tc.expectedHints[r][j])
			})
			if !reflect.DeepEqual(hints[r], tc.expectedHints[r]) {
				t.Errorf("%v: Expected result to be %v, got %v", tc.description, tc.expectedHints[r], hints[r])
			}
		}
	}
}

func TestTopologyAlignedAllocation(t *testing.T) {
	tcases := []struct {
		description                 string
		resource                    string
		request                     int
		devices                     []pluginapi.Device
		allocatedDevices            []string
		hint                        topologymanager.TopologyHint
		getPreferredAllocationFunc  func(available, mustInclude []string, size int) (*pluginapi.PreferredAllocationResponse, error)
		expectedPreferredAllocation []string
		expectedAlignment           map[int]int
	}{
		{
			description: "Single Request, no alignment",
			resource:    "resource",
			request:     1,
			devices: []pluginapi.Device{
				{ID: "Dev1"},
				{ID: "Dev2"},
			},
			hint: topologymanager.TopologyHint{
				NUMANodeAffinity: makeSocketMask(0, 1),
				Preferred:        true,
			},
			expectedAlignment: map[int]int{},
		},
		{
			description: "Request for 1, partial alignment",
			resource:    "resource",
			request:     1,
			devices: []pluginapi.Device{
				{ID: "Dev1"},
				makeNUMADevice("Dev2", 1),
			},
			hint: topologymanager.TopologyHint{
				NUMANodeAffinity: makeSocketMask(1),
				Preferred:        true,
			},
			expectedAlignment: map[int]int{1: 1},
		},
		{
			description: "Single Request, socket 0",
			resource:    "resource",
			request:     1,
			devices: []pluginapi.Device{
				makeNUMADevice("Dev1", 0),
				makeNUMADevice("Dev2", 1),
			},
			hint: topologymanager.TopologyHint{
				NUMANodeAffinity: makeSocketMask(0),
				Preferred:        true,
			},
			expectedAlignment: map[int]int{0: 1},
		},
		{
			description: "Single Request, socket 1",
			resource:    "resource",
			request:     1,
			devices: []pluginapi.Device{
				makeNUMADevice("Dev1", 0),
				makeNUMADevice("Dev2", 1),
			},
			hint: topologymanager.TopologyHint{
				NUMANodeAffinity: makeSocketMask(1),
				Preferred:        true,
			},
			expectedAlignment: map[int]int{1: 1},
		},
		{
			description: "Request for 2, socket 0",
			resource:    "resource",
			request:     2,
			devices: []pluginapi.Device{
				makeNUMADevice("Dev1", 0),
				makeNUMADevice("Dev2", 1),
				makeNUMADevice("Dev3", 0),
				makeNUMADevice("Dev4", 1),
			},
			hint: topologymanager.TopologyHint{
				NUMANodeAffinity: makeSocketMask(0),
				Preferred:        true,
			},
			expectedAlignment: map[int]int{0: 2},
		},
		{
			description: "Request for 2, socket 1",
			resource:    "resource",
			request:     2,
			devices: []pluginapi.Device{
				makeNUMADevice("Dev1", 0),
				makeNUMADevice("Dev2", 1),
				makeNUMADevice("Dev3", 0),
				makeNUMADevice("Dev4", 1),
			},
			hint: topologymanager.TopologyHint{
				NUMANodeAffinity: makeSocketMask(1),
				Preferred:        true,
			},
			expectedAlignment: map[int]int{1: 2},
		},
		{
			description: "Request for 4, unsatisfiable, prefer socket 0",
			resource:    "resource",
			request:     4,
			devices: []pluginapi.Device{
				makeNUMADevice("Dev1", 0),
				makeNUMADevice("Dev2", 1),
				makeNUMADevice("Dev3", 0),
				makeNUMADevice("Dev4", 1),
				makeNUMADevice("Dev5", 0),
				makeNUMADevice("Dev6", 1),
			},
			hint: topologymanager.TopologyHint{
				NUMANodeAffinity: makeSocketMask(0),
				Preferred:        true,
			},
			expectedAlignment: map[int]int{0: 3, 1: 1},
		},
		{
			description: "Request for 4, unsatisfiable, prefer socket 1",
			resource:    "resource",
			request:     4,
			devices: []pluginapi.Device{
				makeNUMADevice("Dev1", 0),
				makeNUMADevice("Dev2", 1),
				makeNUMADevice("Dev3", 0),
				makeNUMADevice("Dev4", 1),
				makeNUMADevice("Dev5", 0),
				makeNUMADevice("Dev6", 1),
			},
			hint: topologymanager.TopologyHint{
				NUMANodeAffinity: makeSocketMask(1),
				Preferred:        true,
			},
			expectedAlignment: map[int]int{0: 1, 1: 3},
		},
		{
			description: "Request for 4, multisocket",
			resource:    "resource",
			request:     4,
			devices: []pluginapi.Device{
				makeNUMADevice("Dev1", 0),
				makeNUMADevice("Dev2", 1),
				makeNUMADevice("Dev3", 2),
				makeNUMADevice("Dev4", 3),
				makeNUMADevice("Dev5", 0),
				makeNUMADevice("Dev6", 1),
				makeNUMADevice("Dev7", 2),
				makeNUMADevice("Dev8", 3),
			},
			hint: topologymanager.TopologyHint{
				NUMANodeAffinity: makeSocketMask(1, 3),
				Preferred:        true,
			},
			expectedAlignment: map[int]int{1: 2, 3: 2},
		},
		{
			description: "Request for 5, socket 0, preferred aligned accepted",
			resource:    "resource",
			request:     5,
			devices: func() []pluginapi.Device {
				devices := []pluginapi.Device{}
				for i := 0; i < 100; i++ {
					id := fmt.Sprintf("Dev%d", i)
					devices = append(devices, makeNUMADevice(id, 0))
				}
				for i := 100; i < 200; i++ {
					id := fmt.Sprintf("Dev%d", i)
					devices = append(devices, makeNUMADevice(id, 1))
				}
				return devices
			}(),
			hint: topologymanager.TopologyHint{
				NUMANodeAffinity: makeSocketMask(0),
				Preferred:        true,
			},
			getPreferredAllocationFunc: func(available, mustInclude []string, size int) (*pluginapi.PreferredAllocationResponse, error) {
				return &pluginapi.PreferredAllocationResponse{
					ContainerResponses: []*pluginapi.ContainerPreferredAllocationResponse{
						{DeviceIDs: []string{"Dev0", "Dev19", "Dev83", "Dev42", "Dev77"}},
					},
				}, nil
			},
			expectedPreferredAllocation: []string{"Dev0", "Dev19", "Dev83", "Dev42", "Dev77"},
			expectedAlignment:           map[int]int{0: 5},
		},
		{
			description: "Request for 5, socket 0, preferred aligned accepted, unaligned ignored",
			resource:    "resource",
			request:     5,
			devices: func() []pluginapi.Device {
				devices := []pluginapi.Device{}
				for i := 0; i < 100; i++ {
					id := fmt.Sprintf("Dev%d", i)
					devices = append(devices, makeNUMADevice(id, 0))
				}
				for i := 100; i < 200; i++ {
					id := fmt.Sprintf("Dev%d", i)
					devices = append(devices, makeNUMADevice(id, 1))
				}
				return devices
			}(),
			hint: topologymanager.TopologyHint{
				NUMANodeAffinity: makeSocketMask(0),
				Preferred:        true,
			},
			getPreferredAllocationFunc: func(available, mustInclude []string, size int) (*pluginapi.PreferredAllocationResponse, error) {
				return &pluginapi.PreferredAllocationResponse{
					ContainerResponses: []*pluginapi.ContainerPreferredAllocationResponse{
						{DeviceIDs: []string{"Dev0", "Dev19", "Dev83", "Dev150", "Dev186"}},
					},
				}, nil
			},
			expectedPreferredAllocation: []string{"Dev0", "Dev19", "Dev83"},
			expectedAlignment:           map[int]int{0: 5},
		},
		{
			description: "Request for 5, socket 1, preferred aligned accepted, bogus ignored",
			resource:    "resource",
			request:     5,
			devices: func() []pluginapi.Device {
				devices := []pluginapi.Device{}
				for i := 0; i < 100; i++ {
					id := fmt.Sprintf("Dev%d", i)
					devices = append(devices, makeNUMADevice(id, 1))
				}
				return devices
			}(),
			hint: topologymanager.TopologyHint{
				NUMANodeAffinity: makeSocketMask(1),
				Preferred:        true,
			},
			getPreferredAllocationFunc: func(available, mustInclude []string, size int) (*pluginapi.PreferredAllocationResponse, error) {
				return &pluginapi.PreferredAllocationResponse{
					ContainerResponses: []*pluginapi.ContainerPreferredAllocationResponse{
						{DeviceIDs: []string{"Dev0", "Dev19", "Dev83", "bogus0", "bogus1"}},
					},
				}, nil
			},
			expectedPreferredAllocation: []string{"Dev0", "Dev19", "Dev83"},
			expectedAlignment:           map[int]int{1: 5},
		},
		{
			description: "Request for 5, multisocket, preferred accepted",
			resource:    "resource",
			request:     5,
			devices: func() []pluginapi.Device {
				devices := []pluginapi.Device{}
				for i := 0; i < 3; i++ {
					id := fmt.Sprintf("Dev%d", i)
					devices = append(devices, makeNUMADevice(id, 0))
				}
				for i := 3; i < 100; i++ {
					id := fmt.Sprintf("Dev%d", i)
					devices = append(devices, makeNUMADevice(id, 1))
				}
				return devices
			}(),
			hint: topologymanager.TopologyHint{
				NUMANodeAffinity: makeSocketMask(0),
				Preferred:        true,
			},
			getPreferredAllocationFunc: func(available, mustInclude []string, size int) (*pluginapi.PreferredAllocationResponse, error) {
				return &pluginapi.PreferredAllocationResponse{
					ContainerResponses: []*pluginapi.ContainerPreferredAllocationResponse{
						{DeviceIDs: []string{"Dev0", "Dev1", "Dev2", "Dev42", "Dev83"}},
					},
				}, nil
			},
			expectedPreferredAllocation: []string{"Dev0", "Dev1", "Dev2", "Dev42", "Dev83"},
			expectedAlignment:           map[int]int{0: 3, 1: 2},
		},
		{
			description: "Request for 5, multisocket, preferred unaligned accepted, bogus ignored",
			resource:    "resource",
			request:     5,
			devices: func() []pluginapi.Device {
				devices := []pluginapi.Device{}
				for i := 0; i < 3; i++ {
					id := fmt.Sprintf("Dev%d", i)
					devices = append(devices, makeNUMADevice(id, 0))
				}
				for i := 3; i < 100; i++ {
					id := fmt.Sprintf("Dev%d", i)
					devices = append(devices, makeNUMADevice(id, 1))
				}
				return devices
			}(),
			hint: topologymanager.TopologyHint{
				NUMANodeAffinity: makeSocketMask(0),
				Preferred:        true,
			},
			getPreferredAllocationFunc: func(available, mustInclude []string, size int) (*pluginapi.PreferredAllocationResponse, error) {
				return &pluginapi.PreferredAllocationResponse{
					ContainerResponses: []*pluginapi.ContainerPreferredAllocationResponse{
						{DeviceIDs: []string{"Dev0", "Dev1", "Dev2", "Dev42", "bogus0"}},
					},
				}, nil
			},
			expectedPreferredAllocation: []string{"Dev0", "Dev1", "Dev2", "Dev42"},
			expectedAlignment:           map[int]int{0: 3, 1: 2},
		},
	}
	for _, tc := range tcases {
		m := ManagerImpl{
			allDevices:            make(map[string]map[string]pluginapi.Device),
			healthyDevices:        make(map[string]sets.String),
			allocatedDevices:      make(map[string]sets.String),
			endpoints:             make(map[string]endpointInfo),
			podDevices:            newPodDevices(),
			sourcesReady:          &sourcesReadyStub{},
			activePods:            func() []*v1.Pod { return []*v1.Pod{} },
			topologyAffinityStore: &mockAffinityStore{tc.hint},
		}

		m.allDevices[tc.resource] = make(map[string]pluginapi.Device)
		m.healthyDevices[tc.resource] = sets.NewString()
		m.endpoints[tc.resource] = endpointInfo{}

		for _, d := range tc.devices {
			m.allDevices[tc.resource][d.ID] = d
			m.healthyDevices[tc.resource].Insert(d.ID)
		}

		if tc.getPreferredAllocationFunc != nil {
			m.endpoints[tc.resource] = endpointInfo{
				e: &MockEndpoint{
					getPreferredAllocationFunc: tc.getPreferredAllocationFunc,
				},
				opts: &pluginapi.DevicePluginOptions{GetPreferredAllocationAvailable: true},
			}
		}

		allocated, err := m.devicesToAllocate("podUID", "containerName", tc.resource, tc.request, sets.NewString())
		if err != nil {
			t.Errorf("Unexpected error: %v", err)
			continue
		}

		if len(allocated) != tc.request {
			t.Errorf("%v. expected allocation size: %v but got: %v", tc.description, tc.request, len(allocated))
		}

		if !allocated.HasAll(tc.expectedPreferredAllocation...) {
			t.Errorf("%v. expected preferred allocation: %v but not present in: %v", tc.description, tc.expectedPreferredAllocation, allocated.UnsortedList())
		}

		alignment := make(map[int]int)
		if m.deviceHasTopologyAlignment(tc.resource) {
			for d := range allocated {
				if m.allDevices[tc.resource][d].Topology != nil {
					alignment[int(m.allDevices[tc.resource][d].Topology.Nodes[0].ID)]++
				}
			}
		}

		if !reflect.DeepEqual(alignment, tc.expectedAlignment) {
			t.Errorf("%v. expected alignment: %v but got: %v", tc.description, tc.expectedAlignment, alignment)
		}
	}
}

func TestGetPreferredAllocationParameters(t *testing.T) {
	tcases := []struct {
		description         string
		resource            string
		request             int
		allDevices          []pluginapi.Device
		allocatedDevices    []string
		reusableDevices     []string
		hint                topologymanager.TopologyHint
		expectedAvailable   []string
		expectedMustInclude []string
		expectedSize        int
	}{
		{
			description: "Request for 1, socket 0, 0 already allocated, 0 reusable",
			resource:    "resource",
			request:     1,
			allDevices: []pluginapi.Device{
				makeNUMADevice("Dev0", 0),
				makeNUMADevice("Dev1", 0),
				makeNUMADevice("Dev2", 0),
				makeNUMADevice("Dev3", 0),
			},
			allocatedDevices: []string{},
			reusableDevices:  []string{},
			hint: topologymanager.TopologyHint{
				NUMANodeAffinity: makeSocketMask(0),
				Preferred:        true,
			},
			expectedAvailable:   []string{"Dev0", "Dev1", "Dev2", "Dev3"},
			expectedMustInclude: []string{},
			expectedSize:        1,
		},
		{
			description: "Request for 4, socket 0, 2 already allocated, 2 reusable",
			resource:    "resource",
			request:     4,
			allDevices: []pluginapi.Device{
				makeNUMADevice("Dev0", 0),
				makeNUMADevice("Dev1", 0),
				makeNUMADevice("Dev2", 0),
				makeNUMADevice("Dev3", 0),
				makeNUMADevice("Dev4", 0),
				makeNUMADevice("Dev5", 0),
				makeNUMADevice("Dev6", 0),
				makeNUMADevice("Dev7", 0),
			},
			allocatedDevices: []string{"Dev0", "Dev5"},
			reusableDevices:  []string{"Dev0", "Dev5"},
			hint: topologymanager.TopologyHint{
				NUMANodeAffinity: makeSocketMask(0),
				Preferred:        true,
			},
			expectedAvailable:   []string{"Dev0", "Dev1", "Dev2", "Dev3", "Dev4", "Dev5", "Dev6", "Dev7"},
			expectedMustInclude: []string{"Dev0", "Dev5"},
			expectedSize:        4,
		},
		{
			description: "Request for 4, socket 0, 4 already allocated, 2 reusable",
			resource:    "resource",
			request:     4,
			allDevices: []pluginapi.Device{
				makeNUMADevice("Dev0", 0),
				makeNUMADevice("Dev1", 0),
				makeNUMADevice("Dev2", 0),
				makeNUMADevice("Dev3", 0),
				makeNUMADevice("Dev4", 0),
				makeNUMADevice("Dev5", 0),
				makeNUMADevice("Dev6", 0),
				makeNUMADevice("Dev7", 0),
			},
			allocatedDevices: []string{"Dev0", "Dev5", "Dev4", "Dev1"},
			reusableDevices:  []string{"Dev0", "Dev5"},
			hint: topologymanager.TopologyHint{
				NUMANodeAffinity: makeSocketMask(0),
				Preferred:        true,
			},
			expectedAvailable:   []string{"Dev0", "Dev2", "Dev3", "Dev5", "Dev6", "Dev7"},
			expectedMustInclude: []string{"Dev0", "Dev5"},
			expectedSize:        4,
		},
		{
			description: "Request for 6, multisocket, 2 already allocated, 2 reusable",
			resource:    "resource",
			request:     6,
			allDevices: []pluginapi.Device{
				makeNUMADevice("Dev0", 0),
				makeNUMADevice("Dev1", 0),
				makeNUMADevice("Dev2", 0),
				makeNUMADevice("Dev3", 0),
				makeNUMADevice("Dev4", 1),
				makeNUMADevice("Dev5", 1),
				makeNUMADevice("Dev6", 1),
				makeNUMADevice("Dev7", 1),
			},
			allocatedDevices: []string{"Dev1", "Dev6"},
			reusableDevices:  []string{"Dev1", "Dev6"},
			hint: topologymanager.TopologyHint{
				NUMANodeAffinity: makeSocketMask(0),
				Preferred:        true,
			},
			expectedAvailable:   []string{"Dev0", "Dev1", "Dev2", "Dev3", "Dev4", "Dev5", "Dev6", "Dev7"},
			expectedMustInclude: []string{"Dev0", "Dev1", "Dev2", "Dev3", "Dev6"},
			expectedSize:        6,
		},
		{
			description: "Request for 6, multisocket, 4 already allocated, 2 reusable",
			resource:    "resource",
			request:     6,
			allDevices: []pluginapi.Device{
				makeNUMADevice("Dev0", 0),
				makeNUMADevice("Dev1", 0),
				makeNUMADevice("Dev2", 0),
				makeNUMADevice("Dev3", 0),
				makeNUMADevice("Dev4", 1),
				makeNUMADevice("Dev5", 1),
				makeNUMADevice("Dev6", 1),
				makeNUMADevice("Dev7", 1),
			},
			allocatedDevices: []string{"Dev0", "Dev1", "Dev6", "Dev7"},
			reusableDevices:  []string{"Dev1", "Dev6"},
			hint: topologymanager.TopologyHint{
				NUMANodeAffinity: makeSocketMask(0),
				Preferred:        true,
			},
			expectedAvailable:   []string{"Dev1", "Dev2", "Dev3", "Dev4", "Dev5", "Dev6"},
			expectedMustInclude: []string{"Dev1", "Dev2", "Dev3", "Dev6"},
			expectedSize:        6,
		},
	}
	for _, tc := range tcases {
		m := ManagerImpl{
			allDevices:            make(map[string]map[string]pluginapi.Device),
			healthyDevices:        make(map[string]sets.String),
			allocatedDevices:      make(map[string]sets.String),
			endpoints:             make(map[string]endpointInfo),
			podDevices:            newPodDevices(),
			sourcesReady:          &sourcesReadyStub{},
			activePods:            func() []*v1.Pod { return []*v1.Pod{} },
			topologyAffinityStore: &mockAffinityStore{tc.hint},
		}

		m.allDevices[tc.resource] = make(map[string]pluginapi.Device)
		m.healthyDevices[tc.resource] = sets.NewString()
		for _, d := range tc.allDevices {
			m.allDevices[tc.resource][d.ID] = d
			m.healthyDevices[tc.resource].Insert(d.ID)
		}

		m.allocatedDevices[tc.resource] = sets.NewString()
		for _, d := range tc.allocatedDevices {
			m.allocatedDevices[tc.resource].Insert(d)
		}

		actualAvailable := []string{}
		actualMustInclude := []string{}
		actualSize := 0
		m.endpoints[tc.resource] = endpointInfo{
			e: &MockEndpoint{
				getPreferredAllocationFunc: func(available, mustInclude []string, size int) (*pluginapi.PreferredAllocationResponse, error) {
					actualAvailable = append(actualAvailable, available...)
					actualMustInclude = append(actualMustInclude, mustInclude...)
					actualSize = size
					return nil, nil
				},
			},
			opts: &pluginapi.DevicePluginOptions{GetPreferredAllocationAvailable: true},
		}

		_, err := m.devicesToAllocate("podUID", "containerName", tc.resource, tc.request, sets.NewString(tc.reusableDevices...))
		if err != nil {
			t.Errorf("Unexpected error: %v", err)
			continue
		}

		if !sets.NewString(actualAvailable...).Equal(sets.NewString(tc.expectedAvailable...)) {
			t.Errorf("%v. expected available: %v but got: %v", tc.description, tc.expectedAvailable, actualAvailable)
		}

		if !sets.NewString(actualAvailable...).Equal(sets.NewString(tc.expectedAvailable...)) {
			t.Errorf("%v. expected mustInclude: %v but got: %v", tc.description, tc.expectedMustInclude, actualMustInclude)
		}

		if actualSize != tc.expectedSize {
			t.Errorf("%v. expected size: %v but got: %v", tc.description, tc.expectedSize, actualSize)
		}
	}
}
