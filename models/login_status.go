package models

import (
	"time"

	"github.com/QOSGroup/qmoon/models/errors"
	"github.com/go-xorm/xorm"
)

type LoginStatus struct {
	Id            int64  `xorm:"pk autoincr BIGINT"`
	AccountId     int64  `xorm:"unique BIGINT"`
	LoginType     int    `xorm:"default 0 SMALLINT"`
	Token         string `xorm:"unique VARCHAR(64)"`
	ExpiredAtUnix int64

	Version int `xorm:"version"`

	UpdatedAt     time.Time `xorm:"-"`
	UpdatedAtUnix int64
	CreatedAt     time.Time `xorm:"-"`
	CreatedAtUnix int64
}

func (ls *LoginStatus) BeforeInsert() {
	ls.CreatedAtUnix = time.Now().Unix()
	ls.UpdatedAtUnix = time.Now().Unix()
}

func (ls *LoginStatus) BeforeUpdate() {
	ls.UpdatedAtUnix = time.Now().Unix()
}

func (ls *LoginStatus) AfterSet(colName string, _ xorm.Cell) {
	switch colName {
	case "created_at_unix":
		ls.CreatedAt = time.Unix(ls.CreatedAtUnix, 0).Local()
	case "updated_at_unix":
		ls.UpdatedAt = time.Unix(ls.UpdatedAtUnix, 0).Local()
	}
}

func (ls *LoginStatus) Insert(chainID string) error {
	_, err := basex.Insert(ls)
	if err != nil {
		return err
	}

	return nil
}

func (ls *LoginStatus) Account() (*Account, error) {
	acc := &Account{Id: ls.AccountId}
	has, err := basex.Get(acc)
	if err != nil {
		return nil, err
	}

	if !has {
		return nil, errors.AccountNotExist{Id: ls.AccountId}
	}

	return acc, nil
}

func LoginStatusByAccountID(id int64) (*LoginStatus, error) {
	ls := &LoginStatus{AccountId: id}
	has, err := basex.Get(ls)
	if err != nil {
		return nil, err
	}

	if !has {
		return nil, errors.NotExist{Obj: "LoginStatus"}
	}

	return ls, nil
}

func LoginStatusByLoginStatusByToken(token string) (*LoginStatus, error) {
	ls := &LoginStatus{Token: token}
	has, err := basex.Get(ls)
	if err != nil {
		return nil, err
	}

	if !has {
		return nil, errors.NotExist{Obj: "LoginStatus"}
	}

	return ls, nil
}
