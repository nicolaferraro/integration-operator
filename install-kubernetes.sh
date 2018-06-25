#!/bin/sh

echo "Installing Integration Operator in Kubernetes"
echo "Minimum Openshift version is v1.9.0"
sleep 2

kubectl replace --force -f https://raw.githubusercontent.com/nicolaferraro/integration-operator/master/deploy/rbac.yaml
kubectl replace --force -f https://raw.githubusercontent.com/nicolaferraro/integration-operator/master/deploy/crd.yaml
kubectl replace --force -f https://raw.githubusercontent.com/nicolaferraro/integration-operator/master/deploy/operator.yaml

#kubectl replace --force -f deploy/rbac.yaml
#kubectl replace --force -f deploy/crd.yaml
#kubectl replace --force -f deploy/operator.yaml
