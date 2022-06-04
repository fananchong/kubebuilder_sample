# kubebuilder_sample

一些使用 kubebuilder 自定义 CRD 的例子

| 例子                      | 说明                         |
| :------------------------ | :--------------------------- |
| [example1](./example1.md) | 演示制作流程、 Spec 定义实现 |
| [example2](./example2.md) | Status 定义实现              |
| [example3](./example3.md) | 控制 K8S 内置资源            |

## 依赖环境


| 工具        | 安装                                                                                                                           |
| :---------- | :----------------------------------------------------------------------------------------------------------------------------- |
| kubectl     | [https://kubernetes.io/docs/tasks/tools/install-kubectl-linux/](https://kubernetes.io/docs/tasks/tools/install-kubectl-linux/) |
| minikube    | [https://minikube.sigs.k8s.io/docs/start/](https://minikube.sigs.k8s.io/docs/start/)                                           |
| kubebuilder | [https://book.kubebuilder.io/quick-start.html#installation](https://book.kubebuilder.io/quick-start.html#installation)         |
| golang      | [https://golang.google.cn/doc/install](https://golang.google.cn/doc/install)                                                   |

最新的安装方法，参考上面 URL 。当前时间点：


**kubectl**

```shell
curl -LO "https://dl.k8s.io/release/$(curl -L -s https://dl.k8s.io/release/stable.txt)/bin/linux/amd64/kubectl"
sudo install -o root -g root -m 0755 kubectl /usr/local/bin/kubectl
kubectl version --client
```


**minikube**

```shell
curl -LO https://storage.googleapis.com/minikube/releases/latest/minikube-linux-amd64
sudo install minikube-linux-amd64 /usr/local/bin/minikube
minikube start
kubectl get po -A
```

**kubebuilder**

```shell
curl -L -o kubebuilder https://go.kubebuilder.io/dl/latest/$(go env GOOS)/$(go env GOARCH)
chmod +x kubebuilder && mv kubebuilder /usr/local/bin/
```


**golang**

```shell
rm -rf /usr/local/go && tar -C /usr/local -xzf go1.18.3.linux-amd64.tar.gz
export PATH=$PATH:/usr/local/go/bin
go version
```