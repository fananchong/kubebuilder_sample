#!/bin/bash

set -ex

mkdir -p example1
pushd example1
go mod init example1
kubebuilder init
kubebuilder create api --group demo --version v1 --kind Example1 --resource true --controller true --namespaced true
popd
