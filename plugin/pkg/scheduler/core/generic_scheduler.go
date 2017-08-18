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

package core

import (
	"fmt"
	"sort"
	"strings"
	"sync"
	"sync/atomic"
	"time"

	"k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/util/errors"
	utiltrace "k8s.io/apiserver/pkg/util/trace"
	"k8s.io/client-go/util/workqueue"
	"k8s.io/kubernetes/plugin/pkg/scheduler/algorithm"
	"k8s.io/kubernetes/plugin/pkg/scheduler/algorithm/predicates"
	schedulerapi "k8s.io/kubernetes/plugin/pkg/scheduler/api"
	"k8s.io/kubernetes/plugin/pkg/scheduler/schedulercache"
	"k8s.io/kubernetes/plugin/pkg/scheduler/util"

	"github.com/golang/glog"
)

type FailedPredicateMap map[string][]algorithm.PredicateFailureReason

type FitError struct {
	Pod              *v1.Pod
	FailedPredicates FailedPredicateMap
}

var ErrNoNodesAvailable = fmt.Errorf("no nodes available to schedule pods")

const NoNodeAvailableMsg = "No nodes are available that match all of the following predicates"

// Error returns detailed information of why the pod failed to fit on each node
func (f *FitError) Error() string {
	reasons := make(map[string]int)
	for _, predicates := range f.FailedPredicates {
		for _, pred := range predicates {
			reasons[pred.GetReason()] += 1
		}
	}

	sortReasonsHistogram := func() []string {
		reasonStrings := []string{}
		for k, v := range reasons {
			reasonStrings = append(reasonStrings, fmt.Sprintf("%v (%v)", k, v))
		}
		sort.Strings(reasonStrings)
		return reasonStrings
	}
	reasonMsg := fmt.Sprintf(NoNodeAvailableMsg+": %v.", strings.Join(sortReasonsHistogram(), ", "))
	return reasonMsg
}

type genericScheduler struct {
	cache                 schedulercache.Cache
	equivalenceCache      *EquivalenceCache
	predicates            map[string]algorithm.FitPredicate
	priorityMetaProducer  algorithm.MetadataProducer
	predicateMetaProducer algorithm.PredicateMetadataProducer
	prioritizers          []algorithm.PriorityConfig
	extenders             []algorithm.SchedulerExtender
	pods                  algorithm.PodLister
	lastNodeIndexLock     sync.Mutex
	lastNodeIndex         uint64

	cachedNodeInfoMap map[string]*schedulercache.NodeInfo
}

// Schedule tries to schedule the given pod to one of node in the node list.
// If it succeeds, it will return the name of the node.
// If it fails, it will return a Fiterror error with reasons.
func (g *genericScheduler) Schedule(pod *v1.Pod, nodeLister algorithm.NodeLister) (string, error) {
	trace := utiltrace.New(fmt.Sprintf("Scheduling %s/%s", pod.Namespace, pod.Name))
	defer trace.LogIfLong(100 * time.Millisecond)

	nodes, err := nodeLister.List()
	if err != nil {
		return "", err
	}
	if len(nodes) == 0 {
		return "", ErrNoNodesAvailable
	}

	// Used for all fit and priority funcs.
	err = g.cache.UpdateNodeNameToInfoMap(g.cachedNodeInfoMap)
	if err != nil {
		return "", err
	}

	trace.Step("Computing predicates")
	filteredNodes, failedPredicateMap, err := findNodesThatFit(pod, g.cachedNodeInfoMap, nodes, g.predicates, g.extenders, g.predicateMetaProducer, g.equivalenceCache)
	if err != nil {
		return "", err
	}

	if len(filteredNodes) == 0 {
		return "", &FitError{
			Pod:              pod,
			FailedPredicates: failedPredicateMap,
		}
	}

	trace.Step("Prioritizing")
	metaPrioritiesInterface := g.priorityMetaProducer(pod, g.cachedNodeInfoMap)
	priorityList, err := PrioritizeNodes(pod, g.cachedNodeInfoMap, metaPrioritiesInterface, g.prioritizers, filteredNodes, g.extenders)
	if err != nil {
		return "", err
	}

	trace.Step("Selecting host")
	return g.selectHost(priorityList)
}

// Prioritizers returns a slice containing all the scheduler's priority
// functions and their config. It is exposed for testing only.
func (g *genericScheduler) Prioritizers() []algorithm.PriorityConfig {
	return g.prioritizers
}

