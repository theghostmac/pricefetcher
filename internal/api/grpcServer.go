package api

import (
	"context"
	"github.com/theghostmac/pricefetcher/internal/app"
	"github.com/theghostmac/pricefetcher/proto"
)

type GRPCPriceFetcherServer struct {
	service app.PriceFetcher
}

func NewGRPCPriceFetcher(service app.PriceFetcher) *GRPCPriceFetcherServer {
	return &GRPCPriceFetcherServer{
		service: service,
	}
}

func (grpcs *GRPCPriceFetcherServer) FetchPrice(ctx context.Context, request *proto.PriceRequest) (*proto.PriceResponse, error) {
	price, err := grpcs.service.FetchPrice(ctx, request.Ticker)
	if err != nil {
		return nil, err
	}

	response := &proto.PriceResponse{
		Ticker: request.Ticker,
		Price:  float32(price),
	}
	return response, err
}
