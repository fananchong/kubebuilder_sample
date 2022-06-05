# example4

example4 演示以下内容：
1. 正式部署 example3 的 CRD manager 

## 部署例子

```
mkdir -p example3
pushd example3
make docker-build docker-push IMG=${REGISTRY}/test111:latest
make deploy IMG=${REGISTRY}/test111:latest
kubectl get po -A | grep example3-controller-manager
popd
```

REGISTRY 变量为 docker 仓库名


## 使用这个 CRD

```shell
kubectl apply -f config/samples/demo_v1_example3.yaml
kubectl get po -A | grep example3-sample
```
