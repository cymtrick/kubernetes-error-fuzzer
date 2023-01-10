#!/usr/bin/env bash

# Copyright 2023 The Kubernetes Authors.
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

# This script checks API-related files for mismatch in docs and field names,
# and outputs a list of fields that their docs and field names are mismatched.
# Usage: `hack/verify-fieldname-docs.sh`.

set -o errexit
set -o nounset
set -o pipefail

KUBE_ROOT=$(dirname "${BASH_SOURCE[0]}")/..
source "${KUBE_ROOT}/hack/lib/init.sh"
source "${KUBE_ROOT}/hack/lib/util.sh"

kube::golang::setup_env

make -C "${KUBE_ROOT}" WHAT=cmd/fieldnamedocscheck

# Find binary
fieldnamedocscheck=$(kube::util::find-binary "fieldnamedocscheck")

result=0

find_files() {
  find . -not \( \
      \( \
        -wholename './output' \
        -o -wholename './_output' \
        -o -wholename './_gopath' \
        -o -wholename './release' \
        -o -wholename './target' \
        -o -wholename '*/third_party/*' \
        -o -wholename '*/vendor/*' \
        -o -wholename './pkg/*' \
      \) -prune \
    \) \
    \( -wholename './staging/src/k8s.io/api/*/v*/types.go' \
       -o -wholename './staging/src/k8s.io/kube-aggregator/pkg/apis/*/v*/types.go' \
       -o -wholename './staging/src/k8s.io/apiextensions-apiserver/pkg/apis/*/v*/types.go' \
    \)
}

versioned_api_files=$(find_files) || true

failure_file="${KUBE_ROOT}/hack/.fieldname_docs_failures"
failing_groups=()
while IFS='' read -r line; do failing_groups+=("${line}"); done < <(cat "${failure_file}")

for file in ${versioned_api_files}; do
	package="${file%"/types.go"}"
    if ! kube::util::array_contains "${package}" "${failing_groups[@]}"; then
        echo "Checking ${package}"
        ${fieldnamedocscheck} -s "${file}" || result=$?
    fi
done

exit ${result}
