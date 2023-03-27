package client

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

func (c *Client) GetQuestions(questionIDs *[]int) (*[]Question, error) {
	ids := make([]string, len(*questionIDs))
	for i, questionID := range *questionIDs {
		ids[i] = strconv.Itoa(questionID)
	}
	log.Printf("%s/%s?team=%s&order=desc&sort=creation&filter=omhz)aiL)ei3-sat(rZKVugTgq0f6)", "questions", strings.Join(ids, ";"), c.TeamName)
	route := fmt.Sprintf("%s/%s?team=%s&order=desc&sort=creation&filter=omhz)aiL)ei3-sat(rZKVugTgq0f6)", "questions", strings.Join(ids, ";"), c.TeamName)
	req, err := http.NewRequest("GET", fmt.Sprintf("%s%s", c.BaseURL, route), nil)
	if err != nil {
		return nil, err
	}

	body, err := c.doRequest(req)
	if err != nil {
		return nil, err
	}

	responseWrapper := Wrapper[Question]{}
	err = json.Unmarshal(body, &responseWrapper)
	if err != nil {
		return nil, err
	}

	questions := responseWrapper.Items

	return &questions, nil
}

func (c *Client) CreateQuestion(question *Question) (*Question, error) {
	formData := GenerateQuestionFormData(question)
	req, err := http.NewRequest("POST", fmt.Sprintf("%s%s?team=%s", c.BaseURL, "questions/add", c.TeamName), strings.NewReader(formData))
	if err != nil {
		return nil, err
	}

	body, err := c.doRequest(req)
	if err != nil {
		return nil, err
	}

	buf := new(strings.Builder)
	io.Copy(buf, strings.NewReader((string(body))))
	log.Printf("Response body: %s", buf.String())

	responseWrapper := Wrapper[Question]{}
	err = json.Unmarshal(body, &responseWrapper)
	if err != nil {
		return nil, err
	}

	if len(responseWrapper.Items) != 1 {
		return nil, fmt.Errorf("response wrapper does not contain expected number of items (1); item length is %d", len(responseWrapper.Items))
	}

	newQuestion := responseWrapper.Items[0]

	return &newQuestion, nil
}

func (c *Client) UpdateQuestion(question *Question) (*Question, error) {
	formData := GenerateQuestionFormData(question)
	req, err := http.NewRequest("POST", fmt.Sprintf("%s%s%s%s?team=%s", c.BaseURL, "questions/", strconv.Itoa(question.ID), "/edit", c.TeamName), strings.NewReader(formData))
	if err != nil {
		return nil, err
	}

	body, err := c.doRequest(req)
	if err != nil {
		return nil, err
	}

	responseWrapper := Wrapper[Question]{}
	err = json.Unmarshal(body, &responseWrapper)
	if err != nil {
		return nil, err
	}

	if len(responseWrapper.Items) != 1 {
		return nil, fmt.Errorf("response wrapper does not contain expected number of items (1); item length is %d", len(responseWrapper.Items))
	}

	newQuestion := responseWrapper.Items[0]

	return &newQuestion, nil
}

func (c *Client) DeleteQuestion(questionId int) error {
	req, err := http.NewRequest("POST", fmt.Sprintf("%s%s%s%s?team=%s&filter=omhz)aiL)ei3-sat(rZKVugTgq0f6)", c.BaseURL, "questions/", strconv.Itoa(questionId), "/delete", c.TeamName), nil)
	if err != nil {
		return err
	}

	_, err2 := c.doRequest(req)
	if err2 != nil {
		return err
	}

	return nil
}

func GenerateQuestionFormData(question *Question) string {
	formData := fmt.Sprintf("title=%s&body=%s&tags=%s&preview=%t&filter=%s", url.QueryEscape(question.Title), url.QueryEscape(question.BodyMarkdown), strings.Join(question.Tags, ","), question.Preview, question.Filter)
	log.Printf("Form data: %s", formData)
	return formData
}
