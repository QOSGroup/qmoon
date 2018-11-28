// Copyright 2018 The QOS Authors

// Package pkg comments for pkg account
// account 管理后台相关功能
package account

import (
	"database/sql"
	"errors"
	"fmt"
	"time"

	"github.com/QOSGroup/qmoon/db"
	"github.com/QOSGroup/qmoon/db/model"
	"github.com/QOSGroup/qmoon/types"
	"github.com/QOSGroup/qmoon/utils"
)

type Account struct {
	ma    *model.Account
	mapps []*model.App

	ID          int64     `json:"id"`
	Mail        string    `json:"mail"`        // mail
	Name        string    `json:"name"`        // name
	Avatar      string    `json:"avatar"`      // avatar
	Description string    `json:"description"` // description
	Status      int64     `json:"status"`      // status
	CreatedAt   time.Time `json:"created_at"`  // created_at
}

type App struct {
	mapp      *model.App
	ID        int64  `json:"id"`
	Name      string `json:"name"`
	SecretKey string `json:"secretKey"`
	Status    int64  `json:"desc"`
	AccountID int64  `json:"accountId"`
}

func covertToAccount(ma *model.Account) *Account {
	return &Account{
		ma:          ma,
		ID:          ma.ID,
		Mail:        ma.Mail.String,
		Name:        ma.Name.String,
		Avatar:      ma.Avatar.String,
		Description: ma.Description.String,
		Status:      ma.Status.Int64,
		CreatedAt:   ma.CreatedAt.Time,
	}
}

func covertToApp(app *model.App) *App {
	return &App{
		mapp:      app,
		ID:        app.ID,
		Name:      app.Name.String,
		SecretKey: app.SecretKey.String,
		Status:    app.Status.Int64,
		AccountID: app.AccountID.Int64,
	}
}

func List(offset, limit int64) ([]*Account, error) {
	mas, err := model.AccountFilter(db.Db, "", offset, limit)
	if err != nil {
		return nil, err
	}

	var as []*Account
	for _, v := range mas {
		as = append(as, covertToAccount(v))
	}

	return as, nil
}

func CreateAccount(mail, password string) (*Account, error) {
	ma := &model.Account{
		Mail:      utils.NullString(mail),
		Password:  utils.NullString(utils.EncryptPwd([]byte(password))),
		Status:    utils.NullInt64(int64(types.AdminAccountStatusChecked)),
		CreatedAt: utils.NullTime(time.Now()),
	}

	err := ma.Insert(db.Db)
	if err != nil {
		return nil, err
	}

	return covertToAccount(ma), nil
}

func RetrieveAccountByID(id int64) (*Account, error) {
	ma, err := model.AccountByID(db.Db, id)
	if err != nil {
		return nil, err
	}

	return covertToAccount(ma), nil
}
func RetrieveAccountByMail(mail string) (*Account, error) {
	ma, err := model.AccountByMail(db.Db, utils.NullString(mail))
	if err != nil {
		return nil, err
	}

	return covertToAccount(ma), nil
}

func (a Account) CheckPassword(pwd string) bool {

	return utils.EncryptPwd([]byte(pwd)) == a.ma.Password.String
}

func (a *Account) UpdatePassword(pwd string) error {
	newPwd := utils.EncryptPwd([]byte(pwd))
	if newPwd == a.ma.Password.String {
		return errors.New("新旧密码不能一样")
	}

	a.ma.Password = utils.NullString(newPwd)

	return a.ma.Update(db.Db)
}

type option struct {
	name        sql.NullString
	avatar      sql.NullString
	description sql.NullString
}
type SetOption func(option *option) error

func NewOption(fs ...SetOption) (*option, error) {
	opt := &option{}
	if fs != nil {
		for _, f := range fs {
			if err := f(opt); err != nil {
				return nil, err
			}
		}
	}

	return opt, nil
}
func SetName(name string) SetOption {
	return func(option *option) error {
		option.name = utils.NullString(name)
		return nil
	}
}

func SetAvatar(avatar string) SetOption {
	return func(option *option) error {
		option.avatar = utils.NullString(avatar)
		return nil
	}
}

func SetDescription(description string) SetOption {
	return func(option *option) error {
		option.description = utils.NullString(description)
		return nil
	}
}

func (a *Account) UpdateProfile(opt *option) error {
	if opt.name.Valid {
		a.ma.Name = opt.name
	}
	if opt.description.Valid {
		a.ma.Description = opt.description
	}
	if opt.avatar.Valid {
		a.ma.Avatar = opt.avatar
	}

	err := a.ma.Update(db.Db)
	if err != nil {
		return err
	}

	a.Name = a.ma.Name.String
	a.Description = a.ma.Description.String
	a.Avatar = a.ma.Avatar.String

	return nil
}

