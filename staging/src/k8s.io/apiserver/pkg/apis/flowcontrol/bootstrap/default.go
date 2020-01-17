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

package bootstrap

import (
	coordinationv1 "k8s.io/api/coordination/v1"
	corev1 "k8s.io/api/core/v1"
	flowcontrol "k8s.io/api/flowcontrol/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apiserver/pkg/authentication/serviceaccount"
	"k8s.io/apiserver/pkg/authentication/user"
)

// The objects that define an apiserver's initial behavior.  The
// registered defaulting procedures make no changes to these
// particular objects (this is verified in the unit tests of the
// internalbootstrap package; it can not be verified in this package
// because that would require importing k8s.io/kubernetes).
var (
	MandatoryPriorityLevelConfigurations = []*flowcontrol.PriorityLevelConfiguration{
		MandatoryPriorityLevelConfigurationExempt,
		MandatoryPriorityLevelConfigurationCatchAll,
	}
	MandatoryFlowSchemas = []*flowcontrol.FlowSchema{
		MandatoryFlowSchemaExempt,
		MandatoryFlowSchemaCatchAll,
	}
)

// The objects that define the current suggested additional configuration
var (
	SuggestedPriorityLevelConfigurations = []*flowcontrol.PriorityLevelConfiguration{
		// "system" priority-level is for the system components that affects self-maintenance of the
		// cluster and the availability of those running pods in the cluster, including kubelet and
		// kube-proxy.
		SuggestedPriorityLevelConfigurationSystem,
		// "leader-election" is dedicated for controllers' leader-election, which majorly affects the
		// availability of any controller runs in the cluster.
		SuggestedPriorityLevelConfigurationLeaderElection,
		// "workload-high" is used by those workloads with higher priority but their failure won't directly
		// impact the existing running pods in the cluster, which includes kube-scheduler, and those well-known
		// built-in workloads such as "deployments", "replicasets" and other low-level custom workload which
		// is important for the cluster.
		SuggestedPriorityLevelConfigurationWorkloadHigh,
		// "workload-low" is used by those workloads with lower priority which availability only has a
		// minor impact on the cluster.
		SuggestedPriorityLevelConfigurationWorkloadLow,
	}
	SuggestedFlowSchemas = []*flowcontrol.FlowSchema{
		SuggestedFlowSchemaSystemNodes,               // references "system" priority-level
		SuggestedFlowSchemaSystemLeaderElection,      // references "leader-election" priority-level
		SuggestedFlowSchemaWorkloadLeaderElection,    // references "leader-election" priority-level
		SuggestedFlowSchemaKubeControllerManager,     // references "workload-high" priority-level
		SuggestedFlowSchemaKubeScheduler,             // references "workload-high" priority-level
		SuggestedFlowSchemaKubeSystemServiceAccounts, // references "workload-high" priority-level
		SuggestedFlowSchemaServiceAccounts,           // references "workload-low" priority-level
	}
)

// Mandatory PriorityLevelConfiguration objects
var (
	MandatoryPriorityLevelConfigurationExempt = newPriorityLevelConfiguration(
		flowcontrol.PriorityLevelConfigurationNameExempt,
		flowcontrol.PriorityLevelConfigurationSpec{
			Type: flowcontrol.PriorityLevelEnablementExempt,
		},
	)
	MandatoryPriorityLevelConfigurationCatchAll = newPriorityLevelConfiguration(
		"catch-all",
		flowcontrol.PriorityLevelConfigurationSpec{
			Type: flowcontrol.PriorityLevelEnablementLimited,
			Limited: &flowcontrol.LimitedPriorityLevelConfiguration{
				AssuredConcurrencyShares: 100,
				LimitResponse: flowcontrol.LimitResponse{
					Type: flowcontrol.LimitResponseTypeQueue,
					Queuing: &flowcontrol.QueuingConfiguration{
						Queues:           128,
						HandSize:         6,
						QueueLengthLimit: 100,
					},
				},
			},
		})
)

