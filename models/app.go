package models

import (
	"time"

	"github.com/QOSGroup/qmoon/models/errors"
	"github.com/go-xorm/xorm"
)

type App struct {
	Id        int64  `xorm:"pk autoincr BIGINT" json:"id"`
	Name      string `xorm:"VARCHAR(64)" json:"name"`
	SecretKey string `xorm:"unique TEXT" json:"secretKey"`
	Status    int    `xorm:"INTEGER" json:"-"`
	AccountId int64  `xorm:"index BIGINT" json:"accountId"`

	Version   int       `xorm:"version" json:"-"`
	DeletedAt time.Time `xorm:"deleted" json:"-"`

	UpdatedAt     time.Time `xorm:"-"`
	UpdatedAtUnix int64
	CreatedAt     time.Time `xorm:"-"`
	CreatedAtUnix int64
}

func (a *App) BeforeInsert() {
	a.CreatedAtUnix = time.Now().Unix()
	a.UpdatedAtUnix = time.Now().Unix()
}

func (a *App) BeforeUpdate() {
	a.UpdatedAtUnix = time.Now().Unix()
}

func (a *App) AfterSet(colName string, _ xorm.Cell) {
	switch colName {
	case "created_at_unix":
		a.CreatedAt = time.Unix(a.CreatedAtUnix, 0).Local()
	case "updated_at_unix":
		a.UpdatedAt = time.Unix(a.UpdatedAtUnix, 0).Local()
	}
}

func (a *App) Insert() error {
	_, err := basex.Insert(a)
	if err != nil {
		return err
	}

	return nil
}

func (a *App) Account() (*Account, error) {
	acc := &Account{Id: a.AccountId}
	has, err := basex.Get(acc)
	if err != nil {
		return nil, err
	}

	if !has {
		return nil, errors.AccountNotExist{Id: a.AccountId}
	}

	return acc, nil
}

func AppById(id int64) (*App, error) {
	a := &App{
		Id: id,
	}
	has, err := basex.Get(a)
	if err != nil {
		return nil, err
	}

	if !has {
		return nil, errors.AppNotExist{ID: id}
	}

	return a, nil
}

func AppBySecretKey(secretKey string) (*App, error) {
	a := &App{
		SecretKey: secretKey,
	}
	has, err := basex.Get(a)
	if err != nil {
		return nil, err
	}

	if !has {
		return nil, errors.AppNotExist{SecretKey: secretKey}
	}

	return a, nil
}

func AppsByAccount(id int64) ([]*App, error) {
	var apps = make([]*App, 0)
	return apps, basex.Where("account_id = ?", id).Find(&apps)
}
