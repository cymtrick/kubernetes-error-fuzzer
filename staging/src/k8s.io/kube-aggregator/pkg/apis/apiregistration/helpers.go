/*
Copyright 2016 The Kubernetes Authors.

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

package apiregistration

import (
	"sort"
	"strings"

	"k8s.io/apimachinery/pkg/runtime/schema"
)

func SortedByGroup(servers []*APIService) [][]*APIService {
	serversByPriority := ByPriority(servers)
	sort.Sort(serversByPriority)

	ret := [][]*APIService{}
	for _, curr := range serversByPriority {
		// check to see if we already have an entry for this group
		existingIndex := -1
		for j, groupInReturn := range ret {
			if groupInReturn[0].Spec.Group == curr.Spec.Group {
				existingIndex = j
				break
			}
		}

		if existingIndex >= 0 {
			ret[existingIndex] = append(ret[existingIndex], curr)
			continue
		}

		ret = append(ret, []*APIService{curr})
	}

	return ret
}

type ByPriority []*APIService

func (s ByPriority) Len() int      { return len(s) }
func (s ByPriority) Swap(i, j int) { s[i], s[j] = s[j], s[i] }
func (s ByPriority) Less(i, j int) bool {
	if s[i].Spec.Priority == s[j].Spec.Priority {
		return s[i].Name < s[j].Name
	}
	return s[i].Spec.Priority < s[j].Spec.Priority
}

// APIServiceNameToGroupVersion returns the GroupVersion for a given apiServiceName.  The name
// must be valid, but any object you get back from an informer will be valid.
func APIServiceNameToGroupVersion(apiServiceName string) schema.GroupVersion {
	tokens := strings.SplitN(apiServiceName, ".", 2)
	return schema.GroupVersion{Group: tokens[1], Version: tokens[0]}
}

// SetAPIServiceCondition sets the status condition.  It either overwrites the existing one or
// creates a new one
func SetAPIServiceCondition(apiService *APIService, newCondition APIServiceCondition) {
	var existingCondition *APIServiceCondition
	for i := range apiService.Status.Conditions {
		if apiService.Status.Conditions[i].Type == newCondition.Type {
			existingCondition = &apiService.Status.Conditions[i]
			break
		}
	}
	if existingCondition == nil {
		apiService.Status.Conditions = append(apiService.Status.Conditions, newCondition)
		return
	}

	if existingCondition.Status != newCondition.Status {
		existingCondition.Status = newCondition.Status
		existingCondition.LastTransitionTime = newCondition.LastTransitionTime
	}

	existingCondition.Reason = newCondition.Reason
	existingCondition.Message = newCondition.Message
}

// IsAPIServiceConditionTrue indicates if the condition is present and strictly true
func IsAPIServiceConditionTrue(apiService *APIService, conditionType APIServiceConditionType) bool {
	for _, condition := range apiService.Status.Conditions {
		if condition.Type == conditionType && condition.Status == ConditionTrue {
			return true
		}
	}
	return false
}
