package ores

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"path/filepath"
	"runtime"
)

var support = map[string]map[Model]bool{}
var models = []Model{ModelDamaging, ModelArticleQuality, ModelGoodFaith}

func init() {
	info := make(map[string]struct {
		Models map[Model]ModelInfo `json:"models"`
	})

	_, b, _, _ := runtime.Caller(0)
	data, err := ioutil.ReadFile(fmt.Sprintf("%s/config/support.json", filepath.Dir(b)))

	if err != nil {
		log.Panic(err)
	}

	err = json.Unmarshal(data, &info)

	if err != nil {
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
