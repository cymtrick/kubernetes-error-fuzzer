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

package status

import (
	"fmt"
	"math/rand"
	"strconv"
	"testing"
	"time"

	clientset "k8s.io/kubernetes/pkg/client/clientset_generated/internalclientset"
	"k8s.io/kubernetes/pkg/client/clientset_generated/internalclientset/fake"
	"k8s.io/kubernetes/pkg/client/testing/core"

	"github.com/stretchr/testify/assert"

	"k8s.io/kubernetes/pkg/api"
	"k8s.io/kubernetes/pkg/api/errors"
	"k8s.io/kubernetes/pkg/api/unversioned"
	kubecontainer "k8s.io/kubernetes/pkg/kubelet/container"
	kubepod "k8s.io/kubernetes/pkg/kubelet/pod"
	podtest "k8s.io/kubernetes/pkg/kubelet/pod/testing"
	kubetypes "k8s.io/kubernetes/pkg/kubelet/types"
	"k8s.io/kubernetes/pkg/runtime"
)

// Generate new instance of test pod with the same initial value.
func getTestPod() *api.Pod {
	return &api.Pod{
		ObjectMeta: api.ObjectMeta{
			UID:       "12345678",
			Name:      "foo",
			Namespace: "new",
		},
	}
}

// After adding reconciliation, if status in pod manager is different from the cached status, a reconciliation
// will be triggered, which will mess up all the old unit test.
// To simplify the implementation of unit test, we add testSyncBatch() here, it will make sure the statuses in
// pod manager the same with cached ones before syncBatch() so as to avoid reconciling.
func (m *manager) testSyncBatch() {
	for uid, status := range m.podStatuses {
		pod, ok := m.podManager.GetPodByUID(uid)
		if ok {
			pod.Status = status.status
		}
		pod, ok = m.podManager.GetMirrorPodByPod(pod)
		if ok {
			pod.Status = status.status
		}
	}
	m.syncBatch()
}

func newTestManager(kubeClient clientset.Interface) *manager {
	podManager := kubepod.NewBasicPodManager(podtest.NewFakeMirrorClient())
	podManager.AddPod(getTestPod())
	return NewManager(kubeClient, podManager).(*manager)
}

func generateRandomMessage() string {
	return strconv.Itoa(rand.Int())
}

func getRandomPodStatus() api.PodStatus {
	return api.PodStatus{
		Message: generateRandomMessage(),
	}
}

func verifyActions(t *testing.T, kubeClient clientset.Interface, expectedActions []core.Action) {
	actions := kubeClient.(*fake.Clientset).Actions()
	if len(actions) != len(expectedActions) {
		t.Fatalf("unexpected actions, got: %+v expected: %+v", actions, expectedActions)
		return
	}
	for i := 0; i < len(actions); i++ {
		e := expectedActions[i]
		a := actions[i]
		if !a.Matches(e.GetVerb(), e.GetResource()) || a.GetSubresource() != e.GetSubresource() {
			t.Errorf("unexpected actions, got: %+v expected: %+v", actions, expectedActions)
		}
	}
}

func verifyUpdates(t *testing.T, manager *manager, expectedUpdates int) {
	// Consume all updates in the channel.
	numUpdates := 0
	for {
		hasUpdate := true
		select {
		case <-manager.podStatusChannel:
			numUpdates++
		default:
			hasUpdate = false
		}

		if !hasUpdate {
			break
		}
	}

	if numUpdates != expectedUpdates {
		t.Errorf("unexpected number of updates %d, expected %d", numUpdates, expectedUpdates)
	}
}

func TestNewStatus(t *testing.T) {
	syncer := newTestManager(&fake.Clientset{})
	testPod := getTestPod()
	syncer.SetPodStatus(testPod, getRandomPodStatus())
	verifyUpdates(t, syncer, 1)

	status := expectPodStatus(t, syncer, testPod)
	if status.StartTime.IsZero() {
		t.Errorf("SetPodStatus did not set a proper start time value")
	}
}

