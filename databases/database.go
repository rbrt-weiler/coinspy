package databases

import (
	"gitlab.com/rbrt-weiler/coinspy/databases/questdb"
	"gitlab.com/rbrt-weiler/coinspy/types"
)

// Database is the common interface used for every specific database.
type Database interface {
	StoreExchangeRates(rates *types.ExchangeRates) (err error)
}

// QuestDB returns an initialized database implementation.
func QuestDB() (db *questdb.QuestDB) {
	database := questdb.New()
	return &database
}
