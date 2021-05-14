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
	"sort"

	"k8s.io/klog/v2"

	"k8s.io/kubernetes/pkg/kubelet/cm/cpumanager/topology"
	"k8s.io/kubernetes/pkg/kubelet/cm/cpuset"
)

type cpuAccumulator struct {
	topo          *topology.CPUTopology
	details       topology.CPUDetails
	numCPUsNeeded int
	result        cpuset.CPUSet
}

func newCPUAccumulator(topo *topology.CPUTopology, availableCPUs cpuset.CPUSet, numCPUs int) *cpuAccumulator {
	return &cpuAccumulator{
		topo:          topo,
		details:       topo.CPUDetails.KeepOnly(availableCPUs),
		numCPUsNeeded: numCPUs,
		result:        cpuset.NewCPUSet(),
	}
}

// Returns true if the supplied NUMANode is fully available in `topoDetails`.
func (a *cpuAccumulator) isNUMANodeFree(numaID int) bool {
	return a.details.CPUsInNUMANodes(numaID).Size() == a.topo.CPUDetails.CPUsInNUMANodes(numaID).Size()
}

// Returns true if the supplied socket is fully available in `topoDetails`.
func (a *cpuAccumulator) isSocketFree(socketID int) bool {
	return a.details.CPUsInSockets(socketID).Size() == a.topo.CPUsPerSocket()
}

// Returns true if the supplied core is fully available in `topoDetails`.
func (a *cpuAccumulator) isCoreFree(coreID int) bool {
	return a.details.CPUsInCores(coreID).Size() == a.topo.CPUsPerCore()
}

// Returns free NUMA Node IDs as a slice sorted by sortAvailableNUMANodes().
func (a *cpuAccumulator) freeNUMANodes() []int {
	free := []int{}
	for _, numa := range a.sortAvailableNUMANodes() {
		if a.isNUMANodeFree(numa) {
			free = append(free, numa)
		}
	}
	return free
}

// Returns free socket IDs as a slice sorted by sortAvailableSockets().
func (a *cpuAccumulator) freeSockets() []int {
	free := []int{}
	for _, socket := range a.sortAvailableSockets() {
		if a.isSocketFree(socket) {
			free = append(free, socket)
		}
	}
	return free
}

// Returns free core IDs as a slice sorted by sortAvailableCores().
func (a *cpuAccumulator) freeCores() []int {
	free := []int{}
	for _, core := range a.sortAvailableCores() {
		if a.isCoreFree(core) {
			free = append(free, core)
		}
	}
	return free
}

// Returns free CPU IDs as a slice sorted by sortAvailableCPUs().
func (a *cpuAccumulator) freeCPUs() []int {
	return a.sortAvailableCPUs()
}

// Sorts the provided list of NUMA nodes/sockets/cores/cpus referenced in 'ids'
// by the number of available CPUs contained within them (smallest to largest).
// The 'getCPU()' paramater defines the function that should be called to
// retrieve the list of available CPUs for the type being referenced. If two
// NUMA nodes/sockets/cores/cpus have the same number of available CPUs, they
// are sorted in ascending order by their id.
func (a *cpuAccumulator) sort(ids []int, getCPUs func(ids ...int) cpuset.CPUSet) {
	sort.Slice(ids,
		func(i, j int) bool {
			iCPUs := getCPUs(ids[i])
			jCPUs := getCPUs(ids[j])
			if iCPUs.Size() < jCPUs.Size() {
				return true
			}
			if iCPUs.Size() > jCPUs.Size() {
				return false
			}
			return ids[i] < ids[j]
		})
}

// Sort all NUMA nodes with free CPUs:
// - If NUMA nodes are higher than sockets in the memory hierarchy then sort
//   them directly using the sort() algorithm defined above.
// - Otherwise sort them:
//   - First by socket using sortAvailableSockets().
//   - Then within each socket, using the sort() algorithm defined above.
func (a *cpuAccumulator) sortAvailableNUMANodes() []int {
	// If NUMA nodes are equal or higher in the memory hierarchy than sockets
	if a.topo.NumSockets >= a.topo.NumNUMANodes {
		numas := a.details.NUMANodes().ToSliceNoSort()
		a.sort(numas, a.details.CPUsInNUMANodes)
		return numas
	}

	// Otherwise each socket has multiple NUMA nodes
	var result []int
	for _, socket := range a.sortAvailableSockets() {
		numas := a.details.NUMANodesInSockets(socket).ToSliceNoSort()
		a.sort(numas, a.details.CPUsInNUMANodes)
		result = append(result, numas...)
	}
	return result
}

// Sort all sockets with free CPUs:
// - If sockets are higher than NUMA nodes in the memory hierarchy then sort
//   them directly using the sort() algorithm defined above.
// - Otherwise sort them:
//   - First by NUMA node using sortAvailableNUMANodes().
//   - Then within each NUMA node, using the sort() algorithm defined above.
func (a *cpuAccumulator) sortAvailableSockets() []int {
	// If sockets are higher than NUMA nodes in the memory hierarchy
	if a.topo.NumNUMANodes >= a.topo.NumSockets {
		sockets := a.details.Sockets().ToSliceNoSort()
		a.sort(sockets, a.details.CPUsInSockets)
		return sockets
	}

	// Otherwise each NUMA Node has multiple sockets
	var result []int
	for _, numa := range a.sortAvailableNUMANodes() {
		sockets := a.details.SocketsInNUMANodes(numa).ToSliceNoSort()
		a.sort(sockets, a.details.CPUsInSockets)
		result = append(result, sockets...)
	}
	return result
}

