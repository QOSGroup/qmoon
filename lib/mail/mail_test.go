// Copyright 2018 The QOS Authors

package mail

import (
	"os"
	"strings"
	"testing"
)

/*
export MailUser=
export MailPassword=
export MailSmtpServer=
*/

func TestMail(t *testing.T) {
	mailSmtpServer := os.Getenv("MailSmtpServer")
	mailUser := os.Getenv("MailUser")
	mailPassword := os.Getenv("MailPassword")
	c := New(mailSmtpServer, mailUser, mailPassword)
	to := []string{"testxxxxx@gmail.com"}
	nickname := "admin"
	user := mailUser
	subject := "QOS验证码"
	contentType := "Content-Type: text/plain; charset=UTF-8"
	body := `
	亲爱的用户：
		您好!

		您本次操作验正码为： 981632, 10分钟内有效, 若非本人操作，请勿理会!
	`
	msg := []byte("To: " + strings.Join(to, ",") + "\r\nFrom: " + nickname +
		"<" + user + ">\r\nSubject: " + subject + "\r\n" + contentType + "\r\n\r\n" + body)

	c.Send(to, msg)
}
