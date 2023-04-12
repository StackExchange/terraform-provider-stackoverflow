package client

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"strings"
)

func (c *Client) GetFilters(filterIDs *[]string) (*[]Filter, error) {
	log.Printf("%s/%s", "filters", strings.Join(*filterIDs, ";"))
	route := fmt.Sprintf("%s/%s?key=", "filters", strings.Join(*filterIDs, ";"))
	req, err := http.NewRequest("GET", fmt.Sprintf("%s%s", c.BaseURL, route), nil)
	if err != nil {
		return nil, err
	}

	body, err := c.doPublicRequest(req)
	if err != nil {
		return nil, err
	}

	responseWrapper := Wrapper[Filter]{}
	err = json.Unmarshal(body, &responseWrapper)
	if err != nil {
		return nil, err
	}

	filters := responseWrapper.Items

	return &filters, nil
}

func (c *Client) CreateFilter(filter *Filter) (*Filter, error) {
	formData := GenerateFilterFormData(filter)
	req, err := http.NewRequest("POST", fmt.Sprintf("%s%s", c.BaseURL, "filters/create?key="), strings.NewReader(formData))
	if err != nil {
		return nil, err
	}

	body, err := c.doPublicRequest(req)
	if err != nil {
		return nil, err
	}

	buf := new(strings.Builder)
	io.Copy(buf, strings.NewReader((string(body))))
	log.Printf("Response body: %s", buf.String())

	responseWrapper := Wrapper[Filter]{}
	err = json.Unmarshal(body, &responseWrapper)
	if err != nil {
		return nil, err
	}

	if len(responseWrapper.Items) != 1 {
		return nil, fmt.Errorf("response wrapper does not contain expected number of items (1); item length is %d", len(responseWrapper.Items))
	}

	newFilter := responseWrapper.Items[0]

	return &newFilter, nil
}

func (c *Client) UpdateFilter(filter *Filter) (*Filter, error) {
	formData := GenerateFilterFormData(filter)
	req, err := http.NewRequest("POST", fmt.Sprintf("%s%s%s%s", c.BaseURL, "filters/", filter.ID, "/edit"), strings.NewReader(formData))
	if err != nil {
		return nil, err
	}

	body, err := c.doPublicRequest(req)
	if err != nil {
		return nil, err
	}

	responseWrapper := Wrapper[Filter]{}
	err = json.Unmarshal(body, &responseWrapper)
	if err != nil {
		return nil, err
	}

	if len(responseWrapper.Items) != 1 {
		return nil, fmt.Errorf("response wrapper does not contain expected number of items (1); item length is %d", len(responseWrapper.Items))
	}

	newFilter := responseWrapper.Items[0]

	return &newFilter, nil
}

func (c *Client) DeleteFilter(filterId string) error {
	req, err := http.NewRequest("POST", fmt.Sprintf("%s%s%s%s", c.BaseURL, "filters/", filterId, "/delete"), nil)
	if err != nil {
		return err
	}

	_, err2 := c.doPublicRequest(req)
	if err2 != nil {
		return err
	}

	return nil
}

func GenerateFilterFormData(filter *Filter) string {
	formData := fmt.Sprintf("include=%s&exclude=%s&unsafe=%t", url.QueryEscape(strings.Join(filter.Include, ";")), url.QueryEscape(strings.Join(filter.Exclude, ";")), filter.Unsafe)
	log.Printf("Form data: %s", formData)
	return formData
}
