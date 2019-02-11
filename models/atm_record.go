package models

import (
	"time"

	"github.com/QOSGroup/qmoon/models/errors"
	"github.com/go-xorm/xorm"
)

type AtmRecord struct {
	Id       int64     `xorm:"pk autoincr BIGINT"`
	Address  string    `xorm:"unique(atm_record_address_idx)"`
	Chainid  string    `xorm:"-"`
	Coin     string    `xorm:"TEXT"`
	Amount   string    `xorm:"TEXT"`
	Height   string    `xorm:"TEXT"`
	Hash     string    `xorm:"TEXT"`
	Createat time.Time `xorm:"unique(atm_record_address_idx)"`

	Version int `xorm:"version"`

	UpdatedAt     time.Time `xorm:"-"`
	UpdatedAtUnix int64
	CreatedAt     time.Time `xorm:"-"`
	CreatedAtUnix int64
}

func (ar *AtmRecord) BeforeInsert() {
	ar.CreatedAtUnix = time.Now().Unix()
	ar.UpdatedAtUnix = time.Now().Unix()
}

func (ar *AtmRecord) BeforeUpdate() {
	ar.UpdatedAtUnix = time.Now().Unix()
}

func (ar *AtmRecord) AfterSet(colName string, _ xorm.Cell) {
	switch colName {
	case "created_at_unix":
		ar.CreatedAt = time.Unix(ar.CreatedAtUnix, 0).Local()
	case "updated_at_unix":
		ar.UpdatedAt = time.Unix(ar.UpdatedAtUnix, 0).Local()
	}
}

func (ar *AtmRecord) Insert(chainID string) error {
	x, err := GetNodeEngine(chainID)
	if err != nil {
		return err
	}

	_, err = x.Insert(ar)
	if err != nil {
		return err
	}

	return nil
}

func (ar *AtmRecord) Update(chainID string) error {
	x, err := GetNodeEngine(chainID)
	if err != nil {
		return err
	}

	_, err = x.ID(ar.Id).Update(ar)
	if err != nil {
		return err
	}

	return nil
}

func RetrieveAtmRecordByAddressCreateat(chainID string, address string, t time.Time) (*AtmRecord, error) {
	x, err := GetNodeEngine(chainID)
	if err != nil {
		return nil, err
	}

	air := &AtmRecord{
		Address:  address,
		Createat: t,
	}

	has, err := x.Get(air)
	if err != nil {
		return nil, err
	}

	if !has {
		return nil, errors.NotExist{Obj: "AtmRecord"}
	}

	return air, nil
}
