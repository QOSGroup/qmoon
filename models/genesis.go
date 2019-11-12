package models

import (
	"time"

	"github.com/QOSGroup/qmoon/models/errors"
	"github.com/go-xorm/xorm"
)

type Genesis struct {
	Id              int64     `xorm:"pk autoincr BIGINT"`
	ChainId         string    `xorm:"-"`
	GenesisTime     time.Time `xorm:"-"`
	GenesisTimeUnix int64
	Data            string `xorm:"TEXT"`
}

func (g *Genesis) BeforeInsert() {
	g.GenesisTimeUnix = g.GenesisTime.Unix()
}

func (g *Genesis) BeforeUpdate() {
	g.GenesisTimeUnix = g.GenesisTime.Unix()
}

func (g *Genesis) AfterSet(colName string, _ xorm.Cell) {
	switch colName {
	case "genesis_time_unix":
		g.GenesisTime = time.Unix(g.GenesisTimeUnix, 0).Local()
	}
}

func (g *Genesis) Insert(chainID string) error {
	x, err := GetNodeEngine(chainID)
	if err != nil {
		return err
	}

	_, err = x.Insert(g)
	if err != nil {
		return err
	}

	return nil
}

func RetrieveGenesis(chainId string) (*Genesis, error) {
	x, err := GetNodeEngine(chainId)
	if err != nil {
		return nil, err
	}

	n := &Genesis{}
	has, err := x.Get(n)
	if err != nil {
		return nil, err
	}

	if !has {
		return nil, errors.NotExist{Obj: "Genesis"}
	}

	return n, nil
}

func CreateGenesis(chainId string, genesisTime time.Time, data string) error {
	x, err := GetNodeEngine(chainId)
	if err != nil {
		return err
	}
	_, err = x.Insert(&Genesis{
		ChainId:     chainId,
		GenesisTime: genesisTime,
		Data:        data,
	})

	return err
}
