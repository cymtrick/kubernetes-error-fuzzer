//go:build windows
// +build windows

/*
Copyright 2023 The Kubernetes Authors.

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

package dns

import (
	"testing"

	runtimeapi "k8s.io/cri-api/pkg/apis/runtime/v1"
)

var (
	defaultResolvConf = hostResolvConf
)

func fakeGetHostDNSConfigCustom(resolverConfig string) (*runtimeapi.DNSConfig, error) {
	return &runtimeapi.DNSConfig{
		Servers:  []string{testHostNameserver},
		Searches: []string{testHostDomain},
	}, nil
}

// getResolvConf returns the hostResolvConf string, which will be used to get the Host's DNS configuration.
func getResolvConf(t *testing.T) (string, func()) {
	return hostResolvConf, func() {}
}
