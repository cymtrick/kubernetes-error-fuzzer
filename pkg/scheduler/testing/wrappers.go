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

package testing

import (
	"fmt"
	"time"

	v1 "k8s.io/api/core/v1"
	resourcev1alpha2 "k8s.io/api/resource/v1alpha2"
	"k8s.io/apimachinery/pkg/api/resource"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/types"
	imageutils "k8s.io/kubernetes/test/utils/image"
	"k8s.io/utils/ptr"
)

var zero int64

// NodeSelectorWrapper wraps a NodeSelector inside.
type NodeSelectorWrapper struct{ v1.NodeSelector }

// MakeNodeSelector creates a NodeSelector wrapper.
func MakeNodeSelector() *NodeSelectorWrapper {
	return &NodeSelectorWrapper{v1.NodeSelector{}}
}

// In injects a matchExpression (with an operator IN) as a selectorTerm
// to the inner nodeSelector.
// NOTE: appended selecterTerms are ORed.
func (s *NodeSelectorWrapper) In(key string, vals []string) *NodeSelectorWrapper {
	expression := v1.NodeSelectorRequirement{
		Key:      key,
		Operator: v1.NodeSelectorOpIn,
		Values:   vals,
	}
	selectorTerm := v1.NodeSelectorTerm{}
	selectorTerm.MatchExpressions = append(selectorTerm.MatchExpressions, expression)
	s.NodeSelectorTerms = append(s.NodeSelectorTerms, selectorTerm)
	return s
}

// NotIn injects a matchExpression (with an operator NotIn) as a selectorTerm
// to the inner nodeSelector.
func (s *NodeSelectorWrapper) NotIn(key string, vals []string) *NodeSelectorWrapper {
	expression := v1.NodeSelectorRequirement{
		Key:      key,
		Operator: v1.NodeSelectorOpNotIn,
		Values:   vals,
	}
	selectorTerm := v1.NodeSelectorTerm{}
	selectorTerm.MatchExpressions = append(selectorTerm.MatchExpressions, expression)
	s.NodeSelectorTerms = append(s.NodeSelectorTerms, selectorTerm)
	return s
}

// Obj returns the inner NodeSelector.
func (s *NodeSelectorWrapper) Obj() *v1.NodeSelector {
	return &s.NodeSelector
}

// LabelSelectorWrapper wraps a LabelSelector inside.
type LabelSelectorWrapper struct{ metav1.LabelSelector }

// MakeLabelSelector creates a LabelSelector wrapper.
func MakeLabelSelector() *LabelSelectorWrapper {
	return &LabelSelectorWrapper{metav1.LabelSelector{}}
}

// Label applies a {k,v} pair to the inner LabelSelector.
func (s *LabelSelectorWrapper) Label(k, v string) *LabelSelectorWrapper {
	if s.MatchLabels == nil {
		s.MatchLabels = make(map[string]string)
	}
	s.MatchLabels[k] = v
	return s
}

// In injects a matchExpression (with an operator In) to the inner labelSelector.
func (s *LabelSelectorWrapper) In(key string, vals []string) *LabelSelectorWrapper {
	expression := metav1.LabelSelectorRequirement{
		Key:      key,
		Operator: metav1.LabelSelectorOpIn,
		Values:   vals,
	}
	s.MatchExpressions = append(s.MatchExpressions, expression)
	return s
}

// NotIn injects a matchExpression (with an operator NotIn) to the inner labelSelector.
func (s *LabelSelectorWrapper) NotIn(key string, vals []string) *LabelSelectorWrapper {
	expression := metav1.LabelSelectorRequirement{
		Key:      key,
		Operator: metav1.LabelSelectorOpNotIn,
		Values:   vals,
	}
	s.MatchExpressions = append(s.MatchExpressions, expression)
	return s
}

// Exists injects a matchExpression (with an operator Exists) to the inner labelSelector.
func (s *LabelSelectorWrapper) Exists(k string) *LabelSelectorWrapper {
	expression := metav1.LabelSelectorRequirement{
		Key:      k,
		Operator: metav1.LabelSelectorOpExists,
	}
	s.MatchExpressions = append(s.MatchExpressions, expression)
	return s
}

// NotExist injects a matchExpression (with an operator NotExist) to the inner labelSelector.
func (s *LabelSelectorWrapper) NotExist(k string) *LabelSelectorWrapper {
	expression := metav1.LabelSelectorRequirement{
		Key:      k,
		Operator: metav1.LabelSelectorOpDoesNotExist,
	}
	s.MatchExpressions = append(s.MatchExpressions, expression)
	return s
}

// Obj returns the inner LabelSelector.
func (s *LabelSelectorWrapper) Obj() *metav1.LabelSelector {
	return &s.LabelSelector
}

// ContainerWrapper wraps a Container inside.
type ContainerWrapper struct{ v1.Container }

// MakeContainer creates a Container wrapper.
func MakeContainer() *ContainerWrapper {
	return &ContainerWrapper{v1.Container{}}
}

// Obj returns the inner Container.
func (c *ContainerWrapper) Obj() v1.Container {
	return c.Container
}

// Name sets `n` as the name of the inner Container.
func (c *ContainerWrapper) Name(n string) *ContainerWrapper {
	c.Container.Name = n
	return c
}

