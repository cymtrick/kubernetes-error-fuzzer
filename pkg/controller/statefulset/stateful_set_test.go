/*
Copyright 2016 The Kubernetes Authors.

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

package statefulset

import (
	"bytes"
	"context"
	"encoding/json"
	"sort"
	"testing"

	apps "k8s.io/api/apps/v1"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/apimachinery/pkg/util/sets"
	"k8s.io/client-go/informers"
	"k8s.io/client-go/kubernetes/fake"
	"k8s.io/client-go/tools/cache"
	"k8s.io/client-go/tools/record"
	"k8s.io/kubernetes/pkg/controller"
	"k8s.io/kubernetes/pkg/controller/history"
)

var parentKind = apps.SchemeGroupVersion.WithKind("StatefulSet")

func alwaysReady() bool { return true }

func TestStatefulSetControllerCreates(t *testing.T) {
	set := newStatefulSet(3)
	ssc, spc, _ := newFakeStatefulSetController(set)
	if err := scaleUpStatefulSetController(set, ssc, spc); err != nil {
		t.Errorf("Failed to turn up StatefulSet : %s", err)
	}
	if obj, _, err := spc.setsIndexer.Get(set); err != nil {
		t.Error(err)
	} else {
		set = obj.(*apps.StatefulSet)
	}
	if set.Status.Replicas != 3 {
		t.Errorf("set.Status.Replicas = %v; want 3", set.Status.Replicas)
	}
}

func TestStatefulSetControllerDeletes(t *testing.T) {
	set := newStatefulSet(3)
	ssc, spc, _ := newFakeStatefulSetController(set)
	if err := scaleUpStatefulSetController(set, ssc, spc); err != nil {
		t.Errorf("Failed to turn up StatefulSet : %s", err)
	}
	if obj, _, err := spc.setsIndexer.Get(set); err != nil {
		t.Error(err)
	} else {
		set = obj.(*apps.StatefulSet)
	}
	if set.Status.Replicas != 3 {
		t.Errorf("set.Status.Replicas = %v; want 3", set.Status.Replicas)
	}
	*set.Spec.Replicas = 0
	if err := scaleDownStatefulSetController(set, ssc, spc); err != nil {
		t.Errorf("Failed to turn down StatefulSet : %s", err)
	}
	if obj, _, err := spc.setsIndexer.Get(set); err != nil {
		t.Error(err)
	} else {
		set = obj.(*apps.StatefulSet)
	}
	if set.Status.Replicas != 0 {
		t.Errorf("set.Status.Replicas = %v; want 0", set.Status.Replicas)
	}
}

func TestStatefulSetControllerRespectsTermination(t *testing.T) {
	set := newStatefulSet(3)
	ssc, spc, _ := newFakeStatefulSetController(set)
	if err := scaleUpStatefulSetController(set, ssc, spc); err != nil {
		t.Errorf("Failed to turn up StatefulSet : %s", err)
	}
	if obj, _, err := spc.setsIndexer.Get(set); err != nil {
		t.Error(err)
	} else {
		set = obj.(*apps.StatefulSet)
	}
	if set.Status.Replicas != 3 {
		t.Errorf("set.Status.Replicas = %v; want 3", set.Status.Replicas)
	}
	_, err := spc.addTerminatingPod(set, 3)
	if err != nil {
		t.Error(err)
	}
	pods, err := spc.addTerminatingPod(set, 4)
	if err != nil {
		t.Error(err)
	}
	ssc.syncStatefulSet(context.TODO(), set, pods)
	selector, err := metav1.LabelSelectorAsSelector(set.Spec.Selector)
	if err != nil {
		t.Error(err)
	}
	pods, err = spc.podsLister.Pods(set.Namespace).List(selector)
	if err != nil {
		t.Error(err)
	}
	if len(pods) != 5 {
		t.Error("StatefulSet does not respect termination")
	}
	sort.Sort(ascendingOrdinal(pods))
	spc.DeleteStatefulPod(set, pods[3])
	spc.DeleteStatefulPod(set, pods[4])
	*set.Spec.Replicas = 0
	if err := scaleDownStatefulSetController(set, ssc, spc); err != nil {
		t.Errorf("Failed to turn down StatefulSet : %s", err)
	}
	if obj, _, err := spc.setsIndexer.Get(set); err != nil {
		t.Error(err)
	} else {
		set = obj.(*apps.StatefulSet)
	}
	if set.Status.Replicas != 0 {
		t.Errorf("set.Status.Replicas = %v; want 0", set.Status.Replicas)
	}
}

func TestStatefulSetControllerBlocksScaling(t *testing.T) {
	set := newStatefulSet(3)
	ssc, spc, _ := newFakeStatefulSetController(set)
	if err := scaleUpStatefulSetController(set, ssc, spc); err != nil {
		t.Errorf("Failed to turn up StatefulSet : %s", err)
	}
	if obj, _, err := spc.setsIndexer.Get(set); err != nil {
		t.Error(err)
	} else {
		set = obj.(*apps.StatefulSet)
	}
	if set.Status.Replicas != 3 {
		t.Errorf("set.Status.Replicas = %v; want 3", set.Status.Replicas)
	}
	*set.Spec.Replicas = 5
	fakeResourceVersion(set)
	spc.setsIndexer.Update(set)
	_, err := spc.setPodTerminated(set, 0)
	if err != nil {
		t.Error("Failed to set pod terminated at ordinal 0")
	}
	ssc.enqueueStatefulSet(set)
	fakeWorker(ssc)
	selector, err := metav1.LabelSelectorAsSelector(set.Spec.Selector)
	if err != nil {
		t.Error(err)
	}
	pods, err := spc.podsLister.Pods(set.Namespace).List(selector)
	if err != nil {
		t.Error(err)
	}
	if len(pods) != 3 {
		t.Error("StatefulSet does not block scaling")
	}
	sort.Sort(ascendingOrdinal(pods))
	spc.DeleteStatefulPod(set, pods[0])
	ssc.enqueueStatefulSet(set)
	fakeWorker(ssc)
	pods, err = spc.podsLister.Pods(set.Namespace).List(selector)
	if err != nil {
		t.Error(err)
	}
	if len(pods) != 3 {
		t.Error("StatefulSet does not resume when terminated Pod is removed")
	}
}

func TestStatefulSetControllerDeletionTimestamp(t *testing.T) {
	set := newStatefulSet(3)
	set.DeletionTimestamp = new(metav1.Time)
	ssc, spc, _ := newFakeStatefulSetController(set)

	spc.setsIndexer.Add(set)

	// Force a sync. It should not try to create any Pods.
	ssc.enqueueStatefulSet(set)
	fakeWorker(ssc)

	selector, err := metav1.LabelSelectorAsSelector(set.Spec.Selector)
	if err != nil {
		t.Fatal(err)
	}
	pods, err := spc.podsLister.Pods(set.Namespace).List(selector)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := len(pods), 0; got != want {
		t.Errorf("len(pods) = %v, want %v", got, want)
	}
}

func TestStatefulSetControllerDeletionTimestampRace(t *testing.T) {
	set := newStatefulSet(3)
	// The bare client says it IS deleted.
	set.DeletionTimestamp = new(metav1.Time)
	ssc, spc, ssh := newFakeStatefulSetController(set)

	// The lister (cache) says it's NOT deleted.
	set2 := *set
	set2.DeletionTimestamp = nil
	spc.setsIndexer.Add(&set2)

	// The recheck occurs in the presence of a matching orphan.
	pod := newStatefulSetPod(set, 1)
	pod.OwnerReferences = nil
	spc.podsIndexer.Add(pod)
	set.Status.CollisionCount = new(int32)
	revision, err := newRevision(set, 1, set.Status.CollisionCount)
	if err != nil {
		t.Fatal(err)
	}
	revision.OwnerReferences = nil
	_, err = ssh.CreateControllerRevision(set, revision, set.Status.CollisionCount)
	if err != nil {
		t.Fatal(err)
	}

	// Force a sync. It should not try to create any Pods.
	ssc.enqueueStatefulSet(set)
	fakeWorker(ssc)

	selector, err := metav1.LabelSelectorAsSelector(set.Spec.Selector)
	if err != nil {
		t.Fatal(err)
	}
	pods, err := spc.podsLister.Pods(set.Namespace).List(selector)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := len(pods), 1; got != want {
		t.Errorf("len(pods) = %v, want %v", got, want)
	}

	// It should not adopt pods.
	for _, pod := range pods {
		if len(pod.OwnerReferences) > 0 {
			t.Errorf("unexpected pod owner references: %v", pod.OwnerReferences)
		}
	}

	// It should not adopt revisions.
	revisions, err := ssh.ListControllerRevisions(set, selector)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := len(revisions), 1; got != want {
		t.Errorf("len(revisions) = %v, want %v", got, want)
	}
	for _, revision := range revisions {
		if len(revision.OwnerReferences) > 0 {
			t.Errorf("unexpected revision owner references: %v", revision.OwnerReferences)
		}
	}
}

func TestStatefulSetControllerAddPod(t *testing.T) {
	ssc, spc, _ := newFakeStatefulSetController()
	set1 := newStatefulSet(3)
	set2 := newStatefulSet(3)
	pod1 := newStatefulSetPod(set1, 0)
	pod2 := newStatefulSetPod(set2, 0)
	spc.setsIndexer.Add(set1)
	spc.setsIndexer.Add(set2)

	ssc.addPod(pod1)
	key, done := ssc.queue.Get()
	if key == nil || done {
		t.Error("failed to enqueue StatefulSet")
	} else if key, ok := key.(string); !ok {
		t.Error("key is not a string")
	} else if expectedKey, _ := controller.KeyFunc(set1); expectedKey != key {
		t.Errorf("expected StatefulSet key %s found %s", expectedKey, key)
	}
	ssc.queue.Done(key)

	ssc.addPod(pod2)
	key, done = ssc.queue.Get()
	if key == nil || done {
		t.Error("failed to enqueue StatefulSet")
	} else if key, ok := key.(string); !ok {
		t.Error("key is not a string")
	} else if expectedKey, _ := controller.KeyFunc(set2); expectedKey != key {
		t.Errorf("expected StatefulSet key %s found %s", expectedKey, key)
	}
	ssc.queue.Done(key)
}

func TestStatefulSetControllerAddPodOrphan(t *testing.T) {
	ssc, spc, _ := newFakeStatefulSetController()
	set1 := newStatefulSet(3)
	set2 := newStatefulSet(3)
	set2.Name = "foo2"
	set3 := newStatefulSet(3)
	set3.Name = "foo3"
	set3.Spec.Selector.MatchLabels = map[string]string{"foo3": "bar"}
	pod := newStatefulSetPod(set1, 0)
	spc.setsIndexer.Add(set1)
	spc.setsIndexer.Add(set2)
	spc.setsIndexer.Add(set3)

	// Make pod an orphan. Expect matching sets to be queued.
	pod.OwnerReferences = nil
	ssc.addPod(pod)
	if got, want := ssc.queue.Len(), 2; got != want {
		t.Errorf("queue.Len() = %v, want %v", got, want)
	}
}

func TestStatefulSetControllerAddPodNoSet(t *testing.T) {
	ssc, _, _ := newFakeStatefulSetController()
	set := newStatefulSet(3)
	pod := newStatefulSetPod(set, 0)
	ssc.addPod(pod)
	ssc.queue.ShutDown()
	key, _ := ssc.queue.Get()
	if key != nil {
		t.Errorf("StatefulSet enqueued key for Pod with no Set %s", key)
	}
}

func TestStatefulSetControllerUpdatePod(t *testing.T) {
	ssc, spc, _ := newFakeStatefulSetController()
	set1 := newStatefulSet(3)
	set2 := newStatefulSet(3)
	set2.Name = "foo2"
	pod1 := newStatefulSetPod(set1, 0)
	pod2 := newStatefulSetPod(set2, 0)
	spc.setsIndexer.Add(set1)
	spc.setsIndexer.Add(set2)

	prev := *pod1
	fakeResourceVersion(pod1)
	ssc.updatePod(&prev, pod1)
	key, done := ssc.queue.Get()
	if key == nil || done {
		t.Error("failed to enqueue StatefulSet")
	} else if key, ok := key.(string); !ok {
		t.Error("key is not a string")
	} else if expectedKey, _ := controller.KeyFunc(set1); expectedKey != key {
		t.Errorf("expected StatefulSet key %s found %s", expectedKey, key)
	}

	prev = *pod2
	fakeResourceVersion(pod2)
	ssc.updatePod(&prev, pod2)
	key, done = ssc.queue.Get()
	if key == nil || done {
		t.Error("failed to enqueue StatefulSet")
	} else if key, ok := key.(string); !ok {
		t.Error("key is not a string")
	} else if expectedKey, _ := controller.KeyFunc(set2); expectedKey != key {
		t.Errorf("expected StatefulSet key %s found %s", expectedKey, key)
	}
}

func TestStatefulSetControllerUpdatePodWithNoSet(t *testing.T) {
	ssc, _, _ := newFakeStatefulSetController()
	set := newStatefulSet(3)
	pod := newStatefulSetPod(set, 0)
	prev := *pod
	fakeResourceVersion(pod)
	ssc.updatePod(&prev, pod)
	ssc.queue.ShutDown()
	key, _ := ssc.queue.Get()
	if key != nil {
		t.Errorf("StatefulSet enqueued key for Pod with no Set %s", key)
	}
}

func TestStatefulSetControllerUpdatePodWithSameVersion(t *testing.T) {
	ssc, spc, _ := newFakeStatefulSetController()
	set := newStatefulSet(3)
	pod := newStatefulSetPod(set, 0)
	spc.setsIndexer.Add(set)
	ssc.updatePod(pod, pod)
	ssc.queue.ShutDown()
	key, _ := ssc.queue.Get()
	if key != nil {
		t.Errorf("StatefulSet enqueued key for Pod with no Set %s", key)
	}
}

func TestStatefulSetControllerUpdatePodOrphanWithNewLabels(t *testing.T) {
	ssc, spc, _ := newFakeStatefulSetController()
	set := newStatefulSet(3)
	pod := newStatefulSetPod(set, 0)
	pod.OwnerReferences = nil
	set2 := newStatefulSet(3)
	set2.Name = "foo2"
	spc.setsIndexer.Add(set)
	spc.setsIndexer.Add(set2)
	clone := *pod
	clone.Labels = map[string]string{"foo2": "bar2"}
	fakeResourceVersion(&clone)
	ssc.updatePod(&clone, pod)
	if got, want := ssc.queue.Len(), 2; got != want {
		t.Errorf("queue.Len() = %v, want %v", got, want)
	}
}

func TestStatefulSetControllerUpdatePodChangeControllerRef(t *testing.T) {
	ssc, spc, _ := newFakeStatefulSetController()
	set := newStatefulSet(3)
	set2 := newStatefulSet(3)
	set2.Name = "foo2"
	pod := newStatefulSetPod(set, 0)
	pod2 := newStatefulSetPod(set2, 0)
	spc.setsIndexer.Add(set)
	spc.setsIndexer.Add(set2)
	clone := *pod
	clone.OwnerReferences = pod2.OwnerReferences
	fakeResourceVersion(&clone)
	ssc.updatePod(&clone, pod)
	if got, want := ssc.queue.Len(), 2; got != want {
		t.Errorf("queue.Len() = %v, want %v", got, want)
	}
}

func TestStatefulSetControllerUpdatePodRelease(t *testing.T) {
	ssc, spc, _ := newFakeStatefulSetController()
	set := newStatefulSet(3)
	set2 := newStatefulSet(3)
	set2.Name = "foo2"
	pod := newStatefulSetPod(set, 0)
	spc.setsIndexer.Add(set)
	spc.setsIndexer.Add(set2)
	clone := *pod
	clone.OwnerReferences = nil
	fakeResourceVersion(&clone)
	ssc.updatePod(pod, &clone)
	if got, want := ssc.queue.Len(), 2; got != want {
		t.Errorf("queue.Len() = %v, want %v", got, want)
	}
}

func TestStatefulSetControllerDeletePod(t *testing.T) {
	ssc, spc, _ := newFakeStatefulSetController()
	set1 := newStatefulSet(3)
	set2 := newStatefulSet(3)
	set2.Name = "foo2"
	pod1 := newStatefulSetPod(set1, 0)
	pod2 := newStatefulSetPod(set2, 0)
	spc.setsIndexer.Add(set1)
	spc.setsIndexer.Add(set2)

	ssc.deletePod(pod1)
	key, done := ssc.queue.Get()
	if key == nil || done {
		t.Error("failed to enqueue StatefulSet")
	} else if key, ok := key.(string); !ok {
		t.Error("key is not a string")
	} else if expectedKey, _ := controller.KeyFunc(set1); expectedKey != key {
		t.Errorf("expected StatefulSet key %s found %s", expectedKey, key)
	}

	ssc.deletePod(pod2)
	key, done = ssc.queue.Get()
	if key == nil || done {
		t.Error("failed to enqueue StatefulSet")
	} else if key, ok := key.(string); !ok {
		t.Error("key is not a string")
	} else if expectedKey, _ := controller.KeyFunc(set2); expectedKey != key {
		t.Errorf("expected StatefulSet key %s found %s", expectedKey, key)
	}
}

func TestStatefulSetControllerDeletePodOrphan(t *testing.T) {
	ssc, spc, _ := newFakeStatefulSetController()
	set1 := newStatefulSet(3)
	set2 := newStatefulSet(3)
	set2.Name = "foo2"
	pod1 := newStatefulSetPod(set1, 0)
	spc.setsIndexer.Add(set1)
	spc.setsIndexer.Add(set2)

	pod1.OwnerReferences = nil
	ssc.deletePod(pod1)
	if got, want := ssc.queue.Len(), 0; got != want {
		t.Errorf("queue.Len() = %v, want %v", got, want)
	}
}

func TestStatefulSetControllerDeletePodTombstone(t *testing.T) {
	ssc, spc, _ := newFakeStatefulSetController()
	set := newStatefulSet(3)
	pod := newStatefulSetPod(set, 0)
	spc.setsIndexer.Add(set)
	tombstoneKey, _ := controller.KeyFunc(pod)
	tombstone := cache.DeletedFinalStateUnknown{Key: tombstoneKey, Obj: pod}
	ssc.deletePod(tombstone)
	key, done := ssc.queue.Get()
	if key == nil || done {
		t.Error("failed to enqueue StatefulSet")
	} else if key, ok := key.(string); !ok {
		t.Error("key is not a string")
	} else if expectedKey, _ := controller.KeyFunc(set); expectedKey != key {
		t.Errorf("expected StatefulSet key %s found %s", expectedKey, key)
	}
}

func TestStatefulSetControllerGetStatefulSetsForPod(t *testing.T) {
	ssc, spc, _ := newFakeStatefulSetController()
	set1 := newStatefulSet(3)
	set2 := newStatefulSet(3)
	set2.Name = "foo2"
	pod := newStatefulSetPod(set1, 0)
	spc.setsIndexer.Add(set1)
	spc.setsIndexer.Add(set2)
	spc.podsIndexer.Add(pod)
	sets := ssc.getStatefulSetsForPod(pod)
	if got, want := len(sets), 2; got != want {
		t.Errorf("len(sets) = %v, want %v", got, want)
	}
}

func TestGetPodsForStatefulSetAdopt(t *testing.T) {
	set := newStatefulSet(5)
	pod1 := newStatefulSetPod(set, 1)
	// pod2 is an orphan with matching labels and name.
	pod2 := newStatefulSetPod(set, 2)
	pod2.OwnerReferences = nil
	// pod3 has wrong labels.
	pod3 := newStatefulSetPod(set, 3)
	pod3.OwnerReferences = nil
	pod3.Labels = nil
	// pod4 has wrong name.
	pod4 := newStatefulSetPod(set, 4)
	pod4.OwnerReferences = nil
	pod4.Name = "x" + pod4.Name

	ssc, spc, _ := newFakeStatefulSetController(set, pod1, pod2, pod3, pod4)

	spc.podsIndexer.Add(pod1)
	spc.podsIndexer.Add(pod2)
	spc.podsIndexer.Add(pod3)
	spc.podsIndexer.Add(pod4)
	selector, err := metav1.LabelSelectorAsSelector(set.Spec.Selector)
	if err != nil {
		t.Fatal(err)
	}
	pods, err := ssc.getPodsForStatefulSet(context.TODO(), set, selector)
	if err != nil {
		t.Fatalf("getPodsForStatefulSet() error: %v", err)
	}
	got := sets.NewString()
	for _, pod := range pods {
		got.Insert(pod.Name)
	}
	// pod2 should be claimed, pod3 and pod4 ignored
	want := sets.NewString(pod1.Name, pod2.Name)
	if !got.Equal(want) {
		t.Errorf("getPodsForStatefulSet() = %v, want %v", got, want)
	}
}

func TestAdoptOrphanRevisions(t *testing.T) {
	ss1 := newStatefulSetWithLabels(3, "ss1", types.UID("ss1"), map[string]string{"foo": "bar"})
	ss1.Status.CollisionCount = new(int32)
	ss1Rev1, err := history.NewControllerRevision(ss1, parentKind, ss1.Spec.Template.Labels, rawTemplate(&ss1.Spec.Template), 1, ss1.Status.CollisionCount)
	if err != nil {
		t.Fatal(err)
	}
	ss1Rev1.Namespace = ss1.Namespace
	ss1.Spec.Template.Annotations = make(map[string]string)
	ss1.Spec.Template.Annotations["ss1"] = "ss1"
	ss1Rev2, err := history.NewControllerRevision(ss1, parentKind, ss1.Spec.Template.Labels, rawTemplate(&ss1.Spec.Template), 2, ss1.Status.CollisionCount)
	if err != nil {
		t.Fatal(err)
	}
	ss1Rev2.Namespace = ss1.Namespace
	ss1Rev2.OwnerReferences = []metav1.OwnerReference{}

	ssc, spc, _ := newFakeStatefulSetController(ss1, ss1Rev1, ss1Rev2)

	spc.revisionsIndexer.Add(ss1Rev1)
	spc.revisionsIndexer.Add(ss1Rev2)

	err = ssc.adoptOrphanRevisions(context.TODO(), ss1)
	if err != nil {
		t.Errorf("adoptOrphanRevisions() error: %v", err)
	}

	if revisions, err := ssc.control.ListRevisions(ss1); err != nil {
		t.Errorf("ListRevisions() error: %v", err)
	} else {
		var adopted bool
		for i := range revisions {
			if revisions[i].Name == ss1Rev2.Name && metav1.GetControllerOf(revisions[i]) != nil {
				adopted = true
			}
		}
		if !adopted {
			t.Error("adoptOrphanRevisions() not adopt orphan revisions")
		}
	}
}

func TestGetPodsForStatefulSetRelease(t *testing.T) {
	set := newStatefulSet(3)
	ssc, spc, _ := newFakeStatefulSetController(set)
	pod1 := newStatefulSetPod(set, 1)
	// pod2 is owned but has wrong name.
	pod2 := newStatefulSetPod(set, 2)
	pod2.Name = "x" + pod2.Name
	// pod3 is owned but has wrong labels.
	pod3 := newStatefulSetPod(set, 3)
	pod3.Labels = nil
	// pod4 is an orphan that doesn't match.
	pod4 := newStatefulSetPod(set, 4)
	pod4.OwnerReferences = nil
	pod4.Labels = nil

	spc.podsIndexer.Add(pod1)
	spc.podsIndexer.Add(pod2)
	spc.podsIndexer.Add(pod3)
	selector, err := metav1.LabelSelectorAsSelector(set.Spec.Selector)
	if err != nil {
		t.Fatal(err)
	}
	pods, err := ssc.getPodsForStatefulSet(context.TODO(), set, selector)
	if err != nil {
		t.Fatalf("getPodsForStatefulSet() error: %v", err)
	}
	got := sets.NewString()
	for _, pod := range pods {
		got.Insert(pod.Name)
	}

	// Expect only pod1 (pod2 and pod3 should be released, pod4 ignored).
	want := sets.NewString(pod1.Name)
	if !got.Equal(want) {
		t.Errorf("getPodsForStatefulSet() = %v, want %v", got, want)
	}
}

func newFakeStatefulSetController(initialObjects ...runtime.Object) (*StatefulSetController, *fakeStatefulPodControl, history.Interface) {
	client := fake.NewSimpleClientset(initialObjects...)
	informerFactory := informers.NewSharedInformerFactory(client, controller.NoResyncPeriodFunc())
	fpc := newFakeStatefulPodControl(informerFactory.Core().V1().Pods(), informerFactory.Apps().V1().StatefulSets(), informerFactory.Apps().V1().ControllerRevisions())
	ssu := newFakeStatefulSetStatusUpdater(informerFactory.Apps().V1().StatefulSets())
	ssc := NewStatefulSetController(
		informerFactory.Core().V1().Pods(),
		informerFactory.Apps().V1().StatefulSets(),
		informerFactory.Core().V1().PersistentVolumeClaims(),
		informerFactory.Apps().V1().ControllerRevisions(),
		client,
	)
	ssh := history.NewFakeHistory(informerFactory.Apps().V1().ControllerRevisions())
	ssc.podListerSynced = alwaysReady
	ssc.setListerSynced = alwaysReady
	recorder := record.NewFakeRecorder(10)
	ssc.control = NewDefaultStatefulSetControl(fpc, ssu, ssh, recorder)

	return ssc, fpc, ssh
}

func fakeWorker(ssc *StatefulSetController) {
	if obj, done := ssc.queue.Get(); !done {
		ssc.sync(context.TODO(), obj.(string))
		ssc.queue.Done(obj)
	}
}

func getPodAtOrdinal(pods []*v1.Pod, ordinal int) *v1.Pod {
	if 0 > ordinal || ordinal >= len(pods) {
		return nil
	}
	sort.Sort(ascendingOrdinal(pods))
	return pods[ordinal]
}

func scaleUpStatefulSetController(set *apps.StatefulSet, ssc *StatefulSetController, spc *fakeStatefulPodControl) error {
	spc.setsIndexer.Add(set)
	ssc.enqueueStatefulSet(set)
	fakeWorker(ssc)
	selector, err := metav1.LabelSelectorAsSelector(set.Spec.Selector)
	if err != nil {
		return err
	}
	for set.Status.ReadyReplicas < *set.Spec.Replicas {
		pods, err := spc.podsLister.Pods(set.Namespace).List(selector)
		if err != nil {
			return err
		}
		ord := len(pods) - 1
		if pods, err = spc.setPodPending(set, ord); err != nil {
			return err
		}
		pod := getPodAtOrdinal(pods, ord)
		ssc.addPod(pod)
		fakeWorker(ssc)
		pod = getPodAtOrdinal(pods, ord)
		prev := *pod
		if pods, err = spc.setPodRunning(set, ord); err != nil {
			return err
		}
		pod = getPodAtOrdinal(pods, ord)
		ssc.updatePod(&prev, pod)
		fakeWorker(ssc)
		pod = getPodAtOrdinal(pods, ord)
		prev = *pod
		if pods, err = spc.setPodReady(set, ord); err != nil {
			return err
		}
		pod = getPodAtOrdinal(pods, ord)
		ssc.updatePod(&prev, pod)
		fakeWorker(ssc)
		if err := assertMonotonicInvariants(set, spc); err != nil {
			return err
		}
		obj, _, err := spc.setsIndexer.Get(set)
		if err != nil {
			return err
		}
		set = obj.(*apps.StatefulSet)

	}
	return assertMonotonicInvariants(set, spc)
}

func scaleDownStatefulSetController(set *apps.StatefulSet, ssc *StatefulSetController, spc *fakeStatefulPodControl) error {
	selector, err := metav1.LabelSelectorAsSelector(set.Spec.Selector)
	if err != nil {
		return err
	}
	pods, err := spc.podsLister.Pods(set.Namespace).List(selector)
	if err != nil {
		return err
	}
	ord := len(pods) - 1
	pod := getPodAtOrdinal(pods, ord)
	prev := *pod
	fakeResourceVersion(set)
	spc.setsIndexer.Add(set)
	ssc.enqueueStatefulSet(set)
	fakeWorker(ssc)
	pods, err = spc.addTerminatingPod(set, ord)
	if err != nil {
		return err
	}
	pod = getPodAtOrdinal(pods, ord)
	ssc.updatePod(&prev, pod)
	fakeWorker(ssc)
	spc.DeleteStatefulPod(set, pod)
	ssc.deletePod(pod)
	fakeWorker(ssc)
	for set.Status.Replicas > *set.Spec.Replicas {
		pods, err = spc.podsLister.Pods(set.Namespace).List(selector)
		if err != nil {
			return err
		}

		ord := len(pods)
		pods, err = spc.addTerminatingPod(set, ord)
		if err != nil {
			return err
		}
		pod = getPodAtOrdinal(pods, ord)
		ssc.updatePod(&prev, pod)
		fakeWorker(ssc)
		spc.DeleteStatefulPod(set, pod)
		ssc.deletePod(pod)
		fakeWorker(ssc)
		obj, _, err := spc.setsIndexer.Get(set)
		if err != nil {
			return err
		}
		set = obj.(*apps.StatefulSet)

	}
	return assertMonotonicInvariants(set, spc)
}

func rawTemplate(template *v1.PodTemplateSpec) runtime.RawExtension {
	buf := new(bytes.Buffer)
	enc := json.NewEncoder(buf)
	if err := enc.Encode(template); err != nil {
		panic(err)
	}
	return runtime.RawExtension{Raw: buf.Bytes()}
}
