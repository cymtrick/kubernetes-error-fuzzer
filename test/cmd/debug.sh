#!/usr/bin/env bash

# Copyright 2020 The Kubernetes Authors.
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

set -o errexit
set -o nounset
set -o pipefail

run_kubectl_debug_pod_tests() {
  set -o nounset
  set -o errexit

  create_and_use_new_namespace
  kube::log::status "Testing kubectl debug (pod tests)"

  ### Pod Troubleshooting by ephemeral containers

  # Pre-Condition: Pod "nginx" is created
  kubectl run target "--image=${IMAGE_NGINX:?}" "${kube_flags[@]:?}"
  kube::test::get_object_assert pod "{{range.items}}{{${id_field:?}}}:{{end}}" 'target:'
  # Command: create a copy of target with a new debug container
  kubectl debug target -it --image=busybox --attach=false -c debug-container "${kube_flags[@]:?}"
  # Post-Conditions
  kube::test::get_object_assert pod/target '{{range.spec.ephemeralContainers}}{{.name}}:{{end}}' 'debug-container:'
  # Clean up
  kubectl delete pod target "${kube_flags[@]:?}"

  ### Pod Troubleshooting by Copy

  # Pre-Condition: Pod "nginx" is created
  kubectl run target "--image=${IMAGE_NGINX:?}" "${kube_flags[@]:?}"
  kube::test::get_object_assert pod "{{range.items}}{{${id_field:?}}}:{{end}}" 'target:'
  # Command: create a copy of target with a new debug container
  kubectl debug target -it --copy-to=target-copy --image=busybox --container=debug-container --attach=false "${kube_flags[@]:?}"
  # Post-Conditions
  kube::test::get_object_assert pod "{{range.items}}{{${id_field:?}}}:{{end}}" 'target:target-copy:'
  kube::test::get_object_assert pod/target-copy '{{range.spec.containers}}{{.name}}:{{end}}' 'target:debug-container:'
  kube::test::get_object_assert pod/target-copy '{{range.spec.containers}}{{.image}}:{{end}}' "${IMAGE_NGINX:?}:busybox:"
  # Clean up
  kubectl delete pod target target-copy "${kube_flags[@]:?}"

  # Pre-Condition: Pod "nginx" is created
  kubectl run target "--image=${IMAGE_NGINX:?}" "${kube_flags[@]:?}"
  kube::test::get_object_assert pod "{{range.items}}{{${id_field:?}}}:{{end}}" 'target:'
  # Command: create a copy of target with a new debug container replacing the previous pod
  kubectl debug target -it --copy-to=target-copy --image=busybox --container=debug-container --attach=false --replace "${kube_flags[@]:?}"
  # Post-Conditions
  kube::test::get_object_assert pod "{{range.items}}{{${id_field:?}}}:{{end}}" 'target-copy:'
  kube::test::get_object_assert pod/target-copy '{{range.spec.containers}}{{.name}}:{{end}}' 'target:debug-container:'
  kube::test::get_object_assert pod/target-copy '{{range.spec.containers}}{{.image}}:{{end}}' "${IMAGE_NGINX:?}:busybox:"
  # Clean up
  kubectl delete pod target-copy "${kube_flags[@]:?}"

  # Pre-Condition: Pod "nginx" is created
  kubectl run target "--image=${IMAGE_NGINX:?}" "${kube_flags[@]:?}"
  kube::test::get_object_assert pod "{{range.items}}{{${id_field:?}}}:{{end}}" 'target:'
  kube::test::get_object_assert pod/target '{{(index .spec.containers 0).name}}' 'target'
  # Command: copy the pod and replace the image of an existing container
  kubectl debug target --image=busybox --container=target --copy-to=target-copy "${kube_flags[@]:?}" -- sleep 1m
  # Post-Conditions
  kube::test::get_object_assert pod "{{range.items}}{{${id_field:?}}}:{{end}}" 'target:target-copy:'
  kube::test::get_object_assert pod/target-copy "{{(len .spec.containers)}}:{{${image_field:?}}}" '1:busybox'
  # Clean up
  kubectl delete pod target target-copy "${kube_flags[@]:?}"

  # Pre-Condition: Pod "nginx" is created
  kubectl run target "--image=${IMAGE_NGINX:?}" "${kube_flags[@]:?}"
  kube::test::get_object_assert pod "{{range.items}}{{${id_field:?}}}:{{end}}" 'target:'
  kube::test::get_object_assert pod/target '{{(index .spec.containers 0).name}}' 'target'
  # Command: copy the pod and replace the image of an existing container
  kubectl get pod/target -o yaml > "${KUBE_TEMP}"/test-pod-debug.yaml
  kubectl debug -f "${KUBE_TEMP}"/test-pod-debug.yaml --image=busybox --container=target --copy-to=target-copy "${kube_flags[@]:?}" -- sleep 1m
  # Post-Conditions
  kube::test::get_object_assert pod "{{range.items}}{{${id_field:?}}}:{{end}}" 'target:target-copy:'
  kube::test::get_object_assert pod/target-copy "{{(len .spec.containers)}}:{{${image_field:?}}}" '1:busybox'
  # Clean up
  kubectl delete pod target target-copy "${kube_flags[@]:?}"

  set +o nounset
  set +o errexit
}

