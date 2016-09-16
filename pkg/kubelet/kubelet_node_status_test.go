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

package kubelet

import (
	"fmt"
	"reflect"
	goruntime "runtime"
	"sort"
	"strconv"
	"testing"
	"time"

	cadvisorapi "github.com/google/cadvisor/info/v1"
	cadvisorapiv2 "github.com/google/cadvisor/info/v2"
	"k8s.io/kubernetes/pkg/api"
	apierrors "k8s.io/kubernetes/pkg/api/errors"
	"k8s.io/kubernetes/pkg/api/resource"
	"k8s.io/kubernetes/pkg/api/unversioned"
	"k8s.io/kubernetes/pkg/client/clientset_generated/internalclientset/fake"
	"k8s.io/kubernetes/pkg/client/testing/core"
	kubecontainer "k8s.io/kubernetes/pkg/kubelet/container"
	"k8s.io/kubernetes/pkg/kubelet/util/sliceutils"
	"k8s.io/kubernetes/pkg/runtime"
	"k8s.io/kubernetes/pkg/util/diff"
	"k8s.io/kubernetes/pkg/util/rand"
	"k8s.io/kubernetes/pkg/util/uuid"
	"k8s.io/kubernetes/pkg/util/wait"
	"k8s.io/kubernetes/pkg/version"
	"k8s.io/kubernetes/pkg/volume/util/volumehelper"
)

const (
	maxImageTagsForTest = 20
)

// generateTestingImageList generate randomly generated image list and corresponding expectedImageList.
func generateTestingImageList(count int) ([]kubecontainer.Image, []api.ContainerImage) {
	// imageList is randomly generated image list
	var imageList []kubecontainer.Image
	for ; count > 0; count-- {
		imageItem := kubecontainer.Image{
			ID:       string(uuid.NewUUID()),
			RepoTags: generateImageTags(),
			Size:     rand.Int63nRange(minImgSize, maxImgSize+1),
		}
		imageList = append(imageList, imageItem)
	}

	// expectedImageList is generated by imageList according to size and maxImagesInNodeStatus
	// 1. sort the imageList by size
	sort.Sort(sliceutils.ByImageSize(imageList))
	// 2. convert sorted imageList to api.ContainerImage list
	var expectedImageList []api.ContainerImage
	for _, kubeImage := range imageList {
		apiImage := api.ContainerImage{
			Names:     kubeImage.RepoTags[0:maxNamesPerImageInNodeStatus],
			SizeBytes: kubeImage.Size,
		}

		expectedImageList = append(expectedImageList, apiImage)
	}
	// 3. only returns the top maxImagesInNodeStatus images in expectedImageList
	return imageList, expectedImageList[0:maxImagesInNodeStatus]
}

func generateImageTags() []string {
	var tagList []string
	// Generate > maxNamesPerImageInNodeStatus tags so that the test can verify
	// that kubelet report up to maxNamesPerImageInNodeStatus tags.
	count := rand.IntnRange(maxNamesPerImageInNodeStatus+1, maxImageTagsForTest+1)
	for ; count > 0; count-- {
		tagList = append(tagList, "gcr.io/google_containers:v"+strconv.Itoa(count))
	}
	return tagList
}

