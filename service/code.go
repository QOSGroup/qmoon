// Copyright 2018 The QOS Authors

package service

import (
	"fmt"
	"strings"
	"time"

	"github.com/QOSGroup/qmoon/lib/mail"
	"github.com/QOSGroup/qmoon/models"
	"github.com/QOSGroup/qmoon/utils"
)

const VerifyCodeExpired = time.Minute * 15
const VerifyCodeDuration = time.Minute

// CheckCode 验证码校验
func CheckCode(email, code string) bool {
	vc, err := models.VerifyCodeByEmail(email)
	if err != nil {
		return false
	}

	now := time.Now()

	if vc.CreatedAt.Add(VerifyCodeExpired).Before(now) {
		return false
	}

	vc.Delete()

	return vc.Code == code
}

func mailMsg(from string, to []string, code string) []byte {
	nickname := "admin"
	user := from
	subject := "QOS验证码"
	contentType := "Content-Type: text/plain; charset=UTF-8"
	body := fmt.Sprintf(`
	亲爱的用户：
		您好!

		您本次操作验正码为：%s, %d分钟内有效, 若非本人操作，请忽略!
	`, code, int(VerifyCodeExpired.Minutes()))
	msg := []byte("To: " + strings.Join(to, ",") + "\r\nFrom: " + nickname +
		"<" + user + ">\r\nSubject: " + subject + "\r\n" + contentType + "\r\n\r\n" + body)

	return msg
}

// SendCode 发送短信验证码
func SendCode(mailSmtpServer, mailUser, mailPassword, email string) error {
	code := utils.RandmonCode(6)
	_, err := models.CreateVerifyCode(email, code)
	if err != nil {
		return err
	}

	msg := mailMsg(mailUser, []string{email}, code)
	err = mail.New(mailSmtpServer, mailUser, mailPassword).Send([]string{email}, msg)
	if err != nil {
		return err
	}

	return nil
}
