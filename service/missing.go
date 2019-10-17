// Copyright 2018 The QOS Authors

package service

import (
	"github.com/QOSGroup/qmoon/models"
	"github.com/QOSGroup/qmoon/types"
)

func convertToMissing(mt *models.Missing) *types.ResultMissing {
	return &types.ResultMissing{
		Time:   mt.CreatedAt.Format("2006-01-02 15:04:05"),
		Height: mt.Height,
	}
}

// Missing missing查询
func (n Node) Missings(validator string) ([]*types.ResultMissing, error) {
	var result []*types.ResultMissing
	if validator == "" {
		return result, nil
	}
	mts, err := models.RetrieveMissings(n.ChainID, validator)
	if err != nil {
		return nil, err
	}

	for _, v := range mts {
		result = append(result, convertToMissing(v))
	}

	return result, err
}
