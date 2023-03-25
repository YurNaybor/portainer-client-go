package portainer

// Note: portainer api does not provide types for networks and endpoints are passed trough to backend (e.g. Docker, K8s etc.)
// Note2: Secrets re read only (at least on Docker)

import (
	"encoding/json"
	"fmt"
	"net/http"
	//"strings"
)

type DockerNetwork struct {
	ID	string `json:"id"`
	Name	string `json:"name"`
}

// GetDockerNetworks - Returns list of dockerNetworks
func (c *Client) GetDockerNetworks(endpointID int) ([]DockerNetwork, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/endpoints/%s/docker/networks", c.ApiUrl, endpointID), nil)
	if err != nil {
		return nil, err
	}

	body, err := c.doRequest(req)
	if err != nil {
		return nil, err
	}

	dockerNetworks := []DockerNetwork{}
	err = json.Unmarshal(body, &dockerNetworks)
	if err != nil {
		return nil, err
	}

	return dockerNetworks, nil
}

// GetDockerNetwork - Returns specific dockerNetwork
func (c *Client) GetDockerNetwork(endpointID int, dockerNetworkID int) (DockerNetwork, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/endpoints/%s/docker/networks/%s", c.ApiUrl, endpointID, dockerNetworkID), nil)
	if err != nil {
		return DockerNetwork{}, err
	}

	body, err := c.doRequest(req)
	if err != nil {
		return DockerNetwork{}, err
	}

	dockerNetwork := DockerNetwork{}
	err = json.Unmarshal(body, &dockerNetwork)
	if err != nil {
		return DockerNetwork{}, err
	}

	return dockerNetwork, nil
}
