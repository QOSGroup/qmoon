// Copyright 2018 The QOS Authors

package types

import (
	"time"
)

type BlockHeader struct {
	ID      int64     `json:"-"`
	ChainID string    `json:"chain_id"`
	Height  int64     `json:"height"`
	Time    time.Time `json:"time"`
	NumTxs  int64     `json:"num_txs"`

	// prev block info
	//LastBlockID BlockID `json:"last_block_id"`
	TotalTxs int64 `json:"total_txs"`

	// hashes of block data
	LastCommitHash string `json:"last_commit_hash"` // commit from validators from the last block
	DataHash       string `json:"data_hash"`        // transactions

	// hashes from the app output from the prev block
	ValidatorsHash  string `json:"validators_hash"`   // validators for the current block
	ConsensusHash   string `json:"consensus_hash"`    // consensus params for current block
	AppHash         string `json:"app_hash"`          // state after txs from the previous block
	LastResultsHash string `json:"last_results_hash"` // root hash of all results from the txs from the previous block

	// consensus info
	EvidenceHash string `json:"evidence_hash"` // evidence included in the block

}

type Block struct {
	Header     BlockHeader
	Txs        []Tx
	Precommits []*BlockValidator
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
	VotingPower      int64     `json:"voting_power"` // voting_power
	Accum            int64     `json:"accum"`        // accum
}
