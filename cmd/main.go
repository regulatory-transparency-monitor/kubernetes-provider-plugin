package main

import (
	"github.com/regulatory-transparency-monitor/kubernetes-provider-plugin/pkg/api"
	"github.com/regulatory-transparency-monitor/kubernetes-provider-plugin/pkg/services"
)

func main() {
	// Create a new resource service
	clusterResources := &api.ClusterResources{}

	// Create a new Provider
	t := &services.KubernetesPlugin{Cluster: clusterResources}
	t.Initialize()
	t.Scan()

}
