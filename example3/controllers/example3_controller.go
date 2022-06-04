/*
Copyright 2022.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package controllers

import (
	"context"

	"github.com/pingcap/errors"
	kapps "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/types"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/controller/controllerutil"
	"sigs.k8s.io/controller-runtime/pkg/log"

	demov1 "example3/api/v1"
)

// Example3Reconciler reconciles a Example3 object
type Example3Reconciler struct {
	client.Client
	Scheme *runtime.Scheme
}

//+kubebuilder:rbac:groups=demo.fananchong.com,resources=example3s,verbs=get;list;watch;create;update;patch;delete
//+kubebuilder:rbac:groups=demo.fananchong.com,resources=example3s/status,verbs=get;update;patch
//+kubebuilder:rbac:groups=demo.fananchong.com,resources=example3s/finalizers,verbs=update
//+kubebuilder:rbac:groups=apps,resources=deployments,verbs=get;list;watch;create;update;patch;delete
//+kubebuilder:rbac:groups=apps,resources=deployments/status,verbs=get

// Reconcile is part of the main kubernetes reconciliation loop which aims to
// move the current state of the cluster closer to the desired state.
// TODO(user): Modify the Reconcile function to compare the state specified by
// the Example3 object against the actual cluster state, and then
// perform operations to make the cluster state reflect the state specified by
// the user.
//
// For more details, check Reconcile and its Result here:
// - https://pkg.go.dev/sigs.k8s.io/controller-runtime@v0.11.2/pkg/reconcile
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

// SetupWithManager sets up the controller with the Manager.
func (r *Example3Reconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&demov1.Example3{}).
		Owns(&kapps.Deployment{}).
		Complete(r)
}

func getDeploymentDef(example *demov1.Example3) *kapps.Deployment {
	return &kapps.Deployment{
		ObjectMeta: metav1.ObjectMeta{
			Namespace: example.Namespace,
			Name:      example.Name,
		},
		Spec: kapps.DeploymentSpec{
			Selector: &metav1.LabelSelector{
				MatchLabels: map[string]string{
					"app": "example-app",
				},
			},
			Template: corev1.PodTemplateSpec{
				ObjectMeta: metav1.ObjectMeta{
					Labels: map[string]string{
						"app": "example-app",
					},
				},
				Spec: corev1.PodSpec{
					Containers: []corev1.Container{
						{
							Name:  "example-app",
							Image: "redis",
						},
					},
				},
			},
		},
	}
}
