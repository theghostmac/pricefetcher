package unit_test

import (
	"context"
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"github.com/theghostmac/pricefetcher/internal/api"
	"github.com/theghostmac/pricefetcher/internal/domain"
	"net/http"
	"net/http/httptest"
	"testing"
)

type mockPriceFetcher struct {
}

func (m *mockPriceFetcher) FetchPrice(ctx context.Context, ticker string) (float64, error) {
	return 1875.91, nil
}

func TestJSONAPIServer_HandleFetchPrice(t *testing.T) {
	mockService := &mockPriceFetcher{}
	jsonServer := &api.JSONAPIServer{Service: mockService}

	request := httptest.NewRequest(http.MethodGet, "/?ticker=ETH", nil)

	recorder := httptest.NewRecorder()

	err := jsonServer.HandleFetchPrice(context.Background(), recorder, request)
	assert.NoError(t, err)

	assert.Equal(t, http.StatusOK, recorder.Code)

	var response domain.PriceResponse
	err = json.NewDecoder(recorder.Body).Decode(&response)
	assert.NoError(t, err)

	assert.Equal(t, "ETH", response.Ticker)
	assert.Equal(t, 1875.91, response.Price)
}
