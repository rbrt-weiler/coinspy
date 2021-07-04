package coingecko

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"strings"
	"sync"

	resty "github.com/go-resty/resty/v2"

	"gitlab.com/rbrt-weiler/coinspy/types"
)

const (
	// ProviderName cotains the common name of the provider.
	ProviderName string = "Coingecko"
	// APIBaseURL points to the basic API endpoint used for all requests.
	APIBaseURL string = "https://api.coingecko.com/api/v3"
)

// Coingecko is a specific implementation of a Provider.
type Coingecko struct {
	client             *resty.Client
	market             string
	providerWithMarket string
	coins              CoinList
	// Error is used to convey possible errors.
	Error error
}

// New initializes and returns a usable Provider object.
func New(c *resty.Client) (p Coingecko) {
	p.client = c
	p.market = "default"
	p.providerWithMarket = ProviderName
	p.Error = p.PopulateCoinList()
	return
}

// PopulateCoinList retrieves all coins known to Coingecko and stores that information for further use.
// This is required for Coingecko because common currency symbols need to be mapped to Coingecko-specific IDs for FetchRate().
func (p *Coingecko) PopulateCoinList() (err error) {
	var resp *resty.Response

	apiURL := fmt.Sprintf("%s/coins/list", APIBaseURL)
	resp, err = p.client.R().Get(apiURL)
	if err != nil {
		err = fmt.Errorf("could not fetch list of coins: %s", err)
		return
	}
	err = json.Unmarshal(resp.Body(), &p.coins)
	if err != nil {
		err = fmt.Errorf("could not unmarshal JSON: %s", err)
		return
	}

	return
}

// SymbolToID returns the Coingecko-specific ID for a common currency symbol.
func (p *Coingecko) SymbolToID(symbol string) (id string, err error) {
	if p.coins == nil {
		err = fmt.Errorf("list of coins not populated")
		return
	}

	lowerSymbol := strings.ToLower(symbol)
	for _, coin := range p.coins {
		if strings.ToLower(coin.Symbol) == lowerSymbol {
			id = coin.ID
		}
	}
	if id == "" {
		err = fmt.Errorf("symbol unknown to %s: %s", ProviderName, symbol)
	}

	return
}

// SetMarket does nothing, but needs to be implemented to satisfy the Provider interface.
func (p *Coingecko) SetMarket(market string) (err error) {
	return nil
}

// FetchRate returns a single ExchangeRate.
func (p *Coingecko) FetchRate(coin string, fiat string) (rate types.ExchangeRate, err error) {
	var coinParts []string
	var owned float64
	var coinID string
	var resp *resty.Response
	var priceList Prices

	apiURL := fmt.Sprintf("%s/simple/price", APIBaseURL)
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

	coinID, err = p.SymbolToID(coin)
	if err != nil {
		err = fmt.Errorf("could not find coin: %s", err)
		return
	}
	fiatID := strings.ToLower(fiat)

	resp, err = p.client.R().
		SetQueryParam("ids", coinID).
		SetQueryParam("vs_currencies", fiatID).
		Get(apiURL)
	if err != nil {
		err = fmt.Errorf("could not fetch rate for %s/%s: %s", coinID, fiatID, err)
		return
	}

	err = json.Unmarshal(resp.Body(), &priceList)
	if err != nil {
		err = fmt.Errorf("could not unmarshal API response: %s", err)
		return
	}

	return types.ExchangeRate{Provider: ProviderName, Market: p.market, ProviderWithMarket: p.providerWithMarket, AsOf: resp.ReceivedAt(), Coin: coin, Owned: owned, Fiat: fiat, Rate: priceList[coinID][fiatID], Error: err}, err
}

// FetchRateSynced is a multi-threading implementation of FetchRate.
func (p *Coingecko) FetchRateSynced(coin string, fiat string, rates *types.ExchangeRates, wg *sync.WaitGroup) {
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
