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

package common

import "k8s.io/kubernetes/test/e2e/framework"

// SIGNetworkDescribe annotates the test with the SIG Network label.
func SIGNetworkDescribe(text string, body func()) bool {
	return framework.KubeDescribe("[sig-network] "+text, body)
}

// SIGNodeDescribe annotates the test with the SIG Node label.
func SIGNodeDescribe(text string, body func()) bool {
	return framework.KubeDescribe("[sig-node] "+text, body)
}

// SIGStorageDescribe annotates the test with the SIG Storage label.
func SIGStorageDescribe(text string, body func()) bool {
	return framework.KubeDescribe("[sig-storage] "+text, body)
}
