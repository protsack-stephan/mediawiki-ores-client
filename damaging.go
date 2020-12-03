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

func (dmg *Damaging) fromMap(score map[string]interface{}) error {
	switch score["prediction"].(type) {
	case bool:
		dmg.Prediction = score["prediction"].(bool)
	default:
		return ErrInvalidDataInterface
	}

	var probability map[string]interface{}

	switch score["probability"].(type) {
	case map[string]interface{}:
		probability = score["probability"].(map[string]interface{})
	default:
		return ErrInvalidDataInterface
	}

	switch probability["true"].(type) {
	case float64:
		dmg.Probability.True = probability["true"].(float64)
	default:
		return ErrInvalidDataInterface
	}

	switch probability["false"].(type) {
	case float64:
		dmg.Probability.False = probability["false"].(float64)
	default:
		return ErrInvalidDataInterface
	}

	return nil
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

	model, err := parse(data, dbName, ModelDamaging, rev)

	if err != nil {
		return score, err
	}

	return score, score.fromMap(model[rev])
}
