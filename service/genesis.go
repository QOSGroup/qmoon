// Copyright 2018 The QOS Authors

package service

import (
	"time"

	"github.com/QOSGroup/qmoon/models"
)

type Genesis struct {
	ID          int64     `json:"id"`           // id
	ChainID     string    `json:"chain_id"`     // chain_id
	GenesisTime time.Time `json:"genesis_time"` // genesis_time
	Data        string    `json:"data"`         // data
	CreatedAt   time.Time `json:"created_at"`   // created_at
}

func convertToGenesis(mg *models.Genesis) *Genesis {
	return &Genesis{
		ID:          mg.Id,
		GenesisTime: mg.GenesisTime,
		Data:        mg.Data,
	}
}

func (n Node) RetrieveGenesis(chainID string) (*Genesis, error) {
	mg, err := models.RetrieveGenesis(n.ChanID)
	if err != nil {
		return nil, err
	}

	return convertToGenesis(mg), nil
}
