// Copyright 2018 The QOS Authors

package types

import (
	"time"
)

type Block struct {
	ID             int64     `json:"-"`
	ChainID        string    `json:"chain_id"`
	Height         int64     `json:"height"`
	NumTxs         int64     `json:"num_txs"`
	TotalTxs       int64     `json:"total_txs"`
	DataHash       string    `json:"data"`
	ValidatorsNum  int64     `json:"validatorsNum"`
	ValidatorsHash string    `json:"validators_hash"`
	Time           time.Time `json:"time"`
}

type BlockValidator struct {
	ChainID          string    `json:"chain_id"`
	Height           int64     `json:"height"`
	ValidatorAddress string    `json:"validator_address"`
	ValidatorIndex   int64     `json:"validator_index"`
	Round            int64     `json:"round"`
	Type             int64     `json:"type"`
	Signature        string    `json:"signature"`
	Timestamp        time.Time `json:"timestamp"`
}
