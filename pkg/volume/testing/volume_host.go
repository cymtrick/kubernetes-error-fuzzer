/*
Copyright 2020 The Kubernetes Authors.

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
	"context"
	"fmt"
	"net"
	"path/filepath"
	"sync"
	"testing"
	"time"

	authenticationv1 "k8s.io/api/authentication/v1"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/apimachinery/pkg/util/wait"
	"k8s.io/client-go/informers"
	clientset "k8s.io/client-go/kubernetes"
	storagelistersv1 "k8s.io/client-go/listers/storage/v1"
	"k8s.io/client-go/tools/cache"
	"k8s.io/client-go/tools/record"
	cloudprovider "k8s.io/cloud-provider"
	proxyutil "k8s.io/kubernetes/pkg/proxy/util"
	. "k8s.io/kubernetes/pkg/volume"
	"k8s.io/kubernetes/pkg/volume/util/hostutil"
	"k8s.io/kubernetes/pkg/volume/util/subpath"
	"k8s.io/mount-utils"
	"k8s.io/utils/exec"
	testingexec "k8s.io/utils/exec/testing"
)

// fakeVolumeHost is useful for testing volume plugins.
type fakeVolumeHost struct {
	rootDir                string
	kubeClient             clientset.Interface
	pluginMgr              *VolumePluginMgr
	cloud                  cloudprovider.Interface
	mounter                mount.Interface
	hostUtil               hostutil.HostUtils
	exec                   *testingexec.FakeExec
	nodeLabels             map[string]string
	nodeName               string
	subpather              subpath.Interface
	csiDriverLister        storagelistersv1.CSIDriverLister
	volumeAttachmentLister storagelistersv1.VolumeAttachmentLister
	informerFactory        informers.SharedInformerFactory
	kubeletErr             error
	mux                    sync.Mutex
	filteredDialOptions    *proxyutil.FilteredDialOptions
}

var _ VolumeHost = &fakeVolumeHost{}
var _ AttachDetachVolumeHost = &fakeVolumeHost{}
var _ KubeletVolumeHost = &fakeVolumeHost{}

func NewFakeVolumeHost(t *testing.T, rootDir string, kubeClient clientset.Interface, plugins []VolumePlugin) *fakeVolumeHost {
	return newFakeVolumeHost(t, rootDir, kubeClient, plugins, nil, nil, "", nil, nil)
}

func NewFakeVolumeHostWithCloudProvider(t *testing.T, rootDir string, kubeClient clientset.Interface, plugins []VolumePlugin, cloud cloudprovider.Interface) *fakeVolumeHost {
	return newFakeVolumeHost(t, rootDir, kubeClient, plugins, cloud, nil, "", nil, nil)
}

func NewFakeVolumeHostWithNodeLabels(t *testing.T, rootDir string, kubeClient clientset.Interface, plugins []VolumePlugin, labels map[string]string) *fakeVolumeHost {
	volHost := newFakeVolumeHost(t, rootDir, kubeClient, plugins, nil, nil, "", nil, nil)
	volHost.nodeLabels = labels
	return volHost
}

func NewFakeVolumeHostWithCSINodeName(t *testing.T, rootDir string, kubeClient clientset.Interface, plugins []VolumePlugin, nodeName string, driverLister storagelistersv1.CSIDriverLister, volumeAttachLister storagelistersv1.VolumeAttachmentLister) *fakeVolumeHost {
	return newFakeVolumeHost(t, rootDir, kubeClient, plugins, nil, nil, nodeName, driverLister, volumeAttachLister)
}

func NewFakeVolumeHostWithMounterFSType(t *testing.T, rootDir string, kubeClient clientset.Interface, plugins []VolumePlugin, pathToTypeMap map[string]hostutil.FileType) *fakeVolumeHost {
	return newFakeVolumeHost(t, rootDir, kubeClient, plugins, nil, pathToTypeMap, "", nil, nil)
}

func newFakeVolumeHost(t *testing.T, rootDir string, kubeClient clientset.Interface, plugins []VolumePlugin, cloud cloudprovider.Interface, pathToTypeMap map[string]hostutil.FileType, nodeName string, driverLister storagelistersv1.CSIDriverLister, volumeAttachLister storagelistersv1.VolumeAttachmentLister) *fakeVolumeHost {
	host := &fakeVolumeHost{rootDir: rootDir, kubeClient: kubeClient, cloud: cloud, nodeName: nodeName, csiDriverLister: driverLister, volumeAttachmentLister: volumeAttachLister}
	host.mounter = mount.NewFakeMounter(nil)
	host.hostUtil = hostutil.NewFakeHostUtil(pathToTypeMap)
	host.exec = &testingexec.FakeExec{DisableScripts: true}
	host.pluginMgr = &VolumePluginMgr{}
	if err := host.pluginMgr.InitPlugins(plugins, nil /* prober */, host); err != nil {
		t.Fatalf("Failed to init plugins while creating fake volume host: %v", err)
	}
	host.subpather = &subpath.FakeSubpath{}
	host.informerFactory = informers.NewSharedInformerFactory(kubeClient, time.Minute)
	// Wait until the InitPlugins setup is finished before returning from this setup func
	if err := host.WaitForKubeletErrNil(); err != nil {
		t.Fatalf("Failed to wait for kubelet err to be nil while creating fake volume host: %v", err)
	}
	return host
}