// Mandatory FlowSchema objects
var (
	// exempt priority-level
	MandatoryFlowSchemaExempt = newFlowSchema(
		"exempt",
		flowcontrol.PriorityLevelConfigurationNameExempt,
		1,  // matchingPrecedence
		"", // distinguisherMethodType
		flowcontrol.PolicyRulesWithSubjects{
			Subjects: groups(user.SystemPrivilegedGroup),
			ResourceRules: []flowcontrol.ResourcePolicyRule{
				resourceRule(
					[]string{flowcontrol.VerbAll},
					[]string{flowcontrol.APIGroupAll},
					[]string{flowcontrol.ResourceAll},
					[]string{flowcontrol.NamespaceEvery},
					true,
				),
			},
			NonResourceRules: []flowcontrol.NonResourcePolicyRule{
				nonResourceRule(
					[]string{flowcontrol.VerbAll},
					[]string{flowcontrol.NonResourceAll},
				),
			},
		},
	)
	// catch-all priority-level
	MandatoryFlowSchemaCatchAll = newFlowSchema(
		"catch-all",
		"catch-all",
		10000, // matchingPrecedence
		flowcontrol.FlowDistinguisherMethodByUserType, // distinguisherMethodType
		flowcontrol.PolicyRulesWithSubjects{
			Subjects: groups(user.AllUnauthenticated, user.AllAuthenticated),
			ResourceRules: []flowcontrol.ResourcePolicyRule{
				resourceRule(
					[]string{flowcontrol.VerbAll},
					[]string{flowcontrol.APIGroupAll},
					[]string{flowcontrol.ResourceAll},
					[]string{flowcontrol.NamespaceEvery},
					true,
				),
			},
			NonResourceRules: []flowcontrol.NonResourcePolicyRule{
				nonResourceRule(
					[]string{flowcontrol.VerbAll},
					[]string{flowcontrol.NonResourceAll},
				),
			},
		},
	)
)

// Suggested PriorityLevelConfiguration objects
var (
	// system priority-level
	SuggestedPriorityLevelConfigurationSystem = newPriorityLevelConfiguration(
		"system",
		flowcontrol.PriorityLevelConfigurationSpec{
			Type: flowcontrol.PriorityLevelEnablementLimited,
			Limited: &flowcontrol.LimitedPriorityLevelConfiguration{
				AssuredConcurrencyShares: 30,
				LimitResponse: flowcontrol.LimitResponse{
					Type: flowcontrol.LimitResponseTypeQueue,
					Queuing: &flowcontrol.QueuingConfiguration{
						Queues:           64,
						HandSize:         6,
						QueueLengthLimit: 1000,
					},
				},
			},
		})
	// leader-election priority-level
	SuggestedPriorityLevelConfigurationLeaderElection = newPriorityLevelConfiguration(
		"leader-election",
		flowcontrol.PriorityLevelConfigurationSpec{
			Type: flowcontrol.PriorityLevelEnablementLimited,
			Limited: &flowcontrol.LimitedPriorityLevelConfiguration{
				AssuredConcurrencyShares: 10,
				LimitResponse: flowcontrol.LimitResponse{
					Type: flowcontrol.LimitResponseTypeQueue,
					Queuing: &flowcontrol.QueuingConfiguration{
						Queues:           16,
						HandSize:         4,
						QueueLengthLimit: 100,
					},
				},
			},
		})
	// workload-high priority-level
	SuggestedPriorityLevelConfigurationWorkloadHigh = newPriorityLevelConfiguration(
		"workload-high",
		flowcontrol.PriorityLevelConfigurationSpec{
			Type: flowcontrol.PriorityLevelEnablementLimited,
			Limited: &flowcontrol.LimitedPriorityLevelConfiguration{
				AssuredConcurrencyShares: 40,
				LimitResponse: flowcontrol.LimitResponse{
					Type: flowcontrol.LimitResponseTypeQueue,
					Queuing: &flowcontrol.QueuingConfiguration{
						Queues:           128,
						HandSize:         6,
						QueueLengthLimit: 100,
					},
				},
			},
		})
	// workload-low priority-level
	SuggestedPriorityLevelConfigurationWorkloadLow = newPriorityLevelConfiguration(
		"workload-low",
		flowcontrol.PriorityLevelConfigurationSpec{
			Type: flowcontrol.PriorityLevelEnablementLimited,
			Limited: &flowcontrol.LimitedPriorityLevelConfiguration{
				AssuredConcurrencyShares: 20,
				LimitResponse: flowcontrol.LimitResponse{
					Type: flowcontrol.LimitResponseTypeQueue,
					Queuing: &flowcontrol.QueuingConfiguration{
						Queues:           128,
						HandSize:         6,
						QueueLengthLimit: 100,
					},
				},
			},
		})
)

