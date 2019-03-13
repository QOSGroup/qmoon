// Copyright 2018 The QOS Authors

package types

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/QOSGroup/qbase/types"
	"github.com/tidwall/gjson"
)

type ResultTime time.Time

const resultTimeFormart = "2006-01-02 15:04:05.379794+08"

func (rt ResultTime) Time() time.Time {
	return time.Time(rt)
}
func (rt *ResultTime) UnmarshalJSON(data []byte) (err error) {
	now, err := time.ParseInLocation(`"`+resultTimeFormart+`"`, string(data), time.Local)
	*rt = ResultTime(now)
	return
}

func (rt ResultTime) MarshalJSON() ([]byte, error) {
	b := make([]byte, 0, len(resultTimeFormart)+2)
	b = append(b, '"')
	b = time.Time(rt).AppendFormat(b, resultTimeFormart)
	b = append(b, '"')
	return b, nil
}

func (rt ResultTime) String() string {
	return time.Time(rt).Format(resultTimeFormart)
}

type PriKey struct {
	Type  string `json:"type"`
	Value string `json:"value"`
}

type ResultValidator struct {
	Validator *Validator        `json:"validator"`
	Blocks    []*BlockValidator `json:"blocks"`
}

func TxCN(t string, tx string, address string) string {
	switch t {
	//qos
	case "qos/txs/TxTransfer":
		if address != "" {
			if gjson.Get(tx, fmt.Sprintf(`senders.#[addr=="%s"].addr`, address)).String() == "" {
				return "转入"
			} else {
				return "转出"
			}
		} else {
			return "转账"
		}
	//qsc

	// cosmos
	case "send":
		if address != "" {
			if gjson.Get(tx, fmt.Sprintf(`#[data.from_address=="%s"].type`, address)).String() == "send" {
				return "转出"
			} else {
				return "转入"
			}
		} else {
			return "转账"
		}

	case "delegate":
		return "委托"
	case "begin_unbonding":
		return "终止委托"
	case "withdraw_delegator_reward":
		return "取回分红"
	default:
		return t
	}
}

type ResultTx struct {
	ChainID   string          `json:"chain_id"`
	Hash      string          `json:"hash"`
	Height    int64           `json:"height"`
	Index     int64           `json:"index"`   // index
	TxType    string          `json:"tx_type"` // tx_type
	TxTypeCN  string          `json:"tx_type_cn"`
	GasWanted int64           `json:"gas_wanted"`
	GasUsed   int64           `json:"gas_used"`
	Fee       string          `json:"fee"`
	TxStatus  string          `json:"tx_status"`
	Status    int             `json:"status"`
	Data      json.RawMessage `json:"data"` // data
	Time      ResultTime      `json:"time"` // time
	Log       string          `json:"log"`
	CreatedAt ResultTime      `json:"created_at"` // created_at
}

// ResultBlockBase 块信息
type ResultBlockBase struct {
	ID             int64      `json:"-"`
	BlockID        string     `json:"block_id"`
	ChainID        string     `json:"chain_id"`
	Height         int64      `json:"height"`
	NumTxs         int64      `json:"num_txs"`
	TotalTxs       int64      `json:"total_txs"`
	Data           string     `json:"data"`
	Time           ResultTime `json:"time"`
	DataHash       string     `json:"data_hash"`
	ValidatorsHash string     `json:"validators_hash"`
	CreatedAt      ResultTime `json:"-"`
}

type ResultBlockDuration struct {
	Height   int64 `json:"height"`
	Duration int64
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
	Moniker    string     `json:"moniker"`
	ID         int64      `json:"-"`
	PeerID     string     `json:"id"`
	ListenAddr string     `json:"listen_addr"`
	Network    string     `json:"network"`
	Version    string     `json:"version"`
	Channels   string     `json:"channels"`
	SendStart  ResultTime `json:"send_start"`
	RecvStart  ResultTime `json:"recv_start"`
	CreateAt   ResultTime `json:"create_at"`
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
	TotalTxs        int64                 `json:"total_txs"`
	GenesisTime     ResultTime            `json:"genesis_time"`
}

type ResultAccount struct {
	Address string          `json:"address"`
	Nonce   int64           `json:"nonce"`
	Coins   types.BaseCoins `json:"coins"`
}

type ResultMatrix struct {
	X string `json:"x"`
	Y string `json:"y"`
}

type ResultFee struct {
	Tx        string `json:"tx"`
	Fee       string `json:"fee"`
	GasWanted int64  `json:"gasWanted"`
	GasUsed   int64  `json:"gasUsed"`
}

// ResultEvidence 双签返回结果
type ResultEvidence struct {
	Time   string `json:"time"`
	Height int64  `json:"height"`
}

// ResultMissing 漏签
type ResultMissing struct {
	Time   string `json:"time"`
	Height int64  `json:"height"`
}