func TestUpdateNewNodeStatus(t *testing.T) {
	// generate one more than maxImagesInNodeStatus in inputImageList
	inputImageList, expectedImageList := generateTestingImageList(maxImagesInNodeStatus + 1)
	testKubelet := newTestKubeletWithImageList(
		t, inputImageList, false /* controllerAttachDetachEnabled */)
	kubelet := testKubelet.kubelet
	kubeClient := testKubelet.fakeKubeClient
	kubeClient.ReactionChain = fake.NewSimpleClientset(&api.NodeList{Items: []api.Node{
		{ObjectMeta: api.ObjectMeta{Name: testKubeletHostname}},
	}}).ReactionChain
	machineInfo := &cadvisorapi.MachineInfo{
		MachineID:      "123",
		SystemUUID:     "abc",
		BootID:         "1b3",
		NumCores:       2,
		MemoryCapacity: 10E9, // 10G
	}
	mockCadvisor := testKubelet.fakeCadvisor
	mockCadvisor.On("Start").Return(nil)
	mockCadvisor.On("MachineInfo").Return(machineInfo, nil)
	versionInfo := &cadvisorapi.VersionInfo{
		KernelVersion:      "3.16.0-0.bpo.4-amd64",
		ContainerOsVersion: "Debian GNU/Linux 7 (wheezy)",
	}
	mockCadvisor.On("VersionInfo").Return(versionInfo, nil)

	// Make kubelet report that it has sufficient disk space.
	if err := updateDiskSpacePolicy(kubelet, mockCadvisor, 500, 500, 200, 200, 100, 100); err != nil {
		t.Fatalf("can't update disk space manager: %v", err)
	}

	expectedNode := &api.Node{
		ObjectMeta: api.ObjectMeta{Name: testKubeletHostname},
		Spec:       api.NodeSpec{},
		Status: api.NodeStatus{
			Conditions: []api.NodeCondition{
				{
					Type:               api.NodeOutOfDisk,
					Status:             api.ConditionFalse,
					Reason:             "KubeletHasSufficientDisk",
					Message:            fmt.Sprintf("kubelet has sufficient disk space available"),
					LastHeartbeatTime:  unversioned.Time{},
					LastTransitionTime: unversioned.Time{},
				},
				{
					Type:               api.NodeMemoryPressure,
					Status:             api.ConditionFalse,
					Reason:             "KubeletHasSufficientMemory",
					Message:            fmt.Sprintf("kubelet has sufficient memory available"),
					LastHeartbeatTime:  unversioned.Time{},
					LastTransitionTime: unversioned.Time{},
				},
				{
					Type:               api.NodeDiskPressure,
					Status:             api.ConditionFalse,
					Reason:             "KubeletHasNoDiskPressure",
					Message:            fmt.Sprintf("kubelet has no disk pressure"),
					LastHeartbeatTime:  unversioned.Time{},
					LastTransitionTime: unversioned.Time{},
				},
				{
					Type:               api.NodeReady,
					Status:             api.ConditionTrue,
					Reason:             "KubeletReady",
					Message:            fmt.Sprintf("kubelet is posting ready status"),
					LastHeartbeatTime:  unversioned.Time{},
					LastTransitionTime: unversioned.Time{},
				},
			},
			NodeInfo: api.NodeSystemInfo{
				MachineID:               "123",
				SystemUUID:              "abc",
				BootID:                  "1b3",
				KernelVersion:           "3.16.0-0.bpo.4-amd64",
				OSImage:                 "Debian GNU/Linux 7 (wheezy)",
				OperatingSystem:         goruntime.GOOS,
				Architecture:            goruntime.GOARCH,
				ContainerRuntimeVersion: "test://1.5.0",
				KubeletVersion:          version.Get().String(),
				KubeProxyVersion:        version.Get().String(),
			},
			Capacity: api.ResourceList{
				api.ResourceCPU:       *resource.NewMilliQuantity(2000, resource.DecimalSI),
				api.ResourceMemory:    *resource.NewQuantity(10E9, resource.BinarySI),
				api.ResourcePods:      *resource.NewQuantity(0, resource.DecimalSI),
				api.ResourceNvidiaGPU: *resource.NewQuantity(0, resource.DecimalSI),
			},
			Allocatable: api.ResourceList{
				api.ResourceCPU:       *resource.NewMilliQuantity(1800, resource.DecimalSI),
				api.ResourceMemory:    *resource.NewQuantity(9900E6, resource.BinarySI),
				api.ResourcePods:      *resource.NewQuantity(0, resource.DecimalSI),
				api.ResourceNvidiaGPU: *resource.NewQuantity(0, resource.DecimalSI),
			},
			Addresses: []api.NodeAddress{
				{Type: api.NodeLegacyHostIP, Address: "127.0.0.1"},
				{Type: api.NodeInternalIP, Address: "127.0.0.1"},
			},
			Images: expectedImageList,
		},
	}

	kubelet.updateRuntimeUp()
	if err := kubelet.updateNodeStatus(); err != nil {
		t.Errorf("unexpected error: %v", err)
	}
	actions := kubeClient.Actions()
	if len(actions) != 2 {
		t.Fatalf("unexpected actions: %v", actions)
	}
	if !actions[1].Matches("update", "nodes") || actions[1].GetSubresource() != "status" {
		t.Fatalf("unexpected actions: %v", actions)
	}
	updatedNode, ok := actions[1].(core.UpdateAction).GetObject().(*api.Node)
	if !ok {
		t.Errorf("unexpected object type")
	}
	for i, cond := range updatedNode.Status.Conditions {
		if cond.LastHeartbeatTime.IsZero() {
			t.Errorf("unexpected zero last probe timestamp for %v condition", cond.Type)
		}
		if cond.LastTransitionTime.IsZero() {
			t.Errorf("unexpected zero last transition timestamp for %v condition", cond.Type)
		}
		updatedNode.Status.Conditions[i].LastHeartbeatTime = unversioned.Time{}
		updatedNode.Status.Conditions[i].LastTransitionTime = unversioned.Time{}
	}

	// Version skew workaround. See: https://github.com/kubernetes/kubernetes/issues/16961
	if updatedNode.Status.Conditions[len(updatedNode.Status.Conditions)-1].Type != api.NodeReady {
		t.Errorf("unexpected node condition order. NodeReady should be last.")
	}

	if maxImagesInNodeStatus != len(updatedNode.Status.Images) {
		t.Errorf("unexpected image list length in node status, expected: %v, got: %v", maxImagesInNodeStatus, len(updatedNode.Status.Images))
	} else {
		if !api.Semantic.DeepEqual(expectedNode, updatedNode) {
			t.Errorf("unexpected objects: %s", diff.ObjectDiff(expectedNode, updatedNode))
		}
	}

}

