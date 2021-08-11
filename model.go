package ores

import (
	"encoding/json"
	"log"
)

var support = map[string]map[Model]bool{}
var models = []Model{ModelDamaging, ModelArticleQuality, ModelGoodFaith}

func init() {
	info := make(map[string]struct {
		Models map[Model]ModelInfo `json:"models"`
	})

	if err := json.Unmarshal([]byte(config), &info); err != nil {
		log.Panic(err)
	}

	for dbName, config := range info {
		support[dbName] = make(map[Model]bool)

		for _, model := range models {
			if _, exists := config.Models[model]; exists {
				support[dbName][model] = true
			}
		}
	}
}

// Supported ores models
const (
	ModelDamaging       Model = "damaging"
	ModelGoodFaith      Model = "goodfaith"
	ModelArticleQuality Model = "articlequality"
)

// ModelInfo all the model meta data
type ModelInfo struct {
	Version string `json:"version"`
}

// Model ORES scoring type
type Model string

// Supports check if model is available for certain database
func (m Model) Supports(dbName string) bool {
	_, exists := support[dbName][m]
	return exists
}
