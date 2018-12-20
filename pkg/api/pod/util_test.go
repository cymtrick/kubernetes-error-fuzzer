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

package pod

import (
	"fmt"
	"reflect"
	"strings"
	"testing"

	"k8s.io/apimachinery/pkg/util/diff"
	"k8s.io/apimachinery/pkg/util/sets"
	"k8s.io/apimachinery/pkg/util/validation/field"
	utilfeature "k8s.io/apiserver/pkg/util/feature"
	utilfeaturetesting "k8s.io/apiserver/pkg/util/feature/testing"
	api "k8s.io/kubernetes/pkg/apis/core"
	"k8s.io/kubernetes/pkg/features"
)

func TestPodSecrets(t *testing.T) {
	// Stub containing all possible secret references in a pod.
	// The names of the referenced secrets match struct paths detected by reflection.
	pod := &api.Pod{
		Spec: api.PodSpec{
			Containers: []api.Container{{
				EnvFrom: []api.EnvFromSource{{
					SecretRef: &api.SecretEnvSource{
						LocalObjectReference: api.LocalObjectReference{
							Name: "Spec.Containers[*].EnvFrom[*].SecretRef"}}}},
				Env: []api.EnvVar{{
					ValueFrom: &api.EnvVarSource{
						SecretKeyRef: &api.SecretKeySelector{
							LocalObjectReference: api.LocalObjectReference{
								Name: "Spec.Containers[*].Env[*].ValueFrom.SecretKeyRef"}}}}}}},
			ImagePullSecrets: []api.LocalObjectReference{{
				Name: "Spec.ImagePullSecrets"}},
			InitContainers: []api.Container{{
				EnvFrom: []api.EnvFromSource{{
					SecretRef: &api.SecretEnvSource{
						LocalObjectReference: api.LocalObjectReference{
							Name: "Spec.InitContainers[*].EnvFrom[*].SecretRef"}}}},
				Env: []api.EnvVar{{
					ValueFrom: &api.EnvVarSource{
						SecretKeyRef: &api.SecretKeySelector{
							LocalObjectReference: api.LocalObjectReference{
								Name: "Spec.InitContainers[*].Env[*].ValueFrom.SecretKeyRef"}}}}}}},
			Volumes: []api.Volume{{
				VolumeSource: api.VolumeSource{
					AzureFile: &api.AzureFileVolumeSource{
						SecretName: "Spec.Volumes[*].VolumeSource.AzureFile.SecretName"}}}, {
				VolumeSource: api.VolumeSource{
					CephFS: &api.CephFSVolumeSource{
						SecretRef: &api.LocalObjectReference{
							Name: "Spec.Volumes[*].VolumeSource.CephFS.SecretRef"}}}}, {
				VolumeSource: api.VolumeSource{
					Cinder: &api.CinderVolumeSource{
						SecretRef: &api.LocalObjectReference{
							Name: "Spec.Volumes[*].VolumeSource.Cinder.SecretRef"}}}}, {
				VolumeSource: api.VolumeSource{
					FlexVolume: &api.FlexVolumeSource{
						SecretRef: &api.LocalObjectReference{
							Name: "Spec.Volumes[*].VolumeSource.FlexVolume.SecretRef"}}}}, {
				VolumeSource: api.VolumeSource{
					Projected: &api.ProjectedVolumeSource{
						Sources: []api.VolumeProjection{{
							Secret: &api.SecretProjection{
								LocalObjectReference: api.LocalObjectReference{
									Name: "Spec.Volumes[*].VolumeSource.Projected.Sources[*].Secret"}}}}}}}, {
				VolumeSource: api.VolumeSource{
					RBD: &api.RBDVolumeSource{
						SecretRef: &api.LocalObjectReference{
							Name: "Spec.Volumes[*].VolumeSource.RBD.SecretRef"}}}}, {
				VolumeSource: api.VolumeSource{
					Secret: &api.SecretVolumeSource{
						SecretName: "Spec.Volumes[*].VolumeSource.Secret.SecretName"}}}, {
				VolumeSource: api.VolumeSource{
					Secret: &api.SecretVolumeSource{
						SecretName: "Spec.Volumes[*].VolumeSource.Secret"}}}, {
				VolumeSource: api.VolumeSource{
					ScaleIO: &api.ScaleIOVolumeSource{
						SecretRef: &api.LocalObjectReference{
							Name: "Spec.Volumes[*].VolumeSource.ScaleIO.SecretRef"}}}}, {
				VolumeSource: api.VolumeSource{
					ISCSI: &api.ISCSIVolumeSource{
						SecretRef: &api.LocalObjectReference{
							Name: "Spec.Volumes[*].VolumeSource.ISCSI.SecretRef"}}}}, {
				VolumeSource: api.VolumeSource{
					StorageOS: &api.StorageOSVolumeSource{
						SecretRef: &api.LocalObjectReference{
							Name: "Spec.Volumes[*].VolumeSource.StorageOS.SecretRef"}}}}},
		},
	}
	extractedNames := sets.NewString()
	VisitPodSecretNames(pod, func(name string) bool {
		extractedNames.Insert(name)
		return true
	})

	// excludedSecretPaths holds struct paths to fields with "secret" in the name that are not actually references to secret API objects
	excludedSecretPaths := sets.NewString(
		"Spec.Volumes[*].VolumeSource.CephFS.SecretFile",
	)
	// expectedSecretPaths holds struct paths to fields with "secret" in the name that are references to secret API objects.
	// every path here should be represented as an example in the Pod stub above, with the secret name set to the path.
	expectedSecretPaths := sets.NewString(
		"Spec.Containers[*].EnvFrom[*].SecretRef",
		"Spec.Containers[*].Env[*].ValueFrom.SecretKeyRef",
		"Spec.ImagePullSecrets",
		"Spec.InitContainers[*].EnvFrom[*].SecretRef",
		"Spec.InitContainers[*].Env[*].ValueFrom.SecretKeyRef",
		"Spec.Volumes[*].VolumeSource.AzureFile.SecretName",
		"Spec.Volumes[*].VolumeSource.CephFS.SecretRef",
		"Spec.Volumes[*].VolumeSource.Cinder.SecretRef",
		"Spec.Volumes[*].VolumeSource.FlexVolume.SecretRef",
		"Spec.Volumes[*].VolumeSource.Projected.Sources[*].Secret",
		"Spec.Volumes[*].VolumeSource.RBD.SecretRef",
		"Spec.Volumes[*].VolumeSource.Secret",
		"Spec.Volumes[*].VolumeSource.Secret.SecretName",
		"Spec.Volumes[*].VolumeSource.ScaleIO.SecretRef",
		"Spec.Volumes[*].VolumeSource.ISCSI.SecretRef",
		"Spec.Volumes[*].VolumeSource.StorageOS.SecretRef",
	)
	secretPaths := collectResourcePaths(t, "secret", nil, "", reflect.TypeOf(&api.Pod{}))
	secretPaths = secretPaths.Difference(excludedSecretPaths)
	if missingPaths := expectedSecretPaths.Difference(secretPaths); len(missingPaths) > 0 {
		t.Logf("Missing expected secret paths:\n%s", strings.Join(missingPaths.List(), "\n"))
		t.Error("Missing expected secret paths. Verify VisitPodSecretNames() is correctly finding the missing paths, then correct expectedSecretPaths")
	}
	if extraPaths := secretPaths.Difference(expectedSecretPaths); len(extraPaths) > 0 {
		t.Logf("Extra secret paths:\n%s", strings.Join(extraPaths.List(), "\n"))
		t.Error("Extra fields with 'secret' in the name found. Verify VisitPodSecretNames() is including these fields if appropriate, then correct expectedSecretPaths")
	}

	if missingNames := expectedSecretPaths.Difference(extractedNames); len(missingNames) > 0 {
		t.Logf("Missing expected secret names:\n%s", strings.Join(missingNames.List(), "\n"))
		t.Error("Missing expected secret names. Verify the pod stub above includes these references, then verify VisitPodSecretNames() is correctly finding the missing names")
	}
	if extraNames := extractedNames.Difference(expectedSecretPaths); len(extraNames) > 0 {
		t.Logf("Extra secret names:\n%s", strings.Join(extraNames.List(), "\n"))
		t.Error("Extra secret names extracted. Verify VisitPodSecretNames() is correctly extracting secret names")
	}
}

