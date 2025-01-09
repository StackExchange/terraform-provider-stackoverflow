package client

import (
	"fmt"
)

func (c *Client) GetCollection(collectionID *int) (*Collection, error) {
	response, err := c.get(fmt.Sprintf("%s/%d", "collections", *collectionID))
	if err != nil {
		return nil, err
	}

	collection, err := ConvertFromJsonString[Collection](response)
	if err != nil {
		return nil, err
	}

	return collection, nil
}

func (c *Client) CreateCollection(collection *Collection) (*Collection, error) {
	body, err := ConvertToJsonString(*collection)
	if err != nil {
		return nil, err
	}

	response, err := c.create("collections", body)
	if err != nil {
		return nil, err
	}

	newCollection, err := ConvertFromJsonString[Collection](response)
	if err != nil {
		return nil, err
	}

	return newCollection, nil
}

func (c *Client) UpdateCollection(collection *Collection) (*Collection, error) {
	body, err := ConvertToJsonString(*collection)
	if err != nil {
		return nil, err
	}

	response, err := c.update(fmt.Sprintf("%s/%d", "collections", (*collection).ID), body)
	if err != nil {
		return nil, err
	}

	newCollection, err := ConvertFromJsonString[Collection](response)
	if err != nil {
		return nil, err
	}

	return newCollection, nil
}

func (c *Client) DeleteCollection(collectionId int) error {
	err := c.delete(fmt.Sprintf("%s/%d", "collections", collectionId))
	if err != nil {
		return err
	}

	return nil
}
