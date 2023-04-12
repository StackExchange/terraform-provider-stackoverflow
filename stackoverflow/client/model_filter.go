package client

type Filter struct {
	ID      string   `json:"filter"`
	Include []string `json:"included_fields"`
	Exclude []string `json:"excluded_fields"`
	Unsafe  bool
}
