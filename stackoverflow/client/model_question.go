package client

type Question[T string | Tag] struct {
	ID           int    `json:"id,omitempty"`
	Body         string `json:"body"`
	BodyMarkdown string `json:"bodyMarkdown"`
	Preview      bool   `json:"preview"`
	Title        string `json:"title"`
	Tags         []T    `json:"tags"`
}
