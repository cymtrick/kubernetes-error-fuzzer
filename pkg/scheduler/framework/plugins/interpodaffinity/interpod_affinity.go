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

package interpodaffinity

import (
	"context"
	"fmt"
	"sync"

	"k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/util/sets"
	"k8s.io/client-go/util/workqueue"
	"k8s.io/klog"
	"k8s.io/kubernetes/pkg/scheduler/algorithm/predicates"
	priorityutil "k8s.io/kubernetes/pkg/scheduler/algorithm/priorities/util"
	"k8s.io/kubernetes/pkg/scheduler/framework/plugins/migration"
	framework "k8s.io/kubernetes/pkg/scheduler/framework/v1alpha1"
	schedulerlisters "k8s.io/kubernetes/pkg/scheduler/listers"
	"k8s.io/kubernetes/pkg/scheduler/nodeinfo"
	schedutil "k8s.io/kubernetes/pkg/scheduler/util"
)

// InterPodAffinity is a plugin that checks inter pod affinity
type InterPodAffinity struct {
	sharedLister          schedulerlisters.SharedLister
	podAffinityChecker    *predicates.PodAffinityChecker
	hardPodAffinityWeight int32
	sync.Mutex
}

// Args holds the args that are used to configure the plugin.
type Args struct {
	HardPodAffinityWeight int32 `json:"hardPodAffinityWeight,omitempty"`
}

var _ framework.PreFilterPlugin = &InterPodAffinity{}
var _ framework.FilterPlugin = &InterPodAffinity{}
var _ framework.PostFilterPlugin = &InterPodAffinity{}
var _ framework.ScorePlugin = &InterPodAffinity{}

const (
	// Name is the name of the plugin used in the plugin registry and configurations.
	Name = "InterPodAffinity"

	// preFilterStateKey is the key in CycleState to InterPodAffinity pre-computed data for Filtering.
	// Using the name of the plugin will likely help us avoid collisions with other plugins.
	preFilterStateKey = "PreFilter" + Name

	// postFilterStateKey is the key in CycleState to InterPodAffinity pre-computed data for Scoring.
	postFilterStateKey = "PostFilter" + Name
)

// preFilterState computed at PreFilter and used at Filter.
type preFilterState struct {
	meta *predicates.PodAffinityMetadata
}

// Clone the prefilter state.
func (s *preFilterState) Clone() framework.StateData {
	copy := &preFilterState{
		meta: s.meta.Clone(),
	}
	return copy
}

// Name returns name of the plugin. It is used in logs, etc.
func (pl *InterPodAffinity) Name() string {
	return Name
}

// PreFilter invoked at the prefilter extension point.
func (pl *InterPodAffinity) PreFilter(ctx context.Context, cycleState *framework.CycleState, pod *v1.Pod) *framework.Status {
	var meta *predicates.PodAffinityMetadata
	var allNodes []*nodeinfo.NodeInfo
	var havePodsWithAffinityNodes []*nodeinfo.NodeInfo
	var err error
	if allNodes, err = pl.sharedLister.NodeInfos().List(); err != nil {
		return framework.NewStatus(framework.Error, fmt.Sprintf("failed to list NodeInfos: %v", err))
	}
	if havePodsWithAffinityNodes, err = pl.sharedLister.NodeInfos().HavePodsWithAffinityList(); err != nil {
		return framework.NewStatus(framework.Error, fmt.Sprintf("failed to list NodeInfos with pods with affinity: %v", err))
	}
	if meta, err = predicates.GetPodAffinityMetadata(pod, allNodes, havePodsWithAffinityNodes); err != nil {
		return framework.NewStatus(framework.Error, fmt.Sprintf("Error calculating podAffinityMetadata: %v", err))
	}

	s := &preFilterState{
		meta: meta,
	}
	cycleState.Write(preFilterStateKey, s)
	return nil
}

// PreFilterExtensions returns prefilter extensions, pod add and remove.
func (pl *InterPodAffinity) PreFilterExtensions() framework.PreFilterExtensions {
	return pl
}

// AddPod from pre-computed data in cycleState.
func (pl *InterPodAffinity) AddPod(ctx context.Context, cycleState *framework.CycleState, podToSchedule *v1.Pod, podToAdd *v1.Pod, nodeInfo *nodeinfo.NodeInfo) *framework.Status {
	meta, err := getPodAffinityMetadata(cycleState)
	if err != nil {
		return framework.NewStatus(framework.Error, err.Error())
	}
	meta.UpdateWithPod(podToAdd, podToSchedule, nodeInfo.Node(), 1)
	return nil
}

