package api

import (
	"encoding/json"
	"fmt"

	"github.com/regulatory-transparency-monitor/kubernetes-provider-plugin/pkg/httpwrapper"
	"github.com/regulatory-transparency-monitor/kubernetes-provider-plugin/pkg/models"
)

type WorkloadResources struct {
	BaseURL string
	Client  *httpwrapper.HTTPClient
}

func (w *WorkloadResources) GetPods() (models.PodList, error) {
	endpoint := fmt.Sprintf("%s/pods", w.BaseURL)

	// Construct the GET request.
	req, err := w.Client.NewRequest("GET", endpoint, nil, nil)
	if err != nil {
		return models.PodList{}, fmt.Errorf("failed to create GetPods request: %w", err)
	}

	// Execute the request.
	res, err := w.Client.Do(req)
	if err != nil {
		return models.PodList{}, fmt.Errorf("request GetPods failed: %w", err)
	}
	/* defer res.Body.Close()
	bodyBytes, _ := ioutil.ReadAll(res.Body)
	fmt.Println("Raw HTTP Response:", string(bodyBytes))
	res.Body = ioutil.NopCloser(bytes.NewBuffer(bodyBytes)) */

	// Decode the response.
	var response models.PodList
	if err := json.NewDecoder(res.Body).Decode(&response); err != nil {
		return models.PodList{}, fmt.Errorf("failed to decode GetPods response: %w", err)
	}

	return response, nil
}
