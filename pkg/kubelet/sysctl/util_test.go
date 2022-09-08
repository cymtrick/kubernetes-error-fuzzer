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

package sysctl

import (
	"testing"

	v1 "k8s.io/api/core/v1"

	"github.com/stretchr/testify/assert"
)

// TestConvertSysctlVariableToDotsSeparator tests whether the sysctl variable
// can be correctly converted to a dot as a separator.
func TestConvertSysctlVariableToDotsSeparator(t *testing.T) {
	type testCase struct {
		in  string
		out string
	}
	valid := []testCase{
		{in: "kernel.shm_rmid_forced", out: "kernel.shm_rmid_forced"},
		{in: "kernel/shm_rmid_forced", out: "kernel.shm_rmid_forced"},
		{in: "net.ipv4.conf.eno2/100.rp_filter", out: "net.ipv4.conf.eno2/100.rp_filter"},
		{in: "net/ipv4/conf/eno2.100/rp_filter", out: "net.ipv4.conf.eno2/100.rp_filter"},
		{in: "net/ipv4/ip_local_port_range", out: "net.ipv4.ip_local_port_range"},
		{in: "kernel/msgmax", out: "kernel.msgmax"},
		{in: "kernel/sem", out: "kernel.sem"},
	}

	for _, test := range valid {
		convertSysctlVal := convertSysctlVariableToDotsSeparator(test.in)
		assert.Equalf(t, test.out, convertSysctlVal, "The sysctl variable was not converted correctly. got: %s, want: %s", convertSysctlVal, test.out)
	}
}

// TestConvertPodSysctlsVariableToDotsSeparator tests whether the sysctls variable
// can be correctly converted to a dot as a separator.
func TestConvertPodSysctlsVariableToDotsSeparator(t *testing.T) {

	sysctls := []v1.Sysctl{
		{
			Name:  "kernel.msgmax",
			Value: "8192",
		},
		{
			Name:  "kernel.shm_rmid_forced",
			Value: "1",
		},
		{
			Name:  "net.ipv4.conf.eno2/100.rp_filter",
			Value: "1",
		},
		{
			Name:  "net/ipv4/ip_local_port_range",
			Value: "1024 65535",
		},
	}
	exceptSysctls := []v1.Sysctl{
		{
			Name:  "kernel.msgmax",
			Value: "8192",
		},
		{
			Name:  "kernel.shm_rmid_forced",
			Value: "1",
		},
		{
			Name:  "net.ipv4.conf.eno2/100.rp_filter",
			Value: "1",
		},
		{
			Name:  "net.ipv4.ip_local_port_range",
			Value: "1024 65535",
		},
	}
	securityContext := &v1.PodSecurityContext{
		Sysctls: sysctls,
	}

	ConvertPodSysctlsVariableToDotsSeparator(securityContext)
	assert.Equalf(t, exceptSysctls, securityContext.Sysctls, "The sysctls name was not converted correctly. got: %s, want: %s", securityContext.Sysctls, exceptSysctls)

}
