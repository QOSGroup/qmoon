// Copyright 2018 The QOS Authors

package service

import (
	"errors"
	"time"

	"github.com/QOSGroup/qmoon/service/account"
	"github.com/QOSGroup/qmoon/types"
)

type session struct {
	Token    string    `json:"token"`
	ExpireAt time.Time `json:"expireAt"`
	Status   int64     `json:"status"` // 用户状态
	Name     string    `json:"name"`   // name
	Avatar   string    `json:"avatar"` // avatar
}

func Login(mail, password string) (*session, error) {
	acc, err := account.RetrieveAccountByMail(mail)
	if err != nil {
		return nil, err
	}

	if !acc.CheckPassword(password) {
		return nil, errors.New("密码错误")
	}

	// 一个月过期时间
	expire := time.Hour * 24 * 30
	t, err := acc.CreateSession(int64(types.LoginWeb), expire)
	if err != nil {
		return nil, err
	}
	s := &session{
		Token:    t,
		ExpireAt: time.Now().Add(expire),
		Status:   acc.Status,
		Name:     acc.Name,
		Avatar:   acc.Avatar,
	}

	return s, nil
}

func Logout(uid int64) error {
	acc, err := account.RetrieveAccountByID(uid)
	if err != nil {
		return err
	}

	return acc.DeleteSession()
}

func CheckSession(token string) (*account.Account, error) {
	t, err := account.RetrieveToken(token)
	if err != nil {
		return nil, err
	}
	acc, err := t.Account()
	if err != nil {
		return nil, err
	}

	return acc, nil
}
