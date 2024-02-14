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
// Source: runtime.go

// Package testing is a generated GoMock package.
package testing

import (
	context "context"
	io "io"
	url "net/url"
	reflect "reflect"
	time "time"

	gomock "github.com/golang/mock/gomock"
	v1 "k8s.io/api/core/v1"
	types "k8s.io/apimachinery/pkg/types"
	remotecommand "k8s.io/client-go/tools/remotecommand"
	flowcontrol "k8s.io/client-go/util/flowcontrol"
	v10 "k8s.io/cri-api/pkg/apis/runtime/v1"
	container "k8s.io/kubernetes/pkg/kubelet/container"
)

// MockVersion is a mock of Version interface.
type MockVersion struct {
	ctrl     *gomock.Controller
	recorder *MockVersionMockRecorder
}

// MockVersionMockRecorder is the mock recorder for MockVersion.
type MockVersionMockRecorder struct {
	mock *MockVersion
}

// NewMockVersion creates a new mock instance.
func NewMockVersion(ctrl *gomock.Controller) *MockVersion {
	mock := &MockVersion{ctrl: ctrl}
	mock.recorder = &MockVersionMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockVersion) EXPECT() *MockVersionMockRecorder {
	return m.recorder
}

// Compare mocks base method.
func (m *MockVersion) Compare(other string) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Compare", other)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Compare indicates an expected call of Compare.
func (mr *MockVersionMockRecorder) Compare(other interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Compare", reflect.TypeOf((*MockVersion)(nil).Compare), other)
}

// String mocks base method.
func (m *MockVersion) String() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "String")
	ret0, _ := ret[0].(string)
	return ret0
}

// String indicates an expected call of String.
func (mr *MockVersionMockRecorder) String() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "String", reflect.TypeOf((*MockVersion)(nil).String))
}

// MockRuntime is a mock of Runtime interface.
type MockRuntime struct {
	ctrl     *gomock.Controller
	recorder *MockRuntimeMockRecorder
}

// MockRuntimeMockRecorder is the mock recorder for MockRuntime.
type MockRuntimeMockRecorder struct {
	mock *MockRuntime
}

// NewMockRuntime creates a new mock instance.
func NewMockRuntime(ctrl *gomock.Controller) *MockRuntime {
	mock := &MockRuntime{ctrl: ctrl}
	mock.recorder = &MockRuntimeMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRuntime) EXPECT() *MockRuntimeMockRecorder {
	return m.recorder
}

// APIVersion mocks base method.
func (m *MockRuntime) APIVersion() (container.Version, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "APIVersion")
	ret0, _ := ret[0].(container.Version)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// APIVersion indicates an expected call of APIVersion.
func (mr *MockRuntimeMockRecorder) APIVersion() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "APIVersion", reflect.TypeOf((*MockRuntime)(nil).APIVersion))
}

// CheckpointContainer mocks base method.
func (m *MockRuntime) CheckpointContainer(ctx context.Context, options *v10.CheckpointContainerRequest) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CheckpointContainer", ctx, options)
	ret0, _ := ret[0].(error)
	return ret0
}

// CheckpointContainer indicates an expected call of CheckpointContainer.
func (mr *MockRuntimeMockRecorder) CheckpointContainer(ctx, options interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CheckpointContainer", reflect.TypeOf((*MockRuntime)(nil).CheckpointContainer), ctx, options)
}

// DeleteContainer mocks base method.
func (m *MockRuntime) DeleteContainer(ctx context.Context, containerID container.ContainerID) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteContainer", ctx, containerID)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteContainer indicates an expected call of DeleteContainer.
func (mr *MockRuntimeMockRecorder) DeleteContainer(ctx, containerID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteContainer", reflect.TypeOf((*MockRuntime)(nil).DeleteContainer), ctx, containerID)
}

