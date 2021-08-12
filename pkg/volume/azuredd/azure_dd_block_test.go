//go:build !providerless
// +build !providerless

/*
Copyright 2018 The Kubernetes Authors.

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

package azuredd

import (
	"os"
	"path/filepath"
	"testing"

	"k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"
	utiltesting "k8s.io/client-go/util/testing"
	"k8s.io/kubernetes/pkg/volume"
	volumetest "k8s.io/kubernetes/pkg/volume/testing"
)

const (
	testDiskName   = "disk1"
	testPVName     = "pv1"
	testGlobalPath = "plugins/kubernetes.io/azure-disk/volumeDevices/disk1"
	testPodPath    = "pods/poduid/volumeDevices/kubernetes.io~azure-disk"
)

func TestGetVolumeSpecFromGlobalMapPath(t *testing.T) {
	// make our test path for fake GlobalMapPath
	// /tmp symbolized our pluginDir
	// /tmp/testGlobalPathXXXXX/plugins/kubernetes.io/azure-disk/volumeDevices/disk1
	tmpVDir, err := utiltesting.MkTmpdir("azureDiskBlockTest")
	if err != nil {
		t.Fatalf("can't make a temp dir: %v", err)
	}
	//deferred clean up
	defer os.RemoveAll(tmpVDir)

	expectedGlobalPath := filepath.Join(tmpVDir, testGlobalPath)

	//Bad Path
	badspec, err := getVolumeSpecFromGlobalMapPath("", "")
	if badspec != nil || err == nil {
		t.Errorf("Expected not to get spec from GlobalMapPath but did")
	}

	// Good Path
	spec, err := getVolumeSpecFromGlobalMapPath(expectedGlobalPath, "")
	if spec == nil || err != nil {
		t.Fatalf("Failed to get spec from GlobalMapPath: %v", err)
	}
	if spec.PersistentVolume.Spec.AzureDisk.DiskName != testDiskName {
		t.Errorf("Invalid pdName from GlobalMapPath spec: %s", spec.PersistentVolume.Spec.AzureDisk.DiskName)
	}
	block := v1.PersistentVolumeBlock
	specMode := spec.PersistentVolume.Spec.VolumeMode
	if specMode == nil {
		t.Errorf("Invalid volumeMode from GlobalMapPath spec: %v expected: %v", specMode, block)
	}
	if *specMode != block {
		t.Errorf("Invalid volumeMode from GlobalMapPath spec: %v expected: %v", *specMode, block)
	}
}

func getTestVolume(readOnly bool, path string, isBlock bool) *volume.Spec {
	pv := &v1.PersistentVolume{
		ObjectMeta: metav1.ObjectMeta{
			Name: testPVName,
		},
		Spec: v1.PersistentVolumeSpec{
			PersistentVolumeSource: v1.PersistentVolumeSource{
				AzureDisk: &v1.AzureDiskVolumeSource{
					DiskName: testDiskName,
				},
			},
		},
	}

	if isBlock {
		blockMode := v1.PersistentVolumeBlock
		pv.Spec.VolumeMode = &blockMode
	}
	return volume.NewSpecFromPersistentVolume(pv, readOnly)
}

func TestGetPodAndPluginMapPaths(t *testing.T) {
	tmpVDir, err := utiltesting.MkTmpdir("azureDiskBlockTest")
	if err != nil {
		t.Fatalf("can't make a temp dir: %v", err)
	}
	//deferred clean up
	defer os.RemoveAll(tmpVDir)

	expectedGlobalPath := filepath.Join(tmpVDir, testGlobalPath)
	expectedPodPath := filepath.Join(tmpVDir, testPodPath)

	spec := getTestVolume(false, tmpVDir, true /*isBlock*/)
	plugMgr := volume.VolumePluginMgr{}
	plugMgr.InitPlugins(ProbeVolumePlugins(), nil /* prober */, volumetest.NewFakeVolumeHost(t, tmpVDir, nil, nil))
	plug, err := plugMgr.FindMapperPluginByName(azureDataDiskPluginName)
	if err != nil {
		os.RemoveAll(tmpVDir)
		t.Fatalf("Can't find the plugin by name: %q", azureDataDiskPluginName)
	}
	if plug.GetPluginName() != azureDataDiskPluginName {
		t.Fatalf("Wrong name: %s", plug.GetPluginName())
	}
	pod := &v1.Pod{ObjectMeta: metav1.ObjectMeta{UID: types.UID("poduid")}}
	mapper, err := plug.NewBlockVolumeMapper(spec, pod, volume.VolumeOptions{})
	if err != nil {
		t.Fatalf("Failed to make a new Mounter: %v", err)
	}
	if mapper == nil {
		t.Fatalf("Got a nil Mounter")
	}

	//GetGlobalMapPath
	gMapPath, err := mapper.GetGlobalMapPath(spec)
	if err != nil || len(gMapPath) == 0 {
		t.Fatalf("Invalid GlobalMapPath from spec: %s, error: %v", spec.PersistentVolume.Spec.AzureDisk.DiskName, err)
	}
	if gMapPath != expectedGlobalPath {
		t.Errorf("Failed to get GlobalMapPath: %s, expected %s", gMapPath, expectedGlobalPath)
	}

	//GetPodDeviceMapPath
	gDevicePath, gVolName := mapper.GetPodDeviceMapPath()
	if gDevicePath != expectedPodPath {
		t.Errorf("Got unexpected pod path: %s, expected %s", gDevicePath, expectedPodPath)
	}
	if gVolName != testPVName {
		t.Errorf("Got unexpected volNamne: %s, expected %s", gVolName, testPVName)
	}
}
