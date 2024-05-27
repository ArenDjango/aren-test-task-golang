package pg

import (
	"context"
	"go.opentelemetry.io/otel"

	"github.com/ArenDjango/golang-test-task/model"
)

type RatesPgRepo struct {
	db *DB
}

func NewRatesRepo(db *DB) *RatesPgRepo {
	return &RatesPgRepo{db: db}
}

func (repo *RatesPgRepo) CreateRate(ctx context.Context, rate *model.DBRates) (*model.DBRates, error) {
	ctx, span := otel.Tracer("").Start(ctx, "RatesPgRepo.CreateRate")
	defer span.End()

	_, err := repo.db.Model(rate).
		Returning("*").
		Insert()
	if err != nil {
		return nil, err
	}
	return rate, nil
}
