package ores

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var damagingTestData = map[string]interface{}{
	"prediction": true,
	"probability": map[string]interface{}{
		"true":  0.9,
		"false": 0.1,
	},
}

func TestDamaging(t *testing.T) {
	dmg := new(Damaging)
	assert.NoError(t, dmg.fromMap(damagingTestData))
	assert.Equal(t, dmg.Prediction, damagingTestData["prediction"].(bool))
	assert.Equal(t, dmg.Probability.True, damagingTestData["probability"].(map[string]interface{})["true"].(float64))
	assert.Equal(t, dmg.Probability.False, damagingTestData["probability"].(map[string]interface{})["false"].(float64))
}
