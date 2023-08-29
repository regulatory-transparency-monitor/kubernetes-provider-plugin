package models

import "time"

type NodeList struct {
	Kind       string   `json:"kind"`
	APIVersion string   `json:"apiVersion"`
	Metadata   Metadata `json:"metadata"`
	Items      []Node   `json:"items"`
}

type Metadata struct {
	ResourceVersion string `json:"resourceVersion"`
}

/*
	 type Node struct {
		Metadata NodeMetadata `json:"metadata"`
		Spec     NodeSpec     `json:"spec"`
		Status   NodeStatus   `json:"status"`
	}
*/
type Node struct {
	Metadata NodeMetadata           `json:"metadata"`
	Spec     map[string]interface{} `json:"spec"`
	Status   map[string]interface{} `json:"status"`
}

type NodeMetadata struct {
	Name              string      `json:"name"`
	UID               string      `json:"uid"`
	ResourceVersion   string      `json:"resourceVersion"`
	CreationTimestamp string      `json:"creationTimestamp"`
	Labels            interface{} `json:"labels"`
	Annotations       interface{} `json:"annotations"`
	ManagedFields     interface{} `json:"managedFields"`
}

type ManagedField struct {
	Manager     string    `json:"manager"`
	Operation   string    `json:"operation"`
	APIVersion  string    `json:"apiVersion"`
	Time        time.Time `json:"time"`
	FieldsType  string    `json:"fieldsType"`
	FieldsV1    FieldsV1  `json:"fieldsV1"`
	Subresource string    `json:"subresource"`
}

type FieldsV1 struct {
	FMetadata FMetadata `json:"f:metadata"`
}

type FMetadata struct {
	FAnnotations FAnnotations `json:"f:annotations"`
	FLabels      FLabels      `json:"f:labels"`
}

type FAnnotations struct {
	FAlphaKubernetesIOProvidedNodeIP                  string `json:"f:alpha.kubernetes.io/provided-node-ip"`
	FVolumesKubernetesIOControllerManagedAttachDetach string `json:"f:volumes.kubernetes.io/controller-managed-attach-detach"`
}

type FLabels struct {
	FBetaKubernetesIOArch string `json:"f:beta.kubernetes.io/arch"`
	FBetaKubernetesIOOS   string `json:"f:beta.kubernetes.io/os"`
	FKubernetesIOArch     string `json:"f:kubernetes.io/arch"`
	FKubernetesIOHostname string `json:"f:kubernetes.io/hostname"`
	FKubernetesIOOS       string `json:"f:kubernetes.io/os"`
}

type NodeSpec struct {
	PodCIDR    string   `json:"podCIDR"`
	PodCIDRs   []string `json:"podCIDRs"`
	ProviderID string   `json:"providerID"`
	Taints     []Taint  `json:"taints"`
}

type Taint struct {
	Key    string `json:"key"`
	Effect string `json:"effect"`
}

type NodeStatus struct {
	Capacity        Capacity        `json:"capacity"`
	Allocatable     Capacity        `json:"allocatable"`
	Conditions      []NodeCondition `json:"conditions"`
	Addresses       []NodeAddress   `json:"addresses"`
	DaemonEndpoints DaemonEndpoints `json:"daemonEndpoints"`
	NodeInfo        NodeInfo        `json:"nodeInfo"`
	Images          []NodeImage     `json:"images"`
}

type Capacity struct {
	CPU              string `json:"cpu"`
	EphemeralStorage string `json:"ephemeral-storage"`
	Hugepages1Gi     string `json:"hugepages-1Gi"`
	Hugepages2Mi     string `json:"hugepages-2Mi"`
	Memory           string `json:"memory"`
	Pods             string `json:"pods"`
}

type NodeCondition struct {
	Type               string    `json:"type"`
	Status             string    `json:"status"`
	LastHeartbeatTime  time.Time `json:"lastHeartbeatTime"`
	LastTransitionTime time.Time `json:"lastTransitionTime"`
	Reason             string    `json:"reason"`
	Message            string    `json:"message"`
}

type NodeAddress struct {
	Type    string `json:"type"`
	Address string `json:"address"`
}

type DaemonEndpoints struct {
	KubeletEndpoint KubeletEndpoint `json:"kubeletEndpoint"`
}

type KubeletEndpoint struct {
	Port int `json:"Port"`
}

type NodeInfo struct {
	MachineID               string `json:"machineID"`
	SystemUUID              string `json:"systemUUID"`
	BootID                  string `json:"bootID"`
	KernelVersion           string `json:"kernelVersion"`
	OSImage                 string `json:"osImage"`
	ContainerRuntimeVersion string `json:"containerRuntimeVersion"`
	KubeletVersion          string `json:"kubeletVersion"`
	KubeProxyVersion        string `json:"kubeProxyVersion"`
	OperatingSystem         string `json:"operatingSystem"`
	Architecture            string `json:"architecture"`
}

type NodeImage struct {
	Names     []string `json:"names"`
	SizeBytes int64    `json:"sizeBytes"`
}
