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
	"fmt"

	"github.com/pingcap/errors"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/types"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/log"

	demov1 "example1/api/v1"
	v1 "example1/api/v1"
)

// Example1Reconciler reconciles a Example1 object
type Example1Reconciler struct {
	client.Client
	Scheme *runtime.Scheme
}

//+kubebuilder:rbac:groups=demo.my.domain,resources=example1s,verbs=get;list;watch;create;update;patch;delete
//+kubebuilder:rbac:groups=demo.my.domain,resources=example1s/status,verbs=get;update;patch
//+kubebuilder:rbac:groups=demo.my.domain,resources=example1s/finalizers,verbs=update

// Reconcile is part of the main kubernetes reconciliation loop which aims to
// move the current state of the cluster closer to the desired state.
// TODO(user): Modify the Reconcile function to compare the state specified by
// the Example1 object against the actual cluster state, and then
// perform operations to make the cluster state reflect the state specified by
// the user.
//
// For more details, check Reconcile and its Result here:
// - https://pkg.go.dev/sigs.k8s.io/controller-runtime@v0.11.2/pkg/reconcile
func (r *Example1Reconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	logger := log.FromContext(ctx)
	example := &v1.Example1{}
	err := r.Get(ctx, types.NamespacedName{Name: req.Name, Namespace: req.Namespace}, example)
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

// SetupWithManager sets up the controller with the Manager.
func (r *Example1Reconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&demov1.Example1{}).
		Complete(r)
}