// RemovePod from pre-computed data in cycleState.
func (pl *InterPodAffinity) RemovePod(ctx context.Context, cycleState *framework.CycleState, podToSchedule *v1.Pod, podToRemove *v1.Pod, nodeInfo *nodeinfo.NodeInfo) *framework.Status {
	meta, err := getPodAffinityMetadata(cycleState)
	if err != nil {
		return framework.NewStatus(framework.Error, err.Error())
	}
	meta.UpdateWithPod(podToRemove, podToSchedule, nodeInfo.Node(), -1)
	return nil
}

func getPodAffinityMetadata(cycleState *framework.CycleState) (*predicates.PodAffinityMetadata, error) {
	c, err := cycleState.Read(preFilterStateKey)
	if err != nil {
		// The metadata wasn't pre-computed in prefilter. We ignore the error for now since
		// Filter is able to handle that by computing it again.
		klog.V(5).Infof("Error reading %q from cycleState: %v", preFilterStateKey, err)
		return nil, nil
	}

	s, ok := c.(*preFilterState)
	if !ok {
		return nil, fmt.Errorf("%+v  convert to interpodaffinity.state error", c)
	}
	return s.meta, nil
}

// Filter invoked at the filter extension point.
func (pl *InterPodAffinity) Filter(ctx context.Context, cycleState *framework.CycleState, pod *v1.Pod, nodeInfo *nodeinfo.NodeInfo) *framework.Status {
	meta, err := getPodAffinityMetadata(cycleState)
	if err != nil {
		return framework.NewStatus(framework.Error, err.Error())
	}
	_, reasons, err := pl.podAffinityChecker.InterPodAffinityMatches(pod, meta, nodeInfo)
	return migration.PredicateResultToFrameworkStatus(reasons, err)
}

// A "processed" representation of v1.WeightedAffinityTerm.
type weightedAffinityTerm struct {
	namespaces  sets.String
	selector    labels.Selector
	weight      int32
	topologyKey string
}

// postFilterState computed at PostFilter and used at Score.
type postFilterState struct {
	topologyScore     map[string]map[string]int64
	affinityTerms     []*weightedAffinityTerm
	antiAffinityTerms []*weightedAffinityTerm
}

// Clone implements the mandatory Clone interface. We don't really copy the data since
// there is no need for that.
func (s *postFilterState) Clone() framework.StateData {
	return s
}

func newWeightedAffinityTerm(pod *v1.Pod, term *v1.PodAffinityTerm, weight int32) (*weightedAffinityTerm, error) {
	namespaces := priorityutil.GetNamespacesFromPodAffinityTerm(pod, term)
	selector, err := metav1.LabelSelectorAsSelector(term.LabelSelector)
	if err != nil {
		return nil, err
	}
	return &weightedAffinityTerm{namespaces: namespaces, selector: selector, topologyKey: term.TopologyKey, weight: weight}, nil
}

func getProcessedTerms(pod *v1.Pod, terms []v1.WeightedPodAffinityTerm) ([]*weightedAffinityTerm, error) {
	if terms == nil {
		return nil, nil
	}

	var processedTerms []*weightedAffinityTerm
	for i := range terms {
		p, err := newWeightedAffinityTerm(pod, &terms[i].PodAffinityTerm, terms[i].Weight)
		if err != nil {
			return nil, err
		}
		processedTerms = append(processedTerms, p)
	}
	return processedTerms, nil
}

func (pl *InterPodAffinity) processTerm(
	state *postFilterState,
	term *weightedAffinityTerm,
	podToCheck *v1.Pod,
	fixedNode *v1.Node,
	multiplier int,
) {
	if len(fixedNode.Labels) == 0 {
		return
	}

	match := priorityutil.PodMatchesTermsNamespaceAndSelector(podToCheck, term.namespaces, term.selector)
	tpValue, tpValueExist := fixedNode.Labels[term.topologyKey]
	if match && tpValueExist {
		pl.Lock()
		if state.topologyScore[term.topologyKey] == nil {
			state.topologyScore[term.topologyKey] = make(map[string]int64)
		}
		state.topologyScore[term.topologyKey][tpValue] += int64(term.weight * int32(multiplier))
		pl.Unlock()
	}
	return
}

func (pl *InterPodAffinity) processTerms(state *postFilterState, terms []*weightedAffinityTerm, podToCheck *v1.Pod, fixedNode *v1.Node, multiplier int) error {
	for _, term := range terms {
		pl.processTerm(state, term, podToCheck, fixedNode, multiplier)
	}
	return nil
}

