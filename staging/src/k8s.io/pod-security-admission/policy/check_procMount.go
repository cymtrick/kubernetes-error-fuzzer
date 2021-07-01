/*
Copyright 2021 The Kubernetes Authors.

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

package policy

import (
	"fmt"

	corev1 "k8s.io/api/core/v1"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/sets"
	"k8s.io/apimachinery/pkg/util/validation/field"
	"k8s.io/pod-security-admission/api"
)

func init() {
	addCheck(CheckProcMount)
}

// CheckProcMount returns a baseline level check that restricts
// setting the value of securityContext.procMount to DefaultProcMount
// in 1.0+
func CheckProcMount() Check {
	return Check{
		ID:    "procMount",
		Level: api.LevelBaseline,
		Versions: []VersionedCheck{
			{
				MinimumVersion: api.MajorMinorVersion(1, 0),
				CheckPod:       checkProcMount_1_0,
			},
		},
	}
}

func checkProcMount_1_0(podMetadata *metav1.ObjectMeta, podSpec *corev1.PodSpec) CheckResult {
	forbiddenContainers := sets.NewString()
	forbiddenProcMountTypes := sets.NewString()
	visitContainersWithPath(podSpec, field.NewPath("spec"), func(container *corev1.Container, path *field.Path) {
		// allow if the security context is nil.
		if container.SecurityContext == nil {
			return
		}
		// allow if proc mount is not set.
		if container.SecurityContext.ProcMount == nil {
			return
		}
		// check if the value of the proc mount type is valid.
		if *container.SecurityContext.ProcMount != v1.DefaultProcMount {
			forbiddenContainers.Insert(container.Name)
			forbiddenProcMountTypes.Insert(string(*container.SecurityContext.ProcMount))
		}
	})
	if len(forbiddenContainers) > 0 {
		return CheckResult{
			Allowed:         false,
			ForbiddenReason: "forbidden procMount",
			ForbiddenDetail: fmt.Sprintf(
				"containers %q have forbidden procMount types %q",
				forbiddenContainers.List(),
				forbiddenProcMountTypes.List(),
			),
		}
	}
	return CheckResult{Allowed: true}
}