// Image sets `image` as the image of the inner Container.
func (c *ContainerWrapper) Image(image string) *ContainerWrapper {
	c.Container.Image = image
	return c
}

// HostPort sets `hostPort` as the host port of the inner Container.
func (c *ContainerWrapper) HostPort(hostPort int32) *ContainerWrapper {
	c.Container.Ports = []v1.ContainerPort{{HostPort: hostPort}}
	return c
}

// ContainerPort sets `ports` as the ports of the inner Container.
func (c *ContainerWrapper) ContainerPort(ports []v1.ContainerPort) *ContainerWrapper {
	c.Container.Ports = ports
	return c
}

// Resources sets the container resources to the given resource map.
func (c *ContainerWrapper) Resources(resMap map[v1.ResourceName]string) *ContainerWrapper {
	res := v1.ResourceList{}
	for k, v := range resMap {
		res[k] = resource.MustParse(v)
	}
	c.Container.Resources = v1.ResourceRequirements{
		Requests: res,
		Limits:   res,
	}
	return c
}

// ResourceRequests sets the container resources requests to the given resource map of requests.
func (c *ContainerWrapper) ResourceRequests(reqMap map[v1.ResourceName]string) *ContainerWrapper {
	res := v1.ResourceList{}
	for k, v := range reqMap {
		res[k] = resource.MustParse(v)
	}
	c.Container.Resources = v1.ResourceRequirements{
		Requests: res,
	}
	return c
}

// ResourceLimits sets the container resource limits to the given resource map.
func (c *ContainerWrapper) ResourceLimits(limMap map[v1.ResourceName]string) *ContainerWrapper {
	res := v1.ResourceList{}
	for k, v := range limMap {
		res[k] = resource.MustParse(v)
	}
	c.Container.Resources = v1.ResourceRequirements{
		Limits: res,
	}
	return c
}

// PodWrapper wraps a Pod inside.
type PodWrapper struct{ v1.Pod }

// MakePod creates a Pod wrapper.
func MakePod() *PodWrapper {
	return &PodWrapper{v1.Pod{}}
}

// Obj returns the inner Pod.
func (p *PodWrapper) Obj() *v1.Pod {
	return &p.Pod
}

// Name sets `s` as the name of the inner pod.
func (p *PodWrapper) Name(s string) *PodWrapper {
	p.SetName(s)
	return p
}

// UID sets `s` as the UID of the inner pod.
func (p *PodWrapper) UID(s string) *PodWrapper {
	p.SetUID(types.UID(s))
	return p
}

// SchedulerName sets `s` as the scheduler name of the inner pod.
func (p *PodWrapper) SchedulerName(s string) *PodWrapper {
	p.Spec.SchedulerName = s
	return p
}

// Namespace sets `s` as the namespace of the inner pod.
func (p *PodWrapper) Namespace(s string) *PodWrapper {
	p.SetNamespace(s)
	return p
}

// OwnerReference updates the owning controller of the pod.
func (p *PodWrapper) OwnerReference(name string, gvk schema.GroupVersionKind) *PodWrapper {
	p.OwnerReferences = []metav1.OwnerReference{
		{
			APIVersion: gvk.GroupVersion().String(),
			Kind:       gvk.Kind,
			Name:       name,
			Controller: ptr.To(true),
		},
	}
	return p
}

// Container appends a container into PodSpec of the inner pod.
func (p *PodWrapper) Container(s string) *PodWrapper {
	name := fmt.Sprintf("con%d", len(p.Spec.Containers))
	p.Spec.Containers = append(p.Spec.Containers, MakeContainer().Name(name).Image(s).Obj())
	return p
}

// Containers sets `containers` to the PodSpec of the inner pod.
func (p *PodWrapper) Containers(containers []v1.Container) *PodWrapper {
	p.Spec.Containers = containers
	return p
}

// PodResourceClaims appends PodResourceClaims into PodSpec of the inner pod.
func (p *PodWrapper) PodResourceClaims(podResourceClaims ...v1.PodResourceClaim) *PodWrapper {
	p.Spec.ResourceClaims = append(p.Spec.ResourceClaims, podResourceClaims...)
	return p
}

// Priority sets a priority value into PodSpec of the inner pod.
func (p *PodWrapper) Priority(val int32) *PodWrapper {
	p.Spec.Priority = &val
	return p
}

// CreationTimestamp sets the inner pod's CreationTimestamp.
func (p *PodWrapper) CreationTimestamp(t metav1.Time) *PodWrapper {
	p.ObjectMeta.CreationTimestamp = t
	return p
}

// Terminating sets the inner pod's deletionTimestamp to current timestamp.
func (p *PodWrapper) Terminating() *PodWrapper {
	now := metav1.Now()
	p.DeletionTimestamp = &now
	return p
}

// ZeroTerminationGracePeriod sets the TerminationGracePeriodSeconds of the inner pod to zero.
func (p *PodWrapper) ZeroTerminationGracePeriod() *PodWrapper {
	p.Spec.TerminationGracePeriodSeconds = &zero
	return p
}

// Node sets `s` as the nodeName of the inner pod.
func (p *PodWrapper) Node(s string) *PodWrapper {
	p.Spec.NodeName = s
	return p
}

