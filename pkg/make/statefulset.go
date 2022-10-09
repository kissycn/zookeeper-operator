package make

import (
	"dtweave.io/zookeeper-operator/api/v1alpha1"
	"dtweave.io/zookeeper-operator/pkg/utils"
	"fmt"
	"github.com/gogo/protobuf/proto"
	v1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func StatefulSet(instance *v1alpha1.Zookeeper) {
	matchLabels := map[string]string{
		"hadoop.dtweave.io/component": "zookeeper",
		"hadoop.dtweave.io/app":       instance.Name,
	}

	statefulSet := &v1.StatefulSet{
		ObjectMeta: metav1.ObjectMeta{
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
			Selector: &metav1.LabelSelector{
				MatchLabels: matchLabels,
			},
			ServiceName: instance.Name + "-headless",
			Template: corev1.PodTemplateSpec{
				ObjectMeta: metav1.ObjectMeta{
					Labels:      utils.MergeMaps(matchLabels, instance.Spec.PodLabels),
					Annotations: instance.Spec.PodAnnotations,
				},
				Spec: corev1.PodSpec{
					// TODO
					ServiceAccountName: "",
					SecurityContext: &corev1.PodSecurityContext{
						RunAsUser:    proto.Int64(1000),
						RunAsGroup:   proto.Int64(1000),
						RunAsNonRoot: proto.Bool(true),
					},
					Containers: []corev1.Container{
						{
							Name:            "zookeeper",
							Image:           instance.Image(),
							ImagePullPolicy: instance.Spec.Image.PullPolicy,
							// TODO webhook set default value
							Resources: instance.Spec.Resources,
							Env:       GetEnv(instance.Spec.Conf, instance.Spec.ExtraEnvVars),
							Ports: []corev1.ContainerPort{
								{
									Name:          "client",
									ContainerPort: instance.Spec.ContainerPorts.Client,
								},
								{
									Name:          "follower",
									ContainerPort: instance.Spec.ContainerPorts.Follower,
								},
								{
									Name:          "election",
									ContainerPort: instance.Spec.ContainerPorts.Election,
								},
								{
									Name:          "tls",
									ContainerPort: instance.Spec.ContainerPorts.Tls,
								},
							},
							ReadinessProbe: &corev1.Probe{
								InitialDelaySeconds: instance.Spec.Readiness.InitialDelaySeconds,
								PeriodSeconds:       instance.Spec.Readiness.PeriodSeconds,
								TimeoutSeconds:      instance.Spec.Readiness.TimeoutSeconds,
								FailureThreshold:    instance.Spec.Readiness.FailureThreshold,
								SuccessThreshold:    instance.Spec.Readiness.SuccessThreshold,
								ProbeHandler: corev1.ProbeHandler{
									Exec: &corev1.ExecAction{Command: []string{fmt.Sprintf("'/bin/bash', '-c', 'echo \"ruok\" | nc -w localhost %d | grep imok'", instance.Spec.ContainerPorts.Client)}},
								},
							},
							LivenessProbe: &corev1.Probe{
								InitialDelaySeconds: instance.Spec.Readiness.InitialDelaySeconds,
								PeriodSeconds:       instance.Spec.Readiness.PeriodSeconds,
								TimeoutSeconds:      instance.Spec.Readiness.TimeoutSeconds,
								FailureThreshold:    instance.Spec.Readiness.FailureThreshold,
								ProbeHandler: corev1.ProbeHandler{
									Exec: &corev1.ExecAction{Command: []string{fmt.Sprintf("'/bin/bash', '-c', 'echo \"ruok\" | nc -w localhost %d | grep imok'", instance.Spec.ContainerPorts.Client)}},
								},
							},
						},
					},
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

}

func GetEnv(conf v1alpha1.ZookeeperConf, params []corev1.EnvVar) []corev1.EnvVar {
	envs := []corev1.EnvVar{
		{
			Name:  "ZOO_DATA_DIR",
			Value: conf.DataDir,
		},
		{
			Name:  "ZOO_DATA_LOG_DIR",
			Value: conf.DataLogDir,
		},
	}
	if nil != params && len(params) > 0 {
		envs = append(envs, params...)
	}

	return envs
}
