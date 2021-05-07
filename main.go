package main

import (
	"fmt"
	"os"
	"sort"
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

	sort.Slice(rates.Rates, func(i, j int) bool {
		left := fmt.Sprintf("%s-%s-%s", rates.Rates[i].Coin, rates.Rates[i].Fiat, rates.Rates[i].Market)
		right := fmt.Sprintf("%s-%s-%s", rates.Rates[j].Coin, rates.Rates[j].Fiat, rates.Rates[j].Market)
		return left < right
	})

	for _, rate := range rates.Rates {
		if rate.Error == nil {
			resultSet = append(resultSet, rate.String())
		}
	}

	for _, line := range resultSet {
		cons.Println(line)
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
