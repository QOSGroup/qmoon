// Copyright 2018 The QOS Authors

package service

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"time"

	"github.com/QOSGroup/qmoon/db"
	"github.com/QOSGroup/qmoon/db/model"
	"github.com/QOSGroup/qmoon/lib"
	"github.com/QOSGroup/qmoon/utils"
)

type Genesis struct {
	ID          int64     `json:"id"`           // id
	ChainID     string    `json:"chain_id"`     // chain_id
	GenesisTime time.Time `json:"genesis_time"` // genesis_time
	Data        string    `json:"data"`         // data
	CreatedAt   time.Time `json:"created_at"`   // created_at
}

func convertToGenesis(mg *model.Genesi) *Genesis {
	return &Genesis{
		ID:          mg.ID,
		ChainID:     mg.ChainID.String,
		GenesisTime: mg.GenesisTime.Time,
		Data:        mg.Data.String,
		CreatedAt:   mg.CreatedAt.Time,
	}
}

func RetrieveGenesis(chainID string) (*Genesis, error) {
	mg, err := model.GenesiByChainID(db.Db, utils.NullString(chainID))
	if err != nil {
		return nil, err
	}

	return convertToGenesis(mg), nil
}

func AddGenesis(remote string) (*model.Genesi, error) {
	tmc := lib.TendermintClient(remote)
	genesis, err := tmc.Genesis()
	if err != nil {
		return nil, fmt.Errorf("retrieve node genesis err:%s", err)
	}

	d, err := json.Marshal(genesis)
	if err != nil {
		return nil, err
	}

	mg, err := model.GenesiByChainID(db.Db, utils.NullString(genesis.Genesis.ChainID))
	if err != nil {
		if err == sql.ErrNoRows {
			mg = &model.Genesi{
				ChainID:     utils.NullString(genesis.Genesis.ChainID),
				GenesisTime: utils.NullTime(genesis.Genesis.GenesisTime),
				Data:        utils.NullString(string(d)),
				CreatedAt:   utils.NullTime(time.Now()),
			}

			err = mg.Insert(db.Db)
			if err != nil {
				return nil, err
			}
		} else {
			return nil, err
		}
	}

	return mg, nil
}
