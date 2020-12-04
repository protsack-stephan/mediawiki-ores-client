package ores

import (
	"context"
	"fmt"
	"net/http"
)

// DamagingRequest request for scores of damaging model
type DamagingRequest struct {
	client *Client
}

// ScoreOne get damaging score for revision
func (dr *DamagingRequest) ScoreOne(ctx context.Context, dbName string, rev int) (*Damaging, error) {
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
