package api

import "github.com/theghostmac/pricefetcher/internal/app"

type GRPCPriceFetcher struct {
	service app.PriceFetcher
}

func NewGRPCPriceFetcher(service app.PriceFetcher) *GRPCPriceFetcher {
	return &GRPCPriceFetcher{
		service: service,
	}
}

func (pr *GRPCPriceFetcher) Name() *Gr
