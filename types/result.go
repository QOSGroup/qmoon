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

type ResultValidator struct {
	Address     string `json:"address"`
	PubKey      PriKey `json:"pub_key"`
	VotingPower string `json:"voting_power"`
	Accum       int64  `json:"accum"`
	Seqin       int64  `json:"seqin"`
	Seqout      int64  `json:"seqout"`
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
	Data           string    `json:"data"`
	Time           time.Time `json:"time"`
	DataHash       string    `json:"data_hash"`
	ValidatorsHash string    `json:"validators_hash"`
	CreatedAt      time.Time `json:"-"`
}

type ResultBlock struct {
	ResultBlockBase

	Txs        []*ResultTx        `json:"txs"`
	Validators []*ResultValidator `json:"validators"`
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
