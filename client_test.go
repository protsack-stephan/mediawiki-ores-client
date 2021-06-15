package ores

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestClient(t *testing.T) {
	assert := assert.New(t)
	client := NewClient()
	assert.NotNil(client)
	assert.NotNil(client.httpClient)
	assert.Equal(baseURL, client.url)
}