// Predicates returns a map containing all the scheduler's predicate
// functions. It is exposed for testing only.
func (g *genericScheduler) Predicates() map[string]algorithm.FitPredicate {
	return g.predicates
}

// selectHost takes a prioritized list of nodes and then picks one
// in a round-robin manner from the nodes that had the highest score.
func (g *genericScheduler) selectHost(priorityList schedulerapi.HostPriorityList) (string, error) {
	if len(priorityList) == 0 {
		return "", fmt.Errorf("empty priorityList")
	}

	sort.Sort(sort.Reverse(priorityList))
	maxScore := priorityList[0].Score
	firstAfterMaxScore := sort.Search(len(priorityList), func(i int) bool { return priorityList[i].Score < maxScore })

	g.lastNodeIndexLock.Lock()
	ix := int(g.lastNodeIndex % uint64(firstAfterMaxScore))
	g.lastNodeIndex++
	g.lastNodeIndexLock.Unlock()

	return priorityList[ix].Host, nil
}

// preempt finds nodes with pods that can be preempted to make room for "pod" to
// schedule. It chooses one of the nodes and preempts the pods on the node and
// returns the node name if such a node is found.
// TODO(bsalamat): This function is under construction! DO NOT USE!
func (g *genericScheduler) preempt(pod *v1.Pod, nodeLister algorithm.NodeLister) (string, error) {
	nodes, err := nodeLister.List()
	if err != nil {
		return "", err
	}
	if len(nodes) == 0 {
		return "", ErrNoNodesAvailable
	}
	nodeToPods := selectNodesForPreemption(pod, g.cachedNodeInfoMap, nodes, g.predicates, g.predicateMetaProducer)
	if len(nodeToPods) == 0 {
		return "", nil
	}
	// TODO: Add a node scoring mechanism and perform preemption
	return "", nil
}

// Filters the nodes to find the ones that fit based on the given predicate functions
// Each node is passed through the predicate functions to determine if it is a fit
func findNodesThatFit(
	pod *v1.Pod,
	nodeNameToInfo map[string]*schedulercache.NodeInfo,
	nodes []*v1.Node,
	predicateFuncs map[string]algorithm.FitPredicate,
	extenders []algorithm.SchedulerExtender,
	metadataProducer algorithm.PredicateMetadataProducer,
	ecache *EquivalenceCache,
) ([]*v1.Node, FailedPredicateMap, error) {
	var filtered []*v1.Node
	failedPredicateMap := FailedPredicateMap{}

	if len(predicateFuncs) == 0 {
		filtered = nodes
	} else {
		// Create filtered list with enough space to avoid growing it
		// and allow assigning.
		filtered = make([]*v1.Node, len(nodes))
		errs := errors.MessageCountMap{}
		var predicateResultLock sync.Mutex
		var filteredLen int32

		// We can use the same metadata producer for all nodes.
		meta := metadataProducer(pod, nodeNameToInfo)
		checkNode := func(i int) {
			nodeName := nodes[i].Name
			fits, failedPredicates, err := podFitsOnNode(pod, meta, nodeNameToInfo[nodeName], predicateFuncs, ecache)
			if err != nil {
				predicateResultLock.Lock()
				errs[err.Error()]++
				predicateResultLock.Unlock()
				return
			}
			if fits {
				filtered[atomic.AddInt32(&filteredLen, 1)-1] = nodes[i]
			} else {
				predicateResultLock.Lock()
				failedPredicateMap[nodeName] = failedPredicates
				predicateResultLock.Unlock()
			}
		}
		workqueue.Parallelize(16, len(nodes), checkNode)
		filtered = filtered[:filteredLen]
		if len(errs) > 0 {
			return []*v1.Node{}, FailedPredicateMap{}, errors.CreateAggregateFromMessageCountMap(errs)
		}
	}

	if len(filtered) > 0 && len(extenders) != 0 {
		for _, extender := range extenders {
			filteredList, failedMap, err := extender.Filter(pod, filtered, nodeNameToInfo)
			if err != nil {
				return []*v1.Node{}, FailedPredicateMap{}, err
			}

			for failedNodeName, failedMsg := range failedMap {
				if _, found := failedPredicateMap[failedNodeName]; !found {
					failedPredicateMap[failedNodeName] = []algorithm.PredicateFailureReason{}
				}
				failedPredicateMap[failedNodeName] = append(failedPredicateMap[failedNodeName], predicates.NewFailureReason(failedMsg))
			}
			filtered = filteredList
			if len(filtered) == 0 {
				break
			}
		}
	}
	return filtered, failedPredicateMap, nil
}

