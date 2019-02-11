package models

import (
	"time"

	"github.com/QOSGroup/qmoon/models/errors"
	"github.com/go-xorm/xorm"
)

type AtmIpRecord struct {
	Id      int64  `xorm:"pk autoincr BIGINT"`
	Chainid string `xorm:"-"`
	Ip      string `xorm:"unique(atm_ip_record_ip_created_at_unix) TEXT"`
	Amount  int

	Version int `xorm:"version"`

	UpdatedAt     time.Time `xorm:"-"`
	UpdatedAtUnix int64
	CreatedAt     time.Time `xorm:"-"`
	CreatedAtUnix int64     `xorm:"unique(atm_ip_record_ip_created_at_unix)"`
}

func (air *AtmIpRecord) BeforeInsert() {
	air.CreatedAtUnix = time.Now().Unix()
	air.UpdatedAtUnix = time.Now().Unix()
}

func (air *AtmIpRecord) BeforeUpdate() {
	air.UpdatedAtUnix = time.Now().Unix()
}

func (air *AtmIpRecord) AfterSet(colName string, _ xorm.Cell) {
	switch colName {
	case "created_at_unix":
		air.CreatedAt = time.Unix(air.CreatedAtUnix, 0).Local()
	case "updated_at_unix":
		air.UpdatedAt = time.Unix(air.UpdatedAtUnix, 0).Local()
	}
}

func (air *AtmIpRecord) Insert(chainID string) error {
	x, err := GetNodeEngine(chainID)
	if err != nil {
		return err
	}

	_, err = x.Insert(air)
	if err != nil {
		return err
	}

	return nil
}

func (air *AtmIpRecord) Update(chainID string) error {
	x, err := GetNodeEngine(chainID)
	if err != nil {
		return err
	}

	_, err = x.ID(air.Id).Update(air)
	if err != nil {
		return err
	}

	return nil
}

func RetrieveAtmIpRecordByIpCreateat(chainID string, ip string, t time.Time) (*AtmIpRecord, error) {
	x, err := GetNodeEngine(chainID)
	if err != nil {
		return nil, err
	}

	air := &AtmIpRecord{
		Ip:            ip,
		CreatedAtUnix: t.Unix(),
	}

	has, err := x.Get(air)
	if err != nil {
		return nil, err
	}

	if !has {
		return nil, errors.NotExist{Obj: "AtmIpRecord"}
	}

	return air, nil
}
