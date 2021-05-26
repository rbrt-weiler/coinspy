package livecoinwatch

// PriceResult stores a single exchange rate as returned by the LiveCoinWatch API.
type PriceResult struct {
	Name              string      `json:"name,omitempty"`
	Symbol            string      `json:"symbol,omitempty"`
	Color             string      `json:"color,omitempty"`
	PNG32             string      `json:"png32,omitempty"`
	PNG64             string      `json:"png64,omitempty"`
	WebP32            string      `json:"webp32,omitempty"`
	WebP64            string      `json:"webp64,omitempty"`
	Exchanges         int         `json:"exchanges,omitempty"`
	Markets           int         `json:"markets,omitempty"`
	Pairs             int         `json:"pairs,omitempty"`
	AllTimeHighUSD    float64     `json:"allTimeHighUSD,omitempty"`
	CirculatingSupply int64       `json:"circulatingSupply,omitempty"`
	TotalSupply       interface{} `json:"totalSupply,omitempty"`
	MaxSupply         interface{} `json:"maxSupply,omitempty"`
	Rate              float64     `json:"rate"`
	Volume            int64       `json:"volume"`
	Cap               int64       `json:"cap"`
	Error             struct {
		Code        int    `json:"code"`
		Status      string `json:"status"`
		Description string `json:"description"`
	} `json:"error,omitempty"`
}
