# k1s

![GitHub Workflow Status](https://img.shields.io/github/workflow/status/1995parham/k1s/ci?label=ci&logo=github&style=flat-square)

## Introduction

Kubernetes is an awesome platform, and I want to have fun with it.
In this repository I have created an example Golang server and then created pod, service and etc. for it.
The `hello-server` is a simple HTTP server that we want to deploy on cloud. It must have replica so everyone can get their hello.

## Step by Step

At the beginning you need an up and running kubernetes cluster.
[Mircok8s](https://microk8s.io/docs) is an awesome platform if you don't know where to start. (For having `kubectl` at your hand with `Mircok8s` check [this](https://microk8s.io/docs/working-with-kubectl))

First of all, create and switch to your desired namespace.

```sh
kubectl create namespace k1s
# kcd is an alias for kubectl context switch
alias kcd='kubectl config set-context $(kubectl config current-context) --namespace'
kcd k1s
```

Then follow these instructions to have your `hello-server` up and running.

The `hello-server` application requires a `config` to work. Possible options are providing `config` through a `config.yaml` file, environment variables, or stick to the default config.
Passing config files and environment variables is through a _ConfigMap_. To test each way, comment out the other one's way of loading in `k1s-deployment.yaml`

1. Create ConfigMap

   ```sh
   kubectl apply -f k1s-config-map.yaml
   kubectl get configmaps
   ```

2. Create Deployment (kubernetes needs _gcr_ and _docker-hub_ so have [proxy](https://microk8s.io/docs/install-proxy) at your pocket)

   ```sh
   kubectl apply -f k1s-deployment.yaml
   kubectl get deployment
   kubectl get pod
   ```

   To visualize the usage of `health-check` you can use `/die` endpoint to make `health-check` fail. and then check what happened with `kubectl get events --watch` to watch the whole process as it happens.

3. Create Service (make sure you did `microk8s enable dns`)

   ```sh
   kubectl apply -f k1s-svc.yaml
   kubectl get svc
   ```

   `k1s` service is used to associate a name for `k1s-deployment` pod's IP addresses inside cluster.
   You can create a pod and access k1s through it.

   ```sh
   kubectl run alpine -ti --image alpine --rm --restart=Never -- sh
   > apk add curl
   > curl htpp://k1s:1378
   ```

   By running this command several times, you can see that this service is working also as a simple `Load Balancer`.
   **Note: make sure `DNS` service is running for k8s.**

4. Create Ingress

   ```sh
   kubectl apply -f k1s-ingress.yaml
   ```

   > [Ingress](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.18/#ingress-v1beta1-networking-k8s-io) exposes HTTP and HTTPS routes from outside the cluster to [services](https://kubernetes.io/docs/concepts/services-networking/service/) within the cluster.

   It simply provides a way to access your pods from Internet by specifying which requests are destined for your `service`.
   **For ingress to work you must first add a [ingress controller](https://kubernetes.io/docs/concepts/services-networking/ingress-controllers). Also `DNS` service must be enabled**
