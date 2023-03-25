package portainer

import (
	"encoding/json"
	"fmt"
	"net/http"
	//"strings"
	portainer "github.com/portainer/portainer/api"
)

// GetStacks - Returns list of stacks
func (c *Client) GetStacks() ([]portainer.Stack, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/stacks", c.ApiUrl), nil)
	if err != nil {
		return nil, err
	}

	body, err := c.doRequest(req)
	if err != nil {
		return nil, err
	}

	stacks := []portainer.Stack{}
	err = json.Unmarshal(body, &stacks)
	if err != nil {
		return nil, err
	}

	return stacks, nil
}

// GetStack - Returns specific stack
func (c *Client) GetStack(stackID string) (portainer.Stack, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/stacks/%s", c.ApiUrl, stackID), nil)
	if err != nil {
		return portainer.Stack{}, err
	}

	body, err := c.doRequest(req)
	if err != nil {
		return portainer.Stack{}, err
	}

	stack := portainer.Stack{}
	err = json.Unmarshal(body, &stack)
	if err != nil {
		return portainer.Stack{}, err
	}

	return stack, nil
}
