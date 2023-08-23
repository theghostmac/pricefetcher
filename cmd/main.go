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
	"os"
	"os/signal"
	"syscall"
)

func main() {
	service := common.NewLoggingService(observability.NewMetricsService(&app.PriceFetched{}))

	var (
		jsonListenAddr = flag.String("listenAddress", ":8080", "Listen address of the JSON transport")
		grpcListenAddr = flag.String("grpcListenAddr", ":4000", "Listen address of the GRPC transport")
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

	// Fetch the price before starting the shutdown process
	price, err := service.FetchPrice(context.Background(), "ETH")
	if err != nil {
		common.LogError(err)
	} else {
		fmt.Printf("%+v\n", price)
	}

	// Start the JSON server
	go jsonServer.Run()

	// Wait for the shutdown signal
	<-stopChan

	// Perform graceful shutdown
	jsonServer.Shutdown()
}
