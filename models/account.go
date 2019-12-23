package models

import (
	"fmt"
	"time"

	"github.com/QOSGroup/qmoon/models/errors"
	"github.com/QOSGroup/qmoon/types"
	"github.com/QOSGroup/qmoon/utils"
	"github.com/go-xorm/xorm"
)

type Account struct {
	Id          int64  `xorm:"pk autoincr BIGINT"`
	Mail        string `xorm:"unique"`
	Name        string `xorm:"VARCHAR(64)"`
	Avatar      string `xorm:"TEXT"`
	Description string `xorm:"TEXT"`
	Status      int    `xorm:"INTEGER"`
	Password    string `xorm:"VARCHAR(128)"`

	Version int `xorm:"version"`

	UpdatedAt     time.Time `xorm:"-"`
	UpdatedAtUnix int64
	CreatedAt     time.Time `xorm:"-"`
	CreatedAtUnix int64
}

func (a *Account) BeforeInsert() {
	a.CreatedAtUnix = time.Now().Unix()
	a.UpdatedAtUnix = time.Now().Unix()
}

func (a *Account) BeforeUpdate() {
	a.UpdatedAtUnix = time.Now().Unix()
}

func (a *Account) AfterSet(colName string, _ xorm.Cell) {
	switch colName {
	case "created_at_unix":
		a.CreatedAt = time.Unix(a.CreatedAtUnix, 0).Local()
	case "updated_at_unix":
		a.UpdatedAt = time.Unix(a.UpdatedAtUnix, 0).Local()
	}
}

type AccountOption struct {
	Name          string
	Avatar        string
	Description   string
	Offset, Limit int
	cols          []string
}

func (a *Account) Insert() error {
	_, err := basex.Insert(a)
	if err != nil {
		return err
	}

	return nil
}

func (a *Account) UpdateProfile(opt *AccountOption) error {
	_, err := basex.ID(a.Id).Update(&Account{
		Name:        opt.Name,
		Avatar:      opt.Avatar,
		Description: opt.Description,
	})
	return err
}

func (a *Account) CheckPassword(pwd string) bool {
	return utils.EncryptPwd([]byte(pwd)) == a.Password
}

func (a *Account) CreateApp(name string) (*App, error) {
	app := &App{
		Name:      name,
		SecretKey: utils.MD5([]byte(time.Now().String())),
		AccountId: a.Id,
	}
	_, err := basex.Insert(app)

	return app, err
}

func (a *Account) Apps() ([]*App, error) {
	return AppsByAccount(a.Id)
}

func (a *Account) DeleteAppByID(id int64) error {
	app, err := AppById(id)
	if err != nil {
		return err
	}

	if app.AccountId != a.Id {
		return nil
	}

	_, err = basex.Delete(app)
	return err
}

func (a *Account) UpdatePassword(pwd string) error {
	newPwd := utils.EncryptPwd([]byte(pwd))
	if newPwd == a.Password {
		return errors.New("新旧密码不能一样")
	}
	a.Password = newPwd
	_, err := basex.ID(a.Id).Cols("password").Update(a)

	return err
}

func (a *Account) CreateSession(loginType int, expire time.Duration) (string, error) {
	now := time.Now()
	token := string(utils.MD5([]byte(fmt.Sprintf("%s:%d", now.Format("2006.01.02"), a.Id))))
	loginSta, err := LoginStatusByAccountID(a.Id)
	if err != nil {
		if !errors.IsNotExist(err) {
			return "", err
		}
		loginSta = &LoginStatus{}
		loginSta.AccountId = a.Id
		loginSta.LoginType = loginType
		loginSta.ExpiredAtUnix = now.Add(expire).Unix()
		loginSta.Token = token
		if _, err := basex.Insert(loginSta); err != nil {
			return "", err
		}
	} else {
		loginSta.ExpiredAtUnix = now.Add(expire).Unix()
		loginSta.Token = token
		if _, err := basex.ID(loginSta.Id).Update(loginSta); err != nil {
			return "", err
		}
	}

	return token, nil
}

func (a *Account) DeleteSession() error {
	loginSta, err := LoginStatusByAccountID(a.Id)
	if err != nil {
		return nil
	}

	_, err = basex.Delete(loginSta)
	return err
}

func Accounts(opt *AccountOption) ([]*Account, error) {
	var accs = make([]*Account, 0)
	sess := basex.NewSession()
	defer sess.Close()
	if opt.Limit > 0 {
		sess = sess.Limit(opt.Limit, opt.Offset)
	}
	return accs, sess.Find(&accs)
}

func CreateAccount(mail, password string) (*Account, error) {
	acc := &Account{}
	acc.Mail = mail
	acc.Password = utils.EncryptPwd([]byte(password))
	acc.Status = int(types.AdminAccountStatusChecked)
	if _, err := basex.Insert(acc); err != nil {
		return nil, err
	}

	return acc, nil
}

func RetrieveAccountByID(id int64) (*Account, error) {
	n := &Account{Id: id}
	has, err := basex.Get(n)
	if err != nil {
		return nil, err
	}

	if !has {
		return nil, errors.AccountNotExist{Id: id}
	}

	return n, nil
}

func RetrieveAccountByMail(mail string) (*Account, error) {
	n := &Account{Mail: mail}
	has, err := basex.Get(n)
	if err != nil {
		return nil, err
	}

	if !has {
		return nil, errors.AccountNotExist{Mail: mail}
	}

	return n, nil
}

func AccountIsExist(mail string) (bool, error) {
	n := &Account{Mail: mail}
	exist, err := basex.Exist(n)
	if err != nil {
		return false, err
	}

	return exist, nil
}