func TestUpdateNewNodeOutOfDiskStatusWithTransitionFrequency(t *testing.T) {
	testKubelet := newTestKubelet(t, false /* controllerAttachDetachEnabled */)
	kubelet := testKubelet.kubelet
	kubeClient := testKubelet.fakeKubeClient
	kubeClient.ReactionChain = fake.NewSimpleClientset(&api.NodeList{Items: []api.Node{
		{ObjectMeta: api.ObjectMeta{Name: testKubeletHostname}},
	}}).ReactionChain
	machineInfo := &cadvisorapi.MachineInfo{
		MachineID:      "123",
		SystemUUID:     "abc",
		BootID:         "1b3",
		NumCores:       2,
		MemoryCapacity: 1024,
	}
	mockCadvisor := testKubelet.fakeCadvisor
	mockCadvisor.On("Start").Return(nil)
	mockCadvisor.On("MachineInfo").Return(machineInfo, nil)
	versionInfo := &cadvisorapi.VersionInfo{
		KernelVersion:      "3.16.0-0.bpo.4-amd64",
		ContainerOsVersion: "Debian GNU/Linux 7 (wheezy)",
	}
	mockCadvisor.On("VersionInfo").Return(versionInfo, nil)

	// Make Kubelet report that it has sufficient disk space.
	if err := updateDiskSpacePolicy(kubelet, mockCadvisor, 500, 500, 200, 200, 100, 100); err != nil {
		t.Fatalf("can't update disk space manager: %v", err)
	}

	kubelet.outOfDiskTransitionFrequency = 10 * time.Second

	expectedNodeOutOfDiskCondition := api.NodeCondition{
		Type:               api.NodeOutOfDisk,
		Status:             api.ConditionFalse,
		Reason:             "KubeletHasSufficientDisk",
		Message:            fmt.Sprintf("kubelet has sufficient disk space available"),
		LastHeartbeatTime:  unversioned.Time{},
		LastTransitionTime: unversioned.Time{},
	}

	kubelet.updateRuntimeUp()
	if err := kubelet.updateNodeStatus(); err != nil {
		t.Errorf("unexpected error: %v", err)
	}
	actions := kubeClient.Actions()
	if len(actions) != 2 {
		t.Fatalf("unexpected actions: %v", actions)
	}
	if !actions[1].Matches("update", "nodes") || actions[1].GetSubresource() != "status" {
		t.Fatalf("unexpected actions: %v", actions)
	}
	updatedNode, ok := actions[1].(core.UpdateAction).GetObject().(*api.Node)
	if !ok {
		t.Errorf("unexpected object type")
	}

	var oodCondition api.NodeCondition
	for i, cond := range updatedNode.Status.Conditions {
		if cond.LastHeartbeatTime.IsZero() {
			t.Errorf("unexpected zero last probe timestamp for %v condition", cond.Type)
		}
		if cond.LastTransitionTime.IsZero() {
			t.Errorf("unexpected zero last transition timestamp for %v condition", cond.Type)
		}
		updatedNode.Status.Conditions[i].LastHeartbeatTime = unversioned.Time{}
		updatedNode.Status.Conditions[i].LastTransitionTime = unversioned.Time{}
		if cond.Type == api.NodeOutOfDisk {
			oodCondition = updatedNode.Status.Conditions[i]
		}
	}

	if !reflect.DeepEqual(expectedNodeOutOfDiskCondition, oodCondition) {
		t.Errorf("unexpected objects: %s", diff.ObjectDiff(expectedNodeOutOfDiskCondition, oodCondition))
	}
}

