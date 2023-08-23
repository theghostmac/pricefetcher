# Price Fetcher Microservice
Pricefetcher is a microservice application that fetches the price of different cryptocurrencies in real-time. 
My main goal is to just use gRPC in communication.

## Features
I'm trying out the following technologies
- JSON API
- GRPC
- Context

Crypto prices for 27th July 2023 used.

## Installation

1. Clone the repository: `git clone https://github.com/theghostmac/pricefetcher.git`
2. Change directory: `cd pricefetcher`
3. Install dependencies: `go mod download`

## Usage
Using Golang
1. Build the application: `go build -o myapp cmd/main.go`
2. Run the application: `./myapp`

Using Makefile to run the application:
```shell
make run
```
You can fetch any price among BTC or ETH like this:
```shell
curl -X GET "http://localhost:8080/?ticker=ETH"
{"ticker":"ETH","price":1875.91}
```
![Works](works.png)

Using Makefile to build the application:
```shell
make build
```
Clean up the binary after you finish:
```shell
make clean
```

While this uses simulated price, I have written the logic to fetch real-time price from
CoinMarketCap, however, the goal of the project is not to do that, but to master the
usage of gRPC as a client and server.

## Configuration

The project supports configuration via environment variables. You can customize the behavior by setting the following environment variables:

- `MYAPP_PORT`: Set the port number on which the application listens. Default is 8080.
- `MYAPP_DEBUG`: Set to "true" to enable debug mode. Default is "false".