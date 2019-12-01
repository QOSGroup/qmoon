package models

import (
	"github.com/QOSGroup/qmoon/types"
	"time"

	"github.com/go-xorm/xorm"
)

type Missing struct {
	Id               int64  `xorm:"pk autoincr BIGINT"`
	Height           int64  `xorm:"unique(height_val_address_idx) BIGINT"`
	ValidatorAddress string `xorm:"unique(height_val_address_idx) TEXT"`

	CreatedAt     time.Time `xorm:"-"`
	CreatedAtUnix int64
}

func (n *Missing) BeforeInsert() {
	if n.CreatedAt.IsZero() {
		n.CreatedAtUnix = time.Now().Unix()
	} else {
		n.CreatedAtUnix = n.CreatedAt.Unix()
	}
}

func (n *Missing) BeforeUpdate() {
}

func (n *Missing) AfterSet(colName string, _ xorm.Cell) {
	switch colName {
	case "created_at_unix":
		n.CreatedAt = time.Unix(n.CreatedAtUnix, 0).Local()
	}
}

func (n *Missing) Insert(chainID string) error {
	x, err := GetNodeEngine(chainID)
	if err != nil {
		return err
	}

	_, err = x.Insert(n)
	if err != nil {
		return err
	}

	return nil
}

func RetrieveMissingValidators(chainID string, height int64) ([]*types.Validator, error) {
	x, err := GetNodeEngine(chainID)
	if err != nil {
		return nil, err
	}
	//
	var missingValidators = make([]*types.Validator, 0)
	return missingValidators, x.SQL("select v.\"id\", v.\"name\", v.\"details\", v.\"identity\", v.\"logo\", v.\"website\", v.\"owner\", v.\"address\", v.\"pub_key_type\", v.\"pub_key_value\", v.\"commission\", v.\"voting_power\", v.\"accum\", v.\"first_block_height\", v.\"first_block_time_unix\", v.\"status\", v.\"inactive_code\", v.\"inactive_time_unix\", v.\"inactive_height\", v.\"bond_height\", v.\"precommit_num\", v.\"bonded_tokens\", v.\"self_bond\", v.\"stake_address\" from validator v, missing m where m.height = ? and m.validator_address = v.address;", height).Limit(10, 0).Find(&missingValidators)
}

func RetrieveMissings(chainID, validator string) ([]*Missing, error) {
	x, err := GetNodeEngine(chainID)
	if err != nil {
		return nil, err
	}

	var missings = make([]*Missing, 0)
	return missings, x.Where("validator_address = ?", validator).OrderBy("created_at_unix desc").Limit(10, 0).Find(&missings)
}