// Checks whether node with a given name and NodeInfo satisfies all predicateFuncs.
func podFitsOnNode(pod *v1.Pod, meta algorithm.PredicateMetadata, info *schedulercache.NodeInfo, predicateFuncs map[string]algorithm.FitPredicate,
	ecache *EquivalenceCache) (bool, []algorithm.PredicateFailureReason, error) {
	var (
		equivalenceHash  uint64
		failedPredicates []algorithm.PredicateFailureReason
		eCacheAvailable  bool
		invalid          bool
		fit              bool
		reasons          []algorithm.PredicateFailureReason
		err              error
	)
	if ecache != nil {
		// getHashEquivalencePod will return immediately if no equivalence pod found
		equivalenceHash = ecache.getHashEquivalencePod(pod)
		eCacheAvailable = (equivalenceHash != 0)
	}
	for predicateKey, predicate := range predicateFuncs {
		// If equivalenceCache is available
		if eCacheAvailable {
			// PredicateWithECache will returns it's cached predicate results
			fit, reasons, invalid = ecache.PredicateWithECache(pod.GetName(), info.Node().GetName(), predicateKey, equivalenceHash)
		}

		if !eCacheAvailable || invalid {
			// we need to execute predicate functions since equivalence cache does not work
			fit, reasons, err = predicate(pod, meta, info)
			if err != nil {
				return false, []algorithm.PredicateFailureReason{}, err
			}

			if eCacheAvailable {
				// update equivalence cache with newly computed fit & reasons
				// TODO(resouer) should we do this in another thread? any race?
				ecache.UpdateCachedPredicateItem(pod.GetName(), info.Node().GetName(), predicateKey, fit, reasons, equivalenceHash)
			}
		}

		if !fit {
			// eCache is available and valid, and predicates result is unfit, record the fail reasons
			failedPredicates = append(failedPredicates, reasons...)
		}
	}
	return len(failedPredicates) == 0, failedPredicates, nil
}

// Prioritizes the nodes by running the individual priority functions in parallel.
// Each priority function is expected to set a score of 0-10
// 0 is the lowest priority score (least preferred node) and 10 is the highest
// Each priority function can also have its own weight
// The node scores returned by the priority function are multiplied by the weights to get weighted scores
// All scores are finally combined (added) to get the total weighted scores of all nodes
func PrioritizeNodes(
	pod *v1.Pod,
	nodeNameToInfo map[string]*schedulercache.NodeInfo,
	meta interface{},
	priorityConfigs []algorithm.PriorityConfig,
	nodes []*v1.Node,
	extenders []algorithm.SchedulerExtender,
) (schedulerapi.HostPriorityList, error) {
	// If no priority configs are provided, then the EqualPriority function is applied
	// This is required to generate the priority list in the required format
	if len(priorityConfigs) == 0 && len(extenders) == 0 {
		result := make(schedulerapi.HostPriorityList, 0, len(nodes))
		for i := range nodes {
			hostPriority, err := EqualPriorityMap(pod, meta, nodeNameToInfo[nodes[i].Name])
			if err != nil {
				return nil, err
			}
			result = append(result, hostPriority)
		}
		return result, nil
	}

	var (
		mu   = sync.Mutex{}
		wg   = sync.WaitGroup{}
		errs []error
	)
	appendError := func(err error) {
		mu.Lock()
		defer mu.Unlock()
		errs = append(errs, err)
	}

	results := make([]schedulerapi.HostPriorityList, len(priorityConfigs), len(priorityConfigs))

	for i, priorityConfig := range priorityConfigs {
		if priorityConfig.Function != nil {
			// DEPRECATED
			wg.Add(1)
			go func(index int, config algorithm.PriorityConfig) {
				defer wg.Done()
				var err error
				results[index], err = config.Function(pod, nodeNameToInfo, nodes)
				if err != nil {
					appendError(err)
				}
			}(i, priorityConfig)
		} else {
			results[i] = make(schedulerapi.HostPriorityList, len(nodes))
		}
	}
	processNode := func(index int) {
		nodeInfo := nodeNameToInfo[nodes[index].Name]
		var err error
		for i := range priorityConfigs {
			if priorityConfigs[i].Function != nil {
				continue
			}
			results[i][index], err = priorityConfigs[i].Map(pod, meta, nodeInfo)
			if err != nil {
				appendError(err)
				return
			}
		}
	}
	workqueue.Parallelize(16, len(nodes), processNode)
	for i, priorityConfig := range priorityConfigs {
		if priorityConfig.Reduce == nil {
			continue
		}
		wg.Add(1)
		go func(index int, config algorithm.PriorityConfig) {
			defer wg.Done()
			if err := config.Reduce(pod, meta, nodeNameToInfo, results[index]); err != nil {
				appendError(err)
			}
		}(i, priorityConfig)
	}
	// Wait for all computations to be finished.
	wg.Wait()
	if len(errs) != 0 {
		return schedulerapi.HostPriorityList{}, errors.NewAggregate(errs)
	}

	// Summarize all scores.
	result := make(schedulerapi.HostPriorityList, 0, len(nodes))

	for i := range nodes {
		result = append(result, schedulerapi.HostPriority{Host: nodes[i].Name, Score: 0})
		for j := range priorityConfigs {
			result[i].Score += results[j][i].Score * priorityConfigs[j].Weight
		}
	}

	if len(extenders) != 0 && nodes != nil {
		combinedScores := make(map[string]int, len(nodeNameToInfo))
		for _, extender := range extenders {
			wg.Add(1)
			go func(ext algorithm.SchedulerExtender) {
				defer wg.Done()
				prioritizedList, weight, err := ext.Prioritize(pod, nodes)
				if err != nil {
					// Prioritization errors from extender can be ignored, let k8s/other extenders determine the priorities
					return
				}
				mu.Lock()
				for i := range *prioritizedList {
					host, score := (*prioritizedList)[i].Host, (*prioritizedList)[i].Score
					combinedScores[host] += score * weight
				}
				mu.Unlock()
			}(extender)
		}
		// wait for all go routines to finish
		wg.Wait()
		for i := range result {
			result[i].Score += combinedScores[result[i].Host]
		}
	}

	if glog.V(10) {
		for i := range result {
			glog.V(10).Infof("Host %s => Score %d", result[i].Host, result[i].Score)
		}
	}
	return result, nil
}

