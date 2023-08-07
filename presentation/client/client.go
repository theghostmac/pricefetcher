package client

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/theghostmac/pricefetcher/internal/domain"
	"net/http"
)

type Client struct {
	endpoint string
}

func NewClient(endpoint string) *Client {
	return &Client{
		endpoint: endpoint,
	}
}

func (c *Client) FetchPrice(ctx context.Context, ticker string) (*domain.PriceResponse, error) {
	endpoint := fmt.Sprintf("%s?ticker=%s", c.endpoint, ticker)

	request, err := http.NewRequest("get", endpoint, nil)
	if err != nil {
		return nil, err
	}
	response, err := http.DefaultClient.Do(request)
	if err != nil {
		return nil, err
	}

	if response.StatusCode != http.StatusOK {
		httpError := map[string]any{}
		if err := json.NewDecoder(response.Body).Decode(&httpError); err != nil {
			return nil, err
		}
		return nil, fmt.Errorf("service responded with non OK status code: %s", httpError["error"])
	}

	priceResponse := new(domain.PriceResponse)
	if err := json.NewDecoder(response.Body).Decode(priceResponse); err != nil {
		return nil, err
	}

	return priceResponse, nil
}
