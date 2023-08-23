package main

import (
	"context"
	"fmt"
	"github.com/theghostmac/pricefetcher/common"
	"github.com/theghostmac/pricefetcher/proto"
	"google.golang.org/grpc"
)

func main() {
	// Set up a connection to the GRPC server.
	conn, err := grpc.Dial("localhost:4000", grpc.WithInsecure())
	if err != nil {
		common.LogError(err)
	}
	defer conn.Close()

	// Create a client instance using the connection.
	client := proto.NewPriceFetcherClient(conn)

	// Define the request.
	request := &proto.PriceRequest{
		Ticker: "ETH",
	}

	// Call the FetchPrice method on the client.
	response, err := client.FetchPrice(context.Background(), request)
	if err != nil {
		common.LogError(err)
	}

	// Print the response.
	fmt.Printf("Price for %s: %f\n", response.Ticker, response.Price)
}
