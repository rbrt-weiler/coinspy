package cryptowatch

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
	"sync"

	resty "github.com/go-resty/resty/v2"

	"gitlab.com/rbrt-weiler/coinspy/types"
)

const (
	ProviderName string = "Cryptowatch"
)

type Cryptowatch struct {
	client             *resty.Client
	market             string
	providerWithMarket string
}

func New() (p Cryptowatch) {
	p.client = resty.New()
	p.SetMarket("Kraken")
	return p
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
