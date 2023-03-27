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

func (c *Client) GetAnswers(answerIDs *[]int) (*[]Answer, error) {
	ids := make([]string, len(*answerIDs))
	for i, answerID := range *answerIDs {
		ids[i] = strconv.Itoa(answerID)
	}
	log.Printf("%s/%s?team=%s&order=desc&sort=creation&filter=omhz)aiL)ei3-sat(rZKVugTgq0f6)", "answers", strings.Join(ids, ";"), c.TeamName)
	route := fmt.Sprintf("%s/%s?team=%s&order=desc&sort=creation&filter=omhz)aiL)ei3-sat(rZKVugTgq0f6)", "answers", strings.Join(ids, ";"), c.TeamName)
	req, err := http.NewRequest("GET", fmt.Sprintf("%s%s", c.BaseURL, route), nil)
	if err != nil {
		return nil, err
	}

	body, err := c.doRequest(req)
	if err != nil {
		return nil, err
	}

	responseWrapper := Wrapper[Answer]{}
	err = json.Unmarshal(body, &responseWrapper)
	if err != nil {
		return nil, err
	}

	answers := responseWrapper.Items

	return &answers, nil
}

func (c *Client) CreateAnswer(answer *Answer) (*Answer, error) {
	formData := GenerateAnswerFormData(answer, true)
	req, err := http.NewRequest("POST", fmt.Sprintf("%s%s%s%s?team=%s", c.BaseURL, "questions/", strconv.Itoa(answer.QuestionID), "/answers/add", c.TeamName), strings.NewReader(formData))
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

	responseWrapper := Wrapper[Answer]{}
	err = json.Unmarshal(body, &responseWrapper)
	if err != nil {
		return nil, err
	}

	if len(responseWrapper.Items) != 1 {
		return nil, fmt.Errorf("response wrapper does not contain expected number of items (1); item length is %d", len(responseWrapper.Items))
	}

	newAnswer := responseWrapper.Items[0]

	return &newAnswer, nil
}

func (c *Client) UpdateAnswer(answer *Answer) (*Answer, error) {
	formData := GenerateAnswerFormData(answer, false)
	req, err := http.NewRequest("POST", fmt.Sprintf("%s%s%s%s?team=%s", c.BaseURL, "answers/", strconv.Itoa(answer.ID), "/edit", c.TeamName), strings.NewReader(formData))
	if err != nil {
		return nil, err
	}

	body, err := c.doRequest(req)
	if err != nil {
		return nil, err
	}

	responseWrapper := Wrapper[Answer]{}
	err = json.Unmarshal(body, &responseWrapper)
	if err != nil {
		return nil, err
	}

	if len(responseWrapper.Items) != 1 {
		return nil, fmt.Errorf("response wrapper does not contain expected number of items (1); item length is %d", len(responseWrapper.Items))
	}

	newAnswer := responseWrapper.Items[0]

	return &newAnswer, nil
}

func (c *Client) DeleteAnswer(answerId int) error {
	req, err := http.NewRequest("POST", fmt.Sprintf("%s%s%s%s?team=%s&filter=omhz)aiL)ei3-sat(rZKVugTgq0f6)", c.BaseURL, "answers/", strconv.Itoa(answerId), "/delete", c.TeamName), nil)
	if err != nil {
		return err
	}

	_, err2 := c.doRequest(req)
	if err2 != nil {
		return err
	}

	return nil
}

func GenerateAnswerFormData(answer *Answer, isCreate bool) string {
	if isCreate {
		formData := fmt.Sprintf("id=%d&body=%s&preview=%t&filter=%s", answer.QuestionID, url.QueryEscape(answer.BodyMarkdown), answer.Preview, answer.Filter)
		log.Printf("Form data: %s", formData)
		return formData
	}

	formData := fmt.Sprintf("id=%d&body=%s&preview=%t&filter=%s", answer.ID, url.QueryEscape(answer.BodyMarkdown), answer.Preview, answer.Filter)
	log.Printf("Form data: %s", formData)
	return formData
}