func TestUpdateExistingNodeStatus(t *testing.T) {
	testKubelet := newTestKubelet(t, false /* controllerAttachDetachEnabled */)
	kubelet := testKubelet.kubelet
	kubeClient := testKubelet.fakeKubeClient
	kubeClient.ReactionChain = fake.NewSimpleClientset(&api.NodeList{Items: []api.Node{
		{
			ObjectMeta: api.ObjectMeta{Name: testKubeletHostname},
			Spec:       api.NodeSpec{},
			Status: api.NodeStatus{
				Conditions: []api.NodeCondition{
					{
						Type:               api.NodeOutOfDisk,
						Status:             api.ConditionTrue,
						Reason:             "KubeletOutOfDisk",
						Message:            "out of disk space",
						LastHeartbeatTime:  unversioned.Date(2012, 1, 1, 0, 0, 0, 0, time.UTC),
						LastTransitionTime: unversioned.Date(2012, 1, 1, 0, 0, 0, 0, time.UTC),
					},
					{
						Type:               api.NodeMemoryPressure,
						Status:             api.ConditionFalse,
						Reason:             "KubeletHasSufficientMemory",
						Message:            fmt.Sprintf("kubelet has sufficient memory available"),
						LastHeartbeatTime:  unversioned.Date(2012, 1, 1, 0, 0, 0, 0, time.UTC),
						LastTransitionTime: unversioned.Date(2012, 1, 1, 0, 0, 0, 0, time.UTC),
					},
					{
						Type:               api.NodeDiskPressure,
						Status:             api.ConditionFalse,
						Reason:             "KubeletHasSufficientDisk",
						Message:            fmt.Sprintf("kubelet has sufficient disk space available"),
						LastHeartbeatTime:  unversioned.Date(2012, 1, 1, 0, 0, 0, 0, time.UTC),
						LastTransitionTime: unversioned.Date(2012, 1, 1, 0, 0, 0, 0, time.UTC),
					},
					{
						Type:               api.NodeReady,
						Status:             api.ConditionTrue,
						Reason:             "KubeletReady",
						Message:            fmt.Sprintf("kubelet is posting ready status"),
						LastHeartbeatTime:  unversioned.Date(2012, 1, 1, 0, 0, 0, 0, time.UTC),
						LastTransitionTime: unversioned.Date(2012, 1, 1, 0, 0, 0, 0, time.UTC),
					},
				},
				Capacity: api.ResourceList{
					api.ResourceCPU:    *resource.NewMilliQuantity(3000, resource.DecimalSI),
					api.ResourceMemory: *resource.NewQuantity(20E9, resource.BinarySI),
					api.ResourcePods:   *resource.NewQuantity(0, resource.DecimalSI),
				},
				Allocatable: api.ResourceList{
					api.ResourceCPU:    *resource.NewMilliQuantity(2800, resource.DecimalSI),
					api.ResourceMemory: *resource.NewQuantity(19900E6, resource.BinarySI),
					api.ResourcePods:   *resource.NewQuantity(0, resource.DecimalSI),
				},
			},
		},
	}}).ReactionChain
	mockCadvisor := testKubelet.fakeCadvisor
	mockCadvisor.On("Start").Return(nil)
	machineInfo := &cadvisorapi.MachineInfo{
		MachineID:      "123",
		SystemUUID:     "abc",
		BootID:         "1b3",
		NumCores:       2,
		MemoryCapacity: 20E9,
	}
	mockCadvisor.On("MachineInfo").Return(machineInfo, nil)
	versionInfo := &cadvisorapi.VersionInfo{
		KernelVersion:      "3.16.0-0.bpo.4-amd64",
		ContainerOsVersion: "Debian GNU/Linux 7 (wheezy)",
	}
	mockCadvisor.On("VersionInfo").Return(versionInfo, nil)

	// Make kubelet report that it is out of disk space.
	if err := updateDiskSpacePolicy(kubelet, mockCadvisor, 500, 500, 50, 50, 100, 100); err != nil {
		t.Fatalf("can't update disk space manager: %v", err)
	}

	expectedNode := &api.Node{
		ObjectMeta: api.ObjectMeta{Name: testKubeletHostname},
		Spec:       api.NodeSpec{},
		Status: api.NodeStatus{
			Conditions: []api.NodeCondition{
				{
					Type:               api.NodeOutOfDisk,
					Status:             api.ConditionTrue,
					Reason:             "KubeletOutOfDisk",
					Message:            "out of disk space",
					LastHeartbeatTime:  unversioned.Time{}, // placeholder
					LastTransitionTime: unversioned.Time{}, // placeholder
				},
				{
					Type:               api.NodeMemoryPressure,
					Status:             api.ConditionFalse,
					Reason:             "KubeletHasSufficientMemory",
					Message:            fmt.Sprintf("kubelet has sufficient memory available"),
					LastHeartbeatTime:  unversioned.Time{},
					LastTransitionTime: unversioned.Time{},
				},
				{
					Type:               api.NodeDiskPressure,
					Status:             api.ConditionFalse,
					Reason:             "KubeletHasSufficientDisk",
					Message:            fmt.Sprintf("kubelet has sufficient disk space available"),
					LastHeartbeatTime:  unversioned.Time{},
					LastTransitionTime: unversioned.Time{},
				},
				{
					Type:               api.NodeReady,
					Status:             api.ConditionTrue,
					Reason:             "KubeletReady",
					Message:            fmt.Sprintf("kubelet is posting ready status"),
					LastHeartbeatTime:  unversioned.Time{}, // placeholder
					LastTransitionTime: unversioned.Time{}, // placeholder
				},
			},
			NodeInfo: api.NodeSystemInfo{
				MachineID:               "123",
				SystemUUID:              "abc",
				BootID:                  "1b3",
				KernelVersion:           "3.16.0-0.bpo.4-amd64",
				OSImage:                 "Debian GNU/Linux 7 (wheezy)",
				OperatingSystem:         goruntime.GOOS,
				Architecture:            goruntime.GOARCH,
				ContainerRuntimeVersion: "test://1.5.0",
				KubeletVersion:          version.Get().String(),
				KubeProxyVersion:        version.Get().String(),
			},
			Capacity: api.ResourceList{
				api.ResourceCPU:       *resource.NewMilliQuantity(2000, resource.DecimalSI),
				api.ResourceMemory:    *resource.NewQuantity(20E9, resource.BinarySI),
				api.ResourcePods:      *resource.NewQuantity(0, resource.DecimalSI),
				api.ResourceNvidiaGPU: *resource.NewQuantity(0, resource.DecimalSI),
			},
			Allocatable: api.ResourceList{
				api.ResourceCPU:       *resource.NewMilliQuantity(1800, resource.DecimalSI),
				api.ResourceMemory:    *resource.NewQuantity(19900E6, resource.BinarySI),
				api.ResourcePods:      *resource.NewQuantity(0, resource.DecimalSI),
				api.ResourceNvidiaGPU: *resource.NewQuantity(0, resource.DecimalSI),
			},
			Addresses: []api.NodeAddress{
				{Type: api.NodeLegacyHostIP, Address: "127.0.0.1"},
				{Type: api.NodeInternalIP, Address: "127.0.0.1"},
			},
			// images will be sorted from max to min in node status.
			Images: []api.ContainerImage{
				{
					Names:     []string{"gcr.io/google_containers:v3", "gcr.io/google_containers:v4"},
					SizeBytes: 456,
				},
				{
					Names:     []string{"gcr.io/google_containers:v1", "gcr.io/google_containers:v2"},
					SizeBytes: 123,
				},
			},
		},
	}

	kubelet.updateRuntimeUp()
	if err := kubelet.updateNodeStatus(); err != nil {
		t.Errorf("unexpected error: %v", err)
	}
	actions := kubeClient.Actions()
	if len(actions) != 2 {
		t.Errorf("unexpected actions: %v", actions)
	}
	updateAction, ok := actions[1].(core.UpdateAction)
	if !ok {
		t.Errorf("unexpected action type.  expected UpdateAction, got %#v", actions[1])
	}
	updatedNode, ok := updateAction.GetObject().(*api.Node)
	if !ok {
		t.Errorf("unexpected object type")
	}
	for i, cond := range updatedNode.Status.Conditions {
		// Expect LastProbeTime to be updated to Now, while LastTransitionTime to be the same.
		if old := unversioned.Date(2012, 1, 1, 0, 0, 0, 0, time.UTC).Time; reflect.DeepEqual(cond.LastHeartbeatTime.Rfc3339Copy().UTC(), old) {
			t.Errorf("Condition %v LastProbeTime: expected \n%v\n, got \n%v", cond.Type, unversioned.Now(), old)
		}
		if got, want := cond.LastTransitionTime.Rfc3339Copy().UTC(), unversioned.Date(2012, 1, 1, 0, 0, 0, 0, time.UTC).Time; !reflect.DeepEqual(got, want) {
			t.Errorf("Condition %v LastTransitionTime: expected \n%#v\n, got \n%#v", cond.Type, want, got)
		}
		updatedNode.Status.Conditions[i].LastHeartbeatTime = unversioned.Time{}
		updatedNode.Status.Conditions[i].LastTransitionTime = unversioned.Time{}
	}

	// Version skew workaround. See: https://github.com/kubernetes/kubernetes/issues/16961
	if updatedNode.Status.Conditions[len(updatedNode.Status.Conditions)-1].Type != api.NodeReady {
		t.Errorf("unexpected node condition order. NodeReady should be last.")
	}

	if !api.Semantic.DeepEqual(expectedNode, updatedNode) {
		t.Errorf("unexpected objects: %s", diff.ObjectDiff(expectedNode, updatedNode))
	}
}

