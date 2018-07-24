#!/bin/sh

# This script builds the docker image for the classic spring-boot runtime
# Ensure you're connected to the correct docker daemon

# Image name: nferraro/camel-classic-runtime-spring-boot:latest

location=$(dirname $0)

cd $location/../runtime/classic
mvn clean install -DskipTests

cd runner-spring-boot
mvn clean install -DskipTests -P kubernetes

