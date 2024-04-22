package databases

import (
	"fmt"
	"net"
	"strings"

	"gitlab.com/rbrt-weiler/coinspy/core"
	"gitlab.com/rbrt-weiler/coinspy/types"
)

// QDBStoreExchangeRates stores a set of exchange rates in a QuestDB database.
func QuestDBStoreExchangeRates(rates *types.ExchangeRates) (err error) {
	var qdbAddr *net.TCPAddr
	var qdbConn *net.TCPConn
	var rate types.ExchangeRate
	var influxLine string

	qdbAddr, err = net.ResolveTCPAddr("tcp", fmt.Sprintf("%s:%d", core.Config.QuestDB.Host, core.Config.QuestDB.Port))
	if err != nil {
		err = fmt.Errorf("could not resolve QuestDB host: %s", err)
		return
	}
	qdbConn, err = net.DialTCP("tcp", nil, qdbAddr)
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
