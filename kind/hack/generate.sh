#!/bin/sh
# Copyright 2018 The Kubernetes Authors.
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

# 'go generate's kind, using tools from vendor (go-bindata)

set -o errexit

REPO_ROOT=$(git rev-parse --show-toplevel)

# add_copyright_header replaces file $2 with an updated file containing
# the header contents in file $1, with YEAR in $1 replaced with the current year
add_copyright_header() {
    local year
    year="$(date '+%Y')"  
    sed "s/YEAR/${year}/" "${1}" > "${2}.new"
    cat "${2}" >> "${2}.new" && mv "${2}.new" "${2}" 
}

# install go-bindata from vendor locally
OUTPUT_GOBIN="${REPO_ROOT}/_output/bin"
GOBIN="${OUTPUT_GOBIN}" go install ./vendor/github.com/jteeuwen/go-bindata/go-bindata

 # go generate (using go-bindata)
# NOTE: go will only take package paths, not absolute directories
cd "${REPO_ROOT}"
PATH="${OUTPUT_GOBIN}:${PATH}" go generate ./kind/...

# fix the lack of copyright header
add_copyright_header "${REPO_ROOT}/kind/hack/copyright.go.txt" "${REPO_ROOT}/kind/pkg/build/sources/images_node_sources.go"

# gofmt the generated file
find ./kind -name "*.go" | xargs gofmt -s -w
