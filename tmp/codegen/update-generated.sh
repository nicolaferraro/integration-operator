#!/usr/bin/env bash

set -o errexit
set -o nounset
set -o pipefail

DOCKER_REPO_ROOT="/go/src/github.com/nicolaferraro/integration-operator"
IMAGE=${IMAGE:-"gcr.io/coreos-k8s-scale-testing/codegen:1.9.3"}

docker run --rm \
  -v "$PWD":"$DOCKER_REPO_ROOT":Z \
  -w "$DOCKER_REPO_ROOT" \
  "$IMAGE" \
  "/go/src/k8s.io/code-generator/generate-groups.sh"  \
  "deepcopy" \
  "github.com/nicolaferraro/integration-operator/pkg/generated" \
  "github.com/nicolaferraro/integration-operator/pkg/apis" \
  "camel:v1alpha1" \
  --go-header-file "./tmp/codegen/boilerplate.go.txt" \
  $@