func TestUpdateExistingNodeOutOfDiskStatusWithTransitionFrequency(t *testing.T) {
	testKubelet := newTestKubelet(t, false /* controllerAttachDetachEnabled */)
	kubelet := testKubelet.kubelet
	clock := testKubelet.fakeClock
	kubeClient := testKubelet.fakeKubeClient
	kubeClient.ReactionChain = fake.NewSimpleClientset(&api.NodeList{Items: []api.Node{
		{
			ObjectMeta: api.ObjectMeta{Name: testKubeletHostname},
			Spec:       api.NodeSpec{},
			Status: api.NodeStatus{
				Conditions: []api.NodeCondition{
					{
						Type:               api.NodeReady,
						Status:             api.ConditionTrue,
						Reason:             "KubeletReady",
						Message:            fmt.Sprintf("kubelet is posting ready status"),
						LastHeartbeatTime:  unversioned.NewTime(clock.Now()),
						LastTransitionTime: unversioned.NewTime(clock.Now()),
					},
					{
						Type:               api.NodeOutOfDisk,
						Status:             api.ConditionTrue,
						Reason:             "KubeletOutOfDisk",
						Message:            "out of disk space",
						LastHeartbeatTime:  unversioned.NewTime(clock.Now()),
						LastTransitionTime: unversioned.NewTime(clock.Now()),
					},
				},
			},
		},
	}}).ReactionChain
	mockCadvisor := testKubelet.fakeCadvisor
	machineInfo := &cadvisorapi.MachineInfo{
		MachineID:      "123",
		SystemUUID:     "abc",
		BootID:         "1b3",
		NumCores:       2,
		MemoryCapacity: 1024,
	}
	fsInfo := cadvisorapiv2.FsInfo{
		Device: "123",
	}
	mockCadvisor.On("Start").Return(nil)
	mockCadvisor.On("MachineInfo").Return(machineInfo, nil)
	mockCadvisor.On("ImagesFsInfo").Return(fsInfo, nil)
	mockCadvisor.On("RootFsInfo").Return(fsInfo, nil)
	versionInfo := &cadvisorapi.VersionInfo{
		KernelVersion:      "3.16.0-0.bpo.4-amd64",
		ContainerOsVersion: "Debian GNU/Linux 7 (wheezy)",
		DockerVersion:      "1.5.0",
	}
	mockCadvisor.On("VersionInfo").Return(versionInfo, nil)

	kubelet.outOfDiskTransitionFrequency = 5 * time.Second

	ood := api.NodeCondition{
		Type:               api.NodeOutOfDisk,
		Status:             api.ConditionTrue,
		Reason:             "KubeletOutOfDisk",
		Message:            "out of disk space",
		LastHeartbeatTime:  unversioned.NewTime(clock.Now()), // placeholder
		LastTransitionTime: unversioned.NewTime(clock.Now()), // placeholder
	}
	noOod := api.NodeCondition{
		Type:               api.NodeOutOfDisk,
		Status:             api.ConditionFalse,
		Reason:             "KubeletHasSufficientDisk",
		Message:            fmt.Sprintf("kubelet has sufficient disk space available"),
		LastHeartbeatTime:  unversioned.NewTime(clock.Now()), // placeholder
		LastTransitionTime: unversioned.NewTime(clock.Now()), // placeholder
	}

	testCases := []struct {
		rootFsAvail   uint64
		dockerFsAvail uint64
		expected      api.NodeCondition
	}{
		{
			// NodeOutOfDisk==false
			rootFsAvail:   200,
			dockerFsAvail: 200,
			expected:      ood,
		},
		{
			// NodeOutOfDisk==true
			rootFsAvail:   50,
			dockerFsAvail: 200,
			expected:      ood,
		},
		{
			// NodeOutOfDisk==false
			rootFsAvail:   200,
			dockerFsAvail: 200,
			expected:      ood,
		},
		{
			// NodeOutOfDisk==true
			rootFsAvail:   200,
			dockerFsAvail: 50,
			expected:      ood,
		},
		{
			// NodeOutOfDisk==false
			rootFsAvail:   200,
			dockerFsAvail: 200,
			expected:      noOod,
		},
	}

	kubelet.updateRuntimeUp()
	for tcIdx, tc := range testCases {
		// Step by a second
		clock.Step(1 * time.Second)

		// Setup expected times.
		tc.expected.LastHeartbeatTime = unversioned.NewTime(clock.Now())
		// In the last case, there should be a status transition for NodeOutOfDisk
		if tcIdx == len(testCases)-1 {
			tc.expected.LastTransitionTime = unversioned.NewTime(clock.Now())
		}

		// Make kubelet report that it has sufficient disk space
		if err := updateDiskSpacePolicy(kubelet, mockCadvisor, 500, 500, tc.rootFsAvail, tc.dockerFsAvail, 100, 100); err != nil {
			t.Fatalf("can't update disk space manager: %v", err)
		}

		if err := kubelet.updateNodeStatus(); err != nil {
			t.Errorf("unexpected error: %v", err)
		}
		actions := kubeClient.Actions()
		if len(actions) != 2 {
			t.Errorf("%d. unexpected actions: %v", tcIdx, actions)
		}
		updateAction, ok := actions[1].(core.UpdateAction)
		if !ok {
			t.Errorf("%d. unexpected action type.  expected UpdateAction, got %#v", tcIdx, actions[1])
		}
		updatedNode, ok := updateAction.GetObject().(*api.Node)
		if !ok {
			t.Errorf("%d. unexpected object type", tcIdx)
		}
		kubeClient.ClearActions()

		var oodCondition api.NodeCondition
		for i, cond := range updatedNode.Status.Conditions {
			if cond.Type == api.NodeOutOfDisk {
				oodCondition = updatedNode.Status.Conditions[i]
			}
		}

		if !reflect.DeepEqual(tc.expected, oodCondition) {
			t.Errorf("%d.\nunexpected objects: %s", tcIdx, diff.ObjectDiff(tc.expected, oodCondition))

		}
	}
}

