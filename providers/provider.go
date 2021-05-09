package providers

import (
	"sync"

	"gitlab.com/rbrt-weiler/coinspy/providers/coingate"
	"gitlab.com/rbrt-weiler/coinspy/providers/coingecko"
	"gitlab.com/rbrt-weiler/coinspy/providers/cryptowatch"
	"gitlab.com/rbrt-weiler/coinspy/types"
)

type Provider interface {
	SetMarket(market string) (err error)
	FetchRate(coin string, fiat string) (rate types.ExchangeRate, err error)
	FetchRateSynced(coin string, fiat string, rates *types.ExchangeRates, wg *sync.WaitGroup)
}

func CoinGate() (p *coingate.CoinGate) {
	provider := coingate.New()
	return &provider
}

func Coingecko() (p *coingecko.Coingecko) {
	provider := coingecko.New()
	return &provider
}

func Cryptowatch() (p *cryptowatch.Cryptowatch) {
	provider := cryptowatch.New()
	return &provider
}
