// Copyright 2018 The QOS Authors

package service

import (
	"fmt"
	"sort"
	"time"

	"github.com/QOSGroup/qmoon/lib"
	"github.com/QOSGroup/qmoon/models"
	"github.com/QOSGroup/qmoon/types"
	"github.com/QOSGroup/qmoon/utils"
	qostypes "github.com/QOSGroup/qos/types"
)

func convertToValidator(bv *models.Validator) *types.Validator {
	statusStr := "Active"
	if bv.Status != 0 {
		statusStr = "Inactive"
	}
	return &types.Validator{
		Name:             bv.Name,
		Owner:            bv.Owner,
		ChainID:          bv.ChainId,
		Address:          bv.Address,
		PubKeyType:       bv.PubKeyType,
		PubKeyValue:      bv.PubKeyValue,
		VotingPower:      bv.VotingPower,
		Accum:            bv.Accum,
		FirstBlockHeight: bv.FirstBlockHeight,
		FirstBlockTime:   bv.FirstBlockTime,
		Status:           int8(bv.Status),
		StatusStr:        statusStr,
		InactiveCode:     qostypes.InactiveCode(bv.InactiveCode),
		InactiveTime:     bv.InactiveTime,
		InactiveHeight:   bv.InactiveHeight,
		BondHeight:       bv.BondHeight,
	}
}

// Validators 查询链所有的validator
func (n Node) Validators() (types.Validators, error) {
	mvs, err := models.Validators(n.ChanID)
	if err != nil {
		return nil, err
	}

	var totoal int64
	var res types.Validators
	for _, v := range mvs {
		if int8(v.Status) == types.Active {
			totoal += v.VotingPower
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
func (n Node) retrieveValidator(address string) (*models.Validator, error) {
	mv, err := models.ValidatorByAddress(n.ChanID, address)
	if err != nil {
		return nil, err
	}

	return mv, nil
}

// RetrieveValidator 单个查询
func (n Node) RetrieveValidator(address string) (*types.Validator, error) {
	mv, err := n.retrieveValidator(address)
	if err != nil {
		return nil, err
	}

	return convertToValidator(mv), nil
}

func (n Node) UpdateValidatorFirstBlock(address string, height int64, t time.Time) error {
	mv, err := n.retrieveValidator(address)
	if err != nil {
		mv = &models.Validator{
			Address:          address,
			FirstBlockHeight: height,
			FirstBlockTime:   t,
		}

		if err := mv.Insert(n.ChanID); err != nil {
			return err
		}
	} else {
		if mv.FirstBlockHeight == 0 {
			mv.FirstBlockHeight = height
			mv.FirstBlockTime = t
			if err := mv.Update(n.ChanID); err != nil {
				return err
			}
		}
	}

	return nil
}

func (n Node) InactiveValidator(address string, status int, inactiveHeight int64, inactiveTime time.Time) error {
	mv, err := n.retrieveValidator(address)
	if err == nil {
		if mv.Status != status {
			mv.Status = status
			mv.InactiveCode = int(types.Inactive)
			mv.InactiveTime = inactiveTime
			mv.InactiveHeight = inactiveHeight

			if err := mv.Update(n.ChanID); err != nil {
				return err
			}
		}
	}

	return nil
}

func (n Node) CreateValidator(vl types.Validator) error {
	address := vl.Address
	mv, err := n.retrieveValidator(address)
	if err != nil {
		mv = &models.Validator{
			Address:        address,
			PubKeyType:     "",
			PubKeyValue:    "",
			VotingPower:    vl.VotingPower,
			Accum:          vl.Accum,
			Status:         int(vl.Status),
			InactiveCode:   int(vl.InactiveCode),
			InactiveTime:   vl.InactiveTime,
			InactiveHeight: vl.InactiveHeight,
			BondHeight:     vl.BondHeight,
			Name:           vl.Name,
			Owner:          vl.Owner,
		}

		if err := mv.Insert(n.ChanID); err != nil {
			return err
		}
	} else {
		mv.VotingPower = vl.VotingPower
		mv.Accum = vl.Accum
		mv.Status = int(vl.Status)
		mv.InactiveCode = int(vl.InactiveCode)
		mv.InactiveTime = vl.InactiveTime
		mv.InactiveHeight = vl.InactiveHeight
		mv.BondHeight = vl.BondHeight
		mv.Name = vl.Name
		mv.Owner = vl.Owner
		if err := mv.Update(n.ChanID); err != nil {
			return err
		}
	}

	return nil
}

// GetQOSValidator 同步最新的validator状态
func (n Node) GetQOSValidator(cli *lib.TmClient, height int64, nodeType types.NodeType) ([]types.Validator, error) {
	var result []types.Validator
	var err error
	if nodeType == types.NodeTypeQOS {
		result, err = cli.QOSValidator(height)
	} else {
		result, err = cli.Validator(height)
	}

	return result, err
}
