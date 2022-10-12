package make

import (
	"dtweave.io/zookeeper-operator/api/v1alpha1"
	"dtweave.io/zookeeper-operator/pkg/utils"
	"fmt"
	"github.com/gogo/protobuf/proto"
	v1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/resource"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

const (
	DATA_VOLUME_NAME     = "data-volume"
	DATA_LOG_VOLUME_NAME = "data-log-volume"
	DATA_PVC_NAME        = "data-pvc"
	DATA_LOG_PVC_NAME    = "data-log-pvc"
	CONFIG_NAME          = "config"
)

func StatefulSet(instance *v1alpha1.Zookeeper) *v1.StatefulSet {
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
							Resources:       instance.Spec.Resources,
							Env:             GetEnv(instance.Spec.Conf, instance.Spec.ExtraEnvVars),
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
									// "'/bin/bash', '-c', 'echo \"ruok\" | nc -w %d localhost %d | grep imok'"
									Exec: &corev1.ExecAction{Command: []string{
										"sh",
										"-c",
										fmt.Sprintf("echo ruok | nc -w %d localhost %d | grep imok",
											instance.Spec.Readiness.ProbeCommandTimeout, instance.Spec.ContainerPorts.Client)}},
								},
							},
							LivenessProbe: &corev1.Probe{
								InitialDelaySeconds: instance.Spec.Liveness.InitialDelaySeconds,
								PeriodSeconds:       instance.Spec.Liveness.PeriodSeconds,
								TimeoutSeconds:      instance.Spec.Liveness.TimeoutSeconds,
								FailureThreshold:    instance.Spec.Liveness.FailureThreshold,
								ProbeHandler: corev1.ProbeHandler{
									Exec: &corev1.ExecAction{Command: []string{
										"sh",
										"-c",
										fmt.Sprintf("echo \"ruok\" | nc -w %d localhost %d | grep imok",
											instance.Spec.Liveness.ProbeCommandTimeout, instance.Spec.ContainerPorts.Client)}},
								},
							},
							VolumeMounts: append([]corev1.VolumeMount{
								{
									Name:      CONFIG_NAME,
									MountPath: "/opt/dtweave/zookeeper/conf/zoo.cfg",
									SubPath:   "zoo.cfg",
								},
								{
									Name:      DATA_VOLUME_NAME,
									MountPath: instance.Spec.Conf.DataDir,
								},
								{
									Name:      DATA_LOG_VOLUME_NAME,
									MountPath: instance.Spec.Conf.DataLogDir,
								},
							}, instance.Spec.ExtraVolumeMounts...),
						},
					},
					Volumes: GetVolumes(instance),
				},
			},
			VolumeClaimTemplates: getPvcTemplate(instance),
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
	if nil != instance.Spec.Image.PullSecrets {
		statefulSet.Spec.Template.Spec.ImagePullSecrets = instance.Spec.Image.PullSecrets
	}

	return statefulSet
}

// GetEnv get customer env
func GetEnv(conf v1alpha1.ZookeeperConf, extraEnv []corev1.EnvVar) []corev1.EnvVar {
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
	if nil != extraEnv && len(extraEnv) > 0 {
		envs = append(envs, extraEnv...)
	}

	return envs
}