func TestNewStatusPreservesPodStartTime(t *testing.T) {
	syncer := newTestManager(&fake.Clientset{})
	pod := &api.Pod{
		ObjectMeta: api.ObjectMeta{
			UID:       "12345678",
			Name:      "foo",
			Namespace: "new",
		},
		Status: api.PodStatus{},
	}
	now := unversioned.Now()
	startTime := unversioned.NewTime(now.Time.Add(-1 * time.Minute))
	pod.Status.StartTime = &startTime
	syncer.SetPodStatus(pod, getRandomPodStatus())

	status := expectPodStatus(t, syncer, pod)
	if !status.StartTime.Time.Equal(startTime.Time) {
		t.Errorf("Unexpected start time, expected %v, actual %v", startTime, status.StartTime)
	}
}

func getReadyPodStatus() api.PodStatus {
	return api.PodStatus{
		Conditions: []api.PodCondition{
			{
				Type:   api.PodReady,
				Status: api.ConditionTrue,
			},
		},
	}
}

func TestNewStatusSetsReadyTransitionTime(t *testing.T) {
	syncer := newTestManager(&fake.Clientset{})
	podStatus := getReadyPodStatus()
	pod := &api.Pod{
		ObjectMeta: api.ObjectMeta{
			UID:       "12345678",
			Name:      "foo",
			Namespace: "new",
		},
		Status: api.PodStatus{},
	}
	syncer.SetPodStatus(pod, podStatus)
	verifyUpdates(t, syncer, 1)
	status := expectPodStatus(t, syncer, pod)
	readyCondition := api.GetPodReadyCondition(status)
	if readyCondition.LastTransitionTime.IsZero() {
		t.Errorf("Unexpected: last transition time not set")
	}
}

func TestChangedStatus(t *testing.T) {
	syncer := newTestManager(&fake.Clientset{})
	testPod := getTestPod()
	syncer.SetPodStatus(testPod, getRandomPodStatus())
	syncer.SetPodStatus(testPod, getRandomPodStatus())
	verifyUpdates(t, syncer, 2)
}

func TestChangedStatusKeepsStartTime(t *testing.T) {
	syncer := newTestManager(&fake.Clientset{})
	testPod := getTestPod()
	now := unversioned.Now()
	firstStatus := getRandomPodStatus()
	firstStatus.StartTime = &now
	syncer.SetPodStatus(testPod, firstStatus)
	syncer.SetPodStatus(testPod, getRandomPodStatus())
	verifyUpdates(t, syncer, 2)
	finalStatus := expectPodStatus(t, syncer, testPod)
	if finalStatus.StartTime.IsZero() {
		t.Errorf("StartTime should not be zero")
	}
	expected := now.Rfc3339Copy()
	if !finalStatus.StartTime.Equal(expected) {
		t.Errorf("Expected %v, but got %v", expected, finalStatus.StartTime)
	}
}

func TestChangedStatusUpdatesLastTransitionTime(t *testing.T) {
	syncer := newTestManager(&fake.Clientset{})
	podStatus := getReadyPodStatus()
	pod := &api.Pod{
		ObjectMeta: api.ObjectMeta{
			UID:       "12345678",
			Name:      "foo",
			Namespace: "new",
		},
		Status: api.PodStatus{},
	}
	syncer.SetPodStatus(pod, podStatus)
	verifyUpdates(t, syncer, 1)
	oldStatus := expectPodStatus(t, syncer, pod)
	anotherStatus := getReadyPodStatus()
	anotherStatus.Conditions[0].Status = api.ConditionFalse
	syncer.SetPodStatus(pod, anotherStatus)
	verifyUpdates(t, syncer, 1)
	newStatus := expectPodStatus(t, syncer, pod)

	oldReadyCondition := api.GetPodReadyCondition(oldStatus)
	newReadyCondition := api.GetPodReadyCondition(newStatus)
	if newReadyCondition.LastTransitionTime.IsZero() {
		t.Errorf("Unexpected: last transition time not set")
	}
	if newReadyCondition.LastTransitionTime.Before(oldReadyCondition.LastTransitionTime) {
		t.Errorf("Unexpected: new transition time %s, is before old transition time %s", newReadyCondition.LastTransitionTime, oldReadyCondition.LastTransitionTime)
	}
}

