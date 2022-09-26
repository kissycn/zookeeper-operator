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
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// ZookeeperSpec defines the desired state of Zookeeper
type ZookeeperSpec struct {
	// Important: Run "make" to regenerate code after modifying this file
	// CommonLabels Add labels to all the deployed resources
	CommonLabels map[string]string `json:"commonLabels,omitempty"`
	// CommonAnnotations Add annotations to all the deployed resources
	CommonAnnotations map[string]string `json:"commonAnnotations,omitempty"`
	// Conf is zookeeper config
	//
	Conf ZookeeperConf `json:"conf,omitempty"`
	// +kubebuilder:default:="fuck"
	DefaultTest string `json:"defaultTest"`
	//Replicas The valid range of size is from 1 to 7.
	//
	// +kubebuilder:validation:Maximum:=7
	// +kubebuilder:validation:Minimum:=1
	// +kubebuilder:default:=1
	// +required
	// +kubebuilder:validation:Required
	ReplicasCount int32 `json:"replicasCount,omitempty"`
	// Labels is pod labels,key-value pairs
	//
	// +kubebuilder:validation:Optional
	Labels map[string]string `json:"labels,omitempty"`
	// Annotations specifies the annotations to attach to pods the operator
	//
	// +kubebuilder:validation:Optional
	Annotations map[string]string `json:"annotations,omitempty"`
	// NodeSelector specifies a map of key-value pairs.
	//
	// +kubebuilder:validation:Optional
	NodeSelector map[string]string `json:"nodeSelector,omitempty"`
	// Image is the  container image.
	//
	Image ContainerImage `json:"image,omitempty"`
	// ContainerPorts
	//
	ContainerPorts ContainerPorts `json:"containerPorts,omitempty"`
	//
	Readiness Probe `json:"readiness,omitempty"`
	//
	Liveness Probe `json:"liveness,omitempty"`
	//Resources is the resource requirements for the container.
	Resources v1.ResourceRequirements `json:"resources,omitempty"`
	// Affinity The scheduling constraints on pods.
	//
	// +kubebuilder:validation:Optional
	Affinity v1.Affinity `json:"affinity,omitempty"`
	//Toleration specifies the pod's
	//
	// +kubebuilder:validation:Optional
	Toleration v1.Toleration `json:"toleration,omitempty"`
	// ExtraVolumeMounts Optionally specify extra list of additional volumeMounts for the ZooKeeper container(s)
	//
	// +kubebuilder:validation:Optional
	ExtraVolumeMounts []v1.VolumeMount `json:"extraVolumeMounts,omitempty"`
	// ExtraVolumes  Optionally specify extra list of additional volumes for the ZooKeeper pod(s)
	//
	// +kubebuilder:validation:Optional
	ExtraVolumes v1.Volume `json:"extraVolumes,omitempty"`
	// InitContainers Add additional init containers to the ZooKeeper pod(s)
	//
	// +kubebuilder:validation:Optional
	InitContainers v1.Container `json:"initContainers,omitempty"`
	// Persistence define Zookeeper persistence
	//
	Persistence ZookeeperPersistence `json:"persistence,omitempty"`
	// Service   Kubernetes Service defines
	//
	Service ServicePolicy `json:"service,omitempty"`
	// TODO containerSecurityContext
	// TODO podSecurityContext
	// TODO  sidecars
}

// ZookeeperStatus defines the observed state of Zookeeper
type ZookeeperStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status

// Zookeeper is the Schema for the zookeepers API
type Zookeeper struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ZookeeperSpec   `json:"spec,omitempty"`
	Status ZookeeperStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// ZookeeperList contains a list of Zookeeper
type ZookeeperList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Zookeeper `json:"items"`
}

func init() {
	SchemeBuilder.Register(&Zookeeper{}, &ZookeeperList{})
}

type ContainerImage struct {
	// Registry ZooKeeper image registry
	// +kubebuilder:default:=docker.io
	Registry string `json:"registry,omitempty"`
	// Repository ZooKeeper image repository
	// +kubebuilder:default:=dtweave/zookeeper
	Repository string `json:"repository,omitempty"`
	// Tag ZooKeeper image tag (immutable tags are recommended)
	// +kubebuilder:default:=v1.0.0
	Tag string `json:"tag,omitempty"`
	// Digest ZooKeeper image digest in the way sha256:aa.... Please note this parameter, if set, will override the tag
	Digest string `json:"digest,omitempty"`
	// PullPolicy ZooKeeper image pull policy
	// +kubebuilder:default:=IfNotPresent
	// +kubebuilder:validation:Enum=Always;Never;IfNotPresent
	PullPolicy v1.PullPolicy `json:"pullPolicy,omitempty"`
	// PullSecrets Specify docker-registry secret names as an array
	PullSecrets []string `json:"pullSecrets,omitempty"`
}

type ContainerPorts struct {
	// Client ZooKeeper client container port
	//
	// +kubebuilder:validation:Maximum:=65535
	// +kubebuilder:validation:Minimum:=1000
	// +kubebuilder:default:=2181
	Client int32 `json:"client,omitempty"`
	// Tls ZooKeeper TLS container port
	//
	// +kubebuilder:validation:Maximum:=65535
	// +kubebuilder:validation:Minimum:=1000
	// +kubebuilder:default:=3181
	Tls int32 `json:"tls,omitempty"`
	// Follower cluster follower connect port
	//
	// +kubebuilder:validation:Maximum:=65535
	// +kubebuilder:validation:Minimum:=1000
	// +kubebuilder:default:=2888
	Follower int32 `json:"follower,omitempty"`
	// Election cluster election port
	//
	// +kubebuilder:validation:Maximum:=65535
	// +kubebuilder:validation:Minimum:=1000
	// +kubebuilder:default:=3888
	Election int32 `json:"election,omitempty"`
}

