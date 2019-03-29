// Copyright 2019 The QOS Authors

package metric

import (
	"testing"
)

func TestQueryValidatorsUptimeV2(t *testing.T) {
	//s := "http://192.168.1.243:8086/"
	//cli, err := client.NewHTTPClient(client.HTTPConfig{
	//	Addr: s,
	//})
	//assert.Nil(t, err)
	//
	//now := time.Now()
	//
	//dataType := "votingPowerPercent"
	//dur := "1d"
	//address := "2199EAE894CA391FA82F01C2C614BFEB103D056C"
	//command := fmt.Sprintf(`SELECT median("%s") FROM "validator" WHERE ("address" = $address) AND time >= $start and time <= $end GROUP BY time(%s) fill(0)`,
	//	dataType, dur)
	//database := "cosmoshub-1"
	//precision := defaultPrecision
	//parameters := map[string]interface{}{
	//	"address": address,
	//	"start":   utils.NDaysAgo(now, 28),
	//	"end":     now,
	//}
	//q := client.NewQueryWithParameters(command, database, precision, parameters)
	//response, err := cli.Query(q)
	//assert.Nil(t, err)
	//assert.Nil(t, response.Error())
	//
	//res := parseInfluxdbValud(address, response.Results)
	//t.Logf("---res:%+v", res)
}
