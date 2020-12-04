package ores

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"strings"
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

	models, err := parse(data, dbName, ModelDamaging, rev)

	if err != nil {
		return score, err
	}

	return score, score.fromMap(models[rev])
}

// ScoreMany get damaging score for multiple revisions
func (dr *DamagingRequest) ScoreMany(ctx context.Context, dbName string, revs ...int) (map[int]*Damaging, error) {
	scores := make(map[int]*Damaging)

	if !ModelDamaging.Supports(dbName) {
		return scores, ErrModelNotSupported
	}

	revids := make([]string, 0)

	for _, rev := range revs {
		revids = append(revids, strconv.Itoa(rev))
	}

	query := make(url.Values)
	query.Add("models", string(ModelDamaging))
	query.Add("revids", strings.Join(revids, "|"))

	data, status, err := req(ctx, dr.client.httpClient, http.MethodGet, dr.client.url+"/"+dbName+"?"+query.Encode(), nil)

	if err != nil {
		return scores, err
	}

	if status != http.StatusOK {
		return scores, fmt.Errorf(errBadRequestMsg, status, string(data))
	}

	models, err := parse(data, dbName, ModelDamaging, revs...)

	if err != nil {
		return scores, err
	}

	for revID, model := range models {
		score := new(Damaging)

		err = score.fromMap(model)

		if err != nil {
			return scores, err
		}

		scores[revID] = score
	}

	return scores, nil
}
