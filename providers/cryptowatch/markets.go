package cryptowatch

// Markets is used to unmarshal an API result containing the available markets.
type Markets struct {
	Result []struct {
		ID       int    `json:"id"`
		Exchange string `json:"exchange"`
		Pair     string `json:"pair"`
		Active   bool   `json:"active"`
		Route    string `json:"route"`
	} `json:"result"`
}
