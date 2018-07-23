#!/bin/sh

set -o errexit

REPO_ROOT=$(git rev-parse --show-toplevel)
OUTPUT_GOBIN="${REPO_ROOT}/_output/bin"

GOBIN="${OUTPUT_GOBIN}" go install ./vendor/github.com/jteeuwen/go-bindata/go-bindata
PATH="${OUTPUT_GOBIN}:${PATH}" go generate k8s.io/test-infra/kind/...
