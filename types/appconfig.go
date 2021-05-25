package types

// AppConfig is used to store application configuration.
type AppConfig struct {
	List struct {
		Providers bool
	}
	Providers string
	Coins     string
	Fiats     string
	QuestDB   struct {
		Host    string
		Port    uint16
		Enabled bool
	}
	Pushover struct {
		Token   string
		User    string
		Enabled bool
	}
	Disable struct {
		QuestDB  bool
		Pushover bool
	}
	Quiet             bool
	CompactOutput     bool
	VeryCompactOutput bool
}