func (a Account) Apps() ([]*App, error) {
	if a.mapps == nil {
		apps, err := model.AppsByAccountID(db.Db, utils.NullInt64(a.ma.ID))
		if err != nil {
			return nil, err
		}

		a.mapps = apps
	}

	var apps []*App
	for _, v := range a.mapps {
		apps = append(apps, covertToApp(v))
	}

	return apps, nil
}

func (a *Account) RetrieveAppByID(id int64) (*App, error) {
	if a.mapps == nil {
		apps, err := model.AppsByAccountID(db.Db, utils.NullInt64(a.ma.ID))
		if err != nil {
			return nil, err
		}

		a.mapps = apps
	}
	for _, v := range a.mapps {
		if v.ID == id {
			return covertToApp(v), nil
		}
	}

	return nil, errors.New("not found")
}

func (a *Account) DeleteAppByID(id int64) error {
	if a.mapps == nil {
		apps, err := model.AppsByAccountID(db.Db, utils.NullInt64(a.ma.ID))
		if err != nil {
			return err
		}

		a.mapps = apps
	}
	for k, v := range a.mapps {
		if v.ID == id {
			err := v.Delete(db.Db)
			if err != nil {
				return err
			}
			a.mapps = append(a.mapps[:k], a.mapps[k+1:]...)
			return nil
		}
	}

	return errors.New("not found")
}

func (a *Account) CreateApp(name string) (*App, error) {
	token := utils.MD5([]byte(time.Now().String()))
	mapp := &model.App{
		Name:      utils.NullString(name),
		SecretKey: utils.NullString(token),
		AccountID: utils.NullInt64(a.ma.ID),
		CreatedAt: utils.NullTime(time.Now()),
	}

	err := mapp.Insert(db.Db)
	if err != nil {
		return nil, err
	}

	return covertToApp(mapp), nil
}

func (a *Account) CreateSession(loginType int64, expire time.Duration) (string, error) {
	now := time.Now()
	token := string(utils.MD5([]byte(fmt.Sprintf("%s:%d", now.Format("2006.01.02"), a.ID))))
	loginSta, err := model.LoginStatusByAccountID(db.Db, utils.NullInt64(a.ID))
	if err != nil {
		loginSta = &model.LoginStatus{}
		loginSta.AccountID = utils.NullInt64(a.ID)
		loginSta.LoginType = utils.NullInt64(loginType)
		loginSta.ExpiredAt = utils.NullTime(now.Add(expire))
		loginSta.Token = utils.NullString(token)
		if err := loginSta.Insert(db.Db); err != nil {
			return "", err
		}
	} else {
		loginSta.ExpiredAt = utils.NullTime(now.Add(expire))
		loginSta.Token = utils.NullString(token)
		if err := loginSta.Update(db.Db); err != nil {
			return "", err
		}
	}

	return token, nil
}

func (a *Account) DeleteSession() error {
	loginSta, err := model.LoginStatusByAccountID(db.Db, utils.NullInt64(a.ID))
	if err != nil {
		return nil
	} else {
		return loginSta.Delete(db.Db)
	}
}

func AppBySecretKey(secretKey string) (*App, error) {
	app, err := model.AppBySecretKey(db.Db, utils.NullString(secretKey))
	if err != nil {
		return nil, err
	}

	return covertToApp(app), nil
}

func (app App) Account() (*Account, error) {
	a, err := app.mapp.Account(db.Db)
	if err != nil {
		return nil, err
	}

	return covertToAccount(a), nil
}

type Token struct {
	ls        *model.LoginStatus
	AccountID int64     `json:"accountId"`
	LoginType int64     `json:"loginType"`
	Token     string    `json:"token"`
	ExpiredAt time.Time `json:"expiredAt"`
}

func covertToToken(ls *model.LoginStatus) *Token {
	return &Token{
		ls:        ls,
		AccountID: ls.AccountID.Int64,
		LoginType: ls.LoginType.Int64,
		Token:     ls.Token.String,
		ExpiredAt: ls.ExpiredAt.Time,
	}
}

func RetrieveToken(token string) (*Token, error) {
	ls, err := model.LoginStatusByToken(db.Db, utils.NullString(token))
	if err != nil {
		return nil, err
	}

	if ls.ExpiredAt.Time.Before(time.Now()) {
		ls.Delete(db.Db)
		return nil, errors.New("登陆过期,请重新登陆")
	}

	return covertToToken(ls), nil
}

func (t Token) Account() (*Account, error) {
	a, err := t.ls.Account(db.Db)
	if err != nil {
		return nil, err
	}

	return covertToAccount(a), nil
}
