// Copyright 2018 The QOS Authors

package metric

import (
	"fmt"
	"strings"

	"github.com/QOSGroup/qmoon/types"
	"github.com/QOSGroup/qmoon/utils"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/push"
)

var (
	validatorVotingPower = prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "validator_voting_power",
	})
)

var (
	pushName    = "qmoon"
	pushGateway = "http://192.168.1.183:9091"
)

func ValidatorVotingPower(vals []types.Validator) {
	var totalPower int64
	for _, v := range vals {
		totalPower += v.VotingPower
	}
	pusher := push.New(pushGateway, pushName)
	for _, v := range vals {
		powerGauge := prometheus.NewGauge(prometheus.GaugeOpts{
			Name: "validator_voting_power_" + strings.Replace(v.ChainID, "-", "_", -1),
			ConstLabels: map[string]string{
				"address": v.Address,
			}})
		powerGauge.Set(float64(v.VotingPower))
		pusher.Collector(powerGauge)

		if totalPower > 0 {
			percentGauge := prometheus.NewGauge(prometheus.GaugeOpts{
				Name: "validator_voting_power_percent_" + strings.Replace(v.ChainID, "-", "_", -1),
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

	if err := pusher.Add(); err != nil {
		fmt.Printf("push err:%s\n", err.Error())
	}
}
