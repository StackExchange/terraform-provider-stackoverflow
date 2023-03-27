package client

type Article struct {
	ID           int      `json:"article_id,omitempty"`
	ArticleType  string   `json:"article_type"`
	BodyMarkdown string   `json:"body_markdown"`
	Title        string   `json:"title"`
	Tags         []string `json:"tags"`
	Filter       string   `json:"filter"`
}
