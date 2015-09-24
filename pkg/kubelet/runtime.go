/*
Copyright 2015 The Kubernetes Authors All rights reserved.

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
	"sync"
	"time"
)

type runtimeState struct {
	sync.Mutex
	lastBaseRuntimeSync      time.Time
	baseRuntimeSyncThreshold time.Duration
	networkError             error
	initError                error
}

func (s *runtimeState) setRuntimeSync(t time.Time) {
	s.Lock()
	defer s.Unlock()
	s.lastBaseRuntimeSync = t
}

func (s *runtimeState) setNetworkError(err error) {
	if err == nil {
		return
	}
	s.Lock()
	defer s.Unlock()
	s.networkError = err
}

func (s *runtimeState) setInitError(err error) {
	s.Lock()
	defer s.Unlock()
	s.initError = err
}

func (s *runtimeState) errors() []string {
	s.Lock()
	defer s.Unlock()
	var ret []string
	if s.initError != nil {
		ret = append(ret, s.initError.Error())
	}
	if s.networkError != nil {
		ret = append(ret, s.networkError.Error())
	}
	if !s.lastBaseRuntimeSync.Add(s.baseRuntimeSyncThreshold).After(time.Now()) {
		ret = append(ret, "container runtime is down")
	}
	return ret
}

func newRuntimeState(runtimeSyncThreshold time.Duration) *runtimeState {
	return &runtimeState{
		lastBaseRuntimeSync:      time.Time{},
		baseRuntimeSyncThreshold: runtimeSyncThreshold,
	}
}
