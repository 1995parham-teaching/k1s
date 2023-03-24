<h1 align="center"> k1s </h1>
<h6 align="center"> there is only one letter here from k to s which is E </h6>

<p align="center">
   <img alt="GitHub Workflow Status" src="https://img.shields.io/github/actions/workflow/status/1995parham-teaching/k1s/ci.yaml?style=for-the-badge">
</p>

## Introduction

Kubernetes is an awesome platform, and we want to have fun with it here with this project.
In this repository we've created an example Golang server and then created pod, service etc. for it.
The `hello-server` is a simple HTTP server that we want to deploy on the cloud.
It must be replicated, so everyone can get their hello even in the peak time.

## Step by Step

At the beginning you need an up and running Kubernetes cluster.
[Mircok8s](https://microk8s.io/docs) is an awesome platform if you don't know where to start.
(For having `kubectl` at your hand with `Mircok8s` check [this](https://microk8s.io/docs/working-with-kubectl))

First, create and switch to your desired namespace.

```bash
kubectl create namespace k1s
# kcd is an alias for kubectl context switch
alias kcd='kubectl config set-context $(kubectl config current-context) --namespace'
kcd k1s
```

Then follow these instructions to have your `hello-server` up and running.

The `hello-server` application requires a `config` to work. Possible options are providing `config` through a `config.yaml` file, environment variables, or stick to the default config.
Passing config files and environment variables is through a _ConfigMap_. To test each way, comment out the other one's way of loading in `k1s-deployment.yaml`

1. Create ConfigMap

   ```bash
   kubectl apply -f k1s-config-map.yaml
   kubectl get configmaps
   ```

2. Create Deployment (kubernetes needs _gcr_ and _docker-hub_ so have [proxy](https://microk8s.io/docs/install-proxy) at your pocket)

   ```bash
   kubectl apply -f k1s-deployment.yaml
   kubectl get deployment
   kubectl get pod
   ```

   To visualize the usage of `health-check` you can use `/die` endpoint to make `health-check` fail,
   and then check what happened with `kubectl get events --watch` to watch the whole process as it happens.

3. Create Service (make sure you did `microk8s enable dns`)

   ```bash
   kubectl apply -f k1s-svc.yaml
   kubectl get svc
   ```

   `k1s` service is used to associate a name for `k1s-deployment` pod's IP addresses inside cluster.
   You can create a pod and access k1s through it.

   ```bash
   kubectl run alpine -ti --image alpine --rm --restart=Never -- sh
   > apk add curl
   > curl htpp://k1s:1378
   ```

   By running this command several times, you can see that this service is working also as a simple `Load Balancer`.
   **Note: make sure `DNS` service is running for k8s.**

4. Create Ingress

   ```bash
   kubectl apply -f k1s-ingress.yaml
   ```

   > [Ingress](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.18/#ingress-v1beta1-networking-k8s-io) exposes HTTP and HTTPS routes from outside the cluster to [services](https://kubernetes.io/docs/concepts/services-networking/service/) within the cluster.

   It simply provides a way to access your pods from Internet by specifying which requests are destined for your `service`.
   **For ingress to work you must first add a [ingress controller](https://kubernetes.io/docs/concepts/services-networking/ingress-controllers). Also `DNS` service must be enabled**
