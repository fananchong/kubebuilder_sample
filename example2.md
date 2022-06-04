# example2

example2 演示以下内容：
1. 如何定义、使用 Status 结构

## 生成 example2 模板

```shell
mkdir -p example2
pushd example2
go mod init example2
kubebuilder init --domain=fananchong.com
kubebuilder create api --group demo --version v1 --kind Example2 --resource true --controller true --namespaced true
popd
```

## 填写 api/v1/example1_types.go

目标，能更新状态值

对应 api/v1/example2_types.go 文件修改：

```go
type Example2Status struct {
	// +optional
	CustomStatus1 string `json:"customStatus1,omitempty"`
	// +optional
	CustomStatus2 *int32 `json:"customStatus2,omitempty"`
}
```

然后执行： 
```shell
make manifests
```

会自动生成对应的配置到 config 目录下 `config/crd/bases/demo.fananchong.com_example2s.yaml`


## 填写 controllers/example1_controller.go


```go
func (r *Example2Reconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	_ = log.FromContext(ctx)
	example := &v1.Example2{}
	err := r.Get(ctx, req.NamespacedName, example)
	if err != nil {
		return ctrl.Result{}, client.IgnoreNotFound(err)
	}
	example.Status.CustomStatus1 = "xxxxxxxxxx"
	if example.Status.CustomStatus2 == nil {
		example.Status.CustomStatus2 = new(int32)
	}
	*(example.Status.CustomStatus2) = 1111111
	if err := r.Status().Update(ctx, example); err != nil {
		return ctrl.Result{}, err
	}
	return ctrl.Result{}, nil
}
```

代码分析：
- r.Status().Update(ctx, example) 更新 example2 资源的状态


## 安装 CRD Example2

```
make install
kubectl api-resources -o wide | grep example2
kubectl get crd | grep example2
```


## 调试验证

1. 执行 example2 ，能实时查看 CRD 的 log
    ```shell
	make build
    ./bin/manager --metrics-bind-address=":7070" --health-probe-bind-address=":7071"
    ```

2. 使用这个 CRD
    ```shell
    kubectl apply -f config/samples/demo_v1_example2.yaml
    kubectl get example2s -o yaml
	kubectl describe -f config/samples/demo_v1_example2.yaml
    kubectl delete -f config/samples/demo_v1_example2.yaml
    ```