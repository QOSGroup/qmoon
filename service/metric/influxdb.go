// Copyright 2019 The QOS Authors

package metric

import (
	"errors"
	"fmt"
	"time"

	"github.com/QOSGroup/qmoon/types"
	"github.com/QOSGroup/qmoon/utils"
	"github.com/influxdata/influxdb/client/v2"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

var influx client.Client

const (
	votingPower          = "votingPower"
	votingPowerPercent   = "votingPowerPercent"
	uptime               = "uptime"
	validatorAddr        = "address"
	validatorName        = "name"
	validatorMeasurement = "validator"
	defaultPrecision     = "s"
)

func getInflux() client.Client {
	if influx != nil {
		return influx
	}
	cli, err := client.NewHTTPClient(client.HTTPConfig{
		Addr:     viper.GetString(types.FlagInfluxdbServer),
		Username: viper.GetString(types.FlagInfluxdbUser),
		Password: viper.GetString(types.FlagInfluxdbPassword),
	})
	logrus.Warnf("FlagInfluxdbServer:%s", viper.GetString(types.FlagInfluxdbServer))
	logrus.Warnf("FlagInfluxdbUser:%s", viper.GetString(types.FlagInfluxdbUser))
	logrus.Warnf("FlagInfluxdbPassword:%s", viper.GetString(types.FlagInfluxdbPassword))
	if err != nil {
		logrus.Errorf("influxdb NewHTTPClient err:%s", err.Error())
	}
	_, _, err = cli.Ping(0)
	if err != nil {
		logrus.Errorf("influxdb Ping err:%s", err.Error())
	}

	influx = cli
	return influx
}

func CreateDatabase(chainID string) error {
	cli := getInflux()
	if cli == nil {
		return nil
	}

	cmd := fmt.Sprintf("CREATE DATABASE \"%s\"", chainID)
	q := client.NewQuery(cmd, "", "")
	response, err := cli.Query(q)
	if err != nil {
		return err
	}

	if response.Error() != nil {
		return response.Error()
	}

	return nil
}

func DeleteDatabase(chainID string) error {
	cli := getInflux()
	if cli == nil {
		return nil
	}

	cmd := fmt.Sprintf("DROP DATABASE \"%s\"", chainID)
	q := client.NewQuery(cmd, "", "")
	response, err := cli.Query(q)
	if err != nil {
		return err
	}

	if response.Error() != nil {
		return response.Error()
	}

	return nil
}

func ValidatorVotingPower(chainID string, t time.Time, vals []types.Validator) {
	cli := getInflux()
	if cli == nil {
		return
	}

	bp, err := client.NewBatchPoints(client.BatchPointsConfig{
		Precision: defaultPrecision,
		Database:  chainID,
	})
	if err != nil {
		logrus.Warnf("influxdb NewBatchPoints err:%v", err.Error())
		return
	}

	var totalPower int64
	for _, v := range vals {
		totalPower += v.VotingPower
	}

	for _, v := range vals {
		tag := map[string]string{
			validatorAddr: v.Address,
			//validatorName: v.Name,
		}

		fields := map[string]interface{}{
			votingPower: v.VotingPower,
			uptime:      v.UptimeFloat,
		}

		if totalPower > 0 {
			var percent float64
			if v.VotingPower > 0 {
				percent = utils.Percent(uint64(v.VotingPower), uint64(totalPower)) * 100
			}
			fields[votingPowerPercent] = percent
		}

		point, err := client.NewPoint(validatorMeasurement, tag, fields, t)
		if err == nil {
			bp.AddPoint(point)
		}
	}

	if err := cli.Write(bp); err != nil {
		logrus.Warnf("influxdb write err:%v", err.Error())
		return
	}
}

func QueryValidatorVotingPower(chainid, addr string, start, end time.Time, step string) ([]types.Matrix, error) {
	var result []types.Matrix
	res, err := queryInfluxdb(votingPower, chainid, addr, start, end, step)
	if err != nil {
		return nil, err
	}
	if len(res) != 0 {
		result = res[0].List
	}

	return result, nil
}

func QueryValidatorVotingPowerPercent(chainid, addr string, start, end time.Time, step string) ([]types.Matrix, error) {
	var result []types.Matrix
	res, err := queryInfluxdb(votingPowerPercent, chainid, addr, start, end, step)
	if err != nil {
		return nil, err
	}
	if len(res) != 0 {
		result = res[0].List
	}

	return result, nil
}

func QueryValidatorUptime(chainid, addr string, start, end time.Time, step string) ([]types.Matrix, error) {
	var result []types.Matrix
	res, err := queryInfluxdb(uptime, chainid, addr, start, end, step)
	if err != nil {
		return nil, err
	}
	if len(res) != 0 {
		result = res[0].List
	}

	return result, nil
}

func QueryValidatorsVotingPowerPercent(chainid string, start, end time.Time, step string) ([]types.ResultMatrix, error) {
	return nil, nil
}

func QueryValidatorsUptime(chainid string, start, end time.Time, step string) ([]types.ResultMatrix, error) {
	return nil, nil
}

func queryInfluxdb(dateType, chainid, addr string, start, end time.Time, step string) ([]types.ResultMatrix, error) {
	cli := getInflux()
	if cli == nil {
		return nil, errors.New("init client error")
	}

	command := fmt.Sprintf(`SELECT median("%s") FROM "validator" WHERE ("address" = $address) AND time >= $start and time <= $end GROUP BY time(%s) fill(0)`,
		dateType, step)
	database := chainid
	precision := defaultPrecision
	parameters := map[string]interface{}{
		"address": addr,
		"start":   start,
		"end":     end,
	}
	q := client.NewQueryWithParameters(command, database, precision, parameters)
	response, err := cli.Query(q)
	if err != nil {
		return nil, err
	}

	if response.Error() != nil {
		return nil, response.Error()
	}

	return parseInfluxdbValud(addr, response.Results), nil
}

func parseInfluxdbValud(addr string, results []client.Result) []types.ResultMatrix {
	var result []types.ResultMatrix
	for _, res := range results {
		for _, serie := range res.Series {
			var d types.ResultMatrix
			d.Title = addr
			for _, v := range serie.Values {
				logrus.Warnf("Values:%+v", v)
				if len(v) == 2 {
					d.List = append(d.List, types.Matrix{
						X: fmt.Sprintf("%v", v[0]),
						Y: fmt.Sprintf("%v", v[1]),
					})
				}
			}
			result = append(result, d)
		}
	}

	return result
}
