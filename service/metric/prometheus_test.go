// Copyright 2018 The QOS Authors

package metric

import (
	"context"
	"testing"
	"time"

	"github.com/QOSGroup/qmoon/types"
	"github.com/QOSGroup/qmoon/utils"
	"github.com/prometheus/client_golang/api"
	"github.com/prometheus/client_golang/api/prometheus/v1"
	"github.com/prometheus/common/model"
	"github.com/stretchr/testify/assert"
)

func TestPush(t *testing.T) {
	var vars []types.Validator
	vars = append(vars, types.Validator{
		ChainID:     "qos_test",
		Address:     "a",
		VotingPower: time.Now().Unix() % 100 * 2,
	})
	vars = append(vars, types.Validator{
		ChainID:     "qos_test",
		Address:     "b",
		VotingPower: time.Now().Unix() % 80 * 2,
	})

	ValidatorVotingPower(vars)
}

func TestQuery(t *testing.T) {
	cli, err := api.NewClient(api.Config{Address: prometheusServer})
	assert.Nil(t, err)

	papi := v1.NewAPI(cli)

	now := time.Now()
	ds, err := papi.QueryRange(context.Background(), "validator_voting_power_capricorn_1000", v1.Range{
		Start: utils.NDaysAgo(now, 1),
		End:   now,
		Step:  time.Minute * 10,
	})
	assert.Nil(t, err)

	if ds.Type() == model.ValMatrix {
		m := ds.(model.Matrix)
		v := m[0]
		t.Logf("---m:%+v", v.Values[0].Timestamp)
		t.Logf("---m:%+v", v.Values[0].Value.String())

	} else {
		t.Logf("------unkown")
	}
}
