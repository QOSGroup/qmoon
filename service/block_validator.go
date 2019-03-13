// Copyright 2018 The QOS Authors

package service

import (
	"errors"
	"log"
	"time"

	"github.com/QOSGroup/qmoon/models"
	"github.com/QOSGroup/qmoon/types"
)

func convertToBlockValidator(bv *models.BlockValidator) *types.BlockValidator {
	return &types.BlockValidator{
		ChainID:          bv.ChainId,
		ValidatorAddress: bv.ValidatorAddress,
		ValidatorIndex:   bv.ValidatorIndex,
		Height:           bv.Height,
		Round:            bv.Round,
		Type:             bv.Type,
		Signature:        bv.Signature,
		Timestamp:        bv.Time,
		//Accum:            bv.Accum,
		//CreatedAt:        bv.CreatedAt.Time,
	}
}

// ListBlockValidatorByHeight 查询
func (n Node) BlockValidatorsByHeight(height int64) ([]*types.BlockValidator, error) {
	mbvs, err := models.BlockValidators(n.ChanID, &models.BlockValidatorOption{
		Height: height,
	})
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
func (n Node) BlockValidatorByAddress(address string, minHeight, maxHeight int64) ([]*types.BlockValidator, error) {
	mbvs, err := models.BlockValidators(n.ChanID, &models.BlockValidatorOption{
		MinHeight: maxHeight, MaxHeight: maxHeight, ValidatorAddress: address,
	})
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
func (n Node) retrieveBlockValidator(height int64, validatorAddress string) (*models.BlockValidator, error) {
	mbvs, err := models.BlockValidators(n.ChanID, &models.BlockValidatorOption{
		Height:           height,
		ValidatorAddress: validatorAddress,
	})
	if err != nil {
		return nil, err
	}

	if len(mbvs) == 0 {
		return nil, errors.New("not found")
	}

	return mbvs[0], nil
}

// Search 单个查询
func (n Node) RetrieveBlockValidator(height int64, validatorAddress string) (*types.BlockValidator, error) {
	mbv, err := n.retrieveBlockValidator(height, validatorAddress)
	if err != nil {
		return nil, err
	}

	return convertToBlockValidator(mbv), nil
}

func (n Node) saveBlockValidator(v *types.BlockValidator) error {
	mbv, err := n.retrieveBlockValidator(v.Height, v.ValidatorAddress)
	if err != nil {
		mbv = &models.BlockValidator{
			Height:           v.Height,
			ValidatorAddress: v.ValidatorAddress,
			ValidatorIndex:   int64(v.ValidatorIndex),
			Type:             int64(v.Type),
			Round:            int64(v.Round),
			Signature:        v.Signature,
			Time:             v.Timestamp,
			VotingPower:      v.VotingPower,
			Accum:            v.Accum,
		}

		if err := mbv.Insert(n.ChanID); err != nil {
			return err
		}
	}

	return nil
}

func (n Node) SaveBlockValidator(vars []*types.BlockValidator) error {
	vm := make(map[string]*types.BlockValidator)
	for _, v := range vars {
		vm[v.ValidatorAddress] = v
	}

	var height int64
	var t time.Time
	for _, v := range vars {
		height = v.Height
		t = v.Timestamp
		if err := n.UpdateValidatorBlock(v.ValidatorAddress, v.Height, v.Timestamp); err != nil {
			log.Printf("UpdateValidatorBlock err:%v", err.Error())
		}

		if err := n.saveBlockValidator(v); err != nil {
			log.Printf("saveBlockValidator err:%v", err.Error())
		}
	}
	allVals, _ := n.Validators()
	for _, v := range allVals {
		if _, ok := vm[v.Address]; !ok {
			missing := &models.Missing{
				Height:           height,
				ValidatorAddress: v.Address,
				CreatedAt:        t,
			}
			missing.Insert(n.ChanID)
		}
	}

	return nil
}
