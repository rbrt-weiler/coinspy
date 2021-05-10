package types

import (
	"fmt"
	"time"
)

type ExchangeRate struct {
	Provider           string
	Market             string
	ProviderWithMarket string
	AsOf               time.Time
	Coin               string
	Fiat               string
	Rate               float64
	Error              error
}

func (r *ExchangeRate) String() string {
	return fmt.Sprintf("1 %s = %f %s (on %s as of %s)", r.Coin, r.Rate, r.Fiat, r.ProviderWithMarket, r.Market, r.AsOf.Format(time.RFC3339))
}

func (r *ExchangeRate) StringCompact() string {
	return fmt.Sprintf("1 %s = %f %s (on %s)", r.Coin, r.Rate, r.Fiat, r.ProviderWithMarket)
}

func (r *ExchangeRate) StringVeryCompact() string {
	return fmt.Sprintf("1 %s = %f %s", r.Coin, r.Rate, r.Fiat)
}
