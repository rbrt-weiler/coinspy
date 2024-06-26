package bitpanda

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"strings"
	"sync"
	"time"

	resty "github.com/go-resty/resty/v2"

	"gitlab.com/rbrt-weiler/coinspy/types"
)

const (
	// ProviderName cotains the common name of the provider.
	ProviderName string = "Bitpanda"
	// APIBaseURL points to the basic API endpoint used for all requests.
	APIBaseURL string = "https://api.bitpanda.com/v1"
)

// Bitpanda is a specific implementation of a Provider.
type Bitpanda struct {
	client             *resty.Client
	market             string
	providerWithMarket string
	prices             Prices
	pricesAsOf         time.Time
	// Error is used to convey possible errors.
	Error error
}

// New initializes and returns a usable Provider object.
func New(c *resty.Client) (p Bitpanda) {
	p.client = c
	p.market = "default"
	p.providerWithMarket = ProviderName
	p.Error = p.FetchPrices()
	return
}

// FetchPrices retrieves all available exchange rates and stores them for further use.
func (p *Bitpanda) FetchPrices() (err error) {
	var resp *resty.Response
	var priceList Prices

	apiURL := fmt.Sprintf("%s/ticker", APIBaseURL)
	resp, err = p.client.R().Get(apiURL)
	if err != nil {
		err = fmt.Errorf("prices could not be fetched: %s", err)
		return
	}

	err = json.Unmarshal(resp.Body(), &priceList)
	if err != nil {
		err = fmt.Errorf("prices could not be parsed: %s", err)
		return
	}

	p.prices = priceList
	p.pricesAsOf = resp.ReceivedAt()

	return
}

// SetMarket does nothing, but needs to be implemented to satisfy the Provider interface.
func (p *Bitpanda) SetMarket(market string) (err error) {
	return nil
}

// FetchRate returns a single ExchangeRate.
func (p *Bitpanda) FetchRate(coin string, fiat string) (rate types.ExchangeRate, err error) {
	var coinParts []string
	var owned float64
	var rateString string
	var ok bool
	var rateFloat float64

	if p.prices == nil {
		err = p.FetchPrices()
		if err != nil {
			err = fmt.Errorf("prices could not be initialized: %s", err)
			return
		}
	}

	coinParts = strings.Split(coin, "=")
	coin = coinParts[0]
	owned = -1
	if len(coinParts) == 2 {
		owned, err = strconv.ParseFloat(coinParts[1], 64)
		if err != nil {
			err = fmt.Errorf("could not parse owned coins for %s/%s: %s", coin, fiat, err)
			return
		}
	}

	fiatID := strings.ToUpper(fiat)
	coinID := strings.ToUpper(coin)
	if rateString, ok = p.prices[coinID][fiatID]; !ok {
		err = fmt.Errorf("rate %s/%s is unknown to %s", coinID, fiatID, p.providerWithMarket)
		return
	}
	rateFloat, err = strconv.ParseFloat(rateString, 64)
	if err != nil {
		err = fmt.Errorf("price could not be parsed to a float: %s", err)
		return
	}

	return types.ExchangeRate{Provider: ProviderName, Market: p.market, ProviderWithMarket: p.providerWithMarket, AsOf: p.pricesAsOf, Coin: coin, Owned: owned, Fiat: fiat, Rate: rateFloat, Error: err}, err
}

// FetchRateSynced is a multi-threading implementation of FetchRate.
func (p *Bitpanda) FetchRateSynced(coin string, fiat string, rates *types.ExchangeRates, wg *sync.WaitGroup) {
	defer wg.Done()
	rate, err := p.FetchRate(coin, fiat)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %s\n", err)
		return
	}
	rates.Mutex.Lock()
	rates.Rates = append(rates.Rates, rate)
	rates.Mutex.Unlock()
}
