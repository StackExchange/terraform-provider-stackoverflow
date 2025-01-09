package client

type Collection struct {
	ID          int                 `json:"id,omitempty"`
	Content     []CollectionContent `json:"content"`
	ContentIds  []int               `json:"contentIds"`
	Description string              `json:"description"`
	Title       string              `json:"title"`
}