// Suggested FlowSchema objects
var (
	SuggestedFlowSchemaSystemNodes = newFlowSchema(
		"system-nodes", "system", 500,
		flowcontrol.FlowDistinguisherMethodByUserType,
		flowcontrol.PolicyRulesWithSubjects{
			Subjects: groups(user.NodesGroup), // the nodes group
			ResourceRules: []flowcontrol.ResourcePolicyRule{resourceRule(
				[]string{flowcontrol.VerbAll},
				[]string{flowcontrol.APIGroupAll},
				[]string{flowcontrol.ResourceAll},
				[]string{flowcontrol.NamespaceEvery},
				true)},
			NonResourceRules: []flowcontrol.NonResourcePolicyRule{
				nonResourceRule(
					[]string{flowcontrol.VerbAll},
					[]string{flowcontrol.NonResourceAll}),
			},
		},
	)
	SuggestedFlowSchemaSystemLeaderElection = newFlowSchema(
		"system-leader-election", "leader-election", 100,
		flowcontrol.FlowDistinguisherMethodByUserType,
		flowcontrol.PolicyRulesWithSubjects{
			Subjects: append(
				users(user.KubeControllerManager, user.KubeScheduler),
				kubeSystemServiceAccount(flowcontrol.NameAll)...),
			ResourceRules: []flowcontrol.ResourcePolicyRule{
				resourceRule(
					[]string{"get", "create", "update"},
					[]string{corev1.GroupName},
					[]string{"endpoints", "configmaps"},
					[]string{"kube-system"},
					false),
				resourceRule(
					[]string{"get", "create", "update"},
					[]string{coordinationv1.GroupName},
					[]string{"leases"},
					[]string{flowcontrol.NamespaceEvery},
					false),
			},
		},
	)
	SuggestedFlowSchemaWorkloadLeaderElection = newFlowSchema(
		"workload-leader-election", "leader-election", 200,
		flowcontrol.FlowDistinguisherMethodByUserType,
		flowcontrol.PolicyRulesWithSubjects{
			Subjects: kubeSystemServiceAccount(flowcontrol.NameAll),
			ResourceRules: []flowcontrol.ResourcePolicyRule{
				resourceRule(
					[]string{"get", "create", "update"},
					[]string{corev1.GroupName},
					[]string{"endpoints", "configmaps"},
					[]string{flowcontrol.NamespaceEvery},
					false),
				resourceRule(
					[]string{"get", "create", "update"},
					[]string{coordinationv1.GroupName},
					[]string{"leases"},
					[]string{flowcontrol.NamespaceEvery},
					false),
			},
		},
	)
	SuggestedFlowSchemaKubeControllerManager = newFlowSchema(
		"kube-controller-manager", "workload-high", 800,
		flowcontrol.FlowDistinguisherMethodByNamespaceType,
		flowcontrol.PolicyRulesWithSubjects{
			Subjects: users(user.KubeControllerManager),
			ResourceRules: []flowcontrol.ResourcePolicyRule{resourceRule(
				[]string{flowcontrol.VerbAll},
				[]string{flowcontrol.APIGroupAll},
				[]string{flowcontrol.ResourceAll},
				[]string{flowcontrol.NamespaceEvery},
				true)},
			NonResourceRules: []flowcontrol.NonResourcePolicyRule{
				nonResourceRule(
					[]string{flowcontrol.VerbAll},
					[]string{flowcontrol.NonResourceAll}),
			},
		},
	)
	SuggestedFlowSchemaKubeScheduler = newFlowSchema(
		"kube-scheduler", "workload-high", 800,
		flowcontrol.FlowDistinguisherMethodByNamespaceType,
		flowcontrol.PolicyRulesWithSubjects{
			Subjects: users(user.KubeScheduler),
			ResourceRules: []flowcontrol.ResourcePolicyRule{resourceRule(
				[]string{flowcontrol.VerbAll},
				[]string{flowcontrol.APIGroupAll},
				[]string{flowcontrol.ResourceAll},
				[]string{flowcontrol.NamespaceEvery},
				true)},
			NonResourceRules: []flowcontrol.NonResourcePolicyRule{
				nonResourceRule(
					[]string{flowcontrol.VerbAll},
					[]string{flowcontrol.NonResourceAll}),
			},
		},
	)
	SuggestedFlowSchemaKubeSystemServiceAccounts = newFlowSchema(
		"kube-system-service-accounts", "workload-high", 900,
		flowcontrol.FlowDistinguisherMethodByNamespaceType,
		flowcontrol.PolicyRulesWithSubjects{
			Subjects: kubeSystemServiceAccount(flowcontrol.NameAll),
			ResourceRules: []flowcontrol.ResourcePolicyRule{resourceRule(
				[]string{flowcontrol.VerbAll},
				[]string{flowcontrol.APIGroupAll},
				[]string{flowcontrol.ResourceAll},
				[]string{flowcontrol.NamespaceEvery},
				true)},
			NonResourceRules: []flowcontrol.NonResourcePolicyRule{
				nonResourceRule(
					[]string{flowcontrol.VerbAll},
					[]string{flowcontrol.NonResourceAll}),
			},
		},
	)
	SuggestedFlowSchemaServiceAccounts = newFlowSchema(
		"service-accounts", "workload-low", 9000,
		flowcontrol.FlowDistinguisherMethodByUserType,
		flowcontrol.PolicyRulesWithSubjects{
			Subjects: groups(serviceaccount.AllServiceAccountsGroup),
			ResourceRules: []flowcontrol.ResourcePolicyRule{resourceRule(
				[]string{flowcontrol.VerbAll},
				[]string{flowcontrol.APIGroupAll},
				[]string{flowcontrol.ResourceAll},
				[]string{flowcontrol.NamespaceEvery},
				true)},
			NonResourceRules: []flowcontrol.NonResourcePolicyRule{
				nonResourceRule(
					[]string{flowcontrol.VerbAll},
					[]string{flowcontrol.NonResourceAll}),
			},
		},
	)
)

