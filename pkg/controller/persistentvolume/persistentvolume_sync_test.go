/*
Copyright 2016 The Kubernetes Authors All rights reserved.

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
	"testing"

	"k8s.io/kubernetes/pkg/api"
)

// Test single call to syncClaim and syncVolume methods.
// 1. Fill in the controller with initial data
// 2. Call the tested function (syncClaim/syncVolume) via
//    controllerTest.testCall *once*.
// 3. Compare resulting volumes and claims with expected volumes and claims.
func TestSync(t *testing.T) {
	tests := []controllerTest{
		// [Unit test set 1] User did not care which PV they get.
		// Test the matching with no claim.Spec.VolumeName and with various
		// volumes.
		{
			// syncClaim binds to a matching unbound volume.
			"1-1 - successful bind",
			newVolumeArray("volume1-1", "1Gi", "", "", api.VolumePending, api.PersistentVolumeReclaimRetain),
			newVolumeArray("volume1-1", "1Gi", "uid1-1", "claim1-1", api.VolumeBound, api.PersistentVolumeReclaimRetain, annBoundByController),
			newClaimArray("claim1-1", "uid1-1", "1Gi", "", api.ClaimPending),
			newClaimArray("claim1-1", "uid1-1", "1Gi", "volume1-1", api.ClaimBound, annBoundByController, annBindCompleted),
			noevents, testSyncClaim,
		},
		{
			// syncClaim does not do anything when there is no matching volume.
			"1-2 - noop",
			newVolumeArray("volume1-2", "1Gi", "", "", api.VolumePending, api.PersistentVolumeReclaimRetain),
			newVolumeArray("volume1-2", "1Gi", "", "", api.VolumePending, api.PersistentVolumeReclaimRetain),
			newClaimArray("claim1-2", "uid1-2", "10Gi", "", api.ClaimPending),
			newClaimArray("claim1-2", "uid1-2", "10Gi", "", api.ClaimPending),
			noevents, testSyncClaim,
		},
		{
			// syncClaim resets claim.Status to Pending when there is no
			// matching volume.
			"1-3 - reset to Pending",
			newVolumeArray("volume1-3", "1Gi", "", "", api.VolumePending, api.PersistentVolumeReclaimRetain),
			newVolumeArray("volume1-3", "1Gi", "", "", api.VolumePending, api.PersistentVolumeReclaimRetain),
			newClaimArray("claim1-3", "uid1-3", "10Gi", "", api.ClaimBound),
			newClaimArray("claim1-3", "uid1-3", "10Gi", "", api.ClaimPending),
			noevents, testSyncClaim,
		},
		{
			// syncClaim binds claims to the smallest matching volume
			"1-4 - smallest volume",
			[]*api.PersistentVolume{
				newVolume("volume1-4_1", "10Gi", "", "", api.VolumePending, api.PersistentVolumeReclaimRetain),
				newVolume("volume1-4_2", "1Gi", "", "", api.VolumePending, api.PersistentVolumeReclaimRetain),
			},
			[]*api.PersistentVolume{
				newVolume("volume1-4_1", "10Gi", "", "", api.VolumePending, api.PersistentVolumeReclaimRetain),
				newVolume("volume1-4_2", "1Gi", "uid1-4", "claim1-4", api.VolumeBound, api.PersistentVolumeReclaimRetain, annBoundByController),
			},
			newClaimArray("claim1-4", "uid1-4", "1Gi", "", api.ClaimPending),
			newClaimArray("claim1-4", "uid1-4", "1Gi", "volume1-4_2", api.ClaimBound, annBoundByController, annBindCompleted),
			noevents, testSyncClaim,
		},
		{
			// syncClaim binds a claim only to volume that points to it (by
			// name), even though a smaller one is available.
			"1-5 - prebound volume by name - success",
			[]*api.PersistentVolume{
				newVolume("volume1-5_1", "10Gi", "", "claim1-5", api.VolumePending, api.PersistentVolumeReclaimRetain),
				newVolume("volume1-5_2", "1Gi", "", "", api.VolumePending, api.PersistentVolumeReclaimRetain),
			},
			[]*api.PersistentVolume{
				newVolume("volume1-5_1", "10Gi", "uid1-5", "claim1-5", api.VolumeBound, api.PersistentVolumeReclaimRetain),
				newVolume("volume1-5_2", "1Gi", "", "", api.VolumePending, api.PersistentVolumeReclaimRetain),
			},
			newClaimArray("claim1-5", "uid1-5", "1Gi", "", api.ClaimPending),
			newClaimArray("claim1-5", "uid1-5", "1Gi", "volume1-5_1", api.ClaimBound, annBoundByController, annBindCompleted),
			noevents, testSyncClaim,
		},
		{
			// syncClaim binds a claim only to volume that points to it (by
			// UID), even though a smaller one is available.
			"1-6 - prebound volume by UID - success",
			[]*api.PersistentVolume{
				newVolume("volume1-6_1", "10Gi", "uid1-6", "claim1-6", api.VolumePending, api.PersistentVolumeReclaimRetain),
				newVolume("volume1-6_2", "1Gi", "", "", api.VolumePending, api.PersistentVolumeReclaimRetain),
			},
			[]*api.PersistentVolume{
				newVolume("volume1-6_1", "10Gi", "uid1-6", "claim1-6", api.VolumeBound, api.PersistentVolumeReclaimRetain),
				newVolume("volume1-6_2", "1Gi", "", "", api.VolumePending, api.PersistentVolumeReclaimRetain),
			},
			newClaimArray("claim1-6", "uid1-6", "1Gi", "", api.ClaimPending),
			newClaimArray("claim1-6", "uid1-6", "1Gi", "volume1-6_1", api.ClaimBound, annBoundByController, annBindCompleted),
			noevents, testSyncClaim,
		},
		{
			// syncClaim does not bind claim to a volume prebound to a claim with
			// same name and different UID
			"1-7 - prebound volume to different claim",
			newVolumeArray("volume1-7", "10Gi", "uid1-777", "claim1-7", api.VolumePending, api.PersistentVolumeReclaimRetain),
			newVolumeArray("volume1-7", "10Gi", "uid1-777", "claim1-7", api.VolumePending, api.PersistentVolumeReclaimRetain),
			newClaimArray("claim1-7", "uid1-7", "1Gi", "", api.ClaimPending),
			newClaimArray("claim1-7", "uid1-7", "1Gi", "", api.ClaimPending),
			noevents, testSyncClaim,
		},
		{
			// syncClaim completes binding - simulates controller crash after
			// PV.ClaimRef is saved
			"1-8 - complete bind after crash - PV bound",
			newVolumeArray("volume1-8", "1Gi", "uid1-8", "claim1-8", api.VolumePending, api.PersistentVolumeReclaimRetain, annBoundByController),
			newVolumeArray("volume1-8", "1Gi", "uid1-8", "claim1-8", api.VolumeBound, api.PersistentVolumeReclaimRetain, annBoundByController),
			newClaimArray("claim1-8", "uid1-8", "1Gi", "", api.ClaimPending),
			newClaimArray("claim1-8", "uid1-8", "1Gi", "volume1-8", api.ClaimBound, annBoundByController, annBindCompleted),
			noevents, testSyncClaim,
		},
		{
			// syncClaim completes binding - simulates controller crash after
			// PV.Status is saved
			"1-9 - complete bind after crash - PV status saved",
			newVolumeArray("volume1-9", "1Gi", "uid1-9", "claim1-9", api.VolumeBound, api.PersistentVolumeReclaimRetain, annBoundByController),
			newVolumeArray("volume1-9", "1Gi", "uid1-9", "claim1-9", api.VolumeBound, api.PersistentVolumeReclaimRetain, annBoundByController),
			newClaimArray("claim1-9", "uid1-9", "1Gi", "", api.ClaimPending),
			newClaimArray("claim1-9", "uid1-9", "1Gi", "volume1-9", api.ClaimBound, annBoundByController, annBindCompleted),
			noevents, testSyncClaim,
		},
		{
			// syncClaim completes binding - simulates controller crash after
			// PVC.VolumeName is saved
			"10 - complete bind after crash - PVC bound",
			newVolumeArray("volume1-10", "1Gi", "uid1-10", "claim1-10", api.VolumeBound, api.PersistentVolumeReclaimRetain, annBoundByController),
			newVolumeArray("volume1-10", "1Gi", "uid1-10", "claim1-10", api.VolumeBound, api.PersistentVolumeReclaimRetain, annBoundByController),
			newClaimArray("claim1-10", "uid1-10", "1Gi", "volume1-10", api.ClaimPending, annBoundByController, annBindCompleted),
			newClaimArray("claim1-10", "uid1-10", "1Gi", "volume1-10", api.ClaimBound, annBoundByController, annBindCompleted),
			noevents, testSyncClaim,
		},
		// [Unit test set 2] User asked for a specific PV.
		// Test the binding when pv.ClaimRef is already set by controller or
		// by user.
		{
			// syncClaim with claim pre-bound to a PV that does not exist
			"2-1 - claim prebound to non-existing volume - noop",
			novolumes,
			novolumes,
			newClaimArray("claim2-1", "uid2-1", "10Gi", "volume2-1", api.ClaimPending),
			newClaimArray("claim2-1", "uid2-1", "10Gi", "volume2-1", api.ClaimPending),
			noevents, testSyncClaim,
		},
		{
			// syncClaim with claim pre-bound to a PV that does not exist.
			// Check that the claim status is reset to Pending
			"2-2 - claim prebound to non-existing volume - reset status",
			novolumes,
			novolumes,
			newClaimArray("claim2-2", "uid2-2", "10Gi", "volume2-2", api.ClaimBound),
			newClaimArray("claim2-2", "uid2-2", "10Gi", "volume2-2", api.ClaimPending),
			noevents, testSyncClaim,
		},
		{
			// syncClaim with claim pre-bound to a PV that exists and is
			// unbound. Check it gets bound and no annBoundByController is set.
			"2-3 - claim prebound to unbound volume",
			newVolumeArray("volume2-3", "1Gi", "", "", api.VolumePending, api.PersistentVolumeReclaimRetain),
			newVolumeArray("volume2-3", "1Gi", "uid2-3", "claim2-3", api.VolumeBound, api.PersistentVolumeReclaimRetain, annBoundByController),
			newClaimArray("claim2-3", "uid2-3", "10Gi", "volume2-3", api.ClaimPending),
			newClaimArray("claim2-3", "uid2-3", "10Gi", "volume2-3", api.ClaimBound, annBindCompleted),
			noevents, testSyncClaim,
		},
		{
			// claim with claim pre-bound to a PV that is pre-bound to the claim
			// by name. Check it gets bound and no annBoundByController is set.
			"2-4 - claim prebound to prebound volume by name",
			newVolumeArray("volume2-4", "1Gi", "", "claim2-4", api.VolumePending, api.PersistentVolumeReclaimRetain),
			newVolumeArray("volume2-4", "1Gi", "uid2-4", "claim2-4", api.VolumeBound, api.PersistentVolumeReclaimRetain),
			newClaimArray("claim2-4", "uid2-4", "10Gi", "volume2-4", api.ClaimPending),
			newClaimArray("claim2-4", "uid2-4", "10Gi", "volume2-4", api.ClaimBound, annBindCompleted),
			noevents, testSyncClaim,
		},
		{
			// syncClaim with claim pre-bound to a PV that is pre-bound to the
			// claim by UID. Check it gets bound and no annBoundByController is
			// set.
			"2-5 - claim prebound to prebound volume by UID",
			newVolumeArray("volume2-5", "1Gi", "uid2-5", "claim2-5", api.VolumePending, api.PersistentVolumeReclaimRetain),
			newVolumeArray("volume2-5", "1Gi", "uid2-5", "claim2-5", api.VolumeBound, api.PersistentVolumeReclaimRetain),
			newClaimArray("claim2-5", "uid2-5", "10Gi", "volume2-5", api.ClaimPending),
			newClaimArray("claim2-5", "uid2-5", "10Gi", "volume2-5", api.ClaimBound, annBindCompleted),
			noevents, testSyncClaim,
		},
		{
			// syncClaim with claim pre-bound to a PV that is bound to different
			// claim. Check it's reset to Pending.
			"2-6 - claim prebound to already bound volume",
			newVolumeArray("volume2-6", "1Gi", "uid2-6_1", "claim2-6_1", api.VolumeBound, api.PersistentVolumeReclaimRetain),
			newVolumeArray("volume2-6", "1Gi", "uid2-6_1", "claim2-6_1", api.VolumeBound, api.PersistentVolumeReclaimRetain),
			newClaimArray("claim2-6", "uid2-6", "10Gi", "volume2-6", api.ClaimBound),
			newClaimArray("claim2-6", "uid2-6", "10Gi", "volume2-6", api.ClaimPending),
			noevents, testSyncClaim,
		},
		{
			// syncClaim with claim bound by controller to a PV that is bound to
			// different claim. Check it throws an error.
			"2-7 - claim bound by controller to already bound volume",
			newVolumeArray("volume2-7", "1Gi", "uid2-7_1", "claim2-7_1", api.VolumeBound, api.PersistentVolumeReclaimRetain),
			newVolumeArray("volume2-7", "1Gi", "uid2-7_1", "claim2-7_1", api.VolumeBound, api.PersistentVolumeReclaimRetain),
			newClaimArray("claim2-7", "uid2-7", "10Gi", "volume2-7", api.ClaimBound, annBoundByController),
			newClaimArray("claim2-7", "uid2-7", "10Gi", "volume2-7", api.ClaimBound, annBoundByController),
			noevents, testSyncClaimError,
		},
		// [Unit test set 3] Syncing bound claim
		{
			// syncClaim with claim  bound and its claim.Spec.VolumeName is
			// removed. Check it's marked as Lost.
			"3-1 - bound claim with missing VolumeName",
			novolumes,
			novolumes,
			newClaimArray("claim3-1", "uid3-1", "10Gi", "", api.ClaimBound, annBoundByController, annBindCompleted),
			newClaimArray("claim3-1", "uid3-1", "10Gi", "", api.ClaimLost, annBoundByController, annBindCompleted),
			[]string{"Warning ClaimLost"}, testSyncClaim,
		},
		{
			// syncClaim with claim bound to non-exising volume. Check it's
			// marked as Lost.
			"3-2 - bound claim with missing volume",
			novolumes,
			novolumes,
			newClaimArray("claim3-2", "uid3-2", "10Gi", "volume3-2", api.ClaimBound, annBoundByController, annBindCompleted),
			newClaimArray("claim3-2", "uid3-2", "10Gi", "volume3-2", api.ClaimLost, annBoundByController, annBindCompleted),
			[]string{"Warning ClaimLost"}, testSyncClaim,
		},
		{
			// syncClaim with claim bound to unbound volume. Check it's bound.
			// Also check that Pending phase is set to Bound
			"3-3 - bound claim with unbound volume",
			newVolumeArray("volume3-3", "10Gi", "", "", api.VolumePending, api.PersistentVolumeReclaimRetain),
			newVolumeArray("volume3-3", "10Gi", "uid3-3", "claim3-3", api.VolumeBound, api.PersistentVolumeReclaimRetain, annBoundByController),
			newClaimArray("claim3-3", "uid3-3", "10Gi", "volume3-3", api.ClaimPending, annBoundByController, annBindCompleted),
			newClaimArray("claim3-3", "uid3-3", "10Gi", "volume3-3", api.ClaimBound, annBoundByController, annBindCompleted),
			noevents, testSyncClaim,
		},
		{
			// syncClaim with claim bound to volume with missing (or different)
			// volume.Spec.ClaimRef.UID. Check that the claim is marked as lost.
			"3-4 - bound claim with prebound volume",
			newVolumeArray("volume3-4", "10Gi", "claim3-4-x", "claim3-4", api.VolumePending, api.PersistentVolumeReclaimRetain),
			newVolumeArray("volume3-4", "10Gi", "claim3-4-x", "claim3-4", api.VolumePending, api.PersistentVolumeReclaimRetain),
			newClaimArray("claim3-4", "uid3-4", "10Gi", "volume3-4", api.ClaimPending, annBoundByController, annBindCompleted),
			newClaimArray("claim3-4", "uid3-4", "10Gi", "volume3-4", api.ClaimLost, annBoundByController, annBindCompleted),
			[]string{"Warning ClaimMisbound"}, testSyncClaim,
		},
		{
			// syncClaim with claim bound to bound volume. Check that the
			// controller does not do anything. Also check that Pending phase is
			// set to Bound
			"3-5 - bound claim with bound volume",
			newVolumeArray("volume3-5", "10Gi", "uid3-5", "claim3-5", api.VolumePending, api.PersistentVolumeReclaimRetain),
			newVolumeArray("volume3-5", "10Gi", "uid3-5", "claim3-5", api.VolumeBound, api.PersistentVolumeReclaimRetain),
			newClaimArray("claim3-5", "uid3-5", "10Gi", "volume3-5", api.ClaimPending, annBindCompleted),
			newClaimArray("claim3-5", "uid3-5", "10Gi", "volume3-5", api.ClaimBound, annBindCompleted),
			noevents, testSyncClaim,
		},
		{
			// syncClaim with claim bound to a volume that is bound to different
			// claim. Check that the claim is marked as lost.
			// TODO: test that an event is emitted
			"3-6 - bound claim with bound volume",
			newVolumeArray("volume3-6", "10Gi", "uid3-6-x", "claim3-6-x", api.VolumePending, api.PersistentVolumeReclaimRetain),
			newVolumeArray("volume3-6", "10Gi", "uid3-6-x", "claim3-6-x", api.VolumePending, api.PersistentVolumeReclaimRetain),
			newClaimArray("claim3-6", "uid3-6", "10Gi", "volume3-6", api.ClaimPending, annBindCompleted),
			newClaimArray("claim3-6", "uid3-6", "10Gi", "volume3-6", api.ClaimLost, annBindCompleted),
			[]string{"Warning ClaimMisbound"}, testSyncClaim,
		},
		// [Unit test set 4] All syncVolume tests.
		{
			// syncVolume with pending volume. Check it's marked as Available.
			"4-1 - pending volume",
			newVolumeArray("volume4-1", "10Gi", "", "", api.VolumePending, api.PersistentVolumeReclaimRetain),
			newVolumeArray("volume4-1", "10Gi", "", "", api.VolumeAvailable, api.PersistentVolumeReclaimRetain),
			noclaims,
			noclaims,
			noevents, testSyncVolume,
		},
		{
			// syncVolume with prebound pending volume. Check it's marked as
			// Available.
			"4-2 - pending prebound volume",
			newVolumeArray("volume4-2", "10Gi", "", "claim4-2", api.VolumePending, api.PersistentVolumeReclaimRetain),
			newVolumeArray("volume4-2", "10Gi", "", "claim4-2", api.VolumeAvailable, api.PersistentVolumeReclaimRetain),
			noclaims,
			noclaims,
			noevents, testSyncVolume,
		},
		{
			// syncVolume with volume bound to missing claim.
			// Check the volume gets Released
			"4-3 - bound volume with missing claim",
			newVolumeArray("volume4-3", "10Gi", "uid4-3", "claim4-3", api.VolumeBound, api.PersistentVolumeReclaimRetain),
			newVolumeArray("volume4-3", "10Gi", "uid4-3", "claim4-3", api.VolumeReleased, api.PersistentVolumeReclaimRetain),
			noclaims,
			noclaims,
			noevents, testSyncVolume,
		},
		{
			// syncVolume with volume bound to claim with different UID.
			// Check the volume gets Released.
			"4-4 - volume bound to claim with different UID",
			newVolumeArray("volume4-4", "10Gi", "uid4-4", "claim4-4", api.VolumeBound, api.PersistentVolumeReclaimRetain),
			newVolumeArray("volume4-4", "10Gi", "uid4-4", "claim4-4", api.VolumeReleased, api.PersistentVolumeReclaimRetain),
			newClaimArray("claim4-4", "uid4-4-x", "10Gi", "volume4-4", api.ClaimBound, annBindCompleted),
			newClaimArray("claim4-4", "uid4-4-x", "10Gi", "volume4-4", api.ClaimBound, annBindCompleted),
			noevents, testSyncVolume,
		},
		{
			// syncVolume with volume bound by controller to unbound claim.
			// Check syncVolume does not do anything.
			"4-5 - volume bound by controller to unbound claim",
			newVolumeArray("volume4-5", "10Gi", "uid4-5", "claim4-5", api.VolumeBound, api.PersistentVolumeReclaimRetain, annBoundByController),
			newVolumeArray("volume4-5", "10Gi", "uid4-5", "claim4-5", api.VolumeBound, api.PersistentVolumeReclaimRetain, annBoundByController),
			newClaimArray("claim4-5", "uid4-5", "10Gi", "", api.ClaimPending),
			newClaimArray("claim4-5", "uid4-5", "10Gi", "", api.ClaimPending),
			noevents, testSyncVolume,
		},
		{
			// syncVolume with volume bound by user to unbound claim.
			// Check syncVolume does not do anything.
			"4-5 - volume bound by user to bound claim",
			newVolumeArray("volume4-5", "10Gi", "uid4-5", "claim4-5", api.VolumeBound, api.PersistentVolumeReclaimRetain),
			newVolumeArray("volume4-5", "10Gi", "uid4-5", "claim4-5", api.VolumeBound, api.PersistentVolumeReclaimRetain),
			newClaimArray("claim4-5", "uid4-5", "10Gi", "", api.ClaimPending),
			newClaimArray("claim4-5", "uid4-5", "10Gi", "", api.ClaimPending),
			noevents, testSyncVolume,
		},
		{
			// syncVolume with volume bound to bound claim.
			// Check that the volume is marked as Bound.
			"4-6 - volume bound by to bound claim",
			newVolumeArray("volume4-6", "10Gi", "uid4-6", "claim4-6", api.VolumeAvailable, api.PersistentVolumeReclaimRetain),
			newVolumeArray("volume4-6", "10Gi", "uid4-6", "claim4-6", api.VolumeBound, api.PersistentVolumeReclaimRetain),
			newClaimArray("claim4-6", "uid4-6", "10Gi", "volume4-6", api.ClaimBound),
			newClaimArray("claim4-6", "uid4-6", "10Gi", "volume4-6", api.ClaimBound),
			noevents, testSyncVolume,
		},
		{
			// syncVolume with volume bound by controller to claim bound to
			// another volume. Check that the volume is rolled back.
			"4-7 - volume bound by controller to claim bound somewhere else",
			newVolumeArray("volume4-7", "10Gi", "uid4-7", "claim4-7", api.VolumeBound, api.PersistentVolumeReclaimRetain, annBoundByController),
			newVolumeArray("volume4-7", "10Gi", "", "", api.VolumeAvailable, api.PersistentVolumeReclaimRetain),
			newClaimArray("claim4-7", "uid4-7", "10Gi", "volume4-7-x", api.ClaimBound),
			newClaimArray("claim4-7", "uid4-7", "10Gi", "volume4-7-x", api.ClaimBound),
			noevents, testSyncVolume,
		},
		{
			// syncVolume with volume bound by user to claim bound to
			// another volume. Check that the volume is marked as Available
			// and its UID is reset.
			"4-8 - volume bound by user to claim bound somewhere else",
			newVolumeArray("volume4-8", "10Gi", "uid4-8", "claim4-8", api.VolumeBound, api.PersistentVolumeReclaimRetain),
			newVolumeArray("volume4-8", "10Gi", "", "claim4-8", api.VolumeAvailable, api.PersistentVolumeReclaimRetain),
			newClaimArray("claim4-8", "uid4-8", "10Gi", "volume4-8-x", api.ClaimBound),
			newClaimArray("claim4-8", "uid4-8", "10Gi", "volume4-8-x", api.ClaimBound),
			noevents, testSyncVolume,
		},
	}
	runSyncTests(t, tests)
}

// Test multiple calls to syncClaim/syncVolume and periodic sync of all
// volume/claims. The test follows this pattern:
// 0. Load the controller with initial data.
// 1. Call controllerTest.testCall() once as in TestSync()
// 2. For all volumes/claims changed by previous syncVolume/syncClaim calls,
//    call appropriate syncVolume/syncClaim (simulating "volume/claim changed"
//    events). Go to 2. if these calls change anything.
// 3. When all changes are processed and no new changes were made, call
//    syncVolume/syncClaim on all volumes/claims (simulating "periodic sync").
// 4. If some changes were done by step 3., go to 2. (simulation of
//    "volume/claim updated" events, eventually performing step 3. again)
// 5. When 3. does not do any changes, finish the tests and compare final set
//    of volumes/claims with expected claims/volumes and report differences.
// Some limit of calls in enforced to prevent endless loops.
func TestMultiSync(t *testing.T) {
	tests := []controllerTest{
		// Test simple binding
		{
			// syncClaim binds to a matching unbound volume.
			"10-1 - successful bind",
			newVolumeArray("volume10-1", "1Gi", "", "", api.VolumePending, api.PersistentVolumeReclaimRetain),
			newVolumeArray("volume10-1", "1Gi", "uid10-1", "claim10-1", api.VolumeBound, api.PersistentVolumeReclaimRetain, annBoundByController),
			newClaimArray("claim10-1", "uid10-1", "1Gi", "", api.ClaimPending),
			newClaimArray("claim10-1", "uid10-1", "1Gi", "volume10-1", api.ClaimBound, annBoundByController, annBindCompleted),
			noevents, testSyncClaim,
		},
		{
			// Two controllers bound two PVs to single claim. Test one of them
			// wins and the second rolls back.
			"10-2 - bind PV race",
			[]*api.PersistentVolume{
				newVolume("volume10-2-1", "1Gi", "uid10-2", "claim10-2", api.VolumeBound, api.PersistentVolumeReclaimRetain, annBoundByController),
				newVolume("volume10-2-2", "1Gi", "uid10-2", "claim10-2", api.VolumeBound, api.PersistentVolumeReclaimRetain, annBoundByController),
			},
			[]*api.PersistentVolume{
				newVolume("volume10-2-1", "1Gi", "uid10-2", "claim10-2", api.VolumeBound, api.PersistentVolumeReclaimRetain, annBoundByController),
				newVolume("volume10-2-2", "1Gi", "", "", api.VolumeAvailable, api.PersistentVolumeReclaimRetain),
			},
			newClaimArray("claim10-2", "uid10-2", "1Gi", "volume10-2-1", api.ClaimBound, annBoundByController, annBindCompleted),
			newClaimArray("claim10-2", "uid10-2", "1Gi", "volume10-2-1", api.ClaimBound, annBoundByController, annBindCompleted),
			noevents, testSyncClaim,
		},
	}

	runMultisyncTests(t, tests)
}
