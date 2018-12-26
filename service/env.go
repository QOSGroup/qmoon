// Copyright 2018 The QOS Authors

package service

import (
	"errors"
	"os"

	"github.com/QOSGroup/qmoon/plugins"
)

type Env struct {
	EmailSmtpServer string
	EmailUser       string
	EmailPassword   string
}

var env *Env

func GetEnv() Env {
	if env == nil {
		env = new(Env)
		env.EmailSmtpServer = os.Getenv("MailSmtpServer")
		env.EmailUser = os.Getenv("MailUser")
		env.EmailPassword = os.Getenv("MailPassword")
	}

	return *env
}

func Doctor() error {
	if err := CheckEnv(); err != nil {
		return err
	}

	if err := plugins.Doctor(); err != nil {
		return err
	}

	return nil
}

func CheckEnv() error {
	if env.EmailPassword == "" {
		return errors.New("MailSmtpServer 未设置")
	}
	if env.EmailUser == "" {
		return errors.New("EmailUser 未设置")
	}
	if env.EmailPassword == "" {
		return errors.New("EmailPassword 未设置")
	}

	return nil
}
