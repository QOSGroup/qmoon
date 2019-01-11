// Copyright 2018 The QOS Authors

package service

import (
	"github.com/QOSGroup/qmoon/lib"
	"github.com/QOSGroup/qmoon/types"
)

// GetQOSValidator 同步最新的validator状态
func GetQOSValidator(cli *lib.TmClient, height int64, nodeType types.NodeType) ([]types.Validator, error) {
	var result []types.Validator
	var err error
	if nodeType == types.NodeTypeQOS {
		result, err = cli.QOSValidator(height)
	} else {
		result, err = cli.Validator(height)
	}

	return result, err
}
