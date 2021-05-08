package coingecko

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
	ProviderName string = "Coingecko"
	APIBaseURL   string = "https://api.coingecko.com/api/v3"
)

type Coingecko struct {
	client *resty.Client
	coins  CoinList
	Error  error
}

func New() (p Coingecko) {
	p.client = resty.New()
	p.Error = p.PopulateCoinList()
	return
}

func (p *Coingecko) PopulateCoinList() (err error) {
	var resp *resty.Response

	apiURL := fmt.Sprintf("%s/coins/list", APIBaseURL)
	resp, err = p.client.R().Get(apiURL)
	if err != nil {
		err = fmt.Errorf("could not fetch exchange rate: %s", err)
		return
	} else {
		err = json.Unmarshal(resp.Body(), &p.coins)
		if err != nil {
			err = fmt.Errorf("could not unmarshal JSON: %s", err)
			return
		}
	}

	return
}

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

	return
}

func (p Coingecko) FetchRate(market string, coin string, fiat string) (rate types.ExchangeRate, err error) {
	var coinID string
	var resp *resty.Response
	var priceList Prices

	apiURL := fmt.Sprintf("%s/simple/price", APIBaseURL)

	coinID, err = p.SymbolToID(coin)
	if err != nil {
		return
	}
	fiatID := strings.ToLower(fiat)

	resp, err = p.client.R().
		SetQueryParam("ids", coinID).
		SetQueryParam("vs_currencies", fiatID).
		Get(apiURL)
	if err != nil {
		return
	}

	err = json.Unmarshal(resp.Body(), &priceList)
	if err != nil {
		return
	}

	return types.ExchangeRate{Provider: ProviderName, Market: market, AsOf: resp.ReceivedAt(), Coin: coin, Fiat: fiat, Rate: priceList[coinID][fiatID], Error: err}, err
}

func (p Coingecko) FetchRateSynced(rates *types.ExchangeRates, market string, coin string, fiat string, wg *sync.WaitGroup) {
	rate, err := p.FetchRate(market, coin, fiat)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %s\n", err)
	}
	rates.Mutex.Lock()
	rates.Rates = append(rates.Rates, rate)
	rates.Mutex.Unlock()
	wg.Done()
}
