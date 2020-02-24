package v1

// ComponentStatus component status
type ComponentStatus string

const (
	//ComponentStatusRunning running
	ComponentStatusRunning = "Running"
	// ComponentStatusIniting initing
	ComponentStatusIniting = "Initing"
	//ComponentStatusCreating creating
	ComponentStatusCreating = "Creating"
	// ComponentStatusTerminating terminal
	ComponentStatusTerminating = "Terminating" // TODO fanyangyang have not found this case
	// ComponentStatusFailed failed
	ComponentStatusFailed = "Failed"
)

// RbdComponentStatus rainbond component status
type RbdComponentStatus struct {
	Name string `json:"name"`

	// Total number of non-terminated pods targeted by this deployment (their labels match the selector).
	// +optional
	Replicas int32 `json:"replicas"`

	// Total number of ready pods targeted by this deployment.
	// +optional
	ReadyReplicas int32 `json:"readyReplicas"`

	Status          ComponentStatus `json:"status"` //translate pod status to component status
	Message         string          `json:"message"`
	Reason          string          `json:"reason"`
	ISInitComponent bool            `json:"isInitComponent"`

	PodStatuses []PodStatus `json:"podStatus"`
}

// PodStatus represents information about the status of a pod, which belongs to RbdComponent.
type PodStatus struct {
	Name              string               `json:"name"`
	Phase             string               `json:"phase"`
	HostIP            string               `json:"hostIP"`
	Reason            string               `json:"reason"`
	Message           string               `json:"message"`
	ContainerStatuses []PodContainerStatus `json:"container_statuses"`
}

// PodContainerStatus -
type PodContainerStatus struct {
	ContainerID string `json:"containerID"`
	Image       string `json:"image"`
	// Specifies whether the container has passed its readiness probe.
	Ready   bool   `json:"ready"`
	State   string `json:"state"`
	Reason  string `json:"reason"`
	Message string `json:"message"`
}

// K8sNode holds the information about a kubernetes node.
type K8sNode struct {
	Name       string `json:"name,omitempty"`
	InternalIP string `json:"internalIP,omitempty"`
	ExternalIP string `json:"externalIP,omitempty"`
}

// AvailableNodes contains nodes available for special rainbond components to run,
// such as rbd-gateway, rbd-chaos.
type AvailableNodes struct {
	// The nodes with user-specified labels.
	SpecifiedNodes []*K8sNode `json:"specifiedNodes,omitempty"`
	// A list of kubernetes master nodes.
	MasterNodes []*K8sNode `json:"masterNodes,omitempty"`
}

// StorageClasses is a List of StorageCass available in the cluster.
// StorageClass storage class
type StorageClass struct {
	Name        string `json:"name"`
	Provisioner string `json:"provisioner"`
	AccessMode  string `json:"accessMode"`
}

// ClusterStatusInfo holds the information of rainbondcluster status.
type ClusterStatusInfo struct {
	// holds some recommend nodes available for rbd-gateway to run.
	GatewayAvailableNodes *AvailableNodes `json:"gatewayAvailableNodes"`
	// holds some recommend nodes available for rbd-chaos to run.
	ChaosAvailableNodes *AvailableNodes `json:"chaosAvailableNodes"`
	StorageClasses      []*StorageClass `json:"storageClasses"`
}
