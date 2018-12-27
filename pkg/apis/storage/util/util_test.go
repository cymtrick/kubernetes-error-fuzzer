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

package util

import (
	"fmt"
	"reflect"
	"testing"

	"k8s.io/apimachinery/pkg/util/diff"
	utilfeature "k8s.io/apiserver/pkg/util/feature"
	utilfeaturetesting "k8s.io/apiserver/pkg/util/feature/testing"
	api "k8s.io/kubernetes/pkg/apis/core"
	"k8s.io/kubernetes/pkg/apis/storage"
	"k8s.io/kubernetes/pkg/features"
)

func TestDropAlphaFields(t *testing.T) {
	bindingMode := storage.VolumeBindingWaitForFirstConsumer
	allowedTopologies := []api.TopologySelectorTerm{
		{
			MatchLabelExpressions: []api.TopologySelectorLabelRequirement{
				{
					Key:    "kubernetes.io/hostname",
					Values: []string{"node1"},
				},
			},
		},
	}

	// Test that field gets dropped when feature gate is not set
	defer utilfeaturetesting.SetFeatureGateDuringTest(t, utilfeature.DefaultFeatureGate, features.VolumeScheduling, false)()
	class := &storage.StorageClass{
		VolumeBindingMode: &bindingMode,
		AllowedTopologies: allowedTopologies,
	}
	DropDisabledFields(class, nil)
	if class.VolumeBindingMode != nil {
		t.Errorf("VolumeBindingMode field didn't get dropped: %+v", class.VolumeBindingMode)
	}
	if class.AllowedTopologies != nil {
		t.Errorf("AllowedTopologies field didn't get dropped: %+v", class.AllowedTopologies)
	}

	// Test that field does not get dropped when feature gate is set
	class = &storage.StorageClass{
		VolumeBindingMode: &bindingMode,
		AllowedTopologies: allowedTopologies,
	}
	defer utilfeaturetesting.SetFeatureGateDuringTest(t, utilfeature.DefaultFeatureGate, features.VolumeScheduling, true)()
	DropDisabledFields(class, nil)
	if class.VolumeBindingMode != &bindingMode {
		t.Errorf("VolumeBindingMode field got unexpectantly modified: %+v", class.VolumeBindingMode)
	}
	if !reflect.DeepEqual(class.AllowedTopologies, allowedTopologies) {
		t.Errorf("AllowedTopologies field got unexpectantly modified: %+v", class.AllowedTopologies)
	}
}

func TestDropAllowVolumeExpansion(t *testing.T) {
	allowVolumeExpansion := false
	scWithoutAllowVolumeExpansion := func() *storage.StorageClass {
		return &storage.StorageClass{}
	}
	scWithAllowVolumeExpansion := func() *storage.StorageClass {
		return &storage.StorageClass{
			AllowVolumeExpansion: &allowVolumeExpansion,
		}
	}

	scInfo := []struct {
		description             string
		hasAllowVolumeExpansion bool
		sc                      func() *storage.StorageClass
	}{
		{
			description:             "StorageClass Without AllowVolumeExpansion",
			hasAllowVolumeExpansion: false,
			sc:                      scWithoutAllowVolumeExpansion,
		},
		{
			description:             "StorageClass With AllowVolumeExpansion",
			hasAllowVolumeExpansion: true,
			sc:                      scWithAllowVolumeExpansion,
		},
		{
			description:             "is nil",
			hasAllowVolumeExpansion: false,
			sc:                      func() *storage.StorageClass { return nil },
		},
	}

	for _, enabled := range []bool{true, false} {
		for _, oldStorageClassInfo := range scInfo {
			for _, newStorageClassInfo := range scInfo {
				oldStorageClassHasAllowVolumeExpansion, oldStorageClass := oldStorageClassInfo.hasAllowVolumeExpansion, oldStorageClassInfo.sc()
				newStorageClassHasAllowVolumeExpansion, newStorageClass := newStorageClassInfo.hasAllowVolumeExpansion, newStorageClassInfo.sc()
				if newStorageClass == nil {
					continue
				}

				t.Run(fmt.Sprintf("feature enabled=%v, old StorageClass %v, new StorageClass %v", enabled, oldStorageClassInfo.description, newStorageClassInfo.description), func(t *testing.T) {
					defer utilfeaturetesting.SetFeatureGateDuringTest(t, utilfeature.DefaultFeatureGate, features.ExpandPersistentVolumes, enabled)()

					DropDisabledFields(newStorageClass, oldStorageClass)

					// old StorageClass should never be changed
					if !reflect.DeepEqual(oldStorageClass, oldStorageClassInfo.sc()) {
						t.Errorf("old StorageClass changed: %v", diff.ObjectReflectDiff(oldStorageClass, oldStorageClassInfo.sc()))
					}

					switch {
					case enabled || oldStorageClassHasAllowVolumeExpansion:
						// new StorageClass should not be changed if the feature is enabled, or if the old StorageClass had AllowVolumeExpansion
						if !reflect.DeepEqual(newStorageClass, newStorageClassInfo.sc()) {
							t.Errorf("new StorageClass changed: %v", diff.ObjectReflectDiff(newStorageClass, newStorageClassInfo.sc()))
						}
					case newStorageClassHasAllowVolumeExpansion:
						// new StorageClass should be changed
						if reflect.DeepEqual(newStorageClass, newStorageClassInfo.sc()) {
							t.Errorf("new StorageClass was not changed")
						}
						// new StorageClass should not have AllowVolumeExpansion
						if !reflect.DeepEqual(newStorageClass, scWithoutAllowVolumeExpansion()) {
							t.Errorf("new StorageClass had StorageClassAllowVolumeExpansion: %v", diff.ObjectReflectDiff(newStorageClass, scWithoutAllowVolumeExpansion()))
						}
					default:
						// new StorageClass should not need to be changed
						if !reflect.DeepEqual(newStorageClass, newStorageClassInfo.sc()) {
							t.Errorf("new StorageClass changed: %v", diff.ObjectReflectDiff(newStorageClass, newStorageClassInfo.sc()))
						}
					}
				})
			}
		}
	}
}
