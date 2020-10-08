# Copyright The Kubernetes Authors.
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#     http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

# #################################################
# # # # # # # # # # # # # # # # # # # # # # # # # #
# This file is autogenerated by kazel. DO NOT EDIT.
# # # # # # # # # # # # # # # # # # # # # # # # # #
# #################################################
#
# The go prefix passed to kazel
go_prefix = "k8s.io/kubernetes"

# The list of codegen tags kazel is configured to find
kazel_configured_tags = ["openapi-gen"]

# tags_values_pkgs is a dictionary mapping {k8s build tag: {tag value: [pkgs including that tag:value]}}
tags_values_pkgs = {"openapi-gen": {
    "false": [
        "staging/src/k8s.io/api/admission/v1",
        "staging/src/k8s.io/api/admission/v1beta1",
        "staging/src/k8s.io/api/core/v1",
        "staging/src/k8s.io/apimachinery/pkg/apis/testapigroup/v1",
        "staging/src/k8s.io/apiserver/pkg/apis/example/v1",
        "staging/src/k8s.io/apiserver/pkg/apis/example2/v1",
    ],
    "true": [
        "pkg/apis/abac/v0",
        "pkg/apis/abac/v1beta1",
        "staging/src/k8s.io/api/admissionregistration/v1",
        "staging/src/k8s.io/api/admissionregistration/v1beta1",
        "staging/src/k8s.io/api/apiserverinternal/v1alpha1",
        "staging/src/k8s.io/api/apps/v1",
        "staging/src/k8s.io/api/apps/v1beta1",
        "staging/src/k8s.io/api/apps/v1beta2",
        "staging/src/k8s.io/api/authentication/v1",
        "staging/src/k8s.io/api/authentication/v1beta1",
        "staging/src/k8s.io/api/authorization/v1",
        "staging/src/k8s.io/api/authorization/v1beta1",
        "staging/src/k8s.io/api/autoscaling/v1",
        "staging/src/k8s.io/api/autoscaling/v2beta1",
        "staging/src/k8s.io/api/autoscaling/v2beta2",
        "staging/src/k8s.io/api/batch/v1",
        "staging/src/k8s.io/api/batch/v1beta1",
        "staging/src/k8s.io/api/batch/v2alpha1",
        "staging/src/k8s.io/api/certificates/v1",
        "staging/src/k8s.io/api/certificates/v1beta1",
        "staging/src/k8s.io/api/coordination/v1",
        "staging/src/k8s.io/api/coordination/v1beta1",
        "staging/src/k8s.io/api/core/v1",
        "staging/src/k8s.io/api/discovery/v1alpha1",
        "staging/src/k8s.io/api/discovery/v1beta1",
        "staging/src/k8s.io/api/events/v1",
        "staging/src/k8s.io/api/events/v1beta1",
        "staging/src/k8s.io/api/extensions/v1beta1",
        "staging/src/k8s.io/api/flowcontrol/v1alpha1",
        "staging/src/k8s.io/api/imagepolicy/v1alpha1",
        "staging/src/k8s.io/api/networking/v1",
        "staging/src/k8s.io/api/networking/v1beta1",
        "staging/src/k8s.io/api/node/v1alpha1",
        "staging/src/k8s.io/api/node/v1beta1",
        "staging/src/k8s.io/api/policy/v1beta1",
        "staging/src/k8s.io/api/rbac/v1",
        "staging/src/k8s.io/api/rbac/v1alpha1",
        "staging/src/k8s.io/api/rbac/v1beta1",
        "staging/src/k8s.io/api/scheduling/v1",
        "staging/src/k8s.io/api/scheduling/v1alpha1",
        "staging/src/k8s.io/api/scheduling/v1beta1",
        "staging/src/k8s.io/api/storage/v1",
        "staging/src/k8s.io/api/storage/v1alpha1",
        "staging/src/k8s.io/api/storage/v1beta1",
        "staging/src/k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1",
        "staging/src/k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1beta1",
        "staging/src/k8s.io/apimachinery/pkg/api/resource",
        "staging/src/k8s.io/apimachinery/pkg/apis/meta/v1",
        "staging/src/k8s.io/apimachinery/pkg/apis/meta/v1beta1",
        "staging/src/k8s.io/apimachinery/pkg/runtime",
        "staging/src/k8s.io/apimachinery/pkg/util/intstr",
        "staging/src/k8s.io/apimachinery/pkg/version",
        "staging/src/k8s.io/apiserver/pkg/apis/audit/v1",
        "staging/src/k8s.io/apiserver/pkg/apis/audit/v1alpha1",
        "staging/src/k8s.io/apiserver/pkg/apis/audit/v1beta1",
        "staging/src/k8s.io/client-go/pkg/apis/clientauthentication/v1alpha1",
        "staging/src/k8s.io/client-go/pkg/apis/clientauthentication/v1beta1",
        "staging/src/k8s.io/client-go/pkg/version",
        "staging/src/k8s.io/cloud-provider/app/apis/config/v1alpha1",
        "staging/src/k8s.io/kube-aggregator/pkg/apis/apiregistration/v1",
        "staging/src/k8s.io/kube-aggregator/pkg/apis/apiregistration/v1beta1",
        "staging/src/k8s.io/kube-controller-manager/config/v1alpha1",
        "staging/src/k8s.io/kube-proxy/config/v1alpha1",
        "staging/src/k8s.io/kube-scheduler/config/v1",
        "staging/src/k8s.io/kube-scheduler/config/v1beta1",
        "staging/src/k8s.io/kubelet/config/v1beta1",
        "staging/src/k8s.io/metrics/pkg/apis/custom_metrics/v1beta1",
        "staging/src/k8s.io/metrics/pkg/apis/custom_metrics/v1beta2",
        "staging/src/k8s.io/metrics/pkg/apis/external_metrics/v1beta1",
        "staging/src/k8s.io/metrics/pkg/apis/metrics/v1alpha1",
        "staging/src/k8s.io/metrics/pkg/apis/metrics/v1beta1",
        "staging/src/k8s.io/sample-apiserver/pkg/apis/wardle/v1alpha1",
        "staging/src/k8s.io/sample-apiserver/pkg/apis/wardle/v1beta1",
    ],
}}