func TestUnchangedStatus(t *testing.T) {
	syncer := newTestManager(&fake.Clientset{})
	testPod := getTestPod()
	podStatus := getRandomPodStatus()
	syncer.SetPodStatus(testPod, podStatus)
	syncer.SetPodStatus(testPod, podStatus)
	verifyUpdates(t, syncer, 1)
}

func TestUnchangedStatusPreservesLastTransitionTime(t *testing.T) {
	syncer := newTestManager(&fake.Clientset{})
	podStatus := getReadyPodStatus()
	pod := &api.Pod{
		ObjectMeta: api.ObjectMeta{
			UID:       "12345678",
			Name:      "foo",
			Namespace: "new",
		},
		Status: api.PodStatus{},
	}
	syncer.SetPodStatus(pod, podStatus)
	verifyUpdates(t, syncer, 1)
	oldStatus := expectPodStatus(t, syncer, pod)
	anotherStatus := getReadyPodStatus()
	syncer.SetPodStatus(pod, anotherStatus)
	// No update.
	verifyUpdates(t, syncer, 0)
	newStatus := expectPodStatus(t, syncer, pod)

	oldReadyCondition := api.GetPodReadyCondition(oldStatus)
	newReadyCondition := api.GetPodReadyCondition(newStatus)
	if newReadyCondition.LastTransitionTime.IsZero() {
		t.Errorf("Unexpected: last transition time not set")
	}
	if !oldReadyCondition.LastTransitionTime.Equal(newReadyCondition.LastTransitionTime) {
		t.Errorf("Unexpected: new transition time %s, is not equal to old transition time %s", newReadyCondition.LastTransitionTime, oldReadyCondition.LastTransitionTime)
	}
}

func TestSyncBatchIgnoresNotFound(t *testing.T) {
	client := fake.Clientset{}
	syncer := newTestManager(&client)
	client.AddReactor("get", "pods", func(action core.Action) (bool, runtime.Object, error) {
		return true, nil, errors.NewNotFound(api.Resource("pods"), "test-pod")
	})
	syncer.SetPodStatus(getTestPod(), getRandomPodStatus())
	syncer.testSyncBatch()

	verifyActions(t, syncer.kubeClient, []core.Action{
		core.GetActionImpl{ActionImpl: core.ActionImpl{Verb: "get", Resource: "pods"}},
	})
}

func TestSyncBatch(t *testing.T) {
	syncer := newTestManager(&fake.Clientset{})
	testPod := getTestPod()
	syncer.kubeClient = fake.NewSimpleClientset(testPod)
	syncer.SetPodStatus(testPod, getRandomPodStatus())
	syncer.testSyncBatch()
	verifyActions(t, syncer.kubeClient, []core.Action{
		core.GetActionImpl{ActionImpl: core.ActionImpl{Verb: "get", Resource: "pods"}},
		core.UpdateActionImpl{ActionImpl: core.ActionImpl{Verb: "update", Resource: "pods", Subresource: "status"}},
	},
	)
}

func TestSyncBatchChecksMismatchedUID(t *testing.T) {
	syncer := newTestManager(&fake.Clientset{})
	pod := getTestPod()
	pod.UID = "first"
	syncer.podManager.AddPod(pod)
	differentPod := getTestPod()
	differentPod.UID = "second"
	syncer.podManager.AddPod(differentPod)
	syncer.kubeClient = fake.NewSimpleClientset(pod)
	syncer.SetPodStatus(differentPod, getRandomPodStatus())
	syncer.testSyncBatch()
	verifyActions(t, syncer.kubeClient, []core.Action{
		core.GetActionImpl{ActionImpl: core.ActionImpl{Verb: "get", Resource: "pods"}},
	})
}

