// Copyright 2018 The QOS Authors

// Package pkg comments for pkg account
// account 管理后台相关功能
package account

import (
	"time"

	"github.com/QOSGroup/qmoon/models"
	"github.com/QOSGroup/qmoon/models/errors"
)

type Account struct {
	ID          int64     `json:"id"`
	Mail        string    `json:"mail"`        // mail
	Name        string    `json:"name"`        // name
	Avatar      string    `json:"avatar"`      // avatar
	Description string    `json:"description"` // description
	Status      int       `json:"status"`      // status
	CreatedAt   time.Time `json:"created_at"`  // created_at

	Password string `json:"-"`
}

type App struct {
	ID        int64  `json:"id"`
	Name      string `json:"name"`
	SecretKey string `json:"secretKey"`
	Status    int    `json:"desc"`
	AccountID int64  `json:"accountId"`
}

func covertToAccount(ma *models.Account) *Account {
	return &Account{
		ID:          ma.Id,
		Mail:        ma.Mail,
		Name:        ma.Name,
		Avatar:      ma.Avatar,
		Description: ma.Description,
		Status:      ma.Status,
		CreatedAt:   ma.CreatedAt,
		Password:    ma.Password,
	}
}

func covertToApp(app *models.App) *App {
	return &App{
		ID:        app.Id,
		Name:      app.Name,
		SecretKey: app.SecretKey,
		Status:    app.Status,
		AccountID: app.AccountId,
	}
}

func List(offset, limit int) ([]*Account, error) {
	maccs, err := models.Accounts(&models.AccountOption{Offset: offset, Limit: limit})
	if err != nil {
		return nil, err
	}

	var as []*Account
	for _, v := range maccs {
		as = append(as, covertToAccount(v))
	}

	return as, nil
}

func CreateAccount(mail, password string) (*Account, error) {
	exist, err := models.AccountIsExist(mail)
	if err != nil {
		return nil, err
	}

	if exist {
		return nil, errors.New("邮箱已经被注册")
	}

	macc, err := models.CreateAccount(mail, password)
	if err != nil {
		return nil, err
	}

	return covertToAccount(macc), nil
}

func RetrieveAccountByID(id int64) (*Account, error) {
	macc, err := models.RetrieveAccountByID(id)
	if err != nil {
		return nil, err
	}

	return covertToAccount(macc), nil
}
func RetrieveAccountByMail(mail string) (*Account, error) {
	macc, err := models.RetrieveAccountByMail(mail)
	if err != nil {
		return nil, err
	}

	return covertToAccount(macc), nil
}

type Option struct {
	Name        string
	Avatar      string
	Description string
}

func UpdateAccountProfile(id int64, opt *Option) error {
	acc, err := models.RetrieveAccountByID(id)
	if err != nil {
		return err
	}

	return acc.UpdateProfile(&models.AccountOption{
		Name:        opt.Name,
		Avatar:      opt.Avatar,
		Description: opt.Description,
	})
}

func (a Account) Apps() ([]*App, error) {
	mapps, err := models.AppsByAccount(a.ID)
	if err != nil {
		return nil, err
	}

	var apps []*App
	for _, v := range mapps {
		apps = append(apps, covertToApp(v))
	}

	return apps, nil
}
