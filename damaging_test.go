package ores

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

const damagingTestDBName = "enwiki"
const damagingTestRev = 1
const damagingTestModel = ModelDamaging
const damagingTestProbabilityFalse = 0.23
const damagingTestProbabilityTrue = 0.73
const damagingTestPrediction = true
const damagingTestResponseTemplate = `{"%s":{"models":{"damaging":{"version":"0.5.1"}},"scores":{"%d":{"damaging":{"score":{"prediction":%t,"probability":{"false":%f,"true":%f}}}}}}}`

var damagingTestURL = fmt.Sprintf("/%s/%d/%s", damagingTestDBName, damagingTestRev, damagingTestModel)
var damagingTestMap = map[string]interface{}{
	"prediction": true,
	"probability": map[string]interface{}{
		"true":  0.9,
		"false": 0.1,
	},
}

func createDamagingServer() http.Handler {
	router := http.NewServeMux()

	router.HandleFunc(damagingTestURL, func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		data := fmt.Sprintf(
			damagingTestResponseTemplate,
			damagingTestDBName,
			damagingTestRev,
			damagingTestPrediction,
			damagingTestProbabilityFalse,
			damagingTestProbabilityTrue)
		_, err := w.Write([]byte(data))

		if err != nil {
			log.Panic(err)
		}
	})

	return router
}

func TestDamagingScoreOne(t *testing.T) {
	srv := httptest.NewServer(createDamagingServer())
	defer srv.Close()

	client := NewClient()
	client.url = srv.URL
	score, err := client.Damaging.ScoreOne(context.Background(), damagingTestDBName, damagingTestRev)

	assert.NoError(t, err)
	assert.Equal(t, damagingTestPrediction, score.Prediction)
	assert.Equal(t, damagingTestProbabilityTrue, score.Probability.True)
	assert.Equal(t, damagingTestProbabilityFalse, score.Probability.False)
}

func TestDamagingFromMap(t *testing.T) {
	dmg := new(Damaging)
	assert.NoError(t, dmg.fromMap(damagingTestMap))
	assert.Equal(t, dmg.Prediction, damagingTestMap["prediction"].(bool))
	assert.Equal(t, dmg.Probability.True, damagingTestMap["probability"].(map[string]interface{})["true"].(float64))
	assert.Equal(t, dmg.Probability.False, damagingTestMap["probability"].(map[string]interface{})["false"].(float64))
}
