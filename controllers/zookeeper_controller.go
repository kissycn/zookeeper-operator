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
	hadoopv1alpha1 "dtweave.io/zookeeper-operator/api/v1alpha1"
	make2 "dtweave.io/zookeeper-operator/pkg/make"
	"fmt"
	"github.com/go-logr/logr"
	v1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/types"
	"reflect"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

// ZookeeperReconciler reconciles a Zookeeper object
type ZookeeperReconciler struct {
	client.Client
	Scheme *runtime.Scheme
	log    logr.Logger
}

//+kubebuilder:rbac:groups=hadoop.dtweave.io,resources=zookeepers,verbs=get;list;watch;create;update;patch;delete
//+kubebuilder:rbac:groups=hadoop.dtweave.io,resources=zookeepers/status,verbs=get;update;patch
//+kubebuilder:rbac:groups=hadoop.dtweave.io,resources=zookeepers/finalizers,verbs=update

// Reconcile is part of the main kubernetes reconciliation loop which aims to
// move the current state of the cluster closer to the desired state.
// TODO(user): Modify the Reconcile function to compare the state specified by
// the Zookeeper object against the actual cluster state, and then
// perform operations to make the cluster state reflect the state specified by
// the user.
//
// For more details, check Reconcile and its Result here:
// - https://pkg.go.dev/sigs.k8s.io/controller-runtime@v0.12.2/pkg/reconcile
func (r *ZookeeperReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	r.log = ctrl.Log.WithName("controller")

	var instance hadoopv1alpha1.Zookeeper
	err := r.Client.Get(ctx, req.NamespacedName, &instance)
	if err != nil {
		if errors.IsNotFound(err) {
			// Request object not found, could have been deleted after reconcile request - return and don't requeue:
			return ctrl.Result{}, nil
		}
		// Error reading the object - requeue the request:
		return ctrl.Result{}, err
	}

	for _, f := range []reconcileFunc{
		r.reconcileConfigMap,
		r.reconcileStatefulSet,
	} {
		if err = f(ctx, &instance); err != nil {
			return ctrl.Result{}, err
		}
	}

	return ctrl.Result{}, nil
}

// SetupWithManager sets up the controller with the Manager.
func (r *ZookeeperReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&hadoopv1alpha1.Zookeeper{}).
		Complete(r)
}

type reconcileFunc func(ctx context.Context, zookeeper *hadoopv1alpha1.Zookeeper) error

func (r *ZookeeperReconciler) reconcileFinalizers(ctx context.Context, zookeeper *hadoopv1alpha1.Zookeeper) (err error) {
	// TODO
	return nil
}

func (r *ZookeeperReconciler) reconcileConfigMap(ctx context.Context, zookeeper *hadoopv1alpha1.Zookeeper) (err error) {
	var foundCM corev1.ConfigMap
	newCM, err := make2.Configmap(zookeeper)
	if err != nil {
		return err
	}
	err = ctrl.SetControllerReference(zookeeper, newCM, r.Scheme)
	if err != nil {
		return err
	}

	err = r.Client.Get(ctx, types.NamespacedName{
		Name:      newCM.Name,
		Namespace: zookeeper.Namespace,
	}, &foundCM)

	if err != nil && errors.IsNotFound(err) {
		//r.log.Info("Creating a zookeeper configmap ", zap.String("Name", zookeeper.Name), zap.String("Namespace", zookeeper.Namespace))
		//r.log.Info("Creating a zookeeper configmap Name:", zookeeper.Name, " Namespace:", zookeeper.Namespace)
		fmt.Println("Creating a zookeeper configmap Name:", zookeeper.Name, " Namespace:", zookeeper.Namespace)
		err = r.Client.Create(ctx, newCM)
		if err != nil {
			return err
		}
	} else if err != nil {
		return err
	} else {
		if !reflect.DeepEqual(foundCM.Data, newCM.Data) {
			foundCM.Data = newCM.Data
			foundCM.BinaryData = newCM.BinaryData
			err = r.Client.Update(ctx, &foundCM)
			if err != nil {
				return err
			}
		} else {
			fmt.Println("not equal!!!")
		}
	}
	return nil
}

func (r *ZookeeperReconciler) reconcileStatefulSet(ctx context.Context, zookeeper *hadoopv1alpha1.Zookeeper) error {
	var foundSts v1.StatefulSet
	sts := make2.StatefulSet(zookeeper)
	err := r.Client.Get(ctx, types.NamespacedName{
		Name:      zookeeper.Name,
		Namespace: zookeeper.Namespace,
	}, &foundSts)

	if nil != err && errors.IsNotFound(err) {
		err := ctrl.SetControllerReference(zookeeper, sts, r.Scheme)
		if err != nil {
			return err
		}

		r.log.Info("Creating a new Zookeeper StatefulSet", "StatefulSet.Namespace", zookeeper.Namespace, "StatefulSet.Name", zookeeper.Name)
		err = r.Client.Create(ctx, sts)
		if err != nil {
			return err
		}
	} else if nil != err {
		return err
	} else {
		r.log.Info("Updating StatefulSet",
			"StatefulSet.Namespace", foundSts.Namespace,
			"StatefulSet.Name", foundSts.Name)
		foundSts.Spec.Template = sts.Spec.Template

		foundSts.Spec.Replicas = sts.Spec.Replicas
		foundSts.Spec.Template = sts.Spec.Template
		foundSts.Spec.PodManagementPolicy = sts.Spec.PodManagementPolicy

		err := r.Client.Update(ctx, &foundSts)
		if err != nil {
			return err
		}

		// TODO status and replicas ready update
	}

	return nil
}

func (r *ZookeeperReconciler) reconcileService(instance *hadoopv1alpha1.Zookeeper) error {
	return nil
}
