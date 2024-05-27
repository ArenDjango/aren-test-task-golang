package service

import (
	"context"

	"github.com/ArenDjango/golang-test-task/model"
)

//go:generate mockery --dir . --name RatesService --output ./mocks
type RatesService interface {
	GetRates(context.Context, string) (*model.Rates, error)
}
