package client

type Image struct {
	ID         int    `json:"id,omitempty"`
	URL        string `json:"url"`
	Base64Data string `json:"base64Data"`
}
