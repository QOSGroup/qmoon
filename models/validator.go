package models

import (
	"time"

	"github.com/QOSGroup/qmoon/models/errors"
	"github.com/go-xorm/xorm"
)

type Validator struct {
	Id                 int64     `xorm:"pk autoincr BIGINT"`
	Name               string    `xorm:"TEXT"`
	Details            string    `xorm:"TEXT"`
	Identity           string    `xorm:"TEXT"`
	Website            string    `xorm:"TEXT"`
	Owner              string    `xorm:"TEXT"`
	ChainId            string    `xorm:"-"`
	Address            string    `xorm:"unique TEXT"`
	PubKeyType         string    `xorm:"TEXT"`
	PubKeyValue        string    `xorm:"TEXT"`
	VotingPower        int64     `xorm:"BIGINT"`
	Accum              int64     `xorm:"BIGINT"`
	FirstBlockHeight   int64     `xorm:"BIGINT"`
	FirstBlockTime     time.Time `xorm:"-"`
	FirstBlockTimeUnix int64
	Status             int       `xorm:"INTEGER"`
	InactiveCode       int       `xorm:"INTEGER"`
	InactiveTime       time.Time `xorm:"-"`
	InactiveTimeUnix   int64
	InactiveHeight     int64 `xorm:"BIGINT"`
	BondHeight         int64 `xorm:"BIGINT"`
	PrecommitNum       int64
}

func (val *Validator) BeforeInsert() {
	if !val.FirstBlockTime.IsZero() {
		val.FirstBlockTimeUnix = val.FirstBlockTime.Unix()
	}
	if !val.InactiveTime.IsZero() {
		val.InactiveTimeUnix = val.InactiveTime.Unix()
	}
}

func (val *Validator) BeforeUpdate() {
	if !val.FirstBlockTime.IsZero() {
		val.FirstBlockTimeUnix = val.FirstBlockTime.Unix()
	}
	if !val.InactiveTime.IsZero() {
		val.InactiveTimeUnix = val.InactiveTime.Unix()
	}
}

func (val *Validator) AfterSet(colName string, _ xorm.Cell) {
	switch colName {
	case "first_block_time_unix":
		val.FirstBlockTime = time.Unix(val.FirstBlockTimeUnix, 0).Local()
	case "inactive_time_unix":
		val.InactiveTime = time.Unix(val.InactiveTimeUnix, 0).Local()
	}
}

func (val *Validator) Insert(chainID string) error {
	x, err := GetNodeEngine(chainID)
	if err != nil {
		return err
	}

	_, err = x.Insert(val)
	if err != nil {
		return err
	}

	return nil
}

func (val *Validator) Update(chainID string) error {
	x, err := GetNodeEngine(chainID)
	if err != nil {
		return err
	}

	_, err = x.ID(val.Id).Update(val)
	if err != nil {
		return err
	}

	return nil
}

func ValidatorByAddress(chainID, address string) (*Validator, error) {
	x, err := GetNodeEngine(chainID)
	if err != nil {
		return nil, err
	}
	val := &Validator{Address: address}
	has, err := x.Get(val)
	if err != nil {
		return nil, err
	}

	if !has {
		return nil, errors.NotExist{Obj: "ValidatorLoop"}
	}

	return val, nil
}

func Validators(chainID string) ([]*Validator, error) {
	x, err := GetNodeEngine(chainID)
	if err != nil {
		return nil, err
	}
	var vals = make([]*Validator, 0)
	return vals, x.Find(&vals)
}
