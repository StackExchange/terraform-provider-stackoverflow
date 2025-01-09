package client

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
	"time"
)

type Client struct {
	BaseURL     string
	TeamName    string
	AccessToken string
	HTTPClient  *http.Client
	DefaultTags []string
}

func NewClient(baseURL *string, accessToken *string) *Client {
	c := Client{
		BaseURL:     *baseURL,
		AccessToken: *accessToken,
		HTTPClient: &http.Client{
			Timeout: time.Minute,
		},
		DefaultTags: []string{},
	}

	return &c
}

func (c *Client) doRequest(req *http.Request, successCode int) ([]byte, error) {
	authzHeader := fmt.Sprintf("Bearer %s", c.AccessToken)
	req.Header.Set("Authorization", authzHeader)
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")

	res, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if res.StatusCode != successCode {
		return nil, fmt.Errorf("status: %d, body: %s", res.StatusCode, body)
	}

	return body, err
}

func (c *Client) get(route string) (*[]byte, error) {
	url := fmt.Sprintf("%s%s", c.BaseURL, route)
	log.Print(url)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	response, err := c.doRequest(req, http.StatusOK)
	if err != nil {
		return nil, err
	}

	buf := new(strings.Builder)
	io.Copy(buf, strings.NewReader((string(response))))
	log.Printf("Response body: %s", buf.String())

	return &response, nil
}

func (c *Client) create(route string, body *string) (*[]byte, error) {
	url := fmt.Sprintf("%s%s", c.BaseURL, route)
	log.Print(url)
	req, err := http.NewRequest("POST", url, strings.NewReader(*body))
	if err != nil {
		return nil, err
	}

	response, err := c.doRequest(req, http.StatusCreated)
	if err != nil {
		return nil, err
	}

	buf := new(strings.Builder)
	io.Copy(buf, strings.NewReader((string(response))))
	log.Printf("Response body: %s", buf.String())

	return &response, nil
}

func (c *Client) update(route string, body *string) (*[]byte, error) {
	url := fmt.Sprintf("%s%s", c.BaseURL, route)
	log.Print(url)
	req, err := http.NewRequest("PUT", url, strings.NewReader(*body))
	if err != nil {
		return nil, err
	}

	response, err := c.doRequest(req, http.StatusOK)
	if err != nil {
		return nil, err
	}

	buf := new(strings.Builder)
	io.Copy(buf, strings.NewReader((string(response))))
	log.Printf("Response body: %s", buf.String())

	return &response, nil
}

func (c *Client) delete(route string) error {
	url := fmt.Sprintf("%s%s", c.BaseURL, route)
	log.Print(url)
	req, err := http.NewRequest("DELETE", url, nil)
	if err != nil {
		return err
	}

	response, err2 := c.doRequest(req, 204)
	if err2 != nil {
		return err2
	}

	buf := new(strings.Builder)
	io.Copy(buf, strings.NewReader((string(response))))
	log.Printf("Response body: %s", buf.String())

	return nil
}

func ConvertFromJsonString[T Answer[Tag] | Article[Tag] | Collection | Tag | Question[Tag]](response *[]byte) (*T, error) {
	var object T
	err := json.Unmarshal(*response, &object)
	if err != nil {
		return nil, err
	}

	return &object, nil
}

func ConvertToJsonString[T Answer[string] | Article[string] | Collection | Tag | Question[string]](object T) (*string, error) {
	bits, err := json.Marshal(object)

	if err != nil {
		return nil, err
	}
	body := string(bits)

	log.Printf("JSON body: %s", body)
	return &body, nil
}
