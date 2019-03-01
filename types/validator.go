// Copyright 2018 The QOS Authors

package types

import (
	"strings"
	"time"

	"github.com/QOSGroup/qos/module/eco/types"
)

const (
	//Active 可获得挖矿奖励状态
	Active int8 = iota

	//Inactive
	Inactive
)

type Validator struct {
	Name             string             `json:"name"`
	Owner            string             `json:"owner"`
	ChainID          string             `json:"chain_id"`
	Address          string             `json:"address"`
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

func ConsensusAddress(nodeType NodeType, addr string) string {
	switch nodeType {
	case NodeTypeQOS:
		return qosConsensusAddress(addr)
	case NodeTypeQSC:
		return qscConsensusAddress(addr)
	case NodeTypeCOSMOS:
		return cosmosConsensusAddress(addr)
	default:
		return addr
	}
}

func qosConsensusAddress(addr string) string {
	return addr
}

const (
	// Bech32PrefixAccAddr defines the Bech32 prefix of an account's address
	COSMOSBech32PrefixAccAddr = "cosmos"
	// Bech32PrefixAccPub defines the Bech32 prefix of an account's public key
	COSMOSBech32PrefixAccPub = "cosmospub"
	// Bech32PrefixValAddr defines the Bech32 prefix of a validator's operator address
	COSMOSBech32PrefixValAddr = "cosmosvaloper"
	// Bech32PrefixValPub defines the Bech32 prefix of a validator's operator public key
	COSMOSBech32PrefixValPub = "cosmosvaloperpub"
	// Bech32PrefixConsAddr defines the Bech32 prefix of a consensus node address
	COSMOSBech32PrefixConsAddr = "cosmosvalcons"
	// Bech32PrefixConsPub defines the Bech32 prefix of a consensus node public key
	COSMOSBech32PrefixConsPub = "cosmosvalconspub"
)

func qscConsensusAddress(addr string) string {
	if strings.HasPrefix(addr, COSMOSBech32PrefixValPub) {

	}
	return addr
}
func cosmosConsensusAddress(addr string) string {
	return addr
}