func (f *fakeVolumeHost) GetPluginDir(podUID string) string {
	return filepath.Join(f.rootDir, "plugins", podUID)
}

func (f *fakeVolumeHost) GetVolumeDevicePluginDir(pluginName string) string {
	return filepath.Join(f.rootDir, "plugins", pluginName, "volumeDevices")
}

func (f *fakeVolumeHost) GetPodsDir() string {
	return filepath.Join(f.rootDir, "pods")
}

func (f *fakeVolumeHost) GetPodVolumeDir(podUID types.UID, pluginName, volumeName string) string {
	return filepath.Join(f.rootDir, "pods", string(podUID), "volumes", pluginName, volumeName)
}

func (f *fakeVolumeHost) GetPodVolumeDeviceDir(podUID types.UID, pluginName string) string {
	return filepath.Join(f.rootDir, "pods", string(podUID), "volumeDevices", pluginName)
}

func (f *fakeVolumeHost) GetPodPluginDir(podUID types.UID, pluginName string) string {
	return filepath.Join(f.rootDir, "pods", string(podUID), "plugins", pluginName)
}

func (f *fakeVolumeHost) GetKubeClient() clientset.Interface {
	return f.kubeClient
}

func (f *fakeVolumeHost) GetCloudProvider() cloudprovider.Interface {
	return f.cloud
}

func (f *fakeVolumeHost) GetMounter(pluginName string) mount.Interface {
	return f.mounter
}

func (f *fakeVolumeHost) GetHostUtil() hostutil.HostUtils {
	return f.hostUtil
}

func (f *fakeVolumeHost) GetSubpather() subpath.Interface {
	return f.subpather
}

func (f *fakeVolumeHost) GetFilteredDialOptions() *proxyutil.FilteredDialOptions {
	return f.filteredDialOptions
}

func (f *fakeVolumeHost) GetPluginMgr() *VolumePluginMgr {
	return f.pluginMgr
}

func (f *fakeVolumeHost) NewWrapperMounter(volName string, spec Spec, pod *v1.Pod, opts VolumeOptions) (Mounter, error) {
	// The name of wrapper volume is set to "wrapped_{wrapped_volume_name}"
	wrapperVolumeName := "wrapped_" + volName
	if spec.Volume != nil {
		spec.Volume.Name = wrapperVolumeName
	}
	plug, err := f.pluginMgr.FindPluginBySpec(&spec)
	if err != nil {
		return nil, err
	}
	return plug.NewMounter(&spec, pod, opts)
}

