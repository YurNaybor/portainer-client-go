package portainer

import (
	"encoding/json"
	"fmt"
	"net/http"
	//"strings"
	portainer "github.com/portainer/portainer/api"
)

// GetEndpoints - Returns list of endpoints
func (c *Client) GetEndpoints() ([]portainer.Endpoint, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/endpoints", c.ApiUrl), nil)
	if err != nil {
		return nil, err
	}

	body, err := c.doRequest(req)
	if err != nil {
		return nil, err
	}

	endpoints := []portainer.Endpoint{}
	err = json.Unmarshal(body, &endpoints)
	if err != nil {
		return nil, err
	}

	return endpoints, nil
}

// GetEndpoint - Returns specific endpoint
func (c *Client) GetEndpoint(endpointID int) (portainer.Endpoint, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/endpoints/%d", c.ApiUrl, endpointID), nil)
	if err != nil {
		return portainer.Endpoint{}, err
	}

	body, err := c.doRequest(req)
	if err != nil {
		return portainer.Endpoint{}, err
	}

	endpoint := portainer.Endpoint{}
	err = json.Unmarshal(body, &endpoint)
	if err != nil {
		return portainer.Endpoint{}, err
	}

	return endpoint, nil
}
