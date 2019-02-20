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

set -o errexit
set -o nounset
set -o pipefail

KUBE_ROOT=$(dirname "${BASH_SOURCE[0]}")/..
source "${KUBE_ROOT}/hack/lib/init.sh"

kube::golang::setup_env

make -C "${KUBE_ROOT}" WHAT=cmd/genswaggertypedocs

# Find binary
genswaggertypedocs=$(kube::util::find-binary "genswaggertypedocs")

if [[ ! -x "$genswaggertypedocs" ]]; then
  {
    echo "It looks as if you don't have a compiled genswaggertypedocs binary"
    echo
    echo "If you are running from a clone of the git repo, please run"
    echo "'make WHAT=cmd/genswaggertypedocs'."
  } >&2
  exit 1
fi

_tmpdir="$(kube::realpath "$(mktemp -d -t swagger-docs.XXXXXX)")"
function swagger_cleanup {
  rm -rf "${_tmpdir}"
}
kube::util::trap_add swagger_cleanup EXIT

# Copy the contents of the kube directory into the nice clean place
_kubetmp="${_tmpdir}/src/k8s.io"
mkdir -p "${_kubetmp}"
# should create ${_kubetmp}/kubernetes
git archive --format=tar --prefix=kubernetes/ "$(git write-tree)" | (cd "${_kubetmp}" && tar xf -)
_kubetmp="${_kubetmp}/kubernetes"
# Do all our work in the new GOPATH
export GOPATH="${_tmpdir}"

find_files() {
  find . -not \( \
      \( \
        -wholename './output' \
        -o -wholename './.git' \
        -o -wholename './_output' \
        -o -wholename './_gopath' \
        -o -wholename './release' \
        -o -wholename './target' \
        -o -wholename '*/third_party/*' \
        -o -wholename '*/vendor/*' \
        -o -wholename './staging/src/k8s.io/client-go/*vendor/*' \
      \) -prune \
    \) -name 'types_swagger_doc_generated.go'
}
while IFS=$'\n' read -r line; do TARGET_FILES+=("$line"); done < <(find_files)

pushd "${_kubetmp}" > /dev/null 2>&1
  # Update the generated swagger docs
  hack/update-generated-swagger-docs.sh
popd > /dev/null 2>&1

ret=0

pushd "${KUBE_ROOT}" > /dev/null 2>&1
  # Test for diffs
  _output=""
  for file in ${TARGET_FILES[*]}; do
    _output="${_output}$(diff -Naupr -I 'Auto generated by' "${KUBE_ROOT}/${file}" "${_kubetmp}/${file}")" || ret=1
  done

  if [[ ${ret} -gt 0 ]]; then
    echo "Generated swagger type documentation is out of date:" >&2
    echo "${_output}" >&2
    exit ${ret}
  fi
popd > /dev/null 2>&1

echo "Generated swagger type documentation up to date."
