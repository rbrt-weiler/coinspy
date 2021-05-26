package livecoinwatch

// PriceQuery encapsulates an API request for an exchange rate.
type PriceQuery struct {
	Currency string `json:"currency"`
	Code     string `json:"code"`
	Meta     bool   `json:"meta"`
}
