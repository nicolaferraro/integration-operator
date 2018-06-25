#!/bin/sh

echo "Installing Integration Operator in Openshift"
echo "Minimum Openshift version is v3.9.0"
echo "Please, ensure you are logged in with cluster-admin privileges, or the installation will fail"
sleep 2

oc replace --force -f https://raw.githubusercontent.com/nicolaferraro/integration-operator/master/deploy/rbac.yaml
oc replace --force -f https://raw.githubusercontent.com/nicolaferraro/integration-operator/master/deploy/crd.yaml
oc replace --force -f https://raw.githubusercontent.com/nicolaferraro/integration-operator/master/deploy/operator.yaml

#oc replace --force -f deploy/rbac.yaml
#oc replace --force -f deploy/crd.yaml
#oc replace --force -f deploy/operator.yaml