type Probe struct {
	// Enabled  livenessProbe on ZooKeeper containers
	//
	// +kubebuilder:default:=false
	Enabled bool `json:"enabled,omitempty"`
	// InitialDelaySeconds Initial delay seconds for Probe
	//
	// +kubebuilder:validation:Minimum=0
	// +kubebuilder:default:=30
	InitialDelaySeconds int32 `json:"initialDelaySeconds"`
	//  PeriodSeconds Period seconds for Probe
	//
	// +kubebuilder:validation:Minimum=0
	// +kubebuilder:default:=10
	PeriodSeconds int32 `json:"periodSeconds"`
	// TimeoutSeconds Timeout seconds for Probe
	//
	// +kubebuilder:validation:Minimum=0
	// +kubebuilder:default:=5
	TimeoutSeconds int32 `json:"timeoutSeconds"`
	// FailureThreshold  Failure threshold for Probe
	//
	// +kubebuilder:validation:Minimum=0
	// +kubebuilder:default:=6
	FailureThreshold int32 `json:"failureThreshold"`
	// SuccessThreshold Success threshold for Probe
	//
	// +kubebuilder:validation:Minimum=0
	// +kubebuilder:default:=1
	SuccessThreshold int32 `json:"successThreshold"`
	// ProbeCommandTimeout Probe command timeout for Probe
	//
	// +kubebuilder:default:=2
	ProbeCommandTimeout int32 `json:"probeCommandTimeout,omitempty"`
}

type ServicePolicy struct {
	// Type Kubernetes Service type
	//
	// +kubebuilder:validation:Enum=ClusterIP;NodePort;LoadBalancer;ExternalName
	// +kubebuilder:default:=ClusterIP
	Type string `json:"type,omitempty"`
	// Ports is service traffic ports
	//
	Ports ServicePort `json:"ports,omitempty"`
	// ClusterIP ZooKeeper service Cluster IP
	//
	// +kubebuilder:validation:Optional
	ClusterIP string `json:"clusterIP,omitempty"`
	// LoadBalancerIP ZooKeeper service Load Balancer IP
	//
	// +kubebuilder:validation:Optional
	LoadBalancerIP string `json:"loadBalancerIP,omitempty"`
	// LoadBalancerSourceRanges ZooKeeper service Load Balancer sources
	//
	LoadBalancerSourceRanges []string `json:"loadBalancerSourceRanges,omitempty"`
	// Annotations Additional custom annotations for ZooKeeper service
	//
	// +kubebuilder:validation:Optional
	Annotations map[string]string `json:"annotations,omitempty"`
	// Headless Service
	//
	Headless HeadlessService `json:"headless,omitempty"`
}

type ServicePort struct {
	// ClientPort ZooKeeper client service port
	//
	// +kubebuilder:default:=2181
	ClientPort int32 `json:"clientPort,omitempty"`
	// TLSPort ZooKeeper TLS service port
	//
	// +kubebuilder:default:=3181
	TLSPort int32 `json:"TLSPort,omitempty"`
	// FollowerPort ZooKeeper follower service port
	//
	// +kubebuilder:default:=2888
	FollowerPort int32 `json:"followerPort,omitempty"`
	// ElectionPort ZooKeeper election service port
	//
	// +kubebuilder:default:=3888
	ElectionPort int32 `json:"electionPort,omitempty"`
}

type NodePort struct {
	// Client Node port for clients
	//
	Client string `json:"client,omitempty"`
	// Tls Node port for TLS
	//
	Tls string `json:"tls,omitempty"`
}

type HeadlessService struct {
	// Annotations for the Headless Service
	//
	Annotations map[string]string `json:"annotations,omitempty"`
	// PublishNotReadyAddresses  If the ZooKeeper headless service should publish DNS records for not ready pods
	//
	// +kubebuilder:default:=true
	PublishNotReadyAddresses bool `json:"publishNotReadyAddresses,omitempty"`
}

type ZookeeperPersistence struct {
	// Enabled Enable ZooKeeper data persistence using PVC. If false, use emptyDir
	//
	// +kubebuilder:default:=true
	Enabled bool `json:"enabled,omitempty"`
	// StorageClass PVC Storage Class for ZooKeeper data volume
	//
	// +kubebuilder:validation:Optional
	StorageClass string `json:"storageClass,omitempty"`
	// Annotation Annotations for the PVC
	//
	// +kubebuilder:validation:Optional
	Annotation map[string]string `json:"annotation,omitempty"`
	// AccessModes PVC Access modes
	//
	// +kubebuilder:validation:Enum:=ReadWriteOnce;ReadOnlyMany;ReadWriteMany;ReadWriteOncePod
	// +kubebuilder:default:=ReadWriteOnce
	AccessModes string `json:"accessModes,omitempty"`
	// Data Zookeeper Data persistence
	//
	// +kubebuilder:validation:Optional
	Data ZookeeperDataPvc `json:"data,omitempty"`
	// DataLog Zookeeper Datalog persistence
	//
	// +kubebuilder:validation:Optional
	DataLog ZookeeperDataPvc `json:"dataLog,omitempty"`
}

type ZookeeperDataPvc struct {
	// Size PVC Storage Request for ZooKeeper data volume
	//
	// +kubebuilder:default:="20Gi"
	Size string `json:"size,omitempty"`
	// Selector to match an existing Persistent Volume for ZooKeeper's data PVC
	//
	// +kubebuilder:validation:Optional
	Selector map[string]string `json:"selector,omitempty"`
	// ExistingClaim Name of an existing PVC to use (only when deploying a single replica)
	//
	// +kubebuilder:validation:Optional
	ExistingClaim string `json:"existingClaim,omitempty"`
}
