package types

type ChainNode struct {
	NodeID int64  `json:"nodeId"`
	Delay  int64  `json:"delay"`
	Remote string `json:"remote"`
}
