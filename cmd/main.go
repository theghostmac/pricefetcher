package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/theghostmac/pricefetcher/common"
	"github.com/theghostmac/pricefetcher/internal/api"
	"github.com/theghostmac/pricefetcher/internal/app"
	"github.com/theghostmac/pricefetcher/internal/observability"
	"github.com/theghostmac/pricefetcher/internal/server"
	"github.com/theghostmac/pricefetcher/presentation/client"
	"github.com/theghostmac/pricefetcher/proto"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	//service := common.NewLoggingService(observability.NewMetricsService(&app.PriceFetched{}))

	var (
		jsonListenAddr = flag.String("listenAddress", ":8080", "Listen address of the JSON transport")
		grpcListenAddr = flag.String("grpcListenAddr", ":4000", "Listen address of the GRPC transport")
		//ctx            = context.Background()
	)
	flag.Parse()

	// Create an instance of PriceFetched as the mock PriceFetcher implementation.
	priceFetcher := &app.PriceFetched{}

	// Create a new MetricsService wrapping the PriceFetcher.
	metricsService := observability.NewMetricsService(priceFetcher)

	// Create an instance of JSONAPIServer with the desired ListenAddr and service.
	jsonServer := &api.JSONAPIServer{
		StartRunner: server.StartRunner{
			ListenAddr: *jsonListenAddr,
		},
		Service: metricsService,
	}

	// Create an instance of the GRPCServer with the desired ListenAddr and service.
	go func() {
		err := api.MakeAndRunGRPCServer(*grpcListenAddr, priceFetcher)
		if err != nil {
			common.LogError(err)
		}
	}()

	// Setup graceful shutdown using SIGINT (Ctrl+C) and SIGTERM signals.
	stopChan := make(chan os.Signal, 1)
	signal.Notify(stopChan, syscall.SIGINT, syscall.SIGTERM)

	// Create a new client instance for fetching price.
	grpcClient, err := client.NewGRPCClient(*grpcListenAddr)
	if err != nil {
		common.LogError(err)
	}

	// Fetch the price using the GRPC client.
	go func() {
		for {
			time.Sleep(3 * time.Second)
			grpcPriceResponse, err := grpcClient.FetchPrice(context.Background(), &proto.PriceRequest{
				Ticker: "ETH",
			})
			if err != nil {
				common.LogError(err)
			} else {
				fmt.Printf("GRPC Price Response: %+v\n", grpcPriceResponse)
			}
		}
	}()

	// Start the JSON server
	go jsonServer.Run()

	// Wait for the shutdown signal
	<-stopChan

	// Perform graceful shutdown
	jsonServer.Shutdown()
}
