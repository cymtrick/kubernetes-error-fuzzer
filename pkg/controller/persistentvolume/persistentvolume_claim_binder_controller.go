/*
Copyright 2014 The Kubernetes Authors All rights reserved.

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

package persistentvolume

import (
	"fmt"
	"sync"
	"time"

	"k8s.io/kubernetes/pkg/api"
	"k8s.io/kubernetes/pkg/api/errors"
	"k8s.io/kubernetes/pkg/api/unversioned"
	"k8s.io/kubernetes/pkg/client/cache"
	client "k8s.io/kubernetes/pkg/client/unversioned"
	"k8s.io/kubernetes/pkg/controller/framework"
	"k8s.io/kubernetes/pkg/conversion"
	"k8s.io/kubernetes/pkg/fields"
	"k8s.io/kubernetes/pkg/labels"
	"k8s.io/kubernetes/pkg/runtime"
	"k8s.io/kubernetes/pkg/watch"

	"github.com/golang/glog"
)

// PersistentVolumeClaimBinder is a controller that synchronizes PersistentVolumeClaims.
type PersistentVolumeClaimBinder struct {
	volumeIndex      *persistentVolumeOrderedIndex
	volumeController *framework.Controller
	claimController  *framework.Controller
	client           binderClient
	stopChannels     map[string]chan struct{}
	lock             sync.RWMutex
}

// NewPersistentVolumeClaimBinder creates a new PersistentVolumeClaimBinder
func NewPersistentVolumeClaimBinder(kubeClient client.Interface, syncPeriod time.Duration) *PersistentVolumeClaimBinder {
	volumeIndex := NewPersistentVolumeOrderedIndex()
	binderClient := NewBinderClient(kubeClient)
	binder := &PersistentVolumeClaimBinder{
		volumeIndex: volumeIndex,
		client:      binderClient,
	}

	_, volumeController := framework.NewInformer(
		&cache.ListWatch{
			ListFunc: func() (runtime.Object, error) {
				return kubeClient.PersistentVolumes().List(labels.Everything(), fields.Everything())
			},
			WatchFunc: func(options unversioned.ListOptions) (watch.Interface, error) {
				return kubeClient.PersistentVolumes().Watch(options)
			},
		},
		&api.PersistentVolume{},
		// TODO: Can we have much longer period here?
		syncPeriod,
		framework.ResourceEventHandlerFuncs{
			AddFunc:    binder.addVolume,
			UpdateFunc: binder.updateVolume,
			DeleteFunc: binder.deleteVolume,
		},
	)
	_, claimController := framework.NewInformer(
		&cache.ListWatch{
			ListFunc: func() (runtime.Object, error) {
				return kubeClient.PersistentVolumeClaims(api.NamespaceAll).List(labels.Everything(), fields.Everything())
			},
			WatchFunc: func(options unversioned.ListOptions) (watch.Interface, error) {
				return kubeClient.PersistentVolumeClaims(api.NamespaceAll).Watch(options)
			},
		},
		&api.PersistentVolumeClaim{},
		// TODO: Can we have much longer period here?
		syncPeriod,
		framework.ResourceEventHandlerFuncs{
			AddFunc:    binder.addClaim,
			UpdateFunc: binder.updateClaim,
			// no DeleteFunc needed.  a claim requires no clean-up.
			// syncVolume handles the missing claim
		},
	)

	binder.claimController = claimController
	binder.volumeController = volumeController

	return binder
}

func (binder *PersistentVolumeClaimBinder) addVolume(obj interface{}) {
	binder.lock.Lock()
	defer binder.lock.Unlock()
	volume := obj.(*api.PersistentVolume)
	err := syncVolume(binder.volumeIndex, binder.client, volume)
	if err != nil {
		glog.Errorf("PVClaimBinder could not add volume %s: %+v", volume.Name, err)
	}
}

func (binder *PersistentVolumeClaimBinder) updateVolume(oldObj, newObj interface{}) {
	binder.lock.Lock()
	defer binder.lock.Unlock()
	newVolume := newObj.(*api.PersistentVolume)
	binder.volumeIndex.Update(newVolume)
	err := syncVolume(binder.volumeIndex, binder.client, newVolume)
	if err != nil {
		glog.Errorf("PVClaimBinder could not update volume %s: %+v", newVolume.Name, err)
	}
}

func (binder *PersistentVolumeClaimBinder) deleteVolume(obj interface{}) {
	binder.lock.Lock()
	defer binder.lock.Unlock()
	volume := obj.(*api.PersistentVolume)
	binder.volumeIndex.Delete(volume)
}

func (binder *PersistentVolumeClaimBinder) addClaim(obj interface{}) {
	binder.lock.Lock()
	defer binder.lock.Unlock()
	claim := obj.(*api.PersistentVolumeClaim)
	err := syncClaim(binder.volumeIndex, binder.client, claim)
	if err != nil {
		glog.Errorf("PVClaimBinder could not add claim %s: %+v", claim.Name, err)
	}
}

func (binder *PersistentVolumeClaimBinder) updateClaim(oldObj, newObj interface{}) {
	binder.lock.Lock()
	defer binder.lock.Unlock()
	newClaim := newObj.(*api.PersistentVolumeClaim)
	err := syncClaim(binder.volumeIndex, binder.client, newClaim)
	if err != nil {
		glog.Errorf("PVClaimBinder could not update claim %s: %+v", newClaim.Name, err)
	}
}

func syncVolume(volumeIndex *persistentVolumeOrderedIndex, binderClient binderClient, volume *api.PersistentVolume) (err error) {
	glog.V(5).Infof("Synchronizing PersistentVolume[%s], current phase: %s\n", volume.Name, volume.Status.Phase)

	// volumes can be in one of the following states:
	//
	// VolumePending -- default value -- not bound to a claim and not yet processed through this controller.
	// VolumeAvailable -- not bound to a claim, but processed at least once and found in this controller's volumeIndex.
	// VolumeBound -- bound to a claim because volume.Spec.ClaimRef != nil.   Claim status may not be correct.
	// VolumeReleased -- volume.Spec.ClaimRef != nil but the claim has been deleted by the user.
	// VolumeFailed -- volume.Spec.ClaimRef != nil and the volume failed processing in the recycler
	currentPhase := volume.Status.Phase
	nextPhase := currentPhase

	_, exists, err := volumeIndex.Get(volume)
	if err != nil {
		return err
	}
	if !exists {
		volumeIndex.Add(volume)
	}

	switch currentPhase {
	case api.VolumePending:

		// 3 possible states:
		//  1.  ClaimRef != nil and Claim exists:   Prebound to claim. Make volume available for binding (it will match PVC).
		//  2.  ClaimRef != nil and Claim !exists:  Recently recycled. Remove bind. Make volume available for new claim.
		//  3.  ClaimRef == nil: Neither recycled nor prebound.  Make volume available for binding.
		nextPhase = api.VolumeAvailable

		if volume.Spec.ClaimRef != nil {
			_, err := binderClient.GetPersistentVolumeClaim(volume.Spec.ClaimRef.Namespace, volume.Spec.ClaimRef.Name)
			if errors.IsNotFound(err) {
				// Pending volumes that have a ClaimRef where the claim is missing were recently recycled.
				// The Recycler set the phase to VolumePending to start the volume at the beginning of this lifecycle.
				// removing ClaimRef unbinds the volume
				clone, err := conversion.NewCloner().DeepCopy(volume)
				if err != nil {
					return fmt.Errorf("Error cloning pv: %v", err)
				}
				volumeClone, ok := clone.(*api.PersistentVolume)
				if !ok {
					return fmt.Errorf("Unexpected pv cast error : %v\n", volumeClone)
				}
				volumeClone.Spec.ClaimRef = nil

				if updatedVolume, err := binderClient.UpdatePersistentVolume(volumeClone); err != nil {
					return fmt.Errorf("Unexpected error saving PersistentVolume: %+v", err)
				} else {
					volume = updatedVolume
					volumeIndex.Update(volume)
				}
			} else if err != nil {
				return fmt.Errorf("Error getting PersistentVolumeClaim[%s/%s]: %v", volume.Spec.ClaimRef.Namespace, volume.Spec.ClaimRef.Name, err)
			}
		}
		glog.V(5).Infof("PersistentVolume[%s] is available\n", volume.Name)

	// available volumes await a claim
	case api.VolumeAvailable:
		if volume.Spec.ClaimRef != nil {
			_, err := binderClient.GetPersistentVolumeClaim(volume.Spec.ClaimRef.Namespace, volume.Spec.ClaimRef.Name)
			if err == nil {
				// change of phase will trigger an update event with the newly bound volume
				glog.V(5).Infof("PersistentVolume[%s] is now bound\n", volume.Name)
				nextPhase = api.VolumeBound
			} else {
				if errors.IsNotFound(err) {
					nextPhase = api.VolumeReleased
				}
			}
		}

	//bound volumes require verification of their bound claims
	case api.VolumeBound:
		if volume.Spec.ClaimRef == nil {
			return fmt.Errorf("PersistentVolume[%s] expected to be bound but found nil claimRef: %+v", volume.Name, volume)
		} else {
			_, err := binderClient.GetPersistentVolumeClaim(volume.Spec.ClaimRef.Namespace, volume.Spec.ClaimRef.Name)
			if err != nil {
				if errors.IsNotFound(err) {
					nextPhase = api.VolumeReleased
				} else {
					return err
				}
			}
		}

	// released volumes require recycling
	case api.VolumeReleased:
		if volume.Spec.ClaimRef == nil {
			return fmt.Errorf("PersistentVolume[%s] expected to be bound but found nil claimRef: %+v", volume.Name, volume)
		} else {
			// another process is watching for released volumes.
			// PersistentVolumeReclaimPolicy is set per PersistentVolume
			//  Recycle - sets the PV to Pending and back under this controller's management
			//  Delete - delete events are handled by this controller's watch. PVs are removed from the index.
		}

	// volumes are removed by processes external to this binder and must be removed from the cluster
	case api.VolumeFailed:
		if volume.Spec.ClaimRef == nil {
			return fmt.Errorf("PersistentVolume[%s] expected to be bound but found nil claimRef: %+v", volume.Name, volume)
		} else {
			glog.V(5).Infof("PersistentVolume[%s] previously failed recycling.  Skipping.\n", volume.Name)
		}
	}

	if currentPhase != nextPhase {
		volume.Status.Phase = nextPhase

		// a change in state will trigger another update through this controller.
		// each pass through this controller evaluates current phase and decides whether or not to change to the next phase
		glog.V(5).Infof("PersistentVolume[%s] changing phase from %s to %s\n", volume.Name, currentPhase, nextPhase)
		volume, err := binderClient.UpdatePersistentVolumeStatus(volume)
		if err != nil {
			// Rollback to previous phase
			volume.Status.Phase = currentPhase
		}
		volumeIndex.Update(volume)
	}

	return nil
}

func syncClaim(volumeIndex *persistentVolumeOrderedIndex, binderClient binderClient, claim *api.PersistentVolumeClaim) (err error) {
	glog.V(5).Infof("Synchronizing PersistentVolumeClaim[%s]\n", claim.Name)

	switch claim.Status.Phase {
	case api.ClaimPending:
		volume, err := volumeIndex.findBestMatchForClaim(claim)
		if err != nil {
			return err
		}
		if volume == nil {
			glog.V(5).Infof("A volume match does not exist for persistent claim: %s", claim.Name)
			return nil
		}

		// create a reference to the claim and assign it to the volume being bound.
		// the volume is a pointer and assigning the reference fixes a race condition where another
		// claim might match this volume but before the claimRef is persistent in the next case statement
		claimRef, err := api.GetReference(claim)
		if err != nil {
			return fmt.Errorf("Unexpected error getting claim reference: %v\n", err)
		}

		// make a binding reference to the claim and ensure to update the local index to prevent dupe bindings
		clone, err := conversion.NewCloner().DeepCopy(volume)
		if err != nil {
			return fmt.Errorf("Error cloning pv: %v", err)
		}
		volumeClone, ok := clone.(*api.PersistentVolume)
		if !ok {
			return fmt.Errorf("Unexpected pv cast error : %v\n", volumeClone)
		}
		volumeClone.Spec.ClaimRef = claimRef
		if updatedVolume, err := binderClient.UpdatePersistentVolume(volumeClone); err != nil {
			return fmt.Errorf("Unexpected error saving PersistentVolume.Status: %+v", err)
		} else {
			volume = updatedVolume
			volumeIndex.Update(updatedVolume)
		}

		// the bind is persisted on the volume above and will always match the claim in a search.
		// claim would remain Pending if the update fails, so processing this state is idempotent.
		// this only needs to be processed once.
		if claim.Spec.VolumeName != volume.Name {
			claim.Spec.VolumeName = volume.Name
			claim, err = binderClient.UpdatePersistentVolumeClaim(claim)
			if err != nil {
				return fmt.Errorf("Error updating claim with VolumeName %s: %+v\n", volume.Name, err)
			}
		}

		claim.Status.Phase = api.ClaimBound
		claim.Status.AccessModes = volume.Spec.AccessModes
		claim.Status.Capacity = volume.Spec.Capacity
		_, err = binderClient.UpdatePersistentVolumeClaimStatus(claim)
		if err != nil {
			return fmt.Errorf("Unexpected error saving claim status: %+v", err)
		}

	case api.ClaimBound:
		// no-op.  Claim is bound, values from PV are set.  PVCs are technically mutable in the API server
		// and we don't want to handle those changes at this time.

	default:
		return fmt.Errorf("Unknown state for PVC: %#v", claim)

	}

	glog.V(5).Infof("PersistentVolumeClaim[%s] is bound\n", claim.Name)
	return nil
}

// Run starts all of this binder's control loops
func (controller *PersistentVolumeClaimBinder) Run() {
	glog.V(5).Infof("Starting PersistentVolumeClaimBinder\n")
	if controller.stopChannels == nil {
		controller.stopChannels = make(map[string]chan struct{})
	}

	if _, exists := controller.stopChannels["volumes"]; !exists {
		controller.stopChannels["volumes"] = make(chan struct{})
		go controller.volumeController.Run(controller.stopChannels["volumes"])
	}

	if _, exists := controller.stopChannels["claims"]; !exists {
		controller.stopChannels["claims"] = make(chan struct{})
		go controller.claimController.Run(controller.stopChannels["claims"])
	}
}

// Stop gracefully shuts down this binder
func (controller *PersistentVolumeClaimBinder) Stop() {
	glog.V(5).Infof("Stopping PersistentVolumeClaimBinder\n")
	for name, stopChan := range controller.stopChannels {
		close(stopChan)
		delete(controller.stopChannels, name)
	}
}

// binderClient abstracts access to PVs and PVCs
type binderClient interface {
	GetPersistentVolume(name string) (*api.PersistentVolume, error)
	UpdatePersistentVolume(volume *api.PersistentVolume) (*api.PersistentVolume, error)
	DeletePersistentVolume(volume *api.PersistentVolume) error
	UpdatePersistentVolumeStatus(volume *api.PersistentVolume) (*api.PersistentVolume, error)
	GetPersistentVolumeClaim(namespace, name string) (*api.PersistentVolumeClaim, error)
	UpdatePersistentVolumeClaim(claim *api.PersistentVolumeClaim) (*api.PersistentVolumeClaim, error)
	UpdatePersistentVolumeClaimStatus(claim *api.PersistentVolumeClaim) (*api.PersistentVolumeClaim, error)
}

func NewBinderClient(c client.Interface) binderClient {
	return &realBinderClient{c}
}

type realBinderClient struct {
	client client.Interface
}

func (c *realBinderClient) GetPersistentVolume(name string) (*api.PersistentVolume, error) {
	return c.client.PersistentVolumes().Get(name)
}

func (c *realBinderClient) UpdatePersistentVolume(volume *api.PersistentVolume) (*api.PersistentVolume, error) {
	return c.client.PersistentVolumes().Update(volume)
}

func (c *realBinderClient) DeletePersistentVolume(volume *api.PersistentVolume) error {
	return c.client.PersistentVolumes().Delete(volume.Name)
}

func (c *realBinderClient) UpdatePersistentVolumeStatus(volume *api.PersistentVolume) (*api.PersistentVolume, error) {
	return c.client.PersistentVolumes().UpdateStatus(volume)
}

func (c *realBinderClient) GetPersistentVolumeClaim(namespace, name string) (*api.PersistentVolumeClaim, error) {
	return c.client.PersistentVolumeClaims(namespace).Get(name)
}

func (c *realBinderClient) UpdatePersistentVolumeClaim(claim *api.PersistentVolumeClaim) (*api.PersistentVolumeClaim, error) {
	return c.client.PersistentVolumeClaims(claim.Namespace).Update(claim)
}

func (c *realBinderClient) UpdatePersistentVolumeClaimStatus(claim *api.PersistentVolumeClaim) (*api.PersistentVolumeClaim, error) {
	return c.client.PersistentVolumeClaims(claim.Namespace).UpdateStatus(claim)
}
