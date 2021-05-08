package main

import (
	"os"
	"strings"
	"sync"
	"time"

	"gitlab.com/rbrt-weiler/coinspy/core"
	"gitlab.com/rbrt-weiler/coinspy/providers"
	"gitlab.com/rbrt-weiler/coinspy/types"
)

func init() {
	core.SetupFlags()
}

func main() {
	var provider providers.Provider
	var rates types.ExchangeRates
	var wg sync.WaitGroup
	var resultSet []string

	config := &core.Config
	cons := &core.Cons

	core.CheckArguments()

	provider = providers.Cryptowatch()
	//provider = providers.Coingecko()
	markets := strings.Split(config.Markets, ",")
	coins := strings.Split(config.Coins, ",")
	fiats := strings.Split(config.Fiats, ",")

	for _, market := range markets {
		for _, coin := range coins {
			for _, fiat := range fiats {
				wg.Add(1)
				go provider.FetchRateSynced(&rates, market, coin, fiat, &wg)
			}
		}
	}
	wg.Wait()

	rates.Sort()
	for _, rate := range rates.Rates {
		if rate.Error == nil {
			resultSet = append(resultSet, rate.String())
		}
	}

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
