// Copyright 2018 The QOS Authors

// Package pkg comments for pkg account
// account 账户
package account

type Account struct {
	SecretID  string `json:"secretId"`
	SecretKey string `json:"secretKey"`
	Desc      string `json:"desc"`
}

func Get(secretID string) (*Account, error) {
	a := new(Account)

	a.SecretID = "123"
	a.SecretKey = "afdcd"

	return a, nil
}

func List(offset, limit int) ([]*Account, error) {
	return nil, nil
}

func (a *Account) Update() error {
	return nil
}

func (a *Account) Delete() error {
	return nil
}
