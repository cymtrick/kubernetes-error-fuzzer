//go:build !dockerless
// +build !dockerless

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
// Source: plugins.go

// Package testing is a generated GoMock package.
package testing

import (
	gomock "github.com/golang/mock/gomock"
	sets "k8s.io/apimachinery/pkg/util/sets"
	config "k8s.io/kubernetes/pkg/kubelet/apis/config"
	container "k8s.io/kubernetes/pkg/kubelet/container"
	network "k8s.io/kubernetes/pkg/kubelet/dockershim/network"
	hostport "k8s.io/kubernetes/pkg/kubelet/dockershim/network/hostport"
	reflect "reflect"
)

// MockNetworkPlugin is a mock of NetworkPlugin interface
type MockNetworkPlugin struct {
	ctrl     *gomock.Controller
	recorder *MockNetworkPluginMockRecorder
}

// MockNetworkPluginMockRecorder is the mock recorder for MockNetworkPlugin
type MockNetworkPluginMockRecorder struct {
	mock *MockNetworkPlugin
}

// NewMockNetworkPlugin creates a new mock instance
func NewMockNetworkPlugin(ctrl *gomock.Controller) *MockNetworkPlugin {
	mock := &MockNetworkPlugin{ctrl: ctrl}
	mock.recorder = &MockNetworkPluginMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockNetworkPlugin) EXPECT() *MockNetworkPluginMockRecorder {
	return m.recorder
}

// Init mocks base method
func (m *MockNetworkPlugin) Init(host network.Host, hairpinMode config.HairpinMode, nonMasqueradeCIDR string, mtu int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Init", host, hairpinMode, nonMasqueradeCIDR, mtu)
	ret0, _ := ret[0].(error)
	return ret0
}

// Init indicates an expected call of Init
func (mr *MockNetworkPluginMockRecorder) Init(host, hairpinMode, nonMasqueradeCIDR, mtu interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Init", reflect.TypeOf((*MockNetworkPlugin)(nil).Init), host, hairpinMode, nonMasqueradeCIDR, mtu)
}

// Event mocks base method
func (m *MockNetworkPlugin) Event(name string, details map[string]interface{}) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Event", name, details)
}

// Event indicates an expected call of Event
func (mr *MockNetworkPluginMockRecorder) Event(name, details interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Event", reflect.TypeOf((*MockNetworkPlugin)(nil).Event), name, details)
}

// Name mocks base method
func (m *MockNetworkPlugin) Name() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Name")
	ret0, _ := ret[0].(string)
	return ret0
}

// Name indicates an expected call of Name
func (mr *MockNetworkPluginMockRecorder) Name() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Name", reflect.TypeOf((*MockNetworkPlugin)(nil).Name))
}

// Capabilities mocks base method
func (m *MockNetworkPlugin) Capabilities() sets.Int {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Capabilities")
	ret0, _ := ret[0].(sets.Int)
	return ret0
}

// Capabilities indicates an expected call of Capabilities
func (mr *MockNetworkPluginMockRecorder) Capabilities() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Capabilities", reflect.TypeOf((*MockNetworkPlugin)(nil).Capabilities))
}

// SetUpPod mocks base method
func (m *MockNetworkPlugin) SetUpPod(namespace, name string, podSandboxID container.ContainerID, annotations, options map[string]string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetUpPod", namespace, name, podSandboxID, annotations, options)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetUpPod indicates an expected call of SetUpPod
func (mr *MockNetworkPluginMockRecorder) SetUpPod(namespace, name, podSandboxID, annotations, options interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetUpPod", reflect.TypeOf((*MockNetworkPlugin)(nil).SetUpPod), namespace, name, podSandboxID, annotations, options)
}

// TearDownPod mocks base method
func (m *MockNetworkPlugin) TearDownPod(namespace, name string, podSandboxID container.ContainerID) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "TearDownPod", namespace, name, podSandboxID)
	ret0, _ := ret[0].(error)
	return ret0
}

// TearDownPod indicates an expected call of TearDownPod
func (mr *MockNetworkPluginMockRecorder) TearDownPod(namespace, name, podSandboxID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TearDownPod", reflect.TypeOf((*MockNetworkPlugin)(nil).TearDownPod), namespace, name, podSandboxID)
}

