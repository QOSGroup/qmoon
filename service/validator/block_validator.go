// Copyright 2018 The QOS Authors

package validator

import (
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/QOSGroup/qmoon/db"
	"github.com/QOSGroup/qmoon/db/model"
	"github.com/QOSGroup/qmoon/types"
	"github.com/QOSGroup/qmoon/utils"
)

const maxLimit = 20

func convertToBlockValidator(bv *model.BlockValidator) *types.BlockValidator {
	return &types.BlockValidator{
		ChainID:          bv.ChainID.String,
		ValidatorAddress: bv.ValidatorAddress.String,
		ValidatorIndex:   bv.ValidatorIndex.Int64,
		Height:           bv.Height.Int64,
		Round:            bv.Round.Int64,
		Type:             bv.Type.Int64,
		Signature:        bv.Signature.String,
		Timestamp:        bv.Time.Time,
		//Accum:            bv.Accum.Int64,
		//CreatedAt:        bv.CreatedAt.Time,
	}
}

// ListBlockValidatorByHeight 查询
func ListBlockValidatorByHeight(chainID string, height int64) ([]*types.BlockValidator, error) {
	const sqlOrder = " order by validator_index "
	var wheres []string
	wheres = append(wheres, fmt.Sprintf(" chain_id = '%s' ", chainID))
	wheres = append(wheres, fmt.Sprintf(" height = %d ", height))

	mbvs, err := model.BlockValidatorFilter(db.Db, strings.Join(wheres, " and "), sqlOrder, 0, -1)
	if err != nil {
		return nil, err
	}

	var res []*types.BlockValidator
	for _, v := range mbvs {
		res = append(res, convertToBlockValidator(v))
	}

	return res, err
}

// ListBlockValidatorByAddress 查询
func ListBlockValidatorByAddress(address string, minHeight, maxHeight int64) ([]*types.BlockValidator, error) {
	const sqlOrder = " order by height desc "
	var wheres []string
	wheres = append(wheres, fmt.Sprintf(" validator_address = '%s' ", address))
	if minHeight != 0 && maxHeight != 0 {
		wheres = append(wheres, fmt.Sprintf(" height >= %d ", minHeight))
		wheres = append(wheres, fmt.Sprintf(" height <= %d ", maxHeight))
	}

	mbvs, err := model.BlockValidatorFilter(db.Db, strings.Join(wheres, " and "), sqlOrder, 0, maxLimit)
	if err != nil {
		return nil, err
	}

	var res []*types.BlockValidator
	for _, v := range mbvs {
		res = append(res, convertToBlockValidator(v))
	}

	return res, err
}

// retrieveBlockValidator 单个查询
func retrieveBlockValidator(chainID string, height int64, validatorAddress string) (*model.BlockValidator, error) {
	var wheres []string
	wheres = append(wheres, fmt.Sprintf(" chain_id = '%s' ", chainID))
	wheres = append(wheres, fmt.Sprintf(" height = %d ", height))
	wheres = append(wheres, fmt.Sprintf(" validator_address = '%s' ", validatorAddress))

	mbvs, err := model.BlockValidatorFilter(db.Db, strings.Join(wheres, " and "), "", 0, -1)
	if err != nil {
		return nil, err
	}

	if len(mbvs) == 0 {
		return nil, errors.New("not found")
	}

	return mbvs[0], nil
}

// Search 单个查询
func RetrieveBlockValidator(chainID string, height int64, validatorAddress string) (*types.BlockValidator, error) {
	mbv, err := retrieveBlockValidator(chainID, height, validatorAddress)
	if err != nil {
		return nil, err
	}

	return convertToBlockValidator(mbv), nil
}

func saveBlockValidator(chainId string, v *types.BlockValidator) error {
	now := time.Now()
	mbv, err := retrieveBlockValidator(chainId, v.Height, v.ValidatorAddress)
	if err != nil {
		mbv = &model.BlockValidator{
			ChainID:          utils.NullString(chainId),
			Height:           utils.NullInt64(v.Height),
			ValidatorAddress: utils.NullString(v.ValidatorAddress),
			ValidatorIndex:   utils.NullInt64(int64(v.ValidatorIndex)),
			Type:             utils.NullInt64(int64(v.Type)),
			Round:            utils.NullInt64(int64(v.Round)),
			Signature:        utils.NullString(v.Signature),
			Time:             utils.NullTime(v.Timestamp),
			VotingPower:      utils.NullInt64(v.VotingPower),
			Accum:            utils.NullInt64(v.Accum),
			CreatedAt:        utils.NullTime(now),
		}

		if err := mbv.Insert(db.Db); err != nil {
			return err
		}
	}

	return nil
}

func SaveBlockValidator(vars []*types.BlockValidator) error {
	vm := make(map[string]*types.BlockValidator)
	for _, v := range vars {
		vm[v.ValidatorAddress] = v
	}

	for _, v := range vars {
		UpdateValidatorFirstBlock(v.ChainID, v.ValidatorAddress, v.Height, v.Timestamp)
		saveBlockValidator(v.ChainID, v)
	}

	return nil
}
