package ores

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var modelTestSupportedDBNames = []string{"enwiki", "ukwiki", "ruwiki", "eswikibooks", "fakewiki"}
var modelTestUnsupportedDBNames = []string{"bnwiki", "elwiki", "euwiki", "viwiki"}

func TestModel(t *testing.T) {
	for _, dbName := range modelTestSupportedDBNames {
		assert.True(t, ModelDamaging.Supports(dbName))
	}

	for _, dbName := range modelTestUnsupportedDBNames {
		assert.False(t, ModelDamaging.Supports(dbName))
	}
}
