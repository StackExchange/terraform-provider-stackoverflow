package client

import (
	"fmt"
	"log"
	"net/url"
)

func (c *Client) GetAnswers(answerIDs *[]int, filter *string) (*[]Answer, error) {
	response, err := c.get("answers", answerIDs, filter)
	if err != nil {
		return nil, err
	}

	answers, err := UnwrapResponseItems[Answer](response)
	if err != nil {
		return nil, err
	}

	return answers, nil
}

func (c *Client) CreateAnswer(answer *Answer) (*Answer, error) {
	response, err := c.create(fmt.Sprintf("%s%d%s", "questions/", answer.QuestionID, "/answers/add"), GenerateAnswerFormData(answer, true))
	if err != nil {
		return nil, err
	}

	answers, err := UnwrapResponseItems[Answer](response)
	if err != nil {
		return nil, err
	}

	if len(*answers) != 1 {
		return nil, fmt.Errorf("response wrapper does not contain expected number of items (1); item length is %d", len(*answers))
	}

	newAnswer := (*answers)[0]

	return &newAnswer, nil
}

func (c *Client) UpdateAnswer(answer *Answer) (*Answer, error) {
	response, err := c.update(fmt.Sprintf("%s%d%s", "answers/", answer.ID, "/edit"), GenerateAnswerFormData(answer, false))
	if err != nil {
		return nil, err
	}

	answers, err := UnwrapResponseItems[Answer](response)
	if err != nil {
		return nil, err
	}

	if len(*answers) != 1 {
		return nil, fmt.Errorf("response wrapper does not contain expected number of items (1); item length is %d", len(*answers))
	}

	newAnswer := (*answers)[0]

	return &newAnswer, nil
}

func (c *Client) DeleteAnswer(answerId int) error {
	err := c.delete(fmt.Sprintf("%s%d%s", "answers/", answerId, "/delete"))
	if err != nil {
		return err
	}

	return nil
}

func GenerateAnswerFormData(answer *Answer, isCreate bool) *string {
	formData := ""
	if isCreate {
		formData = fmt.Sprintf("id=%d&body=%s&preview=%t&filter=%s", answer.QuestionID, url.QueryEscape(answer.BodyMarkdown), answer.Preview, answer.Filter)
	} else {
		formData = fmt.Sprintf("id=%d&body=%s&preview=%t&filter=%s", answer.ID, url.QueryEscape(answer.BodyMarkdown), answer.Preview, answer.Filter)
	}

	log.Printf("Form data: %s", formData)
	return &formData
}
