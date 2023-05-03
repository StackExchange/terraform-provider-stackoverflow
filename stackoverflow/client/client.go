package client

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"
)

const HostURL string = "https://api.stackoverflowteams.com/2.3/"

type Client struct {
	BaseURL     string
	TeamName    string
	AccessToken string
	HTTPClient  *http.Client
	DefaultTags []string
}

func NewClient(baseURL *string, teamName *string, accessToken *string) *Client {
	c := Client{
		BaseURL:     HostURL,
		TeamName:    *teamName,
		AccessToken: *accessToken,
		HTTPClient: &http.Client{
			Timeout: time.Minute,
		},
		DefaultTags: []string{},
	}

	if baseURL != nil {
		c.BaseURL = *baseURL
	}

	return &c
}

func (c *Client) doPublicRequest(req *http.Request) ([]byte, error) {
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

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

func (c *Client) get(route string, identifiers *[]int, filter *string) (*[]byte, error) {
	ids := make([]string, len(*identifiers))
	for i, id := range *identifiers {
		ids[i] = strconv.Itoa(id)
	}
	url := fmt.Sprintf("%s%s/%s?team=%s&filter=%s", c.BaseURL, route, strings.Join(ids, ";"), c.TeamName, *filter)
	log.Print(url)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	body, err := c.doRequest(req)
	if err != nil {
		return nil, err
	}

	buf := new(strings.Builder)
	io.Copy(buf, strings.NewReader((string(body))))
	log.Printf("Response body: %s", buf.String())

	return &body, nil
}

func (c *Client) create(route string, formData *string) (*[]byte, error) {
	url := fmt.Sprintf("%s%s?team=%s", c.BaseURL, route, c.TeamName)
	log.Print(url)
	req, err := http.NewRequest("POST", url, strings.NewReader(*formData))
	if err != nil {
		return nil, err
	}

	body, err := c.doRequest(req)
	if err != nil {
		return nil, err
	}

	buf := new(strings.Builder)
	io.Copy(buf, strings.NewReader((string(body))))
	log.Printf("Response body: %s", buf.String())

	return &body, nil
}

func (c *Client) update(route string, formData *string) (*[]byte, error) {
	url := fmt.Sprintf("%s%s?team=%s", c.BaseURL, route, c.TeamName)
	log.Print(url)
	req, err := http.NewRequest("POST", url, strings.NewReader(*formData))
	if err != nil {
		return nil, err
	}

	body, err := c.doRequest(req)
	if err != nil {
		return nil, err
	}

	buf := new(strings.Builder)
	io.Copy(buf, strings.NewReader((string(body))))
	log.Printf("Response body: %s", buf.String())

	return &body, nil
}

func (c *Client) delete(route string) error {
	url := fmt.Sprintf("%s%s?team=%s", c.BaseURL, route, c.TeamName)
	log.Print(url)
	req, err := http.NewRequest("POST", url, nil)
	if err != nil {
		return err
	}

	body, err2 := c.doRequest(req)
	if err2 != nil {
		return err2
	}

	buf := new(strings.Builder)
	io.Copy(buf, strings.NewReader((string(body))))
	log.Printf("Response body: %s", buf.String())

	return nil
}

func UnwrapResponseItems[T Answer | Article | Filter | Question](response *[]byte) (*[]T, error) {
	responseWrapper := Wrapper[T]{}
	err := json.Unmarshal(*response, &responseWrapper)
	if err != nil {
		return nil, err
	}

	items := responseWrapper.Items

	return &items, nil
}