func TestSyncBatchNoDeadlock(t *testing.T) {
	client := &fake.Clientset{}
	m := newTestManager(client)
	pod := getTestPod()

	// Setup fake client.
	var ret api.Pod
	var err error
	client.AddReactor("*", "pods", func(action core.Action) (bool, runtime.Object, error) {
		switch action := action.(type) {
		case core.GetAction:
			assert.Equal(t, pod.Name, action.GetName(), "Unexpeted GetAction: %+v", action)
		case core.UpdateAction:
			assert.Equal(t, pod.Name, action.GetObject().(*api.Pod).Name, "Unexpeted UpdateAction: %+v", action)
		default:
			assert.Fail(t, "Unexpected Action: %+v", action)
		}
		return true, &ret, err
	})

	pod.Status.ContainerStatuses = []api.ContainerStatus{{State: api.ContainerState{Running: &api.ContainerStateRunning{}}}}

	getAction := core.GetActionImpl{ActionImpl: core.ActionImpl{Verb: "get", Resource: "pods"}}
	updateAction := core.UpdateActionImpl{ActionImpl: core.ActionImpl{Verb: "update", Resource: "pods", Subresource: "status"}}

	// Pod not found.
	ret = *pod
	err = errors.NewNotFound(api.Resource("pods"), pod.Name)
	m.SetPodStatus(pod, getRandomPodStatus())
	m.testSyncBatch()
	verifyActions(t, client, []core.Action{getAction})
	client.ClearActions()

	// Pod was recreated.
	ret.UID = "other_pod"
	err = nil
	m.SetPodStatus(pod, getRandomPodStatus())
	m.testSyncBatch()
	verifyActions(t, client, []core.Action{getAction})
	client.ClearActions()

	// Pod not deleted (success case).
	ret = *pod
	m.SetPodStatus(pod, getRandomPodStatus())
	m.testSyncBatch()
	verifyActions(t, client, []core.Action{getAction, updateAction})
	client.ClearActions()

	// Pod is terminated, but still running.
	pod.DeletionTimestamp = new(unversioned.Time)
	m.SetPodStatus(pod, getRandomPodStatus())
	m.testSyncBatch()
	verifyActions(t, client, []core.Action{getAction, updateAction})
	client.ClearActions()

	// Pod is terminated successfully.
	pod.Status.ContainerStatuses[0].State.Running = nil
	pod.Status.ContainerStatuses[0].State.Terminated = &api.ContainerStateTerminated{}
	m.SetPodStatus(pod, getRandomPodStatus())
	m.testSyncBatch()
	verifyActions(t, client, []core.Action{getAction, updateAction})
	client.ClearActions()

	// Error case.
	err = fmt.Errorf("intentional test error")
	m.SetPodStatus(pod, getRandomPodStatus())
	m.testSyncBatch()
	verifyActions(t, client, []core.Action{getAction})
	client.ClearActions()
}

func TestStaleUpdates(t *testing.T) {
	pod := getTestPod()
	client := fake.NewSimpleClientset(pod)
	m := newTestManager(client)

	status := api.PodStatus{Message: "initial status"}
	m.SetPodStatus(pod, status)
	status.Message = "first version bump"
	m.SetPodStatus(pod, status)
	status.Message = "second version bump"
	m.SetPodStatus(pod, status)
	verifyUpdates(t, m, 3)

	t.Logf("First sync pushes latest status.")
	m.testSyncBatch()
	verifyActions(t, m.kubeClient, []core.Action{
		core.GetActionImpl{ActionImpl: core.ActionImpl{Verb: "get", Resource: "pods"}},
		core.UpdateActionImpl{ActionImpl: core.ActionImpl{Verb: "update", Resource: "pods", Subresource: "status"}},
	})
	client.ClearActions()

	for i := 0; i < 2; i++ {
		t.Logf("Next 2 syncs should be ignored (%d).", i)
		m.testSyncBatch()
		verifyActions(t, m.kubeClient, []core.Action{})
	}

	t.Log("Unchanged status should not send an update.")
	m.SetPodStatus(pod, status)
	verifyUpdates(t, m, 0)

	t.Log("... unless it's stale.")
	m.apiStatusVersions[pod.UID] = m.apiStatusVersions[pod.UID] - 1

	m.SetPodStatus(pod, status)
	m.testSyncBatch()
	verifyActions(t, m.kubeClient, []core.Action{
		core.GetActionImpl{ActionImpl: core.ActionImpl{Verb: "get", Resource: "pods"}},
		core.UpdateActionImpl{ActionImpl: core.ActionImpl{Verb: "update", Resource: "pods", Subresource: "status"}},
	})

	// Nothing stuck in the pipe.
	verifyUpdates(t, m, 0)
}

