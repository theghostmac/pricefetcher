package app

import (
	"context"
	"fmt"
	"time"
)

// PriceFetcher is an interface for functions related to FetchPrice service.
type PriceFetcher interface {
	FetchPrice(context.Context, string) (float64, error)
}

var SimulatedPrice = map[string]float64{
	"BTC": 29_469.55,
	"ETH": 1_875.91,
	"XRP": 0.7162,
}

// PriceFetched is a struct containing the fetched prices from PriceFetcher services.
type PriceFetched struct {
}

func (pf PriceFetched) FetchPrice(ctx context.Context, ticker string) (float64, error) {
	return SimulatedPriceFetcher(ctx, ticker)
}

func SimulatedPriceFetcher(ctx context.Context, ticker string) (float64, error) {
	// Mimic the HTTP roundtrip
	time.Sleep(100 * time.Millisecond)
	price, ok := SimulatedPrice[ticker]
	if !ok {
		return price, fmt.Errorf("the given ticker (%s) is not supported", ticker)
	}
	return price, nil
}
