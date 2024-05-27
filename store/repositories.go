package store

import (
	"context"
	"github.com/ArenDjango/golang-test-task/model"
)

//go:generate mockery --dir . --name RatesRepo --output mocks
type RatesRepo interface {
	CreateRate(ctx context.Context, rate *model.DBRates) (*model.DBRates, error)
}
