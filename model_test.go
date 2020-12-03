package ores

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var modelTestDBName = "enwiki"

func TestModel(t *testing.T) {
	assert.True(t, ModelDamaging.Supports(modelTestDBName))
}
