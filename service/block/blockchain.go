// Copyright 2018 The QOS Authors

// Package pkg comments for pkg block
// block 块相关数据封装
package block

import (
	"github.com/QOSGroup/qmoon/lib/qstartscli"
	tmtypes "github.com/tendermint/tendermint/rpc/core/types"
)

// Blockchain 链查询
func Blockchain(minHeight, maxHeight int64) (*tmtypes.ResultBlockchainInfo, error) {
	opt, err := qstatcli.NewOption()
	if err != nil {
		return nil, err
	}

	_ := qstatcli.NewClient(opt)

	return nil, nil
}
