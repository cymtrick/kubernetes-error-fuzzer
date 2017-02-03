#!/bin/bash

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

set -o errexit
set -o nounset
set -o pipefail

# Send the file content to the server
if command -v curl &>/dev/null; then
    curl -s -k -XPOST --data-binary "@${1}" -o "${1}.result" "${KUBE_EDITOR_CALLBACK}"
elif command -v wget &>/dev/null; then
    wget --post-file="${1}" -O "${1}.result" "${KUBE_EDITOR_CALLBACK}"
else
    echo "curl and wget are unavailable" >&2
    exit 1
fi

# Use the response as the edited version
mv "${1}.result" "${1}"