// collectResourcePaths traverses the object, computing all the struct paths that lead to fields with resourcename in the name.
func collectResourcePaths(t *testing.T, resourcename string, path *field.Path, name string, tp reflect.Type) sets.String {
	resourcename = strings.ToLower(resourcename)
	resourcePaths := sets.NewString()

	if tp.Kind() == reflect.Ptr {
		resourcePaths.Insert(collectResourcePaths(t, resourcename, path, name, tp.Elem()).List()...)
		return resourcePaths
	}

	if strings.Contains(strings.ToLower(name), resourcename) {
		resourcePaths.Insert(path.String())
	}

	switch tp.Kind() {
	case reflect.Ptr:
		resourcePaths.Insert(collectResourcePaths(t, resourcename, path, name, tp.Elem()).List()...)
	case reflect.Struct:
		for i := 0; i < tp.NumField(); i++ {
			field := tp.Field(i)
			resourcePaths.Insert(collectResourcePaths(t, resourcename, path.Child(field.Name), field.Name, field.Type).List()...)
		}
	case reflect.Interface:
		t.Errorf("cannot find %s fields in interface{} field %s", resourcename, path.String())
	case reflect.Map:
		resourcePaths.Insert(collectResourcePaths(t, resourcename, path.Key("*"), "", tp.Elem()).List()...)
	case reflect.Slice:
		resourcePaths.Insert(collectResourcePaths(t, resourcename, path.Key("*"), "", tp.Elem()).List()...)
	default:
		// all primitive types
	}

	return resourcePaths
}

