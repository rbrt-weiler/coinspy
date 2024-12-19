package types

// AppConfig is used to store application configuration.
type AppConfig struct {
	List struct {
		Providers bool
	}
	Providers     string
	Coins         string
	Fiats         string
	LiveCoinWatch struct {
		APIKey string
	}
	DuckDB struct {
		File    string
		Table   string
		Enabled bool
	}
	QuestDB struct {
		Host    string
		Port    uint16
		Table   string
		Timeout uint16
		Enabled bool
	}
	SQLite3 struct {
		File    string
		Table   string
		Enabled bool
	}
	Pushover struct {
		Token        string
		User         string
		IncludeLinks bool
		IncludeHost  bool
		Enabled      bool
	}
	Disable struct {
		DuckDB   bool
		QuestDB  bool
		SQLite3  bool
		Pushover bool
	}
	Quiet                bool
	CompactOutput        bool
	VeryCompactOutput    bool
	PortfolioValueTop    bool
	PortfolioValueBottom bool
}