// shuffle returns a new shuffled list of container statuses.
func shuffle(statuses []api.ContainerStatus) []api.ContainerStatus {
	numStatuses := len(statuses)
	randIndexes := rand.Perm(numStatuses)
	shuffled := make([]api.ContainerStatus, numStatuses)
	for i := 0; i < numStatuses; i++ {
		shuffled[i] = statuses[randIndexes[i]]
	}
	return shuffled
}

func TestStatusEquality(t *testing.T) {
	containerStatus := []api.ContainerStatus{}
	for i := 0; i < 10; i++ {
		s := api.ContainerStatus{
			Name: fmt.Sprintf("container%d", i),
		}
		containerStatus = append(containerStatus, s)
	}
	podStatus := api.PodStatus{
		ContainerStatuses: containerStatus,
	}
	for i := 0; i < 10; i++ {
		oldPodStatus := api.PodStatus{
			ContainerStatuses: shuffle(podStatus.ContainerStatuses),
		}
		normalizeStatus(&oldPodStatus)
		normalizeStatus(&podStatus)
		if !isStatusEqual(&oldPodStatus, &podStatus) {
			t.Fatalf("Order of container statuses should not affect normalized equality.")
		}
	}
}

func TestStaticPodStatus(t *testing.T) {
	staticPod := getTestPod()
	staticPod.Annotations = map[string]string{kubetypes.ConfigSourceAnnotationKey: "file"}
	mirrorPod := getTestPod()
	mirrorPod.UID = "mirror-12345678"
	mirrorPod.Annotations = map[string]string{
		kubetypes.ConfigSourceAnnotationKey: "api",
		kubetypes.ConfigMirrorAnnotationKey: "mirror",
	}
	client := fake.NewSimpleClientset(mirrorPod)
	m := newTestManager(client)
	m.podManager.AddPod(staticPod)
	m.podManager.AddPod(mirrorPod)
	// Verify setup.
	assert.True(t, kubepod.IsStaticPod(staticPod), "SetUp error: staticPod")
	assert.True(t, kubepod.IsMirrorPod(mirrorPod), "SetUp error: mirrorPod")
	assert.Equal(t, m.podManager.TranslatePodUID(mirrorPod.UID), staticPod.UID)

	status := getRandomPodStatus()
	now := unversioned.Now()
	status.StartTime = &now

	m.SetPodStatus(staticPod, status)
	retrievedStatus := expectPodStatus(t, m, staticPod)
	normalizeStatus(&status)
	assert.True(t, isStatusEqual(&status, &retrievedStatus), "Expected: %+v, Got: %+v", status, retrievedStatus)
	retrievedStatus, _ = m.GetPodStatus(mirrorPod.UID)
	assert.True(t, isStatusEqual(&status, &retrievedStatus), "Expected: %+v, Got: %+v", status, retrievedStatus)
	// Should translate mirrorPod / staticPod UID.
	m.testSyncBatch()
	verifyActions(t, m.kubeClient, []core.Action{
		core.GetActionImpl{ActionImpl: core.ActionImpl{Verb: "get", Resource: "pods"}},
		core.UpdateActionImpl{ActionImpl: core.ActionImpl{Verb: "update", Resource: "pods", Subresource: "status"}},
	})
	updateAction := client.Actions()[1].(core.UpdateActionImpl)
	updatedPod := updateAction.Object.(*api.Pod)
	assert.Equal(t, mirrorPod.UID, updatedPod.UID, "Expected mirrorPod (%q), but got %q", mirrorPod.UID, updatedPod.UID)
	assert.True(t, isStatusEqual(&status, &updatedPod.Status), "Expected: %+v, Got: %+v", status, updatedPod.Status)
	client.ClearActions()

	// No changes.
	m.testSyncBatch()
	verifyActions(t, m.kubeClient, []core.Action{})

	// Mirror pod identity changes.
	m.podManager.DeletePod(mirrorPod)
	mirrorPod.UID = "new-mirror-pod"
	mirrorPod.Status = api.PodStatus{}
	m.podManager.AddPod(mirrorPod)
	// Expect update to new mirrorPod.
	m.testSyncBatch()
	verifyActions(t, m.kubeClient, []core.Action{
		core.GetActionImpl{ActionImpl: core.ActionImpl{Verb: "get", Resource: "pods"}},
		core.UpdateActionImpl{ActionImpl: core.ActionImpl{Verb: "update", Resource: "pods", Subresource: "status"}},
	})
	updateAction = client.Actions()[1].(core.UpdateActionImpl)
	updatedPod = updateAction.Object.(*api.Pod)
	assert.Equal(t, mirrorPod.UID, updatedPod.UID, "Expected mirrorPod (%q), but got %q", mirrorPod.UID, updatedPod.UID)
	assert.True(t, isStatusEqual(&status, &updatedPod.Status), "Expected: %+v, Got: %+v", status, updatedPod.Status)
}

