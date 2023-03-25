package portainer

// Note: portainer api does not provide types for secrets and endpoints are passed trough to backend (e.g. Docker, K8s etc.)
// Note2: Secrets re read only (at least on Docker)

import (
	"encoding/json"
	"fmt"
	"net/http"
	//"strings"
)

type DockerSecret struct {
	ID	string `json:"id"`
	Name	string `json:"spec.name"`
	Spec	DockerSecretSpec `json: "spec"`
}

// It also contains labels, but those are nasty to parse
type DockerSecretSpec struct {
	Name	string `json:"name"`
}

// GetDockerSecrets - Returns list of dockerSecrets
func (c *Client) GetDockerSecrets(endpointID int) ([]DockerSecret, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/endpoints/%s/docker/secrets", c.ApiUrl, endpointID), nil)
	if err != nil {
		return nil, err
	}

	body, err := c.doRequest(req)
	if err != nil {
		return nil, err
	}

	dockerSecrets := []DockerSecret{}
	err = json.Unmarshal(body, &dockerSecrets)
	if err != nil {
		return nil, err
	}

	return dockerSecrets, nil
}

// GetDockerSecret - Returns specific dockerSecret
func (c *Client) GetDockerSecret(endpointID int, dockerSecretID string) (DockerSecret, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/endpoints/%s/docker/secrets/%s", c.ApiUrl, endpointID, dockerSecretID), nil)
	if err != nil {
		return DockerSecret{}, err
	}

	body, err := c.doRequest(req)
	if err != nil {
		return DockerSecret{}, err
	}

	dockerSecret := DockerSecret{}
	err = json.Unmarshal(body, &dockerSecret)
	if err != nil {
		return DockerSecret{}, err
	}

	return dockerSecret, nil
}
