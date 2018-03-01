#!/bin/bash

SUPEREGO_ROOT="$(dirname "${BASH_SOURCE}")/.."
kubectl create -f $SUPEREGO_ROOT/hack/k8s/deployment.yaml
kubectl create -f $SUPEREGO_ROOT/hack/k8s/service.yaml
kubectl get pods -l app=superego

minikube service superego --url