func TestSetContainerReadiness(t *testing.T) {
	cID1 := kubecontainer.ContainerID{"test", "1"}
	cID2 := kubecontainer.ContainerID{"test", "2"}
	containerStatuses := []api.ContainerStatus{
		{
			Name:        "c1",
			ContainerID: cID1.String(),
			Ready:       false,
		}, {
			Name:        "c2",
			ContainerID: cID2.String(),
			Ready:       false,
		},
	}
	status := api.PodStatus{
		ContainerStatuses: containerStatuses,
		Conditions: []api.PodCondition{{
			Type:   api.PodReady,
			Status: api.ConditionFalse,
		}},
	}
	pod := getTestPod()
	pod.Spec.Containers = []api.Container{{Name: "c1"}, {Name: "c2"}}

	// Verify expected readiness of containers & pod.
	verifyReadiness := func(step string, status *api.PodStatus, c1Ready, c2Ready, podReady bool) {
		for _, c := range status.ContainerStatuses {
			switch c.ContainerID {
			case cID1.String():
				if c.Ready != c1Ready {
					t.Errorf("[%s] Expected readiness of c1 to be %v but was %v", step, c1Ready, c.Ready)
				}
			case cID2.String():
				if c.Ready != c2Ready {
					t.Errorf("[%s] Expected readiness of c2 to be %v but was %v", step, c2Ready, c.Ready)
				}
			default:
				t.Fatalf("[%s] Unexpected container: %+v", step, c)
			}
		}
		if status.Conditions[0].Type != api.PodReady {
			t.Fatalf("[%s] Unexpected condition: %+v", step, status.Conditions[0])
		} else if ready := (status.Conditions[0].Status == api.ConditionTrue); ready != podReady {
			t.Errorf("[%s] Expected readiness of pod to be %v but was %v", step, podReady, ready)
		}
	}

	m := newTestManager(&fake.Clientset{})
	// Add test pod because the container spec has been changed.
	m.podManager.AddPod(pod)

	t.Log("Setting readiness before status should fail.")
	m.SetContainerReadiness(pod.UID, cID1, true)
	verifyUpdates(t, m, 0)
	if status, ok := m.GetPodStatus(pod.UID); ok {
		t.Errorf("Unexpected PodStatus: %+v", status)
	}

	t.Log("Setting initial status.")
	m.SetPodStatus(pod, status)
	verifyUpdates(t, m, 1)
	status = expectPodStatus(t, m, pod)
	verifyReadiness("initial", &status, false, false, false)

	t.Log("Setting unchanged readiness should do nothing.")
	m.SetContainerReadiness(pod.UID, cID1, false)
	verifyUpdates(t, m, 0)
	status = expectPodStatus(t, m, pod)
	verifyReadiness("unchanged", &status, false, false, false)

	t.Log("Setting container readiness should generate update but not pod readiness.")
	m.SetContainerReadiness(pod.UID, cID1, true)
	verifyUpdates(t, m, 1)
	status = expectPodStatus(t, m, pod)
	verifyReadiness("c1 ready", &status, true, false, false)

	t.Log("Setting both containers to ready should update pod readiness.")
	m.SetContainerReadiness(pod.UID, cID2, true)
	verifyUpdates(t, m, 1)
	status = expectPodStatus(t, m, pod)
	verifyReadiness("all ready", &status, true, true, true)

	t.Log("Setting non-existant container readiness should fail.")
	m.SetContainerReadiness(pod.UID, kubecontainer.ContainerID{"test", "foo"}, true)
	verifyUpdates(t, m, 0)
	status = expectPodStatus(t, m, pod)
	verifyReadiness("ignore non-existant", &status, true, true, true)
}

