package types

import "github.com/QOSGroup/qmoon/lib/qos/base/txs"

// TxResponse defines a structure containing relevant tx data and metadata. The
// tags are stringified and the log is JSON decoded.
type TxResponse struct {
	Height    int64           `json:"height"`
	TxHash    string          `json:"txhash"`
	Code      uint32          `json:"code,omitempty"`
	Data      string          `json:"data,omitempty"`
	RawLog    string          `json:"raw_log,omitempty"`
	Logs      ABCIMessageLogs `json:"logs,omitempty"`
	Info      string          `json:"info,omitempty"`
	GasWanted int64           `json:"gas_wanted,omitempty"`
	GasUsed   int64           `json:"gas_used,omitempty"`
	Events    StringEvents    `json:"events,omitempty"`
	Codespace string          `json:"codespace,omitempty"`
	Tx        txs.TxStd       `json:"tx,omitempty"`
	Timestamp string          `json:"timestamp,omitempty"`
}

type ABCIMessageLogs []ABCIMessageLog

// ABCIMessageLog defines a structure containing an indexed tx ABCI message log.
type ABCIMessageLog struct {
	MsgIndex string `json:"msg_index"`
	Success  bool   `json:"success"`
	Log      string `json:"log"`
}

//type Tx struct {
//	Type  string    `json:"type"`
//	Value txs.TxStd `json:"value"`
//}
