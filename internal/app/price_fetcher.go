package app

import (
	"context"
	"fmt"
)

type PriceFetcher interface {
	FetchPrice(context.Context, string) (float64, error)
}

var SimulatedPrice = map[string]float64{
	"BTC": 29_469.55,
	"ETH": 1_875.91,
	"XRP": 0.7162,
}

type PriceFetched struct {
}

func (pf PriceFetched) FetchPrice(ctx context.Context, ticker string) (float64, error) {
	return 0, nil
}

func SimulatedPriceFetcher(ctx context.Context, ticker string) (float64, error) {
	price, ok := SimulatedPrice[ticker]
	if !ok {
		return price, fmt.Errorf("the given ticker (%s) is not supported", ticker)
	}
	return price, nil
}
