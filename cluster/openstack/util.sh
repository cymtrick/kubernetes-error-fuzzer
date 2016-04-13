#!/bin/bash

# Copyright 2015 The Kubernetes Authors All rights reserved.
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

# A library of helper functions that each provider hosting Kubernetes must implement to use cluster/kube-*.sh scripts.

# exit on any error
set -e

# Use the config file specified in $KUBE_CONFIG_FILE, or default to
# config-default.sh.
KUBE_ROOT=$(dirname "${BASH_SOURCE}")/../..
readonly ROOT=$(dirname "${BASH_SOURCE}")
source "${ROOT}/${KUBE_CONFIG_FILE:-"config-default.sh"}"
source "$KUBE_ROOT/cluster/common.sh"
if [ $CREATE_IMAGE = true ]; then
source "${ROOT}/config-image.sh"
fi

# Verify prereqs on host machine
function verify-prereqs() {
 # Check the OpenStack command-line clients
 for client in swift glance nova heat;
 do
  if which $client >/dev/null 2>&1; then
    echo "$client client installed"
  else
    echo "$client client does not exist"
    echo "Please install $client client, and retry."
    exit 1
  fi
 done
}

# Instantiate a kubernetes cluster
#
# Assumed vars:
#   KUBERNETES_PROVIDER
function kube-up() {
    echo "kube-up for provider $KUBERNETES_PROVIDER"
    create-stack
}

# Periodically checks if cluster is created
#
# Assumed vars:
#   STACK_CREATE_TIMEOUT
#   STACK_NAME
function validate-cluster() {

  while (( --$STACK_CREATE_TIMEOUT >= 0)) ;do
     local status=$(heat stack-show "${STACK_NAME}" | awk '$2=="stack_status" {print $4}')
     if [[ $status ]]; then
        echo "Cluster status ${status}"
        if [ $status = "CREATE_COMPLETE" ]; then
          configure-kubectl
          break
        elif [ $status = "CREATE_FAILED" ]; then
          echo "Cluster not created. Please check stack logs to find the problem"
          break
        fi
     else
       echo "Cluster not created. Please verify if process started correctly"
       break
     fi
     sleep 60
  done
}

# Create stack
#
# Assumed vars:
#   OPENSTACK
#   OPENSTACK_TEMP
#   DNS_SERVER
#   OPENSTACK_IP
#   OPENRC_FILE
function create-stack() {
  echo "[INFO] Execute commands to create Kubernetes cluster"
  # It is required for some cloud provider like CityCloud where swift client has different credentials
  source "${ROOT}/openrc-swift.sh"
  if [[ -z ${OS_PROJECT_ID+x} ]]; then
    SWIFT_PROJECT_ID="${OS_TENANT_ID}"
  else
    SWIFT_PROJECT_ID="${OS_PROJECT_ID}"
  fi
  upload-resources
  source "${ROOT}/openrc-default.sh"

  create-glance-image

  add-keypair
  run-heat-script
}

# Upload kubernetes release tars and heat templates.
#
# Assumed vars:
#   ROOT
#   KUBERNETES_RELEASE_TAR
function upload-resources() {
  swift post kubernetes --read-acl '.r:*,.rlistings'

  echo "[INFO] Upload ${KUBERNETES_RELEASE_TAR}"
  swift upload kubernetes ${ROOT}/../../_output/release-tars/${KUBERNETES_RELEASE_TAR} \
    --object-name kubernetes-server.tar.gz

  echo "[INFO] Upload kubernetes-salt.tar.gz"
  swift upload kubernetes ${ROOT}/../../_output/release-tars/kubernetes-salt.tar.gz \
    --object-name kubernetes-salt.tar.gz
}

# Create a new key pair for use with servers.
#
# Assumed vars:
#   KUBERNETES_KEYPAIR_NAME
#   CLIENT_PUBLIC_KEY_PATH
function add-keypair() {
  local status=$(nova keypair-show ${KUBERNETES_KEYPAIR_NAME})
  if [[ ! $status ]]; then
    nova keypair-add ${KUBERNETES_KEYPAIR_NAME} --pub-key ${CLIENT_PUBLIC_KEY_PATH}
    echo "[INFO] Key pair created"
  else
    echo "[INFO] Key pair already exists"
  fi
}

