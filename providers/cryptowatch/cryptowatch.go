package cryptowatch

import (
	"encoding/json"
	"fmt"
	"os"
	"sort"
	"strings"
	"sync"

	resty "github.com/go-resty/resty/v2"

	"gitlab.com/rbrt-weiler/coinspy/types"
)

const (
	ProviderName string = "Cryptowatch"
	APIBaseURL   string = "https://api.cryptowat.ch"
)

type Cryptowatch struct {
	client             *resty.Client
	market             string
	providerWithMarket string
	Error              error
}

func uniqueStrings(input []string) (output []string) {
	// thanks to https://kylewbanks.com/blog/creating-unique-slices-in-go for this
	m := make(map[string]bool)

	for _, val := range input {
		if _, ok := m[val]; !ok {
			m[val] = true
			output = append(output, val)
		}
	}

	return
}

func New(c *resty.Client) (p Cryptowatch) {
	p.client = c
	p.Error = p.SetMarket("Kraken")
	return p
}

func ListMarkets() (markets []string, err error) {
	var apiURL string
	var resp *resty.Response
	var apiResult Markets

	client := resty.New()

	apiURL = fmt.Sprintf("%s/markets", APIBaseURL)
	resp, err = client.R().Get(apiURL)
	if err != nil {
		err = fmt.Errorf("could not fetch list of markets: %s", err)
	} else {
		err = json.Unmarshal(resp.Body(), &apiResult)
		if err != nil {
			err = fmt.Errorf("could not unmarshal JSON: %s", err)
			return
		}
	}

	for _, market := range apiResult.Result {
		markets = append(markets, market.Exchange)
	}
	markets = uniqueStrings(markets)
	sort.Strings(markets)

	return
}

func (p *Cryptowatch) SetMarket(market string) (err error) {
	p.market = market
	p.providerWithMarket = fmt.Sprintf("%s/%s", ProviderName, p.market)
	return nil
}

func (p *Cryptowatch) FetchRate(coin string, fiat string) (rate types.ExchangeRate, err error) {
	var apiURL string
	var resp *resty.Response
	var apiResult Result
	var rateValue float64

	rateValue = 0
	apiURL = fmt.Sprintf("https://api.cryptowat.ch/markets/%s/%s%s/price", strings.ToLower(p.market), strings.ToLower(coin), strings.ToLower(fiat))
	resp, err = p.client.R().Get(apiURL)
	if err != nil {
		err = fmt.Errorf("could not fetch exchange rate: %s", err)
	} else {
		err = json.Unmarshal(resp.Body(), &apiResult)
		if err != nil {
			err = fmt.Errorf("could not unmarshal JSON: %s", err)
			return
		}
		if apiResult.Error != "" {
			err = fmt.Errorf("%s (%s/%s on %s; %f allowance remaining)", apiResult.Error, coin, fiat, p.providerWithMarket, apiResult.Allowance.Remaining)
		} else {
			rateValue = apiResult.Result.Price
		}
	}

	return types.ExchangeRate{Provider: ProviderName, Market: p.market, ProviderWithMarket: p.providerWithMarket, AsOf: resp.ReceivedAt(), Coin: coin, Fiat: fiat, Rate: rateValue, Error: err}, err
}

func (p *Cryptowatch) FetchRateSynced(coin string, fiat string, rates *types.ExchangeRates, wg *sync.WaitGroup) {
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
