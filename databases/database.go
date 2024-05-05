package databases

import (
	"gitlab.com/rbrt-weiler/coinspy/databases/duckdb"
	"gitlab.com/rbrt-weiler/coinspy/databases/questdb"
	"gitlab.com/rbrt-weiler/coinspy/databases/sqlite3"
	"gitlab.com/rbrt-weiler/coinspy/types"
)

// Database is the common interface used for every specific database.
type Database interface {
	StoreExchangeRates(rates *types.ExchangeRates) (err error)
}

// DuckDB returns an initialized database implementation.
func DuckDB() (db *duckdb.DuckDB) {
	database := duckdb.New()
	return &database
}

// QuestDB returns an initialized database implementation.
func QuestDB() (db *questdb.QuestDB) {
	database := questdb.New()
	return &database
}

// SQLite3 returns an initialized database implementation.
func SQLite3() (db *sqlite3.SQLite3) {
	database := sqlite3.New()
	return &database
}