func TestSyncBatchCleanupVersions(t *testing.T) {
	m := newTestManager(&fake.Clientset{})
	testPod := getTestPod()
	mirrorPod := getTestPod()
	mirrorPod.UID = "mirror-uid"
	mirrorPod.Name = "mirror_pod"
	mirrorPod.Annotations = map[string]string{
		kubetypes.ConfigSourceAnnotationKey: "api",
		kubetypes.ConfigMirrorAnnotationKey: "mirror",
	}

	// Orphaned pods should be removed.
	m.apiStatusVersions[testPod.UID] = 100
	m.apiStatusVersions[mirrorPod.UID] = 200
	m.testSyncBatch()
	if _, ok := m.apiStatusVersions[testPod.UID]; ok {
		t.Errorf("Should have cleared status for testPod")
	}
	if _, ok := m.apiStatusVersions[mirrorPod.UID]; ok {
		t.Errorf("Should have cleared status for mirrorPod")
	}

	// Non-orphaned pods should not be removed.
	m.SetPodStatus(testPod, getRandomPodStatus())
	m.podManager.AddPod(mirrorPod)
	staticPod := mirrorPod
	staticPod.UID = "static-uid"
	staticPod.Annotations = map[string]string{kubetypes.ConfigSourceAnnotationKey: "file"}
	m.podManager.AddPod(staticPod)
	m.apiStatusVersions[testPod.UID] = 100
	m.apiStatusVersions[mirrorPod.UID] = 200
	m.testSyncBatch()
	if _, ok := m.apiStatusVersions[testPod.UID]; !ok {
		t.Errorf("Should not have cleared status for testPod")
	}
	if _, ok := m.apiStatusVersions[mirrorPod.UID]; !ok {
		t.Errorf("Should not have cleared status for mirrorPod")
	}
}

func TestReconcilePodStatus(t *testing.T) {
	testPod := getTestPod()
	client := fake.NewSimpleClientset(testPod)
	syncer := newTestManager(client)
	syncer.SetPodStatus(testPod, getRandomPodStatus())
	// Call syncBatch directly to test reconcile
	syncer.syncBatch() // The apiStatusVersions should be set now

	podStatus, ok := syncer.GetPodStatus(testPod.UID)
	if !ok {
		t.Fatalf("Should find pod status for pod: %+v", testPod)
	}
	testPod.Status = podStatus

	// If the pod status is the same, a reconciliation is not needed,
	// syncBatch should do nothing
	syncer.podManager.UpdatePod(testPod)
	if syncer.needsReconcile(testPod.UID, podStatus) {
		t.Errorf("Pod status is the same, a reconciliation is not needed")
	}
	client.ClearActions()
	syncer.syncBatch()
	verifyActions(t, client, []core.Action{})

	// If the pod status is the same, only the timestamp is in Rfc3339 format (lower precision without nanosecond),
	// a reconciliation is not needed, syncBatch should do nothing.
	// The StartTime should have been set in SetPodStatus().
	// TODO(random-liu): Remove this later when api becomes consistent for timestamp.
	normalizedStartTime := testPod.Status.StartTime.Rfc3339Copy()
	testPod.Status.StartTime = &normalizedStartTime
	syncer.podManager.UpdatePod(testPod)
	if syncer.needsReconcile(testPod.UID, podStatus) {
		t.Errorf("Pod status only differs for timestamp format, a reconciliation is not needed")
	}
	client.ClearActions()
	syncer.syncBatch()
	verifyActions(t, client, []core.Action{})

	// If the pod status is different, a reconciliation is needed, syncBatch should trigger an update
	testPod.Status = getRandomPodStatus()
	syncer.podManager.UpdatePod(testPod)
	if !syncer.needsReconcile(testPod.UID, podStatus) {
		t.Errorf("Pod status is different, a reconciliation is needed")
	}
	client.ClearActions()
	syncer.syncBatch()
	verifyActions(t, client, []core.Action{
		core.GetActionImpl{ActionImpl: core.ActionImpl{Verb: "get", Resource: "pods"}},
		core.UpdateActionImpl{ActionImpl: core.ActionImpl{Verb: "update", Resource: "pods", Subresource: "status"}},
	})
}

