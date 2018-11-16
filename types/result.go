// Copyright 2018 The QOS Authors

package types

import (
	"encoding/json"
	"time"
)

type PriKey struct {
	Type  string `json:"type"`
	Value string `json:"value"`
}

type Validator struct {
	ChainID          string    `json:"chain_id"`
	Address          string    `json:"address"`
	PubKeyType       string    `json:"pub_key_type"`
	PubKeyValue      string    `json:"pub_key_value"`
	VotingPower      int64     `json:"voting_power"`
	Accum            int64     `json:"accum"`
	FirstBlockHeight int64     `json:"first_block_height"`
	FirstBlockTime   time.Time `json:"first_block_time"`
	CreatedAt        time.Time `json:"created_at"`
}

type ResultValidator struct {
	Validator *Validator        `json:"validator"`
	Blocks    []*BlockValidator `json:"blocks"`
}

type BlockValidator struct {
	ChainID          string    `json:"chain_id"`
	ValidatorAddress string    `json:"validator_address"`
	ValidatorIndex   int64     `json:"validator_index"`
	Height           int64     `json:"height"`
	Round            int64     `json:"round"`
	Type             int64     `json:"type"`
	Signature        string    `json:"signature"`
	Timestamp        time.Time `json:"timestamp"`
	Accum            int64     `json:"accum"`
	CreatedAt        time.Time `json:"created_at"`
}

type ResultTx struct {
	ChainID     string          `json:"chain_id"`
	Height      int64           `json:"height"`
	Index       int64           `json:"index"`        // index
	TxType      string          `json:"tx_type"`      // tx_type
	Maxgas      int64           `json:"maxgas"`       // maxgas
	QcpFrom     string          `json:"qcp_from"`     // qcp_from
	QcpTo       string          `json:"qcp_to"`       // qcp_to
	QcpSequence int64           `json:"qcp_sequence"` // qcp_sequence
	QcpTxindex  int64           `json:"qcp_txindex"`  // qcp_txindex
	QcpIsresult bool            `json:"qcp_isresult"` // qcp_isresult
	Data        json.RawMessage `json:"data"`         // data
	Time        time.Time       `json:"time"`         // time
	CreatedAt   time.Time       `json:"created_at"`   // created_at
}

// ResultBlockBase 块信息
type ResultBlockBase struct {
	ID             int64     `json:"-"`
	BlockID        string    `json:"block_id"`
	ChainID        string    `json:"chain_id"`
	Height         int64     `json:"height"`
	NumTxs         int64     `json:"num_txs"`
	TotalTxs       int64     `json:"total_txs"`
	Data           string    `json:"data"`
	Time           time.Time `json:"time"`
	DataHash       string    `json:"data_hash"`
	ValidatorsHash string    `json:"validators_hash"`
	CreatedAt      time.Time `json:"-"`
}

type ResultBlock struct {
	Block      *ResultBlockBase  `json:"block"`
	Txs        []*ResultTx       `json:"txs"`
	Validators []*BlockValidator `json:"validators"`
}

type Sequence struct {
	Name string `json:"name"`
	In   int64  `json:"in"`
	Out  int64  `json:"out"`
}

type ResultSequence struct {
	Apps []*Sequence `json:"apps"`
}

type ResultPeer struct {
	Moniker    string    `json:"moniker"`
	ID         int64     `json:"-"`
	PeerID     string    `json:"id"`
	ListenAddr string    `json:"listen_addr"`
	Network    string    `json:"network"`
	Version    string    `json:"version"`
	Channels   string    `json:"channels"`
	SendStart  time.Time `json:"send_start"`
	RecvStart  time.Time `json:"recv_start"`
	CreateAt   time.Time `json:"create_at"`
}

type ResultPeers struct {
	NPeers int64         `json:"n_peers"`
	Peers  []*ResultPeer `json:"peers"`
}

type ResultTxs struct {
	Txs []*ResultTx `json:"txs"`
}

type ResultConsensusState struct {
	ChainID         string `json:"chain_id"`         // chain_id
	Height          string `json:"height"`           // height
	Round           string `json:"round"`            // round
	Step            string `json:"step"`             // step
	PrevotesNum     int64  `json:"prevotes_num"`     // prevotes_num
	PrevotesValue   string `json:"prevotes_value"`   // prevotes_value
	PrecommitsNum   int64  `json:"precommits_num"`   // precommits_num
	PrecommitsValue string `json:"precommits_value"` // precommits_value
	StartTime       string `json:"start_time"`       // start_time
}

type ResultStatus struct {
	ConsensusState  *ResultConsensusState `json:"consensus_state"`
	TotalValidators int64                 `json:"total_validators"`
}