// GarbageCollect mocks base method.
func (m *MockRuntime) GarbageCollect(ctx context.Context, gcPolicy container.GCPolicy, allSourcesReady, evictNonDeletedPods bool) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GarbageCollect", ctx, gcPolicy, allSourcesReady, evictNonDeletedPods)
	ret0, _ := ret[0].(error)
	return ret0
}

// GarbageCollect indicates an expected call of GarbageCollect.
func (mr *MockRuntimeMockRecorder) GarbageCollect(ctx, gcPolicy, allSourcesReady, evictNonDeletedPods interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GarbageCollect", reflect.TypeOf((*MockRuntime)(nil).GarbageCollect), ctx, gcPolicy, allSourcesReady, evictNonDeletedPods)
}

// GeneratePodStatus mocks base method.
func (m *MockRuntime) GeneratePodStatus(event *v10.ContainerEventResponse) (*container.PodStatus, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GeneratePodStatus", event)
	ret0, _ := ret[0].(*container.PodStatus)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GeneratePodStatus indicates an expected call of GeneratePodStatus.
func (mr *MockRuntimeMockRecorder) GeneratePodStatus(event interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GeneratePodStatus", reflect.TypeOf((*MockRuntime)(nil).GeneratePodStatus), event)
}

// GetContainerLogs mocks base method.
func (m *MockRuntime) GetContainerLogs(ctx context.Context, pod *v1.Pod, containerID container.ContainerID, logOptions *v1.PodLogOptions, stdout, stderr io.Writer) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetContainerLogs", ctx, pod, containerID, logOptions, stdout, stderr)
	ret0, _ := ret[0].(error)
	return ret0
}

// GetContainerLogs indicates an expected call of GetContainerLogs.
func (mr *MockRuntimeMockRecorder) GetContainerLogs(ctx, pod, containerID, logOptions, stdout, stderr interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetContainerLogs", reflect.TypeOf((*MockRuntime)(nil).GetContainerLogs), ctx, pod, containerID, logOptions, stdout, stderr)
}

// GetImageRef mocks base method.
func (m *MockRuntime) GetImageRef(ctx context.Context, image container.ImageSpec) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetImageRef", ctx, image)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetImageRef indicates an expected call of GetImageRef.
func (mr *MockRuntimeMockRecorder) GetImageRef(ctx, image interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetImageRef", reflect.TypeOf((*MockRuntime)(nil).GetImageRef), ctx, image)
}