func TestUpdateNodeStatusWithRuntimeStateError(t *testing.T) {
	testKubelet := newTestKubelet(t, false /* controllerAttachDetachEnabled */)
	kubelet := testKubelet.kubelet
	clock := testKubelet.fakeClock
	kubeClient := testKubelet.fakeKubeClient
	kubeClient.ReactionChain = fake.NewSimpleClientset(&api.NodeList{Items: []api.Node{
		{ObjectMeta: api.ObjectMeta{Name: testKubeletHostname}},
	}}).ReactionChain
	mockCadvisor := testKubelet.fakeCadvisor
	mockCadvisor.On("Start").Return(nil)
	machineInfo := &cadvisorapi.MachineInfo{
		MachineID:      "123",
		SystemUUID:     "abc",
		BootID:         "1b3",
		NumCores:       2,
		MemoryCapacity: 10E9,
	}
	mockCadvisor.On("MachineInfo").Return(machineInfo, nil)
	versionInfo := &cadvisorapi.VersionInfo{
		KernelVersion:      "3.16.0-0.bpo.4-amd64",
		ContainerOsVersion: "Debian GNU/Linux 7 (wheezy)",
	}
	mockCadvisor.On("VersionInfo").Return(versionInfo, nil)

	// Make kubelet report that it has sufficient disk space.
	if err := updateDiskSpacePolicy(kubelet, mockCadvisor, 500, 500, 200, 200, 100, 100); err != nil {
		t.Fatalf("can't update disk space manager: %v", err)
	}

	expectedNode := &api.Node{
		ObjectMeta: api.ObjectMeta{Name: testKubeletHostname},
		Spec:       api.NodeSpec{},
		Status: api.NodeStatus{
			Conditions: []api.NodeCondition{
				{
					Type:               api.NodeOutOfDisk,
					Status:             api.ConditionFalse,
					Reason:             "KubeletHasSufficientDisk",
					Message:            "kubelet has sufficient disk space available",
					LastHeartbeatTime:  unversioned.Time{},
					LastTransitionTime: unversioned.Time{},
				},
				{
					Type:               api.NodeMemoryPressure,
					Status:             api.ConditionFalse,
					Reason:             "KubeletHasSufficientMemory",
					Message:            fmt.Sprintf("kubelet has sufficient memory available"),
					LastHeartbeatTime:  unversioned.Time{},
					LastTransitionTime: unversioned.Time{},
				},
				{
					Type:               api.NodeDiskPressure,
					Status:             api.ConditionFalse,
					Reason:             "KubeletHasNoDiskPressure",
					Message:            fmt.Sprintf("kubelet has no disk pressure"),
					LastHeartbeatTime:  unversioned.Time{},
					LastTransitionTime: unversioned.Time{},
				},
				{}, //placeholder
			},
			NodeInfo: api.NodeSystemInfo{
				MachineID:               "123",
				SystemUUID:              "abc",
				BootID:                  "1b3",
				KernelVersion:           "3.16.0-0.bpo.4-amd64",
				OSImage:                 "Debian GNU/Linux 7 (wheezy)",
				OperatingSystem:         goruntime.GOOS,
				Architecture:            goruntime.GOARCH,
				ContainerRuntimeVersion: "test://1.5.0",
				KubeletVersion:          version.Get().String(),
				KubeProxyVersion:        version.Get().String(),
			},
			Capacity: api.ResourceList{
				api.ResourceCPU:       *resource.NewMilliQuantity(2000, resource.DecimalSI),
				api.ResourceMemory:    *resource.NewQuantity(10E9, resource.BinarySI),
				api.ResourcePods:      *resource.NewQuantity(0, resource.DecimalSI),
				api.ResourceNvidiaGPU: *resource.NewQuantity(0, resource.DecimalSI),
			},
			Allocatable: api.ResourceList{
				api.ResourceCPU:       *resource.NewMilliQuantity(1800, resource.DecimalSI),
				api.ResourceMemory:    *resource.NewQuantity(9900E6, resource.BinarySI),
				api.ResourcePods:      *resource.NewQuantity(0, resource.DecimalSI),
				api.ResourceNvidiaGPU: *resource.NewQuantity(0, resource.DecimalSI),
			},
			Addresses: []api.NodeAddress{
				{Type: api.NodeLegacyHostIP, Address: "127.0.0.1"},
				{Type: api.NodeInternalIP, Address: "127.0.0.1"},
			},
			Images: []api.ContainerImage{
				{
					Names:     []string{"gcr.io/google_containers:v3", "gcr.io/google_containers:v4"},
					SizeBytes: 456,
				},
				{
					Names:     []string{"gcr.io/google_containers:v1", "gcr.io/google_containers:v2"},
					SizeBytes: 123,
				},
			},
		},
	}

	checkNodeStatus := func(status api.ConditionStatus, reason, message string) {
		kubeClient.ClearActions()
		if err := kubelet.updateNodeStatus(); err != nil {
			t.Errorf("unexpected error: %v", err)
		}
		actions := kubeClient.Actions()
		if len(actions) != 2 {
			t.Fatalf("unexpected actions: %v", actions)
		}
		if !actions[1].Matches("update", "nodes") || actions[1].GetSubresource() != "status" {
			t.Fatalf("unexpected actions: %v", actions)
		}
		updatedNode, ok := actions[1].(core.UpdateAction).GetObject().(*api.Node)
		if !ok {
			t.Errorf("unexpected action type.  expected UpdateAction, got %#v", actions[1])
		}

		for i, cond := range updatedNode.Status.Conditions {
			if cond.LastHeartbeatTime.IsZero() {
				t.Errorf("unexpected zero last probe timestamp")
			}
			if cond.LastTransitionTime.IsZero() {
				t.Errorf("unexpected zero last transition timestamp")
			}
			updatedNode.Status.Conditions[i].LastHeartbeatTime = unversioned.Time{}
			updatedNode.Status.Conditions[i].LastTransitionTime = unversioned.Time{}
		}

		// Version skew workaround. See: https://github.com/kubernetes/kubernetes/issues/16961
		lastIndex := len(updatedNode.Status.Conditions) - 1
		if updatedNode.Status.Conditions[lastIndex].Type != api.NodeReady {
			t.Errorf("unexpected node condition order. NodeReady should be last.")
		}
		expectedNode.Status.Conditions[lastIndex] = api.NodeCondition{
			Type:               api.NodeReady,
			Status:             status,
			Reason:             reason,
			Message:            message,
			LastHeartbeatTime:  unversioned.Time{},
			LastTransitionTime: unversioned.Time{},
		}
		if !api.Semantic.DeepEqual(expectedNode, updatedNode) {
			t.Errorf("unexpected objects: %s", diff.ObjectDiff(expectedNode, updatedNode))
		}
	}

	readyMessage := "kubelet is posting ready status"
	downMessage := "container runtime is down"

	// Should report kubelet not ready if the runtime check is out of date
	clock.SetTime(time.Now().Add(-maxWaitForContainerRuntime))
	kubelet.updateRuntimeUp()
	checkNodeStatus(api.ConditionFalse, "KubeletNotReady", downMessage)

	// Should report kubelet ready if the runtime check is updated
	clock.SetTime(time.Now())
	kubelet.updateRuntimeUp()
	checkNodeStatus(api.ConditionTrue, "KubeletReady", readyMessage)

	// Should report kubelet not ready if the runtime check is out of date
	clock.SetTime(time.Now().Add(-maxWaitForContainerRuntime))
	kubelet.updateRuntimeUp()
	checkNodeStatus(api.ConditionFalse, "KubeletNotReady", downMessage)

	// Should report kubelet not ready if the runtime check failed
	fakeRuntime := testKubelet.fakeRuntime
	// Inject error into fake runtime status check, node should be NotReady
	fakeRuntime.StatusErr = fmt.Errorf("injected runtime status error")
	clock.SetTime(time.Now())
	kubelet.updateRuntimeUp()
	checkNodeStatus(api.ConditionFalse, "KubeletNotReady", downMessage)
}

