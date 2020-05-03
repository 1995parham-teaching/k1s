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

1. Create Deployment

```sh
kubectl apply -f k1-deployment.yaml
kubectl get deployment
kubectl get pod
```

2. Create Service

```sh
kubectl apply -f k1-svc.yaml
kubectl get svc
```
