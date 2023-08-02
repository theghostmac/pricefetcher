package client

import (
	"context"
	"github.com/theghostmac/pricefetcher/internal/domain"
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
	return nil, nil
}