func TestPodConfigmaps(t *testing.T) {
	// Stub containing all possible ConfigMap references in a pod.
	// The names of the referenced ConfigMaps match struct paths detected by reflection.
	pod := &api.Pod{
		Spec: api.PodSpec{
			Containers: []api.Container{{
				EnvFrom: []api.EnvFromSource{{
					ConfigMapRef: &api.ConfigMapEnvSource{
						LocalObjectReference: api.LocalObjectReference{
							Name: "Spec.Containers[*].EnvFrom[*].ConfigMapRef"}}}},
				Env: []api.EnvVar{{
					ValueFrom: &api.EnvVarSource{
						ConfigMapKeyRef: &api.ConfigMapKeySelector{
							LocalObjectReference: api.LocalObjectReference{
								Name: "Spec.Containers[*].Env[*].ValueFrom.ConfigMapKeyRef"}}}}}}},
			InitContainers: []api.Container{{
				EnvFrom: []api.EnvFromSource{{
					ConfigMapRef: &api.ConfigMapEnvSource{
						LocalObjectReference: api.LocalObjectReference{
							Name: "Spec.InitContainers[*].EnvFrom[*].ConfigMapRef"}}}},
				Env: []api.EnvVar{{
					ValueFrom: &api.EnvVarSource{
						ConfigMapKeyRef: &api.ConfigMapKeySelector{
							LocalObjectReference: api.LocalObjectReference{
								Name: "Spec.InitContainers[*].Env[*].ValueFrom.ConfigMapKeyRef"}}}}}}},
			Volumes: []api.Volume{{
				VolumeSource: api.VolumeSource{
					Projected: &api.ProjectedVolumeSource{
						Sources: []api.VolumeProjection{{
							ConfigMap: &api.ConfigMapProjection{
								LocalObjectReference: api.LocalObjectReference{
									Name: "Spec.Volumes[*].VolumeSource.Projected.Sources[*].ConfigMap"}}}}}}}, {
				VolumeSource: api.VolumeSource{
					ConfigMap: &api.ConfigMapVolumeSource{
						LocalObjectReference: api.LocalObjectReference{
							Name: "Spec.Volumes[*].VolumeSource.ConfigMap"}}}}},
		},
	}
	extractedNames := sets.NewString()
	VisitPodConfigmapNames(pod, func(name string) bool {
		extractedNames.Insert(name)
		return true
	})

	// expectedPaths holds struct paths to fields with "ConfigMap" in the name that are references to ConfigMap API objects.
	// every path here should be represented as an example in the Pod stub above, with the ConfigMap name set to the path.
	expectedPaths := sets.NewString(
		"Spec.Containers[*].EnvFrom[*].ConfigMapRef",
		"Spec.Containers[*].Env[*].ValueFrom.ConfigMapKeyRef",
		"Spec.InitContainers[*].EnvFrom[*].ConfigMapRef",
		"Spec.InitContainers[*].Env[*].ValueFrom.ConfigMapKeyRef",
		"Spec.Volumes[*].VolumeSource.Projected.Sources[*].ConfigMap",
		"Spec.Volumes[*].VolumeSource.ConfigMap",
	)
	collectPaths := collectResourcePaths(t, "ConfigMap", nil, "", reflect.TypeOf(&api.Pod{}))
	if missingPaths := expectedPaths.Difference(collectPaths); len(missingPaths) > 0 {
		t.Logf("Missing expected paths:\n%s", strings.Join(missingPaths.List(), "\n"))
		t.Error("Missing expected paths. Verify VisitPodConfigmapNames() is correctly finding the missing paths, then correct expectedPaths")
	}
	if extraPaths := collectPaths.Difference(expectedPaths); len(extraPaths) > 0 {
		t.Logf("Extra paths:\n%s", strings.Join(extraPaths.List(), "\n"))
		t.Error("Extra fields with resource in the name found. Verify VisitPodConfigmapNames() is including these fields if appropriate, then correct expectedPaths")
	}

	if missingNames := expectedPaths.Difference(extractedNames); len(missingNames) > 0 {
		t.Logf("Missing expected names:\n%s", strings.Join(missingNames.List(), "\n"))
		t.Error("Missing expected names. Verify the pod stub above includes these references, then verify VisitPodConfigmapNames() is correctly finding the missing names")
	}
	if extraNames := extractedNames.Difference(expectedPaths); len(extraNames) > 0 {
		t.Logf("Extra names:\n%s", strings.Join(extraNames.List(), "\n"))
		t.Error("Extra names extracted. Verify VisitPodConfigmapNames() is correctly extracting resource names")
	}
}

