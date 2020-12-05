package ores

import (
	"context"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

const damagingReqTestDBName = "enwiki"
const damagingReqTestRev = 1
const damagingReqTestModel = ModelDamaging
const damagingReqTestProbabilityFalse = 0.23
const damagingReqTestProbabilityTrue = 0.73
const damagingReqTestPrediction = true
const damagingReqTestResponseTemplate = `{"%s":{"models":{"damaging":{"version":"0.5.1"}},"scores":{"%d":{"damaging":{"score":{"prediction":%t,"probability":{"false":%f,"true":%f}}}}}}}`

var damagingTestURL = fmt.Sprintf("/%s/%d/%s", damagingReqTestDBName, damagingReqTestRev, damagingReqTestModel)

func createDamagingServer() http.Handler {
	router := http.NewServeMux()

	router.HandleFunc(damagingTestURL, func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		data := fmt.Sprintf(
			damagingReqTestResponseTemplate,
			damagingReqTestDBName,
			damagingReqTestRev,
			damagingReqTestPrediction,
			damagingReqTestProbabilityFalse,
			damagingReqTestProbabilityTrue)
		w.Write([]byte(data))
	})

	return router
}

func TestDamagingRequestScoreOne(t *testing.T) {
	srv := httptest.NewServer(createDamagingServer())
	defer srv.Close()

	client := NewClient()
	client.url = srv.URL
	score, err := client.Damaging().ScoreOne(context.Background(), damagingReqTestDBName, damagingReqTestRev)

	assert.NoError(t, err)
	assert.Equal(t, damagingReqTestPrediction, score.Prediction)
	assert.Equal(t, damagingReqTestProbabilityTrue, score.Probability.True)
	assert.Equal(t, damagingReqTestProbabilityFalse, score.Probability.False)
}
