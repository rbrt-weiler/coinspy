package types

import (
	"fmt"
	"time"
)

type ExchangeRate struct {
	Provider string
	Market   string
	AsOf     time.Time
	Coin     string
	Fiat     string
	Rate     float64
	Error    error
}

func (r *ExchangeRate) String() string {
	return fmt.Sprintf("1 %s = %f %s (on %s/%s as of %s)", r.Coin, r.Rate, r.Fiat, r.Provider, r.Market, r.AsOf.Format(time.RFC3339))
}

func (r *ExchangeRate) StringCompact() string {
	return fmt.Sprintf("1 %s = %f %s (on %s/%s)", r.Coin, r.Rate, r.Fiat, r.Provider, r.Market)
}

func (r *ExchangeRate) StringVeryCompact() string {
	return fmt.Sprintf("1 %s = %f %s", r.Coin, r.Rate, r.Fiat)
}
