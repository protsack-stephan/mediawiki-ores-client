package ores

import (
	"context"
	"fmt"
	"net/http"
)

// Damaging struct for model
type Damaging struct {
	Prediction  bool `json:"prediction"`
	Probability struct {
		False float64 `json:"false"`
		True  float64 `json:"true"`
	} `json:"probability"`
}

type damagingRequest struct {
	client *Client
}

// DamagingScore get damaging score for revision
func (dr *damagingRequest) ScoreOne(ctx context.Context, dbName string, rev int) (*Damaging, error) {
	score := new(Damaging)

	if !ModelDamaging.Supports(dbName) {
		return score, ErrModelNotSupported
	}

	data, status, err := req(ctx, dr.client.httpClient, http.MethodGet, dr.client.url+fmt.Sprintf("/%s/%d/%s", dbName, rev, ModelDamaging), nil)

	if err != nil {
		return score, err
	}

	if status != http.StatusOK {
		return score, fmt.Errorf(errBadRequestMsg, status, string(data))
	}

	models, err := parse(data, dbName, ModelDamaging, rev)

	if err != nil {
		return score, err
	}

	model := models[rev]

	switch model["prediction"].(type) {
	case bool:
		score.Prediction = model["prediction"].(bool)
	default:
		return score, ErrInvalidServerResponse
	}

	var probability map[string]interface{}

	switch model["probability"].(type) {
	case map[string]interface{}:
		probability = model["probability"].(map[string]interface{})
	default:
		return score, ErrInvalidServerResponse
	}

	switch probability["true"].(type) {
	case float64:
		score.Probability.True = probability["true"].(float64)
	default:
		return score, ErrInvalidServerResponse
	}

	switch probability["false"].(type) {
	case float64:
		score.Probability.False = probability["false"].(float64)
	default:
		return score, ErrInvalidServerResponse
	}

	return score, nil
}
