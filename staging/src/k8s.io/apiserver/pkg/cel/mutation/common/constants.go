/*
Copyright 2024 The Kubernetes Authors.

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

import (
	"github.com/google/cel-go/common/types/traits"
)

// RootTypeReferenceName is the root reference that all type names should start with.
const RootTypeReferenceName = "Object"

// ObjectTraits is the bitmask that represents traits that an object should have.
const ObjectTraits = traits.ContainerType