run_kubectl_debug_node_tests() {
  set -o nounset
  set -o errexit

  create_and_use_new_namespace
  kube::log::status "Testing kubectl debug (pod tests)"

  ### Node Troubleshooting by Privileged Container

  # Pre-Condition: Pod "nginx" is created
  kube::test::get_object_assert nodes "{{range.items}}{{${id_field:?}}}:{{end}}" '127.0.0.1:'
  # Command: create a new node debugger pod
  output_message=$(kubectl debug node/127.0.0.1 --image=busybox --attach=false "${kube_flags[@]:?}" -- true)
  # Post-Conditions
  kube::test::get_object_assert pod "{{(len .items)}}" '1'
  debugger=$(kubectl get pod -o go-template="{{(index .items 0)${id_field:?}}}")
  kube::test::if_has_string "${output_message:?}" "${debugger:?}"
  kube::test::get_object_assert "pod/${debugger:?}" "{{${image_field:?}}}" 'busybox'
  kube::test::get_object_assert "pod/${debugger:?}" '{{.spec.nodeName}}' '127.0.0.1'
  kube::test::get_object_assert "pod/${debugger:?}" '{{.spec.hostIPC}}' 'true'
  kube::test::get_object_assert "pod/${debugger:?}" '{{.spec.hostNetwork}}' 'true'
  kube::test::get_object_assert "pod/${debugger:?}" '{{.spec.hostPID}}' 'true'
  kube::test::get_object_assert "pod/${debugger:?}" '{{(index (index .spec.containers 0).volumeMounts 0).mountPath}}' '/host'
  kube::test::get_object_assert "pod/${debugger:?}" '{{(index .spec.volumes 0).hostPath.path}}' '/'
  # Clean up
  # pod.spec.nodeName is set by kubectl debug node which causes the delete to hang,
  # presumably waiting for a kubelet that's not present. Force the delete.
  kubectl delete --force pod "${debugger:?}" "${kube_flags[@]:?}"

  set +o nounset
  set +o errexit
}

