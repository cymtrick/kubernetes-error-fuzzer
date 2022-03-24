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
// Source: types.go

// Package testing is a generated GoMock package.
package testing

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	events "github.com/google/cadvisor/events"
	v1 "github.com/google/cadvisor/info/v1"
	v2 "github.com/google/cadvisor/info/v2"
)

// MockInterface is a mock of Interface interface.
type MockInterface struct {
	ctrl     *gomock.Controller
	recorder *MockInterfaceMockRecorder
}

// MockInterfaceMockRecorder is the mock recorder for MockInterface.
type MockInterfaceMockRecorder struct {
	mock *MockInterface
}

// NewMockInterface creates a new mock instance.
func NewMockInterface(ctrl *gomock.Controller) *MockInterface {
	mock := &MockInterface{ctrl: ctrl}
	mock.recorder = &MockInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockInterface) EXPECT() *MockInterfaceMockRecorder {
	return m.recorder
}

// ContainerInfo mocks base method.
func (m *MockInterface) ContainerInfo(name string, req *v1.ContainerInfoRequest) (*v1.ContainerInfo, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ContainerInfo", name, req)
	ret0, _ := ret[0].(*v1.ContainerInfo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ContainerInfo indicates an expected call of ContainerInfo.
func (mr *MockInterfaceMockRecorder) ContainerInfo(name, req interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ContainerInfo", reflect.TypeOf((*MockInterface)(nil).ContainerInfo), name, req)
}

// ContainerInfoV2 mocks base method.
func (m *MockInterface) ContainerInfoV2(name string, options v2.RequestOptions) (map[string]v2.ContainerInfo, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ContainerInfoV2", name, options)
	ret0, _ := ret[0].(map[string]v2.ContainerInfo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ContainerInfoV2 indicates an expected call of ContainerInfoV2.
func (mr *MockInterfaceMockRecorder) ContainerInfoV2(name, options interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ContainerInfoV2", reflect.TypeOf((*MockInterface)(nil).ContainerInfoV2), name, options)
}

// DockerContainer mocks base method.
func (m *MockInterface) DockerContainer(name string, req *v1.ContainerInfoRequest) (v1.ContainerInfo, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DockerContainer", name, req)
	ret0, _ := ret[0].(v1.ContainerInfo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DockerContainer indicates an expected call of DockerContainer.
func (mr *MockInterfaceMockRecorder) DockerContainer(name, req interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DockerContainer", reflect.TypeOf((*MockInterface)(nil).DockerContainer), name, req)
}

// GetDirFsInfo mocks base method.
func (m *MockInterface) GetDirFsInfo(path string) (v2.FsInfo, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetDirFsInfo", path)
	ret0, _ := ret[0].(v2.FsInfo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetDirFsInfo indicates an expected call of GetDirFsInfo.
func (mr *MockInterfaceMockRecorder) GetDirFsInfo(path interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetDirFsInfo", reflect.TypeOf((*MockInterface)(nil).GetDirFsInfo), path)
}

// GetRequestedContainersInfo mocks base method.
func (m *MockInterface) GetRequestedContainersInfo(containerName string, options v2.RequestOptions) (map[string]*v1.ContainerInfo, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetRequestedContainersInfo", containerName, options)
	ret0, _ := ret[0].(map[string]*v1.ContainerInfo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetRequestedContainersInfo indicates an expected call of GetRequestedContainersInfo.
func (mr *MockInterfaceMockRecorder) GetRequestedContainersInfo(containerName, options interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetRequestedContainersInfo", reflect.TypeOf((*MockInterface)(nil).GetRequestedContainersInfo), containerName, options)
}

// ImagesFsInfo mocks base method.
func (m *MockInterface) ImagesFsInfo() (v2.FsInfo, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ImagesFsInfo")
	ret0, _ := ret[0].(v2.FsInfo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ImagesFsInfo indicates an expected call of ImagesFsInfo.
func (mr *MockInterfaceMockRecorder) ImagesFsInfo() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ImagesFsInfo", reflect.TypeOf((*MockInterface)(nil).ImagesFsInfo))
}

// MachineInfo mocks base method.
func (m *MockInterface) MachineInfo() (*v1.MachineInfo, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "MachineInfo")
	ret0, _ := ret[0].(*v1.MachineInfo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// MachineInfo indicates an expected call of MachineInfo.
func (mr *MockInterfaceMockRecorder) MachineInfo() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MachineInfo", reflect.TypeOf((*MockInterface)(nil).MachineInfo))
}

// RootFsInfo mocks base method.
func (m *MockInterface) RootFsInfo() (v2.FsInfo, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RootFsInfo")
	ret0, _ := ret[0].(v2.FsInfo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// RootFsInfo indicates an expected call of RootFsInfo.
func (mr *MockInterfaceMockRecorder) RootFsInfo() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RootFsInfo", reflect.TypeOf((*MockInterface)(nil).RootFsInfo))
}

// Start mocks base method.
func (m *MockInterface) Start() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Start")
	ret0, _ := ret[0].(error)
	return ret0
}

// Start indicates an expected call of Start.
func (mr *MockInterfaceMockRecorder) Start() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Start", reflect.TypeOf((*MockInterface)(nil).Start))
}

// SubcontainerInfo mocks base method.
func (m *MockInterface) SubcontainerInfo(name string, req *v1.ContainerInfoRequest) (map[string]*v1.ContainerInfo, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SubcontainerInfo", name, req)
	ret0, _ := ret[0].(map[string]*v1.ContainerInfo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SubcontainerInfo indicates an expected call of SubcontainerInfo.
func (mr *MockInterfaceMockRecorder) SubcontainerInfo(name, req interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SubcontainerInfo", reflect.TypeOf((*MockInterface)(nil).SubcontainerInfo), name, req)
}

// VersionInfo mocks base method.
func (m *MockInterface) VersionInfo() (*v1.VersionInfo, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "VersionInfo")
	ret0, _ := ret[0].(*v1.VersionInfo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// VersionInfo indicates an expected call of VersionInfo.
func (mr *MockInterfaceMockRecorder) VersionInfo() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "VersionInfo", reflect.TypeOf((*MockInterface)(nil).VersionInfo))
}

// WatchEvents mocks base method.
func (m *MockInterface) WatchEvents(request *events.Request) (*events.EventChannel, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "WatchEvents", request)
	ret0, _ := ret[0].(*events.EventChannel)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// WatchEvents indicates an expected call of WatchEvents.
func (mr *MockInterfaceMockRecorder) WatchEvents(request interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WatchEvents", reflect.TypeOf((*MockInterface)(nil).WatchEvents), request)
}

// MockImageFsInfoProvider is a mock of ImageFsInfoProvider interface.
type MockImageFsInfoProvider struct {
	ctrl     *gomock.Controller
	recorder *MockImageFsInfoProviderMockRecorder
}

// MockImageFsInfoProviderMockRecorder is the mock recorder for MockImageFsInfoProvider.
type MockImageFsInfoProviderMockRecorder struct {
	mock *MockImageFsInfoProvider
}

// NewMockImageFsInfoProvider creates a new mock instance.
func NewMockImageFsInfoProvider(ctrl *gomock.Controller) *MockImageFsInfoProvider {
	mock := &MockImageFsInfoProvider{ctrl: ctrl}
	mock.recorder = &MockImageFsInfoProviderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockImageFsInfoProvider) EXPECT() *MockImageFsInfoProviderMockRecorder {
	return m.recorder
}

// ImageFsInfoLabel mocks base method.
func (m *MockImageFsInfoProvider) ImageFsInfoLabel() (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ImageFsInfoLabel")
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ImageFsInfoLabel indicates an expected call of ImageFsInfoLabel.
func (mr *MockImageFsInfoProviderMockRecorder) ImageFsInfoLabel() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ImageFsInfoLabel", reflect.TypeOf((*MockImageFsInfoProvider)(nil).ImageFsInfoLabel))
}
