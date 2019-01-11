// Copyright 2018 The QOS Authors

package types

import "github.com/tendermint/tendermint/crypto/tmhash"

type TxStatus int

const (
	TxStatusInit TxStatus = iota
	TxStatusSuccess
	TxStatusFaild
)

func (ts TxStatus) MarshalJSON() ([]byte, error) {
	return []byte(ts.String()), nil
}

func (ts TxStatus) String() string {
	switch ts {
	case TxStatusInit:
		return "确认中"
	case TxStatusSuccess:
		return "成功"
	case TxStatusFaild:
		return "失败"
	default:
		return "确认中"
	}
}

type Tx []byte

func (t Tx) Hash() []byte {
	return tmhash.Sum(t)
}
