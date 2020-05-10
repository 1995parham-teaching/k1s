# k1s
[![Drone (cloud)](https://img.shields.io/drone/build/1995parham/k1s.svg?style=flat-square)](https://cloud.drone.io/1995parham/k1s)

## Introduction
Kubernetes is an awesome platform, and I want to have fun with it.
In this repository I have created an example Golang server and then created pod, service and etc. for it.
The `hello-server` is a simple HTTP server that we want to deploy on cloud. It must have replica so everyone
can get their hello.

## Step by Step
First of all switch to your desired namespace. Then follow these instructions
to have your `hello-server` up and running.

The `hello-server` application requires a `config` to work. Possible options are providing `config` through a `config.yaml` file, environment variables, or default config.
Passing config files and environment variables is through a `ConfigMap`. To test each way, comment out the other one's way of loading in `k1s-deployment.yaml` 



1. Create ConfigMap

    ```sh
    kubectl apply -f k1s-config-map.yaml
    kubectl get configmaps
    ```

    By default, both `config.yaml` and environment variables are set and `hello-server` prioritizes env variables against config file. For only loading `config.yaml`, delete the `envFrom:` part from `k1s-deployment.yaml`.

2. Create Deployment

    ```sh
    kubectl apply -f k1s-deployment.yaml
    kubectl get deployment
    kubectl get pod
    ```

    To visualize the usage of `health-check` you can use `/die` endpoint to make `health-check` fail. and then check what happened with `kubectl get events`.

3. Create Service

    ```sh
    kubectl apply -f k1s-svc.yaml
    kubectl get svc
    ```

    `k1s` service is used to associate a name for `k1s-deployment` pod's IP addresses inside cluster. You can create a shell session inside one of pods from `kubectl get pods` and then try  :

    ```shell
    apk add curl
    curl http://k1s:1378
    ```

    By running this command several times, you can see that this service is working also as a simple 	`Load Balancer`.
    **Note: make sure `DNS` service is running for k8s.**
    
4. Create Ingress

    ```sh
    kubectl apply -f k1s-ingress.yaml
    ```

    > [Ingress](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.18/#ingress-v1beta1-networking-k8s-io) exposes HTTP and HTTPS routes from outside the cluster to [services](https://kubernetes.io/docs/concepts/services-networking/service/) within the cluster.

    It simply provides a way to access your pods from Internet by specifying which requests are destined for your `service`.
    **For ingress to work you must first add a [ingress controller](https://kubernetes.io/docs/concepts/services-networking/ingress-controllers). Also `DNS` service must be enabled **
