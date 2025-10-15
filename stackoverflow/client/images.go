package client

import (
	"fmt"
)

func (c *Client) GetImage(imageID *int) (*Image, error) {
	response, err := c.get(fmt.Sprintf("%s/%d", "images", *imageID))
	if err != nil {
		return nil, err
	}

	image, err := ConvertFromJsonString[Image](response)
	if err != nil {
		return nil, err
	}

	return article, nil
}

func (c *Client) CreateArticle(article *Article[string]) (*Article[Tag], error) {
	body, err := ConvertToJsonString(*article)
	if err != nil {
		return nil, err
	}

	response, err := c.create("articles", body)
	if err != nil {
		return nil, err
	}

	newArticle, err := ConvertFromJsonString[Article[Tag]](response)
	if err != nil {
		return nil, err
	}

	return newArticle, nil
}

func (c *Client) UpdateArticle(article *Article[string]) (*Article[Tag], error) {
	body, err := ConvertToJsonString(*article)
	if err != nil {
		return nil, err
	}

	response, err := c.update(fmt.Sprintf("%s/%d", "articles", (*article).ID), body)
	if err != nil {
		return nil, err
	}

	newArticle, err := ConvertFromJsonString[Article[Tag]](response)
	if err != nil {
		return nil, err
	}

	return newArticle, nil
}

func (c *Client) DeleteArticle(articleId int) error {
	err := c.delete(fmt.Sprintf("%s/%d", "articles", articleId))
	if err != nil {
		return err
	}

	return nil
}
