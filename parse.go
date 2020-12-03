package ores

import (
	"encoding/json"
	"errors"
)

func parse(data []byte, dbName string, model Model, revs ...int) (map[int]map[string]interface{}, error) {
	parsed := make(map[int]map[string]interface{})
	res := map[string]response{}

	err := json.Unmarshal(data, &res)

	if err != nil {
		return parsed, err
	}

	info, exists := res[dbName]

	if !exists {
		return parsed, ErrInvalidServerResponse
	}

	for _, rev := range revs {
		scores, exists := info.Scores[rev]

		if !exists {
			return parsed, ErrInvalidServerResponse
		}

		model, exists := scores[model]

		if !exists {
			return parsed, ErrInvalidServerResponse
		}

		if len(model.Error.Message) > 0 {
			return parsed, errors.New(model.Error.Message)
		}

		parsed[rev] = model.Score
	}

	return parsed, nil
}
