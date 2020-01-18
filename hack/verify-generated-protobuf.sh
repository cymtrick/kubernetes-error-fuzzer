#!/usr/bin/env bash

# Copyright 2015 The Kubernetes Authors.
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

# This script checks whether updating of the packages which explicitly
# requesting generation is needed or not. We should run
# `hack/update-generated-protobuf.sh` if those packages are out of date.
# Usage: `hack/verify-generated-protobuf.sh`.

set -o errexit
set -o nounset
set -o pipefail

KUBE_ROOT=$(dirname "${BASH_SOURCE[0]}")/..
source "${KUBE_ROOT}/hack/lib/init.sh"

kube::golang::setup_env

APIROOTS=$({
  # gather the packages explicitly requesting generation
  git grep --files-with-matches -e '// +k8s:protobuf-gen=package' cmd pkg staging | xargs -n 1 dirname
  # add the root apimachinery pkg containing implicitly generated files (--apimachinery-packages)
  echo staging/src/k8s.io/apimachinery/pkg
} | sort | uniq)

_tmp="${KUBE_ROOT}/_tmp"

cleanup() {
  rm -rf "${_tmp}"
}

trap "cleanup" EXIT SIGINT

cleanup
for APIROOT in ${APIROOTS}; do
  mkdir -p "${_tmp}/${APIROOT}"
  cp -a "${KUBE_ROOT}/${APIROOT}"/* "${_tmp}/${APIROOT}/"
done

KUBE_VERBOSE=3 "${KUBE_ROOT}/hack/update-generated-protobuf.sh"
for APIROOT in ${APIROOTS}; do
  TMP_APIROOT="${_tmp}/${APIROOT}"
  echo "diffing ${APIROOT} against freshly generated protobuf"
  ret=0
  diff -Naupr -I 'Auto generated by' -x 'zz_generated.*' -x '.github' "${KUBE_ROOT}/${APIROOT}" "${TMP_APIROOT}" || ret=$?
  cp -a "${TMP_APIROOT}"/* "${KUBE_ROOT}/${APIROOT}/"
  if [[ $ret -eq 0 ]]; then
    echo "${APIROOT} up to date."
  else
    echo "${APIROOT} is out of date. Please run hack/update-generated-protobuf.sh"
    exit 1
  fi
done
