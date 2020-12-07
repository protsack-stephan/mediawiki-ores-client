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

var dmgReqTestOneURL = fmt.Sprintf("/%s/%d/%s", dmgReqTestDBName, dmgReqTestRev, dmgReqTestModel)
var dmgReqTestManyURL = "/"
var dmgReqTestData = map[int]struct {
	Rev              int
	Prediction       bool
	ProbabilityFalse float64
	ProbabilityTrue  float64
}{
	1: {
		1,
		false,
		0.7,
		0.3,
	},
	2: {
		2,
		true,
		0.3,
		0.6,
	},
}

const dmgReqTestDBName = "testwiki"
const dmgReqTestRev = 1
const dmgReqTestModel = ModelDamaging
const dmgReqTestProbabilityFalse = 0.23
const dmgReqTestProbabilityTrue = 0.73
const dmgReqTestPrediction = true
const dmgReqTestOneResponse = `{"%s":{"models":{"damaging":{"version":"0.5.1"}},"scores":{"%d":{"damaging":{"score":{"prediction":%t,"probability":{"false":%f,"true":%f}}}}}}}`
const dmgReqTestManyResponse = `{"%s":{"models":{"damaging":{"version":"0.5.1"}},"scores":{"%d":{"damaging":{"score":{"prediction":%t,"probability":{"false":%f,"true":%f}}}},"%d":{"damaging":{"score":{"prediction":%t,"probability":{"false":%f,"true":%f}}}}}}}`

func createDamagingServer() http.Handler {
	router := http.NewServeMux()

	router.HandleFunc(dmgReqTestOneURL, func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		body := fmt.Sprintf(
			dmgReqTestOneResponse,
			dmgReqTestDBName,
			dmgReqTestRev,
			dmgReqTestPrediction,
			dmgReqTestProbabilityFalse,
			dmgReqTestProbabilityTrue)

		_, err := w.Write([]byte(body))
		if err != nil {
			log.Panic(err)
		}
	})

	router.HandleFunc(dmgReqTestManyURL, func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		body := fmt.Sprintf(
			dmgReqTestManyResponse,
			dmgReqTestDBName,
			dmgReqTestData[1].Rev,
			dmgReqTestData[1].Prediction,
			dmgReqTestData[1].ProbabilityFalse,
			dmgReqTestData[1].ProbabilityTrue,
			dmgReqTestData[2].Rev,
			dmgReqTestData[2].Prediction,
			dmgReqTestData[2].ProbabilityFalse,
			dmgReqTestData[2].ProbabilityTrue)

		_, err := w.Write([]byte(body))
		if err != nil {
			log.Panic(err)
		}
	})

	return router
}

func TestDamagingRequestScoreOne(t *testing.T) {
	srv := httptest.NewServer(createDamagingServer())
	defer srv.Close()

	client := NewClient()
	client.url = srv.URL

	score, err := client.Damaging().ScoreOne(context.Background(), dmgReqTestDBName, dmgReqTestRev)

	assert.NoError(t, err)
	assert.Equal(t, dmgReqTestPrediction, score.Prediction)
	assert.Equal(t, dmgReqTestProbabilityTrue, score.Probability.True)
	assert.Equal(t, dmgReqTestProbabilityFalse, score.Probability.False)
}

func TestDamagingRequestScoreMany(t *testing.T) {
	srv := httptest.NewServer(createDamagingServer())
	defer srv.Close()

	client := NewClient()
	client.url = srv.URL

	revs := []int{}

	for rev := range dmgReqTestData {
		revs = append(revs, rev)
	}

	scores, err := client.Damaging().ScoreMany(context.Background(), dmgReqTestDBName, revs...)
	assert.NoError(t, err)

	for rev, info := range dmgReqTestData {
		assert.Contains(t, scores, rev)
		score := scores[rev]

		assert.Equal(t, info.Prediction, score.Prediction)
		assert.Equal(t, info.ProbabilityTrue, score.Probability.True)
		assert.Equal(t, info.ProbabilityFalse, score.Probability.False)
	}
}
