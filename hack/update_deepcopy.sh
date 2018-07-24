#!/bin/sh

# Generates the deepcopy boilerplate for the custom resources
# Run it after changing types in `pkg/apis/camel/*`.

# Requires deepcopy-gen in the path and kubernetes gengo in the source (go get k8s.io/gengo)

basedir=$(dirname "$0")
deepcopy-gen -O zz_generated.deepcopy -i github.com/nicolaferraro/integration-operator/pkg/apis/camel/v1alpha1/