// GetVolumes get customer volumes and extra volumes
func GetVolumes(instance *v1alpha1.Zookeeper) []corev1.Volume {
	volumes := []corev1.Volume{}

	getCMVolume := func(name string, mapName string) corev1.Volume {
		return corev1.Volume{
			Name: name,
			VolumeSource: corev1.VolumeSource{
				ConfigMap: &corev1.ConfigMapVolumeSource{
					LocalObjectReference: corev1.LocalObjectReference{
						Name: mapName,
					},
				},
			},
		}
	}
	// get volume by pvc
	getPvcVolume := func(name string, claimName string) corev1.Volume {
		return corev1.Volume{
			Name: name,
			VolumeSource: corev1.VolumeSource{
				PersistentVolumeClaim: &corev1.PersistentVolumeClaimVolumeSource{
					ClaimName: claimName,
				},
			},
		}
	}
	// get volume by empty dir
	getEmptyVolume := func(name string) corev1.Volume {
		return corev1.Volume{
			Name: name,
			VolumeSource: corev1.VolumeSource{
				EmptyDir: &corev1.EmptyDirVolumeSource{},
			},
		}
	}

	// check exists configMap
	name := instance.ConfigMapName()
	if "" != instance.Spec.Conf.ExistingCfgConfigmap {
		name = instance.Spec.Conf.ExistingCfgConfigmap
	}
	volumes = append(volumes, getCMVolume(CONFIG_NAME, name))

	// if persistence set pvc else set empty dir
	if instance.Spec.Persistence.Enabled {
		// check exists data pvc
		if "" != instance.Spec.Persistence.Data.ExistingClaim {
			volumes = append(volumes, getPvcVolume(DATA_VOLUME_NAME, instance.Spec.Persistence.Data.ExistingClaim))
		}
		// check exists data log pvc
		if "" != instance.Spec.Persistence.Data.ExistingClaim {
			volumes = append(volumes, getPvcVolume(DATA_LOG_VOLUME_NAME, instance.Spec.Persistence.DataLog.ExistingClaim))
		}
	} else {
		volumes = append(volumes,
			getEmptyVolume(DATA_VOLUME_NAME),
			getEmptyVolume(DATA_LOG_VOLUME_NAME),
		)
	}
	if len(instance.Spec.ExtraVolumes) > 0 {
		volumes = append(volumes, instance.Spec.ExtraVolumes...)
	}

	return volumes
}

func getPvcTemplate(instance *v1alpha1.Zookeeper) []corev1.PersistentVolumeClaim {
	templates := []corev1.PersistentVolumeClaim{}

	if instance.Spec.Persistence.Enabled {
		// no customization found, then use template to create pvc
		if "" == instance.Spec.Persistence.Data.ExistingClaim {
			dataTemplate := corev1.PersistentVolumeClaim{
				ObjectMeta: metav1.ObjectMeta{
					Name:        DATA_VOLUME_NAME,
					Annotations: utils.MergeMaps(instance.Spec.Persistence.Annotation, instance.Spec.CommonAnnotations),
					Labels:      instance.Spec.CommonLabels,
				},
				Spec: corev1.PersistentVolumeClaimSpec{
					AccessModes: []corev1.PersistentVolumeAccessMode{
						instance.Spec.Persistence.AccessModes,
					},
					Resources: corev1.ResourceRequirements{
						Requests: map[corev1.ResourceName]resource.Quantity{
							"storage": resource.MustParse(instance.Spec.Persistence.Data.Size),
						},
					},
					Selector: instance.Spec.Persistence.Data.Selector,
				},
			}
			if nil != instance.Spec.Persistence.StorageClassName {
				dataTemplate.Spec.StorageClassName = instance.Spec.Persistence.StorageClassName
			}

			templates = append(templates, dataTemplate)
		}

		if "" == instance.Spec.Persistence.DataLog.ExistingClaim {
			dataLogTemplate := corev1.PersistentVolumeClaim{
				ObjectMeta: metav1.ObjectMeta{
					Name:        DATA_LOG_VOLUME_NAME,
					Annotations: utils.MergeMaps(instance.Spec.Persistence.Annotation, instance.Spec.CommonAnnotations),
					Labels:      instance.Spec.CommonLabels,
				},
				Spec: corev1.PersistentVolumeClaimSpec{
					AccessModes: []corev1.PersistentVolumeAccessMode{
						instance.Spec.Persistence.AccessModes,
					},
					Resources: corev1.ResourceRequirements{
						Requests: map[corev1.ResourceName]resource.Quantity{
							corev1.ResourceStorage: resource.MustParse(instance.Spec.Persistence.DataLog.Size),
						},
					},
					Selector: instance.Spec.Persistence.DataLog.Selector,
				},
			}
			if nil != instance.Spec.Persistence.StorageClassName {
				dataLogTemplate.Spec.StorageClassName = instance.Spec.Persistence.StorageClassName
			}

			templates = append(templates, dataLogTemplate)
		}
	}

	return templates
}
