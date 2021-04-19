package cryptowatch

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
	"sync"

	"github.com/go-resty/resty/v2"
	"gitlab.com/rbrt-weiler/coinspy/types"
)

const (
	ProviderName string = "Cryptowatch"
)

type Cryptowatch struct {
	client *resty.Client
}

func New() (p Cryptowatch) {
	p.client = resty.New()
	return
}

func (p Cryptowatch) FetchRate(market string, coin string, fiat string) (rate types.ExchangeRate, err error) {
	var apiURL string
	var resp *resty.Response
	var apiResult Result
	var rateValue float64

	rateValue = 0
	apiURL = fmt.Sprintf("https://api.cryptowat.ch/markets/%s/%s%s/price", strings.ToLower(market), strings.ToLower(coin), strings.ToLower(fiat))
	resp, err = p.client.R().Get(apiURL)
	if err != nil {
		err = fmt.Errorf("could not fetch exchange rate: %s", err)
	} else {
		json.Unmarshal(resp.Body(), &apiResult)
		if apiResult.Error != "" {
			err = fmt.Errorf("%s (%s %s/%s; %f allowance remaining)", apiResult.Error, market, coin, fiat, apiResult.Allowance.Remaining)
		} else {
			rateValue = apiResult.Result.Price
		}
	}

	return types.ExchangeRate{Provider: "Cryptowatch", Market: market, AsOf: resp.ReceivedAt(), Coin: coin, Fiat: fiat, Rate: rateValue, Error: err}, err
}

func (p Cryptowatch) FetchRateSynced(rates *types.ExchangeRates, market string, coin string, fiat string, wg *sync.WaitGroup) {
	rate, err := p.FetchRate(market, coin, fiat)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %s\n", err)
	}
	rates.Mutex.Lock()
	rates.Rates = append(rates.Rates, rate)
	rates.Mutex.Unlock()
	wg.Done()
}