// NodeSelector sets `m` as the nodeSelector of the inner pod.
func (p *PodWrapper) NodeSelector(m map[string]string) *PodWrapper {
	p.Spec.NodeSelector = m
	return p
}

// NodeAffinityIn creates a HARD node affinity (with the operator In)
// and injects into the inner pod.
func (p *PodWrapper) NodeAffinityIn(key string, vals []string) *PodWrapper {
	if p.Spec.Affinity == nil {
		p.Spec.Affinity = &v1.Affinity{}
	}
	if p.Spec.Affinity.NodeAffinity == nil {
		p.Spec.Affinity.NodeAffinity = &v1.NodeAffinity{}
	}
	nodeSelector := MakeNodeSelector().In(key, vals).Obj()
	p.Spec.Affinity.NodeAffinity.RequiredDuringSchedulingIgnoredDuringExecution = nodeSelector
	return p
}

// NodeAffinityNotIn creates a HARD node affinity (with the operator NotIn)
// and injects into the inner pod.
func (p *PodWrapper) NodeAffinityNotIn(key string, vals []string) *PodWrapper {
	if p.Spec.Affinity == nil {
		p.Spec.Affinity = &v1.Affinity{}
	}
	if p.Spec.Affinity.NodeAffinity == nil {
		p.Spec.Affinity.NodeAffinity = &v1.NodeAffinity{}
	}
	nodeSelector := MakeNodeSelector().NotIn(key, vals).Obj()
	p.Spec.Affinity.NodeAffinity.RequiredDuringSchedulingIgnoredDuringExecution = nodeSelector
	return p
}

// StartTime sets `t` as .status.startTime for the inner pod.
func (p *PodWrapper) StartTime(t metav1.Time) *PodWrapper {
	p.Status.StartTime = &t
	return p
}

// NominatedNodeName sets `n` as the .Status.NominatedNodeName of the inner pod.
func (p *PodWrapper) NominatedNodeName(n string) *PodWrapper {
	p.Status.NominatedNodeName = n
	return p
}

// Phase sets `phase` as .status.Phase of the inner pod.
func (p *PodWrapper) Phase(phase v1.PodPhase) *PodWrapper {
	p.Status.Phase = phase
	return p
}

// Condition adds a `condition(Type, Status, Reason)` to .Status.Conditions.
func (p *PodWrapper) Condition(t v1.PodConditionType, s v1.ConditionStatus, r string) *PodWrapper {
	p.Status.Conditions = append(p.Status.Conditions, v1.PodCondition{Type: t, Status: s, Reason: r})
	return p
}

// Conditions sets `conditions` as .status.Conditions of the inner pod.
func (p *PodWrapper) Conditions(conditions []v1.PodCondition) *PodWrapper {
	p.Status.Conditions = append(p.Status.Conditions, conditions...)
	return p
}

// Toleration creates a toleration (with the operator Exists)
// and injects into the inner pod.
func (p *PodWrapper) Toleration(key string) *PodWrapper {
	p.Spec.Tolerations = append(p.Spec.Tolerations, v1.Toleration{
		Key:      key,
		Operator: v1.TolerationOpExists,
	})
	return p
}

// HostPort creates a container with a hostPort valued `hostPort`,
// and injects into the inner pod.
func (p *PodWrapper) HostPort(port int32) *PodWrapper {
	p.Spec.Containers = append(p.Spec.Containers, MakeContainer().Name("container").Image("pause").HostPort(port).Obj())
	return p
}

// ContainerPort creates a container with ports valued `ports`,
// and injects into the inner pod.
func (p *PodWrapper) ContainerPort(ports []v1.ContainerPort) *PodWrapper {
	p.Spec.Containers = append(p.Spec.Containers, MakeContainer().Name("container").Image("pause").ContainerPort(ports).Obj())
	return p
}

// PVC creates a Volume with a PVC and injects into the inner pod.
func (p *PodWrapper) PVC(name string) *PodWrapper {
	p.Spec.Volumes = append(p.Spec.Volumes, v1.Volume{
		Name: name,
		VolumeSource: v1.VolumeSource{
			PersistentVolumeClaim: &v1.PersistentVolumeClaimVolumeSource{ClaimName: name},
		},
	})
	return p
}

// Volume creates volume and injects into the inner pod.
func (p *PodWrapper) Volume(volume v1.Volume) *PodWrapper {
	p.Spec.Volumes = append(p.Spec.Volumes, volume)
	return p
}

// Volumes set the volumes and inject into the inner pod.
func (p *PodWrapper) Volumes(volumes []v1.Volume) *PodWrapper {
	p.Spec.Volumes = volumes
	return p
}

// SchedulingGates sets `gates` as additional SchedulerGates of the inner pod.
func (p *PodWrapper) SchedulingGates(gates []string) *PodWrapper {
	for _, gate := range gates {
		p.Spec.SchedulingGates = append(p.Spec.SchedulingGates, v1.PodSchedulingGate{Name: gate})
	}
	return p
}

// PodAffinityKind represents different kinds of PodAffinity.
type PodAffinityKind int