run_kubectl_debug_general_tests() {
  set -o nounset
  set -o errexit

  create_and_use_new_namespace
  kube::log::status "Testing kubectl debug profile general"

  ### Debug by pod copy
  ### probes are removed, sets SYS_PTRACE in debugging container, sets shareProcessNamespace

  # Pre-Condition: Pod "nginx" is created
  kubectl run target "--image=${IMAGE_NGINX:?}" "${kube_flags[@]:?}"
  kube::test::get_object_assert pod "{{range.items}}{{${id_field:?}}}:{{end}}" 'target:'
  # Command: create a copy of target with a new debug container
  kubectl debug --profile general target -it --copy-to=target-copy --image=busybox --container=debug-container --attach=false "${kube_flags[@]:?}"
  # Post-Conditions
  kube::test::get_object_assert pod "{{range.items}}{{${id_field:?}}}:{{end}}" 'target:target-copy:'
  kube::test::get_object_assert pod/target-copy '{{range.spec.containers}}{{.name}}:{{end}}' 'target:debug-container:'
  kube::test::get_object_assert pod/target-copy '{{range.spec.containers}}{{.image}}:{{end}}' "${IMAGE_NGINX:?}:busybox:"
  kube::test::get_object_assert pod/target-copy '{{range.spec.containers}}{{if (index . "livenessProbe")}}:{{end}}{{end}}' ''
  kube::test::get_object_assert pod/target-copy '{{range.spec.containers}}{{if (index . "readinessProbe")}}:{{end}}{{end}}' ''
  kube::test::get_object_assert pod/target-copy '{{(index (index .spec.containers 1).securityContext.capabilities.add 0)}}' 'SYS_PTRACE'
  kube::test::get_object_assert pod/target-copy '{{.spec.shareProcessNamespace}}' 'true'
  # Clean up
  kubectl delete pod target target-copy "${kube_flags[@]:?}"

  ### Debug by EC
  ### sets SYS_PTRACE in ephemeral container

  # Pre-Condition: Pod "nginx" is created
  kubectl run target "--image=${IMAGE_NGINX:?}" "${kube_flags[@]:?}"
  kube::test::get_object_assert pod "{{range.items}}{{${id_field:?}}}:{{end}}" 'target:'
  # Command: create a copy of target with a new debug container
  kubectl debug --profile general target -it --image=busybox --container=debug-container --attach=false "${kube_flags[@]:?}"
  # Post-Conditions
  kube::test::get_object_assert pod/target '{{range.spec.ephemeralContainers}}{{.name}}:{{.image}}{{end}}' 'debug-container:busybox'
  kube::test::get_object_assert pod/target '{{(index (index .spec.ephemeralContainers 0).securityContext.capabilities.add 0)}}' 'SYS_PTRACE'
  # Clean up
  kubectl delete pod target "${kube_flags[@]:?}"

  set +o nounset
  set +o errexit
}

run_kubectl_debug_general_node_tests() {
  set -o nounset
  set -o errexit

  create_and_use_new_namespace
  kube::log::status "Testing kubectl debug profile general (node)"

  ### Debug node
  ### empty securityContext, uses host namespaces, mounts root partition

  # Pre-Condition: node exists
  kube::test::get_object_assert nodes "{{range.items}}{{${id_field:?}}}:{{end}}" '127.0.0.1:'
  # Command: create a new node debugger pod
  output_message=$(kubectl debug --profile general node/127.0.0.1 --image=busybox --attach=false "${kube_flags[@]:?}" -- true)
  # Post-Conditions
  kube::test::get_object_assert pod "{{(len .items)}}" '1'
  debugger=$(kubectl get pod -o go-template="{{(index .items 0)${id_field:?}}}")
  kube::test::if_has_string "${output_message:?}" "${debugger:?}"
  kube::test::get_object_assert "pod/${debugger:?}" "{{${image_field:?}}}" 'busybox'
  kube::test::get_object_assert "pod/${debugger:?}" '{{.spec.nodeName}}' '127.0.0.1'
  kube::test::get_object_assert "pod/${debugger:?}" '{{.spec.hostIPC}}' 'true'
  kube::test::get_object_assert "pod/${debugger:?}" '{{.spec.hostNetwork}}' 'true'
  kube::test::get_object_assert "pod/${debugger:?}" '{{.spec.hostPID}}' 'true'
  kube::test::get_object_assert "pod/${debugger:?}" '{{(index (index .spec.containers 0).volumeMounts 0).mountPath}}' '/host'
  kube::test::get_object_assert "pod/${debugger:?}" '{{(index .spec.volumes 0).hostPath.path}}' '/'
  kube::test::get_object_assert "pod/${debugger:?}" '{{if (index (index .spec.containers 0) "securityContext")}}:{{end}}' ''
  # Clean up
  # pod.spec.nodeName is set by kubectl debug node which causes the delete to hang,
  # presumably waiting for a kubelet that's not present. Force the delete.
  kubectl delete --force pod "${debugger:?}" "${kube_flags[@]:?}"

  set +o nounset
  set +o errexit
}

