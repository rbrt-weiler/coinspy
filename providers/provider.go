package providers

import (
	"sync"

	"gitlab.com/rbrt-weiler/coinspy/providers/coingecko"
	"gitlab.com/rbrt-weiler/coinspy/providers/cryptowatch"
	"gitlab.com/rbrt-weiler/coinspy/types"
)

type Provider interface {
	FetchRate(market string, coin string, fiat string) (types.ExchangeRate, error)
	FetchRateSynced(rates *types.ExchangeRates, market string, coin string, fiat string, wg *sync.WaitGroup)
}

func Coingecko() (p Provider) {
	p = coingecko.New()
	return
}

func Cryptowatch() (p Provider) {
	p = cryptowatch.New()
	return
}
