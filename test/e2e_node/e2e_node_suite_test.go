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

package e2e_node

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"flag"
	"testing"
)

var kubeletHost = flag.String("kubelet-host", "localhost", "Host address of the kubelet")
var kubeletPort = flag.Int("kubelet-port", 10250, "Kubelet port")

var apiServerHost = flag.String("api-server-host", "localhost", "Host address of the api server")
var apiServerPort = flag.Int("api-server-port", 8080, "Api server port")

func TestE2eNode(t *testing.T) {
	flag.Parse()
	RegisterFailHandler(Fail)
	RunSpecs(t, "E2eNode Suite")
}

// Setup the kubelet on the node
var _ = BeforeSuite(func() {
})

// Tear down the kubelet on the node
var _ = AfterSuite(func() {
})
