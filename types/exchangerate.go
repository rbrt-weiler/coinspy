package types

import (
	"fmt"
	"time"
)

// ExchangeRate represents a single exchange rate.
type ExchangeRate struct {
	// Provider contains the name of the provider from which the exchange rate was retrieved.
	Provider string
	// Market contains the market that was used for retrieving the exchange rate.
	Market string
	// ProviderWithMarket is the combination of Provider/Market.
	ProviderWithMarket string
	// AsOf contains the information when the exchange rate was retrieved.
	AsOf time.Time
	// Coin contains the common symbol of the coin.
	Coin string
	// Owned contains the amount of Coin owned.
	Owned float64
	// Fiat contains the common symbol of the fiat.
	Fiat string
	// Rate is the exchange rate "1 Coin = Rate Fiat".
	Rate float64
	// Error is used to convey possible errors.
	Error error
}

// String returns the ExchangeRate in a printable string representation.
func (r *ExchangeRate) String() string {
	if r.Owned > 0 {
		return fmt.Sprintf("%f %s = %f %s (on %s as of %s)", r.Owned, r.Coin, (r.Owned * r.Rate), r.Fiat, r.ProviderWithMarket, r.AsOf.Format(time.RFC3339))
	}
	return fmt.Sprintf("1 %s = %f %s (on %s as of %s)", r.Coin, r.Rate, r.Fiat, r.ProviderWithMarket, r.AsOf.Format(time.RFC3339))
}

// StringCompact is a compact variant of String(), omitting time information.
func (r *ExchangeRate) StringCompact() string {
	if r.Owned > 0 {
		return fmt.Sprintf("%f %s = %f %s (on %s)", r.Owned, r.Coin, (r.Owned * r.Rate), r.Fiat, r.ProviderWithMarket)
	}
	return fmt.Sprintf("1 %s = %f %s (on %s)", r.Coin, r.Rate, r.Fiat, r.ProviderWithMarket)
}

// StringVeryCompact is a very compact variant of String(), omitting time and provider information.
func (r *ExchangeRate) StringVeryCompact() string {
	if r.Owned > 0 {
		return fmt.Sprintf("%f %s = %f %s", r.Owned, r.Coin, (r.Owned * r.Rate), r.Fiat)
	}
	return fmt.Sprintf("1 %s = %f %s", r.Coin, r.Rate, r.Fiat)
}
