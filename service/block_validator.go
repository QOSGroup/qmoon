// Copyright 2018 The QOS Authors

package service

import (
	"errors"
	"fmt"
	"github.com/QOSGroup/qmoon/models"
	"github.com/QOSGroup/qmoon/types"
	"log"
	"strconv"
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
	mbvs, err := models.BlockValidators(n.ChainID, &models.BlockValidatorOption{
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
	mbvs, err := models.BlockValidators(n.ChainID, &models.BlockValidatorOption{
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
	mbvs, err := models.BlockValidators(n.ChainID, &models.BlockValidatorOption{
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
	mbv, err := n.retrieveBlockValidator(v.Height,v.ValidatorAddress)
	if err != nil || mbv == nil {
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

		if err := mbv.Insert(n.ChainID); err != nil {
			return err
		}

		vh := &models.ValidatorHistoryRecord{
				Height: v.Height,
				Address:v.ValidatorAddress,
				VotingPower: v.VotingPower,
				Status:0,
		}
		vh.Insert(n.ChainID)

		// history(votingpower_percent+uptime) reserve for 4 month, thus 17280 * 30 * 4 = 2073600. over 2073600 blocks old, will be purged
		// for blocks more than one week, we only need one block every 720 blocks, in between blocks will be purged. thus 2880. for 21 validators, 60480
		// for blocks less than one week, we reserve every one of them, thus 17280 * 7 * 21 = 2540160, quite a lot.
		if v.Height % (17280 * 30 * 3) == 0 {
			err := models.PurgeOldValidatorHistory(n.ChainID, " where height < " + strconv.FormatInt(v.Height - 2073600, 10))
			if err != nil {
				fmt.Println("Purge failed at ", v.Height, err)
			}
		}
		if v.Height % (17280 * 7) == 0 {
			err = models.PurgeOldValidatorHistory(n.ChainID, " where height < "+strconv.FormatInt(v.Height - 120960, 10)+" and height % 720 != 0")
			if err != nil {
				fmt.Println("Purge failed at ", v.Height, err)
			}
		}
		// for test, purge it 4 hours old
		//if v.Height % 2880 == 0 {
		//	err := models.PurgeOldValidatorHistory(n.ChainID, " where height < " + strconv.FormatInt(v.Height - 2880, 10))
		//	if err != nil {
		//		fmt.Println("Purge failed at ", v.Height, err)
		//	}
		//}
		//if v.Height % (720) == 0 {
		//	err = models.PurgeOldValidatorHistory(n.ChainID, " where height < "+strconv.FormatInt(v.Height - 720, 10)+" and height % 12 != 0")
		//	if err != nil {
		//		fmt.Println("Purge failed at ", v.Height, err)
		//	}
		//}
	}

	return nil
}

func (n Node) SaveBlockValidator(vars []*types.BlockValidator) error {
	if len(vars) <= 0 {
		return nil
	}
	vm := make(map[string]*types.BlockValidator)

	for _, v := range vars {
		vm[v.ValidatorAddress] = v
	}

	height := vars[0].Height
	t := vars[0].Timestamp

	allVals, _ := n.Validators(0)
	total, err := models.TotalVotingPower(n.ChainID)
	if err!=nil || total==0{
		return err
	}
	validatorsInDB:=make(map[string]*types.Validator)
	for _, v := range allVals {
		validatorsInDB[v.Address] = &v
		if _, ok := vm[v.Address]; !ok {
			missing := &models.Missing{
				Height:           height,
				ValidatorAddress: v.Address,
				CreatedAt:        t,
			}
			missing.Insert(n.ChainID)

			vh := &models.ValidatorHistoryRecord{
					Height: height,
					Address: v.Address,
					VotingPower: v.VotingPower,
					TotalPower: total,
					Status: 1,
			}
			err = vh.Insert(n.ChainID)
			if err!= nil {
				fmt.Println("error insert block validator:", vh)
			}
		} else {

			vh := &models.ValidatorHistoryRecord{
				Height: height,
				Address: v.Address,
				VotingPower: v.VotingPower,
				TotalPower: total,
				Status: 0,
			}
			err = vh.Insert(n.ChainID)
			if err!= nil {
				fmt.Println("error insert block validator:", vh)
			}
		}
	}

	for _, v := range vars {
		if _, ok := validatorsInDB[v.ValidatorAddress]; ok {
			v.VotingPower = validatorsInDB[v.ValidatorAddress].VotingPower
		}
		if err := n.saveBlockValidator(v); err != nil {
			log.Printf("saveBlockValidator err:%v", err.Error())
		}
	}

	return nil
}
