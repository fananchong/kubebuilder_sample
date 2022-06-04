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
	InstanceNum *int32 `json:"instanceNum,omitempty"`
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


需要修改 3 处：

**1. 添加访问 Deployment 资源权限**
```go
//+kubebuilder:rbac:groups=apps,resources=deployments,verbs=get;list;watch;create;update;patch;delete
//+kubebuilder:rbac:groups=apps,resources=deployments/status,verbs=get
```


**2. 实现 Reconcile 函数**

```go
func (r *Example3Reconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	log := log.FromContext(ctx)

	// 1. 获取 Example3 CRD 信息
	var example demov1.Example3
	if err := r.Get(ctx, req.NamespacedName, &example); err != nil {
		return ctrl.Result{}, client.IgnoreNotFound(err)
	}

	// 2. 关联 Deployment 资源定义
	deployment := getDeploymentDef(&example)
	deployment.Spec.Replicas = example.Spec.InstanceNum
	if err := controllerutil.SetControllerReference(&example, deployment, r.Scheme); err != nil {
		return ctrl.Result{}, err
	}

	// 3. 扩容 Deployment 达到预期
	foundDeployment := &kapps.Deployment{}
	err := r.Get(ctx, types.NamespacedName{Name: deployment.Name, Namespace: deployment.Namespace}, foundDeployment)
	if err != nil && errors.IsNotFound(err) {
		log.V(1).Info("Creating Deployment", "deployment", deployment.Name)
		err = r.Create(ctx, deployment)
	} else if err == nil {
		if foundDeployment.Spec.Replicas != deployment.Spec.Replicas {
			foundDeployment.Spec.Replicas = deployment.Spec.Replicas
			log.V(1).Info("Updating Deployment", "deployment", deployment.Name)
			err = r.Update(ctx, foundDeployment)
		}
	}

	// 4. 更新 Example3 CRD 状态（略）

	return ctrl.Result{}, err
}
```

代码分析，看注释

**3. 调整下 SetupWithManager 函数**

```go
func (r *Example3Reconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&demov1.Example3{}).
		Owns(&kapps.Deployment{}).
		Complete(r)
}
```

代码分析：
- Owns(&kapps.Deployment{}) 需要添加这行，才能操作 Deployment 对象


**参考资料**

该例子官方文档： [https://book.kubebuilder.io/reference/watching-resources/operator-managed.html](https://book.kubebuilder.io/reference/watching-resources/operator-managed.html)

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
    kubectl get po -A | grep example3
    ```