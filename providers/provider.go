package providers

import (
	"sync"

	"github.com/go-resty/resty/v2"
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

func CoinGate(c *resty.Client) (p *coingate.CoinGate) {
	provider := coingate.New(c)
	return &provider
}

func Coingecko(c *resty.Client) (p *coingecko.Coingecko) {
	provider := coingecko.New(c)
	return &provider
}

func Cryptowatch(c *resty.Client) (p *cryptowatch.Cryptowatch) {
	provider := cryptowatch.New(c)
	return &provider
}
