package api

import (
	"context"
	"encoding/json"
	"github.com/theghostmac/pricefetcher/internal/app"
	"github.com/theghostmac/pricefetcher/internal/server"
	"math/rand"
	"net/http"
)

type APIFunc func(ctx context.Context, writer http.ResponseWriter, request *http.Request) error

type JSONAPIServer struct {
	server.StartRunner
	service app.PriceFetcher
}

type PriceResponse struct {
	Ticker string  `json:"ticker"`
	Price  float64 `json:"price"`
}

func (js *JSONAPIServer) NewJSONAPIServer( /*listenAddr server.StartRunner,*/ service app.PriceFetcher) *JSONAPIServer {
	return &JSONAPIServer{
		StartRunner: server.StartRunner{},
		service:     service,
	}
}

func (js *JSONAPIServer) Run() {
	http.HandleFunc("/", MakeHTTPHandlerFunc(js.HandleFetchPrice))
}

func MakeHTTPHandlerFunc(apiFn APIFunc) http.HandlerFunc {
	ctxBg := context.Background()
	ctx := context.WithValue(ctxBg, "requestID", rand.Intn(100000000))

	return func(writer http.ResponseWriter, request *http.Request) {
		if err := apiFn(ctx, writer, request); err != nil {
			WriteJSON(writer, http.StatusBadRequest, map[string]any{"error": err.Error()})
		}
	}
}

func (js *JSONAPIServer) HandleFetchPrice(ctx context.Context, writer http.ResponseWriter, request *http.Request) error {
	ticker := request.URL.Query().Get("ticker")

	price, err := js.service.FetchPrice(ctx, ticker)
	if err != nil {
		return err
	}

	priceResponse := PriceResponse{
		Price:  price,
		Ticker: ticker,
	}
	return WriteJSON(writer, http.StatusOK, &priceResponse)
}

func WriteJSON(writer http.ResponseWriter, status int, value any) error {
	writer.WriteHeader(status)
	return json.NewEncoder(writer).Encode(value)
}
