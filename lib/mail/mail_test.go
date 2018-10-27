// Copyright 2018 The QOS Authors

package mail

import (
	"os"
	"testing"
)

/*
export MailSmtpServer=
export MailUser=
export MailPassword=
*/

func TestMail(t *testing.T) {
	mailSmtpServer := os.Getenv("MailSmtpServer")
	mailUser := os.Getenv("MailUser")
	mailPassword := os.Getenv("MailPassword")
	c := New(mailSmtpServer, mailUser, mailPassword)
	c.Send([]string{"test@163.com"}, []byte("aaaaaaa"))
}