func TestDropAlphaVolumeDevices(t *testing.T) {
	testPod := api.Pod{
		Spec: api.PodSpec{
			RestartPolicy: api.RestartPolicyNever,
			Containers: []api.Container{
				{
					Name:  "container1",
					Image: "testimage",
					VolumeDevices: []api.VolumeDevice{
						{
							Name:       "myvolume",
							DevicePath: "/usr/test",
						},
					},
				},
			},
			InitContainers: []api.Container{
				{
					Name:  "container1",
					Image: "testimage",
					VolumeDevices: []api.VolumeDevice{
						{
							Name:       "myvolume",
							DevicePath: "/usr/test",
						},
					},
				},
			},
			Volumes: []api.Volume{
				{
					Name: "myvolume",
					VolumeSource: api.VolumeSource{
						HostPath: &api.HostPathVolumeSource{
							Path: "/dev/xvdc",
						},
					},
				},
			},
		},
	}

	// Enable alpha feature BlockVolume
	defer utilfeaturetesting.SetFeatureGateDuringTest(t, utilfeature.DefaultFeatureGate, features.BlockVolume, true)()

	// now test dropping the fields - should not be dropped
	DropDisabledFields(&testPod.Spec, nil)

	// check to make sure VolumeDevices is still present
	// if featureset is set to true
	if testPod.Spec.Containers[0].VolumeDevices == nil {
		t.Error("VolumeDevices in Container should not have been dropped based on feature-gate")
	}
	if testPod.Spec.InitContainers[0].VolumeDevices == nil {
		t.Error("VolumeDevices in InitContainers should not have been dropped based on feature-gate")
	}

	// Disable alpha feature BlockVolume
	defer utilfeaturetesting.SetFeatureGateDuringTest(t, utilfeature.DefaultFeatureGate, features.BlockVolume, false)()

	// now test dropping the fields
	DropDisabledFields(&testPod.Spec, nil)

	// check to make sure VolumeDevices is nil
	// if featureset is set to false
	if testPod.Spec.Containers[0].VolumeDevices != nil {
		t.Error("DropDisabledFields for Containers failed")
	}
	if testPod.Spec.InitContainers[0].VolumeDevices != nil {
		t.Error("DropDisabledFields for InitContainers failed")
	}
}

