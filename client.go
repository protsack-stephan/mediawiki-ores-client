package ores

import (
	"errors"
	"net/http"
)

const baseURL = "https://ores.wikimedia.org/v3/scores"
const maxConnsPerHost = 4

// ErrModelNotSupported model is not supported error
var ErrModelNotSupported = errors.New("model is not supported")

// ErrInvalidServerResponse server sent response in invalid format
var ErrInvalidServerResponse = errors.New("invalid server response")

// NewClient create new ORES client
func NewClient() *Client {
	client := &Client{
		url: baseURL,
		httpClient: &http.Client{
			Transport: &http.Transport{
				MaxConnsPerHost: maxConnsPerHost,
			},
		},
	}

	client.Damaging = &damagingRequest{
		client,
	}

	return client
}

// Client for ORES API
type Client struct {
	url        string
	httpClient *http.Client
	Damaging   *damagingRequest
}