func newPriorityLevelConfiguration(name string, spec flowcontrol.PriorityLevelConfigurationSpec) *flowcontrol.PriorityLevelConfiguration {
	return &flowcontrol.PriorityLevelConfiguration{
		ObjectMeta: metav1.ObjectMeta{Name: name},
		Spec:       spec}
}

func newFlowSchema(name, plName string, matchingPrecedence int32, dmType flowcontrol.FlowDistinguisherMethodType, rules ...flowcontrol.PolicyRulesWithSubjects) *flowcontrol.FlowSchema {
	var dm *flowcontrol.FlowDistinguisherMethod
	if dmType != "" {
		dm = &flowcontrol.FlowDistinguisherMethod{Type: dmType}
	}
	return &flowcontrol.FlowSchema{
		ObjectMeta: metav1.ObjectMeta{Name: name},
		Spec: flowcontrol.FlowSchemaSpec{
			PriorityLevelConfiguration: flowcontrol.PriorityLevelConfigurationReference{
				Name: plName,
			},
			MatchingPrecedence:  matchingPrecedence,
			DistinguisherMethod: dm,
			Rules:               rules},
	}

}

func groups(names ...string) []flowcontrol.Subject {
	ans := make([]flowcontrol.Subject, len(names))
	for idx, name := range names {
		ans[idx] = flowcontrol.Subject{
			Kind: flowcontrol.SubjectKindGroup,
			Group: &flowcontrol.GroupSubject{
				Name: name,
			},
		}
	}
	return ans
}

func users(names ...string) []flowcontrol.Subject {
	ans := make([]flowcontrol.Subject, len(names))
	for idx, name := range names {
		ans[idx] = flowcontrol.Subject{
			Kind: flowcontrol.SubjectKindUser,
			User: &flowcontrol.UserSubject{
				Name: name,
			},
		}
	}
	return ans
}

func kubeSystemServiceAccount(names ...string) []flowcontrol.Subject {
	subjects := []flowcontrol.Subject{}
	for _, name := range names {
		subjects = append(subjects, flowcontrol.Subject{
			Kind: flowcontrol.SubjectKindServiceAccount,
			ServiceAccount: &flowcontrol.ServiceAccountSubject{
				Name:      name,
				Namespace: metav1.NamespaceSystem,
			},
		})
	}
	return subjects
}

func resourceRule(verbs []string, groups []string, resources []string, namespaces []string, clusterScoped bool) flowcontrol.ResourcePolicyRule {
	return flowcontrol.ResourcePolicyRule{
		Verbs:        verbs,
		APIGroups:    groups,
		Resources:    resources,
		Namespaces:   namespaces,
		ClusterScope: clusterScoped,
	}
}

func nonResourceRule(verbs []string, nonResourceURLs []string) flowcontrol.NonResourcePolicyRule {
	return flowcontrol.NonResourcePolicyRule{Verbs: verbs, NonResourceURLs: nonResourceURLs}
}