func TestUpdateNodeStatusError(t *testing.T) {
	testKubelet := newTestKubelet(t, false /* controllerAttachDetachEnabled */)
	kubelet := testKubelet.kubelet
	// No matching node for the kubelet
	testKubelet.fakeKubeClient.ReactionChain = fake.NewSimpleClientset(&api.NodeList{Items: []api.Node{}}).ReactionChain

	if err := kubelet.updateNodeStatus(); err == nil {
		t.Errorf("unexpected non error: %v", err)
	}
	if len(testKubelet.fakeKubeClient.Actions()) != nodeStatusUpdateRetry {
		t.Errorf("unexpected actions: %v", testKubelet.fakeKubeClient.Actions())
	}
}

func TestRegisterWithApiServer(t *testing.T) {
	testKubelet := newTestKubelet(t, false /* controllerAttachDetachEnabled */)
	kubelet := testKubelet.kubelet
	kubeClient := testKubelet.fakeKubeClient
	kubeClient.AddReactor("create", "nodes", func(action core.Action) (bool, runtime.Object, error) {
		// Return an error on create.
		return true, &api.Node{}, &apierrors.StatusError{
			ErrStatus: unversioned.Status{Reason: unversioned.StatusReasonAlreadyExists},
		}
	})
	kubeClient.AddReactor("get", "nodes", func(action core.Action) (bool, runtime.Object, error) {
		// Return an existing (matching) node on get.
		return true, &api.Node{
			ObjectMeta: api.ObjectMeta{Name: testKubeletHostname},
			Spec:       api.NodeSpec{ExternalID: testKubeletHostname},
		}, nil
	})
	kubeClient.AddReactor("*", "*", func(action core.Action) (bool, runtime.Object, error) {
		return true, nil, fmt.Errorf("no reaction implemented for %s", action)
	})
	machineInfo := &cadvisorapi.MachineInfo{
		MachineID:      "123",
		SystemUUID:     "abc",
		BootID:         "1b3",
		NumCores:       2,
		MemoryCapacity: 1024,
	}
	mockCadvisor := testKubelet.fakeCadvisor
	mockCadvisor.On("MachineInfo").Return(machineInfo, nil)
	versionInfo := &cadvisorapi.VersionInfo{
		KernelVersion:      "3.16.0-0.bpo.4-amd64",
		ContainerOsVersion: "Debian GNU/Linux 7 (wheezy)",
		DockerVersion:      "1.5.0",
	}
	mockCadvisor.On("VersionInfo").Return(versionInfo, nil)
	mockCadvisor.On("ImagesFsInfo").Return(cadvisorapiv2.FsInfo{
		Usage:     400 * mb,
		Capacity:  1000 * mb,
		Available: 600 * mb,
	}, nil)
	mockCadvisor.On("RootFsInfo").Return(cadvisorapiv2.FsInfo{
		Usage:    9 * mb,
		Capacity: 10 * mb,
	}, nil)

	done := make(chan struct{})
	go func() {
		kubelet.registerWithApiServer()
		done <- struct{}{}
	}()
	select {
	case <-time.After(wait.ForeverTestTimeout):
		t.Errorf("timed out waiting for registration")
	case <-done:
		return
	}
}