func TestDropSubPath(t *testing.T) {
	podWithSubpaths := func() *api.Pod {
		return &api.Pod{
			Spec: api.PodSpec{
				RestartPolicy:  api.RestartPolicyNever,
				Containers:     []api.Container{{Name: "container1", Image: "testimage", VolumeMounts: []api.VolumeMount{{Name: "a", SubPath: "foo"}, {Name: "a", SubPath: "foo2"}, {Name: "a", SubPath: "foo3"}}}},
				InitContainers: []api.Container{{Name: "container1", Image: "testimage", VolumeMounts: []api.VolumeMount{{Name: "a", SubPath: "foo"}, {Name: "a", SubPath: "foo2"}}}},
				Volumes:        []api.Volume{{Name: "a", VolumeSource: api.VolumeSource{HostPath: &api.HostPathVolumeSource{Path: "/dev/xvdc"}}}},
			},
		}
	}
	podWithoutSubpaths := func() *api.Pod {
		return &api.Pod{
			Spec: api.PodSpec{
				RestartPolicy:  api.RestartPolicyNever,
				Containers:     []api.Container{{Name: "container1", Image: "testimage", VolumeMounts: []api.VolumeMount{{Name: "a", SubPath: ""}, {Name: "a", SubPath: ""}, {Name: "a", SubPath: ""}}}},
				InitContainers: []api.Container{{Name: "container1", Image: "testimage", VolumeMounts: []api.VolumeMount{{Name: "a", SubPath: ""}, {Name: "a", SubPath: ""}}}},
				Volumes:        []api.Volume{{Name: "a", VolumeSource: api.VolumeSource{HostPath: &api.HostPathVolumeSource{Path: "/dev/xvdc"}}}},
			},
		}
	}

	podInfo := []struct {
		description string
		hasSubpaths bool
		pod         func() *api.Pod
	}{
		{
			description: "has subpaths",
			hasSubpaths: true,
			pod:         podWithSubpaths,
		},
		{
			description: "does not have subpaths",
			hasSubpaths: false,
			pod:         podWithoutSubpaths,
		},
		{
			description: "is nil",
			hasSubpaths: false,
			pod:         func() *api.Pod { return nil },
		},
	}

	for _, enabled := range []bool{true, false} {
		for _, oldPodInfo := range podInfo {
			for _, newPodInfo := range podInfo {
				oldPodHasSubpaths, oldPod := oldPodInfo.hasSubpaths, oldPodInfo.pod()
				newPodHasSubpaths, newPod := newPodInfo.hasSubpaths, newPodInfo.pod()
				if newPod == nil {
					continue
				}

				t.Run(fmt.Sprintf("feature enabled=%v, old pod %v, new pod %v", enabled, oldPodInfo.description, newPodInfo.description), func(t *testing.T) {
					defer utilfeaturetesting.SetFeatureGateDuringTest(t, utilfeature.DefaultFeatureGate, features.VolumeSubpath, enabled)()

					var oldPodSpec *api.PodSpec
					if oldPod != nil {
						oldPodSpec = &oldPod.Spec
					}
					DropDisabledFields(&newPod.Spec, oldPodSpec)

					// old pod should never be changed
					if !reflect.DeepEqual(oldPod, oldPodInfo.pod()) {
						t.Errorf("old pod changed: %v", diff.ObjectReflectDiff(oldPod, oldPodInfo.pod()))
					}

					switch {
					case enabled || oldPodHasSubpaths:
						// new pod should not be changed if the feature is enabled, or if the old pod had subpaths
						if !reflect.DeepEqual(newPod, newPodInfo.pod()) {
							t.Errorf("new pod changed: %v", diff.ObjectReflectDiff(newPod, newPodInfo.pod()))
						}
					case newPodHasSubpaths:
						// new pod should be changed
						if reflect.DeepEqual(newPod, newPodInfo.pod()) {
							t.Errorf("new pod was not changed")
						}
						// new pod should not have subpaths
						if !reflect.DeepEqual(newPod, podWithoutSubpaths()) {
							t.Errorf("new pod had subpaths: %v", diff.ObjectReflectDiff(newPod, podWithoutSubpaths()))
						}
					default:
						// new pod should not need to be changed
						if !reflect.DeepEqual(newPod, newPodInfo.pod()) {
							t.Errorf("new pod changed: %v", diff.ObjectReflectDiff(newPod, newPodInfo.pod()))
						}
					}
				})
			}
		}
	}
}

