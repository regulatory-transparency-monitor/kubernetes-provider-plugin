package api

import (
	"encoding/json"
	"fmt"

	"github.com/regulatory-transparency-monitor/kubernetes-provider-plugin/pkg/httpwrapper"
	"github.com/regulatory-transparency-monitor/kubernetes-provider-plugin/pkg/models"
)

type ClusterResources struct {
	BaseURL string
	Client  *httpwrapper.HTTPClient
}

func (c *ClusterResources) GetNodes() (models.NodeList, error) {
	endpoint := "nodes"

	// Construct the GET request.
	req, err := c.Client.NewRequest("GET", endpoint, nil, nil)
	if err != nil {
		return models.NodeList{}, fmt.Errorf("failed to create GetNodes request: %w", err)
	}

	// Execute the request.
	res, err := c.Client.Do(req)
	if err != nil {
		return models.NodeList{}, fmt.Errorf("request GetNodes failed: %w", err)
	}
	defer res.Body.Close()
	/* bodyBytes, _ := ioutil.ReadAll(res.Body)
	fmt.Println("Raw HTTP Response:", string(bodyBytes))
	res.Body = ioutil.NopCloser(bytes.NewBuffer(bodyBytes)) */

	// Decode the response.
	var response models.NodeList
	if err := json.NewDecoder(res.Body).Decode(&response); err != nil {
		return models.NodeList{}, fmt.Errorf("failed to decode GetNodes response: %w", err)
	}

	return response, nil
}
