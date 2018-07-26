#!/bin/sh

# This script builds the docker image for the go runtime
# Ensure you're connected to the correct docker daemon

# Image name: nferraro/camel-go-runtime:latest

location=$(dirname $0)

cd $location/../

docker build -t nferraro/camel-go-runtime:latest -f $location/../runtime/go/Dockerfile .