func TestDropRuntimeClass(t *testing.T) {
	runtimeClassName := "some_container_engine"
	podWithoutRuntimeClass := func() *api.Pod {
		return &api.Pod{
			Spec: api.PodSpec{
				RuntimeClassName: nil,
			},
		}
	}
	podWithRuntimeClass := func() *api.Pod {
		return &api.Pod{
			Spec: api.PodSpec{
				RuntimeClassName: &runtimeClassName,
			},
		}
	}

	podInfo := []struct {
		description            string
		hasPodRuntimeClassName bool
		pod                    func() *api.Pod
	}{
		{
			description:            "pod Without RuntimeClassName",
			hasPodRuntimeClassName: false,
			pod:                    podWithoutRuntimeClass,
		},
		{
			description:            "pod With RuntimeClassName",
			hasPodRuntimeClassName: true,
			pod:                    podWithRuntimeClass,
		},
		{
			description:            "is nil",
			hasPodRuntimeClassName: false,
			pod:                    func() *api.Pod { return nil },
		},
	}

	for _, enabled := range []bool{true, false} {
		for _, oldPodInfo := range podInfo {
			for _, newPodInfo := range podInfo {
				oldPodHasRuntimeClassName, oldPod := oldPodInfo.hasPodRuntimeClassName, oldPodInfo.pod()
				newPodHasRuntimeClassName, newPod := newPodInfo.hasPodRuntimeClassName, newPodInfo.pod()
				if newPod == nil {
					continue
				}

				t.Run(fmt.Sprintf("feature enabled=%v, old pod %v, new pod %v", enabled, oldPodInfo.description, newPodInfo.description), func(t *testing.T) {
					defer utilfeaturetesting.SetFeatureGateDuringTest(t, utilfeature.DefaultFeatureGate, features.RuntimeClass, enabled)()

					var oldPodSpec *api.PodSpec
					if oldPod != nil {
						oldPodSpec = &oldPod.Spec
					}
					DropDisabledFields(&newPod.Spec, oldPodSpec)

					// old pod should never be changed
					if !reflect.DeepEqual(oldPod, oldPodInfo.pod()) {
						t.Errorf("old pod changed: %v", diff.ObjectReflectDiff(oldPod, oldPodInfo.pod()))
					}

					switch {
					case enabled || oldPodHasRuntimeClassName:
						// new pod should not be changed if the feature is enabled, or if the old pod had RuntimeClass
						if !reflect.DeepEqual(newPod, newPodInfo.pod()) {
							t.Errorf("new pod changed: %v", diff.ObjectReflectDiff(newPod, newPodInfo.pod()))
						}
					case newPodHasRuntimeClassName:
						// new pod should be changed
						if reflect.DeepEqual(newPod, newPodInfo.pod()) {
							t.Errorf("new pod was not changed")
						}
						// new pod should not have RuntimeClass
						if !reflect.DeepEqual(newPod, podWithoutRuntimeClass()) {
							t.Errorf("new pod had PodRuntimeClassName: %v", diff.ObjectReflectDiff(newPod, podWithoutRuntimeClass()))
						}
					default:
						// new pod should not need to be changed
						if !reflect.DeepEqual(newPod, newPodInfo.pod()) {
							t.Errorf("new pod changed: %v", diff.ObjectReflectDiff(newPod, newPodInfo.pod()))
						}
					}
				})
			}
		}
	}
}