const (
	// NilPodAffinity is a no-op which doesn't apply any PodAffinity.
	NilPodAffinity PodAffinityKind = iota
	// PodAffinityWithRequiredReq applies a HARD requirement to pod.spec.affinity.PodAffinity.
	PodAffinityWithRequiredReq
	// PodAffinityWithPreferredReq applies a SOFT requirement to pod.spec.affinity.PodAffinity.
	PodAffinityWithPreferredReq
	// PodAffinityWithRequiredPreferredReq applies HARD and SOFT requirements to pod.spec.affinity.PodAffinity.
	PodAffinityWithRequiredPreferredReq
	// PodAntiAffinityWithRequiredReq applies a HARD requirement to pod.spec.affinity.PodAntiAffinity.
	PodAntiAffinityWithRequiredReq
	// PodAntiAffinityWithPreferredReq applies a SOFT requirement to pod.spec.affinity.PodAntiAffinity.
	PodAntiAffinityWithPreferredReq
	// PodAntiAffinityWithRequiredPreferredReq applies HARD and SOFT requirements to pod.spec.affinity.PodAntiAffinity.
	PodAntiAffinityWithRequiredPreferredReq
)

// PodAffinity creates a PodAffinity with topology key and label selector
// and injects into the inner pod.
func (p *PodWrapper) PodAffinity(topologyKey string, labelSelector *metav1.LabelSelector, kind PodAffinityKind) *PodWrapper {
	if kind == NilPodAffinity {
		return p
	}

	if p.Spec.Affinity == nil {
		p.Spec.Affinity = &v1.Affinity{}
	}
	if p.Spec.Affinity.PodAffinity == nil {
		p.Spec.Affinity.PodAffinity = &v1.PodAffinity{}
	}
	term := v1.PodAffinityTerm{LabelSelector: labelSelector, TopologyKey: topologyKey}
	switch kind {
	case PodAffinityWithRequiredReq:
		p.Spec.Affinity.PodAffinity.RequiredDuringSchedulingIgnoredDuringExecution = append(
			p.Spec.Affinity.PodAffinity.RequiredDuringSchedulingIgnoredDuringExecution,
			term,
		)
	case PodAffinityWithPreferredReq:
		p.Spec.Affinity.PodAffinity.PreferredDuringSchedulingIgnoredDuringExecution = append(
			p.Spec.Affinity.PodAffinity.PreferredDuringSchedulingIgnoredDuringExecution,
			v1.WeightedPodAffinityTerm{Weight: 1, PodAffinityTerm: term},
		)
	case PodAffinityWithRequiredPreferredReq:
		p.Spec.Affinity.PodAffinity.RequiredDuringSchedulingIgnoredDuringExecution = append(
			p.Spec.Affinity.PodAffinity.RequiredDuringSchedulingIgnoredDuringExecution,
			term,
		)
		p.Spec.Affinity.PodAffinity.PreferredDuringSchedulingIgnoredDuringExecution = append(
			p.Spec.Affinity.PodAffinity.PreferredDuringSchedulingIgnoredDuringExecution,
			v1.WeightedPodAffinityTerm{Weight: 1, PodAffinityTerm: term},
		)
	}
	return p
}

// PodAntiAffinity creates a PodAntiAffinity with topology key and label selector
// and injects into the inner pod.
func (p *PodWrapper) PodAntiAffinity(topologyKey string, labelSelector *metav1.LabelSelector, kind PodAffinityKind) *PodWrapper {
	if kind == NilPodAffinity {
		return p
	}

	if p.Spec.Affinity == nil {
		p.Spec.Affinity = &v1.Affinity{}
	}
	if p.Spec.Affinity.PodAntiAffinity == nil {
		p.Spec.Affinity.PodAntiAffinity = &v1.PodAntiAffinity{}
	}
	term := v1.PodAffinityTerm{LabelSelector: labelSelector, TopologyKey: topologyKey}
	switch kind {
	case PodAntiAffinityWithRequiredReq:
		p.Spec.Affinity.PodAntiAffinity.RequiredDuringSchedulingIgnoredDuringExecution = append(
			p.Spec.Affinity.PodAntiAffinity.RequiredDuringSchedulingIgnoredDuringExecution,
			term,
		)
	case PodAntiAffinityWithPreferredReq:
		p.Spec.Affinity.PodAntiAffinity.PreferredDuringSchedulingIgnoredDuringExecution = append(
			p.Spec.Affinity.PodAntiAffinity.PreferredDuringSchedulingIgnoredDuringExecution,
			v1.WeightedPodAffinityTerm{Weight: 1, PodAffinityTerm: term},
		)
	case PodAntiAffinityWithRequiredPreferredReq:
		p.Spec.Affinity.PodAntiAffinity.RequiredDuringSchedulingIgnoredDuringExecution = append(
			p.Spec.Affinity.PodAntiAffinity.RequiredDuringSchedulingIgnoredDuringExecution,
			term,
		)
		p.Spec.Affinity.PodAntiAffinity.PreferredDuringSchedulingIgnoredDuringExecution = append(
			p.Spec.Affinity.PodAntiAffinity.PreferredDuringSchedulingIgnoredDuringExecution,
			v1.WeightedPodAffinityTerm{Weight: 1, PodAffinityTerm: term},
		)
	}
	return p
}

