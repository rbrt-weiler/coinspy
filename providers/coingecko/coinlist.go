package coingecko

// CoinList stores the symbol to ID mappings required for Coingecko API usage.
type CoinList []struct {
	ID     string `json:"id"`
	Symbol string `json:"symbol"`
	Name   string `json:"name"`
}
