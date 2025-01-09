package client

import (
	"fmt"
)

func (c *Client) GetQuestion(questionID *int) (*Question[Tag], error) {
	response, err := c.get(fmt.Sprintf("%s/%d", "questions", *questionID))
	if err != nil {
		return nil, err
	}

	question, err := ConvertFromJsonString[Question[Tag]](response)
	if err != nil {
		return nil, err
	}

	return question, nil
}

func (c *Client) CreateQuestion(question *Question[string]) (*Question[Tag], error) {
	body, err := ConvertToJsonString(*question)
	if err != nil {
		return nil, err
	}

	response, err := c.create("questions", body)
	if err != nil {
		return nil, err
	}

	newQuestion, err := ConvertFromJsonString[Question[Tag]](response)
	if err != nil {
		return nil, err
	}

	return newQuestion, nil
}

func (c *Client) UpdateQuestion(question *Question[string]) (*Question[Tag], error) {
	body, err := ConvertToJsonString(*question)
	if err != nil {
		return nil, err
	}

	response, err := c.update(fmt.Sprintf("%s/%d", "questions", (*question).ID), body)
	if err != nil {
		return nil, err
	}

	newQuestion, err := ConvertFromJsonString[Question[Tag]](response)
	if err != nil {
		return nil, err
	}

	return newQuestion, nil
}

func (c *Client) DeleteQuestion(questionId int) error {
	err := c.delete(fmt.Sprintf("%s/%d", "questions", questionId))
	if err != nil {
		return err
	}

	return nil
}