// PodAffinityExists creates a PodAffinity with the operator "Exists"
// and injects into the inner pod.
func (p *PodWrapper) PodAffinityExists(labelKey, topologyKey string, kind PodAffinityKind) *PodWrapper {
	labelSelector := MakeLabelSelector().Exists(labelKey).Obj()
	p.PodAffinity(topologyKey, labelSelector, kind)
	return p
}

// PodAntiAffinityExists creates a PodAntiAffinity with the operator "Exists"
// and injects into the inner pod.
func (p *PodWrapper) PodAntiAffinityExists(labelKey, topologyKey string, kind PodAffinityKind) *PodWrapper {
	labelSelector := MakeLabelSelector().Exists(labelKey).Obj()
	p.PodAntiAffinity(topologyKey, labelSelector, kind)
	return p
}

// PodAffinityNotExists creates a PodAffinity with the operator "NotExists"
// and injects into the inner pod.
func (p *PodWrapper) PodAffinityNotExists(labelKey, topologyKey string, kind PodAffinityKind) *PodWrapper {
	labelSelector := MakeLabelSelector().NotExist(labelKey).Obj()
	p.PodAffinity(topologyKey, labelSelector, kind)
	return p
}

// PodAntiAffinityNotExists creates a PodAntiAffinity with the operator "NotExists"
// and injects into the inner pod.
func (p *PodWrapper) PodAntiAffinityNotExists(labelKey, topologyKey string, kind PodAffinityKind) *PodWrapper {
	labelSelector := MakeLabelSelector().NotExist(labelKey).Obj()
	p.PodAntiAffinity(topologyKey, labelSelector, kind)
	return p
}

// PodAffinityIn creates a PodAffinity with the operator "In"
// and injects into the inner pod.
func (p *PodWrapper) PodAffinityIn(labelKey, topologyKey string, vals []string, kind PodAffinityKind) *PodWrapper {
	labelSelector := MakeLabelSelector().In(labelKey, vals).Obj()
	p.PodAffinity(topologyKey, labelSelector, kind)
	return p
}

// PodAntiAffinityIn creates a PodAntiAffinity with the operator "In"
// and injects into the inner pod.
func (p *PodWrapper) PodAntiAffinityIn(labelKey, topologyKey string, vals []string, kind PodAffinityKind) *PodWrapper {
	labelSelector := MakeLabelSelector().In(labelKey, vals).Obj()
	p.PodAntiAffinity(topologyKey, labelSelector, kind)
	return p
}

// PodAffinityNotIn creates a PodAffinity with the operator "NotIn"
// and injects into the inner pod.
func (p *PodWrapper) PodAffinityNotIn(labelKey, topologyKey string, vals []string, kind PodAffinityKind) *PodWrapper {
	labelSelector := MakeLabelSelector().NotIn(labelKey, vals).Obj()
	p.PodAffinity(topologyKey, labelSelector, kind)
	return p
}

// PodAntiAffinityNotIn creates a PodAntiAffinity with the operator "NotIn"
// and injects into the inner pod.
func (p *PodWrapper) PodAntiAffinityNotIn(labelKey, topologyKey string, vals []string, kind PodAffinityKind) *PodWrapper {
	labelSelector := MakeLabelSelector().NotIn(labelKey, vals).Obj()
	p.PodAntiAffinity(topologyKey, labelSelector, kind)
	return p
}

// SpreadConstraint constructs a TopologySpreadConstraint object and injects
// into the inner pod.
func (p *PodWrapper) SpreadConstraint(maxSkew int, tpKey string, mode v1.UnsatisfiableConstraintAction, selector *metav1.LabelSelector, minDomains *int32, nodeAffinityPolicy, nodeTaintsPolicy *v1.NodeInclusionPolicy, matchLabelKeys []string) *PodWrapper {
	c := v1.TopologySpreadConstraint{
		MaxSkew:            int32(maxSkew),
		TopologyKey:        tpKey,
		WhenUnsatisfiable:  mode,
		LabelSelector:      selector,
		MinDomains:         minDomains,
		NodeAffinityPolicy: nodeAffinityPolicy,
		NodeTaintsPolicy:   nodeTaintsPolicy,
		MatchLabelKeys:     matchLabelKeys,
	}
	p.Spec.TopologySpreadConstraints = append(p.Spec.TopologySpreadConstraints, c)
	return p
}

// Label sets a {k,v} pair to the inner pod label.
func (p *PodWrapper) Label(k, v string) *PodWrapper {
	if p.ObjectMeta.Labels == nil {
		p.ObjectMeta.Labels = make(map[string]string)
	}
	p.ObjectMeta.Labels[k] = v
	return p
}

// Labels sets all {k,v} pair provided by `labels` to the inner pod labels.
func (p *PodWrapper) Labels(labels map[string]string) *PodWrapper {
	for k, v := range labels {
		p.Label(k, v)
	}
	return p
}

// Annotation sets a {k,v} pair to the inner pod annotation.
func (p *PodWrapper) Annotation(key, value string) *PodWrapper {
	metav1.SetMetaDataAnnotation(&p.ObjectMeta, key, value)
	return p
}

// Annotations sets all {k,v} pair provided by `annotations` to the inner pod annotations.
func (p *PodWrapper) Annotations(annotations map[string]string) *PodWrapper {
	for k, v := range annotations {
		p.Annotation(k, v)
	}
	return p
}

