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

package volumebinding

import (
	"math"

	"k8s.io/kubernetes/pkg/controller/volume/scheduling"
	"k8s.io/kubernetes/pkg/scheduler/framework/plugins/helper"
)

type classResourceMap map[string]*scheduling.VolumeResource

type volumeCapacityScorer func(classResourceMap) int64

func buildScorerFunction(scoringFunctionShape helper.FunctionShape) volumeCapacityScorer {
	rawScoringFunction := helper.BuildBrokenLinearFunction(scoringFunctionShape)
	f := func(requested, capacity int64) int64 {
		if capacity == 0 || requested > capacity {
			return rawScoringFunction(maxUtilization)
		}

		return rawScoringFunction(requested * maxUtilization / capacity)
	}
	return func(classResources classResourceMap) int64 {
		var nodeScore, weightSum int64
		for _, resource := range classResources {
			classScore := f(resource.Requested, resource.Capacity)
			nodeScore += classScore
			// in alpha stage, all classes have the same weight
			weightSum += 1
		}
		if weightSum == 0 {
			return 0
		}
		return int64(math.Round(float64(nodeScore) / float64(weightSum)))
	}
}
