// Copyright 2018 The QOS Authors

package validator

import (
	"time"

	"github.com/QOSGroup/qmoon/db"
	"github.com/QOSGroup/qmoon/db/model"
	"github.com/QOSGroup/qmoon/types"
	"github.com/QOSGroup/qmoon/utils"
	tmctypes "github.com/tendermint/tendermint/rpc/core/types"
	tmtypes "github.com/tendermint/tendermint/types"
)

func convertToValidator(bv *model.Validator) *types.Validator {
	return &types.Validator{
		ChainID:          bv.ChainID.String,
		Address:          bv.Address.String,
		PubKeyType:       bv.PubKeyType.String,
		PubKeyValue:      bv.PubKeyValue.String,
		VotingPower:      bv.VotingPower.Int64,
		Accum:            bv.Accum.Int64,
		FirstBlockHeight: bv.FirstBlockHeight.Int64,
		FirstBlockTime:   bv.FirstBlockTime.Time,
		CreatedAt:        bv.CreatedAt.Time,
	}
}

// ListValidatorByChain 查询链所有的validator
func ListValidatorByChain(chainID string) ([]*types.Validator, error) {
	mvs, err := model.ValidatorsByChainID(db.Db, utils.NullString(chainID))
	if err != nil {
		return nil, err
	}

	var res []*types.Validator
	for _, v := range mvs {
		res = append(res, convertToValidator(v))
	}

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

func SaveValidator(chainId string, v *tmtypes.Vote, vl *tmtypes.Validator) error {
	now := time.Now()
	mv, err := retrieveValidator(v.ValidatorAddress.String())
	if err != nil {
		mv = &model.Validator{
			ChainID:          utils.NullString(chainId),
			Address:          utils.NullString(v.ValidatorAddress.String()),
			PubKeyType:       utils.NullString(""),
			PubKeyValue:      utils.NullString(""),
			VotingPower:      utils.NullInt64(vl.VotingPower),
			Accum:            utils.NullInt64(vl.Accum),
			FirstBlockHeight: utils.NullInt64(v.Height),
			FirstBlockTime:   utils.NullTime(v.Timestamp),
			CreatedAt:        utils.NullTime(now),
		}

		if err := mv.Insert(db.Db); err != nil {
			return err
		}
	} else {
		mv.VotingPower = utils.NullInt64(vl.VotingPower)
		mv.Accum = utils.NullInt64(vl.Accum)
		if err := mv.Update(db.Db); err != nil {
			return err
		}
	}

	return nil
}

func Save(b *tmctypes.ResultBlock, vs *tmctypes.ResultValidators) error {
	c := b.Block.LastCommit
	if c.Precommits == nil || len(c.Precommits) == 0 {
		return nil
	}

	vm := make(map[string]*tmtypes.Validator)
	for _, v := range vs.Validators {
		vm[v.Address.String()] = v
	}

	for _, v := range c.Precommits {
		if v.String() == "nil-Vote" {
			continue
		}

		var validator *tmtypes.Validator
		var ok bool
		if validator, ok = vm[v.ValidatorAddress.String()]; !ok {
			validator = new(tmtypes.Validator)
		}

		SaveValidator(b.Block.ChainID, v, validator)
		SaveBlockValidator(b.Block.ChainID, v, validator)
	}

	return nil
}