# tags_pkgs_values is a dictionary mapping {k8s build tag: {pkg: [tag values in pkg]}}
tags_pkgs_values = {"openapi-gen": {
    "pkg/apis/abac/v0": ["true"],
    "pkg/apis/abac/v1beta1": ["true"],
    "staging/src/k8s.io/api/admission/v1": ["false"],
    "staging/src/k8s.io/api/admission/v1beta1": ["false"],
    "staging/src/k8s.io/api/admissionregistration/v1": ["true"],
    "staging/src/k8s.io/api/admissionregistration/v1beta1": ["true"],
    "staging/src/k8s.io/api/apiserverinternal/v1alpha1": ["true"],
    "staging/src/k8s.io/api/apps/v1": ["true"],
    "staging/src/k8s.io/api/apps/v1beta1": ["true"],
    "staging/src/k8s.io/api/apps/v1beta2": ["true"],
    "staging/src/k8s.io/api/authentication/v1": ["true"],
    "staging/src/k8s.io/api/authentication/v1beta1": ["true"],
    "staging/src/k8s.io/api/authorization/v1": ["true"],
    "staging/src/k8s.io/api/authorization/v1beta1": ["true"],
    "staging/src/k8s.io/api/autoscaling/v1": ["true"],
    "staging/src/k8s.io/api/autoscaling/v2beta1": ["true"],
    "staging/src/k8s.io/api/autoscaling/v2beta2": ["true"],
    "staging/src/k8s.io/api/batch/v1": ["true"],
    "staging/src/k8s.io/api/batch/v1beta1": ["true"],
    "staging/src/k8s.io/api/batch/v2alpha1": ["true"],
    "staging/src/k8s.io/api/certificates/v1": ["true"],
    "staging/src/k8s.io/api/certificates/v1beta1": ["true"],
    "staging/src/k8s.io/api/coordination/v1": ["true"],
    "staging/src/k8s.io/api/coordination/v1beta1": ["true"],
    "staging/src/k8s.io/api/core/v1": [
        "false",
        "true",
    ],
    "staging/src/k8s.io/api/discovery/v1alpha1": ["true"],
    "staging/src/k8s.io/api/discovery/v1beta1": ["true"],
    "staging/src/k8s.io/api/events/v1": ["true"],
    "staging/src/k8s.io/api/events/v1beta1": ["true"],
    "staging/src/k8s.io/api/extensions/v1beta1": ["true"],
    "staging/src/k8s.io/api/flowcontrol/v1alpha1": ["true"],
    "staging/src/k8s.io/api/imagepolicy/v1alpha1": ["true"],
    "staging/src/k8s.io/api/networking/v1": ["true"],
    "staging/src/k8s.io/api/networking/v1beta1": ["true"],
    "staging/src/k8s.io/api/node/v1alpha1": ["true"],
    "staging/src/k8s.io/api/node/v1beta1": ["true"],
    "staging/src/k8s.io/api/policy/v1beta1": ["true"],
    "staging/src/k8s.io/api/rbac/v1": ["true"],
    "staging/src/k8s.io/api/rbac/v1alpha1": ["true"],
    "staging/src/k8s.io/api/rbac/v1beta1": ["true"],
    "staging/src/k8s.io/api/scheduling/v1": ["true"],
    "staging/src/k8s.io/api/scheduling/v1alpha1": ["true"],
    "staging/src/k8s.io/api/scheduling/v1beta1": ["true"],
    "staging/src/k8s.io/api/storage/v1": ["true"],
    "staging/src/k8s.io/api/storage/v1alpha1": ["true"],
    "staging/src/k8s.io/api/storage/v1beta1": ["true"],
    "staging/src/k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1": ["true"],
    "staging/src/k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1beta1": ["true"],
    "staging/src/k8s.io/apimachinery/pkg/api/resource": ["true"],
    "staging/src/k8s.io/apimachinery/pkg/apis/meta/v1": ["true"],
    "staging/src/k8s.io/apimachinery/pkg/apis/meta/v1beta1": ["true"],
    "staging/src/k8s.io/apimachinery/pkg/apis/testapigroup/v1": ["false"],
    "staging/src/k8s.io/apimachinery/pkg/runtime": ["true"],
    "staging/src/k8s.io/apimachinery/pkg/util/intstr": ["true"],
    "staging/src/k8s.io/apimachinery/pkg/version": ["true"],
    "staging/src/k8s.io/apiserver/pkg/apis/audit/v1": ["true"],
    "staging/src/k8s.io/apiserver/pkg/apis/audit/v1alpha1": ["true"],
    "staging/src/k8s.io/apiserver/pkg/apis/audit/v1beta1": ["true"],
    "staging/src/k8s.io/apiserver/pkg/apis/example/v1": ["false"],
    "staging/src/k8s.io/apiserver/pkg/apis/example2/v1": ["false"],
    "staging/src/k8s.io/client-go/pkg/apis/clientauthentication/v1alpha1": ["true"],
    "staging/src/k8s.io/client-go/pkg/apis/clientauthentication/v1beta1": ["true"],
    "staging/src/k8s.io/client-go/pkg/version": ["true"],
    "staging/src/k8s.io/cloud-provider/app/apis/config/v1alpha1": ["true"],
    "staging/src/k8s.io/kube-aggregator/pkg/apis/apiregistration/v1": ["true"],
    "staging/src/k8s.io/kube-aggregator/pkg/apis/apiregistration/v1beta1": ["true"],
    "staging/src/k8s.io/kube-controller-manager/config/v1alpha1": ["true"],
    "staging/src/k8s.io/kube-proxy/config/v1alpha1": ["true"],
    "staging/src/k8s.io/kube-scheduler/config/v1": ["true"],
    "staging/src/k8s.io/kube-scheduler/config/v1beta1": ["true"],
    "staging/src/k8s.io/kubelet/config/v1beta1": ["true"],
    "staging/src/k8s.io/metrics/pkg/apis/custom_metrics/v1beta1": ["true"],
    "staging/src/k8s.io/metrics/pkg/apis/custom_metrics/v1beta2": ["true"],
    "staging/src/k8s.io/metrics/pkg/apis/external_metrics/v1beta1": ["true"],
    "staging/src/k8s.io/metrics/pkg/apis/metrics/v1alpha1": ["true"],
    "staging/src/k8s.io/metrics/pkg/apis/metrics/v1beta1": ["true"],
    "staging/src/k8s.io/sample-apiserver/pkg/apis/wardle/v1alpha1": ["true"],
    "staging/src/k8s.io/sample-apiserver/pkg/apis/wardle/v1beta1": ["true"],
}}
