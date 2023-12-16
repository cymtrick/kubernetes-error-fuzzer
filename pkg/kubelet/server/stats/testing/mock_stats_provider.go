/*
Copyright The Kubernetes Authors.

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

// Code generated by MockGen. DO NOT EDIT.
// Source: handler.go

// Package testing is a generated GoMock package.
package testing

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	v1 "github.com/google/cadvisor/info/v1"
	v2 "github.com/google/cadvisor/info/v2"
	v10 "k8s.io/api/core/v1"
	types "k8s.io/apimachinery/pkg/types"
	v1alpha1 "k8s.io/kubelet/pkg/apis/stats/v1alpha1"
	cm "k8s.io/kubernetes/pkg/kubelet/cm"
	volume "k8s.io/kubernetes/pkg/volume"
)

// MockProvider is a mock of Provider interface.
type MockProvider struct {
	ctrl     *gomock.Controller
	recorder *MockProviderMockRecorder
}

// MockProviderMockRecorder is the mock recorder for MockProvider.
type MockProviderMockRecorder struct {
	mock *MockProvider
}

// NewMockProvider creates a new mock instance.
func NewMockProvider(ctrl *gomock.Controller) *MockProvider {
	mock := &MockProvider{ctrl: ctrl}
	mock.recorder = &MockProviderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockProvider) EXPECT() *MockProviderMockRecorder {
	return m.recorder
}

// GetCgroupCPUAndMemoryStats mocks base method.
func (m *MockProvider) GetCgroupCPUAndMemoryStats(cgroupName string, updateStats bool) (*v1alpha1.ContainerStats, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCgroupCPUAndMemoryStats", cgroupName, updateStats)
	ret0, _ := ret[0].(*v1alpha1.ContainerStats)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetCgroupCPUAndMemoryStats indicates an expected call of GetCgroupCPUAndMemoryStats.
func (mr *MockProviderMockRecorder) GetCgroupCPUAndMemoryStats(cgroupName, updateStats interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCgroupCPUAndMemoryStats", reflect.TypeOf((*MockProvider)(nil).GetCgroupCPUAndMemoryStats), cgroupName, updateStats)
}

// GetCgroupStats mocks base method.
func (m *MockProvider) GetCgroupStats(cgroupName string, updateStats bool) (*v1alpha1.ContainerStats, *v1alpha1.NetworkStats, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCgroupStats", cgroupName, updateStats)
	ret0, _ := ret[0].(*v1alpha1.ContainerStats)
	ret1, _ := ret[1].(*v1alpha1.NetworkStats)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// GetCgroupStats indicates an expected call of GetCgroupStats.
func (mr *MockProviderMockRecorder) GetCgroupStats(cgroupName, updateStats interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCgroupStats", reflect.TypeOf((*MockProvider)(nil).GetCgroupStats), cgroupName, updateStats)
}

// GetNode mocks base method.
func (m *MockProvider) GetNode() (*v10.Node, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetNode")
	ret0, _ := ret[0].(*v10.Node)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetNode indicates an expected call of GetNode.
func (mr *MockProviderMockRecorder) GetNode() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetNode", reflect.TypeOf((*MockProvider)(nil).GetNode))
}

// GetNodeConfig mocks base method.
func (m *MockProvider) GetNodeConfig() cm.NodeConfig {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetNodeConfig")
	ret0, _ := ret[0].(cm.NodeConfig)
	return ret0
}

// GetNodeConfig indicates an expected call of GetNodeConfig.
func (mr *MockProviderMockRecorder) GetNodeConfig() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetNodeConfig", reflect.TypeOf((*MockProvider)(nil).GetNodeConfig))
}

// GetPodByCgroupfs mocks base method.
func (m *MockProvider) GetPodByCgroupfs(cgroupfs string) (*v10.Pod, bool) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPodByCgroupfs", cgroupfs)
	ret0, _ := ret[0].(*v10.Pod)
	ret1, _ := ret[1].(bool)
	return ret0, ret1
}

// GetPodByCgroupfs indicates an expected call of GetPodByCgroupfs.
func (mr *MockProviderMockRecorder) GetPodByCgroupfs(cgroupfs interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPodByCgroupfs", reflect.TypeOf((*MockProvider)(nil).GetPodByCgroupfs), cgroupfs)
}

// GetPodByName mocks base method.
func (m *MockProvider) GetPodByName(namespace, name string) (*v10.Pod, bool) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPodByName", namespace, name)
	ret0, _ := ret[0].(*v10.Pod)
	ret1, _ := ret[1].(bool)
	return ret0, ret1
}

// GetPodByName indicates an expected call of GetPodByName.
func (mr *MockProviderMockRecorder) GetPodByName(namespace, name interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPodByName", reflect.TypeOf((*MockProvider)(nil).GetPodByName), namespace, name)
}

// GetPodCgroupRoot mocks base method.
func (m *MockProvider) GetPodCgroupRoot() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPodCgroupRoot")
	ret0, _ := ret[0].(string)
	return ret0
}

// GetPodCgroupRoot indicates an expected call of GetPodCgroupRoot.
func (mr *MockProviderMockRecorder) GetPodCgroupRoot() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPodCgroupRoot", reflect.TypeOf((*MockProvider)(nil).GetPodCgroupRoot))
}

// GetPods mocks base method.
func (m *MockProvider) GetPods() []*v10.Pod {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPods")
	ret0, _ := ret[0].([]*v10.Pod)
	return ret0
}

// GetPods indicates an expected call of GetPods.
func (mr *MockProviderMockRecorder) GetPods() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPods", reflect.TypeOf((*MockProvider)(nil).GetPods))
}

// GetRequestedContainersInfo mocks base method.
func (m *MockProvider) GetRequestedContainersInfo(containerName string, options v2.RequestOptions) (map[string]*v1.ContainerInfo, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetRequestedContainersInfo", containerName, options)
	ret0, _ := ret[0].(map[string]*v1.ContainerInfo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetRequestedContainersInfo indicates an expected call of GetRequestedContainersInfo.
func (mr *MockProviderMockRecorder) GetRequestedContainersInfo(containerName, options interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetRequestedContainersInfo", reflect.TypeOf((*MockProvider)(nil).GetRequestedContainersInfo), containerName, options)
}

// ImageFsStats mocks base method.
func (m *MockProvider) ImageFsStats(ctx context.Context) (*v1alpha1.FsStats, *v1alpha1.FsStats, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ImageFsStats", ctx)
	ret0, _ := ret[0].(*v1alpha1.FsStats)
	ret1, _ := ret[1].(*v1alpha1.FsStats)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// ImageFsStats indicates an expected call of ImageFsStats.
func (mr *MockProviderMockRecorder) ImageFsStats(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ImageFsStats", reflect.TypeOf((*MockProvider)(nil).ImageFsStats), ctx)
}

// ListBlockVolumesForPod mocks base method.
func (m *MockProvider) ListBlockVolumesForPod(podUID types.UID) (map[string]volume.BlockVolume, bool) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListBlockVolumesForPod", podUID)
	ret0, _ := ret[0].(map[string]volume.BlockVolume)
	ret1, _ := ret[1].(bool)
	return ret0, ret1
}

// ListBlockVolumesForPod indicates an expected call of ListBlockVolumesForPod.
func (mr *MockProviderMockRecorder) ListBlockVolumesForPod(podUID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListBlockVolumesForPod", reflect.TypeOf((*MockProvider)(nil).ListBlockVolumesForPod), podUID)
}

// ListPodCPUAndMemoryStats mocks base method.
func (m *MockProvider) ListPodCPUAndMemoryStats(ctx context.Context) ([]v1alpha1.PodStats, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListPodCPUAndMemoryStats", ctx)
	ret0, _ := ret[0].([]v1alpha1.PodStats)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListPodCPUAndMemoryStats indicates an expected call of ListPodCPUAndMemoryStats.
func (mr *MockProviderMockRecorder) ListPodCPUAndMemoryStats(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListPodCPUAndMemoryStats", reflect.TypeOf((*MockProvider)(nil).ListPodCPUAndMemoryStats), ctx)
}

// ListPodStats mocks base method.
func (m *MockProvider) ListPodStats(ctx context.Context) ([]v1alpha1.PodStats, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListPodStats", ctx)
	ret0, _ := ret[0].([]v1alpha1.PodStats)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListPodStats indicates an expected call of ListPodStats.
func (mr *MockProviderMockRecorder) ListPodStats(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListPodStats", reflect.TypeOf((*MockProvider)(nil).ListPodStats), ctx)
}

// ListPodStatsAndUpdateCPUNanoCoreUsage mocks base method.
func (m *MockProvider) ListPodStatsAndUpdateCPUNanoCoreUsage(ctx context.Context) ([]v1alpha1.PodStats, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListPodStatsAndUpdateCPUNanoCoreUsage", ctx)
	ret0, _ := ret[0].([]v1alpha1.PodStats)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListPodStatsAndUpdateCPUNanoCoreUsage indicates an expected call of ListPodStatsAndUpdateCPUNanoCoreUsage.
func (mr *MockProviderMockRecorder) ListPodStatsAndUpdateCPUNanoCoreUsage(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListPodStatsAndUpdateCPUNanoCoreUsage", reflect.TypeOf((*MockProvider)(nil).ListPodStatsAndUpdateCPUNanoCoreUsage), ctx)
}

// ListVolumesForPod mocks base method.
func (m *MockProvider) ListVolumesForPod(podUID types.UID) (map[string]volume.Volume, bool) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListVolumesForPod", podUID)
	ret0, _ := ret[0].(map[string]volume.Volume)
	ret1, _ := ret[1].(bool)
	return ret0, ret1
}

// ListVolumesForPod indicates an expected call of ListVolumesForPod.
func (mr *MockProviderMockRecorder) ListVolumesForPod(podUID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListVolumesForPod", reflect.TypeOf((*MockProvider)(nil).ListVolumesForPod), podUID)
}

// RlimitStats mocks base method.
func (m *MockProvider) RlimitStats() (*v1alpha1.RlimitStats, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RlimitStats")
	ret0, _ := ret[0].(*v1alpha1.RlimitStats)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// RlimitStats indicates an expected call of RlimitStats.
func (mr *MockProviderMockRecorder) RlimitStats() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RlimitStats", reflect.TypeOf((*MockProvider)(nil).RlimitStats))
}

// RootFsStats mocks base method.
func (m *MockProvider) RootFsStats() (*v1alpha1.FsStats, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RootFsStats")
	ret0, _ := ret[0].(*v1alpha1.FsStats)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// RootFsStats indicates an expected call of RootFsStats.
func (mr *MockProviderMockRecorder) RootFsStats() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RootFsStats", reflect.TypeOf((*MockProvider)(nil).RootFsStats))
}
