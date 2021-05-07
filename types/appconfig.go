package types

type AppConfig struct {
	Provider string
	Markets  string
	Coins    string
	Fiats    string
	Pushover struct {
		Token   string
		User    string
		Enabled bool
	}
}
