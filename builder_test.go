package ores

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

var builderTestHTTPClient = &http.Client{}

const builderTestURL = "http://test.org/v3/scores"

func TestBuilder(t *testing.T) {
	client := NewBuilder().
		URL(builderTestURL).
		HTTPClient(builderTestHTTPClient).
		Build()

	assert.NotNil(t, client)
	assert.Equal(t, client.url, builderTestURL)
	assert.Equal(t, client.httpClient, builderTestHTTPClient)
}
