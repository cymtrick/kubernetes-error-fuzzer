/*
Copyright 2020 The Kubernetes Authors.

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
	"fmt"

	v1 "k8s.io/api/core/v1"
	utilfeature "k8s.io/apiserver/pkg/util/feature"
	"k8s.io/client-go/tools/cache"
	"k8s.io/kubernetes/pkg/features"
)

const (
	// PodPVCIndex is the lookup name for the index function, which is to index by pod pvcs.
	PodPVCIndex = "pod-pvc-index"
)

// PodPVCIndexFunc creates an index function that returns PVC keys (=
// namespace/name) for given pod.  If enabled, this includes the PVCs
// that might be created for generic ephemeral volumes.
func PodPVCIndexFunc(genericEphemeralVolumeFeatureEnabled bool) func(obj interface{}) ([]string, error) {
	return func(obj interface{}) ([]string, error) {
		pod, ok := obj.(*v1.Pod)
		if !ok {
			return []string{}, nil
		}
		keys := []string{}
		for _, podVolume := range pod.Spec.Volumes {
			claimName := ""
			if pvcSource := podVolume.VolumeSource.PersistentVolumeClaim; pvcSource != nil {
				claimName = pvcSource.ClaimName
			}
			if ephemeralSource := podVolume.VolumeSource.Ephemeral; genericEphemeralVolumeFeatureEnabled && ephemeralSource != nil {
				claimName = pod.Name + "-" + podVolume.Name
			}
			if claimName != "" {
				keys = append(keys, fmt.Sprintf("%s/%s", pod.Namespace, claimName))
			}
		}
		return keys, nil
	}
}

// AddPodPVCIndexerIfNotPresent adds the PodPVCIndexFunc with the current global setting for GenericEphemeralVolume.
func AddPodPVCIndexerIfNotPresent(indexer cache.Indexer) error {
	return AddIndexerIfNotPresent(indexer, PodPVCIndex,
		PodPVCIndexFunc(utilfeature.DefaultFeatureGate.Enabled(features.GenericEphemeralVolume)))
}

// AddIndexerIfNotPresent adds the index function with the name into the cache indexer if not present
func AddIndexerIfNotPresent(indexer cache.Indexer, indexName string, indexFunc cache.IndexFunc) error {
	indexers := indexer.GetIndexers()
	if _, ok := indexers[indexName]; ok {
		return nil
	}
	return indexer.AddIndexers(cache.Indexers{indexName: indexFunc})
}
