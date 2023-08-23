package api

import (
	"context"
	"github.com/theghostmac/pricefetcher/internal/app"
	"github.com/theghostmac/pricefetcher/proto"
	"google.golang.org/grpc"
	"net"
)

type GRPCPriceFetcherServer struct {
	service app.PriceFetcher
	proto.UnimplementedPriceFetcherServer
}

func MakeAndRunGRPCServer(listenAddr string, service app.PriceFetcher) error {
	grpcPriceFetcher := NewGRPCPriceFetcher(service)

	listener, err := net.Listen("tcp", listenAddr)
	if err != nil {
		return err
	}

	options := []grpc.ServerOption{}
	server := grpc.NewServer(options...)
	proto.RegisterPriceFetcherServer(server, grpcPriceFetcher)
	err = server.Serve(listener)
	if err != nil {
		return err
	}

	return nil
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
