# k1s
## Introduction
Kubernetes is an awesome platform and I want to have fun with it.
In this repository I have created an example Golang server and then created pod, service and etc. for it.

## Step by Step
1. Create ReplicaSet

```sh
kubectl apply -f k1s-rs.yaml
kubectl get rs
kubectl get pod
```

2. Create Service

```sh
kubectl apply -f k1-svc.yaml
kubectl get svc
```
