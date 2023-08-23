package app

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
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
type PriceFetched struct{}

func (pf PriceFetched) FetchPrice(ctx context.Context, ticker string) (float64, error) {
	return SimulatedPriceFetcher(ctx, ticker)
}

func SimulatedPriceFetcher(ctx context.Context, ticker string) (float64, error) {
	// Mimic the HTTP round-trip
	time.Sleep(100 * time.Millisecond)
	price, ok := SimulatedPrice[ticker]
	if !ok {
		return price, fmt.Errorf("the given ticker (%s) is not supported", ticker)
	}
	return price, nil
}

const binanceAPIBaseURl = "https://api.binance.com/api/v3/ticker/price"

type FetchPriceFromBinance struct {
	APIKey string // Binance API key here.
}

func NewFetchPriceFromBinance(apiKey string) *FetchPriceFromBinance {
	return &FetchPriceFromBinance{
		APIKey: apiKey,
	}
}

func (fpfb *FetchPriceFromBinance) FetchPrice(ctx context.Context, ticker string) (float64, error) {
	return fpfb.BinancePriceFetcher(ctx, ticker)
}

type BinancePriceResponse struct {
	Symbol string `json:symbol`
	Price  string `json:price`
}

func (fpfb *FetchPriceFromBinance) BinancePriceFetcher(ctx context.Context, ticker string) (float64, error) {
	client := &http.Client{}
	request, err := http.NewRequest("GET", fmt.Sprintf("%s?symbol=%sUSDT", binanceAPIBaseURl, ticker), nil)
	if err != nil {
		return 0, err
	}

	response, err := client.Do(request)
	if err != nil {
		return 0, err
	}
	defer response.Body.Close()

	var priceResponse BinancePriceResponse
	err = json.NewDecoder(response.Body).Decode(&priceResponse)
	if err != nil {
		return 0, err
	}

	price, err := ParseFloat(priceResponse.Price)
	if err != nil {
		return 0, err
	}

	return price, nil
}

func ParseFloat(input string) (float64, error) {
	return strconv.ParseFloat(input, 64)
}