// Sort all cores with free CPUs:
// - First by socket (or NUMA node) using sortAvailableSockets() (or sortAvailableNUMANodes()).
// - Then within each socket or NUMA node, using the sort() algorithm defined above.
func (a *cpuAccumulator) sortAvailableCores() []int {
	// If NUMA nodes are higher in the memory hierarchy than sockets, then
	// cores sit directly below sockets in the memory hierarchy.
	if a.topo.NumSockets >= a.topo.NumNUMANodes {
		var result []int
		for _, socket := range a.sortAvailableSockets() {
			cores := a.details.CoresInSockets(socket).ToSliceNoSort()
			a.sort(cores, a.details.CPUsInCores)
			result = append(result, cores...)
		}
		return result
	}

	// Otherwise they sit directly below NUMA nodes.
	var result []int
	for _, numa := range a.sortAvailableNUMANodes() {
		cores := a.details.CoresInNUMANodes(numa).ToSliceNoSort()
		a.sort(cores, a.details.CPUsInCores)
		result = append(result, cores...)
	}
	return result
}

// Sort all available CPUs:
// - First by core using sortAvailableCores().
// - Then within each core, using the sort() algorithm defined above.
func (a *cpuAccumulator) sortAvailableCPUs() []int {
	var result []int
	for _, core := range a.sortAvailableCores() {
		cpus := a.details.CPUsInCores(core).ToSliceNoSort()
		sort.Ints(cpus)
		result = append(result, cpus...)
	}
	return result
}

func (a *cpuAccumulator) take(cpus cpuset.CPUSet) {
	a.result = a.result.Union(cpus)
	a.details = a.details.KeepOnly(a.details.CPUs().Difference(a.result))
	a.numCPUsNeeded -= cpus.Size()
}

func (a *cpuAccumulator) takeFullNUMANodes() {
	for _, numa := range a.freeNUMANodes() {
		cpusInNUMANode := a.topo.CPUDetails.CPUsInNUMANodes(numa)
		if !a.needs(cpusInNUMANode.Size()) {
			continue
		}
		klog.V(4).InfoS("takeFullNUMANodes: claiming NUMA node", "numa", numa)
		a.take(cpusInNUMANode)
	}
}

func (a *cpuAccumulator) takeFullSockets() {
	for _, socket := range a.freeSockets() {
		cpusInSocket := a.topo.CPUDetails.CPUsInSockets(socket)
		if !a.needs(cpusInSocket.Size()) {
			continue
		}
		klog.V(4).InfoS("takeFullSockets: claiming socket", "socket", socket)
		a.take(cpusInSocket)
	}
}

func (a *cpuAccumulator) takeFullCores() {
	for _, core := range a.freeCores() {
		cpusInCore := a.topo.CPUDetails.CPUsInCores(core)
		if !a.needs(cpusInCore.Size()) {
			continue
		}
		klog.V(4).InfoS("takeFullCores: claiming core", "core", core)
		a.take(cpusInCore)
	}
}

func (a *cpuAccumulator) takeRemainingCPUs() {
	for _, cpu := range a.sortAvailableCPUs() {
		klog.V(4).InfoS("takeRemainingCPUs: claiming CPU", "cpu", cpu)
		a.take(cpuset.NewCPUSet(cpu))
		if a.isSatisfied() {
			return
		}
	}
}

func (a *cpuAccumulator) needs(n int) bool {
	return a.numCPUsNeeded >= n
}

func (a *cpuAccumulator) isSatisfied() bool {
	return a.numCPUsNeeded < 1
}

func (a *cpuAccumulator) isFailed() bool {
	return a.numCPUsNeeded > a.details.CPUs().Size()
}

func takeByTopology(topo *topology.CPUTopology, availableCPUs cpuset.CPUSet, numCPUs int) (cpuset.CPUSet, error) {
	acc := newCPUAccumulator(topo, availableCPUs, numCPUs)
	if acc.isSatisfied() {
		return acc.result, nil
	}
	if acc.isFailed() {
		return cpuset.NewCPUSet(), fmt.Errorf("not enough cpus available to satisfy request")
	}

	// Algorithm: topology-aware best-fit
	// 1. Acquire whole NUMA nodes and sockets, if available and the container
	//    requires at least a NUMA node or socket's-worth of CPUs. If NUMA
	//    Nodes map to 1 or more sockets, pull from NUMA nodes first.
	//    Otherwise pull from sockets first.
	if acc.topo.NumSockets >= acc.topo.NumNUMANodes {
		acc.takeFullNUMANodes()
		if acc.isSatisfied() {
			return acc.result, nil
		}

		acc.takeFullSockets()
		if acc.isSatisfied() {
			return acc.result, nil
		}
	} else {
		acc.takeFullSockets()
		if acc.isSatisfied() {
			return acc.result, nil
		}

		acc.takeFullNUMANodes()
		if acc.isSatisfied() {
			return acc.result, nil
		}
	}

	// 2. Acquire whole cores, if available and the container requires at least
	//    a core's-worth of CPUs.
	acc.takeFullCores()
	if acc.isSatisfied() {
		return acc.result, nil
	}

	// 3. Acquire single threads, preferring to fill partially-allocated cores
	//    on the same sockets as the whole cores we have already taken in this
	//    allocation.
	acc.takeRemainingCPUs()
	if acc.isSatisfied() {
		return acc.result, nil
	}

	return cpuset.NewCPUSet(), fmt.Errorf("failed to allocate cpus")
}
