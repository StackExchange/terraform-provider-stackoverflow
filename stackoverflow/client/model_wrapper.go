package client

type Wrapper[T any] struct {
	Backoff        int    `json:"backoff"`
	ErrorID        int    `json:"error_id"`
	ErrorMessage   string `json:"error_message"`
	ErrorName      string `json:"error_name"`
	HasMore        bool   `json:"has_more"`
	Items          []T    `json:"items"`
	Page           int    `json:"page"`
	PageSize       int    `json:"page_size"`
	QuotaMax       int    `json:"quota_max"`
	QuotaRemaining int    `json:"quota_remaining"`
	Total          int    `json:"total"`
	ItemType       string `json:"type"`
}
