package common

import "github.com/theghostmac/pricefetcher/internal/app"

type LoggingService struct {
	next app.PriceFetcher
}

func (ls *LoggingService) Fetch() {

}
