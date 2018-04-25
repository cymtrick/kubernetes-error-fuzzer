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

package core

import (
	"reflect"
	"sync"
	"testing"
	"time"

	"k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/sets"
	"k8s.io/apimachinery/pkg/util/wait"
	"k8s.io/kubernetes/pkg/scheduler/algorithm"
	"k8s.io/kubernetes/pkg/scheduler/algorithm/predicates"
	algorithmpredicates "k8s.io/kubernetes/pkg/scheduler/algorithm/predicates"
	"k8s.io/kubernetes/pkg/scheduler/schedulercache"
	schedulertesting "k8s.io/kubernetes/pkg/scheduler/testing"
)

type predicateItemType struct {
	fit     bool
	reasons []algorithm.PredicateFailureReason
}

func TestUpdateCachedPredicateItem(t *testing.T) {
	tests := []struct {
		name               string
		pod                string
		predicateKey       string
		nodeName           string
		fit                bool
		reasons            []algorithm.PredicateFailureReason
		equivalenceHash    uint64
		expectPredicateMap bool
		expectCacheItem    HostPredicate
	}{
		{
			name:               "test 1",
			pod:                "testPod",
			predicateKey:       "GeneralPredicates",
			nodeName:           "node1",
			fit:                true,
			equivalenceHash:    123,
			expectPredicateMap: false,
			expectCacheItem: HostPredicate{
				Fit: true,
			},
		},
		{
			name:               "test 2",
			pod:                "testPod",
			predicateKey:       "GeneralPredicates",
			nodeName:           "node2",
			fit:                false,
			equivalenceHash:    123,
			expectPredicateMap: true,
			expectCacheItem: HostPredicate{
				Fit: false,
			},
		},
	}
	for _, test := range tests {
		// this case does not need to calculate equivalence hash, just pass an empty function
		fakeGetEquivalencePodFunc := func(pod *v1.Pod) interface{} { return nil }
		ecache := NewEquivalenceCache(fakeGetEquivalencePodFunc)
		if test.expectPredicateMap {
			ecache.algorithmCache[test.nodeName] = newAlgorithmCache()
			predicateItem := HostPredicate{
				Fit: true,
			}
			ecache.algorithmCache[test.nodeName].predicatesCache.Add(test.predicateKey,
				PredicateMap{
					test.equivalenceHash: predicateItem,
				})
		}
		ecache.UpdateCachedPredicateItem(
			test.pod,
			test.nodeName,
			test.predicateKey,
			test.fit,
			test.reasons,
			test.equivalenceHash,
			true,
		)

		value, ok := ecache.algorithmCache[test.nodeName].predicatesCache.Get(test.predicateKey)
		if !ok {
			t.Errorf("Failed: %s, can't find expected cache item: %v",
				test.name, test.expectCacheItem)
		} else {
			cachedMapItem := value.(PredicateMap)
			if !reflect.DeepEqual(cachedMapItem[test.equivalenceHash], test.expectCacheItem) {
				t.Errorf("Failed: %s, expected cached item: %v, but got: %v",
					test.name, test.expectCacheItem, cachedMapItem[test.equivalenceHash])
			}
		}
	}
}