func TestDropProcMount(t *testing.T) {
	procMount := api.UnmaskedProcMount
	defaultProcMount := api.DefaultProcMount
	podWithProcMount := func() *api.Pod {
		return &api.Pod{
			Spec: api.PodSpec{
				RestartPolicy:  api.RestartPolicyNever,
				Containers:     []api.Container{{Name: "container1", Image: "testimage", SecurityContext: &api.SecurityContext{ProcMount: &procMount}}},
				InitContainers: []api.Container{{Name: "container1", Image: "testimage", SecurityContext: &api.SecurityContext{ProcMount: &procMount}}},
			},
		}
	}
	podWithoutProcMount := func() *api.Pod {
		return &api.Pod{
			Spec: api.PodSpec{
				RestartPolicy:  api.RestartPolicyNever,
				Containers:     []api.Container{{Name: "container1", Image: "testimage", SecurityContext: &api.SecurityContext{ProcMount: &defaultProcMount}}},
				InitContainers: []api.Container{{Name: "container1", Image: "testimage", SecurityContext: &api.SecurityContext{ProcMount: &defaultProcMount}}},
			},
		}
	}

	podInfo := []struct {
		description  string
		hasProcMount bool
		pod          func() *api.Pod
	}{
		{
			description:  "has ProcMount",
			hasProcMount: true,
			pod:          podWithProcMount,
		},
		{
			description:  "does not have ProcMount",
			hasProcMount: false,
			pod:          podWithoutProcMount,
		},
		{
			description:  "is nil",
			hasProcMount: false,
			pod:          func() *api.Pod { return nil },
		},
	}

	for _, enabled := range []bool{true, false} {
		for _, oldPodInfo := range podInfo {
			for _, newPodInfo := range podInfo {
				oldPodHasProcMount, oldPod := oldPodInfo.hasProcMount, oldPodInfo.pod()
				newPodHasProcMount, newPod := newPodInfo.hasProcMount, newPodInfo.pod()
				if newPod == nil {
					continue
				}

				t.Run(fmt.Sprintf("feature enabled=%v, old pod %v, new pod %v", enabled, oldPodInfo.description, newPodInfo.description), func(t *testing.T) {
					defer utilfeaturetesting.SetFeatureGateDuringTest(t, utilfeature.DefaultFeatureGate, features.ProcMountType, enabled)()

					var oldPodSpec *api.PodSpec
					if oldPod != nil {
						oldPodSpec = &oldPod.Spec
					}
					DropDisabledFields(&newPod.Spec, oldPodSpec)

					// old pod should never be changed
					if !reflect.DeepEqual(oldPod, oldPodInfo.pod()) {
						t.Errorf("old pod changed: %v", diff.ObjectReflectDiff(oldPod, oldPodInfo.pod()))
					}

					switch {
					case enabled || oldPodHasProcMount:
						// new pod should not be changed if the feature is enabled, or if the old pod had ProcMount
						if !reflect.DeepEqual(newPod, newPodInfo.pod()) {
							t.Errorf("new pod changed: %v", diff.ObjectReflectDiff(newPod, newPodInfo.pod()))
						}
					case newPodHasProcMount:
						// new pod should be changed
						if reflect.DeepEqual(newPod, newPodInfo.pod()) {
							t.Errorf("new pod was not changed")
						}
						// new pod should not have ProcMount
						if !reflect.DeepEqual(newPod, podWithoutProcMount()) {
							t.Errorf("new pod had ProcMount: %v", diff.ObjectReflectDiff(newPod, podWithoutProcMount()))
						}
					default:
						// new pod should not need to be changed
						if !reflect.DeepEqual(newPod, newPodInfo.pod()) {
							t.Errorf("new pod changed: %v", diff.ObjectReflectDiff(newPod, newPodInfo.pod()))
						}
					}
				})
			}
		}
	}
}
