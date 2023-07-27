package common

import (
	"context"
	"github.com/sirupsen/logrus"
	"github.com/theghostmac/pricefetcher/internal/app"
	"time"
)

type LoggingService struct {
	next app.PriceFetcher
}

// NewLoggingService is a new logging service for PriceFetcher.
func NewLoggingService(srv app.PriceFetcher) app.PriceFetcher {
	return &LoggingService{
		next: srv,
	}
}

func (ls *LoggingService) FetchPrice(ctx context.Context, ticker string) (price float64, err error) {
	defer func(begin time.Time) {
		logrus.WithFields(logrus.Fields{
			"requestID": ctx.Value("requestID"),
			"took":      time.Since(begin),
			"err":       err,
			"price":     price,
		}).Info("fetchPrice")
	}(time.Now())

	return ls.next.FetchPrice(ctx, ticker)
}

func LogError(err error) {
	logrus.Errorf("Error: %v", err)
}
