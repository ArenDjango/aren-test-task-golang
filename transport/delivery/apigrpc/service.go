package apigrpc

import (
	"context"
	"github.com/ArenDjango/golang-test-task/config"
	"github.com/ArenDjango/golang-test-task/protos/rates"
	"github.com/ArenDjango/golang-test-task/service"
	"time"
)

type APIService struct {
	cfg          *config.Config
	ratesService service.RatesService
	rates.RatesServiceServer
}

func NewAPIService(
	cfg *config.Config,
	ratesService service.RatesService,
) rates.RatesServiceServer {
	return &APIService{
		cfg:          cfg,
		ratesService: ratesService,
	}
}

func (s *APIService) GetRates(
	ctx context.Context,
	request *rates.GetRatesRequest,
) (result *rates.GetRatesResponse, err error) {
	res, err := s.ratesService.GetRates(ctx, s.cfg.GARANTEX_API_URL)
	if err != nil {
		return nil, err
	}
	return &rates.GetRatesResponse{
		AskPrice:  res.AskPrice,
		BidPrice:  res.BidPrice,
		Timestamp: time.Unix(res.TimeStamp.Unix(), 0).Format(time.RFC3339),
	}, nil
}
