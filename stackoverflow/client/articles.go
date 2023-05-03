package client

import (
	"fmt"
	"log"
	"net/url"
	"strings"
)

func (c *Client) GetArticles(articleIDs *[]int, filter *string) (*[]Article, error) {
	response, err := c.get("articles", articleIDs, filter)
	if err != nil {
		return nil, err
	}

	articles, err := UnwrapResponseItems[Article](response)
	if err != nil {
		return nil, err
	}

	return articles, nil
}

func (c *Client) CreateArticle(article *Article) (*Article, error) {
	response, err := c.create("articles/add", GenerateArticleFormData(article))
	if err != nil {
		return nil, err
	}

	articles, err := UnwrapResponseItems[Article](response)
	if err != nil {
		return nil, err
	}

	if len(*articles) != 1 {
		return nil, fmt.Errorf("response wrapper does not contain expected number of items (1); item length is %d", len(*articles))
	}

	newArticle := (*articles)[0]

	return &newArticle, nil
}

func (c *Client) UpdateArticle(article *Article) (*Article, error) {
	response, err := c.update(fmt.Sprintf("%s%d%s", "articles/", article.ID, "/edit"), GenerateArticleFormData(article))
	if err != nil {
		return nil, err
	}

	articles, err := UnwrapResponseItems[Article](response)
	if err != nil {
		return nil, err
	}

	if len(*articles) != 1 {
		return nil, fmt.Errorf("response wrapper does not contain expected number of items (1); item length is %d", len(*articles))
	}

	newArticle := (*articles)[0]

	return &newArticle, nil
}

func (c *Client) DeleteArticle(articleId int) error {
	err := c.delete(fmt.Sprintf("%s%d%s", "articles/", articleId, "/delete"))
	if err != nil {
		return err
	}

	return nil
}

func GenerateArticleFormData(article *Article) *string {
	formData := fmt.Sprintf("title=%s&body=%s&tags=%s&article_type=%s&filter=%s", url.QueryEscape(article.Title), url.QueryEscape(article.BodyMarkdown), strings.Join(article.Tags, ","), article.ArticleType, article.Filter)
	log.Printf("Form data: %s", formData)
	return &formData
}
