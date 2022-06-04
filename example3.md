# example3

example3 演示以下内容：
1. 使 deployment 资源达到目标数量

## 生成 example3 模板

```shell
mkdir -p example3
pushd example3
go mod init example3
kubebuilder init --domain=fananchong.com
kubebuilder create api --group demo --version v1 --kind Example3 --resource true --controller true --namespaced true
popd
```

## 填写 api/v1/example3_types.go

目标，使 deployment 资源达到 InstanceNum 个

对应 api/v1/example3_types.go 文件修改：

```go
type Example3Spec struct {
	// +kubebuilder:validation:Minimum=1
	InstanceNum *int64 `json:"instanceNum,omitempty"`
}
type Example3Status struct {
	// +optional
	RealInstanceNum *int32 `json:"realInstanceNum,omitempty"`
}
```

然后执行： 
```shell
make manifests
```

会自动生成对应的配置到 config 目录下 `config/crd/bases/demo.fananchong.com_example3s.yaml`


## 填写 controllers/example3_controller.go


```go
func (r *Example2Reconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	_ = log.FromContext(ctx)
	example := &v1.Example3{}
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
- r.Status().Update(ctx, example) 更新 example3 资源的状态


## 安装 CRD Example3

```
make install
kubectl api-resources -o wide | grep example3
kubectl get crd | grep example3
```


## 调试验证

1. 执行 example3 ，能实时查看 CRD 的 log
    ```shell
	make build
    ./bin/manager --metrics-bind-address=":7070" --health-probe-bind-address=":7071"
    ```

2. 使用这个 CRD
    ```shell
    kubectl apply -f config/samples/demo_v1_example3.yaml
    kubectl get example3s -o yaml
	kubectl describe -f config/samples/demo_v1_example3.yaml
    kubectl delete -f config/samples/demo_v1_example3.yaml
    ```