package service

import (
	"context"
	"encoding/json"
	"github.com/ArenDjango/golang-test-task/model"
	"github.com/ArenDjango/golang-test-task/store"
	"github.com/google/uuid"
	"github.com/pkg/errors"
	"go.opentelemetry.io/otel"
	"io/ioutil"
	"net/http"
	"strconv"
)

type RatesAPIService struct {
	ctx        context.Context
	store      *store.Store
	httpClient *http.Client
}

func NewRatesAPIService(ctx context.Context, store *store.Store, httpClient *http.Client) *RatesAPIService {
	if httpClient == nil {
		httpClient = http.DefaultClient
	}
	return &RatesAPIService{
		ctx:        ctx,
		store:      store,
		httpClient: httpClient,
	}
}

type Rate struct {
	Price  string `json:"price"`
	Volume string `json:"volume"`
	Amount string `json:"amount"`
	Factor string `json:"factor"`
	Type   string `json:"type"`
}
type RatesResponse struct {
	Timestamp int64  `json:"timestamp"`
	Asks      []Rate `json:"asks"`
	Bids      []Rate `json:"bids"`
}

func (svc *RatesAPIService) GetRates(ctx context.Context, apiUrl string) (*model.Rates, error) {
	ctx, span := otel.Tracer("").Start(ctx, "RatesAPIService.GetRates")
	defer span.End()
	//apiUrl := svc.cfg.GARANTEX_API_URL
	req, err := http.NewRequestWithContext(ctx, "GET", apiUrl, nil)
	if err != nil {
		return nil, errors.Wrap(err, "failed to create new request")
	}

	res, err := svc.httpClient.Do(req)
	if err != nil {
		return nil, errors.Wrap(err, "failed to execute request")
	}
	defer res.Body.Close()

	body, readErr := ioutil.ReadAll(res.Body)
	if readErr != nil {
		return nil, errors.Wrap(readErr, "failed to read response body")
	}

	var apiResponse RatesResponse
	if err := json.Unmarshal(body, &apiResponse); err != nil {
		return nil, errors.Wrap(err, "failed to unmarshal response")
	}

	if len(apiResponse.Asks) == 0 || len(apiResponse.Bids) == 0 {
		return nil, errors.New("no data found")
	}

	ask, err := strconv.ParseFloat(apiResponse.Asks[0].Price, 64)
	if err != nil {
		return nil, errors.Wrap(err, "failed to parse ask price")
	}
	bid, err := strconv.ParseFloat(apiResponse.Bids[0].Price, 64)
	if err != nil {
		return nil, errors.Wrap(err, "failed to parse bid price")
	}

	dbRates := model.Rates{
		ID:       uuid.New(),
		AskPrice: ask,
		BidPrice: bid,
	}

	createdDBRate, err := svc.store.Rates.CreateRate(ctx, dbRates.ToDB())
	if err != nil {
		return nil, errors.Wrap(err, "failed to create rate in database")
	}

	return createdDBRate.ToWeb(), nil
}
