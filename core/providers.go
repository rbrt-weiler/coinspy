package core

// ProviderList returns a map containing all possible providers.
// The key of the map is the provider name, while the value - an array - contains all possible markets for that provider.
func ProviderList() (providers map[string][]string, err error) {
	providers = make(map[string][]string)

	for _, provider := range []string{"Bitpanda", "CoinGate", "Coingecko", "LiveCoinWatch"} {
		providers[provider] = append(providers[provider], "default")
	}

	return
}
