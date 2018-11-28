// Copyright 2018 The QOS Authors

package service

import (
	"os"
	"testing"
)

func TestSendCode(t *testing.T) {
	mailSmtpServer := os.Getenv("MailSmtpServer")
	mailUser := os.Getenv("MailUser")
	mailPassword := os.Getenv("MailPassword")
	to := "testxxxxx@gmail.com"
	SendCode(mailSmtpServer, mailUser, mailPassword, to)
}