// EqualPriority is a prioritizer function that gives an equal weight of one to all nodes
func EqualPriorityMap(_ *v1.Pod, _ interface{}, nodeInfo *schedulercache.NodeInfo) (schedulerapi.HostPriority, error) {
	node := nodeInfo.Node()
	if node == nil {
		return schedulerapi.HostPriority{}, fmt.Errorf("node not found")
	}
	return schedulerapi.HostPriority{
		Host:  node.Name,
		Score: 1,
	}, nil
}

// selectNodesForPreemption finds all the nodes with possible victims for
// preemption in parallel.
func selectNodesForPreemption(pod *v1.Pod,
	nodeNameToInfo map[string]*schedulercache.NodeInfo,
	nodes []*v1.Node,
	predicates map[string]algorithm.FitPredicate,
	metadataProducer algorithm.PredicateMetadataProducer,
) map[string][]*v1.Pod {

	nodeNameToPods := map[string][]*v1.Pod{}
	var resultLock sync.Mutex

	// We can use the same metadata producer for all nodes.
	meta := metadataProducer(pod, nodeNameToInfo)
	checkNode := func(i int) {
		nodeName := nodes[i].Name
		pods, fits := selectVictimsOnNode(pod, meta.ShallowCopy(), nodeNameToInfo[nodeName], predicates)
		if fits && len(pods) != 0 {
			resultLock.Lock()
			nodeNameToPods[nodeName] = pods
			resultLock.Unlock()
		}
	}
	workqueue.Parallelize(16, len(nodes), checkNode)
	return nodeNameToPods
}

