// Copyright 2018 The QOS Authors

package types

import (
	"time"

	"github.com/QOSGroup/qos/module/stake/types"
)

const (
	//Active 可获得挖矿奖励状态
	Active int8 = iota

	//Inactive
	Inactive
)

type Validator struct {
	Name             string             `json:"name"`
	Details          string             `json:"details"`
	Identity         string             `json:"identity"`
	Logo             string             `json:"logo"`
	Website          string             `json:"website"`
	Owner            string             `json:"owner"`
	ChainID          string             `json:"chain_id"`
	Address          string             `json:"address"`
	ConsPubKey       string             `json:"consensus_pubkey"` // the consensus public key of the validator; bech encoded in JSON
	PubKeyType       string             `json:"pub_key_type"`
	PubKeyValue      string             `json:"pub_key_value"`
	VotingPower      int64              `json:"voting_power"`
	Accum            int64              `json:"accum"`
	FirstBlockHeight int64              `json:"first_block_height"`
	FirstBlockTime   time.Time          `json:"first_block_time"`
	CreatedAt        time.Time          `json:"created_at"`
	Status           int8               `json:"status"`
	StatusStr        string             `json:"statusStr"`
	InactiveCode     types.InactiveCode `json:"inactiveCode"`
	InactiveTime     time.Time          `json:"inactiveTime"`
	InactiveHeight   int64              `json:"inactiveHeight"`
	BondHeight       int64              `json:"bondHeight"`
	Percent          string             `json:"percent"`
	PrecommitNum     int64              `json:"precommitNum"`
	Uptime           string             `json:"uptime"`
	UptimeFloat      float64            `json:"-"`
}

type Validators []Validator

func (vals Validators) Len() int {
	return len(vals)
}

func (vals Validators) Less(i, j int) bool {
	if vals[i].Status == Inactive && vals[j].Status == Inactive {
		return vals[i].VotingPower > vals[j].VotingPower
	}

	if vals[i].Status == Inactive {
		return false
	}

	if vals[j].Status == Inactive {
		return true
	}

	return vals[i].VotingPower > vals[j].VotingPower
}

func (vals Validators) Swap(i, j int) {
	vals[i], vals[j] = vals[j], vals[i]
}
