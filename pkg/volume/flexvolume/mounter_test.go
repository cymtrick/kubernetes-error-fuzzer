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

package flexvolume

import (
	"testing"

	"k8s.io/apimachinery/pkg/types"
	"k8s.io/kubernetes/pkg/api/v1"
	"k8s.io/kubernetes/pkg/util/mount"
)

func TestSetUpAt(t *testing.T) {
	spec := fakeVolumeSpec()
	pod := &v1.Pod{}
	mounter := &mount.FakeMounter{}

	plugin, rootDir := testPlugin()
	plugin.unsupportedCommands = []string{"unsupportedCmd"}
	plugin.runner = fakeRunner(
		// first call without fsGroup
		assertDriverCall(t, successOutput(), mountCmd, rootDir+"/mount-dir",
			specJson(plugin, spec, nil)),

		// second test has fsGroup
		assertDriverCall(t, notSupportedOutput(), mountCmd, rootDir+"/mount-dir",
			specJson(plugin, spec, map[string]string{
				optionFSGroup: "42",
			})),
		assertDriverCall(t, fakeVolumeNameOutput("sdx"), getVolumeNameCmd,
			specJson(plugin, spec, nil)),
	)

	m, _ := plugin.newMounterInternal(spec, pod, mounter, plugin.runner)
	m.SetUpAt(rootDir+"/mount-dir", nil)

	fsGroup := types.UnixGroupID(42)
	m.SetUpAt(rootDir+"/mount-dir", &fsGroup)
}
