#!/bin/bash

set -ex

REGISTRY=fananchong

mkdir -p example3
pushd example3
make docker-build docker-push IMG=${REGISTRY}/test111:latest
make deploy IMG=${REGISTRY}/test111:latest
kubectl get po -A | grep example3-controller-manager
popd
