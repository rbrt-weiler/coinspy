package providers

import (
	"sync"

	"github.com/go-resty/resty/v2"
	"gitlab.com/rbrt-weiler/coinspy/providers/bitpanda"
	"gitlab.com/rbrt-weiler/coinspy/providers/coingate"
	"gitlab.com/rbrt-weiler/coinspy/providers/coingecko"
	"gitlab.com/rbrt-weiler/coinspy/providers/livecoinwatch"
	"gitlab.com/rbrt-weiler/coinspy/types"
)

// Provider is the common interface used for every specific provider.
type Provider interface {
	SetMarket(market string) (err error)
	FetchRate(coin string, fiat string) (rate types.ExchangeRate, err error)
	FetchRateSynced(coin string, fiat string, rates *types.ExchangeRates, wg *sync.WaitGroup)
}

// Bitpanda returns an initialized provider implementation.
func Bitpanda(c *resty.Client) (p *bitpanda.Bitpanda) {
	provider := bitpanda.New(c)
	return &provider
}

// CoinGate returns an initialized provider implementation.
func CoinGate(c *resty.Client) (p *coingate.CoinGate) {
	provider := coingate.New(c)
	return &provider
}

// Coingecko returns an initialized provider implementation.
func Coingecko(c *resty.Client) (p *coingecko.Coingecko) {
	provider := coingecko.New(c)
	return &provider
}

// LiveCoinWatch returns an initialized provider implementation.
func LiveCoinWatch(c *resty.Client) (p *livecoinwatch.LiveCoinWatch) {
	provider := livecoinwatch.New(c)
	return &provider
}
