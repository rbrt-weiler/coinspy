package livecoinwatch

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"strings"
	"sync"

	resty "github.com/go-resty/resty/v2"
	"gitlab.com/rbrt-weiler/coinspy/core"
	"gitlab.com/rbrt-weiler/coinspy/types"
)

const (
	// ProviderName cotains the common name of the provider.
	ProviderName string = "LiveCoinWatch"
	// APIBaseURL points to the basic API endpoint used for all requests.
	APIBaseURL string = "https://api.livecoinwatch.com"
)

// LiveCoinWatch is a specific implementation of a Provider.
type LiveCoinWatch struct {
	client             *resty.Client
	market             string
	providerWithMarket string
	// Error is used to convey possible errors.
	Error error
}

// New initializes and returns a usable Provider object.
func New(c *resty.Client) (p LiveCoinWatch) {
	p.client = c
	p.market = "default"
	p.providerWithMarket = ProviderName
	return p
}

// SetMarket does nothing, but needs to be implemented to satisfy the Provider interface.
func (p *LiveCoinWatch) SetMarket(market string) (err error) {
	return nil
}

// FetchRate returns a single ExchangeRate.
func (p *LiveCoinWatch) FetchRate(coin string, fiat string) (rate types.ExchangeRate, err error) {
	var coinParts []string
	var owned float64
	var resp *resty.Response
	var query PriceQuery
	var queryJSON []byte
	var price PriceResult

	apiURL := fmt.Sprintf("%s/coins/single", APIBaseURL)
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

	query.Code = strings.ToUpper(coin)
	query.Currency = strings.ToUpper(fiat)
	query.Meta = false

	queryJSON, err = json.Marshal(query)
	if err != nil {
		err = fmt.Errorf("could not encode API query: %s", err)
		return
	}

	resp, err = p.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("X-API-Key", core.Config.LiveCoinWatch.APIKey).
		SetBody(string(queryJSON)).
		Post(apiURL)
	if err != nil {
		err = fmt.Errorf("could not fetch rate for %s/%s: %s", coin, fiat, err)
		return
	}

	err = json.Unmarshal(resp.Body(), &price)
	if err != nil {
		err = fmt.Errorf("could not unmarshal API response: %s", err)
		return
	}
	if price.Error.Code != 0 {
		err = fmt.Errorf("could not retrieve exchange rate for %s/%s: %s: %s", coin, fiat, price.Error.Status, price.Error.Description)
		return
	}

	return types.ExchangeRate{Provider: ProviderName, Market: p.market, ProviderWithMarket: p.providerWithMarket, AsOf: resp.ReceivedAt(), Coin: coin, Owned: owned, Fiat: fiat, Rate: price.Rate, Error: err}, err
}

// FetchRateSynced is a multi-threading implementation of FetchRate.
func (p *LiveCoinWatch) FetchRateSynced(coin string, fiat string, rates *types.ExchangeRates, wg *sync.WaitGroup) {
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