func TestTryRegisterWithApiServer(t *testing.T) {
	alreadyExists := &apierrors.StatusError{
		ErrStatus: unversioned.Status{Reason: unversioned.StatusReasonAlreadyExists},
	}

	conflict := &apierrors.StatusError{
		ErrStatus: unversioned.Status{Reason: unversioned.StatusReasonConflict},
	}

	newNode := func(cmad bool, externalID string) *api.Node {
		node := &api.Node{
			ObjectMeta: api.ObjectMeta{},
			Spec: api.NodeSpec{
				ExternalID: externalID,
			},
		}

		if cmad {
			node.Annotations = make(map[string]string)
			node.Annotations[volumehelper.ControllerManagedAttachAnnotation] = "true"
		}

		return node
	}

	cases := []struct {
		name            string
		newNode         *api.Node
		existingNode    *api.Node
		createError     error
		getError        error
		updateError     error
		deleteError     error
		expectedResult  bool
		expectedActions int
		testSavedNode   bool
		savedNodeIndex  int
		savedNodeCMAD   bool
	}{
		{
			name:            "success case - new node",
			newNode:         &api.Node{},
			expectedResult:  true,
			expectedActions: 1,
		},
		{
			name:            "success case - existing node - no change in CMAD",
			newNode:         newNode(true, "a"),
			createError:     alreadyExists,
			existingNode:    newNode(true, "a"),
			expectedResult:  true,
			expectedActions: 2,
		},
		{
			name:            "success case - existing node - CMAD disabled",
			newNode:         newNode(false, "a"),
			createError:     alreadyExists,
			existingNode:    newNode(true, "a"),
			expectedResult:  true,
			expectedActions: 3,
			testSavedNode:   true,
			savedNodeIndex:  2,
			savedNodeCMAD:   false,
		},
		{
			name:            "success case - existing node - CMAD enabled",
			newNode:         newNode(true, "a"),
			createError:     alreadyExists,
			existingNode:    newNode(false, "a"),
			expectedResult:  true,
			expectedActions: 3,
			testSavedNode:   true,
			savedNodeIndex:  2,
			savedNodeCMAD:   true,
		},
		{
			name:            "success case - external ID changed",
			newNode:         newNode(false, "b"),
			createError:     alreadyExists,
			existingNode:    newNode(false, "a"),
			expectedResult:  false,
			expectedActions: 3,
		},
		{
			name:            "create failed",
			newNode:         newNode(false, "b"),
			createError:     conflict,
			expectedResult:  false,
			expectedActions: 1,
		},
		{
			name:            "get existing node failed",
			newNode:         newNode(false, "a"),
			createError:     alreadyExists,
			getError:        conflict,
			expectedResult:  false,
			expectedActions: 2,
		},
		{
			name:            "update existing node failed",
			newNode:         newNode(false, "a"),
			createError:     alreadyExists,
			existingNode:    newNode(true, "a"),
			updateError:     conflict,
			expectedResult:  false,
			expectedActions: 3,
		},
		{
			name:            "delete existing node failed",
			newNode:         newNode(false, "b"),
			createError:     alreadyExists,
			existingNode:    newNode(false, "a"),
			deleteError:     conflict,
			expectedResult:  false,
			expectedActions: 3,
		},
	}

	notImplemented := func(action core.Action) (bool, runtime.Object, error) {
		return true, nil, fmt.Errorf("no reaction implemented for %s", action)
	}

	for _, tc := range cases {
		testKubelet := newTestKubelet(t, false /* controllerAttachDetachEnabled is a don't-care for this test */)
		kubelet := testKubelet.kubelet
		kubeClient := testKubelet.fakeKubeClient

		kubeClient.AddReactor("create", "nodes", func(action core.Action) (bool, runtime.Object, error) {
			return true, nil, tc.createError
		})
		kubeClient.AddReactor("get", "nodes", func(action core.Action) (bool, runtime.Object, error) {
			// Return an existing (matching) node on get.
			return true, tc.existingNode, tc.getError
		})
		kubeClient.AddReactor("update", "nodes", func(action core.Action) (bool, runtime.Object, error) {
			if action.GetSubresource() == "status" {
				return true, nil, tc.updateError
			}
			return notImplemented(action)
		})
		kubeClient.AddReactor("delete", "nodes", func(action core.Action) (bool, runtime.Object, error) {
			return true, nil, tc.deleteError
		})
		kubeClient.AddReactor("*", "*", func(action core.Action) (bool, runtime.Object, error) {
			return notImplemented(action)
		})

		result := kubelet.tryRegisterWithApiServer(tc.newNode)
		if e, a := tc.expectedResult, result; e != a {
			t.Errorf("%v: unexpected result; expected %v got %v", tc.name, e, a)
			continue
		}

		actions := kubeClient.Actions()
		if e, a := tc.expectedActions, len(actions); e != a {
			t.Errorf("%v: unexpected number of actions, expected %v, got %v", tc.name, e, a)
		}

		if tc.testSavedNode {
			var savedNode *api.Node
			var ok bool

			t.Logf("actions: %v: %+v", len(actions), actions)
			action := actions[tc.savedNodeIndex]
			if action.GetVerb() == "create" {
				createAction := action.(core.CreateAction)
				savedNode, ok = createAction.GetObject().(*api.Node)
				if !ok {
					t.Errorf("%v: unexpected type; couldn't convert to *api.Node: %+v", tc.name, createAction.GetObject())
					continue
				}
			} else if action.GetVerb() == "update" {
				updateAction := action.(core.UpdateAction)
				savedNode, ok = updateAction.GetObject().(*api.Node)
				if !ok {
					t.Errorf("%v: unexpected type; couldn't convert to *api.Node: %+v", tc.name, updateAction.GetObject())
					continue
				}
			}

			actualCMAD, _ := strconv.ParseBool(savedNode.Annotations[volumehelper.ControllerManagedAttachAnnotation])
			if e, a := tc.savedNodeCMAD, actualCMAD; e != a {
				t.Errorf("%v: unexpected attach-detach value on saved node; expected %v got %v", tc.name, e, a)
			}
		}
	}
}