run_kubectl_debug_baseline_tests() {
  set -o nounset
  set -o errexit

  create_and_use_new_namespace
  kube::log::status "Testing kubectl debug profile baseline"

  ### Debug by pod copy
  ### probes are removed, empty securityContext, sets shareProcessNamespace

  # Pre-Condition: Pod "nginx" is created
  kubectl run target "--image=${IMAGE_NGINX:?}" "${kube_flags[@]:?}"
  kube::test::get_object_assert pod "{{range.items}}{{${id_field:?}}}:{{end}}" 'target:'
  # Command: create a copy of target with a new debug container
  kubectl debug --profile baseline target -it --copy-to=target-copy --image=busybox --container=debug-container --attach=false "${kube_flags[@]:?}"
  # Post-Conditions
  kube::test::get_object_assert pod "{{range.items}}{{${id_field:?}}}:{{end}}" 'target:target-copy:'
  kube::test::get_object_assert pod/target-copy '{{range.spec.containers}}{{.name}}:{{end}}' 'target:debug-container:'
  kube::test::get_object_assert pod/target-copy '{{range.spec.containers}}{{.image}}:{{end}}' "${IMAGE_NGINX:?}:busybox:"
  kube::test::get_object_assert pod/target-copy '{{range.spec.containers}}{{if (index . "livenessProbe")}}:{{end}}{{end}}' ''
  kube::test::get_object_assert pod/target-copy '{{range.spec.containers}}{{if (index . "readinessProbe")}}:{{end}}{{end}}' ''
  kube::test::get_object_assert pod/target-copy '{{if (index (index .spec.containers 0) "securityContext")}}:{{end}}' ''
  kube::test::get_object_assert pod/target-copy '{{.spec.shareProcessNamespace}}' 'true'
  # Clean up
  kubectl delete pod target target-copy "${kube_flags[@]:?}"

  ### Debug by EC
  ### empty securityContext

  # Pre-Condition: Pod "nginx" is created
  kubectl run target "--image=${IMAGE_NGINX:?}" "${kube_flags[@]:?}"
  kube::test::get_object_assert pod "{{range.items}}{{${id_field:?}}}:{{end}}" 'target:'
  # Command: create a copy of target with a new debug container
  kubectl debug --profile baseline target -it --image=busybox --container=debug-container --attach=false "${kube_flags[@]:?}"
  # Post-Conditions
  kube::test::get_object_assert pod/target '{{range.spec.ephemeralContainers}}{{.name}}:{{.image}}{{end}}' 'debug-container:busybox'
  kube::test::get_object_assert pod/target '{{if (index (index .spec.ephemeralContainers 0) "securityContext")}}:{{end}}' ''
  # Clean up
  kubectl delete pod target "${kube_flags[@]:?}"

  set +o nounset
  set +o errexit
}

run_kubectl_debug_baseline_node_tests() {
  set -o nounset
  set -o errexit

  create_and_use_new_namespace
  kube::log::status "Testing kubectl debug profile baseline (node)"

  ### Debug node
  ### empty securityContext, uses isolated namespaces

  # Pre-Condition: node exists
  kube::test::get_object_assert nodes "{{range.items}}{{${id_field:?}}}:{{end}}" '127.0.0.1:'
  # Command: create a new node debugger pod
  output_message=$(kubectl debug --profile baseline node/127.0.0.1 --image=busybox --attach=false "${kube_flags[@]:?}" -- true)
  # Post-Conditions
  kube::test::get_object_assert pod "{{(len .items)}}" '1'
  debugger=$(kubectl get pod -o go-template="{{(index .items 0)${id_field:?}}}")
  kube::test::if_has_string "${output_message:?}" "${debugger:?}"
  kube::test::get_object_assert "pod/${debugger:?}" "{{${image_field:?}}}" 'busybox'
  kube::test::get_object_assert "pod/${debugger:?}" '{{.spec.nodeName}}' '127.0.0.1'
  kube::test::get_object_assert "pod/${debugger:?}" '{{.spec.hostIPC}}' '<no value>'
  kube::test::get_object_assert "pod/${debugger:?}" '{{.spec.hostNetwork}}' '<no value>'
  kube::test::get_object_assert "pod/${debugger:?}" '{{.spec.hostPID}}' '<no value>'
  kube::test::get_object_assert "pod/${debugger:?}" '{{if (index (index .spec.containers 0) "securityContext")}}:{{end}}' ''
  # Clean up
  # pod.spec.nodeName is set by kubectl debug node which causes the delete to hang,
  # presumably waiting for a kubelet that's not present. Force the delete.
  kubectl delete --force pod "${debugger:?}" "${kube_flags[@]:?}"

  set +o nounset
  set +o errexit
}