func TestPredicateWithECache(t *testing.T) {
	tests := []struct {
		name                              string
		podName                           string
		nodeName                          string
		predicateKey                      string
		equivalenceHashForUpdatePredicate uint64
		equivalenceHashForCalPredicate    uint64
		cachedItem                        predicateItemType
		expectedInvalidPredicateKey       bool
		expectedInvalidEquivalenceHash    bool
		expectedPredicateItem             predicateItemType
	}{
		{
			name:     "test 1",
			podName:  "testPod",
			nodeName: "node1",
			equivalenceHashForUpdatePredicate: 123,
			equivalenceHashForCalPredicate:    123,
			predicateKey:                      "GeneralPredicates",
			cachedItem: predicateItemType{
				fit:     false,
				reasons: []algorithm.PredicateFailureReason{predicates.ErrPodNotFitsHostPorts},
			},
			expectedInvalidPredicateKey: true,
			expectedPredicateItem: predicateItemType{
				fit:     false,
				reasons: []algorithm.PredicateFailureReason{},
			},
		},
		{
			name:     "test 2",
			podName:  "testPod",
			nodeName: "node2",
			equivalenceHashForUpdatePredicate: 123,
			equivalenceHashForCalPredicate:    123,
			predicateKey:                      "GeneralPredicates",
			cachedItem: predicateItemType{
				fit: true,
			},
			expectedInvalidPredicateKey: false,
			expectedPredicateItem: predicateItemType{
				fit:     true,
				reasons: []algorithm.PredicateFailureReason{},
			},
		},
		{
			name:     "test 3",
			podName:  "testPod",
			nodeName: "node3",
			equivalenceHashForUpdatePredicate: 123,
			equivalenceHashForCalPredicate:    123,
			predicateKey:                      "GeneralPredicates",
			cachedItem: predicateItemType{
				fit:     false,
				reasons: []algorithm.PredicateFailureReason{predicates.ErrPodNotFitsHostPorts},
			},
			expectedInvalidPredicateKey: false,
			expectedPredicateItem: predicateItemType{
				fit:     false,
				reasons: []algorithm.PredicateFailureReason{predicates.ErrPodNotFitsHostPorts},
			},
		},
		{
			name:     "test 4",
			podName:  "testPod",
			nodeName: "node4",
			equivalenceHashForUpdatePredicate: 123,
			equivalenceHashForCalPredicate:    456,
			predicateKey:                      "GeneralPredicates",
			cachedItem: predicateItemType{
				fit:     false,
				reasons: []algorithm.PredicateFailureReason{predicates.ErrPodNotFitsHostPorts},
			},
			expectedInvalidPredicateKey:    false,
			expectedInvalidEquivalenceHash: true,
			expectedPredicateItem: predicateItemType{
				fit:     false,
				reasons: []algorithm.PredicateFailureReason{},
			},
		},
	}

	for _, test := range tests {
		// this case does not need to calculate equivalence hash, just pass an empty function
		fakeGetEquivalencePodFunc := func(pod *v1.Pod) interface{} { return nil }
		ecache := NewEquivalenceCache(fakeGetEquivalencePodFunc)
		// set cached item to equivalence cache
		ecache.UpdateCachedPredicateItem(
			test.podName,
			test.nodeName,
			test.predicateKey,
			test.cachedItem.fit,
			test.cachedItem.reasons,
			test.equivalenceHashForUpdatePredicate,
			true,
		)
		// if we want to do invalid, invalid the cached item
		if test.expectedInvalidPredicateKey {
			predicateKeys := sets.NewString()
			predicateKeys.Insert(test.predicateKey)
			ecache.InvalidateCachedPredicateItem(test.nodeName, predicateKeys)
		}
		// calculate predicate with equivalence cache
		fit, reasons, invalid := ecache.PredicateWithECache(test.podName,
			test.nodeName,
			test.predicateKey,
			test.equivalenceHashForCalPredicate,
			true,
		)
		// returned invalid should match expectedInvalidPredicateKey or expectedInvalidEquivalenceHash
		if test.equivalenceHashForUpdatePredicate != test.equivalenceHashForCalPredicate {
			if invalid != test.expectedInvalidEquivalenceHash {
				t.Errorf("Failed: %s, expected invalid: %v, but got: %v",
					test.name, test.expectedInvalidEquivalenceHash, invalid)
			}
		} else {
			if invalid != test.expectedInvalidPredicateKey {
				t.Errorf("Failed: %s, expected invalid: %v, but got: %v",
					test.name, test.expectedInvalidPredicateKey, invalid)
			}
		}
		// returned predicate result should match expected predicate item
		if fit != test.expectedPredicateItem.fit {
			t.Errorf("Failed: %s, expected fit: %v, but got: %v", test.name, test.cachedItem.fit, fit)
		}
		if !reflect.DeepEqual(reasons, test.expectedPredicateItem.reasons) {
			t.Errorf("Failed: %s, expected reasons: %v, but got: %v",
				test.name, test.cachedItem.reasons, reasons)
		}
	}
}

