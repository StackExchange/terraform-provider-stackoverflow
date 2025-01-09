package client

type Article[T string | Tag] struct {
	ID           int    `json:"id,omitempty"`
	ArticleType  string `json:"type"`
	Body         string `json:"body"`
	BodyMarkdown string `json:"bodyMarkdown"`
	Title        string `json:"title"`
	Tags         []T    `json:"tags"`
}
