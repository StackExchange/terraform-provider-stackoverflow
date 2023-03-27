package client

import (
	"fmt"
	"io"
	"net/http"
	"time"
)

const HostURL string = "https://api.stackoverflowteams.com/2.3/"

type Client struct {
	BaseURL     string
	TeamName    string
	AccessToken string
	HTTPClient  *http.Client
}

func NewClient(baseURL *string, teamName *string, accessToken *string) *Client {
	c := Client{
		BaseURL:     HostURL,
		TeamName:    *teamName,
		AccessToken: *accessToken,
		HTTPClient: &http.Client{
			Timeout: time.Minute,
		},
	}

	if baseURL != nil {
		c.BaseURL = *baseURL
	}

	return &c
}

func (c *Client) doRequest(req *http.Request) ([]byte, error) {
	req.Header.Set("X-API-Access-Token", c.AccessToken)

	res, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("status: %d, body: %s", res.StatusCode, body)
	}

	return body, err
}
