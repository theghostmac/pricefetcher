package observability

import (
	"context"
	"fmt"
	"github.com/theghostmac/pricefetcher/internal/app"
)

type MetricService struct {
	next app.PriceFetcher
}

func (ms *MetricService) FetchPrice(ctx context.Context, ticker string) (price float64, err error) {
	fmt.Println("Pushing metrics to Prometheus")
	return ms.next.FetchPrice(ctx, ticker)
}

func NewMetricsService(next app.PriceFetcher) app.PriceFetcher {
	return &MetricService{
		next: next,
	}
}
