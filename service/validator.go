// Copyright 2018 The QOS Authors

package service

import (
	"fmt"
	"sort"
	"time"

	"github.com/QOSGroup/qmoon/models"
	"github.com/QOSGroup/qmoon/types"
	"github.com/QOSGroup/qmoon/utils"
	qostypes "github.com/QOSGroup/qos/module/stake/types"
)

func convertToValidator(bv *models.Validator, latestHeight int64) *types.Validator {
	statusStr := "Active"
	if bv.Status != 0 {
		statusStr = "Inactive"
	}

	uptime := float64(bv.PrecommitNum*10000/(latestHeight-bv.FirstBlockHeight)) / 100.00

	return &types.Validator{
		Name:             bv.Name,
		Identity:         bv.Identity,
		Logo:             bv.Logo,
		Website:          bv.Website,
		Details:          bv.Details,
		Owner:            bv.Owner,
		ChainID:          bv.ChainId,
		Address:          bv.Address,
		ConsPubKey:       "",
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
		PrecommitNum:     bv.PrecommitNum,
		Uptime:           fmt.Sprintf("%.2f%%", uptime),
		UptimeFloat:      uptime,
	}
}

// Validators 查询链所有的validator
func (n Node) Validators() (types.Validators, error) {
	latest, err := n.LatestBlock()
	if err != nil {
		return nil, err
	}
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
		res = append(res, *convertToValidator(v, latest.Height))
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
	latest, err := n.LatestBlock()
	if err != nil {
		return nil, err
	}

	mv, err := n.retrieveValidator(address)
	if err != nil {
		return nil, err
	}

	return convertToValidator(mv, latest.Height), nil
}

func (n Node) UpdateValidatorBlock(address string, height int64, t time.Time) error {
	mv, err := n.retrieveValidator(address)
	if err != nil {
		mv = &models.Validator{
			Address:          address,
			FirstBlockHeight: height,
			FirstBlockTime:   t,
			PrecommitNum:     1,
		}

		if err := mv.Insert(n.ChanID); err != nil {
			return err
		}
	} else {
		if mv.FirstBlockHeight == 0 {
			mv.FirstBlockHeight = height
			mv.FirstBlockTime = t
			mv.PrecommitNum = 1
		} else {
			mv.PrecommitNum = mv.PrecommitNum + 1
		}

		if err := mv.Update(n.ChanID); err != nil {
			return err
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

			if err := mv.UpdateStatus(n.ChanID); err != nil {
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
			PubKeyType:     vl.PubKeyType,
			PubKeyValue:    vl.PubKeyValue,
			VotingPower:    vl.VotingPower,
			Accum:          vl.Accum,
			Status:         int(vl.Status),
			InactiveCode:   int(vl.InactiveCode),
			InactiveTime:   vl.InactiveTime,
			InactiveHeight: vl.InactiveHeight,
			BondHeight:     vl.BondHeight,
			Name:           vl.Name,
			Details:        vl.Details,
			Identity:       vl.Identity,
			Logo:           vl.Logo,
			Website:        vl.Website,
			Owner:          vl.Owner,
		}

		if err := mv.Insert(n.ChanID); err != nil {
			return err
		}
	} else {
		mv.PubKeyType = vl.PubKeyType
		mv.PubKeyValue = vl.PubKeyValue
		mv.VotingPower = vl.VotingPower
		mv.Accum = vl.Accum
		mv.Status = int(vl.Status)
		mv.InactiveCode = int(vl.InactiveCode)
		mv.InactiveTime = vl.InactiveTime
		mv.InactiveHeight = vl.InactiveHeight
		mv.BondHeight = vl.BondHeight
		mv.Name = vl.Name
		mv.Logo = vl.Logo
		mv.Details = vl.Details
		mv.Identity = vl.Identity
		mv.Website = vl.Website
		mv.Owner = vl.Owner
		if err := mv.Update(n.ChanID); err != nil {
			return err
		}
	}

	return nil
}
