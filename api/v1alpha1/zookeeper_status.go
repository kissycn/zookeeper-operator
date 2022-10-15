package v1alpha1

import v1 "k8s.io/api/core/v1"

// ZookeeperStatus defines the observed state of Zookeeper
type ZookeeperStatus struct {
	// Replicas is the number of  desired replicas in the cluster
	// +kubebuilder:default:=0
	Replicas int32 `json:"replicas,omitempty"`
	// ReadyReplicas is the number of  ready replicas in the cluster
	// +kubebuilder:default:=0
	ReadyReplicas int32 `json:"readyReplicas,omitempty"`
	// Members is the zookeeper members in the cluster
	Members MemberStatus `json:"members,omitempty"`
	// CurrentVersion is zookeeper cluster current version
	CurrentVersion string `json:"currentVersion,omitempty"`
	// TargetVersion is zookeeper cluster upgrading version
	TargetVersion string `json:"targetVersion,omitempty"`
	// ClusterStatus zookeeper cluster status
	ClusterStatus ConditionType `json:"clusterStatus,omitempty"`
}

// MemberStatus is the status of the members
type MemberStatus struct {
	// Ready pod name
	Ready []string `json:"ready,omitempty"`
	// Unready pod name
	Unready []string `json:"unready,omitempty"`
}

type ConditionType string

const (
	ClusterRunning   ConditionType = "Running"
	ClusterReady                   = "Ready"
	ClusterUpgrading               = "Upgrading"
	ClusterError                   = "Error"
)

type ClusterCondition struct {
	// Type of cluster condition
	Type ConditionType `json:"type,omitempty"`
	// Status of the condition, one of True, False, Unknown.
	Status v1.ConditionStatus `json:"status,omitempty"`
	// Reason The reason for the condition's last transition.
	Reason string `json:"reason,omitempty"`
	// Message indicating details about the transition.
	Message string `json:"message,omitempty"`
}
