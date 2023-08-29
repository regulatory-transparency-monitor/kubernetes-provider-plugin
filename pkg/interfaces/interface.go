package interfaces

import "github.com/regulatory-transparency-monitor/kubernetes-provider-plugin/pkg/models"

type ClusterAPI interface {
	GetNodes() (models.NodeList, error)
	//GetPods() (models.PodList, error)
	/* GetServices() ([]Service, error)
	GetDeployments() ([]Deployment, error)
	GetReplicaSets() ([]ReplicaSet, error)
	GetStatefulSets() ([]StatefulSet, error)
	GetDaemonSets() ([]DaemonSet, error)
	GetJobs() ([]Job, error)
	GetCronJobs() ([]CronJob, error)
	GetPersistentVolumes() ([]PersistentVolume, error)
	GetPersistentVolumeClaims() ([]PersistentVolumeClaim, error)
	GetStorageClasses() ([]StorageClass, error)
	GetConfigMaps() ([]ConfigMap, error)
	GetSecrets() ([]Secret, error)	 */
}

type WorkloadAPI interface {
	GetPods() (models.PodList, error)
}
