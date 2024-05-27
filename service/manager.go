package service

import (
	"context"
	"github.com/ArenDjango/golang-test-task/config"
	"net/http"

	"github.com/ArenDjango/golang-test-task/store"
	"github.com/pkg/errors"
)

type Manager struct {
	Rates RatesService
}

func NewManager(ctx context.Context, store *store.Store, cfg *config.Config) (*Manager, error) {
	if store == nil {
		return nil, errors.New("No store provided")
	}
	return &Manager{
		Rates: NewRatesAPIService(ctx, store, http.DefaultClient),
	}, nil
}
