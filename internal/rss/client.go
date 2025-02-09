package rss

import (
	"net/http"
	"time"
)

// Client -
type Client struct {
	httpClient http.Client
	UserAgent  string
	cache      map[string]interface{}
}

// NewClient -
func NewClient(timeout time.Duration, uAgent string) Client {
	return Client{
		httpClient: http.Client{
			Timeout: timeout,
		},
		UserAgent: uAgent,
		cache:     make(map[string]interface{}),
	}
}