// Res adds a new container to the inner pod with given resource map.
func (p *PodWrapper) Res(resMap map[v1.ResourceName]string) *PodWrapper {
	if len(resMap) == 0 {
		return p
	}

	name := fmt.Sprintf("con%d", len(p.Spec.Containers))
	p.Spec.Containers = append(p.Spec.Containers, MakeContainer().Name(name).Image(imageutils.GetPauseImageName()).Resources(resMap).Obj())
	return p
}

// Req adds a new container to the inner pod with given resource map of requests.
func (p *PodWrapper) Req(reqMap map[v1.ResourceName]string) *PodWrapper {
	if len(reqMap) == 0 {
		return p
	}

	name := fmt.Sprintf("con%d", len(p.Spec.Containers))
	p.Spec.Containers = append(p.Spec.Containers, MakeContainer().Name(name).Image(imageutils.GetPauseImageName()).ResourceRequests(reqMap).Obj())
	return p
}

// Lim adds a new container to the inner pod with given resource map of limits.
func (p *PodWrapper) Lim(limMap map[v1.ResourceName]string) *PodWrapper {
	if len(limMap) == 0 {
		return p
	}

	name := fmt.Sprintf("con%d", len(p.Spec.Containers))
	p.Spec.Containers = append(p.Spec.Containers, MakeContainer().Name(name).Image(imageutils.GetPauseImageName()).ResourceLimits(limMap).Obj())
	return p
}

// InitReq adds a new init container to the inner pod with given resource map.
func (p *PodWrapper) InitReq(resMap map[v1.ResourceName]string) *PodWrapper {
	if len(resMap) == 0 {
		return p
	}

	name := fmt.Sprintf("init-con%d", len(p.Spec.InitContainers))
	p.Spec.InitContainers = append(p.Spec.InitContainers, MakeContainer().Name(name).Image(imageutils.GetPauseImageName()).Resources(resMap).Obj())
	return p
}

// PreemptionPolicy sets the give preemption policy to the inner pod.
func (p *PodWrapper) PreemptionPolicy(policy v1.PreemptionPolicy) *PodWrapper {
	p.Spec.PreemptionPolicy = &policy
	return p
}

// Overhead sets the give ResourceList to the inner pod
func (p *PodWrapper) Overhead(rl v1.ResourceList) *PodWrapper {
	p.Spec.Overhead = rl
	return p
}

// NodeWrapper wraps a Node inside.
type NodeWrapper struct{ v1.Node }

// MakeNode creates a Node wrapper.
func MakeNode() *NodeWrapper {
	w := &NodeWrapper{v1.Node{}}
	return w.Capacity(nil)
}

// Obj returns the inner Node.
func (n *NodeWrapper) Obj() *v1.Node {
	return &n.Node
}

// Name sets `s` as the name of the inner pod.
func (n *NodeWrapper) Name(s string) *NodeWrapper {
	n.SetName(s)
	return n
}

// UID sets `s` as the UID of the inner pod.
func (n *NodeWrapper) UID(s string) *NodeWrapper {
	n.SetUID(types.UID(s))
	return n
}

// Label applies a {k,v} label pair to the inner node.
func (n *NodeWrapper) Label(k, v string) *NodeWrapper {
	if n.Labels == nil {
		n.Labels = make(map[string]string)
	}
	n.Labels[k] = v
	return n
}

// Capacity sets the capacity and the allocatable resources of the inner node.
// Each entry in `resources` corresponds to a resource name and its quantity.
// By default, the capacity and allocatable number of pods are set to 32.
func (n *NodeWrapper) Capacity(resources map[v1.ResourceName]string) *NodeWrapper {
	res := v1.ResourceList{
		v1.ResourcePods: resource.MustParse("32"),
	}
	for name, value := range resources {
		res[name] = resource.MustParse(value)
	}
	n.Status.Capacity, n.Status.Allocatable = res, res
	return n
}

// Images sets the images of the inner node. Each entry in `images` corresponds
// to an image name and its size in bytes.
func (n *NodeWrapper) Images(images map[string]int64) *NodeWrapper {
	var containerImages []v1.ContainerImage
	for name, size := range images {
		containerImages = append(containerImages, v1.ContainerImage{Names: []string{name}, SizeBytes: size})
	}
	n.Status.Images = containerImages
	return n
}

// Taints applies taints to the inner node.
func (n *NodeWrapper) Taints(taints []v1.Taint) *NodeWrapper {
	n.Spec.Taints = taints
	return n
}

// Unschedulable applies the unschedulable field.
func (n *NodeWrapper) Unschedulable(unschedulable bool) *NodeWrapper {
	n.Spec.Unschedulable = unschedulable
	return n
}

// Condition applies the node condition.
func (n *NodeWrapper) Condition(typ v1.NodeConditionType, status v1.ConditionStatus, message, reason string) *NodeWrapper {
	n.Status.Conditions = []v1.NodeCondition{
		{
			Type:               typ,
			Status:             status,
			Message:            message,
			Reason:             reason,
			LastHeartbeatTime:  metav1.Time{Time: time.Now()},
			LastTransitionTime: metav1.Time{Time: time.Now()},
		},
	}
	return n
}

