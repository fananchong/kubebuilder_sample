#!/bin/bash

set -ex

mkdir -p example3
pushd example3
go mod init example3
kubebuilder init --domain=fananchong.com
kubebuilder create api --group demo --version v1 --kind Example3 --resource true --controller true --namespaced true
popd
