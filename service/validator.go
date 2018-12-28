// Copyright 2018 The QOS Authors

package service

import (
	"errors"

	"github.com/QOSGroup/qbase/store"
	"github.com/QOSGroup/qmoon/lib"
	"github.com/QOSGroup/qmoon/types"
	qostypes "github.com/QOSGroup/qos/types"
	"github.com/tendermint/tendermint/rpc/client"
)

func buildQueryOptions(height int64, trust bool) client.ABCIQueryOptions {
	if height <= 0 {
		height = 0
	}

	return client.ABCIQueryOptions{
		Height:  height,
		Trusted: trust,
	}
}

var (
	//keys see docs/spec/staking.md
	validatorKey            = []byte{0x01} // 保存Validator信息. key: ValidatorAddress
	validatorByOwnerKey     = []byte{0x02} // 保存Owner与Validator的映射关系. key: OwnerAddress, value : ValidatorAddress
	validatorByInactiveKey  = []byte{0x03} // 保存处于`inactive`状态的Validator. key: ValidatorInactiveTime + ValidatorAddress
	validatorByVotePowerKey = []byte{0x04} // 按VotePower排序的Validator地址,不包含`pending`状态的Validator. key: VotePower + ValidatorAddress

	currentValidatorAddressKey = []byte("currentValidatorAddressKey")
)

// GetQOSValidator 同步最新的validator状态
func GetQOSValidator(cli *client.HTTP, height int64, node Node) ([]types.Validator, error) {
	var result []types.Validator
	if node.NodeType == types.NodeTypeQOS {
		opts := buildQueryOptions(0, true)
		path := "/store/validator/subspace"
		resp, err := cli.ABCIQueryWithOptions(path, validatorKey, opts)
		if err != nil {
			return nil, err
		}

		valueBz := resp.Response.GetValue()
		if len(valueBz) == 0 {
			return nil, errors.New("")
		}

		var vKVPair []store.KVPair
		if err := lib.Cdc.UnmarshalBinary(valueBz, &vKVPair); err != nil {
			return nil, err
		}
		for _, kv := range vKVPair {
			var val qostypes.Validator
			if err := lib.Cdc.UnmarshalBinaryBare(kv.Value, &val); err == nil {
				result = append(result, types.ParseQOSValidator(node.ChanID, val))
			}
		}
	} else {
		vals, err := cli.Validators(&height)
		if err != nil {
			return nil, err
		}

		for _, v := range vals.Validators {
			result = append(result, types.ParseQSCValidator(node.ChanID, *v))
		}
	}

	return result, nil
}
