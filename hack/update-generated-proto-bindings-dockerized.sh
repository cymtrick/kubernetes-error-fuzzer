#!/usr/bin/env bash

# Copyright 2017 The Kubernetes Authors.
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

# This script generates `*/api.pb.go` files from protobuf files `*/api.proto`.

set -o errexit
set -o nounset
set -o pipefail

KUBE_ROOT="$(cd "$(dirname "${BASH_SOURCE[0]}")/../" && pwd -P)"

source "${KUBE_ROOT}/hack/lib/protoc.sh"

if [ "$#" = 0 ]; then
    echo "usage: $0 <api_dir>..."
    exit 1
fi

for api; do
    # This can't use `git ls-files` because it runs in a container without the
    # .git dir synced.
    find "${api}" -type f -name "api.proto" \
        | while read -r F; do
            D="$(dirname "${F}")"
            kube::protoc::generate_proto "${KUBE_ROOT}/${D}"
        done
done
