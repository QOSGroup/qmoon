// Copyright 2018 The QOS Authors

package validator

import (
	"fmt"
	"sort"
	"time"

	"github.com/QOSGroup/qmoon/db"
	"github.com/QOSGroup/qmoon/db/model"
	"github.com/QOSGroup/qmoon/types"
	"github.com/QOSGroup/qmoon/utils"
	qostypes "github.com/QOSGroup/qos/types"
)

func convertToValidator(bv *model.Validator) *types.Validator {
	statusStr := "Active"
	if bv.Status.Int64 != 0 {
		statusStr = "Inactive"
	}
	return &types.Validator{
		Name:             bv.Name.String,
		Owner:            bv.Owner.String,
		ChainID:          bv.ChainID.String,
		Address:          bv.Address.String,
		PubKeyType:       bv.PubKeyType.String,
		PubKeyValue:      bv.PubKeyValue.String,
		VotingPower:      bv.VotingPower.Int64,
		Accum:            bv.Accum.Int64,
		FirstBlockHeight: bv.FirstBlockHeight.Int64,
		FirstBlockTime:   bv.FirstBlockTime.Time,
		CreatedAt:        bv.CreatedAt.Time,
		Status:           int8(bv.Status.Int64),
		StatusStr:        statusStr,
		InactiveCode:     qostypes.InactiveCode(bv.InactiveCode.Int64),
		InactiveTime:     bv.InactiveTime.Time,
		InactiveHeight:   bv.InactiveHeight.Int64,
		BondHeight:       bv.BondHeight.Int64,
	}
}

// ListValidatorByChain 查询链所有的validator
func ListValidatorByChain(chainID string) (types.Validators, error) {
	mvs, err := model.ValidatorsByChainID(db.Db, utils.NullString(chainID))
	if err != nil {
		return nil, err
	}

	var totoal int64
	var res types.Validators
	for _, v := range mvs {
		if int8(v.Status.Int64) == types.Active {
			totoal += v.VotingPower.Int64
		}
		res = append(res, *convertToValidator(v))
	}

	for i := 0; i < len(res); i++ {
		if res[i].Status == types.Active {
			res[i].Percent = fmt.Sprintf("%.5f", utils.Percent(uint64(res[i].VotingPower), uint64(totoal))*100)
		} else {
			res[i].Percent = "0"
		}
	}

	sort.Sort(res)

	return res, err
}

// retrieveValidator 单个查询
func retrieveValidator(address string) (*model.Validator, error) {
	mv, err := model.ValidatorByAddress(db.Db, utils.NullString(address))
	if err != nil {
		return nil, err
	}

	return mv, nil
}

// RetrieveValidator 单个查询
func RetrieveValidator(address string) (*types.Validator, error) {
	mv, err := retrieveValidator(address)
	if err != nil {
		return nil, err
	}

	return convertToValidator(mv), nil
}

func UpdateValidatorFirstBlock(chainId, address string, height int64, t time.Time) error {
	mv, err := retrieveValidator(address)
	if err != nil {
		mv = &model.Validator{
			ChainID:          utils.NullString(chainId),
			Address:          utils.NullString(address),
			FirstBlockHeight: utils.NullInt64(height),
			FirstBlockTime:   utils.NullTime(t),
			CreatedAt:        utils.NullTime(time.Now()),
		}

		if err := mv.Insert(db.Db); err != nil {
			return err
		}
	} else {
		if mv.FirstBlockHeight.Int64 == 0 {
			mv.FirstBlockHeight = utils.NullInt64(height)
			mv.FirstBlockTime = utils.NullTime(t)
			if err := mv.Update(db.Db); err != nil {
				return err
			}
		}
	}

	return nil
}

func InactiveValidator(address string, status, inactiveCode, inactiveHeight int64, inactiveTime time.Time) error {
	mv, err := retrieveValidator(address)
	if err == nil {
		if mv.Status.Int64 != status {
			mv.Status = utils.NullInt64(status)
			mv.InactiveCode = utils.NullInt64(inactiveCode)
			mv.InactiveTime = utils.NullTime(inactiveTime)
			mv.InactiveHeight = utils.NullInt64(inactiveHeight)

			if err := mv.Update(db.Db); err != nil {
				return err
			}
		}
	}

	return nil
}

func CreateValidator(chainId string, vl types.Validator) error {
	now := time.Now()
	address := vl.Address
	mv, err := retrieveValidator(address)
	if err != nil {
		mv = &model.Validator{
			ChainID:        utils.NullString(chainId),
			Address:        utils.NullString(address),
			PubKeyType:     utils.NullString(""),
			PubKeyValue:    utils.NullString(""),
			VotingPower:    utils.NullInt64(vl.VotingPower),
			Accum:          utils.NullInt64(vl.Accum),
			CreatedAt:      utils.NullTime(now),
			Status:         utils.NullInt64(int64(vl.Status)),
			InactiveCode:   utils.NullInt64(int64(vl.InactiveCode)),
			InactiveTime:   utils.NullTime(vl.InactiveTime),
			InactiveHeight: utils.NullInt64(vl.InactiveHeight),
			BondHeight:     utils.NullInt64(vl.BondHeight),
			Name:           utils.NullString(vl.Name),
			Owner:          utils.NullString(vl.Owner),
		}

		if err := mv.Insert(db.Db); err != nil {
			return err
		}
	} else {
		mv.VotingPower = utils.NullInt64(vl.VotingPower)
		mv.Accum = utils.NullInt64(vl.Accum)
		mv.Status = utils.NullInt64(int64(vl.Status))
		mv.InactiveCode = utils.NullInt64(int64(vl.InactiveCode))
		mv.InactiveTime = utils.NullTime(vl.InactiveTime)
		mv.InactiveHeight = utils.NullInt64(vl.InactiveHeight)
		mv.BondHeight = utils.NullInt64(vl.BondHeight)
		mv.Name = utils.NullString(vl.Name)
		mv.Owner = utils.NullString(vl.Owner)
		if err := mv.Update(db.Db); err != nil {
			return err
		}
	}

	return nil
}
