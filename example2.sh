#!/bin/bash

set -ex

mkdir -p example2
pushd example2
go mod init example2
kubebuilder init --domain=fananchong.com
kubebuilder create api --group demo --version v1 --kind Example2 --resource true --controller true --namespaced true
popd