func TestGetHashEquivalencePod(t *testing.T) {

	testNamespace := "test"

	pvcInfo := predicates.FakePersistentVolumeClaimInfo{
		{
			ObjectMeta: metav1.ObjectMeta{UID: "someEBSVol1", Name: "someEBSVol1", Namespace: testNamespace},
			Spec:       v1.PersistentVolumeClaimSpec{VolumeName: "someEBSVol1"},
		},
		{
			ObjectMeta: metav1.ObjectMeta{UID: "someEBSVol2", Name: "someEBSVol2", Namespace: testNamespace},
			Spec:       v1.PersistentVolumeClaimSpec{VolumeName: "someNonEBSVol"},
		},
		{
			ObjectMeta: metav1.ObjectMeta{UID: "someEBSVol3-0", Name: "someEBSVol3-0", Namespace: testNamespace},
			Spec:       v1.PersistentVolumeClaimSpec{VolumeName: "pvcWithDeletedPV"},
		},
		{
			ObjectMeta: metav1.ObjectMeta{UID: "someEBSVol3-1", Name: "someEBSVol3-1", Namespace: testNamespace},
			Spec:       v1.PersistentVolumeClaimSpec{VolumeName: "anotherPVCWithDeletedPV"},
		},
	}

	// use default equivalence class generator
	ecache := NewEquivalenceCache(predicates.NewEquivalencePodGenerator(pvcInfo))

	isController := true

	pod1 := &v1.Pod{
		ObjectMeta: metav1.ObjectMeta{
			Name:      "pod1",
			Namespace: testNamespace,
			OwnerReferences: []metav1.OwnerReference{
				{
					APIVersion: "v1",
					Kind:       "ReplicationController",
					Name:       "rc",
					UID:        "123",
					Controller: &isController,
				},
			},
		},
		Spec: v1.PodSpec{
			Volumes: []v1.Volume{
				{
					VolumeSource: v1.VolumeSource{
						PersistentVolumeClaim: &v1.PersistentVolumeClaimVolumeSource{
							ClaimName: "someEBSVol1",
						},
					},
				},
				{
					VolumeSource: v1.VolumeSource{
						PersistentVolumeClaim: &v1.PersistentVolumeClaimVolumeSource{
							ClaimName: "someEBSVol2",
						},
					},
				},
			},
		},
	}

	pod2 := &v1.Pod{
		ObjectMeta: metav1.ObjectMeta{
			Name:      "pod2",
			Namespace: testNamespace,
			OwnerReferences: []metav1.OwnerReference{
				{
					APIVersion: "v1",
					Kind:       "ReplicationController",
					Name:       "rc",
					UID:        "123",
					Controller: &isController,
				},
			},
		},
		Spec: v1.PodSpec{
			Volumes: []v1.Volume{
				{
					VolumeSource: v1.VolumeSource{
						PersistentVolumeClaim: &v1.PersistentVolumeClaimVolumeSource{
							ClaimName: "someEBSVol2",
						},
					},
				},
				{
					VolumeSource: v1.VolumeSource{
						PersistentVolumeClaim: &v1.PersistentVolumeClaimVolumeSource{
							ClaimName: "someEBSVol1",
						},
					},
				},
			},
		},
	}

	pod3 := &v1.Pod{
		ObjectMeta: metav1.ObjectMeta{
			Name:      "pod3",
			Namespace: testNamespace,
			OwnerReferences: []metav1.OwnerReference{
				{
					APIVersion: "v1",
					Kind:       "ReplicationController",
					Name:       "rc",
					UID:        "567",
					Controller: &isController,
				},
			},
		},
		Spec: v1.PodSpec{
			Volumes: []v1.Volume{
				{
					VolumeSource: v1.VolumeSource{
						PersistentVolumeClaim: &v1.PersistentVolumeClaimVolumeSource{
							ClaimName: "someEBSVol3-1",
						},
					},
				},
			},
		},
	}

	pod4 := &v1.Pod{
		ObjectMeta: metav1.ObjectMeta{
			Name:      "pod4",
			Namespace: testNamespace,
			OwnerReferences: []metav1.OwnerReference{
				{
					APIVersion: "v1",
					Kind:       "ReplicationController",
					Name:       "rc",
					UID:        "567",
					Controller: &isController,
				},
			},
		},
		Spec: v1.PodSpec{
			Volumes: []v1.Volume{
				{
					VolumeSource: v1.VolumeSource{
						PersistentVolumeClaim: &v1.PersistentVolumeClaimVolumeSource{
							ClaimName: "someEBSVol3-0",
						},
					},
				},
			},
		},
	}

	pod5 := &v1.Pod{
		ObjectMeta: metav1.ObjectMeta{
			Name:      "pod5",
			Namespace: testNamespace,
		},
	}

	pod6 := &v1.Pod{
		ObjectMeta: metav1.ObjectMeta{
			Name:      "pod6",
			Namespace: testNamespace,
			OwnerReferences: []metav1.OwnerReference{
				{
					APIVersion: "v1",
					Kind:       "ReplicationController",
					Name:       "rc",
					UID:        "567",
					Controller: &isController,
				},
			},
		},
		Spec: v1.PodSpec{
			Volumes: []v1.Volume{
				{
					VolumeSource: v1.VolumeSource{
						PersistentVolumeClaim: &v1.PersistentVolumeClaimVolumeSource{
							ClaimName: "no-exists-pvc",
						},
					},
				},
			},
		},
	}

	pod7 := &v1.Pod{
		ObjectMeta: metav1.ObjectMeta{
			Name:      "pod7",
			Namespace: testNamespace,
			OwnerReferences: []metav1.OwnerReference{
				{
					APIVersion: "v1",
					Kind:       "ReplicationController",
					Name:       "rc",
					UID:        "567",
					Controller: &isController,
				},
			},
		},
	}

	type podInfo struct {
		pod         *v1.Pod
		hashIsValid bool
	}

	tests := []struct {
		podInfoList  []podInfo
		isEquivalent bool
	}{
		// pods with same controllerRef and same pvc claim
		{
			podInfoList: []podInfo{
				{pod: pod1, hashIsValid: true},
				{pod: pod2, hashIsValid: true},
			},
			isEquivalent: true,
		},
		// pods with same controllerRef but different pvc claim
		{
			podInfoList: []podInfo{
				{pod: pod3, hashIsValid: true},
				{pod: pod4, hashIsValid: true},
			},
			isEquivalent: false,
		},
		// pod without controllerRef
		{
			podInfoList: []podInfo{
				{pod: pod5, hashIsValid: false},
			},
			isEquivalent: false,
		},
		// pods with same controllerRef but one has non-exists pvc claim
		{
			podInfoList: []podInfo{
				{pod: pod6, hashIsValid: false},
				{pod: pod7, hashIsValid: true},
			},
			isEquivalent: false,
		},
	}

	var (
		targetPodInfo podInfo
		targetHash    uint64
	)

	for _, test := range tests {
		for i, podInfo := range test.podInfoList {
			testPod := podInfo.pod
			eclassInfo := ecache.getEquivalenceClassInfo(testPod)
			if eclassInfo == nil && podInfo.hashIsValid {
				t.Errorf("Failed: pod %v is expected to have valid hash", testPod)
			}

			if eclassInfo != nil {
				// NOTE(harry): the first element will be used as target so
				// this logic can't verify more than two inequivalent pods
				if i == 0 {
					targetHash = eclassInfo.hash
					targetPodInfo = podInfo
				} else {
					if targetHash != eclassInfo.hash {
						if test.isEquivalent {
							t.Errorf("Failed: pod: %v is expected to be equivalent to: %v", testPod, targetPodInfo.pod)
						}
					}
				}
			}
		}
	}
}

