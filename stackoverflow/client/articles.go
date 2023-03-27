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

func (c *Client) GetArticles(articleIDs *[]int) (*[]Article, error) {
	ids := make([]string, len(*articleIDs))
	for i, articleID := range *articleIDs {
		ids[i] = strconv.Itoa(articleID)
	}
	log.Printf("%s/%s?team=%s&order=desc&sort=creation&filter=omhz)aiL)ei3-sat(rZKVugTgq0f6)", "articles", strings.Join(ids, ";"), c.TeamName)
	route := fmt.Sprintf("%s/%s?team=%s&order=desc&sort=creation&filter=omhz)aiL)ei3-sat(rZKVugTgq0f6)", "articles", strings.Join(ids, ";"), c.TeamName)
	req, err := http.NewRequest("GET", fmt.Sprintf("%s%s", c.BaseURL, route), nil)
	if err != nil {
		return nil, err
	}

	body, err := c.doRequest(req)
	if err != nil {
		return nil, err
	}

	responseWrapper := Wrapper[Article]{}
	err = json.Unmarshal(body, &responseWrapper)
	if err != nil {
		return nil, err
	}

	articles := responseWrapper.Items

	return &articles, nil
}

func (c *Client) CreateArticle(article *Article) (*Article, error) {
	//rb, err := json.Marshal(&article)
	//if err != nil {
	//	return nil, err
	//}
	//buf := new(strings.Builder)
	//io.Copy(buf, strings.NewReader((string(rb))))
	//log.Printf("Deserialized JSON bytes: %s", buf.String())
	//req, err := http.NewRequest("POST", fmt.Sprintf("%s%s?team=%s", c.BaseURL, "articles/add", c.TeamName), bytes.NewReader(rb))

	formData := GenerateArticleFormData(article)
	req, err := http.NewRequest("POST", fmt.Sprintf("%s%s?team=%s", c.BaseURL, "articles/add", c.TeamName), strings.NewReader(formData))
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

	responseWrapper := Wrapper[Article]{}
	err = json.Unmarshal(body, &responseWrapper)
	if err != nil {
		return nil, err
	}

	if len(responseWrapper.Items) != 1 {
		return nil, fmt.Errorf("response wrapper does not contain expected number of items (1); item length is %d", len(responseWrapper.Items))
	}

	newArticle := responseWrapper.Items[0]

	return &newArticle, nil
}

func (c *Client) UpdateArticle(article *Article) (*Article, error) {
	//rb, err := json.Marshal(&article)
	//if err != nil {
	//	return nil, err
	//}
	//req, err := http.NewRequest("POST", fmt.Sprintf("%s%s%s%s?team=%s", c.BaseURL, "articles/", strconv.Itoa(article.ID), "/edit", c.TeamName), strings.NewReader(string(rb)))

	formData := GenerateArticleFormData(article)
	req, err := http.NewRequest("POST", fmt.Sprintf("%s%s%s%s?team=%s", c.BaseURL, "articles/", strconv.Itoa(article.ID), "/edit", c.TeamName), strings.NewReader(formData))
	if err != nil {
		return nil, err
	}

	body, err := c.doRequest(req)
	if err != nil {
		return nil, err
	}

	responseWrapper := Wrapper[Article]{}
	err = json.Unmarshal(body, &responseWrapper)
	if err != nil {
		return nil, err
	}

	if len(responseWrapper.Items) != 1 {
		return nil, fmt.Errorf("response wrapper does not contain expected number of items (1); item length is %d", len(responseWrapper.Items))
	}

	newArticle := responseWrapper.Items[0]

	return &newArticle, nil
}

func (c *Client) DeleteArticle(articleId int) error {
	req, err := http.NewRequest("POST", fmt.Sprintf("%s%s%s%s?team=%s&filter=omhz)aiL)ei3-sat(rZKVugTgq0f6)", c.BaseURL, "articles/", strconv.Itoa(articleId), "/delete", c.TeamName), nil)
	if err != nil {
		return err
	}

	_, err2 := c.doRequest(req)
	if err2 != nil {
		return err
	}

	return nil
}

func GenerateArticleFormData(article *Article) string {
	formData := fmt.Sprintf("title=%s&body=%s&tags=%s&article_type=%s&filter=%s", url.QueryEscape(article.Title), url.QueryEscape(article.BodyMarkdown), strings.Join(article.Tags, ","), article.ArticleType, article.Filter)
	log.Printf("Form data: %s", formData)
	return formData
}