// PersistentVolumeClaimWrapper wraps a PersistentVolumeClaim inside.
type PersistentVolumeClaimWrapper struct{ v1.PersistentVolumeClaim }

// MakePersistentVolumeClaim creates a PersistentVolumeClaim wrapper.
func MakePersistentVolumeClaim() *PersistentVolumeClaimWrapper {
	return &PersistentVolumeClaimWrapper{}
}

// Obj returns the inner PersistentVolumeClaim.
func (p *PersistentVolumeClaimWrapper) Obj() *v1.PersistentVolumeClaim {
	return &p.PersistentVolumeClaim
}

// Name sets `s` as the name of the inner PersistentVolumeClaim.
func (p *PersistentVolumeClaimWrapper) Name(s string) *PersistentVolumeClaimWrapper {
	p.SetName(s)
	return p
}

// Namespace sets `s` as the namespace of the inner PersistentVolumeClaim.
func (p *PersistentVolumeClaimWrapper) Namespace(s string) *PersistentVolumeClaimWrapper {
	p.SetNamespace(s)
	return p
}

// Annotation sets a {k,v} pair to the inner PersistentVolumeClaim.
func (p *PersistentVolumeClaimWrapper) Annotation(key, value string) *PersistentVolumeClaimWrapper {
	metav1.SetMetaDataAnnotation(&p.ObjectMeta, key, value)
	return p
}

// VolumeName sets `name` as the volume name of the inner
// PersistentVolumeClaim.
func (p *PersistentVolumeClaimWrapper) VolumeName(name string) *PersistentVolumeClaimWrapper {
	p.PersistentVolumeClaim.Spec.VolumeName = name
	return p
}

// AccessModes sets `accessModes` as the access modes of the inner
// PersistentVolumeClaim.
func (p *PersistentVolumeClaimWrapper) AccessModes(accessModes []v1.PersistentVolumeAccessMode) *PersistentVolumeClaimWrapper {
	p.PersistentVolumeClaim.Spec.AccessModes = accessModes
	return p
}

// Resources sets `resources` as the resource requirements of the inner
// PersistentVolumeClaim.
func (p *PersistentVolumeClaimWrapper) Resources(resources v1.VolumeResourceRequirements) *PersistentVolumeClaimWrapper {
	p.PersistentVolumeClaim.Spec.Resources = resources
	return p
}

// PersistentVolumeWrapper wraps a PersistentVolume inside.
type PersistentVolumeWrapper struct{ v1.PersistentVolume }

// MakePersistentVolume creates a PersistentVolume wrapper.
func MakePersistentVolume() *PersistentVolumeWrapper {
	return &PersistentVolumeWrapper{}
}

// Obj returns the inner PersistentVolume.
func (p *PersistentVolumeWrapper) Obj() *v1.PersistentVolume {
	return &p.PersistentVolume
}

// Name sets `s` as the name of the inner PersistentVolume.
func (p *PersistentVolumeWrapper) Name(s string) *PersistentVolumeWrapper {
	p.SetName(s)
	return p
}

// AccessModes sets `accessModes` as the access modes of the inner
// PersistentVolume.
func (p *PersistentVolumeWrapper) AccessModes(accessModes []v1.PersistentVolumeAccessMode) *PersistentVolumeWrapper {
	p.PersistentVolume.Spec.AccessModes = accessModes
	return p
}

// Capacity sets `capacity` as the resource list of the inner PersistentVolume.
func (p *PersistentVolumeWrapper) Capacity(capacity v1.ResourceList) *PersistentVolumeWrapper {
	p.PersistentVolume.Spec.Capacity = capacity
	return p
}

// HostPathVolumeSource sets `src` as the host path volume source of the inner
// PersistentVolume.
func (p *PersistentVolumeWrapper) HostPathVolumeSource(src *v1.HostPathVolumeSource) *PersistentVolumeWrapper {
	p.PersistentVolume.Spec.HostPath = src
	return p
}

// ResourceClaimWrapper wraps a ResourceClaim inside.
type ResourceClaimWrapper struct{ resourcev1alpha2.ResourceClaim }

// MakeResourceClaim creates a ResourceClaim wrapper.
func MakeResourceClaim() *ResourceClaimWrapper {
	return &ResourceClaimWrapper{resourcev1alpha2.ResourceClaim{}}
}

// FromResourceClaim creates a ResourceClaim wrapper from some existing object.
func FromResourceClaim(other *resourcev1alpha2.ResourceClaim) *ResourceClaimWrapper {
	return &ResourceClaimWrapper{*other.DeepCopy()}
}

// Obj returns the inner ResourceClaim.
func (wrapper *ResourceClaimWrapper) Obj() *resourcev1alpha2.ResourceClaim {
	return &wrapper.ResourceClaim
}

// Name sets `s` as the name of the inner object.
func (wrapper *ResourceClaimWrapper) Name(s string) *ResourceClaimWrapper {
	wrapper.SetName(s)
	return wrapper
}

// UID sets `s` as the UID of the inner object.
func (wrapper *ResourceClaimWrapper) UID(s string) *ResourceClaimWrapper {
	wrapper.SetUID(types.UID(s))
	return wrapper
}

