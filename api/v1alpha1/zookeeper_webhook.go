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

package v1alpha1

import (
	v1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/resource"
	"k8s.io/apimachinery/pkg/runtime"
	"reflect"
	ctrl "sigs.k8s.io/controller-runtime"
	logf "sigs.k8s.io/controller-runtime/pkg/log"
	"sigs.k8s.io/controller-runtime/pkg/webhook"
)

// log is for logging in this package.
var zookeeperlog = logf.Log.WithName("zookeeper-resource")

func (r *Zookeeper) SetupWebhookWithManager(mgr ctrl.Manager) error {
	return ctrl.NewWebhookManagedBy(mgr).
		For(r).
		Complete()
}

// TODO(user): EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!

//+kubebuilder:webhook:path=/mutate-hadoop-dtweave-io-v1alpha1-zookeeper,mutating=true,failurePolicy=fail,sideEffects=None,groups=hadoop.dtweave.io,resources=zookeepers,verbs=create;update,versions=v1alpha1,name=mzookeeper.kb.io,admissionReviewVersions=v1

var _ webhook.Defaulter = &Zookeeper{}

// Default implements webhook.Defaulter so a webhook will be registered for the type
func (r *Zookeeper) Default() {
	zookeeperlog.Info("default", "name", r.Name)
	if reflect.DeepEqual(v1.ResourceRequirements{}, r.Spec.Resources) {
		r.Spec.Resources = v1.ResourceRequirements{
			Requests: map[v1.ResourceName]resource.Quantity{},
			Limits:   map[v1.ResourceName]resource.Quantity{},
		}
	}
	// check resource request by cpu
	if r.Spec.Resources.Requests.Cpu().Value() == 0 {
		r.Spec.Resources.Requests[v1.ResourceCPU] = resource.MustParse("250m")
	}
	// check resource request by memory
	if r.Spec.Resources.Requests.Memory().Value() == 0 {
		r.Spec.Resources.Requests[v1.ResourceMemory] = resource.MustParse("512Mi")
	}
	// check resource limit by cpu
	if r.Spec.Resources.Limits.Cpu().Value() == 0 {
		r.Spec.Resources.Limits[v1.ResourceCPU] = resource.MustParse("512m")
	}
	// check resource limit by memory
	if r.Spec.Resources.Limits.Memory().Value() == 0 {
		r.Spec.Resources.Limits[v1.ResourceMemory] = resource.MustParse("1Gi")
	}

	// set readinessProbe default
	if r.Spec.Readiness.InitialDelaySeconds == 0 {
		r.Spec.Readiness.InitialDelaySeconds = 5
	}
}

// TODO(user): change verbs to "verbs=create;update;delete" if you want to enable deletion validation.
//+kubebuilder:webhook:path=/validate-hadoop-dtweave-io-v1alpha1-zookeeper,mutating=false,failurePolicy=fail,sideEffects=None,groups=hadoop.dtweave.io,resources=zookeepers,verbs=create;update,versions=v1alpha1,name=vzookeeper.kb.io,admissionReviewVersions=v1

var _ webhook.Validator = &Zookeeper{}

// ValidateCreate implements webhook.Validator so a webhook will be registered for the type
func (r *Zookeeper) ValidateCreate() error {
	zookeeperlog.Info("validate create", "name", r.Name)

	// TODO(user): fill in your validation logic upon object creation.
	return nil
}

// ValidateUpdate implements webhook.Validator so a webhook will be registered for the type
func (r *Zookeeper) ValidateUpdate(old runtime.Object) error {
	zookeeperlog.Info("validate update", "name", r.Name)

	// TODO(user): fill in your validation logic upon object update.
	return nil
}

// ValidateDelete implements webhook.Validator so a webhook will be registered for the type
func (r *Zookeeper) ValidateDelete() error {
	zookeeperlog.Info("validate delete", "name", r.Name)

	// TODO(user): fill in your validation logic upon object deletion.
	return nil
}
