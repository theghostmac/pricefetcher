package app

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/theghostmac/pricefetcher/common"
	"io"
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

const coinMarketCapAPIBaseURL = "https://pro-api.coinmarketcap.com/v1/cryptocurrency/listings/latest"

type FetchPriceFromCoinMarketCap struct {
	APIKey string // CoinMarketCap API key here.
}

func NewFetchPriceFromCoinMarketCap(apiKey string) *FetchPriceFromCoinMarketCap {
	return &FetchPriceFromCoinMarketCap{
		APIKey: apiKey,
	}
}

func (fpfcmc *FetchPriceFromCoinMarketCap) FetchPrice(ctx context.Context, ticker string) (float64, error) {
	return fpfcmc.CoinMarketCapPriceFetcher(ctx, ticker)
}

type CoinMarketCapPriceResponse struct {
	Data []struct {
		Symbol string `json:"symbol"`
		Quote  struct {
			USD struct {
				Price string `json:"price"`
			} `json:"USD"`
		} `json:"quote"`
	} `json:"data"`
}

func (fpfcmc *FetchPriceFromCoinMarketCap) CoinMarketCapPriceFetcher(ctx context.Context, ticker string) (float64, error) {
	client := &http.Client{}
	req, err := http.NewRequest("GET", coinMarketCapAPIBaseURL, nil)
	if err != nil {
		return 0, err
	}
	q := req.URL.Query()
	q.Add("start", "1")
	q.Add("limit", "100")
	q.Add("convert", "USD")
	req.URL.RawQuery = q.Encode()

	req.Header.Add("X-CMC_PRO_API_KEY", fpfcmc.APIKey)

	resp, err := client.Do(req)
	if err != nil {
		return 0, err
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			common.LogError(err)
		}
	}(resp.Body)

	var priceResponse CoinMarketCapPriceResponse
	err = json.NewDecoder(resp.Body).Decode(&priceResponse)
	if err != nil {
		return 0, err
	}

	for _, coin := range priceResponse.Data {
		if coin.Symbol == ticker {
			price, err := ParseFloat(coin.Quote.USD.Price)
			if err != nil {
				return 0, err
			}
			return price, nil
		}
	}

	return 0, fmt.Errorf("ticker (%s) not found in CoinMarketCap response", ticker)
}

func ParseFloat(input string) (float64, error) {
	return strconv.ParseFloat(input, 64)
}
