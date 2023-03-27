package client

type Question struct {
	ID           int      `json:"question_id,omitempty"`
	BodyMarkdown string   `json:"body_markdown"`
	Preview      bool     `json:"preview"`
	Title        string   `json:"title"`
	Tags         []string `json:"tags"`
	Filter       string   `json:"filter"`
}
