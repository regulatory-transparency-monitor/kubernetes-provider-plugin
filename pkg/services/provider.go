package services

import (
	"context"
	"flag"
	"fmt"

	shared "github.com/regulatory-transparency-monitor/commons/models"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/klog/v2"
)

type KubernetesPlugin struct {
	Clientset *kubernetes.Clientset
	Config    map[string]interface{}
}

func init() {
	klog.InitFlags(nil)
	_ = flag.Set("logtostderr", "true")
	_ = flag.Set("v", "6")
	flag.Parse()
}

func (provider *KubernetesPlugin) Initialize(config map[string]interface{}) error {

	apiAccess, ok := config["api_access"].(map[string]interface{})
	if !ok {
		return fmt.Errorf("kubernetes api_access configuration is missing or invalid")
	}

	kubeURL, ok := apiAccess["base_url"].(string)
	if !ok || kubeURL == "" {
		return fmt.Errorf("kubernetes kubeURL configuration is missing or invalid")
	}
	token, ok := apiAccess["token"].(string)
	if !ok || token == "" {
		return fmt.Errorf("kubernetes token configuration is missing or invalid")
	}
	fmt.Println("Using API URL:", kubeURL)
	fmt.Println("Using Token:", token)


	kubeConfig := &rest.Config{
		Host: kubeURL,
		TLSClientConfig: rest.TLSClientConfig{
			Insecure: true,
		},
		BearerToken: token,
	}

	clientset, err := kubernetes.NewForConfig(kubeConfig)
	if err != nil {
		return err
	}

	provider.Clientset = clientset
	return nil
}
func (provider *KubernetesPlugin) FetchData() (shared.RawData, error) {
	data := make(shared.RawData)

	nodes, err := provider.FetchNodes()
	if err != nil {
		return data, err
	}

	pods, err := provider.FetchPods()
	if err != nil {
		return data, err
	}

	pvs, err := provider.FetchPersistentVolumes()
	if err != nil {
		return data, err
	}
	nodeInterfaces := convertNodesToInterfaces(nodes)
	podInterfaces := convertPodsToInterfaces(pods)
	pvInterfaces := convertPVsToInterfaces(pvs)

	data["k8s_pv"] = pvInterfaces
	data["k8s_node"] = nodeInterfaces
	data["k8s_pod"] = podInterfaces

	return data, nil
}

func convertNodesToInterfaces(nodes []corev1.Node) []interface{} {
	nodeInterfaces := make([]interface{}, len(nodes))
	for i, node := range nodes {
		nodeInterfaces[i] = node
	}
	return nodeInterfaces
}

func convertPodsToInterfaces(pods []corev1.Pod) []interface{} {
	podInterfaces := make([]interface{}, len(pods))
	for i, pod := range pods {
		podInterfaces[i] = pod
	}
	return podInterfaces
}

func convertPVsToInterfaces(pvs []corev1.PersistentVolume) []interface{} {
	pvInterfaces := make([]interface{}, len(pvs))
	for i, pv := range pvs {
		pvInterfaces[i] = pv
	}
	return pvInterfaces
}
func (provider *KubernetesPlugin) FetchNodes() ([]corev1.Node, error) {
	nodeList, err := provider.Clientset.CoreV1().Nodes().List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		return nil, err
	}
	return nodeList.Items, nil
}

func (provider *KubernetesPlugin) FetchPods() ([]corev1.Pod, error) {
	podList, err := provider.Clientset.CoreV1().Pods("sock-shop").List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		return nil, err
	}
	return podList.Items, nil
}

func (provider *KubernetesPlugin) FetchPersistentVolumes() ([]corev1.PersistentVolume, error) {
	pvList, err := provider.Clientset.CoreV1().PersistentVolumes().List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		return nil, err
	}
	return pvList.Items, nil
}
