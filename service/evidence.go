// Copyright 2018 The QOS Authors

package service

import (
	"github.com/QOSGroup/qmoon/models"
	"github.com/QOSGroup/qmoon/types"
)

func convertToEvidence(mt *models.Evidence) *types.ResultEvidence {
	return &types.ResultEvidence{
		Time:   mt.CreatedAt.Format("2006-01-02 15:04:05"),
		Height: mt.Height,
	}
}

// Evidence evidence查询
func (n Node) Evidences(validator string) ([]*types.ResultEvidence, error) {
	var result []*types.ResultEvidence
	if validator == "" {
		return result, nil
	}
	mts, err := models.RetrieveEvidences(n.ChainID, validator)
	if err != nil {
		return nil, err
	}

	for _, v := range mts {
		result = append(result, convertToEvidence(v))
	}

	return result, err
}
