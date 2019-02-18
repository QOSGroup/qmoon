// Copyright 2018 The QOS Authors

package metric

import (
	"testing"
	"time"

	"github.com/QOSGroup/qmoon/types"
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
