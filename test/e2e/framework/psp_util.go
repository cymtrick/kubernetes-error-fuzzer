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

package framework

import (
	"fmt"
	"sync"

	corev1 "k8s.io/api/core/v1"
	policy "k8s.io/api/policy/v1beta1"
	rbacv1beta1 "k8s.io/api/rbac/v1beta1"
	apierrs "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apiserver/pkg/authentication/serviceaccount"
	"k8s.io/kubernetes/pkg/security/podsecuritypolicy/seccomp"
	"k8s.io/kubernetes/test/e2e/framework/auth"

	"github.com/onsi/ginkgo"
)

const (
	podSecurityPolicyPrivileged = "e2e-test-privileged-psp"
)

var (
	isPSPEnabledOnce sync.Once
	isPSPEnabled     bool
)

// privilegedPSP creates a PodSecurityPolicy that allows everything.
func privilegedPSP(name string) *policy.PodSecurityPolicy {
	allowPrivilegeEscalation := true
	return &policy.PodSecurityPolicy{
		ObjectMeta: metav1.ObjectMeta{
			Name:        name,
			Annotations: map[string]string{seccomp.AllowedProfilesAnnotationKey: seccomp.AllowAny},
		},
		Spec: policy.PodSecurityPolicySpec{
			Privileged:               true,
			AllowPrivilegeEscalation: &allowPrivilegeEscalation,
			AllowedCapabilities:      []corev1.Capability{"*"},
			Volumes:                  []policy.FSType{policy.All},
			HostNetwork:              true,
			HostPorts:                []policy.HostPortRange{{Min: 0, Max: 65535}},
			HostIPC:                  true,
			HostPID:                  true,
			RunAsUser: policy.RunAsUserStrategyOptions{
				Rule: policy.RunAsUserStrategyRunAsAny,
			},
			SELinux: policy.SELinuxStrategyOptions{
				Rule: policy.SELinuxStrategyRunAsAny,
			},
			SupplementalGroups: policy.SupplementalGroupsStrategyOptions{
				Rule: policy.SupplementalGroupsStrategyRunAsAny,
			},
			FSGroup: policy.FSGroupStrategyOptions{
				Rule: policy.FSGroupStrategyRunAsAny,
			},
			ReadOnlyRootFilesystem: false,
			AllowedUnsafeSysctls:   []string{"*"},
		},
	}
}

// IsPodSecurityPolicyEnabled returns true if PodSecurityPolicy is enabled. Otherwise false.
func IsPodSecurityPolicyEnabled(f *Framework) bool {
	isPSPEnabledOnce.Do(func() {
		psps, err := f.ClientSet.PolicyV1beta1().PodSecurityPolicies().List(metav1.ListOptions{})
		if err != nil {
			Logf("Error listing PodSecurityPolicies; assuming PodSecurityPolicy is disabled: %v", err)
			isPSPEnabled = false
		} else if psps == nil || len(psps.Items) == 0 {
			Logf("No PodSecurityPolicies found; assuming PodSecurityPolicy is disabled.")
			isPSPEnabled = false
		} else {
			Logf("Found PodSecurityPolicies; assuming PodSecurityPolicy is enabled.")
			isPSPEnabled = true
		}
	})
	return isPSPEnabled
}

var (
	privilegedPSPOnce sync.Once
)

func createPrivilegedPSPBinding(f *Framework, namespace string) {
	if !IsPodSecurityPolicyEnabled(f) {
		return
	}
	// Create the privileged PSP & role
	privilegedPSPOnce.Do(func() {
		_, err := f.ClientSet.PolicyV1beta1().PodSecurityPolicies().Get(
			podSecurityPolicyPrivileged, metav1.GetOptions{})
		if !apierrs.IsNotFound(err) {
			// Privileged PSP was already created.
			ExpectNoError(err, "Failed to get PodSecurityPolicy %s", podSecurityPolicyPrivileged)
			return
		}

		psp := privilegedPSP(podSecurityPolicyPrivileged)
		psp, err = f.ClientSet.PolicyV1beta1().PodSecurityPolicies().Create(psp)
		if !apierrs.IsAlreadyExists(err) {
			ExpectNoError(err, "Failed to create PSP %s", podSecurityPolicyPrivileged)
		}

		if auth.IsRBACEnabled(f.ClientSet.RbacV1beta1()) {
			// Create the Role to bind it to the namespace.
			_, err = f.ClientSet.RbacV1beta1().ClusterRoles().Create(&rbacv1beta1.ClusterRole{
				ObjectMeta: metav1.ObjectMeta{Name: podSecurityPolicyPrivileged},
				Rules: []rbacv1beta1.PolicyRule{{
					APIGroups:     []string{"extensions"},
					Resources:     []string{"podsecuritypolicies"},
					ResourceNames: []string{podSecurityPolicyPrivileged},
					Verbs:         []string{"use"},
				}},
			})
			if !apierrs.IsAlreadyExists(err) {
				ExpectNoError(err, "Failed to create PSP role")
			}
		}
	})

	if auth.IsRBACEnabled(f.ClientSet.RbacV1beta1()) {
		ginkgo.By(fmt.Sprintf("Binding the %s PodSecurityPolicy to the default service account in %s",
			podSecurityPolicyPrivileged, namespace))
		err := auth.BindClusterRoleInNamespace(f.ClientSet.RbacV1beta1(),
			podSecurityPolicyPrivileged,
			namespace,
			rbacv1beta1.Subject{
				Kind:      rbacv1beta1.ServiceAccountKind,
				Namespace: namespace,
				Name:      "default",
			})
		ExpectNoError(err)
		ExpectNoError(auth.WaitForNamedAuthorizationUpdate(f.ClientSet.AuthorizationV1beta1(),
			serviceaccount.MakeUsername(namespace, "default"), namespace, "use", podSecurityPolicyPrivileged,
			schema.GroupResource{Group: "extensions", Resource: "podsecuritypolicies"}, true))
	}
}
