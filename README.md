# Kamel

Kamel (a.k.a. Camel-K) is a lightweight integration framework built from Apache Camel that runs natively on Kubernetes
and is specifically designed for serverless and microservice architectures.

With Kamel, users can design integrations declaratively and run them on Kubernetes, without caring about platforms or servers.
Kamel creates everything needed to run the integrations and optimizes resource utilization with several techniques, including running integrations on
a [new ligtweight runtime written in Go](https://github.com/lburgazzoli/camel-go) or activating/deactivating them on demand.

Integrations can be used for many tasks: from just connecting a function to a external service, to orchestrating multiple functions and microservices
in arbitrarily complex workflows.

Kamel supports all 200+ components available in Apache Camel for connecting to virtually any cloud service and is built around enterprise integration patterns.

The project is currently in status: **alpha**.

## Roadmap

Towards "beta" release:

- [x] Define a roadmap to startup
- [ ] Package integrations in a builder service that produces runnable artifacts, one for each runtime (this is needed e.g. to enable using the full set of Apache Camel components)
- [ ] Publish and run artifacts produced by a builder service (no matter the runtime). It may be a docker image or a tar file to improve performance.
- [ ] Create the "kamel" binary
- [ ] Define a release strategy and setup a CI/CD pipeline
- [ ] Support integrations written in XML
- [ ] Allow to specify dependencies in a uniform manner (Go packages, Java libraries)
- [ ] Automatically discover Apache Camel dependencies needed by a integration using the catalog
- [ ] Enhance the flow DSL
  - [ ] REST
  - [ ] setHeader, setBody
  - [ ] Simple language
  - [ ] Basic EIPs
- [ ] Define and implement a data mapping strategy

Towards first "stable" release:

- [ ] Documentation and website
- [ ] Add connectors for popular serverless platforms
- [ ] Add connectors for webhooks and cloud events
- [ ] Provide client libraries for popular languages to invoke Kamel integrations from inside functions (e.g. gRPC based)
- [ ] Add a "on-demand" execution model for:
  - [ ] Timer component
  - [ ] Http component
- [ ] Optimize performance of the Java runtime (compilation, configuration)
- [ ] Uniform Go and Java runtime for common components

Towards "future":
- [ ] Extend "on-demand" execution model to other components (messaging)
- [ ] Enhance secret management and evaluate integration with vault services
- [ ] Add tracing capabilities
- [ ] Evaluate integration with Istio
- [ ] Evaluate integration with API management platform
- [ ] Enable integration deployment "on top" of other Kubernetes-based serverless platforms (Knative, Openwhisk)


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
  flows:
   - id: timer
     name: "The Timer"
     steps:
      - type: endpoint
        uri: timer:tick
      - type: endpoint
        uri: log:info
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