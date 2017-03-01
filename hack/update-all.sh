#!/bin/bash

# Copyright 2014 The Kubernetes Authors.
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

# A single script that runs a predefined set of update-* scripts, as they often go together.
set -o errexit
set -o nounset
set -o pipefail

KUBE_ROOT=$(dirname "${BASH_SOURCE}")/..
source "${KUBE_ROOT}/cluster/lib/util.sh"

SILENT=true
ALL=false

while getopts ":va" opt; do
	case $opt in
		a)
			ALL=true
			;;
		v)
			SILENT=false
			;;
		\?)
			echo "Invalid flag: -$OPTARG" >&2
			exit 1
			;;
	esac
done

trap 'exit 1' SIGINT

if $SILENT ; then
	echo "Running in the silent mode, run with -v if you want to see script logs."
fi

if ! $ALL ; then
	echo "Running in short-circuit mode; run with -a to force all scripts to run."
fi

BASH_TARGETS="
	update-generated-protobuf
	update-codegen
	update-codecgen
	update-generated-docs
	update-generated-swagger-docs
	update-swagger-spec
	update-openapi-spec
	update-api-reference-docs
	update-bazel
	update-federation-openapi-spec
	verify-staging-client-go
	verify-staging-godeps"

for t in $BASH_TARGETS
do
	echo -e "${color_yellow}Running $t${color_norm}"
	if $SILENT ; then
		if ! bash "$KUBE_ROOT/hack/$t.sh" 1> /dev/null; then
			echo -e "${color_red}Running $t FAILED${color_norm}"
			if [[ $t == "verify"* ]]; then
				echo -e "${color_red}Run ./hack/update-all-staging.sh to fix it"
			fi
			if ! $ALL; then
				exit 1
			fi
		fi
	else
		if ! bash "$KUBE_ROOT/hack/$t.sh"; then
			echo -e "${color_red}Running $t FAILED${color_norm}"
			if [[ $t == "verify"* ]]; then
				echo -e "${color_red}Run ./hack/update-all-staging.sh to fix it"
			fi
			if ! $ALL; then
				exit 1
			fi
		fi
	fi
done

echo -e "${color_green}Update scripts completed successfully${color_norm}"
