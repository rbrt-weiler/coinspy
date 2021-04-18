package main

import (
	"fmt"
	"os"
	"path"
	"sort"
	"strings"
	"sync"
	"time"

	pflag "github.com/spf13/pflag"

	providers "gitlab.com/rbrt-weiler/coinspy/providers"
	types "gitlab.com/rbrt-weiler/coinspy/types"
	consolehelper "gitlab.com/rbrt-weiler/go-module-consolehelper"
)

const (
	toolName    string = "coinspy"
	toolVersion string = "0.1.0"
	toolID      string = toolName + "/" + toolVersion
	toolURL     string = "https://gitlab.com/rbrt-weiler/coinspy"

	errSuccess int = 0
	errGeneric int = 1
	errUsage   int = 2
)

var (
	config types.AppConfig
	cons   consolehelper.ConsoleHelper
)

func checkArguments() {
	if config.Coins == "" {
		cons.Fprintf(os.Stderr, "Error: No coins provided.\n")
		os.Exit(errGeneric)
	}
	if config.Fiats == "" {
		cons.Fprintf(os.Stderr, "Error: No fiats provided.\n")
		os.Exit(errGeneric)
	}
}

func init() {
	//pflag.StringVarP(&config.Provider, "provider", "P", "Cryptowatch", "Exchange rate provider to use")
	pflag.StringVarP(&config.Markets, "markets", "M", "Kraken", "Markets to use with multi-market providers (comma-seperated)")
	pflag.StringVarP(&config.Coins, "coins", "C", "", "Coins to fetch rates for")
	pflag.StringVarP(&config.Fiats, "fiats", "F", "", "Fiats to fetch rates for")
	pflag.Usage = func() {
		cons.Fprintf(os.Stderr, "%s\n", toolID)
		cons.Fprintf(os.Stderr, "%s\n", toolURL)
		cons.Fprintf(os.Stderr, "\n")
		cons.Fprintf(os.Stderr, "A tool to fetch exchange rates for crypto coins.\n")
		cons.Fprintf(os.Stderr, "\n")
		cons.Fprintf(os.Stderr, "Usage: %s [options]\n", path.Base(os.Args[0]))
		cons.Fprintf(os.Stderr, "\n")
		cons.Fprintf(os.Stderr, "Available options:\n")
		pflag.PrintDefaults()
		os.Exit(errUsage)
	}
	pflag.Parse()
}

func main() {
	var provider providers.Provider
	var rates types.ExchangeRates
	var wg sync.WaitGroup

	checkArguments()

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
			fmt.Printf("1 %s = %f %s (on %s as of %s)\n", rate.Coin, rate.Rate, rate.Fiat, rate.Market, rate.AsOf.Format(time.RFC3339))
		}
	}

	os.Exit(0)
}
