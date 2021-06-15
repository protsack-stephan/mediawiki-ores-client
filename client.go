package ores

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

const baseURL = "https://ores.wikimedia.org/v3/scores"
const maxConnsPerHost = 4
const errBadRequestMsg = "status: '%d' body: '%s'"

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

	return client
}

// Client for ORES API
type Client struct {
	url        string
	httpClient *http.Client
}

func (c *Client) ScoreMany(ctx context.Context, dbName string, models []Model, revs ...int) (*Scores, error) {
	revsq := []string{}

	for _, rev := range revs {
		revsq = append(revsq, strconv.Itoa(rev))
	}

	modelsq := []string{}

	for _, model := range models {
		modelsq = append(modelsq, string(model))
	}

	query := make(url.Values)
	query.Add("models", strings.Join(modelsq, "|"))
	query.Add("revids", strings.Join(revsq, "|"))

	data, status, err := req(ctx, c.httpClient, http.MethodGet, fmt.Sprintf("%s/%s?%s", c.url, dbName, query.Encode()), nil)

	if err != nil {
		return nil, err
	}

	if status != http.StatusOK {
		return nil, fmt.Errorf(errBadRequestMsg, status, string(data))
	}

	res := make(map[string]*Scores)

	if err := json.Unmarshal(data, &res); err != nil {
		return nil, err
	}

	return res[dbName], nil
}
