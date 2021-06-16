package ores

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var modelTestDamagingSupported = []string{"enwiki", "ukwiki", "ruwiki", "eswikibooks", "fakewiki"}
var modelTestDamagingNotSupported = []string{"bnwiki", "elwiki", "euwiki", "viwiki"}

var modelTestArticleQualitySupported = []string{"ukwiki", "trwiki", "testwiki"}
var modelTestArticleQualityNotSupported = []string{"viwiki", "zhwiki", "wikidatawiki", "tawiki"}

var modelTestGoodFaithSupported = []string{"zhwiki", "wikidatawiki", "testwiki", "ukwiki"}
var modelTestGoodFaithNotSupported = []string{"viwiki", "tawiki", "iswiki"}

func TestModel(t *testing.T) {
	assert := assert.New(t)

	t.Run("damaging model true", func(t *testing.T) {
		for _, dbName := range modelTestDamagingSupported {
			assert.True(ModelDamaging.Supports(dbName))
		}
	})

	t.Run("damaging model false", func(t *testing.T) {
		for _, dbName := range modelTestDamagingNotSupported {
			assert.False(ModelDamaging.Supports(dbName))
		}
	})

	t.Run("articlequality model true", func(t *testing.T) {
		for _, dbName := range modelTestArticleQualitySupported {
			assert.True(ModelArticleQuality.Supports(dbName))
		}
	})

	t.Run("articlequality model false", func(t *testing.T) {
		for _, dbName := range modelTestArticleQualityNotSupported {
			assert.False(ModelArticleQuality.Supports(dbName))
		}
	})

	t.Run("goodfaith model true", func(t *testing.T) {
		for _, dbName := range modelTestGoodFaithSupported {
			assert.True(ModelGoodFaith.Supports(dbName))
		}
	})

	t.Run("goodfaith model false", func(t *testing.T) {
		for _, dbName := range modelTestGoodFaithNotSupported {
			assert.False(ModelGoodFaith.Supports(dbName))
		}
	})
}
