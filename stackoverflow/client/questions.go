package client

import (
	"fmt"
	"log"
	"net/url"
	"strings"
)

func (c *Client) GetQuestions(questionIDs *[]int, filter *string) (*[]Question, error) {
	response, err := c.get("questions", questionIDs, filter)
	if err != nil {
		return nil, err
	}

	questions, err := UnwrapResponseItems[Question](response)
	if err != nil {
		return nil, err
	}

	return questions, nil
}

func (c *Client) CreateQuestion(question *Question) (*Question, error) {
	response, err := c.create("questions/add", GenerateQuestionFormData(question))
	if err != nil {
		return nil, err
	}

	questions, err := UnwrapResponseItems[Question](response)
	if err != nil {
		return nil, err
	}

	if len(*questions) != 1 {
		return nil, fmt.Errorf("response wrapper does not contain expected number of items (1); item length is %d", len(*questions))
	}

	newQuestion := (*questions)[0]

	return &newQuestion, nil
}

func (c *Client) UpdateQuestion(question *Question) (*Question, error) {
	response, err := c.update(fmt.Sprintf("%s%d%s", "questions/", question.ID, "/edit"), GenerateQuestionFormData(question))
	if err != nil {
		return nil, err
	}

	questions, err := UnwrapResponseItems[Question](response)
	if err != nil {
		return nil, err
	}

	if len(*questions) != 1 {
		return nil, fmt.Errorf("response wrapper does not contain expected number of items (1); item length is %d", len(*questions))
	}

	newQuestion := (*questions)[0]

	return &newQuestion, nil
}

func (c *Client) DeleteQuestion(questionId int) error {
	err := c.delete(fmt.Sprintf("%s%d%s", "questions/", questionId, "/delete"))
	if err != nil {
		return err
	}

	return nil
}

func GenerateQuestionFormData(question *Question) *string {
	formData := fmt.Sprintf("title=%s&body=%s&tags=%s&preview=%t&filter=%s", url.QueryEscape(question.Title), url.QueryEscape(question.BodyMarkdown), strings.Join(question.Tags, ","), question.Preview, question.Filter)
	log.Printf("Form data: %s", formData)
	return &formData
}