run_kubectl_debug_restricted_tests() {
  set -o nounset
  set -o errexit

  create_and_use_new_namespace  
  kube::log::status "Testing kubectl debug profile restricted"

  ### Pod Troubleshooting by ephemeral containers with restricted profile
  
  # Pre-Condition: Pod "nginx" is created
  kubectl run target "--image=${IMAGE_NGINX:?}" "${kube_flags[@]:?}"
  kube::test::get_object_assert pod "{{range.items}}{{${id_field:?}}}:{{end}}" 'target:'
  # Restricted profile just works in not restricted namespace
  # Command: add a new debug container with restricted profile
  output_message=$(kubectl debug target -it --image=busybox --attach=false -c debug-container --profile=restricted "${kube_flags[@]:?}")
  kube::test::if_has_not_string "${output_message}" 'forbidden: violates PodSecurity'
  # Post-Conditions
  kube::test::get_object_assert pod/target '{{range.spec.ephemeralContainers}}{{.name}}:{{end}}' 'debug-container:'
  # Clean up
  kubectl delete pod target "${kube_flags[@]:?}"

  # Pre-Condition: Pod "nginx" is created
  kubectl run target "--image=${IMAGE_NGINX:?}" "${kube_flags[@]:?}"
  kube::test::get_object_assert pod "{{range.items}}{{${id_field:?}}}:{{end}}" 'target:'
  # Restricted profile just works in not restricted namespace
  # Command: create a copy of target with a new debug container
  kubectl debug target -it --copy-to=target-copy --image=busybox --container=debug-container --attach=false --profile=restricted "${kube_flags[@]:?}"
  # Post-Conditions
  kube::test::get_object_assert pod "{{range.items}}{{${id_field:?}}}:{{end}}" 'target:target-copy:'
  kube::test::get_object_assert pod/target-copy '{{range.spec.containers}}{{.name}}:{{end}}' 'target:debug-container:'
  kube::test::get_object_assert pod/target-copy '{{range.spec.containers}}{{.image}}:{{end}}' "${IMAGE_NGINX:?}:busybox:"
  # Clean up
  kubectl delete pod target target-copy "${kube_flags[@]:?}"

  ns_name="namespace-restricted"
  # Command: create namespace and add a label
  kubectl create namespace "${ns_name}"
  kubectl label namespace "${ns_name}" pod-security.kubernetes.io/enforce=restricted
  output_message=$(kubectl get namespaces "${ns_name}" --show-labels)
  kube::test::if_has_string "${output_message}" 'pod-security.kubernetes.io/enforce=restricted'
 
  # Pre-Condition: Pod "busybox" is created that complies with the restricted policy
  kubectl create -f hack/testdata/pod-restricted-runtime-default.yaml -n "${ns_name}"
  kube::test::get_object_assert "pod -n ${ns_name}" "{{range.items}}{{${id_field:?}}}:{{end}}" 'target:'
  # Restricted profile works when pod's seccompProfile is RuntimeDefault
  # Command: add a new debug container with restricted profile
  output_message=$(kubectl debug target -it --image=busybox --attach=false -c debug-container --profile=restricted -n "${ns_name}" "${kube_flags[@]:?}")
  kube::test::if_has_not_string "${output_message}" 'forbidden: violates PodSecurity'
  # Post-Conditions
  kube::test::get_object_assert "pod/target -n ${ns_name}" '{{range.spec.ephemeralContainers}}{{.name}}:{{end}}' 'debug-container:'
  # Clean up
  kubectl delete pod target -n "${ns_name}" "${kube_flags[@]:?}"

  # Pre-Condition: Pod "nginx" is created
  kubectl create -f hack/testdata/pod-restricted-runtime-default.yaml -n "${ns_name}"
  kube::test::get_object_assert "pod -n ${ns_name}" "{{range.items}}{{${id_field:?}}}:{{end}}" 'target:'
  # Restricted profile works when pod's seccompProfile is RuntimeDefault
  # Command: create a copy of target with a new debug container
  kubectl debug target -it --copy-to=target-copy --image=busybox --container=debug-container --attach=false --profile=restricted -n ${ns_name} "${kube_flags[@]:?}"
  # Post-Conditions
  kube::test::get_object_assert "pod -n ${ns_name}" "{{range.items}}{{${id_field:?}}}:{{end}}" 'target:target-copy:'
  kube::test::get_object_assert "pod/target-copy -n ${ns_name}" '{{range.spec.containers}}{{.name}}:{{end}}' 'target:debug-container:'
  kube::test::get_object_assert "pod/target-copy -n ${ns_name}" '{{range.spec.containers}}{{.image}}:{{end}}' "busybox:busybox:"
  # Clean up
  kubectl delete pod target target-copy -n "${ns_name}" "${kube_flags[@]:?}"

  # Pre-Condition: Pod "busybox" is created that complies with the restricted policy
  kubectl create -f hack/testdata/pod-restricted-localhost.yaml -n "${ns_name}"
  kube::test::get_object_assert "pod -n ${ns_name}" "{{range.items}}{{${id_field:?}}}:{{end}}" 'target:'
  # Restricted profile works when pod's seccompProfile is Localhost
  # Command: add a new debug container with restricted profile
  output_message=$(kubectl debug target -it --image=busybox --attach=false -c debug-container --profile=restricted -n ${ns_name} "${kube_flags[@]:?}")
  kube::test::if_has_not_string "${output_message}" 'forbidden: violates PodSecurity'
  # Post-Conditions
  kube::test::get_object_assert "pod/target -n ${ns_name}" '{{range.spec.ephemeralContainers}}{{.name}}:{{end}}' 'debug-container:'
  # Clean up
  kubectl delete pod target -n ${ns_name} "${kube_flags[@]:?}"

  # Pre-Condition: Pod "nginx" is created
  kubectl create -f hack/testdata/pod-restricted-localhost.yaml -n "${ns_name}"
  kube::test::get_object_assert "pod -n ${ns_name}" "{{range.items}}{{${id_field:?}}}:{{end}}" 'target:'
  # Restricted profile works when pod's seccompProfile is Localhost
  # Command: create a copy of target with a new debug container
  kubectl debug target -it --copy-to=target-copy --image=busybox --container=debug-container --attach=false --profile=restricted -n ${ns_name} "${kube_flags[@]:?}"
  # Post-Conditions
  kube::test::get_object_assert "pod -n ${ns_name}" "{{range.items}}{{${id_field:?}}}:{{end}}" 'target:target-copy:'
  kube::test::get_object_assert "pod/target-copy -n ${ns_name}" '{{range.spec.containers}}{{.name}}:{{end}}' 'target:debug-container:'
  kube::test::get_object_assert "pod/target-copy -n ${ns_name}" '{{range.spec.containers}}{{.image}}:{{end}}' "busybox:busybox:"
  # Clean up
  kubectl delete pod target target-copy -n "${ns_name}" "${kube_flags[@]:?}"

  # Clean up restricted namespace
  kubectl delete namespace "${ns_name}"

  set +o nounset
  set +o errexit
}

