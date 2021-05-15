package core

import (
	"fmt"

	"gitlab.com/rbrt-weiler/coinspy/providers/cryptowatch"
)

func ProviderList() (providers map[string][]string, err error) {
	var providerName string
	var marketList []string

	providers = make(map[string][]string)

	for _, provider := range []string{"CoinGate", "Coingecko"} {
		providers[provider] = append(providers[provider], "default")
	}

	providerName = "Cryptowatch"
	marketList, err = cryptowatch.ListMarkets()
	if err != nil {
		err = fmt.Errorf("could not fetch Cryptowatch markets: %s", err)
		return
	}
	for _, market := range marketList {
		providers[providerName] = append(providers[providerName], market)
	}

	return
}