func expectPodStatus(t *testing.T, m *manager, pod *api.Pod) api.PodStatus {
	status, ok := m.GetPodStatus(pod.UID)
	if !ok {
		t.Fatalf("Expected PodStatus for %q not found", pod.UID)
	}
	return status
}

func TestDeletePods(t *testing.T) {
	pod := getTestPod()
	// Set the deletion timestamp.
	pod.DeletionTimestamp = new(unversioned.Time)
	client := fake.NewSimpleClientset(pod)
	m := newTestManager(client)
	m.podManager.AddPod(pod)

	status := getRandomPodStatus()
	now := unversioned.Now()
	status.StartTime = &now
	m.SetPodStatus(pod, status)

	m.testSyncBatch()
	// Expect to see an delete action.
	verifyActions(t, m.kubeClient, []core.Action{
		core.GetActionImpl{ActionImpl: core.ActionImpl{Verb: "get", Resource: "pods"}},
		core.UpdateActionImpl{ActionImpl: core.ActionImpl{Verb: "update", Resource: "pods", Subresource: "status"}},
		core.DeleteActionImpl{ActionImpl: core.ActionImpl{Verb: "delete", Resource: "pods"}},
	})
}

func TestDoNotDeleteMirrorPods(t *testing.T) {
	staticPod := getTestPod()
	staticPod.Annotations = map[string]string{kubetypes.ConfigSourceAnnotationKey: "file"}
	mirrorPod := getTestPod()
	mirrorPod.UID = "mirror-12345678"
	mirrorPod.Annotations = map[string]string{
		kubetypes.ConfigSourceAnnotationKey: "api",
		kubetypes.ConfigMirrorAnnotationKey: "mirror",
	}
	// Set the deletion timestamp.
	mirrorPod.DeletionTimestamp = new(unversioned.Time)
	client := fake.NewSimpleClientset(mirrorPod)
	m := newTestManager(client)
	m.podManager.AddPod(staticPod)
	m.podManager.AddPod(mirrorPod)
	// Verify setup.
	assert.True(t, kubepod.IsStaticPod(staticPod), "SetUp error: staticPod")
	assert.True(t, kubepod.IsMirrorPod(mirrorPod), "SetUp error: mirrorPod")
	assert.Equal(t, m.podManager.TranslatePodUID(mirrorPod.UID), staticPod.UID)

	status := getRandomPodStatus()
	now := unversioned.Now()
	status.StartTime = &now
	m.SetPodStatus(staticPod, status)

	m.testSyncBatch()
	// Expect not to see an delete action.
	verifyActions(t, m.kubeClient, []core.Action{
		core.GetActionImpl{ActionImpl: core.ActionImpl{Verb: "get", Resource: "pods"}},
		core.UpdateActionImpl{ActionImpl: core.ActionImpl{Verb: "update", Resource: "pods", Subresource: "status"}},
	})
}
