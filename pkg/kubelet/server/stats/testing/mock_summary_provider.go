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
// Source: summary.go

// Package testing is a generated GoMock package.
package testing

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	v1alpha1 "k8s.io/kubelet/pkg/apis/stats/v1alpha1"
)

// MockSummaryProvider is a mock of SummaryProvider interface.
type MockSummaryProvider struct {
	ctrl     *gomock.Controller
	recorder *MockSummaryProviderMockRecorder
}

// MockSummaryProviderMockRecorder is the mock recorder for MockSummaryProvider.
type MockSummaryProviderMockRecorder struct {
	mock *MockSummaryProvider
}

// NewMockSummaryProvider creates a new mock instance.
func NewMockSummaryProvider(ctrl *gomock.Controller) *MockSummaryProvider {
	mock := &MockSummaryProvider{ctrl: ctrl}
	mock.recorder = &MockSummaryProviderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockSummaryProvider) EXPECT() *MockSummaryProviderMockRecorder {
	return m.recorder
}

// Get mocks base method.
func (m *MockSummaryProvider) Get(updateStats bool) (*v1alpha1.Summary, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", updateStats)
	ret0, _ := ret[0].(*v1alpha1.Summary)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get.
func (mr *MockSummaryProviderMockRecorder) Get(updateStats interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockSummaryProvider)(nil).Get), updateStats)
}

// GetCPUAndMemoryStats mocks base method.
func (m *MockSummaryProvider) GetCPUAndMemoryStats() (*v1alpha1.Summary, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCPUAndMemoryStats")
	ret0, _ := ret[0].(*v1alpha1.Summary)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetCPUAndMemoryStats indicates an expected call of GetCPUAndMemoryStats.
func (mr *MockSummaryProviderMockRecorder) GetCPUAndMemoryStats() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCPUAndMemoryStats", reflect.TypeOf((*MockSummaryProvider)(nil).GetCPUAndMemoryStats))
}
