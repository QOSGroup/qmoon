// Copyright 2018 The QOS Authors

package types

import (
	qostypes "github.com/QOSGroup/qos/types"
	tmtypes "github.com/tendermint/tendermint/types"

	"time"
)

const (
	//Active 可获得挖矿奖励状态
	Active int8 = iota

	//Inactive
	Inactive
)

type Validator struct {
	Name             string                `json:"name"`
	Owner            string                `json:"owner"`
	ChainID          string                `json:"chain_id"`
	Address          string                `json:"address"`
	PubKeyType       string                `json:"pub_key_type"`
	PubKeyValue      string                `json:"pub_key_value"`
	VotingPower      int64                 `json:"voting_power"`
	Accum            int64                 `json:"accum"`
	FirstBlockHeight int64                 `json:"first_block_height"`
	FirstBlockTime   time.Time             `json:"first_block_time"`
	CreatedAt        time.Time             `json:"created_at"`
	Status           int8                  `json:"status"`
	StatusStr        string                `json:"statusStr"`
	InactiveCode     qostypes.InactiveCode `json:"inactiveCode"`
	InactiveTime     time.Time             `json:"inactiveTime"`
	InactiveHeight   int64                 `json:"inactiveHeight"`
	BondHeight       int64                 `json:"bondHeight"`
	Percent          string                `json:"percent"`
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

func ParseQOSValidator(chainID string, val qostypes.Validator) Validator {
	return Validator{
		Name:    val.Name,
		Owner:   val.Owner.String(),
		ChainID: chainID,
		Address: val.ValidatorPubKey.Address().String(),
		//PubKeyType       :"",
		PubKeyValue: string(val.ValidatorPubKey.Bytes()),
		VotingPower: int64(val.BondTokens),
		//Accum            :0,
		//FirstBlockHeight :,
		//FirstBlockTime   time.Time             `json:"first_block_time"`
		//CreatedAt        time.Time             `json:"created_at"`
		Status:         val.Status,
		InactiveCode:   val.InactiveCode,
		InactiveTime:   val.InactiveTime,
		InactiveHeight: int64(val.InactiveHeight),
		BondHeight:     int64(val.BondHeight),
	}
}

func ParseQSCValidator(chainID string, val tmtypes.Validator) Validator {
	return Validator{
		//Name:  val.Name,
		//Owner: val.Owner.String(),
		ChainID: chainID,
		Address: val.Address.String(),
		//PubKeyType       :"",
		PubKeyValue: val.PubKey.Address().String(),
		VotingPower: val.VotingPower,
		Accum:       val.Accum,
		//FirstBlockHeight :,
		//FirstBlockTime   time.Time             `json:"first_block_time"`
		//CreatedAt        time.Time             `json:"created_at"`
		//Status:         val.Status,
		//InactiveCode:   val.InactiveCode,
		//InactiveTime:   val.InactiveTime,
		//InactiveHeight: int64(val.InactiveHeight),
		//BondHeight:     int64(val.BondHeight),
	}
}