# Create a new glance image.
#
# Assumed vars:
#   IMAGE_FILE
#   IMAGE_PATH
#   OPENSTACK_IMAGE_NAME
function create-glance-image() {
  if [ $CREATE_IMAGE = true ]; then
    local image_status=$(nova image-show ${OPENSTACK_IMAGE_NAME} | awk '$2=="id" {print $4}')

    if [[ ! $image_status ]]; then
      echo "[INFO] Create image ${OPENSTACK_IMAGE_NAME}"
      glance image-create --name ${OPENSTACK_IMAGE_NAME} --disk-format ${IMAGE_FORMAT} \
        --container-format ${CONTAINER_FORMAT} --file ${IMAGE_PATH}/${IMAGE_FILE}
    else
      echo "[INFO] Image ${OPENSTACK_IMAGE_NAME} already exists"
    fi
  fi
}

# Create a new kubernetes stack.
#
# Assumed vars:
#   STACK_NAME
#   KUBERNETES_KEYPAIR_NAME
#   DNS_SERVER
#   SWIFT_SERVER_URL
#   SWIFT_TENANT_ID
#   OPENSTACK_IMAGE_NAME
#   EXTERNAL_NETWORK
#   IMAGE_ID
#   MASTER_FLAVOR
#   MINION_FLAVOR
#   NUMBER_OF_MINIONS
#   MAX_NUMBER_OF_MINIONS
#   DNS_SERVER
#   STACK_NAME
function run-heat-script() {

  local stack_status=$(heat stack-show ${STACK_NAME})
  local swift_repo_url="${SWIFT_SERVER_URL}/v1/AUTH_${SWIFT_PROJECT_ID}/kubernetes"

  if [ $CREATE_IMAGE = true ]; then
    echo "[INFO] Retrieve new image ID"
    IMAGE_ID=$(nova image-show ${OPENSTACK_IMAGE_NAME} | awk '$2=="id" {print $4}')
    echo "[INFO] Image Id $IMAGE_ID"
  fi

  if [[ ! $stack_status ]]; then
    echo "[INFO] Create stack ${STACK_NAME}"
    (
      cd ${ROOT}/kubernetes-heat
      heat --api-timeout 60 stack-create \
      -P external_network=${EXTERNAL_NETWORK} \
      -P ssh_key_name=${KUBERNETES_KEYPAIR_NAME} \
      -P server_image=${IMAGE_ID} \
      -P master_flavor=${MASTER_FLAVOR} \
      -P minion_flavor=${MINION_FLAVOR} \
      -P number_of_minions=${NUMBER_OF_MINIONS} \
      -P max_number_of_minions=${MAX_NUMBER_OF_MINIONS} \
      -P dns_nameserver=${DNS_SERVER} \
      -P kubernetes_salt_url=${swift_repo_url}/kubernetes-salt.tar.gz \
      -P kubernetes_server_url=${swift_repo_url}/kubernetes-server.tar.gz \
      -P enable_proxy=${ENABLE_PROXY} \
      -P ftp_proxy="${FTP_PROXY}" \
      -P http_proxy="${HTTP_PROXY}" \
      -P https_proxy="${HTTPS_PROXY}" \
      -P socks_proxy="${SOCKS_PROXY}" \
      -P no_proxy="${NO_PROXY}" \
      --template-file kubecluster.yaml \
      ${STACK_NAME}
    )
  else
    echo "[INFO] Stack ${STACK_NAME} already exists"
    heat stack-show ${STACK_NAME}
  fi
}

# Configure kubectl.
#
# Assumed vars:
#   STACK_NAME
function configure-kubectl() {

  export KUBE_MASTER_IP=$(nova show "${STACK_NAME}"-master | awk '$3=="network" {print $6}')
  export CONTEXT="openstack"
  export KUBE_BEARER_TOKEN="TokenKubelet"
  create-kubeconfig
}


# Delete a kubernetes cluster
#
# Assumed vars:
#   STACK_NAME
function kube-down {
  source "${ROOT}/openrc-default.sh"
  heat stack-delete ${STACK_NAME}
}
