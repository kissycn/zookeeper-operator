package make

import (
	"dtweave.io/zookeeper-operator/api/v1alpha1"
	"dtweave.io/zookeeper-operator/pkg/utils"
	v1 "k8s.io/api/apps/v1"
	v13 "k8s.io/api/core/v1"
	v12 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func StatefulSet(instance *v1alpha1.Zookeeper) {
	matchLabels := map[string]string{
		"hadoop.dtweave.io/component": "zookeeper",
		"hadoop.dtweave.io/app":       instance.Name,
	}

	statefulSet := &v1.StatefulSet{
		ObjectMeta: v12.ObjectMeta{
			Name:      instance.Name,
			Namespace: instance.Namespace,
			Labels: utils.MergeMaps(
				instance.Spec.CommonLabels,
				matchLabels,
			),
		},
		Spec: v1.StatefulSetSpec{
			Replicas:            instance.Spec.ReplicasCount,
			PodManagementPolicy: instance.Spec.PodManagementPolicy,
			Selector: &v12.LabelSelector{
				MatchLabels: matchLabels,
			},
			ServiceName: instance.Name + "-headless",
			Template: v13.PodTemplateSpec{
				ObjectMeta: v12.ObjectMeta{
					Labels:      utils.MergeMaps(matchLabels, instance.Spec.PodLabels),
					Annotations: instance.Spec.PodAnnotations,
				},
				Spec: v13.PodSpec{
					// TODO
					ServiceAccountName: "",
				},
			},
		},
	}

	if nil != instance.Spec.CommonAnnotations {
		statefulSet.ObjectMeta.Annotations = instance.Spec.CommonAnnotations
	}
	if nil != instance.Spec.Affinity {
		statefulSet.Spec.Template.Spec.Affinity = instance.Spec.Affinity
	}
	if nil != instance.Spec.Tolerations {
		statefulSet.Spec.Template.Spec.Tolerations = instance.Spec.Tolerations
	}
	if nil != instance.Spec.NodeSelector {
		statefulSet.Spec.Template.Spec.NodeSelector = instance.Spec.NodeSelector
	}
	if nil != instance.Spec.HostAlias {
		statefulSet.Spec.Template.Spec.HostAliases = instance.Spec.HostAlias
	}
	if "" != instance.Spec.PriorityClassName {
		statefulSet.Spec.Template.Spec.PriorityClassName = instance.Spec.PriorityClassName
	}
	if "" != instance.Spec.SchedulerName {
		statefulSet.Spec.Template.Spec.SchedulerName = instance.Spec.SchedulerName
	}
	if nil != instance.Spec.PodSecurityContext {
		statefulSet.Spec.Template.Spec.SecurityContext = instance.Spec.PodSecurityContext
	}
}
