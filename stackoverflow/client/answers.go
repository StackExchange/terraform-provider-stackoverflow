package client

import (
	"fmt"
)

func (c *Client) GetAnswer(questionID *int, answerID *int) (*Answer[Tag], error) {
	response, err := c.get(fmt.Sprintf("%s/%d/%s/%d", "questions", *questionID, "answers", *answerID))
	if err != nil {
		return nil, err
	}

	answer, err := ConvertFromJsonString[Answer[Tag]](response)
	if err != nil {
		return nil, err
	}

	return answer, nil
}

func (c *Client) CreateAnswer(answer *Answer[string]) (*Answer[Tag], error) {
	body, err := ConvertToJsonString(*answer)
	if err != nil {
		return nil, err
	}

	response, err := c.create(fmt.Sprintf("%s/%d/%s", "questions", (*answer).QuestionID, "answers"), body)
	if err != nil {
		return nil, err
	}

	newAnswer, err := ConvertFromJsonString[Answer[Tag]](response)
	if err != nil {
		return nil, err
	}

	return newAnswer, nil
}

func (c *Client) UpdateAnswer(answer *Answer[string]) (*Answer[Tag], error) {
	body, err := ConvertToJsonString(*answer)
	if err != nil {
		return nil, err
	}

	response, err := c.update(fmt.Sprintf("%s/%d/%s/%d", "questions", answer.QuestionID, "answers", (*answer).ID), body)
	if err != nil {
		return nil, err
	}

	newAnswer, err := ConvertFromJsonString[Answer[Tag]](response)
	if err != nil {
		return nil, err
	}

	return newAnswer, nil
}

func (c *Client) DeleteAnswer(questionId int, answerId int) error {
	err := c.delete(fmt.Sprintf("%s/%d/%s/%d", "questions", questionId, "answers", answerId))
	if err != nil {
		return err
	}

	return nil
}
