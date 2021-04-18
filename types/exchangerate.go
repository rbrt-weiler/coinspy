package types

import "time"

type ExchangeRate struct {
	Provider string
	Market   string
	AsOf     time.Time
	Coin     string
	Fiat     string
	Rate     float64
	Error    error
}
