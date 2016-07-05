// +build linux

/*
Copyright 2015 The Kubernetes Authors.

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

package util

import (
	"os"
	"syscall"

	"github.com/opencontainers/runc/libcontainer/cgroups/fs"
	"github.com/opencontainers/runc/libcontainer/configs"
)

// Creates resource-only containerName if it does not already exist and moves
// the current process to it.
//
// containerName must be an absolute container name.
func RunInResourceContainer(containerName string) error {
	manager := fs.Manager{
		Cgroups: &configs.Cgroup{
			Parent: "/",
			Name:   containerName,
			Resources: &configs.Resources{
				AllowAllDevices: true,
			},
		},
	}

	return manager.Apply(os.Getpid())
}

func ApplyRLimitForSelf(maxOpenFiles uint64) {
	syscall.Setrlimit(syscall.RLIMIT_NOFILE, &syscall.Rlimit{Max: maxOpenFiles, Cur: maxOpenFiles})
}
