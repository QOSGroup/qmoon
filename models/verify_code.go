package models

import (
	"time"

	"github.com/QOSGroup/qmoon/models/errors"
	"github.com/go-xorm/xorm"
)

type VerifyCode struct {
	Id    int64  `xorm:"pk autoincr BIGINT"`
	Email string `xorm:"unique TEXT"`
	Code  string `xorm:"TEXT"`

	Version       int       `xorm:"version"`
	CreatedAt     time.Time `xorm:"-"`
	CreatedAtUnix int64
}

func (vc *VerifyCode) BeforeInsert() {
	vc.CreatedAtUnix = time.Now().Unix()
}

func (vc VerifyCode) BeforeUpdate() {
	vc.CreatedAtUnix = time.Now().Unix()
}

func (vc *VerifyCode) AfterSet(colName string, _ xorm.Cell) {
	switch colName {
	case "created_at_unix":
		vc.CreatedAt = time.Unix(vc.CreatedAtUnix, 0).Local()
	}
}

func (vc *VerifyCode) Delete() error {
	_, err := basex.Delete(vc)
	return err
}

const VerifyCodeExpired = time.Minute * 15
const VerifyCodeDuration = time.Minute

func (vc *VerifyCode) Insert() error {
	_, err := basex.Insert(vc)
	if err != nil {
		return err
	}

	return nil
}

func CreateVerifyCode(email, code string) (*VerifyCode, error) {
	vc, err := VerifyCodeByEmail(email)
	if err != nil {
		if !errors.IsNotExist(err) {
			return nil, err
		}
		vc = &VerifyCode{
			Email: email,
			Code:  code}
		_, err = basex.Insert(vc)
	} else {
		if vc.CreatedAt.Add(VerifyCodeDuration).After(time.Now()) {
			return nil, errors.New("验证请求太频繁，请稍后再试")
		}

		vc.Code = code
		_, err = basex.ID(vc.Id).Cols("code", "created_at_unix").Update(vc)
	}

	return vc, err
}

func VerifyCodeByEmail(email string) (*VerifyCode, error) {
	vc := &VerifyCode{Email: email}
	has, err := basex.Get(vc)
	if err != nil {
		return nil, err
	}

	if !has {
		return nil, errors.NotExist{Obj: "VerifyCode"}
	}

	return vc, nil
}
