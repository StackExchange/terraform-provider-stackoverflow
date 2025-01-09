package client

type Answer[T string | Tag] struct {
	ID           int    `json:"id,omitempty"`
	Accepted     bool   `json:"accepted"`
	Body         string `json:"body"`
	BodyMarkdown string `json:"bodyMarkdown"`
	Preview      bool   `json:"preview"`
	QuestionID   int    `json:"questionId"`
	Title        string `json:"title"`
	Tags         []T    `json:"tags"`
}
