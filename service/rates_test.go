package service

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/ArenDjango/golang-test-task/model"
	"github.com/ArenDjango/golang-test-task/store"
	"github.com/ArenDjango/golang-test-task/store/mocks"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestGetRates(t *testing.T) {
	input := &model.Rates{
		ID:        uuid.New(),
		AskPrice:  3.4,
		BidPrice:  6.7,
		TimeStamp: time.Now(),
	}

	tests := []struct {
		name         string
		expectations func(ratesRepo *mocks.RatesRepo)
		input        *model.Rates
		err          error
	}{
		{
			name: "valid and found",
			expectations: func(ratesRepo *mocks.RatesRepo) {
				ratesRepo.On("CreateRate", mock.Anything, mock.Anything).Return(&model.DBRates{
					AskPrice: 1.234,
					BidPrice: 1.123,
				}, nil)
			},
			input: input,
		},
		// also we can add other cases,но лень))
	}
	for _, test := range tests {
		t.Logf("running: %s", test.name)

		ctx := context.Background()

		expectedResponse := RatesResponse{
			Asks: []Rate{{Price: "1.234"}},
			Bids: []Rate{{Price: "1.123"}},
		}

		server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
			response, _ := json.Marshal(expectedResponse)
			rw.Write(response)
		}))
		defer server.Close()

		ratesRepo := &mocks.RatesRepo{}

		svc := NewRatesAPIService(context.Background(), &store.Store{Rates: ratesRepo}, server.Client())
		test.expectations(ratesRepo)

		r, err := svc.GetRates(ctx, "https://garantex.org/api/v2/depth?market=usdtusd")
		fmt.Print(r)
		if err != nil {
			if test.err != nil {
				assert.Equal(t, test.err.Error(), err.Error())
			} else {
				t.Errorf("Expected no error, found: %s", err.Error())
			}
		}
		ratesRepo.AssertExpectations(t)
	}
}