// selectVictimsOnNode finds minimum set of pods on the given node that should
// be preempted in order to make enough room for "pod" to be scheduled. The
// minimum set selected is subject to the constraint that a higher-priority pod
// is never preempted when a lower-priority pod could be (higher/lower relative
// to one another, not relative to the preemptor "pod").
// The algorithm first checks if the pod can be scheduled on the node when all the
// lower priority pods are gone. If so, it sorts all the lower priority pods by
// their priority and starting from the highest priority one, tries to keep as
// many of them as possible while checking that the "pod" can still fit on the node.
// NOTE: This function assumes that it is never called if "pod" cannot be scheduled
// due to pod affinity, node affinity, or node anti-affinity reasons. None of
// these predicates can be satisfied by removing more pods from the node.
func selectVictimsOnNode(pod *v1.Pod, meta algorithm.PredicateMetadata, nodeInfo *schedulercache.NodeInfo, fitPredicates map[string]algorithm.FitPredicate) ([]*v1.Pod, bool) {
	higherPriority := func(pod1, pod2 interface{}) bool {
		return util.GetPodPriority(pod1.(*v1.Pod)) > util.GetPodPriority(pod2.(*v1.Pod))
	}
	potentialVictims := util.SortableList{CompFunc: higherPriority}
	nodeInfoCopy := nodeInfo.Clone()

	removePod := func(rp *v1.Pod) {
		nodeInfoCopy.RemovePod(rp)
		meta.RemovePod(rp)
	}
	addPod := func(ap *v1.Pod) {
		nodeInfoCopy.AddPod(ap)
		meta.AddPod(ap, nodeInfoCopy)
	}
	// As the first step, remove all the lower priority pods from the node and
	// check if the given pod can be scheduled.
	podPriority := util.GetPodPriority(pod)
	for _, p := range nodeInfoCopy.Pods() {
		if util.GetPodPriority(p) < podPriority {
			potentialVictims.Items = append(potentialVictims.Items, p)
			removePod(p)
		}
	}
	potentialVictims.Sort()
	// If the new pod does not fit after removing all the lower priority pods,
	// we are almost done and this node is not suitable for preemption. The only condition
	// that we should check is if the "pod" is failing to schedule due to pod affinity
	// failure.
	if fits, failedPredicates, err := podFitsOnNode(pod, meta, nodeInfoCopy, fitPredicates, nil); !fits {
		if err != nil {
			glog.Warningf("Encountered error while selecting victims on node %v: %v", nodeInfo.Node().Name, err)
			return nil, false
		}
		// If the new pod still cannot be scheduled for any reason other than pod
		// affinity, the new pod will not fit on this node and we are done here.
		affinity := pod.Spec.Affinity
		if affinity == nil || affinity.PodAffinity == nil {
			return nil, false
		}
		for _, failedPred := range failedPredicates {
			if failedPred != predicates.ErrPodAffinityNotMatch {
				return nil, false
			}
		}
		// If we reach here, it means that the pod cannot be scheduled due to pod
		// affinity or anti-affinity. Since failure reason for both affinity and
		// anti-affinity is the same, we cannot say which one caused it. So, we try
		// adding pods one at a time and see if any of them satisfies the affinity rules.
		for i, p := range potentialVictims.Items {
			existingPod := p.(*v1.Pod)
			addPod(existingPod)
			if fits, _, _ = podFitsOnNode(pod, meta, nodeInfoCopy, fitPredicates, nil); !fits {
				removePod(existingPod)
			} else {
				// We found the pod needed to satisfy pod affinity. Let's remove it from
				// potential victims list.
				// NOTE: We assume that pod affinity can be satisfied by only one pod,
				// not multiple pods. This is how scheduler works today.
				potentialVictims.Items = append(potentialVictims.Items[:i], potentialVictims.Items[i+1:]...)
				break
			}
		}
		if !fits {
			return nil, false
		}
	}
	victims := []*v1.Pod{}
	// Try to reprieve as may pods as possible starting from the highest priority one.
	for _, p := range potentialVictims.Items {
		lpp := p.(*v1.Pod)
		addPod(lpp)
		if fits, _, _ := podFitsOnNode(pod, meta, nodeInfoCopy, fitPredicates, nil); !fits {
			removePod(lpp)
			victims = append(victims, lpp)
		}
	}
	return victims, true
}

func NewGenericScheduler(
	cache schedulercache.Cache,
	eCache *EquivalenceCache,
	predicates map[string]algorithm.FitPredicate,
	predicateMetaProducer algorithm.PredicateMetadataProducer,
	prioritizers []algorithm.PriorityConfig,
	priorityMetaProducer algorithm.MetadataProducer,
	extenders []algorithm.SchedulerExtender) algorithm.ScheduleAlgorithm {
	return &genericScheduler{
		cache:                 cache,
		equivalenceCache:      eCache,
		predicates:            predicates,
		predicateMetaProducer: predicateMetaProducer,
		prioritizers:          prioritizers,
		priorityMetaProducer:  priorityMetaProducer,
		extenders:             extenders,
		cachedNodeInfoMap:     make(map[string]*schedulercache.NodeInfo),
	}
}
