package models

type PodList struct {
	Kind       string   `json:"kind"`
	APIVersion string   `json:"apiVersion"`
//	Metadata   Metadata `json:"metadata"`
	Items      []Pod    `json:"items"`
}

type Pod struct {
	Metadata PodMetadata            `json:"metadata"`
	Spec     map[string]interface{} `json:"spec"`
	Status   map[string]interface{} `json:"status"`
}

type PodMetadata struct {
	Name            string `json:"name"`
	GenName         string `json:"generateName"`
	Namespace       string `json:"namespace"`
	UID             string `json:"uid"`
	ResourceVersion string `json:"resourceVersion"`
	CreationTime    string `json:"creationTimestamp"`
	Labels          map[string]string      `json:"labels"`
	Annotations     map[string]string      `json:"annotations"`
	//OwnerReferences map[string]interface{} `json:"ownerReferences"`
	//ManagedFields   map[string]interface{} `json:"managedFields"`
}
