package services

import (
	"encoding/json"
	"fmt"

	"github.com/regulatory-transparency-monitor/commons/models"
	"github.com/regulatory-transparency-monitor/kubernetes-provider-plugin/pkg/api"
	"github.com/regulatory-transparency-monitor/kubernetes-provider-plugin/pkg/httpwrapper"
	"github.com/regulatory-transparency-monitor/kubernetes-provider-plugin/pkg/interfaces"
)

// KubernetesPlugin is a struct holding Keystone and Nova service
type KubernetesPlugin struct {
	Cluster  interfaces.ClusterAPI
	Workload interfaces.WorkloadAPI
}

// TODO pass base url and credentials as parameters from the graph builder service
func (provider *KubernetesPlugin) Initialize() error {
	httpClient := httpwrapper.NewClient("http://localhost:8001/api/v1/")

	provider.Cluster = &api.ClusterResources{
		Client: httpClient,
		// provide url and credentials here
	}
	provider.Workload = &api.WorkloadResources{
		BaseURL: "namespaces/sock-shop/",
		Client:  httpClient,
		// provide url and credentials here
	}

	return nil

}

// Scan returns fetched data from OpenStack API
func (provider *KubernetesPlugin) Scan() (models.CombinedResources, error) {
	var resources models.CombinedResources
	resources.Source = "kubernetes"

	clusterResources, err := provider.FetchClusterData()
	if err != nil {
		return resources, err

	}
	resources.Data = append(resources.Data, *clusterResources)

	workloadResources, err := provider.FetchWorkloadData()
	if err != nil {
		return resources, err
	}
	resources.Data = append(resources.Data, *workloadResources)

	return resources, nil
}

func (provider *KubernetesPlugin) FetchClusterData() (*models.ServiceData, error) {

	nodeListModel, err := provider.Cluster.GetNodes()
	if err != nil {
		return nil, err
	}

	return &models.ServiceData{
		ServiceSource: "cluster",
		Data:          []interface{}{nodeListModel},
	}, nil
}

func (provider *KubernetesPlugin) FetchWorkloadData() (*models.ServiceData, error) {

	podListModel, err := provider.Workload.GetPods()
	if err != nil {
		return nil, err
	}

	return &models.ServiceData{
		ServiceSource: "workload",
		Data:          []interface{}{podListModel},
	}, nil

}

func PrintAsJSON(v interface{}) {
	data, err := json.MarshalIndent(v, "", "  ")
	if err != nil {
		fmt.Println("Error marshalling to JSON:", err)
	}

	fmt.Println(string(data))
}