func (pl *InterPodAffinity) processExistingPod(state *postFilterState, existingPod *v1.Pod, existingPodNodeInfo *nodeinfo.NodeInfo, incomingPod *v1.Pod) error {
	existingPodAffinity := existingPod.Spec.Affinity
	existingHasAffinityConstraints := existingPodAffinity != nil && existingPodAffinity.PodAffinity != nil
	existingHasAntiAffinityConstraints := existingPodAffinity != nil && existingPodAffinity.PodAntiAffinity != nil
	existingPodNode := existingPodNodeInfo.Node()

	// For every soft pod affinity term of <pod>, if <existingPod> matches the term,
	// increment <p.counts> for every node in the cluster with the same <term.TopologyKey>
	// value as that of <existingPods>`s node by the term`s weight.
	pl.processTerms(state, state.affinityTerms, existingPod, existingPodNode, 1)

	// For every soft pod anti-affinity term of <pod>, if <existingPod> matches the term,
	// decrement <p.counts> for every node in the cluster with the same <term.TopologyKey>
	// value as that of <existingPod>`s node by the term`s weight.
	pl.processTerms(state, state.antiAffinityTerms, existingPod, existingPodNode, -1)

	if existingHasAffinityConstraints {
		// For every hard pod affinity term of <existingPod>, if <pod> matches the term,
		// increment <p.counts> for every node in the cluster with the same <term.TopologyKey>
		// value as that of <existingPod>'s node by the constant <ipa.hardPodAffinityWeight>
		if pl.hardPodAffinityWeight > 0 {
			terms := existingPodAffinity.PodAffinity.RequiredDuringSchedulingIgnoredDuringExecution
			// TODO: Uncomment this block when implement RequiredDuringSchedulingRequiredDuringExecution.
			//if len(existingPodAffinity.PodAffinity.RequiredDuringSchedulingRequiredDuringExecution) != 0 {
			//	terms = append(terms, existingPodAffinity.PodAffinity.RequiredDuringSchedulingRequiredDuringExecution...)
			//}
			for i := range terms {
				term := &terms[i]
				processedTerm, err := newWeightedAffinityTerm(existingPod, term, pl.hardPodAffinityWeight)
				if err != nil {
					return err
				}
				pl.processTerm(state, processedTerm, incomingPod, existingPodNode, 1)
			}
		}
		// For every soft pod affinity term of <existingPod>, if <pod> matches the term,
		// increment <p.counts> for every node in the cluster with the same <term.TopologyKey>
		// value as that of <existingPod>'s node by the term's weight.
		terms, err := getProcessedTerms(existingPod, existingPodAffinity.PodAffinity.PreferredDuringSchedulingIgnoredDuringExecution)
		if err != nil {
			klog.Error(err)
			return nil
		}

		pl.processTerms(state, terms, incomingPod, existingPodNode, 1)
	}
	if existingHasAntiAffinityConstraints {
		// For every soft pod anti-affinity term of <existingPod>, if <pod> matches the term,
		// decrement <pm.counts> for every node in the cluster with the same <term.TopologyKey>
		// value as that of <existingPod>'s node by the term's weight.
		terms, err := getProcessedTerms(existingPod, existingPodAffinity.PodAntiAffinity.PreferredDuringSchedulingIgnoredDuringExecution)
		if err != nil {
			return err
		}
		pl.processTerms(state, terms, incomingPod, existingPodNode, -1)
	}
	return nil
}

// PostFilter builds and writes cycle state used by Score and NormalizeScore.
func (pl *InterPodAffinity) PostFilter(
	pCtx context.Context,
	cycleState *framework.CycleState,
	pod *v1.Pod,
	nodes []*v1.Node,
	_ framework.NodeToStatusMap,
) *framework.Status {
	if len(nodes) == 0 {
		// No nodes to score.
		return nil
	}

	if pl.sharedLister == nil {
		return framework.NewStatus(framework.Error, fmt.Sprintf("BuildTopologyPairToScore with empty shared lister"))
	}

	affinity := pod.Spec.Affinity
	hasAffinityConstraints := affinity != nil && affinity.PodAffinity != nil
	hasAntiAffinityConstraints := affinity != nil && affinity.PodAntiAffinity != nil

	// Unless the pod being scheduled has affinity terms, we only
	// need to process nodes hosting pods with affinity.
	allNodes, err := pl.sharedLister.NodeInfos().HavePodsWithAffinityList()
	if err != nil {
		framework.NewStatus(framework.Error, fmt.Sprintf("get pods with affinity list error, err: %v", err))
	}
	if hasAffinityConstraints || hasAntiAffinityConstraints {
		allNodes, err = pl.sharedLister.NodeInfos().List()
		if err != nil {
			framework.NewStatus(framework.Error, fmt.Sprintf("get all nodes from shared lister error, err: %v", err))
		}
	}

	var affinityTerms []*weightedAffinityTerm
	var antiAffinityTerms []*weightedAffinityTerm
	if hasAffinityConstraints {
		if affinityTerms, err = getProcessedTerms(pod, affinity.PodAffinity.PreferredDuringSchedulingIgnoredDuringExecution); err != nil {
			klog.Error(err)
			return nil
		}
	}
	if hasAntiAffinityConstraints {
		if antiAffinityTerms, err = getProcessedTerms(pod, affinity.PodAntiAffinity.PreferredDuringSchedulingIgnoredDuringExecution); err != nil {
			klog.Error(err)
			return nil
		}
	}

	state := &postFilterState{
		topologyScore:     make(map[string]map[string]int64),
		affinityTerms:     affinityTerms,
		antiAffinityTerms: antiAffinityTerms,
	}

	errCh := schedutil.NewErrorChannel()
	ctx, cancel := context.WithCancel(pCtx)
	processNode := func(i int) {
		nodeInfo := allNodes[i]
		if nodeInfo.Node() == nil {
			return
		}
		// Unless the pod being scheduled has affinity terms, we only
		// need to process pods with affinity in the node.
		podsToProcess := nodeInfo.PodsWithAffinity()
		if hasAffinityConstraints || hasAntiAffinityConstraints {
			// We need to process all the pods.
			podsToProcess = nodeInfo.Pods()
		}

		for _, existingPod := range podsToProcess {
			if err := pl.processExistingPod(state, existingPod, nodeInfo, pod); err != nil {
				errCh.SendErrorWithCancel(err, cancel)
				return
			}
		}
	}
	workqueue.ParallelizeUntil(ctx, 16, len(allNodes), processNode)
	if err := errCh.ReceiveError(); err != nil {
		return framework.NewStatus(framework.Error, err.Error())
	}

	cycleState.Write(postFilterStateKey, state)
	return nil
}