// Namespace sets `s` as the namespace of the inner object.
func (wrapper *ResourceClaimWrapper) Namespace(s string) *ResourceClaimWrapper {
	wrapper.SetNamespace(s)
	return wrapper
}

// OwnerReference updates the owning controller of the object.
func (wrapper *ResourceClaimWrapper) OwnerReference(name, uid string, gvk schema.GroupVersionKind) *ResourceClaimWrapper {
	wrapper.OwnerReferences = []metav1.OwnerReference{
		{
			APIVersion: gvk.GroupVersion().String(),
			Kind:       gvk.Kind,
			Name:       name,
			UID:        types.UID(uid),
			Controller: ptr.To(true),
		},
	}
	return wrapper
}

// AllocationMode sets the allocation mode of the inner object.
func (wrapper *ResourceClaimWrapper) AllocationMode(a resourcev1alpha2.AllocationMode) *ResourceClaimWrapper {
	wrapper.ResourceClaim.Spec.AllocationMode = a
	return wrapper
}

// ResourceClassName sets the resource class name of the inner object.
func (wrapper *ResourceClaimWrapper) ResourceClassName(name string) *ResourceClaimWrapper {
	wrapper.ResourceClaim.Spec.ResourceClassName = name
	return wrapper
}

// Allocation sets the allocation of the inner object.
func (wrapper *ResourceClaimWrapper) Allocation(allocation *resourcev1alpha2.AllocationResult) *ResourceClaimWrapper {
	wrapper.ResourceClaim.Status.Allocation = allocation
	return wrapper
}

// DeallocationRequested sets that field of the inner object.
func (wrapper *ResourceClaimWrapper) DeallocationRequested(deallocationRequested bool) *ResourceClaimWrapper {
	wrapper.ResourceClaim.Status.DeallocationRequested = deallocationRequested
	return wrapper
}

// ReservedFor sets that field of the inner object.
func (wrapper *ResourceClaimWrapper) ReservedFor(consumers ...resourcev1alpha2.ResourceClaimConsumerReference) *ResourceClaimWrapper {
	wrapper.ResourceClaim.Status.ReservedFor = consumers
	return wrapper
}

// PodSchedulingWrapper wraps a PodSchedulingContext inside.
type PodSchedulingWrapper struct {
	resourcev1alpha2.PodSchedulingContext
}

// MakePodSchedulingContexts creates a PodSchedulingContext wrapper.
func MakePodSchedulingContexts() *PodSchedulingWrapper {
	return &PodSchedulingWrapper{resourcev1alpha2.PodSchedulingContext{}}
}

// FromPodSchedulingContexts creates a PodSchedulingContext wrapper from an existing object.
func FromPodSchedulingContexts(other *resourcev1alpha2.PodSchedulingContext) *PodSchedulingWrapper {
	return &PodSchedulingWrapper{*other.DeepCopy()}
}

// Obj returns the inner object.
func (wrapper *PodSchedulingWrapper) Obj() *resourcev1alpha2.PodSchedulingContext {
	return &wrapper.PodSchedulingContext
}

// Name sets `s` as the name of the inner object.
func (wrapper *PodSchedulingWrapper) Name(s string) *PodSchedulingWrapper {
	wrapper.SetName(s)
	return wrapper
}

// UID sets `s` as the UID of the inner object.
func (wrapper *PodSchedulingWrapper) UID(s string) *PodSchedulingWrapper {
	wrapper.SetUID(types.UID(s))
	return wrapper
}

// Namespace sets `s` as the namespace of the inner object.
func (wrapper *PodSchedulingWrapper) Namespace(s string) *PodSchedulingWrapper {
	wrapper.SetNamespace(s)
	return wrapper
}

// OwnerReference updates the owning controller of the inner object.
func (wrapper *PodSchedulingWrapper) OwnerReference(name, uid string, gvk schema.GroupVersionKind) *PodSchedulingWrapper {
	wrapper.OwnerReferences = []metav1.OwnerReference{
		{
			APIVersion:         gvk.GroupVersion().String(),
			Kind:               gvk.Kind,
			Name:               name,
			UID:                types.UID(uid),
			Controller:         ptr.To(true),
			BlockOwnerDeletion: ptr.To(true),
		},
	}
	return wrapper
}

// Label applies a {k,v} label pair to the inner object
func (wrapper *PodSchedulingWrapper) Label(k, v string) *PodSchedulingWrapper {
	if wrapper.Labels == nil {
		wrapper.Labels = make(map[string]string)
	}
	wrapper.Labels[k] = v
	return wrapper
}

// SelectedNode sets that field of the inner object.
func (wrapper *PodSchedulingWrapper) SelectedNode(s string) *PodSchedulingWrapper {
	wrapper.Spec.SelectedNode = s
	return wrapper
}

// PotentialNodes sets that field of the inner object.
func (wrapper *PodSchedulingWrapper) PotentialNodes(nodes ...string) *PodSchedulingWrapper {
	wrapper.Spec.PotentialNodes = nodes
	return wrapper
}

// ResourceClaims sets that field of the inner object.
func (wrapper *PodSchedulingWrapper) ResourceClaims(statuses ...resourcev1alpha2.ResourceClaimSchedulingStatus) *PodSchedulingWrapper {
	wrapper.Status.ResourceClaims = statuses
	return wrapper
}
