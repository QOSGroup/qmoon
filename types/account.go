// Copyright 2018 The QOS Authors

package types

const AuthUserKey = "user"

type AdminAccountStatus int

const (
	AdminAccountStatusInit AdminAccountStatus = iota
	AdminAccountStatusChecked
)

func (aas AdminAccountStatus) String() string {
	switch aas {
	case AdminAccountStatusInit:
		return "未验证邮箱"
	default:
		return ""
	}
}
