package ores

import "net/http"

// NewBuilder new builder instance
func NewBuilder() *ClientBuilder {
	return &ClientBuilder{
		NewClient(),
	}
}

// ClientBuilder configurator for ORES client
type ClientBuilder struct {
	client *Client
}

// URL setting client base URL
func (cb *ClientBuilder) URL(url string) *ClientBuilder {
	cb.client.url = url
	return cb
}

// HTTPClient set custom HTTP client if needed
func (cb *ClientBuilder) HTTPClient(httpClient *http.Client) *ClientBuilder {
	cb.client.httpClient = httpClient
	return cb
}

// Build create new client instance based on provided configuration
func (cb *ClientBuilder) Build() *Client {
	return cb.client
}
