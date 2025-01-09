package client

import (
	"fmt"
)

func (c *Client) GetTag(tagID *int) (*Tag, error) {
	response, err := c.get(fmt.Sprintf("%s/%d", "tags", *tagID))
	if err != nil {
		return nil, err
	}

	tag, err := ConvertFromJsonString[Tag](response)
	if err != nil {
		return nil, err
	}

	return tag, nil
}
