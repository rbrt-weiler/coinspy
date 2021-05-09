package coingate

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
	ProviderName string = "CoinGate"
	APIBaseURL   string = "https://api.coingate.com/v2"
)

type CoinGate struct {
	client     *resty.Client
	market     string
	prices     Prices
	pricesAsOf time.Time
	Error      error
}

func New() (p CoinGate) {
	p.client = resty.New()
	p.market = "default"
	p.FetchPrices()
	return
}

func (p *CoinGate) FetchPrices() (err error) {
	var resp *resty.Response
	var priceList Prices

	apiURL := fmt.Sprintf("%s/rates/merchant", APIBaseURL)
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

func (p *CoinGate) SetMarket(market string) (err error) {
	return nil
}

func (p *CoinGate) FetchRate(coin string, fiat string) (rate types.ExchangeRate, err error) {
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

	fiatID := strings.ToUpper(fiat)
	coinID := strings.ToUpper(coin)
	if rateString, ok = p.prices[coinID][fiatID]; !ok {
		err = fmt.Errorf("rate %s/%s is unknown to %s/%s", coinID, fiatID, ProviderName, p.market)
		return
	}
	rateFloat, err = strconv.ParseFloat(rateString, 64)
	if err != nil {
		err = fmt.Errorf("price could not be parsed to a float: %s", err)
		return
	}

	return types.ExchangeRate{Provider: ProviderName, Market: p.market, AsOf: p.pricesAsOf, Coin: coin, Fiat: fiat, Rate: rateFloat, Error: err}, err
}

func (p *CoinGate) FetchRateSynced(coin string, fiat string, rates *types.ExchangeRates, wg *sync.WaitGroup) {
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