func TestInvalidateCachedPredicateItemOfAllNodes(t *testing.T) {
	testPredicate := "GeneralPredicates"
	// tests is used to initialize all nodes
	tests := []struct {
		podName                           string
		nodeName                          string
		predicateKey                      string
		equivalenceHashForUpdatePredicate uint64
		cachedItem                        predicateItemType
	}{
		{
			podName:  "testPod",
			nodeName: "node1",
			equivalenceHashForUpdatePredicate: 123,
			cachedItem: predicateItemType{
				fit: false,
				reasons: []algorithm.PredicateFailureReason{
					predicates.ErrPodNotFitsHostPorts,
				},
			},
		},
		{
			podName:  "testPod",
			nodeName: "node2",
			equivalenceHashForUpdatePredicate: 456,
			cachedItem: predicateItemType{
				fit: false,
				reasons: []algorithm.PredicateFailureReason{
					predicates.ErrPodNotFitsHostPorts,
				},
			},
		},
		{
			podName:  "testPod",
			nodeName: "node3",
			equivalenceHashForUpdatePredicate: 123,
			cachedItem: predicateItemType{
				fit: true,
			},
		},
	}
	// this case does not need to calculate equivalence hash, just pass an empty function
	fakeGetEquivalencePodFunc := func(pod *v1.Pod) interface{} { return nil }
	ecache := NewEquivalenceCache(fakeGetEquivalencePodFunc)

	for _, test := range tests {
		// set cached item to equivalence cache
		ecache.UpdateCachedPredicateItem(
			test.podName,
			test.nodeName,
			testPredicate,
			test.cachedItem.fit,
			test.cachedItem.reasons,
			test.equivalenceHashForUpdatePredicate,
			true,
		)
	}

	// invalidate cached predicate for all nodes
	ecache.InvalidateCachedPredicateItemOfAllNodes(sets.NewString(testPredicate))

	// there should be no cached predicate any more
	for _, test := range tests {
		if algorithmCache, exist := ecache.algorithmCache[test.nodeName]; exist {
			if _, exist := algorithmCache.predicatesCache.Get(testPredicate); exist {
				t.Errorf("Failed: cached item for predicate key: %v on node: %v should be invalidated",
					testPredicate, test.nodeName)
				break
			}
		}
	}
}

func TestInvalidateAllCachedPredicateItemOfNode(t *testing.T) {
	testPredicate := "GeneralPredicates"
	// tests is used to initialize all nodes
	tests := []struct {
		podName                           string
		nodeName                          string
		predicateKey                      string
		equivalenceHashForUpdatePredicate uint64
		cachedItem                        predicateItemType
	}{
		{
			podName:  "testPod",
			nodeName: "node1",
			equivalenceHashForUpdatePredicate: 123,
			cachedItem: predicateItemType{
				fit:     false,
				reasons: []algorithm.PredicateFailureReason{predicates.ErrPodNotFitsHostPorts},
			},
		},
		{
			podName:  "testPod",
			nodeName: "node2",
			equivalenceHashForUpdatePredicate: 456,
			cachedItem: predicateItemType{
				fit:     false,
				reasons: []algorithm.PredicateFailureReason{predicates.ErrPodNotFitsHostPorts},
			},
		},
		{
			podName:  "testPod",
			nodeName: "node3",
			equivalenceHashForUpdatePredicate: 123,
			cachedItem: predicateItemType{
				fit: true,
			},
		},
	}
	// this case does not need to calculate equivalence hash, just pass an empty function
	fakeGetEquivalencePodFunc := func(pod *v1.Pod) interface{} { return nil }
	ecache := NewEquivalenceCache(fakeGetEquivalencePodFunc)

	for _, test := range tests {
		// set cached item to equivalence cache
		ecache.UpdateCachedPredicateItem(
			test.podName,
			test.nodeName,
			testPredicate,
			test.cachedItem.fit,
			test.cachedItem.reasons,
			test.equivalenceHashForUpdatePredicate,
			true,
		)
	}

	for _, test := range tests {
		// invalidate cached predicate for all nodes
		ecache.InvalidateAllCachedPredicateItemOfNode(test.nodeName)
		if _, exist := ecache.algorithmCache[test.nodeName]; exist {
			t.Errorf("Failed: cached item for node: %v should be invalidated", test.nodeName)
			break
		}
	}
}

// syncingMockCache delegates method calls to an actual Cache,
// but calls to UpdateNodeNameToInfoMap synchronize with the test.
type syncingMockCache struct {
	schedulercache.Cache
	cycleStart, cacheInvalidated chan struct{}
	once                         sync.Once
}

