package databases

import (
	"fmt"
	"net"
	"strings"
	"time"

	"gitlab.com/rbrt-weiler/coinspy/core"
	"gitlab.com/rbrt-weiler/coinspy/types"
)

// QDBStoreExchangeRates stores a set of exchange rates in a QuestDB database.
func QuestDBStoreExchangeRates(rates *types.ExchangeRates) (err error) {
	var qdbConn net.Conn
	var rate types.ExchangeRate
	var influxLine string
	var timeout time.Duration

	_, err = net.LookupHost(core.Config.QuestDB.Host)
	if err != nil {
		err = fmt.Errorf("could not resolve QuestDB host: %s", err)
		return
	}
	timeout, err = time.ParseDuration("10s")
	if err != nil {
		err = fmt.Errorf("could not set timeout: %s", err)
		return
	}
	qdbConn, err = net.DialTimeout("tcp", fmt.Sprintf("%s:%d", core.Config.QuestDB.Host, core.Config.QuestDB.Port), timeout)
	if err != nil {
		err = fmt.Errorf("could not connect to QuestDB host: %s", err)
		return
	}
	defer qdbConn.Close()

	for _, rate = range rates.Rates {
		influxLine = fmt.Sprintf(`exchange_rates,provider=%s,market=%s coin="%s",fiat="%s",rate=%f %d`, strings.ToLower(rate.Provider), strings.ToLower(rate.Market), strings.ToUpper(rate.Coin), strings.ToUpper(rate.Fiat), rate.Rate, rate.AsOf.UTC().UnixNano())
		influxLine = fmt.Sprintf("%s\n", influxLine)
		_, err = qdbConn.Write([]byte(influxLine))
		if err != nil {
			err = fmt.Errorf("could not write to QuestDB: %s", err)
			return
		}
	}

	return
}
