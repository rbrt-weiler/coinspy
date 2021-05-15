package types

type AppConfig struct {
	List struct {
		Providers bool
	}
	Providers string
	Coins     string
	Fiats     string
	Pushover  struct {
		Token   string
		User    string
		Enabled bool
	}
	Disable struct {
		Pushover bool
	}
	Quiet             bool
	CompactOutput     bool
	VeryCompactOutput bool
}