// GetImageSize mocks base method.
func (m *MockRuntime) GetImageSize(ctx context.Context, image container.ImageSpec) (uint64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetImageSize", ctx, image)
	ret0, _ := ret[0].(uint64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetImageSize indicates an expected call of GetImageSize.
func (mr *MockRuntimeMockRecorder) GetImageSize(ctx, image interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetImageSize", reflect.TypeOf((*MockRuntime)(nil).GetImageSize), ctx, image)
}

// GetPodStatus mocks base method.
func (m *MockRuntime) GetPodStatus(ctx context.Context, uid types.UID, name, namespace string) (*container.PodStatus, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPodStatus", ctx, uid, name, namespace)
	ret0, _ := ret[0].(*container.PodStatus)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetPodStatus indicates an expected call of GetPodStatus.
func (mr *MockRuntimeMockRecorder) GetPodStatus(ctx, uid, name, namespace interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPodStatus", reflect.TypeOf((*MockRuntime)(nil).GetPodStatus), ctx, uid, name, namespace)
}

// GetPods mocks base method.
func (m *MockRuntime) GetPods(ctx context.Context, all bool) ([]*container.Pod, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPods", ctx, all)
	ret0, _ := ret[0].([]*container.Pod)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetPods indicates an expected call of GetPods.
func (mr *MockRuntimeMockRecorder) GetPods(ctx, all interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPods", reflect.TypeOf((*MockRuntime)(nil).GetPods), ctx, all)
}

// ImageFsInfo mocks base method.
func (m *MockRuntime) ImageFsInfo(ctx context.Context) (*v10.ImageFsInfoResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ImageFsInfo", ctx)
	ret0, _ := ret[0].(*v10.ImageFsInfoResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ImageFsInfo indicates an expected call of ImageFsInfo.
func (mr *MockRuntimeMockRecorder) ImageFsInfo(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ImageFsInfo", reflect.TypeOf((*MockRuntime)(nil).ImageFsInfo), ctx)
}

// ImageStats mocks base method.
func (m *MockRuntime) ImageStats(ctx context.Context) (*container.ImageStats, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ImageStats", ctx)
	ret0, _ := ret[0].(*container.ImageStats)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ImageStats indicates an expected call of ImageStats.
func (mr *MockRuntimeMockRecorder) ImageStats(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ImageStats", reflect.TypeOf((*MockRuntime)(nil).ImageStats), ctx)
}

// KillPod mocks base method.
func (m *MockRuntime) KillPod(ctx context.Context, pod *v1.Pod, runningPod container.Pod, gracePeriodOverride *int64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "KillPod", ctx, pod, runningPod, gracePeriodOverride)
	ret0, _ := ret[0].(error)
	return ret0
}

// KillPod indicates an expected call of KillPod.
func (mr *MockRuntimeMockRecorder) KillPod(ctx, pod, runningPod, gracePeriodOverride interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "KillPod", reflect.TypeOf((*MockRuntime)(nil).KillPod), ctx, pod, runningPod, gracePeriodOverride)
}

// ListImages mocks base method.
func (m *MockRuntime) ListImages(ctx context.Context) ([]container.Image, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListImages", ctx)
	ret0, _ := ret[0].([]container.Image)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListImages indicates an expected call of ListImages.
func (mr *MockRuntimeMockRecorder) ListImages(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListImages", reflect.TypeOf((*MockRuntime)(nil).ListImages), ctx)
}

// ListMetricDescriptors mocks base method.
func (m *MockRuntime) ListMetricDescriptors(ctx context.Context) ([]*v10.MetricDescriptor, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListMetricDescriptors", ctx)
	ret0, _ := ret[0].([]*v10.MetricDescriptor)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListMetricDescriptors indicates an expected call of ListMetricDescriptors.
func (mr *MockRuntimeMockRecorder) ListMetricDescriptors(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListMetricDescriptors", reflect.TypeOf((*MockRuntime)(nil).ListMetricDescriptors), ctx)
}

// ListPodSandboxMetrics mocks base method.
func (m *MockRuntime) ListPodSandboxMetrics(ctx context.Context) ([]*v10.PodSandboxMetrics, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListPodSandboxMetrics", ctx)
	ret0, _ := ret[0].([]*v10.PodSandboxMetrics)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListPodSandboxMetrics indicates an expected call of ListPodSandboxMetrics.
func (mr *MockRuntimeMockRecorder) ListPodSandboxMetrics(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListPodSandboxMetrics", reflect.TypeOf((*MockRuntime)(nil).ListPodSandboxMetrics), ctx)
}

// PullImage mocks base method.
func (m *MockRuntime) PullImage(ctx context.Context, image container.ImageSpec, pullSecrets []v1.Secret, podSandboxConfig *v10.PodSandboxConfig) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PullImage", ctx, image, pullSecrets, podSandboxConfig)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// PullImage indicates an expected call of PullImage.
func (mr *MockRuntimeMockRecorder) PullImage(ctx, image, pullSecrets, podSandboxConfig interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PullImage", reflect.TypeOf((*MockRuntime)(nil).PullImage), ctx, image, pullSecrets, podSandboxConfig)
}

// RemoveImage mocks base method.
func (m *MockRuntime) RemoveImage(ctx context.Context, image container.ImageSpec) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RemoveImage", ctx, image)
	ret0, _ := ret[0].(error)
	return ret0
}

// RemoveImage indicates an expected call of RemoveImage.
func (mr *MockRuntimeMockRecorder) RemoveImage(ctx, image interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoveImage", reflect.TypeOf((*MockRuntime)(nil).RemoveImage), ctx, image)
}

// Status mocks base method.
func (m *MockRuntime) Status(ctx context.Context) (*container.RuntimeStatus, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Status", ctx)
	ret0, _ := ret[0].(*container.RuntimeStatus)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Status indicates an expected call of Status.
func (mr *MockRuntimeMockRecorder) Status(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Status", reflect.TypeOf((*MockRuntime)(nil).Status), ctx)
}

// SyncPod mocks base method.
func (m *MockRuntime) SyncPod(ctx context.Context, pod *v1.Pod, podStatus *container.PodStatus, pullSecrets []v1.Secret, backOff *flowcontrol.Backoff) container.PodSyncResult {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SyncPod", ctx, pod, podStatus, pullSecrets, backOff)
	ret0, _ := ret[0].(container.PodSyncResult)
	return ret0
}

// SyncPod indicates an expected call of SyncPod.
func (mr *MockRuntimeMockRecorder) SyncPod(ctx, pod, podStatus, pullSecrets, backOff interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SyncPod", reflect.TypeOf((*MockRuntime)(nil).SyncPod), ctx, pod, podStatus, pullSecrets, backOff)
}

// Type mocks base method.
func (m *MockRuntime) Type() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Type")
	ret0, _ := ret[0].(string)
	return ret0
}

// Type indicates an expected call of Type.
func (mr *MockRuntimeMockRecorder) Type() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Type", reflect.TypeOf((*MockRuntime)(nil).Type))
}

// UpdatePodCIDR mocks base method.
func (m *MockRuntime) UpdatePodCIDR(ctx context.Context, podCIDR string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdatePodCIDR", ctx, podCIDR)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdatePodCIDR indicates an expected call of UpdatePodCIDR.
func (mr *MockRuntimeMockRecorder) UpdatePodCIDR(ctx, podCIDR interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdatePodCIDR", reflect.TypeOf((*MockRuntime)(nil).UpdatePodCIDR), ctx, podCIDR)
}

// Version mocks base method.
func (m *MockRuntime) Version(ctx context.Context) (container.Version, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Version", ctx)
	ret0, _ := ret[0].(container.Version)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Version indicates an expected call of Version.
func (mr *MockRuntimeMockRecorder) Version(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Version", reflect.TypeOf((*MockRuntime)(nil).Version), ctx)
}

// MockStreamingRuntime is a mock of StreamingRuntime interface.
type MockStreamingRuntime struct {
	ctrl     *gomock.Controller
	recorder *MockStreamingRuntimeMockRecorder
}

// MockStreamingRuntimeMockRecorder is the mock recorder for MockStreamingRuntime.
type MockStreamingRuntimeMockRecorder struct {
	mock *MockStreamingRuntime
}

// NewMockStreamingRuntime creates a new mock instance.
func NewMockStreamingRuntime(ctrl *gomock.Controller) *MockStreamingRuntime {
	mock := &MockStreamingRuntime{ctrl: ctrl}
	mock.recorder = &MockStreamingRuntimeMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockStreamingRuntime) EXPECT() *MockStreamingRuntimeMockRecorder {
	return m.recorder
}

// GetAttach mocks base method.
func (m *MockStreamingRuntime) GetAttach(ctx context.Context, id container.ContainerID, stdin, stdout, stderr, tty bool) (*url.URL, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAttach", ctx, id, stdin, stdout, stderr, tty)
	ret0, _ := ret[0].(*url.URL)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAttach indicates an expected call of GetAttach.
func (mr *MockStreamingRuntimeMockRecorder) GetAttach(ctx, id, stdin, stdout, stderr, tty interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAttach", reflect.TypeOf((*MockStreamingRuntime)(nil).GetAttach), ctx, id, stdin, stdout, stderr, tty)
}

// GetExec mocks base method.
func (m *MockStreamingRuntime) GetExec(ctx context.Context, id container.ContainerID, cmd []string, stdin, stdout, stderr, tty bool) (*url.URL, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetExec", ctx, id, cmd, stdin, stdout, stderr, tty)
	ret0, _ := ret[0].(*url.URL)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetExec indicates an expected call of GetExec.
func (mr *MockStreamingRuntimeMockRecorder) GetExec(ctx, id, cmd, stdin, stdout, stderr, tty interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetExec", reflect.TypeOf((*MockStreamingRuntime)(nil).GetExec), ctx, id, cmd, stdin, stdout, stderr, tty)
}

// GetPortForward mocks base method.
func (m *MockStreamingRuntime) GetPortForward(ctx context.Context, podName, podNamespace string, podUID types.UID, ports []int32) (*url.URL, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPortForward", ctx, podName, podNamespace, podUID, ports)
	ret0, _ := ret[0].(*url.URL)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetPortForward indicates an expected call of GetPortForward.
func (mr *MockStreamingRuntimeMockRecorder) GetPortForward(ctx, podName, podNamespace, podUID, ports interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPortForward", reflect.TypeOf((*MockStreamingRuntime)(nil).GetPortForward), ctx, podName, podNamespace, podUID, ports)
}

// MockImageService is a mock of ImageService interface.
type MockImageService struct {
	ctrl     *gomock.Controller
	recorder *MockImageServiceMockRecorder
}

// MockImageServiceMockRecorder is the mock recorder for MockImageService.
type MockImageServiceMockRecorder struct {
	mock *MockImageService
}

// NewMockImageService creates a new mock instance.
func NewMockImageService(ctrl *gomock.Controller) *MockImageService {
	mock := &MockImageService{ctrl: ctrl}
	mock.recorder = &MockImageServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockImageService) EXPECT() *MockImageServiceMockRecorder {
	return m.recorder
}

// GetImageRef mocks base method.
func (m *MockImageService) GetImageRef(ctx context.Context, image container.ImageSpec) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetImageRef", ctx, image)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetImageRef indicates an expected call of GetImageRef.
func (mr *MockImageServiceMockRecorder) GetImageRef(ctx, image interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetImageRef", reflect.TypeOf((*MockImageService)(nil).GetImageRef), ctx, image)
}

// GetImageSize mocks base method.
func (m *MockImageService) GetImageSize(ctx context.Context, image container.ImageSpec) (uint64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetImageSize", ctx, image)
	ret0, _ := ret[0].(uint64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetImageSize indicates an expected call of GetImageSize.
func (mr *MockImageServiceMockRecorder) GetImageSize(ctx, image interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetImageSize", reflect.TypeOf((*MockImageService)(nil).GetImageSize), ctx, image)
}

// ImageFsInfo mocks base method.
func (m *MockImageService) ImageFsInfo(ctx context.Context) (*v10.ImageFsInfoResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ImageFsInfo", ctx)
	ret0, _ := ret[0].(*v10.ImageFsInfoResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ImageFsInfo indicates an expected call of ImageFsInfo.
func (mr *MockImageServiceMockRecorder) ImageFsInfo(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ImageFsInfo", reflect.TypeOf((*MockImageService)(nil).ImageFsInfo), ctx)
}

// ImageStats mocks base method.
func (m *MockImageService) ImageStats(ctx context.Context) (*container.ImageStats, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ImageStats", ctx)
	ret0, _ := ret[0].(*container.ImageStats)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ImageStats indicates an expected call of ImageStats.
func (mr *MockImageServiceMockRecorder) ImageStats(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ImageStats", reflect.TypeOf((*MockImageService)(nil).ImageStats), ctx)
}

// ListImages mocks base method.
func (m *MockImageService) ListImages(ctx context.Context) ([]container.Image, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListImages", ctx)
	ret0, _ := ret[0].([]container.Image)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListImages indicates an expected call of ListImages.
func (mr *MockImageServiceMockRecorder) ListImages(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListImages", reflect.TypeOf((*MockImageService)(nil).ListImages), ctx)
}

// PullImage mocks base method.
func (m *MockImageService) PullImage(ctx context.Context, image container.ImageSpec, pullSecrets []v1.Secret, podSandboxConfig *v10.PodSandboxConfig) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PullImage", ctx, image, pullSecrets, podSandboxConfig)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// PullImage indicates an expected call of PullImage.
func (mr *MockImageServiceMockRecorder) PullImage(ctx, image, pullSecrets, podSandboxConfig interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PullImage", reflect.TypeOf((*MockImageService)(nil).PullImage), ctx, image, pullSecrets, podSandboxConfig)
}

// RemoveImage mocks base method.
func (m *MockImageService) RemoveImage(ctx context.Context, image container.ImageSpec) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RemoveImage", ctx, image)
	ret0, _ := ret[0].(error)
	return ret0
}

// RemoveImage indicates an expected call of RemoveImage.
func (mr *MockImageServiceMockRecorder) RemoveImage(ctx, image interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoveImage", reflect.TypeOf((*MockImageService)(nil).RemoveImage), ctx, image)
}

// MockAttacher is a mock of Attacher interface.
type MockAttacher struct {
	ctrl     *gomock.Controller
	recorder *MockAttacherMockRecorder
}

// MockAttacherMockRecorder is the mock recorder for MockAttacher.
type MockAttacherMockRecorder struct {
	mock *MockAttacher
}

// NewMockAttacher creates a new mock instance.
func NewMockAttacher(ctrl *gomock.Controller) *MockAttacher {
	mock := &MockAttacher{ctrl: ctrl}
	mock.recorder = &MockAttacherMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAttacher) EXPECT() *MockAttacherMockRecorder {
	return m.recorder
}

// AttachContainer mocks base method.
func (m *MockAttacher) AttachContainer(ctx context.Context, id container.ContainerID, stdin io.Reader, stdout, stderr io.WriteCloser, tty bool, resize <-chan remotecommand.TerminalSize) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AttachContainer", ctx, id, stdin, stdout, stderr, tty, resize)
	ret0, _ := ret[0].(error)
	return ret0
}

// AttachContainer indicates an expected call of AttachContainer.
func (mr *MockAttacherMockRecorder) AttachContainer(ctx, id, stdin, stdout, stderr, tty, resize interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AttachContainer", reflect.TypeOf((*MockAttacher)(nil).AttachContainer), ctx, id, stdin, stdout, stderr, tty, resize)
}

// MockCommandRunner is a mock of CommandRunner interface.
type MockCommandRunner struct {
	ctrl     *gomock.Controller
	recorder *MockCommandRunnerMockRecorder
}

// MockCommandRunnerMockRecorder is the mock recorder for MockCommandRunner.
type MockCommandRunnerMockRecorder struct {
	mock *MockCommandRunner
}

// NewMockCommandRunner creates a new mock instance.
func NewMockCommandRunner(ctrl *gomock.Controller) *MockCommandRunner {
	mock := &MockCommandRunner{ctrl: ctrl}
	mock.recorder = &MockCommandRunnerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockCommandRunner) EXPECT() *MockCommandRunnerMockRecorder {
	return m.recorder
}

// RunInContainer mocks base method.
func (m *MockCommandRunner) RunInContainer(ctx context.Context, id container.ContainerID, cmd []string, timeout time.Duration) ([]byte, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RunInContainer", ctx, id, cmd, timeout)
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// RunInContainer indicates an expected call of RunInContainer.
func (mr *MockCommandRunnerMockRecorder) RunInContainer(ctx, id, cmd, timeout interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RunInContainer", reflect.TypeOf((*MockCommandRunner)(nil).RunInContainer), ctx, id, cmd, timeout)
}