func getPostFilterState(cycleState *framework.CycleState) (*postFilterState, error) {
	c, err := cycleState.Read(postFilterStateKey)
	if err != nil {
		return nil, fmt.Errorf("Error reading %q from cycleState: %v", preFilterStateKey, err)
	}

	s, ok := c.(*postFilterState)
	if !ok {
		return nil, fmt.Errorf("%+v  convert to interpodaffinity.postFilterState error", c)
	}
	return s, nil
}

// Score invoked at the Score extension point.
// The "score" returned in this function is the matching number of pods on the `nodeName`,
// it is normalized later.
func (pl *InterPodAffinity) Score(ctx context.Context, cycleState *framework.CycleState, pod *v1.Pod, nodeName string) (int64, *framework.Status) {
	nodeInfo, err := pl.sharedLister.NodeInfos().Get(nodeName)
	if err != nil || nodeInfo.Node() == nil {
		return 0, framework.NewStatus(framework.Error, fmt.Sprintf("getting node %q from Snapshot: %v, node is nil: %v", nodeName, err, nodeInfo.Node() == nil))
	}
	node := nodeInfo.Node()

	s, err := getPostFilterState(cycleState)
	if err != nil {
		return 0, framework.NewStatus(framework.Error, err.Error())
	}
	var score int64
	for tpKey, tpValues := range s.topologyScore {
		if v, exist := node.Labels[tpKey]; exist {
			score += tpValues[v]
		}
	}

	return score, nil
}

// NormalizeScore normalizes the score for each filteredNode.
// The basic rule is: the bigger the score(matching number of pods) is, the smaller the
// final normalized score will be.
func (pl *InterPodAffinity) NormalizeScore(ctx context.Context, cycleState *framework.CycleState, pod *v1.Pod, scores framework.NodeScoreList) *framework.Status {
	s, err := getPostFilterState(cycleState)
	if err != nil {
		return framework.NewStatus(framework.Error, err.Error())
	}
	if len(s.topologyScore) == 0 {
		return nil
	}

	var maxCount, minCount int64
	for i := range scores {
		score := scores[i].Score
		if score > maxCount {
			maxCount = score
		}
		if score < minCount {
			minCount = score
		}
	}

	maxMinDiff := maxCount - minCount
	for i := range scores {
		fScore := float64(0)
		if maxMinDiff > 0 {
			fScore = float64(framework.MaxNodeScore) * (float64(scores[i].Score-minCount) / float64(maxMinDiff))
		}

		scores[i].Score = int64(fScore)
	}

	return nil
}

// ScoreExtensions of the Score plugin.
func (pl *InterPodAffinity) ScoreExtensions() framework.ScoreExtensions {
	return pl
}

// New initializes a new plugin and returns it.
func New(plArgs *runtime.Unknown, h framework.FrameworkHandle) (framework.Plugin, error) {
	if h.SnapshotSharedLister() == nil {
		return nil, fmt.Errorf("SnapshotSharedlister is nil")
	}

	args := &Args{}
	if err := framework.DecodeInto(plArgs, args); err != nil {
		return nil, err
	}

	return &InterPodAffinity{
		sharedLister:          h.SnapshotSharedLister(),
		podAffinityChecker:    predicates.NewPodAffinityChecker(h.SnapshotSharedLister()),
		hardPodAffinityWeight: args.HardPodAffinityWeight,
	}, nil
}
