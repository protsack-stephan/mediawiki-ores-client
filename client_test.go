package ores

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"
	"strconv"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

const clientTestURL = "/scores"
const clientTestDBName = "testwiki"
const clientTestFailedDBName = "uknownwiki"

func createTestClientServer(assert func(models []Model, revs []int)) http.Handler {
	router := http.NewServeMux()

	router.HandleFunc(fmt.Sprintf("%s/%s", clientTestURL, clientTestDBName), func(rw http.ResponseWriter, r *http.Request) {
		models := []Model{}

		for _, model := range strings.Split(r.URL.Query().Get("models"), "|") {
			models = append(models, Model(model))
		}

		revs := []int{}

		for _, rev := range strings.Split(r.URL.Query().Get("revids"), "|") {
			if rev, err := strconv.Atoi(rev); err == nil {
				revs = append(revs, rev)
			}
		}

		assert(models, revs)
		data, err := ioutil.ReadFile("./testdata/response.json")

		if err != nil {
			log.Panic(err)
		}

		_, _ = rw.Write(data)
	})

	return router
}

func TestClient(t *testing.T) {
	assert := assert.New(t)
	ctx := context.Background()

	t.Run("score many success", func(t *testing.T) {
		errRev := 122
		scoreModels := []Model{ModelDamaging, ModelArticleQuality, ModelGoodFaith}
		scoreRevs := []int{1, 12, errRev}

		srv := httptest.NewServer(createTestClientServer(func(models []Model, revs []int) {
			assert.Equal(scoreModels, models)
			assert.Equal(scoreRevs, revs)
		}))
		defer srv.Close()

		client := NewClient()
		client.url = fmt.Sprintf("%s%s", srv.URL, clientTestURL)
		res, err := client.ScoreMany(ctx, clientTestDBName, scoreModels, scoreRevs...)
		assert.NoError(err)

		for _, rev := range scoreRevs {
			assert.Contains(res.Scores, rev)
			assert.NotNil(res.Scores[rev].Damaging)
			assert.NotNil(res.Scores[rev].Articlequality)
			assert.NotNil(res.Scores[rev].Goodfaith)
		}

		for _, model := range scoreModels {
			assert.Contains(res.Models, model)
		}

		assert.Nil(res.Scores[errRev].Damaging.Score)
		assert.Nil(res.Scores[errRev].Articlequality.Score)
		assert.Nil(res.Scores[errRev].Goodfaith.Score)
		assert.NotNil(res.Scores[errRev].Damaging.Error)
		assert.NotNil(res.Scores[errRev].Articlequality.Error)
		assert.NotNil(res.Scores[errRev].Goodfaith.Error)
	})

	t.Run("score many error", func(t *testing.T) {
		scoreModels := []Model{ModelDamaging, ModelArticleQuality, ModelGoodFaith}
		scoreRevs := []int{1, 12, 122}

		srv := httptest.NewServer(createTestClientServer(nil))
		defer srv.Close()

		client := NewClient()
		client.url = fmt.Sprintf("%s%s", srv.URL, clientTestURL)
		scores, err := client.ScoreMany(ctx, clientTestFailedDBName, scoreModels, scoreRevs...)
		assert.Nil(scores)
		assert.Error(err)
		assert.Contains(err.Error(), strconv.Itoa(http.StatusNotFound))
	})
}
