package portainer

import (
	"net/http"
	"time"
	"io/ioutil"
	"fmt"
	"crypto/tls"
)

type Client struct {
	ApiUrl string
	HTTPClient *http.Client
	ApiToken string
}

func NewClient(apiUrl, apiToken string) (*Client, error) {
	s := &tls.Config{InsecureSkipVerify: true}

	t := &http.Transport{TLSClientConfig: s}

	c := Client{
		HTTPClient: &http.Client{Timeout: 10 * time.Second, Transport: t},
	}

	c.ApiUrl = apiUrl
	c.ApiToken = apiToken

	return &c, nil
}

func (c *Client) doRequest(req *http.Request) ([]byte, error) {
	req.Header.Set("X-Api-Key", c.ApiToken)

	res, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("status: %d, body: %s", res.StatusCode, body)
	}

	return body, err
}
