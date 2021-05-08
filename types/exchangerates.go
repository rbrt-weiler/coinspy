package types

import (
	"fmt"
	"sort"
	"sync"
)

type ExchangeRates struct {
	Mutex sync.Mutex
	Rates []ExchangeRate
}

func (r *ExchangeRates) Sort() {
	r.Mutex.Lock()
	sort.Slice(r.Rates, func(i, j int) bool {
		left := fmt.Sprintf("%s-%s-%s-%s", r.Rates[i].Coin, r.Rates[i].Fiat, r.Rates[i].Provider, r.Rates[i].Market)
		right := fmt.Sprintf("%s-%s-%s-%s", r.Rates[j].Coin, r.Rates[j].Fiat, r.Rates[j].Provider, r.Rates[j].Market)
		return left < right
	})
	r.Mutex.Unlock()
}