// UpdateNodeNameToInfoMap delegates to the real implementation, but on the first call, it
// synchronizes with the test.
//
// Since UpdateNodeNameToInfoMap is one of the first steps of (*genericScheduler).Schedule, we use
// this point to signal to the test that a scheduling cycle has started.
func (c *syncingMockCache) UpdateNodeNameToInfoMap(infoMap map[string]*schedulercache.NodeInfo) error {
	err := c.Cache.UpdateNodeNameToInfoMap(infoMap)
	c.once.Do(func() {
		c.cycleStart <- struct{}{}
		<-c.cacheInvalidated
	})
	return err
}

// TestEquivalenceCacheInvalidationRace tests that equivalence cache invalidation is correctly
// handled when an invalidation event happens early in a scheduling cycle. Specifically, the event
// occurs after schedulercache is snapshotted and before equivalence cache lock is acquired.
func TestEquivalenceCacheInvalidationRace(t *testing.T) {
	// Create a predicate that returns false the first time and true on subsequent calls.
	podWillFit := false
	var callCount int
	testPredicate := func(pod *v1.Pod,
		meta algorithm.PredicateMetadata,
		nodeInfo *schedulercache.NodeInfo) (bool, []algorithm.PredicateFailureReason, error) {
		callCount++
		if !podWillFit {
			podWillFit = true
			return false, []algorithm.PredicateFailureReason{algorithmpredicates.ErrFakePredicate}, nil
		}
		return true, nil, nil
	}

	// Set up the mock cache.
	cache := schedulercache.New(time.Duration(0), wait.NeverStop)
	cache.AddNode(&v1.Node{ObjectMeta: metav1.ObjectMeta{Name: "machine1"}})
	mockCache := &syncingMockCache{
		Cache:            cache,
		cycleStart:       make(chan struct{}),
		cacheInvalidated: make(chan struct{}),
	}

	fakeGetEquivalencePod := func(pod *v1.Pod) interface{} { return pod }
	eCache := NewEquivalenceCache(fakeGetEquivalencePod)
	// Ensure that equivalence cache invalidation happens after the scheduling cycle starts, but before
	// the equivalence cache would be updated.
	go func() {
		<-mockCache.cycleStart
		pod := &v1.Pod{
			ObjectMeta: metav1.ObjectMeta{Name: "new-pod", UID: "new-pod"},
			Spec:       v1.PodSpec{NodeName: "machine1"}}
		if err := cache.AddPod(pod); err != nil {
			t.Errorf("Could not add pod to cache: %v", err)
		}
		eCache.InvalidateAllCachedPredicateItemOfNode("machine1")
		mockCache.cacheInvalidated <- struct{}{}
	}()

	// Set up the scheduler.
	predicates := map[string]algorithm.FitPredicate{"testPredicate": testPredicate}
	algorithmpredicates.SetPredicatesOrdering([]string{"testPredicate"})
	prioritizers := []algorithm.PriorityConfig{{Map: EqualPriorityMap, Weight: 1}}
	pvcLister := schedulertesting.FakePersistentVolumeClaimLister([]*v1.PersistentVolumeClaim{})
	scheduler := NewGenericScheduler(
		mockCache,
		eCache,
		NewSchedulingQueue(),
		predicates,
		algorithm.EmptyPredicateMetadataProducer,
		prioritizers,
		algorithm.EmptyPriorityMetadataProducer,
		nil, nil, pvcLister, true, false)

	// First scheduling attempt should fail.
	nodeLister := schedulertesting.FakeNodeLister(makeNodeList([]string{"machine1"}))
	pod := &v1.Pod{ObjectMeta: metav1.ObjectMeta{Name: "test-pod"}}
	machine, err := scheduler.Schedule(pod, nodeLister)
	if machine != "" || err == nil {
		t.Error("First scheduling attempt did not fail")
	}

	// Second scheduling attempt should succeed because cache was invalidated.
	_, err = scheduler.Schedule(pod, nodeLister)
	if err != nil {
		t.Errorf("Second scheduling attempt failed: %v", err)
	}
	if callCount != 2 {
		t.Errorf("Predicate should have been called twice. Was called %d times.", callCount)
	}
}
