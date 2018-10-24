// Copyright 2018 The QOS Authors

// Package pkg comments for pkg block
// block 块相关数据封装
package block

import (
	"github.com/QOSGroup/qmoon/lib/tmcli"
	tmtypes "github.com/tendermint/tendermint/rpc/core/types"
)

// Blockchain 链查询
func Blockchain(minHeight, maxHeight int64) (*tmtypes.ResultBlockchainInfo, error) {
	opt, err := tmcli.NewOption()
	if err != nil {
		return nil, err
	}

	res, err := tmcli.NewClient(opt).Blockchain.List(nil, &tmcli.BlockchainOption{MaxHeight: maxHeight, MinHeight: minHeight})

	return res, err
}
