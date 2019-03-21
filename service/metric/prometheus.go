// Copyright 2018 The QOS Authors

package metric

import (
	"context"
	"errors"
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"

	"github.com/QOSGroup/qmoon/types"
	"github.com/QOSGroup/qmoon/utils"
	"github.com/prometheus/client_golang/api"
	"github.com/prometheus/client_golang/api/prometheus/v1"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/push"
	"github.com/prometheus/common/model"
	"github.com/spf13/viper"
)

var (
	validatorVotingPower = prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "validator_voting_power",
	})
)

const (
	votingPowerPrefix        = "validator_voting_power_"
	votingPowerPercentPrefix = "validator_voting_power_percent_"
	uptimePrefix             = "validator_uptime_"
)

func ValidatorVotingPower(chainID string, vals []types.Validator) {
	pushName := viper.GetString(types.FlagPrometheusPushName)
	pushGateway := viper.GetString(types.FlagPrometheusPushGateway)

	var totalPower int64
	for _, v := range vals {
		totalPower += v.VotingPower
	}
	pusher := push.New(pushGateway, pushName)
	for _, v := range vals {
		powerGauge := prometheus.NewGauge(prometheus.GaugeOpts{
			Name: votingPowerPrefix + strings.Replace(chainID, "-", "_", -1),
			ConstLabels: map[string]string{
				"address": v.Address,
			}})
		powerGauge.Set(float64(v.VotingPower))
		pusher.Collector(powerGauge)

		uptimeGauge := prometheus.NewGauge(prometheus.GaugeOpts{
			Name: uptimePrefix + strings.Replace(chainID, "-", "_", -1),
			ConstLabels: map[string]string{
				"address": v.Address,
			}})
		d, _ := strconv.ParseFloat(v.Uptime, 64)
		uptimeGauge.Set(d)
		pusher.Collector(uptimeGauge)

		if totalPower > 0 {
			percentGauge := prometheus.NewGauge(prometheus.GaugeOpts{
				Name: votingPowerPercentPrefix + strings.Replace(chainID, "-", "_", -1),
				ConstLabels: map[string]string{
					"address": v.Address,
				}})

			var percent float64
			if v.VotingPower > 0 {
				percent = utils.Percent(uint64(v.VotingPower), uint64(totalPower)) * 100
			}
			percentGauge.Set(percent)
			pusher.Collector(percentGauge)
		}
	}

	err := pusher.Add()
	if err != nil {
		log.Printf("push err:%s", err.Error())
	}
}

func newQuery(prefix, chainid, addr string) string {
	q := prefix + strings.Replace(chainid, "-", "_", -1)

	if addr != "" {
		q = q + fmt.Sprintf(`{address="%s"}`, addr)
	}

	return q
}

func QueryValidatorVotingPower(chainid, addr string, start, end time.Time, step time.Duration) ([]types.ResultMatrix, error) {
	prometheusServer := viper.GetString(types.FlagPrometheusServer)

	cli, err := api.NewClient(api.Config{Address: prometheusServer})
	if err != nil {
		return nil, err
	}
	papi := v1.NewAPI(cli)

	ds, err := papi.QueryRange(context.Background(), newQuery(votingPowerPrefix, chainid, addr), v1.Range{
		Start: start,
		End:   end,
		Step:  step,
	})
	if err != nil {
		return nil, err
	}

	if ds.Type() == model.ValMatrix {
		return parseMatrix(ds.(model.Matrix)), nil
	} else {
		return nil, errors.New("unsupport data")
	}
}

func QueryValidatorVotingPowerPercent(chainid, addr string, start, end time.Time, step time.Duration) ([]types.ResultMatrix, error) {
	prometheusServer := viper.GetString(types.FlagPrometheusServer)

	cli, err := api.NewClient(api.Config{Address: prometheusServer})
	if err != nil {
		return nil, err
	}
	papi := v1.NewAPI(cli)

	ds, err := papi.QueryRange(context.Background(), newQuery(votingPowerPercentPrefix, chainid, addr), v1.Range{
		Start: start,
		End:   end,
		Step:  step,
	})
	if err != nil {
		return nil, err
	}

	if ds.Type() == model.ValMatrix {
		return parseMatrix(ds.(model.Matrix)), nil
	} else {
		return nil, errors.New("unsupport data")
	}
}

func QueryValidatorUptime(chainid, addr string, start, end time.Time, step time.Duration) ([]types.ResultMatrix, error) {
	prometheusServer := viper.GetString(types.FlagPrometheusServer)

	cli, err := api.NewClient(api.Config{Address: prometheusServer})
	if err != nil {
		return nil, err
	}
	papi := v1.NewAPI(cli)

	ds, err := papi.QueryRange(context.Background(), newQuery(uptimePrefix, chainid, addr), v1.Range{
		Start: start,
		End:   end,
		Step:  step,
	})
	if err != nil {
		return nil, err
	}

	if ds.Type() == model.ValMatrix {
		return parseMatrix(ds.(model.Matrix)), nil
	} else {
		return nil, errors.New("unsupport data")
	}
}

func parseMatrix(m model.Matrix) []types.ResultMatrix {
	var result []types.ResultMatrix
	for _, v := range m {
		for _, vv := range v.Values {
			result = append(result, types.ResultMatrix{
				X: vv.Timestamp.String(),
				Y: vv.Value.String(),
			})
		}

	}

	return result
}
