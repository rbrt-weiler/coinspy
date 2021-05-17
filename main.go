package main

import (
	"fmt"
	"os"
	"strings"
	"sync"
	"time"

	resty "github.com/go-resty/resty/v2"

	"gitlab.com/rbrt-weiler/coinspy/core"
	"gitlab.com/rbrt-weiler/coinspy/providers"
	"gitlab.com/rbrt-weiler/coinspy/types"
)

func listProviders() {
	cons := &core.Cons

	prov, err := core.ProviderList()
	if err != nil {
		cons.Fprintf(os.Stderr, "Error: %s\n", err)
		os.Exit(core.ErrGeneric)
	}
	for p, mar := range prov {
		for _, m := range mar {
			if m != "default" {
				cons.Printf("%s/%s\n", p, m)
			} else {
				cons.Printf("%s\n", p)
			}
		}
	}

	os.Exit(core.ErrSuccess)
}

func initializeProvider(providerName string, httpClient *resty.Client) (provider providers.Provider, err error) {
	switch strings.ToLower(providerName) {
	case "coingate":
		provider = providers.CoinGate(httpClient)
	case "coingecko":
		provider = providers.Coingecko(httpClient)
	case "cryptowatch":
		provider = providers.Cryptowatch(httpClient)
	default:
		err = fmt.Errorf("provider %s is unknown", providerName)
	}

	return
}

func fetchRates() (rates types.ExchangeRates) {
	var providerName string
	var markets []string
	var provider providers.Provider
	var err error
	var wg sync.WaitGroup

	config := &core.Config
	cons := &core.Cons

	coins := strings.Split(config.Coins, ",")
	fiats := strings.Split(config.Fiats, ",")
	client := resty.New()
	client.SetHeader("User-Agent", core.ToolID)

	for _, singleProvider := range strings.Split(config.Providers, ",") {
		if strings.Contains(singleProvider, "/") {
			parts := strings.Split(singleProvider, "/")
			providerName = parts[0]
			markets = parts[1:]
		} else {
			providerName = singleProvider
			markets = []string{"default"}
		}
		provider, err = initializeProvider(providerName, client)
		if err != nil {
			cons.Fprintf(os.Stderr, "Error: %s\n", err)
			os.Exit(core.ErrGeneric)
		}
		for _, market := range markets {
			err = provider.SetMarket(market)
			if err != nil {
				cons.Fprintf(os.Stderr, "Error: %s\n", err)
				os.Exit(core.ErrGeneric)
			}
			for _, coin := range coins {
				for _, fiat := range fiats {
					wg.Add(1)
					go provider.FetchRateSynced(coin, fiat, &rates, &wg)
				}
			}
		}
	}
	wg.Wait()

	return
}

func ratesToStrings(rates *types.ExchangeRates) (resultSet []string) {
	config := &core.Config

	rates.Sort()
	for _, rate := range rates.Rates {
		if rate.Error == nil {
			if config.VeryCompactOutput {
				resultSet = append(resultSet, rate.StringVeryCompact())
			} else if config.CompactOutput {
				resultSet = append(resultSet, rate.StringCompact())
			} else {
				resultSet = append(resultSet, rate.String())
			}
		}
	}

	return
}

func init() {
	core.SetupFlags()
}

func main() {
	var rates types.ExchangeRates
	var resultSet []string

	config := &core.Config
	cons := &core.Cons

	if config.List.Providers {
		listProviders()
	}

	core.CheckArguments()

	rates = fetchRates()
	resultSet = ratesToStrings(&rates)

	if !config.Quiet {
		for _, line := range resultSet {
			cons.Println(line)
		}
	}
	if config.Pushover.Enabled {
		poErr := core.SendPushoverMessage(config.Pushover.Token, config.Pushover.User, strings.Join(resultSet, "\r\n"), time.Now())
		if poErr != nil {
			cons.Fprintf(os.Stderr, "Error: %s\n", poErr)
			os.Exit(core.ErrGeneric)
		}
	}

	os.Exit(core.ErrSuccess)
}
