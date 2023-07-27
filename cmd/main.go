package main

import (
	"context"
	"fmt"
	"github.com/theghostmac/pricefetcher/common"
	"github.com/theghostmac/pricefetcher/internal/app"
	"github.com/theghostmac/pricefetcher/internal/observability"
	"github.com/theghostmac/pricefetcher/internal/server"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	// Create an instance of StartRunner with the desired ListenAddr.
	runner := &server.StartRunner{
		ListenAddr: "localhost:8080", // Change this to the address where you want your server to listen.
	}

	// Call the Run method to start the server.
	if err := runner.Run(); err != nil {
		panic(err)
	}

	// Setup graceful shutdown using SIGINT (Ctrl+C) and SIGTERM signals.
	stopChan := make(chan os.Signal, 1)
	signal.Notify(stopChan, syscall.SIGINT, syscall.SIGTERM)
	<-stopChan

	// Perform graceful shutdown
	runner.Shutdown()

	service := common.NewLoggingService(observability.NewMetricsService(&app.PriceFetched{}))

	price, err := service.FetchPrice(context.Background(), "ETH")
	if err != nil {
		common.LogError(err)
	}
	fmt.Println(price)
}

// main.go checked
//
