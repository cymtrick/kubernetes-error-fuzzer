/*
Copyright 2016 The Kubernetes Authors All rights reserved.

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

package fake

import (
	extensions_unversioned "k8s.io/kubernetes/pkg/client/typed/generated/extensions/unversioned"
	extensions_unversioned_fake "k8s.io/kubernetes/pkg/client/typed/generated/extensions/unversioned/fake"
	legacy_unversioned "k8s.io/kubernetes/pkg/client/typed/generated/legacy/unversioned"
	legacy_unversioned_fake "k8s.io/kubernetes/pkg/client/typed/generated/legacy/unversioned/fake"
)

// Legacy retrieves the LegacyClient
func (c *Clientset) Legacy() legacy_unversioned.LegacyInterface {
	return &legacy_unversioned_fake.FakeLegacy{&c.Fake}
}

// Extensions retrieves the ExtensionsClient
func (c *Clientset) Extensions() extensions_unversioned.ExtensionsInterface {
	return &extensions_unversioned_fake.FakeExtensions{&c.Fake}
}
