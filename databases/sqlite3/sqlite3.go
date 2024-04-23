package sqlite3

import (
	"database/sql"
	"fmt"
	"strings"

	_ "github.com/mattn/go-sqlite3"
	"gitlab.com/rbrt-weiler/coinspy/core"
	"gitlab.com/rbrt-weiler/coinspy/types"
)

// SQLite3 is a specific implementation of a Database.
type SQLite3 struct {
	// Error is used to convey possible errors.
	Error error
}

// New initializes and returns a usable Database object.
func New() (db SQLite3) {
	db.Error = nil
	return
}

// StoreExchangeRates stores a set of exchange rates in a SQLite3 database.
func (db *SQLite3) StoreExchangeRates(rates *types.ExchangeRates) (err error) {
	var sqlDb *sql.DB

	sqlDb, err = sql.Open("sqlite3", core.Config.SQLite3.File)
	if err != nil {
		err = fmt.Errorf("could not open SQLite3 database: %s", err)
		return
	}
	defer sqlDb.Close()

	err = createTable(sqlDb)
	if err != nil {
		err = fmt.Errorf("could not create table: %s", err)
		return
	}

	err = insertValues(sqlDb, rates)
	if err != nil {
		err = fmt.Errorf("could not insert values: %s", err)
		return
	}

	return
}

// sanitizeTableName returns a safe variant of the table name passed via arguments.
func sanitizeTableName() (tblName string) {
	// This is extremely basic.
	tblName = strings.ReplaceAll(core.Config.SQLite3.Table, "'", "")
	return
}

// createTable contains all code to create the required table in the database along with indices.
func createTable(sqlDb *sql.DB) (err error) {
	var tableName string
	var tx *sql.Tx
	var stmt *sql.Stmt

	tableName = sanitizeTableName()

	tx, err = sqlDb.Begin()
	if err != nil {
		err = fmt.Errorf("could not create database transaction: %s", err)
		return
	}
	stmt, err = tx.Prepare(fmt.Sprintf(`
		CREATE TABLE IF NOT EXISTS '%s' (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			provider VARCHAR(100),
			market VARCHAR(100),
			coin VARCHAR(50),
			fiat VARCHAR(50),
			rate FLOAT,
			rateAsOf TEXT
		);
		CREATE INDEX IF NOT EXISTS idx_coin ON '%s' (coin);
		CREATE INDEX IF NOT EXISTS idx_fiat ON '%s' (fiat);
		CREATE INDEX IF NOT EXISTS idx_rateasof ON '%s' (rateAsOf);
	`, tableName, tableName, tableName, tableName))
	if err != nil {
		err = fmt.Errorf("could not prepare SQL statement: %s", err)
		return
	}
	defer stmt.Close()
	_, err = stmt.Exec()
	if err != nil {
		err = fmt.Errorf("could not execute SQL statement: %s", err)
		return
	}
	err = tx.Commit()
	if err != nil {
		err = fmt.Errorf("could not commit SQL statement: %s", err)
		return
	}

	return
}

// insertValues actually inserts exchange rates into the database.
func insertValues(sqlDb *sql.DB, rates *types.ExchangeRates) (err error) {
	var tableName string
	var tx *sql.Tx
	var stmt *sql.Stmt
	var rate types.ExchangeRate

	tableName = sanitizeTableName()

	tx, err = sqlDb.Begin()
	if err != nil {
		err = fmt.Errorf("could not create database transaction: %s", err)
		return
	}
	stmt, err = tx.Prepare(fmt.Sprintf(`
		INSERT INTO '%s' (
			provider,
			market,
			coin,
			fiat,
			rate,
			rateAsOf
		) VALUES (
			?,
			?,
			?,
			?,
			?,
			?
		);
	`, tableName))
	if err != nil {
		err = fmt.Errorf("could not prepare SQL statement: %s", err)
		return
	}
	defer stmt.Close()
	for _, rate = range rates.Rates {
		_, err = stmt.Exec(strings.ToLower(rate.Provider), strings.ToLower(rate.Market), strings.ToUpper(rate.Coin), strings.ToUpper(rate.Fiat), rate.Rate, rate.AsOf.Format("2006-01-02T15:04:05.000000000-07:00"))
		if err != nil {
			err = fmt.Errorf("could not execute SQL statement: %s", err)
			return
		}
	}
	err = tx.Commit()
	if err != nil {
		err = fmt.Errorf("could not commit SQL statement: %s", err)
		return
	}

	return
}
