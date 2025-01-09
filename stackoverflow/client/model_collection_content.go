package client

type CollectionContent struct {
	ID    int    `json:"id,omitempty"`
	Type  string `json:"type"`
	Title string `json:"title"`
}