// GetPodNetworkStatus mocks base method
func (m *MockNetworkPlugin) GetPodNetworkStatus(namespace, name string, podSandboxID container.ContainerID) (*network.PodNetworkStatus, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPodNetworkStatus", namespace, name, podSandboxID)
	ret0, _ := ret[0].(*network.PodNetworkStatus)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetPodNetworkStatus indicates an expected call of GetPodNetworkStatus
func (mr *MockNetworkPluginMockRecorder) GetPodNetworkStatus(namespace, name, podSandboxID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPodNetworkStatus", reflect.TypeOf((*MockNetworkPlugin)(nil).GetPodNetworkStatus), namespace, name, podSandboxID)
}

// Status mocks base method
func (m *MockNetworkPlugin) Status() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Status")
	ret0, _ := ret[0].(error)
	return ret0
}

// Status indicates an expected call of Status
func (mr *MockNetworkPluginMockRecorder) Status() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Status", reflect.TypeOf((*MockNetworkPlugin)(nil).Status))
}

// MockHost is a mock of Host interface
type MockHost struct {
	ctrl     *gomock.Controller
	recorder *MockHostMockRecorder
}

// MockHostMockRecorder is the mock recorder for MockHost
type MockHostMockRecorder struct {
	mock *MockHost
}

// NewMockHost creates a new mock instance
func NewMockHost(ctrl *gomock.Controller) *MockHost {
	mock := &MockHost{ctrl: ctrl}
	mock.recorder = &MockHostMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockHost) EXPECT() *MockHostMockRecorder {
	return m.recorder
}

// GetNetNS mocks base method
func (m *MockHost) GetNetNS(containerID string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetNetNS", containerID)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetNetNS indicates an expected call of GetNetNS
func (mr *MockHostMockRecorder) GetNetNS(containerID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetNetNS", reflect.TypeOf((*MockHost)(nil).GetNetNS), containerID)
}

// GetPodPortMappings mocks base method
func (m *MockHost) GetPodPortMappings(containerID string) ([]*hostport.PortMapping, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPodPortMappings", containerID)
	ret0, _ := ret[0].([]*hostport.PortMapping)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetPodPortMappings indicates an expected call of GetPodPortMappings
func (mr *MockHostMockRecorder) GetPodPortMappings(containerID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPodPortMappings", reflect.TypeOf((*MockHost)(nil).GetPodPortMappings), containerID)
}

// MockNamespaceGetter is a mock of NamespaceGetter interface
type MockNamespaceGetter struct {
	ctrl     *gomock.Controller
	recorder *MockNamespaceGetterMockRecorder
}

// MockNamespaceGetterMockRecorder is the mock recorder for MockNamespaceGetter
type MockNamespaceGetterMockRecorder struct {
	mock *MockNamespaceGetter
}

// NewMockNamespaceGetter creates a new mock instance
func NewMockNamespaceGetter(ctrl *gomock.Controller) *MockNamespaceGetter {
	mock := &MockNamespaceGetter{ctrl: ctrl}
	mock.recorder = &MockNamespaceGetterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockNamespaceGetter) EXPECT() *MockNamespaceGetterMockRecorder {
	return m.recorder
}

// GetNetNS mocks base method
func (m *MockNamespaceGetter) GetNetNS(containerID string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetNetNS", containerID)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetNetNS indicates an expected call of GetNetNS
func (mr *MockNamespaceGetterMockRecorder) GetNetNS(containerID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetNetNS", reflect.TypeOf((*MockNamespaceGetter)(nil).GetNetNS), containerID)
}

// MockPortMappingGetter is a mock of PortMappingGetter interface
type MockPortMappingGetter struct {
	ctrl     *gomock.Controller
	recorder *MockPortMappingGetterMockRecorder
}

// MockPortMappingGetterMockRecorder is the mock recorder for MockPortMappingGetter
type MockPortMappingGetterMockRecorder struct {
	mock *MockPortMappingGetter
}

// NewMockPortMappingGetter creates a new mock instance
func NewMockPortMappingGetter(ctrl *gomock.Controller) *MockPortMappingGetter {
	mock := &MockPortMappingGetter{ctrl: ctrl}
	mock.recorder = &MockPortMappingGetterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockPortMappingGetter) EXPECT() *MockPortMappingGetterMockRecorder {
	return m.recorder
}

// GetPodPortMappings mocks base method
func (m *MockPortMappingGetter) GetPodPortMappings(containerID string) ([]*hostport.PortMapping, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPodPortMappings", containerID)
	ret0, _ := ret[0].([]*hostport.PortMapping)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetPodPortMappings indicates an expected call of GetPodPortMappings
func (mr *MockPortMappingGetterMockRecorder) GetPodPortMappings(containerID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPodPortMappings", reflect.TypeOf((*MockPortMappingGetter)(nil).GetPodPortMappings), containerID)
}