func (f *fakeVolumeHost) NewWrapperUnmounter(volName string, spec Spec, podUID types.UID) (Unmounter, error) {
	// The name of wrapper volume is set to "wrapped_{wrapped_volume_name}"
	wrapperVolumeName := "wrapped_" + volName
	if spec.Volume != nil {
		spec.Volume.Name = wrapperVolumeName
	}
	plug, err := f.pluginMgr.FindPluginBySpec(&spec)
	if err != nil {
		return nil, err
	}
	return plug.NewUnmounter(spec.Name(), podUID)
}

// Returns the hostname of the host kubelet is running on
func (f *fakeVolumeHost) GetHostName() string {
	return "fakeHostName"
}

// Returns host IP or nil in the case of error.
func (f *fakeVolumeHost) GetHostIP() (net.IP, error) {
	return nil, fmt.Errorf("GetHostIP() not implemented")
}

func (f *fakeVolumeHost) GetNodeAllocatable() (v1.ResourceList, error) {
	return v1.ResourceList{}, nil
}

func (f *fakeVolumeHost) GetSecretFunc() func(namespace, name string) (*v1.Secret, error) {
	return func(namespace, name string) (*v1.Secret, error) {
		return f.kubeClient.CoreV1().Secrets(namespace).Get(context.TODO(), name, metav1.GetOptions{})
	}
}

func (f *fakeVolumeHost) GetExec(pluginName string) exec.Interface {
	return f.exec
}

func (f *fakeVolumeHost) GetConfigMapFunc() func(namespace, name string) (*v1.ConfigMap, error) {
	return func(namespace, name string) (*v1.ConfigMap, error) {
		return f.kubeClient.CoreV1().ConfigMaps(namespace).Get(context.TODO(), name, metav1.GetOptions{})
	}
}

func (f *fakeVolumeHost) GetServiceAccountTokenFunc() func(string, string, *authenticationv1.TokenRequest) (*authenticationv1.TokenRequest, error) {
	return func(namespace, name string, tr *authenticationv1.TokenRequest) (*authenticationv1.TokenRequest, error) {
		return f.kubeClient.CoreV1().ServiceAccounts(namespace).CreateToken(context.TODO(), name, tr, metav1.CreateOptions{})
	}
}

func (f *fakeVolumeHost) DeleteServiceAccountTokenFunc() func(types.UID) {
	return func(types.UID) {}
}

func (f *fakeVolumeHost) GetNodeLabels() (map[string]string, error) {
	if f.nodeLabels == nil {
		f.nodeLabels = map[string]string{"test-label": "test-value"}
	}
	return f.nodeLabels, nil
}

func (f *fakeVolumeHost) GetNodeName() types.NodeName {
	return types.NodeName(f.nodeName)
}

func (f *fakeVolumeHost) GetEventRecorder() record.EventRecorder {
	return nil
}

func (f *fakeVolumeHost) ScriptCommands(scripts []CommandScript) {
	ScriptCommands(f.exec, scripts)
}

func (f *fakeVolumeHost) CSIDriverLister() storagelistersv1.CSIDriverLister {
	return f.csiDriverLister
}

func (f *fakeVolumeHost) VolumeAttachmentLister() storagelistersv1.VolumeAttachmentLister {
	return f.volumeAttachmentLister
}

func (f *fakeVolumeHost) CSIDriversSynced() cache.InformerSynced {
	// not needed for testing
	return nil
}

func (f *fakeVolumeHost) CSINodeLister() storagelistersv1.CSINodeLister {
	// not needed for testing
	return nil
}

func (f *fakeVolumeHost) GetInformerFactory() informers.SharedInformerFactory {
	return f.informerFactory
}

func (f *fakeVolumeHost) IsAttachDetachController() bool {
	return true
}

func (f *fakeVolumeHost) SetKubeletError(err error) {
	f.mux.Lock()
	defer f.mux.Unlock()
	f.kubeletErr = err
	return
}

func (f *fakeVolumeHost) WaitForCacheSync() error {
	return nil
}

func (f *fakeVolumeHost) WaitForKubeletErrNil() error {
	return wait.PollImmediate(10*time.Millisecond, 10*time.Second, func() (bool, error) {
		f.mux.Lock()
		defer f.mux.Unlock()
		return f.kubeletErr == nil, nil
	})
}
