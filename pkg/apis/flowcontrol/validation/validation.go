/*
Copyright 2019 The Kubernetes Authors.

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

package validation

import (
	"fmt"
	"k8s.io/apimachinery/pkg/api/validation"
	apimachineryvalidation "k8s.io/apimachinery/pkg/api/validation"
	"k8s.io/apimachinery/pkg/util/sets"
	"k8s.io/apimachinery/pkg/util/validation/field"
	"k8s.io/apiserver/pkg/util/shufflesharding"
	apivalidation "k8s.io/kubernetes/pkg/apis/core/validation"
	"k8s.io/kubernetes/pkg/apis/flowcontrol"
)

// ValidateFlowSchemaName validates name for flow-schema.
var ValidateFlowSchemaName = apimachineryvalidation.NameIsDNSSubdomain

// ValidatePriorityLevelConfigurationName validates name for priority-level-configuration.
var ValidatePriorityLevelConfigurationName = apimachineryvalidation.NameIsDNSSubdomain

var supportedDistinguisherMethods = sets.NewString(
	string(flowcontrol.FlowDistinguisherMethodByNamespaceType),
	string(flowcontrol.FlowDistinguisherMethodByUserType),
)

var priorityLevelConfigurationQueuingMaxQueues int32 = 10 * 1000 * 1000 // 10^7

var supportedVerbs = sets.NewString(
	"get",
	"list",
	"create",
	"update",
	"delete",
	"deletecollection",
	"patch",
	"watch",
	"proxy",
)

var supportedSubjectKinds = sets.NewString(
	string(flowcontrol.SubjectKindServiceAccount),
	string(flowcontrol.SubjectKindGroup),
	string(flowcontrol.SubjectKindUser),
)

var supportedQueuingType = sets.NewString(
	string(flowcontrol.PriorityLevelQueuingTypeQueueing),
	string(flowcontrol.PriorityLevelQueuingTypeExempt),
)

// ValidateFlowSchema validates the content of flow-schema
func ValidateFlowSchema(fs *flowcontrol.FlowSchema) field.ErrorList {
	allErrs := apivalidation.ValidateObjectMeta(&fs.ObjectMeta, false, ValidateFlowSchemaName, field.NewPath("metadata"))
	allErrs = append(allErrs, ValidateFlowSchemaSpec(&fs.Spec, field.NewPath("spec"))...)
	allErrs = append(allErrs, ValidateFlowSchemaStatus(&fs.Status, field.NewPath("status"))...)
	return allErrs
}

// ValidateFlowSchemaSpec validates the content of flow-schema's spec
func ValidateFlowSchemaSpec(spec *flowcontrol.FlowSchemaSpec, fldPath *field.Path) field.ErrorList {
	var allErrs field.ErrorList
	if spec.MatchingPrecedence <= 0 {
		allErrs = append(allErrs, field.Invalid(fldPath.Child("matchingPrecedence"), spec.MatchingPrecedence, "must be positive value"))
	}
	if spec.DistinguisherMethod != nil {
		if !supportedDistinguisherMethods.Has(string(spec.DistinguisherMethod.Type)) {
			allErrs = append(allErrs, field.NotSupported(fldPath.Child("distinguisherMethod").Child("type"), spec.DistinguisherMethod, supportedDistinguisherMethods.List()))
		}
	}
	if len(spec.PriorityLevelConfiguration.Name) > 0 {
		for _, msg := range ValidatePriorityLevelConfigurationName(spec.PriorityLevelConfiguration.Name, false) {
			allErrs = append(allErrs, field.Invalid(fldPath.Child("priorityLevelConfiguration").Child("name"), spec.PriorityLevelConfiguration.Name, msg))
		}
	} else {
		allErrs = append(allErrs, field.Required(fldPath.Child("priorityLevelConfiguration").Child("name"), "must reference a priority level"))
	}
	for i, rule := range spec.Rules {
		allErrs = append(allErrs, ValidateFlowSchemaPolicyRulesWithSubjects(&rule, fldPath.Child("rules").Index(i))...)
	}
	return allErrs
}

// ValidateFlowSchemaPolicyRulesWithSubjects validates policy-rule-with-subjects object.
func ValidateFlowSchemaPolicyRulesWithSubjects(rule *flowcontrol.PolicyRulesWithSubjects, fldPath *field.Path) field.ErrorList {
	var allErrs field.ErrorList
	if len(rule.Subjects) > 0 {
		for i, subject := range rule.Subjects {
			allErrs = append(allErrs, ValidateFlowSchemaSubject(&subject, fldPath.Child("subjects").Index(i))...)
		}
	} else {
		allErrs = append(allErrs, field.Required(fldPath.Child("subjects"), "subjects must contain at least one value"))
	}

	if len(rule.ResourceRules) == 0 && len(rule.NonResourceRules) == 0 {
		allErrs = append(allErrs, field.Required(fldPath, "at least one of resourceRules and nonResourceRules has to be non-empty"))
	}
	for i, resourceRule := range rule.ResourceRules {
		allErrs = append(allErrs, ValidateFlowSchemaResourcePolicyRule(&resourceRule, fldPath.Child("resourceRules").Index(i))...)
	}
	for i, nonResourceRule := range rule.NonResourceRules {
		allErrs = append(allErrs, ValidateFlowSchemaNonResourcePolicyRule(&nonResourceRule, fldPath.Child("nonResourceRules").Index(i))...)
	}
	return allErrs
}

// ValidateFlowSchemaSubject validates flow-schema's subject object.
func ValidateFlowSchemaSubject(subject *flowcontrol.Subject, fldPath *field.Path) field.ErrorList {
	var allErrs field.ErrorList
	switch subject.Kind {
	case flowcontrol.SubjectKindServiceAccount:
		allErrs = append(allErrs, ValidateServiceAccountSubject(subject.ServiceAccount, fldPath.Child("serviceAccount"))...)
	case flowcontrol.SubjectKindUser:
		allErrs = append(allErrs, ValidateUserSubject(subject.User, fldPath.Child("user"))...)
	case flowcontrol.SubjectKindGroup:
		allErrs = append(allErrs, ValidateGroupSubject(subject.Group, fldPath.Child("group"))...)
	default:
		allErrs = append(allErrs, field.NotSupported(fldPath.Child("kind"), subject.Kind, supportedSubjectKinds.List()))
	}
	return allErrs
}

// ValidateServiceAccountSubject validates subject of "ServiceAccount" kind
func ValidateServiceAccountSubject(subject *flowcontrol.ServiceAccountSubject, fldPath *field.Path) field.ErrorList {
	var allErrs field.ErrorList
	if len(subject.Name) == 0 {
		allErrs = append(allErrs, field.Required(fldPath.Child("name"), ""))
	} else if subject.Name != flowcontrol.NameAll {
		for _, msg := range validation.ValidateServiceAccountName(subject.Name, false) {
			allErrs = append(allErrs, field.Invalid(fldPath.Child("name"), subject.Name, msg))
		}
	}

	if len(subject.Namespace) > 0 {
		for _, msg := range apimachineryvalidation.ValidateNamespaceName(subject.Namespace, false) {
			allErrs = append(allErrs, field.Invalid(fldPath.Child("namespace"), subject.Namespace, msg))
		}
	} else {
		allErrs = append(allErrs, field.Required(fldPath.Child("namespace"), "must specify namespace for service account"))
	}

	return allErrs
}

// ValidateUserSubject validates subject of "User" kind
func ValidateUserSubject(subject *flowcontrol.UserSubject, fldPath *field.Path) field.ErrorList {
	var allErrs field.ErrorList
	if len(subject.Name) == 0 {
		allErrs = append(allErrs, field.Required(fldPath.Child("name"), ""))
	}
	return allErrs
}

// ValidateGroupSubject validates subject of "Group" kind
func ValidateGroupSubject(subject *flowcontrol.GroupSubject, fldPath *field.Path) field.ErrorList {
	var allErrs field.ErrorList
	if len(subject.Name) == 0 {
		allErrs = append(allErrs, field.Required(fldPath.Child("name"), ""))
	}
	return allErrs
}

// ValidateFlowSchemaNonResourcePolicyRule validates non-resource policy-rule in the flow-schema.
func ValidateFlowSchemaNonResourcePolicyRule(rule *flowcontrol.NonResourcePolicyRule, fldPath *field.Path) field.ErrorList {
	var allErrs field.ErrorList

	if len(rule.Verbs) == 0 {
		allErrs = append(allErrs, field.Required(fldPath.Child("verbs"), "verbs must contain at least one value"))
	} else if hasWildcard(rule.Verbs) {
		if len(rule.Verbs) > 1 {
			allErrs = append(allErrs, field.Invalid(fldPath.Child("verbs"), rule.Verbs, "if '*' is present, must not specify other verbs"))
		}
	} else if !supportedVerbs.IsSuperset(sets.NewString(rule.Verbs...)) {
		// only supported verbs are allowed
		allErrs = append(allErrs, field.NotSupported(fldPath.Child("verbs"), rule.Verbs, supportedVerbs.List()))
	}

	if len(rule.NonResourceURLs) == 0 {
		allErrs = append(allErrs, field.Required(fldPath.Child("nonResourceURLs"), "nonResourceURLs must contain at least one value"))
	} else if len(rule.NonResourceURLs) > 1 && hasWildcard(rule.NonResourceURLs) {
		allErrs = append(allErrs, field.Invalid(fldPath.Child("nonResourceURLs"), rule.NonResourceURLs, "if '*' is present, must not specify other non-resource URLs"))
	}

	return allErrs
}

// ValidateFlowSchemaResourcePolicyRule validates resource policy-rule in the flow-schema.
func ValidateFlowSchemaResourcePolicyRule(rule *flowcontrol.ResourcePolicyRule, fldPath *field.Path) field.ErrorList {
	var allErrs field.ErrorList

	if len(rule.Verbs) == 0 {
		allErrs = append(allErrs, field.Required(fldPath.Child("verbs"), "verbs must contain at least one value"))
	} else if hasWildcard(rule.Verbs) {
		if len(rule.Verbs) > 1 {
			allErrs = append(allErrs, field.Invalid(fldPath.Child("verbs"), rule.Verbs, "if '*' is present, must not specify other verbs"))
		}
	} else if !supportedVerbs.IsSuperset(sets.NewString(rule.Verbs...)) {
		// only supported verbs are allowed
		allErrs = append(allErrs, field.NotSupported(fldPath.Child("verbs"), rule.Verbs, supportedVerbs.List()))
	}

	if len(rule.APIGroups) == 0 {
		allErrs = append(allErrs, field.Required(fldPath.Child("apiGroups"), "resource rules must supply at least one api group"))
	} else if len(rule.APIGroups) > 1 && hasWildcard(rule.APIGroups) {
		allErrs = append(allErrs, field.Invalid(fldPath.Child("apiGroups"), rule.APIGroups, "if '*' is present, must not specify other api groups"))
	}

	if len(rule.Resources) == 0 {
		allErrs = append(allErrs, field.Required(fldPath.Child("resources"), "resource rules must supply at least one resource"))
	} else if len(rule.Resources) > 1 && hasWildcard(rule.Resources) {
		allErrs = append(allErrs, field.Invalid(fldPath.Child("resources"), rule.Resources, "if '*' is present, must not specify other resources"))
	}

	return allErrs
}

// ValidateFlowSchemaStatus validates status for the flow-schema.
func ValidateFlowSchemaStatus(status *flowcontrol.FlowSchemaStatus, fldPath *field.Path) field.ErrorList {
	var allErrs field.ErrorList
	keys := sets.NewString()
	for i, condition := range status.Conditions {
		if keys.Has(string(condition.Type)) {
			allErrs = append(allErrs, field.Duplicate(fldPath.Child("conditions").Index(i).Child("type"), condition.Type))
		}
		keys.Insert(string(condition.Type))
		allErrs = append(allErrs, ValidateFlowSchemaCondition(&condition, fldPath.Child("conditions").Index(i))...)
	}
	return allErrs
}

// ValidateFlowSchemaCondition validates condition in the flow-schema's status.
func ValidateFlowSchemaCondition(condition *flowcontrol.FlowSchemaCondition, fldPath *field.Path) field.ErrorList {
	var allErrs field.ErrorList
	if len(condition.Type) == 0 {
		allErrs = append(allErrs, field.Required(fldPath.Child("type"), "must not be empty"))
	}
	return allErrs
}

// ValidatePriorityLevelConfiguration validates priority-level-configuration.
func ValidatePriorityLevelConfiguration(pl *flowcontrol.PriorityLevelConfiguration) field.ErrorList {
	allErrs := apivalidation.ValidateObjectMeta(&pl.ObjectMeta, false, ValidatePriorityLevelConfigurationName, field.NewPath("metadata"))
	allErrs = append(allErrs, ValidatePriorityLevelConfigurationSpec(&pl.Spec, pl.Name, field.NewPath("spec"))...)
	allErrs = append(allErrs, ValidatePriorityLevelConfigurationStatus(&pl.Status, field.NewPath("status"))...)
	return allErrs
}

// ValidatePriorityLevelConfigurationSpec validates priority-level-configuration's spec.
func ValidatePriorityLevelConfigurationSpec(spec *flowcontrol.PriorityLevelConfigurationSpec, name string, fldPath *field.Path) field.ErrorList {
	var allErrs field.ErrorList
	switch spec.Type {
	case flowcontrol.PriorityLevelQueuingTypeExempt:
		if spec.Queuing != nil {
			allErrs = append(allErrs, field.Forbidden(fldPath.Child("queuing"), "must be nil if the type is not Queuing"))
		}
	case flowcontrol.PriorityLevelQueuingTypeQueueing:
		if spec.Queuing == nil {
			allErrs = append(allErrs, field.Required(fldPath.Child("queuing"), "must not be empty"))
		} else {
			allErrs = append(allErrs, ValidatePriorityLevelQueuingConfiguration(spec.Queuing, fldPath.Child("queuing"))...)
		}
	default:
		allErrs = append(allErrs, field.NotSupported(fldPath.Child("type"), spec.Type, supportedQueuingType.List()))
	}
	return allErrs
}

// ValidatePriorityLevelQueuingConfiguration validates queuing-configuration for a priority-level
func ValidatePriorityLevelQueuingConfiguration(queuing *flowcontrol.QueuingConfiguration, fldPath *field.Path) field.ErrorList {
	var allErrs field.ErrorList
	if queuing.AssuredConcurrencyShares <= 0 {
		allErrs = append(allErrs, field.Invalid(fldPath.Child("assuredConcurrencyShares"), queuing.AssuredConcurrencyShares, "must be positive"))
	}
	if queuing.QueueLengthLimit <= 0 {
		allErrs = append(allErrs, field.Invalid(fldPath.Child("queueLengthLimit"), queuing.QueueLengthLimit, "must be positive"))
	}

	// validate input arguments for shuffle-sharding
	if queuing.Queues <= 0 {
		allErrs = append(allErrs, field.Invalid(fldPath.Child("queues"), queuing.Queues, "must be positive"))
	} else if queuing.Queues > priorityLevelConfigurationQueuingMaxQueues {
		allErrs = append(allErrs, field.Invalid(fldPath.Child("queues"), queuing.Queues,
			fmt.Sprintf("must not be greater than %d", priorityLevelConfigurationQueuingMaxQueues)))
	}

	if queuing.HandSize <= 0 {
		allErrs = append(allErrs, field.Invalid(fldPath.Child("handSize"), queuing.HandSize, "must be positive"))
	} else if queuing.HandSize > queuing.Queues {
		allErrs = append(allErrs, field.Invalid(fldPath.Child("handSize"), queuing.HandSize,
			fmt.Sprintf("should not be greater than queues (%d)", queuing.Queues)))
	} else if entropy := shufflesharding.RequiredEntropyBits(int(queuing.Queues), int(queuing.HandSize)); entropy > shufflesharding.MaxHashBits {
		allErrs = append(allErrs, field.Invalid(fldPath.Child("handSize"), queuing.HandSize,
			fmt.Sprintf("required entropy bits of deckSize %d and handSize %d should not be greater than %d", queuing.Queues, queuing.HandSize, shufflesharding.MaxHashBits)))
	}
	return allErrs
}

// ValidatePriorityLevelConfigurationStatus validates priority-level-configuration's status.
func ValidatePriorityLevelConfigurationStatus(status *flowcontrol.PriorityLevelConfigurationStatus, fldPath *field.Path) field.ErrorList {
	var allErrs field.ErrorList
	keys := sets.NewString()
	for i, condition := range status.Conditions {
		if keys.Has(string(condition.Type)) {
			allErrs = append(allErrs, field.Duplicate(fldPath.Child("conditions").Index(i).Child("type"), condition.Type))
		}
		keys.Insert(string(condition.Type))
		allErrs = append(allErrs, ValidatePriorityLevelConfigurationCondition(&condition, fldPath.Child("conditions").Index(i))...)
	}
	return allErrs
}

// ValidatePriorityLevelConfigurationCondition validates condition in priority-level-configuration's status.
func ValidatePriorityLevelConfigurationCondition(condition *flowcontrol.PriorityLevelConfigurationCondition, fldPath *field.Path) field.ErrorList {
	var allErrs field.ErrorList
	if len(condition.Type) == 0 {
		allErrs = append(allErrs, field.Required(fldPath.Child("type"), "must not be empty"))
	}
	return allErrs
}

func hasWildcard(operations []string) bool {
	for _, o := range operations {
		if o == "*" {
			return true
		}
	}
	return false
}
