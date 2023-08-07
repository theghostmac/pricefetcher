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
	"os"
	"os/signal"
	"syscall"
)

func main() {
	// Create a new client to receive data from the server.
	newClient := client.NewClient("http://localhost:8080")
	//service := common.NewLoggingService(observability.NewMetricsService(&app.PriceFetched{}))

	listenAddr := flag.String("listenAddress", ":8080", "listening on the default port")
	flag.Parse()

	// Create an instance of PriceFetched as the mock PriceFetcher implementation.
	priceFetcher := &app.PriceFetched{}

	// Create a new LoggingService wrapping the PriceFetcher.
	loggingService := common.NewLoggingService(priceFetcher)

	// Create a new MetricsService wrapping the LoggingService.
	metricsService := observability.NewMetricsService(loggingService)

	// Create an instance of JSONAPIServer with the desired ListenAddr and service.
	apiServer := &api.JSONAPIServer{
		StartRunner: server.StartRunner{
			ListenAddr: *listenAddr, // Change this to the address where you want your server to listen.
		},
		Service: metricsService,
	}

	// Call the Run method to start the server.
	go apiServer.Run()

	// Setup graceful shutdown using SIGINT (Ctrl+C) and SIGTERM signals.
	stopChan := make(chan os.Signal, 1)
	signal.Notify(stopChan, syscall.SIGINT, syscall.SIGTERM)
	<-stopChan

	// Perform graceful shutdown
	apiServer.Shutdown()

	price, err := newClient.FetchPrice(context.Background(), "ETH")
	if err != nil {
		common.LogError(err)
	}
	fmt.Printf("%+v\n", price)
}
