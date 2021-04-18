package types

import "sync"

type ExchangeRates struct {
	Mutex sync.Mutex
	Rates []ExchangeRate
}
