package client

type Answer struct {
	ID           int      `json:"answer_id,omitempty"`
	Accepted     bool     `json:"accepted"`
	BodyMarkdown string   `json:"body_markdown"`
	Preview      bool     `json:"preview"`
	QuestionID   int      `json:"question_id"`
	Title        string   `json:"title"`
	Tags         []string `json:"tags"`
	Filter       string   `json:"filter"`
}
