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

package volumerestrictions

import (
	"context"

	v1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/util/sets"
	"k8s.io/kubernetes/pkg/scheduler/framework"
	"k8s.io/kubernetes/pkg/scheduler/framework/plugins/names"
)

// VolumeRestrictions is a plugin that checks volume restrictions.
type VolumeRestrictions struct{}

var _ framework.FilterPlugin = &VolumeRestrictions{}
var _ framework.EnqueueExtensions = &VolumeRestrictions{}

// Name is the name of the plugin used in the plugin registry and configurations.
const Name = names.VolumeRestrictions

const (
	// ErrReasonDiskConflict is used for NoDiskConflict predicate error.
	ErrReasonDiskConflict = "node(s) had no available disk"
)

// Name returns name of the plugin. It is used in logs, etc.
func (pl *VolumeRestrictions) Name() string {
	return Name
}

func isVolumeConflict(volume *v1.Volume, pod *v1.Pod) bool {
	for _, existingVolume := range pod.Spec.Volumes {
		// Same GCE disk mounted by multiple pods conflicts unless all pods mount it read-only.
		if volume.GCEPersistentDisk != nil && existingVolume.GCEPersistentDisk != nil {
			disk, existingDisk := volume.GCEPersistentDisk, existingVolume.GCEPersistentDisk
			if disk.PDName == existingDisk.PDName && !(disk.ReadOnly && existingDisk.ReadOnly) {
				return true
			}
		}

		if volume.AWSElasticBlockStore != nil && existingVolume.AWSElasticBlockStore != nil {
			if volume.AWSElasticBlockStore.VolumeID == existingVolume.AWSElasticBlockStore.VolumeID {
				return true
			}
		}

		if volume.ISCSI != nil && existingVolume.ISCSI != nil {
			iqn := volume.ISCSI.IQN
			eiqn := existingVolume.ISCSI.IQN
			// two ISCSI volumes are same, if they share the same iqn. As iscsi volumes are of type
			// RWO or ROX, we could permit only one RW mount. Same iscsi volume mounted by multiple Pods
			// conflict unless all other pods mount as read only.
			if iqn == eiqn && !(volume.ISCSI.ReadOnly && existingVolume.ISCSI.ReadOnly) {
				return true
			}
		}

		if volume.RBD != nil && existingVolume.RBD != nil {
			mon, pool, image := volume.RBD.CephMonitors, volume.RBD.RBDPool, volume.RBD.RBDImage
			emon, epool, eimage := existingVolume.RBD.CephMonitors, existingVolume.RBD.RBDPool, existingVolume.RBD.RBDImage
			// two RBDs images are the same if they share the same Ceph monitor, are in the same RADOS Pool, and have the same image name
			// only one read-write mount is permitted for the same RBD image.
			// same RBD image mounted by multiple Pods conflicts unless all Pods mount the image read-only
			if haveOverlap(mon, emon) && pool == epool && image == eimage && !(volume.RBD.ReadOnly && existingVolume.RBD.ReadOnly) {
				return true
			}
		}
	}

	return false
}

// haveOverlap searches two arrays and returns true if they have at least one common element; returns false otherwise.
func haveOverlap(a1, a2 []string) bool {
	if len(a1) > len(a2) {
		a1, a2 = a2, a1
	}
	m := make(sets.String)

	for _, val := range a1 {
		m.Insert(val)
	}
	for _, val := range a2 {
		if _, ok := m[val]; ok {
			return true
		}
	}

	return false
}

// Filter invoked at the filter extension point.
// It evaluates if a pod can fit due to the volumes it requests, and those that
// are already mounted. If there is already a volume mounted on that node, another pod that uses the same volume
// can't be scheduled there.
// This is GCE, Amazon EBS, ISCSI and Ceph RBD specific for now:
// - GCE PD allows multiple mounts as long as they're all read-only
// - AWS EBS forbids any two pods mounting the same volume ID
// - Ceph RBD forbids if any two pods share at least same monitor, and match pool and image, and the image is read-only
// - ISCSI forbids if any two pods share at least same IQN and ISCSI volume is read-only
func (pl *VolumeRestrictions) Filter(ctx context.Context, _ *framework.CycleState, pod *v1.Pod, nodeInfo *framework.NodeInfo) *framework.Status {
	for i := range pod.Spec.Volumes {
		v := &pod.Spec.Volumes[i]
		// fast path if there is no conflict checking targets.
		if v.GCEPersistentDisk == nil && v.AWSElasticBlockStore == nil && v.RBD == nil && v.ISCSI == nil {
			continue
		}

		for _, ev := range nodeInfo.Pods {
			if isVolumeConflict(v, ev.Pod) {
				return framework.NewStatus(framework.Unschedulable, ErrReasonDiskConflict)
			}
		}
	}
	return nil
}

// EventsToRegister returns the possible events that may make a Pod
// failed by this plugin schedulable.
func (pl *VolumeRestrictions) EventsToRegister() []framework.ClusterEvent {
	return []framework.ClusterEvent{
		// Pods may fail to schedule because of volumes conflicting with other pods on same node.
		// Once running pods are deleted and volumes have been released, the unschedulable pod will be schedulable.
		// Due to immutable fields `spec.volumes`, pod update events are ignored.
		{Resource: framework.Pod, ActionType: framework.Delete},
		// A new Node may make a pod schedulable.
		{Resource: framework.Node, ActionType: framework.Add},
	}
}

// New initializes a new plugin and returns it.
func New(_ runtime.Object, _ framework.Handle) (framework.Plugin, error) {
	return &VolumeRestrictions{}, nil
}
