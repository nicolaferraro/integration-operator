# Integration Operator

This operator allows to deploy integrations as first-class citizen resources in Kubernetes or Openshift.
Integrations are based on the popular integration framework Apache Camel.

## Installing the Operator

The operator **requires Kubernetes 1.9+ or Openshift 3.9+**. Detailed instructions follow.

The installation process will add to the cluster a new custom resource type named `Integration` with API version `camel.apache.org/v1alpha1`.

### Installing on Kubernetes

To install the operator on Kubernetes, execute the following command:

```
curl https://raw.githubusercontent.com/nicolaferraro/integration-operator/master/install-kubernetes.sh | sh
```

### Installing on Openshift

For simplicity, we will use a admin user to create the operator and deploy integrations.
Configuration for normal users is just more complex.

Login into Openshift with **cluster-admin** privileges. In *Minishift* this can be done with:

```
oc login -u system:admin
```

Once you're logged in, you can install the operator with:

```
curl https://raw.githubusercontent.com/nicolaferraro/integration-operator/master/install-openshift.sh | sh
```

## Create a Sample Integration

The following integration creates a timer that displays a message periodically:

```
kind: "Integration"
apiVersion: "camel.apache.org/v1alpha1"
metadata:
  name: "example"
spec:
  replicas: 1
  routes:
   - id: timer
     route:
      - timer:tick
      - log:info
```

You can save it into a file, e.g. `example.yaml`, then execute the following command:

```
kubectl create -f example.yaml

## For Openshift
# oc create -f example.yaml
```

## Check Installed Integrations

You can check all deployed integrations by executing:

```
kubectl get integrations

## For Openshift
# oc get integrations
```

For each integration running you will see a pod that has the integration name as prefix. Type the following commands to list all pods:

```
kubectl get pods

## For Openshift
# oc get pods
```

Get the name of the pod, e.g. `example-xa3r3s2q`, then execute the following command to watch its logs:

```
kubectl logs -f example-xa3r3s2q

## For Openshift
# oc logs -f example-xa3r3s2q
```

## Edit a Integration

You can edit the integration by executing the following command:

```
kubectl edit integration example

## For Openshift
# oc edit integration example
```

You can change the routes in the `spec` part, save and see the integration automatically redeployed.

## Delete a Integration

You can finally delete the integration by executing the following command:

```
kubectl delete integration example

## For Openshift
# oc delete integration example
```