run_kubectl_debug_restricted_node_tests() {
  set -o nounset
  set -o errexit

  create_and_use_new_namespace  
  kube::log::status "Testing kubectl debug profile restricted (node)"

  ### Debug node with restrected profile
  
  # Pre-Condition: node exists
  kube::test::get_object_assert nodes "{{range.items}}{{${id_field:?}}}:{{end}}" '127.0.0.1:'
  # Restricted profile just works in not restricted namespace
  # Command: create a new node debugger pod
  output_message=$(kubectl debug --profile restricted node/127.0.0.1 --image=busybox --attach=false "${kube_flags[@]:?}" -- true)
  kube::test::if_has_not_string "${output_message}" 'forbidden: violates PodSecurity'
  # Post-Conditions
  kube::test::get_object_assert pod "{{(len .items)}}" '1'
  debugger=$(kubectl get pod -o go-template="{{(index .items 0)${id_field:?}}}")
  kube::test::if_has_string "${output_message:?}" "${debugger:?}"
  kube::test::get_object_assert "pod/${debugger:?}" "{{${image_field:?}}}" 'busybox'
  kube::test::get_object_assert "pod/${debugger:?}" '{{.spec.nodeName}}' '127.0.0.1'
  kube::test::get_object_assert "pod/${debugger:?}" '{{.spec.hostIPC}}' '<no value>'
  kube::test::get_object_assert "pod/${debugger:?}" '{{.spec.hostNetwork}}' '<no value>'
  kube::test::get_object_assert "pod/${debugger:?}" '{{.spec.hostPID}}' '<no value>'
  kube::test::get_object_assert "pod/${debugger:?}" '{{index .spec.containers 0 "securityContext" "allowPrivilegeEscalation"}}' 'false'
  kube::test::get_object_assert "pod/${debugger:?}" '{{index .spec.containers 0 "securityContext" "capabilities" "drop"}}' '\[ALL\]'
  kube::test::get_object_assert "pod/${debugger:?}" '{{if (index (index .spec.containers 0) "securityContext" "capabilities" "add") }}:{{end}}' ''
  kube::test::get_object_assert "pod/${debugger:?}" '{{index .spec.containers 0 "securityContext" "runAsNonRoot"}}' 'true'
  kube::test::get_object_assert "pod/${debugger:?}" '{{index .spec.containers 0 "securityContext" "seccompProfile" "type"}}' 'RuntimeDefault'
  # Clean up
  # pod.spec.nodeName is set by kubectl debug node which causes the delete to hang,
  # presumably waiting for a kubelet that's not present. Force the delete.
  kubectl delete --force pod "${debugger:?}" "${kube_flags[@]:?}"

  ns_name="namespace-restricted"
  # Command: create namespace and add a label
  kubectl create namespace "${ns_name}"
  kubectl label namespace "${ns_name}" pod-security.kubernetes.io/enforce=restricted
  output_message=$(kubectl get namespaces "${ns_name}" --show-labels)
  kube::test::if_has_string "${output_message}" 'pod-security.kubernetes.io/enforce=restricted'

  # Pre-Condition: node exists
  kube::test::get_object_assert nodes "{{range.items}}{{${id_field:?}}}:{{end}}" '127.0.0.1:'
  # Restricted profile works in restricted namespace
  # Command: create a new node debugger pod
  output_message=$(kubectl debug --profile restricted node/127.0.0.1 --image=busybox --attach=false -n ${ns_name} "${kube_flags[@]:?}" -- true)
  kube::test::if_has_not_string "${output_message}" 'forbidden: violates PodSecurity'
  # Post-Conditions
  kube::test::get_object_assert "pod -n ${ns_name}" "{{(len .items)}}" '1'
  debugger=$(kubectl get pod -n ${ns_name} -o go-template="{{(index .items 0)${id_field:?}}}")
  kube::test::if_has_string "${output_message:?}" "${debugger:?}"
  kube::test::get_object_assert "pod/${debugger:?} -n ${ns_name}" "{{${image_field:?}}}" 'busybox'
  kube::test::get_object_assert "pod/${debugger:?} -n ${ns_name}" '{{.spec.nodeName}}' '127.0.0.1'
  kube::test::get_object_assert "pod/${debugger:?} -n ${ns_name}" '{{.spec.hostIPC}}' '<no value>'
  kube::test::get_object_assert "pod/${debugger:?} -n ${ns_name}" '{{.spec.hostNetwork}}' '<no value>'
  kube::test::get_object_assert "pod/${debugger:?} -n ${ns_name}" '{{.spec.hostPID}}' '<no value>'
  kube::test::get_object_assert "pod/${debugger:?} -n ${ns_name}" '{{index .spec.containers 0 "securityContext" "allowPrivilegeEscalation"}}' 'false'
  kube::test::get_object_assert "pod/${debugger:?} -n ${ns_name}" '{{index .spec.containers 0 "securityContext" "capabilities" "drop"}}' '\[ALL\]'
  kube::test::get_object_assert "pod/${debugger:?} -n ${ns_name}" '{{if (index (index .spec.containers 0) "securityContext" "capabilities" "add") }}:{{end}}' ''
  kube::test::get_object_assert "pod/${debugger:?} -n ${ns_name}" '{{index .spec.containers 0 "securityContext" "runAsNonRoot"}}' 'true'
  kube::test::get_object_assert "pod/${debugger:?} -n ${ns_name}" '{{index .spec.containers 0 "securityContext" "seccompProfile" "type"}}' 'RuntimeDefault'
  # Clean up
  # pod.spec.nodeName is set by kubectl debug node which causes the delete to hang,
  # presumably waiting for a kubelet that's not present. Force the delete.
  kubectl delete --force pod "${debugger:?}" -n ${ns_name} "${kube_flags[@]:?}"

  # Clean up restricted namespace
  kubectl delete namespace "${ns_name}"

  set +o nounset
  set +o errexit
}