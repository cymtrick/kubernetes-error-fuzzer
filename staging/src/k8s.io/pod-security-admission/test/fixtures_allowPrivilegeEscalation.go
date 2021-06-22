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

package test

import (
	corev1 "k8s.io/api/core/v1"
	"k8s.io/pod-security-admission/api"
	"k8s.io/utils/pointer"
)

/*
TODO: include field paths in reflect-based unit test

containerFields: []string{
	`securityContext.allowPrivilegeEscalation`,
},

*/

func init() {
	fixtureData_1_8 := fixtureGenerator{
		generatePass: func(p *corev1.Pod) []*corev1.Pod {
			return []*corev1.Pod{
				// only valid pod is to explicitly set allowPrivilegeEscalation to false in all containers
				p,
			}
		},
		generateFail: func(p *corev1.Pod) []*corev1.Pod {
			return []*corev1.Pod{
				// explicit true
				tweak(p, func(p *corev1.Pod) {
					p.Spec.Containers[0].SecurityContext.AllowPrivilegeEscalation = pointer.BoolPtr(true)
				}),
				tweak(p, func(p *corev1.Pod) {
					p.Spec.InitContainers[0].SecurityContext.AllowPrivilegeEscalation = pointer.BoolPtr(true)
				}),
				// nil AllowPrivilegeEscalation
				tweak(p, func(p *corev1.Pod) { p.Spec.Containers[0].SecurityContext.AllowPrivilegeEscalation = nil }),
				tweak(p, func(p *corev1.Pod) { p.Spec.InitContainers[0].SecurityContext.AllowPrivilegeEscalation = nil }),
				// nil security context
				tweak(p, func(p *corev1.Pod) { p.Spec.Containers[0].SecurityContext = nil }),
				tweak(p, func(p *corev1.Pod) { p.Spec.InitContainers[0].SecurityContext = nil }),
			}
		},
	}

	registerFixtureGenerator(
		fixtureKey{level: api.LevelRestricted, version: api.MajorMinorVersion(1, 8), check: "allowPrivilegeEscalation"},
		fixtureData_1_8,
	)
}
