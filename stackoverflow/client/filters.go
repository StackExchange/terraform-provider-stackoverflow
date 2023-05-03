package client

import (
	"fmt"
	"log"
	"net/http"
	"net/url"
	"strings"
)

func (c *Client) GetFilters(filterIDs *[]string) (*[]Filter, error) {
	url := fmt.Sprintf("%s%s/%s?key=", c.BaseURL, "filters", strings.Join(*filterIDs, ";"))
	log.Print(url)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	body, err := c.doPublicRequest(req)
	if err != nil {
		return nil, err
	}

	filters, err := UnwrapResponseItems[Filter](&body)
	if err != nil {
		return nil, err
	}

	return filters, nil
}

func (c *Client) CreateFilter(filter *Filter) (*Filter, error) {
	url := fmt.Sprintf("%s%s", c.BaseURL, "filters/create?key=")
	log.Print(url)
	formData := GenerateFilterFormData(filter)
	req, err := http.NewRequest("POST", url, strings.NewReader(*formData))
	if err != nil {
		return nil, err
	}

	body, err := c.doPublicRequest(req)
	if err != nil {
		return nil, err
	}

	filters, err := UnwrapResponseItems[Filter](&body)
	if err != nil {
		return nil, err
	}

	if len(*filters) != 1 {
		return nil, fmt.Errorf("response wrapper does not contain expected number of items (1); item length is %d", len(*filters))
	}

	newFilter := (*filters)[0]

	return &newFilter, nil
}

func (c *Client) UpdateFilter(filter *Filter) (*Filter, error) {
	url := fmt.Sprintf("%s%s%s%s", c.BaseURL, "filters/", filter.ID, "/edit")
	log.Print(url)
	formData := GenerateFilterFormData(filter)
	req, err := http.NewRequest("POST", url, strings.NewReader(*formData))
	if err != nil {
		return nil, err
	}

	body, err := c.doPublicRequest(req)
	if err != nil {
		return nil, err
	}

	filters, err := UnwrapResponseItems[Filter](&body)
	if err != nil {
		return nil, err
	}

	if len(*filters) != 1 {
		return nil, fmt.Errorf("response wrapper does not contain expected number of items (1); item length is %d", len(*filters))
	}

	newFilter := (*filters)[0]

	return &newFilter, nil
}

func (c *Client) DeleteFilter(filterId string) error {
	url := fmt.Sprintf("%s%s%s%s", c.BaseURL, "filters/", filterId, "/delete")
	log.Print(url)
	req, err := http.NewRequest("POST", url, nil)
	if err != nil {
		return err
	}

	_, err2 := c.doPublicRequest(req)
	if err2 != nil {
		return err
	}

	return nil
}

func GenerateFilterFormData(filter *Filter) *string {
	formData := fmt.Sprintf("include=%s&exclude=%s&unsafe=%t", url.QueryEscape(strings.Join(filter.Include, ";")), url.QueryEscape(strings.Join(filter.Exclude, ";")), filter.Unsafe)
	log.Printf("Form data: %s", formData)
	return &formData
}
