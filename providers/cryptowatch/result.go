package cryptowatch

type Result struct {
	Result struct {
		Price float64 `json:"price"`
	} `json:"result,omitempty"`
	Allowance struct {
		Cost      float64 `json:"cost"`
		Remaining float64 `json:"remaining"`
	} `json:"allowance,omitempty"`
	Error string `json:"error,omitempty"`
}
