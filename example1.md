# example1

example1 主要过一遍制作流程：
1. 如何填写 api/v1/example1_types.go ，简单定义一个 GVK
2. 如何填写 controllers/example1_controller.go ，这里打印下信息
3. 如何安装 Example1 这个 CRD
4. 如何调试验证


## 生成 example1 模板

```shell
mkdir -p example1
pushd example1
go mod init example1
kubebuilder init
kubebuilder create api --group demo --version v1 --kind Example1 --resource true --controller true --namespaced true
popd
```

## 填写 api/v1/example1_types.go

这里的例子，定义一个简单 CRD ，主要过一遍流程

目标，能正常 apply/delete 这个 CRD：

```yaml
apiVersion: demo.my.domain/v1
kind: Example1
metadata:
  name: example-1
  namespace: default
spec:
  custom1: xxxx
  custom2: 10
```


对应 api/v1/example1_types.go 文件修改：

```go
type Example1Spec struct {
	//+kubebuilder:validation:Required
	//+kubebuilder:validation:Type=string
	Custom1 string `json:"custom1,omitempty"`

	//+kubebuilder:validation:Required
	//+kubebuilder:validation:Type=integer
	//+kubebuilder:validation:Minimum=0
	Custom2 *int32 `json:"custom2,omitempty"`
}
```

然后执行： 
```shell
make manifests
```

会自动生成对应的配置到 config 目录下 `config/crd/bases/demo.my.domain_example1s.yaml`

注释，如：
```go
//+kubebuilder:validation:Required
//+kubebuilder:validation:Type=string
```
是可以让 K8S 做些字段检查。参考文档： [https://book.kubebuilder.io/reference/markers/crd-validation.html](https://book.kubebuilder.io/reference/markers/crd-validation.html)


## 填写 controllers/example1_controller.go


同样目的是过一遍流程，主要功能是简单的打印下日志：

```go
func (r *Example1Reconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	logger := log.FromContext(ctx)
	example := &v1.Example1{}
	err := r.Get(ctx, req.NamespacedName, example)
	if err != nil {
		if errors.IsNotFound(err) {
			logger.Info(fmt.Sprintf("[delete] Namespace:%v Name:%v", req.Namespace, req.Name))
			return ctrl.Result{}, nil
		}
		logger.Error(err, "Error occurred while fetching the resource")
		return ctrl.Result{}, err
	}
	logger.Info(fmt.Sprintf("[apply] Namespace:%v Name:%v custom1:%v custom2:%v", example.Namespace, example.Name, example.Spec.Custom1, example.Spec.Custom2))
	return ctrl.Result{}, nil
}
```

代码分析：
- r.Get 获取 CRD 对象


## 安装 CRD Example1

```
make install
kubectl api-resources -o wide | grep demo
kubectl get crd
```


## 调试验证

1. 执行 example1 ，能实时查看 CRD 的 log
    ```shell
	go build
    ./example1 --metrics-bind-address=":7070" --health-probe-bind-address=":7071"
    ```

2. 使用这个 CRD
    ```shell
    kubectl apply -f example1.yaml
    kubectl get example1s
    kubectl get example1s.demo.my.domain
    kubectl delete -f example1.yaml
